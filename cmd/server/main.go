package main

import (
	"context"
	"gitlab.com/synthesis-backend/auth-service/pkg/protocol/grpc"
	"gitlab.com/synthesis-backend/auth-service/pkg/service/v1"
)

func main() {
	grpc.RunServer(context.Background(), &v1.Server{}, "7777")
}
