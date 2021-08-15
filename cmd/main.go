package main

import (
	"github.com/ariefrpm/movies2/pkg/repository"
	"github.com/ariefrpm/movies2/pkg/server"
	service2 "github.com/ariefrpm/movies2/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	repo := repository.CreateRepository()
	service := service2.CreateServices(repo)

	httpServer := server.NewHttpServer(service, 8080)
	grpcServer := server.NewGrpcServer(service, 9000)

	go httpServer.Run()
	go grpcServer.Run()

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	select {
	case o := <-term:
		log.Printf("exiting gracefully %s", o.String())
	case er := <-httpServer.ListenError():
		log.Printf("error starting http server, exiting gracefully %s", er.Error())
	case er := <-grpcServer.ListenError():
		log.Printf("error starting grpc server, exiting gracefully %s", er.Error())
	}

}
