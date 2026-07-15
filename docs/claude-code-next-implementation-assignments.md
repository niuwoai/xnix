# Claude Code Next Implementation Assignments

> Last updated: 2026-07-16 | Baseline: v0.2.231

This document is a short assignment board for giving Claude Code large, relatively independent Xnix implementation work.

Use it when the current problem is: "the contract exists, but durable implementation evidence is still thin." It is intentionally shorter than the full package maps so one assignment can be copied into Claude Code without handing it the whole roadmap.

This document does not replace or modify `docs/claude-code-implementation-packages.md`. For deeper background, read:

- `docs/claude-code-domain-dispatch.md`
- `docs/claude-code-large-empty-domain-assignments.md`
- `docs/claude-code-empty-domain-implementation-packages.md`
- `docs/claude-code-contract-gap-work-packages.md`
- `docs/claude-code-mainline-implementation-plan.md`

## Global Rules for Every Assignment

Every Claude Code branch must:

- Implement exactly one assignment from this file unless explicitly told otherwise.
- Keep source, comments, fixtures, tests, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior.
- Use C only for low-level D-Bus transport, ABI-shaped records, or already-owned Runtime policy surfaces.
- Use Ruby for tests, smoke scripts, reports, and developer tooling.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Add success, failure, and blocked-state tests.
- Run the assignment-specific verification commands plus `ruby scripts/verify_layout.rb`.
- Run `ruby scripts/runtime_contract_drift_report.rb --format json` when touching D-Bus XML, Runtime route manifests, owner dispatch, smoke adapters, D-Bus clients, or Runtime CLI route commands.
- Stop and report if the assignment requires behavior outside the selected package.

Unsafe behavior remains disabled unless a future assignment explicitly enables a constrained test-only path:

- Production D-Bus ownership.
- Runtime write methods.
- Real Wine, Proton, VM, or compatibility backend launch.
- Real Portal transport calls.
- Network artifact fetch by default.
- Host package-manager calls.
- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Host-root mutation.
- Raw backend command, executable path, storage path, profile name, host path, file content, secret, token, or private-key exposure in KDE-facing output.

## How to Assign Work

1. Pick one assignment ID from the table.
2. Copy the assignment prompt into Claude Code.
3. Tell Claude Code to stop if it needs unsafe behavior or cross-package ownership.
4. Review the handoff using the completion template at the end of this file.

Recommended first wave:

- A1 Runtime owner process boundary
- A2 Trust, artifact, cache, and install pipeline
- A3 Runtime state root and backend lifecycle
- A8 Developer evidence and drift harness

These four reduce the risk that later branches only add more preview contracts.

## Assignment Index

| ID | Assignment | Primary owner | Can start now | Main evidence |
| --- | --- | --- | --- | --- |
| A1 | Runtime owner process boundary | Go Runtime | Yes | Private session-bus owner smoke and closed write responses |
| A2 | Trust, artifact, cache, and install pipeline | Go Runtime | Yes | Digest-verified fixtures, cache/staging receipts, install readiness |
| A3 | Runtime state root and backend lifecycle | Go Runtime | Yes | Durable lifecycle states without backend launch |
| A4 | Portal, snapshot, and rollback safety | Go Runtime | Yes | Fake Portal records, snapshot receipts, rollback evidence |
| A5 | KDE activation materialization | Go Runtime plus KDE integration | Yes, staged roots only | Desktop/MIME/service-menu/tray receipt evidence |
| A6 | Execution transaction ledger | Go Runtime | After A2-A4 are stronger | Reviewed transaction records that still block launch |
| A7 | Diagnostics, repair, and AI boundary | Go Runtime | Yes, fixture/no-provider mode | Diagnostic run records and review-only repair recommendations |
| A8 | Developer evidence and drift harness | Ruby tooling plus Go fixtures | Yes | Failing gates for orphan contracts and preview-only regressions |
| A9 | Atomic KDE image and QEMU acceptance | Build/image tooling | Later | Constrained image smoke and product-level acceptance reports |

## A1: Runtime Owner Process Boundary

### Current gap

Read-only Runtime ownership is represented through route manifests, in-process Go dispatch, owner smoke-batch records, and C D-Bus smoke bridge evidence. The missing implementation is a constrained long-running Go owner that can own a private test session-bus name and serve reads through the same owner table.

