package option

import "github.com/GabrielHCataldo/go-helper/helper"

// Replace represents options that can be used to configure a 'ReplaceOne' operation.
type Replace struct {
	// BypassDocumentValidation If true, writes executed as part of the operation will opt out of document-level
	// validation on the server. This option is valid for MongoDB versions >= 3.2 and is ignored for previous server
	// versions. The default value is false.
	// See https://www.mongodb.com/docs/manual/core/schema-validation/ for more information about document validation.
	BypassDocumentValidation *bool
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// Comment A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// Hint The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.2. Server versions >= 3.4 will return an error
	// if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option
	// is specified. The driver will return an error if this option is specified during an unacknowledged write
	// operation. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint any
	// Upsert If true, a new document will be inserted if the filter does not match any documents in the collection. The
	// default value is false.
	Upsert *bool
	// Let Specifies parameters for the update expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let any
	// DisableAutoRollbackSession disable auto rollback if an error occurs.
	DisableAutoRollbackSession *bool
	// DisableAutoCloseSession Disable automatic closing session, if true, we automatically close session according to
	// the result, if an error occurs, we abort the transaction, otherwise, we commit the transaction.
	// default is false
	DisableAutoCloseSession *bool
	// ForceRecreateSession Force the creation of the session, if any session is still open, we close it automatically,
	// committing the transactions, and continue creating a new session.
	// default is false
	ForceRecreateSession *bool
}

// NewReplace creates a new Replace instance.
func NewReplace() *Replace {
	return &Replace{}
}

// SetBypassDocumentValidation sets value for the BypassDocumentValidation field.
func (r *Replace) SetBypassDocumentValidation(b bool) *Replace {
	r.BypassDocumentValidation = &b
	return r
}

// SetCollation sets value for the Collation field.
func (r *Replace) SetCollation(c *Collation) *Replace {
	r.Collation = c
	return r
}

// SetComment sets value for the Comment field.
func (r *Replace) SetComment(a any) *Replace {
	r.Comment = a
	return r
}

// SetHint sets value for the Hint field.
func (r *Replace) SetHint(a any) *Replace {
	r.Hint = a
	return r
}

// SetLet sets value for the Let field.
func (r *Replace) SetLet(a any) *Replace {
	r.Let = a
	return r
}

// SetDisableAutoRollbackSession creates a new DisableAutoRollbackSession instance.
func (r *Replace) SetDisableAutoRollbackSession(b bool) *Replace {
	r.DisableAutoRollbackSession = &b
	return r
}

// SetDisableAutoCloseSession creates a new DisableAutoCloseSession instance.
func (r *Replace) SetDisableAutoCloseSession(b bool) *Replace {
	r.DisableAutoCloseSession = &b
	return r
}

// SetForceRecreateSession sets value for the ForceRecreateSession field.
func (r *Replace) SetForceRecreateSession(b bool) *Replace {
	r.ForceRecreateSession = &b
	return r
}

// MergeReplaceByParams assembles the Replace object from optional parameters.
func MergeReplaceByParams(opts []*Replace, global *Global) *Replace {
	result := &Replace{}
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
		if helper.IsNotNil(opt.DisableAutoRollbackSession) {
			result.DisableAutoRollbackSession = opt.DisableAutoRollbackSession
		}
		if helper.IsNotNil(opt.DisableAutoCloseSession) {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
		}
		if helper.IsNotNil(opt.ForceRecreateSession) {
			result.ForceRecreateSession = opt.ForceRecreateSession
		}
	}
	if helper.IsNil(result.BypassDocumentValidation) {
		result.BypassDocumentValidation = helper.ConvertToPointer(global.BypassDocumentValidation)
	}
	if helper.IsNil(result.Comment) {
		result.Comment = global.Comment
	}
	if helper.IsNil(result.DisableAutoRollbackSession) {
		result.DisableAutoRollbackSession = helper.ConvertToPointer(global.DisableAutoRollbackSession)
	}
	if helper.IsNil(result.DisableAutoCloseSession) {
		result.DisableAutoCloseSession = helper.ConvertToPointer(global.DisableAutoCloseSession)
	}
	if helper.IsNil(result.ForceRecreateSession) {
		result.ForceRecreateSession = helper.ConvertToPointer(global.ForceRecreateSession)
	}
	return result
}
