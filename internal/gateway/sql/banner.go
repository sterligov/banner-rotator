package sql

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sterligov/banner-rotator/internal/model"
)

type (
	BannerGateway struct {
		db *sqlx.DB
	}

	Banner struct {
		ID          int64
		Description string
	}
)

func NewGateway(db *sqlx.DB) *BannerGateway {
	return &BannerGateway{db: db}
}

func (bg *BannerGateway) FindByID(ctx context.Context, id int64) (*model.Banner, error) {
	const query = `SELECT * FROM banner WHERE id = ?`

	b := new(Banner)
	err := bg.db.QueryRowxContext(ctx, query, id).StructScan(b)
	if err != nil {
		return nil, fmt.Errorf("delete banner exec: %w", err)
	}

	return toBanner(b), nil
}

func (bg *BannerGateway) Create(ctx context.Context, b model.Banner) (int64, error) {
	const query = `INSERT INTO banner(description) VALUES(?)`

	res, err := bg.db.ExecContext(ctx, query, b.Description)
	if err != nil {
		return 0, fmt.Errorf("create banner exec: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("banner insterted id: %w", err)
	}

	return insertedID, nil
}

func (bg *BannerGateway) UpdateByID(ctx context.Context, b Banner) (int64, error) {
	const query = `UPDATE banner SET description = ? WHERE id = ?`

	res, err := bg.db.ExecContext(ctx, query, banner.Description, b.ID)
	if err != nil {
		return 0, fmt.Errorf("update banner exec: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("banner insterted id: %w", err)
	}

	return insertedID, nil
}

func (bg *BannerGateway) DeleteByID(ctx context.Context, id int64) (int64, error) {
	const query = `DELETE FROM banner WHERE id = ?`

	res, err := bg.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, fmt.Errorf("delete banner exec: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("delete banner affected: %w", err)
	}

	return affected, nil
}

func toBanner(b *Banner) *model.Banner {
	return &model.Banner{
		ID:          b.ID,
		Description: b.Description,
	}
}
