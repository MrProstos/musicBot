package main

import (
	"github.com/MrProstos/musicBot/config"
	"github.com/MrProstos/musicBot/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
