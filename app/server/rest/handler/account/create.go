package account

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jsperandio/transaction/app/domain/service"
	"github.com/jsperandio/transaction/app/server/rest/model/request"
	"github.com/jsperandio/transaction/app/server/rest/model/response"
	"github.com/labstack/echo/v4"
)

type CreateHandler struct {
	svc       service.AccountCreator
	validator *validator.Validate
}

func NewCreateHandler(ac service.AccountCreator, vld *validator.Validate) *CreateHandler {
	return &CreateHandler{
		svc:       ac,
		validator: vld,
	}
}

func (ch CreateHandler) RegisterRoute(e *echo.Echo) {
	e.POST("/accounts", ch.Handle)
}

// Create godoc
//
//	@Summary		Create an account
//	@Description	Create an account by given document number
//	@Accept			json
//	@Produce		json
//	@Tags			Account
//	@Param			account	body		request.CreateAccount	true	"document number for account"
//	@Success		201	{object}	response.Account
//	@Failure		400	{object}	error
//	@Failure		406	{object}	error
//	@Router			/accounts [post]
func (ch CreateHandler) Handle(e echo.Context) error {
	var req request.CreateAccount

	ctx := e.Request().Context()
	err := e.Bind(&req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "invalid data received")
	}

	err = ch.validate(req)
	if err != nil {
		return ch.JSONValidateError(e, err)
	}

	acc, err := ch.svc.Create(ctx, req.ToDomainModel())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	return e.JSONPretty(http.StatusCreated, response.NewAccountFromDomain(*acc), "	")
}

func (ch *CreateHandler) validate(r request.CreateAccount) error {
	err := ch.validator.Struct(r)
	if err != nil {
		return err
	}

	return nil
}

func (ch *CreateHandler) JSONValidateError(c echo.Context, err error) error {
	fmtVldtErr := response.NewFormattedValidationError(err)
	return c.JSONPretty(http.StatusUnprocessableEntity, fmtVldtErr, "	")
}
