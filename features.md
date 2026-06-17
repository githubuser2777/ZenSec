# ZenSec Features Detail

## 🔐 Core Cryptographic Primitives
* **AES-256-GCM (Galois/Counter Mode):** The gold standard for authenticated encryption. Ensures data confidentiality and guarantees that any tampering will cause decryption to fail immediately.
* **Argon2id Key Derivation:** Uses the winner of the Password Hashing Competition (PHC) to derive a mathematically secure 32-byte key from human-readable passwords or raw keyfiles. Configured to resist GPU and ASIC brute-force attacks.

## 🛡️ Attack Mitigations
* **Memory-Efficient Chunking (64KB):** Files are processed in 64KB blocks. This prevents Out-Of-Memory (OOM) crashes when processing 50GB+ files on machines with limited RAM.
* **Chunk Reordering Protection:** Each chunk's nonce and Additional Authenticated Data (AAD) is mathematically bound to a strictly incrementing sequence number. Attackers cannot swap chunk 1 with chunk 5 without breaking the authentication tag.
* **Truncation Protection:** The final chunk is flagged with a special `isLast` byte inside the AAD. If an attacker deletes the final chunk of an encrypted file to cut off data, the decryption process will realize the `isLast` flag is missing from the *new* final chunk and will halt with a tampering error.

## 💻 Interface & Automation
* **Zero-Dependency CLI:** Built entirely on the Go standard library `flag` package to keep the compiled binary footprint under 5MB.
* **Secure Prompts:** Uses `golang.org/x/term` to prevent passwords from echoing to the terminal screen during manual entry.
* **Keyfile Authentication:** Allows substituting passwords with raw files. Crucial for scripting, automated backups, or high-security physical token workflows (e.g., storing the keyfile on a USB).
* **OS Integrations:** 
  * Windows Context Menu integration via `.reg` files.
  * Multi-file Batch Processing via `batch_zensec.bat`.
