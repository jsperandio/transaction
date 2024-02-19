package client

import (
	"database/sql"
	"fmt"
	"log/slog"

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

	db, err := sql.Open("postgres", cs)
	if err != nil {
		slog.Error("error on connect", "cs", cs)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		slog.Error("error on ping", "cs", cs)
		return nil, err
	}

	return &Connection{
		DB:      sqlx.NewDb(db, "postgres"),
		Options: opt,
	}, nil
}
