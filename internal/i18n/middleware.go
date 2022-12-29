package i18n

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/text/language"
)

type KeyContext string

func DetectLanguageWithConfig() echo.MiddlewareFunc {

	var matcher = language.NewMatcher([]language.Tag{
		language.MustParse("en-US"),
		language.MustParse("ru-RU"),
	})

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			accept := e.Request().Header.Get("Accept-Language")
			tag, _ := language.MatchStrings(matcher, accept)

			lang := "en"
			if tag.String() == "ru-RU" {
				lang = "ru"
			} else {
				lang = "en"
			}

			ctx := context.WithValue(e.Request().Context(), KeyContext("language"), lang)
			req := e.Request().WithContext(ctx)
			e.SetRequest(req)

			return next(e)
		}
	}
}
