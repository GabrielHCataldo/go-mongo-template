package mongo

import (
	"context"
	"errors"
	"github.com/GabrielHCataldo/go-logger/logger"
	"github.com/GabrielHCataldo/go-mongo-template/internal/util"
	"github.com/GabrielHCataldo/go-mongo-template/mongo/option"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strconv"
	"strings"
)

// Pipeline is a type that makes creating aggregation pipelines easier. It is a
// helper and is intended for serializing to BSON.
//
// Example usage:
//
//	Pipeline{
//		{{"$group", bson.D{{"_id", "$state"}, {"totalPop", bson.D{{"$sum", "$pop"}}}}}},
//		{{"$match", bson.D{{"totalPop", bson.D{{"$gte", 10*1000*1000}}}}}},
//	}
type Pipeline []bson.D

// DeleteResult is the result type returned by DeleteOne and DeleteMany operations.
type DeleteResult struct {
	// The number of documents deleted.
	DeletedCount int64
}

// UpdateResult is the result type returned from UpdateOne, UpdateMany, and ReplaceOne operations.
type UpdateResult struct {
	// The number of documents matched by the filter.
	MatchedCount int64
	// The number of documents modified by the operation.
	ModifiedCount int64
	// The number of documents upserted by the operation.
	UpsertedCount int64
	// The _id field of the upserted document, or nil if no upsert was done.
	UpsertedID any
}

// IndexSpecification represents an index in a database. This type is returned by the IndexView.ListSpecifications
// function and is also used in the CollectionSpecification type.
type IndexSpecification struct {
	// The index name.
	Name string
	// The namespace for the index. This is a string in the format "databaseName.collectionName".
	Namespace string
	// The keys specification document for the index.
	KeysDocument bson.Raw
	// The index version.
	Version int32
	// The length of time, in seconds, for documents to remain in the collection. The default value is 0, which means
	// that documents will remain in the collection until they're explicitly deleted or the collection is dropped.
	ExpireAfterSeconds *int32
	// If true, the index will only reference documents that contain the fields specified in the index. The default is
	// false.
	Sparse *bool
	// If true, the collection will not accept insertion or update of documents where the index key value matches an
	// existing value in the index. The default is false.
	Unique *bool
	// The clustered index.
	Clustered *bool
}

type Template struct {
	client  *mongo.Client
	session mongo.Session
}

func NewTemplate(ctx context.Context, opts ...*options.ClientOptions) (*Template, error) {
	conn, err := mongo.Connect(ctx, opts...)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &Template{
		client: conn,
	}, nil
}

// InsertOne executes an insert command to insert a single document into the collection.
//
// The document parameter must be the document to be inserted. It cannot be nil. If the document does not have an _id
// field when transformed into BSON, one will be added automatically to the marshalled document. The original document
// will not be modified. The _id can be retrieved from the InsertedID field of the returned InsertOneResult.
//
// The opts parameter can be used to specify options for the operation (see the option.InsertOne documentation.)
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/insert/.
func (t *Template) InsertOne(ctx context.Context, document any, opts ...option.InsertOne) error {
	opt := option.GetInsertOneOptionByParams(opts)
	err := t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			err = t.insertOne(sc, document, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
		})
	}
	return err
}

// InsertMany executes an insert command to insert multiple documents into the collection. If recording errors occur
// during operation you can configure automatic rollback, see the option.InsertMany documentation.
//
// The documents parameter must be a pointer slice of the documents to be inserted. The slice cannot be null or empty.
// All elements must be non-zero. For any document that does not have the _id field when transformed into BSON,
// the automatically generated value for the document will be added.
//
// The opts parameter can be used to specify options for the operation (see the option.InsertMany documentation.)
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/insert/.
func (t *Template) InsertMany(ctx context.Context, documents any, opts ...option.InsertMany) error {
	opt := option.GetInsertManyOptionByParams(opts)
	err := t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			err = t.insertMany(sc, documents, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, opt.DisableAutoRollbackSession, err)
		})
	}
	return err

}

// DeleteOne executes a delete command to delete at most one document from the collection.
//
// The filter parameter must be a document containing query operators and can be used to select the document to be
// deleted. It cannot be null. If the filter does not match any documents, the operation succeeds and a DeleteResult
// with a DeletedCount of 0 will be returned. If the filter matches multiple documents, one will be selected from the list
// matching set.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.Delete documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/delete/.
func (t *Template) DeleteOne(ctx context.Context, filter, ref any, opts ...option.Delete) (*DeleteResult, error) {
	var result *DeleteResult
	var err error
	opt := option.GetDeleteOptionByParams(opts)
	err = t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			result, err = t.deleteOne(sc, filter, ref, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
		})
	}
	return result, err
}

// DeleteOneById executes an update command to update the document whose _id value matches the provided ID in the collection.
// This is equivalent to running DeleteOne(ctx, bson.D{{"_id", id}}, ref, opts...).
//
// The id parameter is the _id of the document to be updated. It cannot be nil. If the ID does not match any documents,
// the operation will succeed and an UpdateResult with a MatchedCount of 0 will be returned.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.Delete documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/delete/.
func (t *Template) DeleteOneById(ctx context.Context, id, ref any, opts ...option.Delete) (*DeleteResult, error) {
	var result *DeleteResult
	var err error
	opt := option.GetDeleteOptionByParams(opts)
	err = t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			result, err = t.deleteOne(sc, bson.D{{"_id", id}}, ref, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
		})
	}
	return result, err
}

// DeleteMany executes a delete command to delete documents from the collection.
//
// The filter parameter must be a document containing query operators and can be used to select the documents to
// be deleted. It cannot be nil. An empty document (e.g. bson.D{}) should be used to delete all documents in the
// collection. If the filter does not match any documents, the operation will succeed and a DeleteResult with a
// DeletedCount of 0 will be returned.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.Delete documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/delete/.
func (t *Template) DeleteMany(ctx context.Context, filter, ref any, opts ...option.Delete) (*DeleteResult, error) {
	var result *DeleteResult
	var err error
	opt := option.GetDeleteOptionByParams(opts)
	err = t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			result, err = t.deleteMany(sc, filter, ref, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, opt.DisableAutoRollbackSession, err)
		})
	}
	return result, err
}

