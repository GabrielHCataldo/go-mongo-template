package main

import (
	"context"
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"github.com/GabrielHCataldo/go-mongo-template/mongo"
	"go.mongodb.org/mongo-driver/bson"
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
	replaceOne()
	replaceOneById()
}

func replaceOne() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	filter := bson.M{"_id": bson.M{"$exists": true}}
	replacement := test{
		Random:    rand.Int(),
		Name:      "Foo Bar",
		BirthDate: primitive.NewDateTimeFromTime(time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local)),
		Emails:    []string{"foobar@gmail.com", "foobar3@hotmail.com"},
		Balance:   190.12,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	updateResult, err := mongoTemplate.ReplaceOne(ctx, filter, replacement, test{})
	if helper.IsNotNil(err) {
		logger.Error("error replace document:", err)
	} else {
		logger.Info("document replaced successfully:", updateResult)
	}
}

func replaceOneById() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	objectId, _ := primitive.ObjectIDFromHex("6585db26633e225cbeadf553")
	replacement := test{
		Random:    rand.Int(),
		Name:      "Foo Bar",
		BirthDate: primitive.NewDateTimeFromTime(time.Date(1999, 1, 21, 0, 0, 0, 0, time.Local)),
		Emails:    []string{"foobar@gmail.com", "foobar3@hotmail.com"},
		Balance:   190.12,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	updateResult, err := mongoTemplate.ReplaceOneById(ctx, objectId, replacement, test{})
	if helper.IsNotNil(err) {
		logger.Error("error replace document:", err)
	} else {
		logger.Info("document replaced successfully:", updateResult)
	}
}
