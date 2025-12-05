package handlers

import (
	"github.com/dhanavadh/dhanav-api/internal/repository"
	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	repo *repository.PostRepository
}

func NewBlogHandler(repo *repository.PostRepository) *BlogHandler {
	return &BlogHandler{repo: repo}
}

func (h *BlogHandler) GetBlogs(c *gin.Context) {
	posts, err := h.repo.FindAll("blog")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": posts})
}

func (h *BlogHandler) GetBlog(c *gin.Context) {
	slug := c.Param("slug")
	post, err := h.repo.FindBySlug(slug, "blog")
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, gin.H{"data": post})
}
