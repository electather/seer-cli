package tests

import (
	"net/http"
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

func TestSearchCompanyHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/v1/search/company" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"page":1,"totalPages":1,"totalResults":1,"results":[{"id":1,"name":"Warner Bros."}]}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.SearchCompanyHandler()
	result := callTool(t, handler, map[string]any{"query": "Warner"})
	text := resultText(t, result)

	assert.Contains(t, text, "Warner Bros.")
}

func TestSearchKeywordHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/api/v1/search/keyword" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"page":1,"totalPages":1,"totalResults":1,"results":[{"id":878,"name":"science fiction"}]}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.SearchKeywordHandler()
	result := callTool(t, handler, map[string]any{"query": "sci-fi"})
	text := resultText(t, result)

	assert.Contains(t, text, "science fiction")
}
