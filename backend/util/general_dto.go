package util

import (
	"reflect"
	"strings"

	"github.com/fatih/structs"

	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

// list query param

type DtoQryParam interface {
	FilterParam() any
	SortParam() any
	PageParam() DtoPageParam
}

// list response

func NewDtoListResponse[T any](list []T, page *DtoPageResponse) DtoListResponse[T] {
	if list == nil {
		list = make([]T, 0)
	}
	return DtoListResponse[T]{
		DtoPageResponse: page,
		List:            list,
	}
}

type DtoListResponse[T any] struct {
	*DtoPageResponse
	List []T `json:"list"`
}

// sort

const (
	SortNone SortKind = ""
	SortDesc SortKind = "desc"
	SortAsc  SortKind = "asc"
)

type SortKind string

func (s SortKind) IsValid() bool {
	err := s.validate()
	if err != nil {
		return false
	}
	return true
}

func (s SortKind) validate() error {
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
	return s.validate()
}

func ValidateSort(sort any) error {
	if sort == nil {
		return nil
	}

	fields := structs.New(sort).Fields()

	for _, field := range fields {
		if field.IsZero() {
			continue
		}

		var value any
		if field.Kind() == reflect.Pointer {
			value = reflect.ValueOf(field.Value()).Elem().Interface()
		} else {
			value = field.Value()
		}

		fn, ok := value.(SortKind)
		if !ok {
			continue
		}

		err := fn.validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// page

type DtoPageParam struct {
	Page *int64 `json:"page" form:"page"`
	Size *int64 `json:"size" form:"size"`
}

func (p DtoPageParam) RestrictedSize() bool {
	return p.IsValueAssigned()
}

func (p DtoPageParam) IsValueAssigned() bool {
	return !(p.Page == nil && p.Size == nil)
}

func (p DtoPageParam) Validate() error {
	if p.Page != nil && p.Size != nil && *p.Page > 0 && *p.Size > 0 {
		return nil
	}
	return errors.WrapWithMessage(errors.ErrDeveloper, "page param is invalid")
}

func (p *DtoPageParam) SetDefault(maxPageSize int64) {
	var (
		defaultPage int64 = 1
		defaultSize int64 = 20
	)

	err := p.Validate()
	if err != nil {
		p.Page = &defaultPage
		p.Size = &defaultSize
		return
	}

	if maxPageSize > 0 && *p.Size > maxPageSize {
		*p.Size = maxPageSize
	}
}

func (p DtoPageParam) OffsetOrSkip() int64 {
	return (*p.Page - 1) * (*p.Size)
}

func NewDtoPageResponse(param DtoPageParam, totalSize int64) (*DtoPageResponse, error) {
	err := param.Validate()
	if err != nil {
		return nil, err
	}

	quotient := totalSize / *param.Size
	remainder := totalSize % *param.Size

	var totalPage int64
	switch {
	case quotient > 0 && remainder == 0:
		totalPage = quotient

	case quotient > 0 && remainder != 0:
		totalPage = quotient + 1

	case quotient <= 0:
		totalPage = 1
	}

	var size int64
	switch {
	case totalPage > *param.Page:
		size = *param.Size

	case totalPage == *param.Page && totalSize != 0 && remainder == 0:
		size = *param.Size

	case totalPage == *param.Page && totalSize != 0 && remainder != 0:
		size = remainder

	case totalPage == *param.Page && totalSize == 0:
		size = 0

	case totalPage < *param.Page:
		size = 0
	}

	return &DtoPageResponse{
		TotalPage: &totalPage,
		TotalSize: &totalSize,
		Page:      param.Page,
		Size:      &size,
	}, nil
}

type DtoPageResponse struct {
	TotalPage *int64 `json:"total_page,omitempty"`
	TotalSize *int64 `json:"total_size,omitempty"`
	Page      *int64 `json:"page,omitempty"`
	Size      *int64 `json:"size,omitempty"`
}
