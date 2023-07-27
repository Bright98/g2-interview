// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0--rc2
// source: sso.proto

package grpc

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
	SSOService_CheckSSOValidation_FullMethodName = "/SSOGrpc.SSOService/CheckSSOValidation"
	SSOService_InsertSSOToken_FullMethodName     = "/SSOGrpc.SSOService/InsertSSOToken"
)

// SSOServiceClient is the client API for SSOService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SSOServiceClient interface {
	CheckSSOValidation(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*SSOValidationResponse, error)
	InsertSSOToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type sSOServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSSOServiceClient(cc grpc.ClientConnInterface) SSOServiceClient {
	return &sSOServiceClient{cc}
}

func (c *sSOServiceClient) CheckSSOValidation(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*SSOValidationResponse, error) {
	out := new(SSOValidationResponse)
	err := c.cc.Invoke(ctx, SSOService_CheckSSOValidation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSOServiceClient) InsertSSOToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, SSOService_InsertSSOToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SSOServiceServer is the server API for SSOService service.
// All implementations must embed UnimplementedSSOServiceServer
// for forward compatibility
type SSOServiceServer interface {
	CheckSSOValidation(context.Context, *TokenRequest) (*SSOValidationResponse, error)
	InsertSSOToken(context.Context, *TokenRequest) (*TokenResponse, error)
	mustEmbedUnimplementedSSOServiceServer()
}

// UnimplementedSSOServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSSOServiceServer struct {
}

func (UnimplementedSSOServiceServer) CheckSSOValidation(context.Context, *TokenRequest) (*SSOValidationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSSOValidation not implemented")
}
func (UnimplementedSSOServiceServer) InsertSSOToken(context.Context, *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertSSOToken not implemented")
}
func (UnimplementedSSOServiceServer) mustEmbedUnimplementedSSOServiceServer() {}

// UnsafeSSOServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SSOServiceServer will
// result in compilation errors.
type UnsafeSSOServiceServer interface {
	mustEmbedUnimplementedSSOServiceServer()
}

func RegisterSSOServiceServer(s grpc.ServiceRegistrar, srv SSOServiceServer) {
	s.RegisterService(&SSOService_ServiceDesc, srv)
}

func _SSOService_CheckSSOValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).CheckSSOValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_CheckSSOValidation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).CheckSSOValidation(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSOService_InsertSSOToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSOServiceServer).InsertSSOToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SSOService_InsertSSOToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSOServiceServer).InsertSSOToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SSOService_ServiceDesc is the grpc.ServiceDesc for SSOService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SSOService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SSOGrpc.SSOService",
	HandlerType: (*SSOServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckSSOValidation",
			Handler:    _SSOService_CheckSSOValidation_Handler,
		},
		{
			MethodName: "InsertSSOToken",
			Handler:    _SSOService_InsertSSOToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sso.proto",
}