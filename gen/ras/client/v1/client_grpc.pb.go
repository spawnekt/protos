// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package clientv1

import (
	context "context"
	protocol "github.com/v8platform/protos/gen/ras/protocol"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	// Инициализация клиента RAS
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*StatusInfo, error)
	// Возвращает состояние клиента
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusInfo, error)
	// Открытие новой точки работы с RAS
	NewEndpoint(ctx context.Context, in *NewEndpointRequest, opts ...grpc.CallOption) (*protocol.Endpoint, error)
	// Закрытие точки работы с RAS
	CloseEndpoint(ctx context.Context, in *protocol.Endpoint, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Поток запросов
	// EndpointRequest.value должно реализовавать интерфейс сообщения точки работы
	// * MessageType()
	// * Format()
	Requests(ctx context.Context, opts ...grpc.CallOption) (ClientService_RequestsClient, error)
	// Основное метод раборы с запросами к RAS
	// EndpointRequest должно реализовавать интерфейс сообщения точки работы
	// * MessageType()
	// * Format()
	Request(ctx context.Context, in *EndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error)
	// Низкоуровневые методы клиента
	// Для ручного использования
	Negotiate(ctx context.Context, in *protocol.NegotiateMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Connect(ctx context.Context, in *protocol.ConnectMessage, opts ...grpc.CallOption) (*protocol.ConnectMessageAck, error)
	Disconnect(ctx context.Context, in *protocol.DisconnectMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EndpointOpen(ctx context.Context, in *protocol.EndpointOpen, opts ...grpc.CallOption) (*protocol.EndpointOpenAck, error)
	EndpointClose(ctx context.Context, in *protocol.EndpointClose, opts ...grpc.CallOption) (*emptypb.Empty, error)
	EndpointMessage(ctx context.Context, in *protocol.EndpointMessage, opts ...grpc.CallOption) (*protocol.EndpointMessage, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*StatusInfo, error) {
	out := new(StatusInfo)
	err := c.cc.Invoke(ctx, "/ras.client.v1.ClientService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatusInfo, error) {
	out := new(StatusInfo)
	err := c.cc.Invoke(ctx, "/ras.client.v1.ClientService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) NewEndpoint(ctx context.Context, in *NewEndpointRequest, opts ...grpc.CallOption) (*protocol.Endpoint, error) {
	out := new(protocol.Endpoint)
	err := c.cc.Invoke(ctx, "/ras.client.v1.ClientService/NewEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) CloseEndpoint(ctx context.Context, in *protocol.Endpoint, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ras.client.v1.ClientService/CloseEndpoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Requests(ctx context.Context, opts ...grpc.CallOption) (ClientService_RequestsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientService_ServiceDesc.Streams[0], "/ras.client.v1.ClientService/Requests", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceRequestsClient{stream}
	return x, nil
}

type ClientService_RequestsClient interface {
	Send(*EndpointRequest) error
	Recv() (*EndpointResponse, error)
	grpc.ClientStream
}

type clientServiceRequestsClient struct {
	grpc.ClientStream
}

func (x *clientServiceRequestsClient) Send(m *EndpointRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientServiceRequestsClient) Recv() (*EndpointResponse, error) {
	m := new(EndpointResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) Request(ctx context.Context, in *EndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error) {
	out := new(EndpointResponse)
	err := c.cc.Invoke(ctx, "/ras.client.v1.ClientService/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Negotiate(ctx context.Context, in *protocol.NegotiateMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ras.client.v1.ClientService/Negotiate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Connect(ctx context.Context, in *protocol.ConnectMessage, opts ...grpc.CallOption) (*protocol.ConnectMessageAck, error) {
	out := new(protocol.ConnectMessageAck)
	err := c.cc.Invoke(ctx, "/ras.client.v1.ClientService/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Disconnect(ctx context.Context, in *protocol.DisconnectMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ras.client.v1.ClientService/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) EndpointOpen(ctx context.Context, in *protocol.EndpointOpen, opts ...grpc.CallOption) (*protocol.EndpointOpenAck, error) {
	out := new(protocol.EndpointOpenAck)
	err := c.cc.Invoke(ctx, "/ras.client.v1.ClientService/EndpointOpen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) EndpointClose(ctx context.Context, in *protocol.EndpointClose, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ras.client.v1.ClientService/EndpointClose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) EndpointMessage(ctx context.Context, in *protocol.EndpointMessage, opts ...grpc.CallOption) (*protocol.EndpointMessage, error) {
	out := new(protocol.EndpointMessage)
	err := c.cc.Invoke(ctx, "/ras.client.v1.ClientService/EndpointMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations should embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	// Инициализация клиента RAS
	Init(context.Context, *InitRequest) (*StatusInfo, error)
	// Возвращает состояние клиента
	Status(context.Context, *emptypb.Empty) (*StatusInfo, error)
	// Открытие новой точки работы с RAS
	NewEndpoint(context.Context, *NewEndpointRequest) (*protocol.Endpoint, error)
	// Закрытие точки работы с RAS
	CloseEndpoint(context.Context, *protocol.Endpoint) (*emptypb.Empty, error)
	// Поток запросов
	// EndpointRequest.value должно реализовавать интерфейс сообщения точки работы
	// * MessageType()
	// * Format()
	Requests(ClientService_RequestsServer) error
	// Основное метод раборы с запросами к RAS
	// EndpointRequest должно реализовавать интерфейс сообщения точки работы
	// * MessageType()
	// * Format()
	Request(context.Context, *EndpointRequest) (*EndpointResponse, error)
	// Низкоуровневые методы клиента
	// Для ручного использования
	Negotiate(context.Context, *protocol.NegotiateMessage) (*emptypb.Empty, error)
	Connect(context.Context, *protocol.ConnectMessage) (*protocol.ConnectMessageAck, error)
	Disconnect(context.Context, *protocol.DisconnectMessage) (*emptypb.Empty, error)
	EndpointOpen(context.Context, *protocol.EndpointOpen) (*protocol.EndpointOpenAck, error)
	EndpointClose(context.Context, *protocol.EndpointClose) (*emptypb.Empty, error)
	EndpointMessage(context.Context, *protocol.EndpointMessage) (*protocol.EndpointMessage, error)
}

// UnimplementedClientServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) Init(context.Context, *InitRequest) (*StatusInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedClientServiceServer) Status(context.Context, *emptypb.Empty) (*StatusInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedClientServiceServer) NewEndpoint(context.Context, *NewEndpointRequest) (*protocol.Endpoint, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewEndpoint not implemented")
}
func (UnimplementedClientServiceServer) CloseEndpoint(context.Context, *protocol.Endpoint) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseEndpoint not implemented")
}
func (UnimplementedClientServiceServer) Requests(ClientService_RequestsServer) error {
	return status.Errorf(codes.Unimplemented, "method Requests not implemented")
}
func (UnimplementedClientServiceServer) Request(context.Context, *EndpointRequest) (*EndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (UnimplementedClientServiceServer) Negotiate(context.Context, *protocol.NegotiateMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Negotiate not implemented")
}
func (UnimplementedClientServiceServer) Connect(context.Context, *protocol.ConnectMessage) (*protocol.ConnectMessageAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedClientServiceServer) Disconnect(context.Context, *protocol.DisconnectMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedClientServiceServer) EndpointOpen(context.Context, *protocol.EndpointOpen) (*protocol.EndpointOpenAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndpointOpen not implemented")
}
func (UnimplementedClientServiceServer) EndpointClose(context.Context, *protocol.EndpointClose) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndpointClose not implemented")
}
func (UnimplementedClientServiceServer) EndpointMessage(context.Context, *protocol.EndpointMessage) (*protocol.EndpointMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndpointMessage not implemented")
}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.client.v1.ClientService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.client.v1.ClientService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_NewEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).NewEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.client.v1.ClientService/NewEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).NewEndpoint(ctx, req.(*NewEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_CloseEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol.Endpoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).CloseEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.client.v1.ClientService/CloseEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).CloseEndpoint(ctx, req.(*protocol.Endpoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Requests_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientServiceServer).Requests(&clientServiceRequestsServer{stream})
}

type ClientService_RequestsServer interface {
	Send(*EndpointResponse) error
	Recv() (*EndpointRequest, error)
	grpc.ServerStream
}

type clientServiceRequestsServer struct {
	grpc.ServerStream
}

func (x *clientServiceRequestsServer) Send(m *EndpointResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientServiceRequestsServer) Recv() (*EndpointRequest, error) {
	m := new(EndpointRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClientService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.client.v1.ClientService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Request(ctx, req.(*EndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Negotiate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol.NegotiateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Negotiate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.client.v1.ClientService/Negotiate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Negotiate(ctx, req.(*protocol.NegotiateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol.ConnectMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.client.v1.ClientService/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Connect(ctx, req.(*protocol.ConnectMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol.DisconnectMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.client.v1.ClientService/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).Disconnect(ctx, req.(*protocol.DisconnectMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_EndpointOpen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol.EndpointOpen)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).EndpointOpen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.client.v1.ClientService/EndpointOpen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).EndpointOpen(ctx, req.(*protocol.EndpointOpen))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_EndpointClose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol.EndpointClose)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).EndpointClose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.client.v1.ClientService/EndpointClose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).EndpointClose(ctx, req.(*protocol.EndpointClose))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_EndpointMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protocol.EndpointMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).EndpointMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.client.v1.ClientService/EndpointMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).EndpointMessage(ctx, req.(*protocol.EndpointMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ras.client.v1.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _ClientService_Init_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _ClientService_Status_Handler,
		},
		{
			MethodName: "NewEndpoint",
			Handler:    _ClientService_NewEndpoint_Handler,
		},
		{
			MethodName: "CloseEndpoint",
			Handler:    _ClientService_CloseEndpoint_Handler,
		},
		{
			MethodName: "Request",
			Handler:    _ClientService_Request_Handler,
		},
		{
			MethodName: "Negotiate",
			Handler:    _ClientService_Negotiate_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _ClientService_Connect_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _ClientService_Disconnect_Handler,
		},
		{
			MethodName: "EndpointOpen",
			Handler:    _ClientService_EndpointOpen_Handler,
		},
		{
			MethodName: "EndpointClose",
			Handler:    _ClientService_EndpointClose_Handler,
		},
		{
			MethodName: "EndpointMessage",
			Handler:    _ClientService_EndpointMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Requests",
			Handler:       _ClientService_Requests_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ras/client/v1/client.proto",
}