# Claude Code Empty-Domain Implementation Packages

> Last updated: 2026-07-16 | Baseline: v0.2.293

This document splits the contract-heavy parts of Xnix into larger, relatively independent implementation packages that can be handed to Claude Code one branch at a time.

Use this file when the problem is not "add one more preview contract", but "turn a contracted or stub-shaped domain into durable implementation evidence".

This document intentionally does not replace:

- `docs/claude-code-implementation-packages.md`, which is treated as an existing package guide.
- `docs/claude-code-large-empty-domain-assignments.md`, which is the coarse dispatcher for branch-sized empty-domain assignments.
- `docs/claude-code-open-domain-work-packages.md`, which is the broad backlog map.
- `docs/claude-code-independent-implementation-briefs.md`, which is the smaller prompt-ready brief list.
- `docs/claude-code-contract-implementation-handoff.md`, which groups contract-to-implementation workstreams.

## How to Assign Work

Assign exactly one package per Claude Code branch unless the package explicitly lists a required dependency.

Each branch must:

- Keep all source, comments, fixtures, CLI output, documentation, and test descriptions in English.
- Prefer Go for durable Runtime product behavior.
- Use C only for low-level Runtime surfaces, ABI-shaped records, or already-owned C contracts.
- Use Ruby for tests, smoke scripts, and developer tooling.
- Keep host-machine impact minimal: no privileged containers, host networking, Docker socket mounts, broad host-directory mounts, host-root mutation, or real backend launch unless the package explicitly enables the exact behavior.
- Keep production D-Bus ownership, write methods, real Portal calls, and compatibility backend execution disabled unless the package explicitly enables a gated test-only path.
- Avoid exposing raw backend commands, executable paths, compatibility storage paths, profile terminology, file contents, secrets, tokens, or private keys in KDE/user-facing output.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Run the package-specific tests plus `ruby scripts/verify_layout.rb`.
- Run `ruby scripts/runtime_contract_drift_report.rb --format json` when touching D-Bus XML, Runtime route manifests, owner dispatch, smoke adapters, or Runtime CLI commands.
- Run full build and QEMU smoke at every tenth code version, following repository policy.

## Package Selection Guide

| Priority | Package | Best for | Main evidence produced |
| --- | --- | --- | --- |
| 1 | P1 Runtime owner service | Replacing preview-only ownership with a constrained owner process | Session-bus owner smoke, read dispatch, closed write responses |
| 2 | P2 Recipe and artifact trust pipeline | Making install inputs real without downloads or host mutation | Verified recipes, cache/stage receipts, trust diagnostics |
| 3 | P3 Environment lifecycle state | Turning backend plans into persisted Runtime state | State-root lifecycle records and readiness transitions |
| 4 | P4 Portal and snapshot control plane | Building safety prerequisites before execution | Fake Portal objects, snapshot receipts, rollback verification |
| 5 | P5 KDE activation and shell materialization | Making Windows apps appear like Linux apps safely | Staged/target-root desktop, MIME, service-menu, tray, and receipt evidence |
| 6 | P6 Execution transaction ledger | Preparing launch without starting backends | Review, preflight, grant, transaction, and session records |
| 7 | P7 Diagnostics, repair, and AI boundary | Replacing diagnostic previews with reproducible records | Fixture test runs, repair recommendations, provider-gated AI output |
| 8 | P8 Atomic KDE image and QEMU acceptance | Proving the product shape boots and can be tested | Buildable image manifest, constrained QEMU/KDE smoke reports |
| 9 | P9 Developer verification harness | Preventing contract drift and empty implementations from returning | JSON/Markdown reports and CI-friendly gates |

## Current Baseline Evidence

Claude Code should treat the following as already-started implementation evidence, not as greenfield work:

