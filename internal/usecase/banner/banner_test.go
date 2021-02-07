package banner

import (
	"context"
	"fmt"
	"testing"

	"github.com/sterligov/banner-rotator/internal/mocks"
	"github.com/sterligov/banner-rotator/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterClick(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		eventGw := &mocks.EventGateway{}
		bannerGw := &mocks.BannerGateway{}
		statisticGW := &mocks.StatisticGateway{}
		bandit := &mocks.Bandit{}

		var (
			slotID   int64 = 1
			groupID  int64 = 2
			bannerID int64 = 0
		)

		ctx := context.Background()
		statisticGW.
			On("IncrementClicks", ctx, bannerID, slotID, groupID).
			Return(nil).
			Once()
		defer statisticGW.AssertExpectations(t)

		eventGw.On("Publish", mock.MatchedBy(func(e model.Event) bool {
			return e.SlotID == slotID &&
				e.GroupID == groupID &&
				e.BannerID == bannerID &&
				e.Type == model.EventClick
		})).Return(nil).Once()
		defer eventGw.AssertExpectations(t)

		uc := NewUseCase(bannerGw, eventGw, statisticGW, bandit)
		err := uc.RegisterClick(ctx, bannerID, slotID, groupID)
		require.NoError(t, err)
	})

	t.Run("statistic gateway error", func(t *testing.T) {
		eventGw := &mocks.EventGateway{}
		bannerGw := &mocks.BannerGateway{}
		statisticGW := &mocks.StatisticGateway{}
		bandit := &mocks.Bandit{}

		var (
			slotID   int64 = 1
			groupID  int64 = 2
			bannerID int64 = 0
		)

		ctx := context.Background()
		statisticGW.
			On("IncrementClicks", ctx, bannerID, slotID, groupID).
			Return(fmt.Errorf("error")).
			Once()
		defer statisticGW.AssertExpectations(t)

		uc := NewUseCase(bannerGw, eventGw, statisticGW, bandit)
		err := uc.RegisterClick(ctx, bannerID, slotID, groupID)
		require.Error(t, err)
	})

	t.Run("event gateway error", func(t *testing.T) {
		eventGw := &mocks.EventGateway{}
		bannerGw := &mocks.BannerGateway{}
		statisticGW := &mocks.StatisticGateway{}
		bandit := &mocks.Bandit{}

		var (
			slotID   int64 = 1
			groupID  int64 = 2
			bannerID int64 = 0
		)

		ctx := context.Background()
		statisticGW.
			On("IncrementClicks", ctx, bannerID, slotID, groupID).
			Return(nil).
			Once()
		defer statisticGW.AssertExpectations(t)

		eventGw.On("Publish", mock.MatchedBy(func(e model.Event) bool {
			return e.SlotID == slotID &&
				e.GroupID == groupID &&
				e.BannerID == bannerID &&
				e.Type == model.EventClick
		})).Return(fmt.Errorf("error")).Once()
		defer eventGw.AssertExpectations(t)

		uc := NewUseCase(bannerGw, eventGw, statisticGW, bandit)
		err := uc.RegisterClick(ctx, bannerID, slotID, groupID)
		require.Error(t, err)
	})
}

func TestSelectBanner(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		eventGw := &mocks.EventGateway{}
		bannerGw := &mocks.BannerGateway{}
		statisticGW := &mocks.StatisticGateway{}
		bandit := &mocks.Bandit{}

		var (
			slotID           int64 = 1
			groupID          int64 = 1
			selectedBannerID int64 = 1
		)

		stats := []model.Statistic{
			{
				BannerID: 1,
				GroupID:  1,
				SlotID:   1,
				Clicks:   10,
				Shows:    10,
			},
			{
				BannerID: 2,
				GroupID:  1,
				SlotID:   1,
				Clicks:   20,
				Shows:    20,
			},
		}

		statisticGW.
			On("FindStatistic", mock.Anything, slotID, groupID).
			Return(stats, nil).
			Once()
		defer statisticGW.AssertExpectations(t)

		statisticGW.
			On("IncrementShows", mock.Anything, selectedBannerID, slotID, groupID).
			Return(nil).
			Once()

		bandit.On("SelectBanner", mock.MatchedBy(func(s []model.Statistic) bool {
			return assert.ElementsMatch(t, stats, s)
		})).Return(selectedBannerID).Once()
		defer bandit.AssertExpectations(t)

		eventGw.On("Publish", mock.MatchedBy(func(e model.Event) bool {
			return e.SlotID == slotID &&
				e.GroupID == groupID &&
				e.BannerID == selectedBannerID &&
				e.Type == model.EventSelect
		})).Return(nil).Once()
		defer eventGw.AssertExpectations(t)

		uc := NewUseCase(bannerGw, eventGw, statisticGW, bandit)
		actualBannerID, err := uc.SelectBanner(context.Background(), slotID, groupID)
		require.NoError(t, err)
		require.Equal(t, selectedBannerID, actualBannerID)
	})

	t.Run("find statistic error", func(t *testing.T) {
		eventGw := &mocks.EventGateway{}
		bannerGw := &mocks.BannerGateway{}
		statisticGW := &mocks.StatisticGateway{}
		bandit := &mocks.Bandit{}

		var (
			slotID  int64 = 1
			groupID int64 = 2
		)

		statisticGW.
			On("FindStatistic", mock.Anything, slotID, groupID).
			Return(nil, fmt.Errorf("error")).
			Once()
		defer statisticGW.AssertExpectations(t)

		uc := NewUseCase(bannerGw, eventGw, statisticGW, bandit)
		actualBannerID, err := uc.SelectBanner(context.Background(), slotID, groupID)
		require.Error(t, err)
		require.Empty(t, actualBannerID)
	})
}
