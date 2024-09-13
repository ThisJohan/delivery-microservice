// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: logistics_service.proto

package api

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

type TodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TodoRequest) Reset() {
	*x = TodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistics_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoRequest) ProtoMessage() {}

func (x *TodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_logistics_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoRequest.ProtoReflect.Descriptor instead.
func (*TodoRequest) Descriptor() ([]byte, []int) {
	return file_logistics_service_proto_rawDescGZIP(), []int{0}
}

type TodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *TodoResponse) Reset() {
	*x = TodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logistics_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoResponse) ProtoMessage() {}

func (x *TodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_logistics_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoResponse.ProtoReflect.Descriptor instead.
func (*TodoResponse) Descriptor() ([]byte, []int) {
	return file_logistics_service_proto_rawDescGZIP(), []int{1}
}

func (x *TodoResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_logistics_service_proto protoreflect.FileDescriptor

var file_logistics_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x0d,
	0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a,
	0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x32, 0x41, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x10, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x69, 0x73, 0x4a, 0x6f, 0x68, 0x61, 0x6e, 0x2f, 0x73, 0x6e, 0x61,
	0x70, 0x70, 0x2d, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logistics_service_proto_rawDescOnce sync.Once
	file_logistics_service_proto_rawDescData = file_logistics_service_proto_rawDesc
)

func file_logistics_service_proto_rawDescGZIP() []byte {
	file_logistics_service_proto_rawDescOnce.Do(func() {
		file_logistics_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_logistics_service_proto_rawDescData)
	})
	return file_logistics_service_proto_rawDescData
}

var file_logistics_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_logistics_service_proto_goTypes = []any{
	(*TodoRequest)(nil),  // 0: api.TodoRequest
	(*TodoResponse)(nil), // 1: api.TodoResponse
}
var file_logistics_service_proto_depIdxs = []int32{
	0, // 0: api.LogisticsService.Todo:input_type -> api.TodoRequest
	1, // 1: api.LogisticsService.Todo:output_type -> api.TodoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logistics_service_proto_init() }
func file_logistics_service_proto_init() {
	if File_logistics_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logistics_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TodoRequest); i {
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
		file_logistics_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TodoResponse); i {
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
			RawDescriptor: file_logistics_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_logistics_service_proto_goTypes,
		DependencyIndexes: file_logistics_service_proto_depIdxs,
		MessageInfos:      file_logistics_service_proto_msgTypes,
	}.Build()
	File_logistics_service_proto = out.File
	file_logistics_service_proto_rawDesc = nil
	file_logistics_service_proto_goTypes = nil
	file_logistics_service_proto_depIdxs = nil
}