// UpdateOneById executes an update command to update the document whose _id value matches the provided ID in the collection.
// This is equivalent to running UpdateOne(ctx, bson.D{{"_id", id}}, update, ref, opts...).
//
// The id parameter is the _id of the document to be updated. It cannot be nil. If the ID does not match any documents,
// the operation will succeed and an UpdateResult with a MatchedCount of 0 will be returned.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The update parameter must be a document containing update operators
// (https://www.mongodb.com/docs/manual/reference/operator/update/) and can be used to specify the modifications to be
// made to the selected document. It cannot be nil or empty.
//
// The opts parameter can be used to specify options for the operation (see the option.Update documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
func (t *Template) UpdateOneById(ctx context.Context, id, update, ref any, opts ...option.Update) (*UpdateResult, error) {
	var result *UpdateResult
	var err error
	opt := option.GetUpdateOptionByParams(opts)
	err = t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			result, err = t.updateOne(sc, bson.D{{"_id", id}}, update, ref, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
		})
	}
	return result, err
}

// UpdateOne executes an update command to update at most one document in the collection.
//
// The filter parameter must be a document containing query operators and can be used to select the document to be
// updated. It cannot be nil. If the filter does not match any documents, the operation will succeed and an UpdateResult
// with a MatchedCount of 0 will be returned. If the filter matches multiple documents, one will be selected from the
// matched set and MatchedCount will equal 1.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The update parameter must be a document containing update operators
// (https://www.mongodb.com/docs/manual/reference/operator/update/) and can be used to specify the modifications to be
// made to the selected document. It cannot be nil or empty.
//
// The opts parameter can be used to specify options for the operation (see the option.Update documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
func (t *Template) UpdateOne(ctx context.Context, filter any, update, ref any, opts ...option.Update) (*UpdateResult,
	error) {
	var result *UpdateResult
	var err error
	opt := option.GetUpdateOptionByParams(opts)
	err = t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			result, err = t.updateOne(sc, filter, update, ref, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
		})
	}
	return result, err
}

// UpdateMany executes an update command to update documents in the collection.
//
// The filter parameter must be a document containing query operators and can be used to select the documents to be
// updated. It cannot be nil. If the filter does not match any documents, the operation will succeed and an UpdateResult
// with a MatchedCount of 0 will be returned.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The update parameter must be a document containing update operators
// (https://www.mongodb.com/docs/manual/reference/operator/update/) and can be used to specify the modifications to be made
// to the selected documents. It cannot be nil or empty.
//
// The opts parameter can be used to specify options for the operation (see the option.Update documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
func (t *Template) UpdateMany(ctx context.Context, filter any, update, ref any, opts ...option.Update) (*UpdateResult,
	error) {
	var result *UpdateResult
	var err error
	opt := option.GetUpdateOptionByParams(opts)
	err = t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			result, err = t.updateMany(sc, filter, update, ref, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
		})
	}
	return result, err
}

// ReplaceOne executes an update command to replace at most one document in the collection.
//
// The filter parameter must be a document containing query operators and can be used to select the document to be
// replaced. It cannot be nil. If the filter does not match any documents, the operation will succeed and an
// UpdateResult with a MatchedCount of 0 will be returned. If the filter matches multiple documents, one will be
// selected from the matched set and MatchedCount will equal 1.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The replacement parameter must be a document that will be used to replace the selected document. It cannot be nil
// and cannot contain any update operators (https://www.mongodb.com/docs/manual/reference/operator/update/).
//
// The opts parameter can be used to specify options for the operation (see the option.Replace documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
func (t *Template) ReplaceOne(ctx context.Context, filter any, update, ref any, opts ...option.Replace) (*UpdateResult,
	error) {
	var result *UpdateResult
	var err error
	opt := option.GetReplaceOptionByParams(opts)
	err = t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			result, err = t.replaceOne(sc, filter, update, ref, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
		})
	}
	return result, err
}

// ReplaceOneById executes an update command to update the document whose _id value matches the provided ID in the collection.
// This is equivalent to running ReplaceOne(ctx, bson.D{{"_id", id}}, replacement, ref, opts...).
//
// The id parameter is the _id of the document to be updated. It cannot be nil. If the ID does not match any documents,
// the operation will succeed and an UpdateResult with a MatchedCount of 0 will be returned.
//
// The opts parameter can be used to specify options for the operation (see the option.Replace documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
func (t *Template) ReplaceOneById(ctx context.Context, id, replacement, ref any, opts ...option.Replace) (*UpdateResult,
	error) {
	var result *UpdateResult
	var err error
	opt := option.GetReplaceOptionByParams(opts)
	err = t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			result, err = t.replaceOne(sc, bson.D{{"_id", id}}, replacement, ref, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
		})
	}
	return result, err
}

// FindOneById executes a search command whose _id value matches the ID given in the collection.
// This is equivalent to running FindOne(ctx, bson.D{{"_id", id}}, dest, opts...).
//
// The id parameter must be a document containing query operators and can be used to select the document to be
// returned. It cannot be null. If the id does not match any document, the dest parameter will not be modified.
//
// The dest parameter must be a pointer to the return expected by the operation, it is important to have the
// database and collection tags configured.
//
// The opts parameter can be used to specify options for this operation (see the option.FindOneById documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/find/.
func (t *Template) FindOneById(ctx context.Context, id, dest any, opts ...option.FindOneById) error {
	opt := option.GetFindOneByIdOptionByParams(opts)
	return t.findOne(ctx, bson.D{{"_id", id}}, dest, option.FindOne{
		AllowPartialResults: opt.AllowPartialResults,
		Collation:           opt.Collation,
		Comment:             opt.Comment,
		Hint:                opt.Hint,
		Max:                 opt.Max,
		MaxTime:             opt.MaxTime,
		Min:                 opt.Min,
		Projection:          opt.Projection,
		ReturnKey:           opt.ReturnKey,
		ShowRecordID:        opt.ShowRecordID,
	})
}

