# Xnix

Xnix is an atomic KDE Plasma desktop project focused on making existing Windows applications feel native on Linux.

The project is currently at `v0.2.614`.

## Product Direction

- KDE Plasma is the first flagship desktop shell.
- The Xnix AI Compatibility Runtime owns recipes, Wine/VM backend policy, permissions, snapshots, rollback, launch planning, diagnostics, and execution gates.
- KDE-facing components display state, collect user intent, and forward actions; they do not own compatibility decisions or raw backend launch commands.
- Important Runtime product logic is Go-first. C remains for low-level or already-owned Runtime policy surfaces. Ruby is the test and development-tool harness.
- Host impact must stay narrow: no privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation in normal development checks.

## Current Checkpoint

v0.2.614 adds the Go-owned `known-app-kde-runtime-status-launch-execution` entrypoint. It consumes the KDE Runtime-status launch request, accepts the Runtime state root only at the Runtime boundary, revalidates the launch authorization receipt, session-gated review receipt, controlled execution session, post-review dispatch state, and managed guest boundary, then invokes an existing managed launcher executable with state-root injected into argv.

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
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/appidentity -run 'TestPrepareKnownAppKDERuntimeStatusLaunchExecution|TestPreviewKnownAppKDERuntimeStatusLaunchRequest' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./cmd/xnix-runtime-go -run 'TestKnownAppKDERuntimeStatusLaunchExecutionCommandInvokesManagedLauncherWithRuntimeStateRoot|TestKnownAppKDERuntimeStatusLaunchRequestPreviewCommandCollectsOpaqueIDs' -count=1
ruby scripts/verify_layout.rb
```

Run full Buildroot/QEMU smoke only at the configured checkpoint cadence. The next formal full checkpoint is `v0.2.620`.
