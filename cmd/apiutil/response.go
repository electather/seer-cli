package apiutil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// HandleResponse marshals res to indented JSON and prints it.
// In verbose mode it also prints the request URL and HTTP status.
func HandleResponse(cmd *cobra.Command, r *http.Response, err error, res interface{}, isVerbose bool, method string) error {
	if isVerbose && r != nil {
		cmd.Printf("Request URL: %s %s\n", r.Request.Method, r.Request.URL.String())
	}
	if err != nil {
		if isVerbose && r != nil {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			return fmt.Errorf("error when calling %s: %w\nFull HTTP response: %v", method, err, r)
		}
		return fmt.Errorf("error when calling %s: %w", method, err)
	}
	jsonRes, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}
	if isVerbose {
		if r != nil {
			cmd.Printf("HTTP Status: %s\n", r.Status)
		}
		cmd.Printf("Response from %s:\n%s\n", method, string(jsonRes))
	} else {
		cmd.Println(string(jsonRes))
	}
	return nil
}

// HandleRawResponse reads the response body and prints it as-is.
func HandleRawResponse(cmd *cobra.Command, r *http.Response, err error, isVerbose bool, method string) error {
	if isVerbose && r != nil {
		cmd.Printf("Request URL: %s %s\n", r.Request.Method, r.Request.URL.String())
	}
	if err != nil {
		if isVerbose && r != nil {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			return fmt.Errorf("error when calling %s: %w\nFull HTTP response: %v", method, err, r)
		}
		return fmt.Errorf("error when calling %s: %w", method, err)
	}
	if isVerbose && r != nil {
		cmd.Printf("HTTP Status: %s\n", r.Status)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	defer r.Body.Close()
	cmd.Println(string(body))
	return nil
}

// Handle204Response handles responses with no body (HTTP 204).
// It prints {"status":"ok"} on success.
func Handle204Response(cmd *cobra.Command, r *http.Response, err error, verbose bool, method string) error {
	if err != nil {
		if verbose && r != nil {
			return fmt.Errorf("error when calling %s: %w\nFull HTTP response: %v", method, err, r)
		}
		return fmt.Errorf("error when calling %s: %w", method, err)
	}
	if verbose {
		cmd.Printf("HTTP Status: %s\n", r.Status)
	}
	cmd.Println(`{"status": "ok"}`)
	return nil
}
