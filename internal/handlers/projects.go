package handlers

import (
	"github.com/dhanavadh/dhanav-api/internal/repository"
	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	repo *repository.PostRepository
}

func NewProjectHandler(repo *repository.PostRepository) *ProjectHandler {
	return &ProjectHandler{repo: repo}
}

func (h *ProjectHandler) GetProjects(c *gin.Context) {
	projects, err := h.repo.FindAll("project")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": projects})
}

func (h *ProjectHandler) GetProject(c *gin.Context) {
	slug := c.Param("slug")
	project, err := h.repo.FindBySlug(slug, "project")
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, gin.H{"data": project})
}

func GetProject(c *gin.Context) {
	slug := c.Param("slug")
	c.JSON(200, gin.H{"message": "get project by" + slug})
}
