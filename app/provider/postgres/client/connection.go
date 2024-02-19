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
		slog.Error("error on connect")
		return nil, err
	}
	db.Ping()

	db2 := sqlx.NewDb(db, "postgres")

	return &Connection{
		DB:      db2,
		Options: opt,
	}, nil
}
