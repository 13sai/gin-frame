package db

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: viper.GetString("redis.auth"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.MaxActive"), DialTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		PoolTimeout:  10 * time.Second,
	})

	pong, err := RedisClient.Ping(Ctx).Result()
	fmt.Println(pong, err)
}
