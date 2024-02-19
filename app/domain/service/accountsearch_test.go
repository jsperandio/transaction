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

type AccountSearchSuite struct {
	suite.Suite
}

func TestNewAccountSearchSuite(t *testing.T) {
	suite.Run(t, new(AccountSearchSuite))
}

func (acs *AccountSearchSuite) Test_NewAccountSearch() {
	type args struct {
		repository repository.Account
	}

	type test struct {
		name string
		args args
		want *AccountSearch
	}

	tests := []test{
		{
			name: "should return a new AccountSearch",
			args: args{
				repository: new(mocks.Account),
			},
			want: &AccountSearch{
				repository: new(mocks.Account),
			},
		},
		{
			name: "should return an empty AccountSearch",
			args: args{},
			want: &AccountSearch{},
		},
	}

	for _, tc := range tests {
		acs.Run(
			tc.name,
			func() {
				ac := NewAccountSearch(tc.args.repository)
				acs.Assert().Equal(tc.want, ac)
			})
	}
}

func (acs *AccountSearchSuite) Test_AccountSearch_GetByID() {
	type fields struct {
		account *mocks.Account
	}

	type args struct {
		ctx context.Context
		id  int64
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
			name: "when Get return an error, should return nil account and an error ",
			fields: fields{
				account: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want:      nil,
			wantErr:   true,
			wantedErr: errors.New("error"),
			mock: func(account *mocks.Account) {
				account.On("Get", context.Background(), int64(1)).Return(nil, errors.New("error"))
			},
		},
		{
			name: "when Get return success but no account, should return nil account and ErrAccountNotFound",
			fields: fields{
				account: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				id:  2,
			},
			want:      nil,
			wantErr:   true,
			wantedErr: model.ErrAccountNotFound,
			mock: func(account *mocks.Account) {
				account.On("Get", context.Background(), int64(2)).Return(nil, nil)
			},
		},
		{
			name: "when success, should return account and nil error",
			fields: fields{
				account: new(mocks.Account),
			},
			args: args{
				ctx: context.Background(),
				id:  3,
			},
			want: &model.Account{
				ID:             3,
				DocumentNumber: "3333333",
			},
			wantErr:   false,
			wantedErr: nil,
			mock: func(account *mocks.Account) {
				account.On("Get", context.Background(), int64(3)).Return(&model.Account{
					ID:             3,
					DocumentNumber: "3333333",
				}, nil)
			},
		},
	}

	for _, tc := range tests {
		acs.Run(
			tc.name,
			func() {
				tc.mock(tc.fields.account)
				ac := NewAccountSearch(tc.fields.account)

				got, err := ac.GetByID(tc.args.ctx, tc.args.id)
				acs.Assert().Equal(tc.wantErr, err != nil)
				acs.Assert().Equal(tc.wantedErr, err)
				acs.Assert().Equal(tc.want, got, " want = %v , got = %v ", tc.want, got)
				tc.fields.account.AssertExpectations(acs.T())
			})
	}
}
