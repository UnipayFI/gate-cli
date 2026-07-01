package futures

import (
	"github.com/UnipayFI/gate-cli/exchange"
	"github.com/UnipayFI/go-gate/v4/futures"
)

// Client wraps the shared exchange.Client and carries the settle currency
// (usdt / btc) that every futures endpoint is scoped to.
type Client struct {
	*exchange.Client
	Settle futures.Settle
}

func (c *Client) futuresClient() *futures.FuturesClient {
	return c.NewFuturesClient()
}
