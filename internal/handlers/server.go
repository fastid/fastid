package handlers

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/services"
	"github.com/fastid/fastid/internal/validators"
	"github.com/ggwhite/go-masker"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type ServerHandler interface {
	Register(router *echo.Group)
	get() echo.HandlerFunc
	//post() echo.HandlerFunc
	//patch() echo.HandlerFunc
}

type serverHandler struct {
	cfg    *config.Config
	logger logger.Logger
	srv    services.Services
}

func NewServerHandler(cfg *config.Config, logger logger.Logger, srv services.Services) ServerHandler {
	return &serverHandler{cfg: cfg, logger: logger, srv: srv}
}

func (h *serverHandler) Register(router *echo.Group) {
	router.Add("GET", "/server/", h.get())
	router.Add("POST", "/server/", h.post())
	//router.Add("PATCH", "/server/", h.patch())
}

func (h *serverHandler) get() echo.HandlerFunc {
	//Ok        string = "ok"
	//Locked    string = "locked"

	const (
		NeedSetup string = "need_setup"
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

	type Response struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	return func(e echo.Context) error {
		u := new(Request)

		if err := e.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		ctx := e.Request().Context()

		u.Email = strings.TrimLeft(u.Email, " ")
		u.Email = strings.TrimRight(u.Email, " ")
		u.Password = strings.TrimLeft(u.Password, " ")
		u.Password = strings.TrimRight(u.Password, " ")

		if err := e.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, validators.Parse(ctx, err))
		}

		username := h.cfg.ADMIN.USERNAME
		email := u.Email
		password := u.Password

		h.logger.Infof(ctx, "Create super user (username:%s email:%s, password:%s)", username, email, masker.Password(password))
		return e.JSON(http.StatusCreated, &Response{Username: username, Email: email})
	}
}

//func (h *serverHandler) patch() echo.HandlerFunc {
//	type Response struct {
//	}
//
//	type Request struct {
//		Key string `json:"key" validate:"required"`
//	}
//
//	return func(e echo.Context) error {
//		u := new(Request)
//		if err := e.Bind(u); err != nil {
//			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//		}
//
//		u.Key = strings.TrimLeft(u.Key, " ")
//		u.Key = strings.TrimRight(u.Key, " ")
//
//		if err := e.Validate(u); err != nil {
//			return echo.NewHTTPError(http.StatusBadRequest, validators.Parse(e.Request().Context(), err))
//		}
//
//		ctx := e.Request().Context()
//
//		h.srv.Server().UnlockDatabase(ctx, u.Key)
//
//		return e.JSON(http.StatusOK, &Response{})
//	}
//
//}
