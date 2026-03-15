package issue

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Get issue counts",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, err := apiClient.IssueAPI.IssueCountGet(ctx).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "IssueCountGet")
	},
}

func init() {
	Cmd.AddCommand(countCmd)
}
