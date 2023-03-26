package database

import (
	"context"
)

// TransactionFactory that lifecycle is equal to process scope, and is goroutine safe.
type TransactionFactory interface {
	CreateTransaction(ctx context.Context) Transaction
}

// Transaction lifecycle is equal to request scope, and not goroutine safe.
// You should use only one Transaction per goroutine
//
// 參數 fn 中的 txCtx,
// 保證 txCtx 一定有 tx 元件,
// 比如 *gorm.DB, mongo.Session.
type Transaction interface {
	AutoCommit(fn func(ctx context.Context) error) error
}
