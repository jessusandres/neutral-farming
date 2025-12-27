package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"looker.com/neutral-farming/internal/config"
	"looker.com/neutral-farming/internal/controller"
	"looker.com/neutral-farming/internal/http"
	"looker.com/neutral-farming/internal/http/middlewares"
	"looker.com/neutral-farming/internal/repository/gorm"
	"looker.com/neutral-farming/internal/service"
)

func main() {
	port := strconv.Itoa(config.EnvConfig.Port)
	hostname := config.EnvConfig.Host
	address := hostname + ":" + port

	log.Printf("🌐 Running locally with address: %s:", address)

	router := buildGin()

	err := router.Run(address)

	if err != nil {
		panic(err)
	}
}

func buildGin() *gin.Engine {
	log.Printf("🚀 Initializing GIN server")

	router := gin.Default()

	// Middlewares
	router.Use(middlewares.NewState())
	router.Use(middlewares.HandleErr())

	// Dependencies
	db := gorm.NewDB()
	repos := gorm.NewRepositories(db)
	services := service.NewServices(repos)
	controllers := controller.NewControllers(services)

	http.SetupRouter(router, controllers)

	router.NoRoute(controller.NotFound)

	return router
}
