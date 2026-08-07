# FND-001 to FND-005 Checkpoint

> **内部里程碑备忘（非产品文档）。** 面向访客请阅读仓库根目录 `README.md`。

Date: 2026-08-07

## Completed

| ID | Deliverable |
|----|-------------|
| FND-001 | `docs/terminology.md` |
| FND-002 | `docs/id-conventions.md` |
| FND-003 | `docs/event-envelope.md` + `schemas/v0.1/event-envelope.json` |
| FND-004 | `schemas/v0.1/evidence.json` + examples |
| FND-005 | `schemas/v0.1/change-event.json` + examples |

Also shipped supporting docs required by the control brief: versioning, security, architecture, ADR-0001, and draft schemas for Resource / Observation / Snapshot / Relationship / Approval / Release.

## Later milestones (completed separately)

| ID | Notes |
|----|-------|
| FND-006 | `tests/schematest/` negative + forward-compat; `docs/compatibility.md` |
| FND-007 | Checked-in Go/TS types under `gen/`; `scripts/generate-types.md` |
| FND-008 | `tests/redaction` scanner; security doc updated |
