# ZenSec Project Overview

## What is ZenSec?
ZenSec is a modern, cross-platform file encryption utility designed to bring enterprise-grade security to individual users. Its core philosophy revolves around making strong cryptography accessible and easy to use without compromising on security.

## Core Objectives
1.  **Security First:** Use robust, proven cryptographic primitives (AES-256-GCM, Argon2id).
2.  **Performance:** Efficiently handle very large files (e.g., video files, backups) using stream processing to prevent out-of-memory errors.
3.  **Accessibility:** Provide multiple interfaces:
    *   **CLI** for server environments and automation.
    *   **TUI** for terminal power users who want an interactive experience.
    *   **GUI** for general users who prefer visual, drag-and-drop workflows.

## Architecture Highlights
*   **Crypto Core (`internal/crypto`):** An independent module that handles all encryption/decryption logic. It is entirely decoupled from the UI.
*   **Interfaces:**
    *   **CLI Layer (`cmd/zensec`):** Shell-based operations.
    *   **TUI Layer (`internal/ui/tui`):** Terminal interface, likely utilizing Bubbletea.
    *   **GUI Layer (`cmd/zensec-gui`):** Windowed application, likely utilizing Wails or Fyne.
