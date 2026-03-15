package movies

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "movies",
	Short: "Get movie details, recommendations, and more",
	Long:  `Get movie details, recommendations, similar movies, and ratings from the Seerr API.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
