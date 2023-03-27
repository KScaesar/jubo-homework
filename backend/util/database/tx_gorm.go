package database

import (
	"context"

	"gorm.io/gorm"

	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

func NewGormTxFactory(db *WrapperGorm) *GormTxFactory {
	return &GormTxFactory{db: db}
}

type GormTxFactory struct {
	db *WrapperGorm
}

func (f *GormTxFactory) CreateTransaction(ctx context.Context) Transaction {
	if ctx == nil {
		ctx = context.Background()
	}

	return &gormTxAdapter{db: f.db, ctx: ctx}
}

type gormTxAdapter struct {
	ctx context.Context
	db  *WrapperGorm
}

func (adapter *gormTxAdapter) AutoCommit(fn func(ctx context.Context) error) error {
	ctx := adapter.ctx

	if adapter.db.ExistTransaction(ctx) {
		return fn(ctx)
	}

	gormTxFn := func(tx *gorm.DB) error {
		txCtx := adapter.db.ContextWithTx(ctx, tx)
		return fn(txCtx)
	}

	err := adapter.db.Unwrap().Transaction(gormTxFn)
	if err != nil {
		_, ok := errors.ExtractCustomError(err)
		if !ok {
			return errors.Join3rdParty(errors.ErrSystem, err)
		}
		return err
	}

	return nil
}
