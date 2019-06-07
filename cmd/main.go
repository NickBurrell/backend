package main

import (
	"context"
	"github.com/zero-frost/auth-service/pkg/protocol/grpc"
	"github.com/zero-frost/auth-service/pkg/service/v1"
)

func main() {
	grpc.RunServer(context.Background(), &v1.Server{}, "7777")
}
