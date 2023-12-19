package option

type Enum interface {
	IsEnumValid() bool
}

// CursorType specifies whether a cursor should close when the last data is retrieved. See
// NonTailable, Tailable, and TailableAwait.
type CursorType int8

// ReturnDocument specifies whether a findAndUpdate operation should return the document as it was
// before the update or as it is after the update.
type ReturnDocument int8

// FullDocument specifies how a change stream should return the modified document.
type FullDocument string

const (
	// FullDocumentDefault does not include a document copy.
	FullDocumentDefault FullDocument = "default"
	// FullDocumentOff is the same as sending no value for fullDocumentBeforeChange.
	FullDocumentOff FullDocument = "off"
	// FullDocumentRequired is the same as WhenAvailable but raises a server-side error if the post-image is not available.
	FullDocumentRequired FullDocument = "required"
	// FullDocumentUpdateLookup includes a delta describing the changes to the document and a copy of the entire document that
	// was changed.
	FullDocumentUpdateLookup FullDocument = "updateLookup"
	// FullDocumentWhenAvailable includes a post-image of the modified document for replace and update change events
	// if the post-image for this event is available.
	FullDocumentWhenAvailable FullDocument = "whenAvailable"
)
const (
	// ReturnDocumentBefore specifies that findAndUpdate should return the document as it was before the update.
	ReturnDocumentBefore ReturnDocument = iota
	// ReturnDocumentAfter specifies that findAndUpdate should return the document as it is after the update.
	ReturnDocumentAfter
)
const (
	// CursorTypeNonTailable specifies that a cursor should close after retrieving the last data.
	CursorTypeNonTailable CursorType = iota
	// CursorTypeTailable specifies that a cursor should not close when the last data is retrieved and can be resumed later.
	CursorTypeTailable
	// CursorTypeTailableAwait specifies that a cursor should not close when the last data is retrieved and
	// that it should block for a certain amount of time for new data before returning no data.
	CursorTypeTailableAwait
)

func (c CursorType) IsEnumValid() bool {
	switch c {
	case CursorTypeNonTailable, CursorTypeTailable, CursorTypeTailableAwait:
		return true
	}
	return false
}

func (r ReturnDocument) IsEnumValid() bool {
	switch r {
	case ReturnDocumentBefore, ReturnDocumentAfter:
		return true
	}
	return false
}

func (f FullDocument) IsEnumValid() bool {
	switch f {
	case FullDocumentDefault, FullDocumentOff, FullDocumentRequired, FullDocumentUpdateLookup, FullDocumentWhenAvailable:
		return true
	}
	return false
}
