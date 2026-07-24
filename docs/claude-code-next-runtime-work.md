# Claude Code Next Runtime Work

> Last updated: 2026-07-24 | Baseline: v0.2.635

This document is an implementation handoff for Claude Code. It is intentionally separate from `docs/claude-code-implementation-packages.md`; do not edit that file while working from this handoff.

## Product Direction

Xnix is now optimizing for the best Linux experience for existing Windows applications. KDE Plasma is the first mature desktop shell, but it must remain replaceable. The independent Xnix AI Compatibility Runtime owns compatibility decisions, Wine or VM backend policy, recipes, permissions, snapshots, rollback, diagnostics, and execution gates.

Implementation ownership rules:

- Use Go for important Runtime product logic and stable service boundaries.
- Keep existing C where it already owns low-level adapter surfaces.
- Use Ruby for tests, reports, smoke orchestration, and low-frequency development tooling.
- Do not move product policy into KDE plugins, Ruby scripts, or the C smoke adapter.
- Do not expose raw backend commands, state-root paths, broad host paths, launcher paths, raw launcher output, or compatibility internals to normal desktop-facing output.
- Do not use privileged containers, host networking, Docker socket mounts, broad host-directory mounts, or host-root mutation in normal development checks.

## Current Baseline

The current baseline is `v0.2.635`.

Already done:

- The real staged launcher lane can launch the managed known Windows app through `xnix-runtime-owner --service-call ShowRuntimeControlledLaunch evidence-relative-path <relative>` when QEMU/Wine prerequisites are present.
- `known-app-runtime-status-launch-owner-trigger-preview` reads verified Runtime-status launch evidence and emits desktop/owner-service trigger metadata.
- `staged_launcher_dispatch_smoke.rb` consumes `owner_service_cli_args` from the Go trigger preview.
- `dbus_controlled_launch_owner_fixture_smoke.rb` consumes the same Go trigger preview before calling `org.xnix.Compatibility1.ShowRuntimeControlledLaunch`.

Remaining near-term gap:

- The C D-Bus smoke adapter now consumes Go owner-trigger metadata for the controlled-launch dispatch, but the next KDE-facing action entrypoint still needs to consume the same trigger shape.
- KDE-facing action wiring is still smoke-level; it has not become a real KDE entrypoint for launching the managed Windows app.
- The current smoke path proves the trigger shape, but the next step should remove more adapter-owned launch knowledge and move toward a real desktop-triggered staged `xnix-compat-launch` path.

## Work Package A: Go-Owned D-Bus Dispatch Payload

Goal: make the D-Bus controlled-launch adapter consume Go-provided trigger metadata instead of hardcoding the owner-service dispatch shape in C where practical.

Baseline status: the v0.2.635 implementation already moved the controlled-launch C smoke adapter onto `known-app-runtime-status-launch-owner-trigger-preview` for owner-service method, handoff kind, route, and call type. Treat this package as a hardening task, not as fresh work.

Suggested approach:

1. Add or extend a Go command that can render a D-Bus adapter dispatch payload for `ShowRuntimeControlledLaunch` from a verified Runtime-status launch evidence handle.
2. Keep the command evidence-only at the desktop boundary:
   - input may include `--evidence-id` or `--evidence-relative-path`;
   - state root and owner-only inputs must come from Runtime owner configuration or a controlled test fixture boundary;
   - output must include only desktop-safe trigger and dispatch metadata.
3. Update the C smoke adapter only as a thin caller/renderer of Go-owned dispatch metadata where possible.
4. Preserve the existing public D-Bus method:
   - `org.xnix.Compatibility1.ShowRuntimeControlledLaunch`
   - argument: safe relative evidence path
   - result: desktop-safe controlled-launch action evidence

Acceptance criteria:

- C no longer owns method/evidence dispatch policy beyond accepting the public D-Bus call and delegating to Go-owned metadata.
- Go remains the source of truth for:
  - Runtime method name;
  - owner-service call type;
  - evidence handoff semantics;
  - disabled unsafe gates;
  - KDE evidence-only behavior.
- `dbus_controlled_launch_owner_fixture_smoke.rb` still passes or cleanly skips.
- The D-Bus response still includes `runtime-controlled-launch-dbus-action`, `desktop-action-dispatch`, `kde-dbus-runtime-status-action`, and nested `go_owner_service_call_json`.

Suggested files:

- `internal/runtime/owner/`
- `cmd/xnix-runtime-owner/`
- `cmd/xnix-runtime-go/`
- `runtime/dbus/xnix_compatd_smoke.c`
- `scripts/dbus_controlled_launch_owner_fixture_smoke.rb`
- `test/test_dbus_controlled_launch_owner_fixture_smoke_script.rb`
- `scripts/verify_layout.rb`

Targeted tests:

```text
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/owner -run 'TestServiceCallDispatchesShowRuntimeControlledLaunch|TestServiceCallServesReadDispatchInProcess' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./cmd/xnix-runtime-owner -run 'TestRuntimeOwnerCommandRendersShowRuntimeControlledLaunchServiceCall' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./cmd/xnix-runtime-go -run 'TestKnownAppRuntimeStatusLaunchOwnerTriggerPreviewCommandConsumesVerifiedHandoff|TestKnownAppRuntimeStatusLaunchOwnerTriggerPreviewCommandRejectsMissingStateRoot' -count=1
ruby -Ilib test/test_dbus_controlled_launch_owner_fixture_smoke_script.rb
ruby scripts/verify_layout.rb
git diff --check
```

