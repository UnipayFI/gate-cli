package unified

import (
	"log"

	"github.com/UnipayFI/gate-cli/printer"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

func InitLoanCmds() []*cobra.Command {
	borrowCmd := &cobra.Command{
		Use:   "borrow",
		Short: "Borrow a currency on the unified account",
		Long: `Borrow a currency on the unified account. The amount must sit between the
currency's minimum and the platform/user borrow limit.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#borrow-or-repay`,
		Run: borrow,
	}

	repayCmd := &cobra.Command{
		Use:   "repay",
		Short: "Repay a borrowed currency on the unified account",
		Long: `Repay a borrowed currency. Pass --amount for a partial repayment, or --all to
clear the full outstanding balance.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#borrow-or-repay`,
		Run: repay,
	}

	loansCmd := &cobra.Command{
		Use:   "loans",
		Short: "List outstanding borrowings",
		Long: `List the unified account's outstanding borrowings.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#list-loans`,
		Run: listLoans,
	}

	loanRecordsCmd := &cobra.Command{
		Use:   "loan-records",
		Short: "List borrow/repay history",
		Long: `List the unified account's borrow and repayment history.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#get-loan-records`,
		Run: listLoanRecords,
	}

	interestRecordsCmd := &cobra.Command{
		Use:   "interest-records",
		Short: "List interest-deduction history",
		Long: `List the unified account's interest-deduction history.

Docs Link: https://www.gate.io/docs/developers/apiv4/en/#interest-records-for-the-unified-account`,
		Run: listInterestRecords,
	}

	borrowCmd.Flags().StringP("currency", "c", "", "currency, e.g. USDT")
	borrowCmd.Flags().StringP("amount", "a", "", "amount to borrow (decimal string)")
	borrowCmd.MarkFlagRequired("currency")
	borrowCmd.MarkFlagRequired("amount")

	repayCmd.Flags().StringP("currency", "c", "", "currency, e.g. USDT")
	repayCmd.Flags().StringP("amount", "a", "", "amount to repay (decimal string); omit with --all")
	repayCmd.Flags().Bool("all", false, "repay the full outstanding balance")
	repayCmd.MarkFlagRequired("currency")

	loansCmd.Flags().StringP("currency", "c", "", "filter a single currency")
	loansCmd.Flags().StringP("type", "t", "", "loan type: platform or margin")
	loansCmd.Flags().IntP("limit", "l", 100, "max records (default 100)")
	loansCmd.Flags().Int("page", 1, "page number")

	loanRecordsCmd.Flags().StringP("currency", "c", "", "filter a single currency")
	loanRecordsCmd.Flags().StringP("type", "t", "", "record type: borrow or repay")
	loanRecordsCmd.Flags().IntP("limit", "l", 100, "max records (default 100)")
	loanRecordsCmd.Flags().Int("page", 1, "page number")

	interestRecordsCmd.Flags().StringP("currency", "c", "", "filter a single currency")
	interestRecordsCmd.Flags().StringP("type", "t", "", "loan type: platform or margin")
	interestRecordsCmd.Flags().IntP("limit", "l", 100, "max records (default 100)")
	interestRecordsCmd.Flags().Int("page", 1, "page number")
	interestRecordsCmd.Flags().StringP("from", "a", "", "start time (unix seconds or \"YYYY-MM-DD HH:MM:SS\")")
	interestRecordsCmd.Flags().StringP("to", "e", "", "end time (unix seconds or \"YYYY-MM-DD HH:MM:SS\")")

	return []*cobra.Command{borrowCmd, repayCmd, loansCmd, loanRecordsCmd, interestRecordsCmd}
}

func borrow(cmd *cobra.Command, _ []string) {
	client := newClient()
	currency, _ := cmd.Flags().GetString("currency")
	amountStr, _ := cmd.Flags().GetString("amount")
	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		log.Fatal(err)
	}
	result, err := client.BorrowOrRepay(currency, amount, "borrow", false)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(result)
}

func repay(cmd *cobra.Command, _ []string) {
	client := newClient()
	currency, _ := cmd.Flags().GetString("currency")
	all, _ := cmd.Flags().GetBool("all")
	amountStr, _ := cmd.Flags().GetString("amount")

	amount := decimal.Zero
	if amountStr != "" {
		var err error
		amount, err = decimal.NewFromString(amountStr)
		if err != nil {
			log.Fatal(err)
		}
	}
	if !all && amount.IsZero() {
		log.Fatal("either --amount or --all is required")
	}

	result, err := client.BorrowOrRepay(currency, amount, "repay", all)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(result)
}

func listLoans(cmd *cobra.Command, _ []string) {
	client := newClient()
	currency, _ := cmd.Flags().GetString("currency")
	loanType, _ := cmd.Flags().GetString("type")
	limit, _ := cmd.Flags().GetInt("limit")
	page, _ := cmd.Flags().GetInt("page")

	loans, err := client.GetLoans(currency, loanType, page, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(loans)
}

func listLoanRecords(cmd *cobra.Command, _ []string) {
	client := newClient()
	currency, _ := cmd.Flags().GetString("currency")
	recordType, _ := cmd.Flags().GetString("type")
	limit, _ := cmd.Flags().GetInt("limit")
	page, _ := cmd.Flags().GetInt("page")

	records, err := client.GetLoanRecords(currency, recordType, page, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(records)
}

func listInterestRecords(cmd *cobra.Command, _ []string) {
	client := newClient()
	currency, _ := cmd.Flags().GetString("currency")
	loanType, _ := cmd.Flags().GetString("type")
	limit, _ := cmd.Flags().GetInt("limit")
	page, _ := cmd.Flags().GetInt("page")
	from := parseTimeOrFatal(cmd, "--from", "from")
	to := parseTimeOrFatal(cmd, "--to", "to")

	records, err := client.GetInterestRecords(currency, loanType, from, to, page, limit)
	if err != nil {
		log.Fatal(err)
	}
	printer.Print(records)
}
