# Claude Code Domain Dispatch

> Last updated: 2026-07-16 | Baseline: v0.2.235

This document is the dispatch board for assigning large, relatively independent Xnix implementation domains to Claude Code.

Use this when the repository has contracts, previews, or read models for a domain, but not enough durable implementation evidence. This file is intentionally practical: pick one domain, copy its prompt, and keep the branch focused until it produces a reviewable implementation step.

## Relationship to Other Claude Code Documents

- `docs/claude-code-mainline-implementation-plan.md` gives the strategic sequence.
- `docs/claude-code-next-implementation-assignments.md` is the short copyable board for handing the next large implementation packages to Claude Code.
- `docs/claude-code-contract-to-implementation-task-board.md` is the practical task board for converting empty or thin contract domains into durable implementation evidence.
- `docs/claude-code-large-empty-domain-assignments.md` gives branch-sized empty-domain assignments.
- `docs/claude-code-empty-domain-implementation-packages.md` defines evidence levels and completion rules.
- `docs/claude-code-contract-gap-work-packages.md` lists finer contract gaps.
- `docs/claude-code-open-domain-work-packages.md` lists open follow-up areas.
- `docs/claude-code-independent-implementation-briefs.md` contains smaller prompt-ready briefs.
- `docs/windows-app-compatibility-implementation-brief.md` explains the product-level Windows application compatibility direction.

This file does not replace those documents. It is the first page to read when deciding what another agent should implement next.

## Global Branch Rules

Every Claude Code branch must:

- Implement exactly one dispatch domain unless explicitly told otherwise.
- Keep source, comments, fixtures, tests, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior.
- Use C only for low-level transport, ABI-shaped records, or already-owned C Runtime policy surfaces.
- Use Ruby for tests, smoke scripts, reports, and developer tooling.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Add targeted success, failure, and blocked-state tests.
- Run domain-specific tests plus `ruby scripts/verify_layout.rb`.
- Run `ruby scripts/runtime_contract_drift_report.rb --format json` when touching D-Bus XML, Runtime routes, owner dispatch, smoke adapters, or Runtime CLI route commands.
- Stop and report if the work requires unsafe behavior outside the selected domain.

Unsafe behavior remains disabled unless a future dispatch explicitly enables a gated test-only path:

- Production D-Bus ownership.
- Runtime write methods.
- Real Wine, Proton, VM, or backend launch.
- Real Portal transport calls.
- Default network artifact fetch.
- Host package-manager calls.
- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Host-root mutation.
- Raw backend command, executable path, storage path, profile terminology, file contents, secret, token, private key, or host path exposure in KDE-facing output.

## Dispatch Overview

| Domain | Can start now | Primary owner | Depends on | Main result |
| --- | --- | --- | --- | --- |
| D1 Runtime owner process boundary | Yes | Go Runtime | Existing owner dispatch | Private smoke bus owner for read-only methods, writes fail closed |
| D2 Recipe, artifact, cache, and install trust | Yes | Go Runtime | Registry fixtures | Local trust pipeline with digest checks and controlled staging |
| D3 Runtime state root and backend lifecycle | Yes | Go Runtime | D2 for install evidence | Durable lifecycle states without backend launch |
| D4 Portal, permission, snapshot, and rollback safety | Yes | Go Runtime | D3 for state-root integration | Review-first permission and snapshot safety plane |
| D5 KDE activation and seven entry points | Yes, read/stage only | Go Runtime plus KDE integration | D2 for trusted inputs | Desktop entries, MIME, tray, notification, settings, and receipts under explicit roots |
| D6 Execution transaction pipeline | Later | Go Runtime | D2, D3, D4 | Reviewed transaction ledger that still blocks launch |
| D7 Diagnostics, repair, and AI boundary | Yes, no-provider mode | Go Runtime | D3 useful but not required for fixtures | Fixture diagnostics and review-only repair recommendations |
| D8 Developer evidence and drift harness | Yes | Ruby tooling plus Go fixtures | None | Reports fail on orphan contracts and preview-only regressions |
| D9 Atomic KDE image and QEMU acceptance | Later | Build/image tooling | D1 through D5 enough to validate product evidence | Constrained image smoke validates KDE-first Runtime presence |
| D10 Runtime packaging and service binding | Later | Build/system integration | D1 and D9 | Installable service files and activation records without production enablement |

## Recommended Parallelization

