package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seer-cli",
	Short: "A CLI to interact with the Seer API",
	Long:  `A command line interface to call endpoints defined in the Seer OpenAPI specification.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior when no subcommands are provided
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
