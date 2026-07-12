# Changelog

Xnix follows Semantic Versioning.

## [0.1.3] - 2026-07-12

### Added

- Resource-constrained Docker command builder for Xnix build and offline runtime containers.
- Unit tests that enforce container isolation and resource-limit arguments without invoking Docker.

## [0.1.2] - 2026-07-12

### Added

- Pinned Buildroot 2025.02.15 source URL and SHA-256 lock data.
- Container-only Buildroot source retrieval utility with offline lock verification.

## [0.1.1] - 2026-07-12

### Added

- Buildroot external-tree skeleton and x86_64 system configuration.
- Restricted Docker build definition with QEMU software emulation tooling.
- Focused Ruby layout verification for the initial system scaffold.

## [0.1.0] - 2026-07-12

### Added

- Initial Xnix project rules, automated-contributor guide, product overview, and ignore rules.
- Linux LTS and Buildroot primary implementation path.
- Constrained Colima and Docker test policy with focused tests for small versions and full QEMU smoke tests every tenth code version.

### Changed

- Standardized all project-facing content on English only.
