package sql

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sterligov/banner-rotator/internal/model"
)

type (
	GroupGateway struct {
		db *sqlx.DB
	}

	SocialGroup struct {
		ID          int64
		Description string
	}
)

func NewGroupGateway(db *sqlx.DB) *GroupGateway {
	return &GroupGateway{db: db}
}

func (sgg *GroupGateway) FindByID(ctx context.Context, id int64) (*model.Banner, error) {
	const query = `SELECT * FROM social_group WHERE id = ?`

	s := new(SocialGroup)
	err := gg.db.QueryRowxContext(ctx, query, id).StructScan(s)
	if err != nil {
		return nil, fmt.Errorf("delete social group exec: %w", err)
	}

	return toGroup(s), nil
}

func (sgg *GroupGateway) Create(ctx context.Context, sg *model.SocialGroup) (int64, error) {
	const query = `INSERT INTO social_group(description) VALUES(?)`

	res, err := gg.db.ExecContext(ctx, query, sg.Description)
	if err != nil {
		return 0, fmt.Errorf("create social group exec: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("social group insterted id: %w", err)
	}

	return insertedID, nil
}

func (sgg *GroupGateway) UpdateByID(ctx context.Context, sg SocialGroup) (int64, error) {
	const query = `UPDATE social_group SET description = ? WHERE id = ?`

	res, err := gg.db.ExecContext(ctx, query, sg.Description, sg.ID)
	if err != nil {
		return 0, fmt.Errorf("update social group exec: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("social group insterted id: %w", err)
	}

	return insertedID, nil
}

func (sgg *GroupGateway) DeleteByID(ctx context.Context, id int64) (int64, error) {
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

func toGroup(sg *Group) *model.Group {
	return &model.Group{
		ID:          sg.ID,
		Description: sg.Description,
	}
}
