# Generating / maintaining types (FND-007)

**Source of truth:** `schemas/v0.1/*.json`  
**Checked-in types (v0.1):** reviewed wire types with executable schema, artifact-hash, and round-trip gates.

## Output paths

| Language | Path | Notes |
|----------|------|--------|
| Go | `gen/go/lifecycle/v0_1/` | Import `github.com/GeneJie199/lifecycle-spec/gen/go/lifecycle/v0_1` |
| TypeScript | `gen/ts/v0.1/index.ts` | Export interfaces/types from `index.ts` |

`gen/` **must be committed** (not gitignored).

## v0.1 maintenance process

1. Edit JSON Schema under `schemas/v0.1/`.
2. Update examples under `examples/v0.1/` and `tests/schematest/`.
3. Update the Go structs and TypeScript interfaces to match.
4. Run `go test ./...`. Every public example is decoded into its Go type, encoded again, checked for meaningful data loss, and validated against its Schema.
5. Run `go run ./cmd/lifecycle-gen --write` after the tests pass. This refreshes `gen/schema-manifest.json`, which binds every public Schema to its Go/TypeScript type and records the exact source/artifact digests.
6. Run `go generate ./gen/go/lifecycle/v0_1/...` to verify the committed manifest is current.
7. Commit Schema, examples, tests, generated types, and the manifest together.

## Go generate hint

`gen/go/lifecycle/v0_1/types.go` includes:

```go
//go:generate go run ../../../../cmd/lifecycle-gen --check
```

`go generate` fails when a Schema, Go type, or TypeScript type changes without a reviewed manifest refresh. It also fails when a new public Schema has no Go/TypeScript mapping. Use `go run ./cmd/lifecycle-gen --write` only after the all-schema round-trip suite passes.

## Guarantees and boundary

JSON Schema remains normative. `lifecycle-gen` intentionally does not overwrite reviewed public API names or documentation comments. It makes drift visible and testable: schema hashes, both language artifacts, complete type mappings, loss-aware Go wire round trips, and schema validation all have to agree before CI passes.
