package Models

import "time"

type PlayListItem struct {
	Id         uint      `gorm:"primaryKey"`
	PlaylistId uint      `gorm:"not null"`
	FileId     uint      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"index;not null;default:NOW()"`

	File FileStorage `gorm:"foreignKey:FileId"`
}
