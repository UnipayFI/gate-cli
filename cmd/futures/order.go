package futures

import (
	"log"
	"strconv"

	"github.com/UnipayFI/gate-cli/exchange/futures"
	"github.com/UnipayFI/gate-cli/printer"
	"github.com/spf13/cobra"
)

func InitOrderCmds() []*cobra.Command {
	orderCmd := &cobra.Command{
		Use:   "order",
		Short: "Create, cancel, list and query futures orders",
	}

	orderCreateCmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create an order",
		Long: `Create a new futures order.
* Required: --contract, --size (signed: positive = long, negative = short).
* Pass --price for a limit order; omit price with --tif ioc for a market order.

Docs Link: https://www.gate.com/docs/developers/apiv4/en/futures/#place-futures-order`,
		Run: createOrder,
	}

	orderListCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List orders by status",
		Long: `List futures orders filtered by status (open or finished) and optional contract.

Docs Link: https://www.gate.com/docs/developers/apiv4/en/futures/#query-futures-order-list`,
		Run: listOrders,
	}

	orderGetCmd := &cobra.Command{
		Use:   "get",
		Short: "Query a single order",
		Long: `Query a single futures order by its id.

Docs Link: https://www.gate.com/docs/developers/apiv4/en/futures/#query-single-order-details`,
		Run: getOrder,
	}

	orderCancelCmd := &cobra.Command{
		Use:   "cancel",
		Short: "Cancel an order, or all open orders of a contract",
		Long: `Cancel an order by --order-id, or cancel all open orders of --contract when no id is given.

Docs Link: https://www.gate.com/docs/developers/apiv4/en/futures/#cancel-single-order`,
		Run: cancelOrder,
	}

	orderCreateCmd.Flags().StringP("contract", "c", "", "contract, e.g. BTC_USDT")
	orderCreateCmd.Flags().Int64P("size", "z", 0, "order size in contracts (signed: + long, - short)")
	orderCreateCmd.Flags().String("price", "", "order price (decimal string); omit for market")
	orderCreateCmd.Flags().StringP("tif", "T", "", "time in force: gtc, ioc, poc, fok")
	orderCreateCmd.Flags().Bool("reduce-only", false, "reduce-only order")
	orderCreateCmd.Flags().Bool("close", false, "close position (size must be 0)")
	orderCreateCmd.Flags().String("text", "", "custom order label (must start with t-)")
	orderCreateCmd.MarkFlagRequired("contract")

	orderListCmd.Flags().StringP("status", "s", "open", "order status: open or finished")
	orderListCmd.Flags().StringP("contract", "c", "", "filter a single contract")
	orderListCmd.Flags().IntP("limit", "l", 100, "max records (default 100)")

	orderGetCmd.Flags().StringP("order-id", "i", "", "order id")
	orderGetCmd.MarkFlagRequired("order-id")

	orderCancelCmd.Flags().StringP("order-id", "i", "", "order id (omit to cancel all open orders of a contract)")
	orderCancelCmd.Flags().StringP("contract", "c", "", "contract (required when canceling all)")

	orderCmd.AddCommand(orderCreateCmd, orderListCmd, orderGetCmd, orderCancelCmd)
	return []*cobra.Command{orderCmd}
}

func createOrder(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	order, err := client.CreateOrder(buildOrderParams(cmd))
	if err != nil {
		log.Fatal(err)
	}
	orders := futures.OrderList{*order}
	printer.Print(&orders)
}

// buildOrderParams reads the cobra-parsed flags and returns the long-name keyed
// map that exchange/futures.Client.CreateOrder expects.
func buildOrderParams(cmd *cobra.Command) map[string]string {
	params := map[string]string{}
	contract, _ := cmd.Flags().GetString("contract")
	size, _ := cmd.Flags().GetInt64("size")
	params["contract"] = contract
	params["size"] = strconv.FormatInt(size, 10)
	if v, _ := cmd.Flags().GetString("price"); v != "" {
		params["price"] = v
	}
	if v, _ := cmd.Flags().GetString("tif"); v != "" {
		params["timeInForce"] = v
	}
	if v, _ := cmd.Flags().GetString("text"); v != "" {
		params["text"] = v
	}
	if v, _ := cmd.Flags().GetBool("reduce-only"); v {
		params["reduceOnly"] = "true"
	}
	if v, _ := cmd.Flags().GetBool("close"); v {
		params["close"] = "true"
	}
	return params
}

func listOrders(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	status, _ := cmd.Flags().GetString("status")
	contract, _ := cmd.Flags().GetString("contract")
	limit, _ := cmd.Flags().GetInt("limit")

	orders, err := client.GetOrders(status, contract, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(orders)
}

func getOrder(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	orderID, _ := cmd.Flags().GetString("order-id")

	order, err := client.GetOrder(orderID)
	if err != nil {
		log.Fatal(err)
	}
	orders := futures.OrderList{*order}
	printer.Print(&orders)
}

func cancelOrder(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	orderID, _ := cmd.Flags().GetString("order-id")
	contract, _ := cmd.Flags().GetString("contract")

	if orderID == "" {
		if contract == "" {
			log.Fatal("--contract is required when canceling all open orders")
		}
		orders, err := client.CancelOpenOrders(contract)
		if err != nil {
			log.Fatal(err)
		}
		printer.Print(orders)
		return
	}
	order, err := client.CancelOrder(orderID)
	if err != nil {
		log.Fatal(err)
	}
	orders := futures.OrderList{*order}
	printer.Print(&orders)
}
