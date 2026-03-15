package mcp

import (
	"context"
	"fmt"

	"seerr-cli/cmd/apiutil"
	api "seerr-cli/pkg/api"

	mcplib "github.com/mark3labs/mcp-go/mcp"
	"github.com/spf13/cobra"
)

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
	Short: "Run a Model Context Protocol (MCP) server exposing the Seerr API",
	Long:  `Start an MCP server that exposes the Seerr API as tools for use by AI agents.`,
	Example: `  # Start the MCP server using stdio transport (for Claude Desktop)
  seerr-cli mcp serve

  # Start the MCP server over HTTP with a Bearer token
  seerr-cli mcp serve --transport http --auth-token mysecret`,
}

// newAPIClientWithKey builds a client using apiKey, falling back to Viper when empty.
func newAPIClientWithKey(apiKey string) *api.APIClient {
	return apiutil.NewAPIClientWithKey(apiKey)
}

func newAPIClient() (*api.APIClient, context.Context) {
	return apiutil.NewAPIClientWithKey(""), context.Background()
}

// NewAPIClientForTest is an exported alias used by tests.
func NewAPIClientForTest() (*api.APIClient, context.Context) {
	return newAPIClient()
}

// apiToolError converts a Seerr API error into a visible MCP tool result error.
// Using mcp.NewToolResultError (IsError=true) instead of returning a Go error
// makes the actual Seerr error message visible in the MCP client (e.g. Claude)
// rather than the generic "Error occurred during tool execution" wrapper.
func apiToolError(label string, err error) (*mcplib.CallToolResult, error) {
	msg := fmt.Sprintf("%s: %v", label, err)
	if e, ok := err.(*api.GenericOpenAPIError); ok && len(e.Body()) > 0 {
		msg = fmt.Sprintf("%s: %s (HTTP %v)", label, e.Body(), err)
	}
	if mcpLog != nil {
		mcpLog.Warn("tool error", "label", label, "error", err)
	}
	return mcplib.NewToolResultError(msg), nil
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
