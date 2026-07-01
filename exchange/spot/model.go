package spot

import (
	"github.com/UnipayFI/gate-cli/common"
	"github.com/UnipayFI/gate-cli/printer"
)

var _ printer.TableWriter = (*AccountList)(nil)

func (a *AccountList) Header() []string {
	return []string{"Currency", "Available", "Locked"}
}

func (a *AccountList) Row() [][]any {
	rows := [][]any{}
	for _, acc := range *a {
		rows = append(rows, []any{acc.Currency, acc.Available, acc.Locked})
	}
	return rows
}

var _ printer.TableWriter = (*Fee)(nil)

func (f *Fee) Header() []string {
	return []string{"Currency Pair", "Maker Fee", "Taker Fee", "GT Discount", "GT Maker Fee", "GT Taker Fee"}
}

func (f *Fee) Row() [][]any {
	return [][]any{
		{f.CurrencyPair, f.MakerFee, f.TakerFee, f.GTDiscount, f.GTMakerFee, f.GTTakerFee},
	}
}

var _ printer.TableWriter = (*OrderList)(nil)

func (o *OrderList) Header() []string {
	return []string{"ID", "Pair", "Account", "Side", "Type", "Status", "Price", "Amount", "Left", "Filled", "Avg Price", "TIF", "Create Time"}
}

func (o *OrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *o {
		rows = append(rows, []any{
			order.ID, order.CurrencyPair, order.Account, order.Side, order.Type, order.Status,
			order.Price, order.Amount, order.Left, order.FilledAmount, order.AvgDealPrice,
			order.TimeInForce, common.FormatTime(order.CreateTime),
		})
	}
	return rows
}

var _ printer.TableWriter = (*TradeList)(nil)

func (t *TradeList) Header() []string {
	return []string{"ID", "Order ID", "Pair", "Side", "Role", "Price", "Amount", "Fee", "Create Time"}
}

func (t *TradeList) Row() [][]any {
	rows := [][]any{}
	for _, trade := range *t {
		fee := trade.Fee.String()
		if trade.FeeCurrency != "" {
			fee = trade.Fee.String() + " " + trade.FeeCurrency
		}
		rows = append(rows, []any{
			trade.ID, trade.OrderID, trade.CurrencyPair, trade.Side, trade.Role,
			trade.Price, trade.Amount, fee, common.FormatTime(trade.CreateTime),
		})
	}
	return rows
}

var _ printer.TableWriter = (*AccountBookList)(nil)

func (a *AccountBookList) Header() []string {
	return []string{"ID", "Currency", "Change", "Balance", "Type", "Text", "Time"}
}

func (a *AccountBookList) Row() [][]any {
	rows := [][]any{}
	for _, rec := range *a {
		rows = append(rows, []any{
			rec.ID, rec.Currency, rec.Change, rec.Balance, rec.Type, rec.Text, common.FormatTime(rec.Time),
		})
	}
	return rows
}
