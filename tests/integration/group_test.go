//+build integration

//nolint:dupl
package integration

import (
	"context"
	"database/sql"
	"errors"

	internalsql "github.com/sterligov/banner-rotator/internal/gateway/sql"
	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
)

func (s *Suite) TestFindAllGroups() {
	groupService := pb.NewGroupServiceClient(s.grpcConn)

	resp, err := groupService.FindAllGroups(context.Background(), &pb.FindAllGroupsRequest{})
	s.Require().NoError(err)

	rows, err := s.db.Queryx("SELECT * FROM social_group")
	s.Require().NoError(err)
	defer func() {
		s.Require().NoError(rows.Close())
	}()

	m := make(map[int64]*pb.Group, len(resp.Groups))
	for _, b := range resp.Groups {
		m[b.Id] = b
	}

	var (
		group   internalsql.Group
		nGroups = len(resp.Groups)
	)

	for rows.Next() {
		err = rows.StructScan(&group)
		s.Require().NoError(err)
		s.Require().Contains(m, group.ID)
		s.Require().Equal(m[group.ID].Id, group.ID)
		s.Require().Equal(m[group.ID].Description, group.Description)
		delete(m, group.ID)
		nGroups--
	}

	s.Require().Empty(nGroups)
	s.Require().NoError(rows.Err())
}

func (s *Suite) TestFindGroupByID() {
	groupService := pb.NewGroupServiceClient(s.grpcConn)

	var groupID int64 = 1
	resp, err := groupService.FindGroupByID(context.Background(), &pb.FindGroupByIDRequest{Id: groupID})
	s.Require().NoError(err)

	group := s.fetchGroup(groupID)

	s.Require().Equal(group.ID, resp.Group.Id)
	s.Require().Equal(group.Description, resp.Group.Description)
}

func (s *Suite) TestFindGroupByID_NotExistingGroup() {
	groupService := pb.NewGroupServiceClient(s.grpcConn)
	resp, err := groupService.FindGroupByID(context.Background(), &pb.FindGroupByIDRequest{Id: 100500})
	s.Require().Nil(resp)
	s.Require().Error(err)
}

func (s *Suite) TestCreateGroup() {
	groupService := pb.NewGroupServiceClient(s.grpcConn)

	expectedDescr := "new group"
	resp, err := groupService.CreateGroup(context.Background(), &pb.CreateGroupRequest{
		Group: &pb.Group{Description: expectedDescr},
	})
	s.Require().NoError(err)

	group := s.fetchGroup(resp.InsertedId)

	s.Require().Equal(resp.InsertedId, group.ID)
	s.Require().Equal(expectedDescr, group.Description)
}

func (s *Suite) TestUpdateGroup() {
	groupService := pb.NewGroupServiceClient(s.grpcConn)

	group := &pb.Group{Id: 4, Description: "updated description"}
	resp, err := groupService.UpdateGroup(context.Background(), &pb.UpdateGroupRequest{
		Id:    group.Id,
		Group: group,
	})
	s.Require().NoError(err)
	s.Require().Equal(int64(1), resp.Affected)

	actual := s.fetchGroup(group.Id)

	s.Require().Equal(group.Description, actual.Description)
}

func (s *Suite) TestUpdateGroup_NotExistingGroup() {
	groupService := pb.NewGroupServiceClient(s.grpcConn)

	group := &pb.Group{Id: 100500, Description: "updated description"}
	resp, err := groupService.UpdateGroup(context.Background(), &pb.UpdateGroupRequest{
		Id:    group.Id,
		Group: group,
	})
	s.Require().NoError(err)
	s.Require().Equal(int64(0), resp.Affected)
}

func (s *Suite) TestDeleteGroup() {
	groupService := pb.NewGroupServiceClient(s.grpcConn)

	var groupID int64 = 5
	resp, err := groupService.DeleteGroup(context.Background(), &pb.DeleteGroupRequest{Id: groupID})
	s.Require().NoError(err)
	s.Require().Equal(int64(1), resp.Affected)

	var id int64
	err = s.db.
		QueryRow("SELECT id FROM social_group WHERE id = ?", groupID).
		Scan(&id)
	s.Require().True(errors.Is(err, sql.ErrNoRows))
}

func (s *Suite) TestDeleteGroup_NotExistingGroup() {
	groupService := pb.NewGroupServiceClient(s.grpcConn)
	resp, err := groupService.DeleteGroup(context.Background(), &pb.DeleteGroupRequest{Id: 100500})
	s.Require().NoError(err)
	s.Require().Equal(int64(0), resp.Affected)
}

func (s *Suite) fetchGroup(id int64) *internalsql.Group {
	group := new(internalsql.Group)
	err := s.db.
		QueryRowx("SELECT * FROM social_group WHERE id = ?", id).
		StructScan(group)
	s.Require().NoError(err)

	return group
}
