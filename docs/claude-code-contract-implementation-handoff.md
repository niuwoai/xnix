# Claude Code Contract Implementation Handoff

> Last updated: 2026-07-16 | Baseline: v0.2.253

This document splits Xnix's contract-heavy roadmap into coarse, independently implementable workstreams for Claude Code.

Use this when a domain already has contracts, previews, smoke tests, or documentation, but still lacks durable implementation. Each package below should become one Claude Code branch unless the package explicitly says otherwise.

This document complements:

- `docs/claude-code-open-domain-work-packages.md` for the broad backlog map.
- `docs/claude-code-empty-domain-implementation-packages.md` for large empty-domain packages where contracts need durable implementation evidence.
- `docs/claude-code-independent-implementation-briefs.md` for smaller prompt-ready briefs.
- `docs/claude-code-implementation-packages.md` for the original implementation package guide.

## Global Handoff Rules

- Use one branch per package.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Runtime product logic.
- Use C only for low-level Runtime surfaces, ABI-shaped contracts, or already-owned C records.
- Use Ruby for tests, smoke scripts, and developer tooling.
- Do not introduce production D-Bus ownership, real backend launch, privileged containers, host networking, Docker socket mounts, broad host-directory mounts, or host-root mutation unless the package explicitly enables that exact action.
- Do not expose raw backend commands, raw executable paths, compatibility storage paths, profile terminology, file contents, secrets, tokens, or private keys in normal KDE/user-facing output.
- Keep write methods disabled unless the package explicitly implements a gated write path.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Run the package tests and `ruby scripts/verify_layout.rb` before handoff.
- If the version lands on every tenth code version, run the full build and QEMU smoke gate.

## Package Map

| Package | Domain | Independence level | Main implementation output |
| --- | --- | --- | --- |
| H1 | Runtime owner service | High | Long-running Go owner candidate and session-bus smoke owner |
| H2 | Recipe trust and artifact pipeline | Medium | Production-shaped trust, acquisition, cache, and staging pipeline |
| H3 | Environment lifecycle and state root | High | Runtime-owned lifecycle state machine with safe persistence |
| H4 | Portal and snapshot control plane | Medium | Durable Portal request broker plus snapshot/rollback workflow |
| H5 | KDE materialization and activation | Medium | Gated desktop, MIME, tray, notification, and rollback writers |
| H6 | KDE live shell adapters | Medium | Minimal runnable KRunner, Dolphin, tray, task, and KWin integration adapters |
| H7 | Execution transaction pipeline | High | Gated launch/session transaction records without backend launch by default |
| H8 | Test, repair, and AI diagnostics | High | Fixture-driven compatibility test and repair recommendation pipeline |
| H9 | Atomic KDE image and QEMU acceptance | Medium | Reproducible desktop image path and milestone smoke gate |
| H10 | Developer quality gates | High | Contract drift, route parity, and report automation |

## H1: Runtime Owner Service

### Mission

Replace the current preview-and-adapter ownership model with a constrained Go Runtime owner process that can serve read-only Runtime methods on a session bus in smoke mode.

### Existing contracts

- `runtime/dbus/org.xnix.Compatibility1.xml`
- Runtime service binding previews.
- Runtime live owner gate previews.
- Runtime owner process readiness previews.
- Runtime owner smoke plan previews.
- Runtime method parity manifests.
- Runtime owner route manifests.
- Runtime write gates.

### Deliverables

- `cmd/xnix-runtime-owner/` as the owner entry point.
- `internal/runtime/owner/` as the owner package.
- A smoke-owner mode that may claim a constrained session-bus name inside container tests.
- Read-only method routing to existing Go Runtime preview handlers.
- Deterministic disabled errors for write methods.
- Owner readiness output that distinguishes `preview-only`, `smoke-owner`, and `production-owner`.
- Logs for startup, route-table version, bus-claim mode, and shutdown reason without backend-detail leakage.

### Explicit non-goals

