# ZenSec Features Detail

This document provides a comprehensive list of features that are currently implemented or planned for ZenSec.

## 🔐 Core Cryptographic Features
*   **AES-256-GCM Authenticated Encryption:** Ensures that data is both securely encrypted and protected against any tampering or corruption.
*   **Argon2id Key Derivation:** Uses the winner of the Password Hashing Competition (PHC) to provide maximum resistance against GPU and ASIC brute-force attacks.
*   **Memory-Efficient Chunking:** Processes files in a streaming fashion (e.g., 64KB chunks), allowing the encryption of massive files (GBs or TBs) without running out of RAM.
*   **Cryptographically Secure Nonces:** Implements sequential nonce generation intertwined with the chunk counter to prevent chunk reordering, substitution, and truncation attacks.

## 💻 Interfaces
*   **Command-Line Interface (CLI):** 
    *   Fast, scriptable interface for system administrators and automated backups.
    *   Secure, hidden password prompts.
*   **Terminal User Interface (TUI):** *(Planned)*
    *   Interactive terminal navigation without needing a windowing system.
    *   Built-in file browser to easily select files for encryption/decryption.
    *   Real-time progress bars for large file processing.
*   **Graphical User Interface (GUI):** *(Planned)*
    *   User-friendly application for everyday use.
    *   Drag-and-drop file support.
    *   Native OS integration (system tray, native dialogs).

## 🚀 Advanced Functionality (Planned)
*   **Keyfile Support:** Use a physical file (e.g., stored on a USB drive) as the encryption key instead of or in addition to a password.
*   **Batch Processing:** Recursively encrypt or decrypt entire directories.
*   **Header Obfuscation:** Option to hide metadata so that the resulting file does not have a recognizable signature (plausible deniability).
*   **OS Context Menu Integration:** Right-click a file in Windows Explorer, macOS Finder, or Linux File Manager and select "Encrypt with ZenSec".
