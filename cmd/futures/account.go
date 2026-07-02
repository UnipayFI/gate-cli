package futures

import (
	"log"
	"strings"
	"time"

	"github.com/UnipayFI/gate-cli/common"
	"github.com/UnipayFI/gate-cli/config"
	"github.com/UnipayFI/gate-cli/exchange"
	"github.com/UnipayFI/gate-cli/exchange/futures"
	"github.com/UnipayFI/gate-cli/printer"
	gatefutures "github.com/UnipayFI/go-gate/v4/futures"
	"github.com/spf13/cobra"
)

var (
	accountCmd = &cobra.Command{
		Use:   "account",
		Short: "Show futures account summary",
		Long: `Show the futures account summary (equity, margin, unrealised PnL) for the settle currency.

Docs Link: https://www.gate.com/docs/developers/apiv4/en/futures/#get-futures-account`,
		Run: showAccount,
	}

	feeCmd = &cobra.Command{
		Use:   "fee",
		Short: "Show trading fee rates",
		Long: `Query user trading fee rates, keyed by contract. Use --contract to filter.

Docs Link: https://www.gate.com/docs/developers/apiv4/en/futures/#query-futures-market-trading-fee-rates`,
		Run: showFee,
	}

	accountBookCmd = &cobra.Command{
		Use:   "account-book",
		Short: "Query account balance change history",
		Long: `Query futures account balance change history (pnl, fees, funding, etc.).

Docs Link: https://www.gate.com/docs/developers/apiv4/en/futures/#query-futures-account-change-history`,
		Run: showAccountBook,
	}
)

func InitAccountCmds() []*cobra.Command {
	feeCmd.Flags().StringP("contract", "c", "", "contract, e.g. BTC_USDT")

	accountBookCmd.Flags().StringP("contract", "c", "", "filter a single contract")
	accountBookCmd.Flags().StringP("type", "t", "", "change type: dnw, pnl, fee, refr, fund, etc.")
	accountBookCmd.Flags().IntP("limit", "l", 100, "max records (default 100)")
	accountBookCmd.Flags().StringP("from", "a", "", "start time (unix seconds or \"YYYY-MM-DD HH:MM:SS\")")
	accountBookCmd.Flags().StringP("to", "e", "", "end time (unix seconds or \"YYYY-MM-DD HH:MM:SS\")")

	return []*cobra.Command{accountCmd, feeCmd, accountBookCmd}
}

// newClient builds a futures exchange client scoped to the --settle currency.
func newClient(cmd *cobra.Command) futures.Client {
	settleRaw, _ := cmd.Flags().GetString("settle")
	settle := gatefutures.Settle(strings.ToLower(settleRaw))
	return futures.Client{
		Client: exchange.NewClient(config.Config.APIKey, config.Config.APISecret),
		Settle: settle,
	}
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

func showAccount(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	account, err := client.GetAccount()
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(account)
}

func showFee(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	contract, _ := cmd.Flags().GetString("contract")
	fees, err := client.GetFee(contract)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(fees)
}

func showAccountBook(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	contract, _ := cmd.Flags().GetString("contract")
	changeType, _ := cmd.Flags().GetString("type")
	limit, _ := cmd.Flags().GetInt("limit")
	from := parseTimeOrFatal(cmd, "--from", "from")
	to := parseTimeOrFatal(cmd, "--to", "to")

	records, err := client.GetAccountBook(contract, changeType, from, to, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(records)
}
