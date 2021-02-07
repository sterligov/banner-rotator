package config

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	HTTP struct {
		Addr           string        `yaml:"addr"`
		ReadTimeout    time.Duration `yaml:"read_timeout"`
		WriteTimeout   time.Duration `yaml:"write_timeout"`
		HandlerTimeout time.Duration `yaml:"handler_timeout"`
	} `yaml:"http"`

	GRPC struct {
		Addr string `yaml:"addr"`
	} `yaml:"grpc"`

	Database struct {
		Addr                string        `yaml:"connection_addr"`
		Driver              string        `yaml:"driver"`
		MaxReconnectRetries int           `yaml:"max_reconnect_retries"`
		ReconnectTime       time.Duration `yaml:"reconnect_time"`
	} `yaml:"database"`

	Logger struct {
		Path  string `yaml:"path"`
		Level string `yaml:"level"`
	} `yaml:"logger"`

	Queue struct {
		Subject string `yaml:"subject"`
	} `yaml:"queue"`

	Nats struct {
		URL                 string        `yaml:"url"`
		MaxReconnectRetries int           `yaml:"max_reconnect_retries"`
		ConnectTimeout      time.Duration `yaml:"connect_timeout"`
		ConnectTimeWait     time.Duration `yaml:"connect_time_wait"`
		ReconnectTime       time.Duration `yaml:"reconnect_time"`
	} `yaml:"nats"`
}

func New(cfgFilename string) (*Config, error) {
	f, err := os.Open(cfgFilename)
	if err != nil {
		return nil, fmt.Errorf("open config file failed: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			zap.L().Warn("config file close failed", zap.Error(err))
		}
	}()

	cfg := &Config{}

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(cfg); err != nil {
		return nil, fmt.Errorf("decode config file: %w", err)
	}

	return cfg, err
}
