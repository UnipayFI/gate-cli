package unified

import (
	"context"

	"github.com/UnipayFI/go-gate/v4/unified"
)

// GetAccount returns the unified-account overview, optionally filtered to a
// single currency's balance.
func (c *Client) GetAccount(currency string) (*Account, error) {
	service := c.NewUnifiedClient().NewListUnifiedAccountsService()
	if currency != "" {
		service.SetCurrency(currency)
	}
	account, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &Account{UnifiedAccount: *account}, nil
}

// GetMode returns the account's unified margin mode and its toggles.
func (c *Client) GetMode() (*Mode, error) {
	mode, err := c.NewUnifiedClient().NewGetUnifiedModeService().Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &Mode{UnifiedMode: *mode}, nil
}

// GetBorrowable returns the maximum amount of a currency the account can borrow.
func (c *Client) GetBorrowable(currency string) (*Borrowable, error) {
	b, err := c.NewUnifiedClient().NewGetUnifiedBorrowableService(currency).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &Borrowable{UnifiedBorrowable: *b}, nil
}

// GetTransferable returns the maximum amount of a currency that can be
// transferred out of the unified account.
func (c *Client) GetTransferable(currency string) (*Transferable, error) {
	t, err := c.NewUnifiedClient().NewGetUnifiedTransferableService(currency).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &Transferable{UnifiedTransferable: *t}, nil
}

type (
	// Account wraps unified.UnifiedAccount so it can implement printer.TableWriter.
	Account struct{ unified.UnifiedAccount }
	// Mode wraps unified.UnifiedMode for table rendering.
	Mode struct{ unified.UnifiedMode }
	// Borrowable wraps unified.UnifiedBorrowable for table rendering.
	Borrowable struct{ unified.UnifiedBorrowable }
	// Transferable wraps unified.UnifiedTransferable for table rendering.
	Transferable struct{ unified.UnifiedTransferable }
)
