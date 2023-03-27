package application

import (
	"context"

	"github.com/KScaesar/jubo-homework/backend/domain"
	"github.com/KScaesar/jubo-homework/backend/util"
)

type OrderService interface {
	QueryOrderList(ctx context.Context, dto *domain.DtoQryOrderParam) (util.DtoListResponse[domain.DtoOrderResponse], error)
	QueryOrderById(ctx context.Context, orderId string) (domain.DtoOrderResponse, error)
	CreateOrder(ctx context.Context, dto *domain.DtoCreateOrder) (id string, err error)
	UpdateOrderInfo(ctx context.Context, orderId string, dto *domain.DtoUpdateOrderInfo) error
}

func NewOrderUseCase(orderRepo domain.OrderRepo, patientRepo domain.PatientRepo) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:   orderRepo,
		patientRepo: patientRepo,
	}
}

type OrderUseCase struct {
	orderRepo   domain.OrderRepo
	patientRepo domain.PatientRepo
}

func (uc *OrderUseCase) QueryOrderList(ctx context.Context, dto *domain.DtoQryOrderParam) (util.DtoListResponse[domain.DtoOrderResponse], error) {
	dto.DtoSortOrderParam.SetDefault()
	return uc.orderRepo.QueryOrderList(ctx, dto)
}

func (uc *OrderUseCase) QueryOrderById(ctx context.Context, orderId string) (domain.DtoOrderResponse, error) {
	return uc.orderRepo.QueryOrderById(ctx, orderId)
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, dto *domain.DtoCreateOrder) (string, error) {
	_, err := uc.patientRepo.QueryPatientById(ctx, dto.PatientId)
	if err != nil {
		return "", err
	}

	order, err := domain.NewOrder(dto)
	if err != nil {
		return "", err
	}

	err = uc.orderRepo.CreateOrder(ctx, &order)
	if err != nil {
		return "", err
	}

	return order.Id, nil
}

func (uc *OrderUseCase) UpdateOrderInfo(ctx context.Context, orderId string, dto *domain.DtoUpdateOrderInfo) error {
	order, err := uc.orderRepo.LockOrderById(ctx, orderId)
	if err != nil {
		return err
	}

	err = order.UpdateInfo(dto)
	if err != nil {
		return err
	}

	err = uc.orderRepo.UpdateOrder(ctx, &order)
	if err != nil {
		return err
	}

	return nil
}
