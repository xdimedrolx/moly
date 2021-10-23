package main

import (
	"context"
	"emperror.dev/errors/match"
	"fmt"
	"github.com/cloudflare/tableflip"
	"github.com/oklog/run"
	"github.com/xdimedrolx/moly/internal/app"
	papp "github.com/xdimedrolx/moly/pkg/platform/app"
	"github.com/xdimedrolx/moly/pkg/platform/config"
	"github.com/xdimedrolx/moly/pkg/platform/log"
	"github.com/xdimedrolx/moly/pkg/platform/log/common"
	_ "logur.dev/logur"
	"os"
	"os/signal"
	"syscall"
	"time"

	"emperror.dev/emperror"
	logurhandler "emperror.dev/handler/logur"

	"github.com/pkg/errors"
	appkiterrors "github.com/sagikazarmark/appkit/errors"
	appkitrun "github.com/sagikazarmark/appkit/run"
	"github.com/urfave/cli/v2"
)

func main() {
	var configPath string

	defaultLogger := log.NewLogger(log.NewDefaultConfig())

	cliApp := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "./config.toml",
				Usage:       "Load configuration from `FILE`",
				Destination: &configPath,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Action: func(c *cli.Context) error {
					cfg := initConfig(configPath)

					// Create logger (first thing after configuration loading)
					logger := log.NewLogger(cfg.Log)
					// Override the global standard library logger to make sure everything uses our logger
					log.SetStandardLogger(logger)

					err := cfg.Validate()
					if err != nil {
						logger.Error(err.Error())
						os.Exit(3)
					}

					// Configure error handler
					errorHandler := logurhandler.New(logger)
					defer emperror.HandleRecover(errorHandler)

					// configure graceful restart
					upg, _ := tableflip.New(tableflip.Options{})

					// Do an upgrade on SIGHUP
					go func() {
						ch := make(chan os.Signal, 1)
						signal.Notify(ch, syscall.SIGHUP)
						for range ch {
							logger.Info("graceful reloading")

							_ = upg.Upgrade()
						}
					}()

					var group run.Group

					// Set up application
					{
						logger := common.NewContextAwareLogger(logger, papp.ContextExtractor)
						errorHandler := emperror.WithFilter(
							emperror.WithContextExtractor(errorHandler, papp.ContextExtractor),
							appkiterrors.IsServiceError, // filter out service errors
						)

						application := app.InitializeApp(cfg.App, logger, errorHandler)
						defer application.Dispose()

						group.Add(
							func() error { return application.HttpServer().Start(fmt.Sprintf(":%d", cfg.HttpPort)) },
							func(err error) {
								name := "httpServer"
								ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
								defer cancel()

								logger.Debug(fmt.Sprintf("starting graceful shutdown (%s)", name))
								if err := application.HttpServer().Shutdown(ctx); err != nil {
									logger.Error(fmt.Sprintf("graceful shutdown (%s): %v", name, err))
								}
							},
						)
					}

					// Setup signal handler
					group.Add(run.SignalHandler(context.Background(), syscall.SIGINT, syscall.SIGTERM))

					// Setup graceful restart
					group.Add(appkitrun.GracefulRestart(context.Background(), upg))

					err = group.Run()
					emperror.WithFilter(errorHandler, match.As(&run.SignalError{}).MatchError).Handle(err)

					return nil
				},
			},
		},
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		defaultLogger.Error(err.Error())
	}
}

func initConfig(configPath string) Config {
	cfg := NewDefaultConfig()

	err := config.Load(configPath, &cfg)
	emperror.Panic(errors.Wrap(err, "failed to load configuration"))

	return cfg
}
