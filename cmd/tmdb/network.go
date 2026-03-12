package tmdb

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var networkCmd = &cobra.Command{
	Use:   "network <networkId>",
	Short: "Get TV network details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid network ID: %s", args[0])
		}
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, apiErr := apiClient.TmdbAPI.NetworkNetworkIdGet(ctx, float32(id)).Execute()
		return handleResponse(cmd, r, apiErr, res, isVerbose, "NetworkNetworkIdGet")
	},
}

func init() {
	Cmd.AddCommand(networkCmd)
}
