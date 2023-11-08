// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: gRPC/proto.proto

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
	MutualExclusion_ConnectNode_FullMethodName = "/MutualExclusion/ConnectNode"
)

// MutualExclusionClient is the client API for MutualExclusion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MutualExclusionClient interface {
	ConnectNode(ctx context.Context, in *NodeConnection, opts ...grpc.CallOption) (MutualExclusion_ConnectNodeClient, error)
}

type mutualExclusionClient struct {
	cc grpc.ClientConnInterface
}

func NewMutualExclusionClient(cc grpc.ClientConnInterface) MutualExclusionClient {
	return &mutualExclusionClient{cc}
}

func (c *mutualExclusionClient) ConnectNode(ctx context.Context, in *NodeConnection, opts ...grpc.CallOption) (MutualExclusion_ConnectNodeClient, error) {
	stream, err := c.cc.NewStream(ctx, &MutualExclusion_ServiceDesc.Streams[0], MutualExclusion_ConnectNode_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &mutualExclusionConnectNodeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MutualExclusion_ConnectNodeClient interface {
	Recv() (*ServerMessage, error)
	grpc.ClientStream
}

type mutualExclusionConnectNodeClient struct {
	grpc.ClientStream
}

func (x *mutualExclusionConnectNodeClient) Recv() (*ServerMessage, error) {
	m := new(ServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MutualExclusionServer is the server API for MutualExclusion service.
// All implementations must embed UnimplementedMutualExclusionServer
// for forward compatibility
type MutualExclusionServer interface {
	ConnectNode(*NodeConnection, MutualExclusion_ConnectNodeServer) error
	mustEmbedUnimplementedMutualExclusionServer()
}

// UnimplementedMutualExclusionServer must be embedded to have forward compatible implementations.
type UnimplementedMutualExclusionServer struct {
}

func (UnimplementedMutualExclusionServer) ConnectNode(*NodeConnection, MutualExclusion_ConnectNodeServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectNode not implemented")
}
func (UnimplementedMutualExclusionServer) mustEmbedUnimplementedMutualExclusionServer() {}

// UnsafeMutualExclusionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MutualExclusionServer will
// result in compilation errors.
type UnsafeMutualExclusionServer interface {
	mustEmbedUnimplementedMutualExclusionServer()
}

func RegisterMutualExclusionServer(s grpc.ServiceRegistrar, srv MutualExclusionServer) {
	s.RegisterService(&MutualExclusion_ServiceDesc, srv)
}

func _MutualExclusion_ConnectNode_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NodeConnection)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MutualExclusionServer).ConnectNode(m, &mutualExclusionConnectNodeServer{stream})
}

type MutualExclusion_ConnectNodeServer interface {
	Send(*ServerMessage) error
	grpc.ServerStream
}

type mutualExclusionConnectNodeServer struct {
	grpc.ServerStream
}

func (x *mutualExclusionConnectNodeServer) Send(m *ServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

// MutualExclusion_ServiceDesc is the grpc.ServiceDesc for MutualExclusion service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MutualExclusion_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MutualExclusion",
	HandlerType: (*MutualExclusionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectNode",
			Handler:       _MutualExclusion_ConnectNode_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gRPC/proto.proto",
}
