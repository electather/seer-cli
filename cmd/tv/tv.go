package tv

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "tv",
	Short: "Get TV show details, recommendations, and more",
	Long:  `Get TV show details, season info, recommendations, similar shows, and ratings from the Seerr API.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
