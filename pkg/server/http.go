package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ariefrpm/movies2/pkg/library/router"
	"github.com/ariefrpm/movies2/pkg/moviedetail"
	"github.com/ariefrpm/movies2/pkg/saldodetail"
	"github.com/ariefrpm/movies2/pkg/service"
)

type httpServer struct {
	svc   *service.Services
	port  int
	errCh chan error
}

func NewHttpServer(svc *service.Services, port int) Server {
	return &httpServer{
		svc:  svc,
		port: port,
	}
}

func (h *httpServer) Run() {
	log.Printf("start running http server on port %d\n", h.port)

	route := router.NewDefaultRouter()

	moviedetail.RestHandler(h.svc.MovieDetailService, route)
	saldodetail.RestHandler(h.svc.SaldoDetailService, route)
	err := http.ListenAndServe(fmt.Sprintf(":%d", h.port), route.Handler())

	if err != nil {
		h.errCh <- err
	}
}

func (h *httpServer) ListenError() <-chan error {
	return h.errCh
}