### Start from

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `internal/runtime/appidentity/runtime_owner_*.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `runtime/dbus/xnix_compatd_smoke.c`
- `scripts/runtime_owner_candidate_smoke.rb`
- `scripts/dbus_session_smoke.rb`
- `scripts/runtime_contract_drift_report.rb`

### Deliver

- Smoke-owner mode that can claim the Runtime bus name only on a private test session bus.
- Read-only D-Bus routing through Go owner handlers or explicit unsupported-read markers.
- Deterministic disabled errors for every write method.
- Safe lifecycle logs for startup, route-table version, bus mode, readiness, dispatch result, and shutdown.
- Contract parity checks across XML, Go owner routes, owner dispatch, smoke-batch records, C smoke bridge, D-Bus client, and CLI route commands.

### Keep disabled

- Production bus ownership.
- System service installation.
- Runtime write methods.
- Backend launch.
- Host-root writes.
- KDE-owned Runtime policy.

### Verification

```text
go test ./cmd/xnix-runtime-owner ./internal/runtime/owner ./cmd/xnix-runtime-go ./internal/runtime/appidentity
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/container.rb runtime-owner-candidate-smoke
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement A1 from docs/claude-code-next-implementation-assignments.md.

Target outcome:
- Add a constrained Go Runtime smoke-owner process that owns a private test session-bus name, serves read-only Runtime methods through the existing owner route table, and fails every write method closed.

Hard constraints:
- Do not enable production D-Bus ownership, system service installation, write methods, backend launch, network fetch, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add success, unsupported-read, disabled-write, lifecycle, and blocked-production-owner tests.
- Run the A1 verification commands.
```

## A2: Trust, Artifact, Cache, and Install Pipeline

### Current gap

Recipes, package sources, artifact manifests, acquisition preflight, and install plans exist. The missing implementation is a fixture-first pipeline that proves trust, digest verification, cache placement, staging receipts, and install readiness without network fetch or host mutation.

### Start from

- `internal/runtime/recipe/`
- `internal/runtime/artifact/`
- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `internal/runtime/appidentity/package_source.go`
- `internal/runtime/appidentity/acquisition_preflight.go`
- `internal/runtime/appidentity/artifact_manifest.go`
- `internal/runtime/appidentity/install_plan.go`
- `runtime/recipes/`
- `cmd/xnix-runtime-go/`

### Deliver

- Read-only recipe store with local roots.
- Registry validation for schema, identifiers, relative paths, SHA-256 digests, and declared signing state.
- Replaceable signature verifier boundary without committed production keys.
- Fixture artifact manifest parser and digest verifier.
- Runtime-root-scoped cache and staging receipts.
- Install readiness that joins recipe trust, artifact verification, cache state, staging state, and owner readiness.

### Keep disabled

- Default network fetch.
- Private keys or signing secrets.
- Host package-manager calls.
- Host-root mutation.
- Desktop activation writes.
- Backend process starts.

### Verification

```text
go test ./internal/runtime/recipe ./internal/runtime/artifact ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_recipe_registry.rb
ruby -Ilib test/test_recipe_trust_policy.rb
ruby -Ilib test/test_recipe_install_gate.rb
ruby -Ilib test/test_compatibility_package_source.rb
ruby -Ilib test/test_compatibility_acquisition_preflight.rb
ruby -Ilib test/test_compatibility_artifact_manifest.rb
ruby -Ilib test/test_compatibility_install_plan.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement A2 from docs/claude-code-next-implementation-assignments.md.

Target outcome:
- Build a fixture-first trust, artifact, cache, staging, and install-readiness pipeline that fails closed and never mutates the host root.

