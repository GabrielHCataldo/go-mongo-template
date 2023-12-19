package mongo

import (
	"context"
	"github.com/GabrielHCataldo/go-logger/logger"
	"go-mongo/mongo/option"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"testing"
	"time"
)

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

var mongoTemplate Template

type testStruct struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty" database:"test" collection:"test"`
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
	t.Run()
	disconnectMongoTemplate()
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

func initListTestNewTemplate() []testNewTemplate {
	return []testNewTemplate{
		{
			name:            "success",
			options:         options.Client().ApplyURI(os.Getenv("MONGODB_URL")),
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
			options:         options.Client().ApplyURI(os.Getenv("MONGODB_URL")),
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

func initOptionInsertOne() option.InsertOne {
	return option.NewInsertOne().
		SetBypassDocumentValidation(true).
		SetForceRecreateSession(true).
		SetComment("comment insert golang unit test").
		SetDisableAutoCloseTransaction(false)
}

func disconnectMongoTemplate() {
	if mongoTemplate == nil {
		return
	}
	mongoTemplate.Disconnect()
	mongoTemplate = nil
}
