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
	TeamService_GetAllTeams_FullMethodName         = "/teams.TeamService/GetAllTeams"
	TeamService_GetActiveTeams_FullMethodName      = "/teams.TeamService/GetActiveTeams"
	TeamService_GetById_FullMethodName             = "/teams.TeamService/GetById"
	TeamService_GetByNickname_FullMethodName       = "/teams.TeamService/GetByNickname"
	TeamService_CreateFromFaceit_FullMethodName    = "/teams.TeamService/CreateFromFaceit"
	TeamService_Update_FullMethodName              = "/teams.TeamService/Update"
	TeamService_FindTeamsByPlayerId_FullMethodName = "/teams.TeamService/FindTeamsByPlayerId"
	TeamService_GetTeamFromFaceit_FullMethodName   = "/teams.TeamService/GetTeamFromFaceit"
)

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamServiceClient interface {
	GetAllTeams(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TeamList, error)
	GetActiveTeams(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TeamList, error)
	GetById(ctx context.Context, in *GetTeamByIdRequest, opts ...grpc.CallOption) (*Team, error)
	GetByNickname(ctx context.Context, in *GetTeamByNicknameRequest, opts ...grpc.CallOption) (*Team, error)
	CreateFromFaceit(ctx context.Context, in *NewTeamFromFaceitRequest, opts ...grpc.CallOption) (*Team, error)
	Update(ctx context.Context, in *NewTeamFromFaceitRequest, opts ...grpc.CallOption) (*Team, error)
	FindTeamsByPlayerId(ctx context.Context, in *GetTeamByPlayerIdRequest, opts ...grpc.CallOption) (*TeamList, error)
	GetTeamFromFaceit(ctx context.Context, in *GetTeamFromFaceitRequest, opts ...grpc.CallOption) (*Team, error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) GetAllTeams(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TeamList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TeamList)
	err := c.cc.Invoke(ctx, TeamService_GetAllTeams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetActiveTeams(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TeamList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TeamList)
	err := c.cc.Invoke(ctx, TeamService_GetActiveTeams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetById(ctx context.Context, in *GetTeamByIdRequest, opts ...grpc.CallOption) (*Team, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetByNickname(ctx context.Context, in *GetTeamByNicknameRequest, opts ...grpc.CallOption) (*Team, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_GetByNickname_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) CreateFromFaceit(ctx context.Context, in *NewTeamFromFaceitRequest, opts ...grpc.CallOption) (*Team, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_CreateFromFaceit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) Update(ctx context.Context, in *NewTeamFromFaceitRequest, opts ...grpc.CallOption) (*Team, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) FindTeamsByPlayerId(ctx context.Context, in *GetTeamByPlayerIdRequest, opts ...grpc.CallOption) (*TeamList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TeamList)
	err := c.cc.Invoke(ctx, TeamService_FindTeamsByPlayerId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetTeamFromFaceit(ctx context.Context, in *GetTeamFromFaceitRequest, opts ...grpc.CallOption) (*Team, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Team)
	err := c.cc.Invoke(ctx, TeamService_GetTeamFromFaceit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServiceServer is the server API for TeamService service.
// All implementations must embed UnimplementedTeamServiceServer
// for forward compatibility.
type TeamServiceServer interface {
	GetAllTeams(context.Context, *Empty) (*TeamList, error)
	GetActiveTeams(context.Context, *Empty) (*TeamList, error)
	GetById(context.Context, *GetTeamByIdRequest) (*Team, error)
	GetByNickname(context.Context, *GetTeamByNicknameRequest) (*Team, error)
	CreateFromFaceit(context.Context, *NewTeamFromFaceitRequest) (*Team, error)
	Update(context.Context, *NewTeamFromFaceitRequest) (*Team, error)
	FindTeamsByPlayerId(context.Context, *GetTeamByPlayerIdRequest) (*TeamList, error)
	GetTeamFromFaceit(context.Context, *GetTeamFromFaceitRequest) (*Team, error)
	mustEmbedUnimplementedTeamServiceServer()
}

// UnimplementedTeamServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTeamServiceServer struct{}

func (UnimplementedTeamServiceServer) GetAllTeams(context.Context, *Empty) (*TeamList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTeams not implemented")
}
func (UnimplementedTeamServiceServer) GetActiveTeams(context.Context, *Empty) (*TeamList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveTeams not implemented")
}
func (UnimplementedTeamServiceServer) GetById(context.Context, *GetTeamByIdRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedTeamServiceServer) GetByNickname(context.Context, *GetTeamByNicknameRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByNickname not implemented")
}
func (UnimplementedTeamServiceServer) CreateFromFaceit(context.Context, *NewTeamFromFaceitRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFromFaceit not implemented")
}
func (UnimplementedTeamServiceServer) Update(context.Context, *NewTeamFromFaceitRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTeamServiceServer) FindTeamsByPlayerId(context.Context, *GetTeamByPlayerIdRequest) (*TeamList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTeamsByPlayerId not implemented")
}
func (UnimplementedTeamServiceServer) GetTeamFromFaceit(context.Context, *GetTeamFromFaceitRequest) (*Team, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamFromFaceit not implemented")
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

func _TeamService_GetAllTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetAllTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetAllTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetAllTeams(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetActiveTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetActiveTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetActiveTeams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetActiveTeams(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetById(ctx, req.(*GetTeamByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetByNickname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamByNicknameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetByNickname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetByNickname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetByNickname(ctx, req.(*GetTeamByNicknameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_CreateFromFaceit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTeamFromFaceitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).CreateFromFaceit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_CreateFromFaceit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).CreateFromFaceit(ctx, req.(*NewTeamFromFaceitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTeamFromFaceitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).Update(ctx, req.(*NewTeamFromFaceitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_FindTeamsByPlayerId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamByPlayerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).FindTeamsByPlayerId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_FindTeamsByPlayerId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).FindTeamsByPlayerId(ctx, req.(*GetTeamByPlayerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetTeamFromFaceit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamFromFaceitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetTeamFromFaceit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TeamService_GetTeamFromFaceit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetTeamFromFaceit(ctx, req.(*GetTeamFromFaceitRequest))
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
			MethodName: "GetAllTeams",
			Handler:    _TeamService_GetAllTeams_Handler,
		},
		{
			MethodName: "GetActiveTeams",
			Handler:    _TeamService_GetActiveTeams_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _TeamService_GetById_Handler,
		},
		{
			MethodName: "GetByNickname",
			Handler:    _TeamService_GetByNickname_Handler,
		},
		{
			MethodName: "CreateFromFaceit",
			Handler:    _TeamService_CreateFromFaceit_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TeamService_Update_Handler,
		},
		{
			MethodName: "FindTeamsByPlayerId",
			Handler:    _TeamService_FindTeamsByPlayerId_Handler,
		},
		{
			MethodName: "GetTeamFromFaceit",
			Handler:    _TeamService_GetTeamFromFaceit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/teams.proto",
}
