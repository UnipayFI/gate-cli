package cmd

import (
	"github.com/UnipayFI/gate-cli/cmd/futures"
	"github.com/UnipayFI/gate-cli/cmd/spot"
	"github.com/spf13/cobra"
)

var ClassicCmd = &cobra.Command{
	Use:   "classic",
	Short: "Classic account commands",
	Long:  `Classic account: independent spot and futures accounts. Subcommands: spot, futures.`,
}

func init() {
	ClassicCmd.AddCommand(spot.NewSpotCmd(true, ""))
	ClassicCmd.AddCommand(futures.NewFuturesCmd(true))
	RootCmd.AddCommand(ClassicCmd)
}
