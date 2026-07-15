# Claude Code Contract Gap Work Packages

> Last updated: 2026-07-16 | Baseline: v0.2.216

This document is a branch-sized implementation backlog for Claude Code. It focuses on domains where Xnix already has contracts, previews, smoke scripts, or tests, but still lacks durable implementation.

Use one package per Claude Code branch. Do not ask a branch to implement multiple packages unless this file explicitly lists that package as a dependency.

## Global Rules for Every Package

- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Compatibility Runtime product logic.
- Use C only for low-level Runtime policy surfaces, ABI-shaped records, or already-owned C paths.
- Use Ruby for tests, smoke scripts, reports, and developer tooling.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Keep production write methods disabled unless the package explicitly enables a gated write path.
- Do not introduce privileged containers, host networking, Docker socket mounts, broad host-directory mounts, host-root mutation, production D-Bus ownership, real backend launch, or real network fetch by default.
- Do not expose raw backend commands, executable paths, storage paths, profile names, file contents, secrets, tokens, or private keys in KDE-facing or user-facing output.
- Every package must add or update targeted tests and run `ruby scripts/verify_layout.rb` before handoff.
- If the version lands on a tenth code version, run the full build and QEMU smoke gate required by repository policy.

## Package Dependency Map

```text
P1 Runtime Owner Read Service
  -> P2 Contract Parity and Drift Gates
  -> P3 Recipe Trust Store
       -> P4 Artifact Cache and Staging
            -> P5 Environment Lifecycle
                 -> P8 Execution Transaction Records

P6 Portal Permission Broker
  -> P7 Snapshot and Rollback Control Plane
       -> P8 Execution Transaction Records

P9 KDE Materialization Writers
  -> P10 KDE Live Shell Adapters

P11 Test, Repair, and AI Diagnostics
  -> P8 Execution Transaction Records

P12 Atomic KDE Image and QEMU Acceptance
  validates all production-shaped packages
```

## P1: Runtime Owner Read Service

### Mission

Move from preview-only owner candidates to a constrained Go owner service that can serve read-only Runtime methods on a test session bus.

