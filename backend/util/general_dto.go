package util

import (
	"strings"

	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

// list response

func NewListResponseWithPagination[T any](list []T, page *PageResponse) ListResponse[T] {
	if list == nil {
		list = make([]T, 0)
	}
	return ListResponse[T]{
		PageInfo: page,
		List:     list,
	}
}

func NewListResponseWithoutPagination[T any](list []T) ListResponse[T] {
	if list == nil {
		list = make([]T, 0)
	}
	return ListResponse[T]{
		PageInfo: nil,
		List:     list,
	}
}

type ListResponse[T any] struct {
	PageInfo *PageResponse `json:"page_info,omitempty"`
	List     []T           `json:"list"`
}

func (l *ListResponse[T]) IsPagination() bool {
	return l.PageInfo != nil
}

// sort

type SortParam struct {
	SortKey string   `json:"sort_key" form:"sort_key"`
	SortV   SortKind `json:"sort_v" form:"sort_v" validate:"sort"`
}

func (dto *SortParam) SetDefaultIfInvalid(sortKey string, sortV SortKind) {
	if dto.SortKey == "" {
		dto.SortKey = sortKey
	}
	if dto.SortV == SortDesc || dto.SortV == SortAsc {
		return
	}
	dto.SortV = sortV
}

var sortBuilder = strings.Builder{}

func (dto SortParam) KeyValueString() string {
	defer sortBuilder.Reset()
	sortBuilder.WriteString(dto.SortKey)
	sortBuilder.WriteString(" ")
	sortBuilder.WriteString(string(dto.SortV))
	return sortBuilder.String()
}

const (
	SortNone SortKind = ""
	SortDesc SortKind = "desc"
	SortAsc  SortKind = "asc"
)

type SortKind string

func (s SortKind) Validate() error {
	value := SortKind(strings.ToLower(string(s)))

	switch value {
	case SortNone, SortDesc, SortAsc:
		return nil
	default:
		return errors.WrapWithMessage(errors.ErrInvalidParams, "not match sort kind: value = %v", value)
	}
}

func (s *SortKind) UnmarshalText(text []byte) error {
	*s = SortKind(text)
	return s.Validate()
}

// page

func NewWithoutPagination() PageParam {
	return PageParam{}
}

type PageParam struct {
	Page uint64 `json:"page" form:"page"`
	Size uint64 `json:"size" form:"size"`
}

func (p PageParam) Validate() error {
	if p.Page > 0 && p.Size > 0 {
		return nil
	}
	return errors.WrapWithMessage(errors.ErrDeveloper, "page dto is invalid")
}

func (p PageParam) IsPagination() bool {
	return p.Page == 0 || p.Size == 0
}

func (p *PageParam) SetWithoutPagination() {
	p.Page = 0
	p.Size = 0
}

func (p *PageParam) SetDefaultIfInvalid() {
	const (
		defaultMax  = 3000
		defaultSize = 20
	)
	p.SetDefaultAndMaxSizeIfInvalid(defaultSize, defaultMax)
}

func (p *PageParam) SetDefaultAndMaxSizeIfInvalid(defaultSize, maxPageSize uint64) {
	const (
		defaultPage = 1
	)

	err := p.Validate()
	if err != nil {
		p.Page = defaultPage
		p.Size = defaultSize
		return
	}

	if p.Size > maxPageSize {
		p.Size = maxPageSize
	}
}

func (p PageParam) OffsetOrSkip() int64 {
	return int64((p.Page - 1) * (p.Size))
}

func NewPageResponse(param PageParam, totalSize uint64) PageResponse {
	quotient := totalSize / param.Size
	remainder := totalSize % param.Size

	var totalPage uint64
	switch {
	case quotient > 0 && remainder == 0:
		totalPage = quotient

	case quotient > 0 && remainder != 0:
		totalPage = quotient + 1

	case quotient <= 0:
		totalPage = 1
	}

	var size uint64
	switch {
	case totalPage > param.Page:
		size = param.Size

	case totalPage == param.Page && totalSize != 0 && remainder == 0:
		size = param.Size

	case totalPage == param.Page && totalSize != 0 && remainder != 0:
		size = remainder

	case totalPage == param.Page && totalSize == 0:
		size = 0

	case totalPage < param.Page:
		size = 0
	}

	return PageResponse{
		TotalPage: totalPage,
		TotalSize: totalSize,
		Page:      param.Page,
		Size:      size,
	}
}

type PageResponse struct {
	TotalPage uint64 `json:"total_page"`
	TotalSize uint64 `json:"total_size"`
	Page      uint64 `json:"page"`
	Size      uint64 `json:"size"`
}
