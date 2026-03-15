package other

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "other",
	Short: "Miscellaneous lookups (keywords, watch providers, certifications)",
	Long:  `Look up keywords, watch provider regions and availability, and content certifications.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
