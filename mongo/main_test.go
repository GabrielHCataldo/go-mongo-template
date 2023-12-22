package mongo

import (
	"context"
	"github.com/GabrielHCataldo/go-logger/logger"
	"go-mongo/mongo/option"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"os"
	"testing"
	"time"
)

const MongoDBUrl = "MONGODB_URL"
const MongoDBTestId = "MONGODB_TEST_ID"
const MongoDBIndexName = "MONGODB_INDEX_NAME"

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

type testFindOneById struct {
	name            string
	id              any
	dest            any
	option          option.FindOneById
	durationTimeout time.Duration
	wantErr         bool
}

type testFindOne struct {
	name            string
	filter          any
	dest            any
	option          option.FindOne
	durationTimeout time.Duration
	wantErr         bool
}

type testFindOneAndDelete struct {
	name            string
	filter          any
	dest            any
	option          option.FindOneAndDelete
	durationTimeout time.Duration
	wantErr         bool
}

type testFindOneAndReplace struct {
	name            string
	filter          any
	replacement     any
	dest            any
	option          option.FindOneAndReplace
	durationTimeout time.Duration
	wantErr         bool
}

type testFindOneAndUpdate struct {
	name            string
	filter          any
	update          any
	dest            any
	option          option.FindOneAndUpdate
	durationTimeout time.Duration
	wantErr         bool
}

type testFind struct {
	name            string
	filter          any
	dest            any
	option          option.Find
	durationTimeout time.Duration
	wantErr         bool
}

type testFindPageable struct {
	name            string
	filter          any
	pageInput       PageInput
	option          option.FindPageable
	durationTimeout time.Duration
	wantErr         bool
}

type testExists struct {
	name            string
	filter          any
	ref             any
	option          option.Exists
	durationTimeout time.Duration
	wantErr         bool
}

