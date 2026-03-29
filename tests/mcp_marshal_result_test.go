package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMCPRequestCreateFallbackOnDecodeError verifies that request_create returns
// the raw response body as a success result when the HTTP call succeeds (201)
// but the generated client fails to decode the response. This happens in the
// wild because Overseerr's POST /api/v1/request response embeds a UserSettings
// object with an "id" field that the generated model rejects via
// DisallowUnknownFields.
func TestMCPRequestCreateFallbackOnDecodeError(t *testing.T) {
	// The nested settings.id field is not in the generated UserSettings struct.
	// The strict decoder rejects it even though the HTTP response is 201.
	overseerrBody := `{
		"id": 50,
		"status": 2,
		"type": "tv",
		"media": {"mediaType": "tv"},
		"requestedBy": {
			"id": 1,
			"settings": {
				"id": 1,
				"locale": "",
				"discoverRegion": ""
			}
		}
	}`

	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/api/v1/request" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(overseerrBody))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.RequestCreateHandler()
	result := callTool(t, handler, map[string]any{
		"mediaType": "tv",
		"mediaId":   float64(83100),
		"seasons":   "all",
	})

	// Must not be an error — the request was created successfully.
	require.False(t, result.IsError,
		"2xx response with unknown fields must not produce a tool error")

	text := resultText(t, result)

	// The raw body must be returned so the caller has the full response.
	var parsed map[string]interface{}
	require.NoError(t, json.Unmarshal([]byte(text), &parsed))
	assert.Equal(t, float64(50), parsed["id"])
	assert.Equal(t, float64(2), parsed["status"])
}

// TestMCPRequestCreateRealErrorStillFails verifies that a genuine Overseerr
// error (non-2xx) is still surfaced as a tool error, not silently swallowed by
// the fallback path.
func TestMCPRequestCreateRealErrorStillFails(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/api/v1/request" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"message":"Invalid request"}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.RequestCreateHandler()
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{
		"mediaType": "tv",
		"mediaId":   float64(83100),
		"seasons":   "all",
	}
	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.IsError, "non-2xx response must produce a tool error")
}
