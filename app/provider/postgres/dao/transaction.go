package dao

import (
	"context"

	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/jsperandio/transaction/app/provider/postgres/client"
	dbmodel "github.com/jsperandio/transaction/app/provider/postgres/model"
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
	txn := dbmodel.NewTransactionFromDomain(t)
	res, err := tx.conn.DB.NamedQuery(`	INSERT INTO transaction  (account_id, operation_type_id, amount, event_date) 
														  VALUES (:account_id, :operation_type_id, :amount, :event_date) 
										RETURNING id`, txn)
	if err != nil {
		return nil, err
	}

	for res.Next() {
		err = res.Scan(&txn.ID)
		if err != nil {
			return nil, err
		}
	}

	return txn.ToDomain(), nil
}

func (tx *Transaction) Find(ctx context.Context, ID int64) (*model.Transaction, error) {
	txn := dbmodel.Transaction{}
	err := tx.conn.DB.Get(&txn, `SELECT id, account_id, operation_type_id, amount, event_date  FROM transaction WHERE id = $1`, ID)
	if err != nil {
		return nil, err
	}

	return txn.ToDomain(), nil
}
