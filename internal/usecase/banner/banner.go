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
		SelectBanner(stats []model.Statistic) int64
	}

	StatisticGateway interface {
		IncrementClicks(ctx context.Context, bannerID, slotID, groupID int64) error
		IncrementShows(ctx context.Context, bannerID, slotID, groupID int64) error
		FindStatistic(ctx context.Context, slotID, groupID int64) ([]model.Statistic, error)
	}

	BannerGateway interface { //nolint:golint
		FindBannerByID(ctx context.Context, id int64) (model.Banner, error)
		FindAllBanners(ctx context.Context) ([]model.Banner, error)
		FindAllBannersBySlotID(ctx context.Context, slotID int64) ([]model.Banner, error)
		CreateBannerSlotRelation(ctx context.Context, bannerID, slotID int64) (int64, error)
		DeleteBannerSlotRelation(ctx context.Context, bannerID, slotID int64) (int64, error)
		CreateBanner(ctx context.Context, banner model.Banner) (int64, error)
		DeleteBannerByID(ctx context.Context, id int64) (int64, error)
		UpdateBanner(ctx context.Context, banner model.Banner) (int64, error)
	}

	UseCase struct {
		bannerGateway    BannerGateway
		statisticGateway StatisticGateway
		eventGateway     EventGateway
		bandit           Bandit
	}
)

func NewUseCase(
	bannerGateway BannerGateway,
	eventGateway EventGateway,
	statisticGateway StatisticGateway,
	bandit Bandit,
) *UseCase {
	return &UseCase{
		bandit:           bandit,
		bannerGateway:    bannerGateway,
		eventGateway:     eventGateway,
		statisticGateway: statisticGateway,
	}
}

func (uc *UseCase) CreateBannerSlotRelation(ctx context.Context, bannerID, slotID int64) (int64, error) {
	return uc.bannerGateway.CreateBannerSlotRelation(ctx, bannerID, slotID)
}

func (uc *UseCase) DeleteBannerSlotRelation(ctx context.Context, bannerID, slotID int64) (int64, error) {
	return uc.bannerGateway.DeleteBannerSlotRelation(ctx, bannerID, slotID)
}

func (uc *UseCase) RegisterClick(ctx context.Context, bannerID, slotID, groupID int64) error {
	if err := uc.statisticGateway.IncrementClicks(ctx, bannerID, slotID, groupID); err != nil {
		return fmt.Errorf("register click: %w", err)
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
	stats, err := uc.statisticGateway.FindStatistic(ctx, slotID, groupID)
	if err != nil {
		return 0, fmt.Errorf("find statistic by slot: %w", err)
	}

	bannerID := uc.bandit.SelectBanner(stats)

	err = uc.statisticGateway.IncrementShows(ctx, bannerID, slotID, groupID)
	if err != nil {
		return 0, fmt.Errorf("increment shows: %w", err)
	}

	err = uc.eventGateway.Publish(model.Event{
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

func (uc *UseCase) FindAllBannersBySlotID(ctx context.Context, slotID int64) ([]model.Banner, error) {
	return uc.bannerGateway.FindAllBannersBySlotID(ctx, slotID)
}

func (uc *UseCase) CreateBanner(ctx context.Context, banner model.Banner) (int64, error) {
	return uc.bannerGateway.CreateBanner(ctx, banner)
}

func (uc *UseCase) UpdateBanner(ctx context.Context, banner model.Banner) (int64, error) {
	return uc.bannerGateway.UpdateBanner(ctx, banner)
}

func (uc *UseCase) DeleteBannerByID(ctx context.Context, id int64) (int64, error) {
	return uc.bannerGateway.DeleteBannerByID(ctx, id)
}

func (uc *UseCase) FindBannerByID(ctx context.Context, id int64) (model.Banner, error) {
	return uc.bannerGateway.FindBannerByID(ctx, id)
}

func (uc *UseCase) FindAllBanners(ctx context.Context) ([]model.Banner, error) {
	return uc.bannerGateway.FindAllBanners(ctx)
}
