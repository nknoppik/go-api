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

// ObjectLoadServiceClient is the client API for ObjectLoadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectLoadServiceClient interface {
	CreateUploadLink(ctx context.Context, in *CreateUploadLinkRequest, opts ...grpc.CallOption) (*CreateUploadLinkResponse, error)
	CreateDownloadLink(ctx context.Context, in *CreateDownloadLinkRequest, opts ...grpc.CallOption) (*CreateDownloadLinkResponse, error)
	CreateDownloadLinksStream(ctx context.Context, opts ...grpc.CallOption) (ObjectLoadService_CreateDownloadLinksStreamClient, error)
	StartMultipartUpload(ctx context.Context, in *StartMultipartUploadRequest, opts ...grpc.CallOption) (*StartMultipartUploadResponse, error)
	GetMultipartUploadLink(ctx context.Context, in *GetMultipartUploadLinkRequest, opts ...grpc.CallOption) (*GetMultipartUploadLinkResponse, error)
	CompleteMultipartUpload(ctx context.Context, in *CompleteMultipartUploadRequest, opts ...grpc.CallOption) (*CompleteMultipartUploadResponse, error)
}

type objectLoadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectLoadServiceClient(cc grpc.ClientConnInterface) ObjectLoadServiceClient {
	return &objectLoadServiceClient{cc}
}

