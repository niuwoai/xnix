# Claude Code Mainline Implementation Plan

> Last updated: 2026-07-16 | Baseline: v0.2.218

This document is the mainline handoff plan for giving large, relatively independent Xnix implementation packages to Claude Code.

Use it when the goal is to convert contract-only or preview-heavy domains into durable implementation evidence without opening side quests. Each package should be implemented on one branch and should produce one reviewable, testable product step.

This file is intentionally shorter than the detailed package catalogs:

- `docs/claude-code-empty-domain-implementation-packages.md`
- `docs/claude-code-contract-gap-work-packages.md`
- `docs/claude-code-open-domain-work-packages.md`
- `docs/claude-code-independent-implementation-briefs.md`

Those files provide deeper package details. This file provides the mainline dispatch order and copyable task boundaries.

## Non-Negotiable Rules

Every Claude Code implementation branch must:

- Implement exactly one package from this document unless explicitly told otherwise.
- Keep all source, comments, fixtures, tests, CLI output, and documentation in English.
- Prefer Go for durable Compatibility Runtime product behavior.
- Use C only for low-level Runtime contracts, ABI-shaped records, or already-owned C surfaces.
- Use Ruby for tests, smoke scripts, reports, and developer tooling.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Add targeted tests for success, failure, and blocked states.
- Run package-specific tests plus `ruby scripts/verify_layout.rb`.
- Run `ruby scripts/runtime_contract_drift_report.rb --format json` when touching D-Bus XML, Runtime routes, owner dispatch, smoke adapters, or Runtime CLI route commands.
- Stop and report if the package would require network fetch, host-root writes, production D-Bus ownership, real Portal calls, real backend launch, privileged containers, host networking, Docker socket mounts, or broad host-directory mounts.

Claude Code must not:

- Add more preview contracts as a substitute for implementation evidence.
- Enable production write methods before the package explicitly owns a gated write path.
- Expose raw backend commands, executable paths, compatibility storage paths, profile terminology, file contents, secrets, tokens, private keys, or host paths in KDE-facing output.
- Mutate the host root.
- Start Wine, Proton, VM, or other compatibility backends unless a later package explicitly defines a safe gated test-only path.
- Touch unrelated packages or `docs/claude-code-implementation-packages.md` unless explicitly asked.

## Mainline Dependency Shape

```text
M1 Runtime owner read service
  -> M2 Recipe and artifact trust pipeline
      -> M3 Runtime state root and environment lifecycle
          -> M4 Portal and snapshot safety plane
              -> M6 Execution transaction ledger

M5 KDE activation materialization
  depends on M2 for trusted inputs
  can proceed before M6 because it must not launch backends

M7 Diagnostics, repair, and AI boundary
  can proceed after M3
  feeds M6 readiness and Compatibility Center evidence

M8 Developer verification harness
  can proceed at any time
  should be strengthened whenever new contracts appear

M9 Atomic KDE image and QEMU acceptance
  validates integrated milestones
  should not be the first implementation branch
```

## Recommended First Wave

Start with these four packages. They create the foundation without enabling unsafe execution.

| Order | Package | Suggested branch | Why now | Minimal mergeable result |
| --- | --- | --- | --- | --- |
| 1 | M1 Runtime Owner Read Service | `codex/runtime-owner-read-service` | Most other work needs a Runtime-owned read path rather than preview-only commands. | A constrained session-bus owner serves representative read methods and fails all write methods closed. |
| 2 | M2 Recipe and Artifact Trust Pipeline | `codex/recipe-artifact-trust-pipeline` | Install, environment, activation, and execution must start from trusted local inputs. | Local recipes and fixture artifacts verify digests, stage under a controlled root, and produce blocked receipts for invalid inputs. |
| 3 | M3 Runtime State Root and Environment Lifecycle | `codex/environment-lifecycle-state` | Execution readiness needs durable state before any launch path becomes meaningful. | A state-root lifecycle store records missing, planned, staged, ready, repair-required, and blocked states. |
| 4 | M8 Developer Verification Harness | `codex/implementation-evidence-harness` | The repository already has many contracts; this prevents more contract-only surfaces from accumulating silently. | JSON and Markdown reports classify domain evidence and fail on orphan contract-only surfaces. |

