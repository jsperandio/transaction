package dao

import (
	"context"

	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/jsperandio/transaction/app/provider/postgres/client"
)

type Transaction struct {
	conn *client.Connection
}

func NewTransaction(c *client.Connection) *Transaction {
	return &Transaction{
		conn: c,
	}
}

func (tx *Transaction) Save(ctx context.Context, t *model.Transaction) (*model.Transaction, error) {
	return nil, nil
}

func (tx *Transaction) Find(ctx context.Context, ID string) (*model.Transaction, error) {
	return nil, nil
}
