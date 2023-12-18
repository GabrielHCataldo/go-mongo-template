package option

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Aggregate struct {
	// If true, the operation can write to temporary files in the _tmp subdirectory of the database directory path on
	// the server. The default value is false.
	AllowDiskUse bool
	// The maximum number of documents to be included in each batch returned by the server.
	BatchSize int32
	// If true, writes executed as part of the operation will opt out of document-level validation on the server. This
	// option is valid for MongoDB versions >= 3.2 and is ignored for previous server versions. The default value is
	// false. See https://www.mongodb.com/docs/manual/core/schema-validation/ for more information about document
	// validation.
	BypassDocumentValidation bool
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used
	// in its place to control the amount of time that a single operation can run before returning an error. MaxTime
	// is ignored if Timeout is set on the client.
	MaxTime time.Duration
	// The maximum amount of time that the server should wait for new documents to satisfy a cursor query.
	// This option is only valid for MongoDB versions >= 3.2 and is ignored for previous server versions.
	MaxAwaitTime time.Duration
	// A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is nil, which means that no comment will be included in the logs.
	Comment string
	// The index to use for the aggregation. This should either be the index name as a string or the index specification
	// as a document. The hint does not apply to $lookup and $graphLookup aggregation stages. The driver will return an
	// error if the hint parameter is a multi-key map. The default value is nil, which means that no hint will be sent.
	Hint any
	// Specifies parameters for the aggregate expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let any
	// Custom options to be added to aggregate expression. Key-value pairs of the BSON map should correlate with desired
	// option names and values. Values must be Marshaller. Custom options may conflict with non-custom options, and custom
	// options bypass client-side validation. Prefer using non-custom options where possible.
	Custom bson.M
}

type Count struct {
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default is nil, which means that no comment will be included in the logs.
	Comment string
	// The index to use for the aggregation. This should either be the index name as a string or the index specification
	// as a document. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint any
	// The maximum number of documents to count. The default value is 0, which means that there is no limit and all
	// documents matching the filter will be counted.
	Limit int64
	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there is
	// no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime time.Duration
	// The number of documents to skip before counting. The default value is 0.
	Skip int64
}

type EstimatedDocumentCount struct {
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default is nil, which means that no comment will be included in the logs.
	Comment any
	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used
	// in its place to control the amount of time that a single operation can run before returning an error. MaxTime
	// is ignored if Timeout is set on the client.
	MaxTime time.Duration
}

type Distinct struct {
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation. The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be
	// used in its place to control the amount of time that a single operation can run before returning an error.
	// MaxTime is ignored if Timeout is set on the client.
	MaxTime time.Duration
}

func NewAggregate() Aggregate {
	return Aggregate{}
}

func NewCount() Count {
	return Count{}
}

func NewEstimatedDocumentCount() EstimatedDocumentCount {
	return EstimatedDocumentCount{}
}

func NewDistinct() Distinct {
	return Distinct{}
}

func (a Aggregate) SetAllowDiskUse(b bool) Aggregate {
	a.AllowDiskUse = b
	return a
}

func (a Aggregate) SetBatchSize(i int32) Aggregate {
	a.BatchSize = i
	return a
}

func (a Aggregate) SetBypassDocumentValidation(b bool) Aggregate {
	a.BypassDocumentValidation = b
	return a
}

func (a Aggregate) SetCollation(c *Collation) Aggregate {
	a.Collation = c
	return a
}

func (a Aggregate) SetMaxTime(d time.Duration) Aggregate {
	a.MaxTime = d
	return a
}

func (a Aggregate) SetMaxAwaitTime(d time.Duration) Aggregate {
	a.MaxAwaitTime = d
	return a
}

func (a Aggregate) SetComment(s string) Aggregate {
	a.Comment = s
	return a
}

func (a Aggregate) SetHint(v any) Aggregate {
	a.Hint = v
	return a
}

func (a Aggregate) SetLet(v any) Aggregate {
	a.Let = v
	return a
}

func (a Aggregate) SetCustom(b bson.M) Aggregate {
	a.Custom = b
	return a
}

func (c Count) SetCollation(collation *Collation) Count {
	c.Collation = collation
	return c
}

func (c Count) SetComment(comment string) Count {
	c.Comment = comment
	return c
}

func (c Count) SetHint(a any) Count {
	c.Hint = a
	return c
}

func (c Count) SetLimit(i int64) Count {
	c.Limit = i
	return c
}

func (c Count) SetMaxTime(d time.Duration) Count {
	c.MaxTime = d
	return c
}

func (c Count) SetSkip(i int64) Count {
	c.Skip = i
	return c
}

func (e EstimatedDocumentCount) SetMaxTime(d time.Duration) EstimatedDocumentCount {
	e.MaxTime = d
	return e
}

func (e EstimatedDocumentCount) SetComment(comment any) EstimatedDocumentCount {
	e.Comment = comment
	return e
}

func (d Distinct) SetCollation(c *Collation) Distinct {
	d.Collation = c
	return d
}

func (d Distinct) SetMaxTime(duration time.Duration) Distinct {
	d.MaxTime = duration
	return d
}

func (d Distinct) SetComment(comment any) Distinct {
	d.Comment = comment
	return d
}

func GetAggregateOptionByParams(opts []Aggregate) Aggregate {
	result := Aggregate{}
	for _, opt := range opts {
		if opt.AllowDiskUse {
			result.AllowDiskUse = opt.AllowDiskUse
		}
		if opt.BatchSize != 0 {
			result.BatchSize = opt.BatchSize
		}
		if opt.Collation != nil {
			result.Collation = opt.Collation
		}
		if len(opt.Comment) != 0 {
			result.Comment = opt.Comment
		}
		if opt.Hint != nil {
			result.Hint = opt.Hint
		}
		if opt.Let != nil {
			result.Let = opt.Let
		}
		if opt.MaxTime > 0 {
			result.MaxTime = opt.MaxTime
		}
		if opt.MaxAwaitTime > 0 {
			result.MaxAwaitTime = opt.MaxAwaitTime
		}
		if opt.Custom != nil {
			result.Custom = opt.Custom
		}
	}
	return result
}

func GetCountOptionByParams(opts []Count) Count {
	result := Count{}
	for _, opt := range opts {
		if opt.Collation != nil {
			result.Collation = opt.Collation
		}
		if len(opt.Comment) != 0 {
			result.Comment = opt.Comment
		}
		if opt.Hint != nil {
			result.Hint = opt.Hint
		}
		if opt.MaxTime > 0 {
			result.MaxTime = opt.MaxTime
		}
		if opt.Limit > 0 {
			result.Limit = opt.Limit
		}
		if opt.Skip > 0 {
			result.Skip = opt.Skip
		}
	}
	return result
}

func GetEstimatedDocumentCountOptionByParams(opts []EstimatedDocumentCount) EstimatedDocumentCount {
	result := EstimatedDocumentCount{}
	for _, opt := range opts {
		if opt.Comment != nil {
			result.Comment = opt.Comment
		}
		if opt.MaxTime > 0 {
			result.MaxTime = opt.MaxTime
		}
	}
	return result
}

func GetDistinctOptionByParams(opts []Distinct) Distinct {
	result := Distinct{}
	for _, opt := range opts {
		if opt.Collation != nil {
			result.Collation = opt.Collation
		}
		if opt.Comment != nil {
			result.Comment = opt.Comment
		}
		if opt.MaxTime > 0 {
			result.MaxTime = opt.MaxTime
		}
	}
	return result
}
