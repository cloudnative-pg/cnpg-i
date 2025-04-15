// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v6.30.1
// source: proto/postgres.proto

package postgres

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PostgresCapability_RPC_Type int32

const (
	PostgresCapability_RPC_TYPE_UNSPECIFIED PostgresCapability_RPC_Type = 0
	// TYPE_BACKUP indicates that the Plugin is able to
	// enrich the PostgreSQL configuration via the
	// EnrichConfiguration endpoint
	PostgresCapability_RPC_TYPE_ENRICH_CONFIGURATION PostgresCapability_RPC_Type = 1
)

// Enum value maps for PostgresCapability_RPC_Type.
var (
	PostgresCapability_RPC_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_ENRICH_CONFIGURATION",
	}
	PostgresCapability_RPC_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":          0,
		"TYPE_ENRICH_CONFIGURATION": 1,
	}
)

func (x PostgresCapability_RPC_Type) Enum() *PostgresCapability_RPC_Type {
	p := new(PostgresCapability_RPC_Type)
	*p = x
	return p
}

func (x PostgresCapability_RPC_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PostgresCapability_RPC_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_postgres_proto_enumTypes[0].Descriptor()
}

func (PostgresCapability_RPC_Type) Type() protoreflect.EnumType {
	return &file_proto_postgres_proto_enumTypes[0]
}

func (x PostgresCapability_RPC_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PostgresCapability_RPC_Type.Descriptor instead.
func (PostgresCapability_RPC_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_postgres_proto_rawDescGZIP(), []int{2, 0, 0}
}

type OperationType_Type int32

const (
	OperationType_TYPE_UNSPECIFIED OperationType_Type = 0
	OperationType_TYPE_INIT        OperationType_Type = 1
	OperationType_TYPE_RESTORE     OperationType_Type = 2
	OperationType_TYPE_RECONCILE   OperationType_Type = 3
	OperationType_TYPE_UPGRADE     OperationType_Type = 4
)

// Enum value maps for OperationType_Type.
var (
	OperationType_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_INIT",
		2: "TYPE_RESTORE",
		3: "TYPE_RECONCILE",
		4: "TYPE_UPGRADE",
	}
	OperationType_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_INIT":        1,
		"TYPE_RESTORE":     2,
		"TYPE_RECONCILE":   3,
		"TYPE_UPGRADE":     4,
	}
)

func (x OperationType_Type) Enum() *OperationType_Type {
	p := new(OperationType_Type)
	*p = x
	return p
}

func (x OperationType_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_postgres_proto_enumTypes[1].Descriptor()
}

func (OperationType_Type) Type() protoreflect.EnumType {
	return &file_proto_postgres_proto_enumTypes[1]
}

func (x OperationType_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType_Type.Descriptor instead.
func (OperationType_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_postgres_proto_rawDescGZIP(), []int{3, 0}
}

type PostgresCapabilitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostgresCapabilitiesRequest) Reset() {
	*x = PostgresCapabilitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_postgres_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostgresCapabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresCapabilitiesRequest) ProtoMessage() {}

