package spot

import (
	"github.com/UnipayFI/gate-cli/exchange"
)

// Client wraps the shared exchange.Client and exposes spot-specific helpers.
type Client struct {
	*exchange.Client
}
