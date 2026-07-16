# Claude Code Mainline Task Batch

> Last updated: 2026-07-16 | Baseline: v0.2.293

This document contains copyable Claude Code tasks for the KDE-first Windows application compatibility mainline.

Use one task per Claude Code branch. Do not ask one branch to implement the entire product.

Do not modify `docs/claude-code-implementation-packages.md` from any task in this batch.

## Global Rules for Every Task

Every Claude Code branch must:

- Keep source, comments, fixtures, tests, CLI output, and documentation in English.
- Prefer Go for durable Runtime product logic.
- Use C only for low-level D-Bus transport, ABI-shaped records, or already-owned Runtime policy surfaces.
- Use Ruby for tests, smoke scripts, reports, and lightweight developer tooling.
- Keep KDE as the shell and presentation layer only.
- Keep the Xnix Compatibility Runtime as the policy and product-logic owner.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for code changes.
- Add success, failure, and blocked-state tests.
- Run targeted verification plus `ruby scripts/verify_layout.rb`.
- Stop and report if the task needs cross-lane ownership or unsafe host behavior.

Unsafe behavior remains disabled:

- Production D-Bus ownership.
- Runtime write methods.
- Real Wine, Proton, VM, or compatibility backend launch.
- Real XDG Portal transport calls.
- Default network artifact fetch.
- Host package-manager calls.
- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Host-root mutation.
- Raw backend commands in KDE-facing output.
- Raw executable paths, Runtime state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials in user-facing output.

## Recommended Dispatch Order

| Order | Task | Lane | Why now |
| --- | --- | --- | --- |
| 1 | `CB1` Runtime owner private bus service hardening | `CW1 / A1` | Later work needs a real constrained read boundary. |
| 2 | `CB2` Recipe and artifact receipt end-to-end gate | `CW2 / A2` | Desktop and execution work must depend on verified inputs. |
| 3 | `CB3` Backend manager lifecycle repair states | `CW3 / A3` | KDE status and execution readiness need durable backend state. |
| 4 | `CB4` Portal permission review ledger | `CW5 / A4` | Execution and file access must have reviewable permission evidence. |
| 5 | `CB5` Snapshot and rollback receipt store | `CW6 / A4` | Risky compatibility changes need reversible state before launch. |
| 6 | `CB6` KDE entry-point evidence consumers | `CW4 / A5` | KDE should read durable Runtime evidence, not static previews. |
| 7 | `CB7` AI diagnostic privacy boundary | `CW7 / A7` | AI features need redaction and review gates before provider calls. |
| 8 | `CB8` Mainline integration review gate | `CW10 / A8` | Mixed Claude/Codex output needs lane-level review evidence. |
| 9 | `CB9` Restricted product smoke readiness | `CW11 / A9` | Product-level smoke should wait until Runtime evidence is meaningful. |

## CB1: Runtime Owner Private Bus Service Hardening

### Copyable Prompt

```text
Implement CB1 from docs/claude-code-mainline-task-batch.md.

Goal:
- Strengthen the constrained Go Runtime owner read boundary so a private session-bus smoke path proves read-only ownership, disabled write behavior, unsupported-read behavior, and shutdown safety.

Start from:
- cmd/xnix-runtime-owner/
- internal/runtime/owner/
- internal/runtime/appidentity/runtime_owner_*.go
- scripts/runtime_contract_drift_report.rb
- scripts/runtime_owner_candidate_smoke.rb
- test/test_runtime_contract_drift_report.rb
- test/test_runtime_owner_candidate_smoke_script.rb

Deliver:
- Stronger private session-bus lifecycle evidence for startup, bus claim, route-table readiness, read dispatch, disabled writes, unsupported reads, and shutdown.
- A blocked production-owner state that is explicit in JSON output.
- Tests proving production D-Bus ownership, write methods, backend launch, network access, and host-root mutation remain disabled.
- Drift report coverage for any owner route, CLI, or smoke transcript change.

Do not deliver:
- Production D-Bus ownership.
- System service installation.
- Write enablement.
- Backend launch.
- Host-root mutation.

Required verification:
- go test ./cmd/xnix-runtime-owner ./internal/runtime/owner ./internal/runtime/appidentity
- ruby scripts/runtime_contract_drift_report.rb --format json
- ruby -Ilib test/test_runtime_contract_drift_report.rb
- ruby -Ilib test/test_runtime_owner_candidate_smoke_script.rb
- ruby scripts/verify_layout.rb
```

