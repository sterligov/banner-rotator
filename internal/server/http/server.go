package internalhttp

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/sterligov/banner-rotator/internal/config"
)

type Server struct {
	httpServer http.Server
}

func NewServer(cfg *config.Config, h http.Handler) (*Server, error) {
	server := &Server{
		httpServer: http.Server{
			Addr:         cfg.HTTP.Addr,
			ReadTimeout:  cfg.HTTP.ReadTimeout,
			WriteTimeout: cfg.HTTP.WriteTimeout,
			Handler:      http.TimeoutHandler(h, cfg.HTTP.HandlerTimeout, "request timeout"),
		},
	}

	return server, nil
}

func (s *Server) Start() error {
	zap.L().Info("Start http server...")

	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	zap.L().Info("Stop http server...")

	return s.httpServer.Shutdown(ctx)
}
