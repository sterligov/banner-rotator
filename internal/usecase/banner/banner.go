package banner

import (
	"context"
	"fmt"
	"time"

	"github.com/sterligov/banner-rotator/internal/model"
)

type (
	EventGateway interface {
		Publish(e model.Event) error
	}

	Bandit interface {
		Pull(ctx context.Context) error
	}

	BannerGateway interface {
		FindByID(ctx context.Context, id int64) (model.Banner, error)
		FindAll(ctx context.Context) ([]model.Banner, error)
		CreateBannerSlotRelation(ctx context.Context, bannerID, slotID int64) error
		DeleteBannerSlotRelation(ctx context.Context, bannerID, slotID int64) error
		Create(ctx context.Context, banner model.Banner) (int64, error)
		DeleteByID(ctx context.Context, id int64) (int64, error)
		Update(ctx context.Context, banner model.Banner) (int64, error)
		IncrementShows(ctx context.Context, bannerID, slotID, groupID int64) error
	}

	UseCase struct {
		bannerGateway BannerGateway
		eventGateway  EventGateway
	}
)

func NewUseCase(
	bannerGateway BannerGateway,
	eventGateway EventGateway,
) *UseCase {
	return &UseCase{
		bannerGateway: bannerGateway,
		eventGateway:  eventGateway,
	}
}

func (uc *UseCase) CreateBannerSlotRelation(ctx context.Context, bannerID, slotID int64) error {
	return uc.bannerGateway.CreateBannerSlotRelation(ctx, bannerID, slotID)
}

func (uc *UseCase) DeleteBannerSlotRelation(ctx context.Context, bannerID, slotID int64) error {
	return uc.bannerGateway.DeleteBannerSlotRelation(ctx, bannerID, slotID)
}

func (uc *UseCase) RegisterClick(ctx context.Context, bannerID, slotID, groupID int64) error {
	if err := uc.bannerGateway.IncrementShows(ctx, bannerID, slotID, groupID); err != nil {
		return fmt.Errorf("register click gateway: %w", err)
	}

	err := uc.eventGateway.Publish(model.Event{
		Type:     model.EventClick,
		SlotID:   slotID,
		BannerID: bannerID,
		GroupID:  groupID,
		Date:     time.Now(),
	})
	if err != nil {
		return fmt.Errorf("notify click banner: %w", err)
	}

	return nil
}

func (uc *UseCase) SelectBanner(ctx context.Context, slotID, groupID int64) (int64, error) {
	var bannerID int64

	err := uc.eventGateway.Publish(model.Event{
		Type:     model.EventSelect,
		SlotID:   slotID,
		BannerID: bannerID,
		GroupID:  groupID,
		Date:     time.Now(),
	})
	if err != nil {
		return 0, fmt.Errorf("notify select banner: %w", err)
	}

	return bannerID, nil
}

func (uc *UseCase) CreateBanner(ctx context.Context, banner model.Banner) (int64, error) {
	return uc.bannerGateway.Create(ctx, banner)
}

func (uc *UseCase) UpdateBanner(ctx context.Context, banner model.Banner) (int64, error) {
	return uc.bannerGateway.Create(ctx, banner)
}

func (uc *UseCase) DeleteBannerByID(ctx context.Context, id int64) (int64, error) {
	return uc.bannerGateway.DeleteByID(ctx, id)
}

func (uc *UseCase) FindBannerByID(ctx context.Context, id int64) (model.Banner, error) {
	return uc.bannerGateway.FindByID(ctx, id)
}

func (uc *UseCase) FindAllBanners(ctx context.Context) ([]model.Banner, error) {
	return uc.bannerGateway.FindAll(ctx)
}
