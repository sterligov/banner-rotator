package service

import (
	"context"

	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
)

type (
	Pinger interface {
		PingContext(ctx context.Context) error
	}

	HealthService struct {
		pb.UnimplementedHealthServiceServer

		pinger Pinger
	}
)

func NewHealthService(pinger Pinger) *HealthService {
	return &HealthService{pinger: pinger}
}

func (hs *HealthService) Check(ctx context.Context, _ *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	if err := hs.pinger.PingContext(ctx); err != nil {
		return &pb.HealthCheckResponse{Status: pb.HealthCheckResponse_NOT_ALIVE}, nil
	}

	return &pb.HealthCheckResponse{Status: pb.HealthCheckResponse_ALIVE}, nil
}
