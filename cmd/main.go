package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Sheridanlk/Service/internal/app"
	"github.com/Sheridanlk/Service/internal/config"
	"github.com/Sheridanlk/Service/internal/logger"
	"github.com/Sheridanlk/Service/internal/storage/postgresql"
)

func main() {
	cfg := config.Load()
	log := logger.SetupLogger(cfg.Env)

	storage, err := postgresql.New(cfg.PostgreSQL.Host, cfg.PostgreSQL.UserName, cfg.PostgreSQL.Password, cfg.PostgreSQL.DBName, cfg.PostgreSQL.Port)
	if err != nil {
		panic(err)
	}

	// Регистрация в оснвном сервсие и получение ключей / загрузка данных из state.yaml

	// Проверка ключей / и запуск / перезапуск сервера

	// Синхронизация инстансов из конфига в бд
	// TODO: Добавить инстансы в конфиг

	// Запуск воркера для healthcheck
	// TODO: Написать воркер

	app := app.New(log, storage, cfg.HTTPServer)

	go app.Server.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	app.Server.Stop()

	log.Info("Application stopped")
}
