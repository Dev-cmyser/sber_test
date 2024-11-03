package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Dev-cmyser/calc_ipoteka/config"
)

// @title           Calculate Ipoteca Service
// @version         1.0

// @contact.name   Kirill Novgorodtsev
// @contact.email  kirrallt@mail.ru

// @host      localhost:8080
// @BasePath  /

func Run(configPath string) {
	log := SetLogger()

	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal().Msgf("Config error: %s", err)
	}

	SetLoggerLevel(cfg.Log.Level)

	// Services dependencies
	log.Info().Msg("Initializing services...")
	// deps := service.ServicesDependencies{
	// 	Repos:    repositories,
	// 	GDrive:   gdrive.New(cfg.WebAPI.GDriveJSONFilePath),
	// 	Hasher:   hasher.NewSHA1Hasher(cfg.Hasher.Salt),
	// 	SignKey:  cfg.JWT.SignKey,
	// 	TokenTTL: cfg.JWT.TokenTTL,
	// }
	// services := service.NewServices(deps)

	// Echo handler
	// log.Info("Initializing handlers and routes...")
	// handler := echo.New()
	// setup handler validator as lib validator
	// handler.Validator = validator.NewCustomValidator()
	// v1.NewRouter(handler, services)

	// HTTP server
	// log.Info("Starting http server...")
	// log.Debug("Server port: %s", cfg.HTTP.Port)
	// httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	// log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info().Msg("app - Run - signal: " + s.String())
		// case err = <-httpServer.Notify():
		log.Error().Msgf("app - Run - httpServer.Notify: %w", err)
	}

	// Graceful shutdown
	log.Info().Msg("Shutting down...")
	// err = httpServer.Shutdown()
	if err != nil {
		log.Error().Msgf("app - Run - httpServer.Shutdown: %w", err)
	}
}
