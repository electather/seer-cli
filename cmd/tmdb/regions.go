package tmdb

import "github.com/spf13/cobra"

var regionsCmd = &cobra.Command{
	Use:   "regions",
	Short: "Get regions supported by TMDB",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.TmdbAPI.RegionsGet(ctx).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "RegionsGet")
	},
}

func init() {
	Cmd.AddCommand(regionsCmd)
}