### Existing contracts

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `internal/runtime/appidentity/runtime_owner_*.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `scripts/runtime_owner_candidate_smoke.rb`

### Deliverables

- A long-running smoke-owner mode for `xnix-runtime-owner`.
- Session-bus ownership only inside constrained smoke tests.
- Read-only method routing to Go Runtime handlers.
- Stable disabled errors for every write method.
- Owner lifecycle logs for startup, route table version, bus claim mode, readiness, and shutdown.
- Owner readiness states for `preview-only`, `smoke-owner`, and `production-owner-gated`.

### Non-goals

- Production bus ownership.
- System service installation.
- Backend launch.
- Host root writes.
- KDE-owned Runtime policy.

### Acceptance

- A container smoke can start the Go owner on a private session bus.
- Read methods match existing preview payloads or return explicit unsupported-route readiness.
- Write methods fail closed with `WriteMethodDisabled`.
- Owner output remains backend-detail safe.

### Required tests

```text
go test ./...
ruby scripts/container.rb runtime-owner-candidate-smoke
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/verify_layout.rb
```

## P2: Contract Parity and Drift Gates

### Mission

Turn the current contract sprawl into automated parity gates so D-Bus XML, Go routes, C smoke adapters, Ruby clients, CLI previews, and documentation do not drift silently.

### Existing contracts

- `runtime/dbus/org.xnix.Compatibility1.xml`
- `runtime/dbus/xnix_compatd_runtime_models.inc`
- `runtime/dbus/xnix_compatd_smoke.c`
- `cmd/xnix-runtime-go/runtime_owner_commands.go`
- `internal/runtime/appidentity/runtime_method_parity_manifest.go`
- `scripts/verify_layout.rb`

### Deliverables

- A machine-readable route manifest generated or verified from source.
- A parity test that compares D-Bus XML methods, Go owner routes, Go CLI commands, C smoke coverage, Ruby D-Bus client coverage, and documented write gates.
- A markdown and JSON report for parity failures.
- Clear classification for `go-owned`, `c-owned`, `adapter-owned`, `preview-only`, `write-gated`, and `unsupported`.
- CI-friendly exit codes without running QEMU.

### Non-goals

- New Runtime product behavior.
- Production D-Bus ownership.
- Relaxing write gates.

### Acceptance

- Adding a Runtime method in only one layer fails the parity gate.
- Write methods are recognized as intentionally gated rather than missing.
- Reports are deterministic and safe to commit.

### Required tests

```text
go test ./...
ruby scripts/verify_layout.rb
ruby -Ilib test/test_runtime_method_parity_manifest.rb
```

## P3: Production-Shaped Recipe Trust Store

### Mission

Replace development-only recipe assumptions with a read-only trust store and verifier boundary that can fail closed.

### Existing contracts

- `internal/runtime/recipe/`
- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `runtime/recipes/registry.json`
- `lib/xnix/compatibility/recipe_*`

### Deliverables

- A Go recipe store interface with local read-only roots.
- Registry metadata validation for schema, IDs, paths, digests, and declared signing state.
- A replaceable signature verifier boundary.
- Trust states: `production-trusted`, `development-only`, `unsigned`, `invalid`, and `blocked`.
- Trust diagnostics feeding install gates, owner readiness, and Compatibility Center pages.
- Fixture registries that remain useful but cannot pass production readiness.

### Non-goals

- Remote registry downloads.
- Private signing keys.
- Production trust without a real verifier.
- Host root mutation.

### Acceptance

- Invalid digest fails closed.
- Invalid or missing production signature fails production readiness.
- Development fixtures are visibly non-production.
- The same trust result is consumed by install gates and owner readiness.

### Required tests

```text
go test ./...
ruby -Ilib test/test_recipe_registry.rb
ruby -Ilib test/test_recipe_trust_policy.rb
ruby -Ilib test/test_recipe_install_gate.rb
ruby scripts/verify_layout.rb
```

## P4: Artifact Cache and Staging Pipeline

### Mission

Implement a safe local artifact pipeline that can parse manifests, verify digests, cache fixture artifacts, and stage receipts without mutating the host root.

### Existing contracts

- `internal/runtime/artifact/`
- `internal/runtime/appidentity/artifact_manifest.go`
- `internal/runtime/appidentity/acquisition_preflight.go`
- `internal/runtime/appidentity/install_plan.go`
- `test/test_compatibility_artifact_manifest.rb`
- `test/test_compatibility_acquisition_preflight.rb`
- `test/test_compatibility_install_plan.rb`

### Deliverables

- Artifact manifest parsing and validation.
- Digest verification for local fixture artifacts.
- Runtime state-root scoped cache namespace.
- Staging receipts that describe artifact identity, digest, source trust, and target role without exposing host paths.
- Separate readiness for acquisition, cache, staging, and install permission.
- Failure modes for digest mismatch, missing artifact, unsupported source, and trust-blocked artifact.

### Non-goals

- Network downloads by default.
- Host package-manager calls.
- Desktop activation writes.
- Backend process starts.

### Acceptance

- Digest mismatch blocks staging.
- Cache and staging stay under the configured Runtime or test root.
- Install remains gated by recipe trust and owner readiness.
- Receipts are safe for KDE-facing summaries.

### Required tests

```text
go test ./...
ruby -Ilib test/test_compatibility_package_source.rb
ruby -Ilib test/test_compatibility_acquisition_preflight.rb
ruby -Ilib test/test_compatibility_artifact_manifest.rb
ruby -Ilib test/test_compatibility_install_plan.rb
ruby scripts/verify_layout.rb
```

## P5: Environment Lifecycle and State Root

### Mission

Turn static backend and environment plans into a Runtime-owned lifecycle state machine with safe persistence.

### Existing contracts

- `internal/runtime/environment/`
- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `internal/runtime/appidentity/state_root.go`

### Deliverables

- Lifecycle states such as `missing`, `planned`, `staged`, `ready`, `repair-required`, and `blocked`.
- State persistence under a controlled Runtime state root.
- Transitions driven by recipe trust, artifact staging, Portal permission state, and snapshot readiness.
- A fixture store that can simulate installed, missing, broken, and repair-required environments.
- KDE-safe readiness explanations.

### Non-goals

- Starting Wine, Proton, VM, or other backend processes.
- Exposing shell commands, executable paths, profile names, or storage paths.
- Creating host-level state.

### Acceptance

- Lifecycle state survives inside a test-controlled state root.
- Readiness changes when trust, staging, permission, or snapshot inputs change.
- Execution remains disabled until P8 owns transaction records.

### Required tests

```text
go test ./...
ruby -Ilib test/test_application_state_root.rb
ruby -Ilib test/test_compatibility_backend_environment_plan.rb
ruby -Ilib test/test_compatibility_backend_binding.rb
ruby -Ilib test/test_compatibility_backend_lifecycle.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/verify_layout.rb
```

## P6: Portal Permission Broker

### Mission

Turn Portal request previews into a durable fake-first permission broker with stable request records.

### Existing contracts

- `internal/runtime/portal/`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/portal_request.go`
- `lib/xnix/compatibility/portal_*`
- `runtime/dbus/org.xnix.Compatibility1.xml`

