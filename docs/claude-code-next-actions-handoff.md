# Claude Code Next Actions Handoff

> Last updated: 2026-07-24 | Baseline: v0.2.640-rc5

This document is the short, copyable handoff for asking Claude Code to continue Xnix mainline work without drifting into unsafe or overly broad implementation.

Use this document when the human operator wants Claude Code to implement the next bounded branch while Codex keeps the mainline review, merge, release gate, and host-safety decisions.

## Current Situation

Xnix is currently at the local `v0.2.640-rc5` checkpoint candidate.

The formal `v0.2.640` release is not promoted yet because the full-smoke gate still needs a passing run:

```text
ruby scripts/full_smoke.rb
```

The latest known blocker is external to the project code: Colima Docker cannot pull `debian:bookworm-slim` from Docker Hub and fails with EOF while resolving Docker Hub authentication or manifest endpoints.

Claude Code should not repair Colima, Docker Hub access, Docker daemon state, host package managers, or network configuration from this handoff. Those are operator-owned environment tasks.

## Non-Negotiable Boundaries

Every Claude Code task from this handoff must:

- Implement exactly one task.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior.
- Use Ruby only for tests, smoke scripts, reports, and lightweight developer tooling.
- Keep KDE as shell and presentation only.
- Keep the Xnix Compatibility Runtime as the owner of policy, evidence, receipts, state roots, launch planning, and backend dispatch.
- Avoid editing `docs/claude-code-implementation-packages.md`.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` if code changes.
- Add targeted tests for success, blocked input, malformed input, redaction, and no-side-effect behavior.
- Run task-specific tests, `ruby scripts/verify_layout.rb`, and `git diff --check`.
- Report exact commands run and exact commands skipped.

Every Claude Code task from this handoff must not:

- Run Docker, QEMU, Wine, Proton, or a real compatibility backend.
- Pull images or fetch packages from the network.
- Restart Colima, Docker, or host services.
- Use host package-manager commands.
- Use privileged containers, host networking, Docker socket mounts, or broad host-directory mounts.
- Mutate the host root.
- Enable production D-Bus ownership.
- Add real Portal transport calls.
- Write Runtime production state unless the selected task explicitly asks for a state-root-scoped preview writer.
- Write KDE configuration.
- Reconstruct Runtime receipts, sessions, state roots, cache roots, launcher paths, executable paths, backend commands, or owner-only arguments in KDE-facing code.
- Expose secrets, tokens, credentials, usernames, environment variables, host paths, state-root paths, raw executable paths, raw launcher output, backend names, backend commands, or backend details in user-facing output.

## Dispatch Order

Use this order unless a reviewer explicitly asks for a repair branch:

| Order | Task | Suggested branch | Outcome |
| --- | --- | --- | --- |
| Done | `C8W2` Desktop-trigger staged invocation readiness packet | `codex/desktop-trigger-staged-invocation-readiness` | Present locally in v0.2.640-rc3. Do not dispatch again unless a reviewer requests repair. |
| Done | `C8W3` Runtime owner service launch request envelope guard | `codex/owner-service-launch-envelope-guard` | Present locally in v0.2.640-rc4. Do not dispatch again unless a reviewer requests repair. |
| Done | `C8W4` KDE controlled-launch action surface audit | `codex/kde-controlled-launch-action-surface-audit` | Present locally in v0.2.640-rc5. Do not dispatch again unless a reviewer requests repair. |
| 1 | `C8W5` Managed launcher acceptance report | `codex/managed-launcher-acceptance-report` | Tie known-app smoke evidence, launcher bridge evidence, and Compatibility Center projection into a user-safe acceptance report. |
| 2 | `C8W6` Post-checkpoint promotion checklist | `codex/post-checkpoint-promotion-checklist` | Produce a deterministic promotion checklist for `v0.2.640` after full smoke passes. |

Prefer `C8W5` next. It connects the managed launcher evidence chain without needing Docker, QEMU, network access, or host changes.

## Task C8W2: Desktop-Trigger Staged Invocation Readiness Packet

### Mission

Add a Go-owned read-only readiness packet that consumes the existing KDE controlled-launch action preview, D-Bus fixture smoke plan, Runtime-status launch evidence handoff, owner service trigger preview, managed launcher bridge metadata, and known-app guest-smoke evidence.

The packet should answer one product question:

```text
Can a human operator attempt the next desktop-triggered staged launch smoke after the full checkpoint is promoted?
```

### Start From

- `cmd/xnix-runtime-go/kde_controlled_launch_action_commands.go`
- `cmd/xnix-runtime-go/kde_controlled_launch_session_bus_smoke_plan_commands.go`
- `cmd/xnix-runtime-go/known_app_runtime_status_launch_owner_trigger_commands.go`
- `cmd/xnix-runtime-go/windows_compatibility_commands.go`
- `internal/runtime/appidentity`
- `internal/runtime/winapp`
- `scripts/kde_controlled_launch_action_smoke.rb`
- `scripts/staged_launcher_dispatch_smoke.rb`

### Deliver

- A CLI command such as `desktop-trigger-staged-invocation-readiness-preview`.
- Readiness sections for KDE action metadata, public D-Bus route, owner-service trigger, Runtime-status evidence handoff, managed launcher request, known-app artifact verification, guest smoke boundary, and release-gate dependency.
- Deterministic section states: `ready`, `blocked`, `missing-evidence`, `stale-evidence`, `unsafe`, `unsupported`, or `needs-full-checkpoint`.
- A user-safe summary naming the next human-authorized smoke without exposing owner-service arguments, host paths, state-root paths, raw executable paths, backend commands, backend names, backend details, or credentials.
- Tests for ready fixture evidence, missing action evidence, stale digest evidence, missing known-app cache evidence, blocked full checkpoint, malformed evidence, redaction, and no-side-effect behavior.

### Required Verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/winapp ./cmd/xnix-runtime-go
ruby -Ilib test/test_kde_controlled_launch_action_smoke_script.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Prompt

```text
Implement C8W2 from docs/claude-code-next-actions-handoff.md.