| Domain | Baseline evidence at v0.2.293 | What remains useful to implement |
| --- | --- | --- |
| Runtime owner | Owner candidate, full in-process read dispatch coverage for every D-Bus read-only Runtime method, owner-local probes, route manifests, service binding previews, lifecycle JSONL, and write gates exist. | Constrained long-running session-bus event loop, smoke-owned readiness, bus-call parity, and closed write-method evidence through the owner process. |
| Recipe and artifacts | Recipe registry/trust previews exist. Local artifact manifest parsing, digest verification, cache acquisition, and `artifact-stage-record` stage receipts exist for fixture artifacts. | Production-shaped trust-store boundary, signature verifier interface, state-root namespace integration, and richer install-gate diagnostics. |
| KDE activation | Desktop activation previews and a target-root staging writer exist for desktop entries, MIME, Dolphin service menus, activation manifests, and rollback receipts. | Drift checks against live materialized state, optional digest-verified icons, KDE status consumption of receipt state, and stricter transaction rollback UX. |
| Execution | Execution previews and `execution-ledger-record` exist for reviewed/preflighted state-root transaction records without backend launch. | Fake commit/list/inspect flows, joined dependency gates from trust/environment/Portal/snapshot state, and session-status records. |
| Portal and snapshots | Fake Portal model previews and constrained snapshot store evidence exist. | Durable Portal request state machine, disabled real transport boundary, joined safety view, and rollback receipt integration. |
| Diagnostics, repair, and AI | Fixture diagnostics, fake/disabled AI provider boundaries, `diagnostic-run-record` state-root receipts, `diagnostic-run-history` summaries, and `diagnostic-history-preview` Compatibility Center projections exist. | Broader repair workflows, approval receipt integration, provider configuration policy, and deeper Compatibility Center consumption of diagnostic history. |
| Developer gates | Contract drift and implementation evidence reports exist. | Stronger orphan-preview detection, CI-friendly package ownership checks, and evidence regression guards. |

When a package overlaps existing evidence, extend the current package instead of replacing the surface. The preferred path is to add one new durable state transition, receipt, or smoke-owned behavior per branch.

## Assignment Protocol

Use this protocol when handing a package to Claude Code:

1. Assign exactly one package ID and one short outcome.
2. Name the existing files to start from.
3. State the unsafe actions that must remain disabled.
4. Ask for a small semantic version bump, changelog entry, product overview update, targeted tests, and layout verification.
5. Ask Claude Code to stop and report blockers if the package would require network fetch, host-root writes, production D-Bus ownership, real Portal calls, or backend launch.

## Implementation Evidence Standard

Claude Code should not treat a new contract, preview, CLI stub, route declaration, or documentation paragraph as implementation by itself. A package becomes mergeable only when it turns at least one existing contract into evidence that can fail, persist, or be exercised through a constrained smoke path.

Use this scale when reviewing a branch:

| Evidence level | Meaning | Merge expectation |
| --- | --- | --- |
| Contract-only | Interfaces, XML, model names, preview fields, or docs exist but nothing durable happens. | Not enough for these packages unless the branch only updates handoff docs. |
| Fixture-implemented | The behavior runs against controlled fixtures and has success and failure tests. | Acceptable for package foundations that must avoid host mutation. |
| State-root-implemented | The behavior writes or reads durable records under an explicit test/runtime state root. | Preferred for lifecycle, Portal, snapshot, execution, diagnostics, and staging work. |
| Smoke-owned | A constrained container or private session-bus smoke can exercise the behavior end to end. | Required for Runtime owner and D-Bus bridge progress. |
| Production-gated | The production path is represented, observable, and blocked until prerequisites are real. | Required whenever a package approaches writes, launch, Portal transport, backend processes, or host integration. |

Every implementation package should add at least one of these durable evidence types:

- A state-root record with a schema version, app id, operation id, created timestamp, result state, blocked reasons, and relative receipt paths.
- A local fixture pipeline with verified digests and deterministic mismatch failures.
- A constrained smoke path that exercises the real adapter boundary without production privileges.
- A report entry that classifies the domain and fails on orphan contract-only surfaces.
- A KDE-safe read model that consumes durable receipts instead of restating static previews.

If a branch only adds more preview fields, more D-Bus methods, or more documentation without one of those evidence types, ask Claude Code to stop and narrow the branch. The point is to make empty domains heavier, not merely louder.

## Package Completion Definition

Each package handoff should include a short completion statement in this form:

```text
Converted contract(s):
- <contract or preview surface>

Implementation evidence added:
- <fixture/state-root/smoke/report evidence>

Still gated:
- <unsafe production behavior that remains disabled>

Verification run:
- <commands>
```

This keeps package reviews concrete and prevents broad branches from hiding behind "foundation" language.

Copyable template:

