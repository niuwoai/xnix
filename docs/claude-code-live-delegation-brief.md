# Claude Code Live Delegation Brief

> Last updated: 2026-07-24 | Mainline checkpoint: v0.2.640-rc8 | Formal release: v0.2.640 pending full smoke

This is the short live brief for handing the next Xnix work to Claude Code. It is meant to be copied into a Claude Code session one task at a time.

Use this brief when Codex is keeping mainline review, merge decisions, version promotion, and host-safety decisions, while Claude Code implements one bounded branch.

## Current Workspace Gate

Before dispatching Claude Code, check the local workspace:

```text
git status --short --branch
```

If the workspace contains Codex-owned work in progress, do not ask Claude Code to start a different task in the same checkout. Either let Codex finish and commit that work first, or explicitly assign Claude Code to finish only that in-progress task.

No Codex-owned dirty lane should be present after the local v0.2.640-rc8 checkpoint. If C8W3, C8W4, C8W5, C8W7, or C8W8 files are dirty, treat that as interrupted repair work rather than permission to start a new lane:

```text
C8W3 Runtime owner service launch request envelope guard
C8W4 KDE controlled-launch action surface audit
C8W5 Managed launcher acceptance report
C8W7 Desktop-trigger dry-run request review
C8W8 Desktop-trigger service call materialization
```

If those files are dirty, Claude Code must not start `C8W6` in the same checkout.

## Global Rules for Every Claude Code Task

Claude Code must:

- Implement exactly one selected task.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior.
- Use Ruby only for tests, smoke scripts, reports, and lightweight developer tooling.
- Keep KDE as a replaceable presentation shell.
- Keep the Xnix Compatibility Runtime as the owner of policy, evidence, receipts, state roots, launch planning, backend selection, and dispatch.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Update `docs/xnix-current-mainline.md` when the mainline state changes.
- Add targeted tests for success, blocked input, malformed input, redaction, and no-side-effect behavior.
- Run the task-specific verification commands, `ruby scripts/verify_layout.rb`, and `git diff --check`.
- Report exact commands run and exact commands skipped.

Claude Code must not:

- Modify `docs/claude-code-implementation-packages.md`.
- Run Docker, QEMU, Wine, Proton, or a real compatibility backend unless the selected task explicitly says so and the human operator approves it first.
- Pull images, fetch packages, restart Colima, restart Docker, change host networking, or run host package-manager commands.
- Use privileged containers, host networking, Docker socket mounts, broad host-directory mounts, or host-root mutation.
- Enable production D-Bus ownership.
- Add real Portal transport calls.
- Write Runtime production state unless the task explicitly asks for a state-root-scoped preview writer.
- Write KDE configuration.
- Move Runtime policy into KDE, Ruby smoke scripts, or the C smoke adapter.
- Expose state-root paths, cache-root paths, raw launcher paths, executable paths, backend names, backend commands, raw launcher output, host paths, secrets, tokens, credentials, usernames, or environment variables in desktop-facing output.

## Dispatch Order

Use this order unless Codex or the human operator explicitly chooses a repair branch:

| Order | Task | Suggested branch | Safe to start while C8W3 is dirty | Outcome |
| --- | --- | --- | --- | --- |
| Done | `C8W3` Runtime owner service launch request envelope guard | `codex/owner-service-launch-envelope-guard` | Repair only | Present locally in v0.2.640-rc4. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W4` KDE controlled-launch action surface audit | `codex/kde-controlled-launch-action-surface-audit` | Repair only | Present locally in v0.2.640-rc5. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W5` Managed launcher acceptance report | `codex/managed-launcher-acceptance-report` | Repair only | Present locally in v0.2.640-rc6. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W7` Desktop-trigger dry-run request review | `codex/desktop-trigger-dry-run-request-review` | Repair only | Present locally in v0.2.640-rc7. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W8` Desktop-trigger service call materialization | `codex/desktop-trigger-service-call-materialization` | Repair only | Present locally in v0.2.640-rc8. Do not dispatch again unless a reviewer asks for repair. |
| 1 | `C8W6` Post-checkpoint promotion checklist | `codex/post-checkpoint-promotion-checklist` | No | Produce a deterministic checklist for promoting `v0.2.640` after full smoke passes. |

