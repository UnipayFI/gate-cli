package exchange

import (
	"github.com/UnipayFI/gate-cli/config"
	gate "github.com/UnipayFI/go-gate/v4"
	"github.com/UnipayFI/go-gate/v4/client"
	"github.com/UnipayFI/go-gate/v4/futures"
	"github.com/UnipayFI/go-gate/v4/spot"
	"github.com/UnipayFI/go-gate/v4/unified"
)

// Client carries the API credentials and constructs per-product SDK clients on
// demand. Gate signs private requests with HMAC-SHA512 over the API key/secret.
type Client struct {
	APIKey    string
	APISecret string
}

func NewClient(apiKey, apiSecret string) *Client {
	return &Client{APIKey: apiKey, APISecret: apiSecret}
}

func (c *Client) baseOptions() []client.Options {
	opts := []client.Options{client.WithAuth(c.APIKey, c.APISecret)}
	if config.Config.BaseURL != "" {
		opts = append(opts, client.WithBaseURL(config.Config.BaseURL))
	}
	if config.Config.Proxy != "" {
		opts = append(opts, client.WithProxy(config.Config.Proxy))
	}
	return opts
}

func (c *Client) NewSpotClient() *spot.SpotClient {
	return gate.NewSpotClient(c.baseOptions()...)
}

func (c *Client) NewFuturesClient() *futures.FuturesClient {
	return gate.NewFuturesClient(c.baseOptions()...)
}

func (c *Client) NewUnifiedClient() *unified.UnifiedClient {
	return gate.NewUnifiedClient(c.baseOptions()...)
}
