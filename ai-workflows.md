# AI Workflows & Repository Management

This document defines the standard operating procedures for managing the ZenSec repository, particularly when utilizing AI coding assistants. These workflows ensure consistency, traceability, and maintainability.

## 1. Issue Tracking
*   **Creating Issues:** Every new feature, bug fix, or documentation update must begin with a documented Issue.
*   **Issue Formatting:**
    *   **Prefixes:** Use `[Feature]`, `[Bug]`, `[Chore]`, or `[Docs]`.
    *   **Description:** Clearly state the goal, expected behavior, and acceptance criteria.
*   **AI Context:** When asking an AI to implement a change, provide the Issue number or link so the AI understands the broader context and can reference it in commit messages.

## 2. Branching Strategy
We follow a simplified Git Flow approach to keep the repository clean:
*   `main` branch: Stable, production-ready code.
*   `develop` branch: Integration branch for ongoing development.
*   **Feature Branches:** 
    *   Format: `feature/<issue-number>-<short-description>` (e.g., `feature/12-argon2id-impl`).
    *   Created from `develop`.
*   **Bugfix Branches:**
    *   Format: `bugfix/<issue-number>-<short-description>`.
    *   Created from `develop` (or `main` if it's a critical hotfix).

## 3. Commit Messages
*   We strongly enforce **Conventional Commits**.
*   Format: `<type>(<scope>): <subject>`
*   **Types:** `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`.
*   **Example:** `feat(crypto): implement chunked AES-GCM encryption (#12)`
*   **AI Instruction:** AI agents must generate commit messages adhering strictly to this format to allow for automated changelog generation.

## 4. Pull Requests & Code Review
*   **PR Creation:** PRs should merge feature or bugfix branches into `develop`.
*   **Title:** Match the conventional commit format (e.g., `feat: add TUI file browser`).
*   **Review Process:** 
    *   AI agents can be used to perform initial code reviews (e.g., checking for edge cases, security flaws, or linting errors).
    *   *Critical:* Cryptographic changes (`internal/crypto`) must undergo explicit human review before merging.

## 5. Tagging & Releases
*   **Semantic Versioning:** Use SemVer (`vMAJOR.MINOR.PATCH`).
    *   `MAJOR`: Incompatible API or CLI changes.
    *   `MINOR`: Backwards-compatible new functionality.
    *   `PATCH`: Backwards-compatible bug fixes.
*   **Tags:** Create a Git tag for every release (e.g., `v1.0.0`). AI agents should be instructed to draft release notes when a tag is created.
*   **Changelog:** An auto-generated or manually maintained `CHANGELOG.md` must be updated prior to tagging a release, summarizing all conventional commits since the last tag.
