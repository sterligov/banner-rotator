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

	"go.uber.org/zap"

	"github.com/sterligov/banner-rotator/internal/config"
	"github.com/sterligov/banner-rotator/internal/logger"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/calendar_config.yml", "Path to configuration file")
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

	go func() {
		if err := server.GRPC.Start(); err != nil {
			zap.L().Warn("grpc server start failed: %s", zap.Error(err))
		}
	}()

	go func() {
		if err := server.HTTP.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.L().Warn("http server start failed: %s", zap.Error(err))
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	<-signals
	signal.Stop(signals)

	if err := server.HTTP.Stop(context.Background()); err != nil && !errors.Is(err, http.ErrServerClosed) {
		zap.L().Warn("http server stop failed: %s", zap.Error(err))
	}

	server.GRPC.Stop()
}
