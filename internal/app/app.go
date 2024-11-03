package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Dev-cmyser/calc_ipoteka/config"
	v1 "github.com/Dev-cmyser/calc_ipoteka/internal/controller/http/v1"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/httpserver"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @title           Calculate Ipoteca Service
// @version         1.0

// @contact.name   Kirill Novgorodtsev
// @contact.email  kirrallt@mail.ru

// @host      localhost:8080
// @BasePath  /

func Run(cfg *config.Config) {
	log := logger.New(cfg.Log.Level)
	log.Info("Initializing services...")

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, log, translationUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
