// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.2
// source: proto/operator_lifecycle.proto

package lifecycle

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

// The operator type corresponds to the Kubernetes API method
type OperatorOperationType_Type int32

const (
	OperatorOperationType_TYPE_UNSPECIFIED OperatorOperationType_Type = 0
	OperatorOperationType_TYPE_PATCH       OperatorOperationType_Type = 1
	OperatorOperationType_TYPE_UPDATE      OperatorOperationType_Type = 2
	OperatorOperationType_TYPE_CREATE      OperatorOperationType_Type = 3
	OperatorOperationType_TYPE_DELETE      OperatorOperationType_Type = 4
	OperatorOperationType_TYPE_DEREGISTER  OperatorOperationType_Type = 5
)

// Enum value maps for OperatorOperationType_Type.
var (
	OperatorOperationType_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_PATCH",
		2: "TYPE_UPDATE",
		3: "TYPE_CREATE",
		4: "TYPE_DELETE",
		5: "TYPE_DEREGISTER",
	}
	OperatorOperationType_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_PATCH":       1,
		"TYPE_UPDATE":      2,
		"TYPE_CREATE":      3,
		"TYPE_DELETE":      4,
		"TYPE_DEREGISTER":  5,
	}
)

func (x OperatorOperationType_Type) Enum() *OperatorOperationType_Type {
	p := new(OperatorOperationType_Type)
	*p = x
	return p
}

func (x OperatorOperationType_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperatorOperationType_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_operator_lifecycle_proto_enumTypes[0].Descriptor()
}

func (OperatorOperationType_Type) Type() protoreflect.EnumType {
	return &file_proto_operator_lifecycle_proto_enumTypes[0]
}

func (x OperatorOperationType_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperatorOperationType_Type.Descriptor instead.
func (OperatorOperationType_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_operator_lifecycle_proto_rawDescGZIP(), []int{3, 0}
}

type OperatorLifecycleCapabilitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OperatorLifecycleCapabilitiesRequest) Reset() {
	*x = OperatorLifecycleCapabilitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_operator_lifecycle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorLifecycleCapabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorLifecycleCapabilitiesRequest) ProtoMessage() {}

func (x *OperatorLifecycleCapabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorLifecycleCapabilitiesRequest.ProtoReflect.Descriptor instead.
func (*OperatorLifecycleCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_proto_operator_lifecycle_proto_rawDescGZIP(), []int{0}
}

type OperatorLifecycleCapabilitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This message is OPTIONAL, containing the list of resources
	// for which the lifecycle hook is called
	LifecycleCapabilities []*OperatorLifecycleCapabilities `protobuf:"bytes,1,rep,name=lifecycle_capabilities,json=lifecycleCapabilities,proto3" json:"lifecycle_capabilities,omitempty"`
}

func (x *OperatorLifecycleCapabilitiesResponse) Reset() {
	*x = OperatorLifecycleCapabilitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_operator_lifecycle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorLifecycleCapabilitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorLifecycleCapabilitiesResponse) ProtoMessage() {}

func (x *OperatorLifecycleCapabilitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorLifecycleCapabilitiesResponse.ProtoReflect.Descriptor instead.
func (*OperatorLifecycleCapabilitiesResponse) Descriptor() ([]byte, []int) {
	return file_proto_operator_lifecycle_proto_rawDescGZIP(), []int{1}
}

func (x *OperatorLifecycleCapabilitiesResponse) GetLifecycleCapabilities() []*OperatorLifecycleCapabilities {
	if x != nil {
		return x.LifecycleCapabilities
	}
	return nil
}

type OperatorLifecycleCapabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Kubernetes resource group (such as "apps" or empty for core resources)
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// The Kubernetes Kind name (such as "Cluster" or "Pod")
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// The operation type
	OperationTypes []*OperatorOperationType `protobuf:"bytes,3,rep,name=operation_types,json=operationTypes,proto3" json:"operation_types,omitempty"`
}

func (x *OperatorLifecycleCapabilities) Reset() {
	*x = OperatorLifecycleCapabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_operator_lifecycle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorLifecycleCapabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorLifecycleCapabilities) ProtoMessage() {}

func (x *OperatorLifecycleCapabilities) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorLifecycleCapabilities.ProtoReflect.Descriptor instead.
func (*OperatorLifecycleCapabilities) Descriptor() ([]byte, []int) {
	return file_proto_operator_lifecycle_proto_rawDescGZIP(), []int{2}
}

