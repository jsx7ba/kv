// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0--rc1
// source: internal/proto/kv.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_OK    Status = 0
	Status_ERROR Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	Status_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_kv_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_internal_proto_kv_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{0}
}

type OpType int32

const (
	OpType_UNSPECIFIED OpType = 0
	OpType_ALL         OpType = 1
	OpType_PUT         OpType = 2
	OpType_DELETE      OpType = 3
)

// Enum value maps for OpType.
var (
	OpType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "ALL",
		2: "PUT",
		3: "DELETE",
	}
	OpType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"ALL":         1,
		"PUT":         2,
		"DELETE":      3,
	}
)

func (x OpType) Enum() *OpType {
	p := new(OpType)
	*p = x
	return p
}

func (x OpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_kv_proto_enumTypes[1].Descriptor()
}

func (OpType) Type() protoreflect.EnumType {
	return &file_internal_proto_kv_proto_enumTypes[1]
}

func (x OpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpType.Descriptor instead.
func (OpType) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{1}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status     `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	Value  *anypb.Any `protobuf:"bytes,2,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *GetResponse) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *anypb.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{3}
}

func (x *PutRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutRequest) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type WatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	WatchType OpType `protobuf:"varint,2,opt,name=watchType,proto3,enum=OpType" json:"watchType,omitempty"`
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{5}
}

func (x *WatchRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *WatchRequest) GetWatchType() OpType {
	if x != nil {
		return x.WatchType
	}
	return OpType_UNSPECIFIED
}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WatchType OpType     `protobuf:"varint,1,opt,name=watchType,proto3,enum=OpType" json:"watchType,omitempty"`
	Key       string     `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value     *anypb.Any `protobuf:"bytes,3,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{6}
}

func (x *WatchResponse) GetWatchType() OpType {
	if x != nil {
		return x.WatchType
	}
	return OpType_UNSPECIFIED
}

func (x *WatchResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *WatchResponse) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type StringSliceWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *StringSliceWrapper) Reset() {
	*x = StringSliceWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringSliceWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringSliceWrapper) ProtoMessage() {}

func (x *StringSliceWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringSliceWrapper.ProtoReflect.Descriptor instead.
func (*StringSliceWrapper) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{7}
}

func (x *StringSliceWrapper) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type Int32SliceWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []int32 `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *Int32SliceWrapper) Reset() {
	*x = Int32SliceWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int32SliceWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32SliceWrapper) ProtoMessage() {}

func (x *Int32SliceWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int32SliceWrapper.ProtoReflect.Descriptor instead.
func (*Int32SliceWrapper) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{8}
}

func (x *Int32SliceWrapper) GetValue() []int32 {
	if x != nil {
		return x.Value
	}
	return nil
}

type Int64SliceWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []int64 `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *Int64SliceWrapper) Reset() {
	*x = Int64SliceWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64SliceWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64SliceWrapper) ProtoMessage() {}

func (x *Int64SliceWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64SliceWrapper.ProtoReflect.Descriptor instead.
func (*Int64SliceWrapper) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{9}
}

func (x *Int64SliceWrapper) GetValue() []int64 {
	if x != nil {
		return x.Value
	}
	return nil
}

type Float32SliceWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []float32 `protobuf:"fixed32,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *Float32SliceWrapper) Reset() {
	*x = Float32SliceWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Float32SliceWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Float32SliceWrapper) ProtoMessage() {}

func (x *Float32SliceWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Float32SliceWrapper.ProtoReflect.Descriptor instead.
func (*Float32SliceWrapper) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{10}
}

func (x *Float32SliceWrapper) GetValue() []float32 {
	if x != nil {
		return x.Value
	}
	return nil
}

type Float64SliceWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []float64 `protobuf:"fixed64,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *Float64SliceWrapper) Reset() {
	*x = Float64SliceWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_kv_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Float64SliceWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Float64SliceWrapper) ProtoMessage() {}