- Production bus ownership.
- System service installation.
- Real backend launch.
- Host root writes.
- KDE-owned Runtime policy.

### Acceptance

- A constrained container can start the owner on a session bus.
- Every read-only method either returns the existing safe Go preview shape or an explicit unsupported-route readiness marker.
- Every write method fails closed with a stable `WriteMethodDisabled` error.
- Runtime owner readiness reports smoke-owner progress while production ownership remains gated.

### Required tests

```text
go test ./...
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
ruby scripts/verify_layout.rb
```

## H2: Recipe Trust and Artifact Pipeline

### Mission

Turn recipe trust, package acquisition, artifact verification, cache planning, and staging into a production-shaped pipeline while keeping network fetch and host mutation disabled by default.

### Existing contracts

- Recipe registry validation.
- Recipe trust policy previews.
- Recipe install gates.
- Package source previews.
- Acquisition preflight previews.
- Artifact manifest previews.
- Compatibility install previews.

### Deliverables

- `internal/runtime/recipe/` for read-only recipe storage and trust evaluation.
- `internal/runtime/artifact/` for artifact manifests, digests, local fixture acquisition, cache receipts, and staging receipts.
- Trust states such as `production-trusted`, `development-only`, `unsigned`, `invalid`, and `blocked`.
- Fixture-only acquisition mode.
- Runtime state-root scoped cache and staging namespaces.
- Receipts that can feed install gates, owner readiness, and Compatibility Center pages.

### Explicit non-goals

- Remote artifact downloads by default.
- Production signing keys.
- Host package-manager calls.
- Desktop activation writes.
- Backend process starts.

### Acceptance

- Invalid recipe digests fail closed.
- Invalid or missing production signatures cannot pass production readiness.
- Digest mismatch blocks artifact staging.
- Cache and staging paths remain under a configured Runtime/test root.
- Development fixtures remain usable but visibly non-production.

### Required tests

```text
go test ./...
ruby -Ilib test/test_recipe_registry.rb
ruby -Ilib test/test_recipe_trust_policy.rb
ruby -Ilib test/test_recipe_install_gate.rb
ruby -Ilib test/test_compatibility_package_source.rb
ruby -Ilib test/test_compatibility_acquisition_preflight.rb
ruby -Ilib test/test_compatibility_artifact_manifest.rb
ruby -Ilib test/test_compatibility_install_plan.rb
ruby scripts/verify_layout.rb
```

## H3: Environment Lifecycle and State Root

### Mission

Convert static backend/environment plans into a Runtime-owned lifecycle state machine with safe persistence under a controlled state root.

### Existing contracts

- Backend selection plans.
- Backend environment plans.
- Backend binding plans.
- Backend capability matrix previews.
- Backend lifecycle previews.
- Execution readiness previews.
- Application state-root previews.

### Deliverables

- `internal/runtime/environment/` with lifecycle states such as `missing`, `planned`, `staged`, `ready`, `repair-required`, and `blocked`.
- Persistent lifecycle records under a configured Runtime state root.
- State transitions driven by recipe trust, artifact staging, Portal permission state, and snapshot readiness.
- Safe readiness explanations for KDE-facing consumers.
- A test fixture store that can simulate installed, missing, broken, and repair-required environments.

### Explicit non-goals

- Starting Wine, Proton, VM, or other compatibility backend processes.
- Exposing shell commands, profile names, backend executable paths, or host storage paths.
- Creating real user home directories or host-level state.

### Acceptance

- Lifecycle state survives inside a test-controlled state root.
- Readiness changes when trust, staging, permission, or snapshot inputs change.
- KDE-facing output remains implementation-detail safe.
- Execution stays disabled until H7 owns launch transactions.

### Required tests

```text
go test ./...
ruby -Ilib test/test_compatibility_backend_environment_plan.rb
ruby -Ilib test/test_compatibility_backend_binding.rb
ruby -Ilib test/test_compatibility_backend_lifecycle.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/verify_layout.rb
```

## H4: Portal and Snapshot Control Plane

### Mission

