// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// GroupClient is the client API for Group service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupClient interface {
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error)
	ListGroup(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupResponse, error)
	AddGroupItems(ctx context.Context, in *AddGroupItemsRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	ListGroupItems(ctx context.Context, in *ListGroupItemsRequest, opts ...grpc.CallOption) (*ListGroupItemsResponse, error)
	DelGroupItems(ctx context.Context, in *DelGroupItemsRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	AddGroupExt(ctx context.Context, in *AddGroupExtRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdateGroupExt(ctx context.Context, in *UpdateGroupExtRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	DelGroupExt(ctx context.Context, in *DelGroupExtRequest, opts ...grpc.CallOption) (*CommonResponse, error)
}

type groupClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupClient(cc grpc.ClientConnInterface) GroupClient {
	return &groupClient{cc}
}

func (c *groupClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/api.group.v1.Group/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/api.group.v1.Group/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/api.group.v1.Group/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error) {
	out := new(GetGroupResponse)
	err := c.cc.Invoke(ctx, "/api.group.v1.Group/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) ListGroup(ctx context.Context, in *ListGroupRequest, opts ...grpc.CallOption) (*ListGroupResponse, error) {
	out := new(ListGroupResponse)
	err := c.cc.Invoke(ctx, "/api.group.v1.Group/ListGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) AddGroupItems(ctx context.Context, in *AddGroupItemsRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/api.group.v1.Group/AddGroupItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) ListGroupItems(ctx context.Context, in *ListGroupItemsRequest, opts ...grpc.CallOption) (*ListGroupItemsResponse, error) {
	out := new(ListGroupItemsResponse)
	err := c.cc.Invoke(ctx, "/api.group.v1.Group/ListGroupItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) DelGroupItems(ctx context.Context, in *DelGroupItemsRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/api.group.v1.Group/DelGroupItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) AddGroupExt(ctx context.Context, in *AddGroupExtRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/api.group.v1.Group/AddGroupExt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) UpdateGroupExt(ctx context.Context, in *UpdateGroupExtRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/api.group.v1.Group/UpdateGroupExt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) DelGroupExt(ctx context.Context, in *DelGroupExtRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/api.group.v1.Group/DelGroupExt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServer is the server API for Group service.
// All implementations must embed UnimplementedGroupServer
// for forward compatibility
type GroupServer interface {
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	UpdateGroup(context.Context, *UpdateGroupRequest) (*CommonResponse, error)
	DeleteGroup(context.Context, *DeleteGroupRequest) (*CommonResponse, error)
	GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error)
	ListGroup(context.Context, *ListGroupRequest) (*ListGroupResponse, error)
	AddGroupItems(context.Context, *AddGroupItemsRequest) (*CommonResponse, error)
	ListGroupItems(context.Context, *ListGroupItemsRequest) (*ListGroupItemsResponse, error)
	DelGroupItems(context.Context, *DelGroupItemsRequest) (*CommonResponse, error)
	AddGroupExt(context.Context, *AddGroupExtRequest) (*CommonResponse, error)
	UpdateGroupExt(context.Context, *UpdateGroupExtRequest) (*CommonResponse, error)
	DelGroupExt(context.Context, *DelGroupExtRequest) (*CommonResponse, error)
	mustEmbedUnimplementedGroupServer()
}

// UnimplementedGroupServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServer struct {
}

func (UnimplementedGroupServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedGroupServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedGroupServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedGroupServer) GetGroup(context.Context, *GetGroupRequest) (*GetGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedGroupServer) ListGroup(context.Context, *ListGroupRequest) (*ListGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroup not implemented")
}
func (UnimplementedGroupServer) AddGroupItems(context.Context, *AddGroupItemsRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroupItems not implemented")
}
func (UnimplementedGroupServer) ListGroupItems(context.Context, *ListGroupItemsRequest) (*ListGroupItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroupItems not implemented")
}
func (UnimplementedGroupServer) DelGroupItems(context.Context, *DelGroupItemsRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelGroupItems not implemented")
}
func (UnimplementedGroupServer) AddGroupExt(context.Context, *AddGroupExtRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroupExt not implemented")
}
func (UnimplementedGroupServer) UpdateGroupExt(context.Context, *UpdateGroupExtRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupExt not implemented")
}
func (UnimplementedGroupServer) DelGroupExt(context.Context, *DelGroupExtRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelGroupExt not implemented")
}
func (UnimplementedGroupServer) mustEmbedUnimplementedGroupServer() {}

// UnsafeGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServer will
// result in compilation errors.
type UnsafeGroupServer interface {
	mustEmbedUnimplementedGroupServer()
}

func RegisterGroupServer(s grpc.ServiceRegistrar, srv GroupServer) {
	s.RegisterService(&Group_ServiceDesc, srv)
}

func _Group_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.group.v1.Group/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.group.v1.Group/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.group.v1.Group/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.group.v1.Group/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_ListGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).ListGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.group.v1.Group/ListGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).ListGroup(ctx, req.(*ListGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_AddGroupItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).AddGroupItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.group.v1.Group/AddGroupItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).AddGroupItems(ctx, req.(*AddGroupItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_ListGroupItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).ListGroupItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.group.v1.Group/ListGroupItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).ListGroupItems(ctx, req.(*ListGroupItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_DelGroupItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelGroupItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).DelGroupItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.group.v1.Group/DelGroupItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).DelGroupItems(ctx, req.(*DelGroupItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_AddGroupExt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupExtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).AddGroupExt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.group.v1.Group/AddGroupExt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).AddGroupExt(ctx, req.(*AddGroupExtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_UpdateGroupExt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupExtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).UpdateGroupExt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.group.v1.Group/UpdateGroupExt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).UpdateGroupExt(ctx, req.(*UpdateGroupExtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_DelGroupExt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelGroupExtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).DelGroupExt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.group.v1.Group/DelGroupExt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).DelGroupExt(ctx, req.(*DelGroupExtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Group_ServiceDesc is the grpc.ServiceDesc for Group service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Group_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.group.v1.Group",
	HandlerType: (*GroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _Group_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _Group_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _Group_DeleteGroup_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _Group_GetGroup_Handler,
		},
		{
			MethodName: "ListGroup",
			Handler:    _Group_ListGroup_Handler,
		},
		{
			MethodName: "AddGroupItems",
			Handler:    _Group_AddGroupItems_Handler,
		},
		{
			MethodName: "ListGroupItems",
			Handler:    _Group_ListGroupItems_Handler,
		},
		{
			MethodName: "DelGroupItems",
			Handler:    _Group_DelGroupItems_Handler,
		},
		{
			MethodName: "AddGroupExt",
			Handler:    _Group_AddGroupExt_Handler,
		},
		{
			MethodName: "UpdateGroupExt",
			Handler:    _Group_UpdateGroupExt_Handler,
		},
		{
			MethodName: "DelGroupExt",
			Handler:    _Group_DelGroupExt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/group/v1/group.proto",
}
