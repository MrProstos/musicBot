package main

import (
	"MrProstos/download_utils/config"
	"MrProstos/download_utils/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
