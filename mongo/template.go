package mongo

import (
	"context"
	"errors"
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

type template struct {
	client  *mongo.Client
	session mongo.Session
}

type Template interface {
	// InsertOne executes an insert command to insert a single document into the collection.
	//
	// The document parameter must be the document to be inserted. It cannot be nil. If the document does not have an _id
	// field when transformed into BSON, one will be added automatically to the marshalled document. The original document
	// will not be modified. The _id can be retrieved from the InsertedID field of the returned InsertOneResult.
	//
	// The opts parameter can be used to specify options for the operation (see the option.InsertOne documentation.)
	//
	// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/insert/.
	InsertOne(ctx context.Context, document any, opts ...option.InsertOne) error
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
	InsertMany(ctx context.Context, documents any, opts ...option.InsertMany) error
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
	DeleteOne(ctx context.Context, filter, ref any, opts ...option.Delete) (*mongo.DeleteResult, error)
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
	DeleteOneById(ctx context.Context, id, ref any, opts ...option.Delete) (*mongo.DeleteResult, error)
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
	DeleteMany(ctx context.Context, filter, ref any, opts ...option.Delete) (*mongo.DeleteResult, error)
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
	UpdateOneById(ctx context.Context, id, update, ref any, opts ...option.Update) (*mongo.UpdateResult, error)
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
	UpdateOne(ctx context.Context, filter, update, ref any, opts ...option.Update) (*mongo.UpdateResult, error)
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
	UpdateMany(ctx context.Context, filter, update, ref any, opts ...option.Update) (*mongo.UpdateResult, error)
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
	ReplaceOne(ctx context.Context, filter, replacement, ref any, opts ...option.Replace) (*mongo.UpdateResult, error)
	// ReplaceOneById executes an update command to update the document whose _id value matches the provided ID in the collection.
	// This is equivalent to running ReplaceOne(ctx, bson.D{{"_id", id}}, replacement, ref, opts...).
	//
	// The id parameter is the _id of the document to be updated. It cannot be nil. If the ID does not match any documents,
	// the operation will succeed and an UpdateResult with a MatchedCount of 0 will be returned.
	//
	// The opts parameter can be used to specify options for the operation (see the option.Replace documentation).
	//
	// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/update/.
	ReplaceOneById(ctx context.Context, id, replacement, ref any, opts ...option.Replace) (*mongo.UpdateResult, error)
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
	FindOne(ctx context.Context, filter, dest any, opts ...option.FindOne) error
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
	FindOneById(ctx context.Context, id, dest any, opts ...option.FindOneById) error
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
	FindOneAndDelete(ctx context.Context, filter, dest any, opts ...option.FindOneAndDelete) error
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
	FindOneAndReplace(ctx context.Context, filter, replacement, dest any, opts ...option.FindOneAndReplace) error
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
	FindOneAndUpdate(ctx context.Context, filter, update, dest any, opts ...option.FindOneAndUpdate) error
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
	Find(ctx context.Context, filter, dest any, opts ...option.Find) error
	// FindAll execute a search command. This is equivalent to running Find(ctx, bson.D{}, dest, opts...).
	//
	// The dest parameter must be a pointer to the return expected by the operation, it is important to have the
	// database and collection tags configured.
	//
	// The opts parameter can be used to specify options for the operation (see the option.Find documentation).
	//
	// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/find/.
	FindAll(ctx context.Context, dest any, opts ...option.Find) error
	// FindPageable executes a find command, if successful, returns the paginated documents in the
	// corresponding PageOutput structure in the collection on the target parameter with null return error.
	// Otherwise, it will return the corresponding error.
	//
	// The filter parameter must be a document containing query operators and can be used to select which documents are
	// included in the result. It cannot be nil. If the filter does not match any document, the return structure columns
	// will be empty.
	//
	// The opts parameter can be used to specify options for the operation (see the option.FindPageable documentation).
	//
	// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/find/.
	FindPageable(ctx context.Context, filter any, input PageInput, opts ...option.FindPageable) (*PageOutput, error)
	// Exists executes the count command, if the quantity is greater than 0 with a limit of 1, true is returned,
	// otherwise false is returned.
	//
	// The ref parameter must be the collection structure with database and collection tags configured.
	//
	// The opts parameter can be used to specify options for the operation (see the option.Exists documentation).
	Exists(ctx context.Context, filter, ref any, opts ...option.Exists) (bool, error)
	// ExistsById executes a count command whose _id value matches the ID given in the collection.
	// This is equivalent to running Exists(ctx, bson.D{{"_id", id}}, dest, opts...).
	//
	// The ref parameter must be the collection structure with database and collection tags configured.
	//
	// The opts parameter can be used to specify options for the operation (see the option.Exists documentation).
	ExistsById(ctx context.Context, id, ref any, opts ...option.Exists) (bool, error)
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
	Aggregate(ctx context.Context, pipeline, dest any, opts ...option.Aggregate) error
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
	CountDocuments(ctx context.Context, filter, ref any, opts ...option.Count) (int64, error)
	// EstimatedDocumentCount executes a count command and returns an estimate of the number of documents in the collection
	// using collection metadata.
	//
	// The ref parameter must be the collection structure with database and collection tags configured.
	//
	// The opts parameter can be used to specify options for the operation (see the option.EstimatedDocumentCount
	// documentation).
	//
	// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/count/.
	EstimatedDocumentCount(ctx context.Context, ref any, opts ...option.EstimatedDocumentCount) (int64, error)
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
	Distinct(ctx context.Context, fieldName string, filter, dest, ref any, opts ...option.Distinct) error
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
	Watch(ctx context.Context, pipeline any, opts ...option.Watch) (*mongo.ChangeStream, error)
	// WatchHandler is a function that facilitates the reading of watch events, it triggers the Watch function and
	// when an event occurs it reads this event transforming all the data obtained by mongoDB into a Context, right
	// after this conversion, we call the handler parameter passing the context with all the information, so you can
	// process it in the way you see fit.
	//
	// The opts parameter can be used to specify options for change stream creation (see the option.WatchHandler
	// documentation).
	WatchHandler(ctx context.Context, pipeline any, handler HandlerWatch, opts ...option.WatchHandler) error
	// DropCollection drops the collection on the server. This method ignores "namespace not found" errors,
	// so it is safe to drop a collection that does not exist on the server.
	//
	// The ref parameter must be the collection structure with database and collection tags configured.
	DropCollection(ctx context.Context, ref any) error
	// DropDatabase drops the database on the server. This method ignores "namespace not found" errors,
	// so it is safe to drop a database that does not exist on the server.
	//
	// The ref parameter must be the collection structure with database and collection tags configured.
	DropDatabase(ctx context.Context, ref any) error
	// CreateOneIndex executes a createIndexes command to create an index on the collection and returns the name of the new
	// index. See the CreateManyIndex documentation for more information and an example. For this function's response,
	// the name of the index is returned as a string, and if an error occurs, it is returned in the second return parameter
	//
	// The opts parameter can be used to specify options for this operation (see the option.Index documentation).
	//
	// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/createIndexes/.
	CreateOneIndex(ctx context.Context, input IndexInput) (string, error)
	// CreateManyIndex executes a createIndexes command to create multiple indexes on the collection and returns the names of
	// the new indexes.
	//
	// For each IndexInput in the models parameter, the index name can be specified via the Options field. If a name is not
	// given, it will be generated from the Keys document.
	//
	// The opts parameter can be used to specify options for this operation (see the option.Index documentation).
	//
	// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/createIndexes/.
	CreateManyIndex(ctx context.Context, inputs []IndexInput) ([]string, error)
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
	DropOneIndex(ctx context.Context, name string, ref any, opts ...option.DropIndex) error
	// DropAllIndexes executes a dropIndexes operation to drop all indexes on the collection. If the operation succeeds, this
	// returns a BSON document in the form {nIndexesWas: <int32>}. The "nIndexesWas" field in the response contains the
	// number of indexes that existed prior to the drop.
	//
	// The ref parameter must be the collection structure with database and collection tags configured.
	//
	// The opts parameter can be used to specify options for this operation (see the option.DropIndex documentation).
	//
	// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/dropIndexes/.
	DropAllIndexes(ctx context.Context, ref any, opts ...option.DropIndex) error
	// ListIndexes executes a listIndexes command and returns a cursor over the indexes in the collection.
	//
	// The opts parameter can be used to specify options for this operation (see the option.ListIndexes documentation).
	//
	// The ref parameter must be the collection structure with database and collection tags configured.
	//
	// For more information about the command, see https://www.mongodb.com/docs/manual/reference/command/listIndexes/.
	ListIndexes(ctx context.Context, ref any, opts ...option.ListIndexes) ([]IndexOutput, error)
	// ListIndexSpecifications executes a List command and returns a slice of returned IndexSpecifications.
	//
	// The ref parameter must be the collection structure with database and collection tags configured.
	ListIndexSpecifications(ctx context.Context, ref any, opts ...option.ListIndexes) ([]*mongo.IndexSpecification, error)
	// CloseSession closes session and transaction, if param abort is false it will commit the changes,
	// otherwise it will abort all transactions.
	CloseSession(ctx context.Context, abort bool) error
	// CommitTransaction commit all transactions on session
	CommitTransaction(ctx context.Context) error
	// AbortTransaction abort all transactions on session
	AbortTransaction(ctx context.Context) error
	// Disconnect
	// closes the mongodb connection client without return error
	Disconnect(ctx context.Context)
	// DisconnectWithErr
	// closes the mongodb connection client with return error
	DisconnectWithErr(ctx context.Context) error
}

func NewTemplate(ctx context.Context, opts ...*options.ClientOptions) (Template, error) {
	conn, err := mongo.Connect(ctx, opts...)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &template{
		client: conn,
	}, nil
}

func (t *template) InsertOne(ctx context.Context, document any, opts ...option.InsertOne) error {
	opt := option.GetInsertOneOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		err := t.insertOne(sc, document, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
}

func (t *template) InsertMany(ctx context.Context, documents any, opts ...option.InsertMany) error {
	opt := option.GetInsertManyOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		err := t.insertMany(sc, documents, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, opt.DisableAutoRollbackSession, err)
	})
}

