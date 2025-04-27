package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
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
func GetRedisValue(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := redisDB.Get(ctx, "mykey").Result()
	if err != nil {
		c.String(500, "Error in getting value %v", err)
	}
	c.String(200, val)
}

// setting new key and value in redis
func SetRedisValue(c *gin.Context) {
	ctx, cancell := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancell()

	err := redisDB.Set(ctx, "", "", 0).Err()
	if err != nil {
		c.String(500, "failed to insert data in redis %v", err)
	}

	c.String(200, "Data saved in redis")
}
