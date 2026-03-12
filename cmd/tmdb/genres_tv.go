package tmdb

import "github.com/spf13/cobra"

var genresTvCmd = &cobra.Command{
	Use:   "genres-tv",
	Short: "Get list of official TMDB TV genres",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		req := apiClient.TmdbAPI.GenresTvGet(ctx)
		if cmd.Flags().Changed("language") {
			v, _ := cmd.Flags().GetString("language")
			req = req.Language(v)
		}
		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "GenresTvGet")
	},
}

func init() {
	genresTvCmd.Flags().String("language", "", "Language code (e.g. en)")
	Cmd.AddCommand(genresTvCmd)
}