### Deliverables

- Durable request records under a controlled Runtime state root.
- Request states: `planned`, `requested`, `granted`, `denied`, `cancelled`, `expired`, and `failed`.
- A transport interface with fake transport for tests and disabled real transport boundary.
- Request correlation handles suitable for future D-Bus callers.
- Permission summaries for Compatibility Center and execution readiness.

### Non-goals

- Real Portal calls by default.
- Direct file, camera, clipboard, printer, screenshot, or remote desktop access.
- Permission grants without an explicit fake or future real transport result.

### Acceptance

- Fake transport can grant and deny requests deterministically.
- Request state persists under a test root.
- Unsupported operations fail closed.
- KDE-facing output never exposes selected file contents or host paths.

### Required tests

```text
go test ./...
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
ruby scripts/verify_layout.rb
```

## P7: Snapshot and Rollback Control Plane

### Mission

Integrate constrained snapshots into install, settings, permission, repair, and execution safety flows.

### Existing contracts

- `internal/runtime/snapshot/`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/repair_plan.go`
- `lib/xnix/compatibility/compatibility_snapshot_plan.rb`
- `lib/xnix/compatibility/desktop_activation_rollback.rb`

### Deliverables

- Snapshot records for managed Runtime state under a controlled root.
- Snapshot verification receipts with digest and scope metadata.
- Rollback planning and rollback receipt records.
- Restore preflight checks consumed by settings, repair, install, and execution flows.
- Clear distinction between Runtime state rollback and OS-level atomic rollback.

### Non-goals

- Host filesystem snapshots.
- OS deployment rollback implementation.
- Restoring arbitrary user files.
- Write methods enabled by default.

### Acceptance

- Snapshot records can be created and verified inside a test root.
- Corrupt or missing snapshot data blocks restore readiness.
- Repair and execution plans can require a snapshot without performing restore.

### Required tests

```text
go test ./...
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby scripts/verify_layout.rb
```

## P8: Execution Transaction Records

### Mission

Create a durable launch transaction model that records readiness and user decisions without starting real compatibility backends.

### Existing contracts

- `internal/runtime/appidentity/execution_request.go`
- `internal/runtime/appidentity/execution_review.go`
- `internal/runtime/appidentity/execution_decision.go`
- `internal/runtime/appidentity/execution_preflight.go`
- `internal/runtime/appidentity/execution_resource_grant.go`
- `internal/runtime/appidentity/execution_transaction.go`
- `internal/runtime/appidentity/execution_session.go`
- `internal/runtime/appidentity/execution_session_status.go`

### Deliverables

- Transaction records under a controlled Runtime state root.
- States such as `draft`, `reviewed`, `preflight-blocked`, `ready`, `approved`, `denied`, `expired`, and `failed`.
- Inputs from recipe trust, artifact staging, environment lifecycle, Portal permissions, and snapshot readiness.
- Stable disabled launch write-gate behavior until all gates pass.
- Session identity records that remain synthetic and safe.

### Non-goals

- Starting backend processes.
- Creating real windows or observing live windows.
- Granting file permissions directly.
- Enabling production `Launch`.

### Acceptance

- Transaction state changes deterministically in tests.
- Missing trust, staging, permission, or snapshot data blocks readiness.
- Launch remains disabled unless a later package explicitly enables a guarded fake launch mode.

### Required tests

```text
go test ./...
ruby -Ilib test/test_launch_request.rb
ruby -Ilib test/test_compatibility_run_plan.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/verify_layout.rb
```

## P9: KDE Materialization Writers

### Mission

Move from desktop activation previews to gated writers that can stage desktop integration artifacts and produce rollback receipts under a test root.

### Existing contracts

- `internal/runtime/appidentity/desktop_activation_*.go`
- `lib/xnix/compatibility/desktop_activation_installer.rb`
- `lib/xnix/compatibility/desktop_activation_rollback.rb`
- `kde/dolphin/servicemenus/`
- `kde/plasmoids/`

### Deliverables

- Test-root writers for desktop files, MIME association fragments, service menus, tray metadata, notification metadata, and activation manifests.
- Activation receipts with relative paths, modes, digests, and rollback instructions.
- Refusal to overwrite existing user files unless an explicit merge strategy exists.
- Clear production vs development activation modes.

### Non-goals

- Host desktop writes by default.
- KDE cache refresh on the host.
- Backend launch.
- MIME default overwrite without merge support.

### Acceptance

- Writers can materialize into a temporary root only.
- Rollback receipts can remove or restore staged files inside the same root.
- Existing file conflicts fail closed.
- KDE-facing identities stay normal Linux application identities.

### Required tests

```text
go test ./...
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby -Ilib test/test_file_association_model.rb
ruby scripts/verify_layout.rb
```

## P10: KDE Live Shell Adapters

### Mission

Build minimal runnable KDE-side adapters that consume Runtime read models without owning Runtime policy.

### Existing contracts

- `kde/plasmoids/org.xnix.compatibilitycenter/`
- `kde/dolphin/servicemenus/`
- `internal/runtime/appidentity/kde_*`
- `internal/runtime/appidentity/window_identity_routes.go`
- `lib/xnix/compatibility/krunner_model.rb`
- `lib/xnix/compatibility/tray_status_model.rb`

### Deliverables

- Minimal KRunner adapter for Runtime query plans.
- Minimal Dolphin service menu flow to Runtime file-open preview or future launch transaction draft.
- Minimal tray/status adapter consuming Runtime status.
- Minimal task-manager and KWin identity adapter data flow.
- A KDE smoke that verifies adapters can render or emit safe requests without backend launch.

### Non-goals

- KDE owning compatibility policy.
- Direct backend calls from KDE components.
- Host-level KDE config mutation outside test roots.
- Real live process observation unless explicitly fake-scoped.

### Acceptance

- KDE adapters consume Runtime read models or fake Runtime responses.
- Normal user output avoids backend implementation terminology.
- The smoke can run in a constrained container or documented local fake mode.

### Required tests

```text
go test ./...
ruby scripts/kde_first_presence_smoke.rb
ruby -Ilib test/test_kde_center_model.rb
ruby -Ilib test/test_krunner_model.rb
ruby -Ilib test/test_tray_status_model.rb
ruby scripts/verify_layout.rb
```

## P11: Test, Repair, and AI Diagnostics Pipeline

### Mission

Turn compatibility test plans, test results, repair plans, and AI diagnostics into a safe fixture-driven pipeline.

### Existing contracts

- `internal/runtime/appidentity/test_plan.go`
- `internal/runtime/appidentity/test_result.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/diagnostics.go`
- `internal/runtime/appidentity/ai_diagnostics.go`
- `lib/xnix/compatibility/ai_*`
- `lib/xnix/compatibility/compatibility_repair_plan.rb`

### Deliverables

- Fixture-driven test runner records that do not start real backends.
- Test result persistence under a controlled Runtime state root.
- Repair recommendations derived from safe test signals.
- AI provider interface with disabled production default and fake provider for tests.
- Privacy filters that exclude file contents, host paths, backend logs, secrets, and raw commands.
- Approval gates connecting AI recommendations to repair plans and snapshots.

### Non-goals

- Network AI calls by default.
- Automatic repair execution.
- Reading user documents.
- Parsing raw backend logs into user output.

### Acceptance

- Fake test runs produce deterministic results and repair suggestions.
- AI diagnostics can run with a fake provider and are disabled otherwise.
- Repair remains review-first and snapshot-gated.

### Required tests

```text
go test ./...
ruby -Ilib test/test_compatibility_test_plan.rb
ruby -Ilib test/test_compatibility_test_result.rb
ruby -Ilib test/test_compatibility_repair_plan.rb
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_ai_repair_approval_gate.rb
ruby scripts/verify_layout.rb
```

## P12: Atomic KDE Image and QEMU Acceptance

### Mission

Create a reproducible KDE desktop image acceptance path that proves the Runtime can be packaged and smoke-tested without risking the host.

### Existing contracts

- `image/kinoite/`
- `docs/kde-image-pipeline.md`
- `scripts/build_kde_image.rb`
- `scripts/boot_kde_image.rb`
- `scripts/build_kde_disk.rb`
- `lib/xnix/image/`
- `lib/xnix/qemu.rb`

### Deliverables

- Deterministic image manifest validation.
- A build path that can run in constrained Docker or documented Colima mode.
- QEMU boot smoke for the KDE image or a clearly staged image milestone.
- Runtime service activation files included in the image artifact.
- Serial log persistence and failure reports.
- Resource limits and loopback-only networking.

### Non-goals

- Privileged host installers.
- Host network exposure.
- Docker socket mounts.
- Host root writes.
- Shipping secrets or local machine paths.

### Acceptance

- A clean checkout can validate the image manifest.
- The image build path documents and enforces host-safety limits.
- QEMU smoke persists logs and exits with useful diagnostics.
- Runtime activation is present but write methods remain gated.

### Required tests

```text
ruby -Ilib test/test_kde_image.rb
ruby -Ilib test/test_kde_disk.rb
ruby -Ilib test/test_qemu.rb
ruby scripts/verify_layout.rb
```

Run the full image build and QEMU smoke only when explicitly authorized, because it is heavier than normal package tests.
