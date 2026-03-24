package app

import (
	"log/slog"

	serverapp "github.com/Sheridanlk/Service/internal/app/server"
	"github.com/Sheridanlk/Service/internal/config"
	"github.com/Sheridanlk/Service/internal/http/router/chirouter"
	"github.com/Sheridanlk/Service/internal/storage/postgresql"
)

type App struct {
	Storage *postgresql.Storage
	Server  *serverapp.App
	// client
}

func New(log *slog.Logger, storage *postgresql.Storage, serverCfg config.HTTPServer) *App {

	router := chirouter.Setup(log, storage)

	server := serverapp.New(log, router, serverCfg.Address, serverCfg.Timeout, serverCfg.Timeout, serverCfg.IdleTimeout)

	return &App{
		Storage: storage,
		Server:  server,
	}
}
