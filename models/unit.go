package models

import (
	"time"

	"gorm.io/gorm"
)

// Unit table
type Unit struct {
	ApiUnit       `gorm:"embedded"`
	Name          string `gorm:"not null;column:UNIT_NAME"`
	Description   string `gorm:"not null;column:UNIT_DESC"`
	RankStructure string `gorm:"not null;column:RANK_TYPE"`
	Affiliation   string `gorm:"not null;column:AFFILIATION"`
	Branch        string `gorm:"not null;column:BRANCH"`
	Leader        string `gorm:"not null;column:LEADER"`
	CreatedBy     string `gorm:"not null;column:CREATED_BY"`
	UpdateBy      string `gorm:"column:UPDATED_BY"`
}

type ApiUnit struct {
	gorm.Model
	ID            uint8          `gorm:"primaryKey"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index.html"`
	Name          string
	Affiliation   string
	Branch        string
	RankStructure string
	NumUsers      uint8
}