Start with D1, D2, D3, and D8.

- D1 gives other branches a real owner boundary.
- D2 makes install inputs trustworthy.
- D3 gives later execution work a durable state model.
- D8 prevents more contract-only expansion while the other branches work.

D5 can run in parallel with D2 if it only stages activation artifacts under explicit roots and does not consume untrusted production inputs.

D6 should wait until D2, D3, and D4 have enough evidence to block unsafe launches.

D9 and D10 should wait until the product image can validate real Runtime evidence rather than only preview contracts.

## D1: Runtime Owner Process Boundary

### Mission

Turn the Runtime owner from in-process previews and C smoke bridge payloads into a constrained Go smoke-owner process that can own a private test session-bus name, serve read-only Runtime methods, and fail write methods closed.

### Start from

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `internal/runtime/appidentity/runtime_owner_*.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `runtime/dbus/xnix_compatd_smoke.c`
- `scripts/dbus_session_smoke.rb`
- `scripts/runtime_contract_drift_report.rb`

### Deliver

- Long-running smoke-owner mode.
- Private session-bus ownership only in constrained tests.
- Read-only D-Bus method routing through the Go owner table.
- Explicit unsupported-read results for any intentionally unimplemented route.
- Stable disabled responses for every write method.
- Lifecycle logs for startup, route table, readiness, bus claim, dispatch, and shutdown.

### Acceptance

- A restricted container smoke can start the owner on a private session bus.
- Representative D-Bus reads return safe Go Runtime payloads.
- Every write method returns a stable disabled error.
- Contract drift reporting fails if D-Bus XML, owner routes, smoke adapter, D-Bus client, or CLI route commands diverge.

### Required verification

```text
go test ./...
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/container.rb runtime-owner-candidate-smoke
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement D1 from docs/claude-code-domain-dispatch.md.

Target outcome:
- Add a constrained Go Runtime smoke-owner process that can claim a private test session-bus name, serve read-only Runtime methods through the existing owner route table, and fail every write method closed.

Hard constraints:
- No production D-Bus ownership, system service installation, write enablement, backend launch, network fetch, privileged container, host networking, Docker socket mount, broad host mount, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, unsupported-read, disabled-write, lifecycle, and blocked-production-owner tests.
- Run the required D1 verification commands.
```

## D2: Recipe, Artifact, Cache, and Install Trust

### Mission

Turn recipe trust, package sources, artifact manifests, acquisition preflight, cache planning, staging receipts, and install readiness into one fixture-first trust pipeline.

### Start from

- `internal/runtime/recipe/`
- `internal/runtime/artifact/`
- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `internal/runtime/appidentity/package_source.go`
- `internal/runtime/appidentity/acquisition_preflight.go`
- `internal/runtime/appidentity/artifact_manifest.go`
- `internal/runtime/appidentity/install_plan.go`
- `runtime/recipes/registry.json`
- `cmd/xnix-runtime-go/`

### Deliver

- Read-only recipe store with local roots.
- Registry validation for schema, IDs, relative paths, digests, and declared signing state.
- Replaceable signature verifier boundary without committed production keys.
- Fixture artifact manifest parser and SHA-256 verifier.
- Runtime-root-scoped cache and staging receipts.
- Install readiness that joins recipe trust, artifact verification, cache state, staging state, and owner readiness.

### Acceptance

- Invalid recipe digests fail closed.
- Invalid artifact digests block staging.
- Development fixtures remain usable but cannot pass production trust.
- Cache and staging paths stay under explicit Runtime or test roots.
- KDE-facing output exposes relative receipts and user-safe status only.

### Required verification

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

### Copyable prompt

```text
Implement D2 from docs/claude-code-domain-dispatch.md.

Target outcome:
- Build a fixture-first recipe, artifact, cache, staging, and install-readiness trust pipeline that verifies digests, stages only under explicit roots, and fails closed for invalid inputs.

