package integration

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/nats-io/nats.go"
	internalsql "github.com/sterligov/banner-rotator/internal/gateway/sql"
	"github.com/sterligov/banner-rotator/internal/model"
	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
)

func (s *Suite) TestSimulateUCBWork() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)
	ctx := context.Background()

	var (
		slotID  int64 = 2
		groupID int64 = 2
	)
	slot, err := bannerService.FindAllBannersBySlotID(ctx, &pb.FindAllBannersBySlotIDRequest{
		SlotId: slotID,
	})
	s.Require().NoError(err)

	banners := make(map[int64]int)
	for _, b := range slot.Banners {
		banners[b.Id] = 0
	}

	// проверяем, что каждый баннер был показан хотя бы один раз
	for i := 0; i < len(slot.Banners); i++ {
		resp, err := bannerService.SelectBanner(ctx, &pb.SelectBannerRequest{
			GroupId: groupID,
			SlotId:  slotID,
		})

		s.Require().NoError(err)
		s.Require().Contains(banners, resp.BannerId)
		delete(banners, resp.BannerId)
	}
	s.Require().Equal(0, len(banners), "каждый баннер должен быть показан хотя бы один раз")

	// кликаем на каждый баннер по несколько раз
	nDefaultShows := 10
	for _, b := range slot.Banners {
		for i := 0; i < nDefaultShows; i++ {
			_, err := bannerService.RegisterClick(ctx, &pb.RegisterClickRequest{
				BannerId: b.Id,
				SlotId:   slotID,
				GroupId:  groupID,
			})
			s.Require().NoError(err)
		}
	}

	// кликаем на один и тот же баннер много раз
	clickableBanner := slot.Banners[0]
	for i := 0; i < 100; i++ {
		_, err := bannerService.RegisterClick(ctx, &pb.RegisterClickRequest{
			BannerId: clickableBanner.Id,
			SlotId:   slotID,
			GroupId:  groupID,
		})
		s.Require().NoError(err)
	}

	// теперь он должен выбираться намного чаще, чем остальные
	// при этом остальные баннеры тоже должны выбираться, хоть и меньшее кол-во раз
	for i := 0; i < 200; i++ {
		resp, err := bannerService.SelectBanner(ctx, &pb.SelectBannerRequest{
			GroupId: groupID,
			SlotId:  slotID,
		})
		s.Require().NoError(err)
		banners[resp.BannerId]++
	}

	mostShows := banners[clickableBanner.Id]
	for i := 1; i < len(slot.Banners); i++ {
		s.Require().NotEqual(nDefaultShows, banners[slot.Banners[i].Id], "остальные баннеры тоже должны выбираться")
		s.Require().True(mostShows >= 4*banners[slot.Banners[i].Id], "кол-во показов самого кликабельного баннера должно быть минимум в 4 раза больше")
	}
}

func (s *Suite) TestClickBannerQueueEvent() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)

	var (
		bannerID int64 = 3
		slotID   int64 = 1
		groupID  int64 = 1
	)

	sub, err := s.natsConn.Subscribe("rotator", func(msg *nats.Msg) {
		event := new(model.Event)
		err := json.Unmarshal(msg.Data, event)
		s.Require().NoError(err)

		s.Require().True(event.Date.After(time.Now().Add(-time.Minute)))
		s.Require().True(event.Date.Before(time.Now().Add(time.Minute)))
		s.Require().Equal(bannerID, event.BannerID)
		s.Require().Equal(slotID, event.SlotID)
		s.Require().Equal(groupID, event.GroupID)
		s.Require().Equal(byte(model.EventClick), event.Type)
	})
	s.Require().NoError(err)
	s.Require().NoError(sub.AutoUnsubscribe(1))

	_, err = bannerService.RegisterClick(context.Background(), &pb.RegisterClickRequest{
		BannerId: bannerID,
		SlotId:   slotID,
		GroupId:  groupID,
	})
	s.Require().NoError(err)
}

func (s *Suite) TestSelectBannerQueueEvent() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)

	var (
		slotID  int64 = 1
		groupID int64 = 1
	)

	sub, err := s.natsConn.Subscribe("rotator", func(msg *nats.Msg) {
		event := new(model.Event)
		err := json.Unmarshal(msg.Data, event)
		s.Require().NoError(err)

		s.Require().True(event.Date.After(time.Now().Add(-time.Minute)))
		s.Require().True(event.Date.Before(time.Now().Add(time.Minute)))
		s.Require().NotEmpty(event.BannerID)
		s.Require().Equal(slotID, event.SlotID)
		s.Require().Equal(groupID, event.GroupID)
		s.Require().Equal(byte(model.EventSelect), event.Type)
	})
	s.Require().NoError(err)
	s.Require().NoError(sub.AutoUnsubscribe(1))

	_, err = bannerService.SelectBanner(context.Background(), &pb.SelectBannerRequest{
		SlotId:  slotID,
		GroupId: groupID,
	})
	s.Require().NoError(err)
}

func (s *Suite) TestCreateBannerSlotRelation() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)

	resp, err := bannerService.CreateBannerSlotRelation(context.Background(), &pb.CreateBannerSlotRelationRequest{
		BannerId: 1,
		SlotId:   1,
	})
	s.Require().NoError(err)

	var id int64
	err = s.db.
		QueryRow("SELECT id FROM banner_slot WHERE id = ?", resp.InsertedId).
		Scan(&id)
	s.Require().NoError(err)
}

func (s *Suite) TestDeleteBannerSlotRelation() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)

	var (
		bannerID int64 = 8
		slotID   int64 = 1
	)
	resp, err := bannerService.DeleteBannerSlotRelation(context.Background(), &pb.DeleteBannerSlotRelationRequest{
		BannerId: bannerID,
		SlotId:   slotID,
	})
	s.Require().NoError(err)
	s.Require().Equal(int64(1), resp.Affected)

	var id int64
	err = s.db.
		QueryRow("SELECT id FROM banner_slot WHERE slot_id = ? AND banner_id = ?", slotID, bannerID).
		Scan(&id)
	s.Require().True(errors.Is(err, sql.ErrNoRows))
}

func (s *Suite) TestFindBannerByID() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)

	var bannerID int64 = 1
	resp, err := bannerService.FindBannerByID(context.Background(), &pb.FindBannerByIDRequest{Id: bannerID})
	s.Require().NoError(err)

	banner := s.fetchBanner(bannerID)

	s.Require().Equal(banner.ID, resp.Banner.Id)
	s.Require().Equal(banner.Description, resp.Banner.Description)
}

func (s *Suite) TestFindBannerByID_NotExistingBanner() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)
	resp, err := bannerService.FindBannerByID(context.Background(), &pb.FindBannerByIDRequest{Id: 100500})
	s.Require().Nil(resp)
	s.Require().Error(err)
}

func (s *Suite) TestCreateBanner() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)

	expectedDescr := "new banner"
	resp, err := bannerService.CreateBanner(context.Background(), &pb.CreateBannerRequest{
		Banner: &pb.Banner{Description: expectedDescr},
	})
	s.Require().NoError(err)

	banner := s.fetchBanner(resp.InsertedId)

	s.Require().Equal(resp.InsertedId, banner.ID)
	s.Require().Equal(expectedDescr, banner.Description)
}

func (s *Suite) TestUpdateBanner() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)

	banner := &pb.Banner{Id: 8, Description: "updated description"}
	resp, err := bannerService.UpdateBanner(context.Background(), &pb.UpdateBannerRequest{
		Id:     banner.Id,
		Banner: banner,
	})
	s.Require().NoError(err)
	s.Require().Equal(int64(1), resp.Affected)

	actual := s.fetchBanner(banner.Id)

	s.Require().Equal(banner.Description, actual.Description)
}

func (s *Suite) TestUpdateBanner_NotExistingBanner() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)

	banner := &pb.Banner{Id: 8, Description: "updated description"}
	resp, err := bannerService.UpdateBanner(context.Background(), &pb.UpdateBannerRequest{
		Id:     banner.Id,
		Banner: banner,
	})
	s.Require().NoError(err)
	s.Require().Equal(int64(0), resp.Affected)
}

func (s *Suite) TestDeleteBanner() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)

	var bannerID int64 = 7
	resp, err := bannerService.DeleteBanner(context.Background(), &pb.DeleteBannerRequest{Id: bannerID})
	s.Require().NoError(err)
	s.Require().Equal(int64(1), resp.Affected)

	var id int64
	err = s.db.
		QueryRow("SELECT id FROM banner WHERE id = ?", bannerID).
		Scan(&id)
	s.Require().True(errors.Is(err, sql.ErrNoRows))
}

func (s *Suite) TestDeleteBanner_NotExistingBanner() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)
	resp, err := bannerService.DeleteBanner(context.Background(), &pb.DeleteBannerRequest{Id: 100500})
	s.Require().NoError(err)
	s.Require().Equal(int64(0), resp.Affected)
}

func (s *Suite) fetchBanner(id int64) *internalsql.Banner {
	banner := new(internalsql.Banner)
	err := s.db.
		QueryRowx("SELECT * FROM banner WHERE id = ?", id).
		StructScan(banner)
	s.Require().NoError(err)

	return banner
}
