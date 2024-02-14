package repository

import (
	"context"

	"github.com/jsperandio/transaction/app/domain/model"
)

type Account interface {
	Save(ctx context.Context, acc *model.Account) (*model.Account, error)
	Find(ctx context.Context, ID string) (*model.Account, error)
	FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Account, error)
}
