// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/operator.proto

package operator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OperatorCapability_RPC_Type int32

const (
	OperatorCapability_RPC_TYPE_UNSPECIFIED OperatorCapability_RPC_Type = 0
	// TYPE_VALIDATE_CLUSTER_CREATE indicates that the Plugin is able to
	// reply to the ValidateClusterCreate RPC request
	OperatorCapability_RPC_TYPE_VALIDATE_CLUSTER_CREATE OperatorCapability_RPC_Type = 1
	// TYPE_VALIDATE_CLUSTER_CHANGE indicates that the Plugin is able to
	// reply to the ValidateClusterChange RPC request
	OperatorCapability_RPC_TYPE_VALIDATE_CLUSTER_CHANGE OperatorCapability_RPC_Type = 2
	// TYPE_MUTATE_CLUSTER indicates that the Plugin is able to
	// reply to the MutateCluster RPC request
	OperatorCapability_RPC_TYPE_MUTATE_CLUSTER OperatorCapability_RPC_Type = 3
	// TYPE_SET_STATUS_IN_CLUSTER indicates that the Plugin is able to
	// reply to the SetStatusInCluster RPC request
	OperatorCapability_RPC_TYPE_SET_STATUS_IN_CLUSTER OperatorCapability_RPC_Type = 5
	// TYPE_DEREGISTER indicates that the Plugin is able to execute the cleanup logic once it is removed from
	// the cluster definition
	OperatorCapability_RPC_TYPE_DEREGISTER OperatorCapability_RPC_Type = 6
)

// Enum value maps for OperatorCapability_RPC_Type.
var (
	OperatorCapability_RPC_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_VALIDATE_CLUSTER_CREATE",
		2: "TYPE_VALIDATE_CLUSTER_CHANGE",
		3: "TYPE_MUTATE_CLUSTER",
		5: "TYPE_SET_STATUS_IN_CLUSTER",
		6: "TYPE_DEREGISTER",
	}
	OperatorCapability_RPC_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":             0,
		"TYPE_VALIDATE_CLUSTER_CREATE": 1,
		"TYPE_VALIDATE_CLUSTER_CHANGE": 2,
		"TYPE_MUTATE_CLUSTER":          3,
		"TYPE_SET_STATUS_IN_CLUSTER":   5,
		"TYPE_DEREGISTER":              6,
	}
)

func (x OperatorCapability_RPC_Type) Enum() *OperatorCapability_RPC_Type {
	p := new(OperatorCapability_RPC_Type)
	*p = x
	return p
}

func (x OperatorCapability_RPC_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperatorCapability_RPC_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_operator_proto_enumTypes[0].Descriptor()
}

func (OperatorCapability_RPC_Type) Type() protoreflect.EnumType {
	return &file_proto_operator_proto_enumTypes[0]
}

func (x OperatorCapability_RPC_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperatorCapability_RPC_Type.Descriptor instead.
func (OperatorCapability_RPC_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{11, 0, 0}
}

type OperatorCapabilitiesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorCapabilitiesRequest) Reset() {
	*x = OperatorCapabilitiesRequest{}
	mi := &file_proto_operator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorCapabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorCapabilitiesRequest) ProtoMessage() {}

func (x *OperatorCapabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorCapabilitiesRequest.ProtoReflect.Descriptor instead.
func (*OperatorCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{0}
}

type OperatorCapabilitiesResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// All the capabilities that the controller service supports. This
	// field is OPTIONAL.
	Capabilities  []*OperatorCapability `protobuf:"bytes,1,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorCapabilitiesResult) Reset() {
	*x = OperatorCapabilitiesResult{}
	mi := &file_proto_operator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorCapabilitiesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorCapabilitiesResult) ProtoMessage() {}

func (x *OperatorCapabilitiesResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorCapabilitiesResult.ProtoReflect.Descriptor instead.
func (*OperatorCapabilitiesResult) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{1}
}

func (x *OperatorCapabilitiesResult) GetCapabilities() []*OperatorCapability {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

type OperatorValidateClusterCreateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is REQUIRED. Value of this field is the JSON
	// serialization of the Cluster that is being created
	Definition    []byte `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorValidateClusterCreateRequest) Reset() {
	*x = OperatorValidateClusterCreateRequest{}
	mi := &file_proto_operator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorValidateClusterCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorValidateClusterCreateRequest) ProtoMessage() {}

