# Gate CLI

A command-line tool for the Gate.io (Gate.com) v4 API developed in Go,
supporting both the **classic** and **unified** account, across spot and
futures (perpetual) trading.

Built on the Gate v4 API (APIv4 `KEY` / `SIGN` / `Timestamp` HMAC-SHA512
signing). This phase covers the main **private** (API-key) endpoints for spot,
futures and the unified account; public market-data endpoints are out of scope.

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
# Optional. Route all REST traffic through a proxy. Supports http, https,
# socks5, and socks5h schemes, with optional user:pass@ auth.
export HTTPS_PROXY="http://127.0.0.1:7890"
```

### Output format
Every command supports a global `--json` flag. Without it, results render as a
table; with it, the raw API response is printed as indented JSON, e.g.:
```shell
./gate-cli classic spot account --json
```

## Usage
All commands follow this format:
```
./gate-cli [Account] [Module] [Subcommand] [Arguments]

Available Commands:
  classic     Classic account commands (spot, futures)
  help        Help about any command
  unified     Unified account commands
```

The top level selects the **account type**: `classic` (independent spot and
futures accounts) or `unified` (aggregated equity/margin, plus borrowing).

Each leaf subcommand's `-h` output includes a `Docs Link:` pointing to the
official Gate.io API documentation page for that endpoint.

### Classic Account
Exec: `./gate-cli classic [Module] [Subcommand] [Arguments]` — independent spot
and futures accounts.

#### Spot
Exec: `./gate-cli classic spot [Subcommand] [Arguments]`
```shell
Available Commands:
  account       Show spot account balances
  account-book  Query account balance change history
  fee           Show trading fee rates
  order         Create, cancel, list and query spot orders
  trade         Query personal trading history
```
**[View detailed documentation](docs/spot.md)**

#### Futures
Exec: `./gate-cli classic futures [Subcommand] [Arguments]`

All futures commands accept a `--settle` flag (default `usdt`, also `btc`)
selecting the settle currency.
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

### Unified Account
Exec: `./gate-cli unified [Subcommand] [Arguments]` — aggregated equity/margin
across products, plus borrowing. The spot/futures trading commands are mirrored
here and route to the unified account automatically (spot pins `account=unified`;
futures issue the same requests as classic, the mode being decided server-side).
```shell
Available Commands:
  account           Show unified account overview
  mode              Show the unified account mode (read-only)
  borrowable        Show the maximum borrowable amount of a currency
  transferable      Show the maximum transferable amount of a currency
  borrow            Borrow a currency
  repay             Repay a borrowed currency
  loans             List outstanding borrowings
  loan-records      List borrow/repay history
  interest-records  List interest-deduction history
  spot              Spot trading (account=unified)
  futures           Futures trading
```
**[View detailed documentation](docs/unified.md)**
