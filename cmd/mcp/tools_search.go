package mcp

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerSearchTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("search_multi",
			mcp.WithDescription("Search for movies, TV shows, and people"),
			mcp.WithOpenWorldHintAnnotation(true),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithString("query", mcp.Required(), mcp.Description("Search query")),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchMultiHandler(),
	)

	s.AddTool(
		mcp.NewTool("search_discover_movies",
			mcp.WithDescription("Discover movies"),
			mcp.WithOpenWorldHintAnnotation(true),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchDiscoverMoviesHandler(),
	)

	s.AddTool(
		mcp.NewTool("search_discover_tv",
			mcp.WithDescription("Discover TV shows"),
			mcp.WithOpenWorldHintAnnotation(true),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchDiscoverTVHandler(),
	)

	s.AddTool(
		mcp.NewTool("search_trending",
			mcp.WithDescription("Get trending movies and TV shows"),
			mcp.WithOpenWorldHintAnnotation(true),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchTrendingHandler(),
	)
}

func SearchMultiHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query := req.GetString("query", "")
		if query == "" {
			return nil, fmt.Errorf("query is required")
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))

		// Fetch all genres concurrently while the search request is in flight.
		genresCh := make(chan GenreMap, 1)
		go func() { genresCh <- FetchAllGenres(callCtx, client) }()

		// Use a raw HTTP request to avoid the broken union-type unmarshal in the
		// generated client (TV results are incorrectly parsed as PersonResult) and
		// to ensure spaces are encoded as %20 rather than + in the query string.
		params := url.Values{}
		params.Set("query", query)
		if page := req.GetFloat("page", 0); page > 0 {
			params.Set("page", strconv.Itoa(int(page)))
		}
		b, err := apiutil.RawGet(callCtx, client, "/search", params)
		if err != nil {
			return apiToolError("SearchGet failed", err)
		}
		enriched, err := EnrichResultsPage(b, <-genresCh)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(enriched)), nil
	}
}

func SearchDiscoverMoviesHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))

		// Fetch movie genres concurrently while the discover request is in flight.
		genresCh := make(chan GenreMap, 1)
		go func() { genresCh <- FetchMovieGenres(callCtx, client) }()

		params := url.Values{}
		if page := req.GetFloat("page", 0); page > 0 {
			params.Set("page", strconv.Itoa(int(page)))
		}
		b, err := apiutil.RawGet(callCtx, client, "/discover/movies", params)
		if err != nil {
			return apiToolError("DiscoverMoviesGet failed", err)
		}
		enriched, err := EnrichResultsPage(b, <-genresCh)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(enriched)), nil
	}
}

func SearchDiscoverTVHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))

		// Fetch TV genres concurrently while the discover request is in flight.
		genresCh := make(chan GenreMap, 1)
		go func() { genresCh <- FetchTVGenres(callCtx, client) }()

		params := url.Values{}
		if page := req.GetFloat("page", 0); page > 0 {
			params.Set("page", strconv.Itoa(int(page)))
		}
		b, err := apiutil.RawGet(callCtx, client, "/discover/tv", params)
		if err != nil {
			return apiToolError("DiscoverTvGet failed", err)
		}
		enriched, err := EnrichResultsPage(b, <-genresCh)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(enriched)), nil
	}
}

func SearchTrendingHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))

		// Fetch all genres concurrently while the trending request is in flight.
		genresCh := make(chan GenreMap, 1)
		go func() { genresCh <- FetchAllGenres(callCtx, client) }()

		params := url.Values{}
		if page := req.GetFloat("page", 0); page > 0 {
			params.Set("page", strconv.Itoa(int(page)))
		}
		b, err := apiutil.RawGet(callCtx, client, "/discover/trending", params)
		if err != nil {
			return apiToolError("DiscoverTrendingGet failed", err)
		}
		enriched, err := EnrichResultsPage(b, <-genresCh)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(enriched)), nil
	}
}
