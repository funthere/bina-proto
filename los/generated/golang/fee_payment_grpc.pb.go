// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: fee_payment.proto

package golang

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
	FeePayment_CreateFeePayment_FullMethodName = "/protobuf.los.FeePayment/CreateFeePayment"
)

// FeePaymentClient is the client API for FeePayment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeePaymentClient interface {
	CreateFeePayment(ctx context.Context, in *ReqFeePayment, opts ...grpc.CallOption) (*ResFeePayment, error)
}

type feePaymentClient struct {
	cc grpc.ClientConnInterface
}

func NewFeePaymentClient(cc grpc.ClientConnInterface) FeePaymentClient {
	return &feePaymentClient{cc}
}

func (c *feePaymentClient) CreateFeePayment(ctx context.Context, in *ReqFeePayment, opts ...grpc.CallOption) (*ResFeePayment, error) {
	out := new(ResFeePayment)
	err := c.cc.Invoke(ctx, FeePayment_CreateFeePayment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeePaymentServer is the server API for FeePayment service.
// All implementations should embed UnimplementedFeePaymentServer
// for forward compatibility
type FeePaymentServer interface {
	CreateFeePayment(context.Context, *ReqFeePayment) (*ResFeePayment, error)
}

// UnimplementedFeePaymentServer should be embedded to have forward compatible implementations.
type UnimplementedFeePaymentServer struct {
}

func (UnimplementedFeePaymentServer) CreateFeePayment(context.Context, *ReqFeePayment) (*ResFeePayment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeePayment not implemented")
}

// UnsafeFeePaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeePaymentServer will
// result in compilation errors.
type UnsafeFeePaymentServer interface {
	mustEmbedUnimplementedFeePaymentServer()
}

func RegisterFeePaymentServer(s grpc.ServiceRegistrar, srv FeePaymentServer) {
	s.RegisterService(&FeePayment_ServiceDesc, srv)
}

func _FeePayment_CreateFeePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqFeePayment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeePaymentServer).CreateFeePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeePayment_CreateFeePayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeePaymentServer).CreateFeePayment(ctx, req.(*ReqFeePayment))
	}
	return interceptor(ctx, in, info, handler)
}

// FeePayment_ServiceDesc is the grpc.ServiceDesc for FeePayment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeePayment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.los.FeePayment",
	HandlerType: (*FeePaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeePayment",
			Handler:    _FeePayment_CreateFeePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fee_payment.proto",
}
