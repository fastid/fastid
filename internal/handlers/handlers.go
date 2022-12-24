package handlers

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/services"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Handlers interface {
	Register(router *echo.Group)
	HealthCheck() HealthCheckHandler
	Key() KeyHandler
	Server() ServerHandler
}

type handlers struct {
	cfg         *config.Config
	log         *log.Logger
	srv         services.Services
	healthCheck HealthCheckHandler
	key         KeyHandler
	server      ServerHandler
}

func New(cfg *config.Config, log *log.Logger, srv services.Services) Handlers {
	healthCheck := NewHealthCheckHandler(cfg, log, srv)
	key := NewKeyHandler(cfg, log, srv)
	server := NewServerHandler(cfg, log, srv)

	return &handlers{
		cfg:         cfg,
		log:         log,
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
