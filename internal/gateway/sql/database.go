package sql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" //nolint
	"github.com/jmoiron/sqlx"
	"github.com/sterligov/banner-rotator/internal/config"
	"go.uber.org/zap"
)

var ErrDatabaseConnection = fmt.Errorf("database connection failed")

func NewDatabase(cfg *config.Config) (*sqlx.DB, func(), error) {
	for i := 0; i < cfg.Database.MaxReconnectRetries; i++ {
		db, err := sqlx.Connect(cfg.Database.Driver, cfg.Database.Addr)
		if err != nil {
			zap.L().Warn("reconnect to database", zap.String("address", cfg.Database.Addr))
			time.Sleep(cfg.Database.ReconnectTime)
			continue
		}

		db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
		db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
		db.SetConnMaxLifetime(cfg.Database.MaxConnLifetime)

		return db, func() {
			if err := db.Close(); err != nil {
				zap.L().Warn("close nats connection failed")
			}
		}, nil
	}

	return nil, nil, ErrDatabaseConnection
}
