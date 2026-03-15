package status

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statusSystemCmd = &cobra.Command{
	Use:   "system",
	Short: "Get the system status of the Seerr API",
	Long:  `Call the status endpoint defined in the OpenAPI specification to get the service version and commit details.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		if isVerbose {
			cmd.Printf("Calling /status endpoint on %s...\n", viper.GetString("seerr.server"))
		}

		res, r, err := apiClient.PublicAPI.StatusGet(ctx).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "StatusGet")
	},
}

func init() {
	Cmd.AddCommand(statusSystemCmd)
}
