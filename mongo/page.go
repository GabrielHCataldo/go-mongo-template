package mongo

import (
	"github.com/GabrielHCataldo/go-mongo/internal/util"
	"math"
	"time"
)

type PageInput struct {
	Page     int64
	PageSize int64
	Ref      any
	Sort     any
}

type PageOutput struct {
	Page          int64     `json:"page"`
	PageSize      int64     `json:"pageSize"`
	PageTotal     int64     `json:"pageTotal"`
	TotalElements int64     `json:"totalElements"`
	Content       any       `json:"content,omitempty"`
	LastQueryAt   time.Time `json:"lastQueryAt,omitempty"`
}

func NewPageOutput(pageInput PageInput, result any, countTotal int64) *PageOutput {
	minPageTotal := 1
	if util.IsZero(result) {
		minPageTotal = 0
	}
	pageTotal := util.MinInt64(int64(math.Ceil(float64(countTotal)/float64(pageInput.PageSize))), int64(minPageTotal))
	lastQueryAt := time.Now().UTC()
	return &PageOutput{
		Page:          pageInput.Page,
		PageSize:      pageInput.PageSize,
		PageTotal:     pageTotal,
		TotalElements: countTotal,
		Content:       result,
		LastQueryAt:   lastQueryAt,
	}
}
