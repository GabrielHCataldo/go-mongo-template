package mongo

import (
	"context"
	"github.com/GabrielHCataldo/go-mongo-template/mongo/option"
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
	context.Context `json:"-"`
	Event           WatchEvent
}

type Handler func(ctx *ContextWatch)

func processToWatchNext(handler Handler, event WatchEvent, opt option.WatchWithHandler) {
	ctx, cancel := context.WithTimeout(context.TODO(), opt.ContextFuncTimeout)
	defer cancel()
	signal := make(chan struct{}, 1)
	go executeWatchHandler(ctx, handler, event, &signal)
	select {
	case <-ctx.Done():
	case <-signal:
		break
	}
}

func executeWatchHandler(ctx context.Context, handler Handler, event WatchEvent, signal *chan struct{}) {
	handler(&ContextWatch{
		Context: ctx,
		Event:   event,
	})
	*signal <- struct{}{}
}
