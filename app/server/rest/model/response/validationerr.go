package response

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jsperandio/transaction/app/domain/model"
	"github.com/labstack/echo/v4"
)

type Error struct {
	HTTPStatusCode int    `json:"httpStatusCode"`
	Message        string `json:"message"`
}

type ValidationError struct {
	FieldPath  string      `json:"path"`
	Field      string      `json:"field"`
	FieldType  string      `json:"-"`
	Value      interface{} `json:"value"`
	ValidParam string      `json:"-"`
	Tag        string      `json:"-"`
	Message    string      `json:"message"`
}

type FormattedValidationError struct {
	Error
	ValidationErrors []ValidationError `json:"validationErrors,omitempty"`
}

func NewFormattedValidationError(err error) *FormattedValidationError {
	fve := &FormattedValidationError{
		Error: Error{
			HTTPStatusCode: http.StatusUnprocessableEntity,
			Message:        "The server understands the content type of the request entity but was unable to process the contained instructions.",
		},
		ValidationErrors: nil,
	}

	vve := validator.ValidationErrors{}
	if errors.As(err, &vve) {
		fve.ValidationErrors = append(fve.ValidationErrors, fve.BuildValidationError(vve)...)
	}

	return fve
}

func (fme *FormattedValidationError) BuildValidationError(vve validator.ValidationErrors) []ValidationError {
	var listve []ValidationError

	for i := range vve {
		nfve := ValidationError{
			FieldPath:  vve[i].StructNamespace(),
			Field:      vve[i].Field(),
			FieldType:  vve[i].Type().String(),
			Value:      vve[i].Value(),
			Tag:        vve[i].Tag(),
			ValidParam: vve[i].Param(),
		}
		nfve.BuildErrorMessage()
		listve = append(listve, nfve)
	}

	return listve
}

func (ve *ValidationError) BuildErrorMessage() {
	ve.Message = ve.TagErrorDict() + "."
}

func (ve *ValidationError) TagErrorDict() string {
	errMap := map[string]string{
		// Comparisons:
		"gt":  fmt.Sprintf("{%v} must be greater than %v", ve.Field, ve.ValidParam),             // Greater than
		"gte": fmt.Sprintf("{%v} must be greater than or equal to %v", ve.Field, ve.ValidParam), // Greater than or equal
		// Other:
		"len":      fmt.Sprintf("{%v} must have length of %v", ve.Field, ve.ValidParam),              // String length
		"max":      fmt.Sprintf("{%v} must have length less than %v", ve.Field, ve.ValidParam),       // String max length
		"min":      fmt.Sprintf("{%v} must have length greater than %v", ve.Field, ve.ValidParam),    // String min length
		"required": fmt.Sprintf("{%v} is a required field with type %v", ve.Field, ve.FieldType),     // Required
		"oneof":    fmt.Sprintf("{%v} must be one of desired values: [%v]", ve.Field, ve.ValidParam), // One of
	}

	return errMap[ve.Tag]
}

func JSONValidateError(e echo.Context, err error) error {
	fmtVldtErr := NewFormattedValidationError(err)
	return e.JSONPretty(http.StatusUnprocessableEntity, fmtVldtErr, "	")
}

func JSONMappedError(e echo.Context, err error) error {
	fmterr := Error{
		Message: err.Error(),
	}

	switch {
	case errors.Is(err, model.ErrAccountAlreadyExists):
		fmterr.HTTPStatusCode = http.StatusConflict
	case errors.Is(err, model.ErrAccountNotFound):
		fmterr.HTTPStatusCode = http.StatusNotFound
	case errors.Is(err, model.ErrInvalidTransaction):
		fmterr.HTTPStatusCode = http.StatusBadRequest
	default:
		fmterr.HTTPStatusCode = http.StatusInternalServerError
	}

	return e.JSON(fmterr.HTTPStatusCode, &FormattedValidationError{
		Error:            fmterr,
		ValidationErrors: nil,
	})
}