## CB2: Recipe and Artifact Receipt End-to-End Gate

### Copyable Prompt

```text
Implement CB2 from docs/claude-code-mainline-task-batch.md.

Goal:
- Make compatibility install readiness consume verified recipe and artifact staging receipts end to end.

Start from:
- internal/runtime/recipe/
- internal/runtime/artifact/
- internal/runtime/appidentity/install_plan.go
- internal/runtime/appidentity/runtime_owner_recipe_trust.go
- cmd/xnix-runtime-go/install_plan_commands.go
- cmd/xnix-runtime-go/artifact_stage_commands.go
- runtime/recipes/

Deliver:
- A single local fixture flow that verifies recipe identity, artifact digest, staging receipt integrity, required artifact coverage, and install-readiness blocking reasons.
- A tampered receipt flow that fails closed for digest mismatch, path escape, missing staged artifacts, unsafe side-effect flags, and app-id mismatch.
- CLI coverage showing compatibility-install-preview can consume an artifact receipt without exposing cache roots, fixture roots, host paths, backend commands, or raw executable paths.
- Evidence report tokens for the end-to-end trust path.

Do not deliver:
- Network artifact fetch.
- Host package-manager calls.
- Desktop activation writes.
- Runtime write methods.
- Backend launch.
- Host-root mutation.

Required verification:
- go test ./internal/runtime/recipe ./internal/runtime/artifact ./internal/runtime/appidentity ./cmd/xnix-runtime-go
- ruby -Ilib test/test_recipe_registry.rb
- ruby -Ilib test/test_recipe_trust_policy.rb
- ruby -Ilib test/test_recipe_install_gate.rb
- ruby -Ilib test/test_compatibility_artifact_manifest.rb
- ruby -Ilib test/test_compatibility_install_plan.rb
- ruby scripts/implementation_evidence_report.rb --format json
- ruby scripts/verify_layout.rb
```

## CB3: Backend Manager Lifecycle Repair States

### Copyable Prompt

```text
Implement CB3 from docs/claude-code-mainline-task-batch.md.

Goal:
- Strengthen Runtime-owned backend manager and lifecycle records so KDE and execution readiness can consume durable state without launching Wine, Proton, or Windows VM backends.

Start from:
- internal/runtime/appidentity/backend_manager.go
- internal/runtime/appidentity/backend_lifecycle.go
- internal/runtime/environment/
- cmd/xnix-runtime-go/backend_group_cli_test.go
- test/test_compatibility_backend_lifecycle.rb
- test/test_compatibility_backend_selection_plan.rb

Deliver:
- State-root records for backend inventory and lifecycle transitions.
- Repair-required and blocked states with stable reasons and user-safe summaries.
- CLI actions for inspect, plan, stage, satisfy-gate, mark-ready, flag-repair, block, and retire if they are not already complete.
- Tests proving state-root paths, backend commands, profile paths, and raw backend names are not exposed to KDE-facing output.

Do not deliver:
- Wine, Proton, or VM process start.
- Backend downloads or installs.
- Network access.
- Host-root mutation.
- Raw backend command exposure.

Required verification:
- go test ./internal/runtime/environment ./internal/runtime/appidentity ./cmd/xnix-runtime-go
- ruby -Ilib test/test_compatibility_backend_lifecycle.rb
- ruby -Ilib test/test_compatibility_backend_selection_plan.rb
- ruby -Ilib test/test_compatibility_backend_environment_plan.rb
- ruby scripts/implementation_evidence_report.rb --format json
- ruby scripts/verify_layout.rb
```

## CB4: Portal Permission Review Ledger

### Copyable Prompt

