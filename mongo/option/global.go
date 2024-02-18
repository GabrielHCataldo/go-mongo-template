package option

// Global represents options that can be used for all operations as a default form, it is important to highlight that it
// will not overwrite operation options.
type Global struct {
	// BypassDocumentValidation If true, writes executed as part of the operation will opt out of document-level
	// validation on the server. This option is valid for MongoDB versions >= 3.2 and is ignored for previous server
	// versions. The default value is false.
	// See https://www.mongodb.com/docs/manual/core/schema-validation/ for more information about document validation.
	BypassDocumentValidation bool
	// Comment A string or document that will be included in server logs, profiling logs, and currentOp queries to help
	// trace the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// DisableAutoRollbackSession disable auto rollback if an error occurs.
	DisableAutoRollbackSession bool
	// DisableAutoCloseSession Disable automatic closing session, if true, we automatically close session according to
	// the result, if an error occurs, we abort the transaction, otherwise, we commit the transaction.
	// default is false
	DisableAutoCloseSession bool
	// ForceRecreateSession Force the creation of the session, if any session is still open, we close it automatically,
	// aborting all open transactions, and continue creating a new session.
	// default is false
	ForceRecreateSession bool
}