func (c *objectLoadServiceClient) CreateUploadLink(ctx context.Context, in *CreateUploadLinkRequest, opts ...grpc.CallOption) (*CreateUploadLinkResponse, error) {
	out := new(CreateUploadLinkResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.ObjectLoadService/CreateUploadLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectLoadServiceClient) CreateDownloadLink(ctx context.Context, in *CreateDownloadLinkRequest, opts ...grpc.CallOption) (*CreateDownloadLinkResponse, error) {
	out := new(CreateDownloadLinkResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.ObjectLoadService/CreateDownloadLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectLoadServiceClient) CreateDownloadLinksStream(ctx context.Context, opts ...grpc.CallOption) (ObjectLoadService_CreateDownloadLinksStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ObjectLoadService_ServiceDesc.Streams[0], "/api.services.v1.ObjectLoadService/CreateDownloadLinksStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectLoadServiceCreateDownloadLinksStreamClient{stream}
	return x, nil
}

type ObjectLoadService_CreateDownloadLinksStreamClient interface {
	Send(*CreateDownloadLinksStreamRequest) error
	Recv() (*CreateDownloadLinksStreamResponse, error)
	grpc.ClientStream
}

type objectLoadServiceCreateDownloadLinksStreamClient struct {
	grpc.ClientStream
}

func (x *objectLoadServiceCreateDownloadLinksStreamClient) Send(m *CreateDownloadLinksStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *objectLoadServiceCreateDownloadLinksStreamClient) Recv() (*CreateDownloadLinksStreamResponse, error) {
	m := new(CreateDownloadLinksStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectLoadServiceClient) StartMultipartUpload(ctx context.Context, in *StartMultipartUploadRequest, opts ...grpc.CallOption) (*StartMultipartUploadResponse, error) {
	out := new(StartMultipartUploadResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.ObjectLoadService/StartMultipartUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectLoadServiceClient) GetMultipartUploadLink(ctx context.Context, in *GetMultipartUploadLinkRequest, opts ...grpc.CallOption) (*GetMultipartUploadLinkResponse, error) {
	out := new(GetMultipartUploadLinkResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.ObjectLoadService/GetMultipartUploadLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectLoadServiceClient) CompleteMultipartUpload(ctx context.Context, in *CompleteMultipartUploadRequest, opts ...grpc.CallOption) (*CompleteMultipartUploadResponse, error) {
	out := new(CompleteMultipartUploadResponse)
	err := c.cc.Invoke(ctx, "/api.services.v1.ObjectLoadService/CompleteMultipartUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectLoadServiceServer is the server API for ObjectLoadService service.
// All implementations should embed UnimplementedObjectLoadServiceServer
// for forward compatibility
type ObjectLoadServiceServer interface {
	CreateUploadLink(context.Context, *CreateUploadLinkRequest) (*CreateUploadLinkResponse, error)
	CreateDownloadLink(context.Context, *CreateDownloadLinkRequest) (*CreateDownloadLinkResponse, error)
	CreateDownloadLinksStream(ObjectLoadService_CreateDownloadLinksStreamServer) error
	StartMultipartUpload(context.Context, *StartMultipartUploadRequest) (*StartMultipartUploadResponse, error)
	GetMultipartUploadLink(context.Context, *GetMultipartUploadLinkRequest) (*GetMultipartUploadLinkResponse, error)
	CompleteMultipartUpload(context.Context, *CompleteMultipartUploadRequest) (*CompleteMultipartUploadResponse, error)
}

// UnimplementedObjectLoadServiceServer should be embedded to have forward compatible implementations.
type UnimplementedObjectLoadServiceServer struct {
}

func (UnimplementedObjectLoadServiceServer) CreateUploadLink(context.Context, *CreateUploadLinkRequest) (*CreateUploadLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUploadLink not implemented")
}
func (UnimplementedObjectLoadServiceServer) CreateDownloadLink(context.Context, *CreateDownloadLinkRequest) (*CreateDownloadLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDownloadLink not implemented")
}
func (UnimplementedObjectLoadServiceServer) CreateDownloadLinksStream(ObjectLoadService_CreateDownloadLinksStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateDownloadLinksStream not implemented")
}
func (UnimplementedObjectLoadServiceServer) StartMultipartUpload(context.Context, *StartMultipartUploadRequest) (*StartMultipartUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMultipartUpload not implemented")
}
func (UnimplementedObjectLoadServiceServer) GetMultipartUploadLink(context.Context, *GetMultipartUploadLinkRequest) (*GetMultipartUploadLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultipartUploadLink not implemented")
}
func (UnimplementedObjectLoadServiceServer) CompleteMultipartUpload(context.Context, *CompleteMultipartUploadRequest) (*CompleteMultipartUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteMultipartUpload not implemented")
}

// UnsafeObjectLoadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectLoadServiceServer will
// result in compilation errors.
type UnsafeObjectLoadServiceServer interface {
	mustEmbedUnimplementedObjectLoadServiceServer()
}

func RegisterObjectLoadServiceServer(s grpc.ServiceRegistrar, srv ObjectLoadServiceServer) {
	s.RegisterService(&ObjectLoadService_ServiceDesc, srv)
}

func _ObjectLoadService_CreateUploadLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUploadLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectLoadServiceServer).CreateUploadLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.ObjectLoadService/CreateUploadLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectLoadServiceServer).CreateUploadLink(ctx, req.(*CreateUploadLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectLoadService_CreateDownloadLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDownloadLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectLoadServiceServer).CreateDownloadLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.ObjectLoadService/CreateDownloadLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectLoadServiceServer).CreateDownloadLink(ctx, req.(*CreateDownloadLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectLoadService_CreateDownloadLinksStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ObjectLoadServiceServer).CreateDownloadLinksStream(&objectLoadServiceCreateDownloadLinksStreamServer{stream})
}

type ObjectLoadService_CreateDownloadLinksStreamServer interface {
	Send(*CreateDownloadLinksStreamResponse) error
	Recv() (*CreateDownloadLinksStreamRequest, error)
	grpc.ServerStream
}

type objectLoadServiceCreateDownloadLinksStreamServer struct {
	grpc.ServerStream
}

func (x *objectLoadServiceCreateDownloadLinksStreamServer) Send(m *CreateDownloadLinksStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *objectLoadServiceCreateDownloadLinksStreamServer) Recv() (*CreateDownloadLinksStreamRequest, error) {
	m := new(CreateDownloadLinksStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ObjectLoadService_StartMultipartUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMultipartUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectLoadServiceServer).StartMultipartUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.ObjectLoadService/StartMultipartUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectLoadServiceServer).StartMultipartUpload(ctx, req.(*StartMultipartUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectLoadService_GetMultipartUploadLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMultipartUploadLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectLoadServiceServer).GetMultipartUploadLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.ObjectLoadService/GetMultipartUploadLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectLoadServiceServer).GetMultipartUploadLink(ctx, req.(*GetMultipartUploadLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectLoadService_CompleteMultipartUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteMultipartUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectLoadServiceServer).CompleteMultipartUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.services.v1.ObjectLoadService/CompleteMultipartUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectLoadServiceServer).CompleteMultipartUpload(ctx, req.(*CompleteMultipartUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectLoadService_ServiceDesc is the grpc.ServiceDesc for ObjectLoadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectLoadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.services.v1.ObjectLoadService",
	HandlerType: (*ObjectLoadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUploadLink",
			Handler:    _ObjectLoadService_CreateUploadLink_Handler,
		},
		{
			MethodName: "CreateDownloadLink",
			Handler:    _ObjectLoadService_CreateDownloadLink_Handler,
		},
		{
			MethodName: "StartMultipartUpload",
			Handler:    _ObjectLoadService_StartMultipartUpload_Handler,
		},
		{
			MethodName: "GetMultipartUploadLink",
			Handler:    _ObjectLoadService_GetMultipartUploadLink_Handler,
		},
		{
			MethodName: "CompleteMultipartUpload",
			Handler:    _ObjectLoadService_CompleteMultipartUpload_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateDownloadLinksStream",
			Handler:       _ObjectLoadService_CreateDownloadLinksStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/services/v1/object_load.proto",
}
