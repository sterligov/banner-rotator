package sql

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sterligov/banner-rotator/internal/config"
)

func NewDatabase(cfg *config.Config) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, cfg.Database.Driver, cfg.Database.Addr)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	return db, nil
}
