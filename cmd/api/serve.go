package main

import (
	gin "github.com/gin-gonic/gin"
	"github.com/sebastienferry/gokr/core/repositories"
	"github.com/sebastienferry/gokr/handlers"
)

func main() {
	router := gin.Default()

	okrRepository := repositories.NewOkrRepositoryInMemory()
	okrHandler := handlers.NewOkrHandler(okrRepository)

	router.GET("/api/okrs", okrHandler.GetOkrs)
	router.GET("/api/okrs/:id", okrHandler.GetOkr)
	router.PUT("/api/okrs", okrHandler.PutOkr)
	router.POST("/api/okrs:id", okrHandler.PostOkr)
	router.DELETE("/api/okrs:id", okrHandler.DeleteOkr)

	router.Run("localhost:8080")
}