// FindOne executes a find command, if successful it returns the corresponding documents in the collection in the dest
// parameter with return error nil. Otherwise, it returns corresponding error.
//
// The id parameter must be a document id that is used to select the document to be
// returned. It cannot be null. If the id does not match any document, the dest parameter will not be modified.
//
// The dest parameter must be a pointer to the return expected by the operation, it is important to have the
// database and collection tags configured.
//
// The opts parameter can be used to specify options for this operation (see the option.FindOne documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/find/.
func (t *Template) FindOne(ctx context.Context, filter, dest any, opts ...option.FindOne) error {
	return t.findOne(ctx, filter, dest, opts...)
}

// FindOneAndDelete executes a findAndModify command to delete at most one document from the collection. and returns the
// document as it appeared before deletion in the dest parameter.
//
// The filter parameter must be a document containing query operators and can be used to select the document to be
// deleted. It cannot be null. If the filter does not match any documents, an error set to ErrNoDocuments will be
// returned. If the filter matches multiple documents, one will be selected from the matching set.
//
// The dest parameter must be a pointer to the return expected by the operation, it is important to have the
// database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.FindOneAndDelete documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/findAndModify/.
func (t *Template) FindOneAndDelete(ctx context.Context, filter, dest any, opts ...option.FindOneAndDelete) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	opt := option.GetFindOneAndDeleteOptionByParams(opts)
	err := t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			err = t.findOneAndDelete(sc, filter, dest, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
		})
	}
	return err
}

// FindOneAndReplace executes a findAndModify command to replace at most one document in the collection
// and returns the document as it appeared before replacement in the dest parameter.
//
// The filter parameter must be a document containing query operators and can be used to select the document to be
// replaced. It cannot be null. If the filter does not match any documents, an error set to ErrNoDocuments will be
// returned. If the filter matches multiple documents, one will be selected from the matching set.
//
// The replacement parameter must be a document that will be used to replace the selected document. It cannot be nil
// and cannot contain any update operators (https://www.mongodb.com/docs/manual/reference/operator/update/).
//
// The dest parameter must be a pointer to the return expected by the operation, it is important to have the
// database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.FindOneAndReplace documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/findAndModify/.
func (t *Template) FindOneAndReplace(ctx context.Context, filter, replacement, dest any, opts ...option.FindOneAndReplace) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	opt := option.GetFindOneAndReplaceOptionByParams(opts)
	err := t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			err = t.findOneAndReplace(sc, filter, replacement, dest, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
		})
	}
	return err
}

// FindOneAndUpdate executes a findAndModify command to update at most one document in the collection and returns the
// document as it appeared before updating.
//
// The filter parameter must be a document containing query operators and can be used to select the document to be
// updated. It cannot be null. If the filter does not match any documents, an error set to ErrNoDocuments will be
// returned. If the filter matches multiple documents, one will be selected from the matching set.
//
// The update parameter must be a document containing update operators
// (https://www.mongodb.com/docs/manual/reference/operator/update/) and can be used to specify the modifications to be made
// to the selected document. It cannot be nil or empty.
//
// The opts parameter can be used to specify options for the operation (see the options.FindOneAndUpdateOptions
// documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/findAndModify/.
func (t *Template) FindOneAndUpdate(ctx context.Context, filter, update, dest any, opts ...option.FindOneAndUpdate) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	opt := option.GetFindOneAndUpdateOptionByParams(opts)
	err := t.startSession(ctx, opt.ForceRecreateSession)
	if err == nil {
		err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
			err = t.findOneAndUpdate(sc, filter, update, dest, opt)
			return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
		})
	}
	return err
}

// Find executes a find command, if successful it returns the corresponding documents in the collection in the dest
// parameter with return error nil. Otherwise, it returns corresponding error.
//
// The filter parameter must be a document containing query operators and can be used to select which documents are
// included in the result. If the filter does not match any document, the dest parameter will not be modified.
//
// The dest parameter must be a pointer to the return expected by the operation, it is important to have the
// database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.Find documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/find/.
func (t *Template) Find(ctx context.Context, filter, dest any, opts ...option.Find) error {
	return t.find(ctx, filter, dest, opts...)
}

// FindAll execute a search command. This is equivalent to running Find(ctx, bson.D{}, dest, opts...).
//
// The dest parameter must be a pointer to the return expected by the operation, it is important to have the
// database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.Find documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/find/.
func (t *Template) FindAll(ctx context.Context, dest any, opts ...option.Find) error {
	return t.find(ctx, bson.D{}, dest, opts...)
}

