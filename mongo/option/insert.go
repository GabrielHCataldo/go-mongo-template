package option

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
	// DisableAutoCloseSession Disable automatic closing session, if true, we automatically close session according to
	// the result, if an error occurs, we abort the transaction, otherwise, we commit the transaction.
	DisableAutoCloseSession bool
	// ForceRecreateSession Force the creation of the session, if any session is still open, we close it automatically,
	// aborting all open transactions, and continue creating a new session.
	ForceRecreateSession bool
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
	DisableAutoRollbackSession bool
	// DisableAutoCloseSession Disable automatic closing session, if true, we automatically close session according to
	// the result, if an error occurs, we abort the transaction, otherwise, we commit the transaction.
	DisableAutoCloseSession bool
	// ForceRecreateSession Force the creation of the session, if any session is still open, we close it automatically,
	// committing the transactions, and continue creating a new session.
	ForceRecreateSession bool
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

// SetDisableAutoCloseSession creates a new DisableAutoCloseSession instance.
func (i *InsertOne) SetDisableAutoCloseSession(b bool) *InsertOne {
	i.DisableAutoCloseSession = b
	return i
}

// SetForceRecreateSession sets value for the ForceRecreateSession field.
func (i *InsertOne) SetForceRecreateSession(b bool) *InsertOne {
	i.ForceRecreateSession = b
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
	i.DisableAutoRollbackSession = b
	return i
}

// SetDisableAutoCloseSession creates a new DisableAutoCloseSession instance.
func (i *InsertMany) SetDisableAutoCloseSession(b bool) *InsertMany {
	i.DisableAutoCloseSession = b
	return i
}

// SetForceRecreateSession sets value for the ForceRecreateSession field.
func (i *InsertMany) SetForceRecreateSession(b bool) *InsertMany {
	i.ForceRecreateSession = b
	return i
}

// GetInsertOneOptionByParams assembles the InsertOne object from optional parameters.
func GetInsertOneOptionByParams(opts []*InsertOne) *InsertOne {
	result := &InsertOne{}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.BypassDocumentValidation != nil {
			result.BypassDocumentValidation = opt.BypassDocumentValidation
		}
		if opt.Comment != nil {
			result.Comment = opt.Comment
		}
		if opt.DisableAutoCloseSession {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
		}
		if opt.ForceRecreateSession {
			result.ForceRecreateSession = opt.ForceRecreateSession
		}
	}
	return result
}

// GetInsertManyOptionByParams assembles the InsertMany object from optional parameters.
func GetInsertManyOptionByParams(opts []*InsertMany) *InsertMany {
	result := &InsertMany{}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		if opt.ForceRecreateSession {
			result.ForceRecreateSession = opt.ForceRecreateSession
		}
		if opt.BypassDocumentValidation != nil {
			result.BypassDocumentValidation = opt.BypassDocumentValidation
		}
		if opt.Comment != nil {
			result.Comment = opt.Comment
		}
		if opt.DisableAutoRollbackSession {
			result.DisableAutoRollbackSession = opt.DisableAutoRollbackSession
		}
		if opt.DisableAutoCloseSession {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
		}
	}
	return result
}
