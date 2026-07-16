# Claude Code Contract-to-Implementation Task Board

> Last updated: 2026-07-16 | Baseline: v0.2.253

This document is a practical task board for handing large, relatively independent Xnix implementation domains to Claude Code.

Use it when a domain already has contracts, previews, D-Bus methods, CLI routes, smoke scripts, or tests, but still lacks durable implementation evidence. The goal is to give Claude Code one bounded branch at a time without letting the work expand into unrelated side quests.

## How This File Differs From the Other Handoff Docs

- `docs/claude-code-domain-dispatch.md` is the broad dispatch board.
- `docs/claude-code-mainline-implementation-plan.md` gives the mainline sequence.
- `docs/claude-code-next-implementation-assignments.md` is the short next-assignment board.
- This file is the "contract-to-implementation" board: it focuses on empty or thin domains and states exactly what implementation evidence must be produced.

Do not modify `docs/claude-code-implementation-packages.md` when using this board unless explicitly asked.

## Rules for Every Claude Code Branch

Every branch must:

- Implement exactly one task from this board.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Compatibility Runtime product behavior.
- Use C only for low-level D-Bus transport, ABI-shaped records, or already-owned Runtime policy surfaces.
- Use Ruby for tests, smoke scripts, reports, and developer tooling.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Add success, failure, and blocked-state tests.
- Run the task-specific verification commands plus `ruby scripts/verify_layout.rb`.
- Run `ruby scripts/runtime_contract_drift_report.rb --format json` when touching D-Bus XML, Runtime route manifests, owner dispatch, smoke adapters, D-Bus clients, or Runtime CLI route commands.
- Stop and report if the task would require unsafe behavior outside the selected package.

Unsafe behavior remains disabled unless a future task explicitly enables a constrained test-only path:

- Production D-Bus ownership.
- Runtime write methods.
- Real Wine, Proton, VM, or compatibility backend launch.
- Real Portal transport calls.
- Default network artifact fetch.
- Host package-manager calls.
- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Host-root mutation.
- Raw backend command, executable path, storage path, profile name, host path, file content, secret, token, or private-key exposure in KDE-facing output.

## What Counts as Implementation Evidence

Claude Code should not treat a new contract, preview field, D-Bus method, CLI route, or documentation paragraph as implementation evidence by itself.

A mergeable branch must add at least one of these:

- A state-root record with schema version, app id, operation id, created timestamp, result state, blocked reasons, and relative receipt paths.
- A fixture pipeline with deterministic success and failure paths.
- A constrained smoke path that exercises a real adapter boundary.
- A report gate that fails on contract drift or orphan contract-only surfaces.
- A KDE-safe read model that consumes durable receipts instead of restating static previews.

If a branch only adds more preview text, ask Claude Code to stop and narrow the task. The repository already has enough contracts; the missing ingredient is evidence that can fail.

## Recommended Assignment Wave

Start with T1, T2, T3, and T8.

| Task | Domain | Why first | Minimal mergeable result |
| --- | --- | --- | --- |
| T1 | Runtime owner process boundary | Other domains need a real owner boundary instead of preview-only dispatch. | A constrained session-bus owner serves read-only routes and fails writes closed. |
| T2 | Recipe, artifact, cache, and install trust | Install, activation, lifecycle, and execution all depend on trusted inputs. | Local recipes and fixture artifacts verify digests, stage under controlled roots, and produce blocked receipts. |
| T3 | Runtime state root and backend lifecycle | Execution readiness needs durable state before launch can be meaningful. | Lifecycle records persist under a controlled state root. |
| T8 | Developer evidence and drift harness | Contracts are multiplying quickly. | Reports fail on orphan contracts and preview-only regressions. |

T4 and T5 can follow once T2 and T3 have enough state evidence. T6 should wait until T2, T3, and T4 can block unsafe launch paths. T9 should wait until the image smoke can validate real Runtime evidence rather than only previews.

## T1: Runtime Owner Process Boundary

### Current gap

The Runtime owner has route manifests, in-process Go read dispatch, owner smoke-batch records, and C D-Bus smoke bridge evidence. The missing piece is a constrained long-running Go owner process that can claim a private test session-bus name and serve reads through the same owner table.

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

### Copyable Claude Code prompt

