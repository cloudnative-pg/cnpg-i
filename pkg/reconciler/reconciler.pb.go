// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.2
// source: proto/reconciler.proto

package reconciler

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

type ReconcilerHooksCapability_Kind int32

const (
	ReconcilerHooksCapability_KIND_UNSPECIFIED ReconcilerHooksCapability_Kind = 0
	// KIND_CLUSTER indicates that the plugin will plug the
	// Cluster reconciler
	ReconcilerHooksCapability_KIND_CLUSTER ReconcilerHooksCapability_Kind = 1
	// KIND_BACKUP indicates that the plugin will plug the
	// Backup reconciler
	ReconcilerHooksCapability_KIND_BACKUP ReconcilerHooksCapability_Kind = 2
)

// Enum value maps for ReconcilerHooksCapability_Kind.
var (
	ReconcilerHooksCapability_Kind_name = map[int32]string{
		0: "KIND_UNSPECIFIED",
		1: "KIND_CLUSTER",
		2: "KIND_BACKUP",
	}
	ReconcilerHooksCapability_Kind_value = map[string]int32{
		"KIND_UNSPECIFIED": 0,
		"KIND_CLUSTER":     1,
		"KIND_BACKUP":      2,
	}
)

func (x ReconcilerHooksCapability_Kind) Enum() *ReconcilerHooksCapability_Kind {
	p := new(ReconcilerHooksCapability_Kind)
	*p = x
	return p
}

func (x ReconcilerHooksCapability_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReconcilerHooksCapability_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_reconciler_proto_enumTypes[0].Descriptor()
}

func (ReconcilerHooksCapability_Kind) Type() protoreflect.EnumType {
	return &file_proto_reconciler_proto_enumTypes[0]
}

func (x ReconcilerHooksCapability_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReconcilerHooksCapability_Kind.Descriptor instead.
func (ReconcilerHooksCapability_Kind) EnumDescriptor() ([]byte, []int) {
	return file_proto_reconciler_proto_rawDescGZIP(), []int{2, 0}
}

type ReconcilerHooksResult_Behavior int32

const (
	ReconcilerHooksResult_BEHAVIOR_UNSPECIFIED ReconcilerHooksResult_Behavior = 0
	// BEHAVIOR_CONTINUE indicates that this reconciliation loop will
	// proceed running.
	// BEHAVIOR_CONTINUE is useful when the plugin executes changes on internal status or resources not directly managed
	// by the main reconciliation loop
	ReconcilerHooksResult_BEHAVIOR_CONTINUE ReconcilerHooksResult_Behavior = 1
	// BEHAVIOR_REQUEUE indicates that this reconciliation loop will
	// be stopped and a new one will be requested.
	// BEHAVIOR_REQUEUE should always be set when the plugin applies changes on resources that are directly managed
	// by the main reconciliation loop
	ReconcilerHooksResult_BEHAVIOR_REQUEUE ReconcilerHooksResult_Behavior = 2
	// BEHAVIOR_TERMINATE indicates that the main reconciliation loop needs to
	// be marked as succeeded and no further operations needs to be taken.
	// This should be used when the invoked Reconcile hook act as a substitute for the main CNPG reconciliation loop.
	// An example would be a plugin that substitutes the `Backup` logic of CNPG.
	ReconcilerHooksResult_BEHAVIOR_TERMINATE ReconcilerHooksResult_Behavior = 3
)

// Enum value maps for ReconcilerHooksResult_Behavior.
var (
	ReconcilerHooksResult_Behavior_name = map[int32]string{
		0: "BEHAVIOR_UNSPECIFIED",
		1: "BEHAVIOR_CONTINUE",
		2: "BEHAVIOR_REQUEUE",
		3: "BEHAVIOR_TERMINATE",
	}
	ReconcilerHooksResult_Behavior_value = map[string]int32{
		"BEHAVIOR_UNSPECIFIED": 0,
		"BEHAVIOR_CONTINUE":    1,
		"BEHAVIOR_REQUEUE":     2,
		"BEHAVIOR_TERMINATE":   3,
	}
)

func (x ReconcilerHooksResult_Behavior) Enum() *ReconcilerHooksResult_Behavior {
	p := new(ReconcilerHooksResult_Behavior)
	*p = x
	return p
}

func (x ReconcilerHooksResult_Behavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReconcilerHooksResult_Behavior) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_reconciler_proto_enumTypes[1].Descriptor()
}