## Task C8W3: Finish Runtime Owner Service Launch Request Envelope Guard

Use this task only if Claude Code is explicitly assigned to repair or finish interrupted C8W3 work.

### Mission

Finish a fail-closed owner-service guard that validates desktop-triggered launch request envelopes before any future service method can consume them.

The guard must prove that KDE can only forward opaque evidence handles. KDE must not replay, expand, or reconstruct Runtime-owned launch inputs.

### Expected files

- `internal/runtime/owner/launch_envelope_guard.go`
- `internal/runtime/owner/launch_envelope_guard_test.go`
- `cmd/xnix-runtime-go/owner_service_launch_envelope_guard_commands.go`
- `cmd/xnix-runtime-go/owner_service_launch_envelope_guard_cli_test.go`
- `cmd/xnix-runtime-go/main.go`
- `scripts/verify_layout.rb`
- `scripts/mainline_integration_review.rb`
- `VERSION`
- `CHANGELOG.md`
- `PRODUCT_OVERVIEW.md`
- `docs/xnix-current-mainline.md`
- `docs/claude-code-current-dispatch-picks.md`
- `docs/claude-code-next-actions-handoff.md`

### Required behavior

- Accept only one safe evidence handle:
  - `evidence-id`
  - `evidence-relative-path`
- Validate route id, method id, evidence handle shape, digest binding, expected action id, freshness state, replay marker, and caller role.
- Block KDE-facing owner-only inputs:
  - state root
  - cache root
  - launcher path
  - timeout value
  - raw executable path
  - backend command
  - receipt id
  - session id
  - dispatch id
- Keep these side effects disabled:
  - receipt writes
  - production authorization acceptance
  - service call dispatch
  - D-Bus ownership
  - desktop launch
  - backend launch
  - execution start
  - Runtime state writes
  - network access
  - host-root mutation

### Required verification

```text
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/owner ./internal/runtime/appidentity ./cmd/xnix-runtime-go -run 'TestPreviewLaunchEnvelopeGuard|TestOwnerServiceLaunchEnvelopeGuard' -count=1
ruby -Ilib test/test_dbus_controlled_launch_owner_fixture_smoke_script.rb
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable prompt

```text
Implement C8W3 from docs/claude-code-live-delegation-brief.md.

Target outcome:
- Finish the Runtime owner service launch request envelope guard and keep the desktop boundary evidence-only.

Hard constraints:
- Do not start C8W4, C8W5, C8W6, or any unrelated lane.
- Do not write receipts, accept production authorization, dispatch service calls, enable D-Bus ownership, launch desktop actions, start compatibility backends, run Docker, run QEMU, run Wine, fetch from the network, or mutate the host root.
- Do not allow KDE-facing input or output to include state roots, cache roots, launcher paths, timeout values, raw executable paths, backend commands, receipt fields, session ids, dispatch ids, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, PRODUCT_OVERVIEW.md, docs/xnix-current-mainline.md, and layout/review gates as needed.
- Add accepted, missing-evidence, mismatched-route, stale, replay, owner-args, malformed, redaction, and no-side-effect tests.
- Run the C8W3 verification commands and report exact commands run and skipped.
```

## Task C8W4: KDE Controlled-Launch Action Surface Audit

### Mission

Add a KDE-facing audit preview that checks whether the controlled-launch action remains presentation-only before a real desktop-triggered staged launch smoke is attempted.

### Start from

- `kde/actions/xnix-runtime-status-controlled-launch.desktop`
- `cmd/xnix-runtime-go/kde_controlled_launch_action_commands.go`
- `internal/runtime/appidentity`
- `scripts/kde_controlled_launch_action_smoke.rb`

### Deliver

- A CLI command such as `kde-controlled-launch-action-surface-audit-preview`.
- Checks for desktop action id, label, public method, evidence-only argument shape, no owner-service arguments, no state-root access, no raw executable path, no backend details, and no execution enablement.
- A redacted KDE-safe report that says whether the action surface is safe for the next human-authorized smoke.
- Tests for safe metadata, unsafe owner arguments, unsafe backend terms, malformed desktop metadata, missing evidence handle, redaction, and no-side-effect behavior.

### Required verification

```text
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go -run 'TestKDEControlledLaunchAction|TestKDEControlledLaunchActionSurfaceAudit' -count=1
ruby -Ilib test/test_kde_controlled_launch_action_smoke_script.rb
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable prompt