```text
Implement <PACKAGE_ID> from docs/claude-code-empty-domain-implementation-packages.md.

Target outcome:
- <one durable implementation outcome>

Constraints:
- Keep production D-Bus ownership, real backend launch, real Portal calls, network fetch, privileged containers, broad host mounts, Docker socket mounts, and host-root mutation disabled unless this exact package explicitly enables a gated test-only path.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Runtime behavior and Ruby for tests/tooling.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, failure, and blocked-state tests.
- Run the package tests and ruby scripts/verify_layout.rb.
- Run runtime contract drift reporting if Runtime routes, D-Bus XML, owner dispatch, smoke adapters, or CLI route commands change.
- Do not touch unrelated packages or docs/claude-code-implementation-packages.md unless explicitly asked.
```

## P1: Runtime Owner Service

### Mission

Turn the Runtime owner from preview/adapters into a constrained Go owner process that can claim a test session-bus name, serve read-only methods, and fail write methods closed.

### Start from

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `internal/runtime/appidentity/runtime_owner_*.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `runtime/dbus/xnix_compatd_smoke.c`
- `scripts/dbus_session_smoke.rb`
- `scripts/runtime_owner_candidate_smoke.rb`

### Deliver

- A smoke-owner mode that can run inside the constrained container session bus.
- A route table that serves the existing full D-Bus read-only method set through Go handlers or explicit unsupported-read markers.
- Deterministic `WriteMethodDisabled` responses for all write methods.
- Owner readiness evidence that distinguishes `preview-only`, `smoke-owner`, and `production-owner`.
- Logs for startup, route-table version, bus mode, dispatch result, and shutdown reason.
- Tests that prove write methods remain disabled and production bus ownership remains gated.

### Do not deliver

- Production system-bus ownership.
- System service installation.
- Real backend launch.
- Host root writes.
- KDE-owned Runtime policy.

### Acceptance

- A constrained container can start the owner and call representative read-only methods while in-process coverage remains complete for every D-Bus read-only Runtime method.
- The owner returns the same safe shapes as the existing Go preview routes or a stable unsupported-read marker.
- Every write method returns a stable disabled error.
- Runtime contract drift reporting remains clean.

### Required tests

```text
go test ./cmd/xnix-runtime-owner ./internal/runtime/owner ./cmd/xnix-runtime-go ./internal/runtime/appidentity
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
ruby scripts/verify_layout.rb
```

### Prompt for Claude Code

```text
Implement P1 from docs/claude-code-empty-domain-implementation-packages.md. Work on one branch only. Keep production ownership and write methods disabled. Produce constrained session-bus owner evidence, tests, and contract drift proof.
```

## P2: Recipe and Artifact Trust Pipeline

### Mission

Turn recipe trust, package sources, artifact manifests, cache planning, and staging into a real local pipeline while keeping network fetch and host mutation disabled by default.

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

### Deliver

- A read-only recipe store with digest verification and production-shaped trust states.
- A verifier boundary for signed recipe metadata without committing keys or secrets.
- A local fixture artifact source that verifies SHA-256 before staging.
- A Runtime state-root scoped cache namespace.
- Artifact staging receipts with relative paths, digests, and blocked reasons.
- Diagnostics consumable by install gates, owner readiness, and Compatibility Center pages.

### Do not deliver

- Remote downloads by default.
- Production signing keys.
- Host package-manager calls.
- Host root mutation.
- Desktop activation writes.
- Backend process starts.

### Acceptance

- Invalid recipe digests fail closed.
- Missing or invalid production signatures cannot pass production readiness.
- Artifact digest mismatch blocks staging.
- Cache and staging paths stay inside the configured state/test root.
- Development fixtures remain usable but are visibly non-production.

### Required tests

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

### Prompt for Claude Code

```text
Implement P2 from docs/claude-code-empty-domain-implementation-packages.md. Build a local-only recipe and artifact trust pipeline. No network fetch, host package manager, host root mutation, desktop activation writes, or backend launch.
```

## P3: Environment Lifecycle State

### Mission

Convert backend/environment previews into a Runtime-owned lifecycle state machine with safe persistence under a controlled Runtime state root.

### Start from

- `internal/runtime/environment/`
- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/backend_capability_matrix.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `internal/runtime/appidentity/state_root.go`

### Deliver

- Persisted lifecycle states such as `missing`, `planned`, `staged`, `ready`, `repair-required`, and `blocked`.
- Test fixtures that simulate missing, staged, broken, and repaired environments.
- Readiness transitions driven by recipe trust, artifact staging, Portal state, and snapshot state.
- KDE-safe explanations for why an environment is ready or blocked.
- No raw backend commands, executable paths, profile names, or host storage paths in desktop-facing output.

### Do not deliver

- Starting Wine, Proton, VM, or any compatibility backend.
- Creating host-level home directories or system state.
- Real Portal requests.
- Real snapshot operations outside a controlled state root.

### Acceptance

- Lifecycle state survives inside a test-controlled root.
- Readiness changes when upstream state changes.
- Execution remains disabled unless P6 owns a gated transaction path.
- User-facing summaries remain implementation-detail safe.

### Required tests

```text
go test ./internal/runtime/environment ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_backend_environment_plan.rb
ruby -Ilib test/test_compatibility_backend_binding.rb
ruby -Ilib test/test_compatibility_backend_lifecycle.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby -Ilib test/test_application_state_root.rb
ruby scripts/verify_layout.rb
```

### Prompt for Claude Code

```text
Implement P3 from docs/claude-code-empty-domain-implementation-packages.md. Add a persisted Runtime environment lifecycle state machine under a controlled state root. Do not start real compatibility backends or expose backend details.
```

## P4: Portal and Snapshot Control Plane

### Mission

Make permission requests and restore points real enough for execution readiness while remaining fake-first and state-root scoped.

### Start from

- `internal/runtime/portal/`
- `internal/runtime/snapshot/`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/execution_preflight.go`
- `internal/runtime/appidentity/execution_resource_grant.go`
- `internal/runtime/appidentity/execution_readiness.go`