// FindPageable executes a find command, if successful, returns the paginated documents in the
// corresponding PageResult structure in the collection on the target parameter with null return error.
// Otherwise, it will return the corresponding error.
//
// The filter parameter must be a document containing query operators and can be used to select which documents are
// included in the result. It cannot be nil. If the filter does not match any document, the return structure columns
// will be empty.
//
// The opts parameter can be used to specify options for the operation (see the option.FindPageable documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/find/.
func (t *Template) FindPageable(ctx context.Context, filter any, input PageInput, opts ...option.FindPageable) (
	*PageResult, error) {
	if util.IsPointer(input.Ref) {
		return nil, errors.New("mongo: input.Ref cannot be a pointer")
	} else if util.IsInvalid(input.Ref) {
		return nil, errors.New("mongo: invalid type input.Ref")
	}
	databaseName, collectionName, err := getMongoInfosByAny(input.Ref)
	if err != nil {
		return nil, err
	}
	opt := option.GetFindPageableOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName)
	cursor, err := collection.Find(ctx, filter, &options.FindOptions{
		AllowDiskUse:        opt.AllowDiskUse,
		AllowPartialResults: opt.AllowPartialResults,
		BatchSize:           opt.BatchSize,
		Collation:           option.ParseCollationMongoOptions(opt.Collation),
		Comment:             opt.Comment,
		CursorType:          option.ParseCursorType(opt.CursorType),
		Hint:                opt.Hint,
		Limit:               &input.PageSize,
		Max:                 opt.Max,
		MaxAwaitTime:        opt.MaxAwaitTime,
		MaxTime:             opt.MaxTime,
		Min:                 opt.Min,
		NoCursorTimeout:     opt.NoCursorTimeout,
		Projection:          opt.Projection,
		ReturnKey:           opt.ReturnKey,
		ShowRecordID:        opt.ShowRecordID,
		Skip:                &input.Page,
		Sort:                input.Sort,
		Let:                 opt.Let,
	})
	if err != nil {
		return nil, err
	}
	dest := input.Ref
	err = cursor.All(ctx, &dest)
	if err != nil {
		return nil, err
	}
	countTotal, _ := collection.CountDocuments(ctx, filter)
	return newPageResult(input, dest, countTotal), nil
}

// Exists executes the count command, if the quantity is greater than 0 with a limit of 1, true is returned,
// otherwise false is returned.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.Exists documentation).
func (t *Template) Exists(ctx context.Context, filter, ref any, opts ...option.Exists) (bool, error) {
	opt := option.GetExistsOptionByParams(opts)
	count, err := t.CountDocuments(ctx, filter, ref, option.Count{
		Collation: opt.Collation,
		Comment:   opt.Comment,
		Hint:      opt.Hint,
		Limit:     util.ConvertToPointer[int64](1),
		MaxTime:   opt.MaxTime,
		Skip:      nil,
	})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistsById executes a count command whose _id value matches the ID given in the collection.
// This is equivalent to running Exists(ctx, bson.D{{"_id", id}}, dest, opts...).
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.Exists documentation).
func (t *Template) ExistsById(ctx context.Context, id, ref any, opts ...option.Exists) (bool, error) {
	return t.Exists(ctx, bson.D{{"_id", id}}, ref, opts...)
}

// Aggregate executes a find command, if successful it returns the corresponding documents in the collection in the dest
// parameter with return error nil. Otherwise, it returns corresponding error.
//
// The pipeline parameter must be an array of documents, each representing an aggregation stage. The pipeline cannot
// be nil but can be empty. The stage documents must all be non-nil. For a pipeline of bson.D documents, the
// Pipeline type can be used. See
// https://www.mongodb.com/docs/manual/reference/operator/aggregation-pipeline/#db-collection-aggregate-stages for a list of
// valid stages in aggregations.
//
// The dest parameter must be a pointer to the return expected by the operation, it is important to have the
// database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.Aggregate documentation.)
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/aggregate/.
func (t *Template) Aggregate(ctx context.Context, pipeline any, dest any, opts ...option.Aggregate) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	}
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	opt := option.GetAggregateOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName)
	cursor, err := collection.Aggregate(ctx, pipeline, &options.AggregateOptions{
		AllowDiskUse:             opt.AllowDiskUse,
		BatchSize:                opt.BatchSize,
		BypassDocumentValidation: opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		MaxTime:                  opt.MaxTime,
		MaxAwaitTime:             opt.MaxAwaitTime,
		Comment:                  opt.Comment,
		Hint:                     opt.Hint,
		Let:                      opt.Let,
		Custom:                   opt.Custom,
	})
	if err != nil {
		return err
	}
	err = cursor.All(ctx, dest)
	if err != nil {
		return err
	}
	return nil
}

// CountDocuments returns the number of documents in the collection. For a fast count of the documents in the
// collection, see the EstimatedDocumentCount method.
//
// The filter parameter must be a document and can be used to select which documents contribute to the count. It
// cannot be nil. An empty document (e.g. bson.D{}) should be used to count all documents in the collection. This will
// result in a full collection scan.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.Count documentation).
func (t *Template) CountDocuments(ctx context.Context, filter, ref any, opts ...option.Count) (int64, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return 0, err
	}
	opt := option.GetCountOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName)
	return collection.CountDocuments(ctx, filter, &options.CountOptions{
		Collation: option.ParseCollationMongoOptions(opt.Collation),
		Comment:   opt.Comment,
		Hint:      opt.Hint,
		Limit:     opt.Limit,
		MaxTime:   opt.MaxTime,
		Skip:      opt.Skip,
	})
}

// EstimatedDocumentCount executes a count command and returns an estimate of the number of documents in the collection
// using collection metadata.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the option.EstimatedDocumentCount
// documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/count/.
func (t *Template) EstimatedDocumentCount(ctx context.Context, ref any, opts ...option.EstimatedDocumentCount) (int64,
	error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return 0, err
	}
	opt := option.GetEstimatedDocumentCountOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName)
	return collection.EstimatedDocumentCount(ctx, &options.EstimatedDocumentCountOptions{
		Comment: opt.Comment,
		MaxTime: opt.MaxTime,
	})
}

// Distinct executes a distinct command to find the unique values for a specified field in the collection.
//
// The fieldName parameter specifies the field name for which distinct values should be returned.
//
// The filter parameter must be a document containing query operators and can be used to select which documents are
// considered. It cannot be nil. An empty document (e.g. bson.D{}) should be used to select all documents.
//
// The dest parameter must be a pointer to the return expected by the operation.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The opts parameter can be used to specify options for the operation (see the options.DistinctOptions documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/distinct/.
func (t *Template) Distinct(ctx context.Context, fieldName string, filter, dest, ref any, opts ...option.Distinct) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	}
	opt := option.GetDistinctOptionByParams(opts)
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName)
	result, err := collection.Distinct(ctx, fieldName, filter, &options.DistinctOptions{
		Collation: option.ParseCollationMongoOptions(opt.Collation),
		Comment:   opt.Comment,
		MaxTime:   opt.MaxTime,
	})
	if err != nil {
		return err
	}
	return util.ParseAnyJsonDest(result, dest)
}

