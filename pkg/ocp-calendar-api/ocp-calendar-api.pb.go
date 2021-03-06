// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/ocp-calendar-api/ocp-calendar-api.proto

package ocp_calendar_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateCalendarRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Type   uint64 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	Link   string `protobuf:"bytes,3,opt,name=Link,proto3" json:"Link,omitempty"`
}

func (x *CreateCalendarRequestV1) Reset() {
	*x = CreateCalendarRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCalendarRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCalendarRequestV1) ProtoMessage() {}

func (x *CreateCalendarRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCalendarRequestV1.ProtoReflect.Descriptor instead.
func (*CreateCalendarRequestV1) Descriptor() ([]byte, []int) {
	return file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCalendarRequestV1) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateCalendarRequestV1) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CreateCalendarRequestV1) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type MultiCreateCalendarRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64                     `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Items []*CreateCalendarRequestV1 `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *MultiCreateCalendarRequestV1) Reset() {
	*x = MultiCreateCalendarRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiCreateCalendarRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiCreateCalendarRequestV1) ProtoMessage() {}

func (x *MultiCreateCalendarRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiCreateCalendarRequestV1.ProtoReflect.Descriptor instead.
func (*MultiCreateCalendarRequestV1) Descriptor() ([]byte, []int) {
	return file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescGZIP(), []int{1}
}

func (x *MultiCreateCalendarRequestV1) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MultiCreateCalendarRequestV1) GetItems() []*CreateCalendarRequestV1 {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdateCalendarRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Type   uint64 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Link   string `protobuf:"bytes,4,opt,name=Link,proto3" json:"Link,omitempty"`
}

func (x *UpdateCalendarRequestV1) Reset() {
	*x = UpdateCalendarRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCalendarRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCalendarRequestV1) ProtoMessage() {}

func (x *UpdateCalendarRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCalendarRequestV1.ProtoReflect.Descriptor instead.
func (*UpdateCalendarRequestV1) Descriptor() ([]byte, []int) {
	return file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCalendarRequestV1) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCalendarRequestV1) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateCalendarRequestV1) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *UpdateCalendarRequestV1) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type DescribeCalendarRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DescribeCalendarRequestV1) Reset() {
	*x = DescribeCalendarRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeCalendarRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCalendarRequestV1) ProtoMessage() {}

func (x *DescribeCalendarRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCalendarRequestV1.ProtoReflect.Descriptor instead.
func (*DescribeCalendarRequestV1) Descriptor() ([]byte, []int) {
	return file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeCalendarRequestV1) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DescribeCalendarResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Type   uint64 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Link   string `protobuf:"bytes,4,opt,name=Link,proto3" json:"Link,omitempty"`
}

func (x *DescribeCalendarResponseV1) Reset() {
	*x = DescribeCalendarResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeCalendarResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCalendarResponseV1) ProtoMessage() {}

func (x *DescribeCalendarResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCalendarResponseV1.ProtoReflect.Descriptor instead.
func (*DescribeCalendarResponseV1) Descriptor() ([]byte, []int) {
	return file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeCalendarResponseV1) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DescribeCalendarResponseV1) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DescribeCalendarResponseV1) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *DescribeCalendarResponseV1) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type ListCalendarRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	UserId uint64 `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Type   uint64 `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *ListCalendarRequestV1) Reset() {
	*x = ListCalendarRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCalendarRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCalendarRequestV1) ProtoMessage() {}

func (x *ListCalendarRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCalendarRequestV1.ProtoReflect.Descriptor instead.
func (*ListCalendarRequestV1) Descriptor() ([]byte, []int) {
	return file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListCalendarRequestV1) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListCalendarRequestV1) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListCalendarRequestV1) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListCalendarRequestV1) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

type ListCalendarResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*DescribeCalendarResponseV1 `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListCalendarResponseV1) Reset() {
	*x = ListCalendarResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCalendarResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCalendarResponseV1) ProtoMessage() {}

func (x *ListCalendarResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCalendarResponseV1.ProtoReflect.Descriptor instead.
func (*ListCalendarResponseV1) Descriptor() ([]byte, []int) {
	return file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListCalendarResponseV1) GetItems() []*DescribeCalendarResponseV1 {
	if x != nil {
		return x.Items
	}
	return nil
}

type RemoveCalendarRequestV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *RemoveCalendarRequestV1) Reset() {
	*x = RemoveCalendarRequestV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCalendarRequestV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCalendarRequestV1) ProtoMessage() {}

func (x *RemoveCalendarRequestV1) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCalendarRequestV1.ProtoReflect.Descriptor instead.
func (*RemoveCalendarRequestV1) Descriptor() ([]byte, []int) {
	return file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveCalendarRequestV1) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_ocp_calendar_api_ocp_calendar_api_proto protoreflect.FileDescriptor

var file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6f,
	0x63, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a,
	0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x1f, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20,
	0x00, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x6f, 0x0a, 0x1c, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x05, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6f, 0x63, 0x70, 0x5f,
	0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x56, 0x31, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x17,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x17, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69,
	0x6e, 0x6b, 0x22, 0x2b, 0x0a, 0x19, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x22,
	0x6c, 0x0a, 0x1a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x71, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x5c, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12, 0x42, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x63, 0x70, 0x5f,
	0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x29,
	0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x32, 0x83, 0x06, 0x0a, 0x0e, 0x4f, 0x63,
	0x70, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x41, 0x70, 0x69, 0x12, 0x6e, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x56, 0x31,
	0x12, 0x29, 0x2e, 0x6f, 0x63, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x79, 0x0a, 0x15,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x56, 0x31, 0x12, 0x2e, 0x2e, 0x6f, 0x63, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x65,
	0x6e, 0x64, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x89, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x56, 0x31, 0x12, 0x29, 0x2e, 0x6f,
	0x63, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x2c, 0x2e, 0x6f, 0x63, 0x70, 0x5f, 0x63, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x1a, 0x11, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x7b, 0x49, 0x64, 0x7d,
	0x3a, 0x01, 0x2a, 0x12, 0x8a, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x56, 0x31, 0x12, 0x2b, 0x2e, 0x6f, 0x63, 0x70,
	0x5f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x2c, 0x2e, 0x6f, 0x63, 0x70, 0x5f, 0x63, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x7b, 0x49, 0x64, 0x7d,
	0x12, 0x7b, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x73, 0x56, 0x31, 0x12, 0x27, 0x2e, 0x6f, 0x63, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x28, 0x2e, 0x6f,
	0x63, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x73, 0x12, 0x70, 0x0a,
	0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x56,
	0x31, 0x12, 0x29, 0x2e, 0x6f, 0x63, 0x70, 0x5f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x31, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x7b, 0x49, 0x64, 0x7d, 0x42,
	0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a,
	0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x63, 0x70, 0x5f, 0x63, 0x61, 0x6c,
	0x65, 0x6e, 0x64, 0x61, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescOnce sync.Once
	file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescData = file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDesc
)

func file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescGZIP() []byte {
	file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescOnce.Do(func() {
		file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescData)
	})
	return file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDescData
}

var file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_ocp_calendar_api_ocp_calendar_api_proto_goTypes = []interface{}{
	(*CreateCalendarRequestV1)(nil),      // 0: ocp_calendar_api.CreateCalendarRequestV1
	(*MultiCreateCalendarRequestV1)(nil), // 1: ocp_calendar_api.MultiCreateCalendarRequestV1
	(*UpdateCalendarRequestV1)(nil),      // 2: ocp_calendar_api.UpdateCalendarRequestV1
	(*DescribeCalendarRequestV1)(nil),    // 3: ocp_calendar_api.DescribeCalendarRequestV1
	(*DescribeCalendarResponseV1)(nil),   // 4: ocp_calendar_api.DescribeCalendarResponseV1
	(*ListCalendarRequestV1)(nil),        // 5: ocp_calendar_api.ListCalendarRequestV1
	(*ListCalendarResponseV1)(nil),       // 6: ocp_calendar_api.ListCalendarResponseV1
	(*RemoveCalendarRequestV1)(nil),      // 7: ocp_calendar_api.RemoveCalendarRequestV1
	(*emptypb.Empty)(nil),                // 8: google.protobuf.Empty
}
var file_api_ocp_calendar_api_ocp_calendar_api_proto_depIdxs = []int32{
	0, // 0: ocp_calendar_api.MultiCreateCalendarRequestV1.Items:type_name -> ocp_calendar_api.CreateCalendarRequestV1
	4, // 1: ocp_calendar_api.ListCalendarResponseV1.items:type_name -> ocp_calendar_api.DescribeCalendarResponseV1
	0, // 2: ocp_calendar_api.OcpCalendarApi.CreateCalendarV1:input_type -> ocp_calendar_api.CreateCalendarRequestV1
	1, // 3: ocp_calendar_api.OcpCalendarApi.MultiCreateCalendarV1:input_type -> ocp_calendar_api.MultiCreateCalendarRequestV1
	2, // 4: ocp_calendar_api.OcpCalendarApi.UpdateCalendarV1:input_type -> ocp_calendar_api.UpdateCalendarRequestV1
	3, // 5: ocp_calendar_api.OcpCalendarApi.DescribeCalendarV1:input_type -> ocp_calendar_api.DescribeCalendarRequestV1
	5, // 6: ocp_calendar_api.OcpCalendarApi.ListCalendarsV1:input_type -> ocp_calendar_api.ListCalendarRequestV1
	7, // 7: ocp_calendar_api.OcpCalendarApi.RemoveCalendarV1:input_type -> ocp_calendar_api.RemoveCalendarRequestV1
	8, // 8: ocp_calendar_api.OcpCalendarApi.CreateCalendarV1:output_type -> google.protobuf.Empty
	8, // 9: ocp_calendar_api.OcpCalendarApi.MultiCreateCalendarV1:output_type -> google.protobuf.Empty
	4, // 10: ocp_calendar_api.OcpCalendarApi.UpdateCalendarV1:output_type -> ocp_calendar_api.DescribeCalendarResponseV1
	4, // 11: ocp_calendar_api.OcpCalendarApi.DescribeCalendarV1:output_type -> ocp_calendar_api.DescribeCalendarResponseV1
	6, // 12: ocp_calendar_api.OcpCalendarApi.ListCalendarsV1:output_type -> ocp_calendar_api.ListCalendarResponseV1
	8, // 13: ocp_calendar_api.OcpCalendarApi.RemoveCalendarV1:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_ocp_calendar_api_ocp_calendar_api_proto_init() }
func file_api_ocp_calendar_api_ocp_calendar_api_proto_init() {
	if File_api_ocp_calendar_api_ocp_calendar_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCalendarRequestV1); i {
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
		file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiCreateCalendarRequestV1); i {
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
		file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCalendarRequestV1); i {
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
		file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeCalendarRequestV1); i {
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
		file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeCalendarResponseV1); i {
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
		file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCalendarRequestV1); i {
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
		file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCalendarResponseV1); i {
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
		file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCalendarRequestV1); i {
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
			RawDescriptor: file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ocp_calendar_api_ocp_calendar_api_proto_goTypes,
		DependencyIndexes: file_api_ocp_calendar_api_ocp_calendar_api_proto_depIdxs,
		MessageInfos:      file_api_ocp_calendar_api_ocp_calendar_api_proto_msgTypes,
	}.Build()
	File_api_ocp_calendar_api_ocp_calendar_api_proto = out.File
	file_api_ocp_calendar_api_ocp_calendar_api_proto_rawDesc = nil
	file_api_ocp_calendar_api_ocp_calendar_api_proto_goTypes = nil
	file_api_ocp_calendar_api_ocp_calendar_api_proto_depIdxs = nil
}
