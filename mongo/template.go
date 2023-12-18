package mongo

import (
	"context"
	"encoding/json"
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

type template struct {
	client         *mongo.Client
	session        mongo.Session
	sessionContext mongo.SessionContext
}

type Template interface {
	InsertOne(ctx context.Context, document any, opts ...option.InsertOne) error
	InsertMany(ctx context.Context, documents []any, opts ...option.InsertMany) error
	DeleteOne(ctx context.Context, filter any, ref struct{}, opts ...option.Delete) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter any, ref struct{}, opts ...option.Delete) (*mongo.DeleteResult, error)
	UpdateOneById(ctx context.Context, id, update any, ref struct{}, opts ...option.Update) (*mongo.UpdateResult, error)
	UpdateOne(ctx context.Context, filter, update any, ref struct{}, opts ...option.Update) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter, update any, ref struct{}, opts ...option.Update) (*mongo.UpdateResult, error)
	ReplaceOne(ctx context.Context, filter, replacement any, ref struct{}, opts ...option.Replace) (*mongo.UpdateResult, error)
	Aggregate(ctx context.Context, pipeline, dest any, opts ...option.Aggregate) error
	CountDocuments(ctx context.Context, filter any, ref struct{}, opts ...option.Count) (int64, error)
	EstimatedDocumentCount(ctx context.Context, ref struct{}, opts ...option.EstimatedDocumentCount) (int64, error)
	Distinct(ctx context.Context, fieldName string, filter, dest any, opts ...option.Distinct) error
	FindOne(ctx context.Context, filter, dest any, opts ...option.FindOne) error
	FindOneAndDelete(ctx context.Context, filter, dest any, opts ...option.FindOneAndDelete) error
	FindOneAndReplace(ctx context.Context, filter, replacement, dest any, opts ...option.FindOneAndReplace) error
	FindOneAndUpdate(ctx context.Context, filter, update, dest any, opts ...option.FindOneAndUpdate) error
	Find(ctx context.Context, filter, dest any, opts ...option.Find) error
	FindPageable(ctx context.Context, filter any, input PageInput, opts ...option.FindPageable) (*PageOutput, error)
	CloseTransaction(err error)
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
	session, err := conn.StartSession()
	if err != nil {
		return nil, err
	}
	return template{
		client:  conn,
		session: session,
	}, nil
}

func (t template) InsertOne(ctx context.Context, document any, opts ...option.InsertOne) error {
	err := t.session.StartTransaction()
	if err != nil {
		return err
	}
	opt := option.GetInsertOneOptionByParams(opts)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		t.sessionContext = sc
		if !opt.DisableAutoCloseTransaction {
			defer t.CloseTransaction(err)
		}
		err = t.insertOne(sc, document, opt)
		return err
	})
}

func (t template) InsertMany(ctx context.Context, documents []any, opts ...option.InsertMany) error {
	err := t.session.StartTransaction()
	if err != nil {
		return err
	}
	opt := option.GetInsertManyOptionByParams(opts)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		t.sessionContext = sc
		if !opt.DisableAutoCloseTransaction {
			errClose := err
			if !opt.DisableAutoRollback {
				errClose = nil
			}
			defer t.CloseTransaction(errClose)
		}
		errs := t.insertMany(sc, documents, opt)
		if len(errs) != 0 {
			var b strings.Builder
			for i, errResult := range errs {
				if i != 0 {
					b.WriteString(", ")
				}
				b.WriteString(errResult.Error())
			}
			err = errors.New(b.String())
		}
		return err
	})
}

func (t template) DeleteOne(ctx context.Context, filter any, doc struct{}, opts ...option.Delete) (
	*mongo.DeleteResult, error) {
	err := t.session.StartTransaction()
	if err != nil {
		return nil, err
	}
	var result *mongo.DeleteResult
	opt := option.GetDeleteOptionByParams(opts)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		t.sessionContext = sc
		if !opt.DisableAutoCloseTransaction {
			defer t.CloseTransaction(err)
		}
		result, err = t.deleteOne(sc, filter, doc, opt)
		return err
	})
	return result, nil
}

