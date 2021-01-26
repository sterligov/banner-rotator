package banner

import (
	"context"
	"fmt"
	"time"

	"github.com/sterligov/banner-rotator/internal/model"

	"github.com/sterligov/banner-rotator/internal/gateway/sql"
)

type (
	Notifier interface {
		Notify(ctx context.Context, e *model.Event) error
	}

	Bandit interface {
		Pull(ctx context.Context) error
	}

	UseCase struct {
		bannerGateway sql.BannerGateway
		notifier      Notifier
	}
)

func NewUseCase(
	bannerGateway sql.BannerGateway,
	notifier Notifier,
) *UseCase {
	return &UseCase{
		bannerGateway: bannerGateway,
		notifier:      notifier,
	}
}

func (uc *UseCase) RegisterClick(ctx context.Context, bannerID, slotID, socialGroupID int64) error {
	err := uc.notifier.Notify(ctx, &model.Event{
		Type:          model.Click,
		SlotID:        slotID,
		BannerID:      bannerID,
		SocialGroupID: socialGroupID,
		Date:          time.Now(),
	})
	if err != nil {
		return fmt.Errorf("notify click banner: %w", err)
	}

	return nil
}

func (uc *UseCase) SelectBanner(ctx context.Context, slotID, socialGroupID int64) (int64, error) {
	var bannerID int64
	err := uc.notifier.Notify(ctx, &model.Event{
		Type:          model.Select,
		SlotID:        slotID,
		BannerID:      bannerID,
		SocialGroupID: socialGroupID,
		Date:          time.Now(),
	})
	if err != nil {
		return 0, fmt.Errorf("notify select banner: %w", err)
	}
	return 0, nil
}
