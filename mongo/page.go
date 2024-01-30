package mongo

import (
	"github.com/GabrielHCataldo/go-errors/errors"
	"github.com/GabrielHCataldo/go-helper/helper"
	"math"
	"time"
)

type PageInput struct {
	Page     int64
	PageSize int64
	Ref      any
	Sort     any
}

type PageResult struct {
	Page          int64       `json:"page"`
	PageSize      int64       `json:"pageSize"`
	PageTotal     int64       `json:"pageTotal"`
	TotalElements int64       `json:"totalElements"`
	Content       pageContent `json:"content,omitempty"`
	LastQueryAt   time.Time   `json:"lastQueryAt,omitempty"`
}

type pageContent []pageItemContent
type pageItemContent map[string]any

func (p pageContent) Parse(dest any) error {
	if helper.IsNotSlice(dest) {
		return errors.NewSkipCaller(2, "mongo: dest is not a slice or array")
	}
	err := helper.ConvertToDest(p, dest)
	return errors.NewSkipCaller(2, err)
}

func (p pageItemContent) Parse(dest any) error {
	err := helper.ConvertToDest(p, dest)
	return errors.NewSkipCaller(2, err)
}

func newPageResult(pageInput PageInput, result any, countTotal int64) *PageResult {
	minPageTotal := 1
	if helper.IsEmpty(result) {
		minPageTotal = 0
	}
	fPageTotal := math.Ceil(float64(countTotal) / float64(pageInput.PageSize))
	pageTotal := helper.MinInt(int(fPageTotal), minPageTotal)
	lastQueryAt := time.Now().UTC()
	var content pageContent
	_ = helper.ConvertToDest(result, &content)
	return &PageResult{
		Page:          pageInput.Page,
		PageSize:      pageInput.PageSize,
		PageTotal:     int64(pageTotal),
		TotalElements: countTotal,
		Content:       content,
		LastQueryAt:   lastQueryAt,
	}
}
