package transaction

import (
	"context"

	"gorm.io/gorm"
)

type (
	trxKey string

	bizFunc func(txCtx context.Context) error
)

const gormTrxKey trxKey = "gorm_trx"

type TransactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

func (tm *TransactionManager) Execute(ctx context.Context, bizFn bizFunc) error {
	return tm.db.Transaction(func(tx *gorm.DB) error {
		txCtx := tm.setGormTransaction(ctx, tx)
		return bizFn(txCtx)
	})
}

func (tm *TransactionManager) GetGormTransaction(ctx context.Context) *gorm.DB {
	tx := ctx.Value(gormTrxKey)

	gormTx, ok := tx.(*gorm.DB)
	if !ok {
		return tm.db
	}

	return gormTx
}

func (tm *TransactionManager) setGormTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, gormTrxKey, tx)
}
