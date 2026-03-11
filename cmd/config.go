package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CLI configuration",
	Long:  `View or update the configuration for the Seer CLI.`,
}

func init() {
	rootCmd.AddCommand(configCmd)
}
