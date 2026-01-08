package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Sheridanlk/Service/internal/app"
	"github.com/Sheridanlk/Service/internal/config"
	"github.com/Sheridanlk/Service/internal/logger"
)

func main() {
	cfg := config.Load()
	log := logger.SetupLogger(cfg.Env)

	app := app.New(log, cfg.PostgreSQL, cfg.HTTPServer)

	go app.Server.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	app.Server.Stop()

	log.Info("Application stopped")
}
