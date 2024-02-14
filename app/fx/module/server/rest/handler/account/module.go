package account

import (
	"github.com/jsperandio/transaction/app/server/rest"
	hacc "github.com/jsperandio/transaction/app/server/rest/handler/account"
	"go.uber.org/fx"
)

func CreateModule() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				hacc.NewCreateHandler,
				fx.As(new(rest.Handler)),
				fx.ResultTags(`group:"handlers"`),
			),
		),
	)
}
