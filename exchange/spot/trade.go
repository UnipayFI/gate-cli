package spot

import (
	"context"
	"time"

	"github.com/UnipayFI/go-gate/v4/spot"
)

// GetMyTrades returns the personal trading history for a currency pair.
func (c *Client) GetMyTrades(currencyPair, orderID, account string, from, to time.Time, limit int) (*TradeList, error) {
	service := c.NewSpotClient().NewListMyTradesService()
	if currencyPair != "" {
		service.SetCurrencyPair(currencyPair)
	}
	if orderID != "" {
		service.SetOrderID(orderID)
	}
	if account != "" {
		service.SetAccount(spot.Account(account))
	}
	if !from.IsZero() {
		service.SetFrom(from)
	}
	if !to.IsZero() {
		service.SetTo(to)
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

type TradeList []spot.MyTrade
