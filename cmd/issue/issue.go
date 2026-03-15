package issue

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "issue",
	Short: "Manage issues",
	Long:  `Manage issues: list, get, create, delete, update status, and manage comments.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
