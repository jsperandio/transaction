package response

import "github.com/jsperandio/transaction/app/domain/model"

type Account struct {
	ID             int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func NewAccountFromDomain(acc model.Account) *Account {
	return &Account{
		ID:             acc.ID,
		DocumentNumber: acc.DocumentNumber,
	}
}
