package models

import "time"

const AudioPath = "storage/audio/"

type AudioStorage struct {
	Id         uint `gorm:"primaryKey"`
	PlaylistId uint
	VideoId    string    `gorm:"type:varchar(255);index;not null"`
	Title      string    `gorm:"type:varchar(255);index;not null"`
	Author     string    `gorm:"type:varchar(255);index;not null"`
	FilePath   string    `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time `gorm:"index;not null;default:NOW()"`
}
