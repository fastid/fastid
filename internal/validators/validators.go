package validators

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"strings"
)

// Errors - structure for error response
type Errors struct {
	Message   string `json:"message"`
	Field     string `json:"field"`
	Tag       string `json:"tag"`
	ActualTag string `json:"actual_tag"`
}

type Error struct {
	Errors  []Errors `json:"errors,omitempty"`
	Message string   `json:"message,omitempty"`
}

type Validators interface {
	Validate(i interface{}) error
	Register(e *echo.Echo)
}

type validators struct {
	validator *validator.Validate
}

func New() Validators {
	return &validators{validator: validator.New()}
}

func (v *validators) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func (v *validators) Register(e *echo.Echo) {
	e.Validator = v
}

func Parse(ctx context.Context, err error) *Error {
	var errs []Errors
	var errMessage string

	for _, err := range err.(validator.ValidationErrors) {

		if err.Field() == "Email" && err.Tag() == "required" {
			errMessage = `The "Email" field is not filled`
		}

		if err.Field() == "Email" && err.Tag() == "email" {
			errMessage = "The email address is incorrect"
		}

		if err.Field() == "Password" && err.Tag() == "required" {
			errMessage = `The "Password" field is not filled`
		}

		errs = append(errs, Errors{
			errMessage,
			strings.ToLower(err.Field()),
			err.Tag(),
			err.ActualTag(),
		})
	}
	return &Error{Message: "", Errors: errs}
}
