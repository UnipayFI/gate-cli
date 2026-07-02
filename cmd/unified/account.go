package unified

import (
	"log"
	"time"

	"github.com/UnipayFI/gate-cli/common"
	"github.com/UnipayFI/gate-cli/config"
	"github.com/UnipayFI/gate-cli/exchange"
	"github.com/UnipayFI/gate-cli/exchange/unified"
	"github.com/UnipayFI/gate-cli/printer"
	"github.com/spf13/cobra"
)

func InitAccountCmds() []*cobra.Command {
	accountCmd := &cobra.Command{
		Use:   "account",
		Short: "Show unified account overview",
		Long: `Show the unified account overview: per-currency balances plus account-wide
equity, margin and liabilities. Use --currency to filter a single currency, and
--json to see the account-wide totals (equity, margin rate, borrowed).

Docs Link: https://www.gate.com/docs/developers/apiv4/en/unified/#get-unified-account-information`,
		Run: showAccount,
	}

	modeCmd := &cobra.Command{
		Use:   "mode",
		Short: "Show the unified account mode",
		Long: `Show the account's unified margin mode (classic, single_currency, multi_currency
or portfolio) and its feature toggles.

Docs Link: https://www.gate.com/docs/developers/apiv4/en/unified/#query-mode-of-the-unified-account`,
		Run: showMode,
	}

	borrowableCmd := &cobra.Command{
		Use:   "borrowable",
		Short: "Show the maximum borrowable amount of a currency",
		Long: `Show the maximum amount of a currency the unified account can still borrow.

Docs Link: https://www.gate.com/docs/developers/apiv4/en/unified/#query-maximum-borrowable-amount-for-unified-account`,
		Run: showBorrowable,
	}

	transferableCmd := &cobra.Command{
		Use:   "transferable",
		Short: "Show the maximum transferable amount of a currency",
		Long: `Show the maximum amount of a currency that can be transferred out of the unified account.

Docs Link: https://www.gate.com/docs/developers/apiv4/en/unified/#query-maximum-transferable-amount-for-unified-account`,
		Run: showTransferable,
	}

	accountCmd.Flags().StringP("currency", "c", "", "filter a single currency, e.g. USDT")

	borrowableCmd.Flags().StringP("currency", "c", "", "currency, e.g. BTC")
	borrowableCmd.MarkFlagRequired("currency")

	transferableCmd.Flags().StringP("currency", "c", "", "currency, e.g. BTC")
	transferableCmd.MarkFlagRequired("currency")

	return []*cobra.Command{accountCmd, modeCmd, borrowableCmd, transferableCmd}
}

func newClient() unified.Client {
	return unified.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
}

func showAccount(cmd *cobra.Command, _ []string) {
	client := newClient()
	currency, _ := cmd.Flags().GetString("currency")
	account, err := client.GetAccount(currency)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(account)
}

func showMode(_ *cobra.Command, _ []string) {
	client := newClient()
	mode, err := client.GetMode()
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(mode)
}

func showBorrowable(cmd *cobra.Command, _ []string) {
	client := newClient()
	currency, _ := cmd.Flags().GetString("currency")
	b, err := client.GetBorrowable(currency)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(b)
}

func showTransferable(cmd *cobra.Command, _ []string) {
	client := newClient()
	currency, _ := cmd.Flags().GetString("currency")
	t, err := client.GetTransferable(currency)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(t)
}

// parseTimeOrFatal reads a datetime flag and converts it to time.Time,
// exiting on error. A missing/empty flag yields the zero time.
func parseTimeOrFatal(cmd *cobra.Command, flagLabel, flagName string) time.Time {
	raw, _ := cmd.Flags().GetString(flagName)
	sec, ok, err := common.ParseTimeFlag(flagLabel, raw)
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		return time.Time{}
	}
	return time.Unix(sec, 0)
}
