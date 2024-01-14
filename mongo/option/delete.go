package option

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
	DisableAutoCloseSession bool
	// DisableAutoRollbackSession disable auto rollback if an error occurs.
	DisableAutoRollbackSession bool
	// ForceRecreateSession Force the creation of the session, if any session is still open, we close it automatically,
	// committing the transactions, and continue creating a new session.
	ForceRecreateSession bool
}

// NewDelete creates a new Delete instance.
func NewDelete() Delete {
	return Delete{}
}

// SetCollation sets value for the Collation field.
func (d Delete) SetCollation(c *Collation) Delete {
	d.Collation = c
	return d
}

// SetComment sets value for the Comment field.
func (d Delete) SetComment(a any) Delete {
	d.Comment = a
	return d
}

// SetHint sets value for the Hint field.
func (d Delete) SetHint(a any) Delete {
	d.Hint = a
	return d
}

// SetLet sets value for the Let field.
func (d Delete) SetLet(a any) Delete {
	d.Let = a
	return d
}

// SetDisableAutoCloseSession sets value for the DisableAutoCloseSession field.
func (d Delete) SetDisableAutoCloseSession(b bool) Delete {
	d.DisableAutoCloseSession = b
	return d
}

// SetForceRecreateSession sets value for the ForceRecreateSession field.
func (d Delete) SetForceRecreateSession(b bool) Delete {
	d.ForceRecreateSession = b
	return d
}

// SetDisableAutoRollbackSession sets value for the DisableAutoRollbackSession field.
func (d Delete) SetDisableAutoRollbackSession(b bool) Delete {
	d.DisableAutoRollbackSession = b
	return d
}

// GetDeleteOptionByParams assembles the Delete object from optional parameters.
func GetDeleteOptionByParams(opts []Delete) Delete {
	result := Delete{}
	for _, opt := range opts {
		if opt.Collation != nil {
			result.Collation = opt.Collation
		}
		if opt.Comment != nil {
			result.Comment = opt.Comment
		}
		if opt.Hint != nil {
			result.Hint = opt.Hint
		}
		if opt.Let != nil {
			result.Let = opt.Let
		}
		if opt.DisableAutoCloseSession {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
		}
		if opt.ForceRecreateSession {
			result.ForceRecreateSession = opt.ForceRecreateSession
		}
		if opt.DisableAutoRollbackSession {
			result.DisableAutoRollbackSession = opt.DisableAutoRollbackSession
		}
	}
	return result
}
