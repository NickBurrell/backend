package main

import (
	"gitlab.com/synthesis-backend/auth-service/pkg/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := v1.NewAuthClient(conn)

	response, err := c.CheckAuth(context.Background(), &v1.AuthRequest{Token: os.Args[1]})
	if err != nil {
		log.Fatalf("Error when calling CheckAuth: %s", err)
	}
	switch response.Status {
	case v1.AuthResponse_APPROVED:
		log.Printf("Success!")
	default:
		log.Fatalf("Failure!")
	}
}
