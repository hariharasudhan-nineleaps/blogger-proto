// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: grpc/proto/blog/blog.proto

package blog

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

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error)
	GetUserBlogs(ctx context.Context, in *UserBlogRequest, opts ...grpc.CallOption) (*UserBlogResponse, error)
	GetUserBlog(ctx context.Context, in *GetUserBlogRequest, opts ...grpc.CallOption) (*UserBlog, error)
	ViewBlog(ctx context.Context, in *ViewBlogRequest, opts ...grpc.CallOption) (*ViewBlogResponse, error)
	MostViewedBlogs(ctx context.Context, in *MostViewedBlogsRequest, opts ...grpc.CallOption) (*MostViewedBlogsResponse, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error) {
	out := new(CreateBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetUserBlogs(ctx context.Context, in *UserBlogRequest, opts ...grpc.CallOption) (*UserBlogResponse, error) {
	out := new(UserBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/GetUserBlogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) GetUserBlog(ctx context.Context, in *GetUserBlogRequest, opts ...grpc.CallOption) (*UserBlog, error) {
	out := new(UserBlog)
	err := c.cc.Invoke(ctx, "/blog.BlogService/GetUserBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ViewBlog(ctx context.Context, in *ViewBlogRequest, opts ...grpc.CallOption) (*ViewBlogResponse, error) {
	out := new(ViewBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ViewBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) MostViewedBlogs(ctx context.Context, in *MostViewedBlogsRequest, opts ...grpc.CallOption) (*MostViewedBlogsResponse, error) {
	out := new(MostViewedBlogsResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/MostViewedBlogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations must embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error)
	GetUserBlogs(context.Context, *UserBlogRequest) (*UserBlogResponse, error)
	GetUserBlog(context.Context, *GetUserBlogRequest) (*UserBlog, error)
	ViewBlog(context.Context, *ViewBlogRequest) (*ViewBlogResponse, error)
	MostViewedBlogs(context.Context, *MostViewedBlogsRequest) (*MostViewedBlogsResponse, error)
	mustEmbedUnimplementedBlogServiceServer()
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (UnimplementedBlogServiceServer) GetUserBlogs(context.Context, *UserBlogRequest) (*UserBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBlogs not implemented")
}
func (UnimplementedBlogServiceServer) GetUserBlog(context.Context, *GetUserBlogRequest) (*UserBlog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBlog not implemented")
}
func (UnimplementedBlogServiceServer) ViewBlog(context.Context, *ViewBlogRequest) (*ViewBlogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewBlog not implemented")
}
func (UnimplementedBlogServiceServer) MostViewedBlogs(context.Context, *MostViewedBlogsRequest) (*MostViewedBlogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MostViewedBlogs not implemented")
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

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetUserBlogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetUserBlogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/GetUserBlogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetUserBlogs(ctx, req.(*UserBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_GetUserBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).GetUserBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/GetUserBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).GetUserBlog(ctx, req.(*GetUserBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ViewBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ViewBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ViewBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ViewBlog(ctx, req.(*ViewBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_MostViewedBlogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MostViewedBlogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).MostViewedBlogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/MostViewedBlogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).MostViewedBlogs(ctx, req.(*MostViewedBlogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "GetUserBlogs",
			Handler:    _BlogService_GetUserBlogs_Handler,
		},
		{
			MethodName: "GetUserBlog",
			Handler:    _BlogService_GetUserBlog_Handler,
		},
		{
			MethodName: "ViewBlog",
			Handler:    _BlogService_ViewBlog_Handler,
		},
		{
			MethodName: "MostViewedBlogs",
			Handler:    _BlogService_MostViewedBlogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/blog/blog.proto",
}
