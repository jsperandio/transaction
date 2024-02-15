package client

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connection struct {
	DB      *sqlx.DB
	Options *Options
}

func NewConnection(opt *Options) (*Connection, error) {
	cs := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		opt.Username,
		opt.Password,
		opt.DatabaseURL,
		opt.DatabaseName)

	con, err := sqlx.Connect("postgres", cs)
	if err != nil {
		return nil, err
	}

	return &Connection{
		DB:      con,
		Options: opt,
	}, nil
}
