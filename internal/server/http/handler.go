package internalhttp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sterligov/banner-rotator/internal/server/grpc/pb"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sterligov/banner-rotator/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewHandler(cfg *config.Config) (http.Handler, error) {
	jsonPb := &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
	}
	gw := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonPb))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	ctx := context.Background()

	err := pb.RegisterBannerServiceHandlerFromEndpoint(ctx, gw, cfg.GRPC.Addr, opts)
	if err != nil {
		return nil, fmt.Errorf("register banner service handler endpoint: %w", err)
	}

	err = pb.RegisterGroupServiceHandlerFromEndpoint(ctx, gw, cfg.GRPC.Addr, opts)
	if err != nil {
		return nil, fmt.Errorf("register group service handler endpoint: %w", err)
	}

	err = pb.RegisterSlotServiceHandlerFromEndpoint(ctx, gw, cfg.GRPC.Addr, opts)
	if err != nil {
		return nil, fmt.Errorf("register slot service handler endpoint: %w", err)
	}

	err = pb.RegisterHealthServiceHandlerFromEndpoint(ctx, gw, cfg.GRPC.Addr, opts)
	if err != nil {
		return nil, fmt.Errorf("register health service handler endpoint: %w", err)
	}

	mux := http.NewServeMux()
	handler := HeadersMiddleware(gw)
	handler = LoggingMiddleware(handler)
	mux.Handle("/", handler)

	return mux, nil
}
