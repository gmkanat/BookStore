package main

import (
	"BookStore/controllers"
	"BookStore/initializers"
	"BookStore/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	server              *gin.Engine
	BookController      controllers.BookController
	BookRouteController routes.BookRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	initializers.ConnectDB(&config)
	BookController = controllers.NewBookController(initializers.DB)
	BookRouteController = routes.NewBookRouteController(BookController)
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	// health check
	server.GET("/health-checker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	router := server.Group("/api")
	BookRouteController.BookRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
