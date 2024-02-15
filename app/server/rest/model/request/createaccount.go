package request

import "github.com/jsperandio/transaction/app/domain/model"

type CreateAccount struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

func (c *CreateAccount) ToDomainModel() *model.Account {
	return &model.Account{
		DocumentNumber: c.DocumentNumber,
	}
}
