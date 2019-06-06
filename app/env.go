package app

import "github.com/jinzhu/gorm"
import "github.com/go-redis/redis"

type Env struct {
	DB    *gorm.DB
	Cache *redis.Client
}
