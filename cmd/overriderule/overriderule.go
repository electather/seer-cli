package overriderule

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "overriderule",
	Short: "Manage override rules",
	Long:  `Manage override rules: list, create, update, and delete.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
