# Claude Code Next Delegation Pack

> Last updated: 2026-07-24 | Baseline: v0.2.640-rc7 | Owner: Codex mainline review

This document is a copy-first handoff pack for assigning the next bounded Xnix work to Claude Code.

It does not replace `docs/xnix-current-mainline.md`. The mainline document remains the planning authority. This pack exists so the human operator can copy one task prompt at a time into Claude Code while Codex keeps review, merge, version promotion, and host-safety decisions.

## Current Baseline

The local checkpoint candidate is `v0.2.640-rc7`.

Already present locally:

- `C8W2` Desktop-trigger staged invocation readiness packet in `v0.2.640-rc3`.
- `C8W3` Runtime owner service launch request envelope guard in `v0.2.640-rc4`.
- `C8W4` KDE controlled-launch action surface audit in `v0.2.640-rc5`.
- `C8W5` Managed launcher acceptance report in `v0.2.640-rc6`.
- `C8W7` Desktop-trigger dry-run request review in `v0.2.640-rc7`.

Do not ask Claude Code to reimplement those completed tasks unless Codex or a reviewer asks for a repair branch.

The formal `v0.2.640` release is still blocked until full smoke passes:

```text
ruby scripts/full_smoke.rb
```

Claude Code must not run that command from this pack. Full smoke, Docker, QEMU, Wine, Colima, and host-environment repair stay operator-owned unless the human explicitly authorizes them in that exact task.

## Dispatch Protocol

Before assigning any task, run:

```text
git status --short --branch
```

If there are dirty files, assign Claude Code only to finish or repair the same dirty lane. Do not start a new task in the same checkout while another lane is dirty.

Assign exactly one task at a time. Claude Code should return after that task and must not chain into the next one.

Claude Code must not modify:

```text
docs/claude-code-implementation-packages.md
```

## Hard Boundaries

Every task from this pack must:

- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior.
- Use Ruby only for tests, smoke scripts, reports, fixtures, and lightweight developer tooling.
- Keep KDE as a replaceable presentation shell.
- Keep the Xnix Compatibility Runtime as the owner of policy, evidence, receipts, state roots, launch planning, backend selection, and dispatch gates.
- Update `VERSION`, `CHANGELOG.md`, `PRODUCT_OVERVIEW.md`, `docs/xnix-current-mainline.md`, and review/layout gates when code changes.
- Add targeted tests for success, blocked input, malformed input, redaction, and no-side-effect behavior.
- Run task-specific verification commands, `ruby scripts/verify_layout.rb`, and `git diff --check`.
- Report exact commands run and exact commands skipped.

Every task from this pack must not:

- Run Docker, QEMU, Wine, Proton, Colima, or a real compatibility backend.
- Pull images, fetch packages, restart host services, use host package managers, or repair host networking.
- Use privileged containers, host networking, Docker socket mounts, broad host-directory mounts, or host-root mutation.
- Enable production D-Bus ownership.
- Add real Portal transport calls.
- Write Runtime production state unless the selected task explicitly asks for a state-root-scoped preview writer.
- Write KDE configuration.
- Move Runtime policy into KDE, Ruby smoke scripts, or the C smoke adapter.
- Expose state-root paths, cache-root paths, launcher paths, executable paths, backend names, backend commands, raw launcher output, host paths, secrets, tokens, credentials, usernames, or environment variables in KDE-facing or user-facing output.

## Recommended Queue

| Order | Task | Suggested branch | Outcome |
| --- | --- | --- | --- |
| Done | `C8W7` Desktop-trigger dry-run request review | `codex/desktop-trigger-dry-run-request-review` | Present locally in v0.2.640-rc7. Do not dispatch again unless a reviewer asks for repair. |
| 1 | `C8W6` Post-checkpoint promotion checklist | `codex/post-checkpoint-promotion-checklist` | Produce a deterministic checklist for promoting `v0.2.640` only after full smoke passes. |

Stop after each task and return the branch for Codex review.

## Task C8W6: Post-Checkpoint Promotion Checklist

### Mission

Create a deterministic promotion checklist for promoting the active `v0.2.640` release candidate only after full-smoke PASS evidence exists.

This is documentation and release-tooling support. It must not claim that full smoke passed.

### Expected implementation shape

Add:

```text
docs/post-checkpoint-promotion-checklist-0640.md
```

Only update scripts if the checklist needs a small, non-executing verification helper.

### Required content

