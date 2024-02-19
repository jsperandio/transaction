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
		slog.Error("operation type is invalid", "operation type", txn.OperationTypeID)
		return nil, model.ErrInvalidOperationType
	}

	_, err := t.accountRepository.Get(ctx, txn.AccountID)
	if err != nil {
		slog.Error("account not found", "account_id", txn.AccountID)
		return nil, err
	}

	if txn.OperationTypeID != model.Pagamento {
		txn.Amount = math.Abs(txn.Amount) * -1
	}
	txn.EventDate = time.Now()

	ret, err := t.transactionRepository.Save(ctx, txn)
	if err != nil {
		slog.Error("error saving transaction", "error", err)
		return nil, err
	}

	slog.Debug("transaction created", "transaction", ret)
	return ret, nil
}
