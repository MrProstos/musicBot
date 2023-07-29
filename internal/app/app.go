package app

import (
	"github.com/MrProstos/musicBot/config"
	"github.com/MrProstos/musicBot/internal/usecase"
	"github.com/MrProstos/musicBot/internal/usecase/repository"
	"github.com/MrProstos/musicBot/pkg/psql"
	"log"
)

func Run(cfg *config.Config) {
	db, err := psql.New(cfg.Psql)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}

	newBot, err := usecase.NewMusicBot(cfg.App)
	if err != nil {
		log.Fatalln(err)
	}

	newBot.DB = repository.New(db)

	newBot.StartListening()
}
