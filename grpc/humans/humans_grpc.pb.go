// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: humans.proto

package routeguide

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

// HumansClient is the client API for Humans service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HumansClient interface {
	GetHuman(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Human, error)
	CreateHuman(ctx context.Context, in *Human, opts ...grpc.CallOption) (*ID, error)
}

type humansClient struct {
	cc grpc.ClientConnInterface
}

func NewHumansClient(cc grpc.ClientConnInterface) HumansClient {
	return &humansClient{cc}
}

func (c *humansClient) GetHuman(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Human, error) {
	out := new(Human)
	err := c.cc.Invoke(ctx, "/humans.Humans/GetHuman", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *humansClient) CreateHuman(ctx context.Context, in *Human, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/humans.Humans/CreateHuman", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HumansServer is the server API for Humans service.
// All implementations must embed UnimplementedHumansServer
// for forward compatibility
type HumansServer interface {
	GetHuman(context.Context, *ID) (*Human, error)
	CreateHuman(context.Context, *Human) (*ID, error)
	mustEmbedUnimplementedHumansServer()
}

// UnimplementedHumansServer must be embedded to have forward compatible implementations.
type UnimplementedHumansServer struct {
}

func (UnimplementedHumansServer) GetHuman(context.Context, *ID) (*Human, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHuman not implemented")
}
func (UnimplementedHumansServer) CreateHuman(context.Context, *Human) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHuman not implemented")
}
func (UnimplementedHumansServer) mustEmbedUnimplementedHumansServer() {}

// UnsafeHumansServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HumansServer will
// result in compilation errors.
type UnsafeHumansServer interface {
	mustEmbedUnimplementedHumansServer()
}

func RegisterHumansServer(s grpc.ServiceRegistrar, srv HumansServer) {
	s.RegisterService(&Humans_ServiceDesc, srv)
}

func _Humans_GetHuman_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HumansServer).GetHuman(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/humans.Humans/GetHuman",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HumansServer).GetHuman(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Humans_CreateHuman_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Human)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HumansServer).CreateHuman(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/humans.Humans/CreateHuman",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HumansServer).CreateHuman(ctx, req.(*Human))
	}
	return interceptor(ctx, in, info, handler)
}

// Humans_ServiceDesc is the grpc.ServiceDesc for Humans service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Humans_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "humans.Humans",
	HandlerType: (*HumansServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHuman",
			Handler:    _Humans_GetHuman_Handler,
		},
		{
			MethodName: "CreateHuman",
			Handler:    _Humans_CreateHuman_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "humans.proto",
}
