package overriderule

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all override rules",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, err := apiClient.OverrideruleAPI.OverrideRuleGet(ctx).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "OverrideRuleGet")
	},
}

func init() {
	Cmd.AddCommand(listCmd)
}
