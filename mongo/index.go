package mongo

import (
	"github.com/GabrielHCataldo/go-mongo-template/mongo/option"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IndexInput represents a new index to be created.
type IndexInput struct {
	// Keys A document describing which keys should be used for the index. It cannot be nil. This must be an order-preserving
	// type such as bson.D. Map types such as bson.M are not valid. See https://www.mongodb.com/docs/manual/indexes/#indexes
	// for examples of valid documents.
	Keys any
	// Options The options to use to create the index.
	Options option.Index
	// Ref Struct reference contained database and collection tag
	Ref any
}

type IndexResult struct {
	Id         any             `bson:"id,omitempty"`
	Ns         string          `bson:"ns,omitempty"`
	FirstBatch FirstBatchIndex `bson:"firstBatch,omitempty"`
}

type FirstBatchIndex struct {
	V    int         `bson:"v,omitempty"`
	Key  primitive.M `bson:"key,omitempty"`
	Name string      `bson:"name,omitempty"`
	Ns   string      `bson:"ns,omitempty"`
}

func parseIndexInputToModel(input IndexInput) mongo.IndexModel {
	return mongo.IndexModel{
		Keys: input.Keys,
		Options: &options.IndexOptions{
			ExpireAfterSeconds:      input.Options.ExpireAfterSeconds,
			Name:                    input.Options.Name,
			Sparse:                  input.Options.Sparse,
			StorageEngine:           input.Options.StorageEngine,
			Unique:                  input.Options.Unique,
			Version:                 input.Options.Version,
			DefaultLanguage:         input.Options.DefaultLanguage,
			LanguageOverride:        input.Options.LanguageOverride,
			TextVersion:             input.Options.TextVersion,
			Weights:                 input.Options.Weights,
			SphereVersion:           input.Options.SphereVersion,
			Bits:                    input.Options.Bits,
			Max:                     input.Options.Max,
			Min:                     input.Options.Min,
			BucketSize:              input.Options.BucketSize,
			PartialFilterExpression: input.Options.PartialFilterExpression,
			Collation:               option.ParseCollationMongoOptions(input.Options.Collation),
			WildcardProjection:      input.Options.WildcardProjection,
			Hidden:                  input.Options.Hidden,
		},
	}
}
