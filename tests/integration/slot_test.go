package integration

import (
	"context"
	"database/sql"
	"errors"

	internalsql "github.com/sterligov/banner-rotator/internal/gateway/sql"

	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
)

func (s *Suite) TestSlotService() {
	slotService := pb.NewSlotServiceClient(s.grpcConn)
	ctx := context.Background()

	s.Run("find by id", func() {
		var slotID int64 = 1
		resp, err := slotService.FindSlotByID(ctx, &pb.FindSlotByIDRequest{Id: slotID})
		s.Require().NoError(err)

		slot := s.fetchSlot(slotID)

		s.Require().Equal(slot.ID, resp.Slot.Id)
		s.Require().Equal(slot.Description, resp.Slot.Description)
	})

	s.Run("find by id not existing slot", func() {
		resp, err := slotService.FindSlotByID(ctx, &pb.FindSlotByIDRequest{Id: 100500})
		s.Require().Nil(resp)
		s.Require().Error(err)
	})

	s.Run("create slot", func() {
		expectedDescr := "new slot"
		resp, err := slotService.CreateSlot(ctx, &pb.CreateSlotRequest{
			Slot: &pb.Slot{Description: expectedDescr},
		})
		s.Require().NoError(err)

		slot := s.fetchSlot(resp.InsertedId)

		s.Require().Equal(resp.InsertedId, slot.ID)
		s.Require().Equal(expectedDescr, slot.Description)
	})

	s.Run("update slot", func() {
		slot := &pb.Slot{Id: 4, Description: "updated description"}
		resp, err := slotService.UpdateSlot(ctx, &pb.UpdateSlotRequest{
			Id:   slot.Id,
			Slot: slot,
		})
		s.Require().NoError(err)
		s.Require().Equal(int64(1), resp.Affected)

		actual := s.fetchSlot(slot.Id)

		s.Require().Equal(slot.Description, actual.Description)
	})

	s.Run("update not existing slot", func() {
		slot := &pb.Slot{Id: 100500, Description: "updated description"}
		resp, err := slotService.UpdateSlot(ctx, &pb.UpdateSlotRequest{
			Id:   slot.Id,
			Slot: slot,
		})
		s.Require().NoError(err)
		s.Require().Equal(int64(0), resp.Affected)
	})

	s.Run("delete slot", func() {
		var slotID int64 = 5
		resp, err := slotService.DeleteSlot(ctx, &pb.DeleteSlotRequest{Id: slotID})
		s.Require().NoError(err)
		s.Require().Equal(int64(1), resp.Affected)

		var id int64
		err = s.db.
			QueryRow("SELECT id FROM slot WHERE id = ?", slotID).
			Scan(&id)
		s.Require().True(errors.Is(err, sql.ErrNoRows))
	})

	s.Run("delete not existing slot", func() {
		resp, err := slotService.DeleteSlot(ctx, &pb.DeleteSlotRequest{Id: 100500})
		s.Require().NoError(err)
		s.Require().Equal(int64(0), resp.Affected)
	})
}

func (s *Suite) fetchSlot(id int64) *internalsql.Slot {
	slot := new(internalsql.Slot)
	err := s.db.
		QueryRowx("SELECT * FROM slot WHERE id = ?", id).
		StructScan(slot)
	s.Require().NoError(err)

	return slot
}
