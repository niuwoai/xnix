# Claude Code Next Work After v0.2.640-rc7

> Last updated: 2026-07-24 | Baseline: v0.2.640-rc7

This handoff is the next Claude Code dispatch sheet for the KDE-first Windows application compatibility mainline after the original `v0.2.640-rc1` checkpoint candidate, the local `v0.2.640-rc2` failure-classification update, the local `v0.2.640-rc3` desktop-trigger readiness packet, the local `v0.2.640-rc4` owner-service launch envelope guard, the local `v0.2.640-rc5` KDE action surface audit, the local `v0.2.640-rc6` managed launcher acceptance report, and the local `v0.2.640-rc7` desktop-trigger dry-run request review.

Use this document for bounded Claude Code implementation branches while Codex or the human operator resolves the external Docker Hub pull blocker. Do not treat this document as permission to run Docker, QEMU, network fetches, package-manager calls, privileged containers, host networking, Docker socket mounts, broad host-directory mounts, or host-root mutation.

`C8W1` is implemented locally in v0.2.640-rc2, `C8W2` is implemented locally in v0.2.640-rc3, `C8W3` is implemented locally in v0.2.640-rc4, `C8W4` is implemented locally in v0.2.640-rc5, `C8W5` is implemented locally in v0.2.640-rc6, and `C8W7` is implemented locally in v0.2.640-rc7. Use this document for `C8W6` unless a reviewer explicitly asks for a repair branch.

The immediate formal release gate is still:

```text
ruby scripts/full_smoke.rb
```

That gate is currently blocked before project build/test steps because Colima Docker cannot pull `debian:bookworm-slim` from Docker Hub and fails with EOF while resolving Docker Hub auth or manifest endpoints. Claude Code should not try to fix the operator's Colima or Docker Hub network state unless explicitly asked in a separate environment-support task.

Do not modify `docs/claude-code-implementation-packages.md` from any task in this handoff. That file remains protected.

## Current Mainline State

The `v0.2.640-rc7` candidate already promotes `kde-controlled-launch-action-dbus-fixture-smoke` into `scripts/full_smoke.rb` as a required PASS lane, adds structured full-smoke failure classification, adds a Go-owned desktop-trigger staged invocation readiness packet, adds a fail-closed owner-service launch envelope guard, audits the KDE controlled-launch action surface, adds a managed launcher acceptance report, and adds a desktop-trigger dry-run request review. The restricted container command runs the plan-consuming KDE controlled-launch action harness in the tested Runtime image with explicit D-Bus fixture execution enabled, no container networking, no privileged container, no Docker socket, managed cache volume access, and the controlled-launch scratch tmpfs.

The formal `v0.2.640` tag must wait until full smoke passes. If full smoke fails after the Docker pull blocker is resolved, fix the concrete project defect and rerun the complete checkpoint before promoting the version.

The next product direction after a passing full checkpoint is to move beyond fixture validation toward a real desktop-triggered staged `xnix-compat-launch` invocation while preserving KDE as presentation-only and Runtime as the owner of execution policy, receipts, state roots, launch planning, and backend dispatch.

## Global Rules for Every Task

Every Claude Code branch must:

- Pick exactly one task from `C8W6`, unless a reviewer explicitly requests a repair for `C8W1` through `C8W5`.
- Keep source, comments, fixtures, tests, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior.
- Use Ruby only for tests, smoke scripts, reports, and lightweight developer tooling.
- Keep KDE as shell and presentation only.
- Keep the Xnix Compatibility Runtime as the policy and product-logic owner.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for code changes.
- Add success, failure, malformed-input, blocked-state, redaction, and no-side-effect tests.
- Run the task-specific verification commands plus `ruby scripts/verify_layout.rb` and `git diff --check`.
- Stop and report if the task needs Docker, QEMU, real network access, host package managers, host-root mutation, production D-Bus ownership, real Portal transport, or real compatibility backend launch.

Unsafe behavior remains disabled:

- Production D-Bus ownership.
- Runtime write methods outside already explicit test-only or state-root-scoped preview writers.
- Real Wine, Proton, VM, or compatibility backend launch from KDE.
- Real XDG Portal transport calls.
- Default network artifact fetch.
- Host package-manager calls.
- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Host-root mutation.
- Raw backend commands in KDE-facing output.
- Raw executable paths, Runtime state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, or environment variables in user-facing output.

