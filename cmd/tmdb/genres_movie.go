package tmdb

import "github.com/spf13/cobra"

var genresMovieCmd = &cobra.Command{
	Use:   "genres-movie",
	Short: "Get list of official TMDB movie genres",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		req := apiClient.TmdbAPI.GenresMovieGet(ctx)
		if cmd.Flags().Changed("language") {
			v, _ := cmd.Flags().GetString("language")
			req = req.Language(v)
		}
		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "GenresMovieGet")
	},
}

func init() {
	genresMovieCmd.Flags().String("language", "", "Language code (e.g. en)")
	Cmd.AddCommand(genresMovieCmd)
}
