// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/sterligov/banner-rotator/internal/bandit/ucb"
	"github.com/sterligov/banner-rotator/internal/config"
	"github.com/sterligov/banner-rotator/internal/gateway/nats"
	"github.com/sterligov/banner-rotator/internal/gateway/sql"
	"github.com/sterligov/banner-rotator/internal/server"
	internalgrpc "github.com/sterligov/banner-rotator/internal/server/grpc"
	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
	"github.com/sterligov/banner-rotator/internal/server/grpc/service"
	internalhttp "github.com/sterligov/banner-rotator/internal/server/http"
	"github.com/sterligov/banner-rotator/internal/usecase/banner"
	"github.com/sterligov/banner-rotator/internal/usecase/group"
	"github.com/sterligov/banner-rotator/internal/usecase/slot"
)

func setup(*config.Config) (*server.Server, func(), error) {
	panic(wire.Build(
		wire.Bind(new(banner.EventGateway), new(*nats.EventGateway)),
		wire.Bind(new(banner.BannerGateway), new(*sql.BannerGateway)),
		wire.Bind(new(banner.StatisticGateway), new(*sql.StatisticGateway)),
		wire.Bind(new(group.GroupGateway), new(*sql.GroupGateway)),
		wire.Bind(new(slot.SlotGateway), new(*sql.SlotGateway)),
		wire.Bind(new(service.BannerUseCase), new(*banner.UseCase)),
		wire.Bind(new(service.SlotUseCase), new(*slot.UseCase)),
		wire.Bind(new(service.GroupUseCase), new(*group.UseCase)),
		wire.Bind(new(service.Pinger), new(*sqlx.DB)),
		wire.Bind(new(pb.BannerServiceServer), new(*service.BannerService)),
		wire.Bind(new(pb.GroupServiceServer), new(*service.GroupService)),
		wire.Bind(new(pb.SlotServiceServer), new(*service.SlotService)),
		wire.Bind(new(pb.HealthServiceServer), new(*service.HealthService)),
		wire.Bind(new(banner.Bandit), new(*ucb.UCB)),

		ucb.New,
		sql.NewDatabase,
		sql.NewBannerGateway,
		sql.NewSlotGateway,
		sql.NewGroupGateway,
		sql.NewStatisticGateway,
		nats.NewNatsConnection,
		nats.NewEventGateway,
		banner.NewUseCase,
		group.NewUseCase,
		slot.NewUseCase,
		service.NewHealthService,
		service.NewBannerService,
		service.NewSlotService,
		service.NewGroupService,
		internalhttp.NewHandler,
		internalhttp.NewServer,
		internalgrpc.NewServer,
		server.NewServer,
	))
}
