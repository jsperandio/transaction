package response

import "github.com/jsperandio/transaction/app/domain/model"

type Account struct {
	ID             int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func NewAccountFromDomain(acc *model.Account) *Account {
	if acc == nil {
		return nil
	}
	return &Account{
		ID:             acc.ID,
		DocumentNumber: acc.DocumentNumber,
	}
}
