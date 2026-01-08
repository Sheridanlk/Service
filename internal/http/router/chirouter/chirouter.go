package chirouter

import (
	"log/slog"
	"net/http"

	"github.com/Sheridanlk/Service/internal/storage/postgresql"
	"github.com/go-chi/chi"
)

func Setup(log *slog.Logger, storage *postgresql.Storage) http.Handler {
	router := chi.NewRouter()

	return router
}
