package spot

import "github.com/spf13/cobra"

// NewSpotCmd builds the spot command subtree. includeWallet adds the
// account/fee/account-book commands (classic account only). defaultAccount, when
// set (e.g. "unified"), pins every order/trade operation to that account and
// hides the --account flag.
func NewSpotCmd(includeWallet bool, defaultAccount string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "spot",
		Short: "Spot trading commands",
		Long:  `Spot trading commands: order, trade` + walletSuffix(includeWallet) + `.`,
	}
	if includeWallet {
		cmd.AddCommand(InitAccountCmds()...)
	}
	cmd.AddCommand(InitOrderCmds(defaultAccount)...)
	cmd.AddCommand(InitTradeCmds(defaultAccount)...)
	return cmd
}

func walletSuffix(includeWallet bool) string {
	if includeWallet {
		return `, account, fee, account-book`
	}
	return ""
}
