// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: server/api.proto

package helloworld

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
	Treasury_GetWithdrawStatus_FullMethodName = "/helloworld.Treasury/GetWithdrawStatus"
	Treasury_CreateWithdraw_FullMethodName    = "/helloworld.Treasury/CreateWithdraw"
	Treasury_ConfirmWithdraw_FullMethodName   = "/helloworld.Treasury/ConfirmWithdraw"
)

// TreasuryClient is the client API for Treasury service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TreasuryClient interface {
	// get withdraw status
	GetWithdrawStatus(ctx context.Context, in *GetWithdrawStatusRequest, opts ...grpc.CallOption) (*GetWithdrawStatusReply, error)
	// create withdraw
	CreateWithdraw(ctx context.Context, in *CreateWithdrawRequest, opts ...grpc.CallOption) (*CreateWithdrawReply, error)
	// confirm withdraw request
	ConfirmWithdraw(ctx context.Context, in *ConfirmWithdrawRequest, opts ...grpc.CallOption) (*ConfirmWithdrawReply, error)
}

type treasuryClient struct {
	cc grpc.ClientConnInterface
}

func NewTreasuryClient(cc grpc.ClientConnInterface) TreasuryClient {
	return &treasuryClient{cc}
}

func (c *treasuryClient) GetWithdrawStatus(ctx context.Context, in *GetWithdrawStatusRequest, opts ...grpc.CallOption) (*GetWithdrawStatusReply, error) {
	out := new(GetWithdrawStatusReply)
	err := c.cc.Invoke(ctx, Treasury_GetWithdrawStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treasuryClient) CreateWithdraw(ctx context.Context, in *CreateWithdrawRequest, opts ...grpc.CallOption) (*CreateWithdrawReply, error) {
	out := new(CreateWithdrawReply)
	err := c.cc.Invoke(ctx, Treasury_CreateWithdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *treasuryClient) ConfirmWithdraw(ctx context.Context, in *ConfirmWithdrawRequest, opts ...grpc.CallOption) (*ConfirmWithdrawReply, error) {
	out := new(ConfirmWithdrawReply)
	err := c.cc.Invoke(ctx, Treasury_ConfirmWithdraw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TreasuryServer is the server API for Treasury service.
// All implementations must embed UnimplementedTreasuryServer
// for forward compatibility
type TreasuryServer interface {
	// get withdraw status
	GetWithdrawStatus(context.Context, *GetWithdrawStatusRequest) (*GetWithdrawStatusReply, error)
	// create withdraw
	CreateWithdraw(context.Context, *CreateWithdrawRequest) (*CreateWithdrawReply, error)
	// confirm withdraw request
	ConfirmWithdraw(context.Context, *ConfirmWithdrawRequest) (*ConfirmWithdrawReply, error)
	mustEmbedUnimplementedTreasuryServer()
}

// UnimplementedTreasuryServer must be embedded to have forward compatible implementations.
type UnimplementedTreasuryServer struct {
}

func (UnimplementedTreasuryServer) GetWithdrawStatus(context.Context, *GetWithdrawStatusRequest) (*GetWithdrawStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawStatus not implemented")
}
func (UnimplementedTreasuryServer) CreateWithdraw(context.Context, *CreateWithdrawRequest) (*CreateWithdrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWithdraw not implemented")
}
func (UnimplementedTreasuryServer) ConfirmWithdraw(context.Context, *ConfirmWithdrawRequest) (*ConfirmWithdrawReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmWithdraw not implemented")
}
func (UnimplementedTreasuryServer) mustEmbedUnimplementedTreasuryServer() {}

// UnsafeTreasuryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TreasuryServer will
// result in compilation errors.
type UnsafeTreasuryServer interface {
	mustEmbedUnimplementedTreasuryServer()
}

func RegisterTreasuryServer(s grpc.ServiceRegistrar, srv TreasuryServer) {
	s.RegisterService(&Treasury_ServiceDesc, srv)
}

func _Treasury_GetWithdrawStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreasuryServer).GetWithdrawStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Treasury_GetWithdrawStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreasuryServer).GetWithdrawStatus(ctx, req.(*GetWithdrawStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Treasury_CreateWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreasuryServer).CreateWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Treasury_CreateWithdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreasuryServer).CreateWithdraw(ctx, req.(*CreateWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Treasury_ConfirmWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmWithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TreasuryServer).ConfirmWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Treasury_ConfirmWithdraw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TreasuryServer).ConfirmWithdraw(ctx, req.(*ConfirmWithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Treasury_ServiceDesc is the grpc.ServiceDesc for Treasury service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Treasury_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Treasury",
	HandlerType: (*TreasuryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWithdrawStatus",
			Handler:    _Treasury_GetWithdrawStatus_Handler,
		},
		{
			MethodName: "CreateWithdraw",
			Handler:    _Treasury_CreateWithdraw_Handler,
		},
		{
			MethodName: "ConfirmWithdraw",
			Handler:    _Treasury_ConfirmWithdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/api.proto",
}