// Watch returns a change stream for all changes on the deployment. See
// https://www.mongodb.com/docs/manual/changeStreams/ for more information about change streams.
//
// The client must be configured with read concern majority or no read concern for a change stream to be created
// successfully.
//
// The pipeline parameter must be an array of documents, each representing a pipeline stage. The pipeline cannot be
// nil or empty. The stage documents must all be non-nil. See https://www.mongodb.com/docs/manual/changeStreams/ for a list
// of pipeline stages that can be used with change streams. For a pipeline of bson.D documents, the mongo.Pipeline{}
// type can be used.
//
// The opts parameter can be used to specify options for change stream creation (see the option.Watch documentation).
func (t *Template) Watch(ctx context.Context, pipeline any, opts ...option.Watch) (*mongo.ChangeStream, error) {
	opt := option.GetWatchOptionByParams(opts)
	var watchChangeEvents *mongo.ChangeStream
	var err error
	optionsChangeStream := &options.ChangeStreamOptions{
		BatchSize:                opt.BatchSize,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		FullDocument:             option.ParseFullDocument(opt.FullDocument),
		FullDocumentBeforeChange: option.ParseFullDocument(opt.FullDocumentBeforeChange),
		MaxAwaitTime:             opt.MaxAwaitTime,
		ResumeAfter:              opt.ResumeAfter,
		ShowExpandedEvents:       opt.ShowExpandedEvents,
		StartAtOperationTime:     opt.StartAtOperationTime,
		StartAfter:               opt.StartAfter,
		Custom:                   opt.Custom,
		CustomPipeline:           opt.CustomPipeline,
	}
	if len(opt.DatabaseName) != 0 {
		database := t.client.Database(opt.DatabaseName)
		if len(opt.CollectionName) != 0 {
			watchChangeEvents, err = database.Collection(opt.CollectionName).Watch(ctx, pipeline, optionsChangeStream)
		} else {
			watchChangeEvents, err = database.Watch(ctx, pipeline, optionsChangeStream)
		}
	} else {
		watchChangeEvents, err = t.client.Watch(ctx, pipeline, optionsChangeStream)
	}
	return watchChangeEvents, err
}

// WatchHandler is a function that facilitates the reading of watch events, it triggers the Watch function and
// when an event occurs it reads this event transforming all the data obtained by mongoDB into a Context, right
// after this conversion, we call the handler parameter passing the context with all the information, so you can
// process it in the way you see fit.
//
// The opts parameter can be used to specify options for change stream creation (see the option.WatchHandler
// documentation).
func (t *Template) WatchHandler(ctx context.Context, pipeline any, handler HandlerWatch, opts ...option.WatchHandler) error {
	if handler == nil {
		return ErrWatchHandlerIsNil
	}
	opt := option.GetWatchHandlerOptionByParams(opts)
	watchChangeEvents, err := t.Watch(ctx, pipeline, option.Watch{
		DatabaseName:             opt.DatabaseName,
		CollectionName:           opt.CollectionName,
		BatchSize:                opt.BatchSize,
		Collation:                opt.Collation,
		Comment:                  opt.Comment,
		FullDocument:             opt.FullDocument,
		FullDocumentBeforeChange: opt.FullDocumentBeforeChange,
		MaxAwaitTime:             opt.MaxAwaitTime,
		ResumeAfter:              opt.ResumeAfter,
		ShowExpandedEvents:       opt.ShowExpandedEvents,
		StartAtOperationTime:     opt.StartAtOperationTime,
		StartAfter:               opt.StartAfter,
		Custom:                   opt.Custom,
		CustomPipeline:           opt.CustomPipeline,
	})
	if err != nil {
		return err
	}
	for watchChangeEvents.Next(ctx) {
		var event WatchEvent
		_ = watchChangeEvents.Decode(&event)
		processWatchNext(handler, event, opt)
	}
	_ = watchChangeEvents.Close(ctx)
	return nil
}

// DropCollection drops the collection on the server. This method ignores "namespace not found" errors,
// so it is safe to drop a collection that does not exist on the server.
//
// The ref parameter must be the collection structure with database and collection tags configured.
func (t *Template) DropCollection(ctx context.Context, ref any) error {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return err
	}
	return t.client.Database(databaseName).Collection(collectionName).Drop(ctx)
}

// DropDatabase drops the database on the server. This method ignores "namespace not found" errors,
// so it is safe to drop a database that does not exist on the server.
//
// The ref parameter must be the collection structure with database and collection tags configured.
func (t *Template) DropDatabase(ctx context.Context, ref any) error {
	databaseName, _, err := getMongoInfosByAny(ref)
	if err != nil {
		return err
	}
	return t.client.Database(databaseName).Drop(ctx)
}

// CreateOneIndex executes a createIndexes command to create an index on the collection and returns the name of the new
// index. See the CreateManyIndex documentation for more information and an example. For this function's response,
// the name of the index is returned as a string, and if an error occurs, it is returned in the second return parameter
//
// The opts parameter can be used to specify options for this operation (see the option.Index documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/createIndexes/.
func (t *Template) CreateOneIndex(ctx context.Context, input IndexInput) (string, error) {
	return t.createOneIndex(ctx, input)
}

// CreateManyIndex executes a createIndexes command to create multiple indexes on the collection and returns the names of
// the new indexes.
//
// For each IndexInput in the models parameter, the index name can be specified via the Options field. If a name is not
// given, it will be generated from the Keys document.
//
// The opts parameter can be used to specify options for this operation (see the option.Index documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/createIndexes/.
func (t *Template) CreateManyIndex(ctx context.Context, inputs []IndexInput) ([]string, error) {
	return t.createManyIndex(ctx, inputs)
}

