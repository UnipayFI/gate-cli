package spot

import (
	"context"
	"strings"
	"time"

	"github.com/UnipayFI/go-gate/v4/spot"
	"github.com/shopspring/decimal"
)

// CreateOrder places a new spot order from the long-name keyed param map that
// cmd/spot.buildOrderParams produces.
func (c *Client) CreateOrder(params map[string]string) (*spot.Order, error) {
	side := spot.Side(strings.ToLower(params["side"]))
	amount, err := decimal.NewFromString(params["amount"])
	if err != nil {
		return nil, err
	}
	service := c.NewSpotClient().NewCreateOrderService(params["currencyPair"], side, amount)

	if v := params["type"]; v != "" {
		service.SetType(spot.OrderType(strings.ToLower(v)))
	}
	if v := params["price"]; v != "" {
		p, err := decimal.NewFromString(v)
		if err != nil {
			return nil, err
		}
		service.SetPrice(p)
	}
	if v := params["account"]; v != "" {
		service.SetAccount(spot.Account(strings.ToLower(v)))
	}
	if v := params["timeInForce"]; v != "" {
		service.SetTimeInForce(spot.TimeInForce(strings.ToLower(v)))
	}
	if v := params["text"]; v != "" {
		service.SetText(v)
	}
	order, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return order, nil
}

// GetOrders lists orders for a currency pair filtered by status (open/finished).
func (c *Client) GetOrders(currencyPair, status string, from, to time.Time, limit int) (*OrderList, error) {
	service := c.NewSpotClient().NewListOrdersService(currencyPair, status)
	if !from.IsZero() {
		service.SetFrom(from)
	}
	if !to.IsZero() {
		service.SetTo(to)
	}
	if limit != 0 {
		service.SetLimit(limit)
	}
	orders, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := OrderList(orders)
	return &list, nil
}

// GetOpenOrders returns open orders grouped by currency pair, flattened into a
// single order list for display.
func (c *Client) GetOpenOrders() (*OrderList, error) {
	groups, err := c.NewSpotClient().NewListAllOpenOrdersService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := OrderList{}
	for _, g := range groups {
		list = append(list, g.Orders...)
	}
	return &list, nil
}

// GetOrder queries a single order by id and currency pair.
func (c *Client) GetOrder(orderID, currencyPair string) (*spot.Order, error) {
	return c.NewSpotClient().NewGetOrderService(orderID, currencyPair).Do(context.Background())
}

// CancelOrder cancels a single order by id and currency pair.
func (c *Client) CancelOrder(orderID, currencyPair string) (*spot.Order, error) {
	return c.NewSpotClient().NewCancelOrderService(orderID, currencyPair).Do(context.Background())
}

// CancelOpenOrders cancels all open orders, optionally scoped to a currency pair.
func (c *Client) CancelOpenOrders(currencyPair string) (*OrderList, error) {
	service := c.NewSpotClient().NewCancelOrdersService()
	if currencyPair != "" {
		service.SetCurrencyPair(currencyPair)
	}
	orders, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := OrderList(orders)
	return &list, nil
}

type OrderList []spot.Order
