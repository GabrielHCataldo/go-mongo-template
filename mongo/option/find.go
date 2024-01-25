package option

import (
	"github.com/GabrielHCataldo/go-helper/helper"
	"time"
)

// Find represents options that can be used to configure a 'Find' operation.
type Find struct {
	// AllowDiskUse specifies whether the server can write temporary data to disk while executing the Find operation.
	// This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.2 will report an error if this option
	// is specified. For server versions < 3.2, the driver will return a client-side error if this option is specified.
	// The default value is false.
	AllowDiskUse *bool
	// AllowPartialResults AllowPartial results specifies whether the Find operation on a sharded cluster can
	// return partial results if some shards are down rather than returning an error. The default value is false.
	AllowPartialResults *bool
	// BatchSize is the maximum number of documents to be included in each batch returned by the server.
	BatchSize *int32
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// Comment A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is nil, which means that no comment will be included in the logs.
	Comment *string
	// CursorType specifies the type of cursor that should be created for the operation. The default is NonTailable, which
	// means that the cursor will be closed by the server when the last batch of documents is retrieved.
	CursorType *CursorType
	// Hint is the index to use for the Find operation. This should either be the index name as a string or the index
	// specification as a document. The driver will return an error if the hint parameter is a multi-key map. The default
	// value is nil, which means that no hint will be sent.
	Hint any
	// Limit is the maximum number of documents to return. The default value is 0, which means that all documents matching the
	// filter will be returned. A negative limit specifies that the resulting documents should be returned in a single
	// batch. The default value is 0.
	Limit *int64
	// Max is a document specifying the exclusive upper bound for a specific index. The default value is nil, which means that
	// there is no maximum value.
	Max any
	// MaxAwaitTime is the maximum amount of time that the server should wait for new documents to satisfy a tailable cursor
	// query. This option is only valid for tailable await cursors (see the CursorType option for more information) and
	// MongoDB versions >= 3.2. For other cursor types or previous server versions, this option is ignored.
	MaxAwaitTime *time.Duration
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
	// Min is a document specifying the inclusive lower bound for a specific index. The default value is 0, which means that
	// there is no minimum value.
	Min any
	// NoCursorTimeout specifies whether the cursor created by the operation will not time out after a period of inactivity.
	// The default value is false.
	NoCursorTimeout *bool
	// Projection Project is a document describing which fields will be included in the documents returned by the Find
	// operation. The default value is nil, which means all fields will be included.
	Projection any
	// ReturnKey specifies whether the documents returned by the Find operation will only contain fields corresponding to the
	// index used. The default value is false.
	ReturnKey *bool
	// ShowRecordID specifies whether a $recordId field with a record identifier will be included in the documents returned by
	// the Find operation. The default value is false.
	ShowRecordID *bool
	// Skip is the number of documents to skip before adding documents to the result. The default value is 0.
	Skip *int64
	// Sort is a document specifying the order in which documents should be returned.  The driver will return an error if the
	// sort parameter is a multi-key map.
	Sort any
	// Let specifies parameters for the find expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let any
}

// FindPageable represents options that can be used to configure a 'FindPageable' operation.
type FindPageable struct {
	// AllowDiskUse specifies whether the server can write temporary data to disk while executing the Find operation.
	// This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.2 will report an error if this option
	// is specified. For server versions < 3.2, the driver will return a client-side error if this option is specified.
	// The default value is false.
	AllowDiskUse *bool
	// AllowPartial results specifies whether the Find operation on a sharded cluster can return partial results if some
	// shards are down rather than returning an error. The default value is false.
	AllowPartialResults *bool
	// BatchSize is the maximum number of documents to be included in each batch returned by the server.
	BatchSize *int32
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is nil, which means that no comment will be included in the logs.
	Comment *string
	// CursorType specifies the type of cursor that should be created for the operation. The default is NonTailable, which
	// means that the cursor will be closed by the server when the last batch of documents is retrieved.
	CursorType *CursorType
	// Hint is the index to use for the Find operation. This should either be the index name as a string or the index
	// specification as a document. The driver will return an error if the hint parameter is a multi-key map. The default
	// value is nil, which means that no hint will be sent.
	Hint any
	// Max is a document specifying the exclusive upper bound for a specific index. The default value is nil, which means that
	// there is no maximum value.
	Max any
	// MaxAwaitTime is the maximum amount of time that the server should wait for new documents to satisfy a tailable cursor
	// query. This option is only valid for tailable await cursors (see the CursorType option for more information) and
	// MongoDB versions >= 3.2. For other cursor types or previous server versions, this option is ignored.
	MaxAwaitTime *time.Duration
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
	// Min is a document specifying the inclusive lower bound for a specific index. The default value is 0, which means that
	// there is no minimum value.
	Min any
	// NoCursorTimeout specifies whether the cursor created by the operation will not time out after a period of inactivity.
	// The default value is false.
	NoCursorTimeout *bool
	// Project is a document describing which fields will be included in the documents returned by the Find operation. The
	// default value is nil, which means all fields will be included.
	Projection any
	// ReturnKey specifies whether the documents returned by the Find operation will only contain fields corresponding to the
	// index used. The default value is false.
	ReturnKey *bool
	// ShowRecordID specifies whether a $recordId field with a record identifier will be included in the documents returned by
	// the Find operation. The default value is false.
	ShowRecordID *bool
	// Let specifies parameters for the find expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let any
}

