// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: app/lottery/cmd/rpc/pb/lottery.proto

package pb

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
	Lottery_SearchLottery_FullMethodName = "/pb.lottery/SearchLottery"
)

// LotteryClient is the client API for Lottery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LotteryClient interface {
	SearchLottery(ctx context.Context, in *SearchLotteryReq, opts ...grpc.CallOption) (*SearchLotteryResp, error)
}

type lotteryClient struct {
	cc grpc.ClientConnInterface
}

func NewLotteryClient(cc grpc.ClientConnInterface) LotteryClient {
	return &lotteryClient{cc}
}

func (c *lotteryClient) SearchLottery(ctx context.Context, in *SearchLotteryReq, opts ...grpc.CallOption) (*SearchLotteryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchLotteryResp)
	err := c.cc.Invoke(ctx, Lottery_SearchLottery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LotteryServer is the server API for Lottery service.
// All implementations must embed UnimplementedLotteryServer
// for forward compatibility.
type LotteryServer interface {
	SearchLottery(context.Context, *SearchLotteryReq) (*SearchLotteryResp, error)
	mustEmbedUnimplementedLotteryServer()
}

// UnimplementedLotteryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLotteryServer struct{}

func (UnimplementedLotteryServer) SearchLottery(context.Context, *SearchLotteryReq) (*SearchLotteryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchLottery not implemented")
}
func (UnimplementedLotteryServer) mustEmbedUnimplementedLotteryServer() {}
func (UnimplementedLotteryServer) testEmbeddedByValue()                 {}

// UnsafeLotteryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LotteryServer will
// result in compilation errors.
type UnsafeLotteryServer interface {
	mustEmbedUnimplementedLotteryServer()
}

func RegisterLotteryServer(s grpc.ServiceRegistrar, srv LotteryServer) {
	// If the following call pancis, it indicates UnimplementedLotteryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Lottery_ServiceDesc, srv)
}

func _Lottery_SearchLottery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchLotteryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LotteryServer).SearchLottery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lottery_SearchLottery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LotteryServer).SearchLottery(ctx, req.(*SearchLotteryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Lottery_ServiceDesc is the grpc.ServiceDesc for Lottery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lottery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.lottery",
	HandlerType: (*LotteryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchLottery",
			Handler:    _Lottery_SearchLottery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/lottery/cmd/rpc/pb/lottery.proto",
}
