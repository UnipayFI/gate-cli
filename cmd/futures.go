package cmd

import (
	"github.com/UnipayFI/gate-cli/cmd/futures"
	"github.com/spf13/cobra"
)

var FuturesCmd = &cobra.Command{
	Use:   "futures",
	Short: "Futures (perpetual) trading commands",
	Long:  `Futures trading commands: account, fee, order, trade, position. Use --settle to pick the settle currency (usdt or btc).`,
}

func init() {
	FuturesCmd.PersistentFlags().String("settle", "usdt", "settle currency: usdt or btc")
	FuturesCmd.AddCommand(futures.InitAccountCmds()...)
	FuturesCmd.AddCommand(futures.InitOrderCmds()...)
	FuturesCmd.AddCommand(futures.InitTradeCmds()...)
	FuturesCmd.AddCommand(futures.InitPositionCmds()...)
	RootCmd.AddCommand(FuturesCmd)
}
