package mcp

import (
	"context"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	api "seer-cli/pkg/api"
)

// OverrideServerURL is used by tests to redirect API calls to a mock server.
var OverrideServerURL string

type contextKey string

const apiKeyCtxKey contextKey = "seerApiKey"

// APIKeyContextKey is exported for tests to inject a key directly.
const APIKeyContextKey = apiKeyCtxKey

func apiKeyFromContext(ctx context.Context) string {
	v, _ := ctx.Value(apiKeyCtxKey).(string)
	return v
}

var Cmd = &cobra.Command{
	Use:   "mcp",
	Short: "Run a Model Context Protocol (MCP) server exposing the Seer API",
	Long:  `Start an MCP server that exposes the Seer API as tools for use by AI agents.`,
	Example: `  # Start the MCP server using stdio transport (for Claude Desktop)
  seer-cli mcp serve

  # Start the MCP server over HTTP with a Bearer token
  seer-cli mcp serve --transport http --auth-token mysecret`,
}

// newAPIClientWithKey builds a client using apiKey, falling back to Viper when empty.
func newAPIClientWithKey(apiKey string) *api.APIClient {
	configuration := api.NewConfiguration()
	serverURL := viper.GetString("server")
	if !strings.HasSuffix(serverURL, "/api/v1") {
		serverURL = strings.TrimSuffix(serverURL, "/") + "/api/v1"
	}
	configuration.Servers = api.ServerConfigurations{{URL: serverURL, Description: "Configured Server"}}
	key := apiKey
	if key == "" {
		key = viper.GetString("api_key")
	}
	if key != "" {
		configuration.AddDefaultHeader("X-Api-Key", key)
	}
	if OverrideServerURL != "" {
		configuration.Servers = api.ServerConfigurations{{URL: OverrideServerURL, Description: "Mock Server"}}
	}
	return api.NewAPIClient(configuration)
}

func newAPIClient() (*api.APIClient, context.Context) {
	return newAPIClientWithKey(""), context.Background()
}

// NewAPIClientForTest is an exported alias used by tests.
func NewAPIClientForTest() (*api.APIClient, context.Context) {
	return newAPIClient()
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
