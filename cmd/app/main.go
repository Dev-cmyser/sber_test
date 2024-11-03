package main

import (
	"log"

	"github.com/Dev-cmyser/calc_ipoteka/config"
	"github.com/Dev-cmyser/calc_ipoteka/internal/app"
)

const configPath = "config/config.yml"

func main() {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
