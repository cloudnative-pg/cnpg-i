// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: proto/wal.proto

package wal

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

const (
	WAL_GetCapabilities_FullMethodName  = "/cnpgi.wal.v1.WAL/GetCapabilities"
	WAL_Archive_FullMethodName          = "/cnpgi.wal.v1.WAL/Archive"
	WAL_Restore_FullMethodName          = "/cnpgi.wal.v1.WAL/Restore"
	WAL_Status_FullMethodName           = "/cnpgi.wal.v1.WAL/Status"
	WAL_SetFirstRequired_FullMethodName = "/cnpgi.wal.v1.WAL/SetFirstRequired"
)

// WALClient is the client API for WAL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WALClient interface {
	// GetCapabilities gets the capabilities of the WAL service
	GetCapabilities(ctx context.Context, in *WALCapabilitiesRequest, opts ...grpc.CallOption) (*WALCapabilitiesResult, error)
	// Archive copies one WAL file into the archive
	Archive(ctx context.Context, in *WALArchiveRequest, opts ...grpc.CallOption) (*WALArchiveResult, error)
	// Restores copies WAL file from the archive to the data directory
	Restore(ctx context.Context, in *WALRestoreRequest, opts ...grpc.CallOption) (*WALRestoreResult, error)
	// Status gets the statistics of the WAL file archive
	Status(ctx context.Context, in *WALStatusRequest, opts ...grpc.CallOption) (*WALStatusResult, error)
	// SetFirstRequired sets the first required WAL for the cluster
	SetFirstRequired(ctx context.Context, in *SetFirstRequiredRequest, opts ...grpc.CallOption) (*SetFirstRequiredResult, error)
}

type wALClient struct {
	cc grpc.ClientConnInterface
}

func NewWALClient(cc grpc.ClientConnInterface) WALClient {
	return &wALClient{cc}
}

func (c *wALClient) GetCapabilities(ctx context.Context, in *WALCapabilitiesRequest, opts ...grpc.CallOption) (*WALCapabilitiesResult, error) {
	out := new(WALCapabilitiesResult)
	err := c.cc.Invoke(ctx, WAL_GetCapabilities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wALClient) Archive(ctx context.Context, in *WALArchiveRequest, opts ...grpc.CallOption) (*WALArchiveResult, error) {
	out := new(WALArchiveResult)
	err := c.cc.Invoke(ctx, WAL_Archive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wALClient) Restore(ctx context.Context, in *WALRestoreRequest, opts ...grpc.CallOption) (*WALRestoreResult, error) {
	out := new(WALRestoreResult)
	err := c.cc.Invoke(ctx, WAL_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wALClient) Status(ctx context.Context, in *WALStatusRequest, opts ...grpc.CallOption) (*WALStatusResult, error) {
	out := new(WALStatusResult)
	err := c.cc.Invoke(ctx, WAL_Status_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wALClient) SetFirstRequired(ctx context.Context, in *SetFirstRequiredRequest, opts ...grpc.CallOption) (*SetFirstRequiredResult, error) {
	out := new(SetFirstRequiredResult)
	err := c.cc.Invoke(ctx, WAL_SetFirstRequired_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WALServer is the server API for WAL service.
// All implementations must embed UnimplementedWALServer
// for forward compatibility
type WALServer interface {
	// GetCapabilities gets the capabilities of the WAL service
	GetCapabilities(context.Context, *WALCapabilitiesRequest) (*WALCapabilitiesResult, error)
	// Archive copies one WAL file into the archive
	Archive(context.Context, *WALArchiveRequest) (*WALArchiveResult, error)
	// Restores copies WAL file from the archive to the data directory
	Restore(context.Context, *WALRestoreRequest) (*WALRestoreResult, error)
	// Status gets the statistics of the WAL file archive
	Status(context.Context, *WALStatusRequest) (*WALStatusResult, error)
	// SetFirstRequired sets the first required WAL for the cluster
	SetFirstRequired(context.Context, *SetFirstRequiredRequest) (*SetFirstRequiredResult, error)
	mustEmbedUnimplementedWALServer()
}

// UnimplementedWALServer must be embedded to have forward compatible implementations.
type UnimplementedWALServer struct {
}

func (UnimplementedWALServer) GetCapabilities(context.Context, *WALCapabilitiesRequest) (*WALCapabilitiesResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCapabilities not implemented")
}
func (UnimplementedWALServer) Archive(context.Context, *WALArchiveRequest) (*WALArchiveResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Archive not implemented")
}
func (UnimplementedWALServer) Restore(context.Context, *WALRestoreRequest) (*WALRestoreResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedWALServer) Status(context.Context, *WALStatusRequest) (*WALStatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedWALServer) SetFirstRequired(context.Context, *SetFirstRequiredRequest) (*SetFirstRequiredResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFirstRequired not implemented")
}
func (UnimplementedWALServer) mustEmbedUnimplementedWALServer() {}

// UnsafeWALServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WALServer will
// result in compilation errors.
type UnsafeWALServer interface {
	mustEmbedUnimplementedWALServer()
}

func RegisterWALServer(s grpc.ServiceRegistrar, srv WALServer) {
	s.RegisterService(&WAL_ServiceDesc, srv)
}

func _WAL_GetCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WALCapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WALServer).GetCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WAL_GetCapabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WALServer).GetCapabilities(ctx, req.(*WALCapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WAL_Archive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WALArchiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WALServer).Archive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WAL_Archive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WALServer).Archive(ctx, req.(*WALArchiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WAL_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WALRestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WALServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WAL_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WALServer).Restore(ctx, req.(*WALRestoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WAL_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WALStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WALServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WAL_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WALServer).Status(ctx, req.(*WALStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WAL_SetFirstRequired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFirstRequiredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WALServer).SetFirstRequired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WAL_SetFirstRequired_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WALServer).SetFirstRequired(ctx, req.(*SetFirstRequiredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WAL_ServiceDesc is the grpc.ServiceDesc for WAL service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WAL_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cnpgi.wal.v1.WAL",
	HandlerType: (*WALServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCapabilities",
			Handler:    _WAL_GetCapabilities_Handler,
		},
		{
			MethodName: "Archive",
			Handler:    _WAL_Archive_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _WAL_Restore_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _WAL_Status_Handler,
		},
		{
			MethodName: "SetFirstRequired",
			Handler:    _WAL_SetFirstRequired_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/wal.proto",
}
