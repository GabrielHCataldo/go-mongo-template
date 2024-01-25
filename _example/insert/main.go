package main

import (
	"context"
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"github.com/GabrielHCataldo/go-mongo-template/mongo"
	"github.com/GabrielHCataldo/go-mongo-template/mongo/option"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"os"
	"time"
)

type test struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" database:"test" collection:"test"`
	Random    int                `json:"random,omitempty" bson:"random,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	BirthDate primitive.DateTime `json:"birthDate,omitempty" bson:"birthDate,omitempty"`
	Emails    []string           `json:"emails,omitempty" bson:"emails,omitempty"`
	Balance   float64            `json:"balance,omitempty" bson:"balance,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

func main() {
	insertOne()
	insertMany()
	insertOneManualCloseSession()
}

func insertOne() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	testDocument := test{
		Random:    rand.Int(),
		Name:      "Foo Bar",
		BirthDate: primitive.NewDateTimeFromTime(time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local)),
		Emails:    []string{"foobar@gmail.com", "foobar3@hotmail.com"},
		Balance:   190.12,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	//new document need a pointer
	err = mongoTemplate.InsertOne(ctx, &testDocument)
	if helper.IsNotNil(err) {
		logger.Error("error insert document:", err)
	} else {
		logger.Info("document inserted successfully:", testDocument)
	}
}

func insertMany() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	testDocuments := []*test{
		{
			Random:    rand.Int(),
			Name:      "Foo Bar",
			BirthDate: primitive.NewDateTimeFromTime(time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local)),
			Emails:    []string{"foobar@gmail.com", "foobar3@hotmail.com"},
			Balance:   190.12,
			CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		},
		{
			Random:    rand.Int(),
			Name:      "Foo Bar 2",
			BirthDate: primitive.NewDateTimeFromTime(time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local)),
			Emails:    []string{"foobar2@gmail.com", "foobar4@hotmail.com"},
			Balance:   290.12,
			CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		},
	}
	//new document need a slice of the pointer
	err = mongoTemplate.InsertMany(ctx, testDocuments)
	if helper.IsNotNil(err) {
		logger.Error("error insert document:", err)
	} else {
		logger.Info("document inserted successfully:", testDocuments)
	}
}

func insertOneManualCloseSession() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	testDocument := test{
		Random:    rand.Int(),
		Name:      "Foo Bar",
		BirthDate: primitive.NewDateTimeFromTime(time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local)),
		Emails:    []string{"foobar@gmail.com", "foobar3@hotmail.com"},
		Balance:   190.12,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	//new document need a pointer
	err = mongoTemplate.InsertOne(ctx, &testDocument, option.NewInsertOne().
		SetForceRecreateSession(false).SetDisableAutoCloseSession(true))
	if helper.IsNotNil(err) {
		logger.Error("error insert document:", err)
	} else {
		logger.Info("document inserted successfully:", testDocument)
	}
	//mongoTemplate.AbortTransaction(ctx) or mongoTemplate.CommitTransaction(ctc)
	abort := helper.IsNotNil(err)
	err = mongoTemplate.CloseSession(ctx, abort)
	if helper.IsNotNil(err) {
		logger.Error("error close session:", err)
	}
}
