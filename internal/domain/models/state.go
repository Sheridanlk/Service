package models

import "time"

type NodeState struct {
	NodeID       string     `yaml:"node_id"`
	ServerSecret string     `yaml:"server_secret"`
	Tokens       AuthTokens `yaml:"tokens"`
}

type AuthTokens struct {
	AccessToken      string    `yaml:"access_token" json:"access_token"`
	RefreshToken     string    `yaml:"refresh_token" json:"refresh_token"`
	AccessExpiresAt  time.Time `yaml:"access_expires_at" json:"access_expires_at"`
	RefreshExpiresAt time.Time `yaml:"refresh_expires_at" json:"refresh_expires_at"`
}
