// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: proto/omikuji.proto

package grpc

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
	OmikujiService_Draw_FullMethodName         = "/omikuji.OmikujiService/Draw"
	OmikujiService_GetHistories_FullMethodName = "/omikuji.OmikujiService/GetHistories"
)

// OmikujiServiceClient is the client API for OmikujiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OmikujiServiceClient interface {
	Draw(ctx context.Context, in *DrawRequest, opts ...grpc.CallOption) (*DrawResponse, error)
	GetHistories(ctx context.Context, in *GetHistoriesRequest, opts ...grpc.CallOption) (*GetHistoriesResponse, error)
}

type omikujiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOmikujiServiceClient(cc grpc.ClientConnInterface) OmikujiServiceClient {
	return &omikujiServiceClient{cc}
}

func (c *omikujiServiceClient) Draw(ctx context.Context, in *DrawRequest, opts ...grpc.CallOption) (*DrawResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DrawResponse)
	err := c.cc.Invoke(ctx, OmikujiService_Draw_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omikujiServiceClient) GetHistories(ctx context.Context, in *GetHistoriesRequest, opts ...grpc.CallOption) (*GetHistoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHistoriesResponse)
	err := c.cc.Invoke(ctx, OmikujiService_GetHistories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OmikujiServiceServer is the server API for OmikujiService service.
// All implementations must embed UnimplementedOmikujiServiceServer
// for forward compatibility.
type OmikujiServiceServer interface {
	Draw(context.Context, *DrawRequest) (*DrawResponse, error)
	GetHistories(context.Context, *GetHistoriesRequest) (*GetHistoriesResponse, error)
	mustEmbedUnimplementedOmikujiServiceServer()
}

// UnimplementedOmikujiServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOmikujiServiceServer struct{}

func (UnimplementedOmikujiServiceServer) Draw(context.Context, *DrawRequest) (*DrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Draw not implemented")
}
func (UnimplementedOmikujiServiceServer) GetHistories(context.Context, *GetHistoriesRequest) (*GetHistoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistories not implemented")
}
func (UnimplementedOmikujiServiceServer) mustEmbedUnimplementedOmikujiServiceServer() {}
func (UnimplementedOmikujiServiceServer) testEmbeddedByValue()                        {}

// UnsafeOmikujiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OmikujiServiceServer will
// result in compilation errors.
type UnsafeOmikujiServiceServer interface {
	mustEmbedUnimplementedOmikujiServiceServer()
}

func RegisterOmikujiServiceServer(s grpc.ServiceRegistrar, srv OmikujiServiceServer) {
	// If the following call pancis, it indicates UnimplementedOmikujiServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OmikujiService_ServiceDesc, srv)
}

func _OmikujiService_Draw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmikujiServiceServer).Draw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmikujiService_Draw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmikujiServiceServer).Draw(ctx, req.(*DrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmikujiService_GetHistories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmikujiServiceServer).GetHistories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmikujiService_GetHistories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmikujiServiceServer).GetHistories(ctx, req.(*GetHistoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OmikujiService_ServiceDesc is the grpc.ServiceDesc for OmikujiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OmikujiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "omikuji.OmikujiService",
	HandlerType: (*OmikujiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Draw",
			Handler:    _OmikujiService_Draw_Handler,
		},
		{
			MethodName: "GetHistories",
			Handler:    _OmikujiService_GetHistories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/omikuji.proto",
}
