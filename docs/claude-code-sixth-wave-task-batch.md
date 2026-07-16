# Claude Code Sixth-Wave Task Batch

> Last updated: 2026-07-16 | Baseline under review: v0.2.293

This document contains a sixth batch of bounded Claude Code tasks for the KDE-first Windows application compatibility mainline.

Use this batch after the current mainline, second-wave, third-wave, fourth-wave, and fifth-wave work has been merged or deliberately skipped. The goal is to turn more contract-heavy Runtime domains into reviewable, offline, KDE-safe product evidence without enabling real execution, real Portal transport, Docker, QEMU, network fetch, privileged containers, host-root mutation, or compatibility backend launch.

Do not modify `docs/claude-code-implementation-packages.md` from any task in this batch. That file is Claude-owned.

## Global Rules for Every Sixth-Wave Task

Every Claude Code branch must:

- Pick exactly one task from `S6W1` through `S6W8`.
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
| 1 | `S6W1` Support case timeline preview | `CW7 / CW8 / CW10` | Supportability needs a user-safe timeline that joins diagnostic history, recommendations, action receipts, and KDE-visible status. |
| 2 | `S6W2` Desktop deactivation dry-run plan | `CW4 / CW6 / CW10` | Install and activation previews need the matching safe removal story before real desktop activation can advance. |
| 3 | `S6W3` Recipe conflict and pin audit | `CW2 / CW3` | Multiple recipe versions, backend constraints, and package-source pins need deterministic conflict evidence before update flows mature. |
| 4 | `S6W4` Portal permission renewal preview | `CW5 / CW8` | Users need review-only renewal, expiry, and revocation explanations before live Portal transport is introduced. |
| 5 | `S6W5` Snapshot restore candidate ranking | `CW6 / CW8` | Snapshot and rollback records need a ranked, KDE-safe restore chooser without executing restore. |
| 6 | `S6W6` Compatibility settings profile migration preview | `CW5 / CW8` | Settings profiles need migration evidence before persistence and profile import/export can become real. |
| 7 | `S6W7` Offline application fixture matrix | `CW10 / CW11` | Acceptance should test several representative app shapes without downloading artifacts or launching backends. |
| 8 | `S6W8` Merge readiness packet aggregator | `CW10 / CW11` | Reviewers need one offline packet that joins layout, drift, implementation evidence, KDE-first smoke, and dirty-lane classification. |

Prefer dispatching `S6W1`, `S6W2`, `S6W3`, and `S6W8` first because they improve supportability, lifecycle completeness, recipe quality, and merge safety without touching host state or real execution.

## S6W1: Support Case Timeline Preview

### Mission

Create a Runtime-owned support timeline that explains what happened for one application across diagnostic history, repair recommendations, action receipts, onboarding blockers, and KDE-visible readiness.

### Start from

- `internal/runtime/appidentity/diagnostic_history.go`
- `internal/runtime/appidentity/ai_diagnostics.go`
- `internal/runtime/appidentity/kde_journey_evidence.go`
- `internal/runtime/appidentity/compatibility_onboarding_checklist.go`
- `internal/runtime/appidentity/kde_action_dependency_graph.go`
- `scripts/implementation_evidence_report.rb`
- `scripts/mainline_integration_review.rb`

### Deliver

- A Go read model and CLI command such as `support-case-timeline-preview`.
- Timeline entries for diagnostic runs, blocked actions, repair recommendation categories, onboarding gaps, and KDE entry-point state changes.
- Stable event ids, event groups, severity levels, user-safe summaries, and next safe read-only checks.
- A `redaction_summary` that proves file contents, host paths, state-root paths, environment variables, usernames, tokens, credentials, and raw commands were omitted.
- Tests proving the preview can operate with no history, malformed history, mixed app ids, and blocked diagnostic evidence.
- Tests proving the timeline never creates a support ticket, exports a bundle, calls an AI provider, applies repair, starts a backend, or mutates host state.

### Keep disabled

- Ticket creation.
- Bundle export.
- AI provider calls.
- Auto-repair.
- Action execution.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/diagnostics ./cmd/xnix-runtime-go
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby scripts/mainline_integration_review.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S6W1 from docs/claude-code-sixth-wave-task-batch.md.

