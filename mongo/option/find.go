package option

import (
	"time"
)

type Find struct {
	// AllowDiskUse specifies whether the server can write temporary data to disk while executing the Find operation.
	// This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.2 will report an error if this option
	// is specified. For server versions < 3.2, the driver will return a client-side error if this option is specified.
	// The default value is false.
	AllowDiskUse bool
	// AllowPartial results specifies whether the Find operation on a sharded cluster can return partial results if some
	// shards are down rather than returning an error. The default value is false.
	AllowPartialResults bool
	// BatchSize is the maximum number of documents to be included in each batch returned by the server.
	BatchSize int32
	// Collation specifies a collation to use for string comparisons during the operation. This option is only valid for
	// MongoDB versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is nil, which means that no comment will be included in the logs.
	Comment string
	// CursorType specifies the type of cursor that should be created for the operation. The default is NonTailable, which
	// means that the cursor will be closed by the server when the last batch of documents is retrieved.
	CursorType CursorType
	// Hint is the index to use for the Find operation. This should either be the index name as a string or the index
	// specification as a document. The driver will return an error if the hint parameter is a multi-key map. The default
	// value is nil, which means that no hint will be sent.
	Hint any
	// Limit is the maximum number of documents to return. The default value is 0, which means that all documents matching the
	// filter will be returned. A negative limit specifies that the resulting documents should be returned in a single
	// batch. The default value is 0.
	Limit int64
	// Max is a document specifying the exclusive upper bound for a specific index. The default value is nil, which means that
	// there is no maximum value.
	Max any
	// MaxAwaitTime is the maximum amount of time that the server should wait for new documents to satisfy a tailable cursor
	// query. This option is only valid for tailable await cursors (see the CursorType option for more information) and
	// MongoDB versions >= 3.2. For other cursor types or previous server versions, this option is ignored.
	MaxAwaitTime time.Duration
	// MaxTime is the maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in its
	// place to control the amount of time that a single operation can run before returning an error. MaxTime is ignored if
	// Timeout is set on the client.
	MaxTime time.Duration
	// Min is a document specifying the inclusive lower bound for a specific index. The default value is 0, which means that
	// there is no minimum value.
	Min any
	// NoCursorTimeout specifies whether the cursor created by the operation will not time out after a period of inactivity.
	// The default value is false.
	NoCursorTimeout bool
	// Project is a document describing which fields will be included in the documents returned by the Find operation. The
	// default value is nil, which means all fields will be included.
	Projection any
	// ReturnKey specifies whether the documents returned by the Find operation will only contain fields corresponding to the
	// index used. The default value is false.
	ReturnKey bool
	// ShowRecordID specifies whether a $recordId field with a record identifier will be included in the documents returned by
	// the Find operation. The default value is false.
	ShowRecordID bool
	// Skip is the number of documents to skip before adding documents to the result. The default value is 0.
	Skip int64
	// Sort is a document specifying the order in which documents should be returned.  The driver will return an error if the
	// sort parameter is a multi-key map.
	Sort any
	// Let specifies parameters for the find expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let any
}

