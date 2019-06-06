package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"gitlab.com/zero_frost/auth-service/pkg/api/v1"
)

func RunServer(ctx context.Context, healthAPI v1.HealthServer, v1API v1.AuthServer, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	v1.RegisterAuthServer(server, v1API)
	v1.RegisterHealthServer(server, healthAPI)

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
