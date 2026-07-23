# Xnix

Xnix is an atomic KDE Plasma desktop project focused on making existing Windows applications feel native on Linux.

The project is currently at `v0.2.627`.

## Product Direction

- KDE Plasma is the first flagship desktop shell.
- The Xnix AI Compatibility Runtime owns recipes, Wine/VM backend policy, permissions, snapshots, rollback, launch planning, diagnostics, and execution gates.
- KDE-facing components display state, collect user intent, and forward actions; they do not own compatibility decisions or raw backend launch commands.
- Important Runtime product logic is Go-first. C remains for low-level or already-owned Runtime policy surfaces. Ruby is the test and development-tool harness.
- Host impact must stay narrow: no privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation in normal development checks.

## Current Checkpoint

v0.2.627 adds a private session-bus wrapper for the Runtime-status owner service smoke. `ruby scripts/container.rb runtime-status-owner-service-session-bus-smoke` runs the staged launcher dispatch smoke inside `dbus-run-session`, safely skips when session-bus or QEMU/Wine prerequisites are unavailable, and otherwise exercises the real staged `xnix-compat-launch` lane through `xnix-runtime-owner --service-call ShowRuntimeControlledLaunch evidence-relative-path <relative>`.

The desktop-safe result keeps state-root paths, raw launcher output, backend details, Docker socket mounts, broad host mounts, and host-root mutation out of KDE-facing JSON.

## Main References

- [PRODUCT_OVERVIEW.md](PRODUCT_OVERVIEW.md) describes the current product state, architecture, and release notes.
- [docs/xnix-current-mainline.md](docs/xnix-current-mainline.md) is the Codex-owned implementation mainline.
- [docs/mainline-integration-checkpoint.md](docs/mainline-integration-checkpoint.md) captures merge-lane review rules.
- [AGENTS.md](AGENTS.md) defines repository contribution and safety rules.

Historical `docs/claude-code-*` files remain repository evidence, but new implementation work should start from the current mainline unless a user explicitly requests a handoff artifact.

## Focused Verification

Run targeted checks for small versions:

```text
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/appidentity -run 'TestPrepareKnownAppKDERuntimeStatusLaunchExecutionFromActionTriggerConsumesHandoff|TestPrepareKnownAppKDERuntimeStatusLaunchExecutionRevalidatesRuntimeOwnedState|TestPreviewKnownAppKDERuntimeStatusLaunchActionTriggerConsumesHandoff' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./cmd/xnix-runtime-go -run 'TestShowRuntimeControlledLaunchCommandForwardsOnlyEvidenceHandoff|TestShowRuntimeControlledLaunchCommandRejectsReconstructedFields|TestKnownAppKDERuntimeStatusLaunchExecutionCommandConsumesActionTriggerHandoff|TestKnownAppKDERuntimeStatusLaunchActionTriggerCommandAssemblesLaunchRequestFromHandoff' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/owner -run 'TestServiceCallDispatchesShowRuntimeControlledLaunch|TestServiceCallServesReadDispatchInProcess' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./cmd/xnix-runtime-owner -run 'TestRuntimeOwnerCommandRendersShowRuntimeControlledLaunchServiceCall|TestRuntimeOwnerCommandRendersRedactedAdapterProfileOwnerLocalReadDispatch' -count=1
ruby -Ilib test/test_staged_launcher_dispatch_smoke_script.rb
ruby -Ilib test/test_runtime_status_owner_service_session_bus_smoke_script.rb
ruby scripts/verify_layout.rb
```

Run full Buildroot/QEMU smoke only at the configured checkpoint cadence. The next formal full checkpoint is `v0.2.640`.
