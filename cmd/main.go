package main

import (
	"fmt"
	"log"

	"github.com/Aiwan01/inventory-go/internal/cache"
	"github.com/Aiwan01/inventory-go/internal/db"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

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

	err := app.Run(":8000")
	if err != nil {
		log.Fatal("fail to start server")
	}
	fmt.Println("Server running on port 8000")
}

// defer
