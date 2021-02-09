//nolint:dupl
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
}

func TestRegisterClick_StatisticGatewayError(t *testing.T) {
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
}

func TestRegisterClick_EventGatewayError(t *testing.T) {
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
}

func TestSelectBanner(t *testing.T) {
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

	eventGw := &mocks.EventGateway{}
	bannerGw := &mocks.BannerGateway{}
	statisticGW := &mocks.StatisticGateway{}
	bandit := &mocks.Bandit{}

	var (
		slotID           int64 = 1
		groupID          int64 = 1
		selectedBannerID int64 = 1
	)

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
}

func TestSelectBanner_IncrementShowsError(t *testing.T) {
	stats := []model.Statistic{
		{
			BannerID: 1,
			GroupID:  1,
			SlotID:   1,
			Clicks:   10,
			Shows:    10,
		},
	}

	eventGw := &mocks.EventGateway{}
	bannerGw := &mocks.BannerGateway{}
	statisticGW := &mocks.StatisticGateway{}
	bandit := &mocks.Bandit{}

	var (
		slotID           int64 = 1
		groupID          int64 = 1
		selectedBannerID int64 = 1
	)

	statisticGW.
		On("FindStatistic", mock.Anything, slotID, groupID).
		Return(stats, nil).
		Once()

	statisticGW.
		On("IncrementShows", mock.Anything, selectedBannerID, slotID, groupID).
		Return(fmt.Errorf("error")).
		Once()
	defer statisticGW.AssertExpectations(t)

	bandit.On("SelectBanner", mock.MatchedBy(func(s []model.Statistic) bool {
		return assert.ElementsMatch(t, stats, s)
	})).Return(selectedBannerID).Once()
	defer bandit.AssertExpectations(t)

	uc := NewUseCase(bannerGw, eventGw, statisticGW, bandit)
	actualBannerID, err := uc.SelectBanner(context.Background(), slotID, groupID)
	require.Error(t, err)
	require.Empty(t, actualBannerID)
}

func TestSelectBanner_FindStatisticError(t *testing.T) {
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
}

func TestSelectBanner_PublishEventError(t *testing.T) {
	var stats []model.Statistic

	eventGw := &mocks.EventGateway{}
	bannerGw := &mocks.BannerGateway{}
	statisticGW := &mocks.StatisticGateway{}
	bandit := &mocks.Bandit{}

	var (
		slotID           int64 = 1
		groupID          int64 = 1
		selectedBannerID int64 = 1
	)

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
	})).Return(fmt.Errorf("error")).Once()
	defer eventGw.AssertExpectations(t)

	uc := NewUseCase(bannerGw, eventGw, statisticGW, bandit)
	actualBannerID, err := uc.SelectBanner(context.Background(), slotID, groupID)
	require.Error(t, err)
	require.Empty(t, actualBannerID)
}

func TestCreateBannerSlotRelation(t *testing.T) {
	tests := []struct {
		slotID     int64
		bannerID   int64
		insertedID int64
		err        error
		name       string
	}{
		{1, 1, 1, nil, "ok"},
		{1, 1, 0, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			bannerGw := &mocks.BannerGateway{}

			ctx := context.Background()
			bannerGw.
				On("CreateBannerSlotRelation", ctx, tst.bannerID, tst.slotID).
				Return(tst.insertedID, tst.err).
				Once()
			defer bannerGw.AssertExpectations(t)

			uc := NewUseCase(bannerGw, nil, nil, nil)
			insertedID, err := uc.CreateBannerSlotRelation(ctx, tst.bannerID, tst.slotID)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.insertedID, insertedID)
		})
	}
}

func TestDeleteBannerSlotRelation(t *testing.T) {
	tests := []struct {
		slotID   int64
		bannerID int64
		affected int64
		err      error
		name     string
	}{
		{1, 1, 1, nil, "ok"},
		{1, 1, 0, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			bannerGw := &mocks.BannerGateway{}

			ctx := context.Background()
			bannerGw.
				On("DeleteBannerSlotRelation", ctx, tst.bannerID, tst.slotID).
				Return(tst.affected, tst.err).
				Once()
			defer bannerGw.AssertExpectations(t)

			uc := NewUseCase(bannerGw, nil, nil, nil)
			affected, err := uc.DeleteBannerSlotRelation(ctx, tst.bannerID, tst.slotID)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.affected, affected)
		})
	}
}

