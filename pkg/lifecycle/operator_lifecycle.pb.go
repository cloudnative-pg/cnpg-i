// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/operator_lifecycle.proto

package lifecycle

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

// The operator type corresponds to the Kubernetes API method
type OperatorOperationType_Type int32

const (
	OperatorOperationType_TYPE_UNSPECIFIED OperatorOperationType_Type = 0
	OperatorOperationType_TYPE_PATCH       OperatorOperationType_Type = 1
	OperatorOperationType_TYPE_UPDATE      OperatorOperationType_Type = 2
	OperatorOperationType_TYPE_CREATE      OperatorOperationType_Type = 3
	OperatorOperationType_TYPE_DELETE      OperatorOperationType_Type = 4
	OperatorOperationType_TYPE_DEREGISTER  OperatorOperationType_Type = 5
	OperatorOperationType_TYPE_EVALUATE    OperatorOperationType_Type = 6
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
		6: "TYPE_EVALUATE",
	}
	OperatorOperationType_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_PATCH":       1,
		"TYPE_UPDATE":      2,
		"TYPE_CREATE":      3,
		"TYPE_DELETE":      4,
		"TYPE_DEREGISTER":  5,
		"TYPE_EVALUATE":    6,
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorLifecycleCapabilitiesRequest) Reset() {
	*x = OperatorLifecycleCapabilitiesRequest{}
	mi := &file_proto_operator_lifecycle_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorLifecycleCapabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorLifecycleCapabilitiesRequest) ProtoMessage() {}

func (x *OperatorLifecycleCapabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[0]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// This message is OPTIONAL, containing the list of resources
	// for which the lifecycle hook is called
	LifecycleCapabilities []*OperatorLifecycleCapabilities `protobuf:"bytes,1,rep,name=lifecycle_capabilities,json=lifecycleCapabilities,proto3" json:"lifecycle_capabilities,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *OperatorLifecycleCapabilitiesResponse) Reset() {
	*x = OperatorLifecycleCapabilitiesResponse{}
	mi := &file_proto_operator_lifecycle_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorLifecycleCapabilitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorLifecycleCapabilitiesResponse) ProtoMessage() {}

func (x *OperatorLifecycleCapabilitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[1]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// The Kubernetes resource group (such as "apps" or empty for core resources)
	Group string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	// The Kubernetes Kind name (such as "Cluster" or "Pod")
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// The operation type
	OperationTypes []*OperatorOperationType `protobuf:"bytes,3,rep,name=operation_types,json=operationTypes,proto3" json:"operation_types,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *OperatorLifecycleCapabilities) Reset() {
	*x = OperatorLifecycleCapabilities{}
	mi := &file_proto_operator_lifecycle_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorLifecycleCapabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorLifecycleCapabilities) ProtoMessage() {}

func (x *OperatorLifecycleCapabilities) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[2]
	if x != nil {
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
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Type          OperatorOperationType_Type `protobuf:"varint,1,opt,name=type,proto3,enum=cnpgi.operator_lifecycle.v1.OperatorOperationType_Type" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorOperationType) Reset() {
	*x = OperatorOperationType{}
	mi := &file_proto_operator_lifecycle_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorOperationType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorOperationType) ProtoMessage() {}

func (x *OperatorOperationType) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[3]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is REQUIRED.
	OperationType *OperatorOperationType `protobuf:"bytes,1,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	// This field is REQUIRED
	ClusterDefinition []byte `protobuf:"bytes,2,opt,name=cluster_definition,json=clusterDefinition,proto3" json:"cluster_definition,omitempty"`
	// This field is REQUIRED.
	ObjectDefinition []byte `protobuf:"bytes,3,opt,name=object_definition,json=objectDefinition,proto3" json:"object_definition,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *OperatorLifecycleRequest) Reset() {
	*x = OperatorLifecycleRequest{}
	mi := &file_proto_operator_lifecycle_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorLifecycleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorLifecycleRequest) ProtoMessage() {}

func (x *OperatorLifecycleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[4]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is OPTIONAL.
	JsonPatch     []byte `protobuf:"bytes,1,opt,name=json_patch,json=jsonPatch,proto3" json:"json_patch,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OperatorLifecycleResponse) Reset() {
	*x = OperatorLifecycleResponse{}
	mi := &file_proto_operator_lifecycle_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OperatorLifecycleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperatorLifecycleResponse) ProtoMessage() {}

