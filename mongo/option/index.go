package option

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"time"
)

// Index represents options that can be used to configure a CreateOneIndex and CreateManyIndex operation.
type Index struct {
	// ExpireAfterSeconds The length of time, in seconds, for documents to remain in the collection. The default value
	// is 0, which means that documents will remain in the collection until they're explicitly deleted or the collection
	// is dropped.
	ExpireAfterSeconds *int32
	// Name The name of the index. The default value is "[field1]_[direction1]_[field2]_[direction2]...". For example, an
	// index with the specification {name: 1, age: -1} will be named "name_1_age_-1".
	Name *string
	// Sparse If true, the index will only reference documents that contain the fields specified in the index.
	// The default is false.
	Sparse *bool
	// StorageEngine Specifies the storage engine to use for the index. The value must be a document in the form
	// {<storage engine name>: <options>}. The default value is nil, which means that the default storage engine
	// will be used. This option is only applicable for MongoDB versions >= 3.0 and is ignored for previous server
	// versions.
	StorageEngine any
	// Unique If true, the collection will not accept insertion or update of documents where the index key value matches an
	// existing value in the index. The default is false.
	Unique *bool
	// Version The index version number, either 0 or 1.
	Version *int32
	// DefaultLanguage The language that determines the list of stop words and the rules for the stemmer and tokenizer.
	// This option is only applicable for text indexes and is ignored for other index types. The default value is "english".
	DefaultLanguage *string
	// LanguageOverride The name of the field in the collection's documents that contains the override language for the
	// document. This option is only applicable for text indexes and is ignored for other index types. The default value
	// is the value of the DefaultLanguage option.
	LanguageOverride *string
	// TextVersion The index version number for a text index.
	// See https://www.mongodb.com/docs/manual/core/index-text/#text-versions for information about different version numbers.
	TextVersion *int32
	// Weights A document that contains field and weight pairs. The weight is an integer ranging from 1 to 99,999, inclusive,
	// indicating the significance of the field relative to the other indexed fields in terms of the score. This option
	// is only applicable for text indexes and is ignored for other index types. The default value is nil, which means
	// that every field will have a weight of 1.
	Weights any
	// SphereVersion The index version number for a 2D sphere index.
	// See https://www.mongodb.com/docs/manual/core/2dsphere/#dsphere-v2 for information about different version numbers.
	SphereVersion *int32
	// Bits The precision of the stored geo hash value of the location data. This option only applies to 2D indexes and
	// is ignored for other index types. The value must be between 1 and 32, inclusive. The default value is 26.
	Bits *int32
	// Max The upper inclusive boundary for longitude and latitude values. This option is only applicable to 2D indexes
	// and is ignored for other index types. The default value is 180.0.
	Max *float64
	// Min The lower inclusive boundary for longitude and latitude values. This option is only applicable to 2D indexes
	// and is ignored for other index types. The default value is -180.0.
	Min *float64
	// BucketSize The number of units within which to group location values. Location values that are within BucketSize
	// units of each other will be grouped in the same bucket. This option is only applicable to geoHaystack indexes and is
	// ignored for other index types. The value must be greater than 0.
	BucketSize *int32
	// PartialFilterExpression A document that defines which collection documents the index should reference.
	// This option is only valid for MongoDB versions >= 3.2 and is ignored for previous server versions.
	PartialFilterExpression any
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// WildcardProjection A document that defines the wildcard projection for the index.
	WildcardProjection any
	// Hidden If true, the index will exist on the target collection but will not be used by the query planner
	// when executing operations. This option is only valid for MongoDB versions >= 4.4. The default value is false.
	Hidden *bool
}

// DropIndex represents options that can be used to configure a DropOneIndex and DropAllIndexes operation.
type DropIndex struct {
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
}

// ListIndexes represents options that can be used to configure a ListIndexes and ListIndexSpecifications operation.
type ListIndexes struct {
	// The maximum number of documents to be included in each batch returned by the server.
	BatchSize *int32
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
}

// NewIndex creates a new Index instance.
func NewIndex() *Index {
	return &Index{}
}

// NewDropIndex creates a new DropIndex instance.
func NewDropIndex() *DropIndex {
	return &DropIndex{}
}

// NewListIndexes creates a new ListIndexes instance.
func NewListIndexes() *ListIndexes {
	return &ListIndexes{}
}

