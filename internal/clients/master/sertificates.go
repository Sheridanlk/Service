package master

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const endpointSubCa = "/vpn/sub-ca"

type SubCaResponse struct {
	Success bool      `json:"success"`
	Data    SubCaData `json:"data"`
}

type SubCaData struct {
	PublicKey  string `json:"public_part"`
	PrivateKey string `json:"private_part"`
}

func (c *MainServiceClient) GetSubCa(
	ctx context.Context,
	accessToken string,
) (string, string, error) {
	const op = "clients.master.GetSubCa"

	url, err := url.JoinPath(c.baseURL, endpointSubCa)
	if err != nil {
		return "", "", fmt.Errorf("%s: invalid url: %w", op, err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to create request: %w", op, err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := c.http.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("%s: failed to send request: %w", op, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("%s: server returned status: %d", op, resp.StatusCode)
	}

	var res SubCaResponse

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", "", err
	}

	if !res.Success {
		return "", "", fmt.Errorf("%s: server error", op)
	}

	return res.Data.PublicKey, res.Data.PrivateKey, nil
}
