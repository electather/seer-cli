package tmdb

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var backdropsCmd = &cobra.Command{
	Use:   "backdrops",
	Short: "Get backdrops of trending items",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, err := apiClient.TmdbAPI.BackdropsGet(ctx).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "BackdropsGet")
	},
}

func init() {
	Cmd.AddCommand(backdropsCmd)
}
