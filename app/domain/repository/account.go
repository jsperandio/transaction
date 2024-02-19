package repository

import (
	"context"

	"github.com/jsperandio/transaction/app/domain/model"
)

type Account interface {
	Save(ctx context.Context, acc *model.Account) (*model.Account, error)
	Get(ctx context.Context, ID int64) (*model.Account, error)
	FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Account, error)
}
