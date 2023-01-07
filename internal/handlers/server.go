package handlers

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
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
}

func (h *serverHandler) get() echo.HandlerFunc {

	type Response struct {
		PasswordMinLength    int    `json:"password_min_length"`
		PasswordMaxLength    int    `json:"password_man_length"`
		PasswordValidatorURL string `json:"password_validator_url"`
		EmailValidatorURL    string `json:"email_validator_url"`
	}

	return func(e echo.Context) error {
		response := &Response{
			PasswordMinLength:    h.cfg.PasswordMinLength,
			PasswordMaxLength:    h.cfg.PasswordMaxLength,
			PasswordValidatorURL: h.cfg.PasswordValidatorURL.String(),
			EmailValidatorURL:    h.cfg.EmailValidatorURL.String(),
		}
		return e.JSON(http.StatusOK, response)
	}
}