### Deliver

- Durable fake Portal request records with states such as `planned`, `requested`, `granted`, `denied`, `cancelled`, `expired`, and `failed`.
- A disabled real Portal transport boundary.
- Snapshot creation, listing, verification, rollback, and rollback receipts inside the configured Runtime state root.
- A joined safety view explaining which permissions and restore points are required before execution.
- Correlation handles suitable for D-Bus callers and Compatibility Center pages.

### Do not deliver

- Real XDG Desktop Portal calls by default.
- Direct access to user documents.
- System snapshots.
- Host-root rollback.
- Backend launch.

### Acceptance

- Fake Portal requests can complete, deny, cancel, expire, and fail deterministically.
- Snapshot operations cannot escape the configured state root.
- Snapshot identifiers and rollback receipts are verified before use.
- Execution readiness can consume permission and snapshot state without starting execution.

### Required tests

```text
go test ./internal/runtime/portal ./internal/runtime/snapshot ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/verify_layout.rb
```

### Prompt for Claude Code

```text
Implement P4 from docs/claude-code-empty-domain-implementation-packages.md. Build fake-first Portal request records and state-root-scoped snapshot/rollback evidence. Keep real Portal calls, system snapshots, host-root rollback, and backend launch disabled.
```

## P5: KDE Activation and Shell Materialization

### Mission

Move KDE integration from previews to auditable, target-root-scoped materialization so compatibility apps can appear as normal Linux apps without touching the host root.

### Start from

- `internal/runtime/activation/`
- `internal/runtime/appidentity/desktop_activation_*.go`
- `internal/runtime/appidentity/kde_*.go`
- `internal/runtime/appidentity/window_identity_routes.go`
- `kde/dolphin/servicemenus/`
- `kde/plasmoids/org.xnix.compatibilitycenter/`
- `scripts/install_runtime_activation.rb`
- `scripts/runtime_activation_smoke.rb`

### Deliver

- A target-root writer for desktop entries, MIME associations, Dolphin service menus, activation manifests, and rollback receipts.
- Drift checks between previewed content and materialized files.
- Optional staged icon handling if source icons are fixture-controlled and digest-verified.
- Compatibility Center and KDE status models that consume receipt state.
- Rollback verification that refuses unknown or mismatched receipts.

### Do not deliver

- Host-root writes.
- Production desktop activation by default.
- Backend launch from desktop files.
- KWin rule application to the live session.
- Live tray bridge enablement.

### Acceptance

- All file writes are under a supplied target/staging root.
- Existing target files are not overwritten without an explicit safe transaction.
- Rollback receipts include relative paths and digests.
- KDE/user-facing output does not expose backend commands or storage paths.

### Required tests

```text
go test ./internal/runtime/activation ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_desktop_entry.rb
ruby -Ilib test/test_file_association_model.rb
ruby -Ilib test/test_dolphin_service_menu.rb
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby scripts/container.rb runtime-activation-smoke
ruby scripts/verify_layout.rb
```

