# Changelog

Xnix follows Semantic Versioning.

## [0.2.57] - 2026-07-13

### Added

- Added C Runtime application catalog records and CLI read methods for listing and querying safe desktop application metadata.
- Added C Runtime catalog tests that compile the C core, verify the sample application record, reject unknown application ids, and guard against backend detail exposure.

### Changed

- Extended Runtime C ownership from identity and write gates into the first application recipe metadata layer while keeping Ruby as the test and development-tool harness.

## [0.2.56] - 2026-07-13

### Added

- Added a C Runtime core library and CLI for stable Runtime identity, reserved write-method enumeration, and write-gate decisions.
- Added C Runtime core tests that compile and execute the C implementation while keeping Ruby as the test harness.

### Changed

- Documented the Runtime implementation direction so important product logic moves into C while Ruby remains focused on tests and development tooling.

## [0.2.55] - 2026-07-13

### Added

- Added `scripts/runtime_activation_smoke.rb`, a constrained packaged Runtime activation smoke that installs into a temporary root and executes the staged libexec wrapper against installed Runtime assets.

### Changed

- Extended Runtime activation installation to include Runtime Ruby libraries, version metadata, recipe assets, D-Bus contract references, and smoke references required by the packaged wrapper.

## [0.2.54] - 2026-07-13

### Added

- Added Runtime-owned write gate modeling through `xnix-runtime-write-gate`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Changed gated Runtime write-method dispatch so install, launch, snapshot, and restore requests now fail with an explicit `WriteMethodDisabled` boundary instead of a generic unsupported-method error.

## [0.2.53] - 2026-07-13

### Added

- Added Runtime-owned read-only method parity manifests through `xnix-runtime-method-parity-manifest`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Extended production owner smoke readiness so XML contract, Runtime dispatch, D-Bus client, smoke adapter, and session smoke coverage must agree on the Runtime read-only method set before write methods are considered.

## [0.2.52] - 2026-07-13

### Added

- Added Runtime-owned production owner smoke planning through `xnix-runtime-owner-smoke-plan`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Extended live-owner readiness so the future production Runtime owner has explicit smoke steps for bus-name ownership, read-only method parity, write-method rejection, and smoke-adapter boundary checks.

## [0.2.51] - 2026-07-13

### Added

- Added Runtime-owned live D-Bus owner gate modeling through `xnix-runtime-live-owner-gate`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Extended Runtime service readiness so KDE can distinguish aligned activation files from gated production D-Bus ownership, while keeping the smoke adapter explicitly non-production.

## [0.2.50] - 2026-07-13

### Added

- Added Runtime-owned Compatibility Center action review receipts through `xnix-compat-action-review`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Extended Compatibility Center action handling so KDE review intent is recorded as a non-executing Runtime receipt that cannot launch backends, persist settings, grant resources, or mutate the host root.

## [0.2.49] - 2026-07-13

### Added

- Added Runtime-owned Compatibility Center action queues through `xnix-compat-action-queue`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Extended KDE-facing Runtime read models so pending install readiness, settings changes, AI repair review, Runtime service ownership, and Portal policy review are grouped into non-executing Compatibility Center action cards.

## [0.2.48] - 2026-07-13

### Added

- Added Runtime-owned compatibility settings change planning through `xnix-compat-settings-change`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, and session-bus smoke coverage.

### Changed

- Extended compatibility settings so KDE can request a safe Runtime plan for setting changes without persisting state, granting resources, mutating the host root, or exposing backend details.

## [0.2.47] - 2026-07-13

### Added

- Added Runtime-backed compatibility settings exposure through diagnostics, read-only dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, and session-bus smoke coverage.

### Changed

- Extended `xnix-compat-settings` so KDE can display Runtime-owned run mode, resource access, device, network, and snapshot settings without owning backend policy or exposing implementation terminology.

## [0.2.46] - 2026-07-13

### Added

- Added `xnix-compat-install-plan`, a Runtime-owned compatibility install plan that joins artifact manifest, acquisition preflight, package source, state root, and recipe install gate readiness before any desktop activation or backend launch.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for install plan status.

### Changed

