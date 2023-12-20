package mongo

import (
	"context"
	"github.com/GabrielHCataldo/go-logger/logger"
	"go-mongo/mongo/option"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testing"
	"time"
)

const MongoDBUrl = "MONGODB_URL"
const MongoDBTestId = "MONGODB_TEST_ID"

type testNewTemplate struct {
	name            string
	options         *options.ClientOptions
	durationTimeout time.Duration
	wantErr         bool
}

type testInsertOne struct {
	name                     string
	value                    any
	option                   option.InsertOne
	durationTimeout          time.Duration
	beforeStartSession       bool
	beforeCloseMongoClient   bool
	forceErrCloseMongoClient bool
	wantErr                  bool
}

type testInsertMany struct {
	name            string
	value           []any
	option          option.InsertMany
	durationTimeout time.Duration
	wantErr         bool
}

type testDelete struct {
	name            string
	filter          any
	ref             any
	option          option.Delete
	durationTimeout time.Duration
	wantErr         bool
}

type testUpdateOneById struct {
	name            string
	id              any
	update          any
	ref             any
	option          option.Update
	durationTimeout time.Duration
	wantErr         bool
}

type testUpdate struct {
	name            string
	filter          any
	update          any
	ref             any
	option          option.Update
	durationTimeout time.Duration
	wantErr         bool
}

type testReplace struct {
	name            string
	filter          any
	replacement     any
	ref             any
	option          option.Replace
	durationTimeout time.Duration
	wantErr         bool
}

type testAggregate struct {
	name            string
	pipeline        any
	dest            any
	option          option.Aggregate
	durationTimeout time.Duration
	wantErr         bool
}

type testCountDocuments struct {
	name            string
	filter          any
	ref             any
	option          option.Count
	durationTimeout time.Duration
	wantErr         bool
}

type testEstimatedDocumentCount struct {
	name            string
	ref             any
	option          option.EstimatedDocumentCount
	durationTimeout time.Duration
	wantErr         bool
}

var mongoTemplate Template

type testStruct struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" database:"test" collection:"test"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	BirthDate primitive.DateTime `json:"birthDate,omitempty" bson:"birthDate,omitempty"`
	Emails    []string           `json:"emails,omitempty" bson:"emails,omitempty"`
	Balance   float64            `json:"balance,omitempty" bson:"balance,omitempty"`
}

type testInvalidStruct struct {
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	BirthDate primitive.DateTime `json:"birthDate,omitempty" bson:"birthDate,omitempty"`
	Emails    []string           `json:"emails,omitempty" bson:"emails,omitempty"`
	Balance   float64            `json:"balance,omitempty" bson:"balance,omitempty"`
}

type testInvalidCollectionStruct struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty" database:"test"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	BirthDate primitive.DateTime `json:"birthDate,omitempty" bson:"birthDate,omitempty"`
	Emails    []string           `json:"emails,omitempty" bson:"emails,omitempty"`
	Balance   float64            `json:"balance,omitempty" bson:"balance,omitempty"`
}

type testEmptyStruct struct {
}

func TestMain(t *testing.M) {
	initMongoTemplate()
	clearCollection()
	t.Run()
	clearCollection()
	disconnectMongoTemplate()
}

func initMongoTemplate() {
	if mongoTemplate != nil {
		return
	}
	nMongoTemplate, err := NewTemplate(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		logger.Error("error new template:", err)
		return
	}
	mongoTemplate = nMongoTemplate
}

func initDocument() {
	initMongoTemplate()
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	test := initTestStruct()
	err := mongoTemplate.InsertOne(ctx, test)
	if err != nil {
		logger.Error("error init document:", err)
		return
	}
	err = os.Setenv(MongoDBTestId, test.Id.Hex())
	if err != nil {
		logger.Error("err set MongoDBTestId env:", err)
	}
}

func initTestStruct() *testStruct {
	return &testStruct{
		Name:      "Test Full Name",
		BirthDate: primitive.NewDateTimeFromTime(time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local)),
		Emails:    []string{"testemail@gmail.com", "fullname@gmail.com", "test@gmail.com"},
		Balance:   100.32,
	}
}

func initTestInvalidStruct() *testInvalidStruct {
	return &testInvalidStruct{
		Name:      "Test Full Name",
		BirthDate: primitive.NewDateTimeFromTime(time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local)),
		Emails:    []string{"testemail@gmail.com", "fullname@gmail.com", "test@gmail.com"},
		Balance:   100.32,
	}
}

func initTestInvalidCollectionStruct() *testInvalidCollectionStruct {
	return &testInvalidCollectionStruct{
		Name:      "Test Full Name",
		BirthDate: primitive.NewDateTimeFromTime(time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local)),
		Emails:    []string{"testemail@gmail.com", "fullname@gmail.com", "test@gmail.com"},
		Balance:   100.32,
	}
}

func initTestEmptyStruct() *testEmptyStruct {
	return &testEmptyStruct{}
}

func initTestString() *string {
	s := "test value string"
	return &s
}

