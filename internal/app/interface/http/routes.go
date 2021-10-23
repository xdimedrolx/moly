package http

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	useCase "github.com/xdimedrolx/moly/internal/app/application/use_case"
	_ "github.com/xdimedrolx/moly/internal/app/interface/http/docs"
	"github.com/xdimedrolx/moly/internal/app/service"
	"github.com/xdimedrolx/moly/pkg/platform/http"
)

// @title Server API
// @version 1.0
type RouteHandlers struct {
	services service.Services
	useCases useCase.Handlers
}

func NewRouteHandlers(services service.Services, useCases useCase.Handlers) *RouteHandlers {
	return &RouteHandlers{services, useCases}
}

func (r *RouteHandlers) RegisterRoutes(s *http.Server) {
	s.GET("/apidoc/*", echoSwagger.WrapHandler)

	webhook := s.Group("/webhook")
	webhook.GET("/echo", r.GetWebhookEcho)
}
