# Claude Code Second-Wave Task Batch

> Last updated: 2026-07-16 | Baseline under review: v0.2.293

This document contains a second batch of bounded Claude Code tasks for Xnix.

Use it after the current mainline integration pass starts merging the first `CB` work. The goal of this batch is to move more KDE-first Windows compatibility behavior from static contracts into Runtime-owned, testable implementation evidence without widening the unsafe execution surface.

Do not modify `docs/claude-code-implementation-packages.md` from any task in this batch. That file is Claude-owned.

## Global Rules for Every Second-Wave Task

Every Claude Code branch must:

- Pick exactly one task from `SW1` through `SW8`.
- Keep source, comments, fixtures, tests, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior.
- Use C only for low-level D-Bus transport, ABI-shaped records, or already-owned Runtime policy surfaces.
- Use Ruby only for tests, smoke scripts, reports, and lightweight developer tooling.
- Keep KDE as shell and presentation only.
- Keep the Xnix Compatibility Runtime as the policy and product-logic owner.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for code changes.
- Add success, failure, and blocked-state tests.
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
- Host-root mutation.
- Raw backend commands in KDE-facing output.
- Raw executable paths, Runtime state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials in user-facing output.

## Dispatch Order

Current local status:

- `SW1` has a local seed implementation in this working tree. Assign it to Claude only for review, hardening, or missing-test follow-up.
- `SW2` has a local seed implementation in this working tree. Assign it to Claude only for review, hardening, or missing-test follow-up.
- `SW3` through `SW8` remain the preferred next implementation dispatch set.

| Order | Task | Primary lane | Why now |
| --- | --- | --- | --- |
| 1 | `SW1` Runtime desktop safety policy preview | `CW4 / CW10` | KDE-safe terminology and settings policy should be Go Runtime-owned, not Ruby-smoke-owned. |
| 2 | `SW2` Application readiness evidence graph | `CW2 / CW3 / CW5 / CW6 / CW8` | KDE and execution views need one safe read model that joins durable receipts. |
| 3 | `SW3` Compatibility Center evidence sections | `CW4` | Center sections should consume Runtime evidence instead of repeating static previews. |
| 4 | `SW4` Desktop activation transaction receipts | `CW9` | Activation staging needs stronger receipt and rollback evidence before any real target-root writes. |
| 5 | `SW5` Session evidence fan-out hardening | `CW8 / CW4` | Task manager, tray, KWin, and Center views should share one execution-session evidence source. |
| 6 | `SW6` Restricted local smoke aggregator | `CW10` | Developers need one safe local smoke summary that avoids Docker, QEMU, and host mutation by default. |
| 7 | `SW7` Image acceptance preflight manifest | `CW11` | QEMU acceptance should have a dry-run readiness manifest before running heavy image tests. |
| 8 | `SW8` Claude intake guardrails | `CW10` | Merging many Claude branches needs stronger protection against cache files, protected docs, and version drift. |

Do not dispatch `SW7` as a real image or QEMU run. `SW7` is only a preflight manifest task unless a human explicitly authorizes restricted Docker or QEMU execution.

## SW1: Runtime Desktop Safety Policy Preview

### Mission

Move KDE-facing terminology and settings safety policy into a Go Runtime-owned read model and make the KDE-first presence smoke consume it.

### Start from

- `internal/runtime/appidentity/`
- `cmd/xnix-runtime-go/`
- `scripts/kde_first_presence_smoke.rb`
- `test/test_kde_first_presence_smoke_script.rb`
- `scripts/implementation_evidence_report.rb`
- `test/test_implementation_evidence_report.rb`
- `scripts/verify_layout.rb`

### Deliver

- A Go model and CLI command named `desktop-safety-policy-preview`.
- A schema such as `xnix.runtime.desktop_safety_policy.v1`.
- Runtime-owned lists for the seven KDE entry points, user-facing settings field ids, forbidden backend/user-facing terms, and disabled safety keys.
- Tests proving Wine, Proton, prefix, bottle, raw executable, host path, state-root path, and Docker socket terminology are not exposed in KDE-facing payloads.
- `scripts/kde_first_presence_smoke.rb` should consume this Go Runtime policy instead of owning the policy only in Ruby constants.
- Implementation evidence and layout checks should know this policy is product evidence, not a loose smoke helper.

### Keep disabled

- Settings persistence.
- Backend launch.
- Runtime write methods.
- Real Portal calls.
- Network access.
- Host-root mutation.

### Required verification

