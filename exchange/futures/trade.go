package futures

import (
	"context"

	"github.com/UnipayFI/go-gate/v4/futures"
)

// GetMyTrades returns the personal trading history for the settle currency.
func (c *Client) GetMyTrades(contract string, orderID int64, limit int) (*TradeList, error) {
	service := c.futuresClient().NewGetMyTradesService(c.Settle)
	if contract != "" {
		service.SetContract(contract)
	}
	if orderID != 0 {
		service.SetOrder(orderID)
	}
	if limit != 0 {
		service.SetLimit(limit)
	}
	trades, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := TradeList(trades)
	return &list, nil
}

type TradeList []futures.FuturesMyTrade
