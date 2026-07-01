package futures

import (
	"context"
	"strconv"
	"strings"

	"github.com/UnipayFI/go-gate/v4/futures"
	"github.com/shopspring/decimal"
)

// CreateOrder places a futures order. size is signed: positive = long/buy,
// negative = short/sell. params is the long-name keyed map built by the cmd layer.
func (c *Client) CreateOrder(params map[string]string) (*futures.FuturesOrder, error) {
	size, err := strconv.ParseInt(params["size"], 10, 64)
	if err != nil {
		return nil, err
	}
	service := c.futuresClient().NewCreateFuturesOrderService(c.Settle, params["contract"], size)

	if v := params["price"]; v != "" {
		p, err := decimal.NewFromString(v)
		if err != nil {
			return nil, err
		}
		service.SetPrice(p)
	}
	if v := params["timeInForce"]; v != "" {
		service.SetTimeInForce(futures.TimeInForce(strings.ToLower(v)))
	}
	if v := params["text"]; v != "" {
		service.SetText(v)
	}
	if params["reduceOnly"] == "true" {
		service.SetReduceOnly(true)
	}
	if params["close"] == "true" {
		service.SetClose(true)
	}
	return service.Do(context.Background())
}

// GetOrders lists orders filtered by status (open/finished) and optional contract.
func (c *Client) GetOrders(status, contract string, limit int) (*OrderList, error) {
	service := c.futuresClient().NewListFuturesOrdersService(c.Settle, futures.OrderStatus(status))
	if contract != "" {
		service.SetContract(contract)
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

// GetOrder queries a single order by id.
func (c *Client) GetOrder(orderID string) (*futures.FuturesOrder, error) {
	return c.futuresClient().NewGetFuturesOrderService(c.Settle, orderID).Do(context.Background())
}

// CancelOrder cancels a single order by id.
func (c *Client) CancelOrder(orderID string) (*futures.FuturesOrder, error) {
	return c.futuresClient().NewCancelFuturesOrderService(c.Settle, orderID).Do(context.Background())
}

// CancelOpenOrders cancels all open orders for a contract.
func (c *Client) CancelOpenOrders(contract string) (*OrderList, error) {
	orders, err := c.futuresClient().NewCancelFuturesOrdersService(c.Settle, contract).Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := OrderList(orders)
	return &list, nil
}

type OrderList []futures.FuturesOrder
