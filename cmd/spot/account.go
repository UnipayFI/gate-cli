package spot

import (
	"log"
	"time"

	"github.com/UnipayFI/gate-cli/common"
	"github.com/UnipayFI/gate-cli/config"
	"github.com/UnipayFI/gate-cli/exchange"
	"github.com/UnipayFI/gate-cli/exchange/spot"
	"github.com/UnipayFI/gate-cli/printer"
	"github.com/spf13/cobra"
)

var (
	accountCmd = &cobra.Command{
		Use:   "account",
		Short: "Show spot account balances",
		Long: `List spot account balances. Use --currency to filter a single currency.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#list-spot-accounts`,
		Run: showAccount,
	}

	feeCmd = &cobra.Command{
		Use:   "fee",
		Short: "Show trading fee rates",
		Long: `Query user trading fee rates. Use --currency-pair for a single pair.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#query-user-trading-fee-rates`,
		Run: showFee,
	}

	accountBookCmd = &cobra.Command{
		Use:   "account-book",
		Short: "Query account balance change history",
		Long: `Query spot account balance change history (deposits, trades, fees, etc.).

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#query-account-book`,
		Run: showAccountBook,
	}
)

func InitAccountCmds() []*cobra.Command {
	accountCmd.Flags().StringP("currency", "c", "", "filter a single currency, e.g. USDT")

	feeCmd.Flags().StringP("currency-pair", "p", "", "currency pair, e.g. BTC_USDT")

	accountBookCmd.Flags().StringP("currency", "c", "", "filter a single currency")
	accountBookCmd.Flags().StringP("type", "t", "", "change type, e.g. deposit, withdraw, trade, fee")
	accountBookCmd.Flags().IntP("limit", "l", 100, "max records (default 100)")
	accountBookCmd.Flags().StringP("from", "a", "", "start time (unix seconds or \"YYYY-MM-DD HH:MM:SS\")")
	accountBookCmd.Flags().StringP("to", "e", "", "end time (unix seconds or \"YYYY-MM-DD HH:MM:SS\")")

	return []*cobra.Command{accountCmd, feeCmd, accountBookCmd}
}

func newClient() spot.Client {
	return spot.Client{Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret)}
}

func showAccount(cmd *cobra.Command, _ []string) {
	client := newClient()
	currency, _ := cmd.Flags().GetString("currency")
	accounts, err := client.GetAccounts(currency)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(accounts)
}

func showFee(cmd *cobra.Command, _ []string) {
	client := newClient()
	pair, _ := cmd.Flags().GetString("currency-pair")
	fee, err := client.GetFee(pair)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(fee)
}

func showAccountBook(cmd *cobra.Command, _ []string) {
	client := newClient()
	currency, _ := cmd.Flags().GetString("currency")
	bookType, _ := cmd.Flags().GetString("type")
	limit, _ := cmd.Flags().GetInt("limit")
	from := parseTimeOrFatal(cmd, "--from", "from")
	to := parseTimeOrFatal(cmd, "--to", "to")

	records, err := client.GetAccountBook(currency, from, to, limit, bookType)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(records)
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
