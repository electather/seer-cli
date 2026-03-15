package person

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "person",
	Short: "Get person details and credits",
	Long:  `Get person details and combined credits from the Seerr API.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
