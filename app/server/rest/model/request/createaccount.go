package request

import "github.com/jsperandio/transaction/app/domain/model"

type CreateAccount struct {
	DocumentNumber      string  `json:"document_number" validate:"required"`
	AvaliableCreditLimi float64 `json:"avaliable_credit_limit" validate:"required,gt=0" `
}

func (c *CreateAccount) ToDomainModel() *model.Account {
	return &model.Account{
		DocumentNumber:       c.DocumentNumber,
		AvaliableCreditLimit: c.AvaliableCreditLimi,
	}
}
