package slot

import (
	"context"
	"fmt"
	"testing"

	"github.com/sterligov/banner-rotator/internal/mocks"
	"github.com/sterligov/banner-rotator/internal/model"
	"github.com/stretchr/testify/require"
)

func TestFindAllSlots(t *testing.T) {
	tests := []struct {
		slots []model.Slot
		err   error
		name  string
	}{
		{[]model.Slot{{ID: 1, Description: "descr"}, {ID: 2, Description: "descr2"}}, nil, "ok"},
		{[]model.Slot{}, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			slotGw := &mocks.SlotGateway{}

			ctx := context.Background()
			slotGw.
				On("FindAllSlots", ctx).
				Return(tst.slots, tst.err).
				Once()
			defer slotGw.AssertExpectations(t)

			uc := NewUseCase(slotGw)
			slots, err := uc.FindAllSlots(ctx)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.slots, slots)
		})
	}
}

func TestFindSlotByID(t *testing.T) {
	tests := []struct {
		slot model.Slot
		err  error
		name string
	}{
		{model.Slot{ID: 1, Description: "descr"}, nil, "ok"},
		{model.Slot{}, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			slotGw := &mocks.SlotGateway{}

			var slotID int64 = 1
			ctx := context.Background()
			slotGw.
				On("FindSlotByID", ctx, slotID).
				Return(tst.slot, tst.err).
				Once()
			defer slotGw.AssertExpectations(t)

			uc := NewUseCase(slotGw)
			slot, err := uc.FindSlotByID(ctx, slotID)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.slot, slot)
		})
	}
}

func TestCreateSlot(t *testing.T) {
	tests := []struct {
		insertedID int64
		err        error
		name       string
	}{
		{1, nil, "ok"},
		{0, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			slotGw := &mocks.SlotGateway{}

			slot := model.Slot{
				Description: "description",
			}

			ctx := context.Background()
			slotGw.
				On("CreateSlot", ctx, slot).
				Return(tst.insertedID, tst.err).
				Once()
			defer slotGw.AssertExpectations(t)

			uc := NewUseCase(slotGw)
			insertedID, err := uc.CreateSlot(ctx, slot)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.insertedID, insertedID)
		})
	}
}

func TestUpdateSlot(t *testing.T) {
	tests := []struct {
		affected int64
		err      error
		name     string
	}{
		{1, nil, "ok"},
		{0, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			slotGw := &mocks.SlotGateway{}

			slot := model.Slot{
				Description: "description",
			}

			ctx := context.Background()
			slotGw.
				On("UpdateSlot", ctx, slot).
				Return(tst.affected, tst.err).
				Once()
			defer slotGw.AssertExpectations(t)

			uc := NewUseCase(slotGw)
			insertedID, err := uc.UpdateSlot(ctx, slot)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.affected, insertedID)
		})
	}
}

func TestDeleteSlotByID(t *testing.T) {
	tests := []struct {
		affected int64
		err      error
		name     string
	}{
		{1, nil, "ok"},
		{0, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			slotGw := &mocks.SlotGateway{}

			var slotID int64 = 1

			ctx := context.Background()
			slotGw.
				On("DeleteSlotByID", ctx, slotID).
				Return(tst.affected, tst.err).
				Once()
			defer slotGw.AssertExpectations(t)

			uc := NewUseCase(slotGw)
			affected, err := uc.DeleteSlotByID(ctx, slotID)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.affected, affected)
		})
	}
}
