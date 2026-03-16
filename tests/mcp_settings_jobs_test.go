package tests

import (
	"net/http"
	"strings"
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

func TestSettingsJobsCancelHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/cancel") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":"plex-sync","name":"Plex Sync","type":"process","running":false}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.SettingsJobsCancelHandler()
	result := callTool(t, handler, map[string]any{"jobId": "plex-sync"})
	text := resultText(t, result)

	assert.Contains(t, text, "plex-sync")
}

func TestSettingsJobsScheduleHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/schedule") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":"plex-sync","name":"Plex Sync","type":"process","running":false}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.SettingsJobsScheduleHandler()
	result := callTool(t, handler, map[string]any{"jobId": "plex-sync", "schedule": "0 */5 * * *"})
	text := resultText(t, result)

	assert.Contains(t, text, "plex-sync")
}
