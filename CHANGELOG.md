# Changelog

Xnix follows Semantic Versioning.

## [0.1.6] - 2026-07-12

### Added

- Restricted x86_64 QEMU TCG command generation for the future Xnix boot smoke test.
- Focused unit coverage for QEMU memory, CPU, console, and network restrictions.

## [0.1.5] - 2026-07-12

### Added

- Buildroot configuration and full-build command generation against the Docker-managed source cache.
- Focused unit coverage for the x86_64 defconfig, external tree, output directory, and offline build isolation.

## [0.1.4] - 2026-07-12

### Added

- Docker-managed Buildroot source cache volume for constrained source retrieval without a host-directory mount.
- Networked source-retrieval container command with targeted isolation unit coverage.

## [0.1.3-rc3] - 2026-07-12

### Fixed

- Removed the runtime bind mount that is unavailable when Colima host-directory mounts are disabled; project files are now copied into the tool image during its build.

## [0.1.3-rc2] - 2026-07-12

### Fixed

- Replaced the invalid `rw` field in the Docker bind-mount specification with Docker's compatible default writable bind mount.

## [0.1.3-rc1] - 2026-07-12

### Fixed

- Removed unsupported Docker Buildx CPU and memory arguments from the image-build command; resource isolation remains enforced by the 1 GiB, one-CPU Colima VM.

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
