package model

import (
	"time"

	"github.com/jsperandio/transaction/app/domain/model"
)

type Transaction struct {
	ID              int64     `db:"id"`
	AccountID       int64     `db:"account_id"`
	OperationTypeID int64     `db:"operation_type_id"`
	Amount          float64   `db:"amount"`
	EventDate       time.Time `db:"event_date"`
}

func NewTransactionFromDomain(t *model.Transaction) *Transaction {
	if t == nil {
		return nil
	}

	return &Transaction{
		AccountID:       t.AccountID,
		OperationTypeID: t.OperationTypeID.Index(),
		Amount:          t.Amount,
		EventDate:       t.EventDate,
	}
}

func (t *Transaction) ToDomain() *model.Transaction {
	if t == nil {
		return nil
	}

	return &model.Transaction{
		ID:              t.ID,
		AccountID:       t.AccountID,
		OperationTypeID: model.OperationType(t.OperationTypeID),
		Amount:          t.Amount,
		EventDate:       t.EventDate,
	}
}