type FindPageable struct {
	// AllowDiskUse specifies whether the server can write temporary data to disk while executing the Find operation.
	// This option is only valid for MongoDB versions >= 4.4. Server versions >= 3.2 will report an error if this option
	// is specified. For server versions < 3.2, the driver will return a client-side error if this option is specified.
	// The default value is false.
	AllowDiskUse bool
	// AllowPartial results specifies whether the Find operation on a sharded cluster can return partial results if some
	// shards are down rather than returning an error. The default value is false.
	AllowPartialResults bool
	// BatchSize is the maximum number of documents to be included in each batch returned by the server.
	BatchSize int32
	// Collation specifies a collation to use for string comparisons during the operation. This option is only valid for
	// MongoDB versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is nil, which means that no comment will be included in the logs.
	Comment string
	// CursorType specifies the type of cursor that should be created for the operation. The default is NonTailable, which
	// means that the cursor will be closed by the server when the last batch of documents is retrieved.
	CursorType CursorType
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
	MaxAwaitTime time.Duration
	// MaxTime is the maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in its
	// place to control the amount of time that a single operation can run before returning an error. MaxTime is ignored if
	// Timeout is set on the client.
	MaxTime time.Duration
	// Min is a document specifying the inclusive lower bound for a specific index. The default value is 0, which means that
	// there is no minimum value.
	Min any
	// NoCursorTimeout specifies whether the cursor created by the operation will not time out after a period of inactivity.
	// The default value is false.
	NoCursorTimeout bool
	// Project is a document describing which fields will be included in the documents returned by the Find operation. The
	// default value is nil, which means all fields will be included.
	Projection any
	// ReturnKey specifies whether the documents returned by the Find operation will only contain fields corresponding to the
	// index used. The default value is false.
	ReturnKey bool
	// ShowRecordID specifies whether a $recordId field with a record identifier will be included in the documents returned by
	// the Find operation. The default value is false.
	ShowRecordID bool
	// Let specifies parameters for the find expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let any
}

type FindOne struct {
	// If true, an operation on a sharded cluster can return partial results if some shards are down rather than
	// returning an error. The default value is false.
	AllowPartialResults bool
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is nil, which means that no comment will be included in the logs.
	Comment string
	// The index to use for the aggregation. This should either be the index name as a string or the index specification
	// as a document. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint any
	// A document specifying the exclusive upper bound for a specific index. The default value is nil, which means that
	// there is no maximum value.
	Max any
	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used
	// in its place to control the amount of time that a single operation can run before returning an error. MaxTime
	// is ignored if Timeout is set on the client.
	MaxTime time.Duration
	// A document specifying the inclusive lower bound for a specific index. The default value is 0, which means that
	// there is no minimum value.
	Min any
	// A document describing which fields will be included in the document returned by the operation. The default value
	// is nil, which means all fields will be included.
	Projection any
	// If true, the document returned by the operation will only contain fields corresponding to the index used. The
	// default value is false.
	ReturnKey bool
	// If true, a $recordId field with a record identifier will be included in the document returned by the operation.
	// The default value is false.
	ShowRecordID bool
	// The number of documents to skip before selecting the document to be returned. The default value is 0.
	Skip int64
	// A document specifying the sort order to apply to the query. The first document in the sorted order will be
	// returned. The driver will return an error if the sort parameter is a multi-key map.
	Sort any
}

type FindOneAndDelete struct {
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used
	// in its place to control the amount of time that a single operation can run before returning an error. MaxTime
	// is ignored if Timeout is set on the client.
	MaxTime time.Duration
	// A document describing which fields will be included in the document returned by the operation. The default value
	// is nil, which means all fields will be included.
	Projection any
	// A document specifying which document should be replaced if the filter used by the operation matches multiple
	// documents in the collection. If set, the first document in the sorted order will be selected for replacement.
	// The driver will return an error if the sort parameter is a multi-key map. The default value is nil.
	Sort any
	// The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.4. MongoDB version 4.2 will report an error if
	// this option is specified. For server versions < 4.2, the driver will return an error if this option is specified.
	// The driver will return an error if this option is used with during an unacknowledged write operation. The driver
	// will return an error if the hint parameter is a multi-key map. The default value is nil, which means that no hint
	// will be sent.
	Hint any
	// Specifies parameters for the find one and delete expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let                     any
	DisableAutoCloseSession bool
}

