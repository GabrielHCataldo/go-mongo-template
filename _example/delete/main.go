package main

import (
	"context"
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"github.com/GabrielHCataldo/go-mongo-template/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	deleteOne()
	deleteOneById()
	deleteMany()
}

func deleteOne() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	filter := bson.M{"_id": bson.M{"$exists": true}}
	deleteResult, err := mongoTemplate.DeleteOne(ctx, filter, test{})
	if helper.IsNotNil(err) {
		logger.Error("error delete document:", err)
	} else {
		logger.Info("document deleted successfully:", deleteResult)
	}
}

func deleteOneById() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	objectId, _ := primitive.ObjectIDFromHex("6585f3a8bf2af8ad9bcab912")
	deleteResult, err := mongoTemplate.DeleteOneById(ctx, objectId, test{})
	if helper.IsNotNil(err) {
		logger.Error("error delete document:", err)
	} else {
		logger.Info("document deleted successfully:", deleteResult)
	}
}

func deleteMany() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	filter := bson.M{"_id": bson.M{"$exists": true}}
	deleteResult, err := mongoTemplate.DeleteMany(ctx, filter, test{})
	if helper.IsNotNil(err) {
		logger.Error("error delete documents:", err)
	} else {
		logger.Info("documents deleted successfully:", deleteResult)
	}
}
