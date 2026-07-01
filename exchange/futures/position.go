package futures

import (
	"context"
	"time"

	"github.com/UnipayFI/go-gate/v4/futures"
)

// GetPositions lists positions, optionally only those currently holding size.
func (c *Client) GetPositions(holdingOnly bool) (*PositionList, error) {
	service := c.futuresClient().NewListPositionsService(c.Settle)
	if holdingOnly {
		service.SetHolding(true)
	}
	positions, err := service.Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := PositionList(positions)
	return &list, nil
}

// GetPosition returns a single position for a contract.
func (c *Client) GetPosition(contract string) (*PositionList, error) {
	position, err := c.futuresClient().NewGetPositionService(c.Settle, contract).Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := PositionList{*position}
	return &list, nil
}

// UpdateMargin changes the isolated margin of a position. change is signed:
// positive adds margin, negative reduces it.
func (c *Client) UpdateMargin(contract, change string) (*PositionList, error) {
	position, err := c.futuresClient().NewUpdatePositionMarginService(c.Settle, contract, change).Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := PositionList{*position}
	return &list, nil
}

// UpdateLeverage changes the leverage of a position. "0" means cross margin.
func (c *Client) UpdateLeverage(contract, leverage string) (*PositionList, error) {
	position, err := c.futuresClient().NewUpdatePositionLeverageService(c.Settle, contract, leverage).Do(context.Background())
	if err != nil {
		return nil, err
	}
	list := PositionList{*position}
	return &list, nil
}

// SetDualMode enables or disables hedge (dual) position mode.
func (c *Client) SetDualMode(enable bool) (*futures.DualModeResult, error) {
	return c.futuresClient().NewSetDualModeService(c.Settle, enable).Do(context.Background())
}

// GetPositionClose returns the position close history.
func (c *Client) GetPositionClose(contract string, from, to time.Time, limit int) (*PositionCloseList, error) {
	service := c.futuresClient().NewListPositionCloseService(c.Settle)
	if contract != "" {
		service.SetContract(contract)
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
	list := PositionCloseList(records)
	return &list, nil
}

type (
	PositionList      []futures.Position
	PositionCloseList []futures.PositionClose
)
