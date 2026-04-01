package master

import (
	"log/slog"
	"net/http"
	"time"
)

type MainServiceClient struct {
	log     *slog.Logger
	baseURL string
	http    *http.Client
}

func New(log *slog.Logger, baseURL string, timeout time.Duration) *MainServiceClient {
	return &MainServiceClient{
		log:     log,
		baseURL: baseURL,
		http: &http.Client{
			Timeout: timeout,
		},
	}
}
