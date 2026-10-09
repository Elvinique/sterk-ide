# Build verification — 2026-10-09

## Passed

- `go test -v ./...` from `services/daemon` — 3 tests passed: path confinement, local token authentication, and verification command allowlist.
- `go vet ./...` from `services/daemon` — passed.
- Go source formatted using `gofmt`.

## Not verified in this environment

- pnpm dependency installation — Corepack could not reach `registry.npmjs.org` (`EAI_AGAIN`).
- TypeScript typecheck, Vite frontend build, UI end-to-end tests — blocked by missing pnpm packages.
- Tauri desktop build — Rust toolchain and platform prerequisites are unavailable here; bundling is disabled.
- Live provider authentication, Supabase schema deployment, and commercial AI features — require your credentials and external environment.

Do not label the release production-tested until all missing gates pass in CI or on a correctly provisioned development machine.
