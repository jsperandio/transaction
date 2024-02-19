package handler

import (
	echo "github.com/labstack/echo/v4"
	es "github.com/swaggo/echo-swagger"
)

type SwaggerHandler struct{}

func NewSwaggerHandler() *SwaggerHandler {
	return &SwaggerHandler{}
}

func (*SwaggerHandler) Handle(e echo.Context) error {
	return nil
}

func (*SwaggerHandler) RegisterRoute(instance *echo.Echo) {
	instance.GET("/swagger/*", es.WrapHandler)
}