Implement durable permission request state and integrate Runtime snapshot/rollback records into the safety model.

### Existing contracts

- Portal access policy previews.
- Portal request models.
- Fake-mode Runtime Portal broker.
- Snapshot plan previews.
- Existing constrained snapshot store.
- Repair plan and rollback-related receipts.

### Deliverables

- Durable Portal request records under a controlled Runtime state root.
- Request states such as `planned`, `requested`, `granted`, `denied`, `cancelled`, `expired`, and `failed`.
- Fake transport for tests and disabled real transport boundary.
- Snapshot creation, verification, rollback, and receipt integration for managed Runtime state.
- A joined safety view that explains which permissions and restore points are required before execution.

### Explicit non-goals

- Real XDG Desktop Portal calls by default.
- Direct access to user documents.
- System snapshots.
- Host root rollback.
- Backend launch.

### Acceptance

- Fake Portal requests can be created, completed, denied, cancelled, expired, and failed deterministically.
- Snapshot operations stay inside the configured Runtime state root.
- Snapshot identifiers and rollback receipts are validated before use.
- Execution readiness can consume permission and snapshot state without starting execution.

### Required tests

```text
go test ./...
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_runtime_snapshot_plan.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/verify_layout.rb
```

## H5: KDE Materialization and Activation

### Mission

Move KDE activation from read-only previews to gated file materialization with auditable receipts and rollback support.

### Existing contracts

- Desktop identity plans.
- Desktop entry previews.
- MIME association previews.
- Desktop activation bundle previews.
- Desktop activation preflight, staging, transaction, status, and manifest previews.
- Desktop integration installer and rollback scripts.
- Compatibility action queue and action review receipts.

### Deliverables

- Gated writers for desktop entries, Dolphin service menus, MIME defaults, tray metadata, notification metadata, and activation manifests under a staging root.
- Activation receipts with relative paths, SHA-256 digests, file modes, activated entry-point IDs, and rollback eligibility.
- Rollback validation that removes only unchanged Runtime-owned files.
- A clear production/development mode split.
- Compatibility Center action cards that explain materialization readiness and blocked reasons.

### Explicit non-goals

- Writing into the host root by default.
- Overwriting existing user MIME defaults without explicit approval state.
- Launching applications.
- Sending real notifications by default.
- Applying KWin rules or task-manager activation directly.

### Acceptance

- Development staging can materialize files under a supplied test root.
- Production mode blocks development-only recipes.
- Rollback refuses to remove files with unexpected digest changes.
- All materialized files come from Runtime-rendered sources or audited fixtures.

### Required tests

```text
go test ./...
ruby -Ilib test/test_desktop_entry_plan.rb
ruby -Ilib test/test_file_association_plan.rb
ruby -Ilib test/test_desktop_activation_manifest.rb
ruby -Ilib test/test_desktop_activation_transaction.rb
ruby -Ilib test/test_desktop_activation_status.rb
ruby scripts/kde_first_presence_smoke.rb
ruby scripts/verify_layout.rb
```

## H6: KDE Live Shell Adapters

### Mission

Create minimal runnable KDE-side adapters that consume Runtime read models without taking ownership of Runtime policy.

### Existing contracts

- KDE Integration Status.
- KDE Shell Integration Plan.
- KDE Application Surface Plan.
- KRunner Query Plan.
- Dolphin file-open, drag-and-drop, and AI analysis previews.
- Task Manager Identity Plan.
- KWin Window Rule Plan.
- Tray Status Plan.
- Notification Plan.

### Deliverables

- Minimal KRunner adapter that reads Runtime query plans.
- Minimal Dolphin service/action adapter that submits managed Runtime intents.
- Minimal tray/status adapter that renders Runtime status and navigation actions.
- Minimal task/window identity adapter fixtures for KWin and task manager tests.
- Shared D-Bus client usage rather than direct backend calls.
- Smokeable adapters with fixture Runtime responses.

### Explicit non-goals

