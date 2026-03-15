package blocklist

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "blocklist",
	Short: "Manage the blocklist",
	Long:  `Manage blocklisted media items: list, add, get, and remove entries.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
