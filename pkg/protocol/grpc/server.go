package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	// "crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/rcrowley/go-metrics"

	"github.com/zero-frost/auth-service/pkg/config"
	middleware "github.com/zero-frost/auth-service/pkg/middleware/metrics"

	"github.com/grpc-ecosystem/go-grpc-middleware"

	"github.com/zero-frost/auth-service/pkg/api/v1"
)

func RunServer(ctx context.Context, c *config.Config, authAPI v1.AuthServer, healthAPI v1.HealthServer, metricsAPI v1.MetricServer) error {
	// lis, err := net.Listen("tcp", ":"+string(c.Port))
	lis, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				middleware.MetricsUnaryInterceptor(),
			),
		)}

	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")

	log.Printf("%v\n", creds)
	if err != nil {
		log.Fatalf("Failed to load SSL while in production mode. '%v'", err)

	} else {
		// opts = append(opts, grpc.Creds(creds))
	}
	opts = append(opts, grpc.Creds(creds))
	server := grpc.NewServer(opts...)
	v1.RegisterAuthServer(server, authAPI)
	v1.RegisterHealthServer(server, healthAPI)
	v1.RegisterMetricServer(server, metricsAPI)

	go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		for range sig {
			log.Println("shutting down gRPC server...")

			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC server...")
	return server.Serve(lis)
}

func RunAuthServer(ctx context.Context, v1API v1.AuthServer, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	v1.RegisterAuthServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")

			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC server...")
	return server.Serve(lis)
}

func RunHealthServer(ctx context.Context, v1API v1.HealthServer, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	v1.RegisterHealthServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC health server...")

			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Println("starting gRPC health server...")
	return server.Serve(lis)
}
