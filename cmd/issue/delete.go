package issue

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:   "delete <issueId>",
	Short: "Delete an issue",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		r, err := apiClient.IssueAPI.IssueIssueIdDelete(ctx, args[0]).Execute()
		return handle204Response(cmd, r, err, isVerbose, "IssueIssueIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCmd)
}
