// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: types/v1/types.proto

package typesv1

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

// === talk service ===
type MessageContentType int32

const (
	MessageContentType_MESSAGE_CONTENT_TYPE_ILLEGAL_UNSPECIFIED MessageContentType = 0
	MessageContentType_MESSAGE_CONTENT_TYPE_TEXT                MessageContentType = 1
	MessageContentType_MESSAGE_CONTENT_TYPE_IMAGE               MessageContentType = 2
	MessageContentType_MESSAGE_CONTENT_TYPE_VIDEO               MessageContentType = 3
	MessageContentType_MESSAGE_CONTENT_TYPE_GIF                 MessageContentType = 4
)

// Enum value maps for MessageContentType.
var (
	MessageContentType_name = map[int32]string{
		0: "MESSAGE_CONTENT_TYPE_ILLEGAL_UNSPECIFIED",
		1: "MESSAGE_CONTENT_TYPE_TEXT",
		2: "MESSAGE_CONTENT_TYPE_IMAGE",
		3: "MESSAGE_CONTENT_TYPE_VIDEO",
		4: "MESSAGE_CONTENT_TYPE_GIF",
	}
	MessageContentType_value = map[string]int32{
		"MESSAGE_CONTENT_TYPE_ILLEGAL_UNSPECIFIED": 0,
		"MESSAGE_CONTENT_TYPE_TEXT":                1,
		"MESSAGE_CONTENT_TYPE_IMAGE":               2,
		"MESSAGE_CONTENT_TYPE_VIDEO":               3,
		"MESSAGE_CONTENT_TYPE_GIF":                 4,
	}
)

func (x MessageContentType) Enum() *MessageContentType {
	p := new(MessageContentType)
	*p = x
	return p
}

func (x MessageContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_types_v1_types_proto_enumTypes[0].Descriptor()
}

func (MessageContentType) Type() protoreflect.EnumType {
	return &file_types_v1_types_proto_enumTypes[0]
}

func (x MessageContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageContentType.Descriptor instead.
func (MessageContentType) EnumDescriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{0}
}

// === operation service ===
type OperationType int32

const (
	OperationType_OPERATION_TYPE_NOOP_UNSPECIFIED OperationType = 0
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "OPERATION_TYPE_NOOP_UNSPECIFIED",
	}
	OperationType_value = map[string]int32{
		"OPERATION_TYPE_NOOP_UNSPECIFIED": 0,
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
	return file_types_v1_types_proto_enumTypes[1].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_types_v1_types_proto_enumTypes[1]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{1}
}

// === user service ===
type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Account) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Account) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DisplayName   string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	IconPath      string `protobuf:"bytes,3,opt,name=icon_path,json=iconPath,proto3" json:"icon_path,omitempty"`
	StatusMessage string `protobuf:"bytes,4,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	Metadata      string `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *Profile) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Profile) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Profile) GetIconPath() string {
	if x != nil {
		return x.IconPath
	}
	return ""
}

func (x *Profile) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *Profile) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowSearch bool `protobuf:"varint,1,opt,name=allow_search,json=allowSearch,proto3" json:"allow_search,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{2}
}

func (x *Settings) GetAllowSearch() bool {
	if x != nil {
		return x.AllowSearch
	}
	return false
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId   string             `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	From        string             `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To          string             `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	ContentType MessageContentType `protobuf:"varint,4,opt,name=content_type,json=contentType,proto3,enum=types.v1.MessageContentType" json:"content_type,omitempty"`
	Text        string             `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	Metadata    string             `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Message) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Message) GetContentType() MessageContentType {
	if x != nil {
		return x.ContentType
	}
	return MessageContentType_MESSAGE_CONTENT_TYPE_ILLEGAL_UNSPECIFIED
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationId int64         `protobuf:"varint,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	Type        OperationType `protobuf:"varint,2,opt,name=type,proto3,enum=types.v1.OperationType" json:"type,omitempty"`
	Source      string        `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Destination []string      `protobuf:"bytes,4,rep,name=destination,proto3" json:"destination,omitempty"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{4}
}

func (x *Operation) GetOperationId() int64 {
	if x != nil {
		return x.OperationId
	}
	return 0
}

func (x *Operation) GetType() OperationType {
	if x != nil {
		return x.Type
	}
	return OperationType_OPERATION_TYPE_NOOP_UNSPECIFIED
}

func (x *Operation) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Operation) GetDestination() []string {
	if x != nil {
		return x.Destination
	}
	return nil
}

var File_types_v1_types_proto protoreflect.FileDescriptor

var file_types_v1_types_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x22, 0x55, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa5, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x63, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x2d, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0xbd,
	0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x3f, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x95,
	0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0xbf, 0x01, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a,
	0x28, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x4d,
	0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x47, 0x49, 0x46, 0x10, 0x04, 0x2a, 0x34, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4f, 0x50,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x42, 0x1e,
	0x5a, 0x1c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_v1_types_proto_rawDescOnce sync.Once
	file_types_v1_types_proto_rawDescData = file_types_v1_types_proto_rawDesc
)

func file_types_v1_types_proto_rawDescGZIP() []byte {
	file_types_v1_types_proto_rawDescOnce.Do(func() {
		file_types_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_v1_types_proto_rawDescData)
	})
	return file_types_v1_types_proto_rawDescData
}

var file_types_v1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_types_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_types_v1_types_proto_goTypes = []interface{}{
	(MessageContentType)(0), // 0: types.v1.MessageContentType
	(OperationType)(0),      // 1: types.v1.OperationType
	(*Account)(nil),         // 2: types.v1.Account
	(*Profile)(nil),         // 3: types.v1.Profile
	(*Settings)(nil),        // 4: types.v1.Settings
	(*Message)(nil),         // 5: types.v1.Message
	(*Operation)(nil),       // 6: types.v1.Operation
}
var file_types_v1_types_proto_depIdxs = []int32{
	0, // 0: types.v1.Message.content_type:type_name -> types.v1.MessageContentType
	1, // 1: types.v1.Operation.type:type_name -> types.v1.OperationType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_types_v1_types_proto_init() }
func file_types_v1_types_proto_init() {
	if File_types_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_types_v1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_types_v1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Settings); i {
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
		file_types_v1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_types_v1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
			RawDescriptor: file_types_v1_types_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_types_proto_goTypes,
		DependencyIndexes: file_types_v1_types_proto_depIdxs,
		EnumInfos:         file_types_v1_types_proto_enumTypes,
		MessageInfos:      file_types_v1_types_proto_msgTypes,
	}.Build()
	File_types_v1_types_proto = out.File
	file_types_v1_types_proto_rawDesc = nil
	file_types_v1_types_proto_goTypes = nil
	file_types_v1_types_proto_depIdxs = nil
}
