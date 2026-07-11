# go-varyd

Go bindings for the varyd wire API: hand-authored protobuf definitions under [api/proto](../api/proto/), buf-generated gRPC contract committed under [src/proto](../src/proto/), shared by the [vaql](https://github.com/sqlrest/vaql) client and the [varyd](https://github.com/sqlrest/varyd) server. Extracted from `sqlrest/varyd` (spec: `sqlrest/_projects/specs/varyd/002-repo-split`).

- **`library.go` is the class marker** — a build-tagged, never-compiled file whose presence classifies the repo as a library. Do not remove it.
- **Never hand-edit `src/proto`** — change the `.proto` under `api/proto`, run `make proto` (ad-hoc buf via [scripts/proto-gen.sh](../scripts/proto-gen.sh); requires `buf`, `protoc-gen-go`, `protoc-gen-go-grpc` on PATH/GOBIN — never a `go.mod` tool stanza), and commit the regenerated tree. The root-package `TestServicesContract` pins the registered `varyd.api.v1` service set.
- **Additive changes only** within `varyd.api.v1`; a breaking change requires a `v2` proto package.
- **README is exactly workflow badges** plus the [docs.go-varyd](https://sqlrest.github.io/docs.go-varyd/) link; public documentation lives there, private specs in `sqlrest/_projects/specs/go-varyd/`.
- Gate: `make check` — gofumpt, vet, staticcheck, govulncheck, gocognit ≤ 7, exactly 100.0% coverage of the owned (non-generated) packages; `Makefile.local` scopes `COVER_PKGS` past `src/proto`. Managed files (Makefile, .golangci.yaml, workflows) are distributed by `nicerobot/tools.repository` — never hand-edit them.
