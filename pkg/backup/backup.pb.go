// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.3
// source: proto/backup.proto

package backup

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

type BackupCapability_RPC_Type int32

const (
	BackupCapability_RPC_TYPE_UNSPECIFIED BackupCapability_RPC_Type = 0
	// TYPE_BACKUP indicates that the Plugin is able to
	// take physical backups. This feature is required for every
	// plugin exposing the Backup service
	BackupCapability_RPC_TYPE_BACKUP BackupCapability_RPC_Type = 1
)

// Enum value maps for BackupCapability_RPC_Type.
var (
	BackupCapability_RPC_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_BACKUP",
	}
	BackupCapability_RPC_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_BACKUP":      1,
	}
)

func (x BackupCapability_RPC_Type) Enum() *BackupCapability_RPC_Type {
	p := new(BackupCapability_RPC_Type)
	*p = x
	return p
}

func (x BackupCapability_RPC_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BackupCapability_RPC_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_backup_proto_enumTypes[0].Descriptor()
}

func (BackupCapability_RPC_Type) Type() protoreflect.EnumType {
	return &file_proto_backup_proto_enumTypes[0]
}

func (x BackupCapability_RPC_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BackupCapability_RPC_Type.Descriptor instead.
func (BackupCapability_RPC_Type) EnumDescriptor() ([]byte, []int) {
	return file_proto_backup_proto_rawDescGZIP(), []int{2, 0, 0}
}

type BackupCapabilitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BackupCapabilitiesRequest) Reset() {
	*x = BackupCapabilitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupCapabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupCapabilitiesRequest) ProtoMessage() {}

func (x *BackupCapabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupCapabilitiesRequest.ProtoReflect.Descriptor instead.
func (*BackupCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_proto_backup_proto_rawDescGZIP(), []int{0}
}

type BackupCapabilitiesResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All the capabilities that the controller service supports. This
	// field is OPTIONAL.
	Capabilities []*BackupCapability `protobuf:"bytes,1,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
}

func (x *BackupCapabilitiesResult) Reset() {
	*x = BackupCapabilitiesResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupCapabilitiesResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupCapabilitiesResult) ProtoMessage() {}

func (x *BackupCapabilitiesResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupCapabilitiesResult.ProtoReflect.Descriptor instead.
func (*BackupCapabilitiesResult) Descriptor() ([]byte, []int) {
	return file_proto_backup_proto_rawDescGZIP(), []int{1}
}

func (x *BackupCapabilitiesResult) GetCapabilities() []*BackupCapability {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

type BackupCapability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*BackupCapability_Rpc
	Type isBackupCapability_Type `protobuf_oneof:"type"`
}

func (x *BackupCapability) Reset() {
	*x = BackupCapability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupCapability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupCapability) ProtoMessage() {}

func (x *BackupCapability) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupCapability.ProtoReflect.Descriptor instead.
func (*BackupCapability) Descriptor() ([]byte, []int) {
	return file_proto_backup_proto_rawDescGZIP(), []int{2}
}

func (m *BackupCapability) GetType() isBackupCapability_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *BackupCapability) GetRpc() *BackupCapability_RPC {
	if x, ok := x.GetType().(*BackupCapability_Rpc); ok {
		return x.Rpc
	}
	return nil
}

type isBackupCapability_Type interface {
	isBackupCapability_Type()
}

type BackupCapability_Rpc struct {
	Rpc *BackupCapability_RPC `protobuf:"bytes,1,opt,name=rpc,proto3,oneof"`
}

func (*BackupCapability_Rpc) isBackupCapability_Type() {}

type BackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field is REQUIRED. Value of this field is the JSON
	// serialization of the Cluster being backed up
	ClusterDefinition []byte `protobuf:"bytes,1,opt,name=cluster_definition,json=clusterDefinition,proto3" json:"cluster_definition,omitempty"`
	// This field is REQUIRED. Value of this field is the JSON
	// serialization of the Backup that is being taken
	BackupDefinition []byte `protobuf:"bytes,2,opt,name=backup_definition,json=backupDefinition,proto3" json:"backup_definition,omitempty"`
	// This field is OPTIONAL. Value of this field is the configuration
	// of this backup as set in the Backup or in the ScheduledBackup object
	Parameters map[string]string `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BackupRequest) Reset() {
	*x = BackupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backup_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupRequest) ProtoMessage() {}

