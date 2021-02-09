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
	SlotGateway struct {
		db     *sqlx.DB
		logger *zap.Logger
	}

	Slot struct {
		ID          int64
		Description string
	}
)

func NewSlotGateway(db *sqlx.DB) *SlotGateway {
	return &SlotGateway{
		db:     db,
		logger: zap.L().Named("slot gateway"),
	}
}

func (gg *SlotGateway) FindSlotByID(ctx context.Context, id int64) (model.Slot, error) {
	const query = `SELECT * FROM slot WHERE id = ?`

	var g Slot
	err := gg.db.QueryRowxContext(ctx, query, id).StructScan(&g)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return toSlot(g), fmt.Errorf("social slot: %w", model.ErrNotFound)
		}

		return toSlot(g), fmt.Errorf("select social slot: %w", err)
	}

	return toSlot(g), nil
}

func (gg *SlotGateway) FindAllSlots(ctx context.Context) ([]model.Slot, error) {
	const query = `SELECT * FROM slot`

	rows, err := gg.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("social slot find all: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			gg.logger.Warn("social slot close rows failed", zap.Error(err))
		}
	}()

	var (
		slots []Slot
		slot  Slot
	)

	for rows.Next() {
		if err := rows.StructScan(&slot); err != nil {
			return nil, fmt.Errorf("find all social slot rows scan: %w", err)
		}
		slots = append(slots, slot)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("find all social slot rows: %w", err)
	}

	return toSlots(slots), nil
}

func (gg *SlotGateway) CreateSlot(ctx context.Context, s model.Slot) (int64, error) {
	const query = `INSERT INTO slot(description) VALUES(?)`

	res, err := gg.db.ExecContext(ctx, query, s.Description)
	if err != nil {
		return 0, fmt.Errorf("create social slot exec: %w", err)
	}

	insertedID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("social slot insterted id: %w", err)
	}

	return insertedID, nil
}

func (gg *SlotGateway) UpdateSlot(ctx context.Context, s model.Slot) (int64, error) {
	const query = `UPDATE slot SET description = ? WHERE id = ?`

	res, err := gg.db.ExecContext(ctx, query, s.Description, s.ID)
	if err != nil {
		return 0, fmt.Errorf("update slot exec: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("slot affected: %w", err)
	}

	return affected, nil
}

func (gg *SlotGateway) DeleteSlotByID(ctx context.Context, id int64) (int64, error) {
	const query = `DELETE FROM slot WHERE id = ?`

	res, err := gg.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, fmt.Errorf("delete slot exec: %w", err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("delete slot affected: %w", err)
	}

	return affected, nil
}

func toSlots(slots []Slot) []model.Slot {
	mslots := make([]model.Slot, len(slots))

	for i, g := range slots {
		mslots[i] = toSlot(g)
	}

	return mslots
}

func toSlot(sg Slot) model.Slot {
	return model.Slot{
		ID:          sg.ID,
		Description: sg.Description,
	}
}
