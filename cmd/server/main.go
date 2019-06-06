package main

import (
	"context"
	"gitlab.com/zero_frost/auth-service/pkg/protocol/grpc"
	"gitlab.com/zero_frost/auth-service/pkg/service/v1"
)

func main() {
	grpc.RunServer(context.Background(), &v1.HealthServer{}, &v1.Server{}, "7777")
}
