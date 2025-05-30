package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aiwan01/inventory-go/internal/cache"
	"github.com/Aiwan01/inventory-go/internal/db"
	"github.com/Aiwan01/inventory-go/internal/orders"
	"github.com/Aiwan01/inventory-go/internal/products"
	"github.com/Aiwan01/inventory-go/internal/users"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
)

func RateLimiter(limit *rate.Limiter) gin.HandlerFunc {
	return func(res *gin.Context) {
		if !limit.Allow() {
			res.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"message": "You have retry too many request, Please try after sometime"})
		}
		res.Next()
	}
}

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	// connected database postgres
	db.DatabaseConnect()

	// connecting redis for cache
	cache.ConnectRedis()
	app := gin.Default()

	// implement rate limiter
	limiter := rate.NewLimiter(1, 5)
	app.Use(RateLimiter(limiter))

	// user router calling
	users.UserRouter(app)

	// product routing
	products.ProductRouter(app)
	// order routing
	orders.OrderRouter(app)

	err := app.Run(":8000")
	if err != nil {
		log.Fatal("fail to start server")
	}
	fmt.Println("Server running on port 8000")
}

// defer
