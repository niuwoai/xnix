# Claude Code Next Dispatch Brief

> Last updated: 2026-07-24 | Baseline: v0.2.640-rc21 | Formal release: v0.2.640 remains blocked until full smoke passes

This document is the short operator handoff for asking Claude Code to implement the next bounded Xnix tasks while Codex keeps review, merge, release-promotion, and host-safety decisions.

Use this brief when the operator wants Claude Code to do one focused implementation branch and return the result for Codex review.

For the current short-form work order, use `docs/claude-code-active-work-order.md`. This longer brief remains the task reference and historical context.

Do not modify `docs/claude-code-implementation-packages.md` from any task in this brief.

## Current State

The local baseline is `v0.2.640-rc21`.

Already present in the current candidate:

- Runtime-owned desktop trigger readiness, envelope guard, KDE action surface audit, managed launcher acceptance, dry-run review, service-call materialization, and request preflight previews.
- Staged launcher dispatch smoke can consume Runtime-materialized owner service call arguments when a human-authorized smoke candidate flag is present.
- The private D-Bus controlled-launch fixture consumes the same Runtime service-call materialization path.
- Full checkpoint promotion, merge readiness, and release evidence packets exist as read-only release decision tools.
- `docs/post-checkpoint-promotion-checklist-0640.md` defines the human-owned promotion checklist.
- `scripts/desktop_trigger_request_preflight_smoke.rb` is present as a targeted smoke for the request preflight lane.
- `scripts/merge_readiness_packet.rb` consumes existing desktop-trigger request preflight smoke JSON evidence without running that smoke by default.
- `scripts/release_evidence_index.rb` classifies existing desktop-trigger request preflight smoke JSON evidence without running that smoke by default.

The formal `v0.2.640` release is still blocked until this operator-owned command passes:

```text
ruby scripts/full_smoke.rb
```

Claude Code must not run that command unless the human operator gives explicit task-local approval.

## Hard Boundaries

Before starting any task, Claude Code must run:

```text
git status --short --branch
```

If the checkout is dirty, Claude Code must stop and ask whether to continue that dirty lane or wait for Codex review.

Every task must:

- Implement exactly one selected task.
- Keep all source, comments, fixtures, tests, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior and stable service boundaries.
- Use Ruby only for tests, smoke orchestration, reports, fixtures, and lightweight tooling.
- Keep KDE as a replaceable presentation shell.
- Keep the Xnix Compatibility Runtime as the policy owner.
- Update `VERSION`, `CHANGELOG.md`, `PRODUCT_OVERVIEW.md`, `docs/xnix-current-mainline.md`, layout checks, and mainline review gates when code changes.
- Add targeted tests for ready, blocked, malformed, redaction, and no-side-effect behavior where relevant.
- Run task-specific verification, `ruby scripts/verify_layout.rb`, and `git diff --check`.
- Report exact commands run and exact commands skipped.

Every task must not:

- Run Docker, QEMU, Wine, Proton, Colima, or a real compatibility backend.
- Run `ruby scripts/full_smoke.rb`.
- Pull images, fetch packages, restart host services, use host package managers, or repair host networking.
- Use privileged containers, host networking, Docker socket mounts, broad host-directory mounts, or host-root mutation.
- Enable production D-Bus ownership.
- Add real XDG Portal transport calls.
- Write Runtime production state unless the selected task explicitly asks for a state-root-scoped preview writer.
- Write KDE configuration.
- Move Runtime policy into KDE, Ruby smoke scripts, or the C smoke adapter.
- Expose state-root paths, cache-root paths, launcher paths, executable paths, backend names, backend commands, raw launcher output, host paths, secrets, tokens, credentials, usernames, or environment variables in KDE-facing or user-facing output.

## Dispatch Order

Assign exactly one task at a time. Stop after each task and return the branch or diff for Codex review.

