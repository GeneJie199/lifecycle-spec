# Changelog

## 0.2.0 - 2026-08-23

- Define the versioned `expected-changes/v1` contract for infrastructure, database, Fleet, and topology changes.
- Extend ReleaseGuard reports with exact change correlation evidence, cross-stage guidance, and the dedicated changes phase.
- Publish synchronized Go and TypeScript wire types plus positive and negative compatibility fixtures.

## 0.1.1 - 2026-08-12

- Run the formatting gate once on Linux while retaining Windows and macOS release tests, avoiding checkout line-ending false positives without weakening release coverage.

## 0.1.0 - 2026-08-12

- Define versioned schemas for resources, relationships, evidence, observations, snapshots, changes, approvals, releases, monitoring plans, Fleet reports, DevCycle candidates, and ReleaseGuard reports.
- Add valid examples, compatibility tests, secret-redaction tests, and an embedded offline validator CLI.
- Publish hand-synchronized Go and TypeScript types for product integration.
- Add optional relationship confidence, discovery evidence, first/last observation, and human confirmation audit fields.
- Document IDs, terminology, versioning, integration contracts, architecture, and data-protection rules.
- Finalize release-candidate source provenance and ReleaseGuard result phases, metric series reducers, and observation evidence fields before the first public release.
- Canonicalize CRLF and LF before hashing generated artifacts so the contract manifest is stable across operating systems.
