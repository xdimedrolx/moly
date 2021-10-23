package middleware

import (
	"github.com/labstack/echo/v4"
	httpContext "github.com/xdimedrolx/moly/pkg/platform/http/context"
)

func Context(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sc := httpContext.New(c, c.Request().Context())
		return next(sc)
	}
}