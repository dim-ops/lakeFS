//
//Run from lakeFS root:
//protoc --proto_path=pkg/plugins/diff --go_out=pkg/plugins/diff --go_opt=paths=source_relative \
//--go-grpc_out=pkg/plugins/diff --go-grpc_opt=paths=source_relative table_diff.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: table_diff.proto

package tablediff

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DiffType int32

const (
	DiffType_CHANGED DiffType = 0
	DiffType_CREATED DiffType = 1
	DiffType_DROPPED DiffType = 2
)

// Enum value maps for DiffType.
var (
	DiffType_name = map[int32]string{
		0: "CHANGED",
		1: "CREATED",
		2: "DROPPED",
	}
	DiffType_value = map[string]int32{
		"CHANGED": 0,
		"CREATED": 1,
		"DROPPED": 2,
	}
)

func (x DiffType) Enum() *DiffType {
	p := new(DiffType)
	*p = x
	return p
}

func (x DiffType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiffType) Descriptor() protoreflect.EnumDescriptor {
	return file_table_diff_proto_enumTypes[0].Descriptor()
}

func (DiffType) Type() protoreflect.EnumType {
	return &file_table_diff_proto_enumTypes[0]
}

func (x DiffType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiffType.Descriptor instead.
func (DiffType) EnumDescriptor() ([]byte, []int) {
	return file_table_diff_proto_rawDescGZIP(), []int{0}
}

type OperationType int32

const (
	OperationType_CREATE OperationType = 0
	OperationType_UPDATE OperationType = 1
	OperationType_DELETE OperationType = 2
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "CREATE",
		1: "UPDATE",
		2: "DELETE",
	}
	OperationType_value = map[string]int32{
		"CREATE": 0,
		"UPDATE": 1,
		"DELETE": 2,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_table_diff_proto_enumTypes[1].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_table_diff_proto_enumTypes[1]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_table_diff_proto_rawDescGZIP(), []int{1}
}

type TablePath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref  string `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *TablePath) Reset() {
	*x = TablePath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_diff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TablePath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TablePath) ProtoMessage() {}

func (x *TablePath) ProtoReflect() protoreflect.Message {
	mi := &file_table_diff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TablePath.ProtoReflect.Descriptor instead.
func (*TablePath) Descriptor() ([]byte, []int) {
	return file_table_diff_proto_rawDescGZIP(), []int{0}
}

func (x *TablePath) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *TablePath) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DiffProps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo           string     `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	LeftTablePath  *TablePath `protobuf:"bytes,2,opt,name=left_table_path,json=leftTablePath,proto3" json:"left_table_path,omitempty"`
	RightTablePath *TablePath `protobuf:"bytes,3,opt,name=right_table_path,json=rightTablePath,proto3" json:"right_table_path,omitempty"`
	BaseTablePath  *TablePath `protobuf:"bytes,4,opt,name=base_table_path,json=baseTablePath,proto3,oneof" json:"base_table_path,omitempty"`
}

func (x *DiffProps) Reset() {
	*x = DiffProps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_diff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiffProps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiffProps) ProtoMessage() {}

func (x *DiffProps) ProtoReflect() protoreflect.Message {
	mi := &file_table_diff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiffProps.ProtoReflect.Descriptor instead.
func (*DiffProps) Descriptor() ([]byte, []int) {
	return file_table_diff_proto_rawDescGZIP(), []int{1}
}

func (x *DiffProps) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *DiffProps) GetLeftTablePath() *TablePath {
	if x != nil {
		return x.LeftTablePath
	}
	return nil
}

func (x *DiffProps) GetRightTablePath() *TablePath {
	if x != nil {
		return x.RightTablePath
	}
	return nil
}

func (x *DiffProps) GetBaseTablePath() *TablePath {
	if x != nil {
		return x.BaseTablePath
	}
	return nil
}

type GatewayConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Secret   string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	Endpoint string `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *GatewayConfig) Reset() {
	*x = GatewayConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_diff_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayConfig) ProtoMessage() {}

func (x *GatewayConfig) ProtoReflect() protoreflect.Message {
	mi := &file_table_diff_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayConfig.ProtoReflect.Descriptor instead.
func (*GatewayConfig) Descriptor() ([]byte, []int) {
	return file_table_diff_proto_rawDescGZIP(), []int{2}
}

func (x *GatewayConfig) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GatewayConfig) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *GatewayConfig) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type DiffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Props         *DiffProps     `protobuf:"bytes,1,opt,name=props,proto3" json:"props,omitempty"`
	GatewayConfig *GatewayConfig `protobuf:"bytes,2,opt,name=gateway_config,json=gatewayConfig,proto3" json:"gateway_config,omitempty"`
}

func (x *DiffRequest) Reset() {
	*x = DiffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_diff_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiffRequest) ProtoMessage() {}

func (x *DiffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_table_diff_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiffRequest.ProtoReflect.Descriptor instead.
func (*DiffRequest) Descriptor() ([]byte, []int) {
	return file_table_diff_proto_rawDescGZIP(), []int{3}
}

func (x *DiffRequest) GetProps() *DiffProps {
	if x != nil {
		return x.Props
	}
	return nil
}

func (x *DiffRequest) GetGatewayConfig() *GatewayConfig {
	if x != nil {
		return x.GatewayConfig
	}
	return nil
}

type DiffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries  []*TableOperation `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	DiffType DiffType          `protobuf:"varint,2,opt,name=diffType,proto3,enum=diff.DiffType" json:"diffType,omitempty"`
}

