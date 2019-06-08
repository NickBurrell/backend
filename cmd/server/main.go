package main

import (
	"context"
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

	grpc.RunServer(context.Background(), v1.NewAuthServer(db), &v1.HealthServer{}, v1.NewMetricServer(), "7777")
}
