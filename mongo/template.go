package mongo

import (
	"context"
	"errors"
	"github.com/GabrielHCataldo/go-logger/logger"
	"go-mongo/internal/util"
	"go-mongo/mongo/option"
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
//	mongo.Pipeline{
//		{{"$group", bson.D{{"_id", "$state"}, {"totalPop", bson.D{{"$sum", "$pop"}}}}}},
//		{{"$match", bson.D{{"totalPop", bson.D{{"$gte", 10*1000*1000}}}}}},
//	}
type Pipeline []bson.D

type template struct {
	client  *mongo.Client
	session mongo.Session
}

type Template interface {
	InsertOne(ctx context.Context, document any, opts ...option.InsertOne) error
	InsertMany(ctx context.Context, documents []any, opts ...option.InsertMany) error
	DeleteOne(ctx context.Context, filter, ref any, opts ...option.Delete) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter, ref any, opts ...option.Delete) (*mongo.DeleteResult, error)
	UpdateOneById(ctx context.Context, id, update, ref any, opts ...option.Update) (*mongo.UpdateResult, error)
	UpdateOne(ctx context.Context, filter, update, ref any, opts ...option.Update) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter, update, ref any, opts ...option.Update) (*mongo.UpdateResult, error)
	ReplaceOne(ctx context.Context, filter, replacement, ref any, opts ...option.Replace) (*mongo.UpdateResult, error)
	FindOne(ctx context.Context, filter, dest any, opts ...option.FindOne) error
	FindOneById(ctx context.Context, id, dest any, opts ...option.FindOne) error
	FindOneAndDelete(ctx context.Context, filter, dest any, opts ...option.FindOneAndDelete) error
	FindOneAndReplace(ctx context.Context, filter, replacement, dest any, opts ...option.FindOneAndReplace) error
	FindOneAndUpdate(ctx context.Context, filter, update, dest any, opts ...option.FindOneAndUpdate) error
	Find(ctx context.Context, filter, dest any, opts ...option.Find) error
	FindPageable(ctx context.Context, filter any, input PageInput, opts ...option.FindPageable) (*PageOutput, error)
	Aggregate(ctx context.Context, pipeline, dest any, opts ...option.Aggregate) error
	CountDocuments(ctx context.Context, filter, ref any, opts ...option.Count) (int64, error)
	EstimatedDocumentCount(ctx context.Context, ref any, opts ...option.EstimatedDocumentCount) (int64, error)
	Distinct(ctx context.Context, fieldName string, filter, dest, ref any, opts ...option.Distinct) error
	Watch(ctx context.Context, pipeline any, opts ...option.Watch) (*mongo.ChangeStream, error)
	WatchHandler(ctx context.Context, pipeline any, handler HandlerWatch, opts ...option.WatchHandler) error
	DropCollection(ctx context.Context, ref any) error
	DropDatabase(ctx context.Context, ref any) error
	// CreateOneIndex
	//
	// # Parameters:
	//
	// - ref:
	//
	// - value: A document describing which keys should be used for the index. It cannot be nil. This must be an
	// order-preserving type such as bson.D. Map types such as bson.M are not valid.
	// See https://www.mongodb.com/docs/manual/indexes/#indexes for examples of valid documents.
	CreateOneIndex(ctx context.Context, input IndexInput) (string, error)
	CreateManyIndex(ctx context.Context, inputs []IndexInput) ([]string, error)
	DropOneIndex(ctx context.Context, name string, ref any, opts ...option.DropIndex) error
	DropAllIndexes(ctx context.Context, ref any, opts ...option.DropIndex) error
	ListIndexes(ctx context.Context, ref any, opts ...option.ListIndexes) ([]IndexOutput, error)
	ListIndexSpecifications(ctx context.Context, ref any, opts ...option.ListIndexes) ([]*mongo.IndexSpecification,
		error)
	StartSession(ctx context.Context, forceSession bool)
	CloseSession(ctx context.Context, err error)
	Disconnect()
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
	t.StartSession(ctx, opt.ForceRecreateSession)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		err := t.insertOne(sc, document, opt)
		if !opt.DisableAutoCloseSession {
			t.CloseSession(sc, err)
		}
		return err
	})
}