func (x *OperatorLifecycleCapabilities) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *OperatorLifecycleCapabilities) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *OperatorLifecycleCapabilities) GetOperationTypes() []*OperatorOperationType {
	if x != nil {
		return x.OperationTypes
	}
	return nil
}

type OperatorOperationType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type OperatorOperationType_Type `protobuf:"varint,1,opt,name=type,proto3,enum=cnpgi.operator_lifecycle.v1.OperatorOperationType_Type" json:"type,omitempty"`
}

func (x *OperatorOperationType) Reset() {
	*x = OperatorOperationType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_operator_lifecycle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorOperationType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorOperationType) ProtoMessage() {}

func (x *OperatorOperationType) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorOperationType.ProtoReflect.Descriptor instead.
func (*OperatorOperationType) Descriptor() ([]byte, []int) {
	return file_proto_operator_lifecycle_proto_rawDescGZIP(), []int{3}
}

func (x *OperatorOperationType) GetType() OperatorOperationType_Type {
	if x != nil {
		return x.Type
	}
	return OperatorOperationType_TYPE_UNSPECIFIED
}

type OperatorLifecycleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field is REQUIRED.
	OperationType *OperatorOperationType `protobuf:"bytes,1,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	// This field is REQUIRED
	ClusterDefinition []byte `protobuf:"bytes,2,opt,name=cluster_definition,json=clusterDefinition,proto3" json:"cluster_definition,omitempty"`
	// This field is REQUIRED.
	ObjectDefinition []byte `protobuf:"bytes,3,opt,name=object_definition,json=objectDefinition,proto3" json:"object_definition,omitempty"`
}

func (x *OperatorLifecycleRequest) Reset() {
	*x = OperatorLifecycleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_operator_lifecycle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorLifecycleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorLifecycleRequest) ProtoMessage() {}

func (x *OperatorLifecycleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorLifecycleRequest.ProtoReflect.Descriptor instead.
func (*OperatorLifecycleRequest) Descriptor() ([]byte, []int) {
	return file_proto_operator_lifecycle_proto_rawDescGZIP(), []int{4}
}

func (x *OperatorLifecycleRequest) GetOperationType() *OperatorOperationType {
	if x != nil {
		return x.OperationType
	}
	return nil
}

func (x *OperatorLifecycleRequest) GetClusterDefinition() []byte {
	if x != nil {
		return x.ClusterDefinition
	}
	return nil
}

func (x *OperatorLifecycleRequest) GetObjectDefinition() []byte {
	if x != nil {
		return x.ObjectDefinition
	}
	return nil
}

type OperatorLifecycleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field is OPTIONAL.
	JsonPatch []byte `protobuf:"bytes,1,opt,name=json_patch,json=jsonPatch,proto3" json:"json_patch,omitempty"`
}

func (x *OperatorLifecycleResponse) Reset() {
	*x = OperatorLifecycleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_operator_lifecycle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperatorLifecycleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorLifecycleResponse) ProtoMessage() {}

func (x *OperatorLifecycleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorLifecycleResponse.ProtoReflect.Descriptor instead.
func (*OperatorLifecycleResponse) Descriptor() ([]byte, []int) {
	return file_proto_operator_lifecycle_proto_rawDescGZIP(), []int{5}
}

func (x *OperatorLifecycleResponse) GetJsonPatch() []byte {
	if x != nil {
		return x.JsonPatch
	}
	return nil
}

var File_proto_operator_lifecycle_proto protoreflect.FileDescriptor

var file_proto_operator_lifecycle_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x26, 0x0a,
	0x24, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9a, 0x01, 0x0a, 0x25, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x71, 0x0a, 0x16, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x5f, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3a, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x15, 0x6c, 0x69, 0x66,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x22, 0xa6, 0x01, 0x0a, 0x1d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x5b,
	0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0xda, 0x01, 0x0a, 0x15,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x74, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54,
	0x45, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x52, 0x45,
	0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x05, 0x22, 0xd1, 0x01, 0x0a, 0x18, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x59, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2b, 0x0a, 0x11, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x19,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x73, 0x6f,
	0x6e, 0x5f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6a,
	0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x32, 0xb3, 0x02, 0x0a, 0x11, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x9a,
	0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x41, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63,
	0x6c, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x80, 0x01, 0x0a, 0x0d,
	0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x35, 0x2e,
	0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6c,
	0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x66, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x70, 0x67, 0x2f, 0x63, 0x6e, 0x70, 0x67,
	0x2d, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x69, 0x66, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_operator_lifecycle_proto_rawDescOnce sync.Once
	file_proto_operator_lifecycle_proto_rawDescData = file_proto_operator_lifecycle_proto_rawDesc
)

func file_proto_operator_lifecycle_proto_rawDescGZIP() []byte {
	file_proto_operator_lifecycle_proto_rawDescOnce.Do(func() {
		file_proto_operator_lifecycle_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_operator_lifecycle_proto_rawDescData)
	})
	return file_proto_operator_lifecycle_proto_rawDescData
}

var file_proto_operator_lifecycle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_operator_lifecycle_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_operator_lifecycle_proto_goTypes = []interface{}{
	(OperatorOperationType_Type)(0),               // 0: cnpgi.operator_lifecycle.v1.OperatorOperationType.Type
	(*OperatorLifecycleCapabilitiesRequest)(nil),  // 1: cnpgi.operator_lifecycle.v1.OperatorLifecycleCapabilitiesRequest
	(*OperatorLifecycleCapabilitiesResponse)(nil), // 2: cnpgi.operator_lifecycle.v1.OperatorLifecycleCapabilitiesResponse
	(*OperatorLifecycleCapabilities)(nil),         // 3: cnpgi.operator_lifecycle.v1.OperatorLifecycleCapabilities
	(*OperatorOperationType)(nil),                 // 4: cnpgi.operator_lifecycle.v1.OperatorOperationType
	(*OperatorLifecycleRequest)(nil),              // 5: cnpgi.operator_lifecycle.v1.OperatorLifecycleRequest
	(*OperatorLifecycleResponse)(nil),             // 6: cnpgi.operator_lifecycle.v1.OperatorLifecycleResponse
}
var file_proto_operator_lifecycle_proto_depIdxs = []int32{
	3, // 0: cnpgi.operator_lifecycle.v1.OperatorLifecycleCapabilitiesResponse.lifecycle_capabilities:type_name -> cnpgi.operator_lifecycle.v1.OperatorLifecycleCapabilities
	4, // 1: cnpgi.operator_lifecycle.v1.OperatorLifecycleCapabilities.operation_types:type_name -> cnpgi.operator_lifecycle.v1.OperatorOperationType
	0, // 2: cnpgi.operator_lifecycle.v1.OperatorOperationType.type:type_name -> cnpgi.operator_lifecycle.v1.OperatorOperationType.Type
	4, // 3: cnpgi.operator_lifecycle.v1.OperatorLifecycleRequest.operation_type:type_name -> cnpgi.operator_lifecycle.v1.OperatorOperationType
	1, // 4: cnpgi.operator_lifecycle.v1.OperatorLifecycle.GetCapabilities:input_type -> cnpgi.operator_lifecycle.v1.OperatorLifecycleCapabilitiesRequest
	5, // 5: cnpgi.operator_lifecycle.v1.OperatorLifecycle.LifecycleHook:input_type -> cnpgi.operator_lifecycle.v1.OperatorLifecycleRequest
	2, // 6: cnpgi.operator_lifecycle.v1.OperatorLifecycle.GetCapabilities:output_type -> cnpgi.operator_lifecycle.v1.OperatorLifecycleCapabilitiesResponse
	6, // 7: cnpgi.operator_lifecycle.v1.OperatorLifecycle.LifecycleHook:output_type -> cnpgi.operator_lifecycle.v1.OperatorLifecycleResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_operator_lifecycle_proto_init() }
func file_proto_operator_lifecycle_proto_init() {
	if File_proto_operator_lifecycle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_operator_lifecycle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorLifecycleCapabilitiesRequest); i {
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
		file_proto_operator_lifecycle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorLifecycleCapabilitiesResponse); i {
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
		file_proto_operator_lifecycle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorLifecycleCapabilities); i {
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
		file_proto_operator_lifecycle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorOperationType); i {
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
		file_proto_operator_lifecycle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorLifecycleRequest); i {
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
		file_proto_operator_lifecycle_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperatorLifecycleResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_operator_lifecycle_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_operator_lifecycle_proto_goTypes,
		DependencyIndexes: file_proto_operator_lifecycle_proto_depIdxs,
		EnumInfos:         file_proto_operator_lifecycle_proto_enumTypes,
		MessageInfos:      file_proto_operator_lifecycle_proto_msgTypes,
	}.Build()
	File_proto_operator_lifecycle_proto = out.File
	file_proto_operator_lifecycle_proto_rawDesc = nil
	file_proto_operator_lifecycle_proto_goTypes = nil
	file_proto_operator_lifecycle_proto_depIdxs = nil
}
