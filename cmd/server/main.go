package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zero-frost/auth-service/models"
	"github.com/zero-frost/auth-service/pkg/protocol/grpc"
	"github.com/zero-frost/auth-service/pkg/protocol/http"
	"github.com/zero-frost/auth-service/pkg/service/v1"
)

func main() {
	go func() {
		_ = http.RunServer(context.Background(), "7777", "8080")
	}()
	db, err := gorm.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(models.User{})
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	grpc.RunServer(context.Background(), &v1.HealthServer{}, v1.NewServer(db, client), "7777")
}
