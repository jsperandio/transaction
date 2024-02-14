package service

import (
	"context"

	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/jsperandio/transaction/app/domain/repository"
)

type AccountCreator interface {
	Create(ctx context.Context, account *model.Account) (*model.Account, error)
}

type AccountCreation struct {
	repository repository.Account
}

func NewAccountCreation(repository repository.Account) *AccountCreation {
	return &AccountCreation{
		repository: repository,
	}
}

func (a *AccountCreation) Create(ctx context.Context, account *model.Account) (*model.Account, error) {
	fnd, err := a.repository.FindByDocumentNumber(ctx, account.DocumentNumber)
	if err != nil {
		return nil, err
	}

	if fnd != nil {
		return nil, model.ErrAccountAlreadyExists
	}

	acc, err := a.repository.Save(ctx, account)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
