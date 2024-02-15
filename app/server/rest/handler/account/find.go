package account

import (
	"fmt"
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

// Find godoc
//
//	@Summary		Find account
//	@Description	Find an account by ID
//	@Accept			json
//	@Produce		json
//	@Tags			Account
//	@Param			accountId	path		int	true	"ID of desired account"
//	@Success		200	{object}	response.Account
//	@Failure		400	{object}	error
//	@Failure		422	{object}	response.FormattedValidationError
//	@Router			/accounts/{accountId} [get]
func (fh FindHandler) Handle(e echo.Context) error {
	var req request.FindAccount

	ctx := e.Request().Context()
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, fmt.Sprintf("invalid data param : %s", err.Error()))
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
