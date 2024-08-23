// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: infraestructure/server/protos/devices/devices.proto

package devices_manager

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

// DevicesServiceClient is the client API for DevicesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevicesServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type devicesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDevicesServiceClient(cc grpc.ClientConnInterface) DevicesServiceClient {
	return &devicesServiceClient{cc}
}

func (c *devicesServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/server.controllers.DevicesService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicesServiceServer is the server API for DevicesService service.
// All implementations must embed UnimplementedDevicesServiceServer
// for forward compatibility
type DevicesServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedDevicesServiceServer()
}

// UnimplementedDevicesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDevicesServiceServer struct {
}

func (UnimplementedDevicesServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDevicesServiceServer) mustEmbedUnimplementedDevicesServiceServer() {}

// UnsafeDevicesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevicesServiceServer will
// result in compilation errors.
type UnsafeDevicesServiceServer interface {
	mustEmbedUnimplementedDevicesServiceServer()
}

func RegisterDevicesServiceServer(s grpc.ServiceRegistrar, srv DevicesServiceServer) {
	s.RegisterService(&DevicesService_ServiceDesc, srv)
}

func _DevicesService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.controllers.DevicesService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DevicesService_ServiceDesc is the grpc.ServiceDesc for DevicesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DevicesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.controllers.DevicesService",
	HandlerType: (*DevicesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _DevicesService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infraestructure/server/protos/devices/devices.proto",
}
