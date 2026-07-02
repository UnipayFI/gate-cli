# Futures Module

> Commands here run under the **classic** account. The unified account mirrors
> them at `./gate-cli unified futures ...` (identical requests, since futures
> orders carry no account field; see [unified.md](unified.md)).
>
> Every command in this module accepts a global `--json` flag that prints the
> raw API response as indented JSON instead of a table.
>
> Every command also accepts a global `--settle` flag (default `usdt`, also
> `btc`) selecting the settle currency.

## Quick Navigation
- [Account](#account---show-futures-account-summary)
- [Fee](#fee---show-trading-fee-rates)
- [Account-book](#account-book---query-account-balance-change-history)
- [Order](#order)
  - [Create Order](#order-create---create-a-futures-order)
  - [List Orders](#order-list---list-futures-orders)
  - [Get Order](#order-get---query-a-single-order)
  - [Cancel Order](#order-cancel---cancel-an-order)
- [Trade](#trade---query-personal-trading-history)
- [Position](#position)

## Account - Show futures account summary
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#query-futures-account>

Exec: `./gate-cli classic futures account [--settle usdt]`
```shell
┌──────────┬────────────────┬────────────────┬─────────────────┬─────────────────┬──────────────┬───────┬───────────┐
│ CURRENCY │     TOTAL      │   AVAILABLE    │ UNREALISED PN L │ POSITION MARGIN │ ORDER MARGIN │ POINT │ DUAL MODE │
├──────────┼────────────────┼────────────────┼─────────────────┼─────────────────┼──────────────┼───────┼───────────┤
│ USDT     │ 50.00000000112 │ 50.00000000112 │ 0               │ 0               │ 0            │ 0     │ false     │
└──────────┴────────────────┴────────────────┴─────────────────┴─────────────────┴──────────────┴───────┴───────────┘
```

## Fee - Show trading fee rates
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#query-user-trading-fee-rates-2>

Exec: `./gate-cli classic futures fee [--contract BTC_USDT]`

## Account-book - Query account balance change history
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#query-account-book-2>

Exec: `./gate-cli classic futures account-book [--contract BTC_USDT] [--type pnl] [--limit 100]`

## Order
Exec: `./gate-cli classic futures order [Subcommand] [Arguments]`

### order create - Create a futures order
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#create-a-futures-order>

`--size` is signed and denominated in contracts: positive = long, negative =
short. Pass `--price` for a limit order; omit price with `--tif ioc` for a
market order. Use `--close` (with size 0) to close a position.

Exec: `./gate-cli classic futures order create -c BTC_USDT -z 1 --price 10000`
```shell
┌───────────────────┬──────────┬──────┬───────┬──────┬────────────┬────────┬───────────┬─────┬─────────────────────┐
│        ID         │ CONTRACT │ SIZE │ PRICE │ LEFT │ FILL PRICE │ STATUS │ FINISH AS │ TIF │     CREATE TIME     │
├───────────────────┼──────────┼──────┼───────┼──────┼────────────┼────────┼───────────┼─────┼─────────────────────┤
│ 36028832624125599 │ BTC_USDT │ 1    │ 10000 │ 1    │ 0          │ open   │           │ gtc │ 2026-07-01 14:22:00 │
└───────────────────┴──────────┴──────┴───────┴──────┴────────────┴────────┴───────────┴─────┴─────────────────────┘
```

### order list - List futures orders
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#list-futures-orders>

Exec: `./gate-cli classic futures order list --status open [--contract BTC_USDT]`

### order get - Query a single order
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#get-a-single-order-3>

Exec: `./gate-cli classic futures order get -i 36028832624125599`

### order cancel - Cancel an order
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#cancel-a-single-order-3>

Exec (single): `./gate-cli classic futures order cancel -i 36028832624125599`

Exec (all open orders of a contract): `./gate-cli classic futures order cancel -c BTC_USDT`

## Trade - Query personal trading history
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#list-personal-trading-history-2>

Exec: `./gate-cli classic futures trade [--contract BTC_USDT] [--order-id 123] [--limit 100]`

## Position
Exec: `./gate-cli classic futures position [Subcommand] [Arguments]`

```shell
Available Commands:
  list           List positions
  get            Get a single position
  margin         Update isolated position margin
  leverage       Update position leverage
  dual-mode      Enable or disable hedge (dual) mode
  close-history  Query position close history
```

### position list
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#list-all-positions-of-a-user>

Exec: `./gate-cli classic futures position list [--holding]`
```shell
┌──────────┬────────┬──────┬──────────┬─────────────┬────────────┬───────────┬────────┬───────┬─────────────────┬───────────────┐
│ CONTRACT │  MODE  │ SIZE │ LEVERAGE │ ENTRY PRICE │ MARK PRICE │ LIQ PRICE │ MARGIN │ VALUE │ UNREALISED PN L │ REALISED PN L │
├──────────┼────────┼──────┼──────────┼─────────────┼────────────┼───────────┼────────┼───────┼─────────────────┼───────────────┤
│ ETH_USDT │ single │ 0    │ 1        │ 0           │ 1611.58    │ 0         │ 0      │ 0     │ 0               │ 0             │
└──────────┴────────┴──────┴──────────┴─────────────┴────────────┴───────────┴────────┴───────┴─────────────────┴───────────────┘
```

### position get
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#get-single-position>

Exec: `./gate-cli classic futures position get -c BTC_USDT`

### position margin
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#update-position-margin>

`--change` is signed: positive adds margin, negative reduces it.

Exec: `./gate-cli classic futures position margin -c BTC_USDT --change 10`

### position leverage
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#update-position-leverage>

Use `0` for cross margin, otherwise the isolated leverage multiple.

Exec: `./gate-cli classic futures position leverage -c BTC_USDT --leverage 10`

### position dual-mode
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#enable-or-disable-dual-mode>

Exec: `./gate-cli classic futures position dual-mode --enable=true`

### position close-history
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#list-position-close-history>

Exec: `./gate-cli classic futures position close-history [--contract BTC_USDT] [--limit 100]`