## Dispatch Order

| Order | Task | Primary lane | Why now |
| --- | --- | --- | --- |
| Done | `C8W1` Full-smoke blocker classifier | full checkpoint / release evidence | Implemented locally in v0.2.640-rc2. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W2` Desktop-trigger staged invocation readiness packet | KDE action / Runtime launch | Implemented locally in v0.2.640-rc3. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W3` Runtime owner service launch request envelope guard | Runtime owner service | Implemented locally in v0.2.640-rc4. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W4` KDE controlled-launch action surface audit | KDE shell presentation | Implemented locally in v0.2.640-rc5. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W5` Managed launcher acceptance report | known Windows app / Compatibility Center | Implemented locally in v0.2.640-rc6. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W7` Desktop-trigger dry-run request review | desktop-trigger gate | Implemented locally in v0.2.640-rc7. Do not dispatch again unless a reviewer asks for repair. |
| 1 | `C8W6` Post-checkpoint promotion checklist | release train | Give Codex a deterministic checklist for promoting `0.2.640-rc7` to `0.2.640` only after full smoke passes. |

Prefer dispatching `C8W6` next. It improves release promotion safety without relying on Docker, QEMU, or the current external registry problem.

## C8W1: Full-Smoke Blocker Classifier

### Mission

Create a structured classifier for `scripts/full_smoke.rb` failures so release evidence can clearly distinguish infrastructure blockers, missing prerequisites, project build failures, QEMU boot failures, guest Wine failures, known-app smoke failures, and KDE action fixture failures.

### Start from

- `scripts/full_smoke.rb`
- `lib/xnix/full_smoke_report.rb`
- `lib/xnix/container.rb`
- `test/test_full_smoke_report.rb`
- `test/test_full_smoke_script.rb`
- `docs/xnix-current-mainline.md`

### Deliver

- A Ruby-owned failure classification model consumed by the full-smoke report.
- JSON and Markdown report fields for `failure_class`, `failed_step`, `operator_action_required`, `project_defect_possible`, and `safe_retry_command`.
- Classification fixtures for Docker daemon unavailable, Docker Hub EOF, missing local base image, Buildroot build failure, QEMU serial timeout, known Windows app smoke failure, fixture Windows app smoke failure, and KDE action fixture failure.
- Tests proving the classifier does not hide the original error text and does not report a formal release-ready state when a step failed.

### Keep disabled

- Automatic Docker pulls.
- Automatic Colima restarts.
- Automatic retries.
- Network probing.
- QEMU execution.
- Full-smoke execution from tests.
- Host configuration changes.

### Required verification

```text
ruby -Ilib test/test_full_smoke_report.rb
ruby -Ilib test/test_full_smoke_script.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement C8W1 from docs/claude-code-next-work-after-0640-rc1.md.

Target outcome:
- Add structured full-smoke failure classification so release evidence separates infrastructure blockers from project defects.

