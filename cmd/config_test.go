package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestConfigCommands(t *testing.T) {
	// Create a temporary directory for the config file
	tempDir, err := os.MkdirTemp("", "seer-cli-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, ".seer-cli.yaml")

	t.Run("config help", func(t *testing.T) {
		viper.Reset()
		b := new(bytes.Buffer)
		rootCmd.SetOut(b)
		rootCmd.SetErr(b)
		rootCmd.SetArgs([]string{"config", "--help"})

		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		out := b.String()
		if !strings.Contains(out, "View or update the configuration for the Seer CLI") {
			t.Errorf("expected help output to contain command description, got: %q", out)
		}
	})

	t.Run("config set", func(t *testing.T) {
		viper.Reset()
		b := new(bytes.Buffer)
		rootCmd.SetOut(b)
		rootCmd.SetErr(b)
		
		testServer := "http://test-server:5055"
		testKey := "test-api-key-12345"
		
		rootCmd.SetArgs([]string{"config", "set", "--server", testServer, "--api-key", testKey, "--config", configPath})

		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		out := b.String()
		if !strings.Contains(out, "Configuration saved successfully") {
			t.Errorf("expected success message, got: %s", out)
		}

		// Verify file contents
		data, err := os.ReadFile(configPath)
		if err != nil {
			t.Fatalf("failed to read config file: %v", err)
		}
		
		content := string(data)
		if !strings.Contains(content, testServer) || !strings.Contains(content, testKey) {
			t.Errorf("config file does not contain expected values: %s", content)
		}
	})

	t.Run("config show", func(t *testing.T) {
		// Use the file created in the previous step
		viper.Reset()
		b := new(bytes.Buffer)
		rootCmd.SetOut(b)
		rootCmd.SetErr(b)
		
		rootCmd.SetArgs([]string{"config", "show", "--config", configPath})

		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		out := b.String()
		if !strings.Contains(out, "http://test-server:5055") {
			t.Errorf("expected output to contain server URL, got: %s", out)
		}
		// Check for masked API key
		if !strings.Contains(out, "test****2345") {
			t.Errorf("expected output to contain masked API key, got: %s", out)
		}
	})

	t.Run("config show empty", func(t *testing.T) {
		viper.Reset()
		emptyConfig := filepath.Join(tempDir, "empty.yaml")
		
		b := new(bytes.Buffer)
		rootCmd.SetOut(b)
		rootCmd.SetErr(b)
		
		rootCmd.SetArgs([]string{"config", "show", "--config", emptyConfig})

		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		out := b.String()
		if !strings.Contains(out, "API Key:     <not set>") {
			t.Errorf("expected output to show API key as not set, got: %s", out)
		}
	})
}
