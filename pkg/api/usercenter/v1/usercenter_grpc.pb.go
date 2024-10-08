//*
// Created by costalong on 2024/8/11

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: usercenter/v1/usercenter.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	UserCenter_Login_FullMethodName      = "/usercenter.v1.UserCenter/Login"
	UserCenter_Logout_FullMethodName     = "/usercenter.v1.UserCenter/Logout"
	UserCenter_CreateUser_FullMethodName = "/usercenter.v1.UserCenter/CreateUser"
)

// UserCenterClient is the client API for UserCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCenterClient interface {
	// Login 登录
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// Logout 登出
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// CreateUser
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserReply, error)
}

type userCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCenterClient(cc grpc.ClientConnInterface) UserCenterClient {
	return &userCenterClient{cc}
}

func (c *userCenterClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, UserCenter_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserCenter_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserReply)
	err := c.cc.Invoke(ctx, UserCenter_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCenterServer is the server API for UserCenter service.
// All implementations must embed UnimplementedUserCenterServer
// for forward compatibility
type UserCenterServer interface {
	// Login 登录
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// Logout 登出
	Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error)
	// CreateUser
	CreateUser(context.Context, *CreateUserRequest) (*UserReply, error)
	mustEmbedUnimplementedUserCenterServer()
}

// UnimplementedUserCenterServer must be embedded to have forward compatible implementations.
type UnimplementedUserCenterServer struct {
}

func (UnimplementedUserCenterServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserCenterServer) Logout(context.Context, *LogoutRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserCenterServer) CreateUser(context.Context, *CreateUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserCenterServer) mustEmbedUnimplementedUserCenterServer() {}

// UnsafeUserCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCenterServer will
// result in compilation errors.
type UnsafeUserCenterServer interface {
	mustEmbedUnimplementedUserCenterServer()
}

func RegisterUserCenterServer(s grpc.ServiceRegistrar, srv UserCenterServer) {
	s.RegisterService(&UserCenter_ServiceDesc, srv)
}

func _UserCenter_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenter_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenter_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserCenter_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCenter_ServiceDesc is the grpc.ServiceDesc for UserCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usercenter.v1.UserCenter",
	HandlerType: (*UserCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserCenter_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserCenter_Logout_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserCenter_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usercenter/v1/usercenter.proto",
}