func (x *DiffResponse) Reset() {
	*x = DiffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_diff_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiffResponse) ProtoMessage() {}

func (x *DiffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_table_diff_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiffResponse.ProtoReflect.Descriptor instead.
func (*DiffResponse) Descriptor() ([]byte, []int) {
	return file_table_diff_proto_rawDescGZIP(), []int{4}
}

func (x *DiffResponse) GetEntries() []*TableOperation {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *DiffResponse) GetDiffType() DiffType {
	if x != nil {
		return x.DiffType
	}
	return DiffType_CHANGED
}

type HistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path *TablePath `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *HistoryRequest) Reset() {
	*x = HistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_diff_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryRequest) ProtoMessage() {}

func (x *HistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_table_diff_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryRequest.ProtoReflect.Descriptor instead.
func (*HistoryRequest) Descriptor() ([]byte, []int) {
	return file_table_diff_proto_rawDescGZIP(), []int{5}
}

func (x *HistoryRequest) GetPath() *TablePath {
	if x != nil {
		return x.Path
	}
	return nil
}

type HistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*TableOperation `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *HistoryResponse) Reset() {
	*x = HistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_diff_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryResponse) ProtoMessage() {}

func (x *HistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_table_diff_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryResponse.ProtoReflect.Descriptor instead.
func (*HistoryResponse) Descriptor() ([]byte, []int) {
	return file_table_diff_proto_rawDescGZIP(), []int{6}
}

func (x *HistoryResponse) GetEntries() []*TableOperation {
	if x != nil {
		return x.Entries
	}
	return nil
}

// Example
// id: "2"
// timestamp: 2023-02-05T01:30:15.01Z
// operation: "DELETE"
// content: {
// "predicate": "[\"(spark_catalog.delta.lakefs://repo/branch/my-delta-lake-table/.`feature` < 5.0D)\"]"}
// }
type TableOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Operation     string                 `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	Content       map[string]string      `protobuf:"bytes,4,rep,name=content,proto3" json:"content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OperationType OperationType          `protobuf:"varint,5,opt,name=operation_type,json=operationType,proto3,enum=diff.OperationType" json:"operation_type,omitempty"`
}

func (x *TableOperation) Reset() {
	*x = TableOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_table_diff_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableOperation) ProtoMessage() {}

func (x *TableOperation) ProtoReflect() protoreflect.Message {
	mi := &file_table_diff_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableOperation.ProtoReflect.Descriptor instead.
func (*TableOperation) Descriptor() ([]byte, []int) {
	return file_table_diff_proto_rawDescGZIP(), []int{7}
}

func (x *TableOperation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TableOperation) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *TableOperation) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *TableOperation) GetContent() map[string]string {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *TableOperation) GetOperationType() OperationType {
	if x != nil {
		return x.OperationType
	}
	return OperationType_CREATE
}

var File_table_diff_proto protoreflect.FileDescriptor

