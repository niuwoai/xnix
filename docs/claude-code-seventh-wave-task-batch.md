# Claude Code Seventh-Wave Task Batch

> Last updated: 2026-07-16 | Baseline under review: v0.2.293

This document contains a seventh batch of bounded Claude Code tasks for the KDE-first Windows application compatibility mainline.

Use this batch after the current mainline, second-wave, third-wave, fourth-wave, fifth-wave, and sixth-wave work has been merged or deliberately skipped. The goal is to turn remaining product-facing Runtime gaps into reviewable, offline, KDE-safe implementation evidence without enabling real execution, real Portal transport, Docker, QEMU, network fetch, privileged containers, host-root mutation, or compatibility backend launch.

Do not modify `docs/claude-code-implementation-packages.md` from any task in this batch. That file is Claude-owned.

## Global Rules for Every Seventh-Wave Task

Every Claude Code branch must:

- Pick exactly one task from `S7W1` through `S7W8`.
- Keep source, comments, fixtures, tests, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior.
- Use C only for low-level D-Bus transport, ABI-shaped records, or already-owned Runtime policy surfaces.
- Use Ruby only for tests, smoke scripts, reports, and lightweight developer tooling.
- Keep KDE as shell and presentation only.
- Keep the Xnix Compatibility Runtime as the policy and product-logic owner.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for code changes.
- Add success, failure, malformed-input, blocked-state, and no-sensitive-output tests.
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
- Raw executable paths, Runtime state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, or environment variables in user-facing output.

## Dispatch Order

| Order | Task | Primary lane | Why now |
| --- | --- | --- | --- |
| 1 | `S7W1` Application upgrade impact preview | `CW2 / CW6 / CW8` | Recipe and artifact updates need a user-safe impact summary before migrations or installs can advance. |
| 2 | `S7W2` Runtime policy explanation cards | `CW5 / CW7 / CW8` | KDE needs consistent explanations for why actions are blocked, review-only, or not yet implemented. |
| 3 | `S7W3` State-root quota and retention preview | `CW3 / CW6` | Per-application storage growth needs dry-run quota and retention evidence before cleanup can become real. |
| 4 | `S7W4` Crash and hang signal summary preview | `CW7 / CW8` | Diagnostic evidence should summarize crashes, hangs, and repeated failed launches without reading private logs. |
| 5 | `S7W5` Compatibility backend fallback preview | `CW3 / CW8` | Automatic mode needs a safe fallback explanation between local compatibility and isolated compatibility. |
| 6 | `S7W6` KDE search visibility plan | `CW4 / CW8` | Launcher, KRunner, file associations, and Compatibility Center search should share one Runtime-owned visibility model. |
| 7 | `S7W7` Permission evidence audit preview | `CW5 / CW10` | Portal and settings permissions need an offline audit before real permission renewal or revocation is allowed. |
| 8 | `S7W8` Release evidence index | `CW10 / CW11` | Reviewers need a stable index of implemented, fixture-only, blocked, and human-authorized release evidence. |

Prefer dispatching `S7W1`, `S7W2`, `S7W4`, and `S7W8` first because they improve upgrade safety, user explanation quality, diagnostics, and release review without touching host state or real execution.

## S7W1: Application Upgrade Impact Preview

### Mission

Create a Runtime-owned upgrade impact preview that compares the currently known application recipe evidence with a candidate recipe and explains what would change before any update, migration, install, or activation is attempted.

### Start from

- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/install_plan.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `internal/runtime/appidentity/compatibility_onboarding_checklist.go`
- `internal/runtime/appidentity/kde_journey_evidence.go`
- `internal/runtime/artifact/stage.go`

### Deliver

