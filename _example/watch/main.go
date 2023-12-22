package main

import (
	"context"
	"github.com/GabrielHCataldo/go-logger/logger"
	"github.com/GabrielHCataldo/go-mongo/mongo"
	"github.com/GabrielHCataldo/go-mongo/mongo/option"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func main() {
	watchHandler()
}

func watchHandler() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	mongoTemplate, err := mongo.NewTemplate(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		logger.Error("error to init mongo template:", err)
		return
	}
	defer disconnect(ctx, mongoTemplate)
	pipeline := mongo.Pipeline{bson.D{{"$match", bson.D{
		{"operationType", bson.M{"$in": []string{"insert", "update", "delete", "replace"}}},
	}}}}
	err = mongoTemplate.WatchHandler(context.TODO(), pipeline, handler, option.NewWatchHandler().SetDatabaseName("test"))
	if err != nil {
		logger.Error("error watch handler:", err)
	} else {
		logger.Info("watch handler complete successfully")
	}
}

func handler(ctx *mongo.ContextWatch) {
	logger.Info("handler watch called:", ctx)
}

func disconnect(ctx context.Context, mongoTemplate mongo.Template) {
	err := mongoTemplate.Disconnect(ctx)
	if err != nil {
		logger.Error("error disconnect mongodb:", err)
	}
}