Hard constraints:
- Do not add default network fetch, package-manager calls, private keys, production trust bypasses, desktop writes, backend launch, privileged containers, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add digest success, digest mismatch, unsigned/development-only, cache-boundary, and staging-blocked tests.
- Run the A2 verification commands.
```

## A3: Runtime State Root and Backend Lifecycle

### Current gap

Backend selection, binding, capability, environment, lifecycle, run plan, and execution readiness surfaces exist. The missing implementation is durable Runtime state that can change when trust, artifact staging, and readiness evidence change.

### Start from

- `internal/runtime/environment/`
- `internal/runtime/appidentity/state_root.go`
- `internal/runtime/appidentity/backend_selection.go`
- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `internal/runtime/appidentity/run_plan.go`

### Deliver

- State-root-scoped lifecycle records with schema version, application id, operation id, created timestamp, lifecycle state, blocked reasons, and relative receipt paths.
- Lifecycle states such as `missing`, `planned`, `staged`, `ready`, `repair-required`, and `blocked`.
- Readiness joins from recipe trust, artifact staging, owner readiness, snapshot requirements, and Portal requirements.
- KDE-safe summaries that avoid raw backend commands, profile names, storage paths, and executable paths.

### Keep disabled

- Backend process creation.
- Real environment provisioning.
- Profile binding commits.
- Host-root mutation.
- Raw backend detail exposure.
- Runtime launch approval.

### Verification

```text
go test ./internal/runtime/environment ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_application_state_root.rb
ruby -Ilib test/test_compatibility_backend_selection_plan.rb
ruby -Ilib test/test_compatibility_backend_environment_plan.rb
ruby -Ilib test/test_compatibility_backend_binding.rb
ruby -Ilib test/test_compatibility_backend_lifecycle.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement A3 from docs/claude-code-next-implementation-assignments.md.

Target outcome:
- Add Runtime state-root lifecycle records and readiness transitions for compatibility environments without starting any backend.

Hard constraints:
- Do not create backend processes, mutate host roots, expose raw backend details, enable launch, enable write methods, or claim production D-Bus ownership.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add lifecycle persistence, blocked-readiness, staged-readiness, and unsafe-detail-redaction tests.
- Run the A3 verification commands.
```

## A4: Portal, Snapshot, and Rollback Safety

### Current gap

Portal request models, permission review plans, desktop resource bridge plans, snapshots, and rollback previews exist. The missing implementation is a stateful safety plane that proves request intent, snapshot receipts, and rollback evidence before execution becomes possible.

### Start from

- `internal/runtime/portal/`
- `internal/runtime/snapshot/`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/desktop_resource_bridge.go`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/desktop_activation_transaction.go`
- `internal/runtime/appidentity/desktop_activation_status.go`

### Deliver

- Fake-mode Portal request records under a controlled state root.
- Snapshot baseline receipts with content-addressed metadata and relative paths.
- Rollback receipts that prove what would be reverted.
- Joined safety summary for Portal review, snapshot availability, rollback availability, and write-gate state.
- Tests that prove real Portal transport remains disabled.

### Keep disabled

- Real Portal D-Bus calls.
- Permission grant commits.
- Snapshotting host paths outside an explicit root.
- Host-root mutation.
- Runtime launch approval.
- Backend launch.

### Verification

```text
go test ./internal/runtime/portal ./internal/runtime/snapshot ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_desktop_resource_bridge_plan.rb
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement A4 from docs/claude-code-next-implementation-assignments.md.

Target outcome:
- Build a fake-mode Portal, snapshot, and rollback safety plane with state-root receipts while keeping real Portal transport and launch disabled.

Hard constraints:
- Do not call real Portal transport, grant permissions, snapshot broad host paths, mutate host roots, enable Runtime writes, or start backends.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add Portal-denied, snapshot-success, snapshot-mismatch, rollback-preview, and real-transport-disabled tests.
- Run the A4 verification commands.
```

## A5: KDE Activation Materialization

### Current gap

KDE activation, entrypoint, desktop-entry, MIME, Dolphin service-menu, tray, notification, settings, and Compatibility Center models exist. The missing implementation is stronger staged materialization and receipt consumption under explicit roots.

### Start from

- `internal/runtime/activation/`
- `internal/runtime/appidentity/desktop_activation_*.go`
- `internal/runtime/appidentity/kde_*`
- `internal/runtime/appidentity/window_identity_routes.go`
- `kde/`
- `lib/xnix/compatibility/desktop_activation_*`
- `lib/xnix/compatibility/kde_*`

### Deliver

- Target-root-only desktop materialization for desktop files, MIME files, Dolphin service menus, icon metadata, and activation manifests.
- Receipt readers that let KDE-safe read models reflect staged or rolled-back state.
- Drift checks between planned activation artifacts and materialized receipts.
- KDE Compatibility Center summaries that consume receipts instead of restating static previews.

### Keep disabled

- Host desktop database mutation.
- Default MIME association writes on the host.
- System service installation.
- Notification delivery.
- KWin rule application.
- Live tray bridging.
- Backend launch.

### Verification

```text
go test ./internal/runtime/activation ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby -Ilib test/test_desktop_entry.rb
ruby -Ilib test/test_dolphin_service_menu.rb
ruby -Ilib test/test_file_association_model.rb
ruby -Ilib test/test_kde_center_model.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement A5 from docs/claude-code-next-implementation-assignments.md.

