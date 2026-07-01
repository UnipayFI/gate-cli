package config

import "os"

// Config holds process-wide runtime configuration. Credentials are read from
// the environment at startup; the global --json flag is bound by cmd/root.go.
var Config struct {
	APIKey     string
	APISecret  string
	BaseURL    string
	OutputJSON bool
}

func init() {
	Config.APIKey = os.Getenv("GATE_API_KEY")
	Config.APISecret = os.Getenv("GATE_API_SECRET")
	// Optional override, e.g. to point at a testnet gateway.
	Config.BaseURL = os.Getenv("GATE_API_BASE_URL")
}