- A Go read model and CLI command such as `application-upgrade-impact-preview`.
- Impact sections for recipe metadata, package source, artifact digests, backend profile, Portal permissions, snapshot requirements, desktop activation, and diagnostics.
- A deterministic state per section: unchanged, changed, needs review, missing evidence, blocked, or unsupported.
- User-safe summaries that explain whether the upgrade is ready for review, blocked by trust evidence, or blocked by missing local fixtures.
- Tests proving candidate recipes can be older, same-version, newer, malformed, untrusted, digest-drifted, or missing required artifacts.
- Tests proving the preview never edits recipes, stages artifacts, downloads packages, changes settings, writes activation files, starts backends, or mutates host state.

### Keep disabled

- Recipe writes.
- Registry migration.
- Artifact staging.
- Network fetch.
- Package-manager calls.
- Settings persistence.
- Desktop activation writes.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/artifact ./cmd/xnix-runtime-go
ruby -Ilib test/test_recipe_registry.rb
ruby -Ilib test/test_recipe_trust_policy.rb
ruby -Ilib test/test_compatibility_install_plan.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S7W1 from docs/claude-code-seventh-wave-task-batch.md.

Target outcome:
- Add an application upgrade impact preview that compares current and candidate recipe evidence before any update, migration, install, or activation can happen.

Hard constraints:
- Do not edit recipes, migrate registries, stage artifacts, fetch packages, call package managers, persist settings, write desktop activation files, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, environment variables, raw backend commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add older-candidate, same-version, newer-candidate, malformed-candidate, untrusted-candidate, digest-drift, missing-artifact, and no-side-effect tests.
- Run the S7W1 verification commands and report exact commands run.
```

## S7W2: Runtime Policy Explanation Cards

### Mission

Create a shared Runtime-owned explanation model for KDE surfaces so blocked actions, review-only actions, not-yet-implemented actions, and safe next steps use the same wording and evidence ids across the launcher, task manager, Dolphin, tray, notifications, Compatibility Center, and settings.

### Start from

- `internal/runtime/appidentity/application_readiness.go`
- `internal/runtime/appidentity/kde_action_dependency_graph.go`
- `internal/runtime/appidentity/kde_journey_evidence.go`
- `internal/runtime/appidentity/compatibility_onboarding_checklist.go`
- `internal/runtime/appidentity/runtime_write_gate.go`
- `internal/runtime/appidentity/desktop_safety_policy.go`

### Deliver

- A Go read model and CLI command such as `runtime-policy-explanation-cards-preview`.
- Explanation cards for install, launch, execution, Portal permission, snapshot, diagnostics, repair, settings, desktop activation, backend readiness, and unsupported production routes.
- Stable card ids, severity, audience, related evidence ids, user-facing summary, technical summary, and next safe read-only check.
- Tests proving cards deduplicate repeated blocker reasons and avoid backend terminology in user-facing fields.
- Tests proving the preview never enables actions, persists cards, creates requests, grants permissions, starts backends, calls AI providers, or mutates host state.

### Keep disabled

- Action enablement.
- Card persistence.
- Request creation.
- Permission grants.
- Settings persistence.
- AI provider calls.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_runtime_write_gate.rb
ruby -Ilib test/test_kde_first_presence_smoke_script.rb
ruby scripts/kde_first_presence_smoke.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S7W2 from docs/claude-code-seventh-wave-task-batch.md.

Target outcome:
- Add Runtime policy explanation cards that give KDE surfaces consistent user-safe blocked, review-only, and not-yet-implemented explanations.

Hard constraints:
- Do not enable actions, persist cards, create request objects, grant permissions, persist settings, call AI providers, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, environment variables, raw backend commands, backend details, or user-facing Wine prefix terminology.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add install, launch, Portal, snapshot, diagnostics, repair, settings, unsupported-route, deduplication, no-backend-terms, and no-side-effect tests.
- Run the S7W2 verification commands and report exact commands run.
```

## S7W3: State-Root Quota and Retention Preview

### Mission

Create a dry-run quota and retention preview for per-application Runtime state roots so KDE can explain storage pressure, stale receipts, old snapshots, diagnostic history, and cleanup candidates without deleting anything.

### Start from

