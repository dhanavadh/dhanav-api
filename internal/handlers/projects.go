package handlers

import "github.com/gin-gonic/gin"

func GetProjects(c *gin.Context) {
	c.JSON(200, gin.H{"message": "list all projects"})
}

func GetProject(c *gin.Context) {
	slug := c.Param("slug")
	c.JSON(200, gin.H{"message": "get project by" + slug})
}
