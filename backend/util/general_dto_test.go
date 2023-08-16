package util

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPageResponse_json_Marshal(t *testing.T) {
	dto := PageParam{Page: 4, Size: 9}
	expected := `{"total_page":4,"total_size":34,"page":4,"size":7}`

	response := NewPageResponse(dto, 34)

	body, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.Equal(t, expected, string(body))
}

func TestNewPageResponse(t *testing.T) {
	type args struct {
		dto       PageParam
		totalSize uint64
	}
	tests := []struct {
		name   string
		args   args
		expect PageResponse
	}{
		{
			name: "1",
			args: args{
				dto:       PageParam{Page: 1, Size: 10},
				totalSize: 34,
			},
			expect: PageResponse{TotalPage: 4, TotalSize: 34, Page: 1, Size: 10},
		},
		{
			name: "2",
			args: args{
				dto:       PageParam{Page: 2, Size: 40},
				totalSize: 34,
			},
			expect: PageResponse{TotalPage: 1, TotalSize: 34, Page: 2, Size: 0},
		},
		{
			name: "3",
			args: args{
				dto:       PageParam{Page: 2, Size: 30},
				totalSize: 34,
			},
			expect: PageResponse{TotalPage: 2, TotalSize: 34, Page: 2, Size: 4},
		},
		{
			name: "4",
			args: args{
				dto:       PageParam{Page: 1, Size: 50},
				totalSize: 34,
			},
			expect: PageResponse{TotalPage: 1, TotalSize: 34, Page: 1, Size: 34},
		},
		{
			name: "5",
			args: args{
				dto:       PageParam{Page: 1, Size: 40},
				totalSize: 40,
			},
			expect: PageResponse{TotalPage: 1, TotalSize: 40, Page: 1, Size: 40},
		},
		{
			name: "6",
			args: args{
				dto:       PageParam{Page: 1, Size: 40},
				totalSize: 0,
			},
			expect: PageResponse{TotalPage: 1, TotalSize: 0, Page: 1, Size: 0},
		},
		{
			name: "8",
			args: args{
				dto:       PageParam{Page: 5, Size: 10},
				totalSize: 24,
			},
			expect: PageResponse{TotalPage: 3, TotalSize: 24, Page: 5, Size: 0},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				actual := NewPageResponse(tt.args.dto, tt.args.totalSize)
				assert.Equal(t, tt.expect, actual)
			},
		)
	}
}
