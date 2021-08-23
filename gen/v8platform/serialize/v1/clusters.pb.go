// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: v8platform/serialize/v1/clusters.proto

package serializev1

import (
	_ "github.com/v8platform/protos/gen/ras/encoding"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClusterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid                       string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ExpirationTimeout          int32  `protobuf:"varint,2,opt,name=expiration_timeout,json=expirationTimeout,proto3" json:"expiration_timeout,omitempty"`
	Host                       string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	LifetimeLimit              int32  `protobuf:"varint,4,opt,name=lifetime_limit,json=lifetimeLimit,proto3" json:"lifetime_limit,omitempty"`
	Port                       int32  `protobuf:"varint,5,opt,name=Port,proto3" json:"Port,omitempty"`
	MaxMemorySize              int32  `protobuf:"varint,6,opt,name=max_memory_size,json=maxMemorySize,proto3" json:"max_memory_size,omitempty"`
	MaxMemoryTimeLimit         int32  `protobuf:"varint,7,opt,name=max_memory_time_limit,json=maxMemoryTimeLimit,proto3" json:"max_memory_time_limit,omitempty"`
	Name                       string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	SecurityLevel              int32  `protobuf:"varint,9,opt,name=security_level,json=securityLevel,proto3" json:"security_level,omitempty"`
	SessionFaultToleranceLevel int32  `protobuf:"varint,10,opt,name=session_fault_tolerance_level,json=sessionFaultToleranceLevel,proto3" json:"session_fault_tolerance_level,omitempty"`
	LoadBalancingMode          int32  `protobuf:"varint,11,opt,name=load_balancing_mode,json=loadBalancingMode,proto3" json:"load_balancing_mode,omitempty"`
	ErrorsCountThreshold       int32  `protobuf:"varint,12,opt,name=errors_count_threshold,json=errorsCountThreshold,proto3" json:"errors_count_threshold,omitempty"`
	KillProblemProcesses       bool   `protobuf:"varint,13,opt,name=kill_problem_processes,json=killProblemProcesses,proto3" json:"kill_problem_processes,omitempty"`
	// version >= 9
	KillByMemoryWithDump bool `protobuf:"varint,14,opt,name=kill_by_memory_with_dump,json=killByMemoryWithDump,proto3" json:"kill_by_memory_with_dump,omitempty"`
}

func (x *ClusterInfo) Reset() {
	*x = ClusterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_serialize_v1_clusters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterInfo) ProtoMessage() {}