Target outcome:
- Strengthen KDE activation materialization under explicit target roots and make KDE-safe reads consume activation receipts.

Hard constraints:
- Do not mutate the host desktop database, write host MIME defaults, install services, deliver notifications, apply KWin rules, enable tray bridging, start backends, or mutate host roots.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add target-root, rollback, receipt-consumption, drift-detection, and host-mutation-blocked tests.
- Run the A5 verification commands.
```

## A6: Execution Transaction Ledger

### Current gap

Launch intent, execution request, review, decision, preflight, resource grant, transaction, session, and session status previews exist. The missing implementation is a state-root ledger that can record reviewed, blocked launch transactions without creating requests, granting resources, or launching backends.

### Start from

- `internal/runtime/execution/`
- `internal/runtime/appidentity/execution_*.go`
- `internal/runtime/appidentity/launch_intent.go`
- `internal/runtime/appidentity/runtime_write_gate.go`
- `cmd/xnix-runtime-go/execution_ledger_commands.go`

### Deliver

- State-root execution ledger records with reviewed decision, preflight state, resource grant state, transaction state, and write-gate result.
- Inspect/list commands for ledger records.
- Deterministic blocked states when trust, Portal, snapshot, environment, or write gates are incomplete.
- KDE-safe transaction summaries that do not expose backend implementation details.

### Keep disabled

- Runtime launch approval.
- Request object creation.
- Permission grant commits.
- Backend process starts.
- Session registration.
- Host-root mutation.

### Verification

```text
go test ./internal/runtime/execution ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_launch_request.rb
ruby -Ilib test/test_compatibility_run_plan.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby -Ilib test/test_runtime_write_gate.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement A6 from docs/claude-code-next-implementation-assignments.md.

Target outcome:
- Add a state-root execution transaction ledger that records reviewed and preflighted launch intent while keeping launch blocked.

Hard constraints:
- Do not approve Runtime launch, create real request objects, grant permissions, start backend processes, register live sessions, enable write methods, or mutate host roots.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add approved/deferred/rejected, missing-gate, ledger-list, ledger-inspect, and launch-disabled tests.
- Run the A6 verification commands.
```

## A7: Diagnostics, Repair, and AI Boundary

### Current gap

Diagnostics, diagnostic history, AI diagnostic input, AI recommendation, AI repair approval gates, repair plans, and test results exist. The missing implementation is broader fixture-driven diagnostic history and review-only repair workflows with provider policy gates.

### Start from

- `internal/runtime/diagnostics/`
- `internal/runtime/appidentity/diagnostics.go`
- `internal/runtime/appidentity/diagnostic_history.go`
- `internal/runtime/appidentity/ai_diagnostics.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/test_plan.go`
- `internal/runtime/appidentity/test_result.go`
- `cmd/xnix-runtime-go/diagnostic_record_commands.go`
- `cmd/xnix-runtime-go/ai_diagnostics_commands.go`

### Deliver

- Fixture diagnostic runners for success, failure, warning, and blocked states.
- Diagnostic history records under a controlled state root.
- Privacy-filtered AI diagnostic input that excludes file contents, host paths, secrets, and raw backend data.
- Review-only repair recommendations with approval receipts.
- Provider configuration policy that keeps live provider calls disabled by default.

### Keep disabled

- Real AI provider calls.
- File content reads by default.
- Auto-repair.
- Backend launch.
- Host-root mutation.
- Secret or private data exposure.

### Verification

```text
go test ./internal/runtime/diagnostics ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_ai_repair_approval_gate.rb
ruby -Ilib test/test_compatibility_repair_plan.rb
ruby -Ilib test/test_compatibility_test_plan.rb
ruby -Ilib test/test_compatibility_test_result.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement A7 from docs/claude-code-next-implementation-assignments.md.

Target outcome:
- Extend fixture-driven diagnostics, diagnostic history, provider-gated AI input, and review-only repair recommendations without live provider calls or auto-repair.

