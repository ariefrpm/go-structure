package moviedetail

import (
	"context"
	"github.com/ariefrpm/movies2/gen/go/proto/v1"
	"github.com/ariefrpm/movies2/pkg/library/router"
	"google.golang.org/grpc"
)

type request struct {
	ID string
}

type response struct {
	*Movie
}

func endpoint(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(request)
		m, err := s.MovieDetail(r.ID)
		return response{m}, err
	}
}

func RestHandler(svc Service, r router.Router) {
	handler := router.NewHandler(endpoint(svc), restDecodeRequest, restEncodeResponse, restEncodeError)
	r.GET("/api/movie_detail", handler)
}

func GrpcHandler(svc Service, grpcServer *grpc.Server) {
	t := &grpcTransport{handler:router.NewGrpcHandler(endpoint(svc), grpcDecodeRequest, grpcEncodeResponse)}
	proto.RegisterMovieDetailServiceServer(grpcServer, t)
}