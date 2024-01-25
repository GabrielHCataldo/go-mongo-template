package main

import (
	"context"
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"github.com/GabrielHCataldo/go-mongo-template/mongo"
	"github.com/GabrielHCataldo/go-mongo-template/mongo/option"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func main() {
	watch()
	watchHandler()
}

func watch() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	pipeline := mongo.Pipeline{bson.D{{"$match", bson.D{
		{"operationType", bson.M{"$in": []string{"insert", "update", "delete", "replace"}}},
	}}}}
	changeStream, err := mongoTemplate.Watch(context.TODO(), pipeline, option.NewWatch().SetDatabaseName("test"))
	if helper.IsNotNil(err) {
		logger.Error("error watch handler:", err)
		return
	}
	for changeStream.Next(context.TODO()) {
		logger.Info("changeStream called:", changeStream)
	}
	err = changeStream.Close(context.TODO())
	if helper.IsNotNil(err) {
		logger.Error("error close watch:", err)
	} else {
		logger.Info("watch complete successfully")
	}
}

func watchHandler() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if helper.IsNotNil(err) {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer mongoTemplate.SimpleDisconnect(ctx)
	pipeline := mongo.Pipeline{bson.D{{"$match", bson.D{
		{"operationType", bson.M{"$in": []string{"insert", "update", "delete", "replace"}}},
	}}}}
	opt := option.NewWatchWithHandler().SetDatabaseName("test")
	err = mongoTemplate.WatchWithHandler(context.TODO(), pipeline, handler, opt)
	if helper.IsNotNil(err) {
		logger.Error("error watch handler:", err)
	} else {
		logger.Info("watch handler complete successfully")
	}
}

func handler(ctx *mongo.EventContext) {
	logger.Info("handler called:", ctx)
}
