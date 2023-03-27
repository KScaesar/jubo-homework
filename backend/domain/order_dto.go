package domain

import (
	"github.com/KScaesar/jubo-homework/backend/util"
)

func TransformOrderModel(order *Order) DtoOrderResponse {
	return DtoOrderResponse{
		Id:        order.Id,
		Message:   order.Message,
		PatientId: order.PatientId,
	}
}

// read dto

type DtoOrderResponse struct {
	Id        string `json:"id" gorm:"column:id"`
	Message   string `json:"message" gorm:"column:message"`
	PatientId string `json:"patient_id" gorm:"column:patient_id"`
}

type DtoQryOrderParam struct {
	DtoFilterOrderParam
	DtoSortOrderParam
	util.DtoPageParam
}

type DtoFilterOrderParam struct {
	PatientId string `form:"patient_id" rdb:"patient_id = ?"`
}

type DtoSortOrderParam struct {
	SortUpdatedAt util.SortKind `form:"sort_updated_at" rdb:"updated_at" validate:"sort"`
}

func (dto *DtoSortOrderParam) SetDefault() {
	if dto.SortUpdatedAt == "" {
		dto.SortUpdatedAt = util.SortDesc
	}
}

func (d *DtoQryOrderParam) FilterParam() any {
	return &d.DtoFilterOrderParam
}

func (d *DtoQryOrderParam) SortParam() any {
	return &d.DtoSortOrderParam
}

func (d *DtoQryOrderParam) PageParam() util.DtoPageParam {
	return d.DtoPageParam
}

// write dto

type DtoCreateOrder struct {
	Message   string `json:"message"`
	PatientId string `json:"patient_id"`
}

type DtoUpdateOrderInfo struct {
	Message *string `json:"message"`
}
