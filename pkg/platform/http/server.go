package http

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/xdimedrolx/moly/pkg/log"
	logAdapter "github.com/xdimedrolx/moly/pkg/platform/http/log"
	"github.com/xdimedrolx/moly/pkg/platform/http/middleware"
	"logur.dev/logur"
)

type Server struct {
	*echo.Echo
	logger logur.Logger
}

func NewServerWithLogger(logger log.Logger) *Server {
	e := echo.New()
	e.Logger = logAdapter.NewLogurAdapter(logger)

	s := &Server{Echo: e, logger: logger}

	e.Use(echoMiddleware.RequestID())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Logger())

	AddValidator(s)
	e.Use(middleware.Validator)

	return s
}