| Order | Task | Suggested branch | Dispatch condition |
| --- | --- | --- | --- |
| Done | `C9W2` Desktop-trigger request preflight smoke | `codex/desktop-trigger-request-preflight-smoke` | Present locally in v0.2.640-rc16. Do not dispatch again unless Codex or a reviewer asks for repair. |
| Done | `C9W3` Merge readiness consumes request preflight smoke evidence | `codex/merge-readiness-preflight-smoke-evidence` | Present locally in v0.2.640-rc17. Do not dispatch again unless Codex or a reviewer asks for repair. |
| Done | `C9W4` Release evidence index consumes request preflight smoke evidence | `codex/release-evidence-preflight-smoke-evidence` | Present locally in v0.2.640-rc18. Do not dispatch again unless Codex or a reviewer asks for repair. |
| 1 | `C9W5` Runtime owner request receipt preview | `codex/runtime-owner-request-receipt-preview` | Dispatch after formal `v0.2.640` promotion or explicit Codex approval to continue pre-release modeling. |
| 2 | `C9W6` KDE Center desktop-trigger preflight card | `codex/kde-center-desktop-trigger-preflight-card` | Dispatch after `C9W5` lands. |

## Task C9W2: Desktop-Trigger Request Preflight Smoke

### Mission

Add a lightweight smoke script that invokes the Go-owned `desktop-trigger-request-preflight-preview` command through the CLI and proves the preflight is usable from developer automation.

This is a targeted local smoke. It must not launch the desktop action or start any backend.

### Expected files

```text
scripts/desktop_trigger_request_preflight_smoke.rb
test/test_desktop_trigger_request_preflight_smoke_script.rb
scripts/container.rb
lib/xnix/container.rb
scripts/verify_layout.rb
VERSION
CHANGELOG.md
PRODUCT_OVERVIEW.md
docs/xnix-current-mainline.md
docs/claude-code-next-mainline-work-pack.md
docs/claude-code-immediate-mainline-handoff.md
```

### Required behavior

- Create a temporary state root under a project-local cache directory.
- Produce or record safe fixture evidence through existing Runtime CLI commands.
- Run `desktop-trigger-request-preflight-preview` without `--full-checkpoint-promoted` and verify:
  - `request_preflight_state=blocked-missing-promotion`;
  - `operator_request_ready=false`;
  - service dispatch, D-Bus calls, desktop launch, backend launch, Runtime writes, KDE configuration writes, and host-root mutation are disabled.
- Run `desktop-trigger-request-preflight-preview` with explicit promoted fixture state and verify:
  - `request_preflight_state=ready-for-operator-request`;
  - `operator_request_ready=true`;
  - `owner_service_call_shape_verified=true`;
  - service dispatch still remains disabled.
- Ensure output does not expose owner service arguments, CLI arguments, state roots, launcher paths, backend details, raw environment values, usernames, secrets, tokens, or host paths.
- Print:

```text
PASS: desktop-trigger request preflight smoke
```

### Required tests

- The script exists and is executable.
- The script invokes `desktop-trigger-request-preflight-preview`.
- The script verifies both blocked and ready preflight states.
- The script verifies redaction and no-side-effect flags.
- The container helper exposes a restricted command for this smoke without privileged containers, host networking, Docker socket mounts, or broad host mounts.

### Verification

Claude Code may run:

```text
ruby -Ilib test/test_desktop_trigger_request_preflight_smoke_script.rb
ruby scripts/desktop_trigger_request_preflight_smoke.rb
ruby -Ilib test/test_container.rb
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
Implement C9W2 from docs/claude-code-next-dispatch-brief.md.

Scope:
- Add a targeted desktop-trigger request preflight smoke that invokes desktop-trigger-request-preflight-preview through the CLI.
- Verify both blocked-missing-promotion and ready-for-operator-request behavior.
- Keep service dispatch, D-Bus calls, desktop launch, backend launch, Runtime writes, KDE configuration writes, network access, containers, QEMU, Wine, Colima, and host-root mutation disabled.
- Do not expose owner service arguments, CLI arguments, state roots, launcher paths, backend details, raw environment values, usernames, secrets, tokens, or host paths.
- Do not modify docs/claude-code-implementation-packages.md.

Expected files:
- scripts/desktop_trigger_request_preflight_smoke.rb
- test/test_desktop_trigger_request_preflight_smoke_script.rb
- scripts/container.rb
- lib/xnix/container.rb
- scripts/verify_layout.rb
- VERSION
- CHANGELOG.md
- PRODUCT_OVERVIEW.md
- docs/xnix-current-mainline.md
- docs/claude-code-next-mainline-work-pack.md
- docs/claude-code-immediate-mainline-handoff.md

Run the C9W2 verification commands from the dispatch brief and report exact commands run and skipped.
```

