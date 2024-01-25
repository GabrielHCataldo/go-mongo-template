package option

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"time"
)

// Distinct represents options that can be used to configure a Distinct operation.
type Distinct struct {
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation. The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
}

// NewDistinct creates a new Distinct instance.
func NewDistinct() *Distinct {
	return &Distinct{}
}

// SetCollation sets value for the Collation field.
func (d *Distinct) SetCollation(c *Collation) *Distinct {
	d.Collation = c
	return d
}

// SetMaxTime sets value for the MaxTime field.
func (d *Distinct) SetMaxTime(duration time.Duration) *Distinct {
	d.MaxTime = &duration
	return d
}

// SetComment sets value for the Comment field.
func (d *Distinct) SetComment(comment any) *Distinct {
	d.Comment = comment
	return d
}

// MergeDistinctByParams assembles the Distinct object from optional parameters.
func MergeDistinctByParams(opts []*Distinct) *Distinct {
	result := &Distinct{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.Collation) {
			result.Collation = opt.Collation
		}
		if helper.IsNotNil(opt.Comment) {
			result.Comment = opt.Comment
		}
		if helper.IsNotNil(opt.MaxTime) {
			result.MaxTime = opt.MaxTime
		}
	}
	return result
}
