package issue

import "github.com/spf13/cobra"

var deleteCommentCmd = &cobra.Command{
	Use:   "delete-comment <commentId>",
	Short: "Delete an issue comment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		r, err := apiClient.IssueAPI.IssueCommentCommentIdDelete(ctx, args[0]).Execute()
		return handle204Response(cmd, r, err, isVerbose, "IssueCommentCommentIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCommentCmd)
}
