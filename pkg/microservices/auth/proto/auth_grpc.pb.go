// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: auth.proto

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
	Authorization_LogInAdmin_FullMethodName = "/proto.Authorization/LogInAdmin"
	Authorization_SignUp_FullMethodName     = "/proto.Authorization/SignUp"
	Authorization_SignIn_FullMethodName     = "/proto.Authorization/SignIn"
	Authorization_LogOut_FullMethodName     = "/proto.Authorization/LogOut"
)

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationClient interface {
	LogInAdmin(ctx context.Context, in *SignInInput, opts ...grpc.CallOption) (*Cookie, error)
	SignUp(ctx context.Context, in *SignUpInput, opts ...grpc.CallOption) (*UserId, error)
	SignIn(ctx context.Context, in *SignInInput, opts ...grpc.CallOption) (*Cookie, error)
	LogOut(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Empty, error)
}

type authorizationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationClient(cc grpc.ClientConnInterface) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) LogInAdmin(ctx context.Context, in *SignInInput, opts ...grpc.CallOption) (*Cookie, error) {
	out := new(Cookie)
	err := c.cc.Invoke(ctx, Authorization_LogInAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) SignUp(ctx context.Context, in *SignUpInput, opts ...grpc.CallOption) (*UserId, error) {
	out := new(UserId)
	err := c.cc.Invoke(ctx, Authorization_SignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) SignIn(ctx context.Context, in *SignInInput, opts ...grpc.CallOption) (*Cookie, error) {
	out := new(Cookie)
	err := c.cc.Invoke(ctx, Authorization_SignIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) LogOut(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Authorization_LogOut_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
// All implementations must embed UnimplementedAuthorizationServer
// for forward compatibility
type AuthorizationServer interface {
	LogInAdmin(context.Context, *SignInInput) (*Cookie, error)
	SignUp(context.Context, *SignUpInput) (*UserId, error)
	SignIn(context.Context, *SignInInput) (*Cookie, error)
	LogOut(context.Context, *Cookie) (*Empty, error)
	mustEmbedUnimplementedAuthorizationServer()
}

// UnimplementedAuthorizationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (UnimplementedAuthorizationServer) LogInAdmin(context.Context, *SignInInput) (*Cookie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogInAdmin not implemented")
}
func (UnimplementedAuthorizationServer) SignUp(context.Context, *SignUpInput) (*UserId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthorizationServer) SignIn(context.Context, *SignInInput) (*Cookie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthorizationServer) LogOut(context.Context, *Cookie) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogOut not implemented")
}
func (UnimplementedAuthorizationServer) mustEmbedUnimplementedAuthorizationServer() {}

// UnsafeAuthorizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationServer will
// result in compilation errors.
type UnsafeAuthorizationServer interface {
	mustEmbedUnimplementedAuthorizationServer()
}

func RegisterAuthorizationServer(s grpc.ServiceRegistrar, srv AuthorizationServer) {
	s.RegisterService(&Authorization_ServiceDesc, srv)
}

func _Authorization_LogInAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).LogInAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authorization_LogInAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).LogInAdmin(ctx, req.(*SignInInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authorization_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).SignUp(ctx, req.(*SignUpInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authorization_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).SignIn(ctx, req.(*SignInInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_LogOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).LogOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authorization_LogOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).LogOut(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

// Authorization_ServiceDesc is the grpc.ServiceDesc for Authorization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authorization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogInAdmin",
			Handler:    _Authorization_LogInAdmin_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _Authorization_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _Authorization_SignIn_Handler,
		},
		{
			MethodName: "LogOut",
			Handler:    _Authorization_LogOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