func (x *PostgresCapabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_postgres_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresCapabilitiesRequest.ProtoReflect.Descriptor instead.
func (*PostgresCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_proto_postgres_proto_rawDescGZIP(), []int{0}
}

type PostgresCapabilitiesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All the capabilities that the controller service supports. This
	// field is OPTIONAL.
	Capabilities []*PostgresCapability `protobuf:"bytes,1,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
}

func (x *PostgresCapabilitiesResult) Reset() {
	*x = PostgresCapabilitiesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_postgres_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostgresCapabilitiesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresCapabilitiesResult) ProtoMessage() {}

func (x *PostgresCapabilitiesResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_postgres_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresCapabilitiesResult.ProtoReflect.Descriptor instead.
func (*PostgresCapabilitiesResult) Descriptor() ([]byte, []int) {
	return file_proto_postgres_proto_rawDescGZIP(), []int{1}
}

func (x *PostgresCapabilitiesResult) GetCapabilities() []*PostgresCapability {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

type PostgresCapability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*PostgresCapability_Rpc
	Type isPostgresCapability_Type `protobuf_oneof:"type"`
}

func (x *PostgresCapability) Reset() {
	*x = PostgresCapability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_postgres_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostgresCapability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresCapability) ProtoMessage() {}

func (x *PostgresCapability) ProtoReflect() protoreflect.Message {
	mi := &file_proto_postgres_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresCapability.ProtoReflect.Descriptor instead.
func (*PostgresCapability) Descriptor() ([]byte, []int) {
	return file_proto_postgres_proto_rawDescGZIP(), []int{2}
}

func (m *PostgresCapability) GetType() isPostgresCapability_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *PostgresCapability) GetRpc() *PostgresCapability_RPC {
	if x, ok := x.GetType().(*PostgresCapability_Rpc); ok {
		return x.Rpc
	}
	return nil
}

type isPostgresCapability_Type interface {
	isPostgresCapability_Type()
}

type PostgresCapability_Rpc struct {
	Rpc *PostgresCapability_RPC `protobuf:"bytes,1,opt,name=rpc,proto3,oneof"`
}

func (*PostgresCapability_Rpc) isPostgresCapability_Type() {}

type OperationType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type OperationType_Type `protobuf:"varint,1,opt,name=type,proto3,enum=cnpgi.identity.v1.OperationType_Type" json:"type,omitempty"`
}

func (x *OperationType) Reset() {
	*x = OperationType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_postgres_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationType) ProtoMessage() {}

func (x *OperationType) ProtoReflect() protoreflect.Message {
	mi := &file_proto_postgres_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationType.ProtoReflect.Descriptor instead.
func (*OperationType) Descriptor() ([]byte, []int) {
	return file_proto_postgres_proto_rawDescGZIP(), []int{3}
}

func (x *OperationType) GetType() OperationType_Type {
	if x != nil {
		return x.Type
	}
	return OperationType_TYPE_UNSPECIFIED
}

type EnrichConfigurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field is REQUIRED and represent the PostgreSQL configuration parameters as
	// generated by the instance manager
	Configs map[string]string `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// This field is REQUIRED
	ClusterDefinition []byte `protobuf:"bytes,2,opt,name=cluster_definition,json=clusterDefinition,proto3" json:"cluster_definition,omitempty"`
	// This field is REQUIRED.
	OperationType *OperationType `protobuf:"bytes,3,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
}

func (x *EnrichConfigurationRequest) Reset() {
	*x = EnrichConfigurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_postgres_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrichConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrichConfigurationRequest) ProtoMessage() {}

func (x *EnrichConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_postgres_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrichConfigurationRequest.ProtoReflect.Descriptor instead.
func (*EnrichConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_proto_postgres_proto_rawDescGZIP(), []int{4}
}

func (x *EnrichConfigurationRequest) GetConfigs() map[string]string {
	if x != nil {
		return x.Configs
	}
	return nil
}

func (x *EnrichConfigurationRequest) GetClusterDefinition() []byte {
	if x != nil {
		return x.ClusterDefinition
	}
	return nil
}

func (x *EnrichConfigurationRequest) GetOperationType() *OperationType {
	if x != nil {
		return x.OperationType
	}
	return nil
}

type EnrichConfigurationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field is OPTIONAL. Returns the new configuration.
	Configs map[string]string `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EnrichConfigurationResult) Reset() {
	*x = EnrichConfigurationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_postgres_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrichConfigurationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrichConfigurationResult) ProtoMessage() {}

func (x *EnrichConfigurationResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_postgres_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrichConfigurationResult.ProtoReflect.Descriptor instead.
func (*EnrichConfigurationResult) Descriptor() ([]byte, []int) {
	return file_proto_postgres_proto_rawDescGZIP(), []int{5}
}

func (x *EnrichConfigurationResult) GetConfigs() map[string]string {
	if x != nil {
		return x.Configs
	}
	return nil
}

type PostgresCapability_RPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type PostgresCapability_RPC_Type `protobuf:"varint,1,opt,name=type,proto3,enum=cnpgi.identity.v1.PostgresCapability_RPC_Type" json:"type,omitempty"`
}

func (x *PostgresCapability_RPC) Reset() {
	*x = PostgresCapability_RPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_postgres_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostgresCapability_RPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresCapability_RPC) ProtoMessage() {}

func (x *PostgresCapability_RPC) ProtoReflect() protoreflect.Message {
	mi := &file_proto_postgres_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresCapability_RPC.ProtoReflect.Descriptor instead.
func (*PostgresCapability_RPC) Descriptor() ([]byte, []int) {
	return file_proto_postgres_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PostgresCapability_RPC) GetType() PostgresCapability_RPC_Type {
	if x != nil {
		return x.Type
	}
	return PostgresCapability_RPC_TYPE_UNSPECIFIED
}

var File_proto_postgres_proto protoreflect.FileDescriptor

var file_proto_postgres_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x1d, 0x0a, 0x1b, 0x50, 0x6f, 0x73,
	0x74, 0x67, 0x72, 0x65, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x67, 0x0a, 0x1a, 0x50, 0x6f, 0x73, 0x74,
	0x67, 0x72, 0x65, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x49, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63,
	0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x22, 0xe4, 0x01, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x03, 0x72, 0x70, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x50, 0x43,
	0x48, 0x00, 0x52, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x86, 0x01, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12,
	0x42, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e,
	0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x50, 0x43, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x3b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1d, 0x0a, 0x19, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x52, 0x49, 0x43, 0x48,
	0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01,
	0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x0d, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x63, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x49, 0x54,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x4f,
	0x52, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x43,
	0x4f, 0x4e, 0x43, 0x49, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x10, 0x04, 0x22, 0xa6, 0x02, 0x0a, 0x1a, 0x45,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x63, 0x6e, 0x70,
	0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x2d, 0x0a, 0x12, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47,
	0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xac, 0x01, 0x0a, 0x19, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x53, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x32, 0xf4, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x12,
	0x72, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x2e, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x13, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x2e, 0x63, 0x6e, 0x70,
	0x67, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x72, 0x69, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6e, 0x70, 0x67,
	0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e,
	0x72, 0x69, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6e, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x2d, 0x70, 0x67, 0x2f, 0x63, 0x6e, 0x70, 0x67, 0x2d, 0x69, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_postgres_proto_rawDescOnce sync.Once
	file_proto_postgres_proto_rawDescData = file_proto_postgres_proto_rawDesc
)

