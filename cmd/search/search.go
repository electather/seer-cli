package search

import (
	"context"
	"net/http"
	"strings"

	"seerr-cli/cmd/apiutil"
	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// encodingRoundTripper is a custom RoundTripper that replaces '+' with '%20' in query parameters.
// It also removes problematic parameters from certain endpoints if needed.
type encodingRoundTripper struct {
	Proxied http.RoundTripper
}

func (ert *encodingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Special handling for /discover/trending: remove mediaType and timeWindow if they are defaults
	// because some server versions don't support them and fail with 400.
	if strings.Contains(req.URL.Path, "/discover/trending") {
		q := req.URL.Query()
		if q.Get("mediaType") == "all" {
			q.Del("mediaType")
		}
		if q.Get("timeWindow") == "day" {
			q.Del("timeWindow")
		}
		req.URL.RawQuery = q.Encode()
	}

	// Replace '+' with '%20' in the RawQuery
	req.URL.RawQuery = strings.ReplaceAll(req.URL.RawQuery, "+", "%20")

	return ert.Proxied.RoundTrip(req)
}

var Cmd = &cobra.Command{
	Use:   "search",
	Short: "Search for movies, TV shows, people, and more",
	Long:  `Search for various resources from the Seerr API including movies, TV shows, people, keywords, and companies.`,
}

func newAPIClient() (*api.APIClient, context.Context, bool) {
	return apiutil.NewAPIClientWithTransport(&encodingRoundTripper{Proxied: http.DefaultTransport}), context.Background(), viper.GetBool("verbose")
}

func init() {
	// Subcommands will be added in their respective files' init() functions
	// but we can also do it here if we want to be explicit.
}