```text
Implement T1 from docs/claude-code-contract-to-implementation-task-board.md.

Target outcome:
- Add a constrained Go Runtime smoke-owner process that owns a private test session-bus name, serves read-only Runtime methods through the existing owner route table, and fails every write method closed.

Hard constraints:
- Do not enable production D-Bus ownership, system service installation, write methods, backend launch, network fetch, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add success, unsupported-read, disabled-write, lifecycle, and blocked-production-owner tests.
- Run the T1 verification commands.
```

## T2: Recipe, Artifact, Cache, and Install Trust

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

### Copyable Claude Code prompt

```text
Implement T2 from docs/claude-code-contract-to-implementation-task-board.md.

Target outcome:
- Build a fixture-first trust, artifact, cache, staging, and install-readiness pipeline that fails closed and never mutates the host root.

Hard constraints:
- Do not add default network fetch, package-manager calls, private keys, production trust bypasses, desktop writes, backend launch, privileged containers, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add digest success, digest mismatch, unsigned/development-only, cache-boundary, and staging-blocked tests.
- Run the T2 verification commands.
```

## T3: Runtime State Root and Backend Lifecycle

### Current gap

Backend selection, binding, capability, environment, lifecycle, run plan, and execution readiness surfaces exist. The missing implementation is durable Runtime state that changes when trust, artifact staging, Portal, snapshot, and readiness evidence changes.

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

### Copyable Claude Code prompt

```text
Implement T3 from docs/claude-code-contract-to-implementation-task-board.md.

Target outcome:
- Add Runtime state-root lifecycle records and readiness transitions for managed compatibility environments without starting any backend.

Hard constraints:
- Do not start Wine, Proton, VM, or any backend process. Do not provision host-level state, expose raw backend details, mutate the host root, or approve launches.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add missing, staged, ready, repair-required, blocked, and root-escape tests.
- Run the T3 verification commands.
```

## T4: Portal, Snapshot, and Rollback Safety

### Current gap

Portal access policies, Portal request models, snapshot plans, repair plans, and rollback receipts exist. The missing implementation is durable fake-first permission and snapshot state that can later gate execution.

### Start from

- `internal/runtime/portal/`
- `internal/runtime/snapshot/`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/portal_request.go`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/repair_plan.go`
- `lib/xnix/compatibility/desktop_activation_rollback.rb`

### Deliver

- Portal request records with states such as `planned`, `requested`, `granted`, `denied`, `cancelled`, `expired`, and `failed`.
- Fake Portal transport for tests and a disabled real transport boundary.
- Snapshot creation, verification, rollback, and receipt integration under controlled roots.
- A joined safety view showing which permissions and restore points block execution.

### Keep disabled

- Real Portal transport calls.
- System snapshot tools.
- Host-root mutation.
- Permission grants without fake broker approval.
- Launch approval.

### Verification

```text
go test ./internal/runtime/portal ./internal/runtime/snapshot ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_compatibility_repair_plan.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement T4 from docs/claude-code-contract-to-implementation-task-board.md.

Target outcome:
- Add durable fake-first Portal request state and controlled snapshot/rollback receipts that can gate future execution.

Hard constraints:
- Do not call the real Portal transport, system snapshot tools, host-root paths, or backend launch paths. Keep permission grants fake-broker-only.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add granted, denied, cancelled, expired, snapshot-verify, rollback, and invalid-path tests.
- Run the T4 verification commands.
```

## T5: KDE Activation Materialization

### Current gap

Desktop activation, desktop entries, MIME associations, Dolphin service menus, tray, notifications, settings, and Compatibility Center models exist. The missing implementation is stronger materialization evidence under explicit target roots and read models that consume receipts rather than static previews.

### Start from

- `internal/runtime/activation/`
- `internal/runtime/appidentity/desktop_activation_*.go`
- `internal/runtime/appidentity/desktop_entry.go`
- `internal/runtime/appidentity/file_association.go`
- `internal/runtime/appidentity/notification.go`
- `internal/runtime/appidentity/tray_status.go`
- `scripts/install_runtime_activation.rb`
- `scripts/runtime_activation_smoke.rb`

### Deliver

- Target-root-only activation staging receipts for desktop files, MIME associations, service menus, manifests, tray hints, and rollback receipts.
- Digest verification before rollback removes staged files.
- KDE-safe status reads that consume staging or activation receipts.
- Explicit blocked production activation when trust or root requirements are not met.

### Keep disabled