func file_proto_postgres_proto_rawDescGZIP() []byte {
	file_proto_postgres_proto_rawDescOnce.Do(func() {
		file_proto_postgres_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_postgres_proto_rawDescData)
	})
	return file_proto_postgres_proto_rawDescData
}

var file_proto_postgres_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_postgres_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_postgres_proto_goTypes = []interface{}{
	(PostgresCapability_RPC_Type)(0),    // 0: cnpgi.identity.v1.PostgresCapability.RPC.Type
	(OperationType_Type)(0),             // 1: cnpgi.identity.v1.OperationType.Type
	(*PostgresCapabilitiesRequest)(nil), // 2: cnpgi.identity.v1.PostgresCapabilitiesRequest
	(*PostgresCapabilitiesResult)(nil),  // 3: cnpgi.identity.v1.PostgresCapabilitiesResult
	(*PostgresCapability)(nil),          // 4: cnpgi.identity.v1.PostgresCapability
	(*OperationType)(nil),               // 5: cnpgi.identity.v1.OperationType
	(*EnrichConfigurationRequest)(nil),  // 6: cnpgi.identity.v1.EnrichConfigurationRequest
	(*EnrichConfigurationResult)(nil),   // 7: cnpgi.identity.v1.EnrichConfigurationResult
	(*PostgresCapability_RPC)(nil),      // 8: cnpgi.identity.v1.PostgresCapability.RPC
	nil,                                 // 9: cnpgi.identity.v1.EnrichConfigurationRequest.ConfigsEntry
	nil,                                 // 10: cnpgi.identity.v1.EnrichConfigurationResult.ConfigsEntry
}
var file_proto_postgres_proto_depIdxs = []int32{
	4,  // 0: cnpgi.identity.v1.PostgresCapabilitiesResult.capabilities:type_name -> cnpgi.identity.v1.PostgresCapability
	8,  // 1: cnpgi.identity.v1.PostgresCapability.rpc:type_name -> cnpgi.identity.v1.PostgresCapability.RPC
	1,  // 2: cnpgi.identity.v1.OperationType.type:type_name -> cnpgi.identity.v1.OperationType.Type
	9,  // 3: cnpgi.identity.v1.EnrichConfigurationRequest.configs:type_name -> cnpgi.identity.v1.EnrichConfigurationRequest.ConfigsEntry
	5,  // 4: cnpgi.identity.v1.EnrichConfigurationRequest.operation_type:type_name -> cnpgi.identity.v1.OperationType
	10, // 5: cnpgi.identity.v1.EnrichConfigurationResult.configs:type_name -> cnpgi.identity.v1.EnrichConfigurationResult.ConfigsEntry
	0,  // 6: cnpgi.identity.v1.PostgresCapability.RPC.type:type_name -> cnpgi.identity.v1.PostgresCapability.RPC.Type
	2,  // 7: cnpgi.identity.v1.Postgres.GetCapabilities:input_type -> cnpgi.identity.v1.PostgresCapabilitiesRequest
	6,  // 8: cnpgi.identity.v1.Postgres.EnrichConfiguration:input_type -> cnpgi.identity.v1.EnrichConfigurationRequest
	3,  // 9: cnpgi.identity.v1.Postgres.GetCapabilities:output_type -> cnpgi.identity.v1.PostgresCapabilitiesResult
	7,  // 10: cnpgi.identity.v1.Postgres.EnrichConfiguration:output_type -> cnpgi.identity.v1.EnrichConfigurationResult
	9,  // [9:11] is the sub-list for method output_type
	7,  // [7:9] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_postgres_proto_init() }
func file_proto_postgres_proto_init() {
	if File_proto_postgres_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_postgres_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostgresCapabilitiesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_postgres_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostgresCapabilitiesResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_postgres_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostgresCapability); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_postgres_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_postgres_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrichConfigurationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_postgres_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrichConfigurationResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_postgres_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostgresCapability_RPC); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_proto_postgres_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PostgresCapability_Rpc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_postgres_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_postgres_proto_goTypes,
		DependencyIndexes: file_proto_postgres_proto_depIdxs,
		EnumInfos:         file_proto_postgres_proto_enumTypes,
		MessageInfos:      file_proto_postgres_proto_msgTypes,
	}.Build()
	File_proto_postgres_proto = out.File
	file_proto_postgres_proto_rawDesc = nil
	file_proto_postgres_proto_goTypes = nil
	file_proto_postgres_proto_depIdxs = nil
}
