package handler

import (
	fxhndacc "github.com/jsperandio/transaction/app/fx/module/server/rest/handler/account"
	"github.com/jsperandio/transaction/app/server/rest"
	"github.com/jsperandio/transaction/app/server/rest/handler"
	"go.uber.org/fx"
)

func HealthHandler() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				handler.NewHealthHandler,
				fx.As(new(rest.Handler)),
				fx.ResultTags(`group:"handlers"`),
			),
		),
	)
}

func SwaggerHandler() fx.Option {
	return fx.Options(fx.Provide(
		fx.Annotate(
			handler.NewSwaggerHandler,
			fx.As(new(rest.Handler)),
			fx.ResultTags(`group:"handlers"`),
		),
	),
	)
}

func Module() fx.Option {
	return fx.Options(
		HealthHandler(),
		SwaggerHandler(),
		fxhndacc.CreateModule(),
	)
}
