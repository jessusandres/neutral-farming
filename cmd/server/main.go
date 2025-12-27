package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"looker.com/neutral-farming/internal/repository/gorm"

	"looker.com/neutral-farming/internal/config"
	"looker.com/neutral-farming/internal/controller"
)

func main() {
	port := strconv.Itoa(config.EnvConfig.Port)
	hostname := config.EnvConfig.Host
	address := hostname + ":" + port

	log.Printf("🌐 Running locally with address: %s:", address)

	router := buildGin()

	gorm.NewDB()

	err := router.Run(address)

	if err != nil {
		panic(err)
	}
}

func buildGin() *gin.Engine {
	log.Printf("🚀 Initializing GIN server")

	router := gin.Default()

	api := router.Group(config.EnvConfig.ApiPrefix)
	{
		api.GET("", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Hello from Gin API!",
			})
		})
	}

	router.NoRoute(controller.NotFound)

	return router
}
