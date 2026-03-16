package tests

import (
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

func TestSafeLogPath(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		routeToken string
		want       string
	}{
		{
			name: "plain mcp path unchanged",
			path: "/mcp",
			want: "/mcp",
		},
		{
			name: "plain mcp sse path unchanged",
			path: "/mcp/sse",
			want: "/mcp/sse",
		},
		{
			name:       "route token in prefix is redacted",
			path:       "/abc123/mcp",
			routeToken: "abc123",
			want:       "/{redacted}/mcp",
		},
		{
			name:       "route token sse path is redacted",
			path:       "/abc123/mcp/sse",
			routeToken: "abc123",
			want:       "/{redacted}/mcp/sse",
		},
		{
			name:       "route token exact match is redacted",
			path:       "/abc123",
			routeToken: "abc123",
			want:       "/{redacted}",
		},
		{
			name:       "unrelated path is unchanged in route token mode",
			path:       "/health",
			routeToken: "abc123",
			want:       "/health",
		},
		{
			name: "no route token returns path unchanged",
			path: "/unexpected/path",
			want: "/unexpected/path",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := cmdmcp.SafeLogPath(tc.path, tc.routeToken)
			assert.Equal(t, tc.want, got)
			// Verify the raw token never appears in the output.
			if tc.routeToken != "" {
				assert.NotContains(t, got, tc.routeToken)
			}
		})
	}
}

func TestSafeLogQuery(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  string
	}{
		{
			name:  "empty query unchanged",
			query: "",
			want:  "",
		},
		{
			name:  "api_key value is redacted",
			query: "api_key=secret123",
			want:  "api_key={redacted}",
		},
		{
			name:  "api_key redacted while other params preserved",
			query: "api_key=secret&page=1",
			want:  "api_key={redacted}&page=1",
		},
		{
			name:  "query without api_key is unchanged",
			query: "page=1&limit=10",
			want:  "page=1&limit=10",
		},
		{
			name:  "api_key in middle of query is redacted",
			query: "page=1&api_key=mysecret&limit=10",
			want:  "page=1&api_key={redacted}&limit=10",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := cmdmcp.SafeLogQuery(tc.query)
			assert.Equal(t, tc.want, got)
		})
	}
}
