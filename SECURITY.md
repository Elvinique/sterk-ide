# Security policy and threat model

Report security issues privately to the repository maintainers. Do not include API keys or exploit data in public reports.

## Trust boundaries

- The daemon binds **only** to `127.0.0.1` and requires `X-Sterk-Token` on API calls.
- Browser origins other than localhost Vite and Tauri origins are rejected, but this is an MVP defense and not a comprehensive browser/IPC security design.
- Paths are lexically confined to the selected workspace; symlinks are rejected along traversed paths. Race conditions against filesystem mutation remain possible and require hardened fd-relative APIs in future releases.
- `go test` and `go vet` can execute build scripts or tests supplied by a repository. **Only use these with trusted repositories.** There is no container isolation in v0.1.
- Provider keys are environment variables in the Go process, not bundled into React; no OS keychain integration yet. Avoid putting keys in repository files or logs.
- Agent suggestions never automatically apply changes. No arbitrary commands or file deletion API exists.
- Cloud database policies in the migration restrict projects and goals by authenticated owner. Supabase is optional and not connected to current local app.

## Remaining blockers for production security

Implement OS keychain, CSRF/IPC hardening, symlink race elimination, process sandbox, egress controls, audit trail, provider-specific adapters, provenance for changes, and independent penetration testing before distributing as a secure autonomous IDE.
