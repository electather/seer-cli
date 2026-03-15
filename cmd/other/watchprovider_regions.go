package other

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var watchproviderRegionsCmd = &cobra.Command{
	Use:     "watchprovider-regions",
	Short:   "List all available watch provider regions",
	Example: `  seerr-cli other watchprovider-regions`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, err := apiClient.OtherAPI.WatchprovidersRegionsGet(ctx).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "WatchprovidersRegionsGet")
	},
}

func init() {
	Cmd.AddCommand(watchproviderRegionsCmd)
}
