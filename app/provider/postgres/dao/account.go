package dao

import (
	"context"

	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/jsperandio/transaction/app/provider/postgres/client"
	dbmodel "github.com/jsperandio/transaction/app/provider/postgres/model"
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
	dba := dbmodel.NewAccountFromDomain(acc)
	err := a.conn.DB.Get(dba, "INSERT INTO account (document_number) VALUES ($1) RETURNING id", dba.DocumentNumber)
	if err != nil {
		return nil, err
	}
	return dba.ToDomainModel(), nil
}

func (a *Account) Get(ctx context.Context, ID int64) (*model.Account, error) {
	acc := dbmodel.Account{}
	err := a.conn.DB.Get(&acc, "SELECT id, document_number FROM account WHERE id = $1", ID)
	if err != nil {
		return nil, err
	}

	return acc.ToDomainModel(), nil
}

func (a *Account) FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Account, error) {
	acc := dbmodel.Account{}
	err := a.conn.DB.Get(&acc, "SELECT id, document_number FROM account WHERE document_number = $1", documentNumber)
	if err != nil {
		return nil, err
	}

	return acc.ToDomainModel(), nil
}
