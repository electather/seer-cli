package issue

import "github.com/spf13/cobra"

var updateStatusCmd = &cobra.Command{
	Use:   "update-status <issueId> <status>",
	Short: "Update an issue's status (open or resolved)",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.IssueAPI.IssueIssueIdStatusPost(ctx, args[0], args[1]).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "IssueIssueIdStatusPost")
	},
}

func init() {
	Cmd.AddCommand(updateStatusCmd)
}
