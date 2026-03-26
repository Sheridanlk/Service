package clients

import (
	"github.com/Sheridanlk/Service/internal/domain/models"
)

type RegisterRequest struct {
	Name              string          `json:"name"`
	Location          models.Location `json:"location"`
	RegistrationToken string          `json:"registration_token"`
}

type RegisterResponse struct {
	Success bool    `json:"success"`
	Data    RegData `json:"data"`
}

type RegData struct {
	ID           string `json:"id"`
	ServerSecret string `json:"server_secret"`
}

type TokenResponse struct {
	Success bool              `json:"success"`
	Tokens  models.AuthTokens `json:"data"`
}
