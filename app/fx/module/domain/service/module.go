package service

import (
	"github.com/jsperandio/transaction/app/domain/service"
	"go.uber.org/fx"
)

func AccountCreationModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				service.NewAccountCreation,
				fx.As(new(service.AccountCreator)),
			),
		),
	)
}

func AccountSearchModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				service.NewAccountSearch,
				fx.As(new(service.AccountSearcher)),
			),
		),
	)
}

func TransactionCreationModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				service.NewTransactionCreation,
				fx.As(new(service.TransactionCreator)),
			),
		),
	)
}

func Module() fx.Option {
	return fx.Options(
		AccountCreationModule(),
		AccountSearchModule(),
		TransactionCreationModule(),
	)
}
