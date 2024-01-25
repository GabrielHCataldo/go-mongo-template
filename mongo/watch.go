package mongo

import (
	"context"
	"github.com/GabrielHCataldo/go-errors/errors"
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-mongo-template/mongo/option"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FullDocument bson.M

type Event struct {
	DocumentKey       documentKey         `bson:"documentKey"`
	NS                ns                  `bson:"ns"`
	OperationType     string              `bson:"operationType"`
	FullDocument      FullDocument        `bson:"fullDocument"`
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

type EventContext struct {
	context.Context
	Event Event
}

type EventHandler func(ctx *EventContext)

// ParseToStruct convert Event.FullDocument to struct
func (f FullDocument) ParseToStruct(dest any) error {
	if helper.IsNotPointer(dest) {
		return errDestIsNotPointer(2)
	} else if helper.IsNotStruct(dest) {
		return errDestIsNotStruct(2)
	}
	return errors.NewSkipCaller(2, helper.ConvertToDest(f, dest))
}

func processNextEvent(handler EventHandler, event Event, opt *option.WatchWithHandler) {
	ctx, cancel := context.WithTimeout(context.TODO(), opt.ContextFuncTimeout)
	defer cancel()
	signal := make(chan struct{}, 1)
	go executeEventHandler(ctx, handler, event, &signal)
	select {
	case <-ctx.Done():
	case <-signal:
		break
	}
}

func executeEventHandler(ctx context.Context, handler EventHandler, event Event, signal *chan struct{}) {
	handler(&EventContext{
		Context: ctx,
		Event:   event,
	})
	*signal <- struct{}{}
}
