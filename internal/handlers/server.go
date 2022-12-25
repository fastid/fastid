package handlers

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type ServerHandler interface {
	Register(router *echo.Group)
	get() echo.HandlerFunc
	post() echo.HandlerFunc
}

type serverHandler struct {
	cfg *config.Config
	log *log.Logger
	srv services.Services
}

func NewServerHandler(cfg *config.Config, log *log.Logger, srv services.Services) ServerHandler {
	return &serverHandler{cfg: cfg, log: log, srv: srv}
}

func (h *serverHandler) Register(router *echo.Group) {
	router.Add("GET", "/server/", h.get())
	router.Add("POST", "/server/", h.post())
}

func (h *serverHandler) get() echo.HandlerFunc {

	const (
		Ok        string = "ok"
		NeedSetup string = "need_setup"
		Locked    string = "locked"
	)

	type Response struct {
		Status               string `json:"status"`
		PasswordMinLength    int    `json:"password_min_length"`
		PasswordMaxLength    int    `json:"password_man_length"`
		PasswordValidatorURL string `json:"password_validator_url"`
		EmailValidatorURL    string `json:"email_validator_url"`
	}

	return func(e echo.Context) error {
		response := &Response{
			Status:               NeedSetup,
			PasswordMinLength:    h.cfg.PasswordMinLength,
			PasswordMaxLength:    h.cfg.PasswordMaxLength,
			PasswordValidatorURL: h.cfg.PasswordValidatorURL.String(),
			EmailValidatorURL:    h.cfg.EmailValidatorURL.String(),
		}
		return e.JSON(http.StatusOK, response)
	}
}

func (h *serverHandler) post() echo.HandlerFunc {

	type Request struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	return func(e echo.Context) error {
		u := new(Request)

		if err := e.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		u.Email = strings.TrimLeft(u.Email, " ")
		u.Email = strings.TrimRight(u.Email, " ")
		u.Password = strings.TrimLeft(u.Password, " ")
		u.Password = strings.TrimRight(u.Password, " ")

		if err := e.Validate(u); err != nil {
			var errs []Errors
			var errMessage string

			for _, err := range err.(validator.ValidationErrors) {
				if err.Field() == "Email" && err.Tag() == "required" {
					errMessage = `The "Email" field is not filled`
				}

				if err.Field() == "Email" && err.Tag() == "email" {
					errMessage = "The email address is incorrect"
				}

				errs = append(errs, Errors{
					errMessage,
					err.Field(),
					err.Tag(),
					err.ActualTag(),
				})
			}
			return echo.NewHTTPError(http.StatusBadRequest, &Error{Message: "", Errors: errs})
		}

		return e.JSON(http.StatusOK, make(map[string]string))
	}
}
