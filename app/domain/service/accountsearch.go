package service

import (
	"context"

	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/jsperandio/transaction/app/domain/repository"
)

type AccountSearcher interface {
	GetByID(ctx context.Context, id int) (*model.Account, error)
}

type AccountSearch struct {
	repository repository.Account
}

func NewAccountSearch(rep repository.Account) *AccountSearch {
	return &AccountSearch{
		repository: rep,
	}
}

func (a *AccountSearch) GetByID(ctx context.Context, id int) (*model.Account, error) {
	fnd, err := a.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return fnd, nil
}
