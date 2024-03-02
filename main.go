package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	httpPort := os.Getenv("PORT")

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := router.Run(fmt.Sprintf(":%s", httpPort))

	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
