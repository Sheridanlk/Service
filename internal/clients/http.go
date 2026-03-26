package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/Sheridanlk/Service/internal/domain/models"
)

const (
	endpointRegister   = "/vpn/register"
	endpointSubCA      = "/vpn/sub-ca"
	endpontServerToken = "/vpn/token"
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

func (c *MainServiceClient) RegisterNode(
	ctx context.Context,
	serverName string,
	location models.Location,
	registerToken string,
) (string, string, error) {
	const op = "clients.http.RegisterNode"

	url, err := url.JoinPath(c.baseURL, endpointRegister)
	if err != nil {
		return "", "", fmt.Errorf("%s: invalid url: %w", op, err)
	}

	reqBody := RegisterRequest{
		Name:              serverName,
		Location:          location,
		RegistrationToken: registerToken,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to marshal request: %w", op, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to create request: %w", op, err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to send request: %w", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("%s: server returned status: %d", op, resp.StatusCode)
	}

	var res RegisterResponse

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", "", err
	}

	if !res.Success {
		return "", "", fmt.Errorf("%s: server error", op)
	}

	return res.Data.ID, res.Data.ServerSecret, nil
}

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