func (x *Float64SliceWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_kv_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Float64SliceWrapper.ProtoReflect.Descriptor instead.
func (*Float64SliceWrapper) Descriptor() ([]byte, []int) {
	return file_internal_proto_kv_proto_rawDescGZIP(), []int{11}
}

func (x *Float64SliceWrapper) GetValue() []float64 {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_internal_proto_kv_proto protoreflect.FileDescriptor

var file_internal_proto_kv_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x69, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x2b, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4a, 0x0a, 0x0a,
	0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x47, 0x0a, 0x0c, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a,
	0x09, 0x77, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x07, 0x2e, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x77, 0x61, 0x74, 0x63, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x4f, 0x70, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x77, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2a, 0x0a, 0x12, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x29, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x53,
	0x6c, 0x69, 0x63, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x29, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x57,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2b, 0x0a, 0x13,
	0x46, 0x6c, 0x6f, 0x61, 0x74, 0x33, 0x32, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x57, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2b, 0x0a, 0x13, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x36, 0x34, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x1b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x01, 0x2a, 0x37, 0x0a, 0x06, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x32, 0x94, 0x01, 0x0a,
	0x02, 0x4b, 0x56, 0x12, 0x1d, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x0b, 0x2e, 0x50, 0x75, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0b, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x0d, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_kv_proto_rawDescOnce sync.Once
	file_internal_proto_kv_proto_rawDescData = file_internal_proto_kv_proto_rawDesc
)

func file_internal_proto_kv_proto_rawDescGZIP() []byte {
	file_internal_proto_kv_proto_rawDescOnce.Do(func() {
		file_internal_proto_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_kv_proto_rawDescData)
	})
	return file_internal_proto_kv_proto_rawDescData
}

var file_internal_proto_kv_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_internal_proto_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_internal_proto_kv_proto_goTypes = []any{
	(Status)(0),                 // 0: Status
	(OpType)(0),                 // 1: OpType
	(*GetRequest)(nil),          // 2: GetRequest
	(*GetResponse)(nil),         // 3: GetResponse
	(*Response)(nil),            // 4: Response
	(*PutRequest)(nil),          // 5: PutRequest
	(*DeleteRequest)(nil),       // 6: DeleteRequest
	(*WatchRequest)(nil),        // 7: WatchRequest
	(*WatchResponse)(nil),       // 8: WatchResponse
	(*StringSliceWrapper)(nil),  // 9: StringSliceWrapper
	(*Int32SliceWrapper)(nil),   // 10: Int32SliceWrapper
	(*Int64SliceWrapper)(nil),   // 11: Int64SliceWrapper
	(*Float32SliceWrapper)(nil), // 12: Float32SliceWrapper
	(*Float64SliceWrapper)(nil), // 13: Float64SliceWrapper
	(*anypb.Any)(nil),           // 14: google.protobuf.Any
}
var file_internal_proto_kv_proto_depIdxs = []int32{
	0,  // 0: GetResponse.status:type_name -> Status
	14, // 1: GetResponse.value:type_name -> google.protobuf.Any
	0,  // 2: Response.status:type_name -> Status
	14, // 3: PutRequest.value:type_name -> google.protobuf.Any
	1,  // 4: WatchRequest.watchType:type_name -> OpType
	1,  // 5: WatchResponse.watchType:type_name -> OpType
	14, // 6: WatchResponse.value:type_name -> google.protobuf.Any
	5,  // 7: KV.Put:input_type -> PutRequest
	2,  // 8: KV.Get:input_type -> GetRequest
	6,  // 9: KV.Delete:input_type -> DeleteRequest
	7,  // 10: KV.Watch:input_type -> WatchRequest
	4,  // 11: KV.Put:output_type -> Response
	3,  // 12: KV.Get:output_type -> GetResponse
	4,  // 13: KV.Delete:output_type -> Response
	8,  // 14: KV.Watch:output_type -> WatchResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_internal_proto_kv_proto_init() }
func file_internal_proto_kv_proto_init() {
	if File_internal_proto_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_kv_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetRequest); i {
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
		file_internal_proto_kv_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetResponse); i {
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
		file_internal_proto_kv_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Response); i {
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
		file_internal_proto_kv_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PutRequest); i {
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
		file_internal_proto_kv_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
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
		file_internal_proto_kv_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*WatchRequest); i {
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
		file_internal_proto_kv_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*WatchResponse); i {
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
		file_internal_proto_kv_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*StringSliceWrapper); i {
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
		file_internal_proto_kv_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Int32SliceWrapper); i {
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
		file_internal_proto_kv_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Int64SliceWrapper); i {
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
		file_internal_proto_kv_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*Float32SliceWrapper); i {
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
		file_internal_proto_kv_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*Float64SliceWrapper); i {
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
	file_internal_proto_kv_proto_msgTypes[1].OneofWrappers = []any{}
	file_internal_proto_kv_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_kv_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_kv_proto_goTypes,
		DependencyIndexes: file_internal_proto_kv_proto_depIdxs,
		EnumInfos:         file_internal_proto_kv_proto_enumTypes,
		MessageInfos:      file_internal_proto_kv_proto_msgTypes,
	}.Build()
	File_internal_proto_kv_proto = out.File
	file_internal_proto_kv_proto_rawDesc = nil
	file_internal_proto_kv_proto_goTypes = nil
	file_internal_proto_kv_proto_depIdxs = nil
}