func (t template) DeleteMany(ctx context.Context, filter any, doc struct{}, opts ...option.Delete) (
	*mongo.DeleteResult, error) {
	err := t.session.StartTransaction()
	if err != nil {
		return nil, err
	}
	var result *mongo.DeleteResult
	opt := option.GetDeleteOptionByParams(opts)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		t.sessionContext = sc
		if !opt.DisableAutoCloseTransaction {
			defer t.CloseTransaction(err)
		}
		result, err = t.deleteMany(sc, filter, doc, opt)
		return err
	})
	return result, nil
}

func (t template) UpdateOneById(ctx context.Context, id any, update any, ref struct{}, opts ...option.Update) (
	*mongo.UpdateResult, error) {
	err := t.session.StartTransaction()
	if err != nil {
		return nil, err
	}
	var result *mongo.UpdateResult
	opt := option.GetUpdateOptionByParams(opts)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		t.sessionContext = sc
		if !opt.DisableAutoCloseTransaction {
			defer t.CloseTransaction(err)
		}
		result, err = t.updateOne(sc, bson.M{"_id": id}, update, ref, opt)
		return err
	})
	return result, nil
}

func (t template) UpdateOne(ctx context.Context, filter any, update any, ref struct{}, opts ...option.Update) (
	*mongo.UpdateResult, error) {
	err := t.session.StartTransaction()
	if err != nil {
		return nil, err
	}
	var result *mongo.UpdateResult
	opt := option.GetUpdateOptionByParams(opts)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		t.sessionContext = sc
		if !opt.DisableAutoCloseTransaction {
			defer t.CloseTransaction(err)
		}
		result, err = t.updateOne(sc, filter, update, ref, opt)
		return err
	})
	return result, nil
}

func (t template) UpdateMany(ctx context.Context, filter any, update any, ref struct{}, opts ...option.Update) (
	*mongo.UpdateResult, error) {
	err := t.session.StartTransaction()
	if err != nil {
		return nil, err
	}
	var result *mongo.UpdateResult
	opt := option.GetUpdateOptionByParams(opts)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		t.sessionContext = sc
		if !opt.DisableAutoCloseTransaction {
			defer t.CloseTransaction(err)
		}
		result, err = t.updateMany(sc, filter, update, ref, opt)
		return err
	})
	return result, nil
}

func (t template) ReplaceOne(ctx context.Context, filter any, update any, ref struct{}, opts ...option.Replace) (
	*mongo.UpdateResult, error) {
	err := t.session.StartTransaction()
	if err != nil {
		return nil, err
	}
	var result *mongo.UpdateResult
	opt := option.GetReplaceOptionByParams(opts)
	err = mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		t.sessionContext = sc
		if !opt.DisableAutoCloseTransaction {
			defer t.CloseTransaction(err)
		}
		result, err = t.replaceOne(sc, filter, update, ref, opt)
		return err
	})
	return result, nil
}

func (t template) FindOne(ctx context.Context, filter, dest any, opts ...option.FindOne) error {
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
		AllowPartialResults: &opt.AllowPartialResults,
		Collation:           option.ParseCollationMongoOptions(opt.Collation),
		Comment:             &opt.Comment,
		Hint:                opt.Hint,
		Max:                 opt.Max,
		MaxTime:             &opt.MaxTime,
		Min:                 opt.Min,
		Projection:          opt.Projection,
		ReturnKey:           &opt.ReturnKey,
		ShowRecordID:        &opt.ShowRecordID,
		Skip:                &opt.Skip,
		Sort:                opt.Sort,
	}).Decode(dest)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	return nil
}

func (t template) FindOneAndDelete(ctx context.Context, filter, dest any, opts ...option.FindOneAndDelete) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	err := t.session.StartTransaction()
	if err != nil {
		return err
	}
	opt := option.GetFindOneAndDeleteOptionByParams(opts)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		t.sessionContext = sc
		if !opt.DisableAutoCloseTransaction {
			defer t.CloseTransaction(err)
		}
		return t.findOneAndDelete(sc, filter, dest, opt)
	})
}

