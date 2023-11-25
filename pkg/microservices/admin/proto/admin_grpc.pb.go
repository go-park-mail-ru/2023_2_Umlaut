// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: admin.proto

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
	Admin_GetAllStatistic_FullMethodName      = "/proto.Admin/GetAllStatistic"
	Admin_CreateStatistic_FullMethodName      = "/proto.Admin/CreateStatistic"
	Admin_CreateRecommendation_FullMethodName = "/proto.Admin/CreateRecommendation"
)

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClient interface {
	GetAllStatistic(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Statistic, error)
	CreateStatistic(ctx context.Context, in *Statistic, opts ...grpc.CallOption) (*Empty, error)
	CreateRecommendation(ctx context.Context, in *Recommendation, opts ...grpc.CallOption) (*Empty, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) GetAllStatistic(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Statistic, error) {
	out := new(Statistic)
	err := c.cc.Invoke(ctx, Admin_GetAllStatistic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CreateStatistic(ctx context.Context, in *Statistic, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Admin_CreateStatistic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CreateRecommendation(ctx context.Context, in *Recommendation, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Admin_CreateRecommendation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
// All implementations must embed UnimplementedAdminServer
// for forward compatibility
type AdminServer interface {
	GetAllStatistic(context.Context, *Empty) (*Statistic, error)
	CreateStatistic(context.Context, *Statistic) (*Empty, error)
	CreateRecommendation(context.Context, *Recommendation) (*Empty, error)
	mustEmbedUnimplementedAdminServer()
}

// UnimplementedAdminServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (UnimplementedAdminServer) GetAllStatistic(context.Context, *Empty) (*Statistic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStatistic not implemented")
}
func (UnimplementedAdminServer) CreateStatistic(context.Context, *Statistic) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStatistic not implemented")
}
func (UnimplementedAdminServer) CreateRecommendation(context.Context, *Recommendation) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecommendation not implemented")
}
func (UnimplementedAdminServer) mustEmbedUnimplementedAdminServer() {}

// UnsafeAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServer will
// result in compilation errors.
type UnsafeAdminServer interface {
	mustEmbedUnimplementedAdminServer()
}

func RegisterAdminServer(s grpc.ServiceRegistrar, srv AdminServer) {
	s.RegisterService(&Admin_ServiceDesc, srv)
}

func _Admin_GetAllStatistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetAllStatistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_GetAllStatistic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetAllStatistic(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CreateStatistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Statistic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateStatistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_CreateStatistic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateStatistic(ctx, req.(*Statistic))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CreateRecommendation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Recommendation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateRecommendation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Admin_CreateRecommendation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateRecommendation(ctx, req.(*Recommendation))
	}
	return interceptor(ctx, in, info, handler)
}

// Admin_ServiceDesc is the grpc.ServiceDesc for Admin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Admin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllStatistic",
			Handler:    _Admin_GetAllStatistic_Handler,
		},
		{
			MethodName: "CreateStatistic",
			Handler:    _Admin_CreateStatistic_Handler,
		},
		{
			MethodName: "CreateRecommendation",
			Handler:    _Admin_CreateRecommendation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
