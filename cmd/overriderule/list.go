package overriderule

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all override rules",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.OverrideruleAPI.OverrideRuleGet(ctx).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "OverrideRuleGet")
	},
}

func init() {
	Cmd.AddCommand(listCmd)
}
