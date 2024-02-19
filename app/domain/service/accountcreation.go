package service

import (
	"context"
	"log/slog"

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
		slog.Error("error on FindByDocumentNumber", "error", err, "documentNumber", account.DocumentNumber)
		return nil, err
	}

	if fnd != nil {
		slog.Info("account already exists", "documentNumber", account.DocumentNumber)
		return nil, model.ErrAccountAlreadyExists
	}

	acc, err := a.repository.Save(ctx, account)
	if err != nil {
		slog.Error("error on Save", "error", err, "account", account)
		return nil, err
	}

	slog.Debug("account created", "account", acc)
	return acc, nil
}