```text
Implement C8W4 from docs/claude-code-live-delegation-brief.md.

Target outcome:
- Add a KDE controlled-launch action surface audit preview that proves the KDE action remains presentation-only.

Hard constraints:
- Do not enable launch, call D-Bus, write KDE configuration, write Runtime state, start backends, run Docker, run QEMU, run Wine, fetch from the network, or mutate the host root.
- Do not expose owner-service arguments, state roots, cache roots, launcher paths, raw executable paths, backend commands, backend names, backend details, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, PRODUCT_OVERVIEW.md, docs/xnix-current-mainline.md, and layout/review gates as needed.
- Add safe, unsafe-owner-args, unsafe-backend-terms, malformed-desktop-metadata, missing-evidence-handle, redaction, and no-side-effect tests.
- Run the C8W4 verification commands and report exact commands run and skipped.
```

## Task C8W5: Managed Launcher Acceptance Report

### Mission

Tie the known Windows app guest-smoke evidence, managed launcher bridge evidence, Runtime-status action handoff, and Compatibility Center projection into a user-safe acceptance report.

The report should answer:

```text
What evidence proves that the managed known Windows app lane is ready for the next human-authorized desktop-triggered smoke?
```

### Start from

- `internal/runtime/appidentity`
- `internal/runtime/winapp`
- `cmd/xnix-runtime-go/windows_compatibility_commands.go`
- `cmd/xnix-runtime-go/kde_desktop_trigger_staged_invocation_readiness_commands.go`
- `scripts/staged_launcher_dispatch_smoke.rb`

### Deliver

- A Go read model and CLI command such as `managed-launcher-acceptance-report-preview`.
- Evidence sections for known-app identity, artifact digest, launch authorization, session-gated review, controlled execution session, managed launcher request, guest smoke boundary, Runtime-status evidence handoff, KDE-safe projection, and full-checkpoint dependency.
- States such as `accepted`, `blocked`, `missing-evidence`, `stale-evidence`, `malformed`, `unsafe`, and `needs-full-checkpoint`.
- Redacted output suitable for Compatibility Center or a support bundle.
- Tests for accepted fixture evidence, missing artifact, stale digest, missing review receipt, blocked full checkpoint, redaction, and no-side-effect behavior.

### Required verification

```text
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/appidentity ./internal/runtime/winapp ./cmd/xnix-runtime-go -run 'TestPreviewManagedLauncherAcceptance|TestManagedLauncherAcceptance|TestDesktopTriggerStagedInvocationReadiness|TestKnownApp' -count=1
ruby -Ilib test/test_staged_launcher_dispatch_smoke_script.rb
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable prompt

```text
Implement C8W5 from docs/claude-code-live-delegation-brief.md.

Target outcome:
- Add a managed launcher acceptance report that ties the known Windows app evidence chain into a user-safe readiness packet.

Hard constraints:
- Do not launch desktop actions, start compatibility backends, run Docker, run QEMU, run Wine, fetch artifacts from the network, call D-Bus, write Runtime production state, write KDE configuration, or mutate the host root.
- Do not expose state roots, cache roots, launcher paths, raw executable paths, backend commands, backend names, backend details, raw launcher output, host paths, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, PRODUCT_OVERVIEW.md, docs/xnix-current-mainline.md, and layout/review gates as needed.
- Add accepted, missing-artifact, stale-digest, missing-review, blocked-full-checkpoint, redaction, and no-side-effect tests.
- Run the C8W5 verification commands and report exact commands run and skipped.
```

## Task C8W6: Post-Checkpoint Promotion Checklist

### Mission

Create a deterministic promotion checklist that Codex can use after `ruby scripts/full_smoke.rb` passes to promote the release candidate to formal `v0.2.640`.

This task is documentation and release tooling only. It must not claim that the full smoke passed.

### Start from

- `scripts/full_smoke.rb`
- `scripts/verify_layout.rb`
- `scripts/mainline_integration_review.rb`
- `CHANGELOG.md`
- `PRODUCT_OVERVIEW.md`
- `docs/xnix-current-mainline.md`

### Deliver

- A checklist document such as `docs/post-checkpoint-promotion-checklist-0640.md`.
- PASS evidence requirements for full smoke, restricted container lanes, QEMU serial smoke, guest Windows app smoke, KDE controlled-launch action smoke, D-Bus fixture smoke, layout verification, mainline integration review, and diff hygiene.
- Exact version-promotion steps from the active release candidate to `0.2.640`, but only after PASS evidence exists.
- A section for commands that must not be run without human approval, including Colima, Docker, QEMU, host package managers, and network repair commands.
- Tests or script-level checks if any release tooling changes.

### Required verification

```text
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable prompt

