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

// DatasetObjectsServiceClient is the client API for DatasetObjectsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasetObjectsServiceClient interface {
	// Creates a new object group
	CreateObjectGroup(ctx context.Context, in *CreateObjectGroupRequest, opts ...grpc.CallOption) (*CreateObjectGroupResponse, error)
	// Batch request of CreateObjectGroup
	// The call will preserve the ordering of the request in the response
	CreateObjectGroupBatch(ctx context.Context, in *CreateObjectGroupBatchRequest, opts ...grpc.CallOption) (*CreateObjectGroupBatchResponse, error)
	//Returns the object group with the given ID
	GetObjectGroup(ctx context.Context, in *GetObjectGroupRequest, opts ...grpc.CallOption) (*GetObjectGroupResponse, error)
	// Finishes the upload process for an object
	// This will change the status of the objects to "available"
	// Experimental, might change this to FinishObjectGroupUpload
	FinishObjectUpload(ctx context.Context, in *FinishObjectUploadRequest, opts ...grpc.CallOption) (*FinishObjectUploadResponse, error)
	// Deletes the given object group
	// This will also delete all associated objects both as metadata objects and the actual objects in the object storage
	DeleteObjectGroup(ctx context.Context, in *DeleteObjectGroupRequest, opts ...grpc.CallOption) (*DeleteObjectGroupResponse, error)
}

type datasetObjectsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetObjectsServiceClient(cc grpc.ClientConnInterface) DatasetObjectsServiceClient {
	return &datasetObjectsServiceClient{cc}
}

func (c *datasetObjectsServiceClient) CreateObjectGroup(ctx context.Context, in *CreateObjectGroupRequest, opts ...grpc.CallOption) (*CreateObjectGroupResponse, error) {
	out := new(CreateObjectGroupResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/CreateObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) CreateObjectGroupBatch(ctx context.Context, in *CreateObjectGroupBatchRequest, opts ...grpc.CallOption) (*CreateObjectGroupBatchResponse, error) {
	out := new(CreateObjectGroupBatchResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/CreateObjectGroupBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) GetObjectGroup(ctx context.Context, in *GetObjectGroupRequest, opts ...grpc.CallOption) (*GetObjectGroupResponse, error) {
	out := new(GetObjectGroupResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/GetObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) FinishObjectUpload(ctx context.Context, in *FinishObjectUploadRequest, opts ...grpc.CallOption) (*FinishObjectUploadResponse, error) {
	out := new(FinishObjectUploadResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/FinishObjectUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetObjectsServiceClient) DeleteObjectGroup(ctx context.Context, in *DeleteObjectGroupRequest, opts ...grpc.CallOption) (*DeleteObjectGroupResponse, error) {
	out := new(DeleteObjectGroupResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.DatasetObjectsService/DeleteObjectGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetObjectsServiceServer is the server API for DatasetObjectsService service.
// All implementations should embed UnimplementedDatasetObjectsServiceServer
// for forward compatibility
type DatasetObjectsServiceServer interface {
	// Creates a new object group
	CreateObjectGroup(context.Context, *CreateObjectGroupRequest) (*CreateObjectGroupResponse, error)
	// Batch request of CreateObjectGroup
	// The call will preserve the ordering of the request in the response
	CreateObjectGroupBatch(context.Context, *CreateObjectGroupBatchRequest) (*CreateObjectGroupBatchResponse, error)
	//Returns the object group with the given ID
	GetObjectGroup(context.Context, *GetObjectGroupRequest) (*GetObjectGroupResponse, error)
	// Finishes the upload process for an object
	// This will change the status of the objects to "available"
	// Experimental, might change this to FinishObjectGroupUpload
	FinishObjectUpload(context.Context, *FinishObjectUploadRequest) (*FinishObjectUploadResponse, error)
	// Deletes the given object group
	// This will also delete all associated objects both as metadata objects and the actual objects in the object storage
	DeleteObjectGroup(context.Context, *DeleteObjectGroupRequest) (*DeleteObjectGroupResponse, error)
}

// UnimplementedDatasetObjectsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDatasetObjectsServiceServer struct {
}

func (UnimplementedDatasetObjectsServiceServer) CreateObjectGroup(context.Context, *CreateObjectGroupRequest) (*CreateObjectGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectGroup not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) CreateObjectGroupBatch(context.Context, *CreateObjectGroupBatchRequest) (*CreateObjectGroupBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectGroupBatch not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) GetObjectGroup(context.Context, *GetObjectGroupRequest) (*GetObjectGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroup not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) FinishObjectUpload(context.Context, *FinishObjectUploadRequest) (*FinishObjectUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishObjectUpload not implemented")
}
func (UnimplementedDatasetObjectsServiceServer) DeleteObjectGroup(context.Context, *DeleteObjectGroupRequest) (*DeleteObjectGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObjectGroup not implemented")
}

// UnsafeDatasetObjectsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasetObjectsServiceServer will
// result in compilation errors.
type UnsafeDatasetObjectsServiceServer interface {
	mustEmbedUnimplementedDatasetObjectsServiceServer()
}

func RegisterDatasetObjectsServiceServer(s grpc.ServiceRegistrar, srv DatasetObjectsServiceServer) {
	s.RegisterService(&DatasetObjectsService_ServiceDesc, srv)
}

func _DatasetObjectsService_CreateObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).CreateObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/CreateObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).CreateObjectGroup(ctx, req.(*CreateObjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_CreateObjectGroupBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectGroupBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).CreateObjectGroupBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/CreateObjectGroupBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).CreateObjectGroupBatch(ctx, req.(*CreateObjectGroupBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_GetObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).GetObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/GetObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).GetObjectGroup(ctx, req.(*GetObjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_FinishObjectUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishObjectUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).FinishObjectUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/FinishObjectUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).FinishObjectUpload(ctx, req.(*FinishObjectUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetObjectsService_DeleteObjectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetObjectsServiceServer).DeleteObjectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.DatasetObjectsService/DeleteObjectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetObjectsServiceServer).DeleteObjectGroup(ctx, req.(*DeleteObjectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatasetObjectsService_ServiceDesc is the grpc.ServiceDesc for DatasetObjectsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatasetObjectsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.services.v1.DatasetObjectsService",
	HandlerType: (*DatasetObjectsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObjectGroup",
			Handler:    _DatasetObjectsService_CreateObjectGroup_Handler,
		},
		{
			MethodName: "CreateObjectGroupBatch",
			Handler:    _DatasetObjectsService_CreateObjectGroupBatch_Handler,
		},
		{
			MethodName: "GetObjectGroup",
			Handler:    _DatasetObjectsService_GetObjectGroup_Handler,
		},
		{
			MethodName: "FinishObjectUpload",
			Handler:    _DatasetObjectsService_FinishObjectUpload_Handler,
		},
		{
			MethodName: "DeleteObjectGroup",
			Handler:    _DatasetObjectsService_DeleteObjectGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/services/v1/dataset_object_service.proto",
}
