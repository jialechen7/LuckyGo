// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.4
// source: app/usercenter/cmd/rpc/pb/usercenter.proto

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
	Usercenter_Login_FullMethodName                = "/pb.usercenter/login"
	Usercenter_Register_FullMethodName             = "/pb.usercenter/register"
	Usercenter_WxMiniAuth_FullMethodName           = "/pb.usercenter/wxMiniAuth"
	Usercenter_GenerateToken_FullMethodName        = "/pb.usercenter/generateToken"
	Usercenter_GetUserInfo_FullMethodName          = "/pb.usercenter/getUserInfo"
	Usercenter_UpdateUserBaseInfo_FullMethodName   = "/pb.usercenter/updateUserBaseInfo"
	Usercenter_GetUserInfoByUserIds_FullMethodName = "/pb.usercenter/getUserInfoByUserIds"
	Usercenter_GetUserAuthByAuthKey_FullMethodName = "/pb.usercenter/getUserAuthByAuthKey"
	Usercenter_AddUserSponsor_FullMethodName       = "/pb.usercenter/AddUserSponsor"
	Usercenter_UpdateUserSponsor_FullMethodName    = "/pb.usercenter/UpdateUserSponsor"
	Usercenter_DelUserSponsor_FullMethodName       = "/pb.usercenter/DelUserSponsor"
	Usercenter_SearchUserSponsor_FullMethodName    = "/pb.usercenter/SearchUserSponsor"
	Usercenter_SponsorDetail_FullMethodName        = "/pb.usercenter/SponsorDetail"
)

// UsercenterClient is the client API for Usercenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsercenterClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	WxMiniAuth(ctx context.Context, in *WxMiniAuthReq, opts ...grpc.CallOption) (*WxMiniAuthResp, error)
	GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	UpdateUserBaseInfo(ctx context.Context, in *UpdateUserBaseInfoReq, opts ...grpc.CallOption) (*UpdateUserBaseInfoResp, error)
	GetUserInfoByUserIds(ctx context.Context, in *GetUserInfoByUserIdsReq, opts ...grpc.CallOption) (*GetUserInfoByUserIdsResp, error)
	GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error)
	AddUserSponsor(ctx context.Context, in *AddUserSponsorReq, opts ...grpc.CallOption) (*AddUserSponsorResp, error)
	UpdateUserSponsor(ctx context.Context, in *UpdateUserSponsorReq, opts ...grpc.CallOption) (*UpdateUserSponsorResp, error)
	DelUserSponsor(ctx context.Context, in *DelUserSponsorReq, opts ...grpc.CallOption) (*DelUserSponsorResp, error)
	SearchUserSponsor(ctx context.Context, in *SearchUserSponsorReq, opts ...grpc.CallOption) (*SearchUserSponsorResp, error)
	SponsorDetail(ctx context.Context, in *SponsorDetailReq, opts ...grpc.CallOption) (*SponsorDetailResp, error)
}

type usercenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUsercenterClient(cc grpc.ClientConnInterface) UsercenterClient {
	return &usercenterClient{cc}
}

func (c *usercenterClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Usercenter_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, Usercenter_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) WxMiniAuth(ctx context.Context, in *WxMiniAuthReq, opts ...grpc.CallOption) (*WxMiniAuthResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WxMiniAuthResp)
	err := c.cc.Invoke(ctx, Usercenter_WxMiniAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateTokenResp)
	err := c.cc.Invoke(ctx, Usercenter_GenerateToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserInfoResp)
	err := c.cc.Invoke(ctx, Usercenter_GetUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UpdateUserBaseInfo(ctx context.Context, in *UpdateUserBaseInfoReq, opts ...grpc.CallOption) (*UpdateUserBaseInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserBaseInfoResp)
	err := c.cc.Invoke(ctx, Usercenter_UpdateUserBaseInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserInfoByUserIds(ctx context.Context, in *GetUserInfoByUserIdsReq, opts ...grpc.CallOption) (*GetUserInfoByUserIdsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserInfoByUserIdsResp)
	err := c.cc.Invoke(ctx, Usercenter_GetUserInfoByUserIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserAuthByAuthKeyResp)
	err := c.cc.Invoke(ctx, Usercenter_GetUserAuthByAuthKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) AddUserSponsor(ctx context.Context, in *AddUserSponsorReq, opts ...grpc.CallOption) (*AddUserSponsorResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserSponsorResp)
	err := c.cc.Invoke(ctx, Usercenter_AddUserSponsor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UpdateUserSponsor(ctx context.Context, in *UpdateUserSponsorReq, opts ...grpc.CallOption) (*UpdateUserSponsorResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserSponsorResp)
	err := c.cc.Invoke(ctx, Usercenter_UpdateUserSponsor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) DelUserSponsor(ctx context.Context, in *DelUserSponsorReq, opts ...grpc.CallOption) (*DelUserSponsorResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelUserSponsorResp)
	err := c.cc.Invoke(ctx, Usercenter_DelUserSponsor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) SearchUserSponsor(ctx context.Context, in *SearchUserSponsorReq, opts ...grpc.CallOption) (*SearchUserSponsorResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchUserSponsorResp)
	err := c.cc.Invoke(ctx, Usercenter_SearchUserSponsor_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) SponsorDetail(ctx context.Context, in *SponsorDetailReq, opts ...grpc.CallOption) (*SponsorDetailResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SponsorDetailResp)
	err := c.cc.Invoke(ctx, Usercenter_SponsorDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsercenterServer is the server API for Usercenter service.
// All implementations must embed UnimplementedUsercenterServer
// for forward compatibility.
type UsercenterServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	WxMiniAuth(context.Context, *WxMiniAuthReq) (*WxMiniAuthResp, error)
	GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error)
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	UpdateUserBaseInfo(context.Context, *UpdateUserBaseInfoReq) (*UpdateUserBaseInfoResp, error)
	GetUserInfoByUserIds(context.Context, *GetUserInfoByUserIdsReq) (*GetUserInfoByUserIdsResp, error)
	GetUserAuthByAuthKey(context.Context, *GetUserAuthByAuthKeyReq) (*GetUserAuthByAuthKeyResp, error)
	AddUserSponsor(context.Context, *AddUserSponsorReq) (*AddUserSponsorResp, error)
	UpdateUserSponsor(context.Context, *UpdateUserSponsorReq) (*UpdateUserSponsorResp, error)
	DelUserSponsor(context.Context, *DelUserSponsorReq) (*DelUserSponsorResp, error)
	SearchUserSponsor(context.Context, *SearchUserSponsorReq) (*SearchUserSponsorResp, error)
	SponsorDetail(context.Context, *SponsorDetailReq) (*SponsorDetailResp, error)
	mustEmbedUnimplementedUsercenterServer()
}

// UnimplementedUsercenterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsercenterServer struct{}

func (UnimplementedUsercenterServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUsercenterServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUsercenterServer) WxMiniAuth(context.Context, *WxMiniAuthReq) (*WxMiniAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxMiniAuth not implemented")
}
func (UnimplementedUsercenterServer) GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedUsercenterServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUsercenterServer) UpdateUserBaseInfo(context.Context, *UpdateUserBaseInfoReq) (*UpdateUserBaseInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserBaseInfo not implemented")
}
func (UnimplementedUsercenterServer) GetUserInfoByUserIds(context.Context, *GetUserInfoByUserIdsReq) (*GetUserInfoByUserIdsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByUserIds not implemented")
}
func (UnimplementedUsercenterServer) GetUserAuthByAuthKey(context.Context, *GetUserAuthByAuthKeyReq) (*GetUserAuthByAuthKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAuthByAuthKey not implemented")
}
func (UnimplementedUsercenterServer) AddUserSponsor(context.Context, *AddUserSponsorReq) (*AddUserSponsorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserSponsor not implemented")
}
func (UnimplementedUsercenterServer) UpdateUserSponsor(context.Context, *UpdateUserSponsorReq) (*UpdateUserSponsorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserSponsor not implemented")
}
func (UnimplementedUsercenterServer) DelUserSponsor(context.Context, *DelUserSponsorReq) (*DelUserSponsorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUserSponsor not implemented")
}
func (UnimplementedUsercenterServer) SearchUserSponsor(context.Context, *SearchUserSponsorReq) (*SearchUserSponsorResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserSponsor not implemented")
}
func (UnimplementedUsercenterServer) SponsorDetail(context.Context, *SponsorDetailReq) (*SponsorDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SponsorDetail not implemented")
}
func (UnimplementedUsercenterServer) mustEmbedUnimplementedUsercenterServer() {}
func (UnimplementedUsercenterServer) testEmbeddedByValue()                    {}

// UnsafeUsercenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsercenterServer will
// result in compilation errors.
type UnsafeUsercenterServer interface {
	mustEmbedUnimplementedUsercenterServer()
}

func RegisterUsercenterServer(s grpc.ServiceRegistrar, srv UsercenterServer) {
	// If the following call pancis, it indicates UnimplementedUsercenterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Usercenter_ServiceDesc, srv)
}

func _Usercenter_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_WxMiniAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WxMiniAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).WxMiniAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_WxMiniAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).WxMiniAuth(ctx, req.(*WxMiniAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_GenerateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GenerateToken(ctx, req.(*GenerateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UpdateUserBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserBaseInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UpdateUserBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_UpdateUserBaseInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UpdateUserBaseInfo(ctx, req.(*UpdateUserBaseInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserInfoByUserIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByUserIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserInfoByUserIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_GetUserInfoByUserIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserInfoByUserIds(ctx, req.(*GetUserInfoByUserIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserAuthByAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAuthByAuthKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserAuthByAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_GetUserAuthByAuthKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserAuthByAuthKey(ctx, req.(*GetUserAuthByAuthKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_AddUserSponsor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserSponsorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).AddUserSponsor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_AddUserSponsor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).AddUserSponsor(ctx, req.(*AddUserSponsorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UpdateUserSponsor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserSponsorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UpdateUserSponsor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_UpdateUserSponsor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UpdateUserSponsor(ctx, req.(*UpdateUserSponsorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_DelUserSponsor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserSponsorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).DelUserSponsor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_DelUserSponsor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).DelUserSponsor(ctx, req.(*DelUserSponsorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_SearchUserSponsor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserSponsorReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).SearchUserSponsor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_SearchUserSponsor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).SearchUserSponsor(ctx, req.(*SearchUserSponsorReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_SponsorDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SponsorDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).SponsorDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_SponsorDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).SponsorDetail(ctx, req.(*SponsorDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Usercenter_ServiceDesc is the grpc.ServiceDesc for Usercenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usercenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.usercenter",
	HandlerType: (*UsercenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "login",
			Handler:    _Usercenter_Login_Handler,
		},
		{
			MethodName: "register",
			Handler:    _Usercenter_Register_Handler,
		},
		{
			MethodName: "wxMiniAuth",
			Handler:    _Usercenter_WxMiniAuth_Handler,
		},
		{
			MethodName: "generateToken",
			Handler:    _Usercenter_GenerateToken_Handler,
		},
		{
			MethodName: "getUserInfo",
			Handler:    _Usercenter_GetUserInfo_Handler,
		},
		{
			MethodName: "updateUserBaseInfo",
			Handler:    _Usercenter_UpdateUserBaseInfo_Handler,
		},
		{
			MethodName: "getUserInfoByUserIds",
			Handler:    _Usercenter_GetUserInfoByUserIds_Handler,
		},
		{
			MethodName: "getUserAuthByAuthKey",
			Handler:    _Usercenter_GetUserAuthByAuthKey_Handler,
		},
		{
			MethodName: "AddUserSponsor",
			Handler:    _Usercenter_AddUserSponsor_Handler,
		},
		{
			MethodName: "UpdateUserSponsor",
			Handler:    _Usercenter_UpdateUserSponsor_Handler,
		},
		{
			MethodName: "DelUserSponsor",
			Handler:    _Usercenter_DelUserSponsor_Handler,
		},
		{
			MethodName: "SearchUserSponsor",
			Handler:    _Usercenter_SearchUserSponsor_Handler,
		},
		{
			MethodName: "SponsorDetail",
			Handler:    _Usercenter_SponsorDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/usercenter/cmd/rpc/pb/usercenter.proto",
}