Hard constraints:
- Do not call real AI providers, read file contents by default, auto-repair, start backends, mutate host roots, or expose secrets, host paths, file contents, or raw backend data.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add diagnostic success/failure/warning/blocked, privacy-filter, provider-disabled, approval-required, and no-auto-repair tests.
- Run the A7 verification commands.
```

## A8: Developer Evidence and Drift Harness

### Current gap

Contract drift and implementation evidence reports exist. The missing implementation is stricter evidence regression detection so new contracts cannot quietly remain preview-only forever.

### Start from

- `scripts/runtime_contract_drift_report.rb`
- `scripts/implementation_evidence_report.rb`
- `test/test_runtime_contract_drift_report.rb`
- `test/test_implementation_evidence_report.rb`
- `docs/claude-code-domain-dispatch.md`
- `docs/claude-code-empty-domain-implementation-packages.md`

### Deliver

- Report rules that classify contract-only, fixture-implemented, state-root-implemented, smoke-owned, and production-gated evidence.
- Orphan-contract detection for D-Bus methods, Go owner routes, CLI commands, smoke adapter methods, D-Bus client methods, and package docs.
- JSON and Markdown report output for Claude Code handoff reviews.
- CI-friendly nonzero exits for regressions.
- Guidance that maps failed evidence gates back to assignment IDs in this file.

### Keep disabled

- Docker, QEMU, network, backend launch, privileged containers, and host-root mutation as report requirements.

### Verification

```text
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/runtime_contract_drift_report.rb --format markdown
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/implementation_evidence_report.rb --format markdown
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby -Ilib test/test_implementation_evidence_report.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement A8 from docs/claude-code-next-implementation-assignments.md.

Target outcome:
- Strengthen the developer evidence and drift harness so orphan contracts and preview-only regressions fail with actionable assignment guidance.

Hard constraints:
- Do not require Docker, QEMU, network, privileged containers, backend launch, production D-Bus ownership, or host-root mutation for the reports.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add JSON, Markdown, passing, failing, orphan-contract, and assignment-guidance tests.
- Run the A8 verification commands.
```

## A9: Atomic KDE Image and QEMU Acceptance

### Current gap

The low-level Buildroot/QEMU baseline exists, and the KDE-first image direction is documented. The missing implementation is repeated product-level acceptance that proves the atomic KDE image includes the Runtime presence, KDE entry points, and safe service binding without weakening host safety.

### Start from

- `image/kinoite/`
- `scripts/build_kde_image.rb`
- `scripts/build_kde_disk.rb`
- `scripts/boot_kde_image.rb`
- `scripts/kde_first_presence_smoke.rb`
- `docs/kde-image-pipeline.md`
- `docs/kde-first-presence-smoke-spec.md`

### Deliver

- Constrained image build metadata that records exact inputs and output manifest.
- QEMU smoke that checks KDE Runtime presence, service binding, Compatibility Center availability, and disabled unsafe operations.
- Persisted smoke logs under controlled output roots.
- Clear skip/block reporting when the local machine lacks safe resources.

### Keep disabled

- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Host-root mutation.
- Production backend launch.

### Verification

```text
go test ./internal/runtime/image ./cmd/xnix-runtime-go
ruby -Ilib test/test_kde_image.rb
ruby -Ilib test/test_kde_disk.rb
ruby -Ilib test/test_kde_first_presence_smoke_script.rb
ruby scripts/kde_first_presence_smoke.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement A9 from docs/claude-code-next-implementation-assignments.md.

Target outcome:
- Add constrained atomic KDE image and QEMU acceptance evidence for Runtime presence and disabled unsafe operations.

Hard constraints:
- Do not use privileged containers, host networking, Docker socket mounts, broad host mounts, host-root mutation, or production backend launch.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add build-manifest, QEMU-log, presence-smoke, unsafe-operation-disabled, and resource-blocked tests.
- Run the A9 verification commands.
```

## Completion Template

Every Claude Code handoff should end with this exact shape:

```text
Assignment implemented:
- <A1-A9 and short title>

Converted contracts:
- <contract or preview surfaces touched>

Implementation evidence added:
- <fixture/state-root/smoke/report evidence>

Still gated:
- <unsafe production behavior that remains disabled>

Verification run:
- <commands and results>

Known follow-up:
- <small next branch, if any>
```

If the branch cannot provide at least one durable evidence type, it should not claim implementation completion.
