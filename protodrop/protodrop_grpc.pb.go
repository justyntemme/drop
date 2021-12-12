// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protodrop

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TimeSlotServiceClient is the client API for TimeSlotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TimeSlotServiceClient interface {
	GetTimeSlotById(ctx context.Context, in *GetTimeSlotByIdRequest, opts ...grpc.CallOption) (*TimeSlotResponse, error)
}

type timeSlotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimeSlotServiceClient(cc grpc.ClientConnInterface) TimeSlotServiceClient {
	return &timeSlotServiceClient{cc}
}

func (c *timeSlotServiceClient) GetTimeSlotById(ctx context.Context, in *GetTimeSlotByIdRequest, opts ...grpc.CallOption) (*TimeSlotResponse, error) {
	out := new(TimeSlotResponse)
	err := c.cc.Invoke(ctx, "/protodrop.TimeSlotService/GetTimeSlotById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeSlotServiceServer is the server API for TimeSlotService service.
// All implementations must embed UnimplementedTimeSlotServiceServer
// for forward compatibility
type TimeSlotServiceServer interface {
	GetTimeSlotById(context.Context, *GetTimeSlotByIdRequest) (*TimeSlotResponse, error)
	mustEmbedUnimplementedTimeSlotServiceServer()
}

// UnimplementedTimeSlotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTimeSlotServiceServer struct {
}

func (UnimplementedTimeSlotServiceServer) GetTimeSlotById(context.Context, *GetTimeSlotByIdRequest) (*TimeSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSlotById not implemented")
}
func (UnimplementedTimeSlotServiceServer) mustEmbedUnimplementedTimeSlotServiceServer() {}

// UnsafeTimeSlotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TimeSlotServiceServer will
// result in compilation errors.
type UnsafeTimeSlotServiceServer interface {
	mustEmbedUnimplementedTimeSlotServiceServer()
}

func RegisterTimeSlotServiceServer(s grpc.ServiceRegistrar, srv TimeSlotServiceServer) {
	s.RegisterService(&TimeSlotService_ServiceDesc, srv)
}

func _TimeSlotService_GetTimeSlotById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSlotByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeSlotServiceServer).GetTimeSlotById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protodrop.TimeSlotService/GetTimeSlotById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeSlotServiceServer).GetTimeSlotById(ctx, req.(*GetTimeSlotByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TimeSlotService_ServiceDesc is the grpc.ServiceDesc for TimeSlotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TimeSlotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protodrop.TimeSlotService",
	HandlerType: (*TimeSlotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimeSlotById",
			Handler:    _TimeSlotService_GetTimeSlotById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protodrop/protodrop.proto",
}
