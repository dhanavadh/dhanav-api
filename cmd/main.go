package main

import (
	"log"

	"github.com/dhanavadh/dhanav-api/internal/config"
	"github.com/dhanavadh/dhanav-api/internal/database"
	"github.com/dhanavadh/dhanav-api/internal/handlers"
	"github.com/dhanavadh/dhanav-api/internal/repository"
	"github.com/dhanavadh/dhanav-api/internal/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r2 := storage.NewR2Storage(cfg)

	postRepo := repository.NewPostRepository(db)
	aboutRepo := repository.NewAboutRepository(db)

	blogHandler := handlers.NewBlogHandler(postRepo)
	projectHandler := handlers.NewProjectHandler(postRepo)
	aboutHandler := handlers.NewAboutHandler(aboutRepo)
	uploadHandler := handlers.NewUploadHandler(r2)

	gin.SetMode(cfg.GinMode)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/blogs", blogHandler.GetBlogs)
	v1.GET("/blogs/:slug", blogHandler.GetBlog)

	v1.GET("/projects", projectHandler.GetProjects)
	v1.GET("/projects/:slug", projectHandler.GetProject)

	v1.GET("/about", aboutHandler.Get)

	v1.POST("/upload", uploadHandler.UploadFile)

	r.GET("/health", handlers.HealthCheck)

	r.Run(cfg.ServerPort)
}
