package models

import (
	"time"
)

// Playlist users playlist
type Playlist struct {
	Id        uint      `gorm:"primaryKey"`
	UserId    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"index;not null;default:NOW()"`

	Musics []AudioStorage `gorm:"foreignKey:PlaylistId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
