// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: ras/protocol/v1/types.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Ras.Protocol.V1 {

  /// <summary>Holder for reflection information generated from ras/protocol/v1/types.proto</summary>
  public static partial class TypesReflection {

    #region Descriptor
    /// <summary>File descriptor for ras/protocol/v1/types.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static TypesReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChtyYXMvcHJvdG9jb2wvdjEvdHlwZXMucHJvdG8SD3Jhcy5wcm90b2NvbC52",
            "MRoWcmFzL2VuY29kaW5nL3Jhcy5wcm90byqcAgoQRW5kcG9pbnREYXRhVHlw",
            "ZRJYCh9FTkRQT0lOVF9EQVRBX1RZUEVfVk9JRF9NRVNTQUdFEAAaM5r16pQO",
            "LXJhc2JpbmFyeS5wcm90b2NvbC52MS5FbmRwb2ludFZvaWREYXRhTWVzc2Fn",
            "ZRJPChpFTkRQT0lOVF9EQVRBX1RZUEVfTUVTU0FHRRABGi+a9eqUDilyYXNi",
            "aW5hcnkucHJvdG9jb2wudjEuRW5kcG9pbnREYXRhTWVzc2FnZRJVChxFTkRQ",
            "T0lOVF9EQVRBX1RZUEVfRVhDRVBUSU9OEP8BGjKa9eqUDixyYXNiaW5hcnku",
            "cHJvdG9jb2wudjEuRW5kcG9pbnRGYWlsdXJlTWVzc2FnZRoGmPXqlA4BKo4H",
            "CgpQYWNrZXRUeXBlEkcKFVBBQ0tFVF9UWVBFX05FR09USUFURRAAGiya9eqU",
            "DiZyYXNiaW5hcnkucHJvdG9jb2wudjEuTmVnb3RpYXRlTWVzc2FnZRJDChNQ",
            "QUNLRVRfVFlQRV9DT05ORUNUEAEaKpr16pQOJHJhc2JpbmFyeS5wcm90b2Nv",
            "bC52MS5Db25uZWN0TWVzc2FnZRJKChdQQUNLRVRfVFlQRV9DT05ORUNUX0FD",
            "SxACGi2a9eqUDidyYXNiaW5hcnkucHJvdG9jb2wudjEuQ29ubmVjdE1lc3Nh",
            "Z2VBY2sSSQoWUEFDS0VUX1RZUEVfRElTQ09OTkVDVBAEGi2a9eqUDidyYXNi",
            "aW5hcnkucHJvdG9jb2wudjEuRGlzY29ubmVjdE1lc3NhZ2USRwoZUEFDS0VU",
            "X1RZUEVfRU5EUE9JTlRfT1BFThALGiia9eqUDiJyYXNiaW5hcnkucHJvdG9j",
            "b2wudjEuRW5kcG9pbnRPcGVuEk4KHVBBQ0tFVF9UWVBFX0VORFBPSU5UX09Q",
            "RU5fQUNLEAwaK5r16pQOJXJhc2JpbmFyeS5wcm90b2NvbC52MS5FbmRwb2lu",
            "dE9wZW5BY2sSSgoaUEFDS0VUX1RZUEVfRU5EUE9JTlRfQ0xPU0UQDRoqmvXq",
            "lA4kcmFzYmluYXJ5LnByb3RvY29sLnYxLkVuZHBvaW50T0Nsb3NlEk0KHFBB",
            "Q0tFVF9UWVBFX0VORFBPSU5UX01FU1NBR0UQDhormvXqlA4lcmFzYmluYXJ5",
            "LnByb3RvY29sLnYxLkVuZHBvaW50TWVzc2FnZRJNChxQQUNLRVRfVFlQRV9F",
            "TkRQT0lOVF9GQUlMVVJFEA8aK5r16pQOJXJhc2JpbmFyeS5wcm90b2NvbC52",
            "MS5FbmRwb2ludEZhaWx1cmUSGgoWUEFDS0VUX1RZUEVfS0VFUF9BTElWRRAQ",
            "IgQIAxADIgQIBRAKKhVQQUNLRVRfVFlQRV9TVEFSVF9UTFMqGlBBQ0tFVF9U",
            "WVBFX1NBU0xfTkVHT1RJQVRFKhVQQUNLRVRfVFlQRV9TQVNMX0FVVEgqGlBB",
            "Q0tFVF9UWVBFX1NBU0xfQ0hBTExFTkdFKhhQQUNLRVRfVFlQRV9TQVNMX1NV",
            "Q0NFU1MqGFBBQ0tFVF9UWVBFX1NBU0xfRkFJTFVSRSoWUEFDS0VUX1RZUEVf",
            "U0FTTF9BQk9SVEK8AQoTY29tLnJhcy5wcm90b2NvbC52MUIKVHlwZXNQcm90",
            "b1ABWjtnaXRodWIuY29tL3Y4cGxhdGZvcm0vcHJvdG9zL2dlbi9yYXMvcHJv",
            "dG9jb2wvdjE7cHJvdG9jb2x2MaICA1JQWKoCD1Jhcy5Qcm90b2NvbC5WMcoC",
            "D1Jhc1xQcm90b2NvbFxWMeICG1Jhc1xQcm90b2NvbFxWMVxHUEJNZXRhZGF0",
            "YeoCEVJhczo6UHJvdG9jb2w6OlYxYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Ras.Encoding.RasReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Ras.Protocol.V1.EndpointDataType), typeof(global::Ras.Protocol.V1.PacketType), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum EndpointDataType {
    [pbr::OriginalName("ENDPOINT_DATA_TYPE_VOID_MESSAGE")] VoidMessage = 0,
    [pbr::OriginalName("ENDPOINT_DATA_TYPE_MESSAGE")] Message = 1,
    [pbr::OriginalName("ENDPOINT_DATA_TYPE_EXCEPTION")] Exception = 255,
  }

  public enum PacketType {
    [pbr::OriginalName("PACKET_TYPE_NEGOTIATE")] Negotiate = 0,
    [pbr::OriginalName("PACKET_TYPE_CONNECT")] Connect = 1,
    [pbr::OriginalName("PACKET_TYPE_CONNECT_ACK")] ConnectAck = 2,
    [pbr::OriginalName("PACKET_TYPE_DISCONNECT")] Disconnect = 4,
    [pbr::OriginalName("PACKET_TYPE_ENDPOINT_OPEN")] EndpointOpen = 11,
    [pbr::OriginalName("PACKET_TYPE_ENDPOINT_OPEN_ACK")] EndpointOpenAck = 12,
    [pbr::OriginalName("PACKET_TYPE_ENDPOINT_CLOSE")] EndpointClose = 13,
    [pbr::OriginalName("PACKET_TYPE_ENDPOINT_MESSAGE")] EndpointMessage = 14,
    [pbr::OriginalName("PACKET_TYPE_ENDPOINT_FAILURE")] EndpointFailure = 15,
    [pbr::OriginalName("PACKET_TYPE_KEEP_ALIVE")] KeepAlive = 16,
  }

  #endregion

}

#endregion Designer generated code