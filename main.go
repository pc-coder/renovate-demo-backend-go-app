package main

import (
	"fmt"
	"renovate-demo-backend-go-app/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	healthController := controller.NewHealthController()

	router.GET("/api/v1/healthz", healthController.GetStatus)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Unexpected exit...")
		return
	}
}