- Forking KDE, Plasma, KWin, or Dolphin.
- Direct Wine/VM invocation.
- Runtime policy inside KDE packages.
- Direct file reads outside Portal-mediated flows.
- Production user-session autostart by default.

### Acceptance

- KDE adapters can render or submit fixture-backed Runtime-safe actions.
- No adapter exposes backend details or bypasses Runtime gates.
- Missing Runtime service is handled as a recoverable UI state.
- First-release entry points stay aligned with KDE-first presence smoke expectations.

### Required tests

```text
go test ./...
ruby scripts/kde_first_presence_smoke.rb
ruby scripts/container.rb kde-center-dbus-smoke
ruby scripts/verify_layout.rb
```

## H7: Execution Transaction Pipeline

### Mission

Replace launch placeholders with an auditable transaction pipeline that can create, review, approve, deny, and record execution sessions without starting real backends by default.

### Existing contracts

- Launch intent previews.
- Execution request, review, decision, preflight, resource grant, transaction, session, and session-status previews.
- Runtime write gates.
- Portal permission plans.
- Snapshot plans.
- Backend lifecycle and readiness previews.

### Deliverables

- `internal/runtime/execution/` with transaction and session state records.
- Explicit states such as `requested`, `waiting-for-review`, `approved`, `denied`, `preflight-blocked`, `resource-granted`, `ready-to-start`, `start-disabled`, `completed`, and `failed`.
- A fake executor for tests that never launches real backends.
- Review and decision receipts.
- Resource grant linkage to Portal and snapshot packages.
- Stable D-Bus-safe read models for KDE.

### Explicit non-goals

- Real backend process start.
- Network access by default.
- Privileged process control.
- Direct user-file access.
- Silent execution after KDE action review.

### Acceptance

- Launch write attempts remain disabled unless an explicit fake-mode test path is selected.
- Review approval is necessary but not sufficient for execution.
- Preflight can block on recipe trust, artifact staging, Portal permissions, snapshot readiness, and lifecycle state.
- Session-status reads can report fake sessions deterministically.

### Required tests

```text
go test ./...
ruby -Ilib test/test_launch_intent.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby -Ilib test/test_runtime_write_gate.rb
ruby -Ilib test/test_action_review_receipt.rb
ruby scripts/verify_layout.rb
```

## H8: Test, Repair, and AI Diagnostics

### Mission

Turn compatibility testing, repair planning, and AI diagnostics from static previews into fixture-driven records with disabled production provider calls.

### Existing contracts

- Compatibility test plan previews.
- Test result previews.
- Repair plan previews.
- AI diagnostic input previews.
- AI diagnostic recommendation previews.
- AI repair approval gate previews.
- Compatibility Center action queues and review receipts.

### Deliverables

- `internal/runtime/testing/` or equivalent package for compatibility test run records.
- Fixture-driven test executor that records pass/fail/blocked/skipped outcomes without launching real applications.
- Repair recommendation records linked to test failures.
- AI provider interface with a disabled production default and a fixture provider for tests.
- Redaction rules for diagnostic input.
- Approval gate linkage so AI never performs repairs directly.

### Explicit non-goals

- Real AI provider calls by default.
- Reading arbitrary user files.
- Sending file contents to any provider.
- Executing repairs automatically.
- Launching real backends.

### Acceptance

- Test records can be created, listed, and summarized from fixtures.
- Repair plans cite test evidence and remain review-only.
- AI diagnostic input is redacted and path-safe.
- AI recommendation output cannot bypass repair approval gates.

### Required tests

```text
go test ./...
ruby -Ilib test/test_compatibility_test_plan.rb
ruby -Ilib test/test_compatibility_test_result.rb
ruby -Ilib test/test_repair_plan.rb
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_ai_repair_approval_gate.rb
ruby scripts/verify_layout.rb
```

## H9: Atomic KDE Image and QEMU Acceptance

### Mission

Create a reproducible flagship KDE image path while preserving the existing Buildroot/QEMU learning baseline.

### Existing contracts

