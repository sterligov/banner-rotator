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
	BannerUseCase interface {
		RegisterClick(ctx context.Context, bannerID, slotID, groupID int64) error
		CreateBannerSlotRelation(ctx context.Context, bannerID, slotID int64) (int64, error)
		DeleteBannerSlotRelation(ctx context.Context, bannerID, slotID int64) (int64, error)
		SelectBanner(ctx context.Context, slotID, groupID int64) (int64, error)
		FindBannerByID(ctx context.Context, id int64) (model.Banner, error)
		FindAllBanners(ctx context.Context) ([]model.Banner, error)
		FindAllBannersBySlotID(ctx context.Context, slotID int64) ([]model.Banner, error)
		CreateBanner(ctx context.Context, b model.Banner) (int64, error)
		UpdateBanner(ctx context.Context, b model.Banner) (int64, error)
		DeleteBannerByID(ctx context.Context, id int64) (int64, error)
	}

	BannerService struct {
		pb.UnimplementedBannerServiceServer

		bannerUC BannerUseCase
	}
)

func NewBannerService(bannerUC BannerUseCase) *BannerService {
	return &BannerService{bannerUC: bannerUC}
}

func (bs *BannerService) RegisterClick(ctx context.Context, r *pb.RegisterClickRequest) (*pb.RegisterClickResponse, error) {
	err := bs.bannerUC.RegisterClick(ctx, r.BannerId, r.SlotId, r.GroupId)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterClickResponse{}, nil
}

func (bs *BannerService) SelectBanner(ctx context.Context, r *pb.SelectBannerRequest) (*pb.SelectBannerResponse, error) {
	bannerID, err := bs.bannerUC.SelectBanner(ctx, r.SlotId, r.GroupId)
	if err != nil {
		return nil, err
	}

	return &pb.SelectBannerResponse{BannerId: bannerID}, nil
}

func (bs *BannerService) CreateBannerSlotRelation(
	ctx context.Context,
	r *pb.CreateBannerSlotRelationRequest,
) (*pb.CreateBannerSlotRelationResponse, error) {
	insertedID, err := bs.bannerUC.CreateBannerSlotRelation(ctx, r.BannerId, r.SlotId)

	return &pb.CreateBannerSlotRelationResponse{InsertedId: insertedID}, err
}

func (bs *BannerService) DeleteBannerSlotRelation(
	ctx context.Context,
	r *pb.DeleteBannerSlotRelationRequest,
) (*pb.DeleteBannerSlotRelationResponse, error) {
	affected, err := bs.bannerUC.DeleteBannerSlotRelation(ctx, r.BannerId, r.SlotId)

	return &pb.DeleteBannerSlotRelationResponse{Affected: affected}, err
}

func (bs *BannerService) FindBannerByID(ctx context.Context, r *pb.FindBannerByIDRequest) (*pb.FindBannerByIDResponse, error) {
	b, err := bs.bannerUC.FindBannerByID(ctx, r.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "banner not found")
		}

		return nil, err
	}

	return &pb.FindBannerByIDResponse{Banner: toBanner(b)}, nil
}

func (bs *BannerService) FindAllBanners(ctx context.Context, _ *pb.FindAllBannersRequest) (*pb.FindAllBannersResponse, error) {
	banners, err := bs.bannerUC.FindAllBanners(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindAllBannersResponse{Banners: toBanners(banners)}, nil
}

func (bs *BannerService) FindAllBannersBySlotID(
	ctx context.Context,
	r *pb.FindAllBannersBySlotIDRequest,
) (*pb.FindAllBannersBySlotIDResponse, error) {
	banners, err := bs.bannerUC.FindAllBannersBySlotID(ctx, r.SlotId)
	if err != nil {
		return nil, err
	}

	return &pb.FindAllBannersBySlotIDResponse{Banners: toBanners(banners)}, nil
}

func (bs *BannerService) CreateBanner(ctx context.Context, r *pb.CreateBannerRequest) (*pb.CreateBannerResponse, error) {
	insertedID, err := bs.bannerUC.CreateBanner(ctx, model.Banner{
		Description: r.Banner.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateBannerResponse{InsertedId: insertedID}, nil
}

func (bs *BannerService) DeleteBanner(ctx context.Context, r *pb.DeleteBannerRequest) (*pb.DeleteBannerResponse, error) {
	affected, err := bs.bannerUC.DeleteBannerByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteBannerResponse{Affected: affected}, nil
}

func (bs *BannerService) UpdateBanner(ctx context.Context, r *pb.UpdateBannerRequest) (*pb.UpdateBannerResponse, error) {
	affected, err := bs.bannerUC.UpdateBanner(ctx, model.Banner{
		ID:          r.Banner.Id,
		Description: r.Banner.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateBannerResponse{Affected: affected}, nil
}

func toBanners(banners []model.Banner) []*pb.Banner {
	pbBanners := make([]*pb.Banner, len(banners))

	for i, b := range banners {
		pbBanners[i] = toBanner(b)
	}

	return pbBanners
}

func toBanner(b model.Banner) *pb.Banner {
	return &pb.Banner{
		Id:          b.ID,
		Description: b.Description,
	}
}
