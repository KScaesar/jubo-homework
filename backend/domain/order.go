package domain

import (
	"context"
	"time"

	"github.com/KScaesar/jubo-homework/backend/util"
	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

type OrderRepo interface {
	QueryOrderList(ctx context.Context, dto *DtoQryOrderParam) (util.DtoListResponse[DtoOrderResponse], error)
	QueryOrderById(ctx context.Context, orderId string) (order DtoOrderResponse, err error)
	LockOrderById(ctx context.Context, orderId string) (order Order, err error)

	CreateOrder(ctx context.Context, order *Order) error
	UpdateOrder(ctx context.Context, order *Order) error
}

func NewOrder(dto *DtoCreateOrder) (Order, error) {
	id := util.NewUlid()
	now := time.Now()

	order := Order{
		Id:        id,
		Message:   dto.Message,
		CreatedAt: now,
		UpdatedAt: now,
		PatientId: dto.PatientId,
	}

	return order, order.Validate()
}

type Order struct {
	Id        string    `gorm:"column:id"`
	Message   string    `gorm:"column:message"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	PatientId string    `gorm:"column:patient_id"`
}

func (o *Order) Validate() error {
	if o.Message == "" {
		return errors.WrapWithMessage(errors.ErrInvalidParams, "message not allow empty")
	}
	return nil
}

func (o *Order) UpdateInfo(dto *DtoUpdateOrderInfo) error {
	if dto.Message != nil {
		o.Message = *dto.Message
	}

	o.UpdatedAt = time.Now()
	return o.Validate()
}