```text
Implement CB4 from docs/claude-code-mainline-task-batch.md.

Goal:
- Build a fake-mode Portal permission review ledger that execution and KDE views can consume before real XDG Portal transport exists.

Start from:
- internal/runtime/portal/
- internal/runtime/appidentity/portal_access_policy.go
- internal/runtime/appidentity/permission or review flow models
- cmd/xnix-runtime-go/runtime_safety_commands.go
- test/test_portal_access_policy.rb
- test/test_portal_request_model.rb

Deliver:
- State-root Portal request records for create, inspect, grant, deny, complete, cancel, and expire flows.
- Receipt fields that execution preflight can consume without granting real host permissions.
- Tests for granted, denied, expired, cancelled, and malformed receipts.
- User-safe KDE summaries such as Documents allowed, Camera blocked, and Network allowed.

Do not deliver:
- Real XDG Portal calls.
- Host permission changes.
- Execution approval.
- Backend launch.
- Host-root mutation.

Required verification:
- go test ./internal/runtime/portal ./internal/runtime/appidentity ./cmd/xnix-runtime-go
- ruby -Ilib test/test_portal_access_policy.rb
- ruby -Ilib test/test_portal_request_model.rb
- ruby -Ilib test/test_compatibility_permission_review_plan.rb
- ruby scripts/implementation_evidence_report.rb --format json
- ruby scripts/verify_layout.rb
```

## CB5: Snapshot and Rollback Receipt Store

### Copyable Prompt

```text
Implement CB5 from docs/claude-code-mainline-task-batch.md.

Goal:
- Make Runtime snapshot and rollback planning depend on content-addressed, state-root-safe receipts before risky compatibility changes can be represented as ready.

Start from:
- internal/runtime/snapshot/
- internal/runtime/appidentity/snapshot_plan.go
- test/test_compatibility_snapshot_plan.rb
- runtime/core snapshot policy files

Deliver:
- Snapshot receipt records with schema version, app id, operation id, content digest, relative receipt path, blocked reasons, and rollback readiness.
- Rollback receipt preview that stays disabled until safe receipt evidence exists.
- Tests for valid receipts, digest mismatch, missing content, path escape, and rollback blocked states.

Do not deliver:
- Real filesystem rollback.
- Host-root mutation.
- Backend stop/start.
- Raw state-root path exposure.

Required verification:
- go test ./internal/runtime/snapshot ./internal/runtime/appidentity
- ruby -Ilib test/test_compatibility_snapshot_plan.rb
- ruby -Ilib test/test_runtime_core.rb
- ruby scripts/implementation_evidence_report.rb --format json
- ruby scripts/verify_layout.rb
```

## CB6: KDE Entry-Point Evidence Consumers

### Copyable Prompt

```text
Implement CB6 from docs/claude-code-mainline-task-batch.md.

Goal:
- Make KDE entry-point previews consume durable Runtime receipts instead of restating static preview data.

Start from:
- internal/runtime/appidentity/kde_*.go
- internal/runtime/appidentity/window_identity_routes.go
- cmd/xnix-runtime-go/kde_* commands and tests
- cmd/xnix-runtime-go/window_identity_commands.go
- kde/plasmoids/org.xnix.compatibilitycenter/
- test/test_kde_*.rb
- test/test_task_manager_identity.rb
- test/test_kwin_window_rule.rb
- test/test_tray_status_model.rb
- test/test_file_association_model.rb

Deliver:
- KDE-safe read models for the seven first-release entry points: start menu, task manager, file manager, tray, notifications, AI Compatibility Center, and unified settings.
- Consumption of durable Runtime evidence where available: activation receipts, execution session records, Portal receipts, backend lifecycle records, and diagnostics history.
- Tests proving KDE output hides Wine prefix terminology, raw executable paths, Runtime state-root paths, host paths, backend commands, and file contents.

Do not deliver:
- KDE source forks.
- KWin rule application.
- Desktop file writes outside explicit staging roots.
- Live tray bridge activation.
- Backend launch.

Required verification:
- go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
- ruby -Ilib test/test_kde_integration_status.rb
- ruby -Ilib test/test_kde_shell_integration_plan.rb
- ruby -Ilib test/test_kde_application_surface_plan.rb
- ruby -Ilib test/test_kde_center_model.rb
- ruby -Ilib test/test_task_manager_identity.rb
- ruby -Ilib test/test_kwin_window_rule.rb
- ruby -Ilib test/test_tray_status_model.rb
- ruby -Ilib test/test_file_association_model.rb
- ruby scripts/verify_layout.rb
```

## CB7: AI Diagnostic Privacy Boundary

### Copyable Prompt

