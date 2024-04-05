package dao

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/jsperandio/transaction/app/provider/postgres/client"
	dbmodel "github.com/jsperandio/transaction/app/provider/postgres/model"
)

type Account struct {
	conn *client.Connection
}

func NewAccount(c *client.Connection) *Account {
	return &Account{
		conn: c,
	}
}

func (a *Account) Save(ctx context.Context, acc *model.Account) (*model.Account, error) {
	dba := dbmodel.NewAccountFromDomain(acc)

	slog.Debug("saving account", "account", dba)

	err := a.conn.DB.Get(dba, "INSERT INTO account (document_number, avaliable_credit_limit) VALUES ($1,$2) RETURNING id", dba.DocumentNumber, dba.AvaliableCreditLimit)
	if err != nil {
		slog.Error("error saving account", "error", err)
		return nil, err
	}

	dm := dba.ToDomainModel()
	slog.Debug("account saved", "account", dm)

	return dm, nil
}

func (a *Account) Update(ctx context.Context, acc *model.Account) error {
	dba := dbmodel.NewAccountFromDomain(acc)

	slog.Debug("udapte account", "account", dba)

	res, err := a.conn.DB.NamedExec("UPDATE account SET avaliable_credit_limit = :avaliable_credit_limit", dba)
	if err != nil {
		slog.Error("error update account", "error", err)
		return err
	}

	cnt, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if cnt < 1 {
		return errors.New("can't update account")
	}

	slog.Debug("account saved")

	return nil
}

func (a *Account) Get(ctx context.Context, ID int64) (*model.Account, error) {
	acc := dbmodel.Account{}

	slog.Debug("getting account", "id", ID)

	err := a.conn.DB.Get(&acc, "SELECT id, document_number,avaliable_credit_limit  FROM account WHERE id = $1", ID)
	if err != nil {
		slog.Error("error getting account", "error", err)
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		slog.Debug("account not found", "id", ID)
		return nil, nil
	}

	dm := acc.ToDomainModel()
	slog.Debug("account found", "account", dm)

	return dm, nil
}

func (a *Account) FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Account, error) {
	acc := dbmodel.Account{}

	slog.Debug("getting account by document number", "document_number", documentNumber)

	err := a.conn.DB.Get(&acc, "SELECT id, document_number,avaliable_credit_limit FROM account WHERE document_number = $1", documentNumber)
	if err != nil {
		slog.Error("error getting account by document number", "error", err)
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		slog.Debug("account not found by document number", "document_number", documentNumber)
		return nil, nil
	}

	dm := acc.ToDomainModel()
	slog.Debug("account found by document number", "account", dm)

	return dm, nil
}
