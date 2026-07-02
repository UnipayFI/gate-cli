package unified

import (
	"context"
	"time"

	"github.com/UnipayFI/go-gate/v4/unified"
	"github.com/shopspring/decimal"
)

// BorrowOrRepay borrows or repays a currency on the unified account. loanType is
// "borrow" or "repay"; for repayment, repaidAll clears the full outstanding
// balance and overrides amount.
func (c *Client) BorrowOrRepay(currency string, amount decimal.Decimal, loanType string, repaidAll bool) (*LoanResult, error) {
	service := c.NewUnifiedClient().NewCreateUnifiedLoanService(currency, amount, loanType)
	if repaidAll {
		service.SetRepaidAll(true)
	}
	result, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &LoanResult{UnifiedLoanResult: *result}, nil
}

// GetLoans returns the account's outstanding borrowings. loanType filters by
// "platform" or "margin".
func (c *Client) GetLoans(currency, loanType string, page, limit int) (*LoanList, error) {
	service := c.NewUnifiedClient().NewListUnifiedLoansService()
	if currency != "" {
		service.SetCurrency(currency)
	}
	if loanType != "" {
		service.SetType(loanType)
	}
	if page != 0 {
		service.SetPage(page)
	}
	if limit != 0 {
		service.SetLimit(limit)
	}
	loans, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := LoanList(loans)
	return &list, nil
}

// GetLoanRecords returns the borrow/repay history. recordType filters by
// "borrow" or "repay".
func (c *Client) GetLoanRecords(currency, recordType string, page, limit int) (*LoanRecordList, error) {
	service := c.NewUnifiedClient().NewListUnifiedLoanRecordsService()
	if currency != "" {
		service.SetCurrency(currency)
	}
	if recordType != "" {
		service.SetType(recordType)
	}
	if page != 0 {
		service.SetPage(page)
	}
	if limit != 0 {
		service.SetLimit(limit)
	}
	records, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := LoanRecordList(records)
	return &list, nil
}

// GetInterestRecords returns the interest-deduction history. loanType filters by
// "platform" or "margin".
func (c *Client) GetInterestRecords(currency, loanType string, from, to time.Time, page, limit int) (*InterestRecordList, error) {
	service := c.NewUnifiedClient().NewListUnifiedLoanInterestRecordsService()
	if currency != "" {
		service.SetCurrency(currency)
	}
	if loanType != "" {
		service.SetType(loanType)
	}
	if !from.IsZero() {
		service.SetFrom(from)
	}
	if !to.IsZero() {
		service.SetTo(to)
	}
	if page != 0 {
		service.SetPage(page)
	}
	if limit != 0 {
		service.SetLimit(limit)
	}
	records, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := InterestRecordList(records)
	return &list, nil
}

type (
	// LoanResult wraps unified.UnifiedLoanResult for table rendering.
	LoanResult         struct{ unified.UnifiedLoanResult }
	LoanList           []unified.UnifiedLoan
	LoanRecordList     []unified.UnifiedLoanRecord
	InterestRecordList []unified.UnifiedLoanInterestRecord
)