type FindOneAndReplace struct {
	// If true, writes executed as part of the operation will opt out of document-level validation on the server. This
	// option is valid for MongoDB versions >= 3.2 and is ignored for previous server versions. The default value is
	// false. See https://www.mongodb.com/docs/manual/core/schema-validation/ for more information about document
	// validation.
	BypassDocumentValidation bool
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used
	// in its place to control the amount of time that a single operation can run before returning an error. MaxTime
	// is ignored if Timeout is set on the client.
	MaxTime time.Duration
	// A document describing which fields will be included in the document returned by the operation. The default value
	// is nil, which means all fields will be included.
	Projection any
	// Specifies whether the original or replaced document should be returned by the operation. The default value is
	// Before, which means the original document will be returned from before the replacement is performed.
	ReturnDocument ReturnDocument
	// A document specifying which document should be replaced if the filter used by the operation matches multiple
	// documents in the collection. If set, the first document in the sorted order will be replaced. The driver will
	// return an error if the sort parameter is a multi-key map. The default value is nil.
	Sort any
	// If true, a new document will be inserted if the filter does not match any documents in the collection. The
	// default value is false.
	Upsert bool
	// The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.4. MongoDB version 4.2 will report an error if
	// this option is specified. For server versions < 4.2, the driver will return an error if this option is specified.
	// The driver will return an error if this option is used with during an unacknowledged write operation. The driver
	// will return an error if the hint parameter is a multi-key map. The default value is nil, which means that no hint
	// will be sent.
	Hint any
	// Specifies parameters for the find one and replace expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let                     any
	DisableAutoCloseSession bool
}

type FindOneAndUpdate struct {
	// A set of filters specifying to which array elements an update should apply. This option is only valid for MongoDB
	// versions >= 3.6. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the update will apply to all array elements.
	ArrayFilters *ArrayFilters
	// If true, writes executed as part of the operation will opt out of document-level validation on the server. This
	// option is valid for MongoDB versions >= 3.2 and is ignored for previous server versions. The default value is
	// false. See https://www.mongodb.com/docs/manual/core/schema-validation/ for more information about document
	// validation.
	BypassDocumentValidation bool
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default value is nil, which means that no comment will be included in the logs.
	Comment any
	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there
	// is no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used
	// in its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime time.Duration
	// A document describing which fields will be included in the document returned by the operation. The default value
	// is nil, which means all fields will be included.
	Projection any
	// Specifies whether the original or replaced document should be returned by the operation. The default value is
	// Before, which means the original document will be returned before the replacement is performed.
	ReturnDocument ReturnDocument
	// A document specifying which document should be updated if the filter used by the operation matches multiple
	// documents in the collection. If set, the first document in the sorted order will be updated. The driver will
	// return an error if the sort parameter is a multi-key map. The default value is nil.
	Sort any
	// If true, a new document will be inserted if the filter does not match any documents in the collection. The
	// default value is false.
	Upsert bool
	// The index to use for the operation. This should either be the index name as a string or the index specification
	// as a document. This option is only valid for MongoDB versions >= 4.4. MongoDB version 4.2 will report an error if
	// this option is specified. For server versions < 4.2, the driver will return an error if this option is specified.
	// The driver will return an error if this option is used with during an unacknowledged write operation. The driver
	// will return an error if the hint parameter is a multi-key map. The default value is nil, which means that no hint
	// will be sent.
	Hint any
	// Specifies parameters for the find one and update expression. This option is only valid for MongoDB versions >= 5.0. Older
	// servers will report an error for using this option. This must be a document mapping parameter names to values.
	// Values must be constant or closed expressions that do not reference document fields. Parameters can then be
	// accessed as variables in an aggregate expression context (e.g. "$$var").
	Let                     any
	DisableAutoCloseSession bool
}

func NewFind() Find {
	return Find{}
}

func NewFindPageable() FindPageable {
	return FindPageable{}
}

func NewFindOne() FindOne {
	return FindOne{}
}

func NewFindOneAndDelete() FindOneAndDelete {
	return FindOneAndDelete{}
}

func NewFindOneAndReplace() FindOneAndReplace {
	return FindOneAndReplace{}
}

func NewFindOneAndUpdate() FindOneAndUpdate {
	return FindOneAndUpdate{}
}

func (f Find) SetAllowDiskUse(b bool) Find {
	f.AllowDiskUse = b
	return f
}

