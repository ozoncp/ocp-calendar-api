// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_calendar_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OcpCalendarApiClient is the client API for OcpCalendarApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpCalendarApiClient interface {
	CreateCalendarV1(ctx context.Context, in *CreateCalendarRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DescribeCalendarV1(ctx context.Context, in *DescribeCalendarRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListCalendarsV1(ctx context.Context, in *ListCalendarRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveCalendarV1(ctx context.Context, in *RemoveCalendarRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type ocpCalendarApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpCalendarApiClient(cc grpc.ClientConnInterface) OcpCalendarApiClient {
	return &ocpCalendarApiClient{cc}
}

func (c *ocpCalendarApiClient) CreateCalendarV1(ctx context.Context, in *CreateCalendarRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ocp_calendar_api.OcpCalendarApi/CreateCalendarV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpCalendarApiClient) DescribeCalendarV1(ctx context.Context, in *DescribeCalendarRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ocp_calendar_api.OcpCalendarApi/DescribeCalendarV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpCalendarApiClient) ListCalendarsV1(ctx context.Context, in *ListCalendarRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ocp_calendar_api.OcpCalendarApi/ListCalendarsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpCalendarApiClient) RemoveCalendarV1(ctx context.Context, in *RemoveCalendarRequestV1, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ocp_calendar_api.OcpCalendarApi/RemoveCalendarV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpCalendarApiServer is the server API for OcpCalendarApi service.
// All implementations must embed UnimplementedOcpCalendarApiServer
// for forward compatibility
type OcpCalendarApiServer interface {
	CreateCalendarV1(context.Context, *CreateCalendarRequestV1) (*emptypb.Empty, error)
	DescribeCalendarV1(context.Context, *DescribeCalendarRequestV1) (*emptypb.Empty, error)
	ListCalendarsV1(context.Context, *ListCalendarRequestV1) (*emptypb.Empty, error)
	RemoveCalendarV1(context.Context, *RemoveCalendarRequestV1) (*emptypb.Empty, error)
	mustEmbedUnimplementedOcpCalendarApiServer()
}

// UnimplementedOcpCalendarApiServer must be embedded to have forward compatible implementations.
type UnimplementedOcpCalendarApiServer struct {
}

func (UnimplementedOcpCalendarApiServer) CreateCalendarV1(context.Context, *CreateCalendarRequestV1) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCalendarV1 not implemented")
}
func (UnimplementedOcpCalendarApiServer) DescribeCalendarV1(context.Context, *DescribeCalendarRequestV1) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCalendarV1 not implemented")
}
func (UnimplementedOcpCalendarApiServer) ListCalendarsV1(context.Context, *ListCalendarRequestV1) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCalendarsV1 not implemented")
}
func (UnimplementedOcpCalendarApiServer) RemoveCalendarV1(context.Context, *RemoveCalendarRequestV1) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCalendarV1 not implemented")
}
func (UnimplementedOcpCalendarApiServer) mustEmbedUnimplementedOcpCalendarApiServer() {}

// UnsafeOcpCalendarApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpCalendarApiServer will
// result in compilation errors.
type UnsafeOcpCalendarApiServer interface {
	mustEmbedUnimplementedOcpCalendarApiServer()
}

func RegisterOcpCalendarApiServer(s grpc.ServiceRegistrar, srv OcpCalendarApiServer) {
	s.RegisterService(&OcpCalendarApi_ServiceDesc, srv)
}

func _OcpCalendarApi_CreateCalendarV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCalendarRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpCalendarApiServer).CreateCalendarV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp_calendar_api.OcpCalendarApi/CreateCalendarV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpCalendarApiServer).CreateCalendarV1(ctx, req.(*CreateCalendarRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpCalendarApi_DescribeCalendarV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCalendarRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpCalendarApiServer).DescribeCalendarV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp_calendar_api.OcpCalendarApi/DescribeCalendarV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpCalendarApiServer).DescribeCalendarV1(ctx, req.(*DescribeCalendarRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpCalendarApi_ListCalendarsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCalendarRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpCalendarApiServer).ListCalendarsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp_calendar_api.OcpCalendarApi/ListCalendarsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpCalendarApiServer).ListCalendarsV1(ctx, req.(*ListCalendarRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpCalendarApi_RemoveCalendarV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCalendarRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpCalendarApiServer).RemoveCalendarV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp_calendar_api.OcpCalendarApi/RemoveCalendarV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpCalendarApiServer).RemoveCalendarV1(ctx, req.(*RemoveCalendarRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpCalendarApi_ServiceDesc is the grpc.ServiceDesc for OcpCalendarApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpCalendarApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp_calendar_api.OcpCalendarApi",
	HandlerType: (*OcpCalendarApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCalendarV1",
			Handler:    _OcpCalendarApi_CreateCalendarV1_Handler,
		},
		{
			MethodName: "DescribeCalendarV1",
			Handler:    _OcpCalendarApi_DescribeCalendarV1_Handler,
		},
		{
			MethodName: "ListCalendarsV1",
			Handler:    _OcpCalendarApi_ListCalendarsV1_Handler,
		},
		{
			MethodName: "RemoveCalendarV1",
			Handler:    _OcpCalendarApi_RemoveCalendarV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ocp-calendar-api/ocp-calendar-api.proto",
}