- Host-root writes.
- KDE cache refresh on the host.
- Launch or backend start.
- Settings persistence outside explicit test roots.
- Raw backend command or path exposure.

### Verification

```text
go test ./internal/runtime/activation ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby scripts/runtime_activation_smoke.rb
ruby -Ilib test/test_runtime_activation.rb
ruby -Ilib test/test_runtime_activation_install.rb
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement T5 from docs/claude-code-contract-to-implementation-task-board.md.

Target outcome:
- Add stronger KDE activation materialization receipts under explicit target roots and make KDE-safe read models consume those receipts.

Hard constraints:
- Do not write to the host root, refresh host KDE caches, launch applications, start backends, persist settings outside test roots, or expose raw backend details.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add staging success, rollback, digest mismatch, blocked production, and root-boundary tests.
- Run the T5 verification commands.
```

## T6: Execution Transaction Ledger

### Current gap

Launch intent, execution request, review, preflight, resource grant, transaction, session, and session status models exist. The missing implementation is a reviewed transaction ledger that joins trust, lifecycle, Portal, and snapshot gates while still blocking real launch.

### Start from

- `internal/runtime/execution/`
- `internal/runtime/appidentity/launch_intent.go`
- `internal/runtime/appidentity/execution_request.go`
- `internal/runtime/appidentity/execution_review.go`
- `internal/runtime/appidentity/execution_preflight.go`
- `internal/runtime/appidentity/execution_resource_grant.go`
- `internal/runtime/appidentity/execution_transaction.go`
- `internal/runtime/appidentity/execution_session_status.go`

### Deliver

- Execution transaction records under a controlled state root.
- Join checks for recipe trust, artifact staging, environment readiness, Portal permission, snapshot preconditions, and write gates.
- Reviewed decisions that can be `approved-for-preflight`, `blocked`, or `cancelled`.
- Stable session status records that explicitly report no backend launch.

### Keep disabled

- Backend process launch.
- Real execution request dispatch.
- Portal or snapshot bypass.
- Write-method enablement.
- Host-root mutation.

### Verification

```text
go test ./internal/runtime/execution ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_launch_request.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby -Ilib test/test_compatibility_run_plan.rb
ruby -Ilib test/test_compatibility_test_plan.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement T6 from docs/claude-code-contract-to-implementation-task-board.md.

Target outcome:
- Add a reviewed execution transaction ledger that joins safety prerequisites and still blocks real backend launch.

Hard constraints:
- Do not launch backends, enable write methods, bypass Portal or snapshot gates, dispatch real execution, mutate the host root, or expose backend details.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add approved-for-preflight, blocked-trust, blocked-portal, blocked-snapshot, cancelled, and no-backend-launch tests.
- Run the T6 verification commands.
```

## T7: Diagnostics, Repair, and AI Boundary

### Current gap

Diagnostic input, diagnostic recommendations, AI repair approval gates, test plans, test results, repair plans, action queues, and review receipts exist. The missing implementation is a fixture-driven diagnostic and repair pipeline with a provider-disabled AI boundary.

### Start from

- `internal/runtime/diagnostics/`
- `internal/runtime/appidentity/ai_diagnostic_input.go`
- `internal/runtime/appidentity/ai_diagnostic_recommendation.go`
- `internal/runtime/appidentity/ai_repair_approval_gate.go`
- `internal/runtime/appidentity/test_plan.go`
- `internal/runtime/appidentity/test_result.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/kde_action_*.go`

### Deliver

- Fixture diagnostic runs with persisted results under a controlled state root.
- Repair recommendations derived from fixture signals.
- AI provider interface with disabled production default and fake provider tests.
- Approval receipts that keep repair execution blocked until explicit review.
- Compatibility Center action queues that consume diagnostic history.

### Keep disabled

- Real AI provider calls by default.
- File-content reads.
- File-path exposure.
- Automatic repair execution.
- Backend launch.
- Network access.

### Verification