## Task C9W3: Merge Readiness Consumes Request Preflight Smoke Evidence

### Mission

Make `scripts/merge_readiness_packet.rb` consume existing request preflight smoke evidence as an optional readiness signal without running the smoke.

### Expected files

```text
scripts/merge_readiness_packet.rb
test/test_merge_readiness_packet.rb
scripts/verify_layout.rb
scripts/mainline_integration_review.rb
VERSION
CHANGELOG.md
PRODUCT_OVERVIEW.md
docs/xnix-current-mainline.md
```

### Required behavior

- Add an optional fixture flag such as `--desktop-trigger-request-preflight-smoke PATH`.
- Read existing JSON evidence only; do not run the smoke.
- Surface `desktop_trigger_request_preflight_smoke_status`.
- Keep missing smoke evidence non-blocking before formal promotion.
- Keep failed or malformed smoke evidence as a release-only blocker when the release candidate claims desktop-trigger request readiness.
- Preserve `automatic_release_tagging_enabled=false`.
- Preserve all no-side-effect flags.

### Verification

```text
ruby -Ilib test/test_merge_readiness_packet.rb
ruby scripts/merge_readiness_packet.rb --format json --offline-only
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable Claude Code Prompt

```text
Implement C9W3 from docs/claude-code-next-dispatch-brief.md.

Scope:
- Make merge readiness optionally consume existing desktop-trigger request preflight smoke evidence.
- Do not run the smoke, full smoke, Docker, QEMU, Wine, Colima, D-Bus, desktop launch, backend launch, network, or host mutation.
- Keep the new signal release-only unless the existing merge readiness contract already defines a stricter blocker class.
- Do not modify docs/claude-code-implementation-packages.md.

Run the C9W3 verification commands from the dispatch brief and report exact commands run and skipped.
```

## Task C9W4: Release Evidence Index Consumes Request Preflight Smoke Evidence

### Mission

Make `scripts/release_evidence_index.rb` classify existing desktop-trigger request preflight smoke evidence as a release evidence claim without running the smoke.

### Required behavior

- Add a `desktop-trigger-request-preflight-smoke` evidence claim.
- Accept existing JSON evidence only.
- Keep formal `v0.2.640` blocked unless the full checkpoint promotion packet allows the release.
- Do not convert old or missing preflight smoke evidence into a current full-smoke PASS claim.
- Keep Docker, QEMU, Wine, Colima, D-Bus, desktop launch, backend launch, network access, and host mutation disabled.

### Verification

```text
ruby -Ilib test/test_release_evidence_index.rb
ruby scripts/release_evidence_index.rb --format json
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable Claude Code Prompt

```text
Implement C9W4 from docs/claude-code-next-dispatch-brief.md.

Scope:
- Add request preflight smoke evidence classification to release_evidence_index.
- Read existing evidence only.
- Do not run full smoke, the targeted preflight smoke, Docker, QEMU, Wine, Colima, D-Bus, desktop launch, backend launch, network, or host mutation.
- Do not claim formal v0.2.640 readiness unless the full checkpoint promotion packet allows it.
- Do not modify docs/claude-code-implementation-packages.md.

Run the C9W4 verification commands from the dispatch brief and report exact commands run and skipped.
```

## Task C9W5: Runtime Owner Request Receipt Preview

### Mission

