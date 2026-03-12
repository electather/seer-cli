package tmdb

import "github.com/spf13/cobra"

var languagesCmd = &cobra.Command{
	Use:   "languages",
	Short: "Get languages supported by TMDB",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.TmdbAPI.LanguagesGet(ctx).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "LanguagesGet")
	},
}

func init() {
	Cmd.AddCommand(languagesCmd)
}
