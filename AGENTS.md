# Seer CLI - Agent Guide

This document serves as a guide for AI agents and developers working on the `seer-cli` project. It describes the project structure, the build process, and the architectural principles to follow when extending the CLI.

## Project Overview

`seer-cli` is a Go-based Command Line Interface application designed to interact with the Seer API. It uses the `cobra` framework for CLI command structuring and `openapi-generator-cli` to automatically generate the API client library directly from an OpenAPI specification file.

## Project Structure

```text
.
├── AGENTS.md               # This guide
├── cmd/                    # Cobra CLI commands containing the logic for the CLI
│   ├── root.go             # Root command definition & initialization
│   └── status.go           # Example API-consuming command
├── generate-api-lib.sh     # Bash script to generate the OpenAPI Go client
├── go.mod                  # Go module definition for the CLI application
├── go.work                 # Go workspace definition to include local pkg/api
├── main.go                 # Application entrypoint
├── open-api.yaml           # Source OpenAPI specification defining the API
└── pkg/
    └── api/                # Auto-generated Go client library (DO NOT EDIT MANUALLY)
```

## Build & Generation Process

### 1. The OpenAPI Client
The `pkg/api` directory contains a completely auto-generated Go client. 

**Updating the API:** When the underlying API changes or `open-api.yaml` is updated, the client library must be regenerated.
**Generation Script:** Run `./generate-api-lib.sh`. This uses a Dockerized `openapitools/openapi-generator-cli` to read `open-api.yaml` and aggressively overwrite the Go module inside `./pkg/api`.

### 2. The Go Workspace
The project uses a `go.work` file. This is crucial because it allows the main `seer-cli` module to depend on the local `seer-cli/pkg/api` sub-module without needing it published to a remote version control repository. 

### 3. Building the CLI
To build the CLI into an executable binary, run:
```bash
go build -o seer-cli main.go
```

## Architecture & Future Commands

### Command Structure
The CLI is built using the [Cobra](https://github.com/spf13/cobra) framework. All new commands must be placed inside the `cmd/` directory.
- **Root Command:** `cmd/root.go` contains the base command and global flags.
- **Top-Level Commands:** Should be placed in their own files (e.g., `cmd/movies.go`, `cmd/users.go`).
- **Subcommands:** For nested commands, use the `<parent>_<child>.go` naming convention and attach them to the parent command in their respective `init()` blocks (e.g., `cmd/users_list.go` attaches to the `usersCmd` defined in `cmd/users.go`).

### Testing
- **Unit Tests are Expected:** When adding complex business logic, parsing rules, or multi-step command flows, write unit tests.
- **Test Files Location:** Unit tests should be placed alongside the command they test, using the standard Go `_test.go` suffix (e.g., `cmd/status_test.go`).
- **Mocking the API:** Do not make live API calls in tests. Either mock the `seer-cli/pkg/api` interfaces or use HTTP testing utilities like `httptest` to stub out server responses.

### Best Practices for Future Agents
When tasked with adding new features or commands to this CLI, adhere to the following rules:

1. **Keep Files Small and Focused:**
   A core architectural principle of this application is maintaining readability. Do not create massive files. Break complex operations, data transformations, and API handling into smaller, testable helper functions. If a command file grows too large, extract the logic into separate dedicated files within the `cmd/` package or a new internal package if appropriate.

2. **Treat `pkg/api/` as strictly read-only and ephemeral:** 
   Never manually edit any files inside `pkg/api/`. If changes are needed in the client-side data models, endpoints, or method signatures, you *must* modify `open-api.yaml` first, and then run `./generate-api-lib.sh`.
   
2. **Follow the Command Pattern:**
   Use the existing `cmd/status.go` as a template for new commands. Create a new file for each distinct top-level resource.

3. **Subcommands for Complex Resources:** 
   For complex API groups (like users, movies, or media), use Cobra subcommands grouping. Organize them predictably:
   - `cmd/users.go` -> parent `users` command (`var usersCmd = &cobra.Command{...}`)
   - `cmd/users_list.go` -> `users list` command attached to `usersCmd`
   - `cmd/users_get.go` -> `users get` command attached to `usersCmd`

4. **API Client Initialization:** 
   Initialize the API client strictly inside the `RunE` method of the Cobra command so that configurations (like global flags for Host URL or API keys) can be evaluated sequentially.
   ```go
   func(cmd *cobra.Command, args []string) error {
       configuration := api.NewConfiguration()
       
       // Example: Update Host based on Viper config or Flags here
       // configuration.Servers[0].URL = "http://localhost:5055/api/v1"
       
       apiClient := api.NewAPIClient(configuration)
       // ...
   }
   ```

5. **Context Implementation:** 
   Always pass a valid context (e.g., `context.Background()` or one tied to the command lifecycle) to the `.Execute()` calls of the generated API.

6. **Output Formatting:** 
   When designing commands that return complex API data, consider standardizing the output structure. Eventually, commands should support a `--json` or `--yaml` flag to dump raw API responses directly to `stdout` for programmatic consumption, instead of just pretty-printing.

## Development Workflow
1. Analyze the user request and locate the relevant endpoints in `open-api.yaml`.
2. Check if the `/pkg/api` generated files accurately represent the desired endpoint. (Regenerate if you edited the OpenAPI spec).
3. Generate the command logic inside a new `cmd/<resource>.go` file.
4. Add the command to the appropriate parent command in its `init()` block.
5. Provide the user with the exact `go build` or `go run` instruction to test their newly added command!
