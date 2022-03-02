package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type contextKey struct{}

var txKey = contextKey{}

type Tx struct {
	*gorm.DB
}

func NewTx(db *gorm.DB) *Tx {
	return &Tx{DB: db}
}

func setTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey, tx)
}

func readTx(ctx context.Context) (*gorm.DB, error) {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if !ok {
		return nil, errors.New("tx is not exist")
	}
	return tx, nil
}

func (t *Tx) DoInTx(ctx context.Context, f func(ctx context.Context) error) error {
	tx := t.DB.Begin()
	ctx = setTx(ctx, tx)
	if err := f(ctx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
