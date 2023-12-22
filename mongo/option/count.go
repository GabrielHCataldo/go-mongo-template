package option

import "time"

// Count represents options that can be used to configure a 'CountDocuments' operation.
type Count struct {
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// Comment A string or document that will be included in server logs, profiling logs, and currentOp queries to help
	// trace the operation. The default is nil, which means that no comment will be included in the logs.
	Comment *string
	// Hint The index to use for the aggregation. This should either be the index name as a string or the index specification
	// as a document. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint any
	// Limit The maximum number of documents to count. The default value is 0, which means that there is no limit and all
	// documents matching the filter will be counted.
	Limit *int64
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
	// Skip
	// The number of documents to skip before counting. The default value is 0.
	Skip *int64
}

// EstimatedDocumentCount represents options that can be used to configure an 'EstimatedDocumentCount' operation.
type EstimatedDocumentCount struct {
	// Comment A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default is nil, which means that no comment will be included in the logs.
	Comment any
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
}

// NewCount creates a new Count instance.
func NewCount() Count {
	return Count{}
}

// NewEstimatedDocumentCount creates a new EstimatedDocumentCount instance.
func NewEstimatedDocumentCount() EstimatedDocumentCount {
	return EstimatedDocumentCount{}
}

// SetCollation sets value for the Collation field.
func (c Count) SetCollation(collation *Collation) Count {
	c.Collation = collation
	return c
}

// SetComment sets value for the Comment field.
func (c Count) SetComment(comment string) Count {
	c.Comment = &comment
	return c
}

// SetHint sets value for the Hint field.
func (c Count) SetHint(a any) Count {
	c.Hint = a
	return c
}

// SetLimit sets value for the Limit field.
func (c Count) SetLimit(i int64) Count {
	c.Limit = &i
	return c
}

// SetMaxTime sets value for the MaxTime field.
func (c Count) SetMaxTime(d time.Duration) Count {
	c.MaxTime = &d
	return c
}

// SetSkip sets value for the Skip field.
func (c Count) SetSkip(i int64) Count {
	c.Skip = &i
	return c
}

// SetMaxTime sets value for the MaxTime field.
func (e EstimatedDocumentCount) SetMaxTime(d time.Duration) EstimatedDocumentCount {
	e.MaxTime = &d
	return e
}

// SetComment sets value for the Comment field.
func (e EstimatedDocumentCount) SetComment(comment any) EstimatedDocumentCount {
	e.Comment = comment
	return e
}

// GetCountOptionByParams assembles the Count object from optional parameters.
func GetCountOptionByParams(opts []Count) Count {
	result := Count{}
	for _, opt := range opts {
		if opt.Collation != nil {
			result.Collation = opt.Collation
		}
		if opt.Comment != nil {
			result.Comment = opt.Comment
		}
		if opt.Hint != nil {
			result.Hint = opt.Hint
		}
		if opt.MaxTime != nil {
			result.MaxTime = opt.MaxTime
		}
		if opt.Limit != nil {
			result.Limit = opt.Limit
		}
		if opt.Skip != nil {
			result.Skip = opt.Skip
		}
	}
	return result
}

// GetEstimatedDocumentCountOptionByParams assembles the EstimatedDocumentCount object from optional parameters.
func GetEstimatedDocumentCountOptionByParams(opts []EstimatedDocumentCount) EstimatedDocumentCount {
	result := EstimatedDocumentCount{}
	for _, opt := range opts {
		if opt.Comment != nil {
			result.Comment = opt.Comment
		}
		if opt.MaxTime != nil {
			result.MaxTime = opt.MaxTime
		}
	}
	return result
}
