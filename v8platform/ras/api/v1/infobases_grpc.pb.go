// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package infobases

import (
	context "context"
	serialize "github.com/v8platform/protos/v8platform/ras/serialize"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InfobasesServiceClient is the client API for InfobasesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InfobasesServiceClient interface {
	Infobases(ctx context.Context, in *InfobasesRequest, opts ...grpc.CallOption) (*InfobasesResponse, error)
	LookupInfobase(ctx context.Context, in *LookupInfobaseRequest, opts ...grpc.CallOption) (*LookupInfobaseResponse, error)
	GetInfobase(ctx context.Context, in *GetInfobaseRequest, opts ...grpc.CallOption) (*serialize.InfobaseInfo, error)
}

type infobasesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInfobasesServiceClient(cc grpc.ClientConnInterface) InfobasesServiceClient {
	return &infobasesServiceClient{cc}
}

func (c *infobasesServiceClient) Infobases(ctx context.Context, in *InfobasesRequest, opts ...grpc.CallOption) (*InfobasesResponse, error) {
	out := new(InfobasesResponse)
	err := c.cc.Invoke(ctx, "/ras.api.v1.InfobasesService/Infobases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infobasesServiceClient) LookupInfobase(ctx context.Context, in *LookupInfobaseRequest, opts ...grpc.CallOption) (*LookupInfobaseResponse, error) {
	out := new(LookupInfobaseResponse)
	err := c.cc.Invoke(ctx, "/ras.api.v1.InfobasesService/LookupInfobase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *infobasesServiceClient) GetInfobase(ctx context.Context, in *GetInfobaseRequest, opts ...grpc.CallOption) (*serialize.InfobaseInfo, error) {
	out := new(serialize.InfobaseInfo)
	err := c.cc.Invoke(ctx, "/ras.api.v1.InfobasesService/GetInfobase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InfobasesServiceServer is the server API for InfobasesService service.
// All implementations should embed UnimplementedInfobasesServiceServer
// for forward compatibility
type InfobasesServiceServer interface {
	Infobases(context.Context, *InfobasesRequest) (*InfobasesResponse, error)
	LookupInfobase(context.Context, *LookupInfobaseRequest) (*LookupInfobaseResponse, error)
	GetInfobase(context.Context, *GetInfobaseRequest) (*serialize.InfobaseInfo, error)
}

// UnimplementedInfobasesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInfobasesServiceServer struct {
}

func (UnimplementedInfobasesServiceServer) Infobases(context.Context, *InfobasesRequest) (*InfobasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Infobases not implemented")
}
func (UnimplementedInfobasesServiceServer) LookupInfobase(context.Context, *LookupInfobaseRequest) (*LookupInfobaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupInfobase not implemented")
}
func (UnimplementedInfobasesServiceServer) GetInfobase(context.Context, *GetInfobaseRequest) (*serialize.InfobaseInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfobase not implemented")
}

// UnsafeInfobasesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InfobasesServiceServer will
// result in compilation errors.
type UnsafeInfobasesServiceServer interface {
	mustEmbedUnimplementedInfobasesServiceServer()
}

func RegisterInfobasesServiceServer(s grpc.ServiceRegistrar, srv InfobasesServiceServer) {
	s.RegisterService(&InfobasesService_ServiceDesc, srv)
}

func _InfobasesService_Infobases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfobasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfobasesServiceServer).Infobases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.InfobasesService/Infobases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfobasesServiceServer).Infobases(ctx, req.(*InfobasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfobasesService_LookupInfobase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupInfobaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfobasesServiceServer).LookupInfobase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.InfobasesService/LookupInfobase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfobasesServiceServer).LookupInfobase(ctx, req.(*LookupInfobaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InfobasesService_GetInfobase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInfobaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InfobasesServiceServer).GetInfobase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.InfobasesService/GetInfobase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InfobasesServiceServer).GetInfobase(ctx, req.(*GetInfobaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InfobasesService_ServiceDesc is the grpc.ServiceDesc for InfobasesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InfobasesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ras.api.v1.InfobasesService",
	HandlerType: (*InfobasesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Infobases",
			Handler:    _InfobasesService_Infobases_Handler,
		},
		{
			MethodName: "LookupInfobase",
			Handler:    _InfobasesService_LookupInfobase_Handler,
		},
		{
			MethodName: "GetInfobase",
			Handler:    _InfobasesService_GetInfobase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v8platform/ras/api/v1/infobases.proto",
}