func (f Find) SetAllowPartialResults(b bool) Find {
	f.AllowPartialResults = b
	return f
}

func (f Find) SetBatchSize(i int32) Find {
	f.BatchSize = i
	return f
}

func (f Find) SetCollation(c *Collation) Find {
	f.Collation = c
	return f
}

func (f Find) SetComment(s string) Find {
	f.Comment = s
	return f
}

func (f Find) SetCursorType(c CursorType) Find {
	f.CursorType = c
	return f
}

func (f Find) SetHint(v any) Find {
	f.Hint = v
	return f
}

func (f Find) SetLimit(i int64) Find {
	f.Limit = i
	return f
}

func (f Find) SetMax(v any) Find {
	f.Max = v
	return f
}

func (f Find) SetMaxAwaitTime(d time.Duration) Find {
	f.MaxAwaitTime = d
	return f
}

func (f Find) SetMaxTime(d time.Duration) Find {
	f.MaxTime = d
	return f
}

func (f Find) SetMin(v any) Find {
	f.Min = v
	return f
}

func (f Find) SetNoCursorTimeout(b bool) Find {
	f.NoCursorTimeout = b
	return f
}

func (f Find) SetProjection(v any) Find {
	f.Projection = v
	return f
}

func (f Find) SetReturnKey(b bool) Find {
	f.ReturnKey = b
	return f
}

func (f Find) SetShowRecordID(b bool) Find {
	f.ShowRecordID = b
	return f
}

func (f Find) SetSkip(i int64) Find {
	f.Skip = i
	return f
}

func (f Find) SetSort(i int64) Find {
	f.Sort = i
	return f
}

func (f Find) SetLet(v any) Find {
	f.Let = v
	return f
}

func (f FindPageable) SetAllowDiskUse(b bool) FindPageable {
	f.AllowDiskUse = b
	return f
}

func (f FindPageable) SetAllowPartialResults(b bool) FindPageable {
	f.AllowPartialResults = b
	return f
}

func (f FindPageable) SetBatchSize(i int32) FindPageable {
	f.BatchSize = i
	return f
}

func (f FindPageable) SetCollation(c *Collation) FindPageable {
	f.Collation = c
	return f
}

func (f FindPageable) SetComment(s string) FindPageable {
	f.Comment = s
	return f
}

func (f FindPageable) SetCursorType(c CursorType) FindPageable {
	f.CursorType = c
	return f
}

func (f FindPageable) SetHint(v any) FindPageable {
	f.Hint = v
	return f
}

func (f FindPageable) SetMax(v any) FindPageable {
	f.Max = v
	return f
}

func (f FindPageable) SetMaxAwaitTime(d time.Duration) FindPageable {
	f.MaxAwaitTime = d
	return f
}

func (f FindPageable) SetMaxTime(d time.Duration) FindPageable {
	f.MaxTime = d
	return f
}

func (f FindPageable) SetMin(v any) FindPageable {
	f.Min = v
	return f
}

func (f FindPageable) SetNoCursorTimeout(b bool) FindPageable {
	f.NoCursorTimeout = b
	return f
}

func (f FindPageable) SetProjection(v any) FindPageable {
	f.Projection = v
	return f
}

func (f FindPageable) SetReturnKey(b bool) FindPageable {
	f.ReturnKey = b
	return f
}

func (f FindPageable) SetShowRecordID(b bool) FindPageable {
	f.ShowRecordID = b
	return f
}

func (f FindPageable) SetLet(v any) FindPageable {
	f.Let = v
	return f
}

func (f FindOne) SetAllowPartialResults(b bool) FindOne {
	f.AllowPartialResults = b
	return f
}

func (f FindOne) SetCollation(c *Collation) FindOne {
	f.Collation = c
	return f
}

func (f FindOne) SetComment(s string) FindOne {
	f.Comment = s
	return f
}

func (f FindOne) SetHint(v any) FindOne {
	f.Hint = v
	return f
}

