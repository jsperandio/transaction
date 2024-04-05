package response

import "github.com/jsperandio/transaction/app/domain/model"

type Account struct {
	ID                  int64   `json:"account_id"`
	DocumentNumber      string  `json:"document_number"`
	AvaliableCreditLimi float64 `json:"avaliable_credit_limit"`
}

func NewAccountFromDomain(acc *model.Account) *Account {
	if acc == nil {
		return nil
	}
	return &Account{
		ID:                  acc.ID,
		DocumentNumber:      acc.DocumentNumber,
		AvaliableCreditLimi: acc.AvaliableCreditLimit,
	}
}
