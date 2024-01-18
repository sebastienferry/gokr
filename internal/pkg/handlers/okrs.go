package handlers

import (
	"fmt"
	"net/http"

	gin "github.com/gin-gonic/gin"
	"github.com/sebastienferry/gokr/core/logic"
	models "github.com/sebastienferry/gokr/core/models"
	"github.com/sebastienferry/gokr/core/repositories"
)

type OkrHandler struct {
	Repository repositories.OkrRepository
	Logic      logic.OkrLogic
}

func NewOkrHandler(repository repositories.OkrRepository) *OkrHandler {
	return &OkrHandler{Repository: repository, Logic: logic.NewOkrLogic(repository)}
}

// Get all OKRs.
func (h *OkrHandler) GetOkrs(c *gin.Context) {
	okrs, err := h.Repository.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, okrs)
}

// Get one OKR by ID.
func (h *OkrHandler) GetOkr(c *gin.Context) {
	sid := c.Param("id")

	var id int64
	n, err := fmt.Sscanf(sid, "%d", &id)
	if err != nil || n != 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	okr, err := h.Repository.Get(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "OKR not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, okr)
}

// Creates a new OKR.
func (h *OkrHandler) PutOkr(c *gin.Context) {
	var newOkr models.Okr

	if err := c.BindJSON(&newOkr); err != nil {
		return
	}

	if err := h.Logic.Validate(newOkr); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newOkr, err := h.Repository.Create(newOkr)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newOkr)
}

// Update an existing OKR.
func (h *OkrHandler) PostOkr(c *gin.Context) {
	var updatedOkr models.Okr

	if err := c.BindJSON(&updatedOkr); err != nil {
		return
	}

	if err := h.Logic.Validate(updatedOkr); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	okr, err := h.Repository.Update(updatedOkr)

	if err == nil {
		c.IndentedJSON(http.StatusOK, okr)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "OKR not found"})
}

// Delete an existing OKR.
// TODO: Prevent deletion of OKRs that have children.
func (h *OkrHandler) DeleteOkr(c *gin.Context) {
	sid := c.Param("id")

	var id int64
	n, err := fmt.Sscanf(sid, "%d", &id)
	if err != nil || n != 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	err = h.Repository.Delete(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "OKR not found"})
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}
