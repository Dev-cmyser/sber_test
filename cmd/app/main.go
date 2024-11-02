package main

import "github.com/Dev-cmyser/calc_ipoteka/internal/app"

const configPath = "config/config.yml"

func main() {
	app.Run(configPath)
}
