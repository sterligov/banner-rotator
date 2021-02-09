package integration //nolint:dupl

import (
	"context"
	"database/sql"
	"errors"

	internalsql "github.com/sterligov/banner-rotator/internal/gateway/sql"
	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
)

func (s *Suite) TestFindSlotByID() {
	slotService := pb.NewSlotServiceClient(s.grpcConn)

	var slotID int64 = 1
	resp, err := slotService.FindSlotByID(context.Background(), &pb.FindSlotByIDRequest{Id: slotID})
	s.Require().NoError(err)

	slot := s.fetchSlot(slotID)

	s.Require().Equal(slot.ID, resp.Slot.Id)
	s.Require().Equal(slot.Description, resp.Slot.Description)
}

func (s *Suite) TestFindSlotByID_NotExistingSlot() {
	slotService := pb.NewSlotServiceClient(s.grpcConn)
	resp, err := slotService.FindSlotByID(context.Background(), &pb.FindSlotByIDRequest{Id: 100500})
	s.Require().Nil(resp)
	s.Require().Error(err)
}

func (s *Suite) TestCreateSlot() {
	slotService := pb.NewSlotServiceClient(s.grpcConn)

	expectedDescr := "new slot"
	resp, err := slotService.CreateSlot(context.Background(), &pb.CreateSlotRequest{
		Slot: &pb.Slot{Description: expectedDescr},
	})
	s.Require().NoError(err)

	slot := s.fetchSlot(resp.InsertedId)

	s.Require().Equal(resp.InsertedId, slot.ID)
	s.Require().Equal(expectedDescr, slot.Description)
}

func (s *Suite) TestUpdateSlot() {
	slotService := pb.NewSlotServiceClient(s.grpcConn)

	slot := &pb.Slot{Id: 4, Description: "updated description"}
	resp, err := slotService.UpdateSlot(context.Background(), &pb.UpdateSlotRequest{
		Id:   slot.Id,
		Slot: slot,
	})
	s.Require().NoError(err)
	s.Require().Equal(int64(1), resp.Affected)

	actual := s.fetchSlot(slot.Id)

	s.Require().Equal(slot.Description, actual.Description)
}

func (s *Suite) TestUpdateSlot_NotExistingSlot() {
	slotService := pb.NewSlotServiceClient(s.grpcConn)

	slot := &pb.Slot{Id: 100500, Description: "updated description"}
	resp, err := slotService.UpdateSlot(context.Background(), &pb.UpdateSlotRequest{
		Id:   slot.Id,
		Slot: slot,
	})
	s.Require().NoError(err)
	s.Require().Equal(int64(0), resp.Affected)
}

func (s *Suite) TestDeleteSlot() {
	slotService := pb.NewSlotServiceClient(s.grpcConn)

	var slotID int64 = 5
	resp, err := slotService.DeleteSlot(context.Background(), &pb.DeleteSlotRequest{Id: slotID})
	s.Require().NoError(err)
	s.Require().Equal(int64(1), resp.Affected)

	var id int64
	err = s.db.
		QueryRow("SELECT id FROM slot WHERE id = ?", slotID).
		Scan(&id)
	s.Require().True(errors.Is(err, sql.ErrNoRows))
}

func (s *Suite) TestDeleteSlot_NotExistingSlot() {
	slotService := pb.NewSlotServiceClient(s.grpcConn)
	resp, err := slotService.DeleteSlot(context.Background(), &pb.DeleteSlotRequest{Id: 100500})
	s.Require().NoError(err)
	s.Require().Equal(int64(0), resp.Affected)
}

func (s *Suite) fetchSlot(id int64) *internalsql.Slot {
	slot := new(internalsql.Slot)
	err := s.db.
		QueryRowx("SELECT * FROM slot WHERE id = ?", id).
		StructScan(slot)
	s.Require().NoError(err)

	return slot
}
