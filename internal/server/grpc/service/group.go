package service //nolint:dupl

import (
	"context"
	"errors"

	"github.com/sterligov/banner-rotator/internal/model"
	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	GroupUseCase interface {
		FindGroupByID(ctx context.Context, id int64) (model.Group, error)
		FindAllGroups(ctx context.Context) ([]model.Group, error)
		CreateGroup(ctx context.Context, b model.Group) (int64, error)
		UpdateGroup(ctx context.Context, b model.Group) (int64, error)
		DeleteGroupByID(ctx context.Context, id int64) (int64, error)
	}

	GroupService struct {
		pb.UnimplementedGroupServiceServer

		groupUC GroupUseCase
	}
)

func NewGroupService(groupUC GroupUseCase) *GroupService {
	return &GroupService{groupUC: groupUC}
}

func (bs *GroupService) FindGroupByID(ctx context.Context, r *pb.FindGroupByIDRequest) (*pb.FindGroupByIDResponse, error) {
	b, err := bs.groupUC.FindGroupByID(ctx, r.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "group not found")
		}

		return nil, err
	}

	return &pb.FindGroupByIDResponse{Group: toGroup(b)}, nil
}

func (bs *GroupService) FindAllGroups(ctx context.Context, _ *pb.FindAllGroupsRequest) (*pb.FindAllGroupsResponse, error) {
	groups, err := bs.groupUC.FindAllGroups(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.FindAllGroupsResponse{Groups: toGroups(groups)}, nil
}

func (bs *GroupService) CreateGroup(ctx context.Context, r *pb.CreateGroupRequest) (*pb.CreateGroupResponse, error) {
	insertedID, err := bs.groupUC.CreateGroup(ctx, model.Group{
		Description: r.Group.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateGroupResponse{InsertedId: insertedID}, nil
}

func (bs *GroupService) DeleteGroup(ctx context.Context, r *pb.DeleteGroupRequest) (*pb.DeleteGroupResponse, error) {
	affected, err := bs.groupUC.DeleteGroupByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteGroupResponse{Affected: affected}, nil
}

func (bs *GroupService) UpdateGroup(ctx context.Context, r *pb.UpdateGroupRequest) (*pb.UpdateGroupResponse, error) {
	affected, err := bs.groupUC.UpdateGroup(ctx, model.Group{
		ID:          r.Group.Id,
		Description: r.Group.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateGroupResponse{Affected: affected}, nil
}

func toGroups(groups []model.Group) []*pb.Group {
	pbGroups := make([]*pb.Group, len(groups))

	for i, b := range groups {
		pbGroups[i] = toGroup(b)
	}

	return pbGroups
}

func toGroup(b model.Group) *pb.Group {
	return &pb.Group{
		Id:          b.ID,
		Description: b.Description,
	}
}
