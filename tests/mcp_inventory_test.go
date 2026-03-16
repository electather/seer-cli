package tests

import (
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"
)

func TestMCPInventoryConstants(t *testing.T) {
	// Verifies that the constants in inventory.go match the expected counts.
	// When adding a new tool/resource/prompt, update inventory.go too.
	if cmdmcp.MCPToolCount != 52 {
		t.Errorf("MCPToolCount = %d, want 52", cmdmcp.MCPToolCount)
	}
	if cmdmcp.MCPResourceCount != 9 {
		t.Errorf("MCPResourceCount = %d, want 9", cmdmcp.MCPResourceCount)
	}
	if cmdmcp.MCPPromptCount != 6 {
		t.Errorf("MCPPromptCount = %d, want 6", cmdmcp.MCPPromptCount)
	}
}
