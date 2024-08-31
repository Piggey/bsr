// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: bsr.proto

package bsr

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Bsr_CreateGame_FullMethodName = "/bsr.Bsr/CreateGame"
	Bsr_JoinGame_FullMethodName   = "/bsr.Bsr/JoinGame"
)

// BsrClient is the client API for Bsr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BsrClient interface {
	CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error)
	JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (*JoinGameResponse, error)
}

type bsrClient struct {
	cc grpc.ClientConnInterface
}

func NewBsrClient(cc grpc.ClientConnInterface) BsrClient {
	return &bsrClient{cc}
}

func (c *bsrClient) CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateGameResponse)
	err := c.cc.Invoke(ctx, Bsr_CreateGame_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bsrClient) JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (*JoinGameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinGameResponse)
	err := c.cc.Invoke(ctx, Bsr_JoinGame_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BsrServer is the server API for Bsr service.
// All implementations must embed UnimplementedBsrServer
// for forward compatibility.
type BsrServer interface {
	CreateGame(context.Context, *CreateGameRequest) (*CreateGameResponse, error)
	JoinGame(context.Context, *JoinGameRequest) (*JoinGameResponse, error)
	mustEmbedUnimplementedBsrServer()
}

// UnimplementedBsrServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBsrServer struct{}

func (UnimplementedBsrServer) CreateGame(context.Context, *CreateGameRequest) (*CreateGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedBsrServer) JoinGame(context.Context, *JoinGameRequest) (*JoinGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (UnimplementedBsrServer) mustEmbedUnimplementedBsrServer() {}
func (UnimplementedBsrServer) testEmbeddedByValue()             {}

// UnsafeBsrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BsrServer will
// result in compilation errors.
type UnsafeBsrServer interface {
	mustEmbedUnimplementedBsrServer()
}

func RegisterBsrServer(s grpc.ServiceRegistrar, srv BsrServer) {
	// If the following call pancis, it indicates UnimplementedBsrServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Bsr_ServiceDesc, srv)
}

func _Bsr_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BsrServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bsr_CreateGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BsrServer).CreateGame(ctx, req.(*CreateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bsr_JoinGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BsrServer).JoinGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bsr_JoinGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BsrServer).JoinGame(ctx, req.(*JoinGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bsr_ServiceDesc is the grpc.ServiceDesc for Bsr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bsr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bsr.Bsr",
	HandlerType: (*BsrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGame",
			Handler:    _Bsr_CreateGame_Handler,
		},
		{
			MethodName: "JoinGame",
			Handler:    _Bsr_JoinGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bsr.proto",
}