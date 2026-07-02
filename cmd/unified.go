package cmd

import (
	"github.com/UnipayFI/gate-cli/cmd/futures"
	"github.com/UnipayFI/gate-cli/cmd/spot"
	"github.com/UnipayFI/gate-cli/cmd/unified"
	"github.com/spf13/cobra"
)

var UnifiedCmd = &cobra.Command{
	Use:   "unified",
	Short: "Unified account commands",
	Long: `Unified account: aggregated equity/margin across products, plus borrowing.

Account-level: account, mode, borrowable, transferable, borrow, repay, loans,
loan-records, interest-records.
Trading (routed to the unified account): spot, futures.`,
}

func init() {
	UnifiedCmd.AddCommand(unified.InitAccountCmds()...)
	UnifiedCmd.AddCommand(unified.InitLoanCmds()...)
	UnifiedCmd.AddCommand(spot.NewSpotCmd(false, "unified"))
	UnifiedCmd.AddCommand(futures.NewFuturesCmd(false))
	RootCmd.AddCommand(UnifiedCmd)
}