- Product overview describing Fedora Kinoite-compatible atomic KDE direction.
- Buildroot/QEMU boot baseline.
- Full smoke and constrained Docker/Colima smoke scripts.
- Acceptance baseline requiring QEMU serial logs and SSH availability through loopback-bound forwarding.

### Deliverables

- A documented image build path for the KDE flagship direction.
- A minimal image configuration that includes the Runtime package boundary and KDE integration artifacts.
- A QEMU smoke gate that records serial logs, service status, and loopback-bound SSH checks.
- A failure-report format suitable for CI artifacts.
- Explicit separation between the legacy Buildroot learning baseline and the flagship KDE image.

### Explicit non-goals

- Host bootloader changes.
- Host network exposure.
- Privileged container requirements.
- Broad host-directory mounts.
- Replacing the learning baseline before the KDE path is reproducible.

### Acceptance

- A clean checkout can run the documented smoke path inside the approved constrained environment.
- QEMU logs are persisted for diagnosis.
- SSH is only exposed through loopback-bound forwarding.
- Runtime service status is visible in the smoke report.

### Required tests

```text
ruby scripts/full_smoke.rb
ruby scripts/verify_layout.rb
```

## H10: Developer Quality Gates

### Mission

Make contract drift, route drift, safety regressions, and review handoffs easier to catch before large feature branches land.

### Existing contracts

- Layout verification.
- Method parity manifests.
- KDE-first presence smoke.
- Container smoke scripts.
- Current CHANGELOG and product overview discipline.

### Deliverables

- A route-contract drift checker that compares XML, Go route manifest, C smoke adapter, Ruby D-Bus client, KDE read models, and smoke scripts.
- JSON and Markdown reports for major smoke runs.
- A safety-regression checker for forbidden terms and flags in KDE-facing output.
- A branch handoff checklist generator for Claude Code branches.
- Clear failure messages that identify the exact missing route, unsafe term, or stale version field.

### Explicit non-goals

- Networked CI setup.
- GitHub workflow creation unless explicitly requested.
- Reformatting unrelated files.
- Replacing package-specific tests.

### Acceptance

- A local command can produce JSON and Markdown handoff reports.
- Drift checks fail on missing Runtime routes or stale method parity data.
- Safety checks fail on backend-detail leakage in user-facing outputs.
- Version, changelog, and product overview mismatches are reported precisely.

### Required tests

```text
go test ./...
ruby scripts/verify_layout.rb
ruby scripts/kde_first_presence_smoke.rb
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
```

## Suggested Claude Code Prompt Template

```text
Implement package <H-number and title> from docs/claude-code-contract-implementation-handoff.md.

Scope:
- Work only on this package.
- Keep all project-facing text in English.
- Preserve existing safety boundaries: no privileged containers, host networking, Docker socket mounts, broad host-directory mounts, host-root mutation, backend launch, or enabled write methods unless the package explicitly permits it.
- Prefer Go for durable Runtime logic, C only for low-level/already-owned Runtime surfaces, and Ruby for tests/tooling.
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for every code change.
- Do not modify docs/claude-code-implementation-packages.md unless explicitly asked.

Deliver:
- The deliverables listed for the package.
- Targeted tests listed for the package.
- A short handoff summary with implemented files, test commands, remaining blocked work, and safety notes.
```

## Recommended Parallel Branch Queue

These packages can be assigned in parallel with low merge risk:

- H1 Runtime Owner Service
- H3 Environment Lifecycle and State Root
- H8 Test, Repair, and AI Diagnostics
- H10 Developer Quality Gates

These packages should wait for or coordinate with another package:

- H2 should coordinate with H3 because staging receipts and lifecycle readiness meet at the state root.
- H4 should coordinate with H7 because Portal grants and snapshots become execution preflight inputs.
- H5 should wait for enough of H2 to decide production versus development activation.
- H6 should consume H5 and H7 read models, but can start with fixture adapters.
- H9 should wait until H1 and H5 have enough Runtime/KDE materialization to make image smoke meaningful.
