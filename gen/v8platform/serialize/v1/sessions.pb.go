// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: v8platform/serialize/v1/sessions.proto

package serializev1

import (
	_ "github.com/v8platform/protoc-gen-go-ras/plugin/ras/encoding"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SessionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid                          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	AppId                         string                 `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	BlockedByDbms                 int32                  `protobuf:"varint,3,opt,name=blocked_by_dbms,json=blockedByDbms,proto3" json:"blocked_by_dbms,omitempty"`
	BlockedByLs                   int32                  `protobuf:"varint,4,opt,name=blocked_by_ls,json=blockedByLs,proto3" json:"blocked_by_ls,omitempty"`
	BytesAll                      int64                  `protobuf:"varint,5,opt,name=bytes_all,json=bytesAll,proto3" json:"bytes_all,omitempty"`
	BytesLast5Min                 int64                  `protobuf:"varint,6,opt,name=bytes_last5min,json=bytesLast5min,proto3" json:"bytes_last5min,omitempty"`
	CallsAll                      int32                  `protobuf:"varint,7,opt,name=calls_all,json=callsAll,proto3" json:"calls_all,omitempty"`
	CallsLast5Min                 int64                  `protobuf:"varint,8,opt,name=calls_last5min,json=callsLast5min,proto3" json:"calls_last5min,omitempty"`
	ConnectionId                  string                 `protobuf:"bytes,9,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	DbmsBytesAll                  int64                  `protobuf:"varint,10,opt,name=dbms_bytes_all,json=dbmsBytesAll,proto3" json:"dbms_bytes_all,omitempty"`
	DbmsBytesLast5Min             int64                  `protobuf:"varint,11,opt,name=dbms_bytes_last5min,json=dbmsBytesLast5min,proto3" json:"dbms_bytes_last5min,omitempty"`
	DbProcInfo                    string                 `protobuf:"bytes,12,opt,name=db_proc_info,json=dbProcInfo,proto3" json:"db_proc_info,omitempty"`
	DbProcTook                    int32                  `protobuf:"varint,13,opt,name=db_proc_took,json=dbProcTook,proto3" json:"db_proc_took,omitempty"`
	DbProcTookAt                  *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=db_proc_took_at,json=dbProcTookAt,proto3" json:"db_proc_took_at,omitempty"`
	DurationAll                   int32                  `protobuf:"varint,15,opt,name=duration_all,json=durationAll,proto3" json:"duration_all,omitempty"`
	DurationAllDbms               int32                  `protobuf:"varint,16,opt,name=duration_all_dbms,json=durationAllDbms,proto3" json:"duration_all_dbms,omitempty"`
	DurationCurrent               int32                  `protobuf:"varint,17,opt,name=duration_current,json=durationCurrent,proto3" json:"duration_current,omitempty"`
	DurationCurrentDbms           int32                  `protobuf:"varint,18,opt,name=duration_current_dbms,json=durationCurrentDbms,proto3" json:"duration_current_dbms,omitempty"`
	DurationLast_5Min             int64                  `protobuf:"varint,19,opt,name=duration_last_5_min,json=durationLast5Min,proto3" json:"duration_last_5_min,omitempty"`
	DurationLast_5MinDbms         int64                  `protobuf:"varint,20,opt,name=duration_last_5_min_dbms,json=durationLast5MinDbms,proto3" json:"duration_last_5_min_dbms,omitempty"`
	Host                          string                 `protobuf:"bytes,21,opt,name=host,proto3" json:"host,omitempty"`
	InfobaseId                    string                 `protobuf:"bytes,22,opt,name=infobase_id,json=infobaseId,proto3" json:"infobase_id,omitempty"`
	LastActiveAt                  *timestamppb.Timestamp `protobuf:"bytes,23,opt,name=last_active_at,json=lastActiveAt,proto3" json:"last_active_at,omitempty"`
	Hibernate                     bool                   `protobuf:"varint,24,opt,name=hibernate,proto3" json:"hibernate,omitempty"`
	PassiveSessionHibernateTime   int32                  `protobuf:"varint,25,opt,name=passive_session_hibernate_time,json=passiveSessionHibernateTime,proto3" json:"passive_session_hibernate_time,omitempty"`
	HibernateSessionTerminateTime int32                  `protobuf:"varint,26,opt,name=hibernate_session_terminate_time,json=hibernateSessionTerminateTime,proto3" json:"hibernate_session_terminate_time,omitempty"`
	Licenses                      []*LicenseInfo         `protobuf:"bytes,27,rep,name=licenses,proto3" json:"licenses,omitempty"`
	Locale                        string                 `protobuf:"bytes,28,opt,name=locale,proto3" json:"locale,omitempty"`
	ProcessId                     string                 `protobuf:"bytes,29,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	Id                            int32                  `protobuf:"varint,30,opt,name=id,proto3" json:"id,omitempty"`
	StartedAt                     *timestamppb.Timestamp `protobuf:"bytes,31,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	UserName                      string                 `protobuf:"bytes,32,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// version >= 4
	MemoryCurrent  int64 `protobuf:"varint,33,opt,name=memory_current,json=memoryCurrent,proto3" json:"memory_current,omitempty"`
	MemoryLast5Min int64 `protobuf:"varint,34,opt,name=memory_last5min,json=memoryLast5min,proto3" json:"memory_last5min,omitempty"`
	MemoryTotal    int64 `protobuf:"varint,35,opt,name=memory_total,json=memoryTotal,proto3" json:"memory_total,omitempty"`
	ReadCurrent    int64 `protobuf:"varint,36,opt,name=read_current,json=readCurrent,proto3" json:"read_current,omitempty"`
	ReadLast5Min   int64 `protobuf:"varint,37,opt,name=read_last5min,json=readLast5min,proto3" json:"read_last5min,omitempty"`
	ReadTotal      int64 `protobuf:"varint,38,opt,name=read_total,json=readTotal,proto3" json:"read_total,omitempty"`
	WriteCurrent   int64 `protobuf:"varint,39,opt,name=write_current,json=writeCurrent,proto3" json:"write_current,omitempty"`
	WriteLast5Min  int64 `protobuf:"varint,40,opt,name=write_last5min,json=writeLast5min,proto3" json:"write_last5min,omitempty"`
	WriteTotal     int64 `protobuf:"varint,41,opt,name=write_total,json=writeTotal,proto3" json:"write_total,omitempty"`
	// version >= 5
	DurationCurrentService  int32  `protobuf:"varint,42,opt,name=duration_current_service,json=durationCurrentService,proto3" json:"duration_current_service,omitempty"`
	DurationLast5MinService int64  `protobuf:"varint,43,opt,name=duration_last5min_service,json=durationLast5minService,proto3" json:"duration_last5min_service,omitempty"`
	DurationAllService      int32  `protobuf:"varint,44,opt,name=duration_all_service,json=durationAllService,proto3" json:"duration_all_service,omitempty"`
	CurrentServiceName      string `protobuf:"bytes,45,opt,name=current_service_name,json=currentServiceName,proto3" json:"current_service_name,omitempty"`
	// version >= 6
	CpuTimeCurrent  int64 `protobuf:"varint,46,opt,name=cpu_time_current,json=cpuTimeCurrent,proto3" json:"cpu_time_current,omitempty"`
	CpuTimeLast5Min int64 `protobuf:"varint,47,opt,name=cpu_time_last5min,json=cpuTimeLast5min,proto3" json:"cpu_time_last5min,omitempty"`
	CpuTimeTotal    int64 `protobuf:"varint,48,opt,name=cpu_time_total,json=cpuTimeTotal,proto3" json:"cpu_time_total,omitempty"`
	// version >= 7
	DataSeparation string `protobuf:"bytes,49,opt,name=data_separation,json=dataSeparation,proto3" json:"data_separation,omitempty"`
	// version >= 10
	ClientIpAddress string `protobuf:"bytes,50,opt,name=client_ip_address,json=clientIpAddress,proto3" json:"client_ip_address,omitempty"`
	// ignore
	ClusterId string `protobuf:"bytes,51,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *SessionInfo) Reset() {
	*x = SessionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v8platform_serialize_v1_sessions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionInfo) ProtoMessage() {}

func (x *SessionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v8platform_serialize_v1_sessions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionInfo.ProtoReflect.Descriptor instead.
func (*SessionInfo) Descriptor() ([]byte, []int) {
	return file_v8platform_serialize_v1_sessions_proto_rawDescGZIP(), []int{0}
}

func (x *SessionInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SessionInfo) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *SessionInfo) GetBlockedByDbms() int32 {
	if x != nil {
		return x.BlockedByDbms
	}
	return 0
}

func (x *SessionInfo) GetBlockedByLs() int32 {
	if x != nil {
		return x.BlockedByLs
	}
	return 0
}

func (x *SessionInfo) GetBytesAll() int64 {
	if x != nil {
		return x.BytesAll
	}
	return 0
}

func (x *SessionInfo) GetBytesLast5Min() int64 {
	if x != nil {
		return x.BytesLast5Min
	}
	return 0
}

func (x *SessionInfo) GetCallsAll() int32 {
	if x != nil {
		return x.CallsAll
	}
	return 0
}

func (x *SessionInfo) GetCallsLast5Min() int64 {
	if x != nil {
		return x.CallsLast5Min
	}
	return 0
}

func (x *SessionInfo) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *SessionInfo) GetDbmsBytesAll() int64 {
	if x != nil {
		return x.DbmsBytesAll
	}
	return 0
}

func (x *SessionInfo) GetDbmsBytesLast5Min() int64 {
	if x != nil {
		return x.DbmsBytesLast5Min
	}
	return 0
}

func (x *SessionInfo) GetDbProcInfo() string {
	if x != nil {
		return x.DbProcInfo
	}
	return ""
}

func (x *SessionInfo) GetDbProcTook() int32 {
	if x != nil {
		return x.DbProcTook
	}
	return 0
}

func (x *SessionInfo) GetDbProcTookAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DbProcTookAt
	}
	return nil
}

func (x *SessionInfo) GetDurationAll() int32 {
	if x != nil {
		return x.DurationAll
	}
	return 0
}

func (x *SessionInfo) GetDurationAllDbms() int32 {
	if x != nil {
		return x.DurationAllDbms
	}
	return 0
}

func (x *SessionInfo) GetDurationCurrent() int32 {
	if x != nil {
		return x.DurationCurrent
	}
	return 0
}

func (x *SessionInfo) GetDurationCurrentDbms() int32 {
	if x != nil {
		return x.DurationCurrentDbms
	}
	return 0
}

func (x *SessionInfo) GetDurationLast_5Min() int64 {
	if x != nil {
		return x.DurationLast_5Min
	}
	return 0
}

func (x *SessionInfo) GetDurationLast_5MinDbms() int64 {
	if x != nil {
		return x.DurationLast_5MinDbms
	}
	return 0
}

func (x *SessionInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SessionInfo) GetInfobaseId() string {
	if x != nil {
		return x.InfobaseId
	}
	return ""
}

func (x *SessionInfo) GetLastActiveAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastActiveAt
	}
	return nil
}

func (x *SessionInfo) GetHibernate() bool {
	if x != nil {
		return x.Hibernate
	}
	return false
}

func (x *SessionInfo) GetPassiveSessionHibernateTime() int32 {
	if x != nil {
		return x.PassiveSessionHibernateTime
	}
	return 0
}

func (x *SessionInfo) GetHibernateSessionTerminateTime() int32 {
	if x != nil {
		return x.HibernateSessionTerminateTime
	}
	return 0
}

func (x *SessionInfo) GetLicenses() []*LicenseInfo {
	if x != nil {
		return x.Licenses
	}
	return nil
}

func (x *SessionInfo) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *SessionInfo) GetProcessId() string {
	if x != nil {
		return x.ProcessId
	}
	return ""
}

func (x *SessionInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SessionInfo) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *SessionInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *SessionInfo) GetMemoryCurrent() int64 {
	if x != nil {
		return x.MemoryCurrent
	}
	return 0
}

func (x *SessionInfo) GetMemoryLast5Min() int64 {
	if x != nil {
		return x.MemoryLast5Min
	}
	return 0
}

func (x *SessionInfo) GetMemoryTotal() int64 {
	if x != nil {
		return x.MemoryTotal
	}
	return 0
}

func (x *SessionInfo) GetReadCurrent() int64 {
	if x != nil {
		return x.ReadCurrent
	}
	return 0
}

func (x *SessionInfo) GetReadLast5Min() int64 {
	if x != nil {
		return x.ReadLast5Min
	}
	return 0
}

func (x *SessionInfo) GetReadTotal() int64 {
	if x != nil {
		return x.ReadTotal
	}
	return 0
}

func (x *SessionInfo) GetWriteCurrent() int64 {
	if x != nil {
		return x.WriteCurrent
	}
	return 0
}

func (x *SessionInfo) GetWriteLast5Min() int64 {
	if x != nil {
		return x.WriteLast5Min
	}
	return 0
}

func (x *SessionInfo) GetWriteTotal() int64 {
	if x != nil {
		return x.WriteTotal
	}
	return 0
}

func (x *SessionInfo) GetDurationCurrentService() int32 {
	if x != nil {
		return x.DurationCurrentService
	}
	return 0
}

func (x *SessionInfo) GetDurationLast5MinService() int64 {
	if x != nil {
		return x.DurationLast5MinService
	}
	return 0
}

func (x *SessionInfo) GetDurationAllService() int32 {
	if x != nil {
		return x.DurationAllService
	}
	return 0
}

func (x *SessionInfo) GetCurrentServiceName() string {
	if x != nil {
		return x.CurrentServiceName
	}
	return ""
}

func (x *SessionInfo) GetCpuTimeCurrent() int64 {
	if x != nil {
		return x.CpuTimeCurrent
	}
	return 0
}

func (x *SessionInfo) GetCpuTimeLast5Min() int64 {
	if x != nil {
		return x.CpuTimeLast5Min
	}
	return 0
}

func (x *SessionInfo) GetCpuTimeTotal() int64 {
	if x != nil {
		return x.CpuTimeTotal
	}
	return 0
}

func (x *SessionInfo) GetDataSeparation() string {
	if x != nil {
		return x.DataSeparation
	}
	return ""
}

func (x *SessionInfo) GetClientIpAddress() string {
	if x != nil {
		return x.ClientIpAddress
	}
	return ""
}

func (x *SessionInfo) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

var File_v8platform_serialize_v1_sessions_proto protoreflect.FileDescriptor

var file_v8platform_serialize_v1_sessions_proto_rawDesc = []byte{
	0x0a, 0x26, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x72, 0x61, 0x73, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x2f, 0x72, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x76, 0x38, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x63, 0x61, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xfa, 0x14, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0e, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10, 0x01,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x02,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x64, 0x62, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x03, 0x52, 0x0d, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x42, 0x79, 0x44, 0x62, 0x6d, 0x73, 0x12, 0x2c, 0x0a, 0x0d, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x42, 0x79, 0x4c, 0x73, 0x12, 0x25, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94,
	0x0e, 0x02, 0x10, 0x05, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x12, 0x2f,
	0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x06,
	0x52, 0x0d, 0x62, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e, 0x12,
	0x25, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x07, 0x52, 0x08, 0x63, 0x61,
	0x6c, 0x6c, 0x73, 0x41, 0x6c, 0x6c, 0x12, 0x2f, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x5f,
	0x6c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x42, 0x08,
	0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x08, 0x52, 0x0d, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x4c,
	0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e, 0x12, 0x33, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e,
	0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10, 0x09, 0x52, 0x0c,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0e,
	0x64, 0x62, 0x6d, 0x73, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x0a, 0x52, 0x0c,
	0x64, 0x62, 0x6d, 0x73, 0x42, 0x79, 0x74, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x12, 0x38, 0x0a, 0x13,
	0x64, 0x62, 0x6d, 0x73, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x35,
	0x6d, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e,
	0x02, 0x10, 0x0b, 0x52, 0x11, 0x64, 0x62, 0x6d, 0x73, 0x42, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x61,
	0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e, 0x12, 0x2a, 0x0a, 0x0c, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f,
	0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5,
	0xea, 0x94, 0x0e, 0x02, 0x10, 0x0c, 0x52, 0x0a, 0x64, 0x62, 0x50, 0x72, 0x6f, 0x63, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x0c, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x5f, 0x74, 0x6f,
	0x6f, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02,
	0x10, 0x0d, 0x52, 0x0a, 0x64, 0x62, 0x50, 0x72, 0x6f, 0x63, 0x54, 0x6f, 0x6f, 0x6b, 0x12, 0x51,
	0x0a, 0x0f, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x5f, 0x74, 0x6f, 0x6f, 0x6b, 0x5f, 0x61,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x0e, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x10, 0x0e, 0x52, 0x0c, 0x64, 0x62, 0x50, 0x72, 0x6f, 0x63, 0x54, 0x6f, 0x6f, 0x6b, 0x41,
	0x74, 0x12, 0x2b, 0x0a, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6c,
	0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10,
	0x0f, 0x52, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x12, 0x34,
	0x0a, 0x11, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x64,
	0x62, 0x6d, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e,
	0x02, 0x10, 0x10, 0x52, 0x0f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c,
	0x44, 0x62, 0x6d, 0x73, 0x12, 0x33, 0x0a, 0x10, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08,
	0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x11, 0x52, 0x0f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x15, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x62,
	0x6d, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02,
	0x10, 0x12, 0x52, 0x13, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x44, 0x62, 0x6d, 0x73, 0x12, 0x37, 0x0a, 0x13, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x35, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x13, 0x52, 0x10,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x73, 0x74, 0x35, 0x4d, 0x69, 0x6e,
	0x12, 0x40, 0x0a, 0x18, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x35, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x64, 0x62, 0x6d, 0x73, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x14, 0x52, 0x14, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x73, 0x74, 0x35, 0x4d, 0x69, 0x6e, 0x44, 0x62,
	0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x15, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x2f, 0x0a, 0x0b, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x10, 0x16, 0x52, 0x0a, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x61, 0x73, 0x65, 0x49,
	0x64, 0x12, 0x50, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0e, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x10, 0x17, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x09, 0x68, 0x69, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x18,
	0x52, 0x09, 0x68, 0x69, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x4d, 0x0a, 0x1e, 0x70,
	0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x68,
	0x69, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x19, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x19, 0x52, 0x1b, 0x70,
	0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x62,
	0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x51, 0x0a, 0x20, 0x68, 0x69,
	0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x1a,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x1a, 0x52, 0x1d,
	0x68, 0x69, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4a, 0x0a,
	0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x1b, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x1b, 0x52,
	0x08, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e,
	0x02, 0x10, 0x1c, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0e, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x10, 0x1d, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x1e,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x49, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x0e, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x08, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x10, 0x1f, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x25, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x20, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x02, 0x10, 0x20, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x21, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a,
	0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x21, 0x18, 0x04, 0x52, 0x0d, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x0f, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e, 0x18, 0x22, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x22, 0x18, 0x04, 0x52, 0x0e,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e, 0x12, 0x2d,
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x23,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x23, 0x18, 0x04,
	0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2d, 0x0a,
	0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x24, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x24, 0x18, 0x04, 0x52,
	0x0b, 0x72, 0x65, 0x61, 0x64, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x0d,
	0x72, 0x65, 0x61, 0x64, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e, 0x18, 0x25, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x25, 0x18, 0x04, 0x52,
	0x0c, 0x72, 0x65, 0x61, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e, 0x12, 0x29, 0x0a,
	0x0a, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x26, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x26, 0x18, 0x04, 0x52, 0x09, 0x72,
	0x65, 0x61, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x27, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x27, 0x18, 0x04, 0x52, 0x0c, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x0e, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x28, 0x18, 0x04, 0x52, 0x0d, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x0b,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x29, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x29, 0x18, 0x04, 0x52, 0x0a, 0x77,
	0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x44, 0x0a, 0x18, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x82, 0xf5, 0xea,
	0x94, 0x0e, 0x04, 0x10, 0x2a, 0x18, 0x05, 0x52, 0x16, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x46, 0x0a, 0x19, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x61, 0x73, 0x74,
	0x35, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x2b, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x2b, 0x18, 0x05, 0x52, 0x17,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x14, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x2c, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x2c, 0x18,
	0x05, 0x52, 0x12, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x14, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x2d, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x2d, 0x18, 0x05, 0x52,
	0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x10, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0x82,
	0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x2e, 0x18, 0x06, 0x52, 0x0e, 0x63, 0x70, 0x75, 0x54, 0x69,
	0x6d, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x11, 0x63, 0x70, 0x75,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69, 0x6e, 0x18, 0x2f,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x2f, 0x18, 0x06,
	0x52, 0x0f, 0x63, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x35, 0x6d, 0x69,
	0x6e, 0x12, 0x30, 0x0a, 0x0e, 0x63, 0x70, 0x75, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x30, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e,
	0x04, 0x10, 0x30, 0x18, 0x06, 0x52, 0x0c, 0x63, 0x70, 0x75, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x70, 0x61,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x31, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x82, 0xf5,
	0xea, 0x94, 0x0e, 0x04, 0x10, 0x31, 0x18, 0x07, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0x82, 0xf5, 0xea, 0x94, 0x0e, 0x04, 0x10, 0x32, 0x18, 0x0a, 0x52,
	0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x33,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x42,
	0xf0, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x0d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x38, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x76, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x53, 0x58, 0xaa, 0x02, 0x17, 0x56,
	0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x56, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x5c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x23, 0x56, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x53, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x56, 0x38, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v8platform_serialize_v1_sessions_proto_rawDescOnce sync.Once
	file_v8platform_serialize_v1_sessions_proto_rawDescData = file_v8platform_serialize_v1_sessions_proto_rawDesc
)

func file_v8platform_serialize_v1_sessions_proto_rawDescGZIP() []byte {
	file_v8platform_serialize_v1_sessions_proto_rawDescOnce.Do(func() {
		file_v8platform_serialize_v1_sessions_proto_rawDescData = protoimpl.X.CompressGZIP(file_v8platform_serialize_v1_sessions_proto_rawDescData)
	})
	return file_v8platform_serialize_v1_sessions_proto_rawDescData
}

var file_v8platform_serialize_v1_sessions_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v8platform_serialize_v1_sessions_proto_goTypes = []interface{}{
	(*SessionInfo)(nil),           // 0: v8platform.serialize.v1.SessionInfo
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*LicenseInfo)(nil),           // 2: v8platform.serialize.v1.LicenseInfo
}
var file_v8platform_serialize_v1_sessions_proto_depIdxs = []int32{
	1, // 0: v8platform.serialize.v1.SessionInfo.db_proc_took_at:type_name -> google.protobuf.Timestamp
	1, // 1: v8platform.serialize.v1.SessionInfo.last_active_at:type_name -> google.protobuf.Timestamp
	2, // 2: v8platform.serialize.v1.SessionInfo.licenses:type_name -> v8platform.serialize.v1.LicenseInfo
	1, // 3: v8platform.serialize.v1.SessionInfo.started_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v8platform_serialize_v1_sessions_proto_init() }
func file_v8platform_serialize_v1_sessions_proto_init() {
	if File_v8platform_serialize_v1_sessions_proto != nil {
		return
	}
	file_v8platform_serialize_v1_licanses_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v8platform_serialize_v1_sessions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionInfo); i {
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
			RawDescriptor: file_v8platform_serialize_v1_sessions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v8platform_serialize_v1_sessions_proto_goTypes,
		DependencyIndexes: file_v8platform_serialize_v1_sessions_proto_depIdxs,
		MessageInfos:      file_v8platform_serialize_v1_sessions_proto_msgTypes,
	}.Build()
	File_v8platform_serialize_v1_sessions_proto = out.File
	file_v8platform_serialize_v1_sessions_proto_rawDesc = nil
	file_v8platform_serialize_v1_sessions_proto_goTypes = nil
	file_v8platform_serialize_v1_sessions_proto_depIdxs = nil
}
