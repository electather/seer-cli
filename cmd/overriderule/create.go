package overriderule

import "github.com/spf13/cobra"

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new override rule",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.OverrideruleAPI.OverrideRulePost(ctx).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "OverrideRulePost")
	},
}

func init() {
	Cmd.AddCommand(createCmd)
}
