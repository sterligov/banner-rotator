package slot

import (
	"context"

	"github.com/sterligov/banner-rotator/internal/model"
)

type (
	SlotGateway interface { //nolint:golint
		FindSlotByID(ctx context.Context, id int64) (model.Slot, error)
		FindAllSlots(ctx context.Context) ([]model.Slot, error)
		CreateSlot(ctx context.Context, slot model.Slot) (int64, error)
		DeleteSlotByID(ctx context.Context, id int64) (int64, error)
		UpdateSlot(ctx context.Context, slot model.Slot) (int64, error)
	}

	UseCase struct {
		slotGateway SlotGateway
	}
)

func NewUseCase(slotGateway SlotGateway) *UseCase {
	return &UseCase{slotGateway: slotGateway}
}

func (uc *UseCase) CreateSlot(ctx context.Context, slot model.Slot) (int64, error) {
	return uc.slotGateway.CreateSlot(ctx, slot)
}

func (uc *UseCase) UpdateSlot(ctx context.Context, slot model.Slot) (int64, error) {
	return uc.slotGateway.UpdateSlot(ctx, slot)
}

func (uc *UseCase) DeleteSlotByID(ctx context.Context, id int64) (int64, error) {
	return uc.slotGateway.DeleteSlotByID(ctx, id)
}

func (uc *UseCase) FindSlotByID(ctx context.Context, id int64) (model.Slot, error) {
	return uc.slotGateway.FindSlotByID(ctx, id)
}

func (uc *UseCase) FindAllSlots(ctx context.Context) ([]model.Slot, error) {
	return uc.slotGateway.FindAllSlots(ctx)
}
