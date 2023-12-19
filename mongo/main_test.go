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

type testInsertOne struct {
	name    string
	value   any
	option  option.InsertOne
	wantErr bool
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

type testEmptyStruct struct {
}

func TestMain(t *testing.M) {
	initMongoTemplate()
	defer mongoTemplate.Disconnect()
	t.Run()
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

func initListTestInsertOne() []testInsertOne {
	return []testInsertOne{
		{
			name:   "success",
			value:  initTestStruct(),
			option: initOptionInsertOne(),
		},
		{
			name:    "failed type value",
			value:   "string value",
			option:  initOptionInsertOne().SetDisableAutoCloseTransaction(true),
			wantErr: true,
		},
		{
			name:    "failed empty value",
			value:   initTestEmptyStruct(),
			option:  initOptionInsertOne().SetDisableAutoCloseTransaction(true),
			wantErr: true,
		},
		{
			name:    "failed non struct value",
			value:   initTestString(),
			option:  initOptionInsertOne().SetDisableAutoCloseTransaction(true),
			wantErr: true,
		},
	}
}

func initOptionInsertOne() option.InsertOne {
	return option.InsertOne{
		BypassDocumentValidation:    true,
		Comment:                     "comment insert golang unit test",
		DisableAutoCloseTransaction: false,
	}
}
