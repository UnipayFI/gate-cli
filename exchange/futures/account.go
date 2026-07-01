package futures

import (
	"context"
	"time"

	"github.com/UnipayFI/go-gate/v4/futures"
)

// GetAccount returns the futures account summary for the settle currency.
func (c *Client) GetAccount() (*Account, error) {
	account, err := c.futuresClient().NewListFuturesAccountsService(c.Settle).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &Account{FuturesAccount: *account}, nil
}

// GetAccountBook returns the futures account balance change history.
func (c *Client) GetAccountBook(contract, changeType string, from, to time.Time, limit int) (*AccountBookList, error) {
	service := c.futuresClient().NewListFuturesAccountBookService(c.Settle)
	if contract != "" {
		service.SetContract(contract)
	}
	if changeType != "" {
		service.SetType(changeType)
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
	records, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := AccountBookList(records)
	return &list, nil
}

// GetFee returns the taker/maker fee rates keyed by contract.
func (c *Client) GetFee(contract string) (*FeeList, error) {
	service := c.futuresClient().NewGetFuturesFeeService(c.Settle)
	if contract != "" {
		service.SetContract(contract)
	}
	fees, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := FeeList{}
	for contractName, fee := range fees {
		list = append(list, FeeEntry{Contract: contractName, FuturesFee: fee})
	}
	return &list, nil
}

type AccountBookList []futures.FuturesAccountBook

// Account wraps futures.FuturesAccount for table rendering.
type Account struct {
	futures.FuturesAccount
}

// FeeEntry pairs a contract name with its fee rates (the SDK returns a map).
type FeeEntry struct {
	Contract string
	futures.FuturesFee
}

type FeeList []FeeEntry
