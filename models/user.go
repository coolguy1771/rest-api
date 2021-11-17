package models

type User struct {
	UUID      string `gorm:"primaryKey"`
	Name      string
	DiscordID string
}
