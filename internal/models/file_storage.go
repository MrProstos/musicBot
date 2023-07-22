package Models

import "time"

const AudioPath = "Storage/Audio/"

type FileStorage struct {
	Id        uint      `gorm:"primaryKey"`
	FileName  string    `gorm:"type:varchar(255);not null"`
	Title     string    `gorm:"type:varchar(255);index;not null"`
	Album     string    `gorm:"type:varchar(255);index;not null"`
	Author    string    `gorm:"type:varchar(255);index;not null"`
	CreatedAt time.Time `gorm:"index;not null;default:NOW()"`
}