// SetExpireAfterSeconds sets value for the ExpireAfterSeconds field.
func (i *Index) SetExpireAfterSeconds(seconds int32) *Index {
	i.ExpireAfterSeconds = &seconds
	return i
}

// SetName sets the value for the Name field.
func (i *Index) SetName(name string) *Index {
	i.Name = &name
	return i
}

// SetSparse sets the value of the Sparse field.
func (i *Index) SetSparse(sparse bool) *Index {
	i.Sparse = &sparse
	return i
}

// SetStorageEngine sets the value for the StorageEngine field.
func (i *Index) SetStorageEngine(engine any) *Index {
	i.StorageEngine = engine
	return i
}

// SetUnique sets the value for the Unique field.
func (i *Index) SetUnique(unique bool) *Index {
	i.Unique = &unique
	return i
}

// SetVersion sets the value for the Version field.
func (i *Index) SetVersion(version int32) *Index {
	i.Version = &version
	return i
}

// SetDefaultLanguage sets the value for the DefaultLanguage field.
func (i *Index) SetDefaultLanguage(language string) *Index {
	i.DefaultLanguage = &language
	return i
}

// SetLanguageOverride sets the value of the LanguageOverride field.
func (i *Index) SetLanguageOverride(override string) *Index {
	i.LanguageOverride = &override
	return i
}

// SetTextVersion sets the value for the TextVersion field.
func (i *Index) SetTextVersion(version int32) *Index {
	i.TextVersion = &version
	return i
}

// SetWeights sets the value for the Weights field.
func (i *Index) SetWeights(weights any) *Index {
	i.Weights = weights
	return i
}

// SetSphereVersion sets the value for the SphereVersion field.
func (i *Index) SetSphereVersion(version int32) *Index {
	i.SphereVersion = &version
	return i
}

// SetBits sets the value for the Bits field.
func (i *Index) SetBits(bits int32) *Index {
	i.Bits = &bits
	return i
}

// SetMax sets the value for the Max field.
func (i *Index) SetMax(max float64) *Index {
	i.Max = &max
	return i
}

// SetMin sets the value for the Min field.
func (i *Index) SetMin(min float64) *Index {
	i.Min = &min
	return i
}

// SetBucketSize sets the value for the BucketSize field
func (i *Index) SetBucketSize(bucketSize int32) *Index {
	i.BucketSize = &bucketSize
	return i
}

// SetPartialFilterExpression sets the value for the PartialFilterExpression field.
func (i *Index) SetPartialFilterExpression(expression any) *Index {
	i.PartialFilterExpression = expression
	return i
}

// SetCollation sets the value for the Collation field.
func (i *Index) SetCollation(collation *Collation) *Index {
	i.Collation = collation
	return i
}

// SetWildcardProjection sets the value for the WildcardProjection field.
func (i *Index) SetWildcardProjection(wildcardProjection any) *Index {
	i.WildcardProjection = wildcardProjection
	return i
}

// SetHidden sets the value for the Hidden field.
func (i *Index) SetHidden(hidden bool) *Index {
	i.Hidden = &hidden
	return i
}

// SetMaxTime creates a new MaxTime instance.
func (d *DropIndex) SetMaxTime(duration time.Duration) *DropIndex {
	d.MaxTime = &duration
	return d
}

// SetMaxTime creates a new MaxTime instance.
func (l *ListIndexes) SetMaxTime(duration time.Duration) *ListIndexes {
	l.MaxTime = &duration
	return l
}

// SetBatchSize creates a new BatchSize instance.
func (l *ListIndexes) SetBatchSize(i int32) *ListIndexes {
	l.BatchSize = &i
	return l
}

// MergeDropIndexByParams assembles the DropIndex object from optional parameters.
func MergeDropIndexByParams(opts []*DropIndex) *DropIndex {
	result := &DropIndex{}
	for _, opt := range opts {
		if helper.IsNotNil(opt) && helper.IsNotNil(opt.MaxTime) {
			result.MaxTime = opt.MaxTime
		}
	}
	return result
}

// MergeListIndexesByParams assembles the ListIndexes object from optional parameters.
func MergeListIndexesByParams(opts []*ListIndexes) *ListIndexes {
	result := &ListIndexes{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.MaxTime) {
			result.MaxTime = opt.MaxTime
		}
		if helper.IsNotNil(opt.BatchSize) {
			result.BatchSize = opt.BatchSize
		}
	}
	return result
}
