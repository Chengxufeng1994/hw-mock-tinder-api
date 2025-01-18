package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/Chengxufeng1994/hw-mock-tinder-api/internal/infrastructure/config"
	"github.com/Chengxufeng1994/hw-mock-tinder-api/pkg/logging"
)

const (
	readHeaderTimeout = 10 * time.Second
	gracefulShutdown  = 5 * time.Second
)

type Server struct {
	logger  logging.Logger
	config  *config.Server
	handler http.Handler
}

func NewServer(logger logging.Logger, config *config.Server, handler http.Handler) *Server {
	return &Server{
		logger:  logger.WithName("Server"),
		config:  config,
		handler: handler,
	}
}

func (s *Server) ListenAndServe(ctx context.Context) error {
	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	srv := http.Server{
		Addr:              addr,
		Handler:           s.handler,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	var eg errgroup.Group
	eg.Go(func() error {
		s.logger.Infof("server started at %s", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.WithError(err).Error("server exited with error")
			return err
		}
		return nil
	})

	eg.Go(func() error {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdown)
		defer cancel()
		s.logger.Info("server shutting down")
		return srv.Shutdown(ctx)
	})

	return eg.Wait()
}
