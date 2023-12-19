package option

import "time"

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

func NewCount() Count {
	return Count{}
}

func NewEstimatedDocumentCount() EstimatedDocumentCount {
	return EstimatedDocumentCount{}
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