Hard constraints:
- Do not run Docker, QEMU, network probes, package managers, Colima restart commands, or full smoke.
- Do not automatically pull images, restart services, retry failed steps, or mutate host configuration.
- Do not expose secrets, tokens, credentials, usernames, environment variables, host paths, raw state-root paths, or raw backend commands in user-facing report fields.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add classifier tests for Docker daemon unavailable, Docker Hub EOF, missing base image, Buildroot build failure, QEMU timeout, known-app failure, fixture-app failure, and KDE action fixture failure.
- Run the C8W1 verification commands and report exact commands run.
```

## C8W2: Desktop-Trigger Staged Invocation Readiness Packet

### Mission

Create a Go-owned readiness packet that consumes the current KDE controlled-launch action preview, D-Bus fixture smoke plan, Runtime-status launch evidence handoff, owner service trigger preview, managed launcher bridge, and known Windows app guest smoke evidence. The packet should explain whether the next real desktop-triggered staged `xnix-compat-launch` smoke can be attempted after full checkpoint promotion.

### Start from

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
- Deterministic states: ready, blocked, missing-evidence, stale-evidence, unsafe, unsupported, or needs-full-checkpoint.
- A user-safe summary that names the next human-authorized smoke without exposing owner-service arguments, host paths, state-root paths, raw `.exe` paths, Wine/QEMU commands, backend details, or credentials.
- Tests for ready fixture evidence, missing action evidence, stale digest evidence, missing known-app cache evidence, blocked full checkpoint, malformed evidence, and no-side-effect behavior.

### Keep disabled

- Desktop launch.
- Backend launch.
- QEMU execution.
- Wine execution.
- D-Bus calls.
- Runtime state writes.
- KDE configuration writes.
- Receipt reconstruction from KDE.
- Network fetch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/winapp ./cmd/xnix-runtime-go
ruby -Ilib test/test_kde_controlled_launch_action_smoke_script.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement C8W2 from docs/claude-code-next-work-after-0640-rc1.md.

Target outcome:
- Add a Runtime-owned desktop-trigger staged invocation readiness packet for the next real desktop-triggered xnix-compat-launch smoke.

Hard constraints:
- Do not launch desktop actions, start backends, run QEMU, run Wine, call D-Bus, write Runtime state, write KDE configuration, reconstruct receipts in KDE, fetch artifacts, or mutate the host root.
- Do not expose owner-service arguments, host paths, Runtime state-root paths, raw .exe paths, Wine/QEMU commands, backend details, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add ready, missing-evidence, stale-digest, missing-cache, blocked-full-checkpoint, malformed-evidence, redaction, and no-side-effect tests.
- Run the C8W2 verification commands and report exact commands run.
```

## C8W3: Runtime Owner Service Launch Request Envelope Guard

### Mission

Add a fail-closed owner-service guard that validates desktop-triggered launch request envelopes before any future service method can consume them. The guard should prove that KDE can only forward opaque evidence handles and cannot replay, expand, or reconstruct Runtime-owned launch inputs.

### Start from

- `cmd/xnix-runtime-go/known_app_runtime_status_launch_owner_trigger_commands.go`
- `cmd/xnix-runtime-go/kde_controlled_launch_action_commands.go`
- `internal/runtime/owner`
- `internal/runtime/appidentity`
- `scripts/dbus_controlled_launch_owner_fixture_smoke.rb`
- `scripts/staged_launcher_dispatch_smoke.rb`

### Deliver

- A Go read model and CLI command such as `owner-service-launch-envelope-guard-preview`.
- Guard checks for route id, method id, evidence handle shape, digest binding, expected action id, freshness window, replay marker, caller role, and disabled owner-only arguments.
- Output states: accepted-for-review, blocked-missing-evidence, blocked-mismatched-route, blocked-stale, blocked-replay, blocked-owner-args, malformed, or unsupported.
- Tests proving KDE-facing input cannot include state root, cache root, launcher path, timeout values, raw executable paths, backend commands, receipt fields, session ids, or dispatch ids.
- Tests proving the guard never writes receipts, accepts production authorization, dispatches service calls, enables D-Bus ownership, or launches a backend.

### Keep disabled

- Receipt writes.
- Receipt acceptance.
- Production authorization.
- Service-call dispatch.
- D-Bus ownership.
- Backend launch.
- Desktop launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/owner ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_dbus_controlled_launch_owner_fixture_smoke_script.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement C8W3 from docs/claude-code-next-work-after-0640-rc1.md.

Target outcome:
- Add a fail-closed Runtime owner service launch request envelope guard for future desktop-triggered launch calls.