```text
Implement CB7 from docs/claude-code-mainline-task-batch.md.

Goal:
- Strengthen the AI diagnostic boundary so diagnostic inputs are redacted, reviewable, fixture-backed, and never sent to a real provider by default.

Start from:
- internal/runtime/diagnostics/
- internal/runtime/appidentity/ai_diagnostics.go
- internal/runtime/appidentity/diagnostic_history.go
- cmd/xnix-runtime-go/diagnostic_record_commands.go
- test/test_ai_diagnostic_input.rb
- test/test_ai_diagnostic_recommendation.rb
- test/test_ai_repair_approval_gate.rb

Deliver:
- State-root diagnostic records with redacted input payloads, privacy filters, signal ids, and repair recommendation receipts.
- Tests proving provider calls, auto-repair, backend launch, file-content exposure, root-path exposure, and secrets exposure remain disabled.
- Compatibility Center summaries that show useful status without exposing logs or raw file contents.

Do not deliver:
- Real AI provider calls.
- Auto-repair execution.
- Secret capture.
- Backend launch.
- Host-root mutation.

Required verification:
- go test ./internal/runtime/diagnostics ./internal/runtime/appidentity ./cmd/xnix-runtime-go
- ruby -Ilib test/test_ai_diagnostic_input.rb
- ruby -Ilib test/test_ai_diagnostic_recommendation.rb
- ruby -Ilib test/test_ai_repair_approval_gate.rb
- ruby -Ilib test/test_compatibility_repair_plan.rb
- ruby scripts/implementation_evidence_report.rb --format json
- ruby scripts/verify_layout.rb
```

## CB8: Mainline Integration Review Gate

### Copyable Prompt

```text
Implement CB8 from docs/claude-code-mainline-task-batch.md.

Goal:
- Turn the current mainline integration checkpoint into a stronger review gate that classifies worktree changes by lane and prevents accidental all-in-one staging.

Start from:
- docs/mainline-integration-checkpoint.md
- scripts/mainline_integration_review.rb
- test/test_mainline_integration_review.rb
- scripts/implementation_evidence_report.rb
- scripts/verify_layout.rb

Deliver:
- JSON and Markdown reports that group changed files into the checkpoint lanes.
- Explicit exclusion of .gocache/, tmp/, build artifacts, local logs, and private config.
- Explicit blocking status when docs/claude-code-implementation-packages.md is modified.
- Tests for fixture-driven status parsing, lane classification, protected-file detection, unclassified files, and Markdown output.

Do not deliver:
- Git staging or committing from the script.
- Docker or QEMU execution.
- Network access.
- Host-root mutation.

Required verification:
- ruby scripts/mainline_integration_review.rb --format json
- ruby scripts/mainline_integration_review.rb --format markdown
- ruby -Ilib test/test_mainline_integration_review.rb
- ruby scripts/implementation_evidence_report.rb --format json
- ruby scripts/verify_layout.rb
```

## CB9: Restricted Product Smoke Readiness

### Copyable Prompt

```text
Implement CB9 from docs/claude-code-mainline-task-batch.md only after CB1, CB2, CB3, CB4, and CB8 are merged.

Goal:
- Prepare the restricted Docker/QEMU product smoke so it validates real Runtime evidence instead of only static contracts.

Start from:
- scripts/full_smoke.rb
- lib/xnix/full_smoke_report.rb
- scripts/container.rb
- scripts/build_kde_image.rb
- scripts/boot_kde_image.rb
- test/test_full_smoke_report.rb
- test/test_full_smoke_script.rb

Deliver:
- A dry-run or fixture-mode smoke report that checks Runtime owner evidence, artifact trust evidence, backend lifecycle evidence, Portal safety evidence, and KDE entry-point evidence without requiring privileged containers.
- Clear report fields for loopback-only networking, no Docker socket mount, no host networking, no broad host mount, no host-root mutation, and persisted serial logs.
- Tests that can run without Docker or QEMU.

Do not deliver:
- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Host-root mutation.
- Real backend launch.

Required verification:
- ruby -Ilib test/test_full_smoke_report.rb
- ruby -Ilib test/test_full_smoke_script.rb
- ruby scripts/implementation_evidence_report.rb --format json
- ruby scripts/verify_layout.rb

Do not run Docker or QEMU unless the user explicitly authorizes that run.
```

## Completion Template for Claude Code

Claude Code should return:

```text
Task:
Branch:
Version:
Files changed:
What moved from contract-only to implemented evidence:
Safety gates kept disabled:
Verification commands run:
Commands not run and why:
Known follow-up:
```
