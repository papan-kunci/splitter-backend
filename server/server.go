package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/papan-kunci/splitter-backend/config"
)

func Init() {
	router := gin.Default()
	port := config.GetEnv("PORT")

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api/v1")
	{
		api.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	}

	err := router.Run(fmt.Sprintf(":%s", port))

	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
