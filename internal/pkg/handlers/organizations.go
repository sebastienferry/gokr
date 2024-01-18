package handlers

import (
	"fmt"
	"net/http"

	gin "github.com/gin-gonic/gin"
	"github.com/sebastienferry/gokr/core/logic"
	models "github.com/sebastienferry/gokr/core/models"
	"github.com/sebastienferry/gokr/core/repositories"
)

type OrganizationHandler struct {
	Repository repositories.OrganizationRepository
	Logic      logic.OrganizationLogic
}

func NewOrganizationHandler(repository repositories.OrganizationRepository) *OrganizationHandler {
	return &OrganizationHandler{Repository: repository, Logic: logic.NewOrganizationLogic(repository)}
}

// Get all organizations.
func (h *OrganizationHandler) GetOrganizations(c *gin.Context) {
	okrs, err := h.Repository.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, okrs)
}

// Get one OKR by ID.
func (h *OrganizationHandler) GetOrganization(c *gin.Context) {
	sid := c.Param("id")

	var id int64
	n, err := fmt.Sscanf(sid, "%d", &id)
	if err != nil || n != 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	org, err := h.Repository.Get(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Organization not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, org)
}

// Creates a new OKR.
func (h *OrganizationHandler) PutOrganization(c *gin.Context) {
	var newOrg models.Organization

	if err := c.BindJSON(&newOrg); err != nil {
		return
	}

	if err := h.Logic.Validate(newOrg); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newOrg, err := h.Repository.Create(newOrg)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newOrg)
}

// Update an existing OKR.
func (h *OrganizationHandler) PostOrganization(c *gin.Context) {
	var updatedOrg models.Organization

	if err := c.BindJSON(&updatedOrg); err != nil {
		return
	}

	if err := h.Logic.Validate(updatedOrg); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	okr, err := h.Repository.Update(updatedOrg)

	if err == nil {
		c.IndentedJSON(http.StatusOK, okr)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Organization not found"})
}

// Delete an existing OKR.
// TODO: Prevent deletion of OKRs that have children.
func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
	sid := c.Param("id")

	var id int64
	n, err := fmt.Sscanf(sid, "%d", &id)
	if err != nil || n != 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	err = h.Repository.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Organization not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
