package client

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connection struct {
	DB      *sqlx.DB
	Options *Options
}

func NewConnection(opt *Options) (*Connection, error) {
	conn, err := sqlx.Connect("postgres", opt.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return &Connection{
		DB:      conn,
		Options: opt,
	}, nil
}
