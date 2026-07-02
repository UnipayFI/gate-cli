package unified

import (
	"github.com/UnipayFI/gate-cli/exchange"
)

// Client wraps the shared exchange.Client and exposes unified-account helpers.
type Client struct {
	*exchange.Client
}
