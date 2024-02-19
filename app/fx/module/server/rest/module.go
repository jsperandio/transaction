package rest

import (
	"github.com/go-playground/validator/v10"
	fxhndlr "github.com/jsperandio/transaction/app/fx/module/server/rest/handler"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(validator.New),
		fxhndlr.Module(),
	)
}