Hard constraints:
- No default network fetch, host package-manager call, private key, production trust bypass, desktop write, backend launch, privileged container, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add valid fixture, invalid digest, invalid path, production-signature-blocked, and root-escape tests.
- Run the required D2 verification commands.
```

## D3: Runtime State Root and Backend Lifecycle

### Mission

Create durable state-root records for application environment lifecycle without creating environments or starting backends.

### Start from

- `internal/runtime/appidentity/state_root.go`
- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `cmd/xnix-runtime-go/`

### Deliver

- State-root schema for application lifecycle records.
- Lifecycle states for missing, planned, staged, ready, repair-required, blocked, and retired.
- Atomic writes under explicit state roots.
- Relative receipt paths in all user-facing output.
- Read models consumed by backend binding, execution readiness, diagnostics, and owner readiness.

### Acceptance

- State-root writes cannot escape the configured root.
- Lifecycle transitions reject invalid jumps.
- Backend launch remains disabled.
- KDE-facing output never exposes state-root paths or backend internals.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_application_state_root.rb
ruby -Ilib test/test_compatibility_backend_binding.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement D3 from docs/claude-code-domain-dispatch.md.

Target outcome:
- Add a state-root-scoped backend lifecycle store that records environment readiness states without creating environments or launching backends.

Hard constraints:
- No backend process start, environment creation outside explicit roots, host-root mutation, raw command exposure, or state-root path exposure in KDE-facing output.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add lifecycle transition, root escape, corrupted receipt, and blocked launch tests.
- Run the required D3 verification commands.
```

## D4: Portal, Permission, Snapshot, and Rollback Safety

### Mission

Build the safety plane that execution must pass through before any real launch can be considered: Portal review records, permission outcomes, snapshot baselines, and rollback receipts.

### Start from

- `internal/runtime/portal/`
- `internal/runtime/snapshot/`
- `internal/runtime/execution/`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/permission_review.go`

### Deliver

- Fake-mode Portal request records with granted, denied, cancelled, failed, and completed states.
- Snapshot baseline records under explicit state roots.
- Rollback receipts that verify object digests before restore.
- Permission review summaries consumable by execution preflight and Compatibility Center.

### Acceptance

- Real Portal transport calls remain disabled by default.
- Snapshot and rollback never touch the host root.
- Permission records are separated from execution approval.
- Corrupted snapshot objects block rollback.

### Required verification

```text
go test ./internal/runtime/portal ./internal/runtime/snapshot ./internal/runtime/execution ./internal/runtime/appidentity
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement D4 from docs/claude-code-domain-dispatch.md.

Target outcome:
- Strengthen the Portal permission, snapshot, and rollback safety plane so execution preflight can depend on durable review and restore evidence.

Hard constraints:
- No real Portal calls, launch approval, backend launch, host-root mutation, user-document reads, or raw path exposure.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add permission outcome, snapshot digest, rollback blocked, root escape, and execution-not-approved tests.
- Run the required D4 verification commands.
```

## D5: KDE Activation and Seven Entry Points

### Mission

Materialize KDE-visible application presence from Runtime evidence while keeping KDE as presentation only and keeping all writes under explicit staging roots.

### Start from

- `internal/runtime/appidentity/desktop_activation_*.go`
- `internal/runtime/appidentity/kde_*.go`
- `internal/runtime/appidentity/desktop_entry.go`
- `internal/runtime/appidentity/file_association.go`
- `internal/runtime/appidentity/notification.go`
- `internal/runtime/appidentity/tray_status.go`
- `internal/runtime/appidentity/krunner.go`
- `cmd/xnix-runtime-go/`

### Deliver

- Desktop entries, MIME fragments, Dolphin service-menu fragments, icon receipts, tray summaries, notification summaries, settings links, and Compatibility Center routes.
- Staging-only writer with no KDE cache refresh by default.
- Receipt manifest with relative paths and SHA-256 digests.
- Compatibility Center read models that cite the staged evidence.

### Acceptance

- Staging refuses filesystem root and existing-file overwrite.
- No launch, backend start, KDE cache refresh, host-root mutation, or raw executable exposure occurs.
- The seven entry points all consume Runtime evidence or staged receipts.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby -Ilib test/test_desktop_entry.rb
ruby -Ilib test/test_file_association_model.rb
ruby -Ilib test/test_task_manager_identity.rb
ruby -Ilib test/test_tray_status_model.rb
ruby -Ilib test/test_notification_request.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement D5 from docs/claude-code-domain-dispatch.md.

Target outcome:
- Convert Runtime application identity and trust evidence into staged KDE activation artifacts and seven-entry-point receipts under explicit roots.

Hard constraints:
- No host-root writes, KDE cache refresh, backend launch, raw executable exposure, production desktop installation, privileged container, or broad host mount.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add staging-root refusal, overwrite refusal, receipt digest, seven-entry-point coverage, and no-launch tests.
- Run the required D5 verification commands.
```

