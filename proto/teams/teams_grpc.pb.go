// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/teams.proto

package teams

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
	TeamService_GetTeams_FullMethodName          = "/teams.TeamService/GetTeams"
	TeamService_GetTeamById_FullMethodName       = "/teams.TeamService/GetTeamById"
	TeamService_GetTeamByNickname_FullMethodName = "/teams.TeamService/GetTeamByNickname"
	TeamService_NewTeam_FullMethodName           = "/teams.TeamService/NewTeam"
)

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamServiceClient interface {
	GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...grpc.CallOption) (*TeamList, error)
	GetTeamById(ctx context.Context, in *NewTeamRequest, opts ...grpc.CallOption) (*Team, error)
	GetTeamByNickname(ctx context.Context, in *NewTeamRequest, opts ...grpc.CallOption) (*Team, error)
	NewTeam(ctx context.Context, in *NewTeamRequest, opts ...grpc.CallOption) (*Team, error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...grpc.CallOption) (*TeamList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TeamList)
	err := c.cc.Invoke(ctx, TeamService_GetTeams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamById(ctx context.Context, in *NewTeamRequest, opts ...grpc.CallOption) (*Team, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_GetTeamById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamByNickname(ctx context.Context, in *NewTeamRequest, opts ...grpc.CallOption) (*Team, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_GetTeamByNickname_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) NewTeam(ctx context.Context, in *NewTeamRequest, opts ...grpc.CallOption) (*Team, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_NewTeam_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServiceServer is the server API for TeamService service.
// All implementations must embed UnimplementedTeamServiceServer
// for forward compatibility.
type TeamServiceServer interface {
	GetTeams(context.Context, *GetTeamsRequest) (*TeamList, error)
	GetTeamById(context.Context, *NewTeamRequest) (*Team, error)
	GetTeamByNickname(context.Context, *NewTeamRequest) (*Team, error)
	NewTeam(context.Context, *NewTeamRequest) (*Team, error)
	mustEmbedUnimplementedTeamServiceServer()
}

// UnimplementedTeamServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTeamServiceServer struct{}

func (UnimplementedTeamServiceServer) GetTeams(context.Context, *GetTeamsRequest) (*TeamList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeams not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamById(context.Context, *NewTeamRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamById not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamByNickname(context.Context, *NewTeamRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamByNickname not implemented")
}
func (UnimplementedTeamServiceServer) NewTeam(context.Context, *NewTeamRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewTeam not implemented")
}
func (UnimplementedTeamServiceServer) mustEmbedUnimplementedTeamServiceServer() {}
func (UnimplementedTeamServiceServer) testEmbeddedByValue()                     {}

// UnsafeTeamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamServiceServer will
// result in compilation errors.
type UnsafeTeamServiceServer interface {
	mustEmbedUnimplementedTeamServiceServer()
}

func RegisterTeamServiceServer(s grpc.ServiceRegistrar, srv TeamServiceServer) {
	// If the following call pancis, it indicates UnimplementedTeamServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TeamService_ServiceDesc, srv)
}

func _TeamService_GetTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeams(ctx, req.(*GetTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamById(ctx, req.(*NewTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamByNickname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamByNickname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamByNickname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamByNickname(ctx, req.(*NewTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_NewTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).NewTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_NewTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).NewTeam(ctx, req.(*NewTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeamService_ServiceDesc is the grpc.ServiceDesc for TeamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teams.TeamService",
	HandlerType: (*TeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeams",
			Handler:    _TeamService_GetTeams_Handler,
		},
		{
			MethodName: "GetTeamById",
			Handler:    _TeamService_GetTeamById_Handler,
		},
		{
			MethodName: "GetTeamByNickname",
			Handler:    _TeamService_GetTeamByNickname_Handler,
		},
		{
			MethodName: "NewTeam",
			Handler:    _TeamService_NewTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/teams.proto",
}