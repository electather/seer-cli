package tmdb

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "tmdb",
	Short: "Query TMDB metadata",
	Long:  `Retrieve TMDB metadata: regions, languages, studios, networks, genres, and backdrops.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
