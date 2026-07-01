package spot

import (
	"log"

	"github.com/UnipayFI/gate-cli/printer"
	"github.com/spf13/cobra"
)

var tradeCmd = &cobra.Command{
	Use:   "trade",
	Short: "Query personal trading history",
	Long: `List personal spot trades. Use --currency-pair to filter a pair and
--order-id (with --currency-pair) to filter a single order.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#list-personal-trading-history`,
	Run: listTrades,
}

func InitTradeCmds() []*cobra.Command {
	tradeCmd.Flags().StringP("currency-pair", "p", "", "currency pair, e.g. BTC_USDT")
	tradeCmd.Flags().StringP("order-id", "i", "", "filter by order id (requires --currency-pair)")
	tradeCmd.Flags().IntP("limit", "l", 100, "max records (default 100)")
	tradeCmd.Flags().StringP("from", "a", "", "start time (unix seconds or \"YYYY-MM-DD HH:MM:SS\")")
	tradeCmd.Flags().StringP("to", "e", "", "end time (unix seconds or \"YYYY-MM-DD HH:MM:SS\")")
	return []*cobra.Command{tradeCmd}
}

func listTrades(cmd *cobra.Command, _ []string) {
	client := newClient()
	pair, _ := cmd.Flags().GetString("currency-pair")
	orderID, _ := cmd.Flags().GetString("order-id")
	limit, _ := cmd.Flags().GetInt("limit")
	from := parseTimeOrFatal(cmd, "--from", "from")
	to := parseTimeOrFatal(cmd, "--to", "to")

	trades, err := client.GetMyTrades(pair, orderID, from, to, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(trades)
}
