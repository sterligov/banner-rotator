package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sterligov/banner-rotator/internal/model"
	"go.uber.org/zap"
)

type (
	Statistic struct {
		BannerID int64         `db:"banner_id"`
		SlotID   int64         `db:"slot_id"`
		GroupID  sql.NullInt64 `db:"social_group_id"`
		Clicks   sql.NullInt64 `db:"clicks"`
		Shows    sql.NullInt64 `db:"shows"`
	}

	StatisticGateway struct {
		db     *sqlx.DB
		logger *zap.Logger
	}
)

func NewStatisticGateway(db *sqlx.DB) *StatisticGateway {
	return &StatisticGateway{
		db:     db,
		logger: zap.L().Named("statistic gateway"),
	}
}

func (sg *StatisticGateway) IncrementShows(ctx context.Context, bannerID, slotID, groupID int64) error {
	query := `
INSERT IGNORE INTO statistic(
	banner_slot_id,
    social_group_id,
    clicks,
    shows
) VALUES(
    (
	    SELECT id
		FROM banner_slot
		WHERE banner_id = ? AND slot_id = ?
	), ?, 0, 0
)`

	_, err := sg.db.ExecContext(ctx, query, bannerID, slotID, groupID)
	if err != nil {
		return fmt.Errorf("insert into statistic exec: %w", err)
	}

	query = `
UPDATE statistic s
JOIN banner_slot bs ON s.banner_slot_id = bs.id
SET s.shows = s.shows + 1
WHERE bs.banner_id = ? AND bs.slot_id = ? AND s.social_group_id = ?`

	_, err = sg.db.ExecContext(ctx, query, bannerID, slotID, groupID)
	if err != nil {
		return fmt.Errorf("update shows exec: %w", err)
	}

	return nil
}

func (sg *StatisticGateway) IncrementClicks(ctx context.Context, bannerID, slotID, groupID int64) error {
	var query = `
UPDATE statistic s
JOIN banner_slot bs ON s.banner_slot_id = bs.id
SET s.clicks = s.clicks + 1
WHERE bs.banner_id = ? AND bs.slot_id = ? AND s.social_group_id = ?`

	_, err := sg.db.ExecContext(ctx, query, bannerID, slotID, groupID)
	if err != nil {
		return fmt.Errorf("increment clicks exec: %w", err)
	}

	return nil
}

func (sg *StatisticGateway) FindStatistic(ctx context.Context, slotID, groupID int64) ([]model.Statistic, error) {
	var query = `
SELECT bs.banner_id,
       bs.slot_id,
       s.social_group_id,
       s.clicks,
       s.shows
FROM statistic s
RIGHT JOIN banner_slot bs on bs.id = s.banner_slot_id
WHERE bs.slot_id = ? AND (s.social_group_id = ? OR s.social_group_id IS NULL)`

	rows, err := sg.db.QueryxContext(ctx, query, slotID, groupID)
	if err != nil {
		return nil, fmt.Errorf("fetch query statistic: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			sg.logger.Warn("fetch statistic close rows failed", zap.Error(err))
		}
	}()

	var (
		stats []Statistic
		s     Statistic
	)

	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			return nil, fmt.Errorf("fetch statistic rows scan: %w", err)
		}
		stats = append(stats, s)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("fetch statistic rows: %w", err)
	}

	return toStats(stats), nil
}

func toStats(stats []Statistic) []model.Statistic {
	mstats := make([]model.Statistic, len(stats))

	for i, s := range stats {
		mstats[i] = toStat(s)
	}

	return mstats
}

func toStat(s Statistic) model.Statistic {
	return model.Statistic{
		BannerID: s.BannerID,
		SlotID:   s.SlotID,
		GroupID:  s.GroupID.Int64,
		Clicks:   s.Clicks.Int64,
		Shows:    s.Shows.Int64,
	}
}