func initListTestNewTemplate() []testNewTemplate {
	return []testNewTemplate{
		{
			name:            "success",
			options:         options.Client().ApplyURI(os.Getenv(MongoDBUrl)),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed client",
			options:         options.Client().ApplyURI("https://google.com/"),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed timeout",
			options:         options.Client().ApplyURI(os.Getenv(MongoDBUrl)),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
	}
}

func initListTestInsertOne() []testInsertOne {
	return []testInsertOne{
		{
			name:            "success",
			value:           initTestStruct(),
			option:          initOptionInsertOne(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:                     "success with error commit transaction",
			value:                    initTestStruct(),
			option:                   initOptionInsertOne().SetDisableAutoCloseTransaction(true),
			forceErrCloseMongoClient: true,
			durationTimeout:          5 * time.Second,
		},
		{
			name:            "failed type value",
			value:           "string value",
			option:          initOptionInsertOne(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed empty value",
			value:           initTestEmptyStruct(),
			option:          initOptionInsertOne(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed invalid database struct config",
			value:           initTestInvalidStruct(),
			option:          initOptionInsertOne(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed invalid collection struct config",
			value:           initTestInvalidCollectionStruct(),
			option:          initOptionInsertOne(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed non struct value",
			value:           initTestString(),
			option:          initOptionInsertOne(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:                     "failed timeout",
			value:                    initTestStruct(),
			option:                   initOptionInsertOne().SetDisableAutoCloseTransaction(true),
			durationTimeout:          1 * time.Millisecond,
			forceErrCloseMongoClient: true,
			wantErr:                  true,
		},
		{
			name:                   "failed start session",
			value:                  initTestStruct(),
			option:                 initOptionInsertOne().SetForceRecreateSession(false),
			durationTimeout:        5 * time.Second,
			beforeCloseMongoClient: true,
			beforeStartSession:     true,
			wantErr:                true,
		},
	}
}

func initListTestInsertMany() []testInsertMany {
	return []testInsertMany{
		{
			name:            "success",
			value:           []any{initTestStruct(), initTestStruct(), nil},
			option:          initOptionInsertMany(),
			durationTimeout: 5 * time.Second,
		},
		{
			name: "success one",
			value: []any{
				initTestStruct(),
				initTestInvalidStruct(),
				initTestInvalidCollectionStruct(),
				initTestEmptyStruct(),
				initTestString(),
				"test string normal",
			},
			option: initOptionInsertMany().
				SetDisableAutoRollback(true).
				SetDisableAutoCloseTransaction(true),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed empty value",
			value:           []any{},
			option:          initOptionInsertMany(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestDelete() []testDelete {
	return []testDelete{
		{
			name:            "success",
			filter:          bson.M{"_id": os.Getenv(MongoDBTestId)},
			ref:             testStruct{},
			option:          initOptionDelete(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed struct ref",
			filter:          bson.M{},
			ref:             initTestInvalidStruct(),
			option:          initOptionDelete(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type ref",
			filter:          bson.M{},
			ref:             initTestString(),
			option:          initOptionDelete().SetDisableAutoCloseTransaction(true),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestUpdateOneById() []testUpdateOneById {
	objectId, _ := primitive.ObjectIDFromHex(os.Getenv(MongoDBTestId))
	return []testUpdateOneById{
		{
			name:            "success",
			id:              objectId,
			update:          bson.M{"$set": bson.M{"name": "Updated Test Name"}},
			ref:             testStruct{},
			option:          initOptionUpdate().SetUpsert(false).SetArrayFilters(nil),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed struct ref",
			id:              "id string",
			update:          nil,
			ref:             initTestInvalidStruct(),
			option:          initOptionUpdate(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type ref",
			id:              "id string",
			update:          nil,
			ref:             initTestString(),
			option:          initOptionUpdate(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed update param",
			id:              objectId,
			update:          nil,
			ref:             testStruct{},
			option:          initOptionUpdate().SetDisableAutoCloseTransaction(true),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestUpdate() []testUpdate {
	return []testUpdate{
		{
			name:            "success",
			filter:          bson.M{"_id": bson.M{"$exists": true}},
			update:          bson.M{"$set": bson.M{"name": "Updated Test Name"}},
			ref:             testStruct{},
			option:          initOptionUpdate().SetUpsert(false).SetArrayFilters(nil),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed struct ref",
			update:          nil,
			ref:             initTestInvalidStruct(),
			option:          initOptionUpdate(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type ref",
			update:          nil,
			ref:             initTestString(),
			option:          initOptionUpdate(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed update param",
			filter:          bson.M{"_id": bson.M{"$exists": true}},
			update:          nil,
			ref:             testStruct{},
			option:          initOptionUpdate().SetDisableAutoCloseTransaction(true),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestReplace() []testReplace {
	objectId, _ := primitive.ObjectIDFromHex(os.Getenv(MongoDBTestId))
	return []testReplace{
		{
			name:            "success",
			filter:          bson.M{"_id": objectId},
			replacement:     *initTestStruct(),
			ref:             testStruct{},
			option:          initOptionReplace(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed struct ref",
			replacement:     nil,
			ref:             initTestInvalidStruct(),
			option:          initOptionReplace(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type ref",
			replacement:     nil,
			ref:             initTestString(),
			option:          initOptionReplace(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed update param",
			filter:          bson.M{"_id": objectId},
			replacement:     nil,
			ref:             testStruct{},
			option:          initOptionReplace().SetDisableAutoCloseTransaction(true),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestAggregate() []testAggregate {
	return []testAggregate{
		{
			name:            "success",
			pipeline:        Pipeline{bson.D{{"$match", bson.D{{"_id", bson.M{"$exists": true}}}}}},
			dest:            &[]testStruct{},
			option:          initOptionAggregate(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:     "failed struct dest",
			pipeline: nil,
			dest:     &testInvalidStruct{},
			option: initOptionAggregate().
				SetAllowDiskUse(true).
				SetCollation(&option.Collation{}).
				SetBatchSize(10),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type dest",
			pipeline:        nil,
			dest:            initTestString(),
			option:          initOptionAggregate(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestCountDocuments() []testCountDocuments {
	return []testCountDocuments{
		{
			name:            "success",
			filter:          bson.D{{"_id", bson.M{"$exists": true}}},
			ref:             &testStruct{},
			option:          initOptionCount(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:   "failed filter",
			filter: "filter string err",
			option: initOptionCount().
				SetCollation(&option.Collation{}).
				SetLimit(10).
				SetSkip(10),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:   "failed struct ref",
			filter: nil,
			ref:    &testInvalidStruct{},
			option: initOptionCount().
				SetCollation(&option.Collation{}).
				SetLimit(10).
				SetSkip(10),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type ref",
			filter:          nil,
			ref:             initTestString(),
			option:          initOptionCount(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestEstimatedDocumentCount() []testEstimatedDocumentCount {
	return []testEstimatedDocumentCount{
		{
			name:            "success",
			ref:             &testStruct{},
			option:          initOptionEstimatedDocumentCount(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed filter",
			option:          initOptionEstimatedDocumentCount(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed struct ref",
			ref:             &testInvalidStruct{},
			option:          initOptionEstimatedDocumentCount(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type ref",
			ref:             initTestString(),
			option:          initOptionEstimatedDocumentCount(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initOptionInsertOne() option.InsertOne {
	return option.NewInsertOne().
		SetBypassDocumentValidation(true).
		SetForceRecreateSession(true).
		SetComment("comment insert golang unit test").
		SetDisableAutoCloseTransaction(false)
}

func initOptionInsertMany() option.InsertMany {
	return option.NewInsertMany().
		SetBypassDocumentValidation(true).
		SetForceRecreateSession(true).
		SetComment("comment insert golang unit test").
		SetDisableAutoCloseTransaction(false).
		SetDisableAutoRollback(false)
}

func initOptionDelete() option.Delete {
	return option.NewDelete().
		SetDisableAutoCloseTransaction(false).
		SetComment("comment delete golang unit test").
		SetCollation(&option.Collation{}).
		SetHint(bson.M{}).
		SetLet(bson.M{})
}

func initOptionUpdate() option.Update {
	return option.NewUpdate().
		SetDisableAutoCloseTransaction(false).
		SetComment("comment update golang unit test").
		SetCollation(&option.Collation{}).
		SetHint(bson.M{}).
		SetLet(bson.M{}).
		SetBypassDocumentValidation(true).
		SetArrayFilters(&option.ArrayFilters{}).
		SetUpsert(true)
}

func initOptionReplace() option.Replace {
	return option.NewReplace().
		SetDisableAutoCloseTransaction(false).
		SetComment("comment replace golang unit test").
		SetCollation(&option.Collation{}).
		SetHint(bson.M{}).
		SetLet(bson.M{}).
		SetBypassDocumentValidation(true)
}

func initOptionAggregate() option.Aggregate {
	return option.NewAggregate().
		SetAllowDiskUse(false).
		SetBatchSize(0).
		SetBypassDocumentValidation(true).
		SetCollation(nil).
		SetMaxTime(5 * time.Second).
		SetMaxAwaitTime(2 * time.Second).
		SetComment("comment aggregate golang unit test").
		SetHint(bson.M{}).
		SetLet(bson.M{}).
		SetCustom(bson.M{})
}

func initOptionCount() option.Count {
	return option.NewCount().
		SetCollation(nil).
		SetComment("comment count golang unit test").
		SetHint(bson.M{}).
		SetMaxTime(5 * time.Second)
}

func initOptionEstimatedDocumentCount() option.EstimatedDocumentCount {
	return option.NewEstimatedDocumentCount().
		SetComment("comment estimated document count golang unit test").
		SetMaxTime(5 * time.Second)
}

func disconnectMongoTemplate() {
	if mongoTemplate == nil {
		return
	}
	mongoTemplate.Disconnect()
	mongoTemplate = nil
}

func clearCollection() {
	if mongoTemplate == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	_, err := mongoTemplate.DeleteMany(ctx, bson.M{"_id": bson.M{"$exists": true}}, testStruct{})
	if err != nil {
		logger.Error("error clean collection:", err)
	} else {
		logger.Info("collection cleaned!")
	}
}
