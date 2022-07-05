// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/protos/v1/simpleservice.proto

package simpleservicev1

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

// SimpleServiceClient is the client API for SimpleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleServiceClient interface {
	// Lists transactions of a address. Returns NOT_FOUND if the address does not exist.
	ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error)
	// Gets balance of a address. Returns NOT_FOUND if the address does not exist.
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*Balance, error)
}

type simpleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleServiceClient(cc grpc.ClientConnInterface) SimpleServiceClient {
	return &simpleServiceClient{cc}
}

func (c *simpleServiceClient) ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error) {
	out := new(ListTransactionsResponse)
	err := c.cc.Invoke(ctx, "/api.protos.v1.SimpleService/ListTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*Balance, error) {
	out := new(Balance)
	err := c.cc.Invoke(ctx, "/api.protos.v1.SimpleService/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleServiceServer is the server API for SimpleService service.
// All implementations should embed UnimplementedSimpleServiceServer
// for forward compatibility
type SimpleServiceServer interface {
	// Lists transactions of a address. Returns NOT_FOUND if the address does not exist.
	ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error)
	// Gets balance of a address. Returns NOT_FOUND if the address does not exist.
	GetBalance(context.Context, *GetBalanceRequest) (*Balance, error)
}

// UnimplementedSimpleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSimpleServiceServer struct {
}

func (UnimplementedSimpleServiceServer) ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactions not implemented")
}
func (UnimplementedSimpleServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*Balance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}

// UnsafeSimpleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleServiceServer will
// result in compilation errors.
type UnsafeSimpleServiceServer interface {
	mustEmbedUnimplementedSimpleServiceServer()
}

func RegisterSimpleServiceServer(s grpc.ServiceRegistrar, srv SimpleServiceServer) {
	s.RegisterService(&SimpleService_ServiceDesc, srv)
}

func _SimpleService_ListTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServiceServer).ListTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.protos.v1.SimpleService/ListTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServiceServer).ListTransactions(ctx, req.(*ListTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.protos.v1.SimpleService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SimpleService_ServiceDesc is the grpc.ServiceDesc for SimpleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.protos.v1.SimpleService",
	HandlerType: (*SimpleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTransactions",
			Handler:    _SimpleService_ListTransactions_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _SimpleService_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protos/v1/simpleservice.proto",
}