- PASS evidence requirements for full smoke, restricted container lanes, QEMU serial smoke, `sshd`, known Windows app smoke, fixture Windows app smoke, KDE controlled-launch action smoke, D-Bus fixture smoke, layout verification, mainline review, and diff hygiene.
- Exact promotion steps from the active release candidate to `0.2.640`.
- A clear statement that promotion is forbidden until `ruby scripts/full_smoke.rb` passes.
- A list of commands Claude Code must not run without human approval.
- A rollback note for failed promotion review.

### Required verification

```text
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable prompt

```text
Implement C8W6 from docs/claude-code-next-delegation-pack.md.

Target outcome:
- Add a deterministic post-checkpoint promotion checklist for promoting the active v0.2.640 release candidate to formal v0.2.640 only after full smoke passes.

Hard constraints:
- Implement only C8W6.
- Do not run full smoke unless the human operator explicitly approves it in this task.
- Do not run Docker, QEMU, Wine, Proton, Colima, host package-manager commands, network repair commands, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Do not claim v0.2.640 is released.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Add docs/post-checkpoint-promotion-checklist-0640.md.
- Update only directly necessary docs or non-executing verification helpers.
- Run the C8W6 verification commands and report exact commands run and skipped.
```

## Task C8W7: Desktop-Trigger Dry-Run Request Review

### Mission

Add a review-only packet for a future human-authorized desktop-triggered staged launch.

The packet must review the proposed desktop-trigger request without dispatching any service call.

### Expected implementation shape

Add or update files in this lane only:

```text
internal/runtime/owner/desktop_trigger_dry_run_request_review.go
internal/runtime/owner/desktop_trigger_dry_run_request_review_test.go
cmd/xnix-runtime-go/desktop_trigger_dry_run_request_review_commands.go
cmd/xnix-runtime-go/desktop_trigger_dry_run_request_review_cli_test.go
cmd/xnix-runtime-go/main.go
scripts/verify_layout.rb
scripts/mainline_integration_review.rb
VERSION
CHANGELOG.md
PRODUCT_OVERVIEW.md
docs/xnix-current-mainline.md
docs/claude-code-current-dispatch-picks.md
docs/claude-code-live-delegation-brief.md
docs/claude-code-next-actions-handoff.md
```

### Required behavior

- Add a CLI command such as `desktop-trigger-dry-run-request-review-preview`.
- Accept only safe evidence handles and public action metadata.
- Consume C8W3 and C8W4 outputs when available.
- Return blocked by default unless full-checkpoint evidence exists and envelope/action audits are safe.
- Keep service-call dispatch, D-Bus calls, Runtime writes, KDE configuration writes, desktop launch, backend launch, network access, and host mutation disabled.

### Required tests

Add tests for:

- Accepted review when full-checkpoint evidence and safe audits are present.
- Blocked missing full checkpoint.
- Blocked unsafe envelope.
- Blocked unsafe KDE action surface.
- Malformed evidence handle.
- Redaction.
- No-side-effect behavior.

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
Implement C8W7 from docs/claude-code-next-delegation-pack.md.

Target outcome:
- Add a review-only desktop-trigger dry-run request packet for the next human-authorized staged launch attempt.

Hard constraints:
- Implement only C8W7.
- Do not dispatch service calls, call D-Bus, enable desktop launch, start compatibility backends, run Docker, run QEMU, run Wine, fetch from the network, write Runtime production state, write KDE configuration, or mutate the host root.
- Do not expose state roots, cache roots, launcher paths, raw executable paths, backend commands, backend names, backend details, raw launcher output, host paths, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, PRODUCT_OVERVIEW.md, docs/xnix-current-mainline.md, scripts/verify_layout.rb, and scripts/mainline_integration_review.rb as needed.
- Add accepted-review, blocked-missing-full-checkpoint, blocked-unsafe-envelope, blocked-unsafe-action, malformed, redaction, and no-side-effect tests.
- Run the C8W7 verification commands and report exact commands run and skipped.
```

## Claude Completion Template

Claude Code should finish each task with this exact shape:

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
7. Run the task-specific verification commands Claude reported.
8. Run `ruby scripts/verify_layout.rb`.
9. Run `git diff --check`.
10. Inspect version metadata, changelog entry, and product overview updates when code changed.
11. Stage only the reviewed task files, version files, and documentation files by explicit path.
12. Do not use `git add .`.
13. Commit one coherent small version after targeted tests pass.
14. Tag only when the branch is accepted as a local checkpoint.

## Stop Rule

If `C8W6` is present locally, stop and ask Codex for a refreshed mainline plan before assigning more work.

Do not invent a new wave task from this document.
