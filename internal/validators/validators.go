package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

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
