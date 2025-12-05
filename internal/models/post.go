package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Slug        string        `gorm:"uniqueIndex;not null" json:"slug"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	PostType    string        `json:"post_type"`
	Status      string        `json:"status"`
	Language    string        `json:"language"`
	Category    string        `json:"category"`
	CreatedAt   time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	Tags        StringArray   `gorm:"type:jsonb;default:'[]'" json:"tags"`
	Content     ContentBlocks `gorm:"type:jsonb;default:'[]'" json:"contents"`
}
