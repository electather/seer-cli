package status

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statusAppdataCmd = &cobra.Command{
	Use:   "appdata",
	Short: "Get application data volume status",
	Long:  `For Docker installs, returns whether or not the volume mount was configured properly. Always returns true for non-Docker installs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		if isVerbose {
			cmd.Printf("Calling /status/appdata endpoint on %s...\n", viper.GetString("seerr.server"))
		}

		res, r, err := apiClient.PublicAPI.StatusAppdataGet(ctx).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "StatusAppdataGet")
	},
}

func init() {
	Cmd.AddCommand(statusAppdataCmd)
}
