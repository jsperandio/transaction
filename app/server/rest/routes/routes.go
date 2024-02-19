package routes

import (
	"github.com/jsperandio/transaction/app/server/rest"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(instance *echo.Echo, handlers []rest.Handler) {
	for _, h := range handlers {
		h.RegisterRoute(instance)
	}
}
