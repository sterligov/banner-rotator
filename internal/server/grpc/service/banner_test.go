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

func TestRegisterClick(t *testing.T) {
	tests := []struct {
		err  error
		code codes.Code
		name string
	}{
		{nil, codes.OK, "ok"},
		{fmt.Errorf("error"), codes.Unknown, "unexpected error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			bannerUC := &mocks.BannerUseCase{}

			r := &pb.RegisterClickRequest{
				SlotId:   1,
				BannerId: 1,
				GroupId:  1,
			}
			ctx := context.Background()

			bannerUC.
				On("RegisterClick", ctx, r.BannerId, r.SlotId, r.GroupId).
				Return(tst.err).
				Once()
			defer bannerUC.AssertExpectations(t)

			service := NewBannerService(bannerUC)
			_, err := service.RegisterClick(ctx, r)

			require.Equal(t, tst.code, status.Code(err))
		})
	}
}

func TestSelectBanner(t *testing.T) {
	tests := []struct {
		bannerID int64
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
			bannerUC := &mocks.BannerUseCase{}

			r := &pb.SelectBannerRequest{
				SlotId:  1,
				GroupId: 1,
			}
			ctx := context.Background()

			bannerUC.
				On("SelectBanner", ctx, r.SlotId, r.GroupId).
				Return(tst.bannerID, tst.err).
				Once()
			defer bannerUC.AssertExpectations(t)

			service := NewBannerService(bannerUC)
			resp, err := service.SelectBanner(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.bannerID, resp.BannerId)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}

func TestCreateBannerSlotRelation(t *testing.T) {
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
			bannerUC := &mocks.BannerUseCase{}

			r := &pb.CreateBannerSlotRelationRequest{
				SlotId:   1,
				BannerId: 1,
			}
			ctx := context.Background()

			bannerUC.
				On("CreateBannerSlotRelation", ctx, r.BannerId, r.SlotId).
				Return(tst.insertedID, tst.err).
				Once()
			defer bannerUC.AssertExpectations(t)

			service := NewBannerService(bannerUC)
			resp, err := service.CreateBannerSlotRelation(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.insertedID, resp.InsertedId)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}

func TestDeleteBannerSlotRelation(t *testing.T) {
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
			bannerUC := &mocks.BannerUseCase{}

			r := &pb.DeleteBannerSlotRelationRequest{
				SlotId:   1,
				BannerId: 1,
			}
			ctx := context.Background()

			bannerUC.
				On("DeleteBannerSlotRelation", ctx, r.BannerId, r.SlotId).
				Return(tst.affected, tst.err).
				Once()
			defer bannerUC.AssertExpectations(t)

			service := NewBannerService(bannerUC)
			resp, err := service.DeleteBannerSlotRelation(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.affected, resp.Affected)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}

func TestFindBannerByID(t *testing.T) {
	tests := []struct {
		banner model.Banner
		err    error
		code   codes.Code
		name   string
	}{
		{model.Banner{ID: 1, Description: "descr"}, nil, codes.OK, "ok"},
		{model.Banner{}, fmt.Errorf("error"), codes.Unknown, "unexpected error"},
		{model.Banner{}, model.ErrNotFound, codes.NotFound, "banner not found"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			bannerUC := &mocks.BannerUseCase{}

			r := &pb.FindBannerByIDRequest{
				Id: 1,
			}
			ctx := context.Background()

			bannerUC.
				On("FindBannerByID", ctx, r.Id).
				Return(tst.banner, tst.err).
				Once()
			defer bannerUC.AssertExpectations(t)

			service := NewBannerService(bannerUC)
			resp, err := service.FindBannerByID(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.code == codes.OK {
				require.Equal(t, tst.banner.ID, resp.Banner.Id)
				require.Equal(t, tst.banner.Description, resp.Banner.Description)
			}
		})
	}
}

func TestFindAllBanners(t *testing.T) {
	tests := []struct {
		banners []model.Banner
		err     error
		code    codes.Code
		name    string
	}{
		{
			[]model.Banner{{ID: 1, Description: "descr"}, {ID: 2, Description: "descr2"}},
			nil,
			codes.OK,
			"ok",
		},
		{
			[]model.Banner{},
			fmt.Errorf("error"),
			codes.Unknown,
			"unexpected error",
		},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			bannerUC := &mocks.BannerUseCase{}

			r := &pb.FindAllBannersRequest{}
			ctx := context.Background()

			bannerUC.
				On("FindAllBanners", ctx).
				Return(tst.banners, tst.err).
				Once()
			defer bannerUC.AssertExpectations(t)

			service := NewBannerService(bannerUC)
			resp, err := service.FindAllBanners(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.code == codes.OK {
				require.Equal(t, len(tst.banners), len(resp.Banners))
				for i := 0; i < len(tst.banners); i++ {
					require.Equal(t, tst.banners[i].ID, resp.Banners[i].Id)
					require.Equal(t, tst.banners[i].Description, resp.Banners[i].Description)
				}
			}
		})
	}
}

