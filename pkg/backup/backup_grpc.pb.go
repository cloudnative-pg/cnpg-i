// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.2
// source: proto/backup.proto

package backup

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
	Backup_GetCapabilities_FullMethodName = "/cnpgi.backup.v1.Backup/GetCapabilities"
	Backup_Backup_FullMethodName          = "/cnpgi.backup.v1.Backup/Backup"
)

// BackupClient is the client API for Backup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackupClient interface {
	// GetCapabilities gets the capabilities of the Backup service
	GetCapabilities(ctx context.Context, in *BackupCapabilitiesRequest, opts ...grpc.CallOption) (*BackupCapabilitiesResult, error)
	// Backup takes a physical backup of PostgreSQL.
	Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResult, error)
}

type backupClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupClient(cc grpc.ClientConnInterface) BackupClient {
	return &backupClient{cc}
}

func (c *backupClient) GetCapabilities(ctx context.Context, in *BackupCapabilitiesRequest, opts ...grpc.CallOption) (*BackupCapabilitiesResult, error) {
	out := new(BackupCapabilitiesResult)
	err := c.cc.Invoke(ctx, Backup_GetCapabilities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupClient) Backup(ctx context.Context, in *BackupRequest, opts ...grpc.CallOption) (*BackupResult, error) {
	out := new(BackupResult)
	err := c.cc.Invoke(ctx, Backup_Backup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServer is the server API for Backup service.
// All implementations must embed UnimplementedBackupServer
// for forward compatibility
type BackupServer interface {
	// GetCapabilities gets the capabilities of the Backup service
	GetCapabilities(context.Context, *BackupCapabilitiesRequest) (*BackupCapabilitiesResult, error)
	// Backup takes a physical backup of PostgreSQL.
	Backup(context.Context, *BackupRequest) (*BackupResult, error)
	mustEmbedUnimplementedBackupServer()
}

// UnimplementedBackupServer must be embedded to have forward compatible implementations.
type UnimplementedBackupServer struct {
}

func (UnimplementedBackupServer) GetCapabilities(context.Context, *BackupCapabilitiesRequest) (*BackupCapabilitiesResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCapabilities not implemented")
}
func (UnimplementedBackupServer) Backup(context.Context, *BackupRequest) (*BackupResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Backup not implemented")
}
func (UnimplementedBackupServer) mustEmbedUnimplementedBackupServer() {}

// UnsafeBackupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackupServer will
// result in compilation errors.
type UnsafeBackupServer interface {
	mustEmbedUnimplementedBackupServer()
}

func RegisterBackupServer(s grpc.ServiceRegistrar, srv BackupServer) {
	s.RegisterService(&Backup_ServiceDesc, srv)
}

func _Backup_GetCapabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupCapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).GetCapabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backup_GetCapabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).GetCapabilities(ctx, req.(*BackupCapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backup_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Backup_Backup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServer).Backup(ctx, req.(*BackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Backup_ServiceDesc is the grpc.ServiceDesc for Backup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Backup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cnpgi.backup.v1.Backup",
	HandlerType: (*BackupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCapabilities",
			Handler:    _Backup_GetCapabilities_Handler,
		},
		{
			MethodName: "Backup",
			Handler:    _Backup_Backup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/backup.proto",
}
