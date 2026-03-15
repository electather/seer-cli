package issue

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <issueId>",
	Short: "Delete an issue",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		r, err := apiClient.IssueAPI.IssueIssueIdDelete(ctx, args[0]).Execute()
		return apiutil.Handle204Response(cmd, r, err, isVerbose, "IssueIssueIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCmd)
}
