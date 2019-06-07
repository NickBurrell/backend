package main

import (
	"github.com/zero-frost/auth-service/pkg/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := v1.NewAuthClient(conn)
	// response, err := c.Login(context.Background(), &v1.LoginRequest{
	// 	Username: "test",
	// 	Password: "test1234",
	// })
	// if err != nil {
	// 	log.Fatalf("Error when calling CheckAuth: %s", err)
	// }

	// log.Printf("%s", response.OpaqueToken)
	response, err := c.CreateUser(context.Background(), &v1.CreateUserRequest{
		Username: "test",
		Email:    "test@gmail.com",
		Password: "test1234",
	})
	if err != nil {
		log.Fatalf("Error when calling CheckAuth: %s", err)
	}
	switch response.ErrorCode {
	case v1.CreateUserResponse_INTERNAL_ERROR:
		log.Printf("Failure!")
	default:
		log.Fatalf("SUCCESS!")
	}
}