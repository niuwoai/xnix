# Claude Code Implementation Packages

> Last updated: 2026-07-15 | Current version: v0.2.194

This document splits the contract-heavy Xnix Runtime roadmap into implementation packages that are large enough to matter but isolated enough for Claude Code to complete without owning the whole product at once.

Each package should be implemented on its own branch and version bump. Do not combine unrelated packages in one commit or pull request.

## Global Rules for Every Package

- Keep all source, comments, tests, documentation, CLI output, and fixtures in English.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Keep user-visible KDE and Runtime output free of backend implementation terms, raw commands, host paths, and executable details unless a package explicitly says otherwise.
- Do not use privileged containers, host networking, Docker socket mounts, broad host directory mounts, or host root mutations.
- Do not start real compatibility backends until the package explicitly enables a gated execution path.
- Preserve read-only D-Bus behavior and explicit write-gate failures until a package owns production write readiness.
- Prefer Go for durable Runtime product logic, C for low-level policy or ABI-shaped surfaces, and Ruby for tests and developer tooling.
- Keep each source file below 2,000 lines when practical and split before 3,000 lines.
- Run targeted tests for the package, `ruby scripts/verify_layout.rb`, and relevant constrained container smoke tests before commit.

## Package Map

| Package | Primary gap | Suggested owner layer | Depends on |
| --- | --- | --- | --- |
| P1 Go owner adapters for C-backed read-only methods | Many D-Bus read methods still route through C adapter policy records | Go Runtime | Current route manifest |
| P2 Production Runtime owner process | Owner readiness is modeled but no long-running Go D-Bus owner is enabled | Go Runtime and D-Bus | P1 partially |
| P3 Recipe storage and trust | Registry digest checks exist, but signed production recipe storage is not implemented | Go Runtime | Current recipe registry |
| P4 Compatibility package acquisition | Acquisition/install contracts exist, but no safe artifact fetch, cache, or install pipeline exists | Go Runtime and C policy | P3 |
| P5 Backend environment lifecycle | Selection/environment/binding contracts exist, but no lifecycle state machine or provisioner exists | Go Runtime | P3, P4 |
| P6 Execution transaction pipeline | Launch intent and transaction contracts exist, but launch remains blocked everywhere | Go Runtime | P2, P5 |
| P7 Snapshot and rollback implementation | Snapshot/test/repair contracts exist, but no real snapshot store exists | Go Runtime and system integration | P5 |
| P8 Portal request broker | Portal policies are modeled, but no broker owns request objects or completion tracking | Go Runtime and XDG Portal | P2 |
| P9 KDE materialization writer | KDE previews exist, but production desktop files, MIME files, and activation receipts are not written by Go | Go Runtime and KDE integration | P2, P3 |
| P10 Atomic image and QEMU product smoke | Learning baseline boots, but the flagship KDE image pipeline is still mostly absent | Build and packaging | P2, P9 |

## P1: Go Owner Adapters for C-Backed Read-Only Methods

### Goal

Move remaining C-backed read-only Runtime methods behind native Go owner handlers or thin Go adapter boundaries, then update `runtime-owner-route-manifest-preview` until the C-backed route count moves steadily toward zero.

### Recommended slice size

Migrate 3 to 5 related methods per version. Good first slices:

- `GetRunPlan`, `GetExecutionReadiness`, `GetLaunchIntent`
- `GetBackendCapabilityMatrix`, `GetBackendLifecycle`, `GetBackendEnvironmentPlan`
- `GetPortalAccessPolicy`, `GetPortalRequestPlan`
- `GetCompatibilityInstallPlan`, `GetCompatibilityArtifactManifest`, `GetCompatibilityAcquisitionPreflight`

### Suggested files

- `internal/runtime/appidentity/*.go`
- `cmd/xnix-runtime-go/*.go`
- `internal/runtime/appidentity/runtime_owner_route_manifest.go`
- `runtime/core/xnix_runtime_core_cli_dispatch.inc`
- `scripts/verify_layout.rb`
- Relevant Ruby tests under `test/`

