package Models

import (
	"time"
)

type Playlist struct {
	Id        uint      `gorm:"primaryKey"`
	UserId    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"index;not null;default:NOW()"`

	Musics []PlayListItem `gorm:"foreignKey:PlaylistId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