// FindOne represents options that can be used to configure a FindOne operation.
type FindOne struct {
	// AllowPartialResults If true, an operation on a sharded cluster can return partial results if some shards are
	// down rather than returning an error. The default value is false.
	AllowPartialResults *bool
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// Comment A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is nil, which means that no comment will be included in the logs.
	Comment *string
	// Hint The index to use for the aggregation. This should either be the index name as a string or the index specification
	// as a document. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint any
	// Max A document specifying the exclusive upper bound for a specific index. The default value is nil, which means that
	// there is no maximum value.
	Max any
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
	// Min A document specifying the inclusive lower bound for a specific index. The default value is 0, which means that
	// there is no minimum value.
	Min any
	// Projection A document describing which fields will be included in the document returned by the operation.
	// The default value is nil, which means all fields will be included.
	Projection any
	// ReturnKey If true, the document returned by the operation will only contain fields corresponding to the index
	// used. The default value is false.
	ReturnKey *bool
	// ShowRecordID If true, a $recordId field with a record identifier will be included in the document returned by
	// the operation. The default value is false.
	ShowRecordID *bool
	// Skip The number of documents to skip before selecting the document to be returned. The default value is 0.
	Skip *int64
	// Sort A document specifying the sort order to apply to the query. The first document in the sorted order will be
	// returned. The driver will return an error if the sort parameter is a multi-key map.
	Sort any
}

// FindOneById represents options that can be used to configure a 'FindOneById' operation.
type FindOneById struct {
	// If true, an operation on a sharded cluster can return partial results if some shards are down rather than
	// returning an error. The default value is false.
	AllowPartialResults *bool
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is nil, which means that no comment will be included in the logs.
	Comment *string
	// The index to use for the aggregation. This should either be the index name as a string or the index specification
	// as a document. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint any
	// A document specifying the exclusive upper bound for a specific index. The default value is nil, which means that
	// there is no maximum value.
	Max any
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
	// A document specifying the inclusive lower bound for a specific index. The default value is 0, which means that
	// there is no minimum value.
	Min any
	// A document describing which fields will be included in the document returned by the operation. The default value
	// is nil, which means all fields will be included.
	Projection any
	// If true, the document returned by the operation will only contain fields corresponding to the index used. The
	// default value is false.
	ReturnKey *bool
	// If true, a $recordId field with a record identifier will be included in the document returned by the operation.
	// The default value is false.
	ShowRecordID *bool
}

// FindOneAndDelete represents options that can be used to configure a FindOneAndDelete operation.
type FindOneAndDelete struct {
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// Comment A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
	// Projection A document describing which fields will be included in the document returned by the operation.
	// The default value is nil, which means all fields will be included.
	Projection any
	// Sort A document specifying which document should be replaced if the filter used by the operation matches multiple
	// documents in the collection. If set, the first document in the sorted order will be selected for replacement.
	// The driver will return an error if the sort parameter is a multi-key map. The default value is nil.
	Sort any
	// Hint The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.4. MongoDB version 4.2 will report an error if
	// this option is specified. For server versions < 4.2, the driver will return an error if this option is specified.
	// The driver will return an error if this option is used with during an unacknowledged write operation. The driver
	// will return an error if the hint parameter is a multi-key map. The default value is nil, which means that no hint
	// will be sent.
	Hint any
	// Let Specifies parameters for the find one and delete expression. This option is only valid for
	// MongoDB versions >= 5.0. Older servers will report an error for using this option. This must be a document
	// mapping parameter names to values. Values must be constant or closed expressions that do not reference document
	// fields. Parameters can then be accessed as variables in an aggregate expression context (e.g. "$$var").
	Let any
	// DisableAutoCloseSession Disable automatic closing session, if true, we automatically close session according to
	// the result, if an error occurs, we abort the transaction, otherwise, we commit the transaction.
	// default is false.
	DisableAutoCloseSession *bool
	// ForceRecreateSession Force the creation of the session, if any session is still open, we close it automatically,
	// committing the transactions, and continue creating a new session.
	// default is false.
	ForceRecreateSession *bool
}

