# ZenSec Roadmap

## Phase 1: Core Cryptography & CLI (Current)
- [ ] Setup Go project structure (`cmd`, `internal`).
- [ ] Define secure encryption file format (handling chunked AES-GCM securely to prevent reordering attacks).
- [ ] Implement Argon2id key derivation logic.
- [ ] Implement chunked stream encryption and decryption.
- [ ] Basic CLI application with `-encrypt` and `-decrypt` flags.
- [ ] Comprehensive unit tests for core crypto functions.

## Phase 2: Terminal User Interface (TUI)
- [ ] Integrate a TUI framework (e.g., Charmbracelet's Bubbletea).
- [ ] Create a file browser component to select files for encryption/decryption.
- [ ] Implement secure password prompt UI.
- [ ] Add progress bars for large file processing.

## Phase 3: Graphical User Interface (GUI)
- [ ] Evaluate and select GUI framework (Wails or Fyne).
- [ ] Design accessible interface with drag-and-drop support.
- [ ] Build desktop installers for Windows, macOS, and Linux.

## Phase 4: Advanced Features
- [ ] Keyfile support (encrypt/decrypt using a file instead of a password).
- [ ] Batch processing (encrypt multiple files or directories).
- [ ] Integration with OS context menus (e.g., "Right-click -> Encrypt with ZenSec").
