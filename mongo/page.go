package mongo

import (
	"github.com/GabrielHCataldo/go-errors/errors"
	"github.com/GabrielHCataldo/go-helper/helper"
	"math"
	"time"
)

type PageInput struct {
	// Page current page (default 0)
	Page int64
	// PageSize page size (required)
	PageSize int64
	// Ref slice of the struct reference contained database and collection configured
	Ref any
	// Sort value sort to result
	Sort any
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

// Decode parse pageResult to dest param
func (p PageResult) Decode(dest any) error {
	err := helper.ConvertToDest(p, dest)
	return errors.NewSkipCaller(2, err)
}

// Decode parse pageResult.Content to dest param
func (p pageContent) Decode(dest any) error {
	err := helper.ConvertToDest(p, dest)
	return errors.NewSkipCaller(2, err)
}

// Decode parse pageResult item to dest param
func (p pageItemContent) Decode(dest any) error {
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
