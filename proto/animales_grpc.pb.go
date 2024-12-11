// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/animales.proto

package proto

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
	Animalesservice_GetAnimalesInfo_FullMethodName = "/animales.animalesservice/GetAnimalesInfo"
)

// AnimalesserviceClient is the client API for Animalesservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnimalesserviceClient interface {
	GetAnimalesInfo(ctx context.Context, in *AnimalRequest, opts ...grpc.CallOption) (*AnimalResponse, error)
}

type animalesserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnimalesserviceClient(cc grpc.ClientConnInterface) AnimalesserviceClient {
	return &animalesserviceClient{cc}
}

func (c *animalesserviceClient) GetAnimalesInfo(ctx context.Context, in *AnimalRequest, opts ...grpc.CallOption) (*AnimalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AnimalResponse)
	err := c.cc.Invoke(ctx, Animalesservice_GetAnimalesInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnimalesserviceServer is the server API for Animalesservice service.
// All implementations must embed UnimplementedAnimalesserviceServer
// for forward compatibility.
type AnimalesserviceServer interface {
	GetAnimalesInfo(context.Context, *AnimalRequest) (*AnimalResponse, error)
	mustEmbedUnimplementedAnimalesserviceServer()
}

// UnimplementedAnimalesserviceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAnimalesserviceServer struct{}

func (UnimplementedAnimalesserviceServer) GetAnimalesInfo(context.Context, *AnimalRequest) (*AnimalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnimalesInfo not implemented")
}
func (UnimplementedAnimalesserviceServer) mustEmbedUnimplementedAnimalesserviceServer() {}
func (UnimplementedAnimalesserviceServer) testEmbeddedByValue()                         {}

// UnsafeAnimalesserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnimalesserviceServer will
// result in compilation errors.
type UnsafeAnimalesserviceServer interface {
	mustEmbedUnimplementedAnimalesserviceServer()
}

func RegisterAnimalesserviceServer(s grpc.ServiceRegistrar, srv AnimalesserviceServer) {
	// If the following call pancis, it indicates UnimplementedAnimalesserviceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Animalesservice_ServiceDesc, srv)
}

func _Animalesservice_GetAnimalesInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalesserviceServer).GetAnimalesInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Animalesservice_GetAnimalesInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalesserviceServer).GetAnimalesInfo(ctx, req.(*AnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Animalesservice_ServiceDesc is the grpc.ServiceDesc for Animalesservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Animalesservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "animales.animalesservice",
	HandlerType: (*AnimalesserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAnimalesInfo",
			Handler:    _Animalesservice_GetAnimalesInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/animales.proto",
}