func (t template) FindOneAndReplace(ctx context.Context, filter, replacement, dest any, opts ...option.FindOneAndReplace) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	err := t.session.StartTransaction()
	if err != nil {
		return err
	}
	opt := option.GetFindOneAndReplaceOptionByParams(opts)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		t.sessionContext = sc
		if !opt.DisableAutoCloseTransaction {
			defer t.CloseTransaction(err)
		}
		return t.findOneAndReplace(sc, filter, replacement, dest, opt)
	})
}

func (t template) FindOneAndUpdate(ctx context.Context, filter, update, dest any, opts ...option.FindOneAndUpdate) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	} else if util.IsNotStruct(dest) {
		return ErrDestIsNotStruct
	}
	err := t.session.StartTransaction()
	if err != nil {
		return err
	}
	opt := option.GetFindOneAndUpdateOptionByParams(opts)
	return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
		t.sessionContext = sc
		if !opt.DisableAutoCloseTransaction {
			defer t.CloseTransaction(err)
		}
		return t.findOneAndUpdate(sc, filter, update, dest, opt)
	})
}

func (t template) Find(ctx context.Context, filter, dest any, opts ...option.Find) error {
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
		AllowDiskUse:        &opt.AllowDiskUse,
		AllowPartialResults: &opt.AllowPartialResults,
		BatchSize:           &opt.BatchSize,
		Collation:           option.ParseCollationMongoOptions(opt.Collation),
		Comment:             &opt.Comment,
		CursorType:          option.ParseCursorType(opt.CursorType),
		Hint:                opt.Hint,
		Limit:               &opt.Limit,
		Max:                 opt.Max,
		MaxAwaitTime:        &opt.MaxAwaitTime,
		MaxTime:             &opt.MaxTime,
		Min:                 opt.Min,
		NoCursorTimeout:     &opt.NoCursorTimeout,
		Projection:          opt.Projection,
		ReturnKey:           &opt.ReturnKey,
		ShowRecordID:        &opt.ShowRecordID,
		Skip:                &opt.Skip,
		Sort:                opt.Sort,
		Let:                 opt.Let,
	})
	if err != nil {
		return err
	}
	return cursor.All(ctx, dest)
}

func (t template) FindPageable(ctx context.Context, filter any, input PageInput, opts ...option.FindPageable) (
	*PageOutput, error) {
	databaseName, collectionName, err := getMongoInfosByAny(input.DocRef)
	if err != nil {
		return nil, err
	}
	opt := option.GetFindPageableOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName)
	cursor, err := collection.Find(ctx, filter, &options.FindOptions{
		AllowDiskUse:        &opt.AllowDiskUse,
		AllowPartialResults: &opt.AllowPartialResults,
		BatchSize:           &opt.BatchSize,
		Collation:           option.ParseCollationMongoOptions(opt.Collation),
		Comment:             &opt.Comment,
		CursorType:          option.ParseCursorType(opt.CursorType),
		Hint:                opt.Hint,
		Limit:               &input.PageSize,
		Max:                 opt.Max,
		MaxAwaitTime:        &opt.MaxAwaitTime,
		MaxTime:             &opt.MaxTime,
		Min:                 opt.Min,
		NoCursorTimeout:     &opt.NoCursorTimeout,
		Projection:          opt.Projection,
		ReturnKey:           &opt.ReturnKey,
		ShowRecordID:        &opt.ShowRecordID,
		Skip:                &input.Page,
		Sort:                input.Sort,
		Let:                 opt.Let,
	})
	dest := input.DocRef
	err = cursor.All(ctx, dest)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	countTotal, err := collection.CountDocuments(ctx, filter)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	return NewPageOutput(input, dest, countTotal), nil
}

func (t template) Aggregate(ctx context.Context, pipeline any, dest any, opts ...option.Aggregate) error {
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
		AllowDiskUse:             &opt.AllowDiskUse,
		BatchSize:                &opt.BatchSize,
		BypassDocumentValidation: &opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		MaxTime:                  &opt.MaxTime,
		MaxAwaitTime:             &opt.MaxAwaitTime,
		Comment:                  &opt.Comment,
		Hint:                     opt.Hint,
		Let:                      opt.Let,
		Custom:                   opt.Custom,
	})
	if err != nil {
		return err
	}
	err = cursor.All(ctx, dest)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	return nil
}

func (t template) CountDocuments(ctx context.Context, filter any, ref struct{}, opts ...option.Count) (int64, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return 0, err
	}
	opt := option.GetCountOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName)
	return collection.CountDocuments(ctx, filter, &options.CountOptions{
		Collation: option.ParseCollationMongoOptions(opt.Collation),
		Comment:   &opt.Comment,
		Hint:      &opt.Hint,
		Limit:     &opt.Limit,
		MaxTime:   &opt.MaxTime,
		Skip:      &opt.Skip,
	})
}

func (t template) EstimatedDocumentCount(ctx context.Context, ref struct{}, opts ...option.EstimatedDocumentCount) (
	int64, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return 0, err
	}
	opt := option.GetEstimatedDocumentCountOptionByParams(opts)
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName)
	return collection.EstimatedDocumentCount(ctx, &options.EstimatedDocumentCountOptions{
		Comment: &opt.Comment,
		MaxTime: &opt.MaxTime,
	})
}

func (t template) Distinct(ctx context.Context, fieldName string, filter, dest any, opts ...option.Distinct) error {
	if util.IsNotPointer(dest) {
		return ErrDestIsNotPointer
	}
	opt := option.GetDistinctOptionByParams(opts)
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	collection := database.Collection(collectionName)
	result, err := collection.Distinct(ctx, fieldName, filter, &options.DistinctOptions{
		Collation: option.ParseCollationMongoOptions(opt.Collation),
		Comment:   &opt.Comment,
		MaxTime:   &opt.MaxTime,
	})
	if err != nil {
		return err
	} else if result == nil || len(result) == 0 {
		return nil
	}
	b, err := json.Marshal(result)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dest)
}

func (t template) CloseTransaction(err error) {
	ctx := t.sessionContext
	t.sessionContext = nil
	if err != nil {
		if err = t.session.AbortTransaction(ctx); err != nil {
			logger.Error("error abort transaction")
		}
		logger.Info("transaction aborted successfully!")
		return
	}
	if err = t.session.CommitTransaction(ctx); err != nil {
		logger.Error("error commit transaction")
		return
	}
	logger.Info("transaction commit successfully!")
}

func (t template) Disconnect() {
	//TODO implement me
	panic("implement me")
}

func (t template) insertOne(sc mongo.SessionContext, document any, opt option.InsertOne) error {
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
		BypassDocumentValidation: &opt.BypassDocumentValidation,
		Comment:                  opt.Comment,
	})
	if err != nil {
		return err
	}
	util.SetInsertedIdOnDocument(result.InsertedID, document)
	return nil
}

func (t template) insertMany(sc mongo.SessionContext, documents []any, opt option.InsertMany) []error {
	if len(documents) == 0 {
		return []error{ErrDocumentsIsEmpty}
	}
	var errs []error
	for i, document := range documents {
		indexStr := strconv.Itoa(i)
		if util.IsNotPointer(document) {
			errs = append(errs, errors.New(ErrDocumentIsNotPointer.Error()+"(index: "+indexStr+")"))
		} else if util.IsNotStruct(document) {
			errs = append(errs, errors.New(ErrDocumentIsNotStruct.Error()+"(index: "+indexStr+")"))
		} else if util.IsZero(document) {
			errs = append(errs, errors.New(ErrDocumentIsEmpty.Error()+"(index: "+indexStr+")"))
		} else {
			err := t.insertOne(sc, documents, option.InsertOne{
				BypassDocumentValidation: opt.BypassDocumentValidation,
				Comment:                  opt.Comment,
			})
			errs = append(errs, err)
		}
	}
	return errs
}