Target outcome:
- Add a Runtime-owned desktop-trigger staged invocation readiness packet for the next real desktop-triggered staged launch smoke.

Hard constraints:
- Do not launch desktop actions, start compatibility backends, run Docker, run QEMU, run Wine, call D-Bus, write Runtime state, write KDE configuration, reconstruct receipts in KDE, fetch artifacts, or mutate the host root.
- Do not expose owner-service arguments, host paths, Runtime state-root paths, raw executable paths, backend commands, backend names, backend details, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add ready, missing-evidence, stale-digest, missing-cache, blocked-full-checkpoint, malformed-evidence, redaction, and no-side-effect tests.
- Run the C8W2 verification commands and report exact commands run and skipped.
```

## Task C8W3: Runtime Owner Service Launch Request Envelope Guard

### Mission

Add a fail-closed owner-service guard that validates desktop-triggered launch request envelopes before any future service method can consume them.

The guard must prove that KDE can only forward opaque evidence handles. KDE must not replay, expand, or reconstruct Runtime-owned launch inputs.

### Start From

- `cmd/xnix-runtime-go/known_app_runtime_status_launch_owner_trigger_commands.go`
- `cmd/xnix-runtime-go/kde_controlled_launch_action_commands.go`
- `internal/runtime/owner`
- `internal/runtime/appidentity`
- `scripts/dbus_controlled_launch_owner_fixture_smoke.rb`
- `scripts/staged_launcher_dispatch_smoke.rb`

### Deliver

- A Go read model and CLI command such as `owner-service-launch-envelope-guard-preview`.
- Guard checks for route id, method id, evidence handle shape, digest binding, expected action id, freshness window, replay marker, caller role, and disabled owner-only arguments.
- Output states: `accepted-for-review`, `blocked-missing-evidence`, `blocked-mismatched-route`, `blocked-stale`, `blocked-replay`, `blocked-owner-args`, `malformed`, or `unsupported`.
- Tests proving KDE-facing input cannot include state root, cache root, launcher path, timeout values, raw executable paths, backend commands, receipt fields, session ids, or dispatch ids.
- Tests proving the guard never writes receipts, accepts production authorization, dispatches service calls, enables D-Bus ownership, or launches a backend.

### Required Verification

```text
go test ./internal/runtime/owner ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_dbus_controlled_launch_owner_fixture_smoke_script.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Prompt