### Acceptance criteria

- New Go CLI preview commands render JSON with stable schemas.
- Route manifest classifies migrated methods as `go-runtime-cli`.
- C adapter route count decreases.
- All previews pass `validateNoBackendTerms`.
- No write methods, backend launch, network, or host mutation are enabled.

### Required tests

- `go test ./...`
- `ruby scripts/verify_layout.rb`
- Relevant Ruby model tests for the migrated methods
- `ruby -Ilib test/test_runtime_method_parity_manifest.rb`
- `ruby -Ilib test/test_runtime_service_binding.rb`

## P2: Production Runtime Owner Process

### Goal

Turn the current owner readiness previews into a constrained long-running Go Runtime owner candidate without making it the default production owner yet.

### Scope

- Add a Go process entry point that can serve read-only Runtime methods from the current preview/adapters.
- Keep write methods returning explicit gated errors.
- Add lifecycle logs and health/readiness probes.
- Keep D-Bus name claiming behind a flag or test-only mode.

### Out of scope

- Real backend launch.
- Real install.
- Host root writes.
- Privileged services.

### Suggested files

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `runtime/systemd/xnix-compatd.service`
- `scripts/dbus_session_smoke.rb`
- `scripts/container.rb`

### Acceptance criteria

- A constrained container smoke can start the Go owner on a session bus.
- Read-only calls return the same safe payloads as the CLI previews.
- Write calls return deterministic disabled errors.
- Owner readiness still reports production ownership as gated unless the smoke flag is active.

### Required tests

- `go test ./...`
- `ruby scripts/verify_layout.rb`
- `ruby scripts/container.rb runtime-dbus-smoke`
- `ruby scripts/container.rb kde-center-dbus-smoke`

## P3: Recipe Storage and Trust

### Goal

Implement production-shaped recipe storage and trust verification while keeping development fixtures usable.

### Scope

- Define a durable recipe store interface in Go.
- Support local read-only recipe roots and signed registry metadata.
- Verify digest and signature status through a replaceable verifier boundary.
- Preserve development-only registry behavior as explicitly non-production.
- Add trust-state diagnostics for Compatibility Center and owner readiness.

### Out of scope

- Remote downloads.
- Secret management.
- Production key material in the repository.

### Suggested files

- `internal/runtime/appidentity/registry.go`
- New `internal/runtime/recipe/`
- `runtime/recipes/registry.json`
- `lib/xnix/compatibility/recipe_*`
- `test/test_recipe_registry.rb`

### Acceptance criteria

- Invalid digests fail closed.
- Missing signatures are development-only, never production trusted.
- No private keys, tokens, or signing secrets enter the repo.
- Recipe trust preview and owner readiness consume the same trust state.

### Required tests

- `go test ./...`
- `ruby -Ilib test/test_recipe_registry.rb`
- `ruby -Ilib test/test_recipe_trust_policy.rb`
- `ruby -Ilib test/test_recipe_install_gate.rb`
- `ruby scripts/verify_layout.rb`

## P4: Compatibility Package Acquisition

### Goal

Replace acquisition/install placeholders with a safe, reviewable artifact acquisition pipeline that can plan, verify, cache, and stage artifacts without mutating the host root.

### Scope

- Implement artifact manifest parsing in Go.
- Add cache namespace and digest verification models.
- Add dry-run acquisition with a local fixture source first.
- Keep network fetch disabled unless an explicit test fixture enables it.
- Stage into a project-controlled or container-controlled root only.

### Out of scope

- System package manager writes.
- Global host cache writes.
- Privileged installation.

### Suggested files

- New `internal/runtime/artifact/`
- `internal/runtime/appidentity/*artifact*`
- `internal/runtime/appidentity/*acquisition*`
- `scripts/container.rb`
- `test/test_compatibility_artifact_manifest.rb`
- `test/test_compatibility_acquisition_preflight.rb`

### Acceptance criteria

- Artifact digest mismatch blocks staging.
- Dry-run acquisition reports planned cache keys and staged files without downloading from the network.
- Runtime install plan consumes acquisition readiness.
- Host root mutation remains false.

