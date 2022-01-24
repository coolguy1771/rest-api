package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ApiUser   `gorm:"embedded"`
	DiscordID string
	Rank      string
	Squad     string
}

type ApiUser struct {
	ID         uint           `gorm:"primaryKey;unique"`
	Name       string         `json:"name"`
	Unit       string         `json:"unit"`
	LastActive time.Time      `json:"lastActive"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index.html"`
}
