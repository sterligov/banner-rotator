package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/sterligov/banner-rotator/internal/mocks"
	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"
	"github.com/stretchr/testify/require"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		status pb.HealthCheckResponse_HealthStatus
		err    error
		name   string
	}{
		{pb.HealthCheckResponse_ALIVE, nil, "alive"},
		{pb.HealthCheckResponse_NOT_ALIVE, fmt.Errorf("error"), "not alive"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			pinger := &mocks.Pinger{}

			ctx := context.Background()
			pinger.
				On("PingContext", ctx).
				Return(tst.err).
				Once()
			defer pinger.AssertExpectations(t)

			service := NewHealthService(pinger)
			resp, err := service.Check(ctx, &pb.HealthCheckRequest{})
			require.NoError(t, err)
			require.Equal(t, tst.status, resp.Status)
		})
	}
}
