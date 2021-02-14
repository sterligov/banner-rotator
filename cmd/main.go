package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sterligov/banner-rotator/internal/config"
	"github.com/sterligov/banner-rotator/internal/logger"
	"go.uber.org/zap"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/rotator/config.yml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	var rerr error
	defer func() {
		if rerr != nil {
			log.Fatalln(rerr)
		}
	}()

	cfg, err := config.New(configFile)
	if err != nil {
		rerr = err
		return
	}

	if err := logger.InitGlobal(cfg); err != nil {
		rerr = err
		return
	}

	server, cleanup, err := setup(cfg)
	if err != nil {
		rerr = err
		return
	}
	defer cleanup()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	go func() {
		if err := server.GRPC.Start(); err != nil {
			zap.L().Warn("grpc server start failed", zap.Error(err))
		}
	}()

	go func() {
		if err := server.HTTP.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Warn("http server start failed", zap.Error(err))
		}
	}()

	<-signals
	signal.Stop(signals)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.HTTP.Stop(ctx); err != nil && !errors.Is(err, http.ErrServerClosed) {
		zap.L().Warn("http server stop failed", zap.Error(err))
	}

	server.GRPC.Stop()
}