// DropOneIndex executes a dropIndexes operation to drop an index on the collection. If the operation succeeds, this returns
// a BSON document in the form {nIndexesWas: <int32>}. The "nIndexesWas" field in the response contains the number of
// indexes that existed prior to the drop.
//
// The name parameter should be the name of the index to drop. If the name is "*", ErrMultipleIndexDrop will be returned
// without running the command because doing so would drop all indexes.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The opts parameter can be used to specify options for this operation (see the option.DropIndex documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/dropIndexes/.
func (t *Template) DropOneIndex(ctx context.Context, name string, ref any, opts ...option.DropIndex) error {
	opt := option.GetDropIndexOptionByParams(opts)
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName).Indexes()
	_, err = collection.DropOne(ctx, name, &options.DropIndexesOptions{
		MaxTime: opt.MaxTime,
	})
	return err
}

// DropAllIndexes executes a dropIndexes operation to drop all indexes on the collection. If the operation succeeds, this
// returns a BSON document in the form {nIndexesWas: <int32>}. The "nIndexesWas" field in the response contains the
// number of indexes that existed prior to the drop.
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// The opts parameter can be used to specify options for this operation (see the option.DropIndex documentation).
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/dropIndexes/.
func (t *Template) DropAllIndexes(ctx context.Context, ref any, opts ...option.DropIndex) error {
	opt := option.GetDropIndexOptionByParams(opts)
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName).Indexes()
	_, err = collection.DropAll(ctx, &options.DropIndexesOptions{
		MaxTime: opt.MaxTime,
	})
	return err
}

// ListIndexes executes a listIndexes command and returns a cursor over the indexes in the collection.
//
// The opts parameter can be used to specify options for this operation (see the option.ListIndexes documentation).
//
// The ref parameter must be the collection structure with database and collection tags configured.
//
// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/listIndexes/.
func (t *Template) ListIndexes(ctx context.Context, ref any, opts ...option.ListIndexes) ([]IndexResult, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	opt := option.GetListIndexesOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName).Indexes()
	cursor, err := collection.List(ctx, &options.ListIndexesOptions{
		BatchSize: opt.BatchSize,
		MaxTime:   opt.MaxTime,
	})
	if err != nil {
		return nil, err
	}
	var results []IndexResult
	err = cursor.All(ctx, &results)
	return results, err
}

// ListIndexSpecifications executes a List command and returns a slice of returned IndexSpecifications.
//
// The ref parameter must be the collection structure with database and collection tags configured.
func (t *Template) ListIndexSpecifications(ctx context.Context, ref any, opts ...option.ListIndexes) (
	[]IndexSpecification, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	opt := option.GetListIndexesOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName).Indexes()
	mongoResult, err := collection.ListSpecifications(ctx, &options.ListIndexesOptions{
		BatchSize: opt.BatchSize,
		MaxTime:   opt.MaxTime,
	})
	var result []IndexSpecification
	for _, v := range mongoResult {
		if v != nil {
			result = append(result, IndexSpecification{
				Name:               v.Name,
				Namespace:          v.Namespace,
				KeysDocument:       v.KeysDocument,
				Version:            v.Version,
				ExpireAfterSeconds: v.ExpireAfterSeconds,
				Sparse:             v.Sparse,
				Unique:             v.Unique,
				Clustered:          v.Clustered,
			})
		}
	}
	return result, err
}

// StartSession creates a new session and a new transaction and stores it in the template itself for the next operations.
func (t *Template) StartSession(ctx context.Context) error {
	return t.startSession(ctx, true)
}

// CloseSession closes session and transaction, if param abort is false it will commit the changes,
// otherwise it will abort all transactions.
func (t *Template) CloseSession(ctx context.Context, abort bool) error {
	var err error
	if abort {
		err = t.AbortTransaction(ctx)
	} else {
		err = t.CommitTransaction(ctx)
	}
	if err != nil {
		return err
	}
	t.endSession(ctx)
	return nil
}

// CommitTransaction commit all transactions on session
func (t *Template) CommitTransaction(ctx context.Context) error {
	if t.session == nil {
		return ErrNoOpenSession
	}
	return t.session.CommitTransaction(ctx)
}

// AbortTransaction abort all transactions on session
func (t *Template) AbortTransaction(ctx context.Context) error {
	if t.session == nil {
		return ErrNoOpenSession
	}
	return t.session.AbortTransaction(ctx)
}

// Disconnect closes the mongodb connection client with return error
func (t *Template) Disconnect(ctx context.Context) error {
	return t.client.Disconnect(ctx)
}

// SimpleDisconnect closes the mongodb connection client without return error
func (t *Template) SimpleDisconnect(ctx context.Context) {
	err := t.Disconnect(ctx)
	if err != nil {
		logger.Error("error close connection to mongoDB:", err)
		return
	}
	logger.Info("connection to mongoDB closed.")
}

func (t *Template) insertOne(sc mongo.SessionContext, document any, opt option.InsertOne) error {
	if util.IsNotPointer(document) {
		return ErrDocumentIsNotPointer
	} else if util.IsNotStruct(document) {
		return ErrDocumentIsNotStruct
	} else if util.IsZero(document) {
		return ErrDocumentIsEmpty
	}
	databaseName, collectionName, err := getMongoInfosByAny(document)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	result, err := coll.InsertOne(sc, document, &options.InsertOneOptions{
		BypassDocumentValidation: opt.BypassDocumentValidation,
		Comment:                  opt.Comment,
	})
	if err != nil {
		return err
	}
	util.SetInsertedIdOnDocument(result.InsertedID, document)
	return nil
}

