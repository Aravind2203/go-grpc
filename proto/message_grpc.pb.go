// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: proto/message.proto

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

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	SayBye(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (GreetService_SayByeClient, error)
	SaystreamHello(ctx context.Context, opts ...grpc.CallOption) (GreetService_SaystreamHelloClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/message.GreetService/sayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) SayBye(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (GreetService_SayByeClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/message.GreetService/sayBye", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayByeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_SayByeClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceSayByeClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayByeClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) SaystreamHello(ctx context.Context, opts ...grpc.CallOption) (GreetService_SaystreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/message.GreetService/saystreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSaystreamHelloClient{stream}
	return x, nil
}

type GreetService_SaystreamHelloClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceSaystreamHelloClient struct {
	grpc.ClientStream
}

func (x *greetServiceSaystreamHelloClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSaystreamHelloClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	SayHello(context.Context, *NoParam) (*HelloResponse, error)
	SayBye(*NoParam, GreetService_SayByeServer) error
	SaystreamHello(GreetService_SaystreamHelloServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) SayHello(context.Context, *NoParam) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetServiceServer) SayBye(*NoParam, GreetService_SayByeServer) error {
	return status.Errorf(codes.Unimplemented, "method SayBye not implemented")
}
func (UnimplementedGreetServiceServer) SaystreamHello(GreetService_SaystreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SaystreamHello not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.GreetService/sayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).SayHello(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_SayBye_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoParam)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).SayBye(m, &greetServiceSayByeServer{stream})
}

type GreetService_SayByeServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greetServiceSayByeServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayByeServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_SaystreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SaystreamHello(&greetServiceSaystreamHelloServer{stream})
}

type GreetService_SaystreamHelloServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetServiceSaystreamHelloServer struct {
	grpc.ServerStream
}

func (x *greetServiceSaystreamHelloServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSaystreamHelloServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sayHello",
			Handler:    _GreetService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sayBye",
			Handler:       _GreetService_SayBye_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "saystreamHello",
			Handler:       _GreetService_SaystreamHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/message.proto",
}