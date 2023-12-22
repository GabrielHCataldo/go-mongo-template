package mongo

import (
	"context"
	"github.com/GabrielHCataldo/go-logger/logger"
	"go-mongo/mongo/option"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WatchEvent struct {
	DocumentKey       documentKey         `bson:"documentKey"`
	NS                ns                  `bson:"ns"`
	OperationType     string              `bson:"operationType"`
	FullDocument      bson.M              `bson:"fullDocument"`
	UpdateDescription updateDescription   `bson:"updateDescription"`
	ClusterTime       primitive.Timestamp `bson:"clusterTime"`
}

type ns struct {
	DB   string `bson:"db"`
	Coll string `bson:"coll"`
}

type documentKey struct {
	ID primitive.ObjectID `bson:"_id"`
}

type updateDescription struct {
	UpdatedFields   map[string]any `bson:"updatedFields"`
	RemovedFields   []string       `bson:"removedFields"`
	TruncatedArrays []string       `bson:"truncatedArrays"`
}

type ContextWatch struct {
	context.Context
	Event WatchEvent
}

type HandlerWatch func(ctx *ContextWatch)

func processWatchNext(handler HandlerWatch, event WatchEvent, opt option.WatchHandler) {
	ctx, cancel := context.WithTimeout(context.TODO(), opt.ContextFuncTimeout)
	defer cancel()
	signal := make(chan struct{}, 1)
	go processWatchHandler(ctx, handler, event, &signal)
	select {
	case <-ctx.Done():
		logger.Error("Error timeout context func:", ctx.Err())
		break
	case <-signal:
		logger.Info("Handler watch processed!")
		break
	}
}

func processWatchHandler(ctx context.Context, handler HandlerWatch, event WatchEvent, signal *chan struct{}) {
	handler(&ContextWatch{
		Context: ctx,
		Event:   event,
	})
	*signal <- struct{}{}
}