func (t *Template) insertMany(sc mongo.SessionContext, a any, opt option.InsertMany) error {
	documents := reflect.ValueOf(a)
	if documents.Kind() != reflect.Slice {
		return errors.New("mongo: document on insert many needs be a slice")
	}
	if documents.Len() == 0 {
		return ErrDocumentsIsEmpty
	}
	var errs []error
	for i := 0; i < documents.Len(); i++ {
		indexValue := documents.Index(i)
		document := indexValue.Interface()
		indexStr := strconv.Itoa(i)
		if util.IsNotPointer(document) {
			errs = append(errs, errors.New(ErrDocumentIsNotPointer.Error()+" (index: "+indexStr+")"))
		} else if util.IsNotStruct(document) {
			errs = append(errs, errors.New(ErrDocumentIsNotStruct.Error()+"(index: "+indexStr+")"))
		} else if util.IsZero(document) {
			errs = append(errs, errors.New(ErrDocumentIsEmpty.Error()+"(index: "+indexStr+")"))
		} else {
			err := t.insertOne(sc, document, option.InsertOne{
				BypassDocumentValidation: opt.BypassDocumentValidation,
				Comment:                  opt.Comment,
			})
			if err != nil {
				errs = append(errs, errors.New(err.Error()+" (index: "+indexStr+")"))
			}
		}
	}
	if len(errs) != 0 {
		var b strings.Builder
		for i, err := range errs {
			if i != 0 {
				b.WriteString(", ")
			}
			b.WriteString(err.Error())
		}
		return errors.New(b.String())
	}
	return nil
}

func (t *Template) deleteOne(sc mongo.SessionContext, filter, ref any, opt option.Delete) (*DeleteResult, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	mongoResult, err := coll.DeleteOne(sc, filter, &options.DeleteOptions{
		Collation: option.ParseCollationMongoOptions(opt.Collation),
		Comment:   opt.Comment,
		Hint:      opt.Hint,
		Let:       opt.Let,
	})
	var result *DeleteResult
	if mongoResult != nil {
		result = &DeleteResult{
			DeletedCount: mongoResult.DeletedCount,
		}
	}
	return result, err
}

func (t *Template) deleteMany(sc mongo.SessionContext, filter, ref any, opt option.Delete) (*DeleteResult, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	mongoResult, err := coll.DeleteMany(sc, filter, &options.DeleteOptions{
		Collation: option.ParseCollationMongoOptions(opt.Collation),
		Comment:   opt.Comment,
		Hint:      opt.Hint,
		Let:       opt.Let,
	})
	var result *DeleteResult
	if mongoResult != nil {
		result = &DeleteResult{
			DeletedCount: mongoResult.DeletedCount,
		}
	}
	return result, err
}

func (t *Template) updateOne(sc mongo.SessionContext, filter, update, ref any, opt option.Update) (*UpdateResult, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	mongoResult, err := coll.UpdateOne(sc, filter, update, &options.UpdateOptions{
		ArrayFilters:             option.ParseArrayFiltersMongoOptions(opt.ArrayFilters),
		BypassDocumentValidation: opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		Hint:                     opt.Hint,
		Upsert:                   opt.Upsert,
		Let:                      opt.Let,
	})
	var result *UpdateResult
	if mongoResult != nil {
		result = &UpdateResult{
			MatchedCount:  mongoResult.MatchedCount,
			ModifiedCount: mongoResult.ModifiedCount,
			UpsertedCount: mongoResult.UpsertedCount,
			UpsertedID:    mongoResult.UpsertedID,
		}
	}
	return result, err
}

func (t *Template) updateMany(sc mongo.SessionContext, filter, update, ref any, opt option.Update) (*UpdateResult, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	mongoResult, err := coll.UpdateMany(sc, filter, update, &options.UpdateOptions{
		ArrayFilters:             option.ParseArrayFiltersMongoOptions(opt.ArrayFilters),
		BypassDocumentValidation: opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		Hint:                     opt.Hint,
		Upsert:                   opt.Upsert,
		Let:                      opt.Let,
	})
	var result *UpdateResult
	if mongoResult != nil {
		result = &UpdateResult{
			MatchedCount:  mongoResult.MatchedCount,
			ModifiedCount: mongoResult.ModifiedCount,
			UpsertedCount: mongoResult.UpsertedCount,
			UpsertedID:    mongoResult.UpsertedID,
		}
	}
	return result, err
}

func (t *Template) replaceOne(sc mongo.SessionContext, filter, update, ref any, opt option.Replace) (*UpdateResult,
	error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	mongoResult, err := coll.ReplaceOne(sc, filter, update, &options.ReplaceOptions{
		BypassDocumentValidation: opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		Hint:                     opt.Hint,
		Upsert:                   opt.Upsert,
		Let:                      opt.Let,
	})
	var result *UpdateResult
	if mongoResult != nil {
		result = &UpdateResult{
			MatchedCount:  mongoResult.MatchedCount,
			ModifiedCount: mongoResult.ModifiedCount,
			UpsertedCount: mongoResult.UpsertedCount,
			UpsertedID:    mongoResult.UpsertedID,
		}
	}
	return result, err
}

func (t *Template) find(ctx context.Context, filter, dest any, opts ...option.Find) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	}
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	opt := option.GetFindOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName)
	cursor, err := collection.Find(ctx, filter, &options.FindOptions{
		AllowDiskUse:        opt.AllowDiskUse,
		AllowPartialResults: opt.AllowPartialResults,
		BatchSize:           opt.BatchSize,
		Collation:           option.ParseCollationMongoOptions(opt.Collation),
		Comment:             opt.Comment,
		CursorType:          option.ParseCursorType(opt.CursorType),
		Hint:                opt.Hint,
		Limit:               opt.Limit,
		Max:                 opt.Max,
		MaxAwaitTime:        opt.MaxAwaitTime,
		MaxTime:             opt.MaxTime,
		Min:                 opt.Min,
		NoCursorTimeout:     opt.NoCursorTimeout,
		Projection:          opt.Projection,
		ReturnKey:           opt.ReturnKey,
		ShowRecordID:        opt.ShowRecordID,
		Skip:                opt.Skip,
		Sort:                opt.Sort,
		Let:                 opt.Let,
	})
	if err != nil {
		return err
	}
	return cursor.All(ctx, dest)
}