- `internal/runtime/appidentity/state_root.go`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/application_readiness.go`
- `internal/runtime/diagnostics/record.go`
- `internal/runtime/execution/ledger.go`
- `internal/runtime/portal/ledger.go`

### Deliver

- A Go read model and CLI command such as `state-root-quota-retention-preview`.
- Dry-run sections for snapshots, diagnostics, execution receipts, Portal receipts, artifact receipts, activation receipts, and unknown records.
- Relative evidence ids, estimated byte counts from fixture metadata, retention reason codes, blocked cleanup reasons, and user-safe recommendations.
- Tests proving unknown, malformed, over-quota, under-quota, missing-state-root, active-session-blocked, and retention-exempt cases.
- Tests proving the preview never deletes files, creates directories, truncates logs, rewrites receipts, exposes state-root paths, or mutates host state.

### Keep disabled

- File deletion.
- Directory creation.
- Log truncation.
- Receipt rewriting.
- Snapshot deletion.
- State-root path exposure.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/diagnostics ./internal/runtime/execution ./internal/runtime/portal ./cmd/xnix-runtime-go
ruby -Ilib test/test_application_state_root.rb
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S7W3 from docs/claude-code-seventh-wave-task-batch.md.

Target outcome:
- Add a state-root quota and retention preview that explains storage pressure and cleanup candidates without deleting or rewriting anything.

Hard constraints:
- Do not delete files, create directories, truncate logs, rewrite receipts, delete snapshots, expose state-root paths, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, environment variables, raw backend commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add missing-root, under-quota, over-quota, malformed-record, unknown-record, active-session-blocked, retention-exempt, and no-side-effect tests.
- Run the S7W3 verification commands and report exact commands run.
```

## S7W4: Crash and Hang Signal Summary Preview

### Mission

Create a privacy-safe diagnostic summary for repeated failed launches, crashes, hangs, timeouts, and repair regressions using fixture and receipt metadata only.

### Start from

- `internal/runtime/appidentity/diagnostic_history.go`
- `internal/runtime/appidentity/ai_diagnostics.go`
- `internal/runtime/appidentity/test_result.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/diagnostics/record.go`
- `scripts/implementation_evidence_report.rb`

### Deliver

- A Go read model and CLI command such as `crash-hang-signal-summary-preview`.
- Signal groups for crash, hang, timeout, missing dependency, permission denial, graphics issue, network issue, and regression after repair.
- Counts, latest known state, recurrence hints, privacy redaction summary, related diagnostic run ids, and next safe read-only checks.
- Tests proving the preview handles no history, malformed history, mixed app ids, duplicate signal ids, privacy-sensitive fixture fields, and blocked repair evidence.
- Tests proving the preview never reads private logs, reads file contents, calls AI providers, applies repairs, starts backends, or mutates host state.

### Keep disabled

- Private log reads.
- File-content reads.
- AI provider calls.
- Repair execution.
- Backend launch.
- Network calls.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/diagnostics ./cmd/xnix-runtime-go
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_compatibility_test_result.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S7W4 from docs/claude-code-seventh-wave-task-batch.md.

Target outcome:
- Add a crash and hang signal summary preview that uses fixture and receipt metadata to explain repeated failures without reading private logs or calling AI providers.

