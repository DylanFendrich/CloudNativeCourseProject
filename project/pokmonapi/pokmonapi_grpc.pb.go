// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pokmonapi

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

// PokmonInfoClient is the client API for PokmonInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokmonInfoClient interface {
	GetMonsterInfo(ctx context.Context, in *MonsterName, opts ...grpc.CallOption) (*MonsterNames, error)
	SetMonsterInfo(ctx context.Context, in *UserAndName, opts ...grpc.CallOption) (*Status, error)
	JoinQueue(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*Status, error)
	//rpc LeaveQueue (UserName) returns (Status) {}                     // user wants to leave the queue because they scared
	LeaveGame(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*Status, error)
	SetUserName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*Status, error)
	//rpc SetCredentials (Credentials) returns (Status) {}              // user sets Credentials to server *(this is extra)* //uncomment credentials lne 47
	GetGameInfo(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*GameStatus, error)
	GetActionInfo(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*AttackActions, error)
}

type pokmonInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewPokmonInfoClient(cc grpc.ClientConnInterface) PokmonInfoClient {
	return &pokmonInfoClient{cc}
}

func (c *pokmonInfoClient) GetMonsterInfo(ctx context.Context, in *MonsterName, opts ...grpc.CallOption) (*MonsterNames, error) {
	out := new(MonsterNames)
	err := c.cc.Invoke(ctx, "/pokmonapi.PokmonInfo/GetMonsterInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokmonInfoClient) SetMonsterInfo(ctx context.Context, in *UserAndName, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/pokmonapi.PokmonInfo/SetMonsterInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokmonInfoClient) JoinQueue(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/pokmonapi.PokmonInfo/JoinQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokmonInfoClient) LeaveGame(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/pokmonapi.PokmonInfo/LeaveGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokmonInfoClient) SetUserName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/pokmonapi.PokmonInfo/SetUserName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokmonInfoClient) GetGameInfo(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*GameStatus, error) {
	out := new(GameStatus)
	err := c.cc.Invoke(ctx, "/pokmonapi.PokmonInfo/GetGameInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokmonInfoClient) GetActionInfo(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*AttackActions, error) {
	out := new(AttackActions)
	err := c.cc.Invoke(ctx, "/pokmonapi.PokmonInfo/GetActionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokmonInfoServer is the server API for PokmonInfo service.
// All implementations must embed UnimplementedPokmonInfoServer
// for forward compatibility
type PokmonInfoServer interface {
	GetMonsterInfo(context.Context, *MonsterName) (*MonsterNames, error)
	SetMonsterInfo(context.Context, *UserAndName) (*Status, error)
	JoinQueue(context.Context, *UserName) (*Status, error)
	//rpc LeaveQueue (UserName) returns (Status) {}                     // user wants to leave the queue because they scared
	LeaveGame(context.Context, *UserName) (*Status, error)
	SetUserName(context.Context, *UserName) (*Status, error)
	//rpc SetCredentials (Credentials) returns (Status) {}              // user sets Credentials to server *(this is extra)* //uncomment credentials lne 47
	GetGameInfo(context.Context, *RequestInfo) (*GameStatus, error)
	GetActionInfo(context.Context, *RequestInfo) (*AttackActions, error)
	mustEmbedUnimplementedPokmonInfoServer()
}

// UnimplementedPokmonInfoServer must be embedded to have forward compatible implementations.
type UnimplementedPokmonInfoServer struct {
}

func (UnimplementedPokmonInfoServer) GetMonsterInfo(context.Context, *MonsterName) (*MonsterNames, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonsterInfo not implemented")
}
func (UnimplementedPokmonInfoServer) SetMonsterInfo(context.Context, *UserAndName) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMonsterInfo not implemented")
}
func (UnimplementedPokmonInfoServer) JoinQueue(context.Context, *UserName) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinQueue not implemented")
}
func (UnimplementedPokmonInfoServer) LeaveGame(context.Context, *UserName) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGame not implemented")
}
func (UnimplementedPokmonInfoServer) SetUserName(context.Context, *UserName) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserName not implemented")
}
func (UnimplementedPokmonInfoServer) GetGameInfo(context.Context, *RequestInfo) (*GameStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameInfo not implemented")
}
func (UnimplementedPokmonInfoServer) GetActionInfo(context.Context, *RequestInfo) (*AttackActions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActionInfo not implemented")
}
func (UnimplementedPokmonInfoServer) mustEmbedUnimplementedPokmonInfoServer() {}

// UnsafePokmonInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokmonInfoServer will
// result in compilation errors.
type UnsafePokmonInfoServer interface {
	mustEmbedUnimplementedPokmonInfoServer()
}

func RegisterPokmonInfoServer(s grpc.ServiceRegistrar, srv PokmonInfoServer) {
	s.RegisterService(&PokmonInfo_ServiceDesc, srv)
}

func _PokmonInfo_GetMonsterInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MonsterName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokmonInfoServer).GetMonsterInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokmonapi.PokmonInfo/GetMonsterInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokmonInfoServer).GetMonsterInfo(ctx, req.(*MonsterName))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokmonInfo_SetMonsterInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAndName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokmonInfoServer).SetMonsterInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokmonapi.PokmonInfo/SetMonsterInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokmonInfoServer).SetMonsterInfo(ctx, req.(*UserAndName))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokmonInfo_JoinQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokmonInfoServer).JoinQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokmonapi.PokmonInfo/JoinQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokmonInfoServer).JoinQueue(ctx, req.(*UserName))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokmonInfo_LeaveGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokmonInfoServer).LeaveGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokmonapi.PokmonInfo/LeaveGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokmonInfoServer).LeaveGame(ctx, req.(*UserName))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokmonInfo_SetUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokmonInfoServer).SetUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokmonapi.PokmonInfo/SetUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokmonInfoServer).SetUserName(ctx, req.(*UserName))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokmonInfo_GetGameInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokmonInfoServer).GetGameInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokmonapi.PokmonInfo/GetGameInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokmonInfoServer).GetGameInfo(ctx, req.(*RequestInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokmonInfo_GetActionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokmonInfoServer).GetActionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pokmonapi.PokmonInfo/GetActionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokmonInfoServer).GetActionInfo(ctx, req.(*RequestInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// PokmonInfo_ServiceDesc is the grpc.ServiceDesc for PokmonInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PokmonInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pokmonapi.PokmonInfo",
	HandlerType: (*PokmonInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMonsterInfo",
			Handler:    _PokmonInfo_GetMonsterInfo_Handler,
		},
		{
			MethodName: "SetMonsterInfo",
			Handler:    _PokmonInfo_SetMonsterInfo_Handler,
		},
		{
			MethodName: "JoinQueue",
			Handler:    _PokmonInfo_JoinQueue_Handler,
		},
		{
			MethodName: "LeaveGame",
			Handler:    _PokmonInfo_LeaveGame_Handler,
		},
		{
			MethodName: "SetUserName",
			Handler:    _PokmonInfo_SetUserName_Handler,
		},
		{
			MethodName: "GetGameInfo",
			Handler:    _PokmonInfo_GetGameInfo_Handler,
		},
		{
			MethodName: "GetActionInfo",
			Handler:    _PokmonInfo_GetActionInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pokmonapi/pokmonapi.proto",
}