// FindOneAndReplace represents options that can be used to configure a FindOneAndReplace operation.
type FindOneAndReplace struct {
	// BypassDocumentValidation If true, writes executed as part of the operation will opt out of document-level
	// validation on the server. This option is valid for MongoDB versions >= 3.2 and is ignored for previous server
	// versions. The default value is false. See https://www.mongodb.com/docs/manual/core/schema-validation/ for more
	// information about document validation.
	BypassDocumentValidation *bool
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// Comment A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
	// Projection A document describing which fields will be included in the document returned by the operation.
	// The default value is nil, which means all fields will be included.
	Projection any
	// ReturnDocument Specifies whether the original or replaced document should be returned by the operation.
	// The default value is Before, which means the original document will be returned from before the replacement is performed.
	ReturnDocument *ReturnDocument
	// Sort A document specifying which document should be replaced if the filter used by the operation matches multiple
	// documents in the collection. If set, the first document in the sorted order will be replaced. The driver will
	// return an error if the sort parameter is a multi-key map. The default value is nil.
	Sort any
	// Upsert If true, a new document will be inserted if the filter does not match any documents in the collection. The
	// default value is false.
	Upsert *bool
	// Hint The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.4. MongoDB version 4.2 will report an error if
	// this option is specified. For server versions < 4.2, the driver will return an error if this option is specified.
	// The driver will return an error if this option is used with during an unacknowledged write operation. The driver
	// will return an error if the hint parameter is a multi-key map. The default value is nil, which means that no hint
	// will be sent.
	Hint any
	// Let Specifies parameters for the find one and replace expression. This option is only valid for MongoDB
	// versions >= 5.0. Older servers will report an error for using this option. This must be a document mapping
	// parameter names to values. Values must be constant or closed expressions that do not reference document fields.
	// Parameters can then be accessed as variables in an aggregate expression context (e.g. "$$var").
	Let any
	// DisableAutoCloseSession Disable automatic closing session, if true, we automatically close session according to
	// the result, if an error occurs, we abort the transaction, otherwise, we commit the transaction.
	// default is false.
	DisableAutoCloseSession *bool
	// ForceRecreateSession Force the creation of the session, if any session is still open, we close it automatically,
	// committing the transactions, and continue creating a new session.
	// default is false.
	ForceRecreateSession *bool
}

