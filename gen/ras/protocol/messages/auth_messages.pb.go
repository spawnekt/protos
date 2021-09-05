// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ras/protocol/messages/auth_messages.proto

package messages

import (
	_ "github.com/v8platform/protoc-gen-go-ras/plugin/ras/encoding"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthenticateAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthenticateAgentRequest) Reset() {
	*x = AuthenticateAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_messages_auth_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateAgentRequest) ProtoMessage() {}

func (x *AuthenticateAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_messages_auth_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateAgentRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateAgentRequest) Descriptor() ([]byte, []int) {
	return file_ras_protocol_messages_auth_messages_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticateAgentRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AuthenticateAgentRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ClusterAuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	User      string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ClusterAuthenticateRequest) Reset() {
	*x = ClusterAuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_messages_auth_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterAuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterAuthenticateRequest) ProtoMessage() {}

func (x *ClusterAuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_messages_auth_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterAuthenticateRequest.ProtoReflect.Descriptor instead.
func (*ClusterAuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_ras_protocol_messages_auth_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterAuthenticateRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ClusterAuthenticateRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ClusterAuthenticateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthenticateInfobaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	User      string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Password  string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthenticateInfobaseRequest) Reset() {
	*x = AuthenticateInfobaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ras_protocol_messages_auth_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateInfobaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateInfobaseRequest) ProtoMessage() {}

func (x *AuthenticateInfobaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ras_protocol_messages_auth_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateInfobaseRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateInfobaseRequest) Descriptor() ([]byte, []int) {
	return file_ras_protocol_messages_auth_messages_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticateInfobaseRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *AuthenticateInfobaseRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AuthenticateInfobaseRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_ras_protocol_messages_auth_messages_proto protoreflect.FileDescriptor

var file_ras_protocol_messages_auth_messages_proto_rawDesc = []byte{
	0x0a, 0x29, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x61, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x1a, 0x16, 0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x2f, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x18, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x01, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x02, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x22, 0x8a, 0xf5, 0xea, 0x94, 0x0e, 0x1c, 0x3a, 0x1a,
	0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x47, 0x45,
	0x4e, 0x54, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x22, 0xad, 0x01, 0x0a, 0x1a, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x82,
	0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10, 0x01, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x02,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02,
	0x10, 0x03, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x1c, 0x8a, 0xf5,
	0xea, 0x94, 0x0e, 0x16, 0x3a, 0x14, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41,
	0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x22, 0xb4, 0x01, 0x0a, 0x1b, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e,
	0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10, 0x01, 0x52, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10,
	0x02, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e,
	0x02, 0x10, 0x03, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x22, 0x8a,
	0xf5, 0xea, 0x94, 0x0e, 0x1c, 0x3a, 0x1a, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45,
	0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x42, 0xdc, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42,
	0x11, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x61, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x52,
	0x50, 0x4d, 0xaa, 0x02, 0x15, 0x52, 0x61, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0xca, 0x02, 0x15, 0x52, 0x61, 0x73,
	0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0xe2, 0x02, 0x21, 0x52, 0x61, 0x73, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x17, 0x52, 0x61, 0x73, 0x3a, 0x3a, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ras_protocol_messages_auth_messages_proto_rawDescOnce sync.Once
	file_ras_protocol_messages_auth_messages_proto_rawDescData = file_ras_protocol_messages_auth_messages_proto_rawDesc
)

func file_ras_protocol_messages_auth_messages_proto_rawDescGZIP() []byte {
	file_ras_protocol_messages_auth_messages_proto_rawDescOnce.Do(func() {
		file_ras_protocol_messages_auth_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_ras_protocol_messages_auth_messages_proto_rawDescData)
	})
	return file_ras_protocol_messages_auth_messages_proto_rawDescData
}

var file_ras_protocol_messages_auth_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ras_protocol_messages_auth_messages_proto_goTypes = []interface{}{
	(*AuthenticateAgentRequest)(nil),    // 0: ras.protocol.messages.AuthenticateAgentRequest
	(*ClusterAuthenticateRequest)(nil),  // 1: ras.protocol.messages.ClusterAuthenticateRequest
	(*AuthenticateInfobaseRequest)(nil), // 2: ras.protocol.messages.AuthenticateInfobaseRequest
}
var file_ras_protocol_messages_auth_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ras_protocol_messages_auth_messages_proto_init() }
func file_ras_protocol_messages_auth_messages_proto_init() {
	if File_ras_protocol_messages_auth_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ras_protocol_messages_auth_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateAgentRequest); i {
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
		file_ras_protocol_messages_auth_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterAuthenticateRequest); i {
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
		file_ras_protocol_messages_auth_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateInfobaseRequest); i {
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
			RawDescriptor: file_ras_protocol_messages_auth_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ras_protocol_messages_auth_messages_proto_goTypes,
		DependencyIndexes: file_ras_protocol_messages_auth_messages_proto_depIdxs,
		MessageInfos:      file_ras_protocol_messages_auth_messages_proto_msgTypes,
	}.Build()
	File_ras_protocol_messages_auth_messages_proto = out.File
	file_ras_protocol_messages_auth_messages_proto_rawDesc = nil
	file_ras_protocol_messages_auth_messages_proto_goTypes = nil
	file_ras_protocol_messages_auth_messages_proto_depIdxs = nil
}
