package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	teardownSuite, _, e, handler, err := setupSuite(t)
	defer teardownSuite(t)

	if err != nil {
		panic(err)
	}

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
