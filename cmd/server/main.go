package main

import (
	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zero-frost/auth-service/pkg/config"
	"github.com/zero-frost/auth-service/pkg/models"
	"github.com/zero-frost/auth-service/pkg/protocol/grpc"
	// "github.com/zero-frost/auth-service/pkg/protocol/http"
	"github.com/zero-frost/auth-service/pkg/service/v1"
)

func main() {
	// go func() {
	// 	_ = http.RunServer(context.Background(), "7777", "8080")
	// }()
	// TODO: create function to generate db from config
	c, err := config.GetConfig("./config", "env.yml")
	db, err := gorm.Open(c.Database.Kind, c.Database.DatabaseName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(models.User{})

	grpc.RunServer(context.Background(), c, v1.NewAuthServer(db, c.ServerSettings), &v1.HealthServer{}, v1.NewMetricServer())

}
