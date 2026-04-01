package master

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Sheridanlk/Service/internal/domain/models"
)

const (
	endpontServerToken = "/vpn/token"
)

func (c *MainServiceClient) GetTokens(
	ctx context.Context,
	authType string,
	tokenValue string,
) (*models.AuthTokens, error) {
	const op = "clients.http.GetTokens"

	u, err := url.Parse(c.baseURL)
	if err != nil {
		return nil, fmt.Errorf("%s: invalid url: %w", op, err)
	}

	u.Path, err = url.JoinPath(u.Path, endpontServerToken)
	if err != nil {
		return nil, fmt.Errorf("%s: invalid url: %w", op, err)
	}

	q := u.Query()
	q.Set("type", authType)
	q.Add(authType, tokenValue)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to create request: %w", op, err)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to send request: %w", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: server returned status: %d", op, resp.StatusCode)
	}

	var res TokenResponse

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, fmt.Errorf("%s: server error", op)
	}

	return &res.Tokens, nil
}
