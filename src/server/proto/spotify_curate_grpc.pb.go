// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: spotify_curate.proto

package proto

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
	SpotifyCurate_UpdateUserData_FullMethodName                 = "/proto.SpotifyCurate/UpdateUserData"
	SpotifyCurate_CreatePlaylistRecentInGenre_FullMethodName    = "/proto.SpotifyCurate/CreatePlaylistRecentInGenre"
	SpotifyCurate_CreatePlaylistRecentInGenreAll_FullMethodName = "/proto.SpotifyCurate/CreatePlaylistRecentInGenreAll"
	SpotifyCurate_GetUnmappedArtistsForUser_FullMethodName      = "/proto.SpotifyCurate/GetUnmappedArtistsForUser"
)

// SpotifyCurateClient is the client API for SpotifyCurate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpotifyCurateClient interface {
	UpdateUserData(ctx context.Context, in *UpdateUserDataRequest, opts ...grpc.CallOption) (*UpdateUserDataResponse, error)
	CreatePlaylistRecentInGenre(ctx context.Context, in *CreatePlaylistRecentInGenreRequest, opts ...grpc.CallOption) (*CreatePlaylistRecentInGenreResponse, error)
	CreatePlaylistRecentInGenreAll(ctx context.Context, in *CreatePlaylistRecentInGenreAllRequest, opts ...grpc.CallOption) (*CreatePlaylistRecentInGenreAllResponse, error)
	GetUnmappedArtistsForUser(ctx context.Context, in *GetUnmappedArtistsForUserRequest, opts ...grpc.CallOption) (*GetUnmappedArtistsForUserResponse, error)
}

type spotifyCurateClient struct {
	cc grpc.ClientConnInterface
}

func NewSpotifyCurateClient(cc grpc.ClientConnInterface) SpotifyCurateClient {
	return &spotifyCurateClient{cc}
}

