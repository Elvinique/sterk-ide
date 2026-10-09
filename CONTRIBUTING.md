# Contributing

Create a short issue explaining the behavior to change. Branch from main; keep Go workspace/security logic outside React components. Add tests for filesystem policy and new API endpoints. Run `go test ./...`, `go vet ./...` from `services/daemon`, and `pnpm typecheck && pnpm build` from the repository root. Do not commit `.env.local`, private keys, or generated artifacts. New agent features require explicit policy evaluation and auditable evidence before a goal is completed.
