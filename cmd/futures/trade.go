package futures

import (
	"log"

	"github.com/UnipayFI/gate-cli/printer"
	"github.com/spf13/cobra"
)

func InitTradeCmds() []*cobra.Command {
	tradeCmd := &cobra.Command{
		Use:   "trade",
		Short: "Query personal trading history",
		Long: `List personal futures trades for the settle currency. Use --contract to filter
a contract and --order-id to filter a single order.

Docs Link: https://www.gate.com/docs/developers/apiv4/en/futures/#query-personal-trading-records`,
		Run: listTrades,
	}

	tradeCmd.Flags().StringP("contract", "c", "", "contract, e.g. BTC_USDT")
	tradeCmd.Flags().Int64P("order-id", "i", 0, "filter by order id")
	tradeCmd.Flags().IntP("limit", "l", 100, "max records (default 100)")
	return []*cobra.Command{tradeCmd}
}

func listTrades(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	contract, _ := cmd.Flags().GetString("contract")
	orderID, _ := cmd.Flags().GetInt64("order-id")
	limit, _ := cmd.Flags().GetInt("limit")

	trades, err := client.GetMyTrades(contract, orderID, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(trades)
}
