package Models

import (
	"MrProstos/download_utils/App/Models/Database"
	"time"
)

type Playlist struct {
	Id        uint      `gorm:"primaryKey"`
	UserId    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"index;not null;default:NOW()"`

	Musics []PlayListItem `gorm:"foreignKey:PlaylistId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (playList *Playlist) Select() {
}

func (playList *Playlist) Insert() uint {
	Database.GetDB().Create(playList)

	return playList.Id
}

func (playList *Playlist) Delete() {

}
