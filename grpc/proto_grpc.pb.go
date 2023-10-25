// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: grpc/proto.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChittyChat_SendClientMessage_FullMethodName = "/ChittyChat/SendClientMessage"
	ChittyChat_ConnectClient_FullMethodName     = "/ChittyChat/ConnectClient"
	ChittyChat_DisconnectClient_FullMethodName  = "/ChittyChat/DisconnectClient"
)

// ChittyChatClient is the client API for ChittyChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittyChatClient interface {
	SendClientMessage(ctx context.Context, in *ClientMessage, opts ...grpc.CallOption) (*Empty, error)
	ConnectClient(ctx context.Context, opts ...grpc.CallOption) (ChittyChat_ConnectClientClient, error)
	DisconnectClient(ctx context.Context, in *Disconnection, opts ...grpc.CallOption) (*Empty, error)
}

type chittyChatClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatClient(cc grpc.ClientConnInterface) ChittyChatClient {
	return &chittyChatClient{cc}
}

func (c *chittyChatClient) SendClientMessage(ctx context.Context, in *ClientMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChittyChat_SendClientMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatClient) ConnectClient(ctx context.Context, opts ...grpc.CallOption) (ChittyChat_ConnectClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChat_ServiceDesc.Streams[0], ChittyChat_ConnectClient_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chittyChatConnectClientClient{stream}
	return x, nil
}

type ChittyChat_ConnectClientClient interface {
	Send(*ServerMessage) error
	Recv() (*ServerMessage, error)
	grpc.ClientStream
}

type chittyChatConnectClientClient struct {
	grpc.ClientStream
}

func (x *chittyChatConnectClientClient) Send(m *ServerMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chittyChatConnectClientClient) Recv() (*ServerMessage, error) {
	m := new(ServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chittyChatClient) DisconnectClient(ctx context.Context, in *Disconnection, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ChittyChat_DisconnectClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChittyChatServer is the server API for ChittyChat service.
// All implementations must embed UnimplementedChittyChatServer
// for forward compatibility
type ChittyChatServer interface {
	SendClientMessage(context.Context, *ClientMessage) (*Empty, error)
	ConnectClient(ChittyChat_ConnectClientServer) error
	DisconnectClient(context.Context, *Disconnection) (*Empty, error)
	mustEmbedUnimplementedChittyChatServer()
}

// UnimplementedChittyChatServer must be embedded to have forward compatible implementations.
type UnimplementedChittyChatServer struct {
}

func (UnimplementedChittyChatServer) SendClientMessage(context.Context, *ClientMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendClientMessage not implemented")
}
func (UnimplementedChittyChatServer) ConnectClient(ChittyChat_ConnectClientServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectClient not implemented")
}
func (UnimplementedChittyChatServer) DisconnectClient(context.Context, *Disconnection) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisconnectClient not implemented")
}
func (UnimplementedChittyChatServer) mustEmbedUnimplementedChittyChatServer() {}

// UnsafeChittyChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittyChatServer will
// result in compilation errors.
type UnsafeChittyChatServer interface {
	mustEmbedUnimplementedChittyChatServer()
}

func RegisterChittyChatServer(s grpc.ServiceRegistrar, srv ChittyChatServer) {
	s.RegisterService(&ChittyChat_ServiceDesc, srv)
}

func _ChittyChat_SendClientMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).SendClientMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittyChat_SendClientMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).SendClientMessage(ctx, req.(*ClientMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChat_ConnectClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChittyChatServer).ConnectClient(&chittyChatConnectClientServer{stream})
}

type ChittyChat_ConnectClientServer interface {
	Send(*ServerMessage) error
	Recv() (*ServerMessage, error)
	grpc.ServerStream
}

type chittyChatConnectClientServer struct {
	grpc.ServerStream
}

func (x *chittyChatConnectClientServer) Send(m *ServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chittyChatConnectClientServer) Recv() (*ServerMessage, error) {
	m := new(ServerMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChittyChat_DisconnectClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Disconnection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServer).DisconnectClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittyChat_DisconnectClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServer).DisconnectClient(ctx, req.(*Disconnection))
	}
	return interceptor(ctx, in, info, handler)
}

// ChittyChat_ServiceDesc is the grpc.ServiceDesc for ChittyChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittyChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChittyChat",
	HandlerType: (*ChittyChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendClientMessage",
			Handler:    _ChittyChat_SendClientMessage_Handler,
		},
		{
			MethodName: "DisconnectClient",
			Handler:    _ChittyChat_DisconnectClient_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectClient",
			Handler:       _ChittyChat_ConnectClient_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/proto.proto",
}
