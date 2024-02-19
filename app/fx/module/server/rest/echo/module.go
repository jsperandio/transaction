package echo

import (
	"log"
	"sync"

	fxctx "github.com/jsperandio/transaction/app/fx/module/server/context"
	e "github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

var once sync.Once

func Module() fx.Option {
	options := fx.Options()

	once.Do(func() {
		options = fx.Options(
			fxctx.Module(),
			fx.Provide(
				func() *e.Echo {
					ec := e.New()
					ec.Use(mw.LoggerWithConfig(mw.LoggerConfig{
						Format: "{time:${time_rfc3339}, status=${status}, method=${method}, uri=${uri}, user_agent:${user_agent}} latency=${latency_human}\n",
						Output: ec.Logger.Output(),
					}))
					ec.Use(mw.Recover())
					return ec
				},
			),
			fx.Invoke(registerRoutes),
		)
	})

	return options
}

func registerRoutes(params Params) {
	params.Echo.OnAddRouteHandler = func(host string, route e.Route, handler e.HandlerFunc, middleware []e.MiddlewareFunc) {
		log.Printf("Register route  |%-6s - %-40s|", route.Method, route.Path)
	}

	for _, h := range params.Handlers {
		h.RegisterRoute(params.Echo)
	}
}
