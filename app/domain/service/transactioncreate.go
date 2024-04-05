package service

import (
	"context"
	"log/slog"
	"math"
	"time"

	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/jsperandio/transaction/app/domain/repository"
)

type TransactionCreator interface {
	Create(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
}

type TransactionCreation struct {
	transactionRepository repository.Transaction
	accountRepository     repository.Account
}

func NewTransactionCreation(trp repository.Transaction, arp repository.Account) *TransactionCreation {
	return &TransactionCreation{
		transactionRepository: trp,
		accountRepository:     arp,
	}
}

func (t *TransactionCreation) Create(ctx context.Context, txn *model.Transaction) (*model.Transaction, error) {
	if txn == nil {
		slog.Error("transaction is nil")
		return nil, model.ErrInvalidTransaction
	}

	if !txn.OperationTypeID.IsValid() {
		slog.Error(model.ErrInvalidOperationType.Error(), "operation type", txn.OperationTypeID)
		return nil, model.ErrInvalidTransaction
	}

	acc, err := t.accountRepository.Get(ctx, txn.AccountID)
	if err != nil {
		slog.Error("account get error", "account_id", txn.AccountID)
		return nil, err
	}

	if acc == nil {
		slog.Error(model.ErrAccountNotFound.Error(), "account_id", txn.AccountID)
		return nil, model.ErrInvalidTransaction
	}

	if txn.OperationTypeID != model.Pagamento {
		txn.Amount = math.Abs(txn.Amount) * -1
	}
	txn.EventDate = time.Now()

	acc.AvaliableCreditLimit += txn.Amount
	if acc.AvaliableCreditLimit < 0 {
		slog.Error("transaction value except account avaliable credit")
		return nil, model.ErrNonAvaliableLimitForAccountTransaction
	}

	ret, err := t.transactionRepository.Save(ctx, txn)
	if err != nil {
		slog.Error("error saving transaction", "error", err)
		return nil, err
	}

	err = t.accountRepository.Update(ctx, acc)
	if err != nil {
		slog.Error("erro on save account new avaliable limit")
		return nil, err
	}

	slog.Debug("transaction created", "transaction", ret)
	return ret, nil
}
