package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h HealthHandler) RegisterRoute(e *echo.Echo) {
	e.GET("/health", h.Handle)
}

// Health godoc
//
//	@Summary		Check health of api
//	@Description	Check health of api
//	@Accept			json
//	@Produce		json
//	@Tags			App
//	@Success		200	{object}	string
//	@Failure		400	{object}	error
//	@Failure		406	{object}	error
//	@Router			/health [get]
func (HealthHandler) Handle(e echo.Context) (err error) {
	return e.JSON(http.StatusOK, "Ok")
}