Target outcome:
- Add a support-case timeline preview that joins diagnostic history, repair recommendation categories, action blockers, onboarding gaps, and KDE journey evidence for one application.

Hard constraints:
- Do not create support tickets, export bundles, call AI providers, apply repairs, execute actions, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, environment variables, raw backend commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add no-history, malformed-history, mixed-app, blocked-evidence, redaction, and no-side-effect tests.
- Run the S6W1 verification commands and report exact commands run.
```

## S6W2: Desktop Deactivation Dry-Run Plan

### Mission

Create a dry-run plan that explains how a managed application would be removed from KDE entry points without deleting files or changing host associations.

### Start from

- `internal/runtime/appidentity/desktop_activation_transaction.go`
- `internal/runtime/appidentity/desktop_activation_manifest.go`
- `internal/runtime/appidentity/kde_journey_evidence.go`
- `internal/runtime/activation/stage.go`
- `lib/xnix/compatibility/desktop_activation_rollback.rb`
- `test/test_desktop_activation_rollback.rb`

### Deliver

- A Go read model and CLI command such as `desktop-deactivation-dry-run-preview`.
- A plan that lists launcher, desktop icon, MIME association, Dolphin service menu, KWin, tray, notification, Compatibility Center, and settings surfaces affected by deactivation.
- Receipt requirements, digest checks, rollback safety status, and relative evidence ids.
- User-safe blocked reasons for missing receipt, digest mismatch, unknown file owner, shared MIME association, and active session evidence.
- Tests proving the preview never deletes files, updates MIME defaults, refreshes KDE caches, starts processes, exposes target paths, or mutates host state.

### Keep disabled

- File deletion.
- MIME writes.
- KDE cache refresh.
- Receipt writes.
- Session termination.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/activation ./cmd/xnix-runtime-go
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S6W2 from docs/claude-code-sixth-wave-task-batch.md.

Target outcome:
- Add a desktop deactivation dry-run preview that explains how KDE-facing activation artifacts would be safely removed or left untouched.

Hard constraints:
- Do not delete files, update MIME defaults, refresh KDE caches, write receipts, terminate sessions, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose absolute target paths, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, credentials, raw commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add receipt-present, missing-receipt, digest-mismatch, shared-association, active-session-blocked, and no-side-effect tests.
- Run the S6W2 verification commands and report exact commands run.
```

## S6W3: Recipe Conflict and Pin Audit

### Mission

Create a deterministic audit that explains recipe version conflicts, package-source pins, backend capability constraints, and install-gate blockers before any update or install is attempted.

### Start from

- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `internal/runtime/appidentity/install_plan.go`
- `internal/runtime/appidentity/package_source.go`
- `internal/runtime/appidentity/backend_capability_matrix.go`
- `internal/runtime/artifact/stage.go`

### Deliver

- A Go read model and CLI command such as `recipe-conflict-audit-preview`.
- Conflict groups for duplicate app ids, stale recipe versions, unsupported capability claims, mismatched package-source pins, trust-policy blockers, and artifact digest drift.
- Deterministic resolution hints that remain review-only.
- Fixture registries covering clean, duplicate, stale, missing-pin, unsupported-capability, digest-drift, and untrusted cases.
- Tests proving the audit never edits recipes, stages artifacts, fetches packages, calls a package manager, starts backends, or mutates host state.

### Keep disabled

- Recipe writes.
- Registry migration.
- Artifact staging.
- Network fetch.
- Package-manager calls.
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
Implement S6W3 from docs/claude-code-sixth-wave-task-batch.md.

Target outcome:
- Add a recipe conflict and pin audit preview that classifies recipe registry conflicts before update, install, or backend work can begin.

Hard constraints:
- Do not edit recipes, migrate registries, stage artifacts, fetch network packages, call package managers, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, credentials, raw commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add clean, duplicate-id, stale-version, missing-pin, unsupported-capability, digest-drift, untrusted, and no-side-effect tests.
- Run the S6W3 verification commands and report exact commands run.
```

## S6W4: Portal Permission Renewal Preview

### Mission

Create a read-only preview that explains which Portal permissions are current, expiring, denied, revoked, or ready for user review.

### Start from

- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/portal_request_plan.go`
- `internal/runtime/portal/ledger.go`
- `internal/runtime/appidentity/compatibility_permission_review_plan.go`
- `internal/runtime/appidentity/kde_center_page.go`