func (t template) deleteOne(sc mongo.SessionContext, filter any, ref struct{}, opt option.Delete) (
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

func (t template) deleteMany(sc mongo.SessionContext, filter any, ref struct{}, opt option.Delete) (
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

func (t template) updateOne(sc mongo.SessionContext, filter, update any, ref struct{}, opt option.Update) (
	*mongo.UpdateResult, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.UpdateOne(sc, filter, update, &options.UpdateOptions{
		ArrayFilters:             option.ParseArrayFiltersMongoOptions(opt.ArrayFilters),
		BypassDocumentValidation: &opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		Hint:                     opt.Hint,
		Upsert:                   &opt.Upsert,
		Let:                      opt.Let,
	})
}

func (t template) updateMany(sc mongo.SessionContext, filter, update any, ref struct{}, opt option.Update) (
	*mongo.UpdateResult, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.UpdateMany(sc, filter, update, &options.UpdateOptions{
		ArrayFilters:             option.ParseArrayFiltersMongoOptions(opt.ArrayFilters),
		BypassDocumentValidation: &opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		Hint:                     opt.Hint,
		Upsert:                   &opt.Upsert,
		Let:                      opt.Let,
	})
}

func (t template) replaceOne(sc mongo.SessionContext, filter, update any, ref struct{}, opt option.Replace) (
	*mongo.UpdateResult, error) {
	databaseName, collectionName, err := getMongoInfosByAny(ref)
	if err != nil {
		return nil, err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.ReplaceOne(sc, filter, update, &options.ReplaceOptions{
		BypassDocumentValidation: &opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		Hint:                     opt.Hint,
		Upsert:                   &opt.Upsert,
		Let:                      opt.Let,
	})
}

func (t template) findOneAndDelete(sc mongo.SessionContext, filter, dest any, opt option.FindOneAndDelete) error {
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.FindOneAndDelete(sc, filter, &options.FindOneAndDeleteOptions{
		Collation:  option.ParseCollationMongoOptions(opt.Collation),
		Comment:    opt.Comment,
		MaxTime:    &opt.MaxTime,
		Projection: opt.Projection,
		Sort:       opt.Sort,
		Hint:       opt.Hint,
		Let:        opt.Let,
	}).Decode(dest)
}

func (t template) findOneAndReplace(sc mongo.SessionContext, filter, replacement, dest any, opt option.FindOneAndReplace) error {
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.FindOneAndReplace(sc, filter, replacement, &options.FindOneAndReplaceOptions{
		BypassDocumentValidation: &opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		MaxTime:                  &opt.MaxTime,
		Projection:               opt.Projection,
		ReturnDocument:           option.ParseReturnDocument(opt.ReturnDocument),
		Sort:                     opt.Sort,
		Upsert:                   &opt.Upsert,
		Hint:                     opt.Hint,
		Let:                      opt.Let,
	}).Decode(dest)
}

func (t template) findOneAndUpdate(sc mongo.SessionContext, filter, update, dest any, opt option.FindOneAndUpdate) error {
	databaseName, collectionName, err := getMongoInfosByAny(dest)
	if err != nil {
		return err
	}
	database := t.client.Database(databaseName)
	coll := database.Collection(collectionName)
	return coll.FindOneAndUpdate(sc, filter, update, &options.FindOneAndUpdateOptions{
		ArrayFilters:             option.ParseArrayFiltersMongoOptions(opt.ArrayFilters),
		BypassDocumentValidation: &opt.BypassDocumentValidation,
		Collation:                option.ParseCollationMongoOptions(opt.Collation),
		Comment:                  opt.Comment,
		MaxTime:                  &opt.MaxTime,
		Projection:               opt.Projection,
		ReturnDocument:           option.ParseReturnDocument(opt.ReturnDocument),
		Sort:                     opt.Sort,
		Upsert:                   &opt.Upsert,
		Hint:                     opt.Hint,
		Let:                      opt.Let,
	}).Decode(dest)
}

func getMongoInfosByAny(a any) (databaseName string, collectionName string, err error) {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		databaseName = util.GetDatabaseNameBySlice(a)
		collectionName = util.GetCollectionNameBySlice(a)
		break
	default:
		databaseName = util.GetDatabaseNameByStruct(a)
		collectionName = util.GetCollectionNameByStruct(a)
	}
	if len(databaseName) == 0 {
		return "", "", ErrDatabaseNotConfigured
	}
	if len(collectionName) == 0 {
		return "", "", ErrCollectionNotConfigured
	}
	return collectionName, databaseName, nil
}
