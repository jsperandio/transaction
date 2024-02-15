package service

import (
	"context"
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
		return nil, model.ErrInvalidTransaction
	}

	if !txn.OperationTypeID.IsValid() {
		return nil, model.ErrInvalidOperationType
	}

	_, err := t.accountRepository.Get(ctx, txn.AccountID)
	if err != nil {
		return nil, err
	}

	if txn.OperationTypeID == model.Pagamento {
		txn.Amount = math.Abs(txn.Amount) * -1
	}
	txn.EventDate = time.Now()

	ret, err := t.transactionRepository.Save(ctx, txn)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
