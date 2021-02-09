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

func TestFindGroupByID(t *testing.T) {
	tests := []struct {
		group model.Group
		err   error
		code  codes.Code
		name  string
	}{
		{model.Group{ID: 1, Description: "descr"}, nil, codes.OK, "ok"},
		{model.Group{}, fmt.Errorf("error"), codes.Unknown, "unexpected error"},
		{model.Group{}, model.ErrNotFound, codes.NotFound, "group not found"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			groupUC := &mocks.GroupUseCase{}

			r := &pb.FindGroupByIDRequest{
				Id: 1,
			}
			ctx := context.Background()

			groupUC.
				On("FindGroupByID", ctx, r.Id).
				Return(tst.group, tst.err).
				Once()
			defer groupUC.AssertExpectations(t)

			service := NewGroupService(groupUC)
			resp, err := service.FindGroupByID(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.code == codes.OK {
				require.Equal(t, tst.group.ID, resp.Group.Id)
				require.Equal(t, tst.group.Description, resp.Group.Description)
			}
		})
	}
}

func TestFindAllGroups(t *testing.T) {
	tests := []struct {
		groups []model.Group
		err    error
		code   codes.Code
		name   string
	}{
		{
			[]model.Group{{ID: 1, Description: "descr"}, {ID: 2, Description: "descr2"}},
			nil,
			codes.OK,
			"ok",
		},
		{
			[]model.Group{},
			fmt.Errorf("error"),
			codes.Unknown,
			"unexpected error",
		},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			groupUC := &mocks.GroupUseCase{}

			r := &pb.FindAllGroupsRequest{}
			ctx := context.Background()

			groupUC.
				On("FindAllGroups", ctx).
				Return(tst.groups, tst.err).
				Once()
			defer groupUC.AssertExpectations(t)

			service := NewGroupService(groupUC)
			resp, err := service.FindAllGroups(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.code == codes.OK {
				require.Equal(t, len(tst.groups), len(resp.Groups))
				for i := 0; i < len(tst.groups); i++ {
					require.Equal(t, tst.groups[i].ID, resp.Groups[i].Id)
					require.Equal(t, tst.groups[i].Description, resp.Groups[i].Description)
				}
			}
		})
	}
}

func TestCreateGroup(t *testing.T) {
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
			groupUC := &mocks.GroupUseCase{}

			r := &pb.CreateGroupRequest{
				Group: &pb.Group{Description: "descr"},
			}
			ctx := context.Background()

			groupUC.
				On("CreateGroup", ctx, mock.MatchedBy(func(e model.Group) bool {
					return e.Description == r.Group.Description
				})).
				Return(tst.insertedID, tst.err).
				Once()
			defer groupUC.AssertExpectations(t)

			service := NewGroupService(groupUC)
			resp, err := service.CreateGroup(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.insertedID, resp.InsertedId)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}

func TestDeleteGroup(t *testing.T) {
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
			groupUC := &mocks.GroupUseCase{}

			r := &pb.DeleteGroupRequest{
				Id: 1,
			}
			ctx := context.Background()

			groupUC.
				On("DeleteGroupByID", ctx, r.Id).
				Return(tst.affected, tst.err).
				Once()
			defer groupUC.AssertExpectations(t)

			service := NewGroupService(groupUC)
			resp, err := service.DeleteGroup(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.affected, resp.Affected)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}

func TestUpdateGroup(t *testing.T) {
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
			groupUC := &mocks.GroupUseCase{}

			r := &pb.UpdateGroupRequest{
				Id: 1,
				Group: &pb.Group{
					Id:          1,
					Description: "descr",
				},
			}
			ctx := context.Background()

			groupUC.
				On("UpdateGroup", ctx, mock.MatchedBy(func(e model.Group) bool {
					return e.Description == r.Group.Description && e.ID == r.Group.Id
				})).
				Return(tst.affected, tst.err).
				Once()
			defer groupUC.AssertExpectations(t)

			service := NewGroupService(groupUC)
			resp, err := service.UpdateGroup(ctx, r)
			require.Equal(t, tst.code, status.Code(err))
			if tst.err == nil {
				require.Equal(t, tst.affected, resp.Affected)
			} else {
				require.Nil(t, resp)
			}
		})
	}
}
