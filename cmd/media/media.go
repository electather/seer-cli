package media

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "media",
	Short: "Manage media items",
	Long:  `Manage media items including listing, deleting, updating status, and viewing watch data.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