```text
Implement C8W3 from docs/claude-code-next-actions-handoff.md.

Target outcome:
- Add a fail-closed Runtime owner service launch request envelope guard for future desktop-triggered launch calls.

Hard constraints:
- Do not write receipts, accept production authorization, dispatch service calls, enable D-Bus ownership, launch desktop actions, start compatibility backends, run Docker, run QEMU, run Wine, or mutate the host root.
- Do not allow KDE-facing input to include state roots, cache roots, launcher paths, timeout values, raw executable paths, backend commands, receipt fields, session ids, dispatch ids, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add accepted, missing-evidence, mismatched-route, stale, replay, owner-args, malformed, redaction, and no-side-effect tests.
- Run the C8W3 verification commands and report exact commands run and skipped.
```

## Task C8W4: KDE Controlled-Launch Action Surface Audit

### Mission

Add a KDE-facing audit preview that checks whether the controlled-launch action remains presentation-only before a real desktop-triggered staged launch smoke is attempted.

### Start From

- `kde/actions/xnix-runtime-status-controlled-launch.desktop`
- `cmd/xnix-runtime-go/kde_controlled_launch_action_commands.go`
- `internal/runtime/appidentity`
- `scripts/kde_controlled_launch_action_smoke.rb`

### Deliver

- A CLI command such as `kde-controlled-launch-action-surface-audit-preview`.
- Checks for desktop action id, label, public method, evidence-only argument shape, no owner-service arguments, no state-root access, no raw executable path, no backend details, and no execution enablement.
- A redacted KDE-safe report that says whether the action surface is safe for the next human-authorized smoke.
- Tests for safe metadata, unsafe owner arguments, unsafe backend terms, malformed desktop metadata, missing evidence handle, and no-side-effect behavior.

### Required Verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_kde_controlled_launch_action_smoke_script.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Prompt

```text
Implement C8W4 from docs/claude-code-next-actions-handoff.md.

Target outcome:
- Add a KDE controlled-launch action surface audit preview that proves the KDE action remains presentation-only.

Hard constraints:
- Do not enable launch, call D-Bus, write KDE configuration, write Runtime state, start backends, run Docker, run QEMU, run Wine, or mutate the host root.
- Do not expose owner-service arguments, state roots, cache roots, launcher paths, raw executable paths, backend commands, backend names, backend details, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add safe, unsafe-owner-args, unsafe-backend-terms, malformed-desktop-metadata, missing-evidence-handle, redaction, and no-side-effect tests.
- Run the C8W4 verification commands and report exact commands run and skipped.
```

## Merge Intake Checklist

When Claude Code returns a branch, Codex should review it before merging:

1. Run `git status --short --branch`.
2. Run `git diff --name-only`.
3. Confirm `docs/claude-code-implementation-packages.md` is unchanged.
4. Run `ruby scripts/mainline_integration_review.rb --format json`.
5. Confirm `protected_claude_file_modified` is `false`.
6. Confirm `unclassified_file_count` is `0`.
7. Run the task-specific verification commands from this document.
8. Run `ruby scripts/verify_layout.rb`.
9. Run `git diff --check`.
10. Inspect version metadata, changelog entry, and product overview updates if code changed.
11. Stage only expected files by path. Do not use `git add .`.
12. Commit one coherent small version.
13. Tag only if the branch is being promoted as an accepted local checkpoint.

## Do Not Ask Claude Code To Do These From This Handoff

- Finish Windows compatibility end-to-end.
- Make KDE integration real in one branch.
- Enable production D-Bus ownership.
- Enable real Wine, Proton, VM, or compatibility backend launch.
- Install Wine, Proton, KDE packages, Docker images, or system packages.
- Run full smoke, QEMU acceptance, or Docker-based heavy smoke.
- Repair Colima, Docker Desktop, Docker Hub access, or host networking.
- Call a real Portal.
- Add production cleanup, deletion, revocation, renewal, or notification dispatch.
- Download packages or artifacts.
- Edit `docs/claude-code-implementation-packages.md`.