### Deliver

- A Go read model and CLI command such as `portal-permission-renewal-preview`.
- Permission rows for files, URIs, print, clipboard, screen, camera, remote desktop, and network-mediated cases.
- Renewal states: current, needs review, expiring soon, denied, revoked, missing receipt, and blocked by policy.
- KDE-safe action labels for review-only flows.
- Tests proving the preview never calls real Portal transport, grants permissions, changes settings, writes receipts, approves execution, or mutates host state.

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
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S6W4 from docs/claude-code-sixth-wave-task-batch.md.

Target outcome:
- Add a Portal permission renewal preview that explains current, expiring, denied, revoked, missing, and policy-blocked permission evidence.

Hard constraints:
- Do not call real Portal transport, grant permissions, revoke permissions, write receipts, persist settings, approve execution, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, credentials, raw commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add current, expiring, denied, revoked, missing-receipt, policy-blocked, malformed-ledger, and no-side-effect tests.
- Run the S6W4 verification commands and report exact commands run.
```

## S6W5: Snapshot Restore Candidate Ranking

### Mission

Create a KDE-safe restore chooser that ranks snapshot candidates and explains why each candidate can or cannot be selected.

### Start from

- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/application_readiness.go`
- `internal/runtime/appidentity/execution_session_record_evidence.go`
- `internal/runtime/appidentity/kde_action_dependency_graph.go`
- `lib/xnix/compatibility/compatibility_snapshot_plan.rb`
- `test/test_compatibility_snapshot_plan.rb`

### Deliver

- A Go read model and CLI command such as `snapshot-restore-candidates-preview`.
- Ranked candidates with relative evidence ids, reason codes, compatibility risk labels, active-session blockers, missing-receipt blockers, and digest status.
- User-safe summaries for latest good, latest tested, last-known-running, blocked, and unknown candidates.
- Tests proving the preview never restores data, deletes snapshots, reads private file contents, terminates sessions, starts backends, or mutates host state.

### Keep disabled

- Restore execution.
- Snapshot deletion.
- File-content reads.
- Session termination.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S6W5 from docs/claude-code-sixth-wave-task-batch.md.

Target outcome:
- Add a snapshot restore candidate ranking preview that lets KDE show safe restore choices without executing restore.

Hard constraints:
- Do not restore data, delete snapshots, read private file contents, terminate sessions, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, credentials, raw commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add no-candidate, ranked-candidate, missing-receipt, digest-mismatch, active-session-blocked, malformed-record, and no-side-effect tests.
- Run the S6W5 verification commands and report exact commands run.
```

## S6W6: Compatibility Settings Profile Migration Preview

### Mission

Create a migration preview that explains how user-facing compatibility settings would move between schema versions without persisting changes.

### Start from

- `internal/runtime/appidentity/settings.go`
- `internal/runtime/appidentity/settings_change_plan.go`
- `internal/runtime/appidentity/compatibility_onboarding_checklist.go`
- `internal/runtime/appidentity/application_readiness.go`
- `lib/xnix/compatibility/settings_model.rb`
- `lib/xnix/compatibility/settings_change_plan.rb`

### Deliver

- A Go read model and CLI command such as `settings-profile-migration-preview`.
- Migration sections for run mode, priority, document access, downloads access, camera access, network access, snapshots, and diagnostics privacy.
- Schema version checks, defaulting rules, user-review-required flags, blocked setting ids, and rollback notes.
- Tests proving the preview never persists settings, grants resources, calls Portal transport, starts backends, exposes backend details, or mutates host state.

### Keep disabled

- Settings persistence.
- Resource grants.
- Real Portal calls.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_settings_model.rb
ruby -Ilib test/test_settings_change_plan.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S6W6 from docs/claude-code-sixth-wave-task-batch.md.

Target outcome:
- Add a compatibility settings profile migration preview that describes schema changes and user-review gates without persisting settings.

Hard constraints:
- Do not persist settings, grant resources, call real Portal transport, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, credentials, raw commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add current-schema, old-schema, future-schema, invalid-setting, blocked-setting, review-required, rollback-note, and no-side-effect tests.
- Run the S6W6 verification commands and report exact commands run.
```

