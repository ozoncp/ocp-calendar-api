package main

import (
	"context"
	"database/sql"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	api "github.com/ozoncp/ocp-calendar-api/internal/app/api"
	"github.com/ozoncp/ocp-calendar-api/internal/pkg/db"
	"github.com/ozoncp/ocp-calendar-api/internal/repo"
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

func run(dbConnection *sql.DB) error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(db.NewInterceptorWithDB(dbConnection)))
	desc.RegisterOcpCalendarApiServer(s, api.NewOcpCalendarApi(repo.NewCalendarRepository()))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func runJSON(ctx context.Context) {
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
	ctx := context.Background()
	dbConnection := db.Connect("postgres://postgres@localhost:5432/postgres?sslmode=disable")

	go runJSON(ctx)

	if err := run(dbConnection); err != nil {
		log.Fatal(err)
	}
}
