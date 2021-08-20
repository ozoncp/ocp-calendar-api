package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	api "github.com/ozoncp/ocp-calendar-api/app/ocp-calendar-api"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const (
	grpcPort           = ":82"
	grpcServerEndpoint = "127.0.0.1:82"
)

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOcpCalendarApiServer(s, api.NewOcpCalendarApi())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterOcpCalendarApiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)

	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}

}

func main() {
	go runJSON()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