Hard constraints:
- Do not write receipts, accept production authorization, dispatch service calls, enable D-Bus ownership, launch desktop actions, start backends, run Docker/QEMU, or mutate the host root.
- Do not allow KDE-facing input to include state roots, cache roots, launcher paths, timeout values, raw executable paths, backend commands, receipt fields, session ids, dispatch ids, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add accepted-for-review, missing-evidence, mismatched-route, stale, replay, owner-args, malformed-input, redaction, and no-side-effect tests.
- Run the C8W3 verification commands and report exact commands run.
```

## C8W4: KDE Controlled-Launch Action Surface Audit

### Mission

Create an offline audit that proves KDE controlled-launch action files, action ids, D-Bus method declarations, labels, icons, and Compatibility Center navigation metadata match Runtime-owned preview output and remain presentation-only.

### Start from

- `kde/actions/xnix-runtime-status-controlled-launch.desktop`
- `kde/plasmoids`
- `cmd/xnix-runtime-go/kde_controlled_launch_action_commands.go`
- `internal/runtime/appidentity/kde_controlled_launch_action.go`
- `scripts/kde_controlled_launch_action_smoke.rb`
- `scripts/verify_layout.rb`

### Deliver

- A CLI command such as `kde-controlled-launch-action-surface-audit-preview`.
- Audit rows for desktop action file metadata, action id, D-Bus route, public method, label, icon, Compatibility Center target, evidence-handle forwarding, disabled receipt reconstruction, disabled KDE state-root access, and disabled execution start.
- Tests proving drifted action files, unknown methods, raw owner-service args, raw paths, backend terms, receipt/session reconstruction fields, and missing Center navigation are blocked.
- Layout guard coverage so the controlled-launch action cannot drift away from Runtime-owned metadata.

### Keep disabled

- KDE file writes.
- D-Bus calls.
- Runtime state writes.
- Receipt reconstruction.
- Execution start.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_kde_controlled_launch_action_smoke_script.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement C8W4 from docs/claude-code-next-work-after-0640-rc1.md.

Target outcome:
- Add an offline KDE controlled-launch action surface audit that proves KDE metadata matches Runtime-owned preview output and remains presentation-only.

Hard constraints:
- Do not write KDE files, call D-Bus, write Runtime state, reconstruct receipts, start execution, launch backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw owner-service args, raw paths, state-root paths, backend commands, backend details, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add matching, drifted-file, unknown-method, raw-owner-args, raw-path, backend-term, receipt-reconstruction, missing-navigation, and no-side-effect tests.
- Run the C8W4 verification commands and report exact commands run.
```

## C8W5: Managed Launcher Acceptance Report

### Mission

Create a user-safe acceptance report for the managed 7-Zip launcher path that joins known-app catalog metadata, artifact verification, guest Wine smoke evidence, managed launcher bridge evidence, Runtime-status handoff evidence, and Compatibility Center projection evidence.

### Start from

- `internal/runtime/winapp`
- `internal/runtime/appidentity`
- `cmd/xnix-runtime-go/windows_compatibility_commands.go`
- `cmd/xnix-compat-launch`
- `scripts/known_winapp_guest_wine_smoke.rb`
- `scripts/staged_launcher_dispatch_smoke.rb`
- `scripts/compatibility_product_smoke.rb`

### Deliver

- A CLI command such as `managed-launcher-acceptance-report-preview`.
- Report sections for known app identity, artifact digest evidence, cache readiness, guest smoke evidence, launcher bridge readiness, Runtime launch gate, desktop trigger handoff, Compatibility Center projection, and remaining blockers.
- JSON and Markdown output modes if an existing report pattern is available.
- Tests for passed fixture evidence, missing artifact, digest mismatch, guest smoke missing, launch gate blocked, desktop trigger missing, malformed evidence, and redaction.

### Keep disabled

- Artifact download.
- Wine execution.
- QEMU execution.
- Desktop launch.
- Backend launch.
- Runtime state writes.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/winapp ./internal/runtime/appidentity ./cmd/xnix-runtime-go ./cmd/xnix-compat-launch
ruby -Ilib test/test_known_winapp_guest_wine_smoke_script.rb
ruby -Ilib test/test_compatibility_product_smoke.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement C8W5 from docs/claude-code-next-work-after-0640-rc1.md.

Target outcome:
- Add a managed launcher acceptance report that joins 7-Zip known-app evidence, guest smoke evidence, managed launcher bridge evidence, Runtime-status handoff evidence, and Compatibility Center projection evidence.

