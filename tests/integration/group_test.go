package integration

import (
	"context"
	"database/sql"
	"errors"

	internalsql "github.com/sterligov/banner-rotator/internal/gateway/sql"
	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
)

func (s *Suite) TestGroupService() {
	groupService := pb.NewGroupServiceClient(s.grpcConn)
	ctx := context.Background()

	s.Run("find by id", func() {
		var groupID int64 = 1
		resp, err := groupService.FindGroupByID(ctx, &pb.FindGroupByIDRequest{Id: groupID})
		s.Require().NoError(err)

		group := s.fetchGroup(groupID)

		s.Require().Equal(group.ID, resp.Group.Id)
		s.Require().Equal(group.Description, resp.Group.Description)
	})

	s.Run("find by id not existing group", func() {
		resp, err := groupService.FindGroupByID(ctx, &pb.FindGroupByIDRequest{Id: 100500})
		s.Require().Nil(resp)
		s.Require().Error(err)
	})

	s.Run("create group", func() {
		expectedDescr := "new group"
		resp, err := groupService.CreateGroup(ctx, &pb.CreateGroupRequest{
			Group: &pb.Group{Description: expectedDescr},
		})
		s.Require().NoError(err)

		group := s.fetchGroup(resp.InsertedId)

		s.Require().Equal(resp.InsertedId, group.ID)
		s.Require().Equal(expectedDescr, group.Description)
	})

	s.Run("update group", func() {
		group := &pb.Group{Id: 4, Description: "updated description"}
		resp, err := groupService.UpdateGroup(ctx, &pb.UpdateGroupRequest{
			Id:    group.Id,
			Group: group,
		})
		s.Require().NoError(err)
		s.Require().Equal(int64(1), resp.Affected)

		actual := s.fetchGroup(group.Id)

		s.Require().Equal(group.Description, actual.Description)
	})

	s.Run("update not existing group", func() {
		group := &pb.Group{Id: 100500, Description: "updated description"}
		resp, err := groupService.UpdateGroup(ctx, &pb.UpdateGroupRequest{
			Id:    group.Id,
			Group: group,
		})
		s.Require().NoError(err)
		s.Require().Equal(int64(0), resp.Affected)
	})

	s.Run("delete group", func() {
		var groupID int64 = 5
		resp, err := groupService.DeleteGroup(ctx, &pb.DeleteGroupRequest{Id: groupID})
		s.Require().NoError(err)
		s.Require().Equal(int64(1), resp.Affected)

		var id int64
		err = s.db.
			QueryRow("SELECT id FROM social_group WHERE id = ?", groupID).
			Scan(&id)
		s.Require().True(errors.Is(err, sql.ErrNoRows))
	})

	s.Run("delete not existing group", func() {
		resp, err := groupService.DeleteGroup(ctx, &pb.DeleteGroupRequest{Id: 100500})
		s.Require().NoError(err)
		s.Require().Equal(int64(0), resp.Affected)
	})
}

func (s *Suite) fetchGroup(id int64) *internalsql.Group {
	group := new(internalsql.Group)
	err := s.db.
		QueryRowx("SELECT * FROM social_group WHERE id = ?", id).
		StructScan(group)
	s.Require().NoError(err)

	return group
}
