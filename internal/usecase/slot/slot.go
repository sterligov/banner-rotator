package slot

import (
	"context"

	"github.com/sterligov/banner-rotator/internal/model"
)

type (
	SlotGateway interface {
		FindByID(ctx context.Context, id int64) (model.Slot, error)
		FindAll(ctx context.Context) ([]model.Slot, error)
		Create(ctx context.Context, slot model.Slot) (int64, error)
		DeleteByID(ctx context.Context, id int64) (int64, error)
		Update(ctx context.Context, slot model.Slot) (int64, error)
	}

	UseCase struct {
		slotGateway SlotGateway
	}
)

func NewUseCase(slotGateway SlotGateway) *UseCase {
	return &UseCase{slotGateway: slotGateway}
}

func (uc *UseCase) CreateSlot(ctx context.Context, slot model.Slot) (int64, error) {
	return uc.slotGateway.Create(ctx, slot)
}

func (uc *UseCase) UpdateSlot(ctx context.Context, slot model.Slot) (int64, error) {
	return uc.slotGateway.Create(ctx, slot)
}

func (uc *UseCase) DeleteSlotByID(ctx context.Context, id int64) (int64, error) {
	return uc.slotGateway.DeleteByID(ctx, id)
}

func (uc *UseCase) FindSlotByID(ctx context.Context, id int64) (model.Slot, error) {
	return uc.slotGateway.FindByID(ctx, id)
}

func (uc *UseCase) FindAllSlots(ctx context.Context) ([]model.Slot, error) {
	return uc.slotGateway.FindAll(ctx)
}
