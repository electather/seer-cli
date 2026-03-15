package request

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <requestId>",
	Short: "Get a specific media request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, err := apiClient.RequestAPI.RequestRequestIdGet(ctx, args[0]).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "RequestRequestIdGet")
	},
}

func init() {
	Cmd.AddCommand(getCmd)
}
