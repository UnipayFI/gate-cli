# Unified Account Module

> Commands here operate on the Gate **unified account**. Exec prefix:
> `./gate-cli unified [Subcommand] [Arguments]`.
>
> Every command accepts a global `--json` flag that prints the raw API response
> as indented JSON instead of a table.

The unified account aggregates equity and margin across products. This module
adds the account-level endpoints the classic account has no equivalent for
(overview, mode, borrowing), and mirrors the spot/futures trading commands so
they route to the unified account without a manual `--account` flag.

## Quick Navigation
- Account-level
  - [Account](#account---show-unified-account-overview)
  - [Mode](#mode---show-the-unified-account-mode)
  - [Borrowable](#borrowable---maximum-borrowable-amount)
  - [Transferable](#transferable---maximum-transferable-amount)
  - [Borrow](#borrow---borrow-a-currency)
  - [Repay](#repay---repay-a-currency)
  - [Loans](#loans---list-outstanding-borrowings)
  - [Loan-records](#loan-records---borrowrepay-history)
  - [Interest-records](#interest-records---interest-deduction-history)
- Trading (routed to the unified account)
  - [Spot](#spot)
  - [Futures](#futures)

## Account - Show unified account overview
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#get-unified-account-information>

Exec: `./gate-cli unified account [--currency USDT]`

The table lists per-currency balances (available, freeze, borrowed, equity,
margin balance, available margin, collateral). Account-wide totals — total
equity, initial/maintenance margin and rate, borrowed — are in the `--json`
output.

## Mode - Show the unified account mode
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#query-mode-of-the-unified-account>

Exec: `./gate-cli unified mode`

Reports the account mode (`classic`, `single_currency`, `multi_currency` or
`portfolio`) and its feature toggles (usdt_futures, spot_hedge, use_funding,
options). This is read-only; switching modes is out of scope for this phase.

## Borrowable - Maximum borrowable amount
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#get-the-maximum-borrowable-amount-of-unified-account>

Exec: `./gate-cli unified borrowable -c BTC`

## Transferable - Maximum transferable amount
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#get-the-maximum-transferable-amount-for-a-specific-currency-of-unified-account>

Exec: `./gate-cli unified transferable -c BTC`

## Borrow - Borrow a currency
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#borrow-or-repay>

Exec: `./gate-cli unified borrow -c USDT -a 100`

## Repay - Repay a currency
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#borrow-or-repay>

Exec (partial): `./gate-cli unified repay -c USDT -a 100`

Exec (full outstanding balance): `./gate-cli unified repay -c USDT --all`

## Loans - List outstanding borrowings
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#list-loans>

Exec: `./gate-cli unified loans [-c USDT] [--type platform|margin] [--limit 100] [--page 1]`

## Loan-records - Borrow/repay history
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#get-loan-records>

Exec: `./gate-cli unified loan-records [-c USDT] [--type borrow|repay] [--limit 100] [--page 1]`

## Interest-records - Interest-deduction history
Docs Link: <https://www.gate.io/docs/developers/apiv4/en/#interest-records-for-the-unified-account>

Exec: `./gate-cli unified interest-records [-c USDT] [--type platform|margin] [--limit 100] [--page 1]`

Time flags `--from` / `--to` accept a Unix-seconds timestamp or a
`"YYYY-MM-DD HH:MM:SS"` datetime.

## Spot
Exec: `./gate-cli unified spot [Subcommand] [Arguments]`

Mirrors the classic spot `order` and `trade` commands, but every operation is
pinned to the unified account (`account=unified`), so there is no `--account`
flag. Flags and behaviour are otherwise identical to the classic spot commands —
see [spot.md](spot.md).

```shell
Available Commands:
  order       Create, cancel, list and query spot orders
  trade       Query personal trading history
```

Exec example: `./gate-cli unified spot order create -p BTC_USDT -S buy -q 0.001 --price 10000`

## Futures
Exec: `./gate-cli unified futures [Subcommand] [Arguments]`

Mirrors the classic futures `order`, `trade` and `position` commands. Futures
orders carry no account field — whether they route to the unified account is
decided server-side by the account mode — so these issue the exact same requests
as `./gate-cli classic futures ...`. Accepts the same `--settle` flag. See
[futures.md](futures.md).

```shell
Available Commands:
  order       Create, cancel, list and query futures orders
  position    Position management (list, get, margin, leverage, dual-mode, close-history)
  trade       Query personal trading history
```
