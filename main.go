package main

import (
	"apis/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", handlers.Health)
	router.GET("/dashboard", handlers.Dashboard)

	router.Run()

}
