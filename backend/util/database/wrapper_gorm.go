package database

import (
	"context"

	"gorm.io/gorm"
)

func NewWrapperGorm(db *gorm.DB) *WrapperGorm {
	return &WrapperGorm{db: db}
}

type WrapperGorm struct {
	db *gorm.DB
}

func (wrapper *WrapperGorm) ContextWithTx(ctx context.Context, tx *gorm.DB) (txCtx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	// key 用來確認是否來自同一個 *gorm.DB
	key := wrapper
	return context.WithValue(ctx, key, tx)
}

// ChooseProcessor
// 如果找不到符合的 tx 元件, 則使用原本的 db 元件,
// 找不到的原因:
//
// 1. 沒有 transaction 需求, 所以沒有傳入 tx 元件
// 2. tx 元件來自不同的 database, 那當然無法達成 transaction 的需求, 只能各自處理 sql operation
//
// 回傳值 processor, 可以代表 tx or db, 兩者型別都是  *gorm.DB
// 其差異請參考下列網址
// https://gorm.io/docs/method_chaining.html#New-Session-Mode
func (wrapper *WrapperGorm) ChooseProcessor(txCtx context.Context) (processor *gorm.DB) {
	if txCtx == nil {
		return wrapper.Unwrap()
	}

	if wrapper.ExistTransaction(txCtx) {
		tx := txCtx.Value(wrapper).(*gorm.DB)
		return tx
	}
	return wrapper.Unwrap()
}

func (wrapper *WrapperGorm) ExistTransaction(ctx context.Context) bool {
	_, ok := ctx.Value(wrapper).(*gorm.DB)
	return ok
}

func (wrapper *WrapperGorm) Unwrap() *gorm.DB {
	return wrapper.db
}
