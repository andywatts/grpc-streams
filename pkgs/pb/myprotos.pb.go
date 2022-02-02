// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: myprotos.proto

package pb

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

type MyUnaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MyUnaryRequest) Reset() {
	*x = MyUnaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myprotos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyUnaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyUnaryRequest) ProtoMessage() {}

func (x *MyUnaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_myprotos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyUnaryRequest.ProtoReflect.Descriptor instead.
func (*MyUnaryRequest) Descriptor() ([]byte, []int) {
	return file_myprotos_proto_rawDescGZIP(), []int{0}
}

func (x *MyUnaryRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type MyUnaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MyUnaryResponse) Reset() {
	*x = MyUnaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myprotos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyUnaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyUnaryResponse) ProtoMessage() {}

func (x *MyUnaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_myprotos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyUnaryResponse.ProtoReflect.Descriptor instead.
func (*MyUnaryResponse) Descriptor() ([]byte, []int) {
	return file_myprotos_proto_rawDescGZIP(), []int{1}
}

func (x *MyUnaryResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type MyStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MyStreamRequest) Reset() {
	*x = MyStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myprotos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyStreamRequest) ProtoMessage() {}

func (x *MyStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_myprotos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyStreamRequest.ProtoReflect.Descriptor instead.
func (*MyStreamRequest) Descriptor() ([]byte, []int) {
	return file_myprotos_proto_rawDescGZIP(), []int{2}
}

func (x *MyStreamRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type MyStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MyStreamResponse) Reset() {
	*x = MyStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myprotos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyStreamResponse) ProtoMessage() {}

func (x *MyStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_myprotos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyStreamResponse.ProtoReflect.Descriptor instead.
func (*MyStreamResponse) Descriptor() ([]byte, []int) {
	return file_myprotos_proto_rawDescGZIP(), []int{3}
}

func (x *MyStreamResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_myprotos_proto protoreflect.FileDescriptor

var file_myprotos_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x26, 0x0a, 0x0e, 0x4d, 0x79, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x27, 0x0a, 0x0f,
	0x4d, 0x79, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x4d, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x28,
	0x0a, 0x10, 0x4d, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x80, 0x01, 0x0a, 0x09, 0x4d, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x61,
	0x72, 0x79, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x55, 0x6e,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x4d, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x6d,
	0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_myprotos_proto_rawDescOnce sync.Once
	file_myprotos_proto_rawDescData = file_myprotos_proto_rawDesc
)

func file_myprotos_proto_rawDescGZIP() []byte {
	file_myprotos_proto_rawDescOnce.Do(func() {
		file_myprotos_proto_rawDescData = protoimpl.X.CompressGZIP(file_myprotos_proto_rawDescData)
	})
	return file_myprotos_proto_rawDescData
}

var file_myprotos_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_myprotos_proto_goTypes = []interface{}{
	(*MyUnaryRequest)(nil),   // 0: pb.MyUnaryRequest
	(*MyUnaryResponse)(nil),  // 1: pb.MyUnaryResponse
	(*MyStreamRequest)(nil),  // 2: pb.MyStreamRequest
	(*MyStreamResponse)(nil), // 3: pb.MyStreamResponse
}
var file_myprotos_proto_depIdxs = []int32{
	0, // 0: pb.MyService.GetUnary:input_type -> pb.MyUnaryRequest
	2, // 1: pb.MyService.GetStream:input_type -> pb.MyStreamRequest
	1, // 2: pb.MyService.GetUnary:output_type -> pb.MyUnaryResponse
	3, // 3: pb.MyService.GetStream:output_type -> pb.MyStreamResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_myprotos_proto_init() }
func file_myprotos_proto_init() {
	if File_myprotos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_myprotos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyUnaryRequest); i {
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
		file_myprotos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyUnaryResponse); i {
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
		file_myprotos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyStreamRequest); i {
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
		file_myprotos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyStreamResponse); i {
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
			RawDescriptor: file_myprotos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_myprotos_proto_goTypes,
		DependencyIndexes: file_myprotos_proto_depIdxs,
		MessageInfos:      file_myprotos_proto_msgTypes,
	}.Build()
	File_myprotos_proto = out.File
	file_myprotos_proto_rawDesc = nil
	file_myprotos_proto_goTypes = nil
	file_myprotos_proto_depIdxs = nil
}