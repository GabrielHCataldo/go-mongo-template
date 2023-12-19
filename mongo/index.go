package mongo

import (
	"go-mongo/mongo/option"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// IndexInput represents a new index to be created.
type IndexInput struct {
	// A document describing which keys should be used for the index. It cannot be nil. This must be an order-preserving
	// type such as bson.D. Map types such as bson.M are not valid. See https://www.mongodb.com/docs/manual/indexes/#indexes
	// for examples of valid documents.
	Keys any
	// The options to use to create the index.
	Options *option.Index
}

type IndexOutput struct {
	Key                bson.M `bson:"key,omitempty"`
	Name               string `bson:"name,omitempty"`
	Ns                 string `bson:"ns,omitempty"`
	ExpireAfterSeconds int32  `bson:"expireAfterSeconds,omitempty"`
	Unique             bool   `bson:"unique,omitempty"`
}

func parseSliceIndexInputToModels(inputs []IndexInput) []mongo.IndexModel {
	var result []mongo.IndexModel
	for _, input := range inputs {
		result = append(result, parseIndexInputToModel(input))
	}
	return result
}

func parseIndexInputToModel(input IndexInput) mongo.IndexModel {
	var opts *options.IndexOptions
	if input.Options != nil {
		opts = &options.IndexOptions{
			ExpireAfterSeconds:      &input.Options.ExpireAfterSeconds,
			Name:                    &input.Options.Name,
			Sparse:                  &input.Options.Sparse,
			StorageEngine:           input.Options.StorageEngine,
			Unique:                  &input.Options.Unique,
			Version:                 &input.Options.Version,
			DefaultLanguage:         &input.Options.DefaultLanguage,
			LanguageOverride:        &input.Options.LanguageOverride,
			TextVersion:             &input.Options.TextVersion,
			Weights:                 input.Options.Weights,
			SphereVersion:           &input.Options.SphereVersion,
			Bits:                    &input.Options.Bits,
			Max:                     &input.Options.Max,
			Min:                     &input.Options.Min,
			BucketSize:              &input.Options.BucketSize,
			PartialFilterExpression: input.Options.PartialFilterExpression,
			Collation:               option.ParseCollationMongoOptions(input.Options.Collation),
			WildcardProjection:      input.Options.WildcardProjection,
			Hidden:                  &input.Options.Hidden,
		}
	}
	return mongo.IndexModel{
		Keys:    input.Keys,
		Options: opts,
	}
}
