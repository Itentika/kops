// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/iam/v1/api_key_service.proto

package iam

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiKeyServiceClient is the client API for ApiKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiKeyServiceClient interface {
	// Retrieves the list of API keys for the specified service account.
	List(ctx context.Context, in *ListApiKeysRequest, opts ...grpc.CallOption) (*ListApiKeysResponse, error)
	// Returns the specified API key.
	//
	// To get the list of available API keys, make a [List] request.
	Get(ctx context.Context, in *GetApiKeyRequest, opts ...grpc.CallOption) (*ApiKey, error)
	// Creates an API key for the specified service account.
	Create(ctx context.Context, in *CreateApiKeyRequest, opts ...grpc.CallOption) (*CreateApiKeyResponse, error)
	// Updates the specified API key.
	Update(ctx context.Context, in *UpdateApiKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified API key.
	Delete(ctx context.Context, in *DeleteApiKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Retrieves the list of operations for the specified API key.
	ListOperations(ctx context.Context, in *ListApiKeyOperationsRequest, opts ...grpc.CallOption) (*ListApiKeyOperationsResponse, error)
}

type apiKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiKeyServiceClient(cc grpc.ClientConnInterface) ApiKeyServiceClient {
	return &apiKeyServiceClient{cc}
}

func (c *apiKeyServiceClient) List(ctx context.Context, in *ListApiKeysRequest, opts ...grpc.CallOption) (*ListApiKeysResponse, error) {
	out := new(ListApiKeysResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.ApiKeyService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyServiceClient) Get(ctx context.Context, in *GetApiKeyRequest, opts ...grpc.CallOption) (*ApiKey, error) {
	out := new(ApiKey)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.ApiKeyService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyServiceClient) Create(ctx context.Context, in *CreateApiKeyRequest, opts ...grpc.CallOption) (*CreateApiKeyResponse, error) {
	out := new(CreateApiKeyResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.ApiKeyService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyServiceClient) Update(ctx context.Context, in *UpdateApiKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.ApiKeyService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyServiceClient) Delete(ctx context.Context, in *DeleteApiKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.ApiKeyService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyServiceClient) ListOperations(ctx context.Context, in *ListApiKeyOperationsRequest, opts ...grpc.CallOption) (*ListApiKeyOperationsResponse, error) {
	out := new(ListApiKeyOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.ApiKeyService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiKeyServiceServer is the server API for ApiKeyService service.
// All implementations should embed UnimplementedApiKeyServiceServer
// for forward compatibility
type ApiKeyServiceServer interface {
	// Retrieves the list of API keys for the specified service account.
	List(context.Context, *ListApiKeysRequest) (*ListApiKeysResponse, error)
	// Returns the specified API key.
	//
	// To get the list of available API keys, make a [List] request.
	Get(context.Context, *GetApiKeyRequest) (*ApiKey, error)
	// Creates an API key for the specified service account.
	Create(context.Context, *CreateApiKeyRequest) (*CreateApiKeyResponse, error)
	// Updates the specified API key.
	Update(context.Context, *UpdateApiKeyRequest) (*operation.Operation, error)
	// Deletes the specified API key.
	Delete(context.Context, *DeleteApiKeyRequest) (*operation.Operation, error)
	// Retrieves the list of operations for the specified API key.
	ListOperations(context.Context, *ListApiKeyOperationsRequest) (*ListApiKeyOperationsResponse, error)
}

// UnimplementedApiKeyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedApiKeyServiceServer struct {
}

func (UnimplementedApiKeyServiceServer) List(context.Context, *ListApiKeysRequest) (*ListApiKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedApiKeyServiceServer) Get(context.Context, *GetApiKeyRequest) (*ApiKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedApiKeyServiceServer) Create(context.Context, *CreateApiKeyRequest) (*CreateApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedApiKeyServiceServer) Update(context.Context, *UpdateApiKeyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedApiKeyServiceServer) Delete(context.Context, *DeleteApiKeyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedApiKeyServiceServer) ListOperations(context.Context, *ListApiKeyOperationsRequest) (*ListApiKeyOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}

// UnsafeApiKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiKeyServiceServer will
// result in compilation errors.
type UnsafeApiKeyServiceServer interface {
	mustEmbedUnimplementedApiKeyServiceServer()
}

func RegisterApiKeyServiceServer(s grpc.ServiceRegistrar, srv ApiKeyServiceServer) {
	s.RegisterService(&ApiKeyService_ServiceDesc, srv)
}

func _ApiKeyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApiKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.ApiKeyService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).List(ctx, req.(*ListApiKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.ApiKeyService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).Get(ctx, req.(*GetApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.ApiKeyService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).Create(ctx, req.(*CreateApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.ApiKeyService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).Update(ctx, req.(*UpdateApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.ApiKeyService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).Delete(ctx, req.(*DeleteApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeyService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApiKeyOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.ApiKeyService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).ListOperations(ctx, req.(*ListApiKeyOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiKeyService_ServiceDesc is the grpc.ServiceDesc for ApiKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.iam.v1.ApiKeyService",
	HandlerType: (*ApiKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ApiKeyService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ApiKeyService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ApiKeyService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ApiKeyService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ApiKeyService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _ApiKeyService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/iam/v1/api_key_service.proto",
}