var file_table_diff_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x64, 0x69, 0x66, 0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x09, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xe5, 0x01, 0x0a,
	0x09, 0x44, 0x69, 0x66, 0x66, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65,
	0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x37,
	0x0a, 0x0f, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x0d, 0x6c, 0x65, 0x66, 0x74, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x39, 0x0a, 0x10, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x52, 0x0e, 0x72, 0x69, 0x67, 0x68, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x3c, 0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x69,
	0x66, 0x66, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x48, 0x00, 0x52, 0x0d,
	0x62, 0x61, 0x73, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01,
	0x42, 0x12, 0x0a, 0x10, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x22, 0x55, 0x0a, 0x0d, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x70, 0x0a, 0x0b, 0x44,
	0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x72,
	0x6f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x69, 0x66, 0x66,
	0x2e, 0x44, 0x69, 0x66, 0x66, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70,
	0x73, 0x12, 0x3a, 0x0a, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x66, 0x66,
	0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x6a, 0x0a,
	0x0c, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2a, 0x0a,
	0x08, 0x64, 0x69, 0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x64, 0x69, 0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x22, 0x35, 0x0a, 0x0e, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x69, 0x66, 0x66,
	0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x41, 0x0a, 0x0f, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x22, 0xad, 0x02, 0x0a, 0x0e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x2a, 0x31, 0x0a, 0x08, 0x44, 0x69, 0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x52, 0x4f,
	0x50, 0x50, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x33, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x32, 0x7d, 0x0a, 0x0b, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x09, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x44, 0x69, 0x66, 0x66, 0x12, 0x11, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x44,
	0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x64, 0x69, 0x66,
	0x66, 0x2e, 0x44, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x77, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x2e,
	0x64, 0x69, 0x66, 0x66, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x2f, 0x6c, 0x61, 0x6b, 0x65, 0x66, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x69, 0x66, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_table_diff_proto_rawDescOnce sync.Once
	file_table_diff_proto_rawDescData = file_table_diff_proto_rawDesc
)

func file_table_diff_proto_rawDescGZIP() []byte {
	file_table_diff_proto_rawDescOnce.Do(func() {
		file_table_diff_proto_rawDescData = protoimpl.X.CompressGZIP(file_table_diff_proto_rawDescData)
	})
	return file_table_diff_proto_rawDescData
}

var file_table_diff_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_table_diff_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_table_diff_proto_goTypes = []interface{}{
	(DiffType)(0),                 // 0: diff.DiffType
	(OperationType)(0),            // 1: diff.OperationType
	(*TablePath)(nil),             // 2: diff.TablePath
	(*DiffProps)(nil),             // 3: diff.DiffProps
	(*GatewayConfig)(nil),         // 4: diff.GatewayConfig
	(*DiffRequest)(nil),           // 5: diff.DiffRequest
	(*DiffResponse)(nil),          // 6: diff.DiffResponse
	(*HistoryRequest)(nil),        // 7: diff.HistoryRequest
	(*HistoryResponse)(nil),       // 8: diff.HistoryResponse
	(*TableOperation)(nil),        // 9: diff.TableOperation
	nil,                           // 10: diff.TableOperation.ContentEntry
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_table_diff_proto_depIdxs = []int32{
	2,  // 0: diff.DiffProps.left_table_path:type_name -> diff.TablePath
	2,  // 1: diff.DiffProps.right_table_path:type_name -> diff.TablePath
	2,  // 2: diff.DiffProps.base_table_path:type_name -> diff.TablePath
	3,  // 3: diff.DiffRequest.props:type_name -> diff.DiffProps
	4,  // 4: diff.DiffRequest.gateway_config:type_name -> diff.GatewayConfig
	9,  // 5: diff.DiffResponse.entries:type_name -> diff.TableOperation
	0,  // 6: diff.DiffResponse.diffType:type_name -> diff.DiffType
	2,  // 7: diff.HistoryRequest.path:type_name -> diff.TablePath
	9,  // 8: diff.HistoryResponse.entries:type_name -> diff.TableOperation
	11, // 9: diff.TableOperation.timestamp:type_name -> google.protobuf.Timestamp
	10, // 10: diff.TableOperation.content:type_name -> diff.TableOperation.ContentEntry
	1,  // 11: diff.TableOperation.operation_type:type_name -> diff.OperationType
	5,  // 12: diff.TableDiffer.TableDiff:input_type -> diff.DiffRequest
	7,  // 13: diff.TableDiffer.ShowHistory:input_type -> diff.HistoryRequest
	6,  // 14: diff.TableDiffer.TableDiff:output_type -> diff.DiffResponse
	8,  // 15: diff.TableDiffer.ShowHistory:output_type -> diff.HistoryResponse
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_table_diff_proto_init() }
func file_table_diff_proto_init() {
	if File_table_diff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_table_diff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TablePath); i {
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
		file_table_diff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiffProps); i {
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
		file_table_diff_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayConfig); i {
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
		file_table_diff_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiffRequest); i {
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
		file_table_diff_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiffResponse); i {
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
		file_table_diff_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryRequest); i {
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
		file_table_diff_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryResponse); i {
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
		file_table_diff_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableOperation); i {
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
	file_table_diff_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_table_diff_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_table_diff_proto_goTypes,
		DependencyIndexes: file_table_diff_proto_depIdxs,
		EnumInfos:         file_table_diff_proto_enumTypes,
		MessageInfos:      file_table_diff_proto_msgTypes,
	}.Build()
	File_table_diff_proto = out.File
	file_table_diff_proto_rawDesc = nil
	file_table_diff_proto_goTypes = nil
	file_table_diff_proto_depIdxs = nil
}
