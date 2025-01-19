package main

import (
	"context"
	"flag"

	"golang.org/x/sync/errgroup"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/server"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/ws"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/signal"
)

var cfgFile string

func main() {
	flag.StringVar(&cfgFile, "config", "", "config file (default is config.yaml)")
	flag.Parse()

	cfg, err := config.InitializationConfig(cfgFile)
	if err != nil {
		return
	}

	logger := logging.InitializationLogger(&cfg.Logging)
	defer logger.Flush()

	ctx := signal.WithContext(context.Background())

	application, cleanup, err := InitializeApplication(logger, cfg)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	var eg errgroup.Group
	eg.Go(func() error {
		return application.srv.ListenAndServe(ctx)
	})

	eg.Go(func() error {
		_ = application.ws.Run(ctx)
		return nil
	})

	if err := eg.Wait(); err != nil {
		logger.WithError(err).Error("application exited with error")
	}
}

type app struct {
	srv *server.Server
	ws  *ws.Hub
}

func newApp(srv *server.Server, ws *ws.Hub) app {
	return app{srv: srv, ws: ws}
}
