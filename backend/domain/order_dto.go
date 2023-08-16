package domain

import (
	"github.com/KScaesar/jubo-homework/backend/util"
)

func TransformOrderModel(order *Order) OrderResponse {
	return OrderResponse{
		Id:        order.Id,
		Message:   order.Message,
		PatientId: order.PatientId,
	}
}

// read dto

type OrderResponse struct {
	Id        string `json:"id" gorm:"column:id"`
	Message   string `json:"message" gorm:"column:message"`
	PatientId string `json:"patient_id" gorm:"column:patient_id"`
}

type QryOrderParam struct {
	FilterOrderParam
	util.SortParam
	util.PageParam
}

func (d *QryOrderParam) PreProcess(isPagination bool) {
	d.SortParam.SetDefaultIfInvalid("updated_at", util.SortDesc)
	if isPagination {
		d.PageParam.SetDefaultIfInvalid()
		return
	}
	d.PageParam.SetWithoutPagination()
}

type FilterOrderParam struct {
	PatientId string `form:"patient_id" rdb:"patient_id = ?"`
}

// write dto

type CreateOrderCmd struct {
	Message   string `json:"message"`
	PatientId string `json:"patient_id"`
}

type UpdateOrderInfoCmd struct {
	Message *string `json:"message"`
}
