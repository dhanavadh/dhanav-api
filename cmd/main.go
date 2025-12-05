package main

import (
	"github.com/dhanavadh/dhanav-api/internal/config"
	"github.com/dhanavadh/dhanav-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	api := r.Group("/api/v1")
	api.GET("/blogs", handlers.GetBlogs)
	api.GET("/blogs/:slug", handlers.GetBlog)

	api.GET("/projects", handlers.GetProjects)
	api.GET("/projects/:slug", handlers.GetProject)

	api.GET("/about", handlers.GetAbout)

	r.GET("/health", handlers.HealthCheck)

	r.Run(cfg.ServerPort)
}
