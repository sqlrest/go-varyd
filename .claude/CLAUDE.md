# template.library

The layout oracle for gomatic **library** repositories: a minimal pure-Go library that every library repo's structure is verified against (`tools.repository` spec 007). Instantiate it for new libraries; keep it perfectly conformant.

- **`library.go` is the class marker** — a build-tagged, never-compiled file whose presence classifies the repo as a library. Do not remove it; adjust only its `package` clause.
- **README is exactly workflow badges** (plus the docs link in instantiated repos — templates themselves have no docs repo). No prose; documentation lives in the repo's `docs.<repo>` sibling.
- Depends only on the stdlib (testify for tests). A library never imports a CLI framework.
- Gate: `make check` — gofumpt, vet, staticcheck, govulncheck, gocognit ≤ 7, exactly 100.0% coverage. Managed files (Makefile, .golangci.yaml, workflows) are distributed by `nicerobot/tools.repository` — never hand-edit them.
