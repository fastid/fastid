package handlers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	teardownSuite, _, e, handler, err := setupSuite(t)
	defer teardownSuite(t)

	if err != nil {
		panic(err)
	}

	t.Run("GET", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/server/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler.Server()
		serverGet := handler.Server().get()

		if assert.NoError(t, serverGet(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var result map[string]any

			err := json.Unmarshal(rec.Body.Bytes(), &result)
			require.NoError(t, err)
			require.NotEmpty(t, result["email_validator_url"])
			require.NotEmpty(t, result["password_man_length"])
			require.NotEmpty(t, result["password_min_length"])
			require.NotEmpty(t, result["password_validator_url"])
			require.NotEmpty(t, result["email_validator_url"])
		}
	})

	//t.Run("POST", func(t *testing.T) {
	//	bodyJSON := `{"email": "user_admin@exmaple.com", "password": "password"}`
	//
	//	req := httptest.NewRequest(http.MethodPost, "/api/v1/server/", strings.NewReader(bodyJSON))
	//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//	rec := httptest.NewRecorder()
	//	c := e.NewContext(req, rec)
	//
	//	serverPost := handler.Server().post()
	//
	//	if assert.NoError(t, serverPost(c)) {
	//		assert.Equal(t, http.StatusCreated, rec.Code)
	//
	//		var result map[string]any
	//
	//		err := json.Unmarshal(rec.Body.Bytes(), &result)
	//		require.NoError(t, err)
	//		require.NotEmpty(t, result["key"])
	//		require.Equal(t, result["email"], "user_admin@exmaple.com")
	//		require.Equal(t, result["username"], cfg.ADMIN.USERNAME)
	//	}
	//})
	//
	//t.Run("PATCH", func(t *testing.T) {
	//	bodyJSON := `{"key": "f09f3a530cb589dc833d0763689defc7b78594651f9fb7b03a07f40a180cda75"}`
	//
	//	req := httptest.NewRequest(http.MethodPatch, "/api/v1/server/", strings.NewReader(bodyJSON))
	//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//	rec := httptest.NewRecorder()
	//	c := e.NewContext(req, rec)
	//
	//	serverPost := handler.Server().patch()
	//
	//	if assert.NoError(t, serverPost(c)) {
	//		assert.Equal(t, http.StatusOK, rec.Code)
	//
	//		var result map[string]any
	//
	//		err := json.Unmarshal(rec.Body.Bytes(), &result)
	//		require.NoError(t, err)
	//		fmt.Println(result)
	//
	//	}
	//})
	//
	//t.Run("POST-Errors", func(t *testing.T) {
	//	bodyJSON := `{}`
	//
	//	req := httptest.NewRequest(http.MethodPost, "/api/v1/server/", strings.NewReader(bodyJSON))
	//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//	rec := httptest.NewRecorder()
	//	c := e.NewContext(req, rec)
	//
	//	serverPost := handler.Server().post()
	//	err := serverPost(c)
	//	require.Error(t, err)
	//})
	//
	//t.Run("POST-Error-Email", func(t *testing.T) {
	//	bodyJSON := `{"email": "userbox"}`
	//
	//	req := httptest.NewRequest(http.MethodPost, "/api/v1/server/", strings.NewReader(bodyJSON))
	//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//	rec := httptest.NewRecorder()
	//	c := e.NewContext(req, rec)
	//
	//	serverPost := handler.Server().post()
	//	err := serverPost(c)
	//	require.Error(t, err)
	//})
	//
	//t.Run("POST-Error-Email-Invalid", func(t *testing.T) {
	//	bodyJSON := `{"email": "fake-email"}`
	//
	//	req := httptest.NewRequest(http.MethodPost, "/api/v1/server/", strings.NewReader(bodyJSON))
	//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//	rec := httptest.NewRecorder()
	//	c := e.NewContext(req, rec)
	//
	//	serverPost := handler.Server().post()
	//	err := serverPost(c)
	//	require.Error(t, err)
	//})

}
