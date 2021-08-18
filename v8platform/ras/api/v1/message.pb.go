// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: v8platform/ras/api/v1/message.proto

package v1

import (
	_ "github.com/v8platform/protos/v8platform/ras"
	_ "github.com/v8platform/protos/v8platform/ras/serialize"
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

type EndpointMessageType int32

const (
	EndpointMessageType_VOID_MESSAGE_KIND EndpointMessageType = 0
	EndpointMessageType_MESSAGE_KIND      EndpointMessageType = 1
	EndpointMessageType_EXCEPTION_KIND    EndpointMessageType = 99
)

// Enum value maps for EndpointMessageType.
var (
	EndpointMessageType_name = map[int32]string{
		0:  "VOID_MESSAGE_KIND",
		1:  "MESSAGE_KIND",
		99: "EXCEPTION_KIND",
	}
	EndpointMessageType_value = map[string]int32{
		"VOID_MESSAGE_KIND": 0,
		"MESSAGE_KIND":      1,
		"EXCEPTION_KIND":    99,
	}
)

func (x EndpointMessageType) Enum() *EndpointMessageType {
	p := new(EndpointMessageType)
	*p = x
	return p
}

func (x EndpointMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EndpointMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_v8platform_ras_api_v1_message_proto_enumTypes[0].Descriptor()
}

func (EndpointMessageType) Type() protoreflect.EnumType {
	return &file_v8platform_ras_api_v1_message_proto_enumTypes[0]
}

func (x EndpointMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EndpointMessageType.Descriptor instead.
func (EndpointMessageType) EnumDescriptor() ([]byte, []int) {
	return file_v8platform_ras_api_v1_message_proto_rawDescGZIP(), []int{0}
}

type EndpointMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId int32               `protobuf:"varint,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Format     int32               `protobuf:"varint,2,opt,name=format,proto3" json:"format,omitempty"` // TODO int16
	Type       EndpointMessageType `protobuf:"varint,3,opt,name=type,proto3,enum=ras.api.v1.EndpointMessageType" json:"type,omitempty"`
	Message    *Message            `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EndpointMessage) Reset() {
	*x = EndpointMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_ras_api_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointMessage) ProtoMessage() {}

func (x *EndpointMessage) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_ras_api_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointMessage.ProtoReflect.Descriptor instead.
func (*EndpointMessage) Descriptor() ([]byte, []int) {
	return file_v8platform_ras_api_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *EndpointMessage) GetEndpointId() int32 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

func (x *EndpointMessage) GetFormat() int32 {
	if x != nil {
		return x.Format
	}
	return 0
}

func (x *EndpointMessage) GetType() EndpointMessageType {
	if x != nil {
		return x.Type
	}
	return EndpointMessageType_VOID_MESSAGE_KIND
}

func (x *EndpointMessage) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_ras_api_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_ras_api_v1_message_proto_msgTypes[1]
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
	return file_v8platform_ras_api_v1_message_proto_rawDescGZIP(), []int{1}
}

var File_v8platform_ras_api_v1_message_proto protoreflect.FileDescriptor

var file_v8platform_ras_api_v1_message_proto_rawDesc = []byte{
	0x0a, 0x23, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x61, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x1a, 0x27, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x61,
	0x73, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x76, 0x38, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0f, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2a, 0x52, 0x0a, 0x13, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x56,
	0x4f, 0x49, 0x44, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x4b, 0x49,
	0x4e, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x10, 0x63, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v8platform_ras_api_v1_message_proto_rawDescOnce sync.Once
	file_v8platform_ras_api_v1_message_proto_rawDescData = file_v8platform_ras_api_v1_message_proto_rawDesc
)

func file_v8platform_ras_api_v1_message_proto_rawDescGZIP() []byte {
	file_v8platform_ras_api_v1_message_proto_rawDescOnce.Do(func() {
		file_v8platform_ras_api_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_v8platform_ras_api_v1_message_proto_rawDescData)
	})
	return file_v8platform_ras_api_v1_message_proto_rawDescData
}

var file_v8platform_ras_api_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v8platform_ras_api_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v8platform_ras_api_v1_message_proto_goTypes = []interface{}{
	(EndpointMessageType)(0), // 0: ras.api.v1.EndpointMessageType
	(*EndpointMessage)(nil),  // 1: ras.api.v1.EndpointMessage
	(*Message)(nil),          // 2: ras.api.v1.Message
}
var file_v8platform_ras_api_v1_message_proto_depIdxs = []int32{
	0, // 0: ras.api.v1.EndpointMessage.type:type_name -> ras.api.v1.EndpointMessageType
	2, // 1: ras.api.v1.EndpointMessage.message:type_name -> ras.api.v1.Message
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v8platform_ras_api_v1_message_proto_init() }
func file_v8platform_ras_api_v1_message_proto_init() {
	if File_v8platform_ras_api_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v8platform_ras_api_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointMessage); i {
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
		file_v8platform_ras_api_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v8platform_ras_api_v1_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v8platform_ras_api_v1_message_proto_goTypes,
		DependencyIndexes: file_v8platform_ras_api_v1_message_proto_depIdxs,
		EnumInfos:         file_v8platform_ras_api_v1_message_proto_enumTypes,
		MessageInfos:      file_v8platform_ras_api_v1_message_proto_msgTypes,
	}.Build()
	File_v8platform_ras_api_v1_message_proto = out.File
	file_v8platform_ras_api_v1_message_proto_rawDesc = nil
	file_v8platform_ras_api_v1_message_proto_goTypes = nil
	file_v8platform_ras_api_v1_message_proto_depIdxs = nil
}
