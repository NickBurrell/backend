package main

import (
	"github.com/zero-frost/auth-service/pkg/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {

	var conn *grpc.ClientConn
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	conn, err = grpc.Dial("localhost:5300", grpc.WithTransportCredentials(creds))
	// conn, err = grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := v1.NewAuthClient(conn)
	response, err := c.Login(context.Background(), &v1.LoginRequest{
		Username: "test",
		Password: "test1234",
	})
	if err != nil {
		log.Fatalf("Error when calling CheckAuth: %s", err)
	}

	log.Printf("%s", response.Token)
	// response, err := c.CreateUser(context.Background(), &v1.CreateUserRequest{
	// 	Username: "test_user_1",
	// 	Email:    "test_email@gmail.com",
	// 	Password: "test12345",
	// })
	// if err != nil {
	// 	log.Fatalf("Error when calling CheckAuth: %s", err)
	// }
	// switch response.ErrorCode {
	// case v1.CreateUserResponse_INTERNAL_ERROR:
	// 	log.Printf("Failure!")
	// default:
	// 	log.Fatalf("SUCCESS!")
	// }
}