func (t *template) InsertMany(ctx context.Context, documents []any, opts ...option.InsertMany) error {
	opt := option.GetInsertManyOptionByParams(opts)
	t.StartSession(ctx, opt.ForceRecreateSession)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		err := t.insertMany(sc, documents, opt)
		if !opt.DisableAutoCloseSession {
			errClose := err
			if !opt.DisableAutoRollbackSession {
				errClose = nil
			}
			t.CloseSession(sc, errClose)
		}
		return err
	})
}

func (t *template) DeleteOne(ctx context.Context, filter, ref any, opts ...option.Delete) (
	*mongo.DeleteResult, error) {
	var result *mongo.DeleteResult
	var err error
	opt := option.GetDeleteOptionByParams(opts)
	t.StartSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.deleteOne(sc, filter, ref, opt)
		if !opt.DisableAutoCloseSession {
			t.CloseSession(sc, err)
		}
		return err
	})
	return result, err
}

func (t *template) DeleteMany(ctx context.Context, filter, ref any, opts ...option.Delete) (
	*mongo.DeleteResult, error) {
	var result *mongo.DeleteResult
	var err error
	opt := option.GetDeleteOptionByParams(opts)
	t.StartSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.deleteMany(sc, filter, ref, opt)
		if !opt.DisableAutoCloseSession {
			t.CloseSession(sc, err)
		}
		return err
	})
	return result, err
}

func (t *template) UpdateOneById(ctx context.Context, id, update, ref any, opts ...option.Update) (
	*mongo.UpdateResult, error) {
	var result *mongo.UpdateResult
	var err error
	opt := option.GetUpdateOptionByParams(opts)
	t.StartSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.updateOne(sc, bson.D{{"_id", id}}, update, ref, opt)
		if !opt.DisableAutoCloseSession {
			t.CloseSession(sc, err)
		}
		return err
	})
	return result, err
}

func (t *template) UpdateOne(ctx context.Context, filter any, update, ref any, opts ...option.Update) (
	*mongo.UpdateResult, error) {
	var result *mongo.UpdateResult
	var err error
	opt := option.GetUpdateOptionByParams(opts)
	t.StartSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.updateOne(sc, filter, update, ref, opt)
		if !opt.DisableAutoCloseSession {
			t.CloseSession(sc, err)
		}
		return err
	})
	return result, err
}

func (t *template) UpdateMany(ctx context.Context, filter any, update, ref any, opts ...option.Update) (
	*mongo.UpdateResult, error) {
	var result *mongo.UpdateResult
	var err error
	opt := option.GetUpdateOptionByParams(opts)
	t.StartSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.updateMany(sc, filter, update, ref, opt)
		if !opt.DisableAutoCloseSession {
			t.CloseSession(sc, err)
		}
		return err
	})
	return result, err
}

func (t *template) ReplaceOne(ctx context.Context, filter any, update, ref any, opts ...option.Replace) (
	*mongo.UpdateResult, error) {
	var result *mongo.UpdateResult
	var err error
	opt := option.GetReplaceOptionByParams(opts)
	t.StartSession(ctx, opt.ForceRecreateSession)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		result, err = t.replaceOne(sc, filter, update, ref, opt)
		if !opt.DisableAutoCloseSession {
			t.CloseSession(sc, err)
		}
		return err
	})
	return result, err
}

func (t *template) FindOneById(ctx context.Context, id, dest any, opts ...option.FindOne) error {
	return t.findOne(ctx, bson.D{{"_id", id}}, dest, opts...)
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
	t.StartSession(ctx, opt.ForceRecreateSession)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		err := t.findOneAndDelete(sc, filter, dest, opt)
		if !opt.DisableAutoCloseSession {
			t.CloseSession(sc, err)
		}
		return err
	})
}

func (t *template) FindOneAndReplace(ctx context.Context, filter, replacement, dest any, opts ...option.FindOneAndReplace) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	opt := option.GetFindOneAndReplaceOptionByParams(opts)
	t.StartSession(ctx, opt.ForceRecreateSession)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		err := t.findOneAndReplace(sc, filter, replacement, dest, opt)
		if !opt.DisableAutoCloseSession {
			t.CloseSession(sc, err)
		}
		return err
	})
}

