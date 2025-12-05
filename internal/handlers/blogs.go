package handlers

import "github.com/gin-gonic/gin"

func GetBlogs(c *gin.Context) {
	c.JSON(200, gin.H{"message": "blogs page"})
}

func GetBlog(c *gin.Context) {
	slug := c.Param("slug")
	c.JSON(200, gin.H{"message": "blog" + slug})
}
