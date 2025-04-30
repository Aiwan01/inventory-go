package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisDB *redis.Client

func ConnectRedis() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	_, err := redisDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis connected successfully")
}

// getting data from redis by key
func GetRedisValue(redisKey string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := redisDB.Get(ctx, redisKey).Result()
	if err != nil {
		return ""
	}
	return val
}

// setting new key and value in redis
func SetRedisValue(redisKey string, value string) bool {
	ctx, cancell := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancell()

	err := redisDB.Set(ctx, "", "", 0).Err()

	return err != nil
}