Hard constraints:
- Do not download artifacts, run Wine, run QEMU, launch desktop actions, start backends, write Runtime state, run Docker, or mutate the host root.
- Do not expose raw .exe paths, host paths, Runtime state-root paths, Wine/QEMU commands, backend details, secrets, tokens, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add passed, missing-artifact, digest-mismatch, missing-smoke, blocked-gate, missing-trigger, malformed-evidence, redaction, and no-side-effect tests.
- Run the C8W5 verification commands and report exact commands run.
```

## C8W6: Post-Checkpoint Promotion Checklist

### Mission

Create a deterministic promotion checklist that Codex can use after `ruby scripts/full_smoke.rb` passes to promote `0.2.640-rc7` to formal `0.2.640` without accidentally skipping evidence, retaining stale release-blocked wording, or claiming success from a partial smoke.

### Start from

- `VERSION`
- `CHANGELOG.md`
- `PRODUCT_OVERVIEW.md`
- `docs/xnix-current-mainline.md`
- `docs/mainline-integration-checkpoint.md`
- `lib/xnix/full_smoke_report.rb`
- `scripts/full_smoke.rb`
- `scripts/verify_layout.rb`

### Deliver

- A documentation-only checklist or a Ruby dry-run checklist command such as `scripts/promote_full_checkpoint_checklist.rb`.
- Required evidence rows for full-smoke JSON, full-smoke Markdown, serial log, builder tools image, Runtime image, Buildroot image, SSH Wine guest image, known-app cache, base QEMU boot smoke, known-app QEMU Wine smoke, fixture QEMU Wine smoke, and KDE controlled-launch action D-Bus fixture smoke.
- Version-promotion steps that replace `0.2.640-rc7` with `0.2.640` only after all required PASS markers exist.
- Tests if code is added; otherwise include a clear manual checklist and no-code verification commands.

### Keep disabled

- Automatic version promotion.
- Automatic commits or tags.
- Full-smoke execution from tests.
- Docker or QEMU execution from the checklist.
- Host-root mutation.

### Required verification

For documentation-only work:

```text
ruby scripts/verify_layout.rb
git diff --check
```

For a Ruby checklist command:

```text
ruby -Ilib test/test_promote_full_checkpoint_checklist.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement C8W6 from docs/claude-code-next-work-after-0640-rc1.md.

Target outcome:
- Add a deterministic post-checkpoint promotion checklist for turning 0.2.640-rc7 into 0.2.640 only after full smoke passes.

Hard constraints:
- Do not automatically promote the version, commit, tag, run full smoke, run Docker, run QEMU, or mutate the host root.
- Do not claim a formal release-ready state from partial smoke evidence or stale release-blocked wording.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md only if code changes require it.
- Add tests if a Ruby checklist command is created.
- Run the C8W6 verification commands and report exact commands run.
```

## Intake Checklist for Returned Claude Branches

When Claude Code returns a branch:

1. Run `git status --short --branch`.
2. Confirm `docs/claude-code-implementation-packages.md` has no diff.
3. Run `ruby scripts/mainline_integration_review.rb --format json` if the branch touches production evidence, release evidence, Runtime owner routes, KDE integration, or any `docs/claude-code-*` dispatch sheet.
4. Confirm `protected_claude_file_modified` is `false`.
5. Confirm `unclassified_file_count` is `0`, or classify every new file before staging.
6. Run the task-specific verification commands from this document.
7. Run `ruby scripts/verify_layout.rb`.
8. Run `git diff --check`.
9. Exclude `.cache/`, `.gocache/`, `tmp/`, logs, build artifacts, local config, secrets, tokens, private keys, and generated output from staging.
10. Stage by lane, never with `git add .`.
11. Commit one coherent semantic-version change per branch.

## Work Not Suitable for Claude Code From This Sheet

Do not dispatch these as Claude Code tasks from this document:

- Fix Colima networking.
- Pull `debian:bookworm-slim`.
- Restart Colima.
- Run `ruby scripts/full_smoke.rb`.
- Promote `0.2.640-rc7` to `0.2.640` without full-smoke PASS evidence.
- Enable production D-Bus ownership.
- Enable Runtime write methods for KDE.
- Start Wine, Proton, QEMU, or VM backends from KDE.
- Call real XDG Portals.
- Download new Windows application artifacts.
- Modify host package-manager state.
- Mount the Docker socket.
- Use privileged containers.
- Use host networking.
- Mutate host-root paths.
