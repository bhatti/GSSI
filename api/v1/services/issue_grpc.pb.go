// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api/v1/services/issue.proto

package services

import (
	context "context"
	types "github.com/bhatti/GSSI/api/v1/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CredentialsIssueServiceClient is the client API for CredentialsIssueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CredentialsIssueServiceClient interface {
	// Gets list of credentials or verifiable credentials
	// GET /credentials
	// 200	Credentials retrieved
	// 400	Bad Request
	// 401	Not Authorized
	// 410	Gone! There is no data here
	// 500	Internal Error
	// 501	Not Implemented
	GetAll(ctx context.Context, in *GetAllCredentialRequest, opts ...grpc.CallOption) (CredentialsIssueService_GetAllClient, error)
	// Gets list of credentials or verifiable credentials
	// GET /credentials/{id}
	// 200	Credential retrieved
	// 400	Bad Request
	// 401	Not Authorized
	// 410	Gone! There is no data here
	// 418	I'm a teapot - MUST not be returned outside of pre-arranged scenarios between both parties
	// 500	Internal Error
	// 501	Not Implemented
	Get(ctx context.Context, in *GetCredentialRequest, opts ...grpc.CallOption) (*types.RefreshableVerifiableCredential, error)
	// Deletes a credential or verifiable credential by ID
	// DELETE /credentials/{id}
	Delete(ctx context.Context, in *DeleteCredentialRequest, opts ...grpc.CallOption) (*types.Bool, error)
	// Issues a credential and returns it in the response body.
	// POST /credentials/issue
	// 201	Credential successfully issued!
	// 400	invalid input!
	// 500	error!
	Issue(ctx context.Context, in *IssueCredentialRequest, opts ...grpc.CallOption) (*types.RefreshableVerifiableCredential, error)
	// Updates the status of an issued credential
	// POST /credentials/status
	// 200	Credential status successfully updated
	// 400	Bad Request
	// 404	Credential not found
	// 500	Internal Server Error
	Update(ctx context.Context, in *UpdateCredentialRequest, opts ...grpc.CallOption) (*types.RefreshableVerifiableCredential, error)
}

type credentialsIssueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCredentialsIssueServiceClient(cc grpc.ClientConnInterface) CredentialsIssueServiceClient {
	return &credentialsIssueServiceClient{cc}
}

func (c *credentialsIssueServiceClient) GetAll(ctx context.Context, in *GetAllCredentialRequest, opts ...grpc.CallOption) (CredentialsIssueService_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &CredentialsIssueService_ServiceDesc.Streams[0], "/api.services.CredentialsIssueService/getAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &credentialsIssueServiceGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CredentialsIssueService_GetAllClient interface {
	Recv() (*types.RefreshableVerifiableCredential, error)
	grpc.ClientStream
}

type credentialsIssueServiceGetAllClient struct {
	grpc.ClientStream
}

