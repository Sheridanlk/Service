package app

import (
	"net/http"

	"github.com/Sheridanlk/Service/internal/storage/postgresql"
)

type app struct {
	Server  *http.Server
	Storage *postgresql.Storage
}
