package app

import (
	"log/slog"
	"net/http"

	"github.com/Sheridanlk/Service/internal/config"
	"github.com/Sheridanlk/Service/internal/httpService/router/chirouter"
	"github.com/Sheridanlk/Service/internal/storage/postgresql"
)

type App struct {
	Storage *postgresql.Storage
	Server  *http.Server
	// client
}

func New(log *slog.Logger, storageCfg config.PostgreSQL, serverCfg config.HTTPServer) *App {
	storage, err := postgresql.New(storageCfg.Host, storageCfg.Port, storageCfg.DBName, storageCfg.UserName, storageCfg.Password)
	if err != nil {
		panic(err)
	}

	router := chirouter.Setup(log, storage)

	server := &http.Server{
		Addr:         serverCfg.Address,
		Handler:      router,
		ReadTimeout:  serverCfg.Timeout,
		WriteTimeout: serverCfg.Timeout,
		IdleTimeout:  serverCfg.IdleTimeout,
	}

	return &App{
		Storage: storage,
		Server:  server,
	}
}
