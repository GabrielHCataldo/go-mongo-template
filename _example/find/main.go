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
	findOne()
	findOneById()
	find()
	findPageable()
	findAll()
}

func findOne() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	filter := bson.M{"_id": bson.M{"$exists": true}}
	var dest test
	err = mongoTemplate.FindOne(ctx, filter, &dest)
	if helper.IsNotNil(err) {
		logger.Error("error find document:", err)
	} else {
		logger.Info("find document successfully:", dest)
	}
}

func findOneById() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	objectId, _ := primitive.ObjectIDFromHex("6585db26633e225cbeadf553")
	//dest need a pointer
	var dest test
	err = mongoTemplate.FindOneById(ctx, objectId, &dest)
	if helper.IsNotNil(err) {
		logger.Error("error find all documents:", err)
	} else {
		logger.Info("find by id document successfully:", dest)
	}
}

func find() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	filter := bson.M{"_id": bson.M{"$exists": true}}
	//dest need a pointer
	var dest []test
	err = mongoTemplate.Find(ctx, filter, &dest)
	if helper.IsNotNil(err) {
		logger.Error("error find documents:", err)
	} else {
		logger.Info("find documents successfully:", dest)
	}
}

func findPageable() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	filter := bson.M{"_id": bson.M{"$exists": true}}
	pageOutput, err := mongoTemplate.FindPageable(ctx, filter, mongo.PageInput{
		Page:     0,
		PageSize: 10,
		Ref:      []test{}, //need a slice of the structure
		Sort:     bson.M{"createdAt": mongo.SortDesc},
	})
	if helper.IsNotNil(err) {
		logger.Error("error find pageable documents:", err)
	} else {
		logger.Info("find pageable documents successfully:", pageOutput)
	}
}

func findAll() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	var dest []test
	err = mongoTemplate.FindAll(ctx, &dest)
	if helper.IsNotNil(err) {
		logger.Error("error find all documents:", err)
	} else {
		logger.Info("find all documents successfully:", dest)
	}
}
