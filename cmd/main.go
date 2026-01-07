package main

import (
	"github.com/Sheridanlk/Service/internal/app"
	"github.com/Sheridanlk/Service/internal/config"
	"github.com/Sheridanlk/Service/internal/logger"
)

func main() {
	cfg := config.Load()
	log := logger.SetupLogger(cfg.Env)

	app := app.New(log, cfg.PostgreSQL, cfg.HTTPServer)

}
