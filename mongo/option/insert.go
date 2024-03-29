package option

import "github.com/GabrielHCataldo/go-helper/helper"

// InsertOne represents options that can be used to configure a 'InsertOne' operation.
type InsertOne struct {
	// BypassDocumentValidation If true, writes executed as part of the operation will opt out of document-level
	// validation on the server. This option is valid for MongoDB versions >= 3.2 and is ignored for previous server
	// versions. The default value is false.
	// See https://www.mongodb.com/docs/manual/core/schema-validation/ for more information about document validation.
	BypassDocumentValidation *bool
	// Comment A string or document that will be included in server logs, profiling logs, and currentOp queries to help
	// trace the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// DisableAutoRollbackSession disable auto rollback if an error occurs.
	DisableAutoRollbackSession *bool
	// DisableAutoCloseSession Disable automatic closing session, if true, we automatically close session according to
	// the result, if an error occurs, we abort the transaction, otherwise, we commit the transaction.
	// default is false
	DisableAutoCloseSession *bool
	// ForceRecreateSession Force the creation of the session, if any session is still open, we close it automatically,
	// aborting all open transactions, and continue creating a new session.
	// default is false
	ForceRecreateSession *bool
}

// InsertMany represents options that can be used to configure a 'InsertMany' operation.
type InsertMany struct {
	// BypassDocumentValidation If true, writes executed as part of the operation will opt out of document-level
	// validation on the server. This option is valid for MongoDB versions >= 3.2 and is ignored for previous server
	// versions. The default value is false.
	// See https://www.mongodb.com/docs/manual/core/schema-validation/ for more information about document validation.
	BypassDocumentValidation *bool
	// Comment A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
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

// NewInsertOne creates a new InsertOne instance.
func NewInsertOne() *InsertOne {
	return &InsertOne{}
}

// NewInsertMany creates a new InsertMany instance.
func NewInsertMany() *InsertMany {
	return &InsertMany{}
}

// SetBypassDocumentValidation sets value for the BypassDocumentValidation field.
func (i *InsertOne) SetBypassDocumentValidation(b bool) *InsertOne {
	i.BypassDocumentValidation = &b
	return i
}

// SetComment sets value for the Comment field.
func (i *InsertOne) SetComment(a any) *InsertOne {
	i.Comment = a
	return i
}

// SetDisableAutoRollbackSession creates a new DisableAutoRollbackSession instance.
func (i *InsertOne) SetDisableAutoRollbackSession(b bool) *InsertOne {
	i.DisableAutoRollbackSession = &b
	return i
}

// SetDisableAutoCloseSession creates a new DisableAutoCloseSession instance.
func (i *InsertOne) SetDisableAutoCloseSession(b bool) *InsertOne {
	i.DisableAutoCloseSession = &b
	return i
}

// SetForceRecreateSession sets value for the ForceRecreateSession field.
func (i *InsertOne) SetForceRecreateSession(b bool) *InsertOne {
	i.ForceRecreateSession = &b
	return i
}

// SetBypassDocumentValidation sets value for the BypassDocumentValidation field.
func (i *InsertMany) SetBypassDocumentValidation(b bool) *InsertMany {
	i.BypassDocumentValidation = &b
	return i
}

// SetComment sets value for the Comment field.
func (i *InsertMany) SetComment(a any) *InsertMany {
	i.Comment = a
	return i
}

// SetDisableAutoRollback sets value for the DisableAutoRollbackSession field.
func (i *InsertMany) SetDisableAutoRollback(b bool) *InsertMany {
	i.DisableAutoRollbackSession = &b
	return i
}

// SetDisableAutoCloseSession creates a new DisableAutoCloseSession instance.
func (i *InsertMany) SetDisableAutoCloseSession(b bool) *InsertMany {
	i.DisableAutoCloseSession = &b
	return i
}

// SetForceRecreateSession sets value for the ForceRecreateSession field.
func (i *InsertMany) SetForceRecreateSession(b bool) *InsertMany {
	i.ForceRecreateSession = &b
	return i
}

// MergeInsertOneByParams assembles the InsertOne object from optional parameters.
func MergeInsertOneByParams(opts []*InsertOne, global *Global) *InsertOne {
	result := &InsertOne{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.BypassDocumentValidation) {
			result.BypassDocumentValidation = opt.BypassDocumentValidation
		}
		if helper.IsNotNil(opt.Comment) {
			result.Comment = opt.Comment
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

// MergeInsertManyByParams assembles the InsertMany object from optional parameters.
func MergeInsertManyByParams(opts []*InsertMany, global *Global) *InsertMany {
	result := &InsertMany{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.ForceRecreateSession) {
			result.ForceRecreateSession = opt.ForceRecreateSession
		}
		if helper.IsNotNil(opt.BypassDocumentValidation) {
			result.BypassDocumentValidation = opt.BypassDocumentValidation
		}
		if helper.IsNotNil(opt.Comment) {
			result.Comment = opt.Comment
		}
		if helper.IsNotNil(opt.DisableAutoRollbackSession) {
			result.DisableAutoRollbackSession = opt.DisableAutoRollbackSession
		}
		if helper.IsNotNil(opt.DisableAutoCloseSession) {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
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