- Extended Runtime and KDE read models so installation readiness is visible as a first-class Compatibility Center signal before artifact download, desktop activation, or backend launch is enabled.

## [0.2.45] - 2026-07-13

### Added

- Added `xnix-compat-artifact-manifest`, a Runtime-owned compatibility artifact manifest model for signed manifest readiness, artifact groups, digest verification, cache namespaces, and rollback references without downloading artifacts or exposing cache paths.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for artifact manifest status.

### Changed

- Extended Runtime and KDE read models so artifact manifest readiness is visible as a first-class Compatibility Center signal before artifact download, installation, or backend launch is enabled.

## [0.2.44] - 2026-07-13

### Added

- Added `xnix-compat-acquisition-preflight`, a Runtime-owned acquisition preflight model for package-source readiness, signed artifact manifests, cache capacity, network policy review, and rollback markers without downloading artifacts or mutating the host root.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for acquisition preflight status.

### Changed

- Extended Runtime and KDE read models so acquisition readiness is visible as a first-class Compatibility Center signal before package download, installation, or backend launch is enabled.

## [0.2.43] - 2026-07-13

### Added

- Added `xnix-compat-package-source`, a Runtime-owned compatibility package source model for signed source policy, source-channel planning, cache preflight, and blocked package installation before launch binding.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for package source status.

### Changed

- Extended Runtime and KDE read models so package source selection is visible as a first-class Compatibility Center signal before backend launch is enabled.

## [0.2.42] - 2026-07-13

### Added

- Added `xnix-compat-state-root`, a Runtime-owned application state root model for per-application state ownership, snapshot eligibility, Portal file boundaries, and restore confirmation without creating host directories or exposing storage paths.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for application state root status.

### Changed

- Extended Runtime and KDE read models so planned state roots are visible as a first-class Compatibility Center signal before backend launch is enabled.

## [0.2.41] - 2026-07-13

### Added

- Added `xnix-compat-backend-binding`, a Runtime-owned managed compatibility backend binding model that reports launch readiness, required preflight, blocked unsafe actions, and desktop-safe status without exposing implementation details.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for backend binding status.

### Changed

- Extended Runtime and KDE read models so pending backend binding is visible as a first-class Compatibility Center signal before any launch backend is enabled.

## [0.2.40] - 2026-07-13

### Added

- Added `xnix-runtime-service-binding`, a Runtime-owned service binding model for D-Bus activation files, systemd hardening, the packaged libexec wrapper, contract exposure, and live-owner readiness.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke exposure, and focused coverage for Runtime service binding status.

### Changed

- Extended the Runtime foundation documentation to distinguish activation binding readiness from live production D-Bus ownership.
- Fixed the full smoke sequence so a clean milestone run builds the constrained Docker image and fetches Buildroot before configuring, downloading package sources, building, and booting QEMU.

## [0.2.39] - 2026-07-13

### Added

- Added `xnix-ai-repair-approval-gate`, a Runtime-owned approval gate model that blocks AI repair execution until Compatibility Center review, Runtime approval, and restore-point preflight gates pass.
- Added focused coverage for blocked gate decisions, required gates, approval-required actions, Runtime diagnostics, dispatch, D-Bus client reads, Compatibility Center summaries, and session-bus smoke exposure.

### Changed

- Extended Runtime diagnostics, read-only Runtime dispatch, D-Bus planning reads, and the KDE Compatibility Center model with AI repair approval gate summaries.

## [0.2.38] - 2026-07-13

### Added

- Added `xnix-ai-diagnostic-recommendation`, a Runtime-owned AI diagnostic recommendation model for review-first compatibility guidance.
- Added focused coverage for user-visible recommendations, approval-required actions, blocked AI tasks, Runtime diagnostics, dispatch, D-Bus client reads, Compatibility Center summaries, and session-bus smoke exposure.

### Changed

- Extended Runtime diagnostics, read-only Runtime dispatch, D-Bus planning reads, and the KDE Compatibility Center model with AI diagnostic recommendation summaries.

## [0.2.37] - 2026-07-13

### Added

