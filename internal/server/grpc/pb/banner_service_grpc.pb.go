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

// BannerServiceClient is the client API for BannerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BannerServiceClient interface {
	FindBannerByID(ctx context.Context, in *FindBannerByIDRequest, opts ...grpc.CallOption) (*FindBannerByIDResponse, error)
	FindAllBanners(ctx context.Context, in *FindAllBannersRequest, opts ...grpc.CallOption) (*FindAllBannersResponse, error)
	FindAllBannersBySlotID(ctx context.Context, in *FindAllBannersBySlotIDRequest, opts ...grpc.CallOption) (*FindAllBannersBySlotIDResponse, error)
	CreateBanner(ctx context.Context, in *CreateBannerRequest, opts ...grpc.CallOption) (*CreateBannerResponse, error)
	DeleteBanner(ctx context.Context, in *DeleteBannerRequest, opts ...grpc.CallOption) (*DeleteBannerResponse, error)
	UpdateBanner(ctx context.Context, in *UpdateBannerRequest, opts ...grpc.CallOption) (*UpdateBannerResponse, error)
	CreateBannerSlotRelation(ctx context.Context, in *CreateBannerSlotRelationRequest, opts ...grpc.CallOption) (*CreateBannerSlotRelationResponse, error)
	DeleteBannerSlotRelation(ctx context.Context, in *DeleteBannerSlotRelationRequest, opts ...grpc.CallOption) (*DeleteBannerSlotRelationResponse, error)
	RegisterClick(ctx context.Context, in *RegisterClickRequest, opts ...grpc.CallOption) (*RegisterClickResponse, error)
	SelectBanner(ctx context.Context, in *SelectBannerRequest, opts ...grpc.CallOption) (*SelectBannerResponse, error)
}

type bannerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBannerServiceClient(cc grpc.ClientConnInterface) BannerServiceClient {
	return &bannerServiceClient{cc}
}

