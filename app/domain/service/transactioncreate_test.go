package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"bou.ke/monkey"

	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/jsperandio/transaction/app/domain/repository"
	"github.com/jsperandio/transaction/app/domain/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type TransactionCreationSuite struct {
	suite.Suite
}

func TestNewTransactionCreationSuite(t *testing.T) {
	suite.Run(t, new(TransactionCreationSuite))
}

func (tcs *TransactionCreationSuite) SetupSuite() {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2023, time.February, 19, 0o1, 31, 0o0, 0, time.UTC)
	})
}

func (tcs *TransactionCreationSuite) Test_NewTransactionCreation() {
	type args struct {
		transactionRepository repository.Transaction
		accountRepository     repository.Account
	}

	type test struct {
		name string
		args args
		want *TransactionCreation
	}

	tests := []test{
		{
			name: "should return a new TransactionCreation",
			args: args{
				transactionRepository: new(mocks.Transaction),
				accountRepository:     new(mocks.Account),
			},
			want: &TransactionCreation{
				transactionRepository: new(mocks.Transaction),
				accountRepository:     new(mocks.Account),
			},
		},
		{
			name: "should return an empty TransactionCreation",
			args: args{},
			want: &TransactionCreation{},
		},
	}

	for _, tc := range tests {
		tcs.Run(
			tc.name,
			func() {
				ac := NewTransactionCreation(tc.args.transactionRepository, tc.args.accountRepository)
				tcs.Assert().Equal(tc.want, ac)
			})
	}
}

func (tcs *TransactionCreationSuite) Test_TransactionCreation_Create() {
	type fields struct {
		trep *mocks.Transaction
		arep *mocks.Account
	}

	type args struct {
		ctx         context.Context
		transaction *model.Transaction
	}

	type test struct {
		name      string
		fields    fields
		args      args
		want      *model.Transaction
		wantErr   bool
		wantedErr error
		mock      func(transaction *mocks.Transaction, account *mocks.Account)
	}

	tests := []test{
		{
			name: "when transaction is nil , should return nil transaction and ErrInvalidTransaction",
			fields: fields{
				trep: new(mocks.Transaction),
				arep: new(mocks.Account),
			},
			args: args{
				ctx:         context.Background(),
				transaction: nil,
			},
			want:      nil,
			wantErr:   true,
			wantedErr: model.ErrInvalidTransaction,
			mock:      func(transaction *mocks.Transaction, account *mocks.Account) {},
		},
		{
			name: "when transaction is invalid , should return nil transaction and ErrInvalidTransaction",
			fields: fields{
				trep: new(mocks.Transaction),
				arep: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID:       10,
					Amount:          1000,
					OperationTypeID: model.OperationType(11),
				},
			},
			want:      nil,
			wantErr:   true,
			wantedErr: model.ErrInvalidTransaction,
			mock:      func(transaction *mocks.Transaction, account *mocks.Account) {},
		},
		{
			name: "when Get account of transaction return error, should return nil transaction and an error",
			fields: fields{
				trep: new(mocks.Transaction),
				arep: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID:       3,
					Amount:          1000,
					OperationTypeID: model.OperationType(1),
				},
			},
			want:      nil,
			wantErr:   true,
			wantedErr: errors.New("error"),
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {
				account.On("Get", context.Background(), int64(3)).Return(nil, errors.New("error")).Once()
			},
		},
		{
			name: "when success Get execute but no account found, should return nil transaction and an ErrInvalidTransaction",
			fields: fields{
				trep: new(mocks.Transaction),
				arep: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID:       3,
					Amount:          1000,
					OperationTypeID: model.OperationType(1),
				},
			},
			want:      nil,
			wantErr:   true,
			wantedErr: model.ErrInvalidTransaction,
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {
				account.On("Get", context.Background(), int64(3)).Return(nil, nil).Once()
			},
		},
		{
			name: "when Save return an error, should return nil transaction and an error",
			fields: fields{
				trep: new(mocks.Transaction),
				arep: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID:       3,
					Amount:          1000,
					OperationTypeID: model.OperationType(1),
				},
			},
			want:      nil,
			wantErr:   true,
			wantedErr: errors.New("error"),
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {
				account.On("Get", context.Background(), int64(3)).Return(&model.Account{
					ID:             3,
					DocumentNumber: "333333",
				}, nil).Once()
				transaction.On("Save", context.Background(), &model.Transaction{
					AccountID:       3,
					Amount:          -1000,
					OperationTypeID: model.OperationType(1),
					EventDate:       time.Now(),
				}).Return(nil, errors.New("error")).Once()
			},
		},
		{
			name: "when success, should return transaction and no error",
			fields: fields{
				trep: new(mocks.Transaction),
				arep: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID:       3,
					Amount:          1000,
					OperationTypeID: model.OperationType(4),
				},
			},
			want: &model.Transaction{
				ID:              1,
				AccountID:       3,
				Amount:          1000,
				OperationTypeID: model.OperationType(4),
				EventDate:       time.Now(),
			},
			wantErr:   false,
			wantedErr: nil,
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {
				account.On("Get", context.Background(), int64(3)).Return(&model.Account{
					ID:             3,
					DocumentNumber: "333333",
				}, nil).Once()
				transaction.On("Save", context.Background(), &model.Transaction{
					AccountID:       3,
					Amount:          1000,
					OperationTypeID: model.OperationType(4),
					EventDate:       time.Now(),
				}).Return(&model.Transaction{
					ID:              1,
					AccountID:       3,
					Amount:          1000,
					OperationTypeID: model.OperationType(4),
					EventDate:       time.Now(),
				}, nil).Once()
			},
		},
		{
			name: "when success with operation type != Pagamento, should return transaction with negative amount and no error",
			fields: fields{
				trep: new(mocks.Transaction),
				arep: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				transaction: &model.Transaction{
					AccountID:       3,
					Amount:          5000.55,
					OperationTypeID: model.OperationType(2),
				},
			},
			want: &model.Transaction{
				ID:              1,
				AccountID:       3,
				Amount:          -5000.55,
				OperationTypeID: model.OperationType(2),
				EventDate:       time.Now(),
			},
			wantErr:   false,
			wantedErr: nil,
			mock: func(transaction *mocks.Transaction, account *mocks.Account) {
				account.On("Get", context.Background(), int64(3)).Return(&model.Account{
					ID:             3,
					DocumentNumber: "333333",
				}, nil).Once()
				transaction.On("Save", context.Background(), &model.Transaction{
					AccountID:       3,
					Amount:          -5000.55,
					OperationTypeID: model.OperationType(2),
					EventDate:       time.Now(),
				}).Return(&model.Transaction{
					ID:              1,
					AccountID:       3,
					Amount:          -5000.55,
					OperationTypeID: model.OperationType(2),
					EventDate:       time.Now(),
				}, nil).Once()
			},
		},
	}

	for _, tc := range tests {
		tcs.Run(
			tc.name,
			func() {
				tc.mock(tc.fields.trep, tc.fields.arep)
				txnc := NewTransactionCreation(tc.fields.trep, tc.fields.arep)

				got, err := txnc.Create(tc.args.ctx, tc.args.transaction)
				tcs.Assert().Equal(tc.wantErr, err != nil)
				tcs.Assert().Equal(tc.wantedErr, err)
				tcs.Assert().Equal(tc.want, got, " want = %v , got = %v ", tc.want, got)
				tc.fields.trep.AssertExpectations(tcs.T())
				tc.fields.arep.AssertExpectations(tcs.T())
			})
	}
}