```text
Implement C8W6 from docs/claude-code-live-delegation-brief.md.

Target outcome:
- Add a deterministic post-checkpoint promotion checklist for promoting the active v0.2.640 release candidate to formal v0.2.640 only after full smoke passes.

Hard constraints:
- Do not run full smoke unless the human operator explicitly approves it.
- Do not run Docker, QEMU, Wine, Colima, host package-manager commands, network repair commands, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Do not claim v0.2.640 is released.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Add the checklist and update mainline docs if needed.
- Run the C8W6 verification commands and report exact commands run and skipped.
```

## Task C8W7: Desktop-Trigger Dry-Run Request Review

### Mission

Add the next review-only packet for a future human-authorized desktop-triggered staged launch. This packet must review the proposed desktop-trigger request without dispatching any service call.

### Start from

- `internal/runtime/owner`
- `internal/runtime/appidentity`
- `cmd/xnix-runtime-go`
- `kde/actions/xnix-runtime-status-controlled-launch.desktop`
- `scripts/kde_controlled_launch_action_smoke.rb`

### Deliver

- A Go read model and CLI command such as `desktop-trigger-dry-run-request-review-preview`.
- Inputs limited to safe evidence handles and public action metadata.
- Checks that consume C8W3 and C8W4 outputs when available.
- A blocked-by-default result unless full-checkpoint PASS evidence exists and the envelope/action audits are safe.
- No service call dispatch, no D-Bus call, no Runtime state write, no KDE config write, and no backend launch.

### Required verification

```text
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/owner ./internal/runtime/appidentity ./cmd/xnix-runtime-go -run 'TestDesktopTriggerDryRunRequestReview|TestPreviewLaunchEnvelopeGuard|TestKDEControlledLaunchActionSurfaceAudit' -count=1
ruby -Ilib test/test_kde_controlled_launch_action_smoke_script.rb
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable prompt

```text
Implement C8W7 from docs/claude-code-live-delegation-brief.md.

Target outcome:
- Add a review-only desktop-trigger dry-run request packet for the next human-authorized staged launch attempt.

Hard constraints:
- Do not dispatch service calls, call D-Bus, enable desktop launch, start compatibility backends, run Docker, run QEMU, run Wine, fetch from the network, write Runtime production state, write KDE configuration, or mutate the host root.
- Do not expose state roots, cache roots, launcher paths, raw executable paths, backend commands, backend names, backend details, raw launcher output, host paths, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, PRODUCT_OVERVIEW.md, docs/xnix-current-mainline.md, and layout/review gates as needed.
- Add accepted-review, blocked-missing-full-checkpoint, blocked-unsafe-envelope, blocked-unsafe-action, malformed, redaction, and no-side-effect tests.
- Run the C8W7 verification commands and report exact commands run and skipped.
```

## Claude Completion Template

Claude Code should return this exact summary shape:

```text
Task:
Branch:
Version:
Changed files:
Tests run:
Tests skipped:
Safety checks:
Known limitations:
```

## Codex Intake Checklist

When Claude Code returns a branch, Codex should:

1. Run `git status --short --branch`.
2. Run `git diff --name-only`.
3. Confirm `docs/claude-code-implementation-packages.md` is unchanged.
4. Run `ruby scripts/mainline_integration_review.rb --format json`.
5. Confirm `protected_claude_file_modified` is `false`.
6. Confirm `unclassified_file_count` is `0`.
7. Run the task-specific verification commands that Claude reported.
8. Run `ruby scripts/verify_layout.rb`.
9. Run `git diff --check`.
10. Stage only the reviewed task files, version files, and documentation files.