## D6: Execution Transaction Pipeline

### Mission

Connect launch intent, execution request, review, decision, preflight, resource grant, transaction, session, and session status into a durable blocked-by-default execution pipeline.

### Start from

- `internal/runtime/execution/`
- `internal/runtime/appidentity/execution_*.go`
- `internal/runtime/appidentity/launch_intent.go`
- `internal/runtime/appidentity/runtime_write_gate.go`
- `cmd/xnix-runtime-go/`

### Deliver

- Request intake records under explicit state roots.
- Review decision records.
- Preflight records that require trust, lifecycle, Portal, snapshot, and write-gate evidence.
- Transaction ledger records with disabled launch by default.
- Session status read models that can show blocked, pending, failed, or completed states without starting a backend.

### Acceptance

- No request can transition to launchable while write gates remain disabled.
- Missing Portal, snapshot, lifecycle, or trust evidence blocks preflight.
- Transaction records expose relative receipts only.
- Backend launch remains disabled.

### Required verification

```text
go test ./internal/runtime/execution ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_launch_request.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby -Ilib test/test_runtime_write_gate.rb
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement D6 from docs/claude-code-domain-dispatch.md.

Target outcome:
- Build a durable execution transaction pipeline that records requests, reviews, preflight, transactions, and session status while keeping launch disabled.

Hard constraints:
- No write-method enablement, backend launch, permission grant creation, real Portal calls, host-root mutation, raw command exposure, or state-root path exposure.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add missing-prerequisite, rejected-review, disabled-write-gate, transaction receipt, and no-backend-launch tests.
- Run the required D6 verification commands.
```

## D7: Diagnostics, Repair, and AI Boundary

### Mission

Convert diagnostics, AI input, recommendation, and repair approval previews into fixture-backed records without calling live AI providers or auto-repairing anything.

### Start from

- `internal/runtime/diagnostic/`
- `internal/runtime/appidentity/diagnostics.go`
- `internal/runtime/appidentity/ai_diagnostic_*.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/test_plan.go`
- `internal/runtime/appidentity/test_result.go`
- `cmd/xnix-runtime-go/`

### Deliver

- Fixture diagnostic run records.
- Privacy-filtered AI input records.
- Review-first repair recommendations.
- Repair approval gates that require explicit human review.
- Diagnostic history summaries for Compatibility Center cards and sections.

### Acceptance

- No AI provider call occurs.
- No file content or host path is exposed.
- Repair remains review-only and execution-disabled.
- Corrupt or unknown diagnostic fixtures fail closed.

### Required verification

```text
go test ./internal/runtime/diagnostic ./internal/runtime/appidentity ./cmd/xnix-runtime-go
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
Implement D7 from docs/claude-code-domain-dispatch.md.

Target outcome:
- Record fixture-backed diagnostics, privacy-filtered AI input, review-only repair recommendations, and Compatibility Center diagnostic history without live providers or auto-repair.

Hard constraints:
- No AI provider call, file-content read, raw path exposure, auto-repair, backend launch, network access, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add valid fixture, corrupt fixture, privacy filter, repair-review-required, and no-provider-call tests.
- Run the required D7 verification commands.
```

## D8: Developer Evidence and Drift Harness

### Mission

Make it difficult for new contracts or previews to enter the repository without implementation evidence, ownership, and tests.

### Start from

- `scripts/runtime_contract_drift_report.rb`
- `scripts/implementation_evidence_report.rb`
- `scripts/verify_layout.rb`
- `test/test_runtime_contract_drift_report.rb`
- `test/test_implementation_evidence_report.rb`
- `docs/claude-code-mainline-implementation-plan.md`
- `docs/claude-code-large-empty-domain-assignments.md`

### Deliver

- JSON and Markdown evidence reports that classify every major domain.
- Orphan contract detection for Runtime read methods and CLI routes.
- Evidence level gates for contract-only, fixture-implemented, state-root-implemented, smoke-owned, and production-gated.
- Next-dispatch guidance that points to this document and the mainline plan.

### Acceptance

- A new read method with no owner, client, smoke, or implementation package fails evidence checks.
- Reports do not require Docker, QEMU, network, privileged containers, backend launch, or host-root mutation.
- Markdown output is human-readable enough for branch assignment.

### Required verification