func (f FindOne) SetMax(v any) FindOne {
	f.Max = v
	return f
}

func (f FindOne) SetMaxTime(d time.Duration) FindOne {
	f.MaxTime = d
	return f
}

func (f FindOne) SetMin(v any) FindOne {
	f.Min = v
	return f
}

func (f FindOne) SetProjection(v any) FindOne {
	f.Projection = v
	return f
}

func (f FindOne) SetReturnKey(b bool) FindOne {
	f.ReturnKey = b
	return f
}

func (f FindOne) SetShowRecordID(b bool) FindOne {
	f.ShowRecordID = b
	return f
}

func (f FindOne) SetSkip(i int64) FindOne {
	f.Skip = i
	return f
}

func (f FindOne) SetSort(i int64) FindOne {
	f.Sort = i
	return f
}

func (f FindOneAndDelete) SetCollation(c *Collation) FindOneAndDelete {
	f.Collation = c
	return f
}

func (f FindOneAndDelete) SetComment(s string) FindOneAndDelete {
	f.Comment = s
	return f
}

func (f FindOneAndDelete) SetHint(v any) FindOneAndDelete {
	f.Hint = v
	return f
}

func (f FindOneAndDelete) SetMaxTime(d time.Duration) FindOneAndDelete {
	f.MaxTime = d
	return f
}

func (f FindOneAndDelete) SetProjection(v any) FindOneAndDelete {
	f.Projection = v
	return f
}

func (f FindOneAndDelete) SetSort(i int64) FindOneAndDelete {
	f.Sort = i
	return f
}

func (f FindOneAndDelete) SetLet(v any) FindOneAndDelete {
	f.Let = v
	return f
}

func (f FindOneAndDelete) SetDisableAutoCloseTransaction(b bool) FindOneAndDelete {
	f.DisableAutoCloseSession = b
	return f
}

func (f FindOneAndReplace) SetDisableAutoCloseTransaction(b bool) FindOneAndReplace {
	f.DisableAutoCloseSession = b
	return f
}

func (f FindOneAndReplace) SetBypassDocumentValidation(b bool) FindOneAndReplace {
	f.BypassDocumentValidation = b
	return f
}

func (f FindOneAndReplace) SetCollation(c *Collation) FindOneAndReplace {
	f.Collation = c
	return f
}

func (f FindOneAndReplace) SetComment(s string) FindOneAndReplace {
	f.Comment = s
	return f
}

func (f FindOneAndReplace) SetHint(v any) FindOneAndReplace {
	f.Hint = v
	return f
}

func (f FindOneAndReplace) SetMaxTime(d time.Duration) FindOneAndReplace {
	f.MaxTime = d
	return f
}

func (f FindOneAndReplace) SetProjection(v any) FindOneAndReplace {
	f.Projection = v
	return f
}

func (f FindOneAndReplace) SetReturnDocument(r ReturnDocument) FindOneAndReplace {
	f.ReturnDocument = r
	return f
}

func (f FindOneAndReplace) SetUpsert(b bool) FindOneAndReplace {
	f.Upsert = b
	return f
}

func (f FindOneAndReplace) SetLet(v any) FindOneAndReplace {
	f.Let = v
	return f
}

func (f FindOneAndReplace) SetSort(i int64) FindOneAndReplace {
	f.Sort = i
	return f
}

func (f FindOneAndUpdate) SetDisableAutoCloseTransaction(b bool) FindOneAndUpdate {
	f.DisableAutoCloseSession = b
	return f
}

func (f FindOneAndUpdate) SetArrayFilters(a *ArrayFilters) FindOneAndUpdate {
	f.ArrayFilters = a
	return f
}

func (f FindOneAndUpdate) SetBypassDocumentValidation(b bool) FindOneAndUpdate {
	f.BypassDocumentValidation = b
	return f
}

func (f FindOneAndUpdate) SetCollation(c *Collation) FindOneAndUpdate {
	f.Collation = c
	return f
}

