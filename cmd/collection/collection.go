package collection

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "collection",
	Short: "Manage collections",
	Long:  `Access collection details from the Seerr API.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
