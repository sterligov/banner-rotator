//nolint:dupl
package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/sterligov/banner-rotator/internal/mocks"
	"github.com/sterligov/banner-rotator/internal/model"
	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFindSlotByID(t *testing.T) {
	tests := []struct {
		slot model.Slot
		err  error
		code codes.Code
		name string
	}{
		{model.Slot{ID: 1, Description: "descr"}, nil, codes.OK, "ok"},
		{model.Slot{}, fmt.Errorf("error"), codes.Unknown, "unexpected error"},
		{model.Slot{}, model.ErrNotFound, codes.NotFound, "slot not found"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			slotUC := &mocks.SlotUseCase{}

			r := &pb.FindSlotByIDRequest{
				Id: 1,
			}
			ctx := context.Background()

			slotUC.
				On("FindSlotByID", ctx, r.Id).
				Return(tst.slot, tst.err).
				Once()
			defer slotUC.AssertExpectations(t)

			service := NewSlotService(slotUC)
			resp, err := service.FindSlotByID(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.code == codes.OK {
				require.Equal(t, tst.slot.ID, resp.Slot.Id)
				require.Equal(t, tst.slot.Description, resp.Slot.Description)
			}
		})
	}
}

func TestFindAllSlots(t *testing.T) {
	tests := []struct {
		slots []model.Slot
		err   error
		code  codes.Code
		name  string
	}{
		{
			[]model.Slot{{ID: 1, Description: "descr"}, {ID: 2, Description: "descr2"}},
			nil,
			codes.OK,
			"ok",
		},
		{
			[]model.Slot{},
			fmt.Errorf("error"),
			codes.Unknown,
			"unexpected error",
		},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			slotUC := &mocks.SlotUseCase{}

			r := &pb.FindAllSlotsRequest{}
			ctx := context.Background()

			slotUC.
				On("FindAllSlots", ctx).
				Return(tst.slots, tst.err).
				Once()
			defer slotUC.AssertExpectations(t)

			service := NewSlotService(slotUC)
			resp, err := service.FindAllSlots(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.code == codes.OK {
				require.Equal(t, len(tst.slots), len(resp.Slots))
				for i := 0; i < len(tst.slots); i++ {
					require.Equal(t, tst.slots[i].ID, resp.Slots[i].Id)
					require.Equal(t, tst.slots[i].Description, resp.Slots[i].Description)
				}
			}
		})
	}
}

func TestCreateSlot(t *testing.T) {
	tests := []struct {
		insertedID int64
		err        error
		code       codes.Code
		name       string
	}{
		{1, nil, codes.OK, "ok"},
		{0, fmt.Errorf("error"), codes.Unknown, "unexpected error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			slotUC := &mocks.SlotUseCase{}

			r := &pb.CreateSlotRequest{
				Slot: &pb.Slot{Description: "descr"},
			}
			ctx := context.Background()

			slotUC.
				On("CreateSlot", ctx, mock.MatchedBy(func(e model.Slot) bool {
					return e.Description == r.Slot.Description
				})).
				Return(tst.insertedID, tst.err).
				Once()
			defer slotUC.AssertExpectations(t)

			service := NewSlotService(slotUC)
			resp, err := service.CreateSlot(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.insertedID, resp.InsertedId)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}

func TestDeleteSlot(t *testing.T) {
	tests := []struct {
		affected int64
		err      error
		code     codes.Code
		name     string
	}{
		{1, nil, codes.OK, "ok"},
		{0, fmt.Errorf("error"), codes.Unknown, "unexpected error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			slotUC := &mocks.SlotUseCase{}

			r := &pb.DeleteSlotRequest{
				Id: 1,
			}
			ctx := context.Background()

			slotUC.
				On("DeleteSlotByID", ctx, r.Id).
				Return(tst.affected, tst.err).
				Once()
			defer slotUC.AssertExpectations(t)

			service := NewSlotService(slotUC)
			resp, err := service.DeleteSlot(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.affected, resp.Affected)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}

func TestUpdateSlot(t *testing.T) {
	tests := []struct {
		affected int64
		err      error
		code     codes.Code
		name     string
	}{
		{1, nil, codes.OK, "ok"},
		{0, fmt.Errorf("error"), codes.Unknown, "unexpected error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			slotUC := &mocks.SlotUseCase{}

			r := &pb.UpdateSlotRequest{
				Id: 1,
				Slot: &pb.Slot{
					Id:          1,
					Description: "descr",
				},
			}
			ctx := context.Background()

			slotUC.
				On("UpdateSlot", ctx, mock.MatchedBy(func(e model.Slot) bool {
					return e.Description == r.Slot.Description && e.ID == r.Slot.Id
				})).
				Return(tst.affected, tst.err).
				Once()
			defer slotUC.AssertExpectations(t)

			service := NewSlotService(slotUC)
			resp, err := service.UpdateSlot(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.affected, resp.Affected)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}
