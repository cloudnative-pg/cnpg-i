// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: proto/restore_job.proto

package job

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
	RestoreJobHooks_GetCapabilities_FullMethodName = "/cnpgi.wal.v1.RestoreJobHooks/GetCapabilities"
	RestoreJobHooks_Restore_FullMethodName         = "/cnpgi.wal.v1.RestoreJobHooks/Restore"
)

// RestoreJobHooksClient is the client API for RestoreJobHooks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestoreJobHooksClient interface {
	// GetCapabilities gets the capabilities of the Backup service
	GetCapabilities(ctx context.Context, in *RestoreJobHooksCapabilitiesRequest, opts ...grpc.CallOption) (*RestoreJobHooksCapabilitiesResult, error)
	// Restore is called to restore a PGDATA
	Restore(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (*RestoreResponse, error)
}

type restoreJobHooksClient struct {
	cc grpc.ClientConnInterface
}

func NewRestoreJobHooksClient(cc grpc.ClientConnInterface) RestoreJobHooksClient {
	return &restoreJobHooksClient{cc}
}

func (c *restoreJobHooksClient) GetCapabilities(ctx context.Context, in *RestoreJobHooksCapabilitiesRequest, opts ...grpc.CallOption) (*RestoreJobHooksCapabilitiesResult, error) {
	out := new(RestoreJobHooksCapabilitiesResult)
	err := c.cc.Invoke(ctx, RestoreJobHooks_GetCapabilities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restoreJobHooksClient) Restore(ctx context.Context, in *RestoreRequest, opts ...grpc.CallOption) (*RestoreResponse, error) {
	out := new(RestoreResponse)
	err := c.cc.Invoke(ctx, RestoreJobHooks_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestoreJobHooksServer is the server API for RestoreJobHooks service.
// All implementations must embed UnimplementedRestoreJobHooksServer
// for forward compatibility
type RestoreJobHooksServer interface {
	// GetCapabilities gets the capabilities of the Backup service
	GetCapabilities(context.Context, *RestoreJobHooksCapabilitiesRequest) (*RestoreJobHooksCapabilitiesResult, error)
	// Restore is called to restore a PGDATA
	Restore(context.Context, *RestoreRequest) (*RestoreResponse, error)
	mustEmbedUnimplementedRestoreJobHooksServer()
}

// UnimplementedRestoreJobHooksServer must be embedded to have forward compatible implementations.
type UnimplementedRestoreJobHooksServer struct {
}

func (UnimplementedRestoreJobHooksServer) GetCapabilities(context.Context, *RestoreJobHooksCapabilitiesRequest) (*RestoreJobHooksCapabilitiesResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCapabilities not implemented")
}
func (UnimplementedRestoreJobHooksServer) Restore(context.Context, *RestoreRequest) (*RestoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedRestoreJobHooksServer) mustEmbedUnimplementedRestoreJobHooksServer() {}

// UnsafeRestoreJobHooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestoreJobHooksServer will
// result in compilation errors.
type UnsafeRestoreJobHooksServer interface {
	mustEmbedUnimplementedRestoreJobHooksServer()
}

func RegisterRestoreJobHooksServer(s grpc.ServiceRegistrar, srv RestoreJobHooksServer) {
	s.RegisterService(&RestoreJobHooks_ServiceDesc, srv)
}

func _RestoreJobHooks_GetCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreJobHooksCapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestoreJobHooksServer).GetCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestoreJobHooks_GetCapabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestoreJobHooksServer).GetCapabilities(ctx, req.(*RestoreJobHooksCapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RestoreJobHooks_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestoreJobHooksServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RestoreJobHooks_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestoreJobHooksServer).Restore(ctx, req.(*RestoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RestoreJobHooks_ServiceDesc is the grpc.ServiceDesc for RestoreJobHooks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestoreJobHooks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cnpgi.wal.v1.RestoreJobHooks",
	HandlerType: (*RestoreJobHooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCapabilities",
			Handler:    _RestoreJobHooks_GetCapabilities_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _RestoreJobHooks_Restore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/restore_job.proto",
}
