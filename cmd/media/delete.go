package media

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <mediaId>",
	Short: "Delete a media item",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		r, err := apiClient.MediaAPI.MediaMediaIdDelete(ctx, args[0]).Execute()
		return apiutil.Handle204Response(cmd, r, err, isVerbose, "MediaMediaIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCmd)
}
