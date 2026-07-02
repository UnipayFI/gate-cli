package futures

import (
	"log"
	"time"

	"github.com/UnipayFI/gate-cli/printer"
	"github.com/spf13/cobra"
)

func InitPositionCmds() []*cobra.Command {
	positionCmd := &cobra.Command{
		Use:   "position",
		Short: "Position management commands",
		Long:  `Manage positions: list, get, margin, leverage, dual-mode, close-history.`,
	}

	positionListCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List positions",
		Long: `List positions for the settle currency. Use --holding to only show open positions.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#list-all-positions-of-a-user`,
		Run: listPositions,
	}

	positionGetCmd := &cobra.Command{
		Use:   "get",
		Short: "Get a single position",
		Long: `Get a single position by contract.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#get-single-position`,
		Run: getPosition,
	}

	positionMarginCmd := &cobra.Command{
		Use:   "margin",
		Short: "Update isolated position margin",
		Long: `Add or reduce isolated position margin. --change is signed: positive adds, negative reduces.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#update-position-margin`,
		Run: updateMargin,
	}

	positionLeverageCmd := &cobra.Command{
		Use:   "leverage",
		Short: "Update position leverage",
		Long: `Update position leverage. Use "0" for cross margin, otherwise the isolated leverage.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#update-position-leverage`,
		Run: updateLeverage,
	}

	positionDualModeCmd := &cobra.Command{
		Use:   "dual-mode",
		Short: "Enable or disable hedge (dual) mode",
		Long: `Enable or disable dual (hedge) position mode for the settle currency.
Cannot be changed while holding positions or open orders.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#enable-or-disable-dual-mode`,
		Run: setDualMode,
	}

	positionCloseCmd := &cobra.Command{
		Use:   "close-history",
		Short: "Query position close history",
		Long: `Query the position close (realised PnL) history.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#list-position-close-history`,
		Run: showPositionClose,
	}

	positionListCmd.Flags().Bool("holding", false, "only positions currently holding size")

	positionGetCmd.Flags().StringP("contract", "c", "", "contract, e.g. BTC_USDT")
	positionGetCmd.MarkFlagRequired("contract")

	positionMarginCmd.Flags().StringP("contract", "c", "", "contract, e.g. BTC_USDT")
	positionMarginCmd.Flags().String("change", "", "signed margin change (e.g. 10 to add, -10 to reduce)")
	positionMarginCmd.MarkFlagRequired("contract")
	positionMarginCmd.MarkFlagRequired("change")

	positionLeverageCmd.Flags().StringP("contract", "c", "", "contract, e.g. BTC_USDT")
	positionLeverageCmd.Flags().String("leverage", "", "leverage multiple, \"0\" for cross margin")
	positionLeverageCmd.MarkFlagRequired("contract")
	positionLeverageCmd.MarkFlagRequired("leverage")

	positionDualModeCmd.Flags().Bool("enable", false, "true to enable hedge mode, false to disable")
	positionDualModeCmd.MarkFlagRequired("enable")

	positionCloseCmd.Flags().StringP("contract", "c", "", "filter a single contract")
	positionCloseCmd.Flags().IntP("limit", "l", 100, "max records (default 100)")

	positionCmd.AddCommand(
		positionListCmd,
		positionGetCmd,
		positionMarginCmd,
		positionLeverageCmd,
		positionDualModeCmd,
		positionCloseCmd,
	)
	return []*cobra.Command{positionCmd}
}

func listPositions(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	holding, _ := cmd.Flags().GetBool("holding")
	positions, err := client.GetPositions(holding)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(positions)
}

func getPosition(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	contract, _ := cmd.Flags().GetString("contract")
	positions, err := client.GetPosition(contract)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(positions)
}

func updateMargin(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	contract, _ := cmd.Flags().GetString("contract")
	change, _ := cmd.Flags().GetString("change")
	positions, err := client.UpdateMargin(contract, change)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(positions)
}

func updateLeverage(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	contract, _ := cmd.Flags().GetString("contract")
	leverage, _ := cmd.Flags().GetString("leverage")
	positions, err := client.UpdateLeverage(contract, leverage)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(positions)
}

func setDualMode(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	enable, _ := cmd.Flags().GetBool("enable")
	result, err := client.SetDualMode(enable)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(result)
}

func showPositionClose(cmd *cobra.Command, _ []string) {
	client := newClient(cmd)
	contract, _ := cmd.Flags().GetString("contract")
	limit, _ := cmd.Flags().GetInt("limit")
	records, err := client.GetPositionClose(contract, time.Time{}, time.Time{}, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(records)
}
