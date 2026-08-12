# Changelog

## 0.1.0 - 2026-08-12

- Define versioned schemas for resources, relationships, evidence, observations, snapshots, changes, approvals, releases, monitoring plans, Fleet reports, DevCycle candidates, and ReleaseGuard reports.
- Add valid examples, compatibility tests, secret-redaction tests, and an embedded offline validator CLI.
- Publish hand-synchronized Go and TypeScript types for product integration.
- Add optional relationship confidence, discovery evidence, first/last observation, and human confirmation audit fields.
- Document IDs, terminology, versioning, integration contracts, architecture, and data-protection rules.
- Finalize release-candidate source provenance and ReleaseGuard result phases, metric series reducers, and observation evidence fields before the first public release.
- Canonicalize CRLF and LF before hashing generated artifacts so the contract manifest is stable across operating systems.
