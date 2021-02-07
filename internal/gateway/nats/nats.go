package nats

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sterligov/banner-rotator/internal/config"
	"go.uber.org/zap"
)

var ErrMaxReconnectRetries = fmt.Errorf("max reconnect retries")

type Nats struct {
	conn                *nats.Conn
	logger              *zap.Logger
	ctx                 context.Context
	connectTimeout      time.Duration
	connectTimeWait     time.Duration
	maxReconnectRetries int
}

func NewNatsConnection(cfg *config.Config) (*Nats, error) {
	logger := zap.L().Named("NATS")

	n := &Nats{
		ctx:                 context.Background(),
		connectTimeout:      cfg.Nats.ConnectTimeout,
		connectTimeWait:     cfg.Nats.ConnectTimeWait,
		maxReconnectRetries: cfg.Nats.MaxReconnectRetries,
		logger:              logger,
	}
	opts := []nats.Option{
		nats.SetCustomDialer(n),
		nats.ReconnectWait(cfg.Nats.ReconnectTime),
		nats.ReconnectHandler(func(c *nats.Conn) {
			logger.Info("reconnect", zap.String("url", c.ConnectedUrl()))
		}),
		nats.DisconnectErrHandler(func(c *nats.Conn, err error) {
			if err != nil {
				logger.Error("close is failed", zap.Error(err))
			}
		}),
		nats.ClosedHandler(func(c *nats.Conn) {
			logger.Info("connection is closed")
		}),
	}

	var err error
	n.conn, err = nats.Connect(cfg.Nats.URL, opts...)
	if err != nil {
		return nil, fmt.Errorf("nats connect: %w", err)
	}

	return n, nil
}

func (n *Nats) Dial(network, address string) (net.Conn, error) {
	var retries int
	for {
		n.logger.Info("attempting to connect", zap.String("address", address))

		retries++
		if retries > n.maxReconnectRetries {
			return nil, ErrMaxReconnectRetries
		}

		select {
		case <-n.ctx.Done():
			return nil, n.ctx.Err()
		default:
			d := &net.Dialer{}
			if conn, err := d.DialContext(n.ctx, network, address); err == nil {
				n.logger.Info("connecting successfully")
				return conn, nil
			}
			time.Sleep(n.connectTimeWait)
		}
	}
}
