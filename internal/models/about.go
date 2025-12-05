package models

import "time"

type About struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Bio         string         `json:"bio"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	Education   EducationList  `gorm:"type:jsonb;default:'[]'" json:"education"`
	Experience  ExperienceList `gorm:"type:jsonb;default:'[]'" json:"experience"`
	Sociallinks SocialLinks    `gorm:"type:jsonb;default:'{}'" json:"social"`
}