func (f FindOneAndUpdate) SetComment(a any) FindOneAndUpdate {
	f.Comment = a
	return f
}

func (f FindOneAndUpdate) SetHint(v any) FindOneAndUpdate {
	f.Hint = v
	return f
}

func (f FindOneAndUpdate) SetLet(v any) FindOneAndUpdate {
	f.Let = v
	return f
}

func (f FindOneAndUpdate) SetMaxTime(d time.Duration) FindOneAndUpdate {
	f.MaxTime = d
	return f
}

func (f FindOneAndUpdate) SetProjection(v any) FindOneAndUpdate {
	f.Projection = v
	return f
}

func (f FindOneAndUpdate) SetReturnDocument(r ReturnDocument) FindOneAndUpdate {
	f.ReturnDocument = r
	return f
}

func (f FindOneAndUpdate) SetSort(i int64) FindOneAndUpdate {
	f.Sort = i
	return f
}

func (f FindOneAndUpdate) SetUpsert(b bool) FindOneAndUpdate {
	f.Upsert = b
	return f
}

func GetFindOptionByParams(opts []Find) Find {
	result := Find{}
	for _, opt := range opts {
		if opt.AllowDiskUse {
			result.AllowDiskUse = opt.AllowDiskUse
		}
		if opt.AllowPartialResults {
			result.AllowPartialResults = opt.AllowPartialResults
		}
		if opt.NoCursorTimeout {
			result.NoCursorTimeout = opt.NoCursorTimeout
		}
		if opt.ReturnKey {
			result.ReturnKey = opt.ReturnKey
		}
		if opt.ShowRecordID {
			result.ShowRecordID = opt.ShowRecordID
		}
		if opt.CursorType.IsEnumValid() {
			result.CursorType = opt.CursorType
		}
		if opt.BatchSize != 0 {
			result.BatchSize = opt.BatchSize
		}
		if opt.Collation != nil {
			result.Collation = opt.Collation
		}
		if len(opt.Comment) != 0 {
			result.Comment = opt.Comment
		}
		if opt.Hint != nil {
			result.Hint = opt.Hint
		}
		if opt.Let != nil {
			result.Let = opt.Let
		}
		if opt.Limit > 0 {
			result.Limit = opt.Limit
		}
		if opt.Skip > 0 {
			result.Skip = opt.Skip
		}
		if opt.Max != nil {
			result.Max = opt.Max
		}
		if opt.Min != nil {
			result.Min = opt.Min
		}
		if opt.Sort != nil {
			result.Sort = opt.Sort
		}
		if opt.Projection != nil {
			result.Projection = opt.Projection
		}
		if opt.MaxTime > 0 {
			result.MaxTime = opt.MaxTime
		}
		if opt.MaxAwaitTime > 0 {
			result.MaxAwaitTime = opt.MaxAwaitTime
		}
	}
	return result
}

func GetFindPageableOptionByParams(opts []FindPageable) FindPageable {
	result := FindPageable{}
	for _, opt := range opts {
		if opt.AllowDiskUse {
			result.AllowDiskUse = opt.AllowDiskUse
		}
		if opt.AllowPartialResults {
			result.AllowPartialResults = opt.AllowPartialResults
		}
		if opt.NoCursorTimeout {
			result.NoCursorTimeout = opt.NoCursorTimeout
		}
		if opt.ReturnKey {
			result.ReturnKey = opt.ReturnKey
		}
		if opt.ShowRecordID {
			result.ShowRecordID = opt.ShowRecordID
		}
		if opt.CursorType.IsEnumValid() {
			result.CursorType = opt.CursorType
		}
		if opt.BatchSize != 0 {
			result.BatchSize = opt.BatchSize
		}
		if opt.Collation != nil {
			result.Collation = opt.Collation
		}
		if len(opt.Comment) != 0 {
			result.Comment = opt.Comment
		}
		if opt.Hint != nil {
			result.Hint = opt.Hint
		}
		if opt.Let != nil {
			result.Let = opt.Let
		}
		if opt.Max != nil {
			result.Max = opt.Max
		}
		if opt.Min != nil {
			result.Min = opt.Min
		}
		if opt.Projection != nil {
			result.Projection = opt.Projection
		}
		if opt.MaxTime > 0 {
			result.MaxTime = opt.MaxTime
		}
		if opt.MaxAwaitTime > 0 {
			result.MaxAwaitTime = opt.MaxAwaitTime
		}
	}
	return result
}