func (x *OperatorValidateClusterCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorValidateClusterCreateRequest.ProtoReflect.Descriptor instead.
func (*OperatorValidateClusterCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{2}
}

func (x *OperatorValidateClusterCreateRequest) GetDefinition() []byte {
	if x != nil {
		return x.Definition
	}
	return nil
}

type OperatorValidateClusterCreateResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is OPTIONAL. Value of this field is a set
	// of validation errors
	ValidationErrors []*ValidationError `protobuf:"bytes,1,rep,name=validation_errors,json=validationErrors,proto3" json:"validation_errors,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *OperatorValidateClusterCreateResult) Reset() {
	*x = OperatorValidateClusterCreateResult{}
	mi := &file_proto_operator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorValidateClusterCreateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorValidateClusterCreateResult) ProtoMessage() {}

func (x *OperatorValidateClusterCreateResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorValidateClusterCreateResult.ProtoReflect.Descriptor instead.
func (*OperatorValidateClusterCreateResult) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{3}
}

func (x *OperatorValidateClusterCreateResult) GetValidationErrors() []*ValidationError {
	if x != nil {
		return x.ValidationErrors
	}
	return nil
}

type OperatorValidateClusterChangeRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is REQUIRED. Value of this field is the JSON
	// serialization of the current Cluster definition
	OldCluster []byte `protobuf:"bytes,1,opt,name=old_cluster,json=oldCluster,proto3" json:"old_cluster,omitempty"`
	// This field is REQUIRED. Value of this field is the JSON
	// serialization of the updated Cluster definition
	NewCluster    []byte `protobuf:"bytes,2,opt,name=new_cluster,json=newCluster,proto3" json:"new_cluster,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorValidateClusterChangeRequest) Reset() {
	*x = OperatorValidateClusterChangeRequest{}
	mi := &file_proto_operator_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorValidateClusterChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorValidateClusterChangeRequest) ProtoMessage() {}

func (x *OperatorValidateClusterChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorValidateClusterChangeRequest.ProtoReflect.Descriptor instead.
func (*OperatorValidateClusterChangeRequest) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{4}
}

func (x *OperatorValidateClusterChangeRequest) GetOldCluster() []byte {
	if x != nil {
		return x.OldCluster
	}
	return nil
}

func (x *OperatorValidateClusterChangeRequest) GetNewCluster() []byte {
	if x != nil {
		return x.NewCluster
	}
	return nil
}

type OperatorValidateClusterChangeResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is OPTIONAL. Value of this field is a set
	// of validation errors
	ValidationErrors []*ValidationError `protobuf:"bytes,1,rep,name=validation_errors,json=validationErrors,proto3" json:"validation_errors,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *OperatorValidateClusterChangeResult) Reset() {
	*x = OperatorValidateClusterChangeResult{}
	mi := &file_proto_operator_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorValidateClusterChangeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorValidateClusterChangeResult) ProtoMessage() {}

func (x *OperatorValidateClusterChangeResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorValidateClusterChangeResult.ProtoReflect.Descriptor instead.
func (*OperatorValidateClusterChangeResult) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{5}
}

func (x *OperatorValidateClusterChangeResult) GetValidationErrors() []*ValidationError {
	if x != nil {
		return x.ValidationErrors
	}
	return nil
}

type OperatorMutateClusterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is REQUIRED. Value of this field is the JSON
	// serialization of the Cluster that should receive the
	// default values
	Definition    []byte `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorMutateClusterRequest) Reset() {
	*x = OperatorMutateClusterRequest{}
	mi := &file_proto_operator_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorMutateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorMutateClusterRequest) ProtoMessage() {}

func (x *OperatorMutateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorMutateClusterRequest.ProtoReflect.Descriptor instead.
func (*OperatorMutateClusterRequest) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{6}
}

func (x *OperatorMutateClusterRequest) GetDefinition() []byte {
	if x != nil {
		return x.Definition
	}
	return nil
}

type DeregisterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is REQUIRED. Value of this field is the JSON
	// serialization of the Cluster that should receive the
	// default values
	Definition    []byte `protobuf:"bytes,1,opt,name=definition,proto3" json:"definition,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeregisterRequest) Reset() {
	*x = DeregisterRequest{}
	mi := &file_proto_operator_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeregisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterRequest) ProtoMessage() {}

func (x *DeregisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterRequest.ProtoReflect.Descriptor instead.
func (*DeregisterRequest) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{7}
}

func (x *DeregisterRequest) GetDefinition() []byte {
	if x != nil {
		return x.Definition
	}
	return nil
}

type DeregisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeregisterResponse) Reset() {
	*x = DeregisterResponse{}
	mi := &file_proto_operator_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeregisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeregisterResponse) ProtoMessage() {}

func (x *DeregisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeregisterResponse.ProtoReflect.Descriptor instead.
func (*DeregisterResponse) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{8}
}

type OperatorMutateClusterResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is OPTIONAL. Value of this field is a JSONPatch
	// to be applied on the passed Cluster definition
	JsonPatch     []byte `protobuf:"bytes,1,opt,name=json_patch,json=jsonPatch,proto3" json:"json_patch,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorMutateClusterResult) Reset() {
	*x = OperatorMutateClusterResult{}
	mi := &file_proto_operator_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorMutateClusterResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorMutateClusterResult) ProtoMessage() {}

func (x *OperatorMutateClusterResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorMutateClusterResult.ProtoReflect.Descriptor instead.
func (*OperatorMutateClusterResult) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{9}
}

func (x *OperatorMutateClusterResult) GetJsonPatch() []byte {
	if x != nil {
		return x.JsonPatch
	}
	return nil
}

type ValidationError struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is REQUIRED. Value of this field is
	PathComponents []string `protobuf:"bytes,1,rep,name=path_components,json=pathComponents,proto3" json:"path_components,omitempty"`
	// This field is REQUIRED. Value of this field is
	// the value that caused a validation error
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// This field is REQUIRED. Value of this field is a
	// description of the validation error
	Message       string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidationError) Reset() {
	*x = ValidationError{}
	mi := &file_proto_operator_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationError) ProtoMessage() {}

func (x *ValidationError) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationError.ProtoReflect.Descriptor instead.
func (*ValidationError) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{10}
}

func (x *ValidationError) GetPathComponents() []string {
	if x != nil {
		return x.PathComponents
	}
	return nil
}

func (x *ValidationError) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ValidationError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type OperatorCapability struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Type:
	//
	//	*OperatorCapability_Rpc
	Type          isOperatorCapability_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorCapability) Reset() {
	*x = OperatorCapability{}
	mi := &file_proto_operator_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorCapability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorCapability) ProtoMessage() {}

func (x *OperatorCapability) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorCapability.ProtoReflect.Descriptor instead.
func (*OperatorCapability) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{11}
}

func (x *OperatorCapability) GetType() isOperatorCapability_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *OperatorCapability) GetRpc() *OperatorCapability_RPC {
	if x != nil {
		if x, ok := x.Type.(*OperatorCapability_Rpc); ok {
			return x.Rpc
		}
	}
	return nil
}

type isOperatorCapability_Type interface {
	isOperatorCapability_Type()
}

type OperatorCapability_Rpc struct {
	Rpc *OperatorCapability_RPC `protobuf:"bytes,1,opt,name=rpc,proto3,oneof"`
}

func (*OperatorCapability_Rpc) isOperatorCapability_Type() {}

type SetStatusInClusterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is REQUIRED. Value of this field is the JSON
	// serialization of the current Cluster definition
	Cluster       []byte `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetStatusInClusterRequest) Reset() {
	*x = SetStatusInClusterRequest{}
	mi := &file_proto_operator_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetStatusInClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStatusInClusterRequest) ProtoMessage() {}

func (x *SetStatusInClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStatusInClusterRequest.ProtoReflect.Descriptor instead.
func (*SetStatusInClusterRequest) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{12}
}

func (x *SetStatusInClusterRequest) GetCluster() []byte {
	if x != nil {
		return x.Cluster
	}
	return nil
}

type SetStatusInClusterResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is OPTIONAL. No value means no-op, otherwise the passed json is set inside the
	// .status.plugins[pluginName] key
	JsonStatus    []byte `protobuf:"bytes,1,opt,name=json_status,json=jsonStatus,proto3" json:"json_status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetStatusInClusterResponse) Reset() {
	*x = SetStatusInClusterResponse{}
	mi := &file_proto_operator_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetStatusInClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStatusInClusterResponse) ProtoMessage() {}

