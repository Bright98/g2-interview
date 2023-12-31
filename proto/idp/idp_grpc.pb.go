// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0--rc2
// source: idp.proto

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
	IdpService_Login_FullMethodName = "/IdpGrpc.IdpService/Login"
)

// IdpServiceClient is the client API for IdpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdpServiceClient interface {
	Login(ctx context.Context, in *LoginInfoRequest, opts ...grpc.CallOption) (*TokenResponse, error)
}

type idpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdpServiceClient(cc grpc.ClientConnInterface) IdpServiceClient {
	return &idpServiceClient{cc}
}

func (c *idpServiceClient) Login(ctx context.Context, in *LoginInfoRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, IdpService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdpServiceServer is the server API for IdpService service.
// All implementations must embed UnimplementedIdpServiceServer
// for forward compatibility
type IdpServiceServer interface {
	Login(context.Context, *LoginInfoRequest) (*TokenResponse, error)
	mustEmbedUnimplementedIdpServiceServer()
}

// UnimplementedIdpServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIdpServiceServer struct {
}

func (UnimplementedIdpServiceServer) Login(context.Context, *LoginInfoRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedIdpServiceServer) mustEmbedUnimplementedIdpServiceServer() {}

// UnsafeIdpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdpServiceServer will
// result in compilation errors.
type UnsafeIdpServiceServer interface {
	mustEmbedUnimplementedIdpServiceServer()
}

func RegisterIdpServiceServer(s grpc.ServiceRegistrar, srv IdpServiceServer) {
	s.RegisterService(&IdpService_ServiceDesc, srv)
}

func _IdpService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdpServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdpService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdpServiceServer).Login(ctx, req.(*LoginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdpService_ServiceDesc is the grpc.ServiceDesc for IdpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "IdpGrpc.IdpService",
	HandlerType: (*IdpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _IdpService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idp.proto",
}
