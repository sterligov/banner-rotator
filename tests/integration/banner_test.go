package integration

import (
	"context"
	"database/sql"
	"errors"

	internalsql "github.com/sterligov/banner-rotator/internal/gateway/sql"
	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
)

func (s *Suite) TestBannerService() {
	bannerService := pb.NewBannerServiceClient(s.grpcConn)
	ctx := context.Background()

	s.Run("simulate ucb work", func() {
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
		// также проверяем, что не были показаны баннеры не привязанные к данному слоту
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

		// кликаем на каждый баннер по 10 раз
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
	})

	s.Run("create banner slot relation", func() {
		resp, err := bannerService.CreateBannerSlotRelation(ctx, &pb.CreateBannerSlotRelationRequest{
			BannerId: 1,
			SlotId:   1,
		})
		s.Require().NoError(err)

		var id int64
		err = s.db.
			QueryRow("SELECT id FROM banner_slot WHERE id = ?", resp.InsertedId).
			Scan(&id)
		s.Require().NoError(err)
	})

	s.Run("delete banner slot relation", func() {
		var (
			bannerID int64 = 8
			slotID   int64 = 1
		)
		resp, err := bannerService.DeleteBannerSlotRelation(ctx, &pb.DeleteBannerSlotRelationRequest{
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
	})

	s.Run("find by id", func() {
		var bannerID int64 = 1
		resp, err := bannerService.FindBannerByID(ctx, &pb.FindBannerByIDRequest{Id: bannerID})
		s.Require().NoError(err)

		banner := s.fetchBanner(bannerID)

		s.Require().Equal(banner.ID, resp.Banner.Id)
		s.Require().Equal(banner.Description, resp.Banner.Description)
	})

	s.Run("find by id not existing banner", func() {
		resp, err := bannerService.FindBannerByID(ctx, &pb.FindBannerByIDRequest{Id: 100500})
		s.Require().Nil(resp)
		s.Require().Error(err)
	})

	s.Run("create banner", func() {
		expectedDescr := "new banner"
		resp, err := bannerService.CreateBanner(ctx, &pb.CreateBannerRequest{
			Banner: &pb.Banner{Description: expectedDescr},
		})
		s.Require().NoError(err)

		banner := s.fetchBanner(resp.InsertedId)

		s.Require().Equal(resp.InsertedId, banner.ID)
		s.Require().Equal(expectedDescr, banner.Description)
	})

	s.Run("update banner", func() {
		banner := &pb.Banner{Id: 8, Description: "updated description"}
		resp, err := bannerService.UpdateBanner(ctx, &pb.UpdateBannerRequest{
			Id:     banner.Id,
			Banner: banner,
		})
		s.Require().NoError(err)
		s.Require().Equal(int64(1), resp.Affected)

		actual := s.fetchBanner(banner.Id)

		s.Require().Equal(banner.Description, actual.Description)
	})

	s.Run("update not existing banner", func() {
		banner := &pb.Banner{Id: 8, Description: "updated description"}
		resp, err := bannerService.UpdateBanner(ctx, &pb.UpdateBannerRequest{
			Id:     banner.Id,
			Banner: banner,
		})
		s.Require().NoError(err)
		s.Require().Equal(int64(0), resp.Affected)
	})

	s.Run("delete banner", func() {
		var bannerID int64 = 7
		resp, err := bannerService.DeleteBanner(ctx, &pb.DeleteBannerRequest{Id: bannerID})
		s.Require().NoError(err)
		s.Require().Equal(int64(1), resp.Affected)

		var id int64
		err = s.db.
			QueryRow("SELECT id FROM banner WHERE id = ?", bannerID).
			Scan(&id)
		s.Require().True(errors.Is(err, sql.ErrNoRows))
	})

	s.Run("delete not existing banner", func() {
		resp, err := bannerService.DeleteBanner(ctx, &pb.DeleteBannerRequest{Id: 100500})
		s.Require().NoError(err)
		s.Require().Equal(int64(0), resp.Affected)
	})
}

func (s *Suite) fetchBanner(id int64) *internalsql.Banner {
	banner := new(internalsql.Banner)
	err := s.db.
		QueryRowx("SELECT * FROM banner WHERE id = ?", id).
		StructScan(banner)
	s.Require().NoError(err)

	return banner
}
