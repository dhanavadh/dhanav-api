package models

import "time"

type About struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Bio       string    `json:"bio"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
