package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sterligov/banner-rotator/internal/model"
	"go.uber.org/zap"
)

type (
	Banner struct {
		ID          int64
		Description string
	}

	BannerGateway struct {
		db     *sqlx.DB
		logger *zap.Logger
	}
)

func NewBannerGateway(db *sqlx.DB) *BannerGateway {
	return &BannerGateway{
		db:     db,
		logger: zap.L().Named("banner gateway"),
	}
}

func (bg *BannerGateway) CreateBannerSlotRelation(ctx context.Context, bannerID, slotID int64) (int64, error) {
	const query = `INSERT INTO banner_slot(banner_id, slot_id) VALUES(?, ?)`

	res, err := bg.db.ExecContext(ctx, query, bannerID, slotID)
	if err != nil {
		return 0, fmt.Errorf("create banner slot relation exec: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("create banner slot relation, last inserted id: %w", err)
	}

	return insertedID, nil
}

func (bg *BannerGateway) DeleteBannerSlotRelation(ctx context.Context, bannerID, slotID int64) (int64, error) {
	const query = `DELETE FROM banner_slot WHERE banner_id = ? AND slot_id = ?`

	res, err := bg.db.ExecContext(ctx, query, bannerID, slotID)
	if err != nil {
		return 0, fmt.Errorf("delete banner slot relation exec: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("delete banner slot relation, affected: %w", err)
	}

	return affected, nil
}

func (bg *BannerGateway) FindByID(ctx context.Context, id int64) (model.Banner, error) {
	const query = `SELECT * FROM banner WHERE id = ?`

	var b Banner
	err := bg.db.QueryRowxContext(ctx, query, id).StructScan(&b)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return toBanner(b), fmt.Errorf("banner: %w", model.ErrNotFound)
		}

		return toBanner(b), fmt.Errorf("select banner: %w", err)
	}

	return toBanner(b), nil
}

func (bg *BannerGateway) FindAll(ctx context.Context) ([]model.Banner, error) {
	const query = `SELECT * FROM banner`

	rows, err := bg.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("banner find all: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			bg.logger.Warn("banner close rows failed", zap.Error(err))
		}
	}()

	var (
		banners []Banner
		banner  Banner
	)

	for rows.Next() {
		if err := rows.Scan(&banner); err != nil {
			return nil, fmt.Errorf("find all banner rows scan: %w", err)
		}
		banners = append(banners, banner)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("find all banner rows: %w", err)
	}

	return toBanners(banners), nil
}

func (bg *BannerGateway) Create(ctx context.Context, b model.Banner) (int64, error) {
	const query = `INSERT INTO banner(description) VALUES(?)`

	res, err := bg.db.ExecContext(ctx, query, b.Description)
	if err != nil {
		return 0, fmt.Errorf("create banner exec: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("banner last insterted id: %w", err)
	}

	return insertedID, nil
}

func (bg *BannerGateway) Update(ctx context.Context, b model.Banner) (int64, error) {
	const query = `UPDATE banner SET description = ? WHERE id = ?`

	res, err := bg.db.ExecContext(ctx, query, b.Description, b.ID)
	if err != nil {
		return 0, fmt.Errorf("update banner exec: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("update banner affected: %w", err)
	}

	return affected, nil
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

func toBanners(banners []Banner) []model.Banner {
	mbanners := make([]model.Banner, len(banners))

	for i, b := range banners {
		mbanners[i] = toBanner(b)
	}

	return mbanners
}

func toBanner(b Banner) model.Banner {
	return model.Banner{
		ID:          b.ID,
		Description: b.Description,
	}
}
