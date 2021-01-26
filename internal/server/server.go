package server

import (
	internalgrpc "github.com/sterligov/banner-rotator/internal/server/grpc"
	internalhttp "github.com/sterligov/banner-rotator/internal/server/http"
)

type Server struct {
	GRPC *internalgrpc.Server
	HTTP *internalhttp.Server
}

func NewServer(grpcServer *internalgrpc.Server, httpServer *internalhttp.Server) *Server {
	return &Server{
		GRPC: grpcServer,
		HTTP: httpServer,
	}
}
