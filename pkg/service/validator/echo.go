package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type EchoValidator struct {
	validator *validator.Validate
}

func NewEchoValidator() *EchoValidator {
	return &EchoValidator{
		validator: validator.New(),
	}
}

func (v *EchoValidator) Validate(i interface{}) error {

	v.validator.RegisterValidation()
	if err := v.validator.Struct(i); err != nil {
		var errors []*ValidationErrorResponse

		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				element := new(ValidationErrorResponse)

				element.Filed = strings.ToLower(err.Field())
				element.Rule = err.Tag()
				element.Value = err.Param()

				errors = append(errors, element)
			}
		}

		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors) // TODO: Add error struct
	}

	return nil
}