func (x *BackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backup_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupRequest.ProtoReflect.Descriptor instead.
func (*BackupRequest) Descriptor() ([]byte, []int) {
	return file_proto_backup_proto_rawDescGZIP(), []int{3}
}

func (x *BackupRequest) GetClusterDefinition() []byte {
	if x != nil {
		return x.ClusterDefinition
	}
	return nil
}

func (x *BackupRequest) GetBackupDefinition() []byte {
	if x != nil {
		return x.BackupDefinition
	}
	return nil
}

func (x *BackupRequest) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type BackupResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field is REQUIRED and contains a machine-readable ID of the
	// backup that is being taken
	BackupId string `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
	// This field is OPTIONAL and contains a human-readable name of the
	// backup that is being taken
	BackupName string `protobuf:"bytes,2,opt,name=backup_name,json=backupName,proto3" json:"backup_name,omitempty"`
	// This field is REQUIRED and contains the Unix timestamp of the start
	// time of the backup
	StartedAt int64 `protobuf:"varint,3,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// This field is REQUIRED and contains the Unix timestamp of the end
	// time of the backup
	StoppedAt int64 `protobuf:"varint,4,opt,name=stopped_at,json=stoppedAt,proto3" json:"stopped_at,omitempty"`
	// This field is OPTIONAL and contains the current WAL when the backup was started
	BeginWal string `protobuf:"bytes,5,opt,name=begin_wal,json=beginWal,proto3" json:"begin_wal,omitempty"`
	// This field is OPTIONAL and contains the current WAL at the end of the backup
	EndWal string `protobuf:"bytes,6,opt,name=end_wal,json=endWal,proto3" json:"end_wal,omitempty"`
	// This field is OPTIONAL and contains the current LSN record when the backup was started
	BeginLsn string `protobuf:"bytes,7,opt,name=begin_lsn,json=beginLsn,proto3" json:"begin_lsn,omitempty"`
	// This field is OPTIONAL and contains the current LSN record when the backup has finished
	EndLsn string `protobuf:"bytes,8,opt,name=end_lsn,json=endLsn,proto3" json:"end_lsn,omitempty"`
	// This field is OPTIONAL and contains the backup label of the backup that have been taken
	BackupLabelFile []byte `protobuf:"bytes,9,opt,name=backup_label_file,json=backupLabelFile,proto3" json:"backup_label_file,omitempty"`
	// This field is OPTIONAL and contains the tablespace map of the backup that have been taken
	TablespaceMapFile []byte `protobuf:"bytes,10,opt,name=tablespace_map_file,json=tablespaceMapFile,proto3" json:"tablespace_map_file,omitempty"`
	// This field is OPTIONAL and contains the ID of the instance that have been backed up
	InstanceId string `protobuf:"bytes,11,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// This field is REQUIRED and is set to true for online/hot backups and to false otherwise.
	Online bool `protobuf:"varint,12,opt,name=online,proto3" json:"online,omitempty"`
	// This field is OPTIONAL and contains all the plugin specific information that needs to be stored
	Metadata map[string]string `protobuf:"bytes,13,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// This field is OPTIONAL and contains the name of the object storage server
	ServerName string `protobuf:"bytes,14,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
}

func (x *BackupResult) Reset() {
	*x = BackupResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backup_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupResult) ProtoMessage() {}

func (x *BackupResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backup_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupResult.ProtoReflect.Descriptor instead.
func (*BackupResult) Descriptor() ([]byte, []int) {
	return file_proto_backup_proto_rawDescGZIP(), []int{4}
}

func (x *BackupResult) GetBackupId() string {
	if x != nil {
		return x.BackupId
	}
	return ""
}

func (x *BackupResult) GetBackupName() string {
	if x != nil {
		return x.BackupName
	}
	return ""
}

func (x *BackupResult) GetStartedAt() int64 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *BackupResult) GetStoppedAt() int64 {
	if x != nil {
		return x.StoppedAt
	}
	return 0
}

func (x *BackupResult) GetBeginWal() string {
	if x != nil {
		return x.BeginWal
	}
	return ""
}

func (x *BackupResult) GetEndWal() string {
	if x != nil {
		return x.EndWal
	}
	return ""
}

func (x *BackupResult) GetBeginLsn() string {
	if x != nil {
		return x.BeginLsn
	}
	return ""
}

func (x *BackupResult) GetEndLsn() string {
	if x != nil {
		return x.EndLsn
	}
	return ""
}

func (x *BackupResult) GetBackupLabelFile() []byte {
	if x != nil {
		return x.BackupLabelFile
	}
	return nil
}

func (x *BackupResult) GetTablespaceMapFile() []byte {
	if x != nil {
		return x.TablespaceMapFile
	}
	return nil
}

func (x *BackupResult) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *BackupResult) GetOnline() bool {
	if x != nil {
		return x.Online
	}
	return false
}

func (x *BackupResult) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *BackupResult) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

type BackupCapability_RPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type BackupCapability_RPC_Type `protobuf:"varint,1,opt,name=type,proto3,enum=cnpgi.backup.v1.BackupCapability_RPC_Type" json:"type,omitempty"`
}

func (x *BackupCapability_RPC) Reset() {
	*x = BackupCapability_RPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_backup_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupCapability_RPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupCapability_RPC) ProtoMessage() {}

func (x *BackupCapability_RPC) ProtoReflect() protoreflect.Message {
	mi := &file_proto_backup_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupCapability_RPC.ProtoReflect.Descriptor instead.
func (*BackupCapability_RPC) Descriptor() ([]byte, []int) {
	return file_proto_backup_proto_rawDescGZIP(), []int{2, 0}
}

func (x *BackupCapability_RPC) GetType() BackupCapability_RPC_Type {
	if x != nil {
		return x.Type
	}
	return BackupCapability_RPC_TYPE_UNSPECIFIED
}

var File_proto_backup_proto protoreflect.FileDescriptor

var file_proto_backup_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x2e, 0x76, 0x31, 0x22, 0x1b, 0x0a, 0x19, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x61, 0x0a, 0x18, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x45,
	0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0xcb, 0x01, 0x0a, 0x10, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x03, 0x72, 0x70,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x50, 0x43, 0x48, 0x00,
	0x52, 0x03, 0x72, 0x70, 0x63, 0x1a, 0x74, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x3e, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6e, 0x70,
	0x67, 0x69, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x50,
	0x43, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2d, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0xfa, 0x01, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x11, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x11, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x10, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xb2, 0x04, 0x0a, 0x0c, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x77, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x57, 0x61, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e,
	0x64, 0x5f, 0x77, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x64,
	0x57, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x6c, 0x73, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x4c, 0x73, 0x6e,
	0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x5f, 0x6c, 0x73, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x6e, 0x64, 0x4c, 0x73, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x11, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4d, 0x61,
	0x70, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x47,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xbf, 0x01, 0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x12, 0x6a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06,
	0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x1e, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6e, 0x70, 0x67, 0x69, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x2d, 0x70, 0x67, 0x2f, 0x63, 0x6e, 0x70, 0x67, 0x2d, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_backup_proto_rawDescOnce sync.Once
	file_proto_backup_proto_rawDescData = file_proto_backup_proto_rawDesc
)

func file_proto_backup_proto_rawDescGZIP() []byte {
	file_proto_backup_proto_rawDescOnce.Do(func() {
		file_proto_backup_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_backup_proto_rawDescData)
	})
	return file_proto_backup_proto_rawDescData
}

var file_proto_backup_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_backup_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_backup_proto_goTypes = []interface{}{
	(BackupCapability_RPC_Type)(0),    // 0: cnpgi.backup.v1.BackupCapability.RPC.Type
	(*BackupCapabilitiesRequest)(nil), // 1: cnpgi.backup.v1.BackupCapabilitiesRequest
	(*BackupCapabilitiesResult)(nil),  // 2: cnpgi.backup.v1.BackupCapabilitiesResult
	(*BackupCapability)(nil),          // 3: cnpgi.backup.v1.BackupCapability
	(*BackupRequest)(nil),             // 4: cnpgi.backup.v1.BackupRequest
	(*BackupResult)(nil),              // 5: cnpgi.backup.v1.BackupResult
	(*BackupCapability_RPC)(nil),      // 6: cnpgi.backup.v1.BackupCapability.RPC
	nil,                               // 7: cnpgi.backup.v1.BackupRequest.ParametersEntry
	nil,                               // 8: cnpgi.backup.v1.BackupResult.MetadataEntry
}
var file_proto_backup_proto_depIdxs = []int32{
	3, // 0: cnpgi.backup.v1.BackupCapabilitiesResult.capabilities:type_name -> cnpgi.backup.v1.BackupCapability
	6, // 1: cnpgi.backup.v1.BackupCapability.rpc:type_name -> cnpgi.backup.v1.BackupCapability.RPC
	7, // 2: cnpgi.backup.v1.BackupRequest.parameters:type_name -> cnpgi.backup.v1.BackupRequest.ParametersEntry
	8, // 3: cnpgi.backup.v1.BackupResult.metadata:type_name -> cnpgi.backup.v1.BackupResult.MetadataEntry
	0, // 4: cnpgi.backup.v1.BackupCapability.RPC.type:type_name -> cnpgi.backup.v1.BackupCapability.RPC.Type
	1, // 5: cnpgi.backup.v1.Backup.GetCapabilities:input_type -> cnpgi.backup.v1.BackupCapabilitiesRequest
	4, // 6: cnpgi.backup.v1.Backup.Backup:input_type -> cnpgi.backup.v1.BackupRequest
	2, // 7: cnpgi.backup.v1.Backup.GetCapabilities:output_type -> cnpgi.backup.v1.BackupCapabilitiesResult
	5, // 8: cnpgi.backup.v1.Backup.Backup:output_type -> cnpgi.backup.v1.BackupResult
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_backup_proto_init() }
func file_proto_backup_proto_init() {
	if File_proto_backup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_backup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupCapabilitiesRequest); i {
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
		file_proto_backup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupCapabilitiesResult); i {
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
		file_proto_backup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupCapability); i {
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
		file_proto_backup_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupRequest); i {
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
		file_proto_backup_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupResult); i {
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
		file_proto_backup_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupCapability_RPC); i {
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
	file_proto_backup_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BackupCapability_Rpc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_backup_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_backup_proto_goTypes,
		DependencyIndexes: file_proto_backup_proto_depIdxs,
		EnumInfos:         file_proto_backup_proto_enumTypes,
		MessageInfos:      file_proto_backup_proto_msgTypes,
	}.Build()
	File_proto_backup_proto = out.File
	file_proto_backup_proto_rawDesc = nil
	file_proto_backup_proto_goTypes = nil
	file_proto_backup_proto_depIdxs = nil
}
