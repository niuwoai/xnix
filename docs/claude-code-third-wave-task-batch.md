# Claude Code Third-Wave Task Batch

> Last updated: 2026-07-16 | Baseline under review: v0.2.293

This document contains a third batch of bounded Claude Code tasks for the KDE-first Windows application compatibility mainline.

Use this batch after the current mainline and second-wave work has been merged or deliberately skipped. The goal is to turn more cross-cutting contracts into Runtime-owned evidence while keeping execution, host mutation, and heavy image tests explicitly gated.

Do not modify `docs/claude-code-implementation-packages.md` from any task in this batch. That file is Claude-owned.

## Global Rules for Every Third-Wave Task

Every Claude Code branch must:

- Pick exactly one task from `TW1` through `TW8`.
- Keep source, comments, fixtures, tests, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior.
- Use C only for low-level D-Bus transport, ABI-shaped records, or already-owned Runtime policy surfaces.
- Use Ruby only for tests, smoke scripts, reports, and lightweight developer tooling.
- Keep KDE as shell and presentation only.
- Keep the Xnix Compatibility Runtime as the policy and product-logic owner.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for code changes.
- Add success, failure, malformed-input, and blocked-state tests.
- Run the task-specific verification commands plus `ruby scripts/verify_layout.rb`.
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
- Docker or QEMU execution unless a human explicitly authorizes a restricted smoke.
- Host-root mutation.
- Raw backend commands in KDE-facing output.
- Raw executable paths, Runtime state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials in user-facing output.

## Dispatch Order

| Order | Task | Primary lane | Why now |
| --- | --- | --- | --- |
| 1 | `TW1` Runtime route convergence plan | `CW1 / CW10` | Many routes are implemented as previews, but merge review needs a concrete migration ledger for remaining C/Ruby-owned routes. |
| 2 | `TW2` Action queue receipt dependency graph | `CW4 / CW5 / CW8` | Compatibility Center actions need durable prerequisites before any action can become executable. |
| 3 | `TW3` Portal-to-execution preflight join | `CW5 / CW8` | Execution readiness should consume fake-mode Portal receipts before launch is ever enabled. |
| 4 | `TW4` Snapshot baseline readiness join | `CW6 / CW8` | Risky app operations should depend on rollback evidence before later execution work. |
| 5 | `TW5` AI diagnostics privacy corpus | `CW7` | AI diagnostics need fixture-based redaction and review boundaries before any provider integration. |
| 6 | `TW6` KDE action-card evidence deck | `CW4` | KDE should render action cards from Runtime evidence instead of independent static summaries. |
| 7 | `TW7` Runtime owner unsupported-route hardening | `CW1 / CW10` | Owner dispatch must fail closed for unknown, deprecated, or write-shaped routes. |
| 8 | `TW8` Restricted acceptance dashboard | `CW10 / CW11` | Developers need one safe dashboard that summarizes local smoke, image preflight, and blocked heavy tests without running Docker or QEMU. |

Do not dispatch `TW8` as a real Docker or QEMU run. `TW8` is a dashboard and readiness-report task only unless a human explicitly authorizes restricted execution.

## TW1: Runtime Route Convergence Plan

### Mission

Create a Runtime-owned convergence report that shows which read-only methods are implemented by Go product logic, C low-level policy, Ruby smoke adapters, or contract-only fixtures, then generate a safe next-migration order.

### Start from

- `internal/runtime/appidentity/runtime_owner_route_manifest.go`
- `internal/runtime/appidentity/runtime_method_parity_manifest.go`
- `cmd/xnix-runtime-go/runtime_owner_route_manifest_cli_test.go`
- `cmd/xnix-runtime-go/runtime_method_parity_manifest_cli_test.go`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/implementation_evidence_report.rb`
- `test/test_runtime_contract_drift_report.rb`
- `test/test_implementation_evidence_report.rb`

### Deliver

- A Go read model and CLI command such as `runtime-route-convergence-preview`.
- Route classifications for Go product logic, C policy bridge, Ruby smoke bridge, fixture-only, contract-only, deprecated, and unsupported.
- Stable blocked reasons for routes that cannot be migrated yet.
- A recommended migration order grouped by domain and safety risk.
- Report coverage that fails closed when a route is unclassified.

### Keep disabled

- Production D-Bus ownership.
- Automatic route migration.
- Write methods.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/implementation_evidence_report.rb --format json
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby -Ilib test/test_implementation_evidence_report.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement TW1 from docs/claude-code-third-wave-task-batch.md.

Target outcome:
- Add a Runtime route convergence preview that classifies every read-only method by implementation ownership and emits a safe next-migration order.

Hard constraints:
- Do not enable production D-Bus ownership, Runtime write methods, backend launch, network fetch, privileged containers, Docker/QEMU execution, or host-root mutation.
- Do not expose raw backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add classified-route, unclassified-route, deprecated-route, unsupported-route, and JSON report tests.
- Run the TW1 verification commands and report exact commands run.
```