```text
go test ./cmd/xnix-runtime-go ./internal/runtime/appidentity
ruby -Ilib test/test_kde_first_presence_smoke_script.rb
ruby -Ilib test/test_implementation_evidence_report.rb
ruby scripts/kde_first_presence_smoke.rb --format json
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement SW1 from docs/claude-code-second-wave-task-batch.md.

Target outcome:
- Add a Go Runtime-owned desktop safety policy preview and make the KDE-first presence smoke consume it as product evidence.

Hard constraints:
- Do not enable settings writes, Runtime write methods, real Portal calls, backend launch, network fetch, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Do not expose Wine, Proton, prefix, bottle, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials in KDE-facing output.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add success, forbidden-term, disabled-safety-key, and smoke-consumer tests.
- Run the SW1 verification commands and report exact commands run.
```

## SW2: Application Readiness Evidence Graph

### Mission

Create a Runtime-owned application readiness read model that joins existing receipt and state evidence into one KDE-safe status without enabling execution.

### Start from

- `internal/runtime/appidentity/install_plan.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/portal*`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `internal/runtime/artifact/`
- `internal/runtime/portal/`
- `internal/runtime/snapshot/`
- `cmd/xnix-runtime-go/`

### Deliver

- A Go read model and CLI command such as `application-readiness-preview`.
- Inputs for local fixture receipts or state-root records where available.
- A joined readiness graph with recipe trust, artifact staging, backend lifecycle, Portal review, snapshot baseline, owner readiness, and write-gate status.
- Stable blocked reasons for missing or malformed evidence.
- KDE-safe summaries that hide backend details, host paths, state-root paths, raw commands, and raw executables.

### Keep disabled

- Launch approval.
- Backend process start.
- Real Portal calls.
- Snapshot creation or rollback.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/artifact ./internal/runtime/portal ./internal/runtime/snapshot ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_install_plan.rb
ruby -Ilib test/test_compatibility_backend_lifecycle.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement SW2 from docs/claude-code-second-wave-task-batch.md.

Target outcome:
- Add a Runtime-owned application readiness preview that joins recipe/artifact, backend lifecycle, Portal, snapshot, owner, and write-gate evidence into one KDE-safe read model.

Hard constraints:
- Do not approve launch, start Wine/Proton/VM, create real Portal requests, create snapshots, restore snapshots, fetch network artifacts, call host package managers, or mutate the host root.
- Do not expose raw backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add success, missing-evidence, malformed-evidence, blocked-readiness, and hidden-path tests.
- Run the SW2 verification commands and report exact commands run.
```

## SW3: Compatibility Center Evidence Sections

### Mission

Make the KDE Compatibility Center page sections consume durable Runtime evidence instead of repeating static preview-only state.

### Start from

- `internal/runtime/appidentity/kde_center_page.go`
- `internal/runtime/appidentity/kde_action_*.go`
- `internal/runtime/appidentity/diagnostic_history.go`
- `internal/runtime/appidentity/execution_session*.go`
- `cmd/xnix-runtime-go/kde_center*`
- `scripts/kde_first_presence_smoke.rb`

### Deliver

- Center page sections for readiness, permissions, snapshots, diagnostics, recent activity, and blocked actions that can consume available Runtime evidence.
- User-safe section summaries and action cards.
- Tests proving the page hides backend details, raw commands, state-root paths, host paths, and file contents.
- Smoke evidence that the Compatibility Center remains one of exactly seven KDE first-release entry points.

### Keep disabled

- Action execution.
- Settings persistence.
- Backend launch.
- Real Portal transport.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_kde_center_model.rb
ruby -Ilib test/test_kde_first_presence_smoke_script.rb
ruby scripts/kde_first_presence_smoke.rb --format json
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement SW3 from docs/claude-code-second-wave-task-batch.md.

Target outcome:
- Make KDE Compatibility Center sections consume Runtime evidence for readiness, permissions, snapshots, diagnostics, recent activity, and blocked actions.

Hard constraints:
- Do not execute actions, persist settings, start backends, call real Portal transport, mutate host permissions, or mutate the host root.
- Do not expose raw backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add evidence-consumer, missing-evidence, blocked-action, and hidden-detail tests.
- Run the SW3 verification commands and report exact commands run.
```

## SW4: Desktop Activation Transaction Receipts

### Mission

Strengthen desktop activation staging so planned KDE desktop files, MIME associations, service menus, icons, status records, and rollback receipts are auditable under explicit target roots only.

