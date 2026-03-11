package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	api "seer-cli/pkg/api"
)

// overrideServerURL is used by tests to mock the API server
var overrideServerURL string

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status of the Seer API",
	Long:  `Call the status endpoint defined in the OpenAPI specification to get the service status.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Initialize the API configuration
		configuration := api.NewConfiguration()
		
		// Set server URL from viper
		configuration.Servers = api.ServerConfigurations{
			{
				URL: viper.GetString("server"),
				Description: "Configured Server",
			},
		}

		// Set API Key from viper if provided
		if apiKey := viper.GetString("api_key"); apiKey != "" {
			configuration.AddDefaultHeader("X-Api-Key", apiKey)
		}
		
		// If overridden by tests, use the mock server
		if overrideServerURL != "" {
			configuration.Servers = api.ServerConfigurations{
				{
					URL: overrideServerURL,
					Description: "Mock Server",
				},
			}
		}
		
		apiClient := api.NewAPIClient(configuration)
		
		// Example context
		ctx := context.Background()
		
		cmd.Println("Calling /status endpoint...")
		
		// Call a status or public endpoint if one exists
		// Note: Replace StatusGet with an actual endpoint from your schema
		res, r, err := apiClient.PublicAPI.StatusGet(ctx).Execute()
		if err != nil {
			return fmt.Errorf("error when calling StatusGet: %w\nFull HTTP response: %v", err, r)
		}
		
		cmd.Printf("Response from StatusGet: %v\n", res)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
