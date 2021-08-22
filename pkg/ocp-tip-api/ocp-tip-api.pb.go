// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/ocp-tip-api/ocp-tip-api.proto

package ocp_tip_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateTipV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ProblemId uint64 `protobuf:"varint,2,opt,name=ProblemId,proto3" json:"ProblemId,omitempty"`
	Text      string `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *CreateTipV1Request) Reset() {
	*x = CreateTipV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTipV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTipV1Request) ProtoMessage() {}

func (x *CreateTipV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTipV1Request.ProtoReflect.Descriptor instead.
func (*CreateTipV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTipV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateTipV1Request) GetProblemId() uint64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *CreateTipV1Request) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CreateTipV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *CreateTipV1Response) Reset() {
	*x = CreateTipV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTipV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTipV1Response) ProtoMessage() {}

func (x *CreateTipV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTipV1Response.ProtoReflect.Descriptor instead.
func (*CreateTipV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTipV1Response) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MultiCreateTipV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tips      []*CreateTipV1Request `protobuf:"bytes,1,rep,name=Tips,proto3" json:"Tips,omitempty"`
	BatchSize uint64                `protobuf:"varint,2,opt,name=BatchSize,proto3" json:"BatchSize,omitempty"`
}

func (x *MultiCreateTipV1Request) Reset() {
	*x = MultiCreateTipV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiCreateTipV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiCreateTipV1Request) ProtoMessage() {}

func (x *MultiCreateTipV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiCreateTipV1Request.ProtoReflect.Descriptor instead.
func (*MultiCreateTipV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{2}
}

func (x *MultiCreateTipV1Request) GetTips() []*CreateTipV1Request {
	if x != nil {
		return x.Tips
	}
	return nil
}

func (x *MultiCreateTipV1Request) GetBatchSize() uint64 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

type MultiCreateTipV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids            []uint64              `protobuf:"varint,1,rep,packed,name=Ids,proto3" json:"Ids,omitempty"`
	NotCreatedTips []*CreateTipV1Request `protobuf:"bytes,2,rep,name=NotCreatedTips,proto3" json:"NotCreatedTips,omitempty"`
}

func (x *MultiCreateTipV1Response) Reset() {
	*x = MultiCreateTipV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiCreateTipV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiCreateTipV1Response) ProtoMessage() {}

func (x *MultiCreateTipV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiCreateTipV1Response.ProtoReflect.Descriptor instead.
func (*MultiCreateTipV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{3}
}

func (x *MultiCreateTipV1Response) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *MultiCreateTipV1Response) GetNotCreatedTips() []*CreateTipV1Request {
	if x != nil {
		return x.NotCreatedTips
	}
	return nil
}

type UpdateTipV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId    uint64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ProblemId uint64 `protobuf:"varint,3,opt,name=ProblemId,proto3" json:"ProblemId,omitempty"`
	Text      string `protobuf:"bytes,4,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *UpdateTipV1Request) Reset() {
	*x = UpdateTipV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTipV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTipV1Request) ProtoMessage() {}

func (x *UpdateTipV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTipV1Request.ProtoReflect.Descriptor instead.
func (*UpdateTipV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTipV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateTipV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateTipV1Request) GetProblemId() uint64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *UpdateTipV1Request) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type UpdateTipV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateTipV1Response) Reset() {
	*x = UpdateTipV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTipV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTipV1Response) ProtoMessage() {}

func (x *UpdateTipV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTipV1Response.ProtoReflect.Descriptor instead.
func (*UpdateTipV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{5}
}

type DescribeTipV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DescribeTipV1Request) Reset() {
	*x = DescribeTipV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTipV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTipV1Request) ProtoMessage() {}