func (x *OperatorLifecycleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operator_lifecycle_proto_msgTypes[5]
	if x != nil {
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

const file_proto_operator_lifecycle_proto_rawDesc = "" +
	"\n" +
	"\x1eproto/operator_lifecycle.proto\x12\x1bcnpgi.operator_lifecycle.v1\"&\n" +
	"$OperatorLifecycleCapabilitiesRequest\"\x9a\x01\n" +
	"%OperatorLifecycleCapabilitiesResponse\x12q\n" +
	"\x16lifecycle_capabilities\x18\x01 \x03(\v2:.cnpgi.operator_lifecycle.v1.OperatorLifecycleCapabilitiesR\x15lifecycleCapabilities\"\xa6\x01\n" +
	"\x1dOperatorLifecycleCapabilities\x12\x14\n" +
	"\x05group\x18\x01 \x01(\tR\x05group\x12\x12\n" +
	"\x04kind\x18\x02 \x01(\tR\x04kind\x12[\n" +
	"\x0foperation_types\x18\x03 \x03(\v22.cnpgi.operator_lifecycle.v1.OperatorOperationTypeR\x0eoperationTypes\"\xee\x01\n" +
	"\x15OperatorOperationType\x12K\n" +
	"\x04type\x18\x01 \x01(\x0e27.cnpgi.operator_lifecycle.v1.OperatorOperationType.TypeR\x04type\"\x87\x01\n" +
	"\x04Type\x12\x14\n" +
	"\x10TYPE_UNSPECIFIED\x10\x00\x12\x0e\n" +
	"\n" +
	"TYPE_PATCH\x10\x01\x12\x0f\n" +
	"\vTYPE_UPDATE\x10\x02\x12\x0f\n" +
	"\vTYPE_CREATE\x10\x03\x12\x0f\n" +
	"\vTYPE_DELETE\x10\x04\x12\x13\n" +
	"\x0fTYPE_DEREGISTER\x10\x05\x12\x11\n" +
	"\rTYPE_EVALUATE\x10\x06\"\xd1\x01\n" +
	"\x18OperatorLifecycleRequest\x12Y\n" +
	"\x0eoperation_type\x18\x01 \x01(\v22.cnpgi.operator_lifecycle.v1.OperatorOperationTypeR\roperationType\x12-\n" +
	"\x12cluster_definition\x18\x02 \x01(\fR\x11clusterDefinition\x12+\n" +
	"\x11object_definition\x18\x03 \x01(\fR\x10objectDefinition\":\n" +
	"\x19OperatorLifecycleResponse\x12\x1d\n" +
	"\n" +
	"json_patch\x18\x01 \x01(\fR\tjsonPatch2\xb3\x02\n" +
	"\x11OperatorLifecycle\x12\x9a\x01\n" +
	"\x0fGetCapabilities\x12A.cnpgi.operator_lifecycle.v1.OperatorLifecycleCapabilitiesRequest\x1aB.cnpgi.operator_lifecycle.v1.OperatorLifecycleCapabilitiesResponse\"\x00\x12\x80\x01\n" +
	"\rLifecycleHook\x125.cnpgi.operator_lifecycle.v1.OperatorLifecycleRequest\x1a6.cnpgi.operator_lifecycle.v1.OperatorLifecycleResponse\"\x00B0Z.github.com/cloudnative-pg/cnpg-i/pkg/lifecycleb\x06proto3"

var (
	file_proto_operator_lifecycle_proto_rawDescOnce sync.Once
	file_proto_operator_lifecycle_proto_rawDescData []byte
)

func file_proto_operator_lifecycle_proto_rawDescGZIP() []byte {
	file_proto_operator_lifecycle_proto_rawDescOnce.Do(func() {
		file_proto_operator_lifecycle_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_operator_lifecycle_proto_rawDesc), len(file_proto_operator_lifecycle_proto_rawDesc)))
	})
	return file_proto_operator_lifecycle_proto_rawDescData
}

var file_proto_operator_lifecycle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_operator_lifecycle_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_operator_lifecycle_proto_goTypes = []any{
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_operator_lifecycle_proto_rawDesc), len(file_proto_operator_lifecycle_proto_rawDesc)),
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
	file_proto_operator_lifecycle_proto_goTypes = nil
	file_proto_operator_lifecycle_proto_depIdxs = nil
}