func (c *bannerServiceClient) FindBannerByID(ctx context.Context, in *FindBannerByIDRequest, opts ...grpc.CallOption) (*FindBannerByIDResponse, error) {
	out := new(FindBannerByIDResponse)
	err := c.cc.Invoke(ctx, "/banner.BannerService/FindBannerByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) FindAllBanners(ctx context.Context, in *FindAllBannersRequest, opts ...grpc.CallOption) (*FindAllBannersResponse, error) {
	out := new(FindAllBannersResponse)
	err := c.cc.Invoke(ctx, "/banner.BannerService/FindAllBanners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) FindAllBannersBySlotID(ctx context.Context, in *FindAllBannersBySlotIDRequest, opts ...grpc.CallOption) (*FindAllBannersBySlotIDResponse, error) {
	out := new(FindAllBannersBySlotIDResponse)
	err := c.cc.Invoke(ctx, "/banner.BannerService/FindAllBannersBySlotID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) CreateBanner(ctx context.Context, in *CreateBannerRequest, opts ...grpc.CallOption) (*CreateBannerResponse, error) {
	out := new(CreateBannerResponse)
	err := c.cc.Invoke(ctx, "/banner.BannerService/CreateBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) DeleteBanner(ctx context.Context, in *DeleteBannerRequest, opts ...grpc.CallOption) (*DeleteBannerResponse, error) {
	out := new(DeleteBannerResponse)
	err := c.cc.Invoke(ctx, "/banner.BannerService/DeleteBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) UpdateBanner(ctx context.Context, in *UpdateBannerRequest, opts ...grpc.CallOption) (*UpdateBannerResponse, error) {
	out := new(UpdateBannerResponse)
	err := c.cc.Invoke(ctx, "/banner.BannerService/UpdateBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) CreateBannerSlotRelation(ctx context.Context, in *CreateBannerSlotRelationRequest, opts ...grpc.CallOption) (*CreateBannerSlotRelationResponse, error) {
	out := new(CreateBannerSlotRelationResponse)
	err := c.cc.Invoke(ctx, "/banner.BannerService/CreateBannerSlotRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) DeleteBannerSlotRelation(ctx context.Context, in *DeleteBannerSlotRelationRequest, opts ...grpc.CallOption) (*DeleteBannerSlotRelationResponse, error) {
	out := new(DeleteBannerSlotRelationResponse)
	err := c.cc.Invoke(ctx, "/banner.BannerService/DeleteBannerSlotRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) RegisterClick(ctx context.Context, in *RegisterClickRequest, opts ...grpc.CallOption) (*RegisterClickResponse, error) {
	out := new(RegisterClickResponse)
	err := c.cc.Invoke(ctx, "/banner.BannerService/RegisterClick", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bannerServiceClient) SelectBanner(ctx context.Context, in *SelectBannerRequest, opts ...grpc.CallOption) (*SelectBannerResponse, error) {
	out := new(SelectBannerResponse)
	err := c.cc.Invoke(ctx, "/banner.BannerService/SelectBanner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BannerServiceServer is the server API for BannerService service.
// All implementations must embed UnimplementedBannerServiceServer
// for forward compatibility
type BannerServiceServer interface {
	FindBannerByID(context.Context, *FindBannerByIDRequest) (*FindBannerByIDResponse, error)
	FindAllBanners(context.Context, *FindAllBannersRequest) (*FindAllBannersResponse, error)
	FindAllBannersBySlotID(context.Context, *FindAllBannersBySlotIDRequest) (*FindAllBannersBySlotIDResponse, error)
	CreateBanner(context.Context, *CreateBannerRequest) (*CreateBannerResponse, error)
	DeleteBanner(context.Context, *DeleteBannerRequest) (*DeleteBannerResponse, error)
	UpdateBanner(context.Context, *UpdateBannerRequest) (*UpdateBannerResponse, error)
	CreateBannerSlotRelation(context.Context, *CreateBannerSlotRelationRequest) (*CreateBannerSlotRelationResponse, error)
	DeleteBannerSlotRelation(context.Context, *DeleteBannerSlotRelationRequest) (*DeleteBannerSlotRelationResponse, error)
	RegisterClick(context.Context, *RegisterClickRequest) (*RegisterClickResponse, error)
	SelectBanner(context.Context, *SelectBannerRequest) (*SelectBannerResponse, error)
	mustEmbedUnimplementedBannerServiceServer()
}

// UnimplementedBannerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBannerServiceServer struct {
}

func (UnimplementedBannerServiceServer) FindBannerByID(context.Context, *FindBannerByIDRequest) (*FindBannerByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBannerByID not implemented")
}
func (UnimplementedBannerServiceServer) FindAllBanners(context.Context, *FindAllBannersRequest) (*FindAllBannersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllBanners not implemented")
}
func (UnimplementedBannerServiceServer) FindAllBannersBySlotID(context.Context, *FindAllBannersBySlotIDRequest) (*FindAllBannersBySlotIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllBannersBySlotID not implemented")
}
func (UnimplementedBannerServiceServer) CreateBanner(context.Context, *CreateBannerRequest) (*CreateBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBanner not implemented")
}
func (UnimplementedBannerServiceServer) DeleteBanner(context.Context, *DeleteBannerRequest) (*DeleteBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBanner not implemented")
}
func (UnimplementedBannerServiceServer) UpdateBanner(context.Context, *UpdateBannerRequest) (*UpdateBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBanner not implemented")
}
func (UnimplementedBannerServiceServer) CreateBannerSlotRelation(context.Context, *CreateBannerSlotRelationRequest) (*CreateBannerSlotRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBannerSlotRelation not implemented")
}
func (UnimplementedBannerServiceServer) DeleteBannerSlotRelation(context.Context, *DeleteBannerSlotRelationRequest) (*DeleteBannerSlotRelationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBannerSlotRelation not implemented")
}
func (UnimplementedBannerServiceServer) RegisterClick(context.Context, *RegisterClickRequest) (*RegisterClickResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClick not implemented")
}
func (UnimplementedBannerServiceServer) SelectBanner(context.Context, *SelectBannerRequest) (*SelectBannerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectBanner not implemented")
}
func (UnimplementedBannerServiceServer) mustEmbedUnimplementedBannerServiceServer() {}

// UnsafeBannerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BannerServiceServer will
// result in compilation errors.
type UnsafeBannerServiceServer interface {
	mustEmbedUnimplementedBannerServiceServer()
}

func RegisterBannerServiceServer(s grpc.ServiceRegistrar, srv BannerServiceServer) {
	s.RegisterService(&_BannerService_serviceDesc, srv)
}

func _BannerService_FindBannerByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBannerByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).FindBannerByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banner.BannerService/FindBannerByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).FindBannerByID(ctx, req.(*FindBannerByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_FindAllBanners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllBannersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).FindAllBanners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banner.BannerService/FindAllBanners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).FindAllBanners(ctx, req.(*FindAllBannersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_FindAllBannersBySlotID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllBannersBySlotIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).FindAllBannersBySlotID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banner.BannerService/FindAllBannersBySlotID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).FindAllBannersBySlotID(ctx, req.(*FindAllBannersBySlotIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_CreateBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).CreateBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banner.BannerService/CreateBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).CreateBanner(ctx, req.(*CreateBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_DeleteBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).DeleteBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banner.BannerService/DeleteBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).DeleteBanner(ctx, req.(*DeleteBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_UpdateBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).UpdateBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banner.BannerService/UpdateBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).UpdateBanner(ctx, req.(*UpdateBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_CreateBannerSlotRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBannerSlotRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).CreateBannerSlotRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banner.BannerService/CreateBannerSlotRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).CreateBannerSlotRelation(ctx, req.(*CreateBannerSlotRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_DeleteBannerSlotRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBannerSlotRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).DeleteBannerSlotRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banner.BannerService/DeleteBannerSlotRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).DeleteBannerSlotRelation(ctx, req.(*DeleteBannerSlotRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_RegisterClick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterClickRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).RegisterClick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banner.BannerService/RegisterClick",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).RegisterClick(ctx, req.(*RegisterClickRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BannerService_SelectBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectBannerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BannerServiceServer).SelectBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banner.BannerService/SelectBanner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BannerServiceServer).SelectBanner(ctx, req.(*SelectBannerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BannerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "banner.BannerService",
	HandlerType: (*BannerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindBannerByID",
			Handler:    _BannerService_FindBannerByID_Handler,
		},
		{
			MethodName: "FindAllBanners",
			Handler:    _BannerService_FindAllBanners_Handler,
		},
		{
			MethodName: "FindAllBannersBySlotID",
			Handler:    _BannerService_FindAllBannersBySlotID_Handler,
		},
		{
			MethodName: "CreateBanner",
			Handler:    _BannerService_CreateBanner_Handler,
		},
		{
			MethodName: "DeleteBanner",
			Handler:    _BannerService_DeleteBanner_Handler,
		},
		{
			MethodName: "UpdateBanner",
			Handler:    _BannerService_UpdateBanner_Handler,
		},
		{
			MethodName: "CreateBannerSlotRelation",
			Handler:    _BannerService_CreateBannerSlotRelation_Handler,
		},
		{
			MethodName: "DeleteBannerSlotRelation",
			Handler:    _BannerService_DeleteBannerSlotRelation_Handler,
		},
		{
			MethodName: "RegisterClick",
			Handler:    _BannerService_RegisterClick_Handler,
		},
		{
			MethodName: "SelectBanner",
			Handler:    _BannerService_SelectBanner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/banner_service.proto",
}
