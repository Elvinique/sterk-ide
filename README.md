# Sterk IDE

**An original, local-first AI-assisted engineering workspace.** Sterk IDE combines a Monaco code editor, a Go workspace service, policy-limited verification, Git inspection, tracked engineering goals, and bring-your-own-model assistance.

> **Current delivery: functional initial vertical slice (v0.1), not the entire commercial IDE in the master specification.** The app does not yet contain autonomous code-writing agents, production sandboxing, LSP/DAP, integrated PTY terminals, multi-model orchestration, streaming responses, or Playwright browser verification. Read [LIMITATIONS.md](docs/LIMITATIONS.md) before deployment.

## Architecture

```text
Tauri 2 desktop shell (optional) / Vite development UI
  └─ React / TypeScript / Monaco
      └─ authenticated loopback HTTP API (127.0.0.1:4187)
          └─ Go 1.23 daemon
              ├─ workspace reads/writes (with traversal + symlink checks)
              ├─ read-only Git status and diff
              ├─ allowlisted Go verification commands
              ├─ persistent local goals
              └─ OpenAI-compatible model request adapter
```

## Requirements

- Go **1.23+**
- Node.js **20+**
- pnpm **9+**, internet access to install packages
- Optional desktop build: Rust toolchain and [Tauri 2 platform prerequisites](https://v2.tauri.app/start/prerequisites/)

## Start in development mode

Open two terminals from the repository root:

**Terminal 1 — Go daemon**

```bash
export STERK_WORKSPACE="$HOME/my-project" # replace with an existing repository
export STERK_LOCAL_TOKEN="$(openssl rand -hex 32)"
# Print/copy the token to enter in the desktop UI.
echo "$STERK_LOCAL_TOKEN"
cd services/daemon && go run .
```

**Terminal 2 — frontend**

```bash
corepack enable
corepack prepare pnpm@9.15.9 --activate
pnpm install
pnpm dev
```

Open http://127.0.0.1:5173 and enter the daemon token. If the daemon uses `STERK_WORKSPACE` pointing to a Go repository, use the Verification panel to run the project's Go tests or Go vet. Git commands work only inside a Git repository.

## Model configuration

The daemon sends Chat Completions requests to a configured OpenAI-compatible endpoint; only endpoints with compatible request/response formats are supported in this slice. Set environment variables **on the daemon**, not in frontend code:

```bash
export STERK_AI_BASE_URL="https://api.openai.com/v1"
export STERK_AI_KEY="YOUR_KEY"
export STERK_AI_MODEL="YOUR_SUPPORTED_MODEL"
```

In the Settings view, enter the same model ID. This also permits compatible local gateways (including some Ollama setups) if the endpoint implements `/chat/completions`. There is **no native Anthropic Messages or Gemini API adapter yet**. The local token is kept in `sessionStorage` during the browser session. The AI key remains only in the daemon process environment and is never returned by API endpoints.

## Workflows

- **Files:** traverse allowed workspace folders, open/edit/save UTF-8 source files (2 MiB limit); compare unsaved work against the last loaded/saved version.
- **Source Control:** inspect status and unstaged diff (no uncontrolled commit/push).
- **Goal Mode:** enter objective and criteria and persist planned goals in the local user config directory.
- **Verification:** run a small, hard-coded command allowlist (`go test ./...`, `go vet ./...`, `git status --short`) with a 45-second timeout.
- **AI assistant:** send a prompt and current file excerpt to one configured OpenAI-compatible endpoint; AI responses are advisory and cannot execute commands or apply edits.

## Testing

The Go test suite has passed in the build environment; the frontend build has **not** been verified because the package registry was unreachable. See [BUILD-VERIFICATION.md](docs/BUILD-VERIFICATION.md).

```bash
cd services/daemon && go test ./... && go vet ./...
# With packages installed:
pnpm typecheck
pnpm build
pnpm test
```

## Desktop build

```bash
pnpm --filter @sterk/desktop desktop:dev
```

Tauri shell assets require a valid app icon and platform dependencies before bundling/installers are enabled. The Go daemon currently launches independently; it is not bundled as a Tauri sidecar.

## Repo layout

- `apps/desktop` — React/Vite/Monaco UI and Tauri configuration
- `services/daemon` — Go local workspace API and tests
- `infrastructure/supabase/migrations` — optional foundation cloud tables with RLS
- `docs` — architecture, security, limitations, and roadmap

## License

MIT. See `LICENSE`. The project name is for the author's application and is not an endorsement by any competing IDE vendor.
