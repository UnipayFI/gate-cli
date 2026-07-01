package config

import "os"

// Config holds process-wide runtime configuration. Credentials are read from
// the environment at startup; the global --json flag is bound by cmd/root.go.
var Config struct {
	APIKey     string
	APISecret  string
	BaseURL    string
	Proxy      string
	OutputJSON bool
}

func init() {
	Config.APIKey = os.Getenv("GATE_API_KEY")
	Config.APISecret = os.Getenv("GATE_API_SECRET")
	// Optional override, e.g. to point at a testnet gateway.
	Config.BaseURL = os.Getenv("GATE_API_BASE_URL")
	// Optional proxy for all REST traffic. Supports http, https, socks5, and
	// socks5h schemes (with optional user:pass@ auth), e.g.
	// "http://127.0.0.1:7890". Accepts the conventional upper- and lower-case
	// HTTPS_PROXY names, preferring the upper-case form.
	Config.Proxy = firstEnv("HTTPS_PROXY", "https_proxy")
}

// firstEnv returns the value of the first environment variable that is set to a
// non-empty value, or "" if none are.
func firstEnv(keys ...string) string {
	for _, k := range keys {
		if v := os.Getenv(k); v != "" {
			return v
		}
	}
	return ""
}
