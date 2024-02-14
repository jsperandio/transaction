package echo

import (
	"github.com/jsperandio/transaction/app/server/rest"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Handlers []rest.Handler `group:"handlers"`
	Echo     *echo.Echo
}
