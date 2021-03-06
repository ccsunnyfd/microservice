// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// MobileClient is the client API for Mobile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MobileClient interface {
	// Send to mobile
	SendMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*MobileReply, error)
}

type mobileClient struct {
	cc grpc.ClientConnInterface
}

func NewMobileClient(cc grpc.ClientConnInterface) MobileClient {
	return &mobileClient{cc}
}

func (c *mobileClient) SendMobile(ctx context.Context, in *MobileRequest, opts ...grpc.CallOption) (*MobileReply, error) {
	out := new(MobileReply)
	err := c.cc.Invoke(ctx, "/message.v1.Mobile/SendMobile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileServer is the server API for Mobile service.
// All implementations must embed UnimplementedMobileServer
// for forward compatibility
type MobileServer interface {
	// Send to mobile
	SendMobile(context.Context, *MobileRequest) (*MobileReply, error)
	mustEmbedUnimplementedMobileServer()
}

// UnimplementedMobileServer must be embedded to have forward compatible implementations.
type UnimplementedMobileServer struct {
}

func (UnimplementedMobileServer) SendMobile(context.Context, *MobileRequest) (*MobileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMobile not implemented")
}
func (UnimplementedMobileServer) mustEmbedUnimplementedMobileServer() {}

// UnsafeMobileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MobileServer will
// result in compilation errors.
type UnsafeMobileServer interface {
	mustEmbedUnimplementedMobileServer()
}

func RegisterMobileServer(s grpc.ServiceRegistrar, srv MobileServer) {
	s.RegisterService(&Mobile_ServiceDesc, srv)
}

func _Mobile_SendMobile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MobileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileServer).SendMobile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.v1.Mobile/SendMobile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileServer).SendMobile(ctx, req.(*MobileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mobile_ServiceDesc is the grpc.ServiceDesc for Mobile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mobile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.v1.Mobile",
	HandlerType: (*MobileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMobile",
			Handler:    _Mobile_SendMobile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/message/service/v1/message.proto",
}

// EmailClient is the client API for Email service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailClient interface {
	// Send to email
	SendEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailReply, error)
}

type emailClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailClient(cc grpc.ClientConnInterface) EmailClient {
	return &emailClient{cc}
}

func (c *emailClient) SendEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailReply, error) {
	out := new(EmailReply)
	err := c.cc.Invoke(ctx, "/message.v1.Email/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServer is the server API for Email service.
// All implementations must embed UnimplementedEmailServer
// for forward compatibility
type EmailServer interface {
	// Send to email
	SendEmail(context.Context, *EmailRequest) (*EmailReply, error)
	mustEmbedUnimplementedEmailServer()
}

// UnimplementedEmailServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServer struct {
}

func (UnimplementedEmailServer) SendEmail(context.Context, *EmailRequest) (*EmailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedEmailServer) mustEmbedUnimplementedEmailServer() {}

// UnsafeEmailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServer will
// result in compilation errors.
type UnsafeEmailServer interface {
	mustEmbedUnimplementedEmailServer()
}

func RegisterEmailServer(s grpc.ServiceRegistrar, srv EmailServer) {
	s.RegisterService(&Email_ServiceDesc, srv)
}

func _Email_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.v1.Email/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).SendEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Email_ServiceDesc is the grpc.ServiceDesc for Email service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Email_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.v1.Email",
	HandlerType: (*EmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _Email_SendEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/message/service/v1/message.proto",
}
