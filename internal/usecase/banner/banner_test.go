package banner

import (
	"context"
	"fmt"
	"testing"

	"github.com/sterligov/banner-rotator/internal/mocks"
	"github.com/sterligov/banner-rotator/internal/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterClick(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		eventGw := &mocks.EventGateway{}
		bannerGw := &mocks.BannerGateway{}

		var (
			slotID   int64 = 1
			groupID  int64 = 2
			bannerID int64 = 0
		)

		ctx := context.Background()
		bannerGw.
			On("RegisterClick", ctx, bannerID, slotID, groupID).
			Return(nil).
			Once()
		defer bannerGw.AssertExpectations(t)

		eventGw.On("Publish", mock.MatchedBy(func(e *model.Event) bool {
			return e.SlotID == slotID &&
				e.GroupID == groupID &&
				e.BannerID == bannerID &&
				e.Type == model.EventClick
		})).Return(nil).Once()
		defer eventGw.AssertExpectations(t)

		uc := NewUseCase(bannerGw, eventGw)
		err := uc.RegisterClick(ctx, bannerID, slotID, groupID)
		require.NoError(t, err)
	})

	t.Run("banner gateway error", func(t *testing.T) {
		eventGw := &mocks.EventGateway{}
		bannerGw := &mocks.BannerGateway{}

		var (
			slotID   int64 = 1
			groupID  int64 = 2
			bannerID int64 = 0
		)

		ctx := context.Background()
		bannerGw.
			On("RegisterClick", ctx, bannerID, slotID, groupID).
			Return(fmt.Errorf("error")).
			Once()
		defer bannerGw.AssertExpectations(t)

		uc := NewUseCase(bannerGw, eventGw)
		err := uc.RegisterClick(ctx, bannerID, slotID, groupID)
		require.Error(t, err)
	})

	t.Run("event gateway error", func(t *testing.T) {
		eventGw := &mocks.EventGateway{}
		bannerGw := &mocks.BannerGateway{}

		var (
			slotID   int64 = 1
			groupID  int64 = 2
			bannerID int64 = 0
		)

		ctx := context.Background()
		bannerGw.
			On("RegisterClick", ctx, bannerID, slotID, groupID).
			Return(nil).
			Once()
		defer bannerGw.AssertExpectations(t)

		eventGw.On("Publish", mock.MatchedBy(func(e *model.Event) bool {
			return e.SlotID == slotID &&
				e.GroupID == groupID &&
				e.BannerID == bannerID &&
				e.Type == model.EventClick
		})).Return(fmt.Errorf("error")).Once()
		defer eventGw.AssertExpectations(t)

		uc := NewUseCase(bannerGw, eventGw)
		err := uc.RegisterClick(ctx, bannerID, slotID, groupID)
		require.Error(t, err)
	})
}

func TestSelectBanner(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		eventGw := &mocks.EventGateway{}
		bannerGw := &mocks.BannerGateway{}

		var (
			slotID   int64 = 1
			groupID  int64 = 2
			bannerID int64 = 0
		)

		eventGw.On("Publish", mock.MatchedBy(func(e *model.Event) bool {
			return e.SlotID == slotID &&
				e.GroupID == groupID &&
				e.BannerID == bannerID &&
				e.Type == model.EventSelect
		})).Return(nil).Once()
		defer eventGw.AssertExpectations(t)

		uc := NewUseCase(bannerGw, eventGw)
		actualBannerID, err := uc.SelectBanner(context.Background(), slotID, groupID)
		require.NoError(t, err)
		require.Equal(t, bannerID, actualBannerID)
	})

	t.Run("event gateway error", func(t *testing.T) {
		eventGw := &mocks.EventGateway{}
		bannerGw := &mocks.BannerGateway{}

		var (
			slotID   int64 = 1
			groupID  int64 = 2
			bannerID int64 = 0
		)

		eventGw.On("Publish", mock.MatchedBy(func(e *model.Event) bool {
			return e.SlotID == slotID &&
				e.GroupID == groupID &&
				e.BannerID == bannerID &&
				e.Type == model.EventSelect
		})).Return(fmt.Errorf("error")).Once()
		defer eventGw.AssertExpectations(t)

		uc := NewUseCase(bannerGw, eventGw)
		actualBannerID, err := uc.SelectBanner(context.Background(), slotID, groupID)
		require.Error(t, err)
		require.Empty(t, actualBannerID)
	})
}