## TW2: Action Queue Receipt Dependency Graph

### Mission

Make Compatibility Center action cards depend on explicit Runtime evidence prerequisites instead of static task descriptions.

### Start from

- `internal/runtime/appidentity/kde_action_*.go`
- `internal/runtime/appidentity/kde_center_page.go`
- `internal/runtime/appidentity/application_readiness.go`
- `internal/runtime/appidentity/execution_session_record_evidence.go`
- `cmd/xnix-runtime-go/kde_center_cli_test.go`
- `cmd/xnix-runtime-go/kde_center_page_cli_test.go`
- `test/test_compatibility_action_queue.rb`
- `test/test_action_review_receipt.rb`
- `test/test_kde_center_model.rb`

### Deliver

- A Runtime action dependency graph for install, permission review, snapshot baseline, repair suggestion, service binding, and execution review actions.
- Action cards that name required evidence, missing evidence, blocked reasons, and next safe read-only checks.
- Receipt validation that rejects mismatched app ids, malformed operation ids, path escape evidence, and unsafe side-effect flags.
- Tests proving action cards never grant permissions, launch apps, persist settings, or mutate host state.

### Keep disabled

- Action execution.
- Settings persistence.
- Permission grants.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_action_queue.rb
ruby -Ilib test/test_action_review_receipt.rb
ruby -Ilib test/test_kde_center_model.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement TW2 from docs/claude-code-third-wave-task-batch.md.

Target outcome:
- Make Compatibility Center action cards consume explicit Runtime evidence prerequisites and blocked reasons.

Hard constraints:
- Do not execute actions, persist settings, grant permissions, start backends, call real Portal transport, or mutate the host root.
- Do not expose raw backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add valid-dependency, missing-evidence, malformed-receipt, mismatched-app, blocked-action, and no-side-effect tests.
- Run the TW2 verification commands and report exact commands run.
```

## TW3: Portal-to-Execution Preflight Join

### Mission

Make execution preflight consume fake-mode Portal permission receipts before any launch path can become ready.

### Start from

- `internal/runtime/portal/`
- `internal/runtime/execution/`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `internal/runtime/appidentity/application_readiness.go`
- `cmd/xnix-runtime-go/runtime_safety_commands.go`
- `cmd/xnix-runtime-go/execution_ledger_commands.go`
- `test/test_portal_request_model.rb`
- `test/test_compatibility_permission_review_plan.rb`
- `test/test_compatibility_execution_readiness.rb`

### Deliver

- A preflight join that maps Portal receipt states into execution readiness: granted, denied, expired, cancelled, pending, malformed, and missing.
- User-safe summaries for Documents, Downloads, Clipboard, Camera, Screen, Remote Desktop, URI, Print, and Network access.
- Tests proving granted fake-mode receipts do not create real host permissions.
- Tests proving denied or missing required Portal receipts block execution readiness.

### Keep disabled

- Real XDG Portal transport.
- Host permission changes.
- Execution approval.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/portal ./internal/runtime/execution ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_compatibility_permission_review_plan.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement TW3 from docs/claude-code-third-wave-task-batch.md.

Target outcome:
- Make execution readiness consume fake-mode Portal permission receipts and block readiness when required permission evidence is missing or denied.

Hard constraints:
- Do not call real XDG Portal transport, change host permissions, approve execution, start backends, enable Runtime writes, or mutate the host root.
- Do not expose raw backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add granted, denied, expired, cancelled, pending, malformed, missing, and no-host-permission tests.
- Run the TW3 verification commands and report exact commands run.
```

## TW4: Snapshot Baseline Readiness Join

### Mission

Make readiness and action planning require safe snapshot baseline evidence before risky compatibility changes are represented as eligible.

### Start from

- `internal/runtime/snapshot/`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/application_readiness.go`
- `internal/runtime/appidentity/kde_action_*.go`
- `cmd/xnix-runtime-go/`
- `test/test_compatibility_snapshot_plan.rb`
- `test/test_compatibility_execution_readiness.rb`
- `test/test_kde_center_model.rb`

### Deliver

- Snapshot baseline receipt validation with schema version, app id, operation id, content digest, relative receipt path, and rollback eligibility.
- Readiness states for baseline-ready, baseline-missing, digest-mismatch, rollback-blocked, stale-baseline, and malformed receipt.
- KDE-safe action cards that tell the user a restore point is needed without exposing state-root or host paths.
- Tests proving no real snapshot or rollback occurs.

### Keep disabled

- Real snapshot creation.
- Real rollback.
- Backend stop/start.
- Host-root mutation.
- Raw state-root path exposure.

### Required verification

```text
go test ./internal/runtime/snapshot ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby -Ilib test/test_kde_center_model.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement TW4 from docs/claude-code-third-wave-task-batch.md.

