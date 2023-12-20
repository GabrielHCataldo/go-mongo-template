package option

type Update struct {
	// A set of filters specifying to which array elements an update should apply. This option is only valid for MongoDB
	// versions >= 3.6. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the update will apply to all array elements.
	ArrayFilters *ArrayFilters
	// If true, writes executed as part of the operation will opt out of document-level validation on the server. This
	// option is valid for MongoDB versions >= 3.2 and is ignored for previous server versions. The default value is
	// false. See https://www.mongodb.com/docs/manual/core/schema-validation/ for more information about document
	// validation.
	BypassDocumentValidation *bool
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.2. Server versions >= 3.4 will return an error
	// if this option is specified. For server versions < 3.4, the driver will return a client-side error if this option
	// is specified. The driver will return an error if this option is specified during an unacknowledged write
	// operation. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint any
	// If true, a new document will be inserted if the filter does not match any documents in the collection. The
	// default value is false.
	Upsert *bool
	// Specifies parameters for the update expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let                     any
	DisableAutoCloseSession bool
}

func NewUpdate() Update {
	return Update{}
}

func (u Update) SetBypassDocumentValidation(b bool) Update {
	u.BypassDocumentValidation = &b
	return u
}

func (u Update) SetArrayFilters(a *ArrayFilters) Update {
	u.ArrayFilters = a
	return u
}

func (u Update) SetCollation(c *Collation) Update {
	u.Collation = c
	return u
}

func (u Update) SetComment(a any) Update {
	u.Comment = a
	return u
}

func (u Update) SetHint(a any) Update {
	u.Hint = a
	return u
}

func (u Update) SetLet(a any) Update {
	u.Let = a
	return u
}

func (u Update) SetDisableAutoCloseTransaction(b bool) Update {
	u.DisableAutoCloseSession = b
	return u
}

func (u Update) SetUpsert(b bool) Update {
	u.Upsert = &b
	return u
}

func GetUpdateOptionByParams(opts []Update) Update {
	result := Update{}
	for _, opt := range opts {
		if opt.ArrayFilters != nil {
			result.ArrayFilters = opt.ArrayFilters
		}
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