```text
go test ./internal/runtime/diagnostics ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_ai_repair_approval_gate.rb
ruby -Ilib test/test_compatibility_test_plan.rb
ruby -Ilib test/test_compatibility_test_result.rb
ruby -Ilib test/test_compatibility_repair_plan.rb
ruby -Ilib test/test_compatibility_action_queue.rb
ruby -Ilib test/test_action_review_receipt.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement T7 from docs/claude-code-contract-to-implementation-task-board.md.

Target outcome:
- Add fixture-driven diagnostic run records, repair recommendations, review receipts, and a provider-disabled AI boundary.

Hard constraints:
- Do not call real AI providers by default, read file contents, expose paths, auto-repair, launch backends, use network access, or mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add fixture success, fixture failure, provider-disabled, fake-provider, approval-required, and no-auto-repair tests.
- Run the T7 verification commands.
```

## T8: Developer Evidence and Drift Harness

### Current gap

Contract drift and implementation evidence reports exist. The missing implementation is stronger failure logic for orphan contracts, preview-only expansion, stale handoff docs, and missing package evidence.

### Start from

- `scripts/runtime_contract_drift_report.rb`
- `scripts/implementation_evidence_report.rb`
- `test/test_runtime_contract_drift_report.rb`
- `test/test_implementation_evidence_report.rb`
- `docs/claude-code-mainline-implementation-plan.md`
- `docs/claude-code-domain-dispatch.md`

### Deliver

- JSON and Markdown evidence reports that classify every large domain.
- Failure modes for orphan Runtime read methods, missing smoke coverage, missing owner route entries, and preview-only domains.
- A package ownership map that links docs, tests, source files, and evidence level.
- CI-friendly exit codes without Docker or QEMU.

### Keep disabled

- Docker, QEMU, network, backend launch, production ownership, and host-root mutation during report generation.
- Broad filesystem writes outside report output paths explicitly requested by the caller.

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

### Copyable Claude Code prompt

```text
Implement T8 from docs/claude-code-contract-to-implementation-task-board.md.

Target outcome:
- Strengthen the developer evidence and drift harness so orphan contracts and preview-only expansions fail locally without Docker or QEMU.

Hard constraints:
- Do not run Docker, QEMU, network fetches, backend launch, production D-Bus ownership, or host-root mutation from the reports.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add orphan-contract, missing-smoke, missing-owner-route, preview-only, and markdown/json report tests.
- Run the T8 verification commands.
```

## T9: Atomic KDE Image and QEMU Acceptance

### Current gap

KDE image manifests, Buildroot/QEMU baseline tests, and smoke scripts exist. The missing implementation is an integrated, reproducible acceptance path that validates the KDE-first Runtime evidence without weakening host safety.

### Start from

- `scripts/build_kde_image.rb`
- `scripts/build_kde_disk.rb`
- `scripts/boot_kde_image.rb`
- `scripts/full_smoke.rb`
- `docs/kde-image-pipeline.md`
- `docs/kde-first-compatibility-acceptance.md`
- `test/test_kde_image.rb`
- `test/test_qemu.rb`
- `test/test_full_smoke_script.rb`

### Deliver

- A reproducible image or image-manifest path that records exact inputs and expected Runtime evidence.
- A constrained QEMU/KDE smoke plan that verifies boot logs, Runtime presence, D-Bus smoke availability, and KDE-first entry-point readiness when feasible.
- JSON and Markdown smoke reports.
- Clear separation between product image validation and host-machine state.

### Keep disabled

- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host mounts.
- Host-root mutation.
- Production backend launch.
- Host package-manager calls.

### Verification

```text
ruby scripts/full_smoke.rb
ruby -Ilib test/test_kde_image.rb
ruby -Ilib test/test_kde_disk.rb
ruby -Ilib test/test_qemu.rb
ruby -Ilib test/test_full_smoke_script.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement T9 from docs/claude-code-contract-to-implementation-task-board.md.

Target outcome:
- Add a reproducible atomic KDE image and QEMU acceptance path that validates real Runtime evidence while preserving host safety.

Hard constraints:
- Do not use privileged containers, host networking, Docker socket mounts, broad host mounts, host-root mutation, host package-manager calls, or production backend launch.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add manifest, missing-input, smoke-report, safe-container, and no-host-mutation tests.
- Run the T9 verification commands.
```

## Completion Template

Ask Claude Code to finish each branch with this exact shape:

```text
Converted contracts:
- <contract or preview surfaces changed>

Implementation evidence added:
- <fixture/state-root/smoke/report evidence>

Still gated:
- <unsafe production behavior that remains disabled>

Verification run:
- <commands>

Files changed:
- <short list>
```

This makes review concrete and keeps broad branches from hiding behind "foundation" language.
