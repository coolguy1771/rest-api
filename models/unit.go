package models

import "time"

// Unit table
type Unit struct {
	ID              uint       `gorm:"primary_key;column:ID"`
	UUID            string     `gorm:"not null;column:UNIT_ID"`
	Name            string     `gorm:"not null;column:UNIT_NAME"`
	Description     string     `gorm:"not null;column:UNIT_DESC"`
	Creator         string     `gorm:"not null;column:CREATOR"`
	CreatedDatetime time.Time  `gorm:"not null;column:CREATED_DATETIME"`
	UpdatedDatetime time.Time  `gorm:"column:UPDATED_DATETIME"`
	CreatedBy       string     `gorm:"not null;column:CREATED_BY"`
	UpdateBy        string     `gorm:"column:UPDATED_BY"`
}