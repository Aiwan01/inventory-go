package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	err := app.Run(":8000")
	if err != nil {
		log.Fatal("fail to start server")
	}

	log.Fatalln("Server running on port 8000")
}