func (x *SetStatusInClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStatusInClusterResponse.ProtoReflect.Descriptor instead.
func (*SetStatusInClusterResponse) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{13}
}

func (x *SetStatusInClusterResponse) GetJsonStatus() []byte {
	if x != nil {
		return x.JsonStatus
	}
	return nil
}

type OperatorCapability_RPC struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Type          OperatorCapability_RPC_Type `protobuf:"varint,1,opt,name=type,proto3,enum=cnpgi.operator.v1.OperatorCapability_RPC_Type" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorCapability_RPC) Reset() {
	*x = OperatorCapability_RPC{}
	mi := &file_proto_operator_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorCapability_RPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorCapability_RPC) ProtoMessage() {}

func (x *OperatorCapability_RPC) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperatorCapability_RPC.ProtoReflect.Descriptor instead.
func (*OperatorCapability_RPC) Descriptor() ([]byte, []int) {
	return file_proto_operator_proto_rawDescGZIP(), []int{11, 0}
}

func (x *OperatorCapability_RPC) GetType() OperatorCapability_RPC_Type {
	if x != nil {
		return x.Type
	}
	return OperatorCapability_RPC_TYPE_UNSPECIFIED
}

var File_proto_operator_proto protoreflect.FileDescriptor

const file_proto_operator_proto_rawDesc = "" +
	"\n" +
	"\x14proto/operator.proto\x12\x11cnpgi.operator.v1\"\x1d\n" +
	"\x1bOperatorCapabilitiesRequest\"g\n" +
	"\x1aOperatorCapabilitiesResult\x12I\n" +
	"\fcapabilities\x18\x01 \x03(\v2%.cnpgi.operator.v1.OperatorCapabilityR\fcapabilities\"F\n" +
	"$OperatorValidateClusterCreateRequest\x12\x1e\n" +
	"\n" +
	"definition\x18\x01 \x01(\fR\n" +
	"definition\"v\n" +
	"#OperatorValidateClusterCreateResult\x12O\n" +
	"\x11validation_errors\x18\x01 \x03(\v2\".cnpgi.operator.v1.ValidationErrorR\x10validationErrors\"h\n" +
	"$OperatorValidateClusterChangeRequest\x12\x1f\n" +
	"\vold_cluster\x18\x01 \x01(\fR\n" +
	"oldCluster\x12\x1f\n" +
	"\vnew_cluster\x18\x02 \x01(\fR\n" +
	"newCluster\"v\n" +
	"#OperatorValidateClusterChangeResult\x12O\n" +
	"\x11validation_errors\x18\x01 \x03(\v2\".cnpgi.operator.v1.ValidationErrorR\x10validationErrors\">\n" +
	"\x1cOperatorMutateClusterRequest\x12\x1e\n" +
	"\n" +
	"definition\x18\x01 \x01(\fR\n" +
	"definition\"3\n" +
	"\x11DeregisterRequest\x12\x1e\n" +
	"\n" +
	"definition\x18\x01 \x01(\fR\n" +
	"definition\"\x14\n" +
	"\x12DeregisterResponse\"<\n" +
	"\x1bOperatorMutateClusterResult\x12\x1d\n" +
	"\n" +
	"json_patch\x18\x01 \x01(\fR\tjsonPatch\"j\n" +
	"\x0fValidationError\x12'\n" +
	"\x0fpath_components\x18\x01 \x03(\tR\x0epathComponents\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\"\xd8\x02\n" +
	"\x12OperatorCapability\x12=\n" +
	"\x03rpc\x18\x01 \x01(\v2).cnpgi.operator.v1.OperatorCapability.RPCH\x00R\x03rpc\x1a\xfa\x01\n" +
	"\x03RPC\x12B\n" +
	"\x04type\x18\x01 \x01(\x0e2..cnpgi.operator.v1.OperatorCapability.RPC.TypeR\x04type\"\xae\x01\n" +
	"\x04Type\x12\x14\n" +
	"\x10TYPE_UNSPECIFIED\x10\x00\x12 \n" +
	"\x1cTYPE_VALIDATE_CLUSTER_CREATE\x10\x01\x12 \n" +
	"\x1cTYPE_VALIDATE_CLUSTER_CHANGE\x10\x02\x12\x17\n" +
	"\x13TYPE_MUTATE_CLUSTER\x10\x03\x12\x1e\n" +
	"\x1aTYPE_SET_STATUS_IN_CLUSTER\x10\x05\x12\x13\n" +
	"\x0fTYPE_DEREGISTER\x10\x06B\x06\n" +
	"\x04type\"5\n" +
	"\x19SetStatusInClusterRequest\x12\x18\n" +
	"\acluster\x18\x01 \x01(\fR\acluster\"=\n" +
	"\x1aSetStatusInClusterResponse\x12\x1f\n" +
	"\vjson_status\x18\x01 \x01(\fR\n" +
	"jsonStatus2\xde\x05\n" +
	"\bOperator\x12r\n" +
	"\x0fGetCapabilities\x12..cnpgi.operator.v1.OperatorCapabilitiesRequest\x1a-.cnpgi.operator.v1.OperatorCapabilitiesResult\"\x00\x12\x8a\x01\n" +
	"\x15ValidateClusterCreate\x127.cnpgi.operator.v1.OperatorValidateClusterCreateRequest\x1a6.cnpgi.operator.v1.OperatorValidateClusterCreateResult\"\x00\x12\x8a\x01\n" +
	"\x15ValidateClusterChange\x127.cnpgi.operator.v1.OperatorValidateClusterChangeRequest\x1a6.cnpgi.operator.v1.OperatorValidateClusterChangeResult\"\x00\x12r\n" +
	"\rMutateCluster\x12/.cnpgi.operator.v1.OperatorMutateClusterRequest\x1a..cnpgi.operator.v1.OperatorMutateClusterResult\"\x00\x12s\n" +
	"\x12SetStatusInCluster\x12,.cnpgi.operator.v1.SetStatusInClusterRequest\x1a-.cnpgi.operator.v1.SetStatusInClusterResponse\"\x00\x12[\n" +
	"\n" +
	"Deregister\x12$.cnpgi.operator.v1.DeregisterRequest\x1a%.cnpgi.operator.v1.DeregisterResponse\"\x00B/Z-github.com/cloudnative-pg/cnpg-i/pkg/operatorb\x06proto3"

var (
	file_proto_operator_proto_rawDescOnce sync.Once
	file_proto_operator_proto_rawDescData []byte
)

func file_proto_operator_proto_rawDescGZIP() []byte {
	file_proto_operator_proto_rawDescOnce.Do(func() {
		file_proto_operator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_operator_proto_rawDesc), len(file_proto_operator_proto_rawDesc)))
	})
	return file_proto_operator_proto_rawDescData
}

var file_proto_operator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_operator_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_operator_proto_goTypes = []any{
	(OperatorCapability_RPC_Type)(0),             // 0: cnpgi.operator.v1.OperatorCapability.RPC.Type
	(*OperatorCapabilitiesRequest)(nil),          // 1: cnpgi.operator.v1.OperatorCapabilitiesRequest
	(*OperatorCapabilitiesResult)(nil),           // 2: cnpgi.operator.v1.OperatorCapabilitiesResult
	(*OperatorValidateClusterCreateRequest)(nil), // 3: cnpgi.operator.v1.OperatorValidateClusterCreateRequest
	(*OperatorValidateClusterCreateResult)(nil),  // 4: cnpgi.operator.v1.OperatorValidateClusterCreateResult
	(*OperatorValidateClusterChangeRequest)(nil), // 5: cnpgi.operator.v1.OperatorValidateClusterChangeRequest
	(*OperatorValidateClusterChangeResult)(nil),  // 6: cnpgi.operator.v1.OperatorValidateClusterChangeResult
	(*OperatorMutateClusterRequest)(nil),         // 7: cnpgi.operator.v1.OperatorMutateClusterRequest
	(*DeregisterRequest)(nil),                    // 8: cnpgi.operator.v1.DeregisterRequest
	(*DeregisterResponse)(nil),                   // 9: cnpgi.operator.v1.DeregisterResponse
	(*OperatorMutateClusterResult)(nil),          // 10: cnpgi.operator.v1.OperatorMutateClusterResult
	(*ValidationError)(nil),                      // 11: cnpgi.operator.v1.ValidationError
	(*OperatorCapability)(nil),                   // 12: cnpgi.operator.v1.OperatorCapability
	(*SetStatusInClusterRequest)(nil),            // 13: cnpgi.operator.v1.SetStatusInClusterRequest
	(*SetStatusInClusterResponse)(nil),           // 14: cnpgi.operator.v1.SetStatusInClusterResponse
	(*OperatorCapability_RPC)(nil),               // 15: cnpgi.operator.v1.OperatorCapability.RPC
}
var file_proto_operator_proto_depIdxs = []int32{
	12, // 0: cnpgi.operator.v1.OperatorCapabilitiesResult.capabilities:type_name -> cnpgi.operator.v1.OperatorCapability
	11, // 1: cnpgi.operator.v1.OperatorValidateClusterCreateResult.validation_errors:type_name -> cnpgi.operator.v1.ValidationError
	11, // 2: cnpgi.operator.v1.OperatorValidateClusterChangeResult.validation_errors:type_name -> cnpgi.operator.v1.ValidationError
	15, // 3: cnpgi.operator.v1.OperatorCapability.rpc:type_name -> cnpgi.operator.v1.OperatorCapability.RPC
	0,  // 4: cnpgi.operator.v1.OperatorCapability.RPC.type:type_name -> cnpgi.operator.v1.OperatorCapability.RPC.Type
	1,  // 5: cnpgi.operator.v1.Operator.GetCapabilities:input_type -> cnpgi.operator.v1.OperatorCapabilitiesRequest
	3,  // 6: cnpgi.operator.v1.Operator.ValidateClusterCreate:input_type -> cnpgi.operator.v1.OperatorValidateClusterCreateRequest
	5,  // 7: cnpgi.operator.v1.Operator.ValidateClusterChange:input_type -> cnpgi.operator.v1.OperatorValidateClusterChangeRequest
	7,  // 8: cnpgi.operator.v1.Operator.MutateCluster:input_type -> cnpgi.operator.v1.OperatorMutateClusterRequest
	13, // 9: cnpgi.operator.v1.Operator.SetStatusInCluster:input_type -> cnpgi.operator.v1.SetStatusInClusterRequest
	8,  // 10: cnpgi.operator.v1.Operator.Deregister:input_type -> cnpgi.operator.v1.DeregisterRequest
	2,  // 11: cnpgi.operator.v1.Operator.GetCapabilities:output_type -> cnpgi.operator.v1.OperatorCapabilitiesResult
	4,  // 12: cnpgi.operator.v1.Operator.ValidateClusterCreate:output_type -> cnpgi.operator.v1.OperatorValidateClusterCreateResult
	6,  // 13: cnpgi.operator.v1.Operator.ValidateClusterChange:output_type -> cnpgi.operator.v1.OperatorValidateClusterChangeResult
	10, // 14: cnpgi.operator.v1.Operator.MutateCluster:output_type -> cnpgi.operator.v1.OperatorMutateClusterResult
	14, // 15: cnpgi.operator.v1.Operator.SetStatusInCluster:output_type -> cnpgi.operator.v1.SetStatusInClusterResponse
	9,  // 16: cnpgi.operator.v1.Operator.Deregister:output_type -> cnpgi.operator.v1.DeregisterResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_operator_proto_init() }
func file_proto_operator_proto_init() {
	if File_proto_operator_proto != nil {
		return
	}
	file_proto_operator_proto_msgTypes[11].OneofWrappers = []any{
		(*OperatorCapability_Rpc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_operator_proto_rawDesc), len(file_proto_operator_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_operator_proto_goTypes,
		DependencyIndexes: file_proto_operator_proto_depIdxs,
		EnumInfos:         file_proto_operator_proto_enumTypes,
		MessageInfos:      file_proto_operator_proto_msgTypes,
	}.Build()
	File_proto_operator_proto = out.File
	file_proto_operator_proto_goTypes = nil
	file_proto_operator_proto_depIdxs = nil
}
