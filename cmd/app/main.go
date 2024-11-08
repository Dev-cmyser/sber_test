// Package main содержит основной запуск приложения и инициализацию сервисов.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

// Dynamic setting env.
func loadEnvFromFile(filePath string) error {
	absPath, err := filepath.Abs(filePath)

	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	absPath = filepath.Clean(absPath)
	if !strings.HasPrefix(absPath, "/path/to/safe/dir/") {
		return fmt.Errorf("file path is outside of the allowed safe directory:  %w", err)
	}

	file, err := os.Open(absPath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}()

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

		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("failed to set environment variable %s: %w", key, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan file %s: %w", filePath, err)
	}
	return nil
}
