// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// MovieDetailServiceClient is the client API for MovieDetailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieDetailServiceClient interface {
	MovieDetail(ctx context.Context, in *MovieDetailRequest, opts ...grpc.CallOption) (*MovieDetailResponse, error)
}

type movieDetailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieDetailServiceClient(cc grpc.ClientConnInterface) MovieDetailServiceClient {
	return &movieDetailServiceClient{cc}
}

func (c *movieDetailServiceClient) MovieDetail(ctx context.Context, in *MovieDetailRequest, opts ...grpc.CallOption) (*MovieDetailResponse, error) {
	out := new(MovieDetailResponse)
	err := c.cc.Invoke(ctx, "/proto.v1.MovieDetailService/MovieDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieDetailServiceServer is the server API for MovieDetailService service.
// All implementations must embed UnimplementedMovieDetailServiceServer
// for forward compatibility
type MovieDetailServiceServer interface {
	MovieDetail(context.Context, *MovieDetailRequest) (*MovieDetailResponse, error)
	mustEmbedUnimplementedMovieDetailServiceServer()
}

// UnimplementedMovieDetailServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMovieDetailServiceServer struct {
}

func (UnimplementedMovieDetailServiceServer) MovieDetail(context.Context, *MovieDetailRequest) (*MovieDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovieDetail not implemented")
}
func (UnimplementedMovieDetailServiceServer) mustEmbedUnimplementedMovieDetailServiceServer() {}

// UnsafeMovieDetailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieDetailServiceServer will
// result in compilation errors.
type UnsafeMovieDetailServiceServer interface {
	mustEmbedUnimplementedMovieDetailServiceServer()
}

func RegisterMovieDetailServiceServer(s grpc.ServiceRegistrar, srv MovieDetailServiceServer) {
	s.RegisterService(&MovieDetailService_ServiceDesc, srv)
}

func _MovieDetailService_MovieDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieDetailServiceServer).MovieDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.v1.MovieDetailService/MovieDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieDetailServiceServer).MovieDetail(ctx, req.(*MovieDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieDetailService_ServiceDesc is the grpc.ServiceDesc for MovieDetailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieDetailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.v1.MovieDetailService",
	HandlerType: (*MovieDetailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MovieDetail",
			Handler:    _MovieDetailService_MovieDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/movie_detail.proto",
}