func (ReconcilerHooksResult_Behavior) Type() protoreflect.EnumType {
	return &file_proto_reconciler_proto_enumTypes[1]
}

func (x ReconcilerHooksResult_Behavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReconcilerHooksResult_Behavior.Descriptor instead.
func (ReconcilerHooksResult_Behavior) EnumDescriptor() ([]byte, []int) {
	return file_proto_reconciler_proto_rawDescGZIP(), []int{4, 0}
}

type ReconcilerHooksCapabilitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReconcilerHooksCapabilitiesRequest) Reset() {
	*x = ReconcilerHooksCapabilitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reconciler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconcilerHooksCapabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconcilerHooksCapabilitiesRequest) ProtoMessage() {}

func (x *ReconcilerHooksCapabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reconciler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconcilerHooksCapabilitiesRequest.ProtoReflect.Descriptor instead.
func (*ReconcilerHooksCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_proto_reconciler_proto_rawDescGZIP(), []int{0}
}

type ReconcilerHooksCapabilitiesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This message is OPTIONAL, containing the list of resources
	// for which the lifecycle hook is called
	ReconcilerCapabilities []*ReconcilerHooksCapability `protobuf:"bytes,1,rep,name=reconciler_capabilities,json=reconcilerCapabilities,proto3" json:"reconciler_capabilities,omitempty"`
}

func (x *ReconcilerHooksCapabilitiesResult) Reset() {
	*x = ReconcilerHooksCapabilitiesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reconciler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconcilerHooksCapabilitiesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconcilerHooksCapabilitiesResult) ProtoMessage() {}

func (x *ReconcilerHooksCapabilitiesResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reconciler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconcilerHooksCapabilitiesResult.ProtoReflect.Descriptor instead.
func (*ReconcilerHooksCapabilitiesResult) Descriptor() ([]byte, []int) {
	return file_proto_reconciler_proto_rawDescGZIP(), []int{1}
}

func (x *ReconcilerHooksCapabilitiesResult) GetReconcilerCapabilities() []*ReconcilerHooksCapability {
	if x != nil {
		return x.ReconcilerCapabilities
	}
	return nil
}

type ReconcilerHooksCapability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// kind is the controller Kind
	Kind ReconcilerHooksCapability_Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=cnpgi.reconciler.v1.ReconcilerHooksCapability_Kind" json:"kind,omitempty"`
}

func (x *ReconcilerHooksCapability) Reset() {
	*x = ReconcilerHooksCapability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reconciler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconcilerHooksCapability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconcilerHooksCapability) ProtoMessage() {}

func (x *ReconcilerHooksCapability) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reconciler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconcilerHooksCapability.ProtoReflect.Descriptor instead.
func (*ReconcilerHooksCapability) Descriptor() ([]byte, []int) {
	return file_proto_reconciler_proto_rawDescGZIP(), []int{2}
}

func (x *ReconcilerHooksCapability) GetKind() ReconcilerHooksCapability_Kind {
	if x != nil {
		return x.Kind
	}
	return ReconcilerHooksCapability_KIND_UNSPECIFIED
}

type ReconcilerHooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field is REQUIRED. Value of this field is the JSON
	// serialization of the Cluster tied to the `Kind` being reconciled
	ClusterDefinition []byte `protobuf:"bytes,1,opt,name=cluster_definition,json=clusterDefinition,proto3" json:"cluster_definition,omitempty"`
	// This field is REQUIRED. Value of this field is the JSON
	// serialization of the resource being reconciled. Please note that in case of Cluster Reconciliation, the
	// `resource_definition` will match the `cluster_definition`
	ResourceDefinition []byte `protobuf:"bytes,2,opt,name=resource_definition,json=resourceDefinition,proto3" json:"resource_definition,omitempty"`
}

func (x *ReconcilerHooksRequest) Reset() {
	*x = ReconcilerHooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reconciler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconcilerHooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconcilerHooksRequest) ProtoMessage() {}