Do not assign M6 execution transactions until M2, M3, and M4 have enough state evidence to block unsafe launches.

Do not assign M9 image/QEMU acceptance until it can validate real product evidence rather than only preview contracts.

## M1: Runtime Owner Read Service

### Mission

Turn the Runtime owner from preview/adapters into a constrained Go owner process that can claim a private test session-bus name, serve read-only methods, and fail write methods closed.

### Start from

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `internal/runtime/appidentity/runtime_owner_*.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `runtime/dbus/xnix_compatd_smoke.c`
- `scripts/dbus_session_smoke.rb`
- `scripts/runtime_owner_candidate_smoke.rb`

### Deliver

- A constrained smoke-owner mode.
- A route table from D-Bus read methods to Go handlers or explicit unsupported-read markers.
- Deterministic disabled responses for every write method.
- Owner readiness states for `preview-only`, `smoke-owner`, and `production-owner-gated`.
- Logs for startup, route-table version, bus mode, dispatch result, and shutdown reason.

### Keep disabled

- Production bus ownership.
- System service installation.
- Real backend launch.
- Host root writes.
- KDE-owned Runtime policy.

### Acceptance

- A constrained container can start the owner on a private session bus.
- Representative read methods return safe Go Runtime read models.
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

### Copyable Claude Code prompt

```text
Implement M1 from docs/claude-code-mainline-implementation-plan.md.

Target outcome:
- Add a constrained Go Runtime owner service that can claim a private test session bus, serve representative read-only Runtime methods, and fail every write method closed.

Hard constraints:
- No production bus ownership, system service installation, backend launch, host-root writes, privileged containers, broad host mounts, Docker socket mounts, or host networking.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, disabled-write, unsupported-read, and blocked-production-owner tests.
- Run the M1 required tests plus runtime contract drift reporting.
```

## M2: Recipe and Artifact Trust Pipeline

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

### Keep disabled

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

### Copyable Claude Code prompt

```text
Implement M2 from docs/claude-code-mainline-implementation-plan.md.

Target outcome:
- Build a local-only recipe and artifact trust pipeline with digest verification, fixture artifact staging, and state-root-scoped receipts.

Hard constraints:
- No network fetch, host package-manager calls, host-root mutation, desktop activation writes, production signing keys, or backend launch.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add valid, digest-mismatch, unsigned, invalid-signature, and blocked-state tests.
- Run the M2 required tests plus ruby scripts/verify_layout.rb.
```

## M3: Runtime State Root and Environment Lifecycle

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

### Keep disabled

- Starting Wine, Proton, VM, or any compatibility backend.
- Creating host-level home directories or system state.
- Real Portal requests.
- Real snapshot operations outside a controlled state root.
- Raw backend commands, executable paths, profile names, or host storage paths in desktop output.

### Acceptance

- Lifecycle state survives inside a test-controlled root.
- Readiness changes when upstream state changes.
- Execution remains disabled unless M6 owns a gated transaction path.
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

### Copyable Claude Code prompt

```text
Implement M3 from docs/claude-code-mainline-implementation-plan.md.

Target outcome:
- Add a Runtime-owned environment lifecycle state store under a controlled state root, with readiness transitions consumed by execution readiness and KDE-safe summaries.

