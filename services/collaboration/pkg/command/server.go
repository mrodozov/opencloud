package command

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/oklog/run"
	"github.com/urfave/cli/v2"
	microstore "go-micro.dev/v4/store"

	"github.com/opencloud-eu/opencloud/pkg/config/configlog"
	"github.com/opencloud-eu/opencloud/pkg/registry"
	"github.com/opencloud-eu/opencloud/pkg/tracing"
	"github.com/opencloud-eu/opencloud/services/collaboration/pkg/config"
	"github.com/opencloud-eu/opencloud/services/collaboration/pkg/config/parser"
	"github.com/opencloud-eu/opencloud/services/collaboration/pkg/connector"
	"github.com/opencloud-eu/opencloud/services/collaboration/pkg/helpers"
	"github.com/opencloud-eu/opencloud/services/collaboration/pkg/logging"
	"github.com/opencloud-eu/opencloud/services/collaboration/pkg/server/debug"
	"github.com/opencloud-eu/opencloud/services/collaboration/pkg/server/grpc"
	"github.com/opencloud-eu/opencloud/services/collaboration/pkg/server/http"
	"github.com/opencloud-eu/reva/v2/pkg/rgrpc/todo/pool"
	"github.com/opencloud-eu/reva/v2/pkg/store"
)

// Server is the entrypoint for the server command.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "server",
		Usage:    fmt.Sprintf("start the %s service without runtime (unsupervised mode)", cfg.Service.Name),
		Category: "server",
		Before: func(c *cli.Context) error {
			return configlog.ReturnFatal(parser.ParseConfig(cfg))
		},
		Action: func(c *cli.Context) error {
			logger := logging.Configure(cfg.Service.Name, cfg.Log)
			traceProvider, err := tracing.GetServiceTraceProvider(cfg.Tracing, cfg.Service.Name)
			if err != nil {
				return err
			}

			gr := run.Group{}
			ctx, cancel := context.WithCancel(c.Context)
			defer cancel()

			// prepare components
			if err := helpers.RegisterOpenCloudService(ctx, cfg, logger); err != nil {
				return err
			}

			tm, err := pool.StringToTLSMode(cfg.CS3Api.GRPCClientTLS.Mode)
			if err != nil {
				return err
			}
			gatewaySelector, err := pool.GatewaySelector(
				cfg.CS3Api.Gateway.Name,
				pool.WithTLSCACert(cfg.CS3Api.GRPCClientTLS.CACert),
				pool.WithTLSMode(tm),
				pool.WithRegistry(registry.GetRegistry()),
				pool.WithTracerProvider(traceProvider),
			)
			if err != nil {
				return err
			}

			appUrls, err := helpers.GetAppURLs(cfg, logger)
			if err != nil {
				return err
			}

			ticker := time.NewTicker(cfg.CS3Api.APPRegistrationInterval)
			defer ticker.Stop()
			go func() {
				for ; true; <-ticker.C {
					if err := helpers.RegisterAppProvider(ctx, cfg, logger, gatewaySelector, appUrls); err != nil {
						logger.Warn().Err(err).Msg("Failed to register app provider")
					}
				}
			}()

			st := store.Create(
				store.Store(cfg.Store.Store),
				store.TTL(cfg.Store.TTL),
				microstore.Nodes(cfg.Store.Nodes...),
				microstore.Database(cfg.Store.Database),
				microstore.Table(cfg.Store.Table),
				store.Authentication(cfg.Store.AuthUsername, cfg.Store.AuthPassword),
			)

			// start GRPC server
			grpcServer, teardown, err := grpc.Server(
				grpc.AppURLs(appUrls),
				grpc.Config(cfg),
				grpc.Logger(logger),
				grpc.TraceProvider(traceProvider),
				grpc.Store(st),
			)
			defer teardown()
			if err != nil {
				logger.Error().Err(err).Str("transport", "grpc").Msg("Failed to initialize server")
				return err
			}

			gr.Add(func() error {
				l, err := net.Listen("tcp", cfg.GRPC.Addr)
				if err != nil {
					return err
				}
				return grpcServer.Serve(l)
			},
				func(err error) {
					if err != nil {
						logger.Info().
							Str("transport", "grpc").
							Str("server", cfg.Service.Name).
							Msg("Shutting down server")
					} else {
						logger.Error().Err(err).
							Str("transport", "grpc").
							Str("server", cfg.Service.Name).
							Msg("Shutting down server")
					}

					cancel()
				})

			// start debug server
			debugServer, err := debug.Server(
				debug.Logger(logger),
				debug.Context(ctx),
				debug.Config(cfg),
			)
			if err != nil {
				logger.Error().Err(err).Str("transport", "debug").Msg("Failed to initialize server")
				return err
			}

			gr.Add(debugServer.ListenAndServe, func(_ error) {
				_ = debugServer.Shutdown(ctx)
				cancel()
			})

			// start HTTP server
			httpServer, err := http.Server(
				http.Adapter(connector.NewHttpAdapter(gatewaySelector, cfg, st)),
				http.Logger(logger),
				http.Config(cfg),
				http.Context(ctx),
				http.TracerProvider(traceProvider),
				http.Store(st),
			)
			if err != nil {
				logger.Error().Err(err).Str("transport", "http").Msg("Failed to initialize server")
				return err
			}
			gr.Add(httpServer.Run, func(_ error) {
				cancel()
			})

			return gr.Run()
		},
	}
}