### Prompt for Claude Code

```text
Implement P5 from docs/claude-code-empty-domain-implementation-packages.md. Extend KDE activation materialization under an explicit target root only. Add drift checks and rollback receipts. Do not write to the host root or enable backend launch.
```

## P6: Execution Transaction Ledger

### Mission

Create a durable execution transaction ledger that records launch intent, review, preflight, resource grants, transaction state, and session identity without starting real backends by default.

### Start from

- `internal/runtime/execution/`
- `internal/runtime/appidentity/launch_intent.go`
- `internal/runtime/appidentity/execution_request.go`
- `internal/runtime/appidentity/execution_review.go`
- `internal/runtime/appidentity/execution_decision.go`
- `internal/runtime/appidentity/execution_preflight.go`
- `internal/runtime/appidentity/execution_resource_grant.go`
- `internal/runtime/appidentity/execution_transaction.go`
- `internal/runtime/appidentity/execution_session.go`
- `internal/runtime/appidentity/execution_session_status.go`
- `internal/runtime/appidentity/runtime_write_gate.go`

### Deliver

- A state-root-scoped transaction ledger.
- Immutable records for request, user review, preflight, grant proposal, transaction decision, and session identity.
- Transaction states such as `draft`, `review-required`, `blocked`, `ready`, `committed-fake`, `failed`, and `cancelled`.
- Integration with environment readiness, Portal permission state, snapshot readiness, and write gates.
- A fake runner that records a fake committed session only when all gates are satisfied.

### Do not deliver

- Real backend process starts by default.
- Raw command exposure.
- Runtime `Launch` write-method enablement in production.
- Portal permission grants.
- Host-root mutation.

### Acceptance

- A fake execution transaction can be created, blocked, approved, committed in fake mode, listed, and inspected.
- Missing permission, missing snapshot, untrusted recipe, or blocked backend state prevents commit.
- Session identity remains KDE-safe and does not expose implementation details.
- Production launch remains disabled.

### Required tests

```text
go test ./internal/runtime/execution ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_launch_request.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby -Ilib test/test_compatibility_run_plan.rb
ruby scripts/kde_first_presence_smoke.rb
ruby scripts/verify_layout.rb
```

### Prompt for Claude Code

```text
Implement P6 from docs/claude-code-empty-domain-implementation-packages.md. Add a state-root-scoped execution transaction ledger with fake commits only. Do not start real backends, enable production Launch, grant real Portal permissions, or mutate the host root.
```

## P7: Diagnostics, Repair, and AI Boundary

### Mission

Turn compatibility test, result, repair, action queue, action review, and AI diagnostic previews into reproducible fixture-driven records with a disabled-by-default provider boundary.

### Start from

- `internal/runtime/diagnostics/`
- `internal/runtime/appidentity/test_plan.go`
- `internal/runtime/appidentity/test_result.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/ai_diagnostics.go`
- `internal/runtime/appidentity/kde_action_*.go`
- `cmd/xnix-runtime-go/ai_diagnostics_commands.go`

### Deliver

- Fixture-driven test runner records for preflight, smoke, repair-readiness, and regression checks.
- Repair recommendation records linked to snapshot and rollback prerequisites.
- AI diagnostic provider interface with fake provider and disabled production default.
- Approval gates that require explicit user review and snapshot readiness before any repair action can be marked executable.
- JSON and Markdown reports for developer handoff.

### Do not deliver

- Real AI provider calls by default.
- Reading arbitrary user file contents.
- Automatic repair execution.
- Backend launch.
- Network dependency.

### Acceptance

- Tests can run from fixtures without real backends.
- Diagnostics produce stable result records and repair recommendations.
- AI recommendations are deterministic in fake mode and fail closed without a provider.
- Risky repair actions remain blocked without approval and snapshot readiness.

### Required tests

```text
go test ./internal/runtime/diagnostics ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_test_plan.rb
ruby -Ilib test/test_compatibility_test_result.rb
ruby -Ilib test/test_compatibility_repair_plan.rb
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_ai_repair_approval_gate.rb
ruby scripts/verify_layout.rb
```

### Prompt for Claude Code

