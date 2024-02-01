// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0--rc1
// source: proto/lifecycle.proto

package lifecycle

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
	Lifecycle_GetCapabilities_FullMethodName = "/cnpgi.identity.v1.Lifecycle/GetCapabilities"
	Lifecycle_LifecycleHook_FullMethodName   = "/cnpgi.identity.v1.Lifecycle/LifecycleHook"
)

// LifecycleClient is the client API for Lifecycle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LifecycleClient interface {
	// GetCapabilities gets the capabilities of the Lifecycle service
	GetCapabilities(ctx context.Context, in *LifecycleCapabilitiesRequest, opts ...grpc.CallOption) (*LifecycleCapabilitiesResponse, error)
	// LifecycleHook allows the plugin to manipulate the Kubernetes resources
	// before the CNPG operator applies them to the Kubernetes cluster.
	LifecycleHook(ctx context.Context, in *LifecycleRequest, opts ...grpc.CallOption) (*LifecycleResponse, error)
}

type lifecycleClient struct {
	cc grpc.ClientConnInterface
}

func NewLifecycleClient(cc grpc.ClientConnInterface) LifecycleClient {
	return &lifecycleClient{cc}
}

func (c *lifecycleClient) GetCapabilities(ctx context.Context, in *LifecycleCapabilitiesRequest, opts ...grpc.CallOption) (*LifecycleCapabilitiesResponse, error) {
	out := new(LifecycleCapabilitiesResponse)
	err := c.cc.Invoke(ctx, Lifecycle_GetCapabilities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lifecycleClient) LifecycleHook(ctx context.Context, in *LifecycleRequest, opts ...grpc.CallOption) (*LifecycleResponse, error) {
	out := new(LifecycleResponse)
	err := c.cc.Invoke(ctx, Lifecycle_LifecycleHook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LifecycleServer is the server API for Lifecycle service.
// All implementations must embed UnimplementedLifecycleServer
// for forward compatibility
type LifecycleServer interface {
	// GetCapabilities gets the capabilities of the Lifecycle service
	GetCapabilities(context.Context, *LifecycleCapabilitiesRequest) (*LifecycleCapabilitiesResponse, error)
	// LifecycleHook allows the plugin to manipulate the Kubernetes resources
	// before the CNPG operator applies them to the Kubernetes cluster.
	LifecycleHook(context.Context, *LifecycleRequest) (*LifecycleResponse, error)
	mustEmbedUnimplementedLifecycleServer()
}

// UnimplementedLifecycleServer must be embedded to have forward compatible implementations.
type UnimplementedLifecycleServer struct {
}

func (UnimplementedLifecycleServer) GetCapabilities(context.Context, *LifecycleCapabilitiesRequest) (*LifecycleCapabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCapabilities not implemented")
}
func (UnimplementedLifecycleServer) LifecycleHook(context.Context, *LifecycleRequest) (*LifecycleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LifecycleHook not implemented")
}
func (UnimplementedLifecycleServer) mustEmbedUnimplementedLifecycleServer() {}

// UnsafeLifecycleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LifecycleServer will
// result in compilation errors.
type UnsafeLifecycleServer interface {
	mustEmbedUnimplementedLifecycleServer()
}

func RegisterLifecycleServer(s grpc.ServiceRegistrar, srv LifecycleServer) {
	s.RegisterService(&Lifecycle_ServiceDesc, srv)
}

func _Lifecycle_GetCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LifecycleCapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecycleServer).GetCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lifecycle_GetCapabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecycleServer).GetCapabilities(ctx, req.(*LifecycleCapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lifecycle_LifecycleHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LifecycleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LifecycleServer).LifecycleHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lifecycle_LifecycleHook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LifecycleServer).LifecycleHook(ctx, req.(*LifecycleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Lifecycle_ServiceDesc is the grpc.ServiceDesc for Lifecycle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lifecycle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cnpgi.identity.v1.Lifecycle",
	HandlerType: (*LifecycleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCapabilities",
			Handler:    _Lifecycle_GetCapabilities_Handler,
		},
		{
			MethodName: "LifecycleHook",
			Handler:    _Lifecycle_LifecycleHook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/lifecycle.proto",
}
