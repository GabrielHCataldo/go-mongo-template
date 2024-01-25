package option

import "github.com/GabrielHCataldo/go-helper/helper"

// Delete represents options that can be used to configure DeleteOne and DeleteMany operations.
type Delete struct {
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// Comment A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation. The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// Hint The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.4 will return an error
	// if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option
	// is specified. The driver will return an error if this option is specified during an unacknowledged write
	// operation. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint any
	// Let Specifies parameters for the delete expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let any
	// DisableAutoCloseSession Disable automatic closing session, if true, we automatically close session according to
	// the result, if an error occurs, we abort the transaction, otherwise, we commit the transaction.
	// default is false
	DisableAutoCloseSession *bool
	// ForceRecreateSession Force the creation of the session, if any session is still open, we close it automatically,
	// committing the transactions, and continue creating a new session.
	// default is false
	ForceRecreateSession *bool
}

// NewDelete creates a new Delete instance.
func NewDelete() *Delete {
	return &Delete{}
}

// SetCollation sets value for the Collation field.
func (d *Delete) SetCollation(c *Collation) *Delete {
	d.Collation = c
	return d
}

// SetComment sets value for the Comment field.
func (d *Delete) SetComment(a any) *Delete {
	d.Comment = a
	return d
}

// SetHint sets value for the Hint field.
func (d *Delete) SetHint(a any) *Delete {
	d.Hint = a
	return d
}

// SetLet sets value for the Let field.
func (d *Delete) SetLet(a any) *Delete {
	d.Let = a
	return d
}

// SetDisableAutoCloseSession sets value for the DisableAutoCloseSession field.
func (d *Delete) SetDisableAutoCloseSession(b bool) *Delete {
	d.DisableAutoCloseSession = &b
	return d
}

// SetForceRecreateSession sets value for the ForceRecreateSession field.
func (d *Delete) SetForceRecreateSession(b bool) *Delete {
	d.ForceRecreateSession = &b
	return d
}

// MergeDeleteByParams assembles the Delete object from optional parameters.
func MergeDeleteByParams(opts []*Delete) *Delete {
	result := &Delete{}
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
		if helper.IsNotNil(opt.Hint) {
			result.Hint = opt.Hint
		}
		if helper.IsNotNil(opt.Let) {
			result.Let = opt.Let
		}
		if helper.IsNotNil(opt.DisableAutoCloseSession) {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
		}
		if helper.IsNotNil(opt.ForceRecreateSession) {
			result.ForceRecreateSession = opt.ForceRecreateSession
		}
	}
	if helper.IsNil(result.DisableAutoCloseSession) {
		result.DisableAutoCloseSession = helper.ConvertToPointer(false)
	}
	if helper.IsNil(result.ForceRecreateSession) {
		result.ForceRecreateSession = helper.ConvertToPointer(false)
	}
	return result
}
