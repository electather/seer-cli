package tests

import (
	"net/http"
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

func TestSearchDiscoverMoviesGenreFilter(t *testing.T) {
	var capturedQuery string
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		switch r.URL.Path {
		case "/api/v1/discover/movies":
			capturedQuery = r.URL.RawQuery
			w.Write([]byte(`{"page":1,"totalPages":1,"totalResults":0,"results":[]}`))
		case "/api/v1/genres/movie":
			w.Write([]byte(`[]`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.SearchDiscoverMoviesHandler()
	callTool(t, handler, map[string]any{"genre": "18"})

	assert.Contains(t, capturedQuery, "genre=18")
}

func TestSearchDiscoverMoviesSortBy(t *testing.T) {
	var capturedQuery string
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		switch r.URL.Path {
		case "/api/v1/discover/movies":
			capturedQuery = r.URL.RawQuery
			w.Write([]byte(`{"page":1,"totalPages":1,"totalResults":0,"results":[]}`))
		case "/api/v1/genres/movie":
			w.Write([]byte(`[]`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.SearchDiscoverMoviesHandler()
	callTool(t, handler, map[string]any{"sortBy": "popularity.desc"})

	assert.Contains(t, capturedQuery, "sortBy=popularity.desc")
}

func TestSearchDiscoverMoviesDateRange(t *testing.T) {
	var capturedQuery string
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		switch r.URL.Path {
		case "/api/v1/discover/movies":
			capturedQuery = r.URL.RawQuery
			w.Write([]byte(`{"page":1,"totalPages":1,"totalResults":0,"results":[]}`))
		case "/api/v1/genres/movie":
			w.Write([]byte(`[]`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.SearchDiscoverMoviesHandler()
	callTool(t, handler, map[string]any{"primaryReleaseDateGte": "2023-01-01"})

	assert.Contains(t, capturedQuery, "primaryReleaseDateGte=2023-01-01")
}

func TestSearchDiscoverTVNetworkFilter(t *testing.T) {
	var capturedQuery string
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		switch r.URL.Path {
		case "/api/v1/discover/tv":
			capturedQuery = r.URL.RawQuery
			w.Write([]byte(`{"page":1,"totalPages":1,"totalResults":0,"results":[]}`))
		case "/api/v1/genres/tv":
			w.Write([]byte(`[]`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.SearchDiscoverTVHandler()
	callTool(t, handler, map[string]any{"network": "1"})

	assert.Contains(t, capturedQuery, "network=1")
}

func TestSearchDiscoverTVLanguageFilter(t *testing.T) {
	var capturedQuery string
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		switch r.URL.Path {
		case "/api/v1/discover/tv":
			capturedQuery = r.URL.RawQuery
			w.Write([]byte(`{"page":1,"totalPages":1,"totalResults":0,"results":[]}`))
		case "/api/v1/genres/tv":
			w.Write([]byte(`[]`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.SearchDiscoverTVHandler()
	callTool(t, handler, map[string]any{"language": "fr"})

	assert.Contains(t, capturedQuery, "language=fr")
}
