// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api/did/services/connection.proto

package services

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

// ConnectionControllerClient is the client API for ConnectionController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionControllerClient interface {
	// rotateDID swagger:route POST /connections/{id}/rotate-did connections rotateDID
	//
	// Rotates the agent's DID in the given connection.
	//
	// Responses:
	//        200: rotateDIDResponse
	RotateDID(ctx context.Context, in *RotateDIDRequest, opts ...grpc.CallOption) (*RotateDIDResponse, error)
	// createConnectionV2 swagger:route POST /connections/create-v2 connections createConnectionV2
	//
	// Creates a DIDComm v2 connection record with the given DIDs.
	//
	// Responses:
	//        200: createConnectionV2Response
	CreateConnectionV2(ctx context.Context, in *CreateConnectionRequestV2, opts ...grpc.CallOption) (*CreateConnectionV2Response, error)
	// setConnectionToV2 swagger:route POST /connections/{id}/use-v2 connections setConnectionToV2
	// SetConnectionToDIDCommV2 sets the didcomm version of the given connection to V2.
	// 200	Succeeded
	// 400	Bad Request
	// 401	Not Authorized
	// 410	Gone! There is no data here
	// 500	Internal Error
	// 501	Not Implemented
	SetConnectionToDIDCommV2(ctx context.Context, in *IDMessage, opts ...grpc.CallOption) (*IDMessage, error)
}

type connectionControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionControllerClient(cc grpc.ClientConnInterface) ConnectionControllerClient {
	return &connectionControllerClient{cc}
}

func (c *connectionControllerClient) RotateDID(ctx context.Context, in *RotateDIDRequest, opts ...grpc.CallOption) (*RotateDIDResponse, error) {
	out := new(RotateDIDResponse)
	err := c.cc.Invoke(ctx, "/api.did.services.ConnectionController/rotateDID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionControllerClient) CreateConnectionV2(ctx context.Context, in *CreateConnectionRequestV2, opts ...grpc.CallOption) (*CreateConnectionV2Response, error) {
	out := new(CreateConnectionV2Response)
	err := c.cc.Invoke(ctx, "/api.did.services.ConnectionController/createConnectionV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionControllerClient) SetConnectionToDIDCommV2(ctx context.Context, in *IDMessage, opts ...grpc.CallOption) (*IDMessage, error) {
	out := new(IDMessage)
	err := c.cc.Invoke(ctx, "/api.did.services.ConnectionController/setConnectionToDIDCommV2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionControllerServer is the server API for ConnectionController service.
// All implementations must embed UnimplementedConnectionControllerServer
// for forward compatibility
type ConnectionControllerServer interface {
	// rotateDID swagger:route POST /connections/{id}/rotate-did connections rotateDID
	//
	// Rotates the agent's DID in the given connection.
	//
	// Responses:
	//        200: rotateDIDResponse
	RotateDID(context.Context, *RotateDIDRequest) (*RotateDIDResponse, error)
	// createConnectionV2 swagger:route POST /connections/create-v2 connections createConnectionV2
	//
	// Creates a DIDComm v2 connection record with the given DIDs.
	//
	// Responses:
	//        200: createConnectionV2Response
	CreateConnectionV2(context.Context, *CreateConnectionRequestV2) (*CreateConnectionV2Response, error)
	// setConnectionToV2 swagger:route POST /connections/{id}/use-v2 connections setConnectionToV2
	// SetConnectionToDIDCommV2 sets the didcomm version of the given connection to V2.
	// 200	Succeeded
	// 400	Bad Request
	// 401	Not Authorized
	// 410	Gone! There is no data here
	// 500	Internal Error
	// 501	Not Implemented
	SetConnectionToDIDCommV2(context.Context, *IDMessage) (*IDMessage, error)
	mustEmbedUnimplementedConnectionControllerServer()
}

// UnimplementedConnectionControllerServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionControllerServer struct {
}

func (UnimplementedConnectionControllerServer) RotateDID(context.Context, *RotateDIDRequest) (*RotateDIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateDID not implemented")
}
func (UnimplementedConnectionControllerServer) CreateConnectionV2(context.Context, *CreateConnectionRequestV2) (*CreateConnectionV2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnectionV2 not implemented")
}
func (UnimplementedConnectionControllerServer) SetConnectionToDIDCommV2(context.Context, *IDMessage) (*IDMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConnectionToDIDCommV2 not implemented")
}
func (UnimplementedConnectionControllerServer) mustEmbedUnimplementedConnectionControllerServer() {}

// UnsafeConnectionControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionControllerServer will
// result in compilation errors.
type UnsafeConnectionControllerServer interface {
	mustEmbedUnimplementedConnectionControllerServer()
}

func RegisterConnectionControllerServer(s grpc.ServiceRegistrar, srv ConnectionControllerServer) {
	s.RegisterService(&ConnectionController_ServiceDesc, srv)
}

func _ConnectionController_RotateDID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RotateDIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionControllerServer).RotateDID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.did.services.ConnectionController/rotateDID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionControllerServer).RotateDID(ctx, req.(*RotateDIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionController_CreateConnectionV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectionRequestV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionControllerServer).CreateConnectionV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.did.services.ConnectionController/createConnectionV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionControllerServer).CreateConnectionV2(ctx, req.(*CreateConnectionRequestV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionController_SetConnectionToDIDCommV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionControllerServer).SetConnectionToDIDCommV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.did.services.ConnectionController/setConnectionToDIDCommV2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionControllerServer).SetConnectionToDIDCommV2(ctx, req.(*IDMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectionController_ServiceDesc is the grpc.ServiceDesc for ConnectionController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.did.services.ConnectionController",
	HandlerType: (*ConnectionControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "rotateDID",
			Handler:    _ConnectionController_RotateDID_Handler,
		},
		{
			MethodName: "createConnectionV2",
			Handler:    _ConnectionController_CreateConnectionV2_Handler,
		},
		{
			MethodName: "setConnectionToDIDCommV2",
			Handler:    _ConnectionController_SetConnectionToDIDCommV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/did/services/connection.proto",
}
