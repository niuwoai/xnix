# Changelog

Xnix follows Semantic Versioning.

## [0.2.20] - 2026-07-13

### Added

- Added Runtime probe reporting for recipe trust status, including registry-backed loading, digest verification, development registry state, and signed-recipe validation status.
- Added focused coverage for recipe trust capability reporting.

### Changed

- Documented the Runtime probe as the observable trust boundary for KDE surfaces and diagnostics.

### Verified

- Completed the twentieth-version full Buildroot build and QEMU serial smoke test.

## [0.2.19] - 2026-07-13

### Added

- Added a registry-backed recipe store that verifies `registry.json` before loading Runtime-managed recipes.
- Added focused coverage for registered recipe loading, digest mismatch rejection, and no-registry development fallback.

### Changed

- Made `xnix-compatd` use registry-backed recipe loading by default when a recipe registry is present.

## [0.2.18] - 2026-07-13

### Added

- Added `xnix-recipe-registry`, a recipe registry verifier for schema version, recipe identifiers, safe relative paths, SHA-256 digests, and signature status.
- Added a development recipe registry for the bundled sample application recipe.
- Added focused coverage for digest verification, digest mismatch rejection, unsafe path rejection, and explicit reporting that production signed-recipe validation is not yet enabled.

### Changed

- Documented recipe registry verification as the next application-library trust layer before production signed recipe storage.

## [0.2.17] - 2026-07-12

### Added

- Added `xnix-rollback-desktop-integration`, a receipt-based rollback command for staged KDE desktop activation files.
- Added SHA-256 activation receipts for generated launchers, Dolphin service menus, and persisted desktop integration manifests.
- Added focused coverage for checksum-verified removal, changed-file preservation, receipt cleanup, and host-root rejection.

### Changed

- Extended staged desktop activation results with rollback receipt metadata so Runtime-managed desktop integration can be audited and reverted.

## [0.2.16] - 2026-07-12

### Added

- Added `xnix-install-desktop-integration`, a staged desktop activation installer that writes generated application launchers, Dolphin service menus, and desktop integration manifests under a target root.
- Added focused coverage for staging-root safety, installed file paths, file modes, manifest persistence, and backend-term filtering.

### Changed

- Documented desktop activation as a staging-root operation so tests can verify KDE integration output without modifying the host desktop.

## [0.2.15] - 2026-07-12

### Added

- Added `xnix-desktop-integration-manifest`, a KDE recipe activation manifest that groups launcher, task manager, Dolphin, tray, notification, Compatibility Center, and settings artifacts.
- Added focused coverage to keep the manifest ordered, portal-aware, privilege-free, and free of backend implementation terminology.

### Changed

- Documented the desktop activation path as a single Runtime-owned manifest over the seven KDE entry-point models.

## [0.2.14] - 2026-07-12

### Added

- Added `xnix-compat-window-identity`, a KDE Task Manager and KWin-facing identity model for compatibility application windows.
- Added focused coverage for desktop file mapping, grouping, pinning, restore behavior, and KWin identity-only metadata.

### Changed

- Marked the Task Manager KDE entry point as an initial integration, completing initial coverage for all seven first-release KDE entry points.

## [0.2.13] - 2026-07-12

### Added

- Added `xnix-compat-tray-status`, a KDE-facing system tray status model for Runtime activity, compatibility attention state, and bridged tray applications.
- Added focused coverage to keep tray status user-facing and free of backend implementation terminology.

### Changed

- Marked the System Tray KDE entry point as an initial integration in the seven-entry-point status model.

## [0.2.12] - 2026-07-12

### Added

- Added `xnix-compat-settings`, a KDE-facing compatibility settings model for run mode, file access, device access, network access, and snapshots.
- Added focused coverage to keep settings user-facing and free of backend implementation terminology.

### Changed

- Marked the Settings KDE entry point as an initial integration in the seven-entry-point status model.

## [0.2.11] - 2026-07-12

### Added

- Added `xnix-compat-notify`, a Runtime event notification request entry point for KDE-facing desktop notifications.
- Added a notification request model for install failures, automatic repairs, compatibility mode changes, and approval-required events.
- Added focused coverage for notification urgency, actions, event validation, and desktop-facing output safety.

### Changed

- Marked the Notifications KDE entry point as an initial integration in the seven-entry-point status model.

## [0.2.10] - 2026-07-12

### Added

- Added a KDE seven-entry-point integration status model covering launcher, task manager, file manager, system tray, notifications, Compatibility Center, and settings.
- Added `xnix-kde-integration-status` for reporting which KDE entry points have an initial Runtime-backed integration and which remain planned.
- Added focused coverage for the first-release KDE scope and entry-point evidence.

### Verified

- Completed the tenth-version full Buildroot build and QEMU serial smoke test.

## [0.2.9] - 2026-07-12

### Added

- Added `xnix-compat-launch`, the managed desktop launcher entry point used by generated application `.desktop` files.
- Added a launch request model that validates Runtime application ids, preserves optional file URIs, marks file launches as portal-mediated, and targets the Runtime `Launch` method without invoking a backend directly.
- Added focused coverage for launcher request validation and desktop-facing output safety.

### Changed

- Documented the KDE launcher path as a Runtime request model rather than a direct Wine or VM command path.

## [0.2.8] - 2026-07-12

### Added

- Added a Dolphin service menu entry for opening selected files through Xnix Compatibility without exposing backend commands.
- Added `xnix-compat-open`, a file-open request entry point that validates `file://` URIs, resolves a Runtime application recipe by file extension, and emits a portal-required Runtime launch request model.
- Added focused coverage for Dolphin service menu packaging and file-open request validation.

### Changed

- Documented the file manager integration as a Runtime request path rather than a direct Wine or VM launcher.

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
