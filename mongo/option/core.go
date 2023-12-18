package option

import (
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collation allows users to specify language-specific rules for string comparison, such as
// rules for letter case and accent marks.
type Collation struct {
	Locale          string `bson:",omitempty"` // The locale
	CaseLevel       bool   `bson:",omitempty"` // The case level
	CaseFirst       string `bson:",omitempty"` // The case ordering
	Strength        int    `bson:",omitempty"` // The number of comparison levels to use
	NumericOrdering bool   `bson:",omitempty"` // Whether to order numbers based on numerical order and not collation order
	Alternate       string `bson:",omitempty"` // Whether spaces and punctuation are considered base characters
	MaxVariable     string `bson:",omitempty"` // Which characters are affected by alternate: "shifted"
	Normalization   bool   `bson:",omitempty"` // Causes text to be normalized into Unicode NFD
	Backwards       bool   `bson:",omitempty"` // Causes secondary differences to be considered in reverse order, as it is done in the French language
}

// ArrayFilters is used to hold filters for the array filters CRUD option. If a registry is nil, bson.DefaultRegistry
// will be used when converting the filter interfaces to BSON.
type ArrayFilters struct {
	// Registry is the registry to use for converting filters. Defaults to bson.DefaultRegistry.
	//
	// Deprecated: Marshaling ArrayFilters to BSON will not be supported in Go Driver 2.0.
	Registry *bsoncodec.Registry
	// The filters to apply
	Filters []interface{}
}

func ParseCollationMongoOptions(c *Collation) *options.Collation {
	if c == nil {
		return nil
	}
	return &options.Collation{
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

func ParseArrayFiltersMongoOptions(a *ArrayFilters) *options.ArrayFilters {
	if a == nil {
		return nil
	}
	return &options.ArrayFilters{
		Registry: a.Registry,
		Filters:  a.Filters,
	}
}

func ParseCursorType(c CursorType) *options.CursorType {
	if !c.IsEnumValid() {
		return nil
	}
	result := options.CursorType(c)
	return &result
}

func ParseReturnDocument(c ReturnDocument) *options.ReturnDocument {
	if !c.IsEnumValid() {
		return nil
	}
	result := options.ReturnDocument(c)
	return &result
}
