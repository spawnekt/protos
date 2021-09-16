// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: v8platform/serialize/v1/servers.proto

package serializev1

import (
	_ "github.com/v8platform/protoc-gen-go-ras/plugin/ras/encoding"
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

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid                                 string       `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	AgentHost                            string       `protobuf:"bytes,2,opt,name=agent_host,json=agentHost,proto3" json:"agent_host,omitempty"`
	AgentPort                            int32        `protobuf:"varint,3,opt,name=agent_port,json=agentPort,proto3" json:"agent_port,omitempty"`
	Name                                 string       `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	MainServer                           bool         `protobuf:"varint,5,opt,name=main_server,json=mainServer,proto3" json:"main_server,omitempty"`
	SafeWorkingProcessesMemoryLimit      int64        `protobuf:"varint,6,opt,name=safe_working_processes_memory_limit,json=safeWorkingProcessesMemoryLimit,proto3" json:"safe_working_processes_memory_limit,omitempty"`
	SafeCallMemoryLimit                  int64        `protobuf:"varint,7,opt,name=safe_call_memory_limit,json=safeCallMemoryLimit,proto3" json:"safe_call_memory_limit,omitempty"`
	InfobasesLimit                       int32        `protobuf:"varint,8,opt,name=infobases_limit,json=infobasesLimit,proto3" json:"infobases_limit,omitempty"`
	MemoryLimit                          int64        `protobuf:"varint,9,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	ConnectionsLimit                     int32        `protobuf:"varint,10,opt,name=connections_limit,json=connectionsLimit,proto3" json:"connections_limit,omitempty"`
	ClusterPort                          int32        `protobuf:"varint,11,opt,name=cluster_port,json=clusterPort,proto3" json:"cluster_port,omitempty"`
	DedicateManagers                     bool         `protobuf:"varint,12,opt,name=dedicate_managers,json=dedicateManagers,proto3" json:"dedicate_managers,omitempty"`
	PortRanges                           []*PortRange `protobuf:"bytes,13,rep,name=port_ranges,json=portRanges,proto3" json:"port_ranges,omitempty"`
	CriticalTotalMemory                  int64        `protobuf:"varint,14,opt,name=critical_total_memory,json=criticalTotalMemory,proto3" json:"critical_total_memory,omitempty"`
	TemporaryAllowedTotalMemory          int64        `protobuf:"varint,15,opt,name=temporary_allowed_total_memory,json=temporaryAllowedTotalMemory,proto3" json:"temporary_allowed_total_memory,omitempty"`
	TemporaryAllowedTotalMemoryTimeLimit int64        `protobuf:"varint,16,opt,name=temporary_allowed_total_memory_time_limit,json=temporaryAllowedTotalMemoryTimeLimit,proto3" json:"temporary_allowed_total_memory_time_limit,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_serialize_v1_servers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_serialize_v1_servers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_v8platform_serialize_v1_servers_proto_rawDescGZIP(), []int{0}
}

func (x *ServerInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ServerInfo) GetAgentHost() string {
	if x != nil {
		return x.AgentHost
	}
	return ""
}

func (x *ServerInfo) GetAgentPort() int32 {
	if x != nil {
		return x.AgentPort
	}
	return 0
}

func (x *ServerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServerInfo) GetMainServer() bool {
	if x != nil {
		return x.MainServer
	}
	return false
}

func (x *ServerInfo) GetSafeWorkingProcessesMemoryLimit() int64 {
	if x != nil {
		return x.SafeWorkingProcessesMemoryLimit
	}
	return 0
}

func (x *ServerInfo) GetSafeCallMemoryLimit() int64 {
	if x != nil {
		return x.SafeCallMemoryLimit
	}
	return 0
}

func (x *ServerInfo) GetInfobasesLimit() int32 {
	if x != nil {
		return x.InfobasesLimit
	}
	return 0
}

func (x *ServerInfo) GetMemoryLimit() int64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *ServerInfo) GetConnectionsLimit() int32 {
	if x != nil {
		return x.ConnectionsLimit
	}
	return 0
}

func (x *ServerInfo) GetClusterPort() int32 {
	if x != nil {
		return x.ClusterPort
	}
	return 0
}

func (x *ServerInfo) GetDedicateManagers() bool {
	if x != nil {
		return x.DedicateManagers
	}
	return false
}

func (x *ServerInfo) GetPortRanges() []*PortRange {
	if x != nil {
		return x.PortRanges
	}
	return nil
}

func (x *ServerInfo) GetCriticalTotalMemory() int64 {
	if x != nil {
		return x.CriticalTotalMemory
	}
	return 0
}

func (x *ServerInfo) GetTemporaryAllowedTotalMemory() int64 {
	if x != nil {
		return x.TemporaryAllowedTotalMemory
	}
	return 0
}

func (x *ServerInfo) GetTemporaryAllowedTotalMemoryTimeLimit() int64 {
	if x != nil {
		return x.TemporaryAllowedTotalMemoryTimeLimit
	}
	return 0
}

type PortRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min int32 `protobuf:"varint,1,opt,name=Min,proto3" json:"Min,omitempty"`
	Max int32 `protobuf:"varint,2,opt,name=Max,proto3" json:"Max,omitempty"`
}

