package option

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Watch struct {
	DatabaseName   string
	CollectionName string
	// Duration time to process the func watch, timeout applied in the past context.
	//
	// default: 5 seconds
	ContextFuncTimeout time.Duration
	// Delay to run the next search for messages in the queue
	//
	// default: 5 seconds
	DelayLoop time.Duration
	// The maximum number of documents to be included in each batch returned by the server.
	BatchSize *int32
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string that will be included in server logs, profiling logs, and currentOp queries to help trace the operation.
	// The default is nil, which means that no comment will be included in the logs.
	Comment *string
	// Specifies how the updated document should be returned in change notifications for update operations. The default
	// is options.Default, which means that only partial update deltas will be included in the change notification.
	FullDocument *FullDocument
	// Specifies how the pre-update document should be returned in change notifications for update operations. The default
	// is options.Off, which means that the pre-update document will not be included in the change notification.
	FullDocumentBeforeChange *FullDocument
	// The maximum amount of time that the server should wait for new documents to satisfy a tailable cursor query.
	MaxAwaitTime *time.Duration
	// A document specifying the logical starting point for the change stream. Only changes corresponding to an oplog
	// entry immediately after the resume token will be returned. If this is specified, StartAtOperationTime and
	// StartAfter must not be set.
	ResumeAfter any
	// ShowExpandedEvents specifies whether the server will return an expanded list of change stream events. Additional
	// events include: createIndexes, dropIndexes, modify, create, shardCollection, reshardCollection and
	// refineCollectionShardKey. This option is only valid for MongoDB versions >= 6.0.
	ShowExpandedEvents *bool
	// If specified, the change stream will only return changes that occurred at or after the given timestamp. This
	// option is only valid for MongoDB versions >= 4.0. If this is specified, ResumeAfter and StartAfter must not be
	// set.
	StartAtOperationTime *primitive.Timestamp
	// A document specifying the logical starting point for the change stream. This is similar to the ResumeAfter
	// option, but allows a resume token from an "invalidate" notification to be used. This allows a change stream on a
	// collection to be resumed after the collection has been dropped and recreated or renamed. Only changes
	// corresponding to an oplog entry immediately after the specified token will be returned. If this is specified,
	// ResumeAfter and StartAtOperationTime must not be set. This option is only valid for MongoDB versions >= 4.1.1.
	StartAfter any
	// Custom options to be added to the initial aggregate for the change stream. Key-value pairs of the BSON map should
	// correlate with desired option names and values. Values must be Marshalable. Custom options may conflict with
	// non-custom options, and custom options bypass client-side validation. Prefer using non-custom options where possible.
	Custom bson.M
	// Custom options to be added to the $changeStream stage in the initial aggregate. Key-value pairs of the BSON map should
	// correlate with desired option names and values. Values must be Marshalable. Custom pipeline options bypass client-side
	// validation. Prefer using non-custom options where possible.
	CustomPipeline bson.M
}

func NewWatch() Watch {
	return Watch{}
}

func (w Watch) SetDatabaseName(s string) Watch {
	w.DatabaseName = s
	return w
}

func (w Watch) SetCollectionName(s string) Watch {
	w.CollectionName = s
	return w
}

func (w Watch) SetContextFuncTimeout(d time.Duration) Watch {
	w.ContextFuncTimeout = d
	return w
}

func (w Watch) SetDelayLoop(d time.Duration) Watch {
	w.DelayLoop = d
	return w
}

func (w Watch) SetBatchSize(i int32) Watch {
	w.BatchSize = &i
	return w
}

func (w Watch) SetCollation(c *Collation) Watch {
	w.Collation = c
	return w
}

func (w Watch) SetComment(s string) Watch {
	w.Comment = &s
	return w
}

func (w Watch) SetFullDocument(f FullDocument) Watch {
	w.FullDocument = &f
	return w
}

func (w Watch) SetFullDocumentBeforeChange(f FullDocument) Watch {
	w.FullDocumentBeforeChange = &f
	return w
}

func (w Watch) SetMaxAwaitTime(d time.Duration) Watch {
	w.MaxAwaitTime = &d
	return w
}

func (w Watch) SetResumeAfter(a any) Watch {
	w.ResumeAfter = a
	return w
}

func (w Watch) SetShowExpandedEvents(b bool) Watch {
	w.ShowExpandedEvents = &b
	return w
}

func (w Watch) SetStartAtOperationTime(t primitive.Timestamp) Watch {
	w.StartAtOperationTime = &t
	return w
}

func (w Watch) SetStartAfter(a any) Watch {
	w.StartAfter = a
	return w
}

func (w Watch) SetCustom(b bson.M) Watch {
	w.Custom = b
	return w
}

func (w Watch) SetCustomPipeline(b bson.M) Watch {
	w.CustomPipeline = b
	return w
}

func GetWatchOptionByParams(opts []Watch) Watch {
	result := Watch{}
	for _, opt := range opts {
		if opt.ContextFuncTimeout > 0 {
			result.ContextFuncTimeout = opt.ContextFuncTimeout
		}
		if opt.DelayLoop > 0 {
			result.DelayLoop = opt.DelayLoop
		}
		if opt.BatchSize != nil {
			result.BatchSize = opt.BatchSize
		}
		if opt.Collation != nil {
			result.Collation = opt.Collation
		}
		if opt.Comment != nil {
			result.Comment = opt.Comment
		}
		if len(opt.DatabaseName) != 0 {
			result.DatabaseName = opt.DatabaseName
		}
		if len(opt.CollectionName) != 0 {
			result.CollectionName = opt.CollectionName
		}
		if opt.MaxAwaitTime != nil {
			result.MaxAwaitTime = opt.MaxAwaitTime
		}
		if opt.Custom != nil {
			result.Custom = opt.Custom
		}
		if opt.ResumeAfter != nil {
			result.ResumeAfter = opt.ResumeAfter
		}
		if opt.FullDocument != nil {
			result.FullDocument = opt.FullDocument
		}
		if opt.FullDocumentBeforeChange != nil {
			result.FullDocumentBeforeChange = opt.FullDocumentBeforeChange
		}
		if opt.ShowExpandedEvents != nil {
			result.ShowExpandedEvents = opt.ShowExpandedEvents
		}
		if opt.StartAtOperationTime != nil {
			result.StartAtOperationTime = opt.StartAtOperationTime
		}
		if opt.StartAfter != nil {
			result.StartAfter = opt.StartAfter
		}
		if opt.Custom != nil {
			result.Custom = opt.Custom
		}
		if opt.CustomPipeline != nil {
			result.CustomPipeline = opt.CustomPipeline
		}
	}
	return result
}
