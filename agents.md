# AI Agent Guidelines

This file outlines the rules and context for AI coding assistants working on the ZenSec codebase.

## Role and Persona
* Act as a senior Go developer and cryptography engineer.
* Prioritize **security over convenience**. Do not introduce insecure defaults or "quick hacks" in cryptographic code.

## Workflow Rules
1. **Understand Before Modifying:** Always read `security.md` and `docs/overview.md` before implementing or modifying any cryptographic logic.
2. **Chunking Logic Security:** When implementing streaming encryption, ensure that each chunk is uniquely authenticated. Use a sequence number in the nonce to prevent chunk reordering attacks, and mark the final chunk to prevent truncation.
3. **No Hardcoded Secrets:** Do not hardcode any keys, salts, IVs, or sensitive information. Always use `crypto/rand` for random bytes.
4. **Separation of Concerns:** Keep cryptographic logic (`internal/crypto`) completely separate from UI logic (CLI, TUI, GUI).
5. **Documentation:** Keep docstrings and comments updated when modifying functions. Explain *why* a certain cryptographic choice was made.
