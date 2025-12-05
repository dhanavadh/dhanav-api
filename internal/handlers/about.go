package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetAbout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "about me"})
}