Target outcome:
- Make application readiness and KDE action planning consume snapshot baseline receipts before risky compatibility changes are represented as eligible.

Hard constraints:
- Do not create snapshots, run rollback, stop or start backends, expose state-root paths, enable Runtime writes, or mutate the host root.
- Do not expose raw backend commands, raw executable paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add baseline-ready, missing, digest-mismatch, rollback-blocked, stale-baseline, malformed-receipt, and no-real-snapshot tests.
- Run the TW4 verification commands and report exact commands run.
```

## TW5: AI Diagnostics Privacy Corpus

### Mission

Build a fixture-based AI diagnostics privacy boundary so diagnostic recommendations can be tested without calling any external AI provider.

### Start from

- `internal/runtime/appidentity/ai_diagnostic_*.go`
- `cmd/xnix-runtime-go/runtime_safety_commands.go`
- `test/test_ai_diagnostic_input.rb`
- `test/test_ai_diagnostic_recommendation.rb`
- `test/test_ai_repair_approval_gate.rb`
- `scripts/implementation_evidence_report.rb`

### Deliver

- A redaction corpus for paths, usernames, emails, tokens, URLs with secrets, raw commands, crash snippets, and file contents.
- Fixture-only diagnostic input and recommendation previews that prove provider calls remain disabled.
- Review-only repair recommendations that require explicit human approval and cannot execute repairs.
- Implementation evidence coverage for redaction, provider-disabled state, and approval-gate state.

### Keep disabled

- External AI provider calls.
- Network access.
- Automatic repair.
- Settings persistence.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_ai_repair_approval_gate.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement TW5 from docs/claude-code-third-wave-task-batch.md.

Target outcome:
- Add a fixture-based AI diagnostics privacy corpus and prove diagnostic recommendations remain redacted, provider-disabled, and review-only.

Hard constraints:
- Do not call external AI providers, fetch network resources, run automatic repair, persist settings, enable Runtime writes, or mutate the host root.
- Do not expose raw backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, emails, or sensitive URLs.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add redaction corpus, provider-disabled, malformed-input, review-only, approval-required, and no-network tests.
- Run the TW5 verification commands and report exact commands run.
```

## TW6: KDE Action-Card Evidence Deck

### Mission

Create a Runtime-owned KDE action-card deck that groups readiness, permission, snapshot, diagnostic, service, and execution evidence into user-safe cards.

### Start from

- `internal/runtime/appidentity/kde_center_page.go`
- `internal/runtime/appidentity/kde_action_*.go`
- `internal/runtime/appidentity/application_readiness.go`
- `internal/runtime/appidentity/runtime_service_binding*.go`
- `cmd/xnix-runtime-go/kde_center_page_cli_test.go`
- `test/test_kde_center_model.rb`
- `test/test_compatibility_action_queue.rb`
- `test/test_runtime_service_binding.rb`

### Deliver

- A deck schema such as `xnix.runtime.kde_action_card_deck.v1`.
- Card groups for readiness, permissions, snapshots, diagnostics, service binding, and execution review.
- Stable card states: ready, blocked, waiting, review-required, disabled, malformed-evidence.
- Tests proving deck output is KDE-safe and never reveals backend details or host paths.

### Keep disabled

- Card action execution.
- Settings persistence.
- Service installation.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_kde_center_model.rb
ruby -Ilib test/test_compatibility_action_queue.rb
ruby -Ilib test/test_runtime_service_binding.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement TW6 from docs/claude-code-third-wave-task-batch.md.

Target outcome:
- Add a Runtime-owned KDE action-card deck that groups readiness, permissions, snapshots, diagnostics, service binding, and execution review evidence into safe user-facing cards.

