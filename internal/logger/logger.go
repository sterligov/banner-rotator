package logger

import (
	"fmt"

	"github.com/sterligov/otus_highload/dating/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
}

var level = map[string]zapcore.Level{
	"info":    zapcore.InfoLevel,
	"error":   zapcore.ErrorLevel,
	"warning": zapcore.WarnLevel,
	"debug":   zapcore.DebugLevel,
}

func InitGlobal(cfg *config.Config) error {
	if _, ok := level[cfg.Logger.Level]; !ok {
		return fmt.Errorf("unexpected logger level %s", cfg.Logger.Level)
	}

	lcfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(level[cfg.Logger.Level]),
		OutputPaths: []string{cfg.Logger.Path},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
			TimeKey:     "time",
			EncodeTime:  zapcore.RFC3339TimeEncoder,
			NameKey:     "name",
			EncodeName:  zapcore.FullNameEncoder,
		},
	}

	zapLogger, err := lcfg.Build()
	if err != nil {
		return fmt.Errorf("build logger failed: %w", err)
	}

	zap.ReplaceGlobals(zapLogger)

	return nil
}
