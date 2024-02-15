package model

import "github.com/jsperandio/transaction/app/domain/model"

type Account struct {
	ID             int64  `db:"id"`
	DocumentNumber string `db:"document_number"`
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

func (a *Account) ToDomainModel() *model.Account {
	if a == nil {
		return nil
	}

	return &model.Account{
		ID:             a.ID,
		DocumentNumber: a.DocumentNumber,
	}
}
