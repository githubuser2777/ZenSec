# ZenSec - Local File Encryption Utility 🔐

A blazing-fast, secure command-line, terminal, and graphical utility built to encrypt and decrypt local files. Written entirely in Go, this project provides multiple interfaces (CLI, TUI, GUI) to suit your workflow. It leverages Go's powerful cryptography libraries to ensure data privacy and integrity.

## Features

* **Multiple Interfaces:** Includes a fast Command-Line Interface (CLI), an interactive Terminal UI (TUI), and a user-friendly Graphical User Interface (GUI).
* **Single Binary Execution:** Compiles to a dependency-free executable (for CLI/TUI), making it incredibly easy to run on any system.
* **AES-GCM Encryption:** Uses the Advanced Encryption Standard with Galois/Counter Mode for authenticated encryption.
* **Secure Key Derivation:** Implements robust key derivation (Argon2id) to safely convert human-readable passwords into mathematically secure 32-byte cryptographic keys.
* **Memory Efficient:** Processes files in chunked byte streams rather than loading the entire file into memory, allowing you to encrypt massive files without crashing your system.

## Documentation

For more detailed information, please refer to our documentation files:
* [Features Detail](features.md) - Comprehensive list of core and planned features.
* [Roadmap](Roadmap.md) - Project timeline and planned features.
* [Security Overview](security.md) - Cryptographic design and security policies.
* [Agent Rules](agents.md) - Guidelines for AI agents working on this repository.
* [Project Rules](rules.md) - Coding standards and contribution guidelines.
* [AI Workflows](ai-workflows.md) - Repository management (tags, issues, branches).
* [Project Overview](docs/overview.md) - In-depth project description and architecture.

## Tech Stack

* **Language:** Go (Golang)
* **Core Packages:** `crypto/aes`, `crypto/cipher`, `crypto/rand`
* **Key Derivation:** `golang.org/x/crypto/argon2`
* **TUI Framework:** (Planned) e.g., Bubbletea
* **GUI Framework:** (Planned) e.g., Wails or Fyne

## Getting Started

### Prerequisites

Ensure you have [Go](https://go.dev/doc/install) installed on your system.

### Installation & Build

Clone the repository to your local machine, initialize the module, and build the binary:

```bash
# Clone the repository
git clone https://github.com/githubuser2777/ZenSec.git
cd ZenSec

# Download necessary supplemental crypto packages
go mod tidy

# Compile the CLI/TUI tool
go build -o zensec main.go
```

## Usage

*(Currently CLI is the primary focus. TUI and GUI usage will be updated in the future).*

### To Encrypt a file (CLI):

```bash
./zensec -mode=encrypt -file=my_secret_notes.txt
```

### To Decrypt a file (CLI):

```bash
./zensec -mode=decrypt -file=my_secret_notes.txt.enc
```

## License

This project is licensed under the [MIT License](LICENSE).
