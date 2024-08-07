package response

import (
	"time"

	"github.com/jsperandio/transaction/app/domain/model"
)

type Transaction struct {
	ID              int64     `json:"transaction_id"`
	AccountID       int64     `json:"account_id"`
	OperationTypeID int64     `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

func NewTransactionFromDomain(txn *model.Transaction) *Transaction {
	if txn == nil {
		return nil
	}

	return &Transaction{
		ID:              txn.ID,
		AccountID:       txn.AccountID,
		OperationTypeID: txn.OperationTypeID.Index(),
		Amount:          txn.Amount,
		EventDate:       txn.EventDate,
	}
}