func (x *ClusterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_serialize_v1_clusters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterInfo.ProtoReflect.Descriptor instead.
func (*ClusterInfo) Descriptor() ([]byte, []int) {
	return file_v8platform_serialize_v1_clusters_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ClusterInfo) GetExpirationTimeout() int32 {
	if x != nil {
		return x.ExpirationTimeout
	}
	return 0
}

func (x *ClusterInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ClusterInfo) GetLifetimeLimit() int32 {
	if x != nil {
		return x.LifetimeLimit
	}
	return 0
}

func (x *ClusterInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ClusterInfo) GetMaxMemorySize() int32 {
	if x != nil {
		return x.MaxMemorySize
	}
	return 0
}

func (x *ClusterInfo) GetMaxMemoryTimeLimit() int32 {
	if x != nil {
		return x.MaxMemoryTimeLimit
	}
	return 0
}

func (x *ClusterInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterInfo) GetSecurityLevel() int32 {
	if x != nil {
		return x.SecurityLevel
	}
	return 0
}

func (x *ClusterInfo) GetSessionFaultToleranceLevel() int32 {
	if x != nil {
		return x.SessionFaultToleranceLevel
	}
	return 0
}

func (x *ClusterInfo) GetLoadBalancingMode() int32 {
	if x != nil {
		return x.LoadBalancingMode
	}
	return 0
}

func (x *ClusterInfo) GetErrorsCountThreshold() int32 {
	if x != nil {
		return x.ErrorsCountThreshold
	}
	return 0
}

func (x *ClusterInfo) GetKillProblemProcesses() bool {
	if x != nil {
		return x.KillProblemProcesses
	}
	return false
}

func (x *ClusterInfo) GetKillByMemoryWithDump() bool {
	if x != nil {
		return x.KillByMemoryWithDump
	}
	return false
}

var File_v8platform_serialize_v1_clusters_proto protoreflect.FileDescriptor

var file_v8platform_serialize_v1_clusters_proto_rawDesc = []byte{
	0x0a, 0x26, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x2f, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x08, 0x0a, 0x0b, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x72,
	0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x75, 0x75, 0x69, 0x64, 0x2c, 0x31,
	0x22, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10, 0x01, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x12, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x1b, 0x9a, 0x84, 0x9e, 0x03, 0x0e, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x3a, 0x22, 0x2c, 0x32, 0x22, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x02, 0x52, 0x11,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x2f, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1b, 0x9a, 0x84, 0x9e, 0x03, 0x0e, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a,
	0x22, 0x2c, 0x33, 0x22, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x03, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x1b, 0x9a, 0x84, 0x9e, 0x03,
	0x0e, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x2c, 0x34, 0x22, 0x82,
	0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x04, 0x52, 0x0d, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x27, 0x9a, 0x84, 0x9e, 0x03, 0x13, 0x72, 0x61, 0x73, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x2c, 0x35, 0x22, 0x82, 0xf5,
	0xea, 0x94, 0x0e, 0x09, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x10, 0x05, 0x52, 0x04, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x43, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x1b, 0x9a, 0x84,
	0x9e, 0x03, 0x0e, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x2c, 0x36,
	0x22, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x06, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x6d, 0x61, 0x78, 0x5f,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x1b, 0x9a, 0x84, 0x9e, 0x03, 0x0e, 0x72, 0x61,
	0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x2c, 0x37, 0x22, 0x82, 0xf5, 0xea, 0x94,
	0x0e, 0x02, 0x10, 0x07, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0x9a, 0x84, 0x9e, 0x03, 0x0e, 0x72, 0x61, 0x73,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x2c, 0x38, 0x22, 0x82, 0xf5, 0xea, 0x94, 0x0e,
	0x02, 0x10, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x1b, 0x9a, 0x84, 0x9e, 0x03, 0x0e, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x3a, 0x22, 0x2c, 0x39, 0x22, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x09, 0x52, 0x0d,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x5f, 0x0a,
	0x1d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x74,
	0x6f, 0x6c, 0x65, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x1c, 0x9a, 0x84, 0x9e, 0x03, 0x0f, 0x72, 0x61, 0x73, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x2c, 0x31, 0x30, 0x22, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02,
	0x10, 0x0a, 0x52, 0x1a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x75, 0x6c, 0x74,
	0x54, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x4c,
	0x0a, 0x13, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x42, 0x1c, 0x9a, 0x84, 0x9e,
	0x03, 0x0f, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x2c, 0x31, 0x31,
	0x22, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x0b, 0x52, 0x11, 0x6c, 0x6f, 0x61, 0x64, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x52, 0x0a, 0x16,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x42, 0x1c, 0x9a, 0x84,
	0x9e, 0x03, 0x0f, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x2c, 0x31,
	0x32, 0x22, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x0c, 0x52, 0x14, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x12, 0x52, 0x0a, 0x16, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x42, 0x1c, 0x9a, 0x84, 0x9e, 0x03, 0x0f, 0x72, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x3a, 0x22, 0x2c, 0x31, 0x33, 0x22, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x0d, 0x52, 0x14,
	0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x18, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x62, 0x79, 0x5f,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x75, 0x6d, 0x70,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x42, 0x20, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x72, 0x61, 0x73,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x3a, 0x22, 0x2c, 0x31, 0x34, 0x2c, 0x39, 0x22, 0x82, 0xf5,
	0xea, 0x94, 0x0e, 0x04, 0x10, 0x0e, 0x18, 0x09, 0x52, 0x14, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x79,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x57, 0x69, 0x74, 0x68, 0x44, 0x75, 0x6d, 0x70, 0x42, 0xf0,
	0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0d,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x53, 0x58, 0xaa, 0x02, 0x17, 0x56, 0x38,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x56, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x5c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x23, 0x56, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x56, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v8platform_serialize_v1_clusters_proto_rawDescOnce sync.Once
	file_v8platform_serialize_v1_clusters_proto_rawDescData = file_v8platform_serialize_v1_clusters_proto_rawDesc
)

func file_v8platform_serialize_v1_clusters_proto_rawDescGZIP() []byte {
	file_v8platform_serialize_v1_clusters_proto_rawDescOnce.Do(func() {
		file_v8platform_serialize_v1_clusters_proto_rawDescData = protoimpl.X.CompressGZIP(file_v8platform_serialize_v1_clusters_proto_rawDescData)
	})
	return file_v8platform_serialize_v1_clusters_proto_rawDescData
}

var file_v8platform_serialize_v1_clusters_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v8platform_serialize_v1_clusters_proto_goTypes = []interface{}{
	(*ClusterInfo)(nil), // 0: v8platform.serialize.v1.ClusterInfo
}
var file_v8platform_serialize_v1_clusters_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v8platform_serialize_v1_clusters_proto_init() }
func file_v8platform_serialize_v1_clusters_proto_init() {
	if File_v8platform_serialize_v1_clusters_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v8platform_serialize_v1_clusters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterInfo); i {
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
			RawDescriptor: file_v8platform_serialize_v1_clusters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v8platform_serialize_v1_clusters_proto_goTypes,
		DependencyIndexes: file_v8platform_serialize_v1_clusters_proto_depIdxs,
		MessageInfos:      file_v8platform_serialize_v1_clusters_proto_msgTypes,
	}.Build()
	File_v8platform_serialize_v1_clusters_proto = out.File
	file_v8platform_serialize_v1_clusters_proto_rawDesc = nil
	file_v8platform_serialize_v1_clusters_proto_goTypes = nil
	file_v8platform_serialize_v1_clusters_proto_depIdxs = nil
}
