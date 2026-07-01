package cmd

import (
	"errors"

	"github.com/UnipayFI/gate-cli/config"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:               "gate-cli",
	Short:             "Gate.io API for CLI version",
	PersistentPreRunE: checkCredentials,
}

func init() {
	initCommandConfig()
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func initCommandConfig() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
	RootCmd.PersistentFlags().BoolVar(&config.Config.OutputJSON, "json", false, "Output JSON instead of a table")
}

func checkCredentials(cmd *cobra.Command, args []string) error {
	if config.Config.APIKey == "" || config.Config.APISecret == "" {
		return errors.New("GATE_API_KEY and GATE_API_SECRET must be set")
	}
	return nil
}
