package unified

import (
	"sort"

	"github.com/UnipayFI/gate-cli/common"
	"github.com/UnipayFI/gate-cli/printer"
)

var _ printer.TableWriter = (*Account)(nil)

// Header for the unified account lists per-currency balances. Account-wide
// totals (equity, margin rate, borrowed) are visible via --json.
func (a *Account) Header() []string {
	return []string{"Currency", "Available", "Freeze", "Borrowed", "Equity", "Margin Balance", "Available Margin", "Collateral"}
}

func (a *Account) Row() [][]any {
	currencies := make([]string, 0, len(a.Balances))
	for currency := range a.Balances {
		currencies = append(currencies, currency)
	}
	sort.Strings(currencies)

	rows := [][]any{}
	for _, currency := range currencies {
		bal := a.Balances[currency]
		rows = append(rows, []any{
			currency, bal.Available, bal.Freeze, bal.Borrowed, bal.Equity,
			bal.MarginBalance, bal.AvailableMargin, bal.EnabledCollateral,
		})
	}
	return rows
}

var _ printer.TableWriter = (*Mode)(nil)

func (m *Mode) Header() []string {
	return []string{"Mode", "USDT Futures", "Spot Hedge", "Use Funding", "Options"}
}

func (m *Mode) Row() [][]any {
	return [][]any{
		{m.Mode, m.Settings.USDTFutures, m.Settings.SpotHedge, m.Settings.UseFunding, m.Settings.Options},
	}
}

var _ printer.TableWriter = (*Borrowable)(nil)

func (b *Borrowable) Header() []string {
	return []string{"Currency", "Max Borrowable"}
}

func (b *Borrowable) Row() [][]any {
	return [][]any{{b.Currency, b.Amount}}
}

var _ printer.TableWriter = (*Transferable)(nil)

func (t *Transferable) Header() []string {
	return []string{"Currency", "Max Transferable"}
}

func (t *Transferable) Row() [][]any {
	return [][]any{{t.Currency, t.Amount}}
}

var _ printer.TableWriter = (*LoanResult)(nil)

func (r *LoanResult) Header() []string {
	return []string{"Transaction ID"}
}

func (r *LoanResult) Row() [][]any {
	return [][]any{{r.TranID}}
}

var _ printer.TableWriter = (*LoanList)(nil)

func (l *LoanList) Header() []string {
	return []string{"Currency", "Currency Pair", "Amount", "Type", "Create Time", "Update Time"}
}

func (l *LoanList) Row() [][]any {
	rows := [][]any{}
	for _, loan := range *l {
		rows = append(rows, []any{
			loan.Currency, loan.CurrencyPair, loan.Amount, loan.Type,
			common.FormatTime(loan.CreateTime), common.FormatTime(loan.UpdateTime),
		})
	}
	return rows
}

var _ printer.TableWriter = (*LoanRecordList)(nil)

func (l *LoanRecordList) Header() []string {
	return []string{"ID", "Type", "Repayment Type", "Borrow Type", "Currency Pair", "Currency", "Amount", "Create Time"}
}

func (l *LoanRecordList) Row() [][]any {
	rows := [][]any{}
	for _, rec := range *l {
		rows = append(rows, []any{
			rec.ID, rec.Type, rec.RepaymentType, rec.BorrowType, rec.CurrencyPair,
			rec.Currency, rec.Amount, common.FormatTime(rec.CreateTime),
		})
	}
	return rows
}

var _ printer.TableWriter = (*InterestRecordList)(nil)

func (l *InterestRecordList) Header() []string {
	return []string{"Currency", "Currency Pair", "Actual Rate", "Interest", "Status", "Type", "Create Time"}
}

func (l *InterestRecordList) Row() [][]any {
	rows := [][]any{}
	for _, rec := range *l {
		rows = append(rows, []any{
			rec.Currency, rec.CurrencyPair, rec.ActualRate, rec.Interest, rec.Status,
			rec.Type, common.FormatTime(rec.CreateTime),
		})
	}
	return rows
}
