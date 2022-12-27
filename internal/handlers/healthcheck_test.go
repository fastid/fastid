package handlers

import (
	"context"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/fastid/fastid/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	// Logger
	log := logger.New(cfg)

	// Context
	ctx := context.Background()

	// DB
	database, err := db.New(cfg, ctx)
	require.NoError(t, err)

	// Storage
	repos := repositories.New(cfg, log, database)

	// Service
	srv := services.New(cfg, log, repos)

	// Handlers
	handler := New(cfg, log, srv)

	// Echo
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	group := e.Group("/api/v1")
	handler.Register(group)

	t.Run("GET", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/healthcheck/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		HealthCheckGet := handler.HealthCheck().get()

		if assert.NoError(t, HealthCheckGet(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

}
