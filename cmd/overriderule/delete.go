package overriderule

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <ruleId>",
	Short: "Delete an override rule",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid rule ID: %s", args[0])
		}
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, apiErr := apiClient.OverrideruleAPI.OverrideRuleRuleIdDelete(ctx, float32(id)).Execute()
		return handleResponse(cmd, r, apiErr, res, isVerbose, "OverrideRuleRuleIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCmd)
}
