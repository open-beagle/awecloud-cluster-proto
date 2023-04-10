// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: cluster_registry.proto

package awecloud_cluster_proto

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

// ClusterRegistryClient is the client API for ClusterRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterRegistryClient interface {
	// Login 登录服务.
	//
	// 输入ID与Secret，匹配后返回集群链接信息与密钥.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Listen .
	//
	// 持续侦听，直至获取需要注册的服务.
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (ClusterRegistry_ListenClient, error)
	// PublishService .
	//
	// 客户端主动注册服务
	PublishService(ctx context.Context, in *PublishServiceRequest, opts ...grpc.CallOption) (*RegisterService, error)
}

type clusterRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterRegistryClient(cc grpc.ClientConnInterface) ClusterRegistryClient {
	return &clusterRegistryClient{cc}
}

func (c *clusterRegistryClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/awecloud.ClusterRegistry/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterRegistryClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (ClusterRegistry_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClusterRegistry_ServiceDesc.Streams[0], "/awecloud.ClusterRegistry/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterRegistryListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClusterRegistry_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type clusterRegistryListenClient struct {
	grpc.ClientStream
}

func (x *clusterRegistryListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterRegistryClient) PublishService(ctx context.Context, in *PublishServiceRequest, opts ...grpc.CallOption) (*RegisterService, error) {
	out := new(RegisterService)
	err := c.cc.Invoke(ctx, "/awecloud.ClusterRegistry/PublishService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterRegistryServer is the server API for ClusterRegistry service.
// All implementations must embed UnimplementedClusterRegistryServer
// for forward compatibility
type ClusterRegistryServer interface {
	// Login 登录服务.
	//
	// 输入ID与Secret，匹配后返回集群链接信息与密钥.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Listen .
	//
	// 持续侦听，直至获取需要注册的服务.
	Listen(*ListenRequest, ClusterRegistry_ListenServer) error
	// PublishService .
	//
	// 客户端主动注册服务
	PublishService(context.Context, *PublishServiceRequest) (*RegisterService, error)
	mustEmbedUnimplementedClusterRegistryServer()
}

// UnimplementedClusterRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedClusterRegistryServer struct {
}

func (UnimplementedClusterRegistryServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedClusterRegistryServer) Listen(*ListenRequest, ClusterRegistry_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedClusterRegistryServer) PublishService(context.Context, *PublishServiceRequest) (*RegisterService, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishService not implemented")
}
func (UnimplementedClusterRegistryServer) mustEmbedUnimplementedClusterRegistryServer() {}

// UnsafeClusterRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterRegistryServer will
// result in compilation errors.
type UnsafeClusterRegistryServer interface {
	mustEmbedUnimplementedClusterRegistryServer()
}

func RegisterClusterRegistryServer(s grpc.ServiceRegistrar, srv ClusterRegistryServer) {
	s.RegisterService(&ClusterRegistry_ServiceDesc, srv)
}

func _ClusterRegistry_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterRegistryServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/awecloud.ClusterRegistry/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterRegistryServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterRegistry_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClusterRegistryServer).Listen(m, &clusterRegistryListenServer{stream})
}

type ClusterRegistry_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type clusterRegistryListenServer struct {
	grpc.ServerStream
}

func (x *clusterRegistryListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ClusterRegistry_PublishService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterRegistryServer).PublishService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/awecloud.ClusterRegistry/PublishService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterRegistryServer).PublishService(ctx, req.(*PublishServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterRegistry_ServiceDesc is the grpc.ServiceDesc for ClusterRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "awecloud.ClusterRegistry",
	HandlerType: (*ClusterRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ClusterRegistry_Login_Handler,
		},
		{
			MethodName: "PublishService",
			Handler:    _ClusterRegistry_PublishService_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _ClusterRegistry_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cluster_registry.proto",
}
