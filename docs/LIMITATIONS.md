# Verified scope and outstanding work

This build is a **real initial product slice**, not a complete implementation of the 54-section master requirements. Do not market it as a fully autonomous or secure cloud IDE.

Implemented: React Monaco editor and diff display; local file reads/saves; folder navigation; Git status/diff; planned-goal creation and JSON persistence; allowlisted build checks; single OpenAI-compatible chat-completions adapter; local token-protected loopback daemon; workspace traversal/symlink checks; UI shell and Tauri source project.

Not implemented: native OS secret keychain, encryption at rest, complete secrets broker, agent/tool calling, policy levels 0–5, sandbox with Docker or VM boundary, PTY terminal, streaming AI, full provider abstractions, model routing and provider settings, LSP, DAP, Tree-sitter, MCP, ACP, repository embeddings, pgvector synchronization, agent worktrees, UI accessibility audits, Playwright, multi-agent planning, application preview, deploys, Git writes, full Supabase organization workflows, subscription controls, hosted execution, browser-based tests, GitHub pull requests, and production packaging.

Security caution: the daemon token and cross-origin restriction protect an MVP loopback service; they are **not** a substitute for robust origin-independent CSRF defenses, OS IPC identity validation, sandboxing, authenticated multi-user API deployment, or authorization. Never listen on public network interfaces with this version. The current daemon does not run arbitrary agent-provided commands. UI text files should be treated as potentially untrusted. There is no sandbox for running `go test` or `go vet`; invoke these only for projects you trust. Do not use as a hosted multi-tenant server.

Go client exposes a single configurable provider endpoint; avoid untrusted `STERK_AI_BASE_URL` values. No actual browser or Rust/Tauri packaging tests were performed unless individually reported in verification output.