```text
Implement P7 from docs/claude-code-empty-domain-implementation-packages.md. Build fixture-driven diagnostics, repair records, and a disabled-by-default AI provider boundary. Do not call real AI providers, read arbitrary user files, auto-repair, or launch backends.
```

## P8: Atomic KDE Image and QEMU Acceptance

### Mission

Create a reproducible product-image path that can prove the KDE-first shape in constrained QEMU without endangering the host.

### Start from

- `image/kinoite/`
- `internal/runtime/image/`
- `scripts/build_kde_image.rb`
- `scripts/build_kde_disk.rb`
- `scripts/boot_kde_image.rb`
- `buildroot/`
- `scripts/full_smoke.rb`
- `scripts/container.rb`

### Deliver

- Image manifest validation with pinned inputs and lock evidence.
- A safe local build plan that refuses privileged containers and broad mounts.
- A QEMU smoke path that persists serial logs and reports boot/KDE/Runtime evidence.
- JSON and Markdown smoke reports.
- Clear separation between the Buildroot learning baseline and the KDE flagship image path.

### Do not deliver

- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Unpinned base images or package sources.
- Host system changes.

### Acceptance

- Image validation can run without network.
- Build/smoke scripts fail closed when unsafe Docker/QEMU options are requested.
- QEMU logs are persisted under a controlled output directory.
- Reports identify whether Runtime owner, KDE shell, activation, and SSH/serial evidence are present.

### Required tests

```text
go test ./internal/runtime/image ./...
ruby -Ilib test/test_kde_image.rb
ruby -Ilib test/test_kde_disk.rb
ruby -Ilib test/test_qemu.rb
ruby -Ilib test/test_serial_log.rb
ruby scripts/verify_layout.rb
```

### Prompt for Claude Code

```text
Implement P8 from docs/claude-code-empty-domain-implementation-packages.md. Build a safe atomic KDE image validation and constrained QEMU smoke path. Do not use privileged containers, host networking, Docker socket mounts, broad host mounts, or host system changes.
```

## P9: Developer Verification Harness

### Mission

Make it difficult for the project to accumulate more contracts without implementation evidence.

### Start from

- `scripts/verify_layout.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/kde_first_presence_smoke.rb`
- `scripts/container.rb`
- `test/test_runtime_contract_drift_report.rb`
- `docs/kde-first-current-gap-audit.md`

### Deliver

- A machine-readable implementation evidence report that classifies domains as `contract-only`, `fixture-implemented`, `state-root-implemented`, `smoke-owned`, or `production-gated`.
- A Markdown report for humans.
- Drift gates for D-Bus XML, owner routes, CLI commands, smoke adapters, and docs references.
- A check that highlights new preview methods that have no implementation package or evidence owner.
- Optional report artifacts under a controlled `tmp/` or `output/` path, never committed by default.

### Do not deliver

- CI secrets.
- Network-only checks.
- Host mutation.
- Broad filesystem scans outside the repository.
- Automatic deletion of user files.

### Acceptance

- The report exits non-zero on contract drift or orphan preview methods.
- The report does not require Docker, QEMU, network, privileged containers, or host-root access.
- Markdown output is stable enough to paste into PRs.
- The report can be used before handing work to Claude Code to pick the next highest-value empty domain.

### Required tests

```text
go test ./...
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/runtime_contract_drift_report.rb --format markdown
ruby scripts/kde_first_presence_smoke.rb
ruby scripts/verify_layout.rb
```

### Prompt for Claude Code

```text
Implement P9 from docs/claude-code-empty-domain-implementation-packages.md. Add implementation-evidence reporting and drift gates that expose contract-only domains. Do not require Docker, QEMU, network, privileged containers, or host-root access for the default report.
```

## Recommended Assignment Order

1. Assign P1 if Runtime ownership is the next milestone.
2. Assign P2 before any real install or execution work.
3. Assign P3 and P4 before P6, because execution needs environment, permission, and snapshot evidence.
4. Assign P5 when the goal is "apps appear as normal Linux apps" without launching them.
5. Assign P6 only after P2, P3, and P4 have enough evidence.
6. Assign P7 whenever diagnostics and repair UX need to become more than static previews.
7. Assign P8 when image/QEMU proof becomes the milestone gate.
8. Assign P9 whenever the project starts adding contracts faster than implementation evidence.

## First Dispatch Wave

Use this wave when the goal is to convert the largest empty domains into implementation evidence without making the project unsafe or too broad to review.

