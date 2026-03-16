package tests

import (
	"net/http"
	"strings"
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

func TestMoviesRatingsCombinedHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasPrefix(r.URL.Path, "/api/v1/movie/") && strings.HasSuffix(r.URL.Path, "/ratingscombined") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"rt":{"criticsScore":94,"criticsRating":"Certified Fresh"},"imdb":{"id":"tt0816692","rating":8.7}}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	handler := cmdmcp.MoviesRatingsCombinedHandler()
	result := callTool(t, handler, map[string]any{"movieId": float64(157336)})
	text := resultText(t, result)

	assert.Contains(t, text, "rt")
	assert.Contains(t, text, "imdb")
}
