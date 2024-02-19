package service

import (
	"context"
	"log/slog"

	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/jsperandio/transaction/app/domain/repository"
)

type AccountSearcher interface {
	GetByID(ctx context.Context, id int64) (*model.Account, error)
}

type AccountSearch struct {
	repository repository.Account
}

func NewAccountSearch(rep repository.Account) *AccountSearch {
	return &AccountSearch{
		repository: rep,
	}
}

func (a *AccountSearch) GetByID(ctx context.Context, id int64) (*model.Account, error) {
	fnd, err := a.repository.Get(ctx, id)
	if err != nil {
		slog.Error("error getting account", "error", err)
		return nil, err
	}

	slog.Debug("account found", "account", fnd)
	return fnd, nil
}
