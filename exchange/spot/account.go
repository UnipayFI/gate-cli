package spot

import (
	"context"
	"time"

	"github.com/UnipayFI/go-gate/v4/spot"
)

// GetAccounts returns spot balances, optionally filtered to a single currency.
func (c *Client) GetAccounts(currency string) (*AccountList, error) {
	service := c.NewSpotClient().NewListSpotAccountsService()
	if currency != "" {
		service.SetCurrency(currency)
	}
	accounts, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := AccountList(accounts)
	return &list, nil
}

// GetFee returns the trading fee rates, optionally for a single currency pair.
func (c *Client) GetFee(currencyPair string) (*Fee, error) {
	service := c.NewSpotClient().NewGetFeeService()
	if currencyPair != "" {
		service.SetCurrencyPair(currencyPair)
	}
	fee, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &Fee{SpotFee: *fee}, nil
}

// GetAccountBook returns the account balance change history.
func (c *Client) GetAccountBook(currency string, from, to time.Time, limit int, bookType string) (*AccountBookList, error) {
	service := c.NewSpotClient().NewListSpotAccountBookService()
	if currency != "" {
		service.SetCurrency(currency)
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
	if bookType != "" {
		service.SetType(bookType)
	}
	records, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := AccountBookList(records)
	return &list, nil
}

type (
	AccountList     []spot.SpotAccount
	AccountBookList []spot.SpotAccountBook
)

// Fee wraps spot.SpotFee so it can implement printer.TableWriter.
type Fee struct {
	spot.SpotFee
}
