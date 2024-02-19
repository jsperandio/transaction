package transaction

import (
	"github.com/jsperandio/transaction/app/server/rest"
	htxn "github.com/jsperandio/transaction/app/server/rest/handler/transaction"
	"go.uber.org/fx"
)

func CreateModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				htxn.NewCreateHandler,
				fx.As(new(rest.Handler)),
				fx.ResultTags(`group:"handlers"`),
			),
		),
	)
}

func Module() fx.Option {
	return fx.Options(
		CreateModule(),
	)
}
