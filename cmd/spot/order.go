package spot

import (
	"log"
	"time"

	"github.com/UnipayFI/gate-cli/exchange/spot"
	"github.com/UnipayFI/gate-cli/printer"
	"github.com/spf13/cobra"
)

var (
	orderCmd = &cobra.Command{
		Use:   "order",
		Short: "Create, cancel, list and query spot orders",
	}

	orderCreateCmd = &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create an order",
		Long: `Create a new spot order.
* Required: --currency-pair, --side, --amount
* For LIMIT orders also pass --price. Default type is limit.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#create-an-order`,
		Run: createOrder,
	}

	orderListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List orders for a currency pair",
		Long: `List orders for a currency pair filtered by status (open or finished).

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#list-orders`,
		Run: listOrders,
	}

	orderOpenCmd = &cobra.Command{
		Use:   "open",
		Short: "List all open orders across pairs",
		Long: `List all open orders for the account, across every currency pair.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#list-all-open-orders`,
		Run: listOpenOrders,
	}

	orderGetCmd = &cobra.Command{
		Use:   "get",
		Short: "Query a single order",
		Long: `Query a single order by its id and currency pair.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#get-a-single-order`,
		Run: getOrder,
	}

	orderCancelCmd = &cobra.Command{
		Use:   "cancel",
		Short: "Cancel an order, or all open orders",
		Long: `Cancel an order by --order-id, or cancel all open orders when no id is given.
If --currency-pair is set without --order-id, only that pair's open orders are canceled.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#cancel-a-single-order`,
		Run: cancelOrder,
	}
)

func InitOrderCmds() []*cobra.Command {
	orderCmd.PersistentFlags().StringP("currency-pair", "p", "", "currency pair, e.g. BTC_USDT")

	orderCreateCmd.Flags().StringP("side", "S", "", "buy or sell")
	orderCreateCmd.Flags().StringP("amount", "q", "", "order amount (decimal string)")
	orderCreateCmd.Flags().StringP("type", "t", "limit", "limit or market")
	orderCreateCmd.Flags().String("price", "", "order price, required for limit orders (decimal string)")
	orderCreateCmd.Flags().StringP("time-in-force", "T", "", "gtc, ioc, poc, fok")
	orderCreateCmd.Flags().StringP("account", "A", "", "spot, margin, cross_margin, unified")
	orderCreateCmd.Flags().String("text", "", "custom order label (must start with t-)")
	orderCreateCmd.MarkFlagRequired("currency-pair")
	orderCreateCmd.MarkFlagRequired("side")
	orderCreateCmd.MarkFlagRequired("amount")

	orderListCmd.Flags().StringP("status", "s", "open", "order status: open or finished")
	orderListCmd.Flags().IntP("limit", "l", 100, "max records (default 100)")
	orderListCmd.MarkFlagRequired("currency-pair")

	orderGetCmd.Flags().StringP("order-id", "i", "", "order id")
	orderGetCmd.MarkFlagRequired("currency-pair")
	orderGetCmd.MarkFlagRequired("order-id")

	orderCancelCmd.Flags().StringP("order-id", "i", "", "order id (omit to cancel all open orders)")

	orderCmd.AddCommand(orderCreateCmd, orderListCmd, orderOpenCmd, orderGetCmd, orderCancelCmd)
	return []*cobra.Command{orderCmd}
}

func createOrder(cmd *cobra.Command, _ []string) {
	client := newClient()
	order, err := client.CreateOrder(buildOrderParams(cmd))
	if err != nil {
		log.Fatal(err)
	}
	orders := spot.OrderList{*order}
	printer.Print(&orders)
}

// buildOrderParams reads the cobra-parsed flags and returns the long-name keyed
// map that exchange/spot.Client.CreateOrder expects.
func buildOrderParams(cmd *cobra.Command) map[string]string {
	params := map[string]string{}
	flags := map[string]string{
		"currency-pair": "currencyPair",
		"side":          "side",
		"amount":        "amount",
		"type":          "type",
		"price":         "price",
		"time-in-force": "timeInForce",
		"account":       "account",
		"text":          "text",
	}
	for flagName, key := range flags {
		if v, _ := cmd.Flags().GetString(flagName); v != "" {
			params[key] = v
		}
	}
	return params
}

func listOrders(cmd *cobra.Command, _ []string) {
	client := newClient()
	pair, _ := cmd.Flags().GetString("currency-pair")
	status, _ := cmd.Flags().GetString("status")
	limit, _ := cmd.Flags().GetInt("limit")

	orders, err := client.GetOrders(pair, status, time.Time{}, time.Time{}, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(orders)
}

func listOpenOrders(cmd *cobra.Command, _ []string) {
	client := newClient()
	orders, err := client.GetOpenOrders()
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(orders)
}

func getOrder(cmd *cobra.Command, _ []string) {
	client := newClient()
	pair, _ := cmd.Flags().GetString("currency-pair")
	orderID, _ := cmd.Flags().GetString("order-id")

	order, err := client.GetOrder(orderID, pair)
	if err != nil {
		log.Fatal(err)
	}
	orders := spot.OrderList{*order}
	printer.Print(&orders)
}

func cancelOrder(cmd *cobra.Command, _ []string) {
	client := newClient()
	pair, _ := cmd.Flags().GetString("currency-pair")
	orderID, _ := cmd.Flags().GetString("order-id")

	if orderID == "" {
		orders, err := client.CancelOpenOrders(pair)
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(orders)
		return
	}
	order, err := client.CancelOrder(orderID, pair)
	if err != nil {
		log.Fatal(err)
	}
	orders := spot.OrderList{*order}
	printer.Print(&orders)
}