- Added `xnix-ai-diagnostic-input`, a Runtime-owned AI diagnostic input model for safe recipe, run-plan, test-result, repair, signal, and privacy-boundary context.
- Added focused coverage for AI diagnostic safety boundaries, allowed and blocked AI tasks, Runtime diagnostics, dispatch, D-Bus client reads, Compatibility Center summaries, and session-bus smoke exposure.

### Changed

- Extended Runtime diagnostics, read-only Runtime dispatch, D-Bus planning reads, and the KDE Compatibility Center model with AI diagnostic input summaries.

## [0.2.36] - 2026-07-13

### Added

- Added `xnix-compat-test-result`, a Runtime-owned compatibility test result model for passed, pending, and blocked test outcomes.
- Added focused coverage for test result counts, diagnostic evidence, Compatibility Center summaries, Runtime diagnostics, dispatch, D-Bus client reads, and session-bus smoke exposure.

### Changed

- Extended Runtime diagnostics, read-only Runtime dispatch, D-Bus planning reads, and the KDE Compatibility Center model with compatibility test result summaries.

## [0.2.35] - 2026-07-13

### Added

- Added `xnix-compat-test-plan`, a Runtime-owned compatibility test plan model for preflight, smoke, and repair-readiness checks.
- Added focused coverage for recipe validation, Portal preflight, snapshot preflight, launch-binding readiness, Compatibility Center summaries, Runtime diagnostics, dispatch, D-Bus client reads, and session-bus smoke exposure.

### Changed

- Extended Runtime diagnostics, read-only Runtime dispatch, D-Bus planning reads, and the KDE Compatibility Center model with compatibility test plan summaries.

## [0.2.34] - 2026-07-13

### Added

- Added `xnix-portal-request-model`, a Runtime-owned XDG Desktop Portal request model for file, URI, print, screenshot, clipboard, camera, and remote-desktop operations.
- Added focused coverage for Portal D-Bus destination, interface, method, request handle tokens, Response completion, denied-request guidance, CLI validation, and backend-detail filtering.

### Changed

- Extended desktop integration manifests so settings artifacts expose both Portal policy evaluation and Portal request modeling.

## [0.2.33] - 2026-07-13

### Added

- Added `xnix-file-association-model`, a Runtime-owned file association model that maps recipe MIME types to generated desktop files and standard `mimeapps.list` content.
- Added activation staging for `usr/share/applications/mimeapps.list` so managed applications can become default handlers for their recipe file types inside a staging root.
- Added focused coverage for MIME defaults, portal-mediated file-open metadata, rollback tracking, and refusal to overwrite an existing `mimeapps.list`.

### Changed

- Extended desktop integration manifests so the file-manager artifact exposes file association generation alongside the Dolphin service menu.

## [0.2.32] - 2026-07-13

### Added

- Added `xnix-kwin-window-rule`, a KDE KWin window rule model that binds Runtime application identity to generated desktop files, task grouping, taskbar visibility, switcher visibility, and restore behavior.
- Added focused coverage for KWin rule identity matching, task-manager preservation, bounded window-manager scope, CLI validation, and backend-detail filtering.

### Changed

- Extended desktop integration manifests so the task-manager artifact exposes both the window identity model and the KWin window rule model.

## [0.2.31] - 2026-07-13

### Added

- Added `xnix-krunner-model`, a KDE KRunner query model that resolves Runtime-managed applications from names and file-oriented natural queries.
- Added focused coverage for KRunner launch delegation, desktop entry mapping, blank query behavior, extension hints, and backend-detail filtering.

### Changed

- Made `DBusRuntimeClient` parse D-Bus string arrays so KDE-facing read paths preserve list fields such as supported file extensions.

## [0.2.30] - 2026-07-13

### Added

- Added `DBusRuntimeClient` wrappers for Runtime planning reads covering engine catalog, run plans, repair plans, snapshot plans, and Portal access policy.
- Added focused D-Bus client coverage for planning reads and D-Bus boolean variant parsing.

### Changed

- Made the KDE-facing D-Bus client parse boolean `true` and `false` values into native booleans instead of strings.

### Verified

- Ran the full Buildroot and QEMU serial smoke test required for the 0.2.30 milestone.

## [0.2.29] - 2026-07-13

