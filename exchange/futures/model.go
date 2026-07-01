package futures

import (
	"github.com/UnipayFI/gate-cli/common"
	"github.com/UnipayFI/gate-cli/printer"
)

var _ printer.TableWriter = (*Account)(nil)

func (a *Account) Header() []string {
	return []string{"Currency", "Total", "Available", "Unrealised PnL", "Position Margin", "Order Margin", "Point", "Dual Mode"}
}

func (a *Account) Row() [][]any {
	return [][]any{
		{a.Currency, a.Total, a.Available, a.UnrealisedPnL, a.PositionMargin, a.OrderMargin, a.Point, a.InDualMode},
	}
}

var _ printer.TableWriter = (*FeeList)(nil)

func (f *FeeList) Header() []string {
	return []string{"Contract", "Maker Fee", "Taker Fee"}
}

func (f *FeeList) Row() [][]any {
	rows := [][]any{}
	for _, fee := range *f {
		rows = append(rows, []any{fee.Contract, fee.MakerFee, fee.TakerFee})
	}
	return rows
}

var _ printer.TableWriter = (*AccountBookList)(nil)

func (a *AccountBookList) Header() []string {
	return []string{"Contract", "Change", "Balance", "Type", "Text", "Time"}
}

func (a *AccountBookList) Row() [][]any {
	rows := [][]any{}
	for _, rec := range *a {
		rows = append(rows, []any{rec.Contract, rec.Change, rec.Balance, rec.Type, rec.Text, common.FormatTime(rec.Time)})
	}
	return rows
}

var _ printer.TableWriter = (*OrderList)(nil)

func (o *OrderList) Header() []string {
	return []string{"ID", "Contract", "Size", "Price", "Left", "Fill Price", "Status", "Finish As", "TIF", "Create Time"}
}

func (o *OrderList) Row() [][]any {
	rows := [][]any{}
	for _, order := range *o {
		rows = append(rows, []any{
			order.ID, order.Contract, order.Size, order.Price, order.Left, order.FillPrice,
			order.Status, order.FinishAs, order.Tif, common.FormatTime(order.CreateTime),
		})
	}
	return rows
}

var _ printer.TableWriter = (*TradeList)(nil)

func (t *TradeList) Header() []string {
	return []string{"ID", "Order ID", "Contract", "Size", "Price", "Role", "Fee", "Create Time"}
}

func (t *TradeList) Row() [][]any {
	rows := [][]any{}
	for _, trade := range *t {
		rows = append(rows, []any{
			trade.ID, trade.OrderID, trade.Contract, trade.Size, trade.Price, trade.Role,
			trade.Fee, common.FormatTime(trade.CreateTime),
		})
	}
	return rows
}

var _ printer.TableWriter = (*PositionList)(nil)

func (p *PositionList) Header() []string {
	return []string{"Contract", "Mode", "Size", "Leverage", "Entry Price", "Mark Price", "Liq Price", "Margin", "Value", "Unrealised PnL", "Realised PnL"}
}

func (p *PositionList) Row() [][]any {
	rows := [][]any{}
	for _, pos := range *p {
		rows = append(rows, []any{
			pos.Contract, pos.Mode, pos.Size, pos.Leverage, pos.EntryPrice, pos.MarkPrice,
			pos.LiqPrice, pos.Margin, pos.Value, pos.UnrealisedPnL, pos.RealisedPnL,
		})
	}
	return rows
}

var _ printer.TableWriter = (*PositionCloseList)(nil)

func (p *PositionCloseList) Header() []string {
	return []string{"Contract", "Side", "PnL", "Max Size", "Accum Size", "Long Price", "Short Price", "First Open Time", "Time"}
}

func (p *PositionCloseList) Row() [][]any {
	rows := [][]any{}
	for _, rec := range *p {
		rows = append(rows, []any{
			rec.Contract, rec.Side, rec.PnL, rec.MaxSize, rec.AccumSize, rec.LongPrice, rec.ShortPrice,
			common.FormatTime(rec.FirstOpenTime), common.FormatTime(rec.Time),
		})
	}
	return rows
}
