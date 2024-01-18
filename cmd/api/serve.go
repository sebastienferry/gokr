package main

import (
	gin "github.com/gin-gonic/gin"
	"github.com/sebastienferry/gokr/core/repositories"
	"github.com/sebastienferry/gokr/handlers"
)

func main() {
	router := gin.Default()

	orgRepository := repositories.NewOrganizationRepositoryInMemory()
	orgHandler := handlers.NewOrganizationHandler(orgRepository)

	// Route organization requests
	router.GET("/api/organizations", orgHandler.GetOrganizations)
	router.GET("/api/organizations/:id", orgHandler.GetOrganization)
	router.PUT("/api/organizations", orgHandler.PutOrganization)
	router.POST("/api/organizations:id", orgHandler.PostOrganization)
	router.DELETE("/api/organizations:id", orgHandler.DeleteOrganization)

	okrRepository := repositories.NewOkrRepositoryInMemory()
	okrHandler := handlers.NewOkrHandler(orgRepository, okrRepository)

	// Route okr requests
	router.GET("/api/okrs", okrHandler.GetOkrs)
	router.GET("/api/okrs/:id", okrHandler.GetOkr)
	router.PUT("/api/okrs", okrHandler.PutOkr)
	router.POST("/api/okrs:id", okrHandler.PostOkr)
	router.DELETE("/api/okrs:id", okrHandler.DeleteOkr)

	router.Run("localhost:8080")
}
