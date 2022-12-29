package handlers

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/services"
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	Register(router *echo.Group)
	HealthCheck() HealthCheckHandler
	Key() KeyHandler
	Server() ServerHandler
}

type handlers struct {
	cfg         *config.Config
	logger      logger.Logger
	srv         services.Services
	healthCheck HealthCheckHandler
	key         KeyHandler
	server      ServerHandler
}

func New(cfg *config.Config, logger logger.Logger, srv services.Services) Handlers {
	healthCheck := NewHealthCheckHandler(cfg, logger, srv)
	key := NewKeyHandler(cfg, logger, srv)
	server := NewServerHandler(cfg, logger, srv)

	return &handlers{
		cfg:         cfg,
		logger:      logger,
		srv:         srv,
		healthCheck: healthCheck,
		key:         key,
		server:      server,
	}
}

func (h *handlers) Register(router *echo.Group) {
	h.healthCheck.Register(router)
	h.key.Register(router)
	h.server.Register(router)
}

func (h *handlers) HealthCheck() HealthCheckHandler {
	return h.healthCheck
}

func (h *handlers) Key() KeyHandler {
	return h.key
}

func (h *handlers) Server() ServerHandler {
	return h.server
}