## Work Package B: KDE-Facing Controlled Launch Action Stub

Goal: add a KDE-facing action stub that consumes the Go-owned trigger shape without implementing policy in KDE.

Suggested approach:

1. Add a small KDE-side script or metadata entrypoint that represents the Runtime-status controlled launch action.
2. The KDE-facing artifact must only forward the evidence handle to the D-Bus method.
3. It must not include state-root paths, launcher paths, backend names, Wine commands, QEMU commands, or reconstructed receipt/session fields.
4. Add static and smoke-level tests proving the KDE artifact is presentation-only.

Acceptance criteria:

- The KDE-side artifact exposes a user-facing controlled launch action label or action id.
- The only executable launch handoff is the safe evidence handle plus `org.xnix.Compatibility1.ShowRuntimeControlledLaunch`.
- Tests prove KDE does not own compatibility decisions, backend selection, owner-service arguments, or Runtime state paths.

Suggested files:

- `kde/`
- `scripts/`
- `test/`
- `scripts/verify_layout.rb`
- `docs/xnix-current-mainline.md`

Targeted tests:

```text
ruby -Ilib test/test_kde_center_dbus_smoke_script.rb
ruby -Ilib test/test_dbus_controlled_launch_owner_fixture_smoke_script.rb
ruby scripts/verify_layout.rb
git diff --check
```

## Work Package C: Real Managed Windows App Evidence Tightening

Goal: improve confidence that the current 7zr managed Windows app path is the concrete product lane, not just a contract projection.

Suggested approach:

1. Find the smallest place where the managed 7zr evidence can be made more concrete without running a full smoke every version.
2. Prefer Go-owned evidence normalization and verification.
3. Keep Ruby responsible for smoke orchestration only.
4. Add targeted checks that prove the evidence came from the real managed launcher lane when prerequisites are present.

Acceptance criteria:

- Evidence explicitly connects:
  - known Windows app identity;
  - managed artifact verification;
  - Runtime-owned launch authorization;
  - session-gated review;
  - controlled execution session;
  - staged `xnix-compat-launch`;
  - owner-service `ShowRuntimeControlledLaunch`.
- Desktop-facing JSON remains redacted and safe.
- The smoke can still SKIP cleanly when QEMU/Wine artifacts are unavailable.

Suggested files:

- `internal/runtime/appidentity/`
- `cmd/xnix-runtime-go/`
- `scripts/staged_launcher_dispatch_smoke.rb`
- `test/test_staged_launcher_dispatch_smoke_script.rb`
- `scripts/verify_layout.rb`

Targeted tests:

```text
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/appidentity -run 'TestPreviewKnownAppRuntimeStatusLaunchOwnerTriggerConsumesVerifiedHandoff|TestPrepareKnownAppKDERuntimeStatusLaunchExecutionFromActionTriggerConsumesHandoff' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./cmd/xnix-runtime-go -run 'TestKnownAppRuntimeStatusLaunchOwnerTriggerPreviewCommandConsumesVerifiedHandoff|TestKnownAppKDERuntimeStatusLaunchExecutionCommandConsumesActionTriggerHandoff' -count=1
ruby -Ilib test/test_staged_launcher_dispatch_smoke_script.rb
ruby scripts/verify_layout.rb
git diff --check
```

## Preferred Order

1. Work Package B first. It starts turning the D-Bus lane into a KDE-facing action without forking KDE or moving policy into the shell.
2. Work Package C second, or in parallel only if it does not touch the same files. It tightens real managed app evidence while preserving the targeted-test cadence.
3. Work Package A only for hardening or cleanup if the C adapter still owns avoidable dispatch details after v0.2.635.

## Versioning and Documentation Rules

For every implementation change:

1. Bump `VERSION` and all existing version mirrors.
2. Update `CHANGELOG.md`.
3. Update `PRODUCT_OVERVIEW.md`.
4. Update `docs/xnix-current-mainline.md` when the mainline changes.
5. Keep project-facing text in English.
6. Do not edit `docs/claude-code-implementation-packages.md`.
7. Do not run full Buildroot/QEMU smoke for every small version. The next formal full checkpoint is `v0.2.640`.

## Safety Checklist

Before committing, verify:

- no `.env`, key, token, password, private key, or local credential was added;
- no Docker socket mount was introduced;
- no privileged container command was introduced;
- no host networking was introduced;
- no broad host-directory mount was introduced;
- no host-root mutation path was introduced;
- no desktop-facing output exposes state root, raw launcher output, backend details, Wine command lines, QEMU command lines, or local absolute launcher paths;
- `docs/claude-code-implementation-packages.md` remains untouched.

## Stop Conditions

Stop and report instead of guessing if:

- the required change would broaden host access;
- the task needs privileged containers, host networking, or Docker socket mounts;
- the KDE artifact would need to own Runtime policy;
- QEMU/Wine prerequisites are missing and the smoke cannot safely SKIP;
- Docker Hub or Colima state blocks container verification after targeted local tests pass.