func (t *template) DeleteOne(ctx context.Context, filter, ref any, opts ...option.Delete) (*mongo.DeleteResult, error) {
	var result *mongo.DeleteResult
	var err error
	opt := option.GetDeleteOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.deleteOne(sc, filter, ref, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
	return result, err
}

func (t *template) DeleteOneById(ctx context.Context, id, ref any, opts ...option.Delete) (*mongo.DeleteResult, error) {
	var result *mongo.DeleteResult
	var err error
	opt := option.GetDeleteOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.deleteOne(sc, bson.D{{"_id", id}}, ref, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
	return result, err
}

func (t *template) DeleteMany(ctx context.Context, filter, ref any, opts ...option.Delete) (*mongo.DeleteResult, error) {
	var result *mongo.DeleteResult
	var err error
	opt := option.GetDeleteOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.deleteMany(sc, filter, ref, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
	return result, err
}

func (t *template) UpdateOneById(ctx context.Context, id, update, ref any, opts ...option.Update) (*mongo.UpdateResult,
	error) {
	var result *mongo.UpdateResult
	var err error
	opt := option.GetUpdateOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.updateOne(sc, bson.D{{"_id", id}}, update, ref, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
	return result, err
}

func (t *template) UpdateOne(ctx context.Context, filter any, update, ref any, opts ...option.Update) (
	*mongo.UpdateResult, error) {
	var result *mongo.UpdateResult
	var err error
	opt := option.GetUpdateOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.updateOne(sc, filter, update, ref, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
	return result, err
}

func (t *template) UpdateMany(ctx context.Context, filter any, update, ref any, opts ...option.Update) (
	*mongo.UpdateResult, error) {
	var result *mongo.UpdateResult
	var err error
	opt := option.GetUpdateOptionByParams(opts)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.updateMany(sc, filter, update, ref, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
	return result, err
}

func (t *template) ReplaceOne(ctx context.Context, filter any, update, ref any, opts ...option.Replace) (
	*mongo.UpdateResult, error) {
	var result *mongo.UpdateResult
	var err error
	opt := option.GetReplaceOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.replaceOne(sc, filter, update, ref, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
	return result, err
}

func (t *template) ReplaceOneById(ctx context.Context, id, replacement, ref any, opts ...option.Replace) (
	*mongo.UpdateResult, error) {
	var result *mongo.UpdateResult
	var err error
	opt := option.GetReplaceOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.replaceOne(sc, bson.D{{"_id", id}}, replacement, ref, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
	return result, err
}

func (t *template) FindOneById(ctx context.Context, id, dest any, opts ...option.FindOneById) error {
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

func (t *template) FindOne(ctx context.Context, filter, dest any, opts ...option.FindOne) error {
	return t.findOne(ctx, filter, dest, opts...)
}

func (t *template) FindOneAndDelete(ctx context.Context, filter, dest any, opts ...option.FindOneAndDelete) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	opt := option.GetFindOneAndDeleteOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		err := t.findOneAndDelete(sc, filter, dest, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
}

func (t *template) FindOneAndReplace(ctx context.Context, filter, replacement, dest any, opts ...option.FindOneAndReplace) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	opt := option.GetFindOneAndReplaceOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		err := t.findOneAndReplace(sc, filter, replacement, dest, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
}

func (t *template) FindOneAndUpdate(ctx context.Context, filter, update, dest any, opts ...option.FindOneAndUpdate) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	opt := option.GetFindOneAndUpdateOptionByParams(opts)
	t.startSession(ctx, opt.ForceRecreateSession)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		err := t.findOneAndUpdate(sc, filter, update, dest, opt)
		return t.closeSession(sc, opt.DisableAutoCloseSession, false, err)
	})
}

func (t *template) Find(ctx context.Context, filter, dest any, opts ...option.Find) error {
	return t.find(ctx, filter, dest, opts...)
}

func (t *template) FindAll(ctx context.Context, dest any, opts ...option.Find) error {
	return t.find(ctx, bson.D{}, dest, opts...)
}

func (t *template) FindPageable(ctx context.Context, filter any, input PageInput, opts ...option.FindPageable) (
	*PageOutput, error) {
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
	return NewPageOutput(input, dest, countTotal), nil
}

func (t *template) Exists(ctx context.Context, filter, ref any, opts ...option.Exists) (bool, error) {
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

func (t *template) ExistsById(ctx context.Context, id, ref any, opts ...option.Exists) (bool, error) {
	return t.Exists(ctx, bson.D{{"_id", id}}, ref, opts...)
}

func (t *template) Aggregate(ctx context.Context, pipeline any, dest any, opts ...option.Aggregate) error {
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

func (t *template) CountDocuments(ctx context.Context, filter, ref any, opts ...option.Count) (int64, error) {
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

func (t *template) EstimatedDocumentCount(ctx context.Context, ref any, opts ...option.EstimatedDocumentCount) (int64,
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

func (t *template) Distinct(ctx context.Context, fieldName string, filter, dest, ref any, opts ...option.Distinct) error {
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

func (t *template) Watch(ctx context.Context, pipeline any, opts ...option.Watch) (*mongo.ChangeStream, error) {
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

func (t *template) WatchHandler(ctx context.Context, pipeline any, handler HandlerWatch, opts ...option.WatchHandler) error {
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

func (t *template) DropCollection(ctx context.Context, ref any) error {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return err
	}
	return t.client.Database(databaseName).Collection(collectionName).Drop(ctx)
}

func (t *template) DropDatabase(ctx context.Context, ref any) error {
	databaseName, _, err := getMongoInfosByAny(ref)
	if err != nil {
		return err
	}
	return t.client.Database(databaseName).Drop(ctx)
}

func (t *template) CreateOneIndex(ctx context.Context, input IndexInput) (string, error) {
	return t.createOneIndex(ctx, input)
}

func (t *template) CreateManyIndex(ctx context.Context, inputs []IndexInput) ([]string, error) {
	return t.createManyIndex(ctx, inputs)
}

func (t *template) DropOneIndex(ctx context.Context, name string, ref any, opts ...option.DropIndex) error {
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

func (t *template) DropAllIndexes(ctx context.Context, ref any, opts ...option.DropIndex) error {
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

func (t *template) ListIndexes(ctx context.Context, ref any, opts ...option.ListIndexes) ([]IndexOutput, error) {
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
	var results []IndexOutput
	err = cursor.All(ctx, &results)
	return results, err
}

func (t *template) ListIndexSpecifications(ctx context.Context, ref any, opts ...option.ListIndexes) (
	[]*mongo.IndexSpecification, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	opt := option.GetListIndexesOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName).Indexes()
	return collection.ListSpecifications(ctx, &options.ListIndexesOptions{
		BatchSize: opt.BatchSize,
		MaxTime:   opt.MaxTime,
	})
}

func (t *template) Disconnect(ctx context.Context) {
	_ = t.DisconnectWithErr(ctx)
}

func (t *template) DisconnectWithErr(ctx context.Context) error {
	return t.client.Disconnect(ctx)
}

func (t *template) CloseSession(ctx context.Context, abort bool) error {
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

func (t *template) CommitTransaction(ctx context.Context) error {
	if t.session == nil {
		return nil
	}
	return t.session.CommitTransaction(ctx)
}

func (t *template) AbortTransaction(ctx context.Context) error {
	if t.session == nil {
		return nil
	}
	return t.session.AbortTransaction(ctx)
}

func (t *template) insertOne(sc mongo.SessionContext, document any, opt option.InsertOne) error {
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

func (t *template) insertMany(sc mongo.SessionContext, a any, opt option.InsertMany) error {
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
		for i, errResult := range errs {
			if i != 0 {
				b.WriteString(", ")
			}
			b.WriteString(errResult.Error())
		}
		return errors.New(b.String())
	}
	return nil
}

func (t *template) deleteOne(sc mongo.SessionContext, filter, ref any, opt option.Delete) (*mongo.DeleteResult, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.DeleteOne(sc, filter, &options.DeleteOptions{
		Collation: option.ParseCollationMongoOptions(opt.Collation),
		Comment:   opt.Comment,
		Hint:      opt.Hint,
		Let:       opt.Let,
	})
}

func (t *template) deleteMany(sc mongo.SessionContext, filter, ref any, opt option.Delete) (*mongo.DeleteResult, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.DeleteMany(sc, filter, &options.DeleteOptions{
		Collation: option.ParseCollationMongoOptions(opt.Collation),
		Comment:   opt.Comment,
		Hint:      opt.Hint,
		Let:       opt.Let,
	})
}

func (t *template) updateOne(sc mongo.SessionContext, filter, update, ref any, opt option.Update) (*mongo.UpdateResult,
	error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.UpdateOne(sc, filter, update, &options.UpdateOptions{
		ArrayFilters:             option.ParseArrayFiltersMongoOptions(opt.ArrayFilters),
		BypassDocumentValidation: opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		Hint:                     opt.Hint,
		Upsert:                   opt.Upsert,
		Let:                      opt.Let,
	})
}

func (t *template) updateMany(sc mongo.SessionContext, filter, update, ref any, opt option.Update) (*mongo.UpdateResult,
	error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.UpdateMany(sc, filter, update, &options.UpdateOptions{
		ArrayFilters:             option.ParseArrayFiltersMongoOptions(opt.ArrayFilters),
		BypassDocumentValidation: opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		Hint:                     opt.Hint,
		Upsert:                   opt.Upsert,
		Let:                      opt.Let,
	})
}

func (t *template) replaceOne(sc mongo.SessionContext, filter, update, ref any, opt option.Replace) (*mongo.UpdateResult,
	error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.ReplaceOne(sc, filter, update, &options.ReplaceOptions{
		BypassDocumentValidation: opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		Hint:                     opt.Hint,
		Upsert:                   opt.Upsert,
		Let:                      opt.Let,
	})
}

func (t *template) find(ctx context.Context, filter, dest any, opts ...option.Find) error {
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

func (t *template) findOne(ctx context.Context, filter, dest any, opts ...option.FindOne) error {
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

func (t *template) findOneAndDelete(sc mongo.SessionContext, filter, dest any, opt option.FindOneAndDelete) error {
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

func (t *template) findOneAndReplace(sc mongo.SessionContext, filter, replacement, dest any, opt option.FindOneAndReplace) error {
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

func (t *template) findOneAndUpdate(sc mongo.SessionContext, filter, update, dest any, opt option.FindOneAndUpdate) error {
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

func (t *template) createOneIndex(ctx context.Context, input IndexInput) (string, error) {
	databaseName, collectionName, err := getMongoInfosByAny(input.Ref)
	if err != nil {
		return "", err
	}
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName).Indexes()
	return collection.CreateOne(ctx, parseIndexInputToModel(input))
}

func (t *template) createManyIndex(ctx context.Context, inputs []IndexInput) ([]string, error) {
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

func (t *template) startSession(ctx context.Context, forceSession bool) {
	if t.session != nil && !forceSession {
		return
	} else if t.session != nil {
		_ = t.CloseSession(ctx, false)
	}
	session, _ := t.client.StartSession()
	if session != nil {
		_ = session.StartTransaction()
		t.session = session
	}
}

func (t *template) closeSession(
	sc mongo.SessionContext,
	DisableAutoCloseSession,
	DisableAutoRollbackSession bool,
	err error,
) error {
	abort := err != nil && !DisableAutoRollbackSession
	if !DisableAutoCloseSession {
		errClose := t.CloseSession(sc, abort)
		if errClose != nil {
			return errClose
		}
	}
	return err
}

func (t *template) endSession(ctx context.Context) {
	if t.session == nil {
		return
	}
	t.session.EndSession(ctx)
	t.session = nil
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
