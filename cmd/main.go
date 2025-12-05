package main

import (
	"log"

	"github.com/dhanavadh/dhanav-api/internal/config"
	"github.com/dhanavadh/dhanav-api/internal/database"
	"github.com/dhanavadh/dhanav-api/internal/handlers"
	"github.com/dhanavadh/dhanav-api/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	postRepo := repository.NewPostRepository(db)
	blogHandler := handlers.NewBlogHandler(postRepo)
	projectHandler := handlers.NewProjectHandler(postRepo)

	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	api := r.Group("/api/v1")
	api.GET("/blogs", blogHandler.GetBlogs)
	api.GET("/blogs/:slug", blogHandler.GetBlog)

	api.GET("/projects", projectHandler.GetProjects)
	api.GET("/projects/:slug", projectHandler.GetProject)

	api.GET("/about", handlers.GetAbout)

	r.GET("/health", handlers.HealthCheck)

	r.Run(cfg.ServerPort)
}
