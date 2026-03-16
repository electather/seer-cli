package mcp

// Inventory counts must match the registrations in registerXxx functions.
// serve.go references these so log lines stay accurate. Tests assert them
// to catch drift when tools are added or removed.
const (
	MCPToolCount     = 52
	MCPResourceCount = 9
	MCPPromptCount   = 6
)
