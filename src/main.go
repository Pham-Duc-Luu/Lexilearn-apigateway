package main

import (
	"api-gateway/src/config"
	"api-gateway/src/controller"
	"api-gateway/src/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Load configuration
	configData, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := gin.Default()

	r.Use(middleware.RecordMetrics)
	// Initialize controllers
	userController := &controller.UserController{
		ServiceConfig: configData.Services["userService"],
	}

	imageController := &controller.ImageController{
		ServiceConfig: configData.Services["imageService"],
	}

	// Register routes
	r.Any("/user/*proxyPath", userController.ProxyRequest)
	r.Any("/image/*proxyPath", imageController.ProxyRequest)
	// Add a simple Ping route for health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Expose Prometheus metrics at /metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Start the API Gateway
	port := configData.Global.Port
	if port[0] != ':' {
		port = ":" + port
	}
	r.Run(port)
}
