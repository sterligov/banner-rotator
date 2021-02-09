//nolint:dupl
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
	GroupGateway struct {
		db     *sqlx.DB
		logger *zap.Logger
	}

	Group struct {
		ID          int64
		Description string
	}
)

func NewGroupGateway(db *sqlx.DB) *GroupGateway {
	return &GroupGateway{
		db:     db,
		logger: zap.L().Named("group gateway"),
	}
}

func (gg *GroupGateway) FindByID(ctx context.Context, id int64) (model.Group, error) {
	const query = `SELECT * FROM social_group WHERE id = ?`

	var g Group
	err := gg.db.QueryRowxContext(ctx, query, id).StructScan(&g)
	if err != nil {
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return toGroup(g), fmt.Errorf("social group: %w", model.ErrNotFound)
			}

			return toGroup(g), fmt.Errorf("select social group: %w", err)
		}
	}

	return toGroup(g), nil
}

func (gg *GroupGateway) FindAll(ctx context.Context) ([]model.Group, error) {
	const query = `SELECT * FROM social_group`

	rows, err := gg.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("social group find all: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			gg.logger.Warn("social group close rows failed", zap.Error(err))
		}
	}()

	var (
		groups []Group
		group  Group
	)

	for rows.Next() {
		if err := rows.StructScan(&group); err != nil {
			return nil, fmt.Errorf("find all social group rows scan: %w", err)
		}
		groups = append(groups, group)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("find all social group rows: %w", err)
	}

	return toGroups(groups), nil
}

func (gg *GroupGateway) Create(ctx context.Context, g model.Group) (int64, error) {
	const query = `INSERT INTO social_group(description) VALUES(?)`

	res, err := gg.db.ExecContext(ctx, query, g.Description)
	if err != nil {
		return 0, fmt.Errorf("create social group exec: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("social group insterted id: %w", err)
	}

	return insertedID, nil
}

func (gg *GroupGateway) Update(ctx context.Context, g model.Group) (int64, error) {
	const query = `UPDATE social_group SET description = ? WHERE id = ?`

	res, err := gg.db.ExecContext(ctx, query, g.Description, g.ID)
	if err != nil {
		return 0, fmt.Errorf("update social group exec: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("social group affected: %w", err)
	}

	return affected, nil
}

func (gg *GroupGateway) DeleteByID(ctx context.Context, id int64) (int64, error) {
	const query = `DELETE FROM social_group WHERE id = ?`

	res, err := gg.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, fmt.Errorf("delete social group exec: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("delete social group affected: %w", err)
	}

	return affected, nil
}

func toGroups(groups []Group) []model.Group {
	mgroups := make([]model.Group, len(groups))

	for i, g := range groups {
		mgroups[i] = toGroup(g)
	}

	return mgroups
}

func toGroup(sg Group) model.Group {
	return model.Group{
		ID:          sg.ID,
		Description: sg.Description,
	}
}
