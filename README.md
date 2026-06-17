# ZenSec - The Minimalist Local File Encryption Utility 🔐

[![Go Report Card](https://goreportcard.com/badge/github.com/githubuser2777/ZenSec)](https://goreportcard.com/report/github.com/githubuser2777/ZenSec)
[![Go Version](https://img.shields.io/github/go-mod/go-version/githubuser2777/ZenSec)](https://golang.org/doc/go1.20)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://github.com/githubuser2777/ZenSec/actions/workflows/build.yml/badge.svg)](https://github.com/githubuser2777/ZenSec/actions)

A blazing-fast, strictly secure, zero-dependency command-line utility built to encrypt and decrypt local files. Written entirely in Go, ZenSec follows the UNIX philosophy: do one thing and do it exceptionally well. It leverages Go's powerful standard cryptography libraries to ensure data privacy and integrity without any unnecessary bloat.

---

## ✨ Features

* **Zero-Dependency CLI:** A straightforward Command-Line Interface built entirely with standard libraries. No bloated UI frameworks.
* **Military-Grade Cryptography:**
  * **AES-256-GCM:** Authenticated streaming encryption prevents data tampering.
  * **Argon2id:** The PHC-winning key derivation function protects against GPU/ASIC brute forcing.
  * **Anti-Tampering Architecture:** Sequence numbers in nonces prevent chunk reordering, while custom AAD flags mathematically prevent truncation attacks.
* **Memory Efficient (Streamed Chunking):** Processes massive files (GBs or TBs) using a tiny 64KB memory footprint. Your RAM will never overflow.
* **Advanced Automation:**
  * **Keyfile Support:** Use any physical file (an image, a PDF, a text file) as your cryptographic key instead of typing a password.
  * **Batch Processing:** Includes native scripts to instantly encrypt/decrypt hundreds of files in a directory.
* **Native OS Integration:** Easily installable Windows Context Menu (Right-click -> Encrypt with ZenSec).

---

## 🚀 Getting Started

### Prerequisites
* [Go 1.20+](https://go.dev/doc/install) (if building from source)

### Installation

Clone the repository and build the single, standalone binary:

```bash
git clone https://github.com/githubuser2777/ZenSec.git
cd ZenSec

# Download dependencies (only golang.org/x/crypto and golang.org/x/term)
go mod tidy

# Build the executable
go build -o zensec.exe cmd/zensec/main.go
```

> **Pro-Tip:** Move `zensec.exe` to a folder in your system's `PATH` (e.g., `C:\Windows\System32` on Windows or `/usr/local/bin` on Linux) so you can run it from anywhere.

---

## 🛠️ Usage Guide

### 1. Standard Encryption & Decryption (Password)

Encrypt a file:
```bash
zensec -encrypt -file my_secret.txt
```
*(You will be securely prompted to type a password, which will be hidden as you type)*

Decrypt a file:
```bash
zensec -decrypt -file my_secret.txt.enc
```

### 2. Keyfile Mode (Passwordless)

Instead of a password, you can use any file as your key. This is incredibly useful for automation or storing your "key" on an external USB drive.

```bash
# Encrypt using an image as the key
zensec -encrypt -file backup.zip -keyfile D:\my_secret_key.jpg

# Decrypt using the same image
zensec -decrypt -file backup.zip.enc -keyfile D:\my_secret_key.jpg
```

### 3. Windows Context Menu (Right-Click)

1. Ensure `zensec.exe` is in your system `PATH`.
2. Double-click the included `install_context_menu.reg` file.
3. You can now right-click any file in Windows Explorer and select **"Encrypt with ZenSec"** or **"Decrypt with ZenSec"**. A secure terminal will pop up asking for your password.

### 4. Batch Processing (Directories)

Need to encrypt a whole folder? Just double-click `batch_zensec.bat`. It will prompt you for the folder path, your keyfile, and automatically process every file inside.

---

## 📚 Documentation

For an in-depth look at the internal architecture, please refer to our documentation files:
* [Features Detail](features.md) - Deep dive into technical capabilities.
* [Security Overview](security.md) - Cryptographic design and threat modeling.
* [Project Overview](docs/overview.md) - Philosophy and structural architecture.
* [Roadmap](Roadmap.md) - Project history and status.

## 📄 License

This project is licensed under the [MIT License](LICENSE).