| Slot | Package | Suggested branch | Why this is first-wave work | Minimal mergeable outcome |
| --- | --- | --- | --- | --- |
| 1 | P1 Runtime Owner Service | `codex/runtime-owner-smoke-service` | Most later work depends on a Runtime-owned read path instead of preview-only commands. | A constrained session-bus owner serves the existing read dispatch table and fails every write method closed. |
| 2 | P2 Recipe and Artifact Trust Pipeline | `codex/recipe-artifact-trust-pipeline` | Install, environment, and execution work need trusted local inputs before they can be meaningful. | Local recipes and fixture artifacts verify digests, stage into a state/test root, and produce blocked receipts for invalid inputs. |
| 3 | P3 Environment Lifecycle State | `codex/environment-lifecycle-state` | Backend readiness is currently mostly declarative; execution should not advance until readiness is persistent. | A state-root lifecycle store records missing, planned, staged, ready, repair-required, and blocked states. |
| 4 | P9 Developer Verification Harness | `codex/implementation-evidence-report` | The project is adding contracts quickly; this keeps empty domains visible and prevents drift. | JSON and Markdown reports classify each domain by evidence level and fail on orphan contract-only surfaces. |

Do not assign P6 execution transactions until P2, P3, and P4 have enough state evidence to block unsafe launches. Do not assign P8 image/QEMU acceptance until the branch can test real product milestones rather than only restating preview contracts.

## Copyable First-Wave Prompts

### P1 prompt

```text
Implement P1 from docs/claude-code-empty-domain-implementation-packages.md.

Target outcome:
- Add a constrained Go Runtime owner service that can claim a private test session bus, serve the existing D-Bus read-only Runtime method dispatch table, and fail every write method closed.

Hard constraints:
- No production bus ownership, system service installation, backend launch, host-root writes, privileged containers, broad host mounts, Docker socket mounts, or host networking.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, disabled-write, and blocked-production-owner tests.
- Run the P1 required tests plus ruby scripts/verify_layout.rb and runtime contract drift reporting.
```

### P2 prompt

```text
Implement P2 from docs/claude-code-empty-domain-implementation-packages.md.

Target outcome:
- Build a local-only recipe and artifact trust pipeline with digest verification, fixture artifact staging, and state-root-scoped receipts.

Hard constraints:
- No network fetch, host package-manager calls, host-root mutation, desktop activation writes, production signing keys, or backend launch.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add valid, digest-mismatch, unsigned, and blocked-state tests.
- Run the P2 required tests plus ruby scripts/verify_layout.rb.
```

### P3 prompt

```text
Implement P3 from docs/claude-code-empty-domain-implementation-packages.md.

Target outcome:
- Add a Runtime-owned environment lifecycle state store under a controlled state root, with readiness transitions that can be consumed by execution readiness and KDE-safe summaries.

Hard constraints:
- Do not start Wine, Proton, VM, or any compatibility backend.
- Do not expose raw backend commands, executable paths, profile names, storage paths, or host paths in user-facing output.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add missing, staged, ready, repair-required, and blocked-state tests.
- Run the P3 required tests plus ruby scripts/verify_layout.rb.
```

### P9 prompt

```text
Implement P9 from docs/claude-code-empty-domain-implementation-packages.md.

Target outcome:
- Add implementation-evidence reporting that classifies domains as contract-only, fixture-implemented, state-root-implemented, smoke-owned, or production-gated.

Hard constraints:
- The default report must not require Docker, QEMU, network, privileged containers, host-root access, or broad filesystem scans outside the repository.
- Do not delete, rewrite, or auto-fix user files.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add JSON, Markdown, orphan-contract, and no-host-mutation tests.
- Run the P9 required tests plus ruby scripts/verify_layout.rb.
```

## Handoff Checklist for Each Claude Code Branch

- [ ] The branch implements exactly one package.
- [ ] The branch states which contracts were converted into implementation evidence.
- [ ] The branch keeps unsafe actions gated or disabled as required.
- [ ] The branch adds targeted tests for success, failure, and blocked states.
- [ ] The branch updates version, changelog, and product overview for code changes.
- [ ] The branch runs package tests and `ruby scripts/verify_layout.rb`.
- [ ] The branch runs contract drift reporting when Runtime routes or D-Bus surfaces change.
- [ ] The branch does not touch unrelated packages, secrets, build artifacts, or host configuration.