### Required tests

- `go test ./...`
- Relevant Ruby artifact/acquisition/install tests
- `ruby scripts/verify_layout.rb`

## P5: Backend Environment Lifecycle

### Goal

Build a real Runtime-owned lifecycle state machine for compatibility environments without starting real backends by default.

### Scope

- Model environment states: `missing`, `planned`, `staged`, `ready`, `repair-required`, `blocked`.
- Store lifecycle state in a test-controlled state root.
- Connect backend selection, environment plan, binding plan, and lifecycle status.
- Add repair hints without executing repair.

### Out of scope

- Starting Wine, Proton, QEMU, or VM processes.
- Writing host user home directories.
- Exposing raw backend commands to KDE.

### Suggested files

- New `internal/runtime/environment/`
- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/runtime_owner_route_manifest.go`
- `test/test_compatibility_backend_lifecycle.rb`

### Acceptance criteria

- Lifecycle state survives within a controlled test state root.
- KDE previews continue to hide backend details.
- Backend binding reports real lifecycle readiness instead of static placeholders.
- Execution remains disabled until P6.

### Required tests

- `go test ./...`
- Backend selection/environment/binding/lifecycle Ruby tests
- `ruby scripts/verify_layout.rb`

## P6: Execution Transaction Pipeline

### Goal

Turn launch intent, review, preflight, resource grant, transaction, session, and session status contracts into one coherent blocked-by-default execution pipeline.

### Scope

- Introduce execution request IDs and non-persistent dry-run transactions.
- Connect write gate, Portal requirements, snapshot baseline, backend binding, and user review state.
- Add deterministic status transitions without launching a backend.
- Keep approval and launch separated.

### Out of scope

- Real process launch.
- Real permission grant.
- Real snapshot creation unless P7 is complete.

### Suggested files

- New `internal/runtime/execution/`
- `internal/runtime/appidentity/execution_*.go`
- `internal/runtime/appidentity/launch_intent.go`
- `test/test_launch_request.rb`
- Execution-related Go CLI tests

### Acceptance criteria

- Launch intent creates a safe request preview.
- Review decision changes transaction status but does not execute.
- Transaction reports exactly why launch is blocked.
- Session status remains preview-only unless a later package enables launch.

### Required tests

- `go test ./...`
- Execution request/review/decision/preflight/resource/transaction/session tests
- `ruby scripts/verify_layout.rb`

## P7: Snapshot and Rollback Implementation

### Goal

Implement a constrained snapshot store that can create, list, verify, and roll back test-controlled Runtime state roots.

### Scope

- Add snapshot metadata and content-addressed state archives.
- Support dry-run and test-root-only modes.
- Connect snapshot baseline readiness to execution preflight and repair planning.
- Add rollback receipts.

### Out of scope

- Host filesystem snapshots.
- Btrfs, ZFS, or system snapshot integration.
- Privileged rollback.

### Suggested files

- New `internal/runtime/snapshot/`
- `internal/runtime/appidentity/*snapshot*`
- `lib/xnix/compatibility/desktop_activation_rollback.rb`
- `test/test_compatibility_snapshot_plan.rb`

### Acceptance criteria

- Snapshot operations refuse paths outside the configured test state root.
- Rollback restores fixture state in tests.
- Execution preflight can require and verify a snapshot baseline.
- No host root mutation occurs.

### Required tests

- `go test ./...`
- Snapshot and rollback Ruby tests
- `ruby scripts/verify_layout.rb`

## P8: Portal Request Broker

### Goal

Implement Runtime-owned XDG Desktop Portal request brokering for preview/test mode, including request object tracking and completion state.

### Scope

- Add request records for file, URI, print, clipboard, screenshot, camera, and remote desktop operations.
- Add a broker interface so real Portal calls can be plugged in later.
- Implement a fake broker for container tests.
- Connect Portal request plans to execution preflight and permission review.

### Out of scope

- Direct desktop permission changes.
- Real user session Portal calls by default.
- Bypassing user review.

### Suggested files

- New `internal/runtime/portal/`
- `internal/runtime/appidentity/portal_request*`
- `lib/xnix/compatibility/portal_*`
- `test/test_portal_request_model.rb`

### Acceptance criteria

- Request object creation is explicit and test-gated.
- Permission state is tracked separately from execution approval.
- Portal failures are recoverable and visible in diagnostics.
- KDE never receives direct host paths or backend details.

### Required tests

- `go test ./...`
- Portal policy/request Ruby tests
- `ruby scripts/verify_layout.rb`

## P9: KDE Materialization Writer

### Goal

Turn KDE activation previews into a constrained writer that stages desktop files, service menus, MIME records, icons, and activation receipts under an unprivileged target root.

### Scope

- Implement Go or Ruby writer for test-controlled activation roots.
- Emit activation receipts for rollback.
- Keep all writes under an explicit target root.
- Keep desktop entry content user-facing and backend-safe.

### Out of scope

- Writing to the real host `/usr`, `/etc`, or user home by default.
- Setting MIME defaults globally.
- Starting applications.

### Suggested files

- `scripts/install_runtime_activation.rb`
- `lib/xnix/compatibility/desktop_activation_installer.rb`
- `internal/runtime/appidentity/desktop_activation_*.go`
- `kde/`
- `test/test_desktop_activation_installer.rb`

### Acceptance criteria

- Staging writes are deterministic and reversible.
- Activation receipt contains all written files.
- Rollback removes only files recorded in the receipt.
- Container activation smoke passes without host mutation.

### Required tests

- `go test ./...`
- Desktop activation installer/rollback tests
- `ruby scripts/container.rb build`
- `ruby scripts/container.rb runtime-dbus-smoke`
- `ruby scripts/verify_layout.rb`

## P10: Atomic Image and QEMU Product Smoke

### Goal

Create a reproducible product-image path that can eventually boot KDE-facing Xnix components while preserving the existing Buildroot/QEMU learning baseline.

### Scope

- Define image build inputs and lock files.
- Add a minimal smoke that verifies Runtime files and service definitions in an image root.
- Keep QEMU launch constrained and loopback-only.
- Preserve Buildroot baseline as a separate learning target.

### Out of scope

- Replacing the current Buildroot smoke.
- Privileged host image installation.
- Host networking.

### Suggested files

- `buildroot/`
- New image metadata under `boot/` or `docs/`
- `scripts/full_smoke.rb`
- `scripts/container.rb`
- `PRODUCT_OVERVIEW.md`

### Acceptance criteria

- Clean checkout can reproduce image metadata and smoke artifacts.
- QEMU smoke stores serial logs.
- SSH, if enabled, remains loopback-bound.
- Runtime product smoke is separate from the learning baseline smoke.

### Required tests

- `ruby scripts/full_smoke.rb` when the package touches the boot/image path.
- `ruby scripts/verify_layout.rb`
- Relevant container smoke tests.

## Suggested Parallelization

These packages can proceed mostly independently:

- P1 can run continuously in small slices.
- P3 can run alongside P1 because it mostly touches recipe trust.
- P8 can start with fake Portal brokering before P6 consumes it.
- P9 can improve staged activation while P2 works on the owner process.

Avoid starting P6 until P5 has a real lifecycle state model. Avoid starting production execution until P2, P3, P5, P7, and P8 provide enough safety gates.

## Handoff Prompt Template

Use this prompt when assigning one package to Claude Code:

```text
Work in /Users/rocky/Sites/xnix.
Implement package <PACKAGE_ID> from docs/claude-code-implementation-packages.md.
Follow AGENTS.md strictly.
Keep all project-facing text in English.
Do not use privileged containers, host networking, Docker socket mounts, broad host mounts, or host root mutation.
Make the smallest coherent version bump, update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md, and run the package-required tests plus ruby scripts/verify_layout.rb.
Do not commit unrelated .claude/ or tmp/ content.
Report changed files, tests, and remaining blockers.
```