Hard constraints:
- Do not start Wine, Proton, VM, or any compatibility backend.
- Do not expose raw backend commands, executable paths, profile names, storage paths, or host paths in user-facing output.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add missing, staged, ready, repair-required, and blocked-state tests.
- Run the M3 required tests plus ruby scripts/verify_layout.rb.
```

## M4: Portal and Snapshot Safety Plane

### Mission

Make permission requests and restore points real enough for execution readiness while remaining fake-first and state-root scoped.

### Start from

- `internal/runtime/portal/`
- `internal/runtime/snapshot/`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/portal_request.go`
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

### Keep disabled

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

### Copyable Claude Code prompt

```text
Implement M4 from docs/claude-code-mainline-implementation-plan.md.

Target outcome:
- Add a fake-first Portal request state machine and state-root-scoped snapshot/rollback safety plane that execution readiness can consume without starting execution.

Hard constraints:
- Do not enable real XDG Desktop Portal calls by default.
- Do not access user documents directly, create system snapshots, perform host-root rollback, mutate the host root, or launch compatibility backends.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add Portal completion, denial, cancellation, expiration, snapshot-create, snapshot-verify, rollback, and path-escape blocked tests.
- Run the M4 required tests plus ruby scripts/verify_layout.rb.
```

## M5: KDE Activation Materialization

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

### Keep disabled

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
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby -Ilib test/test_file_association_model.rb
ruby scripts/container.rb runtime-activation-smoke
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement M5 from docs/claude-code-mainline-implementation-plan.md.

Target outcome:
- Add or strengthen target-root-scoped KDE activation materialization for desktop entries, MIME associations, Dolphin service menus, manifests, and rollback receipts.

Hard constraints:
- Do not write to the host root, enable production desktop activation by default, launch compatibility backends from generated desktop files, apply live KWin rules, or enable live tray bridges.
- Keep generated user-facing output free of raw backend commands, executable paths, storage paths, and profile terminology.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add deterministic write, no-overwrite, manifest-digest, rollback, and path-escape blocked tests.
- Run the M5 required tests plus ruby scripts/verify_layout.rb.
```

## M6: Execution Transaction Ledger

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

### Keep disabled

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

### Copyable Claude Code prompt

```text
Implement M6 from docs/claude-code-mainline-implementation-plan.md.

Target outcome:
- Add a state-root-scoped execution transaction ledger that records request, review, preflight, grant proposal, transaction state, and fake committed session identity without launching real backends.

Hard constraints:
- Do not start Wine, Proton, VM, or any compatibility backend by default.
- Do not expose raw launch commands, enable production Launch writes, grant real Portal permissions, or mutate the host root.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add request, review-required, blocked, ready, fake-commit, failed, cancelled, missing-permission, missing-snapshot, untrusted-recipe, and blocked-backend tests.
- Run the M6 required tests plus ruby scripts/verify_layout.rb.
```

## M7: Diagnostics, Repair, and AI Boundary

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

### Keep disabled

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

### Copyable Claude Code prompt

```text
Implement M7 from docs/claude-code-mainline-implementation-plan.md.

Target outcome:
- Add fixture-driven diagnostic records, repair recommendation records, and a disabled-by-default AI provider boundary that can feed Compatibility Center history and execution readiness.