func GetFindOneOptionByParams(opts []FindOne) FindOne {
	result := FindOne{}
	for _, opt := range opts {
		if opt.AllowPartialResults {
			result.AllowPartialResults = opt.AllowPartialResults
		}
		if opt.ReturnKey {
			result.ReturnKey = opt.ReturnKey
		}
		if opt.ShowRecordID {
			result.ShowRecordID = opt.ShowRecordID
		}
		if opt.Collation != nil {
			result.Collation = opt.Collation
		}
		if len(opt.Comment) != 0 {
			result.Comment = opt.Comment
		}
		if opt.Hint != nil {
			result.Hint = opt.Hint
		}
		if opt.Skip > 0 {
			result.Skip = opt.Skip
		}
		if opt.Max != nil {
			result.Max = opt.Max
		}
		if opt.Min != nil {
			result.Min = opt.Min
		}
		if opt.Sort != nil {
			result.Sort = opt.Sort
		}
		if opt.Projection != nil {
			result.Projection = opt.Projection
		}
		if opt.MaxTime > 0 {
			result.MaxTime = opt.MaxTime
		}
	}
	return result
}

func GetFindOneAndDeleteOptionByParams(opts []FindOneAndDelete) FindOneAndDelete {
	result := FindOneAndDelete{}
	for _, opt := range opts {
		if opt.DisableAutoCloseSession {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
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
		if opt.Sort != nil {
			result.Sort = opt.Sort
		}
		if opt.Projection != nil {
			result.Projection = opt.Projection
		}
		if opt.MaxTime > 0 {
			result.MaxTime = opt.MaxTime
		}
	}
	return result
}

func GetFindOneAndReplaceOptionByParams(opts []FindOneAndReplace) FindOneAndReplace {
	result := FindOneAndReplace{}
	for _, opt := range opts {
		if opt.BypassDocumentValidation {
			result.BypassDocumentValidation = opt.BypassDocumentValidation
		}
		if opt.DisableAutoCloseSession {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
		}
		if opt.Upsert {
			result.Upsert = opt.Upsert
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
		if opt.Sort != nil {
			result.Sort = opt.Sort
		}
		if opt.Projection != nil {
			result.Projection = opt.Projection
		}
		if opt.MaxTime > 0 {
			result.MaxTime = opt.MaxTime
		}
		if opt.ReturnDocument.IsEnumValid() {
			result.ReturnDocument = opt.ReturnDocument
		}
	}
	return result
}

func GetFindOneAndUpdateOptionByParams(opts []FindOneAndUpdate) FindOneAndUpdate {
	result := FindOneAndUpdate{}
	for _, opt := range opts {
		if opt.BypassDocumentValidation {
			result.BypassDocumentValidation = opt.BypassDocumentValidation
		}
		if opt.DisableAutoCloseSession {
			result.DisableAutoCloseSession = opt.DisableAutoCloseSession
		}
		if opt.Upsert {
			result.Upsert = opt.Upsert
		}
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
		if opt.Sort != nil {
			result.Sort = opt.Sort
		}
		if opt.Projection != nil {
			result.Projection = opt.Projection
		}
		if opt.MaxTime > 0 {
			result.MaxTime = opt.MaxTime
		}
		if opt.ReturnDocument.IsEnumValid() {
			result.ReturnDocument = opt.ReturnDocument
		}
	}
	return result
}