Hard constraints:
- Do not read private logs, read file contents, call AI providers, apply repairs, start backends, make network calls, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, environment variables, raw backend commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add no-history, malformed-history, mixed-app, duplicate-signal, privacy-redaction, blocked-repair, regression, and no-side-effect tests.
- Run the S7W4 verification commands and report exact commands run.
```

## S7W5: Compatibility Backend Fallback Preview

### Mission

Create a Runtime-owned fallback preview that explains how Automatic mode would choose between local compatibility and isolated compatibility when evidence is missing, blocked, or risky.

### Start from

- `internal/runtime/appidentity/backend_manager.go`
- `internal/runtime/appidentity/backend_capability_matrix.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/backend_selection_plan.go`
- `internal/runtime/appidentity/application_readiness.go`
- `internal/runtime/appidentity/desktop_safety_policy.go`

### Deliver

- A Go read model and CLI command such as `compatibility-backend-fallback-preview`.
- Fallback candidates for Automatic, Local compatibility, and Isolated compatibility with user-safe reason codes.
- Evidence joins for recipe hints, capability matrix, lifecycle status, Portal needs, snapshot needs, diagnostics history, and safety policy.
- Tests proving fallback stays blocked when both candidates lack evidence, when isolation is required, when local compatibility is allowed, and when policy forbids a profile.
- Tests proving the preview never persists selection, installs engines, starts backends, starts VMs, exposes backend names to KDE-facing fields, or mutates host state.

### Keep disabled

- Selection persistence.
- Engine installation.
- Backend launch.
- VM start.
- Raw backend detail exposure.
- Network calls.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/environment ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_backend_capability_matrix.rb
ruby -Ilib test/test_compatibility_backend_lifecycle.rb
ruby -Ilib test/test_compatibility_backend_selection_plan.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S7W5 from docs/claude-code-seventh-wave-task-batch.md.

Target outcome:
- Add a compatibility backend fallback preview for Automatic mode that explains safe choices between local and isolated compatibility without exposing implementation details.

Hard constraints:
- Do not persist backend selection, install engines, start backends, start VMs, make network calls, run Docker/QEMU, or mutate the host root.
- Do not expose raw backend names in KDE-facing fields, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, environment variables, raw backend commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add automatic-blocked, local-allowed, isolation-required, policy-forbidden, missing-lifecycle, diagnostics-risk, no-backend-terms, and no-side-effect tests.
- Run the S7W5 verification commands and report exact commands run.
```

## S7W6: KDE Search Visibility Plan

### Mission

Create a Runtime-owned visibility plan that explains how an application should appear in KDE search, KRunner, launcher categories, file associations, Compatibility Center search, and settings search.

### Start from

- `internal/runtime/appidentity/identity.go`
- `internal/runtime/appidentity/krunner.go`
- `internal/runtime/appidentity/file_association.go`
- `internal/runtime/appidentity/kde_shell_surface.go`
- `internal/runtime/appidentity/kde_center_page.go`
- `internal/runtime/appidentity/desktop_activation_manifest.go`

### Deliver

- A Go read model and CLI command such as `kde-search-visibility-plan-preview`.
- Visibility rows for launcher, KRunner, file association, Dolphin action, Compatibility Center, settings, and task manager.
- Stable searchable labels, safe synonyms, MIME and extension evidence, disabled action reasons, and activation receipt requirements.
- Tests proving duplicate labels, unsafe synonyms, missing activation receipt, unsupported MIME, hidden app, and malformed recipe cases.
- Tests proving the preview never writes desktop files, refreshes KDE caches, changes MIME defaults, indexes host files, starts backends, or mutates host state.

### Keep disabled

- Desktop file writes.
- KDE cache refresh.
- MIME default writes.
- Host file indexing.
- Search index persistence.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_krunner_model.rb
ruby -Ilib test/test_file_association_model.rb
ruby -Ilib test/test_kde_shell_integration_plan.rb
ruby scripts/kde_first_presence_smoke.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S7W6 from docs/claude-code-seventh-wave-task-batch.md.

Target outcome:
- Add a KDE search visibility plan preview that aligns launcher, KRunner, file associations, Compatibility Center, settings, and task manager visibility.

Hard constraints:
- Do not write desktop files, refresh KDE caches, change MIME defaults, index host files, persist search indexes, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, environment variables, raw backend commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add duplicate-label, unsafe-synonym, missing-activation-receipt, unsupported-MIME, hidden-app, malformed-recipe, no-host-indexing, and no-side-effect tests.
- Run the S7W6 verification commands and report exact commands run.
```

## S7W7: Permission Evidence Audit Preview

### Mission

Create an offline audit that compares settings, Portal receipts, permission review plans, execution preflight evidence, and KDE-facing resource bridge summaries without changing permissions.

### Start from

- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/portal_request_plan.go`
- `internal/runtime/appidentity/compatibility_permission_review_plan.go`
- `internal/runtime/appidentity/execution_preflight.go`
- `internal/runtime/appidentity/desktop_resource_bridge.go`
- `internal/runtime/portal/ledger.go`