Hard constraints:
- Do not call real AI providers by default.
- Do not read arbitrary user file contents, execute automatic repairs, launch compatibility backends, require network access, or expose secrets, host paths, backend commands, or file contents.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add diagnostic fixture, failing result, repair recommendation, provider-disabled, provider-fake, approval-required, snapshot-required, and redaction tests.
- Run the M7 required tests plus ruby scripts/verify_layout.rb.
```

## M8: Developer Verification Harness

### Mission

Make it difficult for the project to accumulate more contracts without implementation evidence.

### Start from

- `scripts/verify_layout.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/implementation_evidence_report.rb`
- `scripts/kde_first_presence_smoke.rb`
- `scripts/container.rb`
- `test/test_runtime_contract_drift_report.rb`
- `test/test_implementation_evidence_report.rb`
- `docs/kde-first-current-gap-audit.md`

### Deliver

- A machine-readable implementation evidence report that classifies domains as `contract-only`, `fixture-implemented`, `state-root-implemented`, `smoke-owned`, or `production-gated`.
- A Markdown report for humans.
- Drift gates for D-Bus XML, owner routes, CLI commands, smoke adapters, and docs references.
- A check that highlights new preview methods with no implementation package or evidence owner.
- Optional report artifacts under a controlled `tmp/` or output path, never committed by default.

### Keep disabled

- CI secrets.
- Network-only checks.
- Host mutation.
- Broad filesystem scans outside the repository.
- Automatic deletion of user files.

### Acceptance

- The report exits non-zero on contract drift or orphan preview methods.
- The report does not require Docker, QEMU, network, privileged containers, or host-root access.
- Markdown output is stable enough to paste into pull requests.
- The report can be used before assigning Claude Code work to pick the next highest-value empty domain.

### Required tests

```text
go test ./...
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby -Ilib test/test_implementation_evidence_report.rb
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/runtime_contract_drift_report.rb --format markdown
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/implementation_evidence_report.rb --format markdown
ruby scripts/kde_first_presence_smoke.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement M8 from docs/claude-code-mainline-implementation-plan.md.

Target outcome:
- Add or strengthen implementation-evidence reporting that classifies domains as contract-only, fixture-implemented, state-root-implemented, smoke-owned, or production-gated.

Hard constraints:
- The default report must not require Docker, QEMU, network, privileged containers, host-root access, or broad filesystem scans outside the repository.
- Do not delete, rewrite, or auto-fix user files.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add JSON, Markdown, orphan-contract, and no-host-mutation tests.
- Run the M8 required tests plus ruby scripts/verify_layout.rb.
```

## M9: Atomic KDE Image and QEMU Acceptance

### Mission

Create a reproducible product-image path that can prove the KDE-first shape in constrained QEMU without endangering the host.

### Start from

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

### Keep disabled

- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Unpinned base images or package sources.
- Host system changes.

### Acceptance

- Image validation can run without network.
- Build and smoke scripts fail closed when unsafe Docker or QEMU options are requested.
- QEMU logs are persisted under a controlled output directory.
- Reports identify whether Runtime owner, KDE shell, activation, SSH, and serial evidence are present.

### Required tests

```text
go test ./internal/runtime/image ./...
ruby -Ilib test/test_kde_image.rb
ruby -Ilib test/test_kde_disk.rb
ruby -Ilib test/test_qemu.rb
ruby -Ilib test/test_serial_log.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement M9 from docs/claude-code-mainline-implementation-plan.md.

Target outcome:
- Add a safe product-image validation and constrained QEMU smoke path that can prove KDE, Runtime owner, activation, SSH loopback, and serial-log evidence without endangering the host.

Hard constraints:
- Do not use privileged containers, host networking, Docker socket mounts, broad host-directory mounts, unpinned base inputs, or host system changes.
- Keep QEMU logs under a controlled output directory and keep SSH forwarding loopback-only.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add image-manifest validation, unsafe-Docker-option blocked, unsafe-QEMU-option blocked, serial-log, loopback-SSH, and report-format tests.
- Run the M9 required tests plus ruby scripts/verify_layout.rb.
```

## Per-Branch Handoff Checklist

Claude Code should include this checklist in each branch handoff:

- [ ] Exactly one mainline package was implemented.
- [ ] The branch explains which contracts were converted into implementation evidence.
- [ ] Unsafe actions remain disabled or gated as required.
- [ ] Targeted tests cover success, failure, and blocked states.
- [ ] `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` were updated for code changes.
- [ ] Package tests and `ruby scripts/verify_layout.rb` passed.
- [ ] Contract drift reporting passed if Runtime routes or D-Bus surfaces changed.
- [ ] The branch did not touch unrelated packages, secrets, build artifacts, host configuration, `.claude/`, or `tmp/`.
