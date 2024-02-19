package echo

import (
	"context"

	"github.com/jsperandio/transaction/app/server"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func ListenAndServe(lc fx.Lifecycle, instance *echo.Echo, opt *server.Options) {
	lc.Append(fx.Hook{
		OnStart: func(context context.Context) error {
			go instance.Start(":" + opt.Port)
			return nil
		},
		OnStop: func(context context.Context) error {
			return instance.Close()
		},
	})
}
