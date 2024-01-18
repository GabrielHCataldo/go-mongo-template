package option

import "time"

// Exists represents options that can be used to configure an 'Exists' operation.
type Exists struct {
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default is nil, which means that no comment will be included in the logs.
	Comment *string
	// The index to use for the aggregation. This should either be the index name as a string or the index specification
	// as a document. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint any
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
}

// NewExists creates a new Exists instance.
func NewExists() *Exists {
	return &Exists{}
}

// SetCollation creates a new Collation instance.
func (e *Exists) SetCollation(collation *Collation) *Exists {
	e.Collation = collation
	return e
}

// SetComment creates a new Comment instance.
func (e *Exists) SetComment(comment string) *Exists {
	e.Comment = &comment
	return e
}

// SetHint creates a new Hint instance.
func (e *Exists) SetHint(a any) *Exists {
	e.Hint = a
	return e
}

// SetMaxTime creates a new MaxTime instance.
func (e *Exists) SetMaxTime(d time.Duration) *Exists {
	e.MaxTime = &d
	return e
}

// GetExistsOptionByParams assembles the Exists object from optional parameters.
func GetExistsOptionByParams(opts []*Exists) *Exists {
	result := &Exists{}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
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
	}
	return result
}