```text
ruby -Ilib test/test_implementation_evidence_report.rb
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/implementation_evidence_report.rb --format markdown
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement D8 from docs/claude-code-domain-dispatch.md.

Target outcome:
- Strengthen evidence and drift reporting so new Runtime contracts, CLI routes, D-Bus clients, smoke adapters, and implementation domains cannot silently remain contract-only.

Hard constraints:
- No Docker, QEMU, network, privileged container, backend launch, host-root mutation, or unrelated implementation work.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add JSON, Markdown, orphan-contract, missing-owner, missing-test, and no-host-mutation tests.
- Run the required D8 verification commands.
```

## D9: Atomic KDE Image and QEMU Acceptance

### Mission

Validate that the KDE-first product path can be represented in a constrained image/QEMU smoke without weakening the original Buildroot learning baseline or touching the host root.

### Start from

- `docs/kde-image-pipeline.md`
- `docs/kde-first-presence-smoke-spec.md`
- `docs/kde-first-compatibility-acceptance.md`
- `buildroot/`
- `scripts/container.rb`
- `lib/xnix/qemu.rb`
- `lib/xnix/serial_log.rb`

### Deliver

- Image smoke plan that separates Buildroot learning baseline from KDE product baseline.
- Restricted QEMU acceptance that persists serial logs and does not use host networking.
- Runtime presence checks that validate real owner/trust/state evidence where available.
- Clear skip and blocked states when KDE image prerequisites are absent.

### Acceptance

- Existing Buildroot/QEMU SSH baseline still passes.
- KDE smoke never requires privileged containers, host networking, Docker socket mounts, or broad host mounts.
- Runtime evidence checks fail closed when only preview contracts exist.

### Required verification

```text
ruby -Ilib test/test_qemu.rb
ruby -Ilib test/test_serial_log.rb
ruby -Ilib test/test_milestone.rb
ruby -Ilib test/test_container.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement D9 from docs/claude-code-domain-dispatch.md.

Target outcome:
- Add a constrained KDE-first image/QEMU acceptance path that validates Runtime presence evidence while preserving the existing Buildroot learning baseline.

Hard constraints:
- No privileged container, host networking, Docker socket mount, broad host mount, host-root mutation, production backend launch, or unstable host dependency.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add skip, blocked, serial-log, no-host-network, no-privileged-container, and Runtime-evidence tests.
- Run the required D9 verification commands.
```

## D10: Runtime Packaging and Service Binding

### Mission

Prepare the Runtime service binding, activation files, and package layout for installable production ownership without enabling production ownership yet.

### Start from

- `buildroot/board/xnix/rootfs-overlay/`
- `runtime/dbus/`
- `internal/runtime/appidentity/runtime_service_binding.go`
- `internal/runtime/appidentity/runtime_live_owner_gate.go`
- `cmd/xnix-runtime-owner/`
- `cmd/xnix-runtime-go/`

### Deliver

- Service binding records for binary path, service file, D-Bus activation file, systemd unit, and ownership gate.
- Install-root-scoped packaging checks.
- Readiness gates for preview-only, smoke-owner, and production-owner-blocked states.
- Documentation for what must be true before production ownership can be enabled.

### Acceptance

- Packaging checks can run against an explicit install root.
- Production ownership remains blocked.
- Service files do not contain host-local absolute paths.
- Missing or mismatched version files fail readiness.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-owner ./cmd/xnix-runtime-go
ruby -Ilib test/test_runtime_service_binding.rb
ruby -Ilib test/test_runtime_live_owner_gate.rb
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement D10 from docs/claude-code-domain-dispatch.md.

Target outcome:
- Build install-root-scoped Runtime packaging and service-binding readiness checks while keeping production D-Bus ownership blocked.

Hard constraints:
- No production service enablement, host service installation, host-root mutation, backend launch, privileged container, or host-local path leakage.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add install-root, missing-file, version-mismatch, production-blocked, and no-host-path tests.
- Run the required D10 verification commands.
```

## Branch Completion Template

Every Claude Code branch should finish with this summary:

```text
Implemented domain:
- D<id> <name>

Implementation evidence:
- <new durable state, owner process, receipt, or harness evidence>

Still disabled:
- <unsafe operations intentionally left disabled>

Tests run:
- <commands and results>

Files changed:
- <short list>

Next recommended dispatch:
- <D-id and reason>
```
