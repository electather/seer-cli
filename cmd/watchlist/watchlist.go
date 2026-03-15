package watchlist

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "watchlist",
	Short: "Manage the watchlist",
	Long:  `Manage watchlist items: add and remove entries.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