Hard constraints:
- Do not execute card actions, persist settings, install services, start backends, enable Runtime writes, or mutate the host root.
- Do not expose raw backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add ready, blocked, waiting, review-required, disabled, malformed-evidence, and hidden-detail tests.
- Run the TW6 verification commands and report exact commands run.
```

## TW7: Runtime Owner Unsupported-Route Hardening

### Mission

Make the constrained Runtime owner fail closed for unknown, deprecated, unsupported, malformed, and write-shaped routes with structured evidence.

### Start from

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `internal/runtime/appidentity/runtime_owner_route_manifest.go`
- `internal/runtime/appidentity/runtime_write_gate.go`
- `scripts/runtime_owner_candidate_smoke.rb`
- `scripts/runtime_contract_drift_report.rb`
- `test/test_runtime_owner_candidate_smoke_script.rb`
- `test/test_runtime_contract_drift_report.rb`

### Deliver

- Structured unsupported-route receipts with route name, classification, blocked reason, write-gate state, and safety flags.
- Tests for unknown reads, deprecated reads, malformed route names, write-shaped route names, and disabled write methods.
- Smoke evidence that unsupported routes cannot fall through to shell commands, backend launch, network access, or host mutation.

### Keep disabled

- Production D-Bus ownership.
- Runtime write methods.
- Shell fallback.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./cmd/xnix-runtime-owner ./internal/runtime/owner ./internal/runtime/appidentity
ruby scripts/runtime_contract_drift_report.rb --format json
ruby -Ilib test/test_runtime_owner_candidate_smoke_script.rb
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement TW7 from docs/claude-code-third-wave-task-batch.md.

Target outcome:
- Make the constrained Runtime owner fail closed for unknown, deprecated, unsupported, malformed, and write-shaped routes with structured evidence.

Hard constraints:
- Do not enable production D-Bus ownership, Runtime write methods, shell fallback, backend launch, network fetch, Docker/QEMU execution, or host-root mutation.
- Do not expose raw backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add unknown-read, deprecated-read, malformed-route, write-shaped-route, disabled-write, and no-fallback tests.
- Run the TW7 verification commands and report exact commands run.
```

## TW8: Restricted Acceptance Dashboard

### Mission

Create a safe local dashboard that summarizes restricted local smoke, contract drift, implementation evidence, mainline review, and image acceptance preflight without running Docker or QEMU by default.

### Start from

- `scripts/restricted_local_smoke.rb`
- `scripts/image_acceptance_preflight.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/implementation_evidence_report.rb`
- `scripts/mainline_integration_review.rb`
- `scripts/full_smoke.rb`
- `test/test_full_smoke_report.rb`
- `test/test_mainline_integration_review.rb`

### Deliver

- A dashboard script such as `scripts/restricted_acceptance_dashboard.rb`.
- JSON and Markdown output.
- Summary sections for local smoke, contract drift, implementation evidence, mainline lane review, image preflight, blocked heavy tests, and recommended next safe command.
- Tests proving the default dashboard never runs Docker, Colima, QEMU, network fetch, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.

### Keep disabled

- Docker or Colima execution by default.
- QEMU execution by default.
- Network fetch.
- Host package-manager calls.
- Destructive cleanup.
- Host-root mutation.

### Required verification

```text
ruby -Ilib test/test_full_smoke_report.rb
ruby -Ilib test/test_mainline_integration_review.rb
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/mainline_integration_review.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement TW8 from docs/claude-code-third-wave-task-batch.md.

Target outcome:
- Add a restricted acceptance dashboard that summarizes local smoke, contract drift, implementation evidence, mainline review, and image preflight without running Docker or QEMU by default.

Hard constraints:
- Do not run Docker, Colima, QEMU, network fetches, host package-manager commands, privileged containers, host networking, Docker socket mounts, broad host mounts, destructive cleanup, or host-root mutation from the default path.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add pass, fail, skipped-heavy-test, blocked-heavy-test, JSON report, Markdown report, and no-execution tests.
- Run the TW8 verification commands and report exact commands run.
```

## Claude Completion Template

Require Claude Code to finish every task with:

```text
Task:
Branch:
Version:
Lane:
Files changed:
Shared files changed and why:
What moved from contract-only to implementation evidence:
Safety gates kept disabled:
Verification commands run:
Commands not run and why:
Known follow-up:
```

## Local Intake Checklist

Before staging a Claude branch from this batch:

1. Confirm the branch implemented exactly one `TW` task.
2. Confirm `docs/claude-code-implementation-packages.md` has no diff.
3. Exclude `.gocache/`, `tmp/`, build artifacts, logs, private config, and generated local output.
4. Confirm `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` are updated for code changes.
5. Confirm all task-specific verification commands are reported.
6. Run `ruby scripts/mainline_integration_review.rb --format json`.
7. Confirm `protected_claude_file_modified` is `false`.
8. Confirm `unclassified_file_count` is `0`.
9. Run `ruby scripts/verify_layout.rb`.
10. Run `git diff --check`.

Never stage with `git add .`.
