package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
	"github.com/xdimedrolx/moly/pkg/platform/correlation"
)

type (
	// RequestIDConfig defines the config for RequestID middleware.
	RequestIDConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper

		// Generator defines a function to generate an ID.
		// Optional. Default value random.String(32).
		Generator func() string

		// RequestIDHandler defines a function which is executed for a request id.
		RequestIDHandler func(echo.Context, string)
	}
)

var (
	// DefaultRequestIDConfig is the default RequestID middleware config.
	DefaultRequestIDConfig = RequestIDConfig{
		Skipper:   middleware.DefaultSkipper,
		Generator: generator,
	}
)

// RequestID returns a X-Request-ID middleware.
func RequestID() echo.MiddlewareFunc {
	return RequestIDWithConfig(DefaultRequestIDConfig)
}

// RequestIDWithConfig returns a X-Request-ID middleware with config.
func RequestIDWithConfig(config RequestIDConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultRequestIDConfig.Skipper
	}
	if config.Generator == nil {
		config.Generator = generator
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			rid := req.Header.Get(echo.HeaderXRequestID)
			if rid == "" {
				rid = config.Generator()
			}

			req = req.WithContext(correlation.ToContext(req.Context(), rid))
			c.SetRequest(req)

			res.Header().Set(echo.HeaderXRequestID, rid)
			if config.RequestIDHandler != nil {
				config.RequestIDHandler(c, rid)
			}

			return next(c)
		}
	}
}

func generator() string {
	return random.String(32)
}