package handlers

import (
	"context"
	"encoding/json"
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/db"
	"github.com/fastid/fastid/internal/logger"
	"github.com/fastid/fastid/internal/repositories"
	"github.com/fastid/fastid/internal/services"
	"github.com/fastid/fastid/internal/validators"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
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

	// Validator
	validator := validators.New()
	validator.Register(e)

	group := e.Group("/api/v1")
	handler.Register(group)

	t.Run("GET", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/server/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		serverGet := handler.Server().get()
		if assert.NoError(t, serverGet(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var result map[string]any

			err := json.Unmarshal(rec.Body.Bytes(), &result)
			require.NoError(t, err)
			require.NotEmpty(t, result["email_validator_url"])
			require.NotEmpty(t, result["password_man_length"])
			require.NotEmpty(t, result["password_min_length"])
		}
	})

	t.Run("POST", func(t *testing.T) {
		bodyJSON := `{"email": "user_admin@exmaple.com", "password": "password"}`

		req := httptest.NewRequest(http.MethodPost, "/api/v1/server/", strings.NewReader(bodyJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		serverPost := handler.Server().post()

		if assert.NoError(t, serverPost(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)

			var result map[string]any

			err := json.Unmarshal(rec.Body.Bytes(), &result)
			require.NoError(t, err)
			require.NotEmpty(t, result["key"])
			require.Equal(t, result["email"], "user_admin@exmaple.com")
			require.Equal(t, result["username"], cfg.ADMIN.USERNAME)
		}
	})

}