func (x *credentialsIssueServiceGetAllClient) Recv() (*types.RefreshableVerifiableCredential, error) {
	m := new(types.RefreshableVerifiableCredential)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *credentialsIssueServiceClient) Get(ctx context.Context, in *GetCredentialRequest, opts ...grpc.CallOption) (*types.RefreshableVerifiableCredential, error) {
	out := new(types.RefreshableVerifiableCredential)
	err := c.cc.Invoke(ctx, "/api.services.CredentialsIssueService/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsIssueServiceClient) Delete(ctx context.Context, in *DeleteCredentialRequest, opts ...grpc.CallOption) (*types.Bool, error) {
	out := new(types.Bool)
	err := c.cc.Invoke(ctx, "/api.services.CredentialsIssueService/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsIssueServiceClient) Issue(ctx context.Context, in *IssueCredentialRequest, opts ...grpc.CallOption) (*types.RefreshableVerifiableCredential, error) {
	out := new(types.RefreshableVerifiableCredential)
	err := c.cc.Invoke(ctx, "/api.services.CredentialsIssueService/issue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsIssueServiceClient) Update(ctx context.Context, in *UpdateCredentialRequest, opts ...grpc.CallOption) (*types.RefreshableVerifiableCredential, error) {
	out := new(types.RefreshableVerifiableCredential)
	err := c.cc.Invoke(ctx, "/api.services.CredentialsIssueService/update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialsIssueServiceServer is the server API for CredentialsIssueService service.
// All implementations must embed UnimplementedCredentialsIssueServiceServer
// for forward compatibility
type CredentialsIssueServiceServer interface {
	// Gets list of credentials or verifiable credentials
	// GET /credentials
	// 200	Credentials retrieved
	// 400	Bad Request
	// 401	Not Authorized
	// 410	Gone! There is no data here
	// 500	Internal Error
	// 501	Not Implemented
	GetAll(*GetAllCredentialRequest, CredentialsIssueService_GetAllServer) error
	// Gets list of credentials or verifiable credentials
	// GET /credentials/{id}
	// 200	Credential retrieved
	// 400	Bad Request
	// 401	Not Authorized
	// 410	Gone! There is no data here
	// 418	I'm a teapot - MUST not be returned outside of pre-arranged scenarios between both parties
	// 500	Internal Error
	// 501	Not Implemented
	Get(context.Context, *GetCredentialRequest) (*types.RefreshableVerifiableCredential, error)
	// Deletes a credential or verifiable credential by ID
	// DELETE /credentials/{id}
	Delete(context.Context, *DeleteCredentialRequest) (*types.Bool, error)
	// Issues a credential and returns it in the response body.
	// POST /credentials/issue
	// 201	Credential successfully issued!
	// 400	invalid input!
	// 500	error!
	Issue(context.Context, *IssueCredentialRequest) (*types.RefreshableVerifiableCredential, error)
	// Updates the status of an issued credential
	// POST /credentials/status
	// 200	Credential status successfully updated
	// 400	Bad Request
	// 404	Credential not found
	// 500	Internal Server Error
	Update(context.Context, *UpdateCredentialRequest) (*types.RefreshableVerifiableCredential, error)
	mustEmbedUnimplementedCredentialsIssueServiceServer()
}

// UnimplementedCredentialsIssueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCredentialsIssueServiceServer struct {
}

func (UnimplementedCredentialsIssueServiceServer) GetAll(*GetAllCredentialRequest, CredentialsIssueService_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCredentialsIssueServiceServer) Get(context.Context, *GetCredentialRequest) (*types.RefreshableVerifiableCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCredentialsIssueServiceServer) Delete(context.Context, *DeleteCredentialRequest) (*types.Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCredentialsIssueServiceServer) Issue(context.Context, *IssueCredentialRequest) (*types.RefreshableVerifiableCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issue not implemented")
}
func (UnimplementedCredentialsIssueServiceServer) Update(context.Context, *UpdateCredentialRequest) (*types.RefreshableVerifiableCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCredentialsIssueServiceServer) mustEmbedUnimplementedCredentialsIssueServiceServer() {
}

// UnsafeCredentialsIssueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CredentialsIssueServiceServer will
// result in compilation errors.
type UnsafeCredentialsIssueServiceServer interface {
	mustEmbedUnimplementedCredentialsIssueServiceServer()
}

func RegisterCredentialsIssueServiceServer(s grpc.ServiceRegistrar, srv CredentialsIssueServiceServer) {
	s.RegisterService(&CredentialsIssueService_ServiceDesc, srv)
}

func _CredentialsIssueService_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllCredentialRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CredentialsIssueServiceServer).GetAll(m, &credentialsIssueServiceGetAllServer{stream})
}

type CredentialsIssueService_GetAllServer interface {
	Send(*types.RefreshableVerifiableCredential) error
	grpc.ServerStream
}

type credentialsIssueServiceGetAllServer struct {
	grpc.ServerStream
}

func (x *credentialsIssueServiceGetAllServer) Send(m *types.RefreshableVerifiableCredential) error {
	return x.ServerStream.SendMsg(m)
}

func _CredentialsIssueService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsIssueServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.CredentialsIssueService/get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsIssueServiceServer).Get(ctx, req.(*GetCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialsIssueService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsIssueServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.CredentialsIssueService/delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsIssueServiceServer).Delete(ctx, req.(*DeleteCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialsIssueService_Issue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsIssueServiceServer).Issue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.CredentialsIssueService/issue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsIssueServiceServer).Issue(ctx, req.(*IssueCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialsIssueService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialsIssueServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.CredentialsIssueService/update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialsIssueServiceServer).Update(ctx, req.(*UpdateCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CredentialsIssueService_ServiceDesc is the grpc.ServiceDesc for CredentialsIssueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CredentialsIssueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.services.CredentialsIssueService",
	HandlerType: (*CredentialsIssueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _CredentialsIssueService_Get_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _CredentialsIssueService_Delete_Handler,
		},
		{
			MethodName: "issue",
			Handler:    _CredentialsIssueService_Issue_Handler,
		},
		{
			MethodName: "update",
			Handler:    _CredentialsIssueService_Update_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getAll",
			Handler:       _CredentialsIssueService_GetAll_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1/services/issue.proto",
}