func TestFindAllBannersBySlotID(t *testing.T) {
	tests := []struct {
		banners []model.Banner
		err     error
		code    codes.Code
		name    string
	}{
		{
			[]model.Banner{{ID: 1, Description: "descr"}, {ID: 2, Description: "descr2"}},
			nil,
			codes.OK,
			"ok",
		},
		{
			[]model.Banner{},
			fmt.Errorf("error"),
			codes.Unknown,
			"unexpected error",
		},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			bannerUC := &mocks.BannerUseCase{}

			r := &pb.FindAllBannersBySlotIDRequest{
				SlotId: 1,
			}
			ctx := context.Background()

			bannerUC.
				On("FindAllBannersBySlotID", ctx, r.SlotId).
				Return(tst.banners, tst.err).
				Once()
			defer bannerUC.AssertExpectations(t)

			service := NewBannerService(bannerUC)
			resp, err := service.FindAllBannersBySlotID(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.code == codes.OK {
				require.Equal(t, len(tst.banners), len(resp.Banners))
				for i := 0; i < len(tst.banners); i++ {
					require.Equal(t, tst.banners[i].ID, resp.Banners[i].Id)
					require.Equal(t, tst.banners[i].Description, resp.Banners[i].Description)
				}
			}
		})
	}
}

func TestCreateBanner(t *testing.T) {
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
			bannerUC := &mocks.BannerUseCase{}

			r := &pb.CreateBannerRequest{
				Banner: &pb.Banner{Description: "descr"},
			}
			ctx := context.Background()

			bannerUC.
				On("CreateBanner", ctx, mock.MatchedBy(func(e model.Banner) bool {
					return e.Description == r.Banner.Description
				})).
				Return(tst.insertedID, tst.err).
				Once()
			defer bannerUC.AssertExpectations(t)

			service := NewBannerService(bannerUC)
			resp, err := service.CreateBanner(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.insertedID, resp.InsertedId)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}

func TestDeleteBanner(t *testing.T) {
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
			bannerUC := &mocks.BannerUseCase{}

			r := &pb.DeleteBannerRequest{
				Id: 1,
			}
			ctx := context.Background()

			bannerUC.
				On("DeleteBannerByID", ctx, r.Id).
				Return(tst.affected, tst.err).
				Once()
			defer bannerUC.AssertExpectations(t)

			service := NewBannerService(bannerUC)
			resp, err := service.DeleteBanner(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.affected, resp.Affected)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}

func TestUpdateBanner(t *testing.T) {
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
			bannerUC := &mocks.BannerUseCase{}

			r := &pb.UpdateBannerRequest{
				Id: 1,
				Banner: &pb.Banner{
					Id:          1,
					Description: "descr",
				},
			}
			ctx := context.Background()

			bannerUC.
				On("UpdateBanner", ctx, mock.MatchedBy(func(e model.Banner) bool {
					return e.Description == r.Banner.Description && e.ID == r.Banner.Id
				})).
				Return(tst.affected, tst.err).
				Once()
			defer bannerUC.AssertExpectations(t)

			service := NewBannerService(bannerUC)
			resp, err := service.UpdateBanner(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.affected, resp.Affected)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}
