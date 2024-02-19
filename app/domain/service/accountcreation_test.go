package service

import (
	"context"
	"errors"
	"testing"

	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/jsperandio/transaction/app/domain/repository"
	"github.com/jsperandio/transaction/app/domain/repository/mocks"
	"github.com/stretchr/testify/suite"
)

type AccountCreationSuite struct {
	suite.Suite
}

func TestNewAccountCreationSuite(t *testing.T) {
	suite.Run(t, new(AccountCreationSuite))
}

func (acs *AccountCreationSuite) Test_NewAccountCreation() {
	type args struct {
		repository repository.Account
	}

	type test struct {
		name string
		args args
		want *AccountCreation
	}

	tests := []test{
		{
			name: "should return a new AccountCreation",
			args: args{
				repository: new(mocks.Account),
			},
			want: &AccountCreation{
				repository: new(mocks.Account),
			},
		},
		{
			name: "should return an empty AccountCreation",
			args: args{},
			want: &AccountCreation{},
		},
	}

	for _, tc := range tests {
		acs.Run(
			tc.name,
			func() {
				ac := NewAccountCreation(tc.args.repository)
				acs.Assert().Equal(tc.want, ac)
			})
	}
}

func (acs *AccountCreationSuite) Test_AccountCreation_Create() {
	type fields struct {
		account *mocks.Account
	}

	type args struct {
		ctx     context.Context
		account *model.Account
	}

	type test struct {
		name      string
		fields    fields
		args      args
		want      *model.Account
		wantErr   bool
		wantedErr error
		mock      func(account *mocks.Account)
	}

	tests := []test{
		{
			name: "when FindByDocumentNumber return an error, should return nil account and an error",
			fields: fields{
				account: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				account: &model.Account{
					DocumentNumber: "11111",
				},
			},
			want:      nil,
			wantErr:   true,
			wantedErr: errors.New("error"),
			mock: func(account *mocks.Account) {
				account.On("FindByDocumentNumber", context.Background(), "11111").Return(nil, errors.New("error")).Once()
			},
		},
		{
			name: "when FindByDocumentNumber find an account, should return nil account and an error",
			fields: fields{
				account: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				account: &model.Account{
					DocumentNumber: "123456",
				},
			},
			want:      nil,
			wantErr:   true,
			wantedErr: model.ErrAccountAlreadyExists,
			mock: func(account *mocks.Account) {
				account.On("FindByDocumentNumber", context.Background(), "123456").
					Return(&model.Account{
						ID:             1,
						DocumentNumber: "123456",
					}, nil).Once()
			},
		},
		{
			name: "when Save return an error, should return nil account and an error",
			fields: fields{
				account: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				account: &model.Account{
					DocumentNumber: "123457",
				},
			},
			want:      nil,
			wantErr:   true,
			wantedErr: errors.New("error"),
			mock: func(account *mocks.Account) {
				account.On("FindByDocumentNumber", context.Background(), "123457").Return(nil, nil).Once()
				account.On("Save", context.Background(), &model.Account{
					DocumentNumber: "123457",
				}).Return(nil, errors.New("error")).Once()
			},
		},
		{
			name: "when success, should return the account and no error",
			fields: fields{
				account: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				account: &model.Account{
					DocumentNumber: "123458",
				},
			},
			want: &model.Account{
				ID:             1,
				DocumentNumber: "123458",
			},
			wantErr:   false,
			wantedErr: nil,
			mock: func(account *mocks.Account) {
				account.On("FindByDocumentNumber", context.Background(), "123458").Return(nil, nil).Once()
				account.On("Save", context.Background(), &model.Account{
					DocumentNumber: "123458",
				}).Return(&model.Account{
					ID:             1,
					DocumentNumber: "123458",
				}, nil).Once()
			},
		},
	}

	for _, tc := range tests {
		acs.Run(
			tc.name,
			func() {
				tc.mock(tc.fields.account)
				ac := NewAccountCreation(tc.fields.account)

				got, err := ac.Create(tc.args.ctx, tc.args.account)
				acs.Assert().Equal(tc.wantErr, err != nil)
				acs.Assert().Equal(tc.wantedErr, err)
				acs.Assert().Equal(tc.want, got, " want = %v , got = %v ", tc.want, got)
				tc.fields.account.AssertExpectations(acs.T())
			})
	}
}
