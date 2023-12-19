package option

type Delete struct {
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.4 will return an error
	// if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option
	// is specified. The driver will return an error if this option is specified during an unacknowledged write
	// operation. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint any
	// Specifies parameters for the delete expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let                     any
	DisableAutoCloseSession bool
}

func NewDelete() Delete {
	return Delete{}
}

func (d Delete) SetCollation(c *Collation) Delete {
	d.Collation = c
	return d
}

func (d Delete) SetComment(a any) Delete {
	d.Comment = a
	return d
}

func (d Delete) SetHint(a any) Delete {
	d.Hint = a
	return d
}

func (d Delete) SetLet(a any) Delete {
	d.Let = a
	return d
}

func (d Delete) SetDisableAutoCloseTransaction(b bool) Delete {
	d.DisableAutoCloseSession = b
	return d
}

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
	}
	return result
}
