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

func Module() fx.Option {
	return fx.Options(
		AccountCreationModule(),
	)
}
