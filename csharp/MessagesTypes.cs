// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: ras/messages/v1/messages_types.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Ras.Messages.V1 {

  /// <summary>Holder for reflection information generated from ras/messages/v1/messages_types.proto</summary>
  public static partial class MessagesTypesReflection {

    #region Descriptor
    /// <summary>File descriptor for ras/messages/v1/messages_types.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static MessagesTypesReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiRyYXMvbWVzc2FnZXMvdjEvbWVzc2FnZXNfdHlwZXMucHJvdG8SD3Jhcy5t",
            "ZXNzYWdlcy52MRoWcmFzL2VuY29kaW5nL3Jhcy5wcm90bxogZ29vZ2xlL3By",
            "b3RvYnVmL2Rlc2NyaXB0b3IucHJvdG8q6yIKC01lc3NhZ2VUeXBlEhwKGEdF",
            "VF9BR0VOVF9BRE1JTlNfUkVRVUVTVBAAEh0KGUdFVF9BR0VOVF9BRE1JTlNf",
            "UkVTUE9OU0UQARIeChpHRVRfQ0xVU1RFUl9BRE1JTlNfUkVRVUVTVBACEh8K",
            "G0dFVF9DTFVTVEVSX0FETUlOU19SRVNQT05TRRADEhsKF1JFR19BR0VOVF9B",
            "RE1JTl9SRVFVRVNUEAQSHQoZUkVHX0NMVVNURVJfQURNSU5fUkVRVUVTVBAF",
            "Eh0KGVVOUkVHX0FHRU5UX0FETUlOX1JFUVVFU1QQBhIfChtVTlJFR19DTFVT",
            "VEVSX0FETUlOX1JFUVVFU1QQBxIeChpBVVRIRU5USUNBVEVfQUdFTlRfUkVR",
            "VUVTVBAIEhgKFEFVVEhFTlRJQ0FURV9SRVFVRVNUEAkSHgoaQUREX0FVVEhF",
            "TlRJQ0FUSU9OX1JFUVVFU1QQChIYChRHRVRfQ0xVU1RFUlNfUkVRVUVTVBAL",
            "EhkKFUdFVF9DTFVTVEVSU19SRVNQT05TRRAMEhwKGEdFVF9DTFVTVEVSX0lO",
            "Rk9fUkVRVUVTVBANEh0KGUdFVF9DTFVTVEVSX0lORk9fUkVTUE9OU0UQDhIX",
            "ChNSRUdfQ0xVU1RFUl9SRVFVRVNUEA8SGAoUUkVHX0NMVVNURVJfUkVTUE9O",
            "U0UQEBIZChVVTlJFR19DTFVTVEVSX1JFUVVFU1QQERIgChxHRVRfQ0xVU1RF",
            "Ul9NQU5BR0VSU19SRVFVRVNUEBISIQodR0VUX0NMVVNURVJfTUFOQUdFUlNf",
            "UkVTUE9OU0UQExIkCiBHRVRfQ0xVU1RFUl9NQU5BR0VSX0lORk9fUkVRVUVT",
            "VBAUEiUKIUdFVF9DTFVTVEVSX01BTkFHRVJfSU5GT19SRVNQT05TRRAVEh8K",
            "G0dFVF9XT1JLSU5HX1NFUlZFUlNfUkVRVUVTVBAWEiAKHEdFVF9XT1JLSU5H",
            "X1NFUlZFUlNfUkVTUE9OU0UQFxIjCh9HRVRfV09SS0lOR19TRVJWRVJfSU5G",
            "T19SRVFVRVNUEBgSJAogR0VUX1dPUktJTkdfU0VSVkVSX0lORk9fUkVTUE9O",
            "U0UQGRIeChpSRUdfV09SS0lOR19TRVJWRVJfUkVRVUVTVBAaEh8KG1JFR19X",
            "T1JLSU5HX1NFUlZFUl9SRVNQT05TRRAbEiAKHFVOUkVHX1dPUktJTkdfU0VS",
            "VkVSX1JFUVVFU1QQHBIhCh1HRVRfV09SS0lOR19QUk9DRVNTRVNfUkVRVUVT",
            "VBAdEiIKHkdFVF9XT1JLSU5HX1BST0NFU1NFU19SRVNQT05TRRAeEiQKIEdF",
            "VF9XT1JLSU5HX1BST0NFU1NfSU5GT19SRVFVRVNUEB8SJQohR0VUX1dPUktJ",
            "TkdfUFJPQ0VTU19JTkZPX1JFU1BPTlNFECASKAokR0VUX1NFUlZFUl9XT1JL",
            "SU5HX1BST0NFU1NFU19SRVFVRVNUECESKQolR0VUX1NFUlZFUl9XT1JLSU5H",
            "X1BST0NFU1NFU19SRVNQT05TRRAiEiAKHEdFVF9DTFVTVEVSX1NFUlZJQ0VT",
            "X1JFUVVFU1QQIxIhCh1HRVRfQ0xVU1RFUl9TRVJWSUNFU19SRVNQT05TRRAk",
            "EhsKF0NSRUFURV9JTkZPQkFTRV9SRVFVRVNUECUSHAoYQ1JFQVRFX0lORk9C",
            "QVNFX1JFU1BPTlNFECYSIQodVVBEQVRFX0lORk9CQVNFX1NIT1JUX1JFUVVF",
            "U1QQJxIbChdVUERBVEVfSU5GT0JBU0VfUkVRVUVTVBAoEhkKFURST1BfSU5G",
            "T0JBU0VfUkVRVUVTVBApEh8KG0dFVF9JTkZPQkFTRVNfU0hPUlRfUkVRVUVT",
            "VBAqEiAKHEdFVF9JTkZPQkFTRVNfU0hPUlRfUkVTUE9OU0UQKxIZChVHRVRf",
            "SU5GT0JBU0VTX1JFUVVFU1QQLBIaChZHRVRfSU5GT0JBU0VTX1JFU1BPTlNF",
            "EC0SIwofR0VUX0lORk9CQVNFX1NIT1JUX0lORk9fUkVRVUVTVBAuEiQKIEdF",
            "VF9JTkZPQkFTRV9TSE9SVF9JTkZPX1JFU1BPTlNFEC8SHQoZR0VUX0lORk9C",
            "QVNFX0lORk9fUkVRVUVTVBAwEh4KGkdFVF9JTkZPQkFTRV9JTkZPX1JFU1BP",
            "TlNFEDESIQodR0VUX0NPTk5FQ1RJT05TX1NIT1JUX1JFUVVFU1QQMhIiCh5H",
            "RVRfQ09OTkVDVElPTlNfU0hPUlRfUkVTUE9OU0UQMxIqCiZHRVRfSU5GT0JB",
            "U0VfQ09OTkVDVElPTlNfU0hPUlRfUkVRVUVTVBA0EisKJ0dFVF9JTkZPQkFT",
            "RV9DT05ORUNUSU9OU19TSE9SVF9SRVNQT05TRRA1EiUKIUdFVF9DT05ORUNU",
            "SU9OX0lORk9fU0hPUlRfUkVRVUVTVBA2EiYKIkdFVF9DT05ORUNUSU9OX0lO",
            "Rk9fU0hPUlRfUkVTUE9OU0UQNxIkCiBHRVRfSU5GT0JBU0VfQ09OTkVDVElP",
            "TlNfUkVRVUVTVBA4EiUKIUdFVF9JTkZPQkFTRV9DT05ORUNUSU9OU19SRVNQ",
            "T05TRRA5EhYKEkRJU0NPTk5FQ1RfUkVRVUVTVBBAEhgKFEdFVF9TRVNTSU9O",
            "U19SRVFVRVNUEEESGQoVR0VUX1NFU1NJT05TX1JFU1BPTlNFEEISIQodR0VU",
            "X0lORk9CQVNFX1NFU1NJT05TX1JFUVVFU1QQQxIiCh5HRVRfSU5GT0JBU0Vf",
            "U0VTU0lPTlNfUkVTUE9OU0UQRBIcChhHRVRfU0VTU0lPTl9JTkZPX1JFUVVF",
            "U1QQRRIdChlHRVRfU0VTU0lPTl9JTkZPX1JFU1BPTlNFEEYSHQoZVEVSTUlO",
            "QVRFX1NFU1NJT05fUkVRVUVTVBBHEhUKEUdFVF9MT0NLU19SRVFVRVNUEEgS",
            "FgoSR0VUX0xPQ0tTX1JFU1BPTlNFEEkSHgoaR0VUX0lORk9CQVNFX0xPQ0tT",
            "X1JFUVVFU1QQShIfChtHRVRfSU5GT0JBU0VfTE9DS1NfUkVTUE9OU0UQSxIg",
            "ChxHRVRfQ09OTkVDVElPTl9MT0NLU19SRVFVRVNUEEwSIQodR0VUX0NPTk5F",
            "Q1RJT05fTE9DS1NfUkVTUE9OU0UQTRIdChlHRVRfU0VTU0lPTl9MT0NLU19S",
            "RVFVRVNUEE4SHgoaR0VUX1NFU1NJT05fTE9DS1NfUkVTUE9OU0UQTxIiCh5B",
            "UFBMWV9BU1NJR05NRU5UX1JVTEVTX1JFUVVFU1QQURIfChtSRUdfQVNTSUdO",
            "TUVOVF9SVUxFX1JFUVVFU1QQUhIgChxSRUdfQVNTSUdOTUVOVF9SVUxFX1JF",
            "U1BPTlNFEFMSIQodVU5SRUdfQVNTSUdOTUVOVF9SVUxFX1JFUVVFU1QQVBIg",
            "ChxHRVRfQVNTSUdOTUVOVF9SVUxFU19SRVFVRVNUEFUSIQodR0VUX0FTU0lH",
            "Tk1FTlRfUlVMRVNfUkVTUE9OU0UQVhIkCiBHRVRfQVNTSUdOTUVOVF9SVUxF",
            "X0lORk9fUkVRVUVTVBBXEiUKIUdFVF9BU1NJR05NRU5UX1JVTEVfSU5GT19S",
            "RVNQT05TRRBYEiEKHUdFVF9TRUNVUklUWV9QUk9GSUxFU19SRVFVRVNUEFkS",
            "IgoeR0VUX1NFQ1VSSVRZX1BST0ZJTEVTX1JFU1BPTlNFEFoSIwofQ1JFQVRF",
            "X1NFQ1VSSVRZX1BST0ZJTEVfUkVRVUVTVBBbEiEKHURST1BfU0VDVVJJVFlf",
            "UFJPRklMRV9SRVFVRVNUEFwSIwofR0VUX1ZJUlRVQUxfRElSRUNUT1JJRVNf",
            "UkVRVUVTVBBdEiQKIEdFVF9WSVJUVUFMX0RJUkVDVE9SSUVTX1JFU1BPTlNF",
            "EF4SJAogQ1JFQVRFX1ZJUlRVQUxfRElSRUNUT1JZX1JFUVVFU1QQXxIiCh5E",
            "Uk9QX1ZJUlRVQUxfRElSRUNUT1JZX1JFUVVFU1QQYBIbChdHRVRfQ09NX0NM",
            "QVNTRVNfUkVRVUVTVBBhEhwKGEdFVF9DT01fQ0xBU1NFU19SRVNQT05TRRBi",
            "EhwKGENSRUFURV9DT01fQ0xBU1NfUkVRVUVTVBBjEhoKFkRST1BfQ09NX0NM",
            "QVNTX1JFUVVFU1QQZBIeChpHRVRfQUxMT1dFRF9BRERJTlNfUkVRVUVTVBBl",
            "Eh8KG0dFVF9BTExPV0VEX0FERElOU19SRVNQT05TRRBmEiAKHENSRUFURV9B",
            "TExPV0VEX0FERElOX1JFUVVFU1QQZxIeChpEUk9QX0FMTE9XRURfQURESU5f",
            "UkVRVUVTVBBoEiAKHEdFVF9FWFRFUk5BTF9NT0RVTEVTX1JFUVVFU1QQaRIh",
            "Ch1HRVRfRVhURVJOQUxfTU9EVUxFU19SRVNQT05TRRBqEiIKHkNSRUFURV9F",
            "WFRFUk5BTF9NT0RVTEVfUkVRVUVTVBBrEiAKHERST1BfRVhURVJOQUxfTU9E",
            "VUxFX1JFUVVFU1QQbBIkCiBHRVRfQUxMT1dFRF9BUFBMSUNBVElPTlNfUkVR",
            "VUVTVBBtEiUKIUdFVF9BTExPV0VEX0FQUExJQ0FUSU9OU19SRVNQT05TRRBu",
            "EiYKIkNSRUFURV9BTExPV0VEX0FQUExJQ0FUSU9OX1JFUVVFU1QQbxIkCiBE",
            "Uk9QX0FMTE9XRURfQVBQTElDQVRJT05fUkVRVUVTVBBwEiIKHkdFVF9JTlRF",
            "Uk5FVF9SRVNPVVJDRVNfUkVRVUVTVBBxEiMKH0dFVF9JTlRFUk5FVF9SRVNP",
            "VVJDRVNfUkVTUE9OU0UQchIkCiBDUkVBVEVfSU5URVJORVRfUkVTT1VSQ0Vf",
            "UkVRVUVTVBBzEiIKHkRST1BfSU5URVJORVRfUkVTT1VSQ0VfUkVRVUVTVBB0",
            "EjEKLUlOVEVSUlVQVF9TRVNTSU9OX0NVUlJFTlRfU0VSVkVSX0NBTExfUkVR",
            "VUVTVBB1EiEKHUdFVF9SRVNPVVJDRV9DT1VOVEVSU19SRVFVRVNUEHYSIgoe",
            "R0VUX1JFU09VUkNFX0NPVU5URVJTX1JFU1BPTlNFEHcSJQohR0VUX1JFU09V",
            "UkNFX0NPVU5URVJfSU5GT19SRVFVRVNUEHgSJgoiR0VUX1JFU09VUkNFX0NP",
            "VU5URVJfSU5GT19SRVNQT05TRRB5EiAKHFJFR19SRVNPVVJDRV9DT1VOVEVS",
            "X1JFUVVFU1QQehIiCh5VTlJFR19SRVNPVVJDRV9DT1VOVEVSX1JFUVVFU1QQ",
            "exIfChtHRVRfUkVTT1VSQ0VfTElNSVRTX1JFUVVFU1QQfBIgChxHRVRfUkVT",
            "T1VSQ0VfTElNSVRTX1JFU1BPTlNFEH0SIwofR0VUX1JFU09VUkNFX0xJTUlU",
            "X0lORk9fUkVRVUVTVBB+EiQKIEdFVF9SRVNPVVJDRV9MSU1JVF9JTkZPX1JF",
            "U1BPTlNFEH8SHwoaUkVHX1JFU09VUkNFX0xJTUlUX1JFUVVFU1QQgAESIQoc",
            "VU5SRUdfUkVTT1VSQ0VfTElNSVRfUkVRVUVTVBCBARIfChpHRVRfQ09VTlRF",
            "Ul9WQUxVRVNfUkVRVUVTVBCCARIgChtHRVRfQ09VTlRFUl9WQUxVRVNfUkVT",
            "UE9OU0UQgwESIAobQ0xFQVJfQ09VTlRFUl9WQUxVRV9SRVFVRVNUEIQBEisK",
            "JkdFVF9DT1VOVEVSX0FDQ1VNVUxBVEVEX1ZBTFVFU19SRVFVRVNUEIUBEiwK",
            "J0dFVF9DT1VOVEVSX0FDQ1VNVUxBVEVEX1ZBTFVFU19SRVNQT05TRRCGARIe",
            "ChlHRVRfQUdFTlRfVkVSU0lPTl9SRVFVRVNUEIcBEh8KGkdFVF9BR0VOVF9W",
            "RVJTSU9OX1JFU1BPTlNFEIgBIgQIOhA/IgQIUBBQOmQKDG1lc3NhZ2VfdHlw",
            "ZRIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxjTrs3iASABKA4y",
            "HC5yYXMubWVzc2FnZXMudjEuTWVzc2FnZVR5cGVSC21lc3NhZ2VUeXBlQsQB",
            "ChNjb20ucmFzLm1lc3NhZ2VzLnYxQhJNZXNzYWdlc1R5cGVzUHJvdG9QAVo7",
            "Z2l0aHViLmNvbS92OHBsYXRmb3JtL3Byb3Rvcy9nZW4vcmFzL21lc3NhZ2Vz",
            "L3YxO21lc3NhZ2VzdjGiAgNSTViqAg9SYXMuTWVzc2FnZXMuVjHKAg9SYXNc",
            "TWVzc2FnZXNcVjHiAhtSYXNcTWVzc2FnZXNcVjFcR1BCTWV0YWRhdGHqAhFS",
            "YXM6Ok1lc3NhZ2VzOjpWMWIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Ras.Encoding.RasReflection.Descriptor, global::Google.Protobuf.Reflection.DescriptorReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Ras.Messages.V1.MessageType), }, new pb::Extension[] { MessagesTypesExtensions.MessageType }, null));
    }
    #endregion

  }
  /// <summary>Holder for extension identifiers generated from the top level of ras/messages/v1/messages_types.proto</summary>
  public static partial class MessagesTypesExtensions {
    public static readonly pb::Extension<global::Google.Protobuf.Reflection.MessageOptions, global::Ras.Messages.V1.MessageType> MessageType =
      new pb::Extension<global::Google.Protobuf.Reflection.MessageOptions, global::Ras.Messages.V1.MessageType>(475223891, pb::FieldCodec.ForEnum(3801791128, x => (int) x, x => (global::Ras.Messages.V1.MessageType) x, global::Ras.Messages.V1.MessageType.GetAgentAdminsRequest));
  }

  #region Enums
  public enum MessageType {
    [pbr::OriginalName("GET_AGENT_ADMINS_REQUEST")] GetAgentAdminsRequest = 0,
    [pbr::OriginalName("GET_AGENT_ADMINS_RESPONSE")] GetAgentAdminsResponse = 1,
    [pbr::OriginalName("GET_CLUSTER_ADMINS_REQUEST")] GetClusterAdminsRequest = 2,
    [pbr::OriginalName("GET_CLUSTER_ADMINS_RESPONSE")] GetClusterAdminsResponse = 3,
    [pbr::OriginalName("REG_AGENT_ADMIN_REQUEST")] RegAgentAdminRequest = 4,
    [pbr::OriginalName("REG_CLUSTER_ADMIN_REQUEST")] RegClusterAdminRequest = 5,
    [pbr::OriginalName("UNREG_AGENT_ADMIN_REQUEST")] UnregAgentAdminRequest = 6,
    [pbr::OriginalName("UNREG_CLUSTER_ADMIN_REQUEST")] UnregClusterAdminRequest = 7,
    /// <summary>
    /// Auth on agent
    /// </summary>
    [pbr::OriginalName("AUTHENTICATE_AGENT_REQUEST")] AuthenticateAgentRequest = 8,
    /// <summary>
    /// Auth on cluster
    /// </summary>
    [pbr::OriginalName("AUTHENTICATE_REQUEST")] AuthenticateRequest = 9,
    /// <summary>
    /// Auth on infobase
    /// </summary>
    [pbr::OriginalName("ADD_AUTHENTICATION_REQUEST")] AddAuthenticationRequest = 10,
    [pbr::OriginalName("GET_CLUSTERS_REQUEST")] GetClustersRequest = 11,
    [pbr::OriginalName("GET_CLUSTERS_RESPONSE")] GetClustersResponse = 12,
    [pbr::OriginalName("GET_CLUSTER_INFO_REQUEST")] GetClusterInfoRequest = 13,
    [pbr::OriginalName("GET_CLUSTER_INFO_RESPONSE")] GetClusterInfoResponse = 14,
    [pbr::OriginalName("REG_CLUSTER_REQUEST")] RegClusterRequest = 15,
    [pbr::OriginalName("REG_CLUSTER_RESPONSE")] RegClusterResponse = 16,
    [pbr::OriginalName("UNREG_CLUSTER_REQUEST")] UnregClusterRequest = 17,
    [pbr::OriginalName("GET_CLUSTER_MANAGERS_REQUEST")] GetClusterManagersRequest = 18,
    [pbr::OriginalName("GET_CLUSTER_MANAGERS_RESPONSE")] GetClusterManagersResponse = 19,
    [pbr::OriginalName("GET_CLUSTER_MANAGER_INFO_REQUEST")] GetClusterManagerInfoRequest = 20,
    [pbr::OriginalName("GET_CLUSTER_MANAGER_INFO_RESPONSE")] GetClusterManagerInfoResponse = 21,
    [pbr::OriginalName("GET_WORKING_SERVERS_REQUEST")] GetWorkingServersRequest = 22,
    [pbr::OriginalName("GET_WORKING_SERVERS_RESPONSE")] GetWorkingServersResponse = 23,
    [pbr::OriginalName("GET_WORKING_SERVER_INFO_REQUEST")] GetWorkingServerInfoRequest = 24,
    [pbr::OriginalName("GET_WORKING_SERVER_INFO_RESPONSE")] GetWorkingServerInfoResponse = 25,
    [pbr::OriginalName("REG_WORKING_SERVER_REQUEST")] RegWorkingServerRequest = 26,
    [pbr::OriginalName("REG_WORKING_SERVER_RESPONSE")] RegWorkingServerResponse = 27,
    [pbr::OriginalName("UNREG_WORKING_SERVER_REQUEST")] UnregWorkingServerRequest = 28,
    [pbr::OriginalName("GET_WORKING_PROCESSES_REQUEST")] GetWorkingProcessesRequest = 29,
    [pbr::OriginalName("GET_WORKING_PROCESSES_RESPONSE")] GetWorkingProcessesResponse = 30,
    [pbr::OriginalName("GET_WORKING_PROCESS_INFO_REQUEST")] GetWorkingProcessInfoRequest = 31,
    [pbr::OriginalName("GET_WORKING_PROCESS_INFO_RESPONSE")] GetWorkingProcessInfoResponse = 32,
    [pbr::OriginalName("GET_SERVER_WORKING_PROCESSES_REQUEST")] GetServerWorkingProcessesRequest = 33,
    [pbr::OriginalName("GET_SERVER_WORKING_PROCESSES_RESPONSE")] GetServerWorkingProcessesResponse = 34,
    [pbr::OriginalName("GET_CLUSTER_SERVICES_REQUEST")] GetClusterServicesRequest = 35,
    [pbr::OriginalName("GET_CLUSTER_SERVICES_RESPONSE")] GetClusterServicesResponse = 36,
    [pbr::OriginalName("CREATE_INFOBASE_REQUEST")] CreateInfobaseRequest = 37,
    [pbr::OriginalName("CREATE_INFOBASE_RESPONSE")] CreateInfobaseResponse = 38,
    [pbr::OriginalName("UPDATE_INFOBASE_SHORT_REQUEST")] UpdateInfobaseShortRequest = 39,
    [pbr::OriginalName("UPDATE_INFOBASE_REQUEST")] UpdateInfobaseRequest = 40,
    [pbr::OriginalName("DROP_INFOBASE_REQUEST")] DropInfobaseRequest = 41,
    [pbr::OriginalName("GET_INFOBASES_SHORT_REQUEST")] GetInfobasesShortRequest = 42,
    [pbr::OriginalName("GET_INFOBASES_SHORT_RESPONSE")] GetInfobasesShortResponse = 43,
    [pbr::OriginalName("GET_INFOBASES_REQUEST")] GetInfobasesRequest = 44,
    [pbr::OriginalName("GET_INFOBASES_RESPONSE")] GetInfobasesResponse = 45,
    [pbr::OriginalName("GET_INFOBASE_SHORT_INFO_REQUEST")] GetInfobaseShortInfoRequest = 46,
    [pbr::OriginalName("GET_INFOBASE_SHORT_INFO_RESPONSE")] GetInfobaseShortInfoResponse = 47,
    [pbr::OriginalName("GET_INFOBASE_INFO_REQUEST")] GetInfobaseInfoRequest = 48,
    [pbr::OriginalName("GET_INFOBASE_INFO_RESPONSE")] GetInfobaseInfoResponse = 49,
    [pbr::OriginalName("GET_CONNECTIONS_SHORT_REQUEST")] GetConnectionsShortRequest = 50,
    [pbr::OriginalName("GET_CONNECTIONS_SHORT_RESPONSE")] GetConnectionsShortResponse = 51,
    [pbr::OriginalName("GET_INFOBASE_CONNECTIONS_SHORT_REQUEST")] GetInfobaseConnectionsShortRequest = 52,
    [pbr::OriginalName("GET_INFOBASE_CONNECTIONS_SHORT_RESPONSE")] GetInfobaseConnectionsShortResponse = 53,
    [pbr::OriginalName("GET_CONNECTION_INFO_SHORT_REQUEST")] GetConnectionInfoShortRequest = 54,
    [pbr::OriginalName("GET_CONNECTION_INFO_SHORT_RESPONSE")] GetConnectionInfoShortResponse = 55,
    [pbr::OriginalName("GET_INFOBASE_CONNECTIONS_REQUEST")] GetInfobaseConnectionsRequest = 56,
    [pbr::OriginalName("GET_INFOBASE_CONNECTIONS_RESPONSE")] GetInfobaseConnectionsResponse = 57,
    [pbr::OriginalName("DISCONNECT_REQUEST")] DisconnectRequest = 64,
    [pbr::OriginalName("GET_SESSIONS_REQUEST")] GetSessionsRequest = 65,
    [pbr::OriginalName("GET_SESSIONS_RESPONSE")] GetSessionsResponse = 66,
    [pbr::OriginalName("GET_INFOBASE_SESSIONS_REQUEST")] GetInfobaseSessionsRequest = 67,
    [pbr::OriginalName("GET_INFOBASE_SESSIONS_RESPONSE")] GetInfobaseSessionsResponse = 68,
    [pbr::OriginalName("GET_SESSION_INFO_REQUEST")] GetSessionInfoRequest = 69,
    [pbr::OriginalName("GET_SESSION_INFO_RESPONSE")] GetSessionInfoResponse = 70,
    [pbr::OriginalName("TERMINATE_SESSION_REQUEST")] TerminateSessionRequest = 71,
    [pbr::OriginalName("GET_LOCKS_REQUEST")] GetLocksRequest = 72,
    [pbr::OriginalName("GET_LOCKS_RESPONSE")] GetLocksResponse = 73,
    [pbr::OriginalName("GET_INFOBASE_LOCKS_REQUEST")] GetInfobaseLocksRequest = 74,
    [pbr::OriginalName("GET_INFOBASE_LOCKS_RESPONSE")] GetInfobaseLocksResponse = 75,
    [pbr::OriginalName("GET_CONNECTION_LOCKS_REQUEST")] GetConnectionLocksRequest = 76,
    [pbr::OriginalName("GET_CONNECTION_LOCKS_RESPONSE")] GetConnectionLocksResponse = 77,
    [pbr::OriginalName("GET_SESSION_LOCKS_REQUEST")] GetSessionLocksRequest = 78,
    [pbr::OriginalName("GET_SESSION_LOCKS_RESPONSE")] GetSessionLocksResponse = 79,
    [pbr::OriginalName("APPLY_ASSIGNMENT_RULES_REQUEST")] ApplyAssignmentRulesRequest = 81,
    [pbr::OriginalName("REG_ASSIGNMENT_RULE_REQUEST")] RegAssignmentRuleRequest = 82,
    [pbr::OriginalName("REG_ASSIGNMENT_RULE_RESPONSE")] RegAssignmentRuleResponse = 83,
    [pbr::OriginalName("UNREG_ASSIGNMENT_RULE_REQUEST")] UnregAssignmentRuleRequest = 84,
    [pbr::OriginalName("GET_ASSIGNMENT_RULES_REQUEST")] GetAssignmentRulesRequest = 85,
    [pbr::OriginalName("GET_ASSIGNMENT_RULES_RESPONSE")] GetAssignmentRulesResponse = 86,
    [pbr::OriginalName("GET_ASSIGNMENT_RULE_INFO_REQUEST")] GetAssignmentRuleInfoRequest = 87,
    [pbr::OriginalName("GET_ASSIGNMENT_RULE_INFO_RESPONSE")] GetAssignmentRuleInfoResponse = 88,
    [pbr::OriginalName("GET_SECURITY_PROFILES_REQUEST")] GetSecurityProfilesRequest = 89,
    [pbr::OriginalName("GET_SECURITY_PROFILES_RESPONSE")] GetSecurityProfilesResponse = 90,
    [pbr::OriginalName("CREATE_SECURITY_PROFILE_REQUEST")] CreateSecurityProfileRequest = 91,
    [pbr::OriginalName("DROP_SECURITY_PROFILE_REQUEST")] DropSecurityProfileRequest = 92,
    [pbr::OriginalName("GET_VIRTUAL_DIRECTORIES_REQUEST")] GetVirtualDirectoriesRequest = 93,
    [pbr::OriginalName("GET_VIRTUAL_DIRECTORIES_RESPONSE")] GetVirtualDirectoriesResponse = 94,
    [pbr::OriginalName("CREATE_VIRTUAL_DIRECTORY_REQUEST")] CreateVirtualDirectoryRequest = 95,
    [pbr::OriginalName("DROP_VIRTUAL_DIRECTORY_REQUEST")] DropVirtualDirectoryRequest = 96,
    [pbr::OriginalName("GET_COM_CLASSES_REQUEST")] GetComClassesRequest = 97,
    [pbr::OriginalName("GET_COM_CLASSES_RESPONSE")] GetComClassesResponse = 98,
    [pbr::OriginalName("CREATE_COM_CLASS_REQUEST")] CreateComClassRequest = 99,
    [pbr::OriginalName("DROP_COM_CLASS_REQUEST")] DropComClassRequest = 100,
    [pbr::OriginalName("GET_ALLOWED_ADDINS_REQUEST")] GetAllowedAddinsRequest = 101,
    [pbr::OriginalName("GET_ALLOWED_ADDINS_RESPONSE")] GetAllowedAddinsResponse = 102,
    [pbr::OriginalName("CREATE_ALLOWED_ADDIN_REQUEST")] CreateAllowedAddinRequest = 103,
    [pbr::OriginalName("DROP_ALLOWED_ADDIN_REQUEST")] DropAllowedAddinRequest = 104,
    [pbr::OriginalName("GET_EXTERNAL_MODULES_REQUEST")] GetExternalModulesRequest = 105,
    [pbr::OriginalName("GET_EXTERNAL_MODULES_RESPONSE")] GetExternalModulesResponse = 106,
    [pbr::OriginalName("CREATE_EXTERNAL_MODULE_REQUEST")] CreateExternalModuleRequest = 107,
    [pbr::OriginalName("DROP_EXTERNAL_MODULE_REQUEST")] DropExternalModuleRequest = 108,
    [pbr::OriginalName("GET_ALLOWED_APPLICATIONS_REQUEST")] GetAllowedApplicationsRequest = 109,
    [pbr::OriginalName("GET_ALLOWED_APPLICATIONS_RESPONSE")] GetAllowedApplicationsResponse = 110,
    [pbr::OriginalName("CREATE_ALLOWED_APPLICATION_REQUEST")] CreateAllowedApplicationRequest = 111,
    [pbr::OriginalName("DROP_ALLOWED_APPLICATION_REQUEST")] DropAllowedApplicationRequest = 112,
    [pbr::OriginalName("GET_INTERNET_RESOURCES_REQUEST")] GetInternetResourcesRequest = 113,
    [pbr::OriginalName("GET_INTERNET_RESOURCES_RESPONSE")] GetInternetResourcesResponse = 114,
    [pbr::OriginalName("CREATE_INTERNET_RESOURCE_REQUEST")] CreateInternetResourceRequest = 115,
    [pbr::OriginalName("DROP_INTERNET_RESOURCE_REQUEST")] DropInternetResourceRequest = 116,
    [pbr::OriginalName("INTERRUPT_SESSION_CURRENT_SERVER_CALL_REQUEST")] InterruptSessionCurrentServerCallRequest = 117,
    [pbr::OriginalName("GET_RESOURCE_COUNTERS_REQUEST")] GetResourceCountersRequest = 118,
    [pbr::OriginalName("GET_RESOURCE_COUNTERS_RESPONSE")] GetResourceCountersResponse = 119,
    [pbr::OriginalName("GET_RESOURCE_COUNTER_INFO_REQUEST")] GetResourceCounterInfoRequest = 120,
    [pbr::OriginalName("GET_RESOURCE_COUNTER_INFO_RESPONSE")] GetResourceCounterInfoResponse = 121,
    [pbr::OriginalName("REG_RESOURCE_COUNTER_REQUEST")] RegResourceCounterRequest = 122,
    [pbr::OriginalName("UNREG_RESOURCE_COUNTER_REQUEST")] UnregResourceCounterRequest = 123,
    [pbr::OriginalName("GET_RESOURCE_LIMITS_REQUEST")] GetResourceLimitsRequest = 124,
    [pbr::OriginalName("GET_RESOURCE_LIMITS_RESPONSE")] GetResourceLimitsResponse = 125,
    [pbr::OriginalName("GET_RESOURCE_LIMIT_INFO_REQUEST")] GetResourceLimitInfoRequest = 126,
    [pbr::OriginalName("GET_RESOURCE_LIMIT_INFO_RESPONSE")] GetResourceLimitInfoResponse = 127,
    [pbr::OriginalName("REG_RESOURCE_LIMIT_REQUEST")] RegResourceLimitRequest = 128,
    [pbr::OriginalName("UNREG_RESOURCE_LIMIT_REQUEST")] UnregResourceLimitRequest = 129,
    [pbr::OriginalName("GET_COUNTER_VALUES_REQUEST")] GetCounterValuesRequest = 130,
    [pbr::OriginalName("GET_COUNTER_VALUES_RESPONSE")] GetCounterValuesResponse = 131,
    [pbr::OriginalName("CLEAR_COUNTER_VALUE_REQUEST")] ClearCounterValueRequest = 132,
    [pbr::OriginalName("GET_COUNTER_ACCUMULATED_VALUES_REQUEST")] GetCounterAccumulatedValuesRequest = 133,
    [pbr::OriginalName("GET_COUNTER_ACCUMULATED_VALUES_RESPONSE")] GetCounterAccumulatedValuesResponse = 134,
    [pbr::OriginalName("GET_AGENT_VERSION_REQUEST")] GetAgentVersionRequest = 135,
    [pbr::OriginalName("GET_AGENT_VERSION_RESPONSE")] GetAgentVersionResponse = 136,
  }

  #endregion

}

#endregion Designer generated code