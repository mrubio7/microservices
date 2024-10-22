// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/players.proto

package players

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
	PlayerService_GetPlayer_FullMethodName           = "/players.PlayerService/GetPlayer"
	PlayerService_GetPlayers_FullMethodName          = "/players.PlayerService/GetPlayers"
	PlayerService_GetProminentPlayers_FullMethodName = "/players.PlayerService/GetProminentPlayers"
	PlayerService_NewPlayer_FullMethodName           = "/players.PlayerService/NewPlayer"
)

// PlayerServiceClient is the client API for PlayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayerServiceClient interface {
	GetPlayer(ctx context.Context, in *GetPlayerRequest, opts ...grpc.CallOption) (*PlayerList, error)
	GetPlayers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlayerList, error)
	GetProminentPlayers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProminentPlayerList, error)
	NewPlayer(ctx context.Context, in *NewPlayerRequest, opts ...grpc.CallOption) (*PlayerResponse, error)
}

type playerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerServiceClient(cc grpc.ClientConnInterface) PlayerServiceClient {
	return &playerServiceClient{cc}
}

func (c *playerServiceClient) GetPlayer(ctx context.Context, in *GetPlayerRequest, opts ...grpc.CallOption) (*PlayerList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlayerList)
	err := c.cc.Invoke(ctx, PlayerService_GetPlayer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) GetPlayers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlayerList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlayerList)
	err := c.cc.Invoke(ctx, PlayerService_GetPlayers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) GetProminentPlayers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProminentPlayerList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProminentPlayerList)
	err := c.cc.Invoke(ctx, PlayerService_GetProminentPlayers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerServiceClient) NewPlayer(ctx context.Context, in *NewPlayerRequest, opts ...grpc.CallOption) (*PlayerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlayerResponse)
	err := c.cc.Invoke(ctx, PlayerService_NewPlayer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerServiceServer is the server API for PlayerService service.
// All implementations must embed UnimplementedPlayerServiceServer
// for forward compatibility.
type PlayerServiceServer interface {
	GetPlayer(context.Context, *GetPlayerRequest) (*PlayerList, error)
	GetPlayers(context.Context, *Empty) (*PlayerList, error)
	GetProminentPlayers(context.Context, *Empty) (*ProminentPlayerList, error)
	NewPlayer(context.Context, *NewPlayerRequest) (*PlayerResponse, error)
	mustEmbedUnimplementedPlayerServiceServer()
}

// UnimplementedPlayerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPlayerServiceServer struct{}

func (UnimplementedPlayerServiceServer) GetPlayer(context.Context, *GetPlayerRequest) (*PlayerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayer not implemented")
}
func (UnimplementedPlayerServiceServer) GetPlayers(context.Context, *Empty) (*PlayerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayers not implemented")
}
func (UnimplementedPlayerServiceServer) GetProminentPlayers(context.Context, *Empty) (*ProminentPlayerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProminentPlayers not implemented")
}
func (UnimplementedPlayerServiceServer) NewPlayer(context.Context, *NewPlayerRequest) (*PlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewPlayer not implemented")
}
func (UnimplementedPlayerServiceServer) mustEmbedUnimplementedPlayerServiceServer() {}
func (UnimplementedPlayerServiceServer) testEmbeddedByValue()                       {}

// UnsafePlayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayerServiceServer will
// result in compilation errors.
type UnsafePlayerServiceServer interface {
	mustEmbedUnimplementedPlayerServiceServer()
}

func RegisterPlayerServiceServer(s grpc.ServiceRegistrar, srv PlayerServiceServer) {
	// If the following call pancis, it indicates UnimplementedPlayerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PlayerService_ServiceDesc, srv)
}

func _PlayerService_GetPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).GetPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayerService_GetPlayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).GetPlayer(ctx, req.(*GetPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_GetPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).GetPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayerService_GetPlayers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).GetPlayers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_GetProminentPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).GetProminentPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayerService_GetProminentPlayers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).GetProminentPlayers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerService_NewPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerServiceServer).NewPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PlayerService_NewPlayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerServiceServer).NewPlayer(ctx, req.(*NewPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayerService_ServiceDesc is the grpc.ServiceDesc for PlayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "players.PlayerService",
	HandlerType: (*PlayerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayer",
			Handler:    _PlayerService_GetPlayer_Handler,
		},
		{
			MethodName: "GetPlayers",
			Handler:    _PlayerService_GetPlayers_Handler,
		},
		{
			MethodName: "GetProminentPlayers",
			Handler:    _PlayerService_GetProminentPlayers_Handler,
		},
		{
			MethodName: "NewPlayer",
			Handler:    _PlayerService_NewPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/players.proto",
}
