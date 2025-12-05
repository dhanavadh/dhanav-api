package handlers

import (
	"github.com/dhanavadh/dhanav-api/internal/repository"
	"github.com/gin-gonic/gin"
)

type AboutHandler struct {
	repo *repository.AboutRepository
}

func NewAboutHandler(repo *repository.AboutRepository) *AboutHandler {
	return &AboutHandler{repo: repo}
}

func (h *AboutHandler) Get(c *gin.Context) {
	about, err := h.repo.Get()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, about)
}