After formal promotion, add a Go-owned preview for the future owner-managed request receipt that will be required before a real desktop-triggered request can move from preflight into a controlled Runtime request object.

This task is a preview only. It must not write production state or dispatch anything.

### Required behavior

- Consume `desktop-trigger-request-preflight-preview`.
- Accept only the ready preflight state.
- Produce an opaque receipt plan, not a durable production receipt.
- Keep receipt persistence, request-object writes, permission grants, service dispatch, D-Bus calls, desktop launch, backend launch, Runtime production writes, KDE configuration writes, and host mutation disabled.
- Add ready, missing-promotion, blocked-preflight, malformed, redaction, replay-marker, and no-side-effect tests.

### Verification

```text
go test ./internal/runtime/owner ./cmd/xnix-runtime-go -run 'Test.*DesktopTrigger.*Request.*Receipt|Test.*DesktopTrigger.*Request.*Preflight' -count=1
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable Claude Code Prompt

```text
Implement C9W5 from docs/claude-code-next-dispatch-brief.md only if Codex confirms formal v0.2.640 promotion or explicitly approves pre-release modeling.

Scope:
- Add a Go-owned desktop-trigger owner request receipt preview.
- Consume the existing request preflight packet and fail closed unless the preflight is ready.
- Do not persist receipts, create request objects, dispatch service calls, call D-Bus, launch desktop actions, start backends, write Runtime production state, write KDE configuration, use network, use containers, or mutate the host.
- Do not modify docs/claude-code-implementation-packages.md.

Run the C9W5 verification commands from the dispatch brief and report exact commands run and skipped.
```

## Task C9W6: KDE Center Desktop-Trigger Preflight Card

### Mission

Expose the desktop-trigger request preflight state as a KDE-safe read model card so KDE can present readiness without owning policy or seeing owner-only arguments.

### Required behavior

- Add or extend the KDE Center page preview to consume redacted preflight state.
- Display only safe readiness state, blocker summary, and operator action hints.
- Do not expose owner service call arguments, CLI arguments, state roots, launcher paths, backend details, raw environment values, usernames, secrets, tokens, or host paths.
- Keep KDE presentation-only.
- Keep desktop launch, backend launch, D-Bus calls, Runtime writes, KDE configuration writes, network access, and host mutation disabled.

### Verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/owner ./cmd/xnix-runtime-go -run 'Test.*KDE.*Center|Test.*DesktopTrigger.*Preflight' -count=1
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable Claude Code Prompt

```text
Implement C9W6 from docs/claude-code-next-dispatch-brief.md.

Scope:
- Add a KDE-safe read model card for desktop-trigger request preflight state.
- KDE must remain presentation-only and must not receive owner service arguments, state roots, launcher paths, backend details, raw environment values, usernames, secrets, tokens, or host paths.
- Do not dispatch service calls, call D-Bus, launch desktop actions, start backends, write Runtime production state, write KDE configuration, use network, use containers, or mutate the host.
- Do not modify docs/claude-code-implementation-packages.md.

Run the C9W6 verification commands from the dispatch brief and report exact commands run and skipped.
```

## Intake Checklist for Returned Claude Work

When Claude Code returns a branch or patch, Codex should review it with this checklist:

1. Run `git status --short --branch`.
2. Confirm the task id and touched files match exactly one selected task.
3. Confirm `docs/claude-code-implementation-packages.md` is unchanged.
4. Confirm no Docker, QEMU, Wine, Colima, full smoke, network fetch, package-manager call, host service restart, privileged container, Docker socket mount, broad host mount, or host-root mutation was introduced.
5. Run the task-specific tests from this document.
6. Run `ruby scripts/verify_layout.rb`.
7. Run `ruby scripts/mainline_integration_review.rb --format json`.
8. Run `git diff --check`.
9. Review redaction tests for state roots, cache roots, launcher paths, backend details, host paths, secrets, tokens, credentials, usernames, and environment variables.
10. Stage only expected files by explicit path; never stage caches, logs, build outputs, local config, or environment files.
