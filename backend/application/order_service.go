package application

import (
	"context"

	"github.com/KScaesar/jubo-homework/backend/domain"
	"github.com/KScaesar/jubo-homework/backend/util"
)

type OrderService interface {
	QueryOrderList(ctx context.Context, dto *domain.QryOrderParam) (util.ListResponse[domain.OrderResponse], error)
	QueryOrderById(ctx context.Context, orderId string) (domain.OrderResponse, error)
	CreateOrder(ctx context.Context, dto *domain.CreateOrderCmd) (id string, err error)
	UpdateOrderInfo(ctx context.Context, orderId string, dto *domain.UpdateOrderInfoCmd) error
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

func (uc *OrderUseCase) QueryOrderList(
	ctx context.Context, dto *domain.QryOrderParam,
) (
	util.ListResponse[domain.OrderResponse], error,
) {
	dto.PreProcess(false)
	return uc.orderRepo.QueryOrderList(ctx, dto)
}

func (uc *OrderUseCase) QueryOrderById(ctx context.Context, orderId string) (domain.OrderResponse, error) {
	return uc.orderRepo.QueryOrderById(ctx, orderId)
}

func (uc *OrderUseCase) CreateOrder(ctx context.Context, dto *domain.CreateOrderCmd) (string, error) {
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

func (uc *OrderUseCase) UpdateOrderInfo(ctx context.Context, orderId string, dto *domain.UpdateOrderInfoCmd) error {
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