### Changed

- Extended the Runtime D-Bus contract with read-only planning methods for engine catalog, run plans, repair plans, snapshot plans, and Portal access policy.
- Routed the new read-only planning methods through `xnix-compatd dispatch`.
- Updated the Linux session-bus smoke adapter and smoke script to expose and call the new read-only planning methods.

### Added

- Added focused contract and dispatch coverage for the new planning methods while keeping write operations unsupported until backends exist.

## [0.2.28] - 2026-07-13

### Added

- Added `xnix-compat-engine-catalog`, a Runtime-owned catalog for automatic, local, and isolated compatibility engine choices.
- Added focused coverage for engine selection, Runtime ownership, pending backend readiness, CLI output, and backend-term filtering.

### Changed

- Made compatibility run plans select their engine summary through the engine catalog.
- Added Runtime probe capability reporting for the compatibility engine catalog.

## [0.2.27] - 2026-07-13

### Added

- Added `xnix-compat-snapshot-plan`, a Runtime-owned snapshot planning model for repair, compatibility-engine changes, and manual restore points.
- Added focused coverage for snapshot scope, restore availability, bounded retention, user-document preservation, CLI validation, and backend-term filtering.

### Changed

- Attached snapshot plan summaries to repair plans, Runtime diagnostics, and the KDE Compatibility Center model when a repair requires a restore point.

## [0.2.26] - 2026-07-13

### Added

- Added `xnix-compat-repair-plan`, a Runtime-owned repair planning model for pending engine setup, Portal approval, recipe trust blocking, and applied repair records.
- Added focused coverage for user approval requirements, snapshot requirements, rollback availability, notification mapping, CLI validation, and backend-term filtering.

### Changed

- Extended Runtime diagnostics and KDE Compatibility Center models with desktop-safe repair summaries.
- Reworded pending engine diagnostics so desktop-facing diagnostics do not expose implementation-specific backend terms.

## [0.2.25] - 2026-07-13

### Added

- Added `xnix-compat-run-plan`, a Runtime-owned compatibility run plan model for automatic, local, and isolated execution strategies.
- Added focused coverage for desktop-safe run strategy mapping, pending backend binding, Portal preflight, snapshot preflight, and backend-term filtering.

### Changed

- Extended launch request models with a desktop-safe run plan summary so KDE launchers can show Runtime planning status without exposing backend implementation details.
- Added Runtime probe capability reporting for compatibility run planning.

## [0.2.24] - 2026-07-13

### Added

- Added `xnix-portal-access-policy`, a Runtime-owned policy model for file, URI, print, screenshot, clipboard, camera, and remote-desktop Portal access decisions.
- Added focused coverage for Portal mediation, direct-access denial, default ask or deny decisions, CLI validation, and backend-term filtering.

### Changed

- Connected the desktop integration manifest settings artifact to the Portal access policy command so KDE settings can discover Runtime-owned sensitive desktop operation policy.

## [0.2.23] - 2026-07-13

### Changed

- Enforced recipe install gate preflight in `xnix-install-desktop-integration` before staged desktop activation.
- Made the desktop activation CLI use registry-backed recipe loading and block development-only registries in production mode by default.
- Added an explicit development install mode for safe local staging after digest verification.

### Added

- Added focused coverage that blocked production activation writes no desktop files and development staging still succeeds under the verified development registry.

## [0.2.22] - 2026-07-13

### Added

- Added `xnix-recipe-install-gate`, an install gate that evaluates a recipe registry, target application id, and production or development mode before activation.
- Added focused coverage for production blocking of development-only registries, development staging allowance, verified production registries, unknown recipe blocking, and desktop-safe output.

### Changed

- Documented recipe install gates as the enforcement layer above recipe trust policy decisions.

## [0.2.21] - 2026-07-13

### Added

- Added `xnix-recipe-trust-policy`, a policy model that evaluates recipe registry trust as production-trusted, development-only, or untrusted.
- Added focused coverage for development-only blocking reasons, production-trusted decisions, and desktop-safe summaries.

### Changed

- Documented recipe trust policy as the desktop-facing explanation layer above raw recipe trust probe signals.

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