func (t *template) FindOneAndUpdate(ctx context.Context, filter, update, dest any, opts ...option.FindOneAndUpdate) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	opt := option.GetFindOneAndUpdateOptionByParams(opts)
	t.StartSession(ctx, opt.ForceRecreateSession)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		err := t.findOneAndUpdate(sc, filter, update, dest, opt)
		if !opt.DisableAutoCloseSession {
			t.CloseSession(sc, err)
		}
		return err
	})
}

func (t *template) Find(ctx context.Context, filter, dest any, opts ...option.Find) error {
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

func (t *template) EstimatedDocumentCount(ctx context.Context, ref any, opts ...option.EstimatedDocumentCount) (
	int64, error) {
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

func (t *template) Disconnect() {
	err := t.client.Disconnect(context.TODO())
	if err != nil {
		logger.Error("Error disconnect MongoDB:", err)
		return
	}
	logger.Info("Connection to MongoDB closed.")
}

func (t *template) StartSession(ctx context.Context, forceSession bool) {
	if t.session != nil && !forceSession {
		return
	} else if t.session != nil {
		t.CloseSession(ctx, nil)
	}
	session, _ := t.client.StartSession()
	_ = session.StartTransaction()
	t.session = session
}

func (t *template) CloseSession(ctx context.Context, err error) {
	if err != nil {
		if err = t.abortTransaction(ctx); err != nil {
			logger.Error("error abort transaction")
		} else {
			logger.Info("transaction aborted successfully!")
		}
		return
	}
	if err = t.commitTransaction(ctx); err != nil {
		logger.Error("error commit transaction")
	} else {
		logger.Info("transaction commit successfully!")
	}
	t.endSession(ctx)
	logger.Info("session closed successfully!")
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

func (t *template) insertMany(sc mongo.SessionContext, documents []any, opt option.InsertMany) error {
	if len(documents) == 0 {
		return ErrDocumentsIsEmpty
	}
	var errs []error
	for i, document := range documents {
		if document == nil {
			continue
		}
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

func (t *template) deleteOne(sc mongo.SessionContext, filter, ref any, opt option.Delete) (
	*mongo.DeleteResult, error) {
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

func (t *template) deleteMany(sc mongo.SessionContext, filter, ref any, opt option.Delete) (
	*mongo.DeleteResult, error) {
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

func (t *template) updateOne(sc mongo.SessionContext, filter, update, ref any, opt option.Update) (
	*mongo.UpdateResult, error) {
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

func (t *template) updateMany(sc mongo.SessionContext, filter, update, ref any, opt option.Update) (
	*mongo.UpdateResult, error) {
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

func (t *template) replaceOne(sc mongo.SessionContext, filter, update, ref any, opt option.Replace) (
	*mongo.UpdateResult, error) {
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
	return coll.FindOneAndDelete(sc, filter, &options.FindOneAndDeleteOptions{
		Collation:  option.ParseCollationMongoOptions(opt.Collation),
		Comment:    opt.Comment,
		MaxTime:    opt.MaxTime,
		Projection: opt.Projection,
		Sort:       opt.Sort,
		Hint:       opt.Hint,
		Let:        opt.Let,
	}).Decode(dest)
}

func (t *template) findOneAndReplace(sc mongo.SessionContext, filter, replacement, dest any, opt option.FindOneAndReplace) error {
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.FindOneAndReplace(sc, filter, replacement, &options.FindOneAndReplaceOptions{
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
}

func (t *template) findOneAndUpdate(sc mongo.SessionContext, filter, update, dest any, opt option.FindOneAndUpdate) error {
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.FindOneAndUpdate(sc, filter, update, &options.FindOneAndUpdateOptions{
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

func (t *template) commitTransaction(ctx context.Context) error {
	if t.session == nil {
		return nil
	}
	return t.session.CommitTransaction(ctx)
}

func (t *template) abortTransaction(ctx context.Context) error {
	if t.session == nil {
		return nil
	}
	return t.session.AbortTransaction(ctx)
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
