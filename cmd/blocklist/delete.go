package blocklist

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <tmdbId>",
	Short: "Remove media from the blocklist",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		r, err := apiClient.BlocklistAPI.BlocklistTmdbIdDelete(ctx, args[0]).Execute()
		return apiutil.Handle204Response(cmd, r, err, isVerbose, "BlocklistTmdbIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCmd)
}