// FindOneAndUpdate represents options that can be used to configure a FindOneAndUpdate operation.
type FindOneAndUpdate struct {
	// ArrayFilters A set of filters specifying to which array elements an update should apply. This option is only
	// valid for MongoDB versions >= 3.6. For previous server versions, the driver will return an error if this option
	// is used. The default value is nil, which means the update will apply to all array elements.
	ArrayFilters *ArrayFilters
	// BypassDocumentValidation If true, writes executed as part of the operation will opt out of document-level
	// validation on the server. This option is valid for MongoDB versions >= 3.2 and is ignored for previous server
	// versions. The default value is false. See https://www.mongodb.com/docs/manual/core/schema-validation/ for more
	// information about document validation.
	BypassDocumentValidation *bool
	// Collation Specifies a collation to use for string comparisons during the operation. This option is only valid
	// for MongoDB  versions >= 3.4. For previous server versions, the driver will return an error if this option is
	// used. The default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// Comment A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// MaxTime The maximum amount of time that the query can run on the server. The default value is nil, meaning that
	// there is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
	// Projection A document describing which fields will be included in the document returned by the operation.
	// The default value is nil, which means all fields will be included.
	Projection any
	// ReturnDocument Specifies whether the original or replaced document should be returned by the operation.
	// The default value is Before, which means the original document will be returned before the replacement is performed.
	ReturnDocument *ReturnDocument
	// Sort A document specifying which document should be updated if the filter used by the operation matches multiple
	// documents in the collection. If set, the first document in the sorted order will be updated. The driver will
	// return an error if the sort parameter is a multi-key map. The default value is nil.
	Sort any
	// Upsert If true, a new document will be inserted if the filter does not match any documents in the collection. The
	// default value is false.
	Upsert *bool
	// Hint The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.4. MongoDB version 4.2 will report an error if
	// this option is specified. For server versions < 4.2, the driver will return an error if this option is specified.
	// The driver will return an error if this option is used with during an unacknowledged write operation. The driver
	// will return an error if the hint parameter is a multi-key map. The default value is nil, which means that no hint
	// will be sent.
	Hint any
	// Let Specifies parameters for the find one and update expression. This option is only valid for MongoDB versions >= 5.0. Older
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

// NewFind creates a new Find instance.
func NewFind() *Find {
	return &Find{}
}

// NewFindPageable creates a new FindPageable instance.
func NewFindPageable() *FindPageable {
	return &FindPageable{}
}

// NewFindOne creates a new FindOne instance.
func NewFindOne() *FindOne {
	return &FindOne{}
}

// NewFindOneById creates a new FindOneById instance.
func NewFindOneById() *FindOneById {
	return &FindOneById{}
}

// NewFindOneAndDelete creates a new FindOneAndDelete instance.
func NewFindOneAndDelete() *FindOneAndDelete {
	return &FindOneAndDelete{}
}

// NewFindOneAndReplace creates a new FindOneAndReplace instance.
func NewFindOneAndReplace() *FindOneAndReplace {
	return &FindOneAndReplace{}
}

// NewFindOneAndUpdate creates a new FindOneAndUpdate instance.
func NewFindOneAndUpdate() *FindOneAndUpdate {
	return &FindOneAndUpdate{}
}

// SetAllowDiskUse creates a new AllowDiskUse instance.
func (f *Find) SetAllowDiskUse(b bool) *Find {
	f.AllowDiskUse = &b
	return f
}

// SetAllowPartialResults creates a new AllowPartialResults instance.
func (f *Find) SetAllowPartialResults(b bool) *Find {
	f.AllowPartialResults = &b
	return f
}

// SetBatchSize creates a new BatchSize instance.
func (f *Find) SetBatchSize(i int32) *Find {
	f.BatchSize = &i
	return f
}

// SetCollation creates a new Collation instance.
func (f *Find) SetCollation(c *Collation) *Find {
	f.Collation = c
	return f
}

// SetComment creates a new Comment instance.
func (f *Find) SetComment(s string) *Find {
	f.Comment = &s
	return f
}

// SetCursorType creates a new CursorType instance.
func (f *Find) SetCursorType(c CursorType) *Find {
	f.CursorType = &c
	return f
}

// SetHint creates a new Hint instance.
func (f *Find) SetHint(v any) *Find {
	f.Hint = v
	return f
}

// SetLimit creates a new Limit instance.
func (f *Find) SetLimit(i int64) *Find {
	f.Limit = &i
	return f
}

// SetMax creates a new Max instance.
func (f *Find) SetMax(v any) *Find {
	f.Max = v
	return f
}

// SetMaxAwaitTime creates a new MaxAwaitTime instance.
func (f *Find) SetMaxAwaitTime(d time.Duration) *Find {
	f.MaxAwaitTime = &d
	return f
}

// SetMaxTime creates a new MaxTime instance.
func (f *Find) SetMaxTime(d time.Duration) *Find {
	f.MaxTime = &d
	return f
}

// SetMin creates a new Min instance.
func (f *Find) SetMin(v any) *Find {
	f.Min = v
	return f
}

// SetNoCursorTimeout creates a new NoCursorTimeout instance.
func (f *Find) SetNoCursorTimeout(b bool) *Find {
	f.NoCursorTimeout = &b
	return f
}

// SetProjection creates a new Projection instance.
func (f *Find) SetProjection(v any) *Find {
	f.Projection = v
	return f
}

// SetReturnKey creates a new ReturnKey instance.
func (f *Find) SetReturnKey(b bool) *Find {
	f.ReturnKey = &b
	return f
}

// SetShowRecordID creates a new ShowRecordID instance.
func (f *Find) SetShowRecordID(b bool) *Find {
	f.ShowRecordID = &b
	return f
}

// SetSkip creates a new Skip instance.
func (f *Find) SetSkip(i int64) *Find {
	f.Skip = &i
	return f
}

// SetSort creates a new Sort instance.
func (f *Find) SetSort(a any) *Find {
	f.Sort = a
	return f
}

// SetLet creates a new Let instance.
func (f *Find) SetLet(v any) *Find {
	f.Let = v
	return f
}

// SetAllowDiskUse creates a new AllowDiskUse instance.
func (f *FindPageable) SetAllowDiskUse(b bool) *FindPageable {
	f.AllowDiskUse = &b
	return f
}

// SetAllowPartialResults creates a new AllowPartialResults instance.
func (f *FindPageable) SetAllowPartialResults(b bool) *FindPageable {
	f.AllowPartialResults = &b
	return f
}

// SetBatchSize creates a new BatchSize instance.
func (f *FindPageable) SetBatchSize(i int32) *FindPageable {
	f.BatchSize = &i
	return f
}

// SetCollation creates a new Collation instance.
func (f *FindPageable) SetCollation(c *Collation) *FindPageable {
	f.Collation = c
	return f
}

// SetComment creates a new Comment instance.
func (f *FindPageable) SetComment(s string) *FindPageable {
	f.Comment = &s
	return f
}

// SetCursorType creates a new CursorType instance.
func (f *FindPageable) SetCursorType(c CursorType) *FindPageable {
	f.CursorType = &c
	return f
}

// SetHint creates a new Hint instance.
func (f *FindPageable) SetHint(v any) *FindPageable {
	f.Hint = v
	return f
}

// SetMax creates a new Max instance.
func (f *FindPageable) SetMax(v any) *FindPageable {
	f.Max = v
	return f
}

// SetMaxAwaitTime creates a new MaxAwaitTime instance.
func (f *FindPageable) SetMaxAwaitTime(d time.Duration) *FindPageable {
	f.MaxAwaitTime = &d
	return f
}

// SetMaxTime creates a new MaxTime instance.
func (f *FindPageable) SetMaxTime(d time.Duration) *FindPageable {
	f.MaxTime = &d
	return f
}

// SetMin creates a new Min instance.
func (f *FindPageable) SetMin(v any) *FindPageable {
	f.Min = v
	return f
}

// SetNoCursorTimeout creates a new NoCursorTimeout instance.
func (f *FindPageable) SetNoCursorTimeout(b bool) *FindPageable {
	f.NoCursorTimeout = &b
	return f
}

// SetProjection creates a new Projection instance.
func (f *FindPageable) SetProjection(v any) *FindPageable {
	f.Projection = v
	return f
}

// SetReturnKey creates a new ReturnKey instance.
func (f *FindPageable) SetReturnKey(b bool) *FindPageable {
	f.ReturnKey = &b
	return f
}

// SetShowRecordID creates a new ShowRecordID instance.
func (f *FindPageable) SetShowRecordID(b bool) *FindPageable {
	f.ShowRecordID = &b
	return f
}

// SetLet creates a new Let instance.
func (f *FindPageable) SetLet(v any) *FindPageable {
	f.Let = v
	return f
}

// SetAllowPartialResults creates a new AllowPartialResults instance.
func (f *FindOne) SetAllowPartialResults(b bool) *FindOne {
	f.AllowPartialResults = &b
	return f
}

// SetCollation creates a new Collation instance.
func (f *FindOne) SetCollation(c *Collation) *FindOne {
	f.Collation = c
	return f
}

// SetComment creates a new Comment instance.
func (f *FindOne) SetComment(s string) *FindOne {
	f.Comment = &s
	return f
}

// SetHint creates a new Hint instance.
func (f *FindOne) SetHint(v any) *FindOne {
	f.Hint = v
	return f
}

// SetMax creates a new Max instance.
func (f *FindOne) SetMax(v any) *FindOne {
	f.Max = v
	return f
}

// SetMaxTime creates a new MaxTime instance.
func (f *FindOne) SetMaxTime(d time.Duration) *FindOne {
	f.MaxTime = &d
	return f
}

// SetMin creates a new Min instance.
func (f *FindOne) SetMin(v any) *FindOne {
	f.Min = v
	return f
}

// SetProjection creates a new Projection instance.
func (f *FindOne) SetProjection(v any) *FindOne {
	f.Projection = v
	return f
}

// SetReturnKey creates a new ReturnKey instance.
func (f *FindOne) SetReturnKey(b bool) *FindOne {
	f.ReturnKey = &b
	return f
}

// SetShowRecordID creates a new ShowRecordID instance.
func (f *FindOne) SetShowRecordID(b bool) *FindOne {
	f.ShowRecordID = &b
	return f
}

// SetSkip creates a new Skip instance.
func (f *FindOne) SetSkip(i int64) *FindOne {
	f.Skip = &i
	return f
}

// SetSort creates a new Sort instance.
func (f *FindOne) SetSort(a any) *FindOne {
	f.Sort = a
	return f
}

// SetAllowPartialResults creates a new AllowPartialResults instance.
func (f *FindOneById) SetAllowPartialResults(b bool) *FindOneById {
	f.AllowPartialResults = &b
	return f
}

// SetCollation creates a new Collation instance.
func (f *FindOneById) SetCollation(c *Collation) *FindOneById {
	f.Collation = c
	return f
}

// SetComment creates a new Comment instance.
func (f *FindOneById) SetComment(s string) *FindOneById {
	f.Comment = &s
	return f
}

// SetHint creates a new Hint instance.
func (f *FindOneById) SetHint(v any) *FindOneById {
	f.Hint = v
	return f
}

// SetMax creates a new Max instance.
func (f *FindOneById) SetMax(v any) *FindOneById {
	f.Max = v
	return f
}

// SetMaxTime creates a new MaxTime instance.
func (f *FindOneById) SetMaxTime(d time.Duration) *FindOneById {
	f.MaxTime = &d
	return f
}

// SetMin creates a new Min instance.
func (f *FindOneById) SetMin(v any) *FindOneById {
	f.Min = v
	return f
}

// SetProjection creates a new Projection instance.
func (f *FindOneById) SetProjection(v any) *FindOneById {
	f.Projection = v
	return f
}

// SetReturnKey creates a new ReturnKey instance.
func (f *FindOneById) SetReturnKey(b bool) *FindOneById {
	f.ReturnKey = &b
	return f
}

// SetShowRecordID creates a new ShowRecordID instance.
func (f *FindOneById) SetShowRecordID(b bool) *FindOneById {
	f.ShowRecordID = &b
	return f
}

// SetCollation creates a new Collation instance.
func (f *FindOneAndDelete) SetCollation(c *Collation) *FindOneAndDelete {
	f.Collation = c
	return f
}

// SetComment creates a new Comment instance.
func (f *FindOneAndDelete) SetComment(s string) *FindOneAndDelete {
	f.Comment = s
	return f
}

// SetHint creates a new Hint instance.
func (f *FindOneAndDelete) SetHint(v any) *FindOneAndDelete {
	f.Hint = v
	return f
}

// SetMaxTime creates a new MaxTime instance.
func (f *FindOneAndDelete) SetMaxTime(d time.Duration) *FindOneAndDelete {
	f.MaxTime = &d
	return f
}

// SetProjection creates a new Projection instance.
func (f *FindOneAndDelete) SetProjection(v any) *FindOneAndDelete {
	f.Projection = v
	return f
}

// SetSort creates a new Sort instance.
func (f *FindOneAndDelete) SetSort(a any) *FindOneAndDelete {
	f.Sort = a
	return f
}

// SetLet creates a new Let instance.
func (f *FindOneAndDelete) SetLet(v any) *FindOneAndDelete {
	f.Let = v
	return f
}

// SetDisableAutoCloseSession creates a new DisableAutoCloseSession instance.
func (f *FindOneAndDelete) SetDisableAutoCloseSession(b bool) *FindOneAndDelete {
	f.DisableAutoCloseSession = &b
	return f
}

// SetForceRecreateSession creates a new ForceRecreateSession instance.
func (f *FindOneAndDelete) SetForceRecreateSession(b bool) *FindOneAndDelete {
	f.ForceRecreateSession = &b
	return f
}

// SetDisableAutoCloseSession creates a new DisableAutoCloseSession instance.
func (f *FindOneAndReplace) SetDisableAutoCloseSession(b bool) *FindOneAndReplace {
	f.DisableAutoCloseSession = &b
	return f
}

// SetForceRecreateSession sets value for the ForceRecreateSession field.
func (f *FindOneAndReplace) SetForceRecreateSession(b bool) *FindOneAndReplace {
	f.ForceRecreateSession = &b
	return f
}

// SetBypassDocumentValidation sets value for the BypassDocumentValidation field.
func (f *FindOneAndReplace) SetBypassDocumentValidation(b bool) *FindOneAndReplace {
	f.BypassDocumentValidation = &b
	return f
}

// SetCollation sets value for the Collation field.
func (f *FindOneAndReplace) SetCollation(c *Collation) *FindOneAndReplace {
	f.Collation = c
	return f
}

// SetComment sets value for the Comment field.
func (f *FindOneAndReplace) SetComment(s string) *FindOneAndReplace {
	f.Comment = s
	return f
}

// SetHint sets value for the Hint field.
func (f *FindOneAndReplace) SetHint(v any) *FindOneAndReplace {
	f.Hint = v
	return f
}

// SetMaxTime creates a new MaxTime instance.
func (f *FindOneAndReplace) SetMaxTime(d time.Duration) *FindOneAndReplace {
	f.MaxTime = &d
	return f
}

// SetProjection creates a new Projection instance.
func (f *FindOneAndReplace) SetProjection(v any) *FindOneAndReplace {
	f.Projection = v
	return f
}

// SetReturnDocument creates a new ReturnDocument instance.
func (f *FindOneAndReplace) SetReturnDocument(r ReturnDocument) *FindOneAndReplace {
	f.ReturnDocument = &r
	return f
}

// SetUpsert creates a new Upsert instance.
func (f *FindOneAndReplace) SetUpsert(b bool) *FindOneAndReplace {
	f.Upsert = &b
	return f
}

// SetLet sets value for the Let field.
func (f *FindOneAndReplace) SetLet(v any) *FindOneAndReplace {
	f.Let = v
	return f
}

// SetSort creates a new Sort instance.
func (f *FindOneAndReplace) SetSort(a any) *FindOneAndReplace {
	f.Sort = a
	return f
}

// SetDisableAutoCloseSession creates a new DisableAutoCloseSession instance.
func (f *FindOneAndUpdate) SetDisableAutoCloseSession(b bool) *FindOneAndUpdate {
	f.DisableAutoCloseSession = &b
	return f
}

// SetForceRecreateSession sets value for the ForceRecreateSession field.
func (f *FindOneAndUpdate) SetForceRecreateSession(b bool) *FindOneAndUpdate {
	f.ForceRecreateSession = &b
	return f
}

func (f *FindOneAndUpdate) SetArrayFilters(a *ArrayFilters) *FindOneAndUpdate {
	f.ArrayFilters = a
	return f
}

// SetBypassDocumentValidation sets value for the BypassDocumentValidation field.
func (f *FindOneAndUpdate) SetBypassDocumentValidation(b bool) *FindOneAndUpdate {
	f.BypassDocumentValidation = &b
	return f
}

// SetCollation sets value for the Collation field.
func (f *FindOneAndUpdate) SetCollation(c *Collation) *FindOneAndUpdate {
	f.Collation = c
	return f
}

// SetComment sets value for the Comment field.
func (f *FindOneAndUpdate) SetComment(a any) *FindOneAndUpdate {
	f.Comment = a
	return f
}

// SetHint sets value for the Hint field.
func (f *FindOneAndUpdate) SetHint(v any) *FindOneAndUpdate {
	f.Hint = v
	return f
}

// SetLet sets value for the Let field.
func (f *FindOneAndUpdate) SetLet(v any) *FindOneAndUpdate {
	f.Let = v
	return f
}

// SetMaxTime creates a new MaxTime instance.
func (f *FindOneAndUpdate) SetMaxTime(d time.Duration) *FindOneAndUpdate {
	f.MaxTime = &d
	return f
}

// SetProjection creates a new Projection instance.
func (f *FindOneAndUpdate) SetProjection(v any) *FindOneAndUpdate {
	f.Projection = v
	return f
}

// SetReturnDocument creates a new ReturnDocument instance.
func (f *FindOneAndUpdate) SetReturnDocument(r ReturnDocument) *FindOneAndUpdate {
	f.ReturnDocument = &r
	return f
}

// SetSort creates a new Sort instance.
func (f *FindOneAndUpdate) SetSort(i any) *FindOneAndUpdate {
	f.Sort = i
	return f
}

// SetUpsert creates a new Upsert instance.
func (f *FindOneAndUpdate) SetUpsert(b bool) *FindOneAndUpdate {
	f.Upsert = &b
	return f
}

// MergeFindByParams assembles the Find object from optional parameters.
func MergeFindByParams(opts []*Find) *Find {
	result := &Find{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.AllowDiskUse) {
			result.AllowDiskUse = opt.AllowDiskUse
		}
		if helper.IsNotNil(opt.AllowPartialResults) {
			result.AllowPartialResults = opt.AllowPartialResults
		}
		if helper.IsNotNil(opt.NoCursorTimeout) {
			result.NoCursorTimeout = opt.NoCursorTimeout
		}
		if helper.IsNotNil(opt.ReturnKey) {
			result.ReturnKey = opt.ReturnKey
		}
		if helper.IsNotNil(opt.ShowRecordID) {
			result.ShowRecordID = opt.ShowRecordID
		}
		if helper.IsNotNil(opt.CursorType) {
			result.CursorType = opt.CursorType
		}
		if helper.IsNotNil(opt.BatchSize) {
			result.BatchSize = opt.BatchSize
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
		if helper.IsNotNil(opt.Limit) {
			result.Limit = opt.Limit
		}
		if helper.IsNotNil(opt.Skip) {
			result.Skip = opt.Skip
		}
		if helper.IsNotNil(opt.Max) {
			result.Max = opt.Max
		}
		if helper.IsNotNil(opt.Min) {
			result.Min = opt.Min
		}
		if helper.IsNotNil(opt.Sort) {
			result.Sort = opt.Sort
		}
		if helper.IsNotNil(opt.Projection) {
			result.Projection = opt.Projection
		}
		if helper.IsNotNil(opt.MaxTime) {
			result.MaxTime = opt.MaxTime
		}
		if helper.IsNotNil(opt.MaxAwaitTime) {
			result.MaxAwaitTime = opt.MaxAwaitTime
		}
	}
	return result
}

// MergeFindPageableByParams assembles the FindPageable object from optional parameters.
func MergeFindPageableByParams(opts []*FindPageable) *FindPageable {
	result := &FindPageable{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.AllowDiskUse) {
			result.AllowDiskUse = opt.AllowDiskUse
		}
		if helper.IsNotNil(opt.AllowPartialResults) {
			result.AllowPartialResults = opt.AllowPartialResults
		}
		if helper.IsNotNil(opt.NoCursorTimeout) {
			result.NoCursorTimeout = opt.NoCursorTimeout
		}
		if helper.IsNotNil(opt.ReturnKey) {
			result.ReturnKey = opt.ReturnKey
		}
		if helper.IsNotNil(opt.ShowRecordID) {
			result.ShowRecordID = opt.ShowRecordID
		}
		if helper.IsNotNil(opt.CursorType) {
			result.CursorType = opt.CursorType
		}
		if helper.IsNotNil(opt.BatchSize) {
			result.BatchSize = opt.BatchSize
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
		if helper.IsNotNil(opt.Max) {
			result.Max = opt.Max
		}
		if helper.IsNotNil(opt.Min) {
			result.Min = opt.Min
		}
		if helper.IsNotNil(opt.Projection) {
			result.Projection = opt.Projection
		}
		if helper.IsNotNil(opt.MaxTime) {
			result.MaxTime = opt.MaxTime
		}
		if helper.IsNotNil(opt.MaxAwaitTime) {
			result.MaxAwaitTime = opt.MaxAwaitTime
		}
	}
	return result
}

// MergeFindOneByParams assembles the FindOne object from optional parameters.
func MergeFindOneByParams(opts []*FindOne) *FindOne {
	result := &FindOne{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.AllowPartialResults) {
			result.AllowPartialResults = opt.AllowPartialResults
		}
		if helper.IsNotNil(opt.ReturnKey) {
			result.ReturnKey = opt.ReturnKey
		}
		if helper.IsNotNil(opt.ShowRecordID) {
			result.ShowRecordID = opt.ShowRecordID
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
		if helper.IsNotNil(opt.Skip) {
			result.Skip = opt.Skip
		}
		if helper.IsNotNil(opt.Max) {
			result.Max = opt.Max
		}
		if helper.IsNotNil(opt.Min) {
			result.Min = opt.Min
		}
		if helper.IsNotNil(opt.Sort) {
			result.Sort = opt.Sort
		}
		if helper.IsNotNil(opt.Projection) {
			result.Projection = opt.Projection
		}
		if helper.IsNotNil(opt.MaxTime) {
			result.MaxTime = opt.MaxTime
		}
	}
	return result
}

// MergeFindOneByIdByParams assembles the FindOneById object from optional parameters.
func MergeFindOneByIdByParams(opts []*FindOneById) *FindOneById {
	result := &FindOneById{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.AllowPartialResults) {
			result.AllowPartialResults = opt.AllowPartialResults
		}
		if helper.IsNotNil(opt.ReturnKey) {
			result.ReturnKey = opt.ReturnKey
		}
		if helper.IsNotNil(opt.ShowRecordID) {
			result.ShowRecordID = opt.ShowRecordID
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
		if helper.IsNotNil(opt.Max) {
			result.Max = opt.Max
		}
		if helper.IsNotNil(opt.Min) {
			result.Min = opt.Min
		}
		if helper.IsNotNil(opt.Projection) {
			result.Projection = opt.Projection
		}
		if helper.IsNotNil(opt.MaxTime) {
			result.MaxTime = opt.MaxTime
		}
	}
	return result
}

// MergeFindOneAndDeleteByParams assembles the FindOneAndDelete object from optional parameters.
func MergeFindOneAndDeleteByParams(opts []*FindOneAndDelete) *FindOneAndDelete {
	result := &FindOneAndDelete{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.DisableAutoCloseSession) {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
		}
		if helper.IsNotNil(opt.ForceRecreateSession) {
			result.ForceRecreateSession = opt.ForceRecreateSession
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
		if helper.IsNotNil(opt.Sort) {
			result.Sort = opt.Sort
		}
		if helper.IsNotNil(opt.Projection) {
			result.Projection = opt.Projection
		}
		if helper.IsNotNil(opt.MaxTime) {
			result.MaxTime = opt.MaxTime
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

// MergeFindOneAndReplaceByParams assembles the FindOneAndReplace object from optional parameters.
func MergeFindOneAndReplaceByParams(opts []*FindOneAndReplace) *FindOneAndReplace {
	result := &FindOneAndReplace{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.BypassDocumentValidation) {
			result.BypassDocumentValidation = opt.BypassDocumentValidation
		}
		if helper.IsNotNil(opt.DisableAutoCloseSession) {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
		}
		if helper.IsNotNil(opt.ForceRecreateSession) {
			result.ForceRecreateSession = opt.ForceRecreateSession
		}
		if helper.IsNotNil(opt.Upsert) {
			result.Upsert = opt.Upsert
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
		if helper.IsNotNil(opt.Sort) {
			result.Sort = opt.Sort
		}
		if helper.IsNotNil(opt.Projection) {
			result.Projection = opt.Projection
		}
		if helper.IsNotNil(opt.MaxTime) {
			result.MaxTime = opt.MaxTime
		}
		if helper.IsNotNil(opt.ReturnDocument) {
			result.ReturnDocument = opt.ReturnDocument
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

// MergeFindOneAndUpdateByParams assembles the FindOneAndUpdate object from optional parameters.
func MergeFindOneAndUpdateByParams(opts []*FindOneAndUpdate) *FindOneAndUpdate {
	result := &FindOneAndUpdate{}
	for _, opt := range opts {
		if helper.IsNil(opt) {
			continue
		}
		if helper.IsNotNil(opt.BypassDocumentValidation) {
			result.BypassDocumentValidation = opt.BypassDocumentValidation
		}
		if helper.IsNotNil(opt.DisableAutoCloseSession) {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
		}
		if helper.IsNotNil(opt.ForceRecreateSession) {
			result.ForceRecreateSession = opt.ForceRecreateSession
		}
		if helper.IsNotNil(opt.Upsert) {
			result.Upsert = opt.Upsert
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
		if helper.IsNotNil(opt.Sort) {
			result.Sort = opt.Sort
		}
		if helper.IsNotNil(opt.Projection) {
			result.Projection = opt.Projection
		}
		if helper.IsNotNil(opt.MaxTime) {
			result.MaxTime = opt.MaxTime
		}
		if helper.IsNotNil(opt.ReturnDocument) {
			result.ReturnDocument = opt.ReturnDocument
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
