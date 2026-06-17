# Project Rules & Coding Standards

## General Guidelines
*   **Language:** Go (1.21+ recommended).
*   **Formatting:** Always run `go fmt` before committing code.
*   **Linting:** Use `golangci-lint` to ensure code quality and consistency.

## Code Structure
*   `cmd/zensec/`: Contains the main application entry points for the CLI and TUI.
*   `cmd/zensec-gui/`: Contains the main entry point for the GUI application.
*   `internal/crypto/`: Core cryptographic logic. Must not depend on any UI packages.
*   `internal/ui/`: UI components (TUI and GUI code).

## Pull Requests & Contributions
*   All PRs must include relevant unit tests, especially for components in `internal/crypto`.
*   Cryptographic functions must have close to 100% test coverage.
*   PR descriptions should clearly state the problem being solved and the approach taken.
*   Do not merge code that fails linting or security checks.
