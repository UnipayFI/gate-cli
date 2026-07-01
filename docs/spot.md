# Spot Module

> Every command in this module accepts a global `--json` flag that prints the
> raw API response as indented JSON instead of a table.

## Quick Navigation
- [Account](#account---show-spot-account-balances)
- [Fee](#fee---show-trading-fee-rates)
- [Account-book](#account-book---query-account-balance-change-history)
- [Order](#order)
  - [Create Order](#order-create---create-an-order)
  - [List Orders](#order-list---list-orders)
  - [List Open Orders](#order-open---list-all-open-orders)
  - [Get Order](#order-get---query-a-single-order)
  - [Cancel Order](#order-cancel---cancel-an-order)
- [Trade](#trade---query-personal-trading-history)

## Account - Show spot account balances
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#list-spot-accounts>

Exec: `./gate-cli spot account [--currency USDT]`
```shell
┌──────────┬──────────────────┬────────┐
│ CURRENCY │    AVAILABLE     │ LOCKED │
├──────────┼──────────────────┼────────┤
│ USDT     │ 950.000000004797 │ 0      │
└──────────┴──────────────────┴────────┘
```

## Fee - Show trading fee rates
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#query-user-trading-fee-rates>

Exec: `./gate-cli spot fee --currency-pair BTC_USDT`

## Account-book - Query account balance change history
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#query-account-book>

Exec: `./gate-cli spot account-book [--currency USDT] [--type trade] [--limit 100]`

Time flags `--from` / `--to` accept a Unix-seconds timestamp or a
`"YYYY-MM-DD HH:MM:SS"` datetime.

## Order
Exec: `./gate-cli spot order [Subcommand] [Arguments]`

The `--currency-pair` flag is shared across all `order` subcommands.

### order create - Create an order
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#create-an-order>

Exec: `./gate-cli spot order create -p BTC_USDT -S buy -q 0.001 --price 10000`
```shell
┌───────────────┬──────────┬─────────┬──────┬───────┬────────┬───────┬────────┬───────┬────────┬───────────┬─────┬─────────────────────┐
│      ID       │   PAIR   │ ACCOUNT │ SIDE │ TYPE  │ STATUS │ PRICE │ AMOUNT │ LEFT  │ FILLED │ AVG PRICE │ TIF │     CREATE TIME     │
├───────────────┼──────────┼─────────┼──────┼───────┼────────┼───────┼────────┼───────┼────────┼───────────┼─────┼─────────────────────┤
│ 1092351564349 │ BTC_USDT │ spot    │ buy  │ limit │ open   │ 10000 │ 0.001  │ 0.001 │ 0      │ 0         │ gtc │ 2026-07-01 14:21:38 │
└───────────────┴──────────┴─────────┴──────┴───────┴────────┴───────┴────────┴───────┴────────┴───────────┴─────┴─────────────────────┘
```
Flags: `--side` (buy/sell), `--amount`, `--type` (limit/market), `--price`,
`--time-in-force` (gtc/ioc/poc/fok), `--account` (spot/margin/cross_margin/unified),
`--text`.

### order list - List orders
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#list-orders>

Exec: `./gate-cli spot order list -p BTC_USDT --status open`

### order open - List all open orders
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#list-all-open-orders>

Exec: `./gate-cli spot order open`

Open orders across every currency pair, flattened into one table.

### order get - Query a single order
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#get-a-single-order>

Exec: `./gate-cli spot order get -p BTC_USDT -i 1092351564349`

### order cancel - Cancel an order
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#cancel-a-single-order>

Exec (single): `./gate-cli spot order cancel -p BTC_USDT -i 1092351564349`

Exec (all open orders, optionally scoped to a pair): `./gate-cli spot order cancel [-p BTC_USDT]`

## Trade - Query personal trading history
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#list-personal-trading-history>

Exec: `./gate-cli spot trade -p BTC_USDT [--order-id 123] [--limit 100]`
