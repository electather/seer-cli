package tmdb

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var genresMovieCmd = &cobra.Command{
	Use:   "genres-movie",
	Short: "Get list of official TMDB movie genres",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		req := apiClient.TmdbAPI.GenresMovieGet(ctx)
		if cmd.Flags().Changed("language") {
			v, _ := cmd.Flags().GetString("language")
			req = req.Language(v)
		}
		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "GenresMovieGet")
	},
}

func init() {
	genresMovieCmd.Flags().String("language", "", "Language code (e.g. en)")
	Cmd.AddCommand(genresMovieCmd)
}