## S6W7: Offline Application Fixture Matrix

### Mission

Create an offline matrix that verifies representative application shapes without downloading artifacts or launching compatibility backends.

### Start from

- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/application_readiness.go`
- `internal/runtime/appidentity/kde_journey_evidence.go`
- `internal/runtime/appidentity/compatibility_onboarding_checklist.go`
- `scripts/kde_first_presence_smoke.rb`
- `scripts/implementation_evidence_report.rb`

### Deliver

- A Ruby smoke/report script and optional Go read model such as `offline-application-fixture-matrix`.
- Fixture rows for document editor, game, installer, launcher, network-heavy app, tray-heavy app, and unsupported app shapes.
- Per-row evidence for recipe trust, artifact readiness, backend profile mapping, Portal needs, snapshot readiness, diagnostic readiness, KDE journey coverage, and blocked unsafe actions.
- JSON and Markdown output.
- Tests proving the matrix never downloads artifacts, calls package managers, launches backends, runs Docker/QEMU, or mutates host state.

### Keep disabled

- Network fetch.
- Package-manager calls.
- Artifact staging unless the task uses local immutable fixtures only.
- Backend launch.
- Docker.
- QEMU.
- Host-root mutation.

### Required verification

```text
ruby -Ilib test/test_kde_first_presence_smoke_script.rb
ruby scripts/kde_first_presence_smoke.rb --format json
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S6W7 from docs/claude-code-sixth-wave-task-batch.md.

Target outcome:
- Add an offline application fixture matrix that covers representative app shapes and reports KDE-safe readiness evidence in JSON and Markdown.

Hard constraints:
- Do not download artifacts, call package managers, launch backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, credentials, raw commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add matrix-shape, missing-fixture, unsupported-app, unsafe-action-disabled, json-output, markdown-output, and no-side-effect tests.
- Run the S6W7 verification commands and report exact commands run.
```

## S6W8: Merge Readiness Packet Aggregator

### Mission

Create a single offline readiness packet for reviewers that joins the layout verifier, implementation evidence report, Runtime contract drift report, KDE-first presence smoke, and mainline integration review.

### Start from

- `scripts/verify_layout.rb`
- `scripts/implementation_evidence_report.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/kde_first_presence_smoke.rb`
- `scripts/mainline_integration_review.rb`
- `test/test_implementation_evidence_report.rb`
- `test/test_runtime_contract_drift_report.rb`
- `test/test_mainline_integration_review.rb`

### Deliver

- A Ruby report command such as `scripts/merge_readiness_packet.rb`.
- JSON and Markdown output containing tool statuses, command strings, lane classification, protected-file status, unsafe-operation status, changed-file counts, required follow-up commands, and release-blocking reasons.
- A `--offline-only` default that does not run Docker, QEMU, network checks, package managers, or host-mutating commands.
- Tests proving the packet handles pass, fail, skipped, missing-command, malformed-json, protected-file, unclassified-file, and unsafe-operation cases.

### Keep disabled

- Staging.
- Committing.
- Tagging.
- Pushing.
- Docker.
- QEMU.
- Network fetch.
- Host-root mutation.

### Required verification

```text
ruby -Ilib test/test_implementation_evidence_report.rb
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby -Ilib test/test_mainline_integration_review.rb
ruby scripts/merge_readiness_packet.rb --format json
ruby scripts/merge_readiness_packet.rb --format markdown
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement S6W8 from docs/claude-code-sixth-wave-task-batch.md.

Target outcome:
- Add an offline merge readiness packet aggregator that joins layout, implementation evidence, Runtime drift, KDE-first smoke, and mainline integration review into one reviewer packet.

Hard constraints:
- Do not stage, commit, tag, push, run Docker/QEMU, fetch network resources, call package managers, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, credentials, raw commands, or backend details.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add pass, fail, skipped, missing-command, malformed-json, protected-file, unclassified-file, unsafe-operation, JSON-output, and Markdown-output tests.
- Run the S6W8 verification commands and report exact commands run.
```