func (x *ReconcilerHooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reconciler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconcilerHooksRequest.ProtoReflect.Descriptor instead.
func (*ReconcilerHooksRequest) Descriptor() ([]byte, []int) {
	return file_proto_reconciler_proto_rawDescGZIP(), []int{3}
}

func (x *ReconcilerHooksRequest) GetClusterDefinition() []byte {
	if x != nil {
		return x.ClusterDefinition
	}
	return nil
}

func (x *ReconcilerHooksRequest) GetResourceDefinition() []byte {
	if x != nil {
		return x.ResourceDefinition
	}
	return nil
}

// ReconcilerHooksResult response is used to instruct then the CNPG controller on which action
// take after the plugin execution.
type ReconcilerHooksResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field is REQUIRED, and indicates the behavior that should
	// be used for the current reconciliation loop.
	Behavior ReconcilerHooksResult_Behavior `protobuf:"varint,1,opt,name=behavior,proto3,enum=cnpgi.reconciler.v1.ReconcilerHooksResult_Behavior" json:"behavior,omitempty"`
	// This field is OPTIONAL. If true, the current reconciliation loop
	// will be stopped and the operator will ensure that another one will
	// be run in the requested number of seconds. IMPORTANT: the new
	// reconciliation loop may start even before the number of specified
	// seconds.
	RequeueAfter int64 `protobuf:"varint,2,opt,name=requeue_after,json=requeueAfter,proto3" json:"requeue_after,omitempty"`
}

func (x *ReconcilerHooksResult) Reset() {
	*x = ReconcilerHooksResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_reconciler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconcilerHooksResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconcilerHooksResult) ProtoMessage() {}

func (x *ReconcilerHooksResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reconciler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconcilerHooksResult.ProtoReflect.Descriptor instead.
func (*ReconcilerHooksResult) Descriptor() ([]byte, []int) {
	return file_proto_reconciler_proto_rawDescGZIP(), []int{4}
}

func (x *ReconcilerHooksResult) GetBehavior() ReconcilerHooksResult_Behavior {
	if x != nil {
		return x.Behavior
	}
	return ReconcilerHooksResult_BEHAVIOR_UNSPECIFIED
}

func (x *ReconcilerHooksResult) GetRequeueAfter() int64 {
	if x != nil {
		return x.RequeueAfter
	}
	return 0
}

var File_proto_reconciler_proto protoreflect.FileDescriptor

var file_proto_reconciler_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x24, 0x0a,
	0x22, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x21, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c,
	0x65, 0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x67, 0x0a, 0x17, 0x72, 0x65, 0x63,
	0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6e, 0x70,
	0x67, 0x69, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x16, 0x72, 0x65, 0x63, 0x6f,
	0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x22, 0xa5, 0x01, 0x0a, 0x19, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65,
	0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x47, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33,
	0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x48,
	0x6f, 0x6f, 0x6b, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x3f, 0x0a, 0x04, 0x4b, 0x69, 0x6e,
	0x64, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x49, 0x4e, 0x44, 0x5f,
	0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x02, 0x22, 0x78, 0x0a, 0x16, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x11, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf8, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69,
	0x6c, 0x65, 0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4f,
	0x0a, 0x08, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x33, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65,
	0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x42, 0x65, 0x68,
	0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x08, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x22, 0x69, 0x0a, 0x08, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x14, 0x42, 0x45, 0x48, 0x41, 0x56, 0x49, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x45,
	0x48, 0x41, 0x56, 0x49, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x45, 0x48, 0x41, 0x56, 0x49, 0x4f, 0x52, 0x5f, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x55, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x45, 0x48, 0x41, 0x56,
	0x49, 0x4f, 0x52, 0x5f, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x10, 0x03, 0x32,
	0xdd, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x48, 0x6f,
	0x6f, 0x6b, 0x73, 0x12, 0x84, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x37, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69,
	0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65,
	0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x03, 0x50, 0x72,
	0x65, 0x12, 0x2b, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c,
	0x65, 0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x48,
	0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x04,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e,
	0x63, 0x69, 0x6c, 0x65, 0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c,
	0x65, 0x72, 0x48, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x70, 0x67, 0x2f, 0x63, 0x6e, 0x70,
	0x67, 0x2d, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_reconciler_proto_rawDescOnce sync.Once
	file_proto_reconciler_proto_rawDescData = file_proto_reconciler_proto_rawDesc
)

