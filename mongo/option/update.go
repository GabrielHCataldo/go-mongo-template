package option

import "github.com/GabrielHCataldo/go-helper/helper"

// Update represents options that can be used to configure a 'UpdateOne' ,'UpdateMany'  or 'UpdateOneById'  operation.
type Update struct {
	// ArrayFilters A set of filters specifying to which array elements an update should apply. This option is only valid for MongoDB
	// versions >= 3.6. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the update will apply to all array elements.
	ArrayFilters *ArrayFilters
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
	// DisableAutoCloseSession Disable automatic closing session, if true, we automatically close session according to
	// the result, if an error occurs, we abort the transaction, otherwise, we commit the transaction.
	// default is false
	DisableAutoCloseSession *bool
	// ForceRecreateSession Force the creation of the session, if any session is still open, we close it automatically,
	// committing the transactions, and continue creating a new session.
	// default is false
	ForceRecreateSession *bool
}

// NewUpdate creates a new Update instance.
func NewUpdate() *Update {
	return &Update{}
}

// SetBypassDocumentValidation sets value for the BypassDocumentValidation field.
func (u *Update) SetBypassDocumentValidation(b bool) *Update {
	u.BypassDocumentValidation = &b
	return u
}

// SetArrayFilters sets value for the ArrayFilters field.
func (u *Update) SetArrayFilters(a *ArrayFilters) *Update {
	u.ArrayFilters = a
	return u
}

// SetCollation sets value for the Collation field.
func (u *Update) SetCollation(c *Collation) *Update {
	u.Collation = c
	return u
}

// SetComment sets value for the Comment field.
func (u *Update) SetComment(a any) *Update {
	u.Comment = a
	return u
}

// SetHint sets value for the Hint field.
func (u *Update) SetHint(a any) *Update {
	u.Hint = a
	return u
}

// SetLet sets value for the Let field.
func (u *Update) SetLet(a any) *Update {
	u.Let = a
	return u
}

// SetDisableAutoCloseSession creates a new DisableAutoCloseSession instance.
func (u *Update) SetDisableAutoCloseSession(b bool) *Update {
	u.DisableAutoCloseSession = &b
	return u
}

// SetForceRecreateSession sets value for the ForceRecreateSession field.
func (u *Update) SetForceRecreateSession(b bool) *Update {
	u.ForceRecreateSession = &b
	return u
}

// SetUpsert creates a new Upsert instance.
func (u *Update) SetUpsert(b bool) *Update {
	u.Upsert = &b
	return u
}

// MergeUpdateByParams assembles the Update object from optional parameters.
func MergeUpdateByParams(opts []*Update) *Update {
	result := &Update{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.ArrayFilters) {
			result.ArrayFilters = opt.ArrayFilters
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
