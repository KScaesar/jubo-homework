package util

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDtoPageResponse_json_Marshal(t *testing.T) {
	param := inputDtoPageParam(4, 9)
	expected := `{"total_page":4,"total_size":34,"page":4,"size":7}`

	response, err := NewDtoPageResponse(param, 34)
	assert.NoError(t, err)

	body, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(body))
}

func TestNewDtoPageResponse(t *testing.T) {
	type args struct {
		param     DtoPageParam
		totalSize int64
	}
	tests := []struct {
		name   string
		args   args
		expect *DtoPageResponse
	}{
		{
			name: "1",
			args: args{
				param:     inputDtoPageParam(1, 10),
				totalSize: 34,
			},
			expect: expectedDtoPageResponse(4, 34, 1, 10),
		},
		{
			name: "2",
			args: args{
				param:     inputDtoPageParam(2, 40),
				totalSize: 34,
			},
			expect: expectedDtoPageResponse(1, 34, 2, 0),
		},
		{
			name: "3",
			args: args{
				param:     inputDtoPageParam(2, 30),
				totalSize: 34,
			},
			expect: expectedDtoPageResponse(2, 34, 2, 4),
		},
		{
			name: "4",
			args: args{
				param:     inputDtoPageParam(1, 50),
				totalSize: 34,
			},
			expect: expectedDtoPageResponse(1, 34, 1, 34),
		},
		{
			name: "5",
			args: args{
				param:     inputDtoPageParam(1, 40),
				totalSize: 40,
			},
			expect: expectedDtoPageResponse(1, 40, 1, 40),
		},
		{
			name: "6",
			args: args{
				param:     inputDtoPageParam(1, 40),
				totalSize: 0,
			},
			expect: expectedDtoPageResponse(1, 0, 1, 0),
		},
		{
			name: "assign value, but value is invalid",
			args: args{
				param:     inputDtoPageParam(0, 0),
				totalSize: 0,
			},
			expect: nil,
		},
		{
			name: "not assign value",
			args: args{
				param:     DtoPageParam{},
				totalSize: 0,
			},
			expect: nil,
		},
		{
			name: "8",
			args: args{
				param:     inputDtoPageParam(5, 10),
				totalSize: 24,
			},
			expect: expectedDtoPageResponse(3, 24, 5, 0),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := NewDtoPageResponse(tt.args.param, tt.args.totalSize)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func expectedDtoPageResponse(totalPage int64, totalSize int64, page int64, size int64) *DtoPageResponse {
	return &DtoPageResponse{TotalPage: &totalPage, TotalSize: &totalSize, Page: &page, Size: &size}
}

func inputDtoPageParam(page int64, size int64) DtoPageParam {
	return DtoPageParam{Page: &page, Size: &size}
}

func TestValidateSort(t *testing.T) {
	type args struct {
		sort any
	}
	tests := []struct {
		name   string
		args   args
		assert assert.ErrorAssertionFunc
	}{
		{
			name: "欄位型別不是 SortKind 會跳過驗證",
			args: args{
				sort: struct {
					SortName  SortKind `rdb:"name" json:"sort_name"`
					SortPhone string   `rdb:"phone" json:"sort_phone"`
				}{
					SortName:  "asc",
					SortPhone: "akfj",
				},
			},
			assert: assert.NoError,
		},
		{
			name: "欄位型別是 SortKind, 可驗證不合理字串",
			args: args{
				sort: struct {
					SortName  SortKind `rdb:"name" json:"sort_name"`
					SortPhone string   `rdb:"phone" json:"sort_phone"`
				}{
					SortName:  "asc1",
					SortPhone: "akfj",
				},
			},
			assert: assert.Error,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.assert(t, ValidateSort(tt.args.sort))
		})
	}
}
