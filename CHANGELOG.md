# Changelog

Xnix follows Semantic Versioning.

## [0.2.7] - 2026-07-12

### Added

- Added a thin Runtime D-Bus client for KDE-facing read models, using the session bus and the stable `org.xnix.Compatibility1` contract.
- Added `--source auto|local|dbus` to `xnix-kde-center-model`, preferring D-Bus when a Runtime session service is available and falling back to the local read model otherwise.
- Added a constrained KDE Compatibility Center D-Bus smoke path that verifies Plasma-facing model data can be read from the Runtime session bus.

### Changed

- Kept the KDE read model desktop-facing while making the Runtime source explicit as either local fallback data or D-Bus session data.

## [0.2.6] - 2026-07-12

### Added

- Added a read-only KDE Compatibility Center model that converts Runtime applications and diagnostics into a safe desktop presentation model.
- Added a `xnix-kde-center-model` CLI for Plasma-facing model smoke tests.
- Added focused coverage to prevent the KDE presentation model from exposing backend storage or implementation terms.

### Changed

- Updated the Plasma package metadata to the current project version and documented the Runtime-backed Compatibility Center model boundary.

## [0.2.5] - 2026-07-12

### Added

- Added a Linux-only D-Bus session smoke adapter that owns `org.xnix.Compatibility1`, exposes `/org/xnix/Compatibility1`, and answers read-only Runtime calls over `gdbus`.
- Added a constrained container command for the Runtime D-Bus session smoke path.
- Added focused coverage for the D-Bus smoke script, adapter source, and container command.

### Changed

- Made the container image build prefer the local base-image cache with `--pull=false` so repeated verification is less sensitive to registry metadata failures.

### Verified

- Built `xnix-builder:0.2.5` and ran the Runtime D-Bus session smoke path in the constrained no-network, read-only-root container.

## [0.2.4] - 2026-07-12

### Fixed

- Made the Runtime activation installer test compatible with the constrained container's `noexec` temporary filesystem by checking executable mode bits instead of attempting an executable-path predicate.

### Verified

- Rebuilt the constrained Docker image with D-Bus tooling and ran the Runtime activation smoke path in a no-network, read-only-root container.

## [0.2.3] - 2026-07-12

### Added

- Added a root-staging installer for Runtime activation files, covering the packaged libexec wrapper, D-Bus system service, and systemd unit.
- Added container image dependencies for future Linux D-Bus smoke tests: `dbus` and `libglib2.0-bin`.
- Added focused coverage for activation-file installation into an unprivileged temporary root.

## [0.2.2] - 2026-07-12

### Added

- Added a Runtime method-dispatch layer for read-only D-Bus contract methods: `ListApplications`, `GetApplication`, and `GetDiagnostics`.
- Added a packaged libexec wrapper that matches the D-Bus and systemd activation path.
- Added focused tests for Runtime method dispatch and activation-path consistency.

### Changed

- Kept launch, install, snapshot, and restore methods rejected until Wine/VM backends and request signaling exist.

## [0.2.1] - 2026-07-12

### Added

- Added a runnable Compatibility Runtime daemon core with JSON probe, application listing, per-application diagnostics, introspection output, and managed recipe-store loading.
- Added a bundled sample recipe and focused tests for recipe loading and runtime daemon behavior.

### Changed

- Documented that the Runtime core is executable while the real D-Bus binding, Wine backend, VM backend, and signed recipe trust policy remain pending.

## [0.2.0] - 2026-07-12

### Added

- Added the independent Compatibility Runtime foundation: validated application recipes, ordinary Linux desktop entry generation, a D-Bus API contract, a hardened systemd service definition, and a Plasma 6 Compatibility Center package skeleton.
- Added the atomic KDE Plasma product architecture, with Wine/Proton and Windows VM backends explicitly separated from the desktop shell.

### Changed

- Repositioned Buildroot/QEMU as the verified learning baseline and selected a Fedora Kinoite-compatible atomic KDE desktop as the flagship product direction.

## [0.1.16-rc1] - 2026-07-12

### Fixed

- Removed the unsupported `UsePAM` directive after the QEMU serial boot showed that Buildroot OpenSSH rejected it during `sshd` startup.

### Verified

- Completed the constrained Buildroot, QEMU serial, DHCP, and loopback key-authenticated `sshd` smoke path.

## [0.1.16] - 2026-07-12

### Added

- Added an isolated QEMU SSH smoke-test path using a disposable test key stored only in the Docker-managed cache volume and a container-loopback-only forwarding rule.

### Changed

- Updated the documented constrained runtime limit to 4 GiB inside the user-authorized 6 GiB Colima VM.

## [0.1.15-rc2] - 2026-07-12

### Fixed

- Raised the constrained build-container memory budget to 4 GiB after Colima was explicitly authorized to use up to 6 GiB, while retaining single-job compilation and a 2 GiB VM safety reserve.

## [0.1.15-rc1] - 2026-07-12

### Fixed

- Limited Buildroot to one job after the initial cross-GCC automata generator was killed by the constrained container memory limit.

## [0.1.15] - 2026-07-12

### Added

- Human-readable serial boot failure summaries for QEMU smoke-test diagnostics.

## [0.1.14] - 2026-07-12

### Added

- Named persistent full-build container command so Buildroot logs remain available after a failed build.

## [0.1.13] - 2026-07-12

### Added

- OpenSSH server hardening overlay that disables password authentication and root password login.

## [0.1.12] - 2026-07-12

### Added

- Loopback-only non-interactive SSH probe command and focused unit coverage.

## [0.1.11] - 2026-07-12

### Added

- OpenSSH `sshd` serial-service evidence verifier and focused unit coverage.

## [0.1.10-rc1] - 2026-07-12

### Fixed

- Added a constrained Buildroot dependency-download phase before the offline full build, allowing the v0.1.10 verification to obtain its declared sources without granting network access to compilation or QEMU.

## [0.1.10] - 2026-07-12

### Added

- Full-build and QEMU serial-smoke runner protected by the tenth-version milestone gate.

## [0.1.9] - 2026-07-12

### Added

- Version gate that permits full-build and smoke-test execution only at every tenth formal code version.
- Focused unit coverage for milestone and pre-release version handling.

## [0.1.8] - 2026-07-12

### Added

- Optional loopback-only QEMU SSH forwarding command and focused port-exposure unit coverage.

## [0.1.7] - 2026-07-12

### Added

- Xnix hostname configuration and serial-log boot evidence verifier.
- Focused unit coverage for successful and incomplete boot logs.

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
