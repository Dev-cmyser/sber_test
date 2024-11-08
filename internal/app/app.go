// Package app contains run func for start app
package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Dev-cmyser/calc_ipoteka/config"
	v1 "github.com/Dev-cmyser/calc_ipoteka/internal/controller/http/v1"
	"github.com/Dev-cmyser/calc_ipoteka/internal/entity"
	uc_mortgage "github.com/Dev-cmyser/calc_ipoteka/internal/usecase/ucmortgage"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/cache"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/httpserver"
	"github.com/Dev-cmyser/calc_ipoteka/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @title           Calculate Ipoteca Service
// @version         1.0

// @contact.name   Kirill Novgorodtsev
// @contact.email  kirrallt@mail.ru

// @host      localhost:8080
// @BasePath  /.

// Run start app.
func Run(cfg *config.Config) {
	log := logger.New(cfg.Log.Level)

	mCache := cache.SetCache[int, entity.CachedMortgage](cfg.Cache.TTL, cfg.Cache.SIZE)
	mortgage := uc_mortgage.New(mCache)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, log, mortgage)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	const (
		signalLogFormat        = "app - Run - signal: %s"
		notifyErrorLogFormat   = "app - Run - httpServer.Notify: %w"
		shutdownErrorLogFormat = "app - Run - httpServer.Shutdown: %w"
	)

	var err error
	select {
	case s := <-interrupt:
		log.Info(signalLogFormat, s.String())
	case err = <-httpServer.Notify():
		log.Error(notifyErrorLogFormat, err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		log.Error(shutdownErrorLogFormat, err)
	}
}
