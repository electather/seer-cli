package tmdb

import "github.com/spf13/cobra"

var backdropsCmd = &cobra.Command{
	Use:   "backdrops",
	Short: "Get backdrops of trending items",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.TmdbAPI.BackdropsGet(ctx).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "BackdropsGet")
	},
}

func init() {
	Cmd.AddCommand(backdropsCmd)
}