### Start from

- `internal/runtime/activation/`
- `internal/runtime/appidentity/desktop_activation_*.go`
- `cmd/xnix-runtime-go/desktop_activation_*`
- `test/test_desktop_activation_installer.rb`
- `test/test_desktop_activation_rollback.rb`
- `kde/`

### Deliver

- Activation transaction receipts with schema version, operation id, app id, relative paths, digests, rollback ids, blocked reasons, and safe summaries.
- Path-escape, missing-file, digest-mismatch, and rollback-blocked tests.
- A preview/status flow that can read receipts without exposing target-root paths.
- Evidence report coverage for staged-root-only activation materialization.

### Keep disabled

- Production desktop writes.
- Host service cache refresh.
- Launch.
- Backend process start.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/activation ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement SW4 from docs/claude-code-second-wave-task-batch.md.

Target outcome:
- Strengthen desktop activation transaction receipts so staged KDE materialization is auditable, rollback-aware, and target-root confined.

Hard constraints:
- Do not write to production desktop locations, refresh host service caches, launch applications, start backends, require privileged containers, or mutate the host root.
- Do not expose target-root paths, raw backend commands, raw executable paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add receipt success, path-escape, digest-mismatch, missing-file, rollback-blocked, and hidden-path tests.
- Run the SW4 verification commands and report exact commands run.
```

## SW5: Session Evidence Fan-Out Hardening

### Mission

Make execution session evidence the single Runtime-owned source for task-manager identity, tray status, KWin hints, notifications, and Compatibility Center recent activity while real execution remains disabled.

### Start from

- `internal/runtime/execution/`
- `internal/runtime/appidentity/execution_session*.go`
- `internal/runtime/appidentity/window_identity_routes.go`
- `internal/runtime/appidentity/kde_center_page.go`
- `internal/runtime/appidentity/kde_shell_surface.go`
- `cmd/xnix-runtime-go/`

### Deliver

- A session evidence record that fan-outs to task-manager, tray, KWin, notification, and Center models.
- Tests for completed, blocked, denied-permission, missing-ledger, and malformed-session states.
- Hidden-detail tests for backend commands, state-root paths, host paths, raw executables, and file contents.

### Keep disabled

- Real execution.
- Window observation.
- KWin rule application.
- Tray bridge activation.
- Notification sending.
- Backend process start.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/execution ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_task_manager_identity.rb
ruby -Ilib test/test_tray_status_model.rb
ruby -Ilib test/test_kwin_window_rule.rb
ruby -Ilib test/test_kde_center_model.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement SW5 from docs/claude-code-second-wave-task-batch.md.

Target outcome:
- Use Runtime execution-session evidence as the single safe source for task-manager, tray, KWin, notification, and Compatibility Center recent-activity views.

Hard constraints:
- Do not start execution, observe live windows, apply KWin rules, activate tray bridges, send notifications, start backends, or mutate the host root.
- Do not expose raw backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add session fan-out, blocked-session, malformed-session, missing-ledger, and hidden-detail tests.
- Run the SW5 verification commands and report exact commands run.
```

## SW6: Restricted Local Smoke Aggregator

### Mission

Create a safe local smoke aggregator that runs only non-Docker, non-QEMU, non-host-mutating checks by default and emits JSON plus Markdown summaries.

### Start from

- `scripts/full_smoke.rb`
- `scripts/kde_first_presence_smoke.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/implementation_evidence_report.rb`
- `scripts/mainline_integration_review.rb`
- `test/test_full_smoke_report.rb`
- `test/test_full_smoke_script.rb`

### Deliver

- A new restricted local smoke mode or script, such as `scripts/restricted_local_smoke.rb`.
- JSON and Markdown output.
- A default command set that avoids Docker, QEMU, network, privileged containers, host networking, Docker socket mounts, broad host mounts, and host-root mutation.
- Exit-code behavior that fails on contract drift, missing implementation evidence, protected Claude document changes, unclassified changed files, layout failures, or smoke failures.
- Tests for pass, fail, skipped-unsafe-command, and report-shape behavior.

### Keep disabled

- Docker or Colima execution by default.
- QEMU execution by default.
- Network fetch.
- Host package-manager calls.
- Host-root mutation.

### Required verification

