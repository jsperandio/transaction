package dao

import (
	"context"

	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/jsperandio/transaction/app/provider/postgres/client"
)

type Account struct {
	conn *client.Connection
}

func NewAccount(c *client.Connection) *Account {
	return &Account{
		conn: c,
	}
}

func (a *Account) Save(ctx context.Context, acc *model.Account) (*model.Account, error) {
	return nil, nil
}

func (a *Account) Get(ctx context.Context, ID int) (*model.Account, error) {
	return nil, nil
}

func (a *Account) FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Account, error) {
	return nil, nil
}
