# Sterk IDE Architecture

## Current implementation

- **Experience plane:** React, TypeScript, Monaco, workspace explorer, code tabs, diff, Git panel, Goals panel, verification output, and AI chat.
- **Execution plane:** single-user Go daemon on loopback, scoped to `STERK_WORKSPACE`. Verification has fixed command allowlist and deadline.
- **Data plane:** source code in the existing filesystem/Git; goal metadata under the OS user config directory; optional Supabase schema intended for future cloud control plane.
- **AI plane:** single OpenAI-compatible Chat Completions adapter; key supplied only to Go. No autonomous tool execution.

## Target service boundaries

The next architecture expands into Go packages with interfaces for: GoalEngine, AgentOrchestrator, ModelGateway, RepositoryBrain, PolicyEngine, SecretsBroker, Sandbox, VerificationEngine, and typed event streams. Current modules are intentionally small; do not assume an empty module placeholder is a feature implementation.

### Goal lifecycle

Proposed: Draft → Planned → Authorized → Executing → Verifying → Human Review → Complete, with failure/cancel states. Current implementation creates **planned** goals only, so no goal is misleadingly marked complete.

### Trust boundaries

Repository content, HTTP inputs, provider responses, and future MCP servers are untrusted. AI suggestions are displayed but cannot issue workspace commands. The Go daemon has host permissions; therefore trusted-only verification is essential pending proper isolation.