func (x *DescribeTipV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTipV1Request.ProtoReflect.Descriptor instead.
func (*DescribeTipV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeTipV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeTipV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tip *TipV1 `protobuf:"bytes,1,opt,name=Tip,proto3" json:"Tip,omitempty"`
}

func (x *DescribeTipV1Response) Reset() {
	*x = DescribeTipV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeTipV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTipV1Response) ProtoMessage() {}

func (x *DescribeTipV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTipV1Response.ProtoReflect.Descriptor instead.
func (*DescribeTipV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeTipV1Response) GetTip() *TipV1 {
	if x != nil {
		return x.Tip
	}
	return nil
}

type ListTipsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *ListTipsV1Request) Reset() {
	*x = ListTipsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTipsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTipsV1Request) ProtoMessage() {}

func (x *ListTipsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTipsV1Request.ProtoReflect.Descriptor instead.
func (*ListTipsV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{8}
}

func (x *ListTipsV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTipsV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListTipsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tips []*TipV1 `protobuf:"bytes,1,rep,name=Tips,proto3" json:"Tips,omitempty"`
}

func (x *ListTipsV1Response) Reset() {
	*x = ListTipsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTipsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTipsV1Response) ProtoMessage() {}

func (x *ListTipsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTipsV1Response.ProtoReflect.Descriptor instead.
func (*ListTipsV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{9}
}

func (x *ListTipsV1Response) GetTips() []*TipV1 {
	if x != nil {
		return x.Tips
	}
	return nil
}

type RemoveTipV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *RemoveTipV1Request) Reset() {
	*x = RemoveTipV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTipV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTipV1Request) ProtoMessage() {}

func (x *RemoveTipV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTipV1Request.ProtoReflect.Descriptor instead.
func (*RemoveTipV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{10}
}

func (x *RemoveTipV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveTipV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool `protobuf:"varint,1,opt,name=Found,proto3" json:"Found,omitempty"`
}

func (x *RemoveTipV1Response) Reset() {
	*x = RemoveTipV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveTipV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveTipV1Response) ProtoMessage() {}

func (x *RemoveTipV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveTipV1Response.ProtoReflect.Descriptor instead.
func (*RemoveTipV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{11}
}

func (x *RemoveTipV1Response) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type TipV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId    uint64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	ProblemId uint64 `protobuf:"varint,3,opt,name=ProblemId,proto3" json:"ProblemId,omitempty"`
	Text      string `protobuf:"bytes,4,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *TipV1) Reset() {
	*x = TipV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TipV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TipV1) ProtoMessage() {}

func (x *TipV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TipV1.ProtoReflect.Descriptor instead.
func (*TipV1) Descriptor() ([]byte, []int) {
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP(), []int{12}
}

func (x *TipV1) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TipV1) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TipV1) GetProblemId() uint64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *TipV1) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_api_ocp_tip_api_ocp_tip_api_proto protoreflect.FileDescriptor

var file_api_ocp_tip_api_ocp_tip_api_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x74, 0x69, 0x70, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x74, 0x69, 0x70, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61, 0x70, 0x69,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x70, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x17, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x54, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x54, 0x69, 0x70, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x75, 0x0a, 0x18, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x03, 0x49, 0x64, 0x73, 0x12, 0x47, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0e, 0x4e, 0x6f, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x70, 0x73, 0x22,
	0x89, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x09, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2f, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69,
	0x70, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x02, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54,
	0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x03,
	0x54, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x74, 0x69, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x03, 0x54,
	0x69, 0x70, 0x22, 0x4a, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x70, 0x73, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3c,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x70, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x54, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x04, 0x54, 0x69, 0x70, 0x73, 0x22, 0x2d, 0x0a, 0x12,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x13, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x61, 0x0a, 0x05, 0x54, 0x69, 0x70, 0x56,
	0x31, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x50, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x32, 0x9b, 0x05, 0x0a, 0x09,
	0x4f, 0x63, 0x70, 0x54, 0x69, 0x70, 0x41, 0x70, 0x69, 0x12, 0x65, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x12, 0x1f, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74,
	0x69, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x74, 0x69, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x70, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x70, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x81, 0x01, 0x0a, 0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x70, 0x56, 0x31, 0x12, 0x24, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6f, 0x63,
	0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x69, 0x70, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x70, 0x56, 0x31, 0x12, 0x1f, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x70, 0x73, 0x2f, 0x7b, 0x49, 0x64, 0x7d, 0x3a, 0x01, 0x2a,
	0x12, 0x6d, 0x0a, 0x0d, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69, 0x70, 0x56,
	0x31, 0x12, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x70, 0x73, 0x2f, 0x7b, 0x49, 0x64, 0x7d, 0x12,
	0x5f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x69, 0x70, 0x73, 0x56, 0x31, 0x12, 0x1e, 0x2e,
	0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x69, 0x70, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x69, 0x70, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x70, 0x73,
	0x12, 0x67, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x12,
	0x1f, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x74, 0x69, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x70, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x69, 0x70, 0x73, 0x2f, 0x7b, 0x49, 0x64, 0x7d, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f,
	0x63, 0x70, 0x2d, 0x74, 0x69, 0x70, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f,
	0x63, 0x70, 0x2d, 0x74, 0x69, 0x70, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x63, 0x70, 0x5f, 0x74,
	0x69, 0x70, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ocp_tip_api_ocp_tip_api_proto_rawDescOnce sync.Once
	file_api_ocp_tip_api_ocp_tip_api_proto_rawDescData = file_api_ocp_tip_api_ocp_tip_api_proto_rawDesc
)

func file_api_ocp_tip_api_ocp_tip_api_proto_rawDescGZIP() []byte {
	file_api_ocp_tip_api_ocp_tip_api_proto_rawDescOnce.Do(func() {
		file_api_ocp_tip_api_ocp_tip_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ocp_tip_api_ocp_tip_api_proto_rawDescData)
	})
	return file_api_ocp_tip_api_ocp_tip_api_proto_rawDescData
}

var file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_ocp_tip_api_ocp_tip_api_proto_goTypes = []interface{}{
	(*CreateTipV1Request)(nil),       // 0: ocp.tip.api.CreateTipV1Request
	(*CreateTipV1Response)(nil),      // 1: ocp.tip.api.CreateTipV1Response
	(*MultiCreateTipV1Request)(nil),  // 2: ocp.tip.api.MultiCreateTipV1Request
	(*MultiCreateTipV1Response)(nil), // 3: ocp.tip.api.MultiCreateTipV1Response
	(*UpdateTipV1Request)(nil),       // 4: ocp.tip.api.UpdateTipV1Request
	(*UpdateTipV1Response)(nil),      // 5: ocp.tip.api.UpdateTipV1Response
	(*DescribeTipV1Request)(nil),     // 6: ocp.tip.api.DescribeTipV1Request
	(*DescribeTipV1Response)(nil),    // 7: ocp.tip.api.DescribeTipV1Response
	(*ListTipsV1Request)(nil),        // 8: ocp.tip.api.ListTipsV1Request
	(*ListTipsV1Response)(nil),       // 9: ocp.tip.api.ListTipsV1Response
	(*RemoveTipV1Request)(nil),       // 10: ocp.tip.api.RemoveTipV1Request
	(*RemoveTipV1Response)(nil),      // 11: ocp.tip.api.RemoveTipV1Response
	(*TipV1)(nil),                    // 12: ocp.tip.api.TipV1
}
var file_api_ocp_tip_api_ocp_tip_api_proto_depIdxs = []int32{
	0,  // 0: ocp.tip.api.MultiCreateTipV1Request.Tips:type_name -> ocp.tip.api.CreateTipV1Request
	0,  // 1: ocp.tip.api.MultiCreateTipV1Response.NotCreatedTips:type_name -> ocp.tip.api.CreateTipV1Request
	12, // 2: ocp.tip.api.DescribeTipV1Response.Tip:type_name -> ocp.tip.api.TipV1
	12, // 3: ocp.tip.api.ListTipsV1Response.Tips:type_name -> ocp.tip.api.TipV1
	0,  // 4: ocp.tip.api.OcpTipApi.CreateTipV1:input_type -> ocp.tip.api.CreateTipV1Request
	2,  // 5: ocp.tip.api.OcpTipApi.MultiCreateTipV1:input_type -> ocp.tip.api.MultiCreateTipV1Request
	4,  // 6: ocp.tip.api.OcpTipApi.UpdateTipV1:input_type -> ocp.tip.api.UpdateTipV1Request
	6,  // 7: ocp.tip.api.OcpTipApi.DescribeTipV1:input_type -> ocp.tip.api.DescribeTipV1Request
	8,  // 8: ocp.tip.api.OcpTipApi.ListTipsV1:input_type -> ocp.tip.api.ListTipsV1Request
	10, // 9: ocp.tip.api.OcpTipApi.RemoveTipV1:input_type -> ocp.tip.api.RemoveTipV1Request
	1,  // 10: ocp.tip.api.OcpTipApi.CreateTipV1:output_type -> ocp.tip.api.CreateTipV1Response
	3,  // 11: ocp.tip.api.OcpTipApi.MultiCreateTipV1:output_type -> ocp.tip.api.MultiCreateTipV1Response
	5,  // 12: ocp.tip.api.OcpTipApi.UpdateTipV1:output_type -> ocp.tip.api.UpdateTipV1Response
	7,  // 13: ocp.tip.api.OcpTipApi.DescribeTipV1:output_type -> ocp.tip.api.DescribeTipV1Response
	9,  // 14: ocp.tip.api.OcpTipApi.ListTipsV1:output_type -> ocp.tip.api.ListTipsV1Response
	11, // 15: ocp.tip.api.OcpTipApi.RemoveTipV1:output_type -> ocp.tip.api.RemoveTipV1Response
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_ocp_tip_api_ocp_tip_api_proto_init() }
func file_api_ocp_tip_api_ocp_tip_api_proto_init() {
	if File_api_ocp_tip_api_ocp_tip_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTipV1Request); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTipV1Response); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiCreateTipV1Request); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiCreateTipV1Response); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTipV1Request); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTipV1Response); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTipV1Request); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeTipV1Response); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTipsV1Request); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTipsV1Response); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTipV1Request); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveTipV1Response); i {
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
		file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TipV1); i {
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
			RawDescriptor: file_api_ocp_tip_api_ocp_tip_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ocp_tip_api_ocp_tip_api_proto_goTypes,
		DependencyIndexes: file_api_ocp_tip_api_ocp_tip_api_proto_depIdxs,
		MessageInfos:      file_api_ocp_tip_api_ocp_tip_api_proto_msgTypes,
	}.Build()
	File_api_ocp_tip_api_ocp_tip_api_proto = out.File
	file_api_ocp_tip_api_ocp_tip_api_proto_rawDesc = nil
	file_api_ocp_tip_api_ocp_tip_api_proto_goTypes = nil
	file_api_ocp_tip_api_ocp_tip_api_proto_depIdxs = nil
}