type testExistsById struct {
	name            string
	id              any
	ref             any
	option          option.Exists
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

type testDistinct struct {
	name            string
	fieldName       string
	filter          any
	dest            any
	ref             any
	option          option.Distinct
	durationTimeout time.Duration
	wantErr         bool
}

type testWatch struct {
	name            string
	pipeline        any
	option          option.Watch
	durationTimeout time.Duration
	wantErr         bool
}

type testWatchHandler struct {
	name            string
	pipeline        any
	handler         HandlerWatch
	option          option.WatchHandler
	durationTimeout time.Duration
	wantErr         bool
}

type testDrop struct {
	name            string
	ref             any
	durationTimeout time.Duration
	wantErr         bool
}

type testCreateOneIndex struct {
	name            string
	input           IndexInput
	durationTimeout time.Duration
	wantErr         bool
}

type testCreateManyIndex struct {
	name            string
	inputs          []IndexInput
	durationTimeout time.Duration
	wantErr         bool
}

type testDropIndex struct {
	name            string
	nameIndex       string
	ref             any
	option          option.DropIndex
	durationTimeout time.Duration
	wantErr         bool
}

type testListIndexes struct {
	name            string
	ref             any
	option          option.ListIndexes
	durationTimeout time.Duration
	wantErr         bool
}

var mongoTemplate Template

type testStruct struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" database:"test" collection:"test"`
	Random    int                `json:"random,omitempty" bson:"random,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	BirthDate primitive.DateTime `json:"birthDate,omitempty" bson:"birthDate,omitempty"`
	Emails    []string           `json:"emails,omitempty" bson:"emails,omitempty"`
	Balance   float64            `json:"balance,omitempty" bson:"balance,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
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

func initIndex() {
	initDocument()
	clearIndexes()
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	result, err := mongoTemplate.CreateOneIndex(ctx, initIndexInput())
	if err != nil {
		logger.Error("error init index:", err)
		return
	}
	err = os.Setenv(MongoDBIndexName, result)
	if err != nil {
		logger.Error("err set MongoDBIndexName env:", err)
	}
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

func clearIndexes() {
	if mongoTemplate == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	err := mongoTemplate.DropAllIndexes(ctx, testStruct{})
	if err != nil {
		logger.Error("error clean indexes:", err)
	} else {
		logger.Info("collection indexes cleaned!")
	}
}

func initTestStruct() *testStruct {
	return &testStruct{
		Random:    rand.Int(),
		Name:      "Test Full Name",
		BirthDate: primitive.NewDateTimeFromTime(time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local)),
		Emails:    []string{"testemail@gmail.com", "fullname@gmail.com", "test@gmail.com"},
		Balance:   100.32,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
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

func initIndexInput() IndexInput {
	return IndexInput{
		Keys: bson.D{{"random", 1}},
		Options: initOptionIndex().
			SetName("test index unique").
			SetUnique(true).
			SetSparse(true),
		Ref: testStruct{},
	}
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
			option:                   initOptionInsertOne().SetDisableAutoCloseSession(true),
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
			option:                   initOptionInsertOne().SetDisableAutoCloseSession(true),
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
				SetDisableAutoCloseSession(true),
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
			name:            "failed",
			filter:          bson.M{"_id": os.Getenv(MongoDBTestId)},
			ref:             testStruct{},
			option:          initOptionDelete(),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
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
			option:          initOptionDelete().SetDisableAutoCloseSession(true),
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
			option:          initOptionUpdate().SetDisableAutoCloseSession(true),
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
			option:          initOptionUpdate().SetDisableAutoCloseSession(true),
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
			option:          initOptionReplace().SetDisableAutoCloseSession(true),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestFindOneById() []testFindOneById {
	objectId, _ := primitive.ObjectIDFromHex(os.Getenv(MongoDBTestId))
	return []testFindOneById{
		{
			name:            "success",
			id:              objectId,
			dest:            &testStruct{},
			option:          initOptionFindOneById(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "success no content",
			id:              primitive.NewObjectID(),
			dest:            &testStruct{},
			option:          initOptionFindOneById(),
			durationTimeout: 5 * time.Second,
		},
		{
			name: "failed",
			id:   objectId,
			dest: &testStruct{},
			option: initOptionFindOneById().
				SetCollation(&option.Collation{}),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:            "failed struct dest",
			id:              nil,
			dest:            &testInvalidStruct{},
			option:          initOptionFindOneById(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type dest",
			id:              nil,
			dest:            initTestString(),
			option:          initOptionFindOneById(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed dest non pointer",
			id:              nil,
			dest:            *initTestString(),
			option:          initOptionFindOneById(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestFindOne() []testFindOne {
	objectId, _ := primitive.ObjectIDFromHex(os.Getenv(MongoDBTestId))
	return []testFindOne{
		{
			name:            "success",
			filter:          bson.D{{"_id", objectId}},
			dest:            &testStruct{},
			option:          initOptionFindOne(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:   "failed",
			filter: bson.D{{"_id", objectId}},
			dest:   &testStruct{},
			option: initOptionFindOne().
				SetCollation(&option.Collation{}).
				SetSkip(2),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:            "failed struct dest",
			filter:          nil,
			dest:            &testInvalidStruct{},
			option:          initOptionFindOne(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type dest",
			filter:          nil,
			dest:            initTestString(),
			option:          initOptionFindOne(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed dest non pointer",
			filter:          nil,
			dest:            *initTestString(),
			option:          initOptionFindOne(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestFindOneAndDelete() []testFindOneAndDelete {
	objectId, _ := primitive.ObjectIDFromHex(os.Getenv(MongoDBTestId))
	return []testFindOneAndDelete{
		{
			name:            "success",
			filter:          bson.D{{"_id", objectId}},
			dest:            &testStruct{},
			option:          initOptionFindOneAndDelete(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed not found",
			filter:          bson.D{{"_id", primitive.NewObjectID()}},
			dest:            &testStruct{},
			option:          initOptionFindOneAndDelete(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:   "failed",
			filter: bson.D{{"_id", objectId}},
			dest:   &testStruct{},
			option: initOptionFindOneAndDelete().
				SetCollation(&option.Collation{}).
				SetDisableAutoCloseSession(true),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:            "failed struct dest",
			filter:          nil,
			dest:            &testInvalidStruct{},
			option:          initOptionFindOneAndDelete(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type dest",
			filter:          nil,
			dest:            initTestString(),
			option:          initOptionFindOneAndDelete(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed dest non pointer",
			filter:          nil,
			dest:            *initTestString(),
			option:          initOptionFindOneAndDelete(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestFindOneAndReplace() []testFindOneAndReplace {
	objectId, _ := primitive.ObjectIDFromHex(os.Getenv(MongoDBTestId))
	return []testFindOneAndReplace{
		{
			name:            "success",
			filter:          bson.D{{"_id", objectId}},
			replacement:     *initTestStruct(),
			dest:            &testStruct{},
			option:          initOptionFindOneAndReplace(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed not found",
			filter:          bson.D{{"_id", primitive.NewObjectID()}},
			replacement:     *initTestStruct(),
			dest:            &testStruct{},
			option:          initOptionFindOneAndReplace(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:        "failed",
			filter:      bson.D{{"_id", objectId}},
			replacement: *initTestStruct(),
			dest:        &testStruct{},
			option: initOptionFindOneAndReplace().
				SetCollation(&option.Collation{}).
				SetReturnDocument(option.ReturnDocumentAfter).
				SetBypassDocumentValidation(true).
				SetDisableAutoCloseSession(true),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:            "failed struct dest",
			filter:          nil,
			replacement:     nil,
			dest:            &testInvalidStruct{},
			option:          initOptionFindOneAndReplace(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type dest",
			filter:          nil,
			replacement:     nil,
			dest:            initTestString(),
			option:          initOptionFindOneAndReplace(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed dest non pointer",
			filter:          nil,
			dest:            *initTestString(),
			option:          initOptionFindOneAndReplace(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestFindOneAndUpdate() []testFindOneAndUpdate {
	objectId, _ := primitive.ObjectIDFromHex(os.Getenv(MongoDBTestId))
	return []testFindOneAndUpdate{
		{
			name:            "success",
			filter:          bson.D{{"_id", objectId}},
			update:          bson.M{"$set": bson.M{"name": "Updated Test Name"}},
			dest:            &testStruct{},
			option:          initOptionFindOneAndUpdate(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed not found",
			filter:          bson.D{{"_id", primitive.NewObjectID()}},
			update:          bson.M{"$set": bson.M{"name": "Updated Test Name"}},
			dest:            &testStruct{},
			option:          initOptionFindOneAndUpdate(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:   "failed struct dest",
			filter: nil,
			update: nil,
			dest:   &testInvalidStruct{},
			option: initOptionFindOneAndUpdate().
				SetArrayFilters(&option.ArrayFilters{}).
				SetCollation(&option.Collation{}).
				SetReturnDocument(option.ReturnDocumentAfter).
				SetBypassDocumentValidation(true).
				SetDisableAutoCloseSession(true),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type dest",
			filter:          nil,
			update:          nil,
			dest:            initTestString(),
			option:          initOptionFindOneAndUpdate(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed dest non pointer",
			filter:          nil,
			update:          nil,
			dest:            *initTestString(),
			option:          initOptionFindOneAndUpdate(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestFind() []testFind {
	return []testFind{
		{
			name:            "success",
			filter:          bson.D{{"_id", bson.D{{"$exists", true}}}},
			dest:            &[]testStruct{},
			option:          initOptionFind(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:   "failed",
			filter: bson.D{{"_id", bson.D{{"$exists", true}}}},
			dest:   &[]testStruct{},
			option: initOptionFind().
				SetReturnKey(true).
				SetBatchSize(10).
				SetCursorType(option.CursorTypeTailable).
				SetCollation(&option.Collation{}).
				SetNoCursorTimeout(true).
				SetLimit(10).
				SetSkip(1),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:            "failed struct dest",
			filter:          nil,
			dest:            &[]testInvalidStruct{},
			option:          initOptionFind(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed slice type dest",
			filter:          nil,
			dest:            &[]string{},
			option:          initOptionFind(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type dest",
			filter:          nil,
			dest:            initTestString(),
			option:          initOptionFind(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed dest non pointer",
			filter:          nil,
			dest:            *initTestString(),
			option:          initOptionFind(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestFindPageable() []testFindPageable {
	return []testFindPageable{
		{
			name:   "success",
			filter: bson.D{{"_id", bson.D{{"$exists", true}}}},
			pageInput: PageInput{
				Page:     0,
				PageSize: 10,
				Ref:      []testStruct{},
				Sort:     bson.M{"createdAt": SortAsc},
			},
			option:          initOptionFindPageable(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:   "success empty",
			filter: bson.D{{"_id", bson.D{{"$exists", false}}}},
			pageInput: PageInput{
				Page:     0,
				PageSize: 10,
				Ref:      []testStruct{},
				Sort:     nil,
			},
			option:          initOptionFindPageable(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:   "failed timeout",
			filter: bson.D{{"_id", bson.D{{"$exists", true}}}},
			pageInput: PageInput{
				Page:     0,
				PageSize: 10,
				Ref:      []testStruct{},
				Sort:     nil,
			},
			option: initOptionFindPageable().
				SetBatchSize(10).
				SetReturnKey(true).
				SetCursorType(option.CursorTypeTailable).
				SetNoCursorTimeout(true).
				SetCollation(&option.Collation{}),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:   "failed struct ref",
			filter: nil,
			pageInput: PageInput{
				Page:     0,
				PageSize: 10,
				Ref:      []testInvalidStruct{},
				Sort:     nil,
			},
			option:          initOptionFindPageable(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:   "failed type ref pointer",
			filter: nil,
			pageInput: PageInput{
				Page:     0,
				PageSize: 10,
				Ref:      &testStruct{},
				Sort:     nil,
			},
			option:          initOptionFindPageable(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:   "failed type ref invalid",
			filter: nil,
			pageInput: PageInput{
				Page:     0,
				PageSize: 10,
				Sort:     nil,
			},
			option:          initOptionFindPageable(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:   "failed type ref",
			filter: bson.D{{"_id", bson.D{{"$exists", true}}}},
			pageInput: PageInput{
				Page:     0,
				PageSize: 10,
				Ref:      testStruct{},
				Sort:     nil,
			},
			option:          initOptionFindPageable(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestExists() []testExists {
	return []testExists{
		{
			name:            "success",
			filter:          bson.D{{"_id", bson.D{{"$exists", true}}}},
			ref:             testStruct{},
			option:          initOptionExists(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "success not exists",
			filter:          bson.D{{"_id", bson.D{{"$exists", false}}}},
			ref:             testStruct{},
			option:          initOptionExists(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed timeout",
			filter:          bson.D{{"_id", bson.D{{"$exists", false}}}},
			ref:             testStruct{},
			option:          initOptionExists().SetCollation(&option.Collation{}),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:            "failed ref",
			filter:          bson.D{{"_id", bson.D{{"$exists", false}}}},
			ref:             testInvalidStruct{},
			option:          initOptionExists(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type ref",
			filter:          bson.D{{"_id", bson.D{{"$exists", false}}}},
			ref:             "",
			option:          initOptionExists(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestExistsById() []testExistsById {
	objectId, _ := primitive.ObjectIDFromHex(os.Getenv(MongoDBTestId))
	return []testExistsById{
		{
			name:            "success",
			id:              objectId,
			ref:             testStruct{},
			option:          initOptionExists(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "success not exists",
			id:              primitive.NewObjectID(),
			ref:             testStruct{},
			option:          initOptionExists(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed timeout",
			id:              primitive.NewObjectID(),
			ref:             testStruct{},
			option:          initOptionExists(),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:            "failed ref",
			id:              primitive.NewObjectID(),
			ref:             testInvalidStruct{},
			option:          initOptionExists(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type ref",
			id:              primitive.NewObjectID(),
			ref:             "",
			option:          initOptionExists(),
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
			name:     "failed timeout",
			pipeline: Pipeline{bson.D{{"$match", bson.D{{"_id", bson.M{"$exists": true}}}}}},
			dest:     &[]testStruct{},
			option: initOptionAggregate().
				SetAllowDiskUse(true).
				SetCollation(&option.Collation{}).
				SetBatchSize(10),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:            "failed cursor dest",
			pipeline:        Pipeline{bson.D{{"$match", bson.D{{"_id", bson.M{"$exists": true}}}}}},
			dest:            &testStruct{},
			option:          initOptionAggregate(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed struct dest",
			pipeline:        nil,
			dest:            &testInvalidStruct{},
			option:          initOptionAggregate(),
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
		{
			name:            "failed dest non pointer",
			pipeline:        nil,
			dest:            *initTestString(),
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
			name:   "failed",
			filter: bson.D{{"_id", bson.M{"$exists": true}}},
			ref:    &testStruct{},
			option: initOptionCount().
				SetCollation(&option.Collation{}).
				SetLimit(10).
				SetSkip(10),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:            "failed filter",
			filter:          "filter string err",
			option:          initOptionCount(),
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

func initListTestDistinct() []testDistinct {
	return []testDistinct{
		{
			name:            "success",
			fieldName:       "_id",
			filter:          bson.D{{"_id", bson.D{{"$exists", true}}}},
			dest:            &[]primitive.ObjectID{},
			ref:             testStruct{},
			option:          initOptionDistinct(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "success no content",
			fieldName:       "_id",
			filter:          bson.D{{"_id", bson.D{{"$exists", false}}}},
			dest:            &[]primitive.ObjectID{},
			ref:             testStruct{},
			option:          initOptionDistinct(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed timeout",
			fieldName:       "_id",
			filter:          bson.D{{"_id", bson.D{{"$exists", true}}}},
			dest:            &[]primitive.ObjectID{},
			ref:             testStruct{},
			option:          initOptionDistinct(),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:      "failed struct dest",
			fieldName: "_id",
			filter:    bson.D{{"_id", bson.D{{"$exists", true}}}},
			dest:      &[]int{},
			ref:       testStruct{},
			option: initOptionDistinct().
				SetCollation(&option.Collation{}),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed type ref",
			fieldName:       "_id",
			filter:          nil,
			dest:            &[]primitive.ObjectID{},
			ref:             initTestString(),
			option:          initOptionDistinct(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed dest non pointer",
			fieldName:       "_id",
			filter:          nil,
			dest:            []primitive.ObjectID{},
			ref:             initTestString(),
			option:          initOptionDistinct(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestWatch() []testWatch {
	return []testWatch{
		{
			name: "success",
			pipeline: Pipeline{bson.D{{"$match", bson.D{
				{"operationType", bson.M{"$in": []string{"insert", "update", "delete", "replace"}}},
			}}}},
			option:          initOptionWatch(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:     "failed database",
			pipeline: nil,
			option: initOptionWatch().
				SetDatabaseName("test").
				SetCollation(&option.Collation{}).
				SetBatchSize(10).
				SetFullDocument(option.FullDocumentDefault).
				SetFullDocumentBeforeChange(option.FullDocumentOff).
				SetResumeAfter(bson.M{}).
				SetStartAtOperationTime(primitive.Timestamp{}).
				SetStartAfter(bson.M{}).
				SetCustom(bson.M{}).
				SetCustomPipeline(bson.M{}),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:     "failed collection",
			pipeline: nil,
			option: initOptionWatch().
				SetDatabaseName("test").
				SetCollectionName("test"),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestWatchHandler() []testWatchHandler {
	return []testWatchHandler{
		{
			name: "success",
			pipeline: Pipeline{bson.D{{"$match", bson.D{
				{"operationType", bson.M{"$in": []string{"insert", "update", "delete", "replace"}}},
			}}}},
			handler: func(ctx *ContextWatch) {
				logger.Info("watch handler ctx:", ctx)
			},
			option:          initOptionWatchHandler(),
			durationTimeout: 5 * time.Second,
		},
		{
			name: "success and failed timeout",
			pipeline: Pipeline{bson.D{{"$match", bson.D{
				{"operationType", bson.M{"$in": []string{"insert", "update", "delete", "replace"}}},
			}}}},
			handler: func(ctx *ContextWatch) {
				logger.Info("watch handler ctx:", ctx)
				time.Sleep(5 * time.Nanosecond)
			},
			option:          initOptionWatchHandler().SetContextFuncTimeout(1 * time.Nanosecond),
			durationTimeout: 5 * time.Second,
		},
		{
			name:     "failed",
			pipeline: nil,
			handler: func(ctx *ContextWatch) {
			},
			option: initOptionWatchHandler().
				SetDatabaseName("test").
				SetCollectionName("test").
				SetShowExpandedEvents(true).
				SetStartAtOperationTime(primitive.Timestamp{}).
				SetCollation(&option.Collation{}).
				SetBatchSize(10).
				SetFullDocument(option.FullDocumentDefault).
				SetFullDocumentBeforeChange(option.FullDocumentOff).
				SetResumeAfter(bson.M{}).
				SetStartAtOperationTime(primitive.Timestamp{}).
				SetStartAfter(bson.M{}).
				SetCustom(bson.M{}).
				SetCustomPipeline(bson.M{}),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed handler nil",
			pipeline:        nil,
			handler:         nil,
			option:          initOptionWatchHandler(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestDrop() []testDrop {
	return []testDrop{
		{
			name:            "success",
			ref:             testStruct{},
			durationTimeout: 5 * time.Minute,
		},
		{
			name:            "failed",
			ref:             testInvalidStruct{},
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:            "failed ref",
			ref:             testInvalidStruct{},
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestCreateOneIndex() []testCreateOneIndex {
	return []testCreateOneIndex{
		{
			name:            "success",
			input:           initIndexInput(),
			durationTimeout: 5 * time.Second,
		},
		{
			name: "failed",
			input: IndexInput{
				Keys: nil,
				Options: initOptionIndex().
					SetExpireAfterSeconds(600000).
					SetName("test failed").
					SetSparse(true).
					SetStorageEngine(bson.M{}).
					SetUnique(true).
					SetVersion(1).
					SetDefaultLanguage("pt-br").
					SetLanguageOverride("pt-br").
					SetTextVersion(1).
					SetWeights(bson.M{}).
					SetSphereVersion(1).
					SetBits(100).
					SetMax(180).
					SetMin(-180).
					SetBucketSize(0).
					SetPartialFilterExpression(bson.M{}).
					SetCollation(&option.Collation{}).
					SetWildcardProjection(bson.M{}).
					SetHidden(true),
				Ref: testStruct{},
			},
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name: "failed ref",
			input: IndexInput{
				Keys:    bson.D{{"random", 1}},
				Options: initOptionIndex().SetName("test index unique").SetUnique(true).SetSparse(true),
				Ref:     testInvalidStruct{},
			},
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestCreateManyIndex() []testCreateManyIndex {
	return []testCreateManyIndex{
		{
			name: "success",
			inputs: []IndexInput{
				{
					Keys:    bson.D{{"random", 1}},
					Options: initOptionIndex().SetName("test index unique").SetUnique(true).SetSparse(true),
					Ref:     testStruct{},
				},
				{
					Keys:    bson.D{{"createdAt", 1}},
					Options: initOptionIndex().SetName("test index expire").SetExpireAfterSeconds(300),
					Ref:     testStruct{},
				},
			},
			durationTimeout: 5 * time.Second,
		},
		{
			name: "success partial",
			inputs: []IndexInput{
				{
					Keys:    bson.D{{"random", 1}},
					Options: initOptionIndex().SetName("test index unique").SetUnique(true).SetSparse(true),
					Ref:     testStruct{},
				},
				{
					Keys:    bson.D{{"createdAt", 1}},
					Options: initOptionIndex().SetName("test index expire").SetExpireAfterSeconds(300),
					Ref:     testStruct{},
				},
				{
					Keys:    bson.D{{"updatedAt", 1}},
					Options: initOptionIndex(),
					Ref:     testInvalidStruct{},
				},
				{
					Keys:    nil,
					Options: initOptionIndex(),
					Ref:     testStruct{},
				},
				{
					Keys:    nil,
					Options: initOptionIndex(),
					Ref:     "",
				},
			},
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name: "failed",
			inputs: []IndexInput{
				{
					Keys: nil,
					Options: initOptionIndex().
						SetExpireAfterSeconds(600000).
						SetName("test failed").
						SetSparse(true).
						SetStorageEngine(bson.M{}).
						SetUnique(true).
						SetVersion(1).
						SetDefaultLanguage("pt-br").
						SetLanguageOverride("pt-br").
						SetTextVersion(1).
						SetWeights(bson.M{}).
						SetSphereVersion(1).
						SetBits(100).
						SetMax(180).
						SetMin(-180).
						SetBucketSize(0).
						SetPartialFilterExpression(bson.M{}).
						SetCollation(&option.Collation{}).
						SetWildcardProjection(bson.M{}).
						SetHidden(true),
					Ref: testStruct{},
				},
			},
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name: "failed ref",
			inputs: []IndexInput{
				{
					Keys: bson.D{{"random", 1}},
					Options: initOptionIndex().
						SetName("test index unique").
						SetUnique(true).
						SetSparse(true),
					Ref: testInvalidStruct{},
				},
			},
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
		{
			name:            "failed empty inputs",
			inputs:          []IndexInput{},
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestListIndexes() []testListIndexes {
	return []testListIndexes{
		{
			name:            "success",
			ref:             testStruct{},
			option:          initOptionListIndexes(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed",
			ref:             testStruct{},
			option:          initOptionListIndexes().SetBatchSize(10),
			durationTimeout: 1 * time.Nanosecond,
			wantErr:         true,
		},
		{
			name:            "failed ref",
			ref:             "",
			option:          initOptionListIndexes(),
			durationTimeout: 5 * time.Second,
			wantErr:         true,
		},
	}
}

func initListTestDropIndex() []testDropIndex {
	return []testDropIndex{
		{
			name:            "success",
			nameIndex:       os.Getenv(MongoDBIndexName),
			ref:             testStruct{},
			option:          initOptionDropIndex(),
			durationTimeout: 5 * time.Second,
		},
		{
			name:            "failed",
			nameIndex:       "",
			option:          initOptionDropIndex(),
			durationTimeout: 1 * time.Millisecond,
			wantErr:         true,
		},
		{
			name:            "failed ref",
			nameIndex:       os.Getenv(MongoDBIndexName),
			ref:             "",
			option:          initOptionDropIndex(),
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
		SetDisableAutoCloseSession(false)
}

func initOptionInsertMany() option.InsertMany {
	return option.NewInsertMany().
		SetBypassDocumentValidation(true).
		SetForceRecreateSession(true).
		SetComment("comment insert golang unit test").
		SetDisableAutoCloseSession(false).
		SetDisableAutoRollback(false)
}

func initOptionDelete() option.Delete {
	return option.NewDelete().
		SetForceRecreateSession(true).
		SetDisableAutoCloseSession(false).
		SetComment("comment delete golang unit test").
		SetCollation(&option.Collation{}).
		SetHint(bson.M{}).
		SetLet(bson.M{})
}

func initOptionUpdate() option.Update {
	return option.NewUpdate().
		SetDisableAutoCloseSession(false).
		SetForceRecreateSession(true).
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
		SetForceRecreateSession(true).
		SetDisableAutoCloseSession(false).
		SetComment("comment replace golang unit test").
		SetCollation(&option.Collation{}).
		SetHint(bson.M{}).
		SetLet(bson.M{}).
		SetBypassDocumentValidation(true)
}

func initOptionFindOne() option.FindOne {
	return option.NewFindOne().
		SetAllowPartialResults(true).
		SetCollation(nil).
		SetMaxTime(5 * time.Second).
		SetComment("comment golang unit test").
		SetHint(bson.M{}).
		SetMax(bson.M{}).
		SetMin(bson.M{}).
		SetProjection(bson.M{}).
		SetReturnKey(true).
		SetShowRecordID(true).
		SetSkip(0).
		SetSort(bson.M{})

}

func initOptionFindOneById() option.FindOneById {
	return option.NewFindOneById().
		SetAllowPartialResults(true).
		SetCollation(nil).
		SetMaxTime(5 * time.Second).
		SetComment("comment golang unit test").
		SetHint(bson.M{}).
		SetMax(bson.M{}).
		SetMin(bson.M{}).
		SetProjection(bson.M{}).
		SetReturnKey(true).
		SetShowRecordID(true)

}

func initOptionFindOneAndDelete() option.FindOneAndDelete {
	return option.NewFindOneAndDelete().
		SetForceRecreateSession(true).
		SetCollation(nil).
		SetComment("comment golang unit test").
		SetHint(bson.M{}).
		SetMaxTime(5 * time.Second).
		SetProjection(bson.M{}).
		SetSort(bson.M{}).
		SetLet(bson.M{}).
		SetDisableAutoCloseSession(false)
}

func initOptionFindOneAndReplace() option.FindOneAndReplace {
	return option.NewFindOneAndReplace().
		SetForceRecreateSession(true).
		SetCollation(nil).
		SetComment("comment golang unit test").
		SetHint(bson.M{}).
		SetMaxTime(5 * time.Second).
		SetProjection(bson.M{}).
		SetSort(bson.M{}).
		SetLet(bson.M{}).
		SetDisableAutoCloseSession(false).
		SetUpsert(true)
}

func initOptionFindOneAndUpdate() option.FindOneAndUpdate {
	return option.NewFindOneAndUpdate().
		SetForceRecreateSession(true).
		SetCollation(nil).
		SetComment("comment golang unit test").
		SetHint(bson.M{}).
		SetMaxTime(5 * time.Second).
		SetProjection(bson.M{}).
		SetSort(bson.M{}).
		SetLet(bson.M{}).
		SetDisableAutoCloseSession(false).
		SetUpsert(true)
}

func initOptionFind() option.Find {
	return option.NewFind().
		SetCollation(nil).
		SetComment("comment golang unit test").
		SetHint(bson.M{}).
		SetMaxAwaitTime(2 * time.Second).
		SetMaxTime(5 * time.Second).
		SetProjection(bson.M{}).
		SetSort(bson.M{"createdAt": SortDesc}).
		SetLet(bson.M{}).
		SetAllowPartialResults(true).
		SetShowRecordID(true).
		SetAllowDiskUse(true).
		SetNoCursorTimeout(false).
		SetMin(bson.M{}).
		SetMax(bson.M{})
}

func initOptionFindPageable() option.FindPageable {
	return option.NewFindPageable().
		SetCollation(nil).
		SetComment("comment golang unit test").
		SetHint(bson.M{}).
		SetMaxTime(5 * time.Second).
		SetMaxAwaitTime(2 * time.Second).
		SetProjection(bson.M{}).
		SetLet(bson.M{}).
		SetAllowPartialResults(true).
		SetShowRecordID(true).
		SetAllowDiskUse(true).
		SetNoCursorTimeout(false).
		SetMin(bson.M{}).
		SetMax(bson.M{})
}

func initOptionExists() option.Exists {
	return option.NewExists().
		SetCollation(nil).
		SetComment("comment count golang unit test").
		SetHint(bson.M{}).
		SetMaxTime(5 * time.Second)
}

func initOptionAggregate() option.Aggregate {
	return option.NewAggregate().
		SetAllowDiskUse(false).
		SetBatchSize(0).
		SetBypassDocumentValidation(true).
		SetCollation(nil).
		SetMaxTime(5 * time.Second).
		SetMaxAwaitTime(2 * time.Second).
		SetComment("comment golang unit test").
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

func initOptionDistinct() option.Distinct {
	return option.NewDistinct().
		SetCollation(nil).
		SetComment("comment golang unit test").
		SetMaxTime(5 * time.Second)
}

func initOptionWatch() option.Watch {
	return option.NewWatch().
		SetBatchSize(0).
		SetCollation(nil).
		SetComment("comment golang unit test").
		SetMaxAwaitTime(2 * time.Second).
		SetShowExpandedEvents(true)
}

func initOptionWatchHandler() option.WatchHandler {
	return option.NewWatchHandler().
		SetContextFuncTimeout(5 * time.Second).
		SetDelayLoop(5 * time.Second).
		SetComment("comment golang unit test").
		SetFullDocument(option.FullDocumentDefault).
		SetFullDocumentBeforeChange(option.FullDocumentOff).
		SetMaxAwaitTime(2 * time.Second)
}

func initOptionIndex() option.Index {
	return option.NewIndex()
}

func initOptionDropIndex() option.DropIndex {
	return option.NewDropIndex().
		SetMaxTime(5 * time.Second)
}

func initOptionListIndexes() option.ListIndexes {
	return option.NewListIndexes().
		SetMaxTime(5 * time.Second)
}
