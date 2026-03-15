package issue

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var deleteCommentCmd = &cobra.Command{
	Use:   "delete-comment <commentId>",
	Short: "Delete an issue comment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		r, err := apiClient.IssueAPI.IssueCommentCommentIdDelete(ctx, args[0]).Execute()
		return apiutil.Handle204Response(cmd, r, err, isVerbose, "IssueCommentCommentIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCommentCmd)
}