func (t *Template) findOne(ctx context.Context, filter, dest any, opts ...option.FindOne) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	opt := option.GetFindOneOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName)
	err = collection.FindOne(ctx, filter, &options.FindOneOptions{
		AllowPartialResults: opt.AllowPartialResults,
		Collation:           option.ParseCollationMongoOptions(opt.Collation),
		Comment:             opt.Comment,
		Hint:                opt.Hint,
		Max:                 opt.Max,
		MaxTime:             opt.MaxTime,
		Min:                 opt.Min,
		Projection:          opt.Projection,
		ReturnKey:           opt.ReturnKey,
		ShowRecordID:        opt.ShowRecordID,
		Skip:                opt.Skip,
		Sort:                opt.Sort,
	}).Decode(dest)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	return nil
}

func (t *Template) findOneAndDelete(sc mongo.SessionContext, filter, dest any, opt option.FindOneAndDelete) error {
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	err = coll.FindOneAndDelete(sc, filter, &options.FindOneAndDeleteOptions{
		Collation:  option.ParseCollationMongoOptions(opt.Collation),
		Comment:    opt.Comment,
		MaxTime:    opt.MaxTime,
		Projection: opt.Projection,
		Sort:       opt.Sort,
		Hint:       opt.Hint,
		Let:        opt.Let,
	}).Decode(dest)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return ErrNoDocuments
	} else if err != nil {
		return err
	}
	return nil
}

func (t *Template) findOneAndReplace(sc mongo.SessionContext, filter, replacement, dest any, opt option.FindOneAndReplace) error {
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	err = coll.FindOneAndReplace(sc, filter, replacement, &options.FindOneAndReplaceOptions{
		BypassDocumentValidation: opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		MaxTime:                  opt.MaxTime,
		Projection:               opt.Projection,
		ReturnDocument:           option.ParseReturnDocument(opt.ReturnDocument),
		Sort:                     opt.Sort,
		Upsert:                   opt.Upsert,
		Hint:                     opt.Hint,
		Let:                      opt.Let,
	}).Decode(dest)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return ErrNoDocuments
	} else if err != nil {
		return err
	}
	return nil
}

func (t *Template) findOneAndUpdate(sc mongo.SessionContext, filter, update, dest any, opt option.FindOneAndUpdate) error {
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	err = coll.FindOneAndUpdate(sc, filter, update, &options.FindOneAndUpdateOptions{
		ArrayFilters:             option.ParseArrayFiltersMongoOptions(opt.ArrayFilters),
		BypassDocumentValidation: opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		MaxTime:                  opt.MaxTime,
		Projection:               opt.Projection,
		ReturnDocument:           option.ParseReturnDocument(opt.ReturnDocument),
		Sort:                     opt.Sort,
		Upsert:                   opt.Upsert,
		Hint:                     opt.Hint,
		Let:                      opt.Let,
	}).Decode(dest)
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {
		return ErrNoDocuments
	} else if err != nil {
		return err
	}
	return nil
}

func (t *Template) createOneIndex(ctx context.Context, input IndexInput) (string, error) {
	databaseName, collectionName, err := getMongoInfosByAny(input.Ref)
	if err != nil {
		return "", err
	}
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName).Indexes()
	return collection.CreateOne(ctx, parseIndexInputToModel(input))
}

func (t *Template) createManyIndex(ctx context.Context, inputs []IndexInput) ([]string, error) {
	var result []string
	if len(inputs) == 0 {
		return result, ErrDocumentsIsEmpty
	}
	var errs []error
	for i, input := range inputs {
		indexStr := strconv.Itoa(i)
		r, err := t.createOneIndex(ctx, input)
		if err != nil {
			errs = append(errs, errors.New(err.Error()+" (index: "+indexStr+")"))
		} else {
			result = append(result, r)
		}
	}
	if len(errs) != 0 {
		var b strings.Builder
		for i, errResult := range errs {
			if i != 0 {
				b.WriteString(", ")
			}
			b.WriteString(errResult.Error())
		}
		return result, errors.New(b.String())
	}
	return result, nil
}

func (t *Template) startSession(ctx context.Context, forceRecreate bool) error {
	if t.session != nil && !forceRecreate {
		return nil
	} else if t.session != nil {
		t.endSession(ctx)
	}
	session, err := t.client.StartSession()
	if err == nil {
		err = session.StartTransaction()
		if err == nil {
			t.session = session
		}
	}
	return err
}

func (t *Template) closeSession(
	sc mongo.SessionContext,
	disableAutoCloseSession,
	disableAutoRollbackSession bool,
	err error,
) error {
	abort := err != nil && !disableAutoRollbackSession
	if !disableAutoCloseSession {
		errClose := t.CloseSession(sc, abort)
		if errClose != nil {
			return errClose
		}
	}
	return err
}

func (t *Template) endSession(ctx context.Context) {
	if t.session != nil {
		t.session.EndSession(ctx)
		t.session = nil
	}
}

func getMongoInfosByAny(a any) (databaseName string, collectionName string, err error) {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		if v.Type().Elem().Kind() == reflect.Struct {
			databaseName = util.GetDatabaseNameBySlice(a)
			collectionName = util.GetCollectionNameBySlice(a)
		} else {
			return "", "", ErrRefDocument
		}
		break
	case reflect.Struct:
		databaseName = util.GetDatabaseNameByStruct(a)
		collectionName = util.GetCollectionNameByStruct(a)
		break
	default:
		return "", "", ErrRefDocument
	}
	if len(databaseName) == 0 {
		return "", "", ErrDatabaseNotConfigured
	}
	if len(collectionName) == 0 {
		return "", "", ErrCollectionNotConfigured
	}
	return collectionName, databaseName, nil
}
