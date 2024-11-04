package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/Dev-cmyser/calc_ipoteka/config"
	_ "github.com/Dev-cmyser/calc_ipoteka/docs"
	"github.com/Dev-cmyser/calc_ipoteka/internal/app"
)

const configPath = "config/config.yml"

func main() {
	err := loadEnvFromFile(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}

// Dynamic setting env
func loadEnvFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
