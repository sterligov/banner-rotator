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
		ID          int64  `db:"id"`
		Description string `db:"description"`
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
		return 0, fmt.Errorf("exec(CreateBannerSlotRelation): %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("last inserted id(CreateBannerSlotRelation): %w", err)
	}

	return insertedID, nil
}

func (bg *BannerGateway) DeleteBannerSlotRelation(ctx context.Context, bannerID, slotID int64) (int64, error) {
	const query = `DELETE FROM banner_slot WHERE banner_id = ? AND slot_id = ?`

	res, err := bg.db.ExecContext(ctx, query, bannerID, slotID)
	if err != nil {
		return 0, fmt.Errorf("exec(DeleteBannerSlotRelation): %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("affected(DeleteBannerSlotRelation): %w", err)
	}

	return affected, nil
}

func (bg *BannerGateway) FindBannerByID(ctx context.Context, id int64) (model.Banner, error) {
	const query = `SELECT * FROM banner WHERE id = ?`

	var b Banner
	err := bg.db.QueryRowxContext(ctx, query, id).StructScan(&b)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return toBanner(b), fmt.Errorf("banner: %w", model.ErrNotFound)
		}

		return toBanner(b), fmt.Errorf("query(FindByID banner): %w", err)
	}

	return toBanner(b), nil
}

func (bg *BannerGateway) FindAllBanners(ctx context.Context) ([]model.Banner, error) {
	const query = `SELECT * FROM banner`

	rows, err := bg.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query(FindAll banners): %w", err)
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
		if err := rows.StructScan(&banner); err != nil {
			return nil, fmt.Errorf("rows scan(FindAll banners): %w", err)
		}
		banners = append(banners, banner)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error(FindAll banners): %w", err)
	}

	return toBanners(banners), nil
}

func (bg *BannerGateway) FindAllBannersBySlotID(ctx context.Context, slotID int64) ([]model.Banner, error) {
	const query = `
SELECT b.*
FROM banner b
JOIN banner_slot bs ON bs.banner_id = b.id
WHERE bs.slot_id = ?`

	rows, err := bg.db.QueryxContext(ctx, query, slotID)
	if err != nil {
		return nil, fmt.Errorf("query(FindAllBannersBySlotID): %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			bg.logger.Warn("close rows failed", zap.Error(err))
		}
	}()

	var (
		banners []Banner
		banner  Banner
	)

	for rows.Next() {
		if err := rows.StructScan(&banner); err != nil {
			return nil, fmt.Errorf("cannot scan rows(FindAllBannersBySlotID): %w", err)
		}
		banners = append(banners, banner)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error(FindAllBannersBySlotID): %w", err)
	}

	return toBanners(banners), nil
}

func (bg *BannerGateway) CreateBanner(ctx context.Context, b model.Banner) (int64, error) {
	const query = `INSERT INTO banner(description) VALUES(?)`

	res, err := bg.db.ExecContext(ctx, query, b.Description)
	if err != nil {
		return 0, fmt.Errorf("exec(Create banner): %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("last insterted id(Create banner): %w", err)
	}

	return insertedID, nil
}

func (bg *BannerGateway) UpdateBanner(ctx context.Context, b model.Banner) (int64, error) {
	const query = `UPDATE banner SET description = ? WHERE id = ?`

	res, err := bg.db.ExecContext(ctx, query, b.Description, b.ID)
	if err != nil {
		return 0, fmt.Errorf("exec(Update banner): %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("affected(Update banner): %w", err)
	}

	return affected, nil
}

func (bg *BannerGateway) DeleteBannerByID(ctx context.Context, id int64) (int64, error) {
	const query = `DELETE FROM banner WHERE id = ?`

	res, err := bg.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, fmt.Errorf("exec(Delete banner): %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("affected(Delete banner): %w", err)
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
