// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: infraestructure/server/protos/devices/devices.proto

package devices_manager

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
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
	GetDeviceById(ctx context.Context, in *DeviceByIdRequest, opts ...grpc.CallOption) (*DeviceData, error)
	GetDevicesByLocation(ctx context.Context, in *LocationRequests, opts ...grpc.CallOption) (*DevicesDataResponse, error)
	CreateDevice(ctx context.Context, in *DeviceData, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteDeviceById(ctx context.Context, in *DeviceByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetStatusByLocation(ctx context.Context, in *LocationRequests, opts ...grpc.CallOption) (*StatusDataResponse, error)
}

type devicesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDevicesServiceClient(cc grpc.ClientConnInterface) DevicesServiceClient {
	return &devicesServiceClient{cc}
}

func (c *devicesServiceClient) GetDeviceById(ctx context.Context, in *DeviceByIdRequest, opts ...grpc.CallOption) (*DeviceData, error) {
	out := new(DeviceData)
	err := c.cc.Invoke(ctx, "/server.controllers.DevicesService/GetDeviceById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) GetDevicesByLocation(ctx context.Context, in *LocationRequests, opts ...grpc.CallOption) (*DevicesDataResponse, error) {
	out := new(DevicesDataResponse)
	err := c.cc.Invoke(ctx, "/server.controllers.DevicesService/GetDevicesByLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) CreateDevice(ctx context.Context, in *DeviceData, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/server.controllers.DevicesService/CreateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) DeleteDeviceById(ctx context.Context, in *DeviceByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/server.controllers.DevicesService/DeleteDeviceById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesServiceClient) GetStatusByLocation(ctx context.Context, in *LocationRequests, opts ...grpc.CallOption) (*StatusDataResponse, error) {
	out := new(StatusDataResponse)
	err := c.cc.Invoke(ctx, "/server.controllers.DevicesService/GetStatusByLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicesServiceServer is the server API for DevicesService service.
// All implementations must embed UnimplementedDevicesServiceServer
// for forward compatibility
type DevicesServiceServer interface {
	GetDeviceById(context.Context, *DeviceByIdRequest) (*DeviceData, error)
	GetDevicesByLocation(context.Context, *LocationRequests) (*DevicesDataResponse, error)
	CreateDevice(context.Context, *DeviceData) (*empty.Empty, error)
	DeleteDeviceById(context.Context, *DeviceByIdRequest) (*empty.Empty, error)
	GetStatusByLocation(context.Context, *LocationRequests) (*StatusDataResponse, error)
	mustEmbedUnimplementedDevicesServiceServer()
}

// UnimplementedDevicesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDevicesServiceServer struct {
}

func (UnimplementedDevicesServiceServer) GetDeviceById(context.Context, *DeviceByIdRequest) (*DeviceData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceById not implemented")
}
func (UnimplementedDevicesServiceServer) GetDevicesByLocation(context.Context, *LocationRequests) (*DevicesDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevicesByLocation not implemented")
}
func (UnimplementedDevicesServiceServer) CreateDevice(context.Context, *DeviceData) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedDevicesServiceServer) DeleteDeviceById(context.Context, *DeviceByIdRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeviceById not implemented")
}
func (UnimplementedDevicesServiceServer) GetStatusByLocation(context.Context, *LocationRequests) (*StatusDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatusByLocation not implemented")
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

func _DevicesService_GetDeviceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).GetDeviceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.controllers.DevicesService/GetDeviceById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).GetDeviceById(ctx, req.(*DeviceByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_GetDevicesByLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).GetDevicesByLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.controllers.DevicesService/GetDevicesByLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).GetDevicesByLocation(ctx, req.(*LocationRequests))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.controllers.DevicesService/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).CreateDevice(ctx, req.(*DeviceData))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_DeleteDeviceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).DeleteDeviceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.controllers.DevicesService/DeleteDeviceById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).DeleteDeviceById(ctx, req.(*DeviceByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevicesService_GetStatusByLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServiceServer).GetStatusByLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.controllers.DevicesService/GetStatusByLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServiceServer).GetStatusByLocation(ctx, req.(*LocationRequests))
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
			MethodName: "GetDeviceById",
			Handler:    _DevicesService_GetDeviceById_Handler,
		},
		{
			MethodName: "GetDevicesByLocation",
			Handler:    _DevicesService_GetDevicesByLocation_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _DevicesService_CreateDevice_Handler,
		},
		{
			MethodName: "DeleteDeviceById",
			Handler:    _DevicesService_DeleteDeviceById_Handler,
		},
		{
			MethodName: "GetStatusByLocation",
			Handler:    _DevicesService_GetStatusByLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "infraestructure/server/protos/devices/devices.proto",
}
