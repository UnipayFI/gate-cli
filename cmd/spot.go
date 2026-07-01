package cmd

import (
	"github.com/UnipayFI/gate-cli/cmd/spot"
	"github.com/spf13/cobra"
)

var SpotCmd = &cobra.Command{
	Use:   "spot",
	Short: "Spot trading commands",
	Long:  `Spot trading commands: account, fee, account-book, order, trade.`,
}

func init() {
	SpotCmd.AddCommand(spot.InitAccountCmds()...)
	SpotCmd.AddCommand(spot.InitOrderCmds()...)
	SpotCmd.AddCommand(spot.InitTradeCmds()...)
	RootCmd.AddCommand(SpotCmd)
}
