package http

import (
	"github.com/labstack/echo/v4"
)

type (
	ServerHeaderConfig struct {
		Header string
	}
)

var (
	DefaultServerHeaderConfig = ServerHeaderConfig{
		Header: "FastID",
	}
)

func ServerHeaderWithConfig(config ServerHeaderConfig) echo.MiddlewareFunc {
	if config.Header == "" {
		config.Header = DefaultServerHeaderConfig.Header
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderServer, config.Header)
			return next(c)
		}
	}
}