```text
ruby -Ilib test/test_full_smoke_report.rb
ruby -Ilib test/test_full_smoke_script.rb
ruby -Ilib test/test_mainline_integration_review.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/mainline_integration_review.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement SW6 from docs/claude-code-second-wave-task-batch.md.

Target outcome:
- Add a restricted local smoke aggregator that emits JSON and Markdown while avoiding Docker, QEMU, network, privileged containers, host networking, Docker socket mounts, broad host mounts, and host-root mutation by default.

Hard constraints:
- Do not run Docker, Colima, QEMU, network fetches, host package-manager commands, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation from the default path.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add pass, fail, skipped-unsafe-command, JSON report, Markdown report, and exit-code tests.
- Run the SW6 verification commands and report exact commands run.
```

## SW7: Image Acceptance Preflight Manifest

### Mission

Add a dry-run image acceptance manifest that explains whether Xnix is ready for restricted Docker or QEMU acceptance without running Docker or QEMU.

### Start from

- `scripts/build_kde_image.rb`
- `scripts/build_kde_disk.rb`
- `scripts/boot_kde_image.rb`
- `scripts/full_smoke.rb`
- `buildroot/`
- `docs/kde-image-pipeline.md`
- `docs/kde-first-compatibility-acceptance.md`
- `test/test_kde_image.rb`
- `test/test_qemu.rb`
- `test/test_sshd.rb`

### Deliver

- A dry-run preflight command or script such as `scripts/image_acceptance_preflight.rb`.
- JSON and Markdown output describing required kernel, initramfs, rootfs, SSH, serial-log, KDE, Runtime, and loopback-only port-forward evidence.
- Clear blocked reasons when artifacts, config, or safety settings are missing.
- Tests that prove the preflight does not run Docker, QEMU, network fetch, privileged containers, host networking, or host-root mutation.

### Keep disabled

- Docker execution.
- QEMU execution.
- Network fetch.
- Privileged containers.
- Host networking.
- Host-root mutation.

### Required verification

```text
ruby -Ilib test/test_kde_image.rb
ruby -Ilib test/test_qemu.rb
ruby -Ilib test/test_sshd.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement SW7 from docs/claude-code-second-wave-task-batch.md.

Target outcome:
- Add a dry-run image acceptance preflight manifest that reports whether restricted Docker or QEMU acceptance is ready without running Docker or QEMU.

Hard constraints:
- Do not run Docker, Colima, QEMU, network fetches, privileged containers, host networking, Docker socket mounts, broad host mounts, host package-manager commands, or host-root mutation.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add ready, blocked, missing-artifact, unsafe-config, JSON report, Markdown report, and no-execution tests.
- Run the SW7 verification commands and report exact commands run.
```

## SW8: Claude Intake Guardrails

### Mission

Strengthen local review tooling so Claude-produced branches are easier to merge safely and harder to accidentally stage with cache files, protected docs, version drift, or unclassified lane changes.

### Start from

- `scripts/mainline_integration_review.rb`
- `test/test_mainline_integration_review.rb`
- `scripts/verify_layout.rb`
- `scripts/implementation_evidence_report.rb`
- `docs/claude-code-dispatch-runbook.md`
- `docs/mainline-integration-checkpoint.md`

### Deliver

- Review output that flags protected document diffs, `.gocache/`, `tmp/`, build artifacts, local logs, private config, missing version updates for code changes, stale version strings, and unclassified files.
- A staging recommendation grouped by lane, with explicit files to exclude.
- JSON and Markdown report coverage.
- Tests for clean, protected-doc, cache-file, stale-version, missing-version-bump, and unclassified-file cases.

### Keep disabled

- Automatic staging.
- Automatic committing.
- Automatic tag creation.
- Destructive cleanup.
- Host-root mutation.

### Required verification

```text
ruby -Ilib test/test_mainline_integration_review.rb
ruby -Ilib test/test_implementation_evidence_report.rb
ruby scripts/mainline_integration_review.rb --format json
ruby scripts/mainline_integration_review.rb --format markdown
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable Claude Code prompt

```text
Implement SW8 from docs/claude-code-second-wave-task-batch.md.

Target outcome:
- Strengthen Claude intake guardrails so local review reports catch protected document diffs, cache/tmp/build artifacts, stale version strings, missing version bumps, and unclassified changed files before staging.

Hard constraints:
- Do not automatically stage, commit, tag, push, delete files, run destructive cleanup, or mutate the host root.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add clean, protected-doc, cache-file, stale-version, missing-version-bump, unclassified-file, JSON report, and Markdown report tests.
- Run the SW8 verification commands and report exact commands run.
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

1. Confirm the branch implemented exactly one `SW` task.
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
