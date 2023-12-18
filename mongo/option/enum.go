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

func (c ReturnDocument) IsEnumValid() bool {
	switch c {
	case ReturnDocumentBefore, ReturnDocumentAfter:
		return true
	}
	return false
}
