package server

import (
	"fmt"
	"log"
	"net"

	"github.com/ariefrpm/movies2/pkg/moviedetail"
	"github.com/ariefrpm/movies2/pkg/saldodetail"
	"github.com/ariefrpm/movies2/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	svc   *service.Services
	port  int
	errCh chan error
}

func NewGrpcServer(svc *service.Services, port int) Server {
	return &grpcServer{
		svc:  svc,
		port: port,
	}
}

func (g *grpcServer) Run() {
	log.Printf("start running grpc server on port %d\n", g.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		g.errCh <- err
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	moviedetail.GrpcHandler(g.svc.MovieDetailService, grpcServer)
	saldodetail.GrpcHandler(g.svc.SaldoDetailService, grpcServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		g.errCh <- err
	}

}

func (g *grpcServer) ListenError() <-chan error {
	return g.errCh
}
