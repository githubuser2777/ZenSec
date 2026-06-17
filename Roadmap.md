# ZenSec Roadmap

> **Status:** All planned phases are currently COMPLETE. The tool has reached a stable v1.0 state.

## ✅ Phase 1: Core Cryptography & Standard CLI
- [x] Setup Go project structure (`cmd`, `internal`).
- [x] Define secure encryption file format (handling chunked AES-GCM securely to prevent reordering attacks with sequence numbers).
- [x] Implement Argon2id key derivation logic.
- [x] Implement chunked stream encryption and decryption.
- [x] Minimal CLI application with `flag` standard library.
- [x] Comprehensive unit tests for core crypto functions.

## ✅ Phase 2: OS Integration & Utils 
- [x] Document batch processing using native shell scripts (`batch_zensec.bat`).
- [x] Provide OS Context Menu integrations (Windows `.reg` file) for "Right-click -> Encrypt with ZenSec".
- [x] Add advanced Keyfile support (`-keyfile` flag).

## 🔮 Future Considerations (Community Driven)
While the core philosophy of ZenSec is "ponytail minimalism" (YAGNI - You Aren't Gonna Need It), future community forks may explore:
- Built-in multi-threading for processing multiple files simultaneously in Go.
- Integration into macOS Automator / Linux `.desktop` files.
