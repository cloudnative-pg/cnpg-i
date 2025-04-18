// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: proto/reconciler.proto

package reconciler

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ReconcilerHooks_GetCapabilities_FullMethodName = "/cnpgi.reconciler.v1.ReconcilerHooks/GetCapabilities"
	ReconcilerHooks_Pre_FullMethodName             = "/cnpgi.reconciler.v1.ReconcilerHooks/Pre"
	ReconcilerHooks_Post_FullMethodName            = "/cnpgi.reconciler.v1.ReconcilerHooks/Post"
)

// ReconcilerHooksClient is the client API for ReconcilerHooks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ReconcilerHooks offers a way for the plugins to directly execute changes on the resources
// through the kube-api server.
type ReconcilerHooksClient interface {
	// GetCapabilities gets the capabilities of the Backup service
	GetCapabilities(ctx context.Context, in *ReconcilerHooksCapabilitiesRequest, opts ...grpc.CallOption) (*ReconcilerHooksCapabilitiesResult, error)
	// Pre is executed before the operator executes the reconciliation loop
	// It is a way for the plugins to directly execute changes on the resources
	// through the kube-api server.
	Pre(ctx context.Context, in *ReconcilerHooksRequest, opts ...grpc.CallOption) (*ReconcilerHooksResult, error)
	// Post is executed after the operator executes the reconciliation loop
	// It is a way for the plugins to directly execute changes on the resources
	// through the kube-api server.
	Post(ctx context.Context, in *ReconcilerHooksRequest, opts ...grpc.CallOption) (*ReconcilerHooksResult, error)
}

type reconcilerHooksClient struct {
	cc grpc.ClientConnInterface
}

func NewReconcilerHooksClient(cc grpc.ClientConnInterface) ReconcilerHooksClient {
	return &reconcilerHooksClient{cc}
}

func (c *reconcilerHooksClient) GetCapabilities(ctx context.Context, in *ReconcilerHooksCapabilitiesRequest, opts ...grpc.CallOption) (*ReconcilerHooksCapabilitiesResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReconcilerHooksCapabilitiesResult)
	err := c.cc.Invoke(ctx, ReconcilerHooks_GetCapabilities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reconcilerHooksClient) Pre(ctx context.Context, in *ReconcilerHooksRequest, opts ...grpc.CallOption) (*ReconcilerHooksResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReconcilerHooksResult)
	err := c.cc.Invoke(ctx, ReconcilerHooks_Pre_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reconcilerHooksClient) Post(ctx context.Context, in *ReconcilerHooksRequest, opts ...grpc.CallOption) (*ReconcilerHooksResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReconcilerHooksResult)
	err := c.cc.Invoke(ctx, ReconcilerHooks_Post_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReconcilerHooksServer is the server API for ReconcilerHooks service.
// All implementations must embed UnimplementedReconcilerHooksServer
// for forward compatibility.
//
// ReconcilerHooks offers a way for the plugins to directly execute changes on the resources
// through the kube-api server.
type ReconcilerHooksServer interface {
	// GetCapabilities gets the capabilities of the Backup service
	GetCapabilities(context.Context, *ReconcilerHooksCapabilitiesRequest) (*ReconcilerHooksCapabilitiesResult, error)
	// Pre is executed before the operator executes the reconciliation loop
	// It is a way for the plugins to directly execute changes on the resources
	// through the kube-api server.
	Pre(context.Context, *ReconcilerHooksRequest) (*ReconcilerHooksResult, error)
	// Post is executed after the operator executes the reconciliation loop
	// It is a way for the plugins to directly execute changes on the resources
	// through the kube-api server.
	Post(context.Context, *ReconcilerHooksRequest) (*ReconcilerHooksResult, error)
	mustEmbedUnimplementedReconcilerHooksServer()
}

// UnimplementedReconcilerHooksServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedReconcilerHooksServer struct{}

func (UnimplementedReconcilerHooksServer) GetCapabilities(context.Context, *ReconcilerHooksCapabilitiesRequest) (*ReconcilerHooksCapabilitiesResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCapabilities not implemented")
}
func (UnimplementedReconcilerHooksServer) Pre(context.Context, *ReconcilerHooksRequest) (*ReconcilerHooksResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pre not implemented")
}
func (UnimplementedReconcilerHooksServer) Post(context.Context, *ReconcilerHooksRequest) (*ReconcilerHooksResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedReconcilerHooksServer) mustEmbedUnimplementedReconcilerHooksServer() {}
func (UnimplementedReconcilerHooksServer) testEmbeddedByValue()                         {}

// UnsafeReconcilerHooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReconcilerHooksServer will
// result in compilation errors.
type UnsafeReconcilerHooksServer interface {
	mustEmbedUnimplementedReconcilerHooksServer()
}

func RegisterReconcilerHooksServer(s grpc.ServiceRegistrar, srv ReconcilerHooksServer) {
	// If the following call pancis, it indicates UnimplementedReconcilerHooksServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ReconcilerHooks_ServiceDesc, srv)
}

func _ReconcilerHooks_GetCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReconcilerHooksCapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReconcilerHooksServer).GetCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReconcilerHooks_GetCapabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReconcilerHooksServer).GetCapabilities(ctx, req.(*ReconcilerHooksCapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReconcilerHooks_Pre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReconcilerHooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReconcilerHooksServer).Pre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReconcilerHooks_Pre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReconcilerHooksServer).Pre(ctx, req.(*ReconcilerHooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReconcilerHooks_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReconcilerHooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReconcilerHooksServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReconcilerHooks_Post_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReconcilerHooksServer).Post(ctx, req.(*ReconcilerHooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReconcilerHooks_ServiceDesc is the grpc.ServiceDesc for ReconcilerHooks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReconcilerHooks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cnpgi.reconciler.v1.ReconcilerHooks",
	HandlerType: (*ReconcilerHooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCapabilities",
			Handler:    _ReconcilerHooks_GetCapabilities_Handler,
		},
		{
			MethodName: "Pre",
			Handler:    _ReconcilerHooks_Pre_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _ReconcilerHooks_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/reconciler.proto",
}
