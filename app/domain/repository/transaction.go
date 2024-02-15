package repository

import (
	"context"

	"github.com/jsperandio/transaction/app/domain/model"
)

type Transaction interface {
	Save(ctx context.Context, t *model.Transaction) (*model.Transaction, error)
	Find(ctx context.Context, ID int64) (*model.Transaction, error)
}
