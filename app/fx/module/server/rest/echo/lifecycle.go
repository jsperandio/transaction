package echo

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

const port = "8081"

func ListenAndServe(lc fx.Lifecycle, instance *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(context context.Context) error {
			go instance.Start(":" + port)
			return nil
		},
		OnStop: func(context context.Context) error {
			return instance.Close()
		},
	})
}
