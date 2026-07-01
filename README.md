# Gate CLI

A command-line tool for the Gate.io (Gate.com) v4 API developed in Go,
supporting spot and futures (perpetual) trading.

Built on the Gate v4 API (APIv4 `KEY` / `SIGN` / `Timestamp` HMAC-SHA512
signing). This first phase covers the main **private** (API-key) spot and
futures endpoints; public market-data endpoints are out of scope.

## Installation and Configuration

### Build
```shell
go build -o gate-cli .
```

### Environment variables
Before using, set your Gate.io API-key credentials (from the API-keys page):
```shell
export GATE_API_KEY="your-api-key"
export GATE_API_SECRET="your-api-secret"
# Optional. Override the REST host (e.g. to point at a proxy).
export GATE_API_BASE_URL=""
```

### Output format
Every command supports a global `--json` flag. Without it, results render as a
table; with it, the raw API response is printed as indented JSON, e.g.:
```shell
./gate-cli spot account --json
```

## Usage
All commands follow this format:
```
./gate-cli [Module] [Subcommand] [Arguments]

Available Commands:
  futures     Futures (perpetual) trading commands
  help        Help about any command
  spot        Spot trading commands
```

Each leaf subcommand's `-h` output includes a `Docs Link:` pointing to the
official Gate.io API documentation page for that endpoint.

### Spot Module
Exec: `./gate-cli spot [Subcommand] [Arguments]`
```shell
Available Commands:
  account       Show spot account balances
  account-book  Query account balance change history
  fee           Show trading fee rates
  order         Create, cancel, list and query spot orders
  trade         Query personal trading history
```
**[View detailed documentation](docs/spot.md)**

### Futures Module
Exec: `./gate-cli futures [Subcommand] [Arguments]`

All futures commands accept a global `--settle` flag (default `usdt`, also
`btc`) selecting the settle currency.
```shell
Available Commands:
  account       Show futures account summary
  account-book  Query account balance change history
  fee           Show trading fee rates
  order         Create, cancel, list and query futures orders
  position      Position management (list, get, margin, leverage, dual-mode, close-history)
  trade         Query personal trading history
```
**[View detailed documentation](docs/futures.md)**
