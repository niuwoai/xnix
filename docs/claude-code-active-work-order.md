# Claude Code Active Work Order

> Last updated: 2026-07-24 | Baseline: v0.2.640-rc20 | Formal release: v0.2.640 remains blocked until full smoke passes

This document is the current short-form work order for asking Claude Code to implement the next bounded Xnix mainline task. It intentionally excludes completed task details so Claude Code does not drift back into already-landed work.

Use this document when the operator wants one implementation branch or patch from Claude Code, then Codex review before the next task.

## Starting Rules

Claude Code must start by running:

```text
git status --short --branch
```

If the checkout is dirty, Claude Code must stop and ask whether to continue that dirty lane or wait for Codex review.

Claude Code must implement exactly one selected task from this document.

Do not modify:

```text
docs/claude-code-implementation-packages.md
```

Completed tasks must not be reimplemented unless Codex or a reviewer explicitly asks for a repair:

- `C9W2` Desktop-trigger request preflight smoke.
- `C9W3` Merge readiness consumes request preflight smoke evidence.
- `C9W4` Release evidence index consumes request preflight smoke evidence.

## Safety Boundaries

Claude Code must not run or introduce:

- Docker, QEMU, Wine, Proton, Colima, or a real compatibility backend.
- `ruby scripts/full_smoke.rb`.
- Network fetches, package-manager calls, host service restarts, or host networking repair.
- Privileged containers, Docker socket mounts, broad host-directory mounts, or host-root mutation.
- Production D-Bus ownership or real XDG Portal transport calls.
- Desktop launch, backend launch, service dispatch, Runtime production writes, or KDE configuration writes.

Claude Code must keep:

- Runtime policy in Go-owned Runtime code.
- KDE as a replaceable presentation shell.
- Ruby limited to tests, smoke orchestration, reports, fixtures, and lightweight tooling.
- All project-facing source, comments, fixtures, tests, CLI output, and documentation in English.
- Redaction for state roots, cache roots, launcher paths, executable paths, backend names, backend commands, raw launcher output, host paths, secrets, tokens, credentials, usernames, and environment variables.

## Active Queue

Dispatch exactly one task at a time.

| Order | Task | Suggested branch | Dispatch condition |
| --- | --- | --- | --- |
| 1 | `C9W5` Runtime owner request receipt preview | `codex/runtime-owner-request-receipt-preview` | Dispatch only after formal `v0.2.640` promotion or explicit Codex/operator approval to continue pre-release modeling. |
| 2 | `C9W6` KDE Center desktop-trigger preflight card | `codex/kde-center-desktop-trigger-preflight-card` | Dispatch only after `C9W5` lands and Codex approves the next handoff. |

## Task C9W5: Runtime Owner Request Receipt Preview

### Mission

Add a Go-owned preview for the future owner-managed request receipt required before a real desktop-triggered request can move from preflight into a controlled Runtime request object.

This is a preview only. It must not persist production state or dispatch anything.

### Required behavior

- Consume the existing `desktop-trigger-request-preflight-preview` result shape.
- Accept only the ready preflight state.
- Reject missing promotion, blocked preflight, malformed input, unsafe side-effect flags, stale evidence, and replay-marker mismatch.
- Produce an opaque receipt plan suitable for review, not a durable production receipt.
- Keep receipt persistence, request-object writes, permission grants, service dispatch, D-Bus calls, desktop launch, backend launch, Runtime production writes, KDE configuration writes, network access, and host mutation disabled.
- Keep all KDE-facing and user-facing output redacted.

### Expected tests

- Ready preflight produces a review-only receipt plan.
- Missing promotion is blocked.
- Blocked preflight is blocked.
- Malformed input is blocked without leaking raw values.
- Replay-marker mismatch is blocked.
- Unsafe side-effect flags are blocked.
- Redaction covers all owner-only and host-specific details.

### Verification

Claude Code may run:

```text
go test ./internal/runtime/owner ./cmd/xnix-runtime-go -run 'Test.*DesktopTrigger.*Request.*Receipt|Test.*DesktopTrigger.*Request.*Preflight' -count=1
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

Claude Code must not run:

```text
ruby scripts/full_smoke.rb
ruby scripts/container.rb full-smoke
docker ...
qemu-system-...
wine ...
colima ...
```

### Copyable Claude Code Prompt

```text
Implement C9W5 from docs/claude-code-active-work-order.md only if Codex or the operator has explicitly approved pre-release modeling, or if formal v0.2.640 promotion is already complete.

Scope:
- Add a Go-owned desktop-trigger owner request receipt preview.
- Consume the existing request preflight packet and fail closed unless the preflight is ready.
- Produce a review-only opaque receipt plan.
- Do not persist receipts, create request objects, dispatch service calls, call D-Bus, launch desktop actions, start backends, write Runtime production state, write KDE configuration, use network, use containers, or mutate the host.
- Do not modify docs/claude-code-implementation-packages.md.

Run the C9W5 verification commands from docs/claude-code-active-work-order.md and report exact commands run and skipped.
```

## Task C9W6: KDE Center Desktop-Trigger Preflight Card

### Mission

Expose the desktop-trigger request preflight state as a KDE-safe read model card so KDE can present readiness without owning policy or seeing owner-only arguments.

### Required behavior

- Add or extend the KDE Center page preview to consume redacted preflight state.
- Display only safe readiness state, blocker summary, and operator action hints.
- Keep KDE presentation-only.
- Do not expose owner service call arguments, CLI arguments, state roots, cache roots, launcher paths, backend details, raw environment values, usernames, secrets, tokens, credentials, or host paths.
- Keep desktop launch, backend launch, D-Bus calls, Runtime writes, KDE configuration writes, network access, and host mutation disabled.

### Expected tests

- Ready preflight maps to a safe KDE card.
- Blocked preflight maps to safe blocker text.
- Missing evidence is presented as unavailable, not successful.
- Owner-only arguments and host paths are redacted.
- KDE does not become the policy owner.
- All side-effect flags stay disabled.

### Verification

Claude Code may run:

```text
go test ./internal/runtime/appidentity ./internal/runtime/owner ./cmd/xnix-runtime-go -run 'Test.*KDE.*Center|Test.*DesktopTrigger.*Preflight' -count=1
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

Claude Code must not run:

```text
ruby scripts/full_smoke.rb
ruby scripts/container.rb full-smoke
docker ...
qemu-system-...
wine ...
colima ...
```

### Copyable Claude Code Prompt

```text
Implement C9W6 from docs/claude-code-active-work-order.md only after C9W5 lands and Codex approves this next handoff.

Scope:
- Add a KDE-safe read model card for desktop-trigger request preflight state.
- KDE must remain presentation-only and must not receive owner service arguments, state roots, cache roots, launcher paths, backend details, raw environment values, usernames, secrets, tokens, credentials, or host paths.
- Do not dispatch service calls, call D-Bus, launch desktop actions, start backends, write Runtime production state, write KDE configuration, use network, use containers, or mutate the host.
- Do not modify docs/claude-code-implementation-packages.md.

Run the C9W6 verification commands from docs/claude-code-active-work-order.md and report exact commands run and skipped.
```

## Return Requirements

Claude Code must return:

- The selected task id.
- A concise implementation summary.
- The exact files changed.
- The exact commands run.
- The exact commands intentionally skipped.
- Any blockers or assumptions.
- Confirmation that `docs/claude-code-implementation-packages.md` was not modified.

Codex must review the returned branch or patch before dispatching another task.
