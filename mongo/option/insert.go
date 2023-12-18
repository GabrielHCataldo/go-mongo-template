package option

type InsertOne struct {
	// If true, writes executed as part of the operation will opt out of document-level validation on the server. This
	// option is valid for MongoDB versions >= 3.2 and is ignored for previous server versions. The default value is
	// false. See https://www.mongodb.com/docs/manual/core/schema-validation/ for more information about document
	// validation.
	BypassDocumentValidation bool
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment                     any
	DisableAutoCloseTransaction bool
}

type InsertMany struct {
	// If true, writes executed as part of the operation will opt out of document-level validation on the server. This
	// option is valid for MongoDB versions >= 3.2 and is ignored for previous server versions. The default value is
	// false. See https://www.mongodb.com/docs/manual/core/schema-validation/ for more information about document
	// validation.
	BypassDocumentValidation bool
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment                     any
	DisableAutoRollback         bool
	DisableAutoCloseTransaction bool
}

func NewInsertOne() InsertOne {
	return InsertOne{}
}

func NewInsertMany() InsertMany {
	return InsertMany{}
}

func (i InsertOne) SetBypassDocumentValidation(b bool) InsertOne {
	i.BypassDocumentValidation = b
	return i
}

func (i InsertOne) SetComment(a any) InsertOne {
	i.Comment = a
	return i
}

func (i InsertOne) SetDisableAutoCloseTransaction(b bool) InsertOne {
	i.DisableAutoCloseTransaction = b
	return i
}

func (i InsertMany) SetBypassDocumentValidation(b bool) InsertMany {
	i.BypassDocumentValidation = b
	return i
}

func (i InsertMany) SetComment(a any) InsertMany {
	i.Comment = a
	return i
}

func (i InsertMany) SetDisableAutoRollback(b bool) InsertMany {
	i.DisableAutoRollback = b
	return i
}

func (i InsertMany) SetDisableAutoCloseTransaction(b bool) InsertMany {
	i.DisableAutoCloseTransaction = b
	return i
}

func GetInsertOneOptionByParams(opts []InsertOne) InsertOne {
	result := InsertOne{}
	for _, opt := range opts {
		if opt.BypassDocumentValidation {
			result.BypassDocumentValidation = opt.BypassDocumentValidation
		}
		if opt.Comment != nil {
			result.Comment = opt.Comment
		}
		if opt.DisableAutoCloseTransaction {
			result.DisableAutoCloseTransaction = opt.DisableAutoCloseTransaction
		}
	}
	return result
}

func GetInsertManyOptionByParams(opts []InsertMany) InsertMany {
	result := InsertMany{}
	for _, opt := range opts {
		if opt.BypassDocumentValidation {
			result.BypassDocumentValidation = opt.BypassDocumentValidation
		}
		if opt.Comment != nil {
			result.Comment = opt.Comment
		}
		if opt.DisableAutoRollback {
			result.DisableAutoRollback = opt.DisableAutoRollback
		}
		if opt.DisableAutoCloseTransaction {
			result.DisableAutoCloseTransaction = opt.DisableAutoCloseTransaction
		}
	}
	return result
}
