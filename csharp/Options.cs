// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: ras/protocol/v1/options.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Ras.Protocol.V1 {

  /// <summary>Holder for reflection information generated from ras/protocol/v1/options.proto</summary>
  public static partial class OptionsReflection {

    #region Descriptor
    /// <summary>File descriptor for ras/protocol/v1/options.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static OptionsReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Ch1yYXMvcHJvdG9jb2wvdjEvb3B0aW9ucy5wcm90bxIPcmFzLnByb3RvY29s",
            "LnYxGiBnb29nbGUvcHJvdG9idWYvZGVzY3JpcHRvci5wcm90bxobcmFzL3By",
            "b3RvY29sL3YxL3R5cGVzLnByb3RvOmEKC3BhY2tldF90eXBlEh8uZ29vZ2xl",
            "LnByb3RvYnVmLk1lc3NhZ2VPcHRpb25zGNGuzeIBIAEoDjIbLnJhcy5wcm90",
            "b2NvbC52MS5QYWNrZXRUeXBlUgpwYWNrZXRUeXBlOnQKEmVuZHBvaW50X2Rh",
            "dGFfdHlwZRIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxjSrs3i",
            "ASABKA4yIS5yYXMucHJvdG9jb2wudjEuRW5kcG9pbnREYXRhVHlwZVIQZW5k",
            "cG9pbnREYXRhVHlwZUK+AQoTY29tLnJhcy5wcm90b2NvbC52MUIMT3B0aW9u",
            "c1Byb3RvUAFaO2dpdGh1Yi5jb20vdjhwbGF0Zm9ybS9wcm90b3MvZ2VuL3Jh",
            "cy9wcm90b2NvbC92MTtwcm90b2NvbHYxogIDUlBYqgIPUmFzLlByb3RvY29s",
            "LlYxygIPUmFzXFByb3RvY29sXFYx4gIbUmFzXFByb3RvY29sXFYxXEdQQk1l",
            "dGFkYXRh6gIRUmFzOjpQcm90b2NvbDo6VjFiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Google.Protobuf.Reflection.DescriptorReflection.Descriptor, global::Ras.Protocol.V1.TypesReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, new pb::Extension[] { OptionsExtensions.PacketType, OptionsExtensions.EndpointDataType }, null));
    }
    #endregion

  }
  /// <summary>Holder for extension identifiers generated from the top level of ras/protocol/v1/options.proto</summary>
  public static partial class OptionsExtensions {
    public static readonly pb::Extension<global::Google.Protobuf.Reflection.MessageOptions, global::Ras.Protocol.V1.PacketType> PacketType =
      new pb::Extension<global::Google.Protobuf.Reflection.MessageOptions, global::Ras.Protocol.V1.PacketType>(475223889, pb::FieldCodec.ForEnum(3801791112, x => (int) x, x => (global::Ras.Protocol.V1.PacketType) x, global::Ras.Protocol.V1.PacketType.Negotiate));
    public static readonly pb::Extension<global::Google.Protobuf.Reflection.MessageOptions, global::Ras.Protocol.V1.EndpointDataType> EndpointDataType =
      new pb::Extension<global::Google.Protobuf.Reflection.MessageOptions, global::Ras.Protocol.V1.EndpointDataType>(475223890, pb::FieldCodec.ForEnum(3801791120, x => (int) x, x => (global::Ras.Protocol.V1.EndpointDataType) x, global::Ras.Protocol.V1.EndpointDataType.VoidMessage));
  }

}

#endregion Designer generated code