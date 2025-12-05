package repository

import (
	"fmt"

	"github.com/dhanavadh/dhanav-api/internal/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) FindAll(postType string) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Where("post_type = ? AND status = ?", postType, "published").Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find posts: %w", err)
	}
	return posts, nil
}

func (r *PostRepository) FindBySlug(slug string, postType string) (*models.Post, error) {
	var post models.Post
	err := r.db.Where("slug = ? AND post_type = ? AND status = ?", slug, postType, "published").First(&post).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find post by slug: %w", err)
	}
	return &post, nil
}
