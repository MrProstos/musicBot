package psql

import (
	"fmt"
	"github.com/MrProstos/musicBot/config"
	Models "github.com/MrProstos/musicBot/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Psql struct {
	*gorm.DB
}

func New(psql config.Psql) (*Psql, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		psql.Host, psql.User, psql.Password, psql.DatabaseName, psql.Port,
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Psql{gormDB}, nil
}

func (db *Psql) AutoMigrate() error {
	err := db.DB.AutoMigrate(
		new(Models.Playlist),
		new(Models.AudioStorage),
	)
	if err != nil {
		return err
	}

	return nil
}
