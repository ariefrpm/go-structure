package saldodetail

import (
	"context"
	"log"

	"github.com/ariefrpm/movies2/gen/go/proto/v1"
	"github.com/ariefrpm/movies2/pkg/library/router"
	"google.golang.org/grpc"
)

type request struct {
	kasId int64
}

type response struct {
	*SaldoTransaksiKasDetail
}

func endpoint(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(request)
		m, err := s.SaldoDetail(r.kasId)
		log.Print(m)
		return response{m}, err
	}
}

func RestHandler(svc Service, r router.Router) {
	handler := router.NewHandler(endpoint(svc), restDecodeRequest, restEncodeResponse, restEncodeError)
	r.GET("/api/saldo_detail", handler)
}

func GrpcHandler(svc Service, grpcServer *grpc.Server) {
	log.Print("GRPC handler")
	t := &grpcTransport{handler: router.NewGrpcHandler(endpoint(svc), grpcDecodeRequest, grpcEncodeResponse)}
	log.Print(t)
	proto.RegisterSaldoDetailServiceServer(grpcServer, t)
}
