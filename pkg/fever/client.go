package fever

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/infituit/freshrss-cli/internal/config"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	endpoint   string
	apiKey     string
	httpClient HTTPClient
}

type AuthResponse struct {
	APIVersion          int    `json:"api_version"`
	Auth                int    `json:"auth"`
	LastRefreshedOnTime string `json:"last_refreshed_on_time"`
}

func NewClient(cfg config.Config, httpClient HTTPClient) (*Client, error) {
	if strings.TrimSpace(cfg.BaseURL) == "" {
		return nil, fmt.Errorf("FreshRSS URL is required")
	}

	apiKey := strings.TrimSpace(cfg.APIKey)
	if apiKey == "" {
		if strings.TrimSpace(cfg.Username) == "" || cfg.APIPassword == "" {
			return nil, fmt.Errorf("Fever API key or username and API password are required")
		}
		apiKey = APIKey(cfg.Username, cfg.APIPassword)
	}

	endpoint, err := feverEndpoint(cfg.BaseURL)
	if err != nil {
		return nil, err
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{endpoint: endpoint, apiKey: apiKey, httpClient: httpClient}, nil
}

func APIKey(username, apiPassword string) string {
	sum := md5.Sum([]byte(username + ":" + apiPassword))
	return hex.EncodeToString(sum[:])
}

func (c *Client) AuthCheck(ctx context.Context) (AuthResponse, error) {
	var auth AuthResponse
	if err := c.post(ctx, nil, &auth); err != nil {
		return AuthResponse{}, err
	}
	return auth, nil
}

func (c *Client) post(ctx context.Context, query url.Values, out any) error {
	endpoint, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	values := endpoint.Query()
	values.Set("api", "")
	for key, entries := range query {
		for _, entry := range entries {
			values.Add(key, entry)
		}
	}
	endpoint.RawQuery = values.Encode()

	form := url.Values{}
	form.Set("api_key", c.apiKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("Fever API returned HTTP %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("decode Fever API response: %w", err)
	}
	return nil
}

func feverEndpoint(rawURL string) (string, error) {
	parsed, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil {
		return "", err
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return "", fmt.Errorf("FreshRSS URL must include scheme and host")
	}
	if !strings.HasSuffix(parsed.Path, "/api/fever.php") {
		parsed.Path = strings.TrimRight(parsed.Path, "/") + "/api/fever.php"
	}
	return parsed.String(), nil
}
