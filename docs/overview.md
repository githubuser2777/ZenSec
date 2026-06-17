# ZenSec Project Overview

## What is ZenSec?
ZenSec is a modern, cross-platform file encryption utility designed to bring enterprise-grade security to individual users. Its core philosophy revolves around making strong cryptography accessible, fast, and minimal. We prioritize the UNIX philosophy: do one thing and do it well, avoiding unnecessary abstractions like complex TUIs or GUIs.

## Core Objectives
1.  **Security First:** Use robust, proven cryptographic primitives (AES-256-GCM, Argon2id) from the Go standard library.
2.  **Performance:** Efficiently handle very large files (e.g., video files, backups) using stream processing to prevent out-of-memory errors.
3.  **Simplicity:** Provide a minimal, zero-dependency CLI interface that seamlessly integrates with existing OS tools and workflows.

## Architecture Highlights
*   **Crypto Core (`internal/crypto`):** An independent module that handles all encryption/decryption logic securely.
*   **CLI Layer (`cmd/zensec`):** Shell-based operations built exclusively with the standard library.
*   **OS Integration Layer:** Native OS scripts and registry modifications to bridge the gap for users who prefer visual interaction (e.g., right-click context menus).
