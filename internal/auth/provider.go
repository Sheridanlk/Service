package auth

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Sheridanlk/Service/internal/clients/master"
	"github.com/Sheridanlk/Service/internal/domain/models"
	"github.com/Sheridanlk/Service/internal/state"
)

type TokenGetter interface {
	GetTokens(ctx context.Context, authType string, tokenValue string) (models.AuthTokens, error)
}

type Provider struct {
	mu          sync.Mutex
	tokenGetter TokenGetter
	state       *state.State
}

func New(tokenGetter TokenGetter, state *state.State) *Provider {
	return &Provider{
		tokenGetter: tokenGetter,
		state:       state,
	}
}

func (p *Provider) GetValidToken(ctx context.Context) (string, error) {
	const op = "auth.Provider.GetValidToken"

	tokens := p.state.GetTokens()

	if tokens.AccessToken != "" && !isTokenExpired(tokens.AccessExpiresAt) {
		return tokens.AccessToken, nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	tokens = p.state.GetTokens()

	update, authType := tokens.RefreshToken, master.AuthTypeRefresh

	if tokens.RefreshToken == "" || isTokenExpired(tokens.RefreshExpiresAt) {
		update, authType = p.state.GetServerSecret(), master.AuthTypeSecret
		if update == "" {
			return "", fmt.Errorf("%s: no valid authentication method available", op)
		}
	}
	newTokens, err := p.tokenGetter.GetTokens(ctx, authType, update)
	if err != nil {
		return "", fmt.Errorf("%s: failed to update tokens: %w", op, err)
	}

	token := newTokens.AccessToken
	if token == "" {
		return "", fmt.Errorf("%s: received empty access token", op)
	}

	p.state.UpdateTokens(newTokens)

	return token, nil
}

func isTokenExpired(expiresAt time.Time) bool {
	return time.Now().After(expiresAt)
}
