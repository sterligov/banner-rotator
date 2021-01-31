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

// SlotServiceClient is the client API for SlotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SlotServiceClient interface {
	FindSlotByID(ctx context.Context, in *FindSlotByIDRequest, opts ...grpc.CallOption) (*FindSlotByIDResponse, error)
	FindAllSlots(ctx context.Context, in *FindAllSlotsRequest, opts ...grpc.CallOption) (*FindAllSlotsResponse, error)
	CreateSlot(ctx context.Context, in *CreateSlotRequest, opts ...grpc.CallOption) (*CreateSlotResponse, error)
	DeleteSlot(ctx context.Context, in *DeleteSlotRequest, opts ...grpc.CallOption) (*DeleteSlotResponse, error)
	UpdateSlot(ctx context.Context, in *UpdateSlotRequest, opts ...grpc.CallOption) (*UpdateSlotResponse, error)
}

type slotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSlotServiceClient(cc grpc.ClientConnInterface) SlotServiceClient {
	return &slotServiceClient{cc}
}

func (c *slotServiceClient) FindSlotByID(ctx context.Context, in *FindSlotByIDRequest, opts ...grpc.CallOption) (*FindSlotByIDResponse, error) {
	out := new(FindSlotByIDResponse)
	err := c.cc.Invoke(ctx, "/slot.SlotService/FindSlotByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) FindAllSlots(ctx context.Context, in *FindAllSlotsRequest, opts ...grpc.CallOption) (*FindAllSlotsResponse, error) {
	out := new(FindAllSlotsResponse)
	err := c.cc.Invoke(ctx, "/slot.SlotService/FindAllSlots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) CreateSlot(ctx context.Context, in *CreateSlotRequest, opts ...grpc.CallOption) (*CreateSlotResponse, error) {
	out := new(CreateSlotResponse)
	err := c.cc.Invoke(ctx, "/slot.SlotService/CreateSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) DeleteSlot(ctx context.Context, in *DeleteSlotRequest, opts ...grpc.CallOption) (*DeleteSlotResponse, error) {
	out := new(DeleteSlotResponse)
	err := c.cc.Invoke(ctx, "/slot.SlotService/DeleteSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slotServiceClient) UpdateSlot(ctx context.Context, in *UpdateSlotRequest, opts ...grpc.CallOption) (*UpdateSlotResponse, error) {
	out := new(UpdateSlotResponse)
	err := c.cc.Invoke(ctx, "/slot.SlotService/UpdateSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SlotServiceServer is the server API for SlotService service.
// All implementations must embed UnimplementedSlotServiceServer
// for forward compatibility
type SlotServiceServer interface {
	FindSlotByID(context.Context, *FindSlotByIDRequest) (*FindSlotByIDResponse, error)
	FindAllSlots(context.Context, *FindAllSlotsRequest) (*FindAllSlotsResponse, error)
	CreateSlot(context.Context, *CreateSlotRequest) (*CreateSlotResponse, error)
	DeleteSlot(context.Context, *DeleteSlotRequest) (*DeleteSlotResponse, error)
	UpdateSlot(context.Context, *UpdateSlotRequest) (*UpdateSlotResponse, error)
	mustEmbedUnimplementedSlotServiceServer()
}

// UnimplementedSlotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSlotServiceServer struct {
}

func (UnimplementedSlotServiceServer) FindSlotByID(context.Context, *FindSlotByIDRequest) (*FindSlotByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSlotByID not implemented")
}
func (UnimplementedSlotServiceServer) FindAllSlots(context.Context, *FindAllSlotsRequest) (*FindAllSlotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllSlots not implemented")
}
func (UnimplementedSlotServiceServer) CreateSlot(context.Context, *CreateSlotRequest) (*CreateSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSlot not implemented")
}
func (UnimplementedSlotServiceServer) DeleteSlot(context.Context, *DeleteSlotRequest) (*DeleteSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSlot not implemented")
}
func (UnimplementedSlotServiceServer) UpdateSlot(context.Context, *UpdateSlotRequest) (*UpdateSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSlot not implemented")
}
func (UnimplementedSlotServiceServer) mustEmbedUnimplementedSlotServiceServer() {}

// UnsafeSlotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SlotServiceServer will
// result in compilation errors.
type UnsafeSlotServiceServer interface {
	mustEmbedUnimplementedSlotServiceServer()
}

func RegisterSlotServiceServer(s grpc.ServiceRegistrar, srv SlotServiceServer) {
	s.RegisterService(&_SlotService_serviceDesc, srv)
}

func _SlotService_FindSlotByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSlotByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).FindSlotByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slot.SlotService/FindSlotByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).FindSlotByID(ctx, req.(*FindSlotByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_FindAllSlots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllSlotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).FindAllSlots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slot.SlotService/FindAllSlots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).FindAllSlots(ctx, req.(*FindAllSlotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_CreateSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).CreateSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slot.SlotService/CreateSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).CreateSlot(ctx, req.(*CreateSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_DeleteSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).DeleteSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slot.SlotService/DeleteSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).DeleteSlot(ctx, req.(*DeleteSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlotService_UpdateSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlotServiceServer).UpdateSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slot.SlotService/UpdateSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlotServiceServer).UpdateSlot(ctx, req.(*UpdateSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SlotService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slot.SlotService",
	HandlerType: (*SlotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindSlotByID",
			Handler:    _SlotService_FindSlotByID_Handler,
		},
		{
			MethodName: "FindAllSlots",
			Handler:    _SlotService_FindAllSlots_Handler,
		},
		{
			MethodName: "CreateSlot",
			Handler:    _SlotService_CreateSlot_Handler,
		},
		{
			MethodName: "DeleteSlot",
			Handler:    _SlotService_DeleteSlot_Handler,
		},
		{
			MethodName: "UpdateSlot",
			Handler:    _SlotService_UpdateSlot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/slot_service.proto",
}