func TestFindAllBannersBySlotID(t *testing.T) {
	tests := []struct {
		banners []model.Banner
		err     error
		name    string
	}{
		{[]model.Banner{{ID: 1, Description: "descr"}, {ID: 2, Description: "descr2"}}, nil, "ok"},
		{[]model.Banner{}, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			bannerGw := &mocks.BannerGateway{}

			var slotID int64 = 1
			ctx := context.Background()
			bannerGw.
				On("FindAllBannersBySlotID", ctx, slotID).
				Return(tst.banners, tst.err).
				Once()
			defer bannerGw.AssertExpectations(t)

			uc := NewUseCase(bannerGw, nil, nil, nil)
			banners, err := uc.FindAllBannersBySlotID(ctx, slotID)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.banners, banners)
		})
	}
}

func TestFindAllBanners(t *testing.T) {
	tests := []struct {
		banners []model.Banner
		err     error
		name    string
	}{
		{[]model.Banner{{ID: 1, Description: "descr"}, {ID: 2, Description: "descr2"}}, nil, "ok"},
		{[]model.Banner{}, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			bannerGw := &mocks.BannerGateway{}

			ctx := context.Background()
			bannerGw.
				On("FindAll", ctx).
				Return(tst.banners, tst.err).
				Once()
			defer bannerGw.AssertExpectations(t)

			uc := NewUseCase(bannerGw, nil, nil, nil)
			banners, err := uc.FindAllBanners(ctx)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.banners, banners)
		})
	}
}

func TestFindBannerByID(t *testing.T) {
	tests := []struct {
		banner model.Banner
		err    error
		name   string
	}{
		{model.Banner{ID: 1, Description: "descr"}, nil, "ok"},
		{model.Banner{}, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			bannerGw := &mocks.BannerGateway{}

			var bannerID int64 = 1
			ctx := context.Background()
			bannerGw.
				On("FindByID", ctx, bannerID).
				Return(tst.banner, tst.err).
				Once()
			defer bannerGw.AssertExpectations(t)

			uc := NewUseCase(bannerGw, nil, nil, nil)
			banner, err := uc.FindBannerByID(ctx, bannerID)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.banner, banner)
		})
	}
}

func TestCreateBanner(t *testing.T) {
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
			bannerGw := &mocks.BannerGateway{}

			banner := model.Banner{
				Description: "description",
			}

			ctx := context.Background()
			bannerGw.
				On("Create", ctx, banner).
				Return(tst.insertedID, tst.err).
				Once()
			defer bannerGw.AssertExpectations(t)

			uc := NewUseCase(bannerGw, nil, nil, nil)
			insertedID, err := uc.CreateBanner(ctx, banner)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.insertedID, insertedID)
		})
	}
}

func TestUpdateBanner(t *testing.T) {
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
			bannerGw := &mocks.BannerGateway{}

			banner := model.Banner{
				Description: "description",
			}

			ctx := context.Background()
			bannerGw.
				On("Update", ctx, banner).
				Return(tst.affected, tst.err).
				Once()
			defer bannerGw.AssertExpectations(t)

			uc := NewUseCase(bannerGw, nil, nil, nil)
			insertedID, err := uc.UpdateBanner(ctx, banner)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.affected, insertedID)
		})
	}
}

func TestDeleteBannerByID(t *testing.T) {
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
			bannerGw := &mocks.BannerGateway{}

			var bannerID int64 = 1

			ctx := context.Background()
			bannerGw.
				On("DeleteByID", ctx, bannerID).
				Return(tst.affected, tst.err).
				Once()
			defer bannerGw.AssertExpectations(t)

			uc := NewUseCase(bannerGw, nil, nil, nil)
			affected, err := uc.DeleteBannerByID(ctx, bannerID)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.affected, affected)
		})
	}
}
