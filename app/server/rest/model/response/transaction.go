package response

import (
	"time"

	"github.com/jsperandio/transaction/app/domain/model"
)

type Transaction struct {
	ID              int       `json:"transaction_id"`
	AccountID       int       `json:"account_id"`
	OperationTypeID int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

func NewTransactionFromDomain(transaction *model.Transaction) *Transaction {
	return &Transaction{
		ID:              transaction.ID,
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID.Index(),
		Amount:          transaction.Amount,
		EventDate:       transaction.EventDate,
	}
}
