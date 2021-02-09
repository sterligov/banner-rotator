//nolint:dupl
package service

import (
	"context"
	"errors"

	"github.com/sterligov/banner-rotator/internal/model"
	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	SlotUseCase interface {
		FindSlotByID(ctx context.Context, id int64) (model.Slot, error)
		FindAllSlots(ctx context.Context) ([]model.Slot, error)
		CreateSlot(ctx context.Context, b model.Slot) (int64, error)
		UpdateSlot(ctx context.Context, b model.Slot) (int64, error)
		DeleteSlotByID(ctx context.Context, id int64) (int64, error)
	}

	SlotService struct {
		pb.UnimplementedSlotServiceServer

		slotUC SlotUseCase
	}
)

func NewSlotService(slotUC SlotUseCase) *SlotService {
	return &SlotService{slotUC: slotUC}
}

func (bs *SlotService) FindSlotByID(ctx context.Context, r *pb.FindSlotByIDRequest) (*pb.FindSlotByIDResponse, error) {
	b, err := bs.slotUC.FindSlotByID(ctx, r.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "slot not found")
		}

		return nil, err
	}

	return &pb.FindSlotByIDResponse{Slot: toSlot(b)}, nil
}

func (bs *SlotService) FindAllSlots(ctx context.Context, _ *pb.FindAllSlotsRequest) (*pb.FindAllSlotsResponse, error) {
	slots, err := bs.slotUC.FindAllSlots(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindAllSlotsResponse{Slots: toSlots(slots)}, nil
}

func (bs *SlotService) CreateSlot(ctx context.Context, r *pb.CreateSlotRequest) (*pb.CreateSlotResponse, error) {
	insertedID, err := bs.slotUC.CreateSlot(ctx, model.Slot{
		Description: r.Slot.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateSlotResponse{InsertedId: insertedID}, nil
}

func (bs *SlotService) DeleteSlot(ctx context.Context, r *pb.DeleteSlotRequest) (*pb.DeleteSlotResponse, error) {
	affected, err := bs.slotUC.DeleteSlotByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteSlotResponse{Affected: affected}, nil
}

func (bs *SlotService) UpdateSlot(ctx context.Context, r *pb.UpdateSlotRequest) (*pb.UpdateSlotResponse, error) {
	affected, err := bs.slotUC.UpdateSlot(ctx, model.Slot{
		ID:          r.Slot.Id,
		Description: r.Slot.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateSlotResponse{Affected: affected}, nil
}

func toSlots(slots []model.Slot) []*pb.Slot {
	pbSlots := make([]*pb.Slot, len(slots))

	for i, b := range slots {
		pbSlots[i] = toSlot(b)
	}

	return pbSlots
}

func toSlot(b model.Slot) *pb.Slot {
	return &pb.Slot{
		Id:          b.ID,
		Description: b.Description,
	}
}
