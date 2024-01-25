package option

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// Aggregate represents options that can be used to configure an 'Aggregate' operation.
type Aggregate struct {
	// AllowDiskUse If true, the operation can write to temporary files in the _tmp subdirectory of the database
	// directory path on the server. The default value is false.
	AllowDiskUse *bool
	// BatchSize
	// The maximum number of documents to be included in each batch returned by the server.
	BatchSize *int32
	// BypassDocumentValidation If true, writes executed as part of the operation will opt out of document-level
	// validation on the server. This option is valid for MongoDB versions >= 3.2 and is ignored for previous server
	// versions. The default value is false. See https://www.mongodb.com/docs/manual/core/schema-validation/ for more
	// information about document validation.
	BypassDocumentValidation *bool
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only
	// valid for MongoDB versions >= 3.4. For previous server versions, the driver will return an error if this option
	// is used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
	// MaxAwaitTime The maximum amount of time that the server should wait for new documents to satisfy a cursor query.
	// This option is only valid for MongoDB versions >= 3.2 and is ignored for previous server versions.
	MaxAwaitTime *time.Duration
	// Comment A string that will be included in server logs, profiling logs, and currentOp queries to help trace the
	// operation. The default is nil, which means that no comment will be included in the logs.
	Comment *string
	// Hint The index to use for the aggregation. This should either be the index name as a string or the index
	// specification as a document. The hint does not apply to $lookup and $graphLookup aggregation stages. The driver
	// will return an error if the hint parameter is a multi-key map. The default value is nil, which means that no hint
	// will be sent.
	Hint any
	// Let Specifies parameters for the aggregate expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let any
	// Custom options to be added to aggregate expression. Key-value pairs of the BSON map should correlate with desired
	// option names and values. Values must be Marshaller. Custom options may conflict with non-custom options, and custom
	// options bypass client-side validation. Prefer using non-custom options where possible.
	Custom bson.M
}

// NewAggregate creates a new Aggregate instance.
func NewAggregate() *Aggregate {
	return &Aggregate{}
}

// SetAllowDiskUse sets value for the AllowDiskUse field.
func (a *Aggregate) SetAllowDiskUse(b bool) *Aggregate {
	a.AllowDiskUse = &b
	return a
}

// SetBatchSize sets value for the BatchSize field.
func (a *Aggregate) SetBatchSize(i int32) *Aggregate {
	a.BatchSize = &i
	return a
}

// SetBypassDocumentValidation sets value for the BypassDocumentValidation field.
func (a *Aggregate) SetBypassDocumentValidation(b bool) *Aggregate {
	a.BypassDocumentValidation = &b
	return a
}

// SetCollation sets value for the Collation field.
func (a *Aggregate) SetCollation(c *Collation) *Aggregate {
	a.Collation = c
	return a
}

// SetMaxTime sets value for the MaxTime field.
func (a *Aggregate) SetMaxTime(d time.Duration) *Aggregate {
	a.MaxTime = &d
	return a
}

// SetMaxAwaitTime sets value for the MaxAwaitTime field.
func (a *Aggregate) SetMaxAwaitTime(d time.Duration) *Aggregate {
	a.MaxAwaitTime = &d
	return a
}

// SetComment sets value for the Comment field.
func (a *Aggregate) SetComment(s string) *Aggregate {
	a.Comment = &s
	return a
}

// SetHint sets value for the Hint field.
func (a *Aggregate) SetHint(v any) *Aggregate {
	a.Hint = v
	return a
}

// SetLet sets value for the Let field.
func (a *Aggregate) SetLet(v any) *Aggregate {
	a.Let = v
	return a
}

// SetCustom sets value for the Custom field.
func (a *Aggregate) SetCustom(b bson.M) *Aggregate {
	a.Custom = b
	return a
}

// MergeAggregateByParams assembles the Aggregate object from optional parameters.
func MergeAggregateByParams(opts []*Aggregate) *Aggregate {
	result := &Aggregate{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.AllowDiskUse) {
			result.AllowDiskUse = opt.AllowDiskUse
		}
		if helper.IsNotNil(opt.BatchSize) {
			result.BatchSize = opt.BatchSize
		}
		if helper.IsNotNil(opt.Collation) {
			result.Collation = opt.Collation
		}
		if helper.IsNotNil(opt.Comment) {
			result.Comment = opt.Comment
		}
		if helper.IsNotNil(opt.Hint) {
			result.Hint = opt.Hint
		}
		if helper.IsNotNil(opt.Let) {
			result.Let = opt.Let
		}
		if helper.IsNotNil(opt.MaxTime) {
			result.MaxTime = opt.MaxTime
		}
		if helper.IsNotNil(opt.MaxAwaitTime) {
			result.MaxAwaitTime = opt.MaxAwaitTime
		}
		if helper.IsNotNil(opt.Custom) {
			result.Custom = opt.Custom
		}
	}
	return result
}
