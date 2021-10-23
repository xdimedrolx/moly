package app

import (
	"github.com/xdimedrolx/moly/internal/app/application/use_case"
	"github.com/xdimedrolx/moly/internal/app/config"
	routes "github.com/xdimedrolx/moly/internal/app/interface/http"
	"github.com/xdimedrolx/moly/internal/app/service"
	"github.com/xdimedrolx/moly/pkg/log"
	papp "github.com/xdimedrolx/moly/pkg/platform/app"
	"github.com/xdimedrolx/moly/pkg/platform/http"
)

type App struct {
	config     config.Config
	logger     Logger
	errorHandler ErrorHandler
	services    service.Services
	httpServer *http.Server
	useCases useCase.Handlers
}

func InitializeApp(
	config config.Config,
	logger Logger,
	errorHandler ErrorHandler,
) *App {
	services := initializeServices(config, logger, errorHandler)

	useCases := initializeUseCases(config, services)

	routeHandlers := routes.NewRouteHandlers(services, useCases)

	httpServer := initializeHttpServer(logger, routeHandlers)

	return &App{
		config,
		logger,
		errorHandler,
		services,
		httpServer,
		useCases,
	}
}

func initializeServices(config config.Config, logger Logger, errorHandler ErrorHandler) service.Services {
	return service.NewContainer(logger, errorHandler)
}

func initializeUseCases(config config.Config, services service.Services) useCase.Handlers {
	return useCase.New(services)
}

func initializeHttpServer(logger log.Logger, routes *routes.RouteHandlers) *http.Server {
	server := http.NewServerWithLogger(logger)
	routes.RegisterRoutes(server)
	return server
}

func (a *App) Dispose() {
	a.logger.Info("disposing...")
}

func (a *App) Logger() log.Logger {
	return a.logger
}

func (a *App) HttpServer() *http.Server {
	return a.httpServer
}

var _ papp.App = (*App)(nil)
