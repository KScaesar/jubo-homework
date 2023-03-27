package infra

import (
	"context"

	"gorm.io/gorm/clause"

	"github.com/KScaesar/jubo-homework/backend/domain"
	"github.com/KScaesar/jubo-homework/backend/util"
	"github.com/KScaesar/jubo-homework/backend/util/database"
	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

const (
	OrderTableName = "orders"
)

func NewOrderRepository(db *database.WrapperGorm) *OrderRepository {
	return &OrderRepository{db: db}
}

type OrderRepository struct {
	db *database.WrapperGorm
}

func (repo *OrderRepository) QueryOrderList(ctx context.Context, dto *domain.DtoQryOrderParam) (util.DtoListResponse[domain.DtoOrderResponse], error) {
	db := repo.db.ChooseProcessor(ctx)

	if dto == nil {
		dto = &domain.DtoQryOrderParam{}
	}
	return database.GormQueryListFromSingleTable[domain.DtoOrderResponse](db, OrderTableName, dto)
}

func (repo *OrderRepository) QueryOrderById(ctx context.Context, id string) (order domain.DtoOrderResponse, err error) {
	db := repo.db.ChooseProcessor(ctx)

	err = db.Table(OrderTableName).Where("id = ?", id).First(&order).Error
	if err != nil {
		err = errors.WrapWithMessage(database.GormError(err), "order_id = %v", id)
		return
	}
	return
}

func (repo *OrderRepository) LockOrderById(ctx context.Context, id string) (order domain.Order, err error) {
	db := repo.db.ChooseProcessor(ctx)

	err = db.Table(OrderTableName).Where("id = ?", id).Clauses(clause.Locking{Strength: "UPDATE"}).First(&order).Error
	if err != nil {
		err = errors.WrapWithMessage(database.GormError(err), "id = %v", id)
		return
	}
	return
}

func (repo *OrderRepository) CreateOrder(ctx context.Context, order *domain.Order) error {
	db := repo.db.ChooseProcessor(ctx)

	err := db.Table(OrderTableName).Create(order).Error
	if err != nil {
		return database.GormError(err)
	}
	return nil
}

func (repo *OrderRepository) UpdateOrder(ctx context.Context, order *domain.Order) error {
	db := repo.db.ChooseProcessor(ctx)

	err := db.Table(OrderTableName).Where("id = ?", order.Id).Save(order).Error
	if err != nil {
		return database.GormError(err)
	}
	return nil
}
