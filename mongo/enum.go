package mongo

type Enum interface {
	IsEnumValid() bool
}

type Sort int

const (
	SortDesc Sort = -1
	SortAsc  Sort = 1
)

func (m Sort) IsEnumValid() bool {
	switch m {
	case SortDesc, SortAsc:
		return true
	}
	return false
}
