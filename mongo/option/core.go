package option

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collation allows users to specify language-specific rules for string comparison, such as
// rules for letter case and accent marks.
type Collation struct {
	// Locale The locale
	Locale string `bson:",omitempty"`
	// CaseLevel The case level
	CaseLevel bool `bson:",omitempty"`
	// CaseFirst The case ordering
	CaseFirst string `bson:",omitempty"`
	// Strength The number of comparison levels to use
	Strength int `bson:",omitempty"`
	// NumericOrdering Whether to order numbers based on numerical order and not collation order
	NumericOrdering bool `bson:",omitempty"`
	// Alternate Whether spaces and punctuation are considered base characters
	Alternate string `bson:",omitempty"`
	// MaxVariable Which characters are affected by alternate: "shifted"
	MaxVariable string `bson:",omitempty"`
	// Normalization Causes text to be normalized into Unicode NFD
	Normalization bool `bson:",omitempty"`
	// Backwards Causes secondary differences to be considered in reverse order, as it is done in the French language
	Backwards bool `bson:",omitempty"`
}

// ArrayFilters is used to hold filters for the array filters CRUD option. If a registry is nil, bson.DefaultRegistry
// will be used when converting the filter interfaces to BSON.
type ArrayFilters struct {
	// Registry is the registry to use for converting filters. Defaults to bson.DefaultRegistry.
	//
	// Deprecated: Marshaling ArrayFilters to BSON will not be supported in Go Driver 2.0.
	Registry *bsoncodec.Registry
	// Filters The filters to apply
	Filters []any
}

// ParseCollationMongoOptions convert Collation to mongo options.Collation
func ParseCollationMongoOptions(c *Collation) *options.Collation {
	var mongoCollation *options.Collation
	if helper.IsNotNil(c) {
		mongoCollation = &options.Collation{
			Locale:          c.Locale,
			CaseLevel:       c.CaseLevel,
			CaseFirst:       c.CaseFirst,
			Strength:        c.Strength,
			NumericOrdering: c.NumericOrdering,
			Alternate:       c.Alternate,
			MaxVariable:     c.MaxVariable,
			Normalization:   c.Normalization,
			Backwards:       c.Backwards,
		}
	}
	return mongoCollation
}

// ParseArrayFiltersMongoOptions convert ArrayFilters to mongo options.ArrayFilters
func ParseArrayFiltersMongoOptions(a *ArrayFilters) *options.ArrayFilters {
	var mongoArrayFilters *options.ArrayFilters
	if helper.IsNotNil(a) {
		mongoArrayFilters = &options.ArrayFilters{
			Registry: a.Registry,
			Filters:  a.Filters,
		}
	}
	return mongoArrayFilters
}

// ParseCursorType convert CursorType to mongo options.CursorType
func ParseCursorType(c *CursorType) *options.CursorType {
	var mongoCursorType *options.CursorType
	if helper.IsNotNil(c) {
		value := options.CursorType(*c)
		mongoCursorType = &value
	}
	return mongoCursorType
}

// ParseReturnDocument convert ReturnDocument to mongo options.ReturnDocument
func ParseReturnDocument(r *ReturnDocument) *options.ReturnDocument {
	var mongoReturnDocument *options.ReturnDocument
	if helper.IsNotNil(r) {
		value := options.ReturnDocument(*r)
		mongoReturnDocument = &value
	}
	return mongoReturnDocument
}

// ParseFullDocument convert FullDocument to mongo options.FullDocument
func ParseFullDocument(f *FullDocument) *options.FullDocument {
	var mongoFullDocument *options.FullDocument
	if helper.IsNotNil(f) {
		value := options.FullDocument(*f)
		mongoFullDocument = &value
	}
	return mongoFullDocument
}
