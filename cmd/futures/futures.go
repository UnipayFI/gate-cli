package futures

import "github.com/spf13/cobra"

// NewFuturesCmd builds the futures command subtree scoped to a --settle currency.
// includeWallet adds the account/fee/account-book commands (classic account only).
// The futures order flow carries no account field: whether it routes to the
// unified account is decided server-side by the account mode, so the unified and
// classic subtrees issue identical requests.
func NewFuturesCmd(includeWallet bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "futures",
		Short: "Futures (perpetual) trading commands",
		Long:  `Futures trading commands: order, trade, position` + walletSuffix(includeWallet) + `. Use --settle to pick the settle currency (usdt or btc).`,
	}
	cmd.PersistentFlags().String("settle", "usdt", "settle currency: usdt or btc")
	if includeWallet {
		cmd.AddCommand(InitAccountCmds()...)
	}
	cmd.AddCommand(InitOrderCmds()...)
	cmd.AddCommand(InitTradeCmds()...)
	cmd.AddCommand(InitPositionCmds()...)
	return cmd
}

func walletSuffix(includeWallet bool) string {
	if includeWallet {
		return `, account, fee, account-book`
	}
	return ""
}