func (c *spotifyCurateClient) UpdateUserData(ctx context.Context, in *UpdateUserDataRequest, opts ...grpc.CallOption) (*UpdateUserDataResponse, error) {
	out := new(UpdateUserDataResponse)
	err := c.cc.Invoke(ctx, SpotifyCurate_UpdateUserData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotifyCurateClient) CreatePlaylistRecentInGenre(ctx context.Context, in *CreatePlaylistRecentInGenreRequest, opts ...grpc.CallOption) (*CreatePlaylistRecentInGenreResponse, error) {
	out := new(CreatePlaylistRecentInGenreResponse)
	err := c.cc.Invoke(ctx, SpotifyCurate_CreatePlaylistRecentInGenre_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotifyCurateClient) CreatePlaylistRecentInGenreAll(ctx context.Context, in *CreatePlaylistRecentInGenreAllRequest, opts ...grpc.CallOption) (*CreatePlaylistRecentInGenreAllResponse, error) {
	out := new(CreatePlaylistRecentInGenreAllResponse)
	err := c.cc.Invoke(ctx, SpotifyCurate_CreatePlaylistRecentInGenreAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotifyCurateClient) GetUnmappedArtistsForUser(ctx context.Context, in *GetUnmappedArtistsForUserRequest, opts ...grpc.CallOption) (*GetUnmappedArtistsForUserResponse, error) {
	out := new(GetUnmappedArtistsForUserResponse)
	err := c.cc.Invoke(ctx, SpotifyCurate_GetUnmappedArtistsForUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpotifyCurateServer is the server API for SpotifyCurate service.
// All implementations must embed UnimplementedSpotifyCurateServer
// for forward compatibility
type SpotifyCurateServer interface {
	UpdateUserData(context.Context, *UpdateUserDataRequest) (*UpdateUserDataResponse, error)
	CreatePlaylistRecentInGenre(context.Context, *CreatePlaylistRecentInGenreRequest) (*CreatePlaylistRecentInGenreResponse, error)
	CreatePlaylistRecentInGenreAll(context.Context, *CreatePlaylistRecentInGenreAllRequest) (*CreatePlaylistRecentInGenreAllResponse, error)
	GetUnmappedArtistsForUser(context.Context, *GetUnmappedArtistsForUserRequest) (*GetUnmappedArtistsForUserResponse, error)
	mustEmbedUnimplementedSpotifyCurateServer()
}

// UnimplementedSpotifyCurateServer must be embedded to have forward compatible implementations.
type UnimplementedSpotifyCurateServer struct {
}

func (UnimplementedSpotifyCurateServer) UpdateUserData(context.Context, *UpdateUserDataRequest) (*UpdateUserDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserData not implemented")
}
func (UnimplementedSpotifyCurateServer) CreatePlaylistRecentInGenre(context.Context, *CreatePlaylistRecentInGenreRequest) (*CreatePlaylistRecentInGenreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlaylistRecentInGenre not implemented")
}
func (UnimplementedSpotifyCurateServer) CreatePlaylistRecentInGenreAll(context.Context, *CreatePlaylistRecentInGenreAllRequest) (*CreatePlaylistRecentInGenreAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlaylistRecentInGenreAll not implemented")
}
func (UnimplementedSpotifyCurateServer) GetUnmappedArtistsForUser(context.Context, *GetUnmappedArtistsForUserRequest) (*GetUnmappedArtistsForUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnmappedArtistsForUser not implemented")
}
func (UnimplementedSpotifyCurateServer) mustEmbedUnimplementedSpotifyCurateServer() {}

// UnsafeSpotifyCurateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpotifyCurateServer will
// result in compilation errors.
type UnsafeSpotifyCurateServer interface {
	mustEmbedUnimplementedSpotifyCurateServer()
}

func RegisterSpotifyCurateServer(s grpc.ServiceRegistrar, srv SpotifyCurateServer) {
	s.RegisterService(&SpotifyCurate_ServiceDesc, srv)
}

func _SpotifyCurate_UpdateUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyCurateServer).UpdateUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpotifyCurate_UpdateUserData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyCurateServer).UpdateUserData(ctx, req.(*UpdateUserDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotifyCurate_CreatePlaylistRecentInGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlaylistRecentInGenreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyCurateServer).CreatePlaylistRecentInGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpotifyCurate_CreatePlaylistRecentInGenre_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyCurateServer).CreatePlaylistRecentInGenre(ctx, req.(*CreatePlaylistRecentInGenreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotifyCurate_CreatePlaylistRecentInGenreAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlaylistRecentInGenreAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyCurateServer).CreatePlaylistRecentInGenreAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpotifyCurate_CreatePlaylistRecentInGenreAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyCurateServer).CreatePlaylistRecentInGenreAll(ctx, req.(*CreatePlaylistRecentInGenreAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotifyCurate_GetUnmappedArtistsForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnmappedArtistsForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyCurateServer).GetUnmappedArtistsForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SpotifyCurate_GetUnmappedArtistsForUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyCurateServer).GetUnmappedArtistsForUser(ctx, req.(*GetUnmappedArtistsForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpotifyCurate_ServiceDesc is the grpc.ServiceDesc for SpotifyCurate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpotifyCurate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SpotifyCurate",
	HandlerType: (*SpotifyCurateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUserData",
			Handler:    _SpotifyCurate_UpdateUserData_Handler,
		},
		{
			MethodName: "CreatePlaylistRecentInGenre",
			Handler:    _SpotifyCurate_CreatePlaylistRecentInGenre_Handler,
		},
		{
			MethodName: "CreatePlaylistRecentInGenreAll",
			Handler:    _SpotifyCurate_CreatePlaylistRecentInGenreAll_Handler,
		},
		{
			MethodName: "GetUnmappedArtistsForUser",
			Handler:    _SpotifyCurate_GetUnmappedArtistsForUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spotify_curate.proto",
}
