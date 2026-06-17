# Security Policy and Design

## Cryptographic Primitives
*   **Encryption Algorithm:** AES-256 in GCM mode (Galois/Counter Mode) for Authenticated Encryption with Associated Data (AEAD).
*   **Key Derivation:** Argon2id (using recommended memory and time costs).
*   **Salt:** 16-byte cryptographically secure random salt generated per file.
*   **Random Number Generation:** Must exclusively use `crypto/rand`.

## Chunked Encryption Security
To support large files memory-efficiently, ZenSec encrypts files in chunks (e.g., 64KB blocks). Standard AES-GCM only authenticates the single block it encrypts. 

**To prevent chunk substitution, truncation, or reordering attacks:**
1.  **Sequential Nonces:** The `Nonce` (12 bytes for AES-GCM) must incorporate a sequential chunk counter.
2.  **Last Chunk Identifier:** The final chunk must be cryptographically distinguishable from intermediate chunks (e.g., by flipping a specific bit in the nonce or associated data) to prevent attackers from silently dropping the end of the file.

## Vulnerability Reporting
If you discover a security vulnerability in ZenSec, please do **NOT** open a public issue. 
Please reach out to the project maintainers directly or use the GitHub Security Advisory feature to privately report the issue.
