package request

import "github.com/jsperandio/transaction/app/domain/model"

type CreateTransaction struct {
	AccountID       int     `json:"account_id" validate:"required,gt=0"`
	OperationTypeID int     `json:"operation_type_id" validate:"required,gt=0"`
	Amount          float64 `json:"amount" validate:"required,gt=0"`
}

func (ct *CreateTransaction) ToDomainModel() *model.Transaction {
	return &model.Transaction{
		AccountID:       ct.AccountID,
		OperationTypeID: model.OperationType(ct.OperationTypeID),
		Amount:          ct.Amount,
	}
}
