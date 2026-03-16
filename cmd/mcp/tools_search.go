package mcp

import (
	"context"
	"encoding/json"
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
			mcp.WithString("genre", mcp.Description("Genre ID (use seerr://genres/movies resource for IDs)")),
			mcp.WithString("studio", mcp.Description("Studio/production company ID")),
			mcp.WithString("language", mcp.Description("ISO 639-1 language code, e.g. 'en'")),
			mcp.WithString("sortBy", mcp.Description("Sort order, e.g. 'popularity.desc', 'vote_average.desc'")),
			mcp.WithString("primaryReleaseDateGte", mcp.Description("Minimum release date (YYYY-MM-DD)")),
			mcp.WithString("primaryReleaseDateLte", mcp.Description("Maximum release date (YYYY-MM-DD)")),
			mcp.WithNumber("voteAverageGte", mcp.Description("Minimum vote average (0-10)")),
			mcp.WithNumber("voteAverageLte", mcp.Description("Maximum vote average (0-10)")),
			mcp.WithNumber("withRuntimeGte", mcp.Description("Minimum runtime in minutes")),
			mcp.WithNumber("withRuntimeLte", mcp.Description("Maximum runtime in minutes")),
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
			mcp.WithString("genre", mcp.Description("Genre ID (use seerr://genres/tv resource for IDs)")),
			mcp.WithString("network", mcp.Description("Network ID")),
			mcp.WithString("language", mcp.Description("ISO 639-1 language code, e.g. 'en'")),
			mcp.WithString("sortBy", mcp.Description("Sort order, e.g. 'popularity.desc', 'vote_average.desc'")),
			mcp.WithString("firstAirDateGte", mcp.Description("Minimum first air date (YYYY-MM-DD)")),
			mcp.WithString("firstAirDateLte", mcp.Description("Maximum first air date (YYYY-MM-DD)")),
			mcp.WithNumber("voteAverageGte", mcp.Description("Minimum vote average (0-10)")),
			mcp.WithNumber("voteAverageLte", mcp.Description("Maximum vote average (0-10)")),
			mcp.WithNumber("withRuntimeGte", mcp.Description("Minimum runtime in minutes")),
			mcp.WithNumber("withRuntimeLte", mcp.Description("Maximum runtime in minutes")),
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

	s.AddTool(
		mcp.NewTool("search_company",
			mcp.WithDescription("Search for production companies"),
			mcp.WithOpenWorldHintAnnotation(true),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithString("query", mcp.Required(), mcp.Description("Search query")),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchCompanyHandler(),
	)

	s.AddTool(
		mcp.NewTool("search_keyword",
			mcp.WithDescription("Search for TMDB keywords"),
			mcp.WithOpenWorldHintAnnotation(true),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithString("query", mcp.Required(), mcp.Description("Search query")),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchKeywordHandler(),
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
		for _, key := range []string{"genre", "studio", "language", "sortBy", "primaryReleaseDateGte", "primaryReleaseDateLte"} {
			if v := req.GetString(key, ""); v != "" {
				params.Set(key, v)
			}
		}
		for _, key := range []string{"voteAverageGte", "voteAverageLte", "withRuntimeGte", "withRuntimeLte"} {
			if v := req.GetFloat(key, 0); v != 0 {
				params.Set(key, strconv.FormatFloat(v, 'f', -1, 64))
			}
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
		for _, key := range []string{"genre", "network", "language", "sortBy", "firstAirDateGte", "firstAirDateLte"} {
			if v := req.GetString(key, ""); v != "" {
				params.Set(key, v)
			}
		}
		for _, key := range []string{"voteAverageGte", "voteAverageLte", "withRuntimeGte", "withRuntimeLte"} {
			if v := req.GetFloat(key, 0); v != 0 {
				params.Set(key, strconv.FormatFloat(v, 'f', -1, 64))
			}
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

func SearchCompanyHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query := req.GetString("query", "")
		if query == "" {
			return nil, fmt.Errorf("query is required")
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.SearchAPI.SearchCompanyGet(callCtx).Query(query)
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("SearchCompanyGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func SearchKeywordHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query := req.GetString("query", "")
		if query == "" {
			return nil, fmt.Errorf("query is required")
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.SearchAPI.SearchKeywordGet(callCtx).Query(query)
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("SearchKeywordGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
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
