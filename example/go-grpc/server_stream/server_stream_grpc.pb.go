// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// StreamServerClient is the client API for StreamServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamServerClient interface {
	// 服务端流式rpc，在响应数据前添加stream
	ListValue(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (StreamServer_ListValueClient, error)
}

type streamServerClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServerClient(cc grpc.ClientConnInterface) StreamServerClient {
	return &streamServerClient{cc}
}

func (c *streamServerClient) ListValue(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (StreamServer_ListValueClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamServer_ServiceDesc.Streams[0], "/StreamServer/ListValue", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServerListValueClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamServer_ListValueClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServerListValueClient struct {
	grpc.ClientStream
}

func (x *streamServerListValueClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServerServer is the server API for StreamServer service.
// All implementations must embed UnimplementedStreamServerServer
// for forward compatibility
type StreamServerServer interface {
	// 服务端流式rpc，在响应数据前添加stream
	ListValue(*SimpleRequest, StreamServer_ListValueServer) error
	mustEmbedUnimplementedStreamServerServer()
}

// UnimplementedStreamServerServer must be embedded to have forward compatible implementations.
type UnimplementedStreamServerServer struct {
}

func (UnimplementedStreamServerServer) ListValue(*SimpleRequest, StreamServer_ListValueServer) error {
	return status.Errorf(codes.Unimplemented, "method ListValue not implemented")
}
func (UnimplementedStreamServerServer) mustEmbedUnimplementedStreamServerServer() {}

// UnsafeStreamServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServerServer will
// result in compilation errors.
type UnsafeStreamServerServer interface {
	mustEmbedUnimplementedStreamServerServer()
}

func RegisterStreamServerServer(s grpc.ServiceRegistrar, srv StreamServerServer) {
	s.RegisterService(&StreamServer_ServiceDesc, srv)
}

func _StreamServer_ListValue_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServerServer).ListValue(m, &streamServerListValueServer{stream})
}

type StreamServer_ListValueServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type streamServerListValueServer struct {
	grpc.ServerStream
}

func (x *streamServerListValueServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// StreamServer_ServiceDesc is the grpc.ServiceDesc for StreamServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StreamServer",
	HandlerType: (*StreamServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListValue",
			Handler:       _StreamServer_ListValue_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server_stream.proto",
}