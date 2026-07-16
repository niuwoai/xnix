# Claude Code Fifth-Wave Task Batch

> Last updated: 2026-07-16 | Baseline under review: v0.2.293

This document contains a fifth batch of bounded Claude Code tasks for the KDE-first Windows application compatibility mainline.

Use this batch after the current mainline, second-wave, third-wave, and fourth-wave work has been merged or deliberately skipped. The goal is to connect Runtime evidence to end-user journeys, offline support bundles, fixture-based acceptance, and merge-readiness gates without enabling real execution, real Portal transport, Docker, QEMU, network fetch, privileged containers, or host-root mutation.

Do not modify `docs/claude-code-implementation-packages.md` from any task in this batch. That file is Claude-owned.

## Global Rules for Every Fifth-Wave Task

Every Claude Code branch must:

- Pick exactly one task from `F5W1` through `F5W8`.
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
- Raw executable paths, Runtime state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials in user-facing output.

## Dispatch Order

| Order | Task | Primary lane | Why now |
| --- | --- | --- | --- |
| 1 | `F5W1` KDE journey evidence stitcher | `CW4 / CW8` | The seven KDE entry points should share one user-safe explanation of why an app is ready, blocked, or needs review. |
| 2 | `F5W2` Compatibility onboarding checklist | `CW2 / CW3 / CW5` | First-run setup should teach the user what is required without creating requests, installing backends, or mutating host state. |
| 3 | `F5W3` Offline support bundle manifest | `CW7 / CW10` | Users and reviewers need a redacted support bundle manifest before any real export tool exists. |
| 4 | `F5W4` Multi-application install queue preview | `CW2 / CW8` | Batch installation planning should stay deterministic and review-only before activation or execution is enabled. |
| 5 | `F5W5` Runtime state maintenance planner | `CW3 / CW6` | State-root cleanup, stale receipt handling, and retention policies need dry-run planning before deletion is possible. |
| 6 | `F5W6` Recipe update migration preview | `CW2 / CW10` | Recipe evolution needs a local fixture migration plan before network registries or production trust are introduced. |
| 7 | `F5W7` KDE notification digest model | `CW4 / CW7 / CW8` | Notification center and tray surfaces need a deduplicated, non-alarming summary of blocked actions and review items. |
| 8 | `F5W8` Restricted acceptance fixture suite | `CW10 / CW11` | The project needs an offline fixture acceptance suite that names blocked Docker/QEMU checks without running them. |

Prefer dispatching `F5W1`, `F5W2`, `F5W3`, and `F5W8` first because they improve user-journey clarity, onboarding, supportability, and merge safety without touching launch, host state, or heavy virtualization.

## F5W1: KDE Journey Evidence Stitcher

### Mission

Create a Runtime-owned journey summary that explains the same application's state across the seven KDE entry points: launcher, task manager, Dolphin, tray, notifications, AI Compatibility Center, and unified settings.

### Start from

- `internal/runtime/appidentity/kde_center_page.go`
- `internal/runtime/appidentity/kde_action_dependency_graph.go`
- `internal/runtime/appidentity/application_readiness.go`
- `internal/runtime/appidentity/window_identity_routes.go`
- `internal/runtime/appidentity/kde_shell_surface.go`
- `scripts/kde_first_presence_smoke.rb`
- `test/test_kde_first_presence_smoke_script.rb`

### Deliver

- A Go read model and CLI command such as `kde-journey-evidence-preview`.
- One entry per KDE entry point with shared readiness status, blocked reasons, next safe read-only checks, and user-facing copy.
- Cross-links to application readiness, action dependency graph, task-manager identity, KWin rule, tray status, notifications, and settings dependency evidence.
- Tests proving all seven entry points agree on app id, display name, readiness state, and disabled unsafe actions.
- Tests proving no backend terms, raw executable paths, state-root paths, host paths, file contents, or secrets appear in KDE-facing output.

### Keep disabled

- Runtime write methods.
- Request creation.
- Permission grants.
- Launch approval.
- Backend launch.
- KWin rule application.
- Tray live bridge activation.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_kde_first_presence_smoke_script.rb
ruby scripts/kde_first_presence_smoke.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement F5W1 from docs/claude-code-fifth-wave-task-batch.md.

Target outcome:
- Add a KDE journey evidence preview that gives the launcher, task manager, Dolphin, tray, notifications, AI Compatibility Center, and unified settings the same Runtime-owned explanation of app readiness and blocked actions.

Hard constraints:
- Do not enable Runtime writes, create requests, grant permissions, approve launch, start backends, apply KWin rules, activate live tray bridging, run Docker/QEMU, or mutate the host root.
- Do not expose backend terms, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials in KDE-facing output.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add seven-entrypoint, mismatched-app, blocked-action, missing-evidence, and no-sensitive-output tests.
- Run the F5W1 verification commands and report exact commands run.
```

## F5W2: Compatibility Onboarding Checklist

### Mission

Create a first-run onboarding checklist that explains what the Runtime needs before Windows applications can be managed safely.

### Start from

- `internal/runtime/appidentity/desktop_safety_policy.go`
- `internal/runtime/appidentity/application_readiness.go`
- `internal/runtime/appidentity/runtime_owner_readiness.go`
- `internal/runtime/appidentity/runtime_service_binding.go`
- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/snapshot_plan.go`

### Deliver

- A Go read model and CLI command such as `compatibility-onboarding-checklist-preview`.
- Checklist sections for Runtime owner readiness, recipe trust, artifact staging, backend lifecycle, Portal review, snapshot baseline, diagnostics privacy, and KDE entry points.
- User-safe states: ready, needs review, missing evidence, blocked, and not yet implemented.
- Tests proving onboarding never creates Portal requests, stages artifacts, changes settings, starts backends, or mutates host state.

### Keep disabled

- Request creation.
- Artifact staging.
- Settings persistence.
- Backend launch.
- Network fetch.
- Host package-manager calls.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/owner ./cmd/xnix-runtime-go
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement F5W2 from docs/claude-code-fifth-wave-task-batch.md.

Target outcome:
- Add a Compatibility onboarding checklist preview that explains first-run Runtime readiness without taking action.

Hard constraints:
- Do not create Portal requests, stage artifacts, persist settings, start backends, fetch network artifacts, invoke package managers, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add ready, needs-review, missing-evidence, blocked, unsupported, and no-side-effect tests.
- Run the F5W2 verification commands and report exact commands run.
```

## F5W3: Offline Support Bundle Manifest

### Mission

Create a redacted support bundle manifest that tells a user what diagnostic evidence could be shared without exporting files or reading private content.

### Start from

- `internal/runtime/appidentity/diagnostic_history.go`
- `internal/runtime/diagnostics/history.go`
- `internal/runtime/diagnostics/record.go`
- `internal/runtime/appidentity/ai_diagnostics.go`
- `scripts/implementation_evidence_report.rb`
- `scripts/mainline_integration_review.rb`

### Deliver

- A Go read model and CLI command such as `support-bundle-manifest-preview`.
- Manifest sections for Runtime version, app id, diagnostic run summaries, failing signal ids, repair recommendation categories, privacy redaction status, and omitted evidence categories.
- Explicit counts for omitted file contents, host paths, environment variables, token-shaped values, command-shaped values, and usernames.
- Tests proving the preview never reads file contents, writes an archive, calls AI providers, exposes state-root paths, or includes secrets.

### Keep disabled

- Archive creation.
- File-content reads.
- AI provider calls.
- Auto-repair.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/diagnostics ./cmd/xnix-runtime-go
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement F5W3 from docs/claude-code-fifth-wave-task-batch.md.

Target outcome:
- Add an offline support bundle manifest preview that summarizes redacted diagnostic evidence without exporting any bundle.

Hard constraints:
- Do not create archives, read file contents, call AI providers, apply repairs, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add manifest, redaction-count, malformed-record, provider-disabled, no-archive, and no-sensitive-output tests.
- Run the F5W3 verification commands and report exact commands run.
```

## F5W4: Multi-Application Install Queue Preview

### Mission

Create a review-only queue that plans multiple application installs while preserving per-application evidence and blocked reasons.

### Start from

- `internal/runtime/appidentity/install_plan.go`
- `internal/runtime/appidentity/application_readiness.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `internal/runtime/artifact/stage.go`
- `internal/runtime/appidentity/kde_action_queue.go`
- `test/test_compatibility_install_plan.rb`
- `test/test_compatibility_action_queue.rb`

### Deliver

- A Go read model and CLI command such as `multi-app-install-queue-preview`.
- Queue items with app id, display name, recipe trust state, artifact verification state, backend lifecycle state, Portal readiness, snapshot readiness, and blocked reasons.
- Deterministic ordering, duplicate app handling, malformed app id handling, and per-item safe summaries.
- Tests proving no artifact staging, package installation, backend launch, desktop activation, or host mutation occurs.

### Keep disabled

- Artifact staging.
- Package-manager calls.
- Desktop activation.
- Backend launch.
- Network fetch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/artifact ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_install_plan.rb
ruby -Ilib test/test_compatibility_action_queue.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement F5W4 from docs/claude-code-fifth-wave-task-batch.md.

Target outcome:
- Add a multi-application install queue preview that aggregates install readiness for several apps without installing anything.

Hard constraints:
- Do not stage artifacts, invoke package managers, activate desktop files, start backends, fetch network artifacts, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add multi-app, duplicate-app, malformed-app, blocked-item, ordering, and no-side-effect tests.
- Run the F5W4 verification commands and report exact commands run.
```

## F5W5: Runtime State Maintenance Planner

### Mission

Create a dry-run maintenance planner for Runtime state roots that identifies stale receipts, orphaned records, expired fake-mode Portal requests, old diagnostics, and snapshot retention pressure without deleting anything.

### Start from

- `internal/runtime/execution/ledger.go`
- `internal/runtime/execution/session.go`
- `internal/runtime/portal/ledger.go`
- `internal/runtime/diagnostics/history.go`
- `internal/runtime/snapshot/store.go`
- `internal/runtime/appidentity/backend_lifecycle.go`

### Deliver

- A Go read model and CLI command such as `runtime-state-maintenance-preview`.
- Dry-run maintenance categories for stale execution records, orphaned sessions, expired Portal requests, old diagnostics, unused artifact receipts, stale backend lifecycle records, and snapshot retention candidates.
- Relative receipt ids and redacted summaries only.
- Tests proving no files are deleted, rewritten, compacted, restored, or moved.

### Keep disabled

- Deletion.
- Compaction.
- Snapshot restore.
- Receipt rewrite.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/execution ./internal/runtime/portal ./internal/runtime/diagnostics ./internal/runtime/snapshot ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement F5W5 from docs/claude-code-fifth-wave-task-batch.md.

Target outcome:
- Add a Runtime state maintenance preview that dry-runs cleanup candidates without deleting or rewriting anything.

Hard constraints:
- Do not delete, compact, restore, rewrite, move files, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add stale-record, orphan-session, expired-portal, old-diagnostic, snapshot-retention, malformed-record, and no-delete tests.
- Run the F5W5 verification commands and report exact commands run.
```

## F5W6: Recipe Update Migration Preview

### Mission

Create a local-fixture migration preview that compares current and proposed recipe registries without fetching network data or trusting production signatures.

### Start from

- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `internal/runtime/appidentity/install_plan.go`
- `runtime/recipes/`
- `test/test_recipe_registry.rb`
- `test/test_recipe_trust_policy.rb`

### Deliver

- A Go read model and CLI command such as `recipe-update-migration-preview`.
- Comparison rows for added, removed, unchanged, digest-changed, signature-state-changed, incompatible-schema, and blocked recipes.
- Safe upgrade impact summaries for desktop activation, backend lifecycle, permissions, snapshots, and diagnostics.
- Tests proving migrations are preview-only and never update registry files.

### Keep disabled

- Registry writes.
- Network registry fetch.
- Production trust claims.
- Artifact downloads.
- Install activation.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_recipe_registry.rb
ruby -Ilib test/test_recipe_trust_policy.rb
ruby -Ilib test/test_compatibility_install_plan.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement F5W6 from docs/claude-code-fifth-wave-task-batch.md.

Target outcome:
- Add a recipe update migration preview that compares local fixture registries and explains compatibility impact without updating anything.

Hard constraints:
- Do not write registries, fetch network registries, claim production trust, download artifacts, activate installs, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add added, removed, unchanged, digest-changed, signature-changed, incompatible-schema, blocked, and no-write tests.
- Run the F5W6 verification commands and report exact commands run.
```

## F5W7: KDE Notification Digest Model

### Mission

Create a KDE-safe notification digest that deduplicates blocked action, diagnostics, permission, snapshot, and readiness events without sending real desktop notifications.

### Start from

- `internal/runtime/appidentity/identity.go`
- `internal/runtime/appidentity/kde_center_page.go`
- `internal/runtime/appidentity/kde_action_dependency_graph.go`
- `internal/runtime/appidentity/diagnostic_history.go`
- `internal/runtime/appidentity/portal_access_policy.go`
- `test/test_notification_request.rb`
- `test/test_tray_status_model.rb`

### Deliver

- A Go read model and CLI command such as `kde-notification-digest-preview`.
- Digest groups for needs review, blocked action, permission attention, diagnostic issue, snapshot warning, and readiness change.
- Deduplication keys, severity, user-safe labels, and next safe read-only route.
- Tests proving no desktop notification is sent, no tray bridge is activated, and no backend detail leaks.

### Keep disabled

- Real notification delivery.
- Tray live bridge activation.
- Request creation.
- Permission grants.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_notification_request.rb
ruby -Ilib test/test_tray_status_model.rb
ruby scripts/kde_first_presence_smoke.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement F5W7 from docs/claude-code-fifth-wave-task-batch.md.

Target outcome:
- Add a KDE notification digest preview that deduplicates Runtime review and blocked-action events without sending notifications.

Hard constraints:
- Do not send desktop notifications, activate live tray bridging, create requests, grant permissions, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose backend terms, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials in KDE-facing output.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add digest-group, deduplication, severity, blocked-route, notification-disabled, and no-sensitive-output tests.
- Run the F5W7 verification commands and report exact commands run.
```

## F5W8: Restricted Acceptance Fixture Suite

### Mission

Create an offline acceptance fixture suite that proves the main KDE-first compatibility journey can be reviewed locally while explicitly naming Docker, QEMU, Colima, and network checks as blocked unless authorized.

### Start from

- `scripts/kde_first_presence_smoke.rb`
- `scripts/mainline_integration_review.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/implementation_evidence_report.rb`
- `docs/kde-first-compatibility-acceptance.md`
- `docs/kde-first-presence-smoke-spec.md`
- `test/test_kde_first_presence_smoke_script.rb`
- `test/test_mainline_integration_review.rb`

### Deliver

- A Ruby report script such as `scripts/restricted_acceptance_fixture_suite.rb` with JSON and Markdown output.
- Offline fixture assertions for KDE seven-entrypoint presence, Runtime read-only route coverage, action dependency graph presence, write-gate closure, protected-file status, and unclassified-file status.
- A `blocked_heavy_checks` section that names Docker, QEMU, Colima, network fetch, package-manager calls, and host-root mutation as not run.
- Tests proving the suite does not execute Docker, QEMU, Colima, network fetch, package managers, git staging, git commits, git tags, git pushes, or host mutation commands.

### Keep disabled

- Docker execution.
- QEMU execution.
- Colima start or stop.
- Network fetch.
- Package-manager calls.
- Git staging, committing, tagging, or pushing.
- Host-root mutation.

### Required verification

```text
ruby -Ilib test/test_kde_first_presence_smoke_script.rb
ruby -Ilib test/test_mainline_integration_review.rb
ruby scripts/restricted_acceptance_fixture_suite.rb --format json
ruby scripts/restricted_acceptance_fixture_suite.rb --format markdown
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement F5W8 from docs/claude-code-fifth-wave-task-batch.md.

Target outcome:
- Add an offline restricted acceptance fixture suite that summarizes KDE-first compatibility evidence and explicitly lists heavy checks as blocked unless human-authorized.

Hard constraints:
- Do not run Docker, QEMU, Colima, network fetches, package managers, git staging, commits, tags, pushes, or host-root mutation.
- Do not expose host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add JSON, Markdown, fixture-input, blocked-heavy-check, protected-file, unclassified-file, and no-command-execution tests.
- Run the F5W8 verification commands and report exact commands run.
```
