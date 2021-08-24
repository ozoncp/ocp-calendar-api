package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	api "github.com/ozoncp/ocp-calendar-api/internal/app/api"
	"github.com/ozoncp/ocp-calendar-api/internal/pkg/db"
	"github.com/ozoncp/ocp-calendar-api/internal/repo"
	desc "github.com/ozoncp/ocp-calendar-api/pkg/ocp-calendar-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
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
	dbConnection := db.Connect("postgres://postgres@localhost:5432/postgres?sslmode=disable")
	listen, err := net.Listen("tcp", grpcPort)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ocpCalendarApi, err := api.NewOcpCalendarApi(repo.NewCalendarRepository())

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(db.NewInterceptorWithDB(dbConnection)))
	desc.RegisterOcpCalendarApiServer(s, ocpCalendarApi)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func runJSON() {
	ctx, cancel := context.WithCancel(context.Background())
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

func initTracer() {
	cfg := jaegercfg.Configuration{
		ServiceName: "ocp_calendar_api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, _ := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	http.HandleFunc("/publish", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		serverSpan := tracer.StartSpan("server", ext.RPCServerOption(spanCtx))
		defer serverSpan.Finish()
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}

func runMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":9100", nil)

	if err != nil {
		log.Fatalf("failed prometheus %v", err)
	}
}

func main() {
	go runJSON()
	go initTracer()
	go runMetrics()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
