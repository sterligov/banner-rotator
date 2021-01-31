package grpc

import (
	"fmt"
	"net"

	"github.com/sterligov/banner-rotator/internal/config"
	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	addr       string
}

func NewServer(
	cfg *config.Config,
	bannerServer pb.BannerServiceServer,
	slotServer pb.SlotServiceServer,
	groupServer pb.GroupServiceServer,
	healthServer pb.HealthServiceServer,
) *Server {
	chainInterceptor := grpc.ChainUnaryInterceptor(
		ErrorInterceptor,
		LoggingInterceptor,
	)
	grpcServer := grpc.NewServer(chainInterceptor)
	pb.RegisterBannerServiceServer(grpcServer, bannerServer)
	pb.RegisterSlotServiceServer(grpcServer, slotServer)
	pb.RegisterGroupServiceServer(grpcServer, groupServer)
	pb.RegisterHealthServiceServer(grpcServer, healthServer)

	return &Server{
		grpcServer: grpcServer,
		addr:       cfg.GRPC.Addr,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("start grpc server failed: %w", err)
	}

	zap.L().Info("Start grpc server...")

	return s.grpcServer.Serve(listener)
}

func (s *Server) Stop() {
	zap.L().Info("Stop grpc server...")

	s.grpcServer.GracefulStop()
}
