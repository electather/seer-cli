package issue

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var getCommentCmd = &cobra.Command{
	Use:   "get-comment <commentId>",
	Short: "Get a single issue comment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, err := apiClient.IssueAPI.IssueCommentCommentIdGet(ctx, args[0]).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "IssueCommentCommentIdGet")
	},
}

func init() {
	Cmd.AddCommand(getCommentCmd)
}
