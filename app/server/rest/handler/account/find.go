package account

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jsperandio/transaction/app/domain/service"
	"github.com/jsperandio/transaction/app/server/rest/model/request"
	"github.com/jsperandio/transaction/app/server/rest/model/response"
	"github.com/labstack/echo/v4"
)

type FindHandler struct {
	service   service.AccountSearcher
	validator *validator.Validate
}

func NewFindHandler(service service.AccountSearcher, vld *validator.Validate) *FindHandler {
	return &FindHandler{
		service:   service,
		validator: vld,
	}
}

func (fh FindHandler) RegisterRoute(e *echo.Echo) {
	e.GET("/accounts/:accountId", fh.Handle)
}

func (fh FindHandler) Handle(e echo.Context) error {
	var req request.FindAccount

	ctx := e.Request().Context()
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "invalid data param")
	}

	err = fh.validate(req)
	if err != nil {
		return response.JSONValidateError(e, err)
	}

	acc, err := fh.service.GetByID(ctx, req.AccountID)
	if err != nil {
		return err
	}
	return e.JSONPretty(http.StatusOK, response.NewAccountFromDomain(acc), "	")
}

func (fh FindHandler) validate(r request.FindAccount) error {
	err := fh.validator.Struct(r)
	if err != nil {
		return err
	}

	return nil
}