func (x *PortRange) Reset() {
	*x = PortRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_serialize_v1_servers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortRange) ProtoMessage() {}

func (x *PortRange) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_serialize_v1_servers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortRange.ProtoReflect.Descriptor instead.
func (*PortRange) Descriptor() ([]byte, []int) {
	return file_v8platform_serialize_v1_servers_proto_rawDescGZIP(), []int{1}
}

func (x *PortRange) GetMin() int32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *PortRange) GetMax() int32 {
	if x != nil {
		return x.Max
	}
	return 0
}

var File_v8platform_serialize_v1_servers_proto protoreflect.FileDescriptor

var file_v8platform_serialize_v1_servers_proto_rawDesc = []byte{
	0x0a, 0x25, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f,
	0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x76, 0x38, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa8, 0x07, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x22, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e,
	0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10, 0x01, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02,
	0x10, 0x02, 0x52, 0x09, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x27, 0x0a,
	0x0a, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x03, 0x52, 0x09, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x04, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e,
	0x02, 0x10, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x5c, 0x0a, 0x23, 0x73, 0x61, 0x66, 0x65, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0x82, 0xf5,
	0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x10, 0x06, 0x52, 0x1f, 0x73, 0x61,
	0x66, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x3d, 0x0a,
	0x16, 0x73, 0x61, 0x66, 0x65, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0x82,
	0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x07, 0x52, 0x13, 0x73, 0x61, 0x66, 0x65, 0x43, 0x61, 0x6c,
	0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x31, 0x0a, 0x0f,
	0x69, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x08, 0x52,
	0x0e, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x2b, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x09, 0x52,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x35, 0x0a, 0x11,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10,
	0x0a, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x2b, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e,
	0x02, 0x10, 0x0b, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x35, 0x0a, 0x11, 0x64, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x42, 0x08, 0x82, 0xf5, 0xea,
	0x94, 0x0e, 0x02, 0x10, 0x0c, 0x52, 0x10, 0x64, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x73, 0x12, 0x4d, 0x0a, 0x0b, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76,
	0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x0d, 0x52, 0x0a, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x15, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63,
	0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x0e, 0x18,
	0x08, 0x52, 0x13, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x4f, 0x0a, 0x1e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x72, 0x79, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a,
	0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x0f, 0x18, 0x08, 0x52, 0x1b, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x63, 0x0a, 0x29, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94,
	0x0e, 0x04, 0x10, 0x10, 0x18, 0x08, 0x52, 0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72,
	0x79, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x51, 0x0a, 0x09,
	0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x4d, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x09, 0x0a, 0x05,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x10, 0x01, 0x52, 0x03, 0x4d, 0x69, 0x6e, 0x12, 0x21, 0x0a, 0x03,
	0x4d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0x82, 0xf5, 0xea, 0x94, 0x0e,
	0x09, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x10, 0x02, 0x52, 0x03, 0x4d, 0x61, 0x78, 0x42,
	0xef, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
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
	file_v8platform_serialize_v1_servers_proto_rawDescOnce sync.Once
	file_v8platform_serialize_v1_servers_proto_rawDescData = file_v8platform_serialize_v1_servers_proto_rawDesc
)

func file_v8platform_serialize_v1_servers_proto_rawDescGZIP() []byte {
	file_v8platform_serialize_v1_servers_proto_rawDescOnce.Do(func() {
		file_v8platform_serialize_v1_servers_proto_rawDescData = protoimpl.X.CompressGZIP(file_v8platform_serialize_v1_servers_proto_rawDescData)
	})
	return file_v8platform_serialize_v1_servers_proto_rawDescData
}

var file_v8platform_serialize_v1_servers_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v8platform_serialize_v1_servers_proto_goTypes = []interface{}{
	(*ServerInfo)(nil), // 0: v8platform.serialize.v1.ServerInfo
	(*PortRange)(nil),  // 1: v8platform.serialize.v1.PortRange
}
var file_v8platform_serialize_v1_servers_proto_depIdxs = []int32{
	1, // 0: v8platform.serialize.v1.ServerInfo.port_ranges:type_name -> v8platform.serialize.v1.PortRange
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v8platform_serialize_v1_servers_proto_init() }
func file_v8platform_serialize_v1_servers_proto_init() {
	if File_v8platform_serialize_v1_servers_proto != nil {
		return
	}
	file_v8platform_serialize_v1_licanses_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v8platform_serialize_v1_servers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
		file_v8platform_serialize_v1_servers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortRange); i {
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
			RawDescriptor: file_v8platform_serialize_v1_servers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v8platform_serialize_v1_servers_proto_goTypes,
		DependencyIndexes: file_v8platform_serialize_v1_servers_proto_depIdxs,
		MessageInfos:      file_v8platform_serialize_v1_servers_proto_msgTypes,
	}.Build()
	File_v8platform_serialize_v1_servers_proto = out.File
	file_v8platform_serialize_v1_servers_proto_rawDesc = nil
	file_v8platform_serialize_v1_servers_proto_goTypes = nil
	file_v8platform_serialize_v1_servers_proto_depIdxs = nil
}
