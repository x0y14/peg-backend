// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: supervisor/v1/supervisor.proto

package supervisorv1

import (
	v1 "backend/gen/types/v1"
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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supervisor_v1_supervisor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supervisor_v1_supervisor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_supervisor_v1_supervisor_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supervisor_v1_supervisor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supervisor_v1_supervisor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_supervisor_v1_supervisor_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type RecordOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation *v1.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *RecordOperationRequest) Reset() {
	*x = RecordOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supervisor_v1_supervisor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordOperationRequest) ProtoMessage() {}

func (x *RecordOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supervisor_v1_supervisor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordOperationRequest.ProtoReflect.Descriptor instead.
func (*RecordOperationRequest) Descriptor() ([]byte, []int) {
	return file_supervisor_v1_supervisor_proto_rawDescGZIP(), []int{2}
}

func (x *RecordOperationRequest) GetOperation() *v1.Operation {
	if x != nil {
		return x.Operation
	}
	return nil
}

type RecordOperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RecordOperationResponse) Reset() {
	*x = RecordOperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supervisor_v1_supervisor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordOperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordOperationResponse) ProtoMessage() {}

func (x *RecordOperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supervisor_v1_supervisor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordOperationResponse.ProtoReflect.Descriptor instead.
func (*RecordOperationResponse) Descriptor() ([]byte, []int) {
	return file_supervisor_v1_supervisor_proto_rawDescGZIP(), []int{3}
}

var File_supervisor_v1_supervisor_proto protoreflect.FileDescriptor

var file_supervisor_v1_supervisor_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x22, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4b, 0x0a, 0x16,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xba, 0x01, 0x0a, 0x11, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69,
	0x73, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a,
	0x0f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76,
	0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x28, 0x5a, 0x26, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_supervisor_v1_supervisor_proto_rawDescOnce sync.Once
	file_supervisor_v1_supervisor_proto_rawDescData = file_supervisor_v1_supervisor_proto_rawDesc
)

func file_supervisor_v1_supervisor_proto_rawDescGZIP() []byte {
	file_supervisor_v1_supervisor_proto_rawDescOnce.Do(func() {
		file_supervisor_v1_supervisor_proto_rawDescData = protoimpl.X.CompressGZIP(file_supervisor_v1_supervisor_proto_rawDescData)
	})
	return file_supervisor_v1_supervisor_proto_rawDescData
}

var file_supervisor_v1_supervisor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_supervisor_v1_supervisor_proto_goTypes = []interface{}{
	(*PingRequest)(nil),             // 0: supervisor.v1.PingRequest
	(*PingResponse)(nil),            // 1: supervisor.v1.PingResponse
	(*RecordOperationRequest)(nil),  // 2: supervisor.v1.RecordOperationRequest
	(*RecordOperationResponse)(nil), // 3: supervisor.v1.RecordOperationResponse
	(*v1.Operation)(nil),            // 4: types.v1.Operation
}
var file_supervisor_v1_supervisor_proto_depIdxs = []int32{
	4, // 0: supervisor.v1.RecordOperationRequest.operation:type_name -> types.v1.Operation
	0, // 1: supervisor.v1.SupervisorService.Ping:input_type -> supervisor.v1.PingRequest
	2, // 2: supervisor.v1.SupervisorService.RecordOperation:input_type -> supervisor.v1.RecordOperationRequest
	1, // 3: supervisor.v1.SupervisorService.Ping:output_type -> supervisor.v1.PingResponse
	3, // 4: supervisor.v1.SupervisorService.RecordOperation:output_type -> supervisor.v1.RecordOperationResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_supervisor_v1_supervisor_proto_init() }
func file_supervisor_v1_supervisor_proto_init() {
	if File_supervisor_v1_supervisor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_supervisor_v1_supervisor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_supervisor_v1_supervisor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_supervisor_v1_supervisor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordOperationRequest); i {
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
		file_supervisor_v1_supervisor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordOperationResponse); i {
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
			RawDescriptor: file_supervisor_v1_supervisor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_supervisor_v1_supervisor_proto_goTypes,
		DependencyIndexes: file_supervisor_v1_supervisor_proto_depIdxs,
		MessageInfos:      file_supervisor_v1_supervisor_proto_msgTypes,
	}.Build()
	File_supervisor_v1_supervisor_proto = out.File
	file_supervisor_v1_supervisor_proto_rawDesc = nil
	file_supervisor_v1_supervisor_proto_goTypes = nil
	file_supervisor_v1_supervisor_proto_depIdxs = nil
}
