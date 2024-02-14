package provider

import (
	"sync"

	"github.com/jsperandio/transaction/app/domain/repository"
	pgconn "github.com/jsperandio/transaction/app/provider/postgres/client"
	pvdr "github.com/jsperandio/transaction/app/provider/postgres/dao"
	"go.uber.org/fx"
)

var once sync.Once

func DBConnectionModule() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			fx.Provide(
				pgconn.DefaultOptions,
				pgconn.NewConnection,
			),
		)
	})

	return options
}

func AccountModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				pvdr.NewAccount,
				fx.As(new(repository.Account)),
			),
		),
	)
}

func TransactionModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				pvdr.NewTransaction,
				fx.As(new(repository.Transaction)),
			),
		),
	)
}

func Module() fx.Option {
	return fx.Options(
		DBConnectionModule(),
		AccountModule(),
		TransactionModule(),
	)
}