### Deliver

- A Go read model and CLI command such as `permission-evidence-audit-preview`.
- Audit rows for documents, downloads, URIs, print, clipboard, screenshot, camera, remote desktop, and network policy.
- Evidence consistency states: consistent, setting-only, receipt-only, expired, denied, missing review, blocked by policy, or unsupported.
- User-safe remediation hints that remain review-only.
- Tests proving stale receipt, denied receipt, missing settings, missing review, policy block, malformed ledger, and unsupported permission cases.
- Tests proving the preview never calls real Portal transport, grants permissions, revokes permissions, writes receipts, changes settings, approves execution, or mutates host state.

### Keep disabled

- Real Portal calls.
- Permission grants.
- Permission revocation.
- Receipt writes.
- Settings persistence.
- Execution approval.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/portal ./cmd/xnix-runtime-go
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_compatibility_permission_review_plan.rb
ruby scripts/mainline_integration_review.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S7W7 from docs/claude-code-seventh-wave-task-batch.md.

Target outcome:
- Add a permission evidence audit preview that compares settings, Portal receipts, review plans, execution preflight, and resource bridge summaries without changing permissions.

Hard constraints:
- Do not call real Portal transport, grant permissions, revoke permissions, write receipts, persist settings, approve execution, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, environment variables, raw backend commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add consistent, setting-only, receipt-only, expired, denied, missing-review, policy-blocked, malformed-ledger, unsupported, and no-side-effect tests.
- Run the S7W7 verification commands and report exact commands run.
```

## S7W8: Release Evidence Index

### Mission

Create a stable offline index that maps every release-critical product claim to implemented evidence, fixture-only evidence, contract-only evidence, blocked evidence, or human-authorized smoke evidence.

### Start from

- `scripts/implementation_evidence_report.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/mainline_integration_review.rb`
- `scripts/kde_first_presence_smoke.rb`
- `docs/mainline-integration-checkpoint.md`
- `docs/windows-app-compatibility-implementation-brief.md`
- `docs/claude-code-dispatch-runbook.md`

### Deliver

- A Ruby report command such as `scripts/release_evidence_index.rb`.
- JSON and Markdown output with release claims, evidence source files, verification commands, current evidence level, unsafe gates, human authorization requirements, and next branch-sized follow-up.
- A default offline mode that does not run Docker, QEMU, network checks, package managers, backend launch, or host-mutating commands.
- Tests proving implemented, fixture-only, contract-only, blocked, skipped, malformed-report, protected-file, unclassified-file, and human-authorization-required cases.
- Documentation that explains how to use the index during Claude branch intake and release review.

### Keep disabled

- Docker.
- QEMU.
- Network checks.
- Package-manager calls.
- Backend launch.
- Host-root mutation.
- Automatic staging.
- Automatic release tagging.

### Required verification

```text
ruby -Ilib test/test_implementation_evidence_report.rb
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby -Ilib test/test_mainline_integration_review.rb
ruby -Ilib test/test_kde_first_presence_smoke_script.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/mainline_integration_review.rb --format json
ruby scripts/kde_first_presence_smoke.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S7W8 from docs/claude-code-seventh-wave-task-batch.md.

Target outcome:
- Add a release evidence index that maps release-critical product claims to implemented, fixture-only, contract-only, blocked, or human-authorized evidence.

Hard constraints:
- Do not run Docker, QEMU, network checks, package managers, backend launch, host-mutating commands, automatic staging, or automatic release tagging.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, environment variables, raw backend commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add implemented, fixture-only, contract-only, blocked, skipped, malformed-report, protected-file, unclassified-file, human-authorization-required, json-output, markdown-output, and no-side-effect tests.
- Run the S7W8 verification commands and report exact commands run.
```
