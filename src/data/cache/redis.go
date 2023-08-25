package cache

import (
	"context"
	"fmt"
	"time"
	"wapi/src/config"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DB:           0,
		DialTimeout:  cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:  cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
		PoolSize:     cfg.Redis.PoolSize,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: cfg.Redis.IdleCheckFrequency * time.Millisecond,
	})
	ctx := context.Background()
	res, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}
