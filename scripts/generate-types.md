# Generating / maintaining types (FND-007)

**Source of truth:** `schemas/v0.1/*.json`  
**Checked-in types (v0.1):** hand-synced is OK until an automated generator is adopted.

## Output paths

| Language | Path | Notes |
|----------|------|--------|
| Go | `gen/go/lifecycle/v0_1/` | Import `github.com/GeneJie199/lifecycle-spec/gen/go/lifecycle/v0_1` |
| TypeScript | `gen/ts/v0.1/index.ts` | Export interfaces/types from `index.ts` |

`gen/` **must be committed** (not gitignored).

## v0.1 maintenance process

1. Edit JSON Schema under `schemas/v0.1/`.
2. Update examples under `examples/v0.1/` and `tests/schematest/`.
3. Hand-sync Go structs and TypeScript interfaces to match.
4. Run `go test ./...` (includes `gen/go/lifecycle/v0_1` round-trip of `change-event.json`).
5. Commit Schema, examples, tests, and `gen/` together.

## Go generate hint

`gen/go/lifecycle/v0_1/types.go` includes:

```go
//go:generate echo See scripts/generate-types.md — v0.1 types are checked in and hand-synced from schemas/v0.1
```

Running `go generate ./gen/go/lifecycle/v0_1/...` only prints the reminder; it does not rewrite files.

## Future (post-0.1)

Optional automation ideas (not required for 0.1):

- [quicktype](https://quicktype.io/) / [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)-style generators from JSON Schema
- CI check that `gen/` matches Schema (hash or golden-file diff)

Until then, treat Schema as normative and keep `gen/` in lockstep manually.
