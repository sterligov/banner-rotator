// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthServiceClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/health.HealthService/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
// All implementations must embed UnimplementedHealthServiceServer
// for forward compatibility
type HealthServiceServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedHealthServiceServer()
}

// UnimplementedHealthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (UnimplementedHealthServiceServer) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedHealthServiceServer) mustEmbedUnimplementedHealthServiceServer() {}

// UnsafeHealthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServiceServer will
// result in compilation errors.
type UnsafeHealthServiceServer interface {
	mustEmbedUnimplementedHealthServiceServer()
}

func RegisterHealthServiceServer(s grpc.ServiceRegistrar, srv HealthServiceServer) {
	s.RegisterService(&_HealthService_serviceDesc, srv)
}

func _HealthService_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health.HealthService/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "health.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _HealthService_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/health_service.proto",
}
