package watchlist

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <tmdbId>",
	Short: "Remove an item from the watchlist",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		r, err := apiClient.WatchlistAPI.WatchlistTmdbIdDelete(ctx, args[0]).Execute()
		return apiutil.Handle204Response(cmd, r, err, isVerbose, "WatchlistTmdbIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCmd)
}
