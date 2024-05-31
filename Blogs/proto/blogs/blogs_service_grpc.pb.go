// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: blogs/blogs_service.proto

package blogs

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
	BlogService_GetBlog_FullMethodName    = "/BlogService/GetBlog"
	BlogService_CreateBlog_FullMethodName = "/BlogService/CreateBlog"
)

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	GetBlog(ctx context.Context, in *GetBlogRequest, opts ...grpc.CallOption) (*Blog, error)
	CreateBlog(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*Blog, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) GetBlog(ctx context.Context, in *GetBlogRequest, opts ...grpc.CallOption) (*Blog, error) {
	out := new(Blog)
	err := c.cc.Invoke(ctx, BlogService_GetBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *Blog, opts ...grpc.CallOption) (*Blog, error) {
	out := new(Blog)
	err := c.cc.Invoke(ctx, BlogService_CreateBlog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations must embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	GetBlog(context.Context, *GetBlogRequest) (*Blog, error)
	CreateBlog(context.Context, *Blog) (*Blog, error)
	mustEmbedUnimplementedBlogServiceServer()
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) GetBlog(context.Context, *GetBlogRequest) (*Blog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlog not implemented")
}
func (UnimplementedBlogServiceServer) CreateBlog(context.Context, *Blog) (*Blog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (UnimplementedBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {}

// UnsafeBlogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServiceServer will
// result in compilation errors.
type UnsafeBlogServiceServer interface {
	mustEmbedUnimplementedBlogServiceServer()
}

func RegisterBlogServiceServer(s grpc.ServiceRegistrar, srv BlogServiceServer) {
	s.RegisterService(&BlogService_ServiceDesc, srv)
}

func _BlogService_GetBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_GetBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetBlog(ctx, req.(*GetBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlogService_CreateBlog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*Blog))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlog",
			Handler:    _BlogService_GetBlog_Handler,
		},
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blogs/blogs_service.proto",
}
