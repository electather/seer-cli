package other

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var certificationsMovieCmd = &cobra.Command{
	Use:     "certifications-movie",
	Short:   "List movie certifications from TMDB",
	Example: `  seerr-cli other certifications-movie`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, err := apiClient.OtherAPI.CertificationsMovieGet(ctx).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "CertificationsMovieGet")
	},
}

func init() {
	Cmd.AddCommand(certificationsMovieCmd)
}
