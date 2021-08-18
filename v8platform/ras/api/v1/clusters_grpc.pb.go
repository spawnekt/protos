// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	serialize "github.com/v8platform/protos/v8platform/ras/serialize"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClustersServiceClient is the client API for ClustersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClustersServiceClient interface {
	Clusters(ctx context.Context, in *GetClustersRequest, opts ...grpc.CallOption) (*GetClustersResponse, error)
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*serialize.ClusterInfo, error)
	AddCluster(ctx context.Context, in *AddClusterRequest, opts ...grpc.CallOption) (*AddClusterResponse, error)
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clustersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClustersServiceClient(cc grpc.ClientConnInterface) ClustersServiceClient {
	return &clustersServiceClient{cc}
}

func (c *clustersServiceClient) Clusters(ctx context.Context, in *GetClustersRequest, opts ...grpc.CallOption) (*GetClustersResponse, error) {
	out := new(GetClustersResponse)
	err := c.cc.Invoke(ctx, "/ras.api.v1.ClustersService/Clusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersServiceClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*serialize.ClusterInfo, error) {
	out := new(serialize.ClusterInfo)
	err := c.cc.Invoke(ctx, "/ras.api.v1.ClustersService/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersServiceClient) AddCluster(ctx context.Context, in *AddClusterRequest, opts ...grpc.CallOption) (*AddClusterResponse, error) {
	out := new(AddClusterResponse)
	err := c.cc.Invoke(ctx, "/ras.api.v1.ClustersService/AddCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersServiceClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ras.api.v1.ClustersService/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClustersServiceServer is the server API for ClustersService service.
// All implementations should embed UnimplementedClustersServiceServer
// for forward compatibility
type ClustersServiceServer interface {
	Clusters(context.Context, *GetClustersRequest) (*GetClustersResponse, error)
	GetCluster(context.Context, *GetClusterRequest) (*serialize.ClusterInfo, error)
	AddCluster(context.Context, *AddClusterRequest) (*AddClusterResponse, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*emptypb.Empty, error)
}

// UnimplementedClustersServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClustersServiceServer struct {
}

func (UnimplementedClustersServiceServer) Clusters(context.Context, *GetClustersRequest) (*GetClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clusters not implemented")
}
func (UnimplementedClustersServiceServer) GetCluster(context.Context, *GetClusterRequest) (*serialize.ClusterInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (UnimplementedClustersServiceServer) AddCluster(context.Context, *AddClusterRequest) (*AddClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCluster not implemented")
}
func (UnimplementedClustersServiceServer) DeleteCluster(context.Context, *DeleteClusterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}

// UnsafeClustersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClustersServiceServer will
// result in compilation errors.
type UnsafeClustersServiceServer interface {
	mustEmbedUnimplementedClustersServiceServer()
}

func RegisterClustersServiceServer(s grpc.ServiceRegistrar, srv ClustersServiceServer) {
	s.RegisterService(&ClustersService_ServiceDesc, srv)
}

func _ClustersService_Clusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServiceServer).Clusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.ClustersService/Clusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServiceServer).Clusters(ctx, req.(*GetClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClustersService_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServiceServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.ClustersService/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServiceServer).GetCluster(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClustersService_AddCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServiceServer).AddCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.ClustersService/AddCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServiceServer).AddCluster(ctx, req.(*AddClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClustersService_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServiceServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.ClustersService/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServiceServer).DeleteCluster(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClustersService_ServiceDesc is the grpc.ServiceDesc for ClustersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClustersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ras.api.v1.ClustersService",
	HandlerType: (*ClustersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Clusters",
			Handler:    _ClustersService_Clusters_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _ClustersService_GetCluster_Handler,
		},
		{
			MethodName: "AddCluster",
			Handler:    _ClustersService_AddCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _ClustersService_DeleteCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v8platform/ras/api/v1/clusters.proto",
}
