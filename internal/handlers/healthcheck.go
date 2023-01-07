package handlers

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthCheckHandler interface {
	Register(router *echo.Group)
	get() echo.HandlerFunc
}

type healthCheckHandler struct {
	cfg    *config.Config
	logger logger.Logger
	srv    services.Services
}

func NewHealthCheckHandler(cfg *config.Config, logger logger.Logger, srv services.Services) HealthCheckHandler {
	return &healthCheckHandler{cfg: cfg, logger: logger, srv: srv}
}

func (h *healthCheckHandler) Register(router *echo.Group) {
	router.Add("GET", "/healthcheck/", h.get())
}

func (h *healthCheckHandler) get() echo.HandlerFunc {
	return func(e echo.Context) error {
		return e.JSON(http.StatusOK, make(map[string]string))
	}
}
