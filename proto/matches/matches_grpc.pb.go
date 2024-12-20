// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/matches.proto

package matches

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
	MatchesService_GetAllMatches_FullMethodName      = "/matches.MatchesService/GetAllMatches"
	MatchesService_GetNearbyMatches_FullMethodName   = "/matches.MatchesService/GetNearbyMatches"
	MatchesService_GetMatchByFaceitId_FullMethodName = "/matches.MatchesService/GetMatchByFaceitId"
	MatchesService_SetStreamToMatch_FullMethodName   = "/matches.MatchesService/SetStreamToMatch"
	MatchesService_GetMatchesByTeamId_FullMethodName = "/matches.MatchesService/GetMatchesByTeamId"
	MatchesService_NewMatch_FullMethodName           = "/matches.MatchesService/NewMatch"
)

// MatchesServiceClient is the client API for MatchesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchesServiceClient interface {
	GetAllMatches(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MatchList, error)
	GetNearbyMatches(ctx context.Context, in *GetNearbyMatchesRequest, opts ...grpc.CallOption) (*MatchList, error)
	GetMatchByFaceitId(ctx context.Context, in *GetMatchRequest, opts ...grpc.CallOption) (*Match, error)
	SetStreamToMatch(ctx context.Context, in *SetStreamRequest, opts ...grpc.CallOption) (*Bool, error)
	GetMatchesByTeamId(ctx context.Context, in *GetMatchRequest, opts ...grpc.CallOption) (*MatchList, error)
	NewMatch(ctx context.Context, in *NewMatchRequest, opts ...grpc.CallOption) (*Match, error)
}

type matchesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchesServiceClient(cc grpc.ClientConnInterface) MatchesServiceClient {
	return &matchesServiceClient{cc}
}

func (c *matchesServiceClient) GetAllMatches(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MatchList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MatchList)
	err := c.cc.Invoke(ctx, MatchesService_GetAllMatches_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchesServiceClient) GetNearbyMatches(ctx context.Context, in *GetNearbyMatchesRequest, opts ...grpc.CallOption) (*MatchList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MatchList)
	err := c.cc.Invoke(ctx, MatchesService_GetNearbyMatches_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchesServiceClient) GetMatchByFaceitId(ctx context.Context, in *GetMatchRequest, opts ...grpc.CallOption) (*Match, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Match)
	err := c.cc.Invoke(ctx, MatchesService_GetMatchByFaceitId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchesServiceClient) SetStreamToMatch(ctx context.Context, in *SetStreamRequest, opts ...grpc.CallOption) (*Bool, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Bool)
	err := c.cc.Invoke(ctx, MatchesService_SetStreamToMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchesServiceClient) GetMatchesByTeamId(ctx context.Context, in *GetMatchRequest, opts ...grpc.CallOption) (*MatchList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MatchList)
	err := c.cc.Invoke(ctx, MatchesService_GetMatchesByTeamId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchesServiceClient) NewMatch(ctx context.Context, in *NewMatchRequest, opts ...grpc.CallOption) (*Match, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Match)
	err := c.cc.Invoke(ctx, MatchesService_NewMatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchesServiceServer is the server API for MatchesService service.
// All implementations must embed UnimplementedMatchesServiceServer
// for forward compatibility.
type MatchesServiceServer interface {
	GetAllMatches(context.Context, *Empty) (*MatchList, error)
	GetNearbyMatches(context.Context, *GetNearbyMatchesRequest) (*MatchList, error)
	GetMatchByFaceitId(context.Context, *GetMatchRequest) (*Match, error)
	SetStreamToMatch(context.Context, *SetStreamRequest) (*Bool, error)
	GetMatchesByTeamId(context.Context, *GetMatchRequest) (*MatchList, error)
	NewMatch(context.Context, *NewMatchRequest) (*Match, error)
	mustEmbedUnimplementedMatchesServiceServer()
}

// UnimplementedMatchesServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMatchesServiceServer struct{}

func (UnimplementedMatchesServiceServer) GetAllMatches(context.Context, *Empty) (*MatchList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMatches not implemented")
}
func (UnimplementedMatchesServiceServer) GetNearbyMatches(context.Context, *GetNearbyMatchesRequest) (*MatchList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNearbyMatches not implemented")
}
func (UnimplementedMatchesServiceServer) GetMatchByFaceitId(context.Context, *GetMatchRequest) (*Match, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatchByFaceitId not implemented")
}
func (UnimplementedMatchesServiceServer) SetStreamToMatch(context.Context, *SetStreamRequest) (*Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStreamToMatch not implemented")
}
func (UnimplementedMatchesServiceServer) GetMatchesByTeamId(context.Context, *GetMatchRequest) (*MatchList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatchesByTeamId not implemented")
}
func (UnimplementedMatchesServiceServer) NewMatch(context.Context, *NewMatchRequest) (*Match, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewMatch not implemented")
}
func (UnimplementedMatchesServiceServer) mustEmbedUnimplementedMatchesServiceServer() {}
func (UnimplementedMatchesServiceServer) testEmbeddedByValue()                        {}

// UnsafeMatchesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchesServiceServer will
// result in compilation errors.
type UnsafeMatchesServiceServer interface {
	mustEmbedUnimplementedMatchesServiceServer()
}

func RegisterMatchesServiceServer(s grpc.ServiceRegistrar, srv MatchesServiceServer) {
	// If the following call pancis, it indicates UnimplementedMatchesServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MatchesService_ServiceDesc, srv)
}

func _MatchesService_GetAllMatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchesServiceServer).GetAllMatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchesService_GetAllMatches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchesServiceServer).GetAllMatches(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchesService_GetNearbyMatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNearbyMatchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchesServiceServer).GetNearbyMatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchesService_GetNearbyMatches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchesServiceServer).GetNearbyMatches(ctx, req.(*GetNearbyMatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchesService_GetMatchByFaceitId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchesServiceServer).GetMatchByFaceitId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchesService_GetMatchByFaceitId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchesServiceServer).GetMatchByFaceitId(ctx, req.(*GetMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchesService_SetStreamToMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchesServiceServer).SetStreamToMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchesService_SetStreamToMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchesServiceServer).SetStreamToMatch(ctx, req.(*SetStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchesService_GetMatchesByTeamId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchesServiceServer).GetMatchesByTeamId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchesService_GetMatchesByTeamId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchesServiceServer).GetMatchesByTeamId(ctx, req.(*GetMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchesService_NewMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchesServiceServer).NewMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MatchesService_NewMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchesServiceServer).NewMatch(ctx, req.(*NewMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchesService_ServiceDesc is the grpc.ServiceDesc for MatchesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "matches.MatchesService",
	HandlerType: (*MatchesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllMatches",
			Handler:    _MatchesService_GetAllMatches_Handler,
		},
		{
			MethodName: "GetNearbyMatches",
			Handler:    _MatchesService_GetNearbyMatches_Handler,
		},
		{
			MethodName: "GetMatchByFaceitId",
			Handler:    _MatchesService_GetMatchByFaceitId_Handler,
		},
		{
			MethodName: "SetStreamToMatch",
			Handler:    _MatchesService_SetStreamToMatch_Handler,
		},
		{
			MethodName: "GetMatchesByTeamId",
			Handler:    _MatchesService_GetMatchesByTeamId_Handler,
		},
		{
			MethodName: "NewMatch",
			Handler:    _MatchesService_NewMatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/matches.proto",
}
