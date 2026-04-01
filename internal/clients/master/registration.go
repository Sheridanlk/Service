package master

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Sheridanlk/Service/internal/domain/models"
)

const (
	endpointRegister = "/vpn/register"
)

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
