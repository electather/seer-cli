package apiutil

import (
	"context"
	"net/http"
	"strings"

	api "seerr-cli/pkg/api"

	"github.com/spf13/viper"
)

// OverrideServerURL is used by tests to redirect API calls to a mock server.
var OverrideServerURL string

// NewAPIClient builds a standard API client using Viper config.
func NewAPIClient() (*api.APIClient, context.Context, bool) {
	return NewAPIClientWithTransport(nil), context.Background(), viper.GetBool("verbose")
}

// NewAPIClientWithTransport builds an API client with an optional custom HTTP transport.
// Pass nil to use the default transport.
func NewAPIClientWithTransport(transport http.RoundTripper) *api.APIClient {
	return NewAPIClientWithKeyAndTransport("", transport)
}

// NewAPIClientWithKey builds a client using apiKey, falling back to Viper when empty.
func NewAPIClientWithKey(apiKey string) *api.APIClient {
	return NewAPIClientWithKeyAndTransport(apiKey, nil)
}

// NewAPIClientWithKeyAndTransport is the base constructor used by all other helpers.
func NewAPIClientWithKeyAndTransport(apiKey string, transport http.RoundTripper) *api.APIClient {
	configuration := api.NewConfiguration()
	serverURL := viper.GetString("seerr.server")
	if !strings.HasSuffix(serverURL, "/api/v1") {
		serverURL = strings.TrimSuffix(serverURL, "/") + "/api/v1"
	}
	configuration.Servers = api.ServerConfigurations{{URL: serverURL, Description: "Configured Server"}}
	key := apiKey
	if key == "" {
		key = viper.GetString("seerr.api_key")
	}
	if key != "" {
		configuration.AddDefaultHeader("X-Api-Key", key)
	}
	if transport != nil {
		configuration.HTTPClient = &http.Client{Transport: transport}
	}
	if OverrideServerURL != "" {
		configuration.Servers = api.ServerConfigurations{{URL: OverrideServerURL, Description: "Mock Server"}}
	}
	return api.NewAPIClient(configuration)
}
