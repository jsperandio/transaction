package transaction

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
	service   service.TransactionCreator
	validator *validator.Validate
}

func NewCreateHandler(tc service.TransactionCreator, vld *validator.Validate) *CreateHandler {
	return &CreateHandler{
		service:   tc,
		validator: vld,
	}
}

func (ch CreateHandler) RegisterRoute(e *echo.Echo) {
	e.POST("/transactions", ch.Handle)
}

// Create godoc
//
//	@Summary		Create an transaction
//	@Description	Create an transaction
//	@Accept			json
//	@Produce		json
//	@Tags			Transaction
//	@Param			transaction	body		request.CreateTransaction	true	"values for transaction"
//	@Success		201			{object}	response.Transaction
//	@Failure		400			{object}	error
//	@Failure		422			{object}	response.FormattedValidationError
//	@Failure		500			{object}	response.FormattedValidationError
//	@Router			/transactions [post]
func (ch CreateHandler) Handle(e echo.Context) error {
	var req request.CreateTransaction

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

	txn, err := ch.service.Create(ctx, req.ToDomainModel())
	if err != nil {
		return response.JSONMappedError(e, err)
	}

	slog.Info("transaction created successfully", "id:", txn.ID)
	return e.JSONPretty(http.StatusCreated, response.NewTransactionFromDomain(txn), "	")
}

func (ch *CreateHandler) validate(r request.CreateTransaction) error {
	err := ch.validator.Struct(r)
	if err != nil {
		slog.Error("validation check error")
		return err
	}

	return nil
}
