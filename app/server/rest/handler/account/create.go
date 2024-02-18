package account

import (
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jsperandio/transaction/app/domain/service"
	"github.com/jsperandio/transaction/app/server/rest/model/request"
	"github.com/jsperandio/transaction/app/server/rest/model/response"
	"github.com/labstack/echo/v4"
)

type CreateHandler struct {
	service   service.AccountCreator
	validator *validator.Validate
}

func NewCreateHandler(ac service.AccountCreator, vld *validator.Validate) *CreateHandler {
	return &CreateHandler{
		service:   ac,
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
//	@Success		201		{object}	response.Account
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Failure		422		{object}	response.FormattedValidationError
//	@Router			/accounts [post]
func (ch CreateHandler) Handle(e echo.Context) error {
	var req request.CreateAccount

	slog.Debug("incomming account create request")

	ctx := e.Request().Context()
	err := e.Bind(&req)
	if err != nil {
		slog.Error("can't bind body data", "err:", err.Error())
		return e.JSON(http.StatusBadRequest, "invalid data received")
	}

	err = ch.validate(req)
	if err != nil {
		return response.JSONValidateError(e, err)
	}

	acc, err := ch.service.Create(ctx, req.ToDomainModel())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}

	slog.Debug("account created successfully", "acc:", acc)
	return e.JSONPretty(http.StatusCreated, response.NewAccountFromDomain(acc), "	")
}

func (ch *CreateHandler) validate(r request.CreateAccount) error {
	err := ch.validator.Struct(r)
	if err != nil {
		slog.Error("validation check error")
		return err
	}

	return nil
}