func file_proto_reconciler_proto_rawDescGZIP() []byte {
	file_proto_reconciler_proto_rawDescOnce.Do(func() {
		file_proto_reconciler_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_reconciler_proto_rawDescData)
	})
	return file_proto_reconciler_proto_rawDescData
}

var file_proto_reconciler_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_reconciler_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_reconciler_proto_goTypes = []interface{}{
	(ReconcilerHooksCapability_Kind)(0),        // 0: cnpgi.reconciler.v1.ReconcilerHooksCapability.Kind
	(ReconcilerHooksResult_Behavior)(0),        // 1: cnpgi.reconciler.v1.ReconcilerHooksResult.Behavior
	(*ReconcilerHooksCapabilitiesRequest)(nil), // 2: cnpgi.reconciler.v1.ReconcilerHooksCapabilitiesRequest
	(*ReconcilerHooksCapabilitiesResult)(nil),  // 3: cnpgi.reconciler.v1.ReconcilerHooksCapabilitiesResult
	(*ReconcilerHooksCapability)(nil),          // 4: cnpgi.reconciler.v1.ReconcilerHooksCapability
	(*ReconcilerHooksRequest)(nil),             // 5: cnpgi.reconciler.v1.ReconcilerHooksRequest
	(*ReconcilerHooksResult)(nil),              // 6: cnpgi.reconciler.v1.ReconcilerHooksResult
}
var file_proto_reconciler_proto_depIdxs = []int32{
	4, // 0: cnpgi.reconciler.v1.ReconcilerHooksCapabilitiesResult.reconciler_capabilities:type_name -> cnpgi.reconciler.v1.ReconcilerHooksCapability
	0, // 1: cnpgi.reconciler.v1.ReconcilerHooksCapability.kind:type_name -> cnpgi.reconciler.v1.ReconcilerHooksCapability.Kind
	1, // 2: cnpgi.reconciler.v1.ReconcilerHooksResult.behavior:type_name -> cnpgi.reconciler.v1.ReconcilerHooksResult.Behavior
	2, // 3: cnpgi.reconciler.v1.ReconcilerHooks.GetCapabilities:input_type -> cnpgi.reconciler.v1.ReconcilerHooksCapabilitiesRequest
	5, // 4: cnpgi.reconciler.v1.ReconcilerHooks.Pre:input_type -> cnpgi.reconciler.v1.ReconcilerHooksRequest
	5, // 5: cnpgi.reconciler.v1.ReconcilerHooks.Post:input_type -> cnpgi.reconciler.v1.ReconcilerHooksRequest
	3, // 6: cnpgi.reconciler.v1.ReconcilerHooks.GetCapabilities:output_type -> cnpgi.reconciler.v1.ReconcilerHooksCapabilitiesResult
	6, // 7: cnpgi.reconciler.v1.ReconcilerHooks.Pre:output_type -> cnpgi.reconciler.v1.ReconcilerHooksResult
	6, // 8: cnpgi.reconciler.v1.ReconcilerHooks.Post:output_type -> cnpgi.reconciler.v1.ReconcilerHooksResult
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_reconciler_proto_init() }
func file_proto_reconciler_proto_init() {
	if File_proto_reconciler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_reconciler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconcilerHooksCapabilitiesRequest); i {
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
		file_proto_reconciler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconcilerHooksCapabilitiesResult); i {
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
		file_proto_reconciler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconcilerHooksCapability); i {
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
		file_proto_reconciler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconcilerHooksRequest); i {
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
		file_proto_reconciler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconcilerHooksResult); i {
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
			RawDescriptor: file_proto_reconciler_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_reconciler_proto_goTypes,
		DependencyIndexes: file_proto_reconciler_proto_depIdxs,
		EnumInfos:         file_proto_reconciler_proto_enumTypes,
		MessageInfos:      file_proto_reconciler_proto_msgTypes,
	}.Build()
	File_proto_reconciler_proto = out.File
	file_proto_reconciler_proto_rawDesc = nil
	file_proto_reconciler_proto_goTypes = nil
	file_proto_reconciler_proto_depIdxs = nil
}
