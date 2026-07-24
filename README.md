# Xnix

Xnix is an atomic KDE Plasma desktop project focused on making existing Windows applications feel native on Linux.

The project is currently at `v0.2.640-rc45`.

## Product Direction

- KDE Plasma is the first flagship desktop shell.
- The Xnix AI Compatibility Runtime owns recipes, Wine/VM backend policy, permissions, snapshots, rollback, launch planning, diagnostics, and execution gates.
- KDE-facing components display state, collect user intent, and forward actions; they do not own compatibility decisions or raw backend launch commands.
- Important Runtime product logic is Go-first. C remains for low-level or already-owned Runtime policy surfaces. Ruby is the test and development-tool harness.
- Host impact must stay narrow: no privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation in normal development checks.

## Current Checkpoint

v0.2.640-rc45 parses the PE machine architecture for direct and profile-backed Windows app smoke inputs, reporting `executable_architecture` and `executable_architecture_supported` before any runner execution. The previous v0.2.640-rc44 checkpoint made direct `windows-app-run-smoke --exe <path>` validate the executable's Windows `MZ` signature.

The previous v0.2.640-rc20 checkpoint added JSON and Markdown evidence reports to `scripts/winapp_smoke.rb`. Text mode remains the quick developer PASS/SKIP path, while JSON and Markdown default to `windows-app-run-smoke --redact-output` so future merge and release tooling can consume real Windows executable smoke evidence without raw stdout or stderr.

The previous v0.2.640-rc19 checkpoint added `--redact-output` to the Go-owned `windows-app-run-smoke` path. The local Windows app smoke still keeps raw output by default for developer marker checks, but product-facing consumers can now request marker, byte, and line-count evidence without raw stdout or stderr, moving the real executable runner closer to KDE-safe Compatibility Center consumption.

The previous v0.2.640-rc18 checkpoint made `scripts/release_evidence_index.rb` optionally classify existing desktop-trigger request preflight smoke JSON evidence through `--desktop-trigger-request-preflight-smoke`. The index never runs the smoke, emits a `desktop-trigger-request-preflight-smoke` claim, treats missing evidence as skipped, and blocks failed or malformed supplied evidence while keeping release evidence indexing offline and side-effect free. `docs/claude-code-active-work-order.md` is the short current handoff for dispatching the next Claude Code task without reworking completed C9W2 through C9W4 lanes.

The previous v0.2.640-rc17 checkpoint made `scripts/merge_readiness_packet.rb` optionally consume existing desktop-trigger request preflight smoke JSON evidence through `--desktop-trigger-request-preflight-smoke`. The packet never runs the smoke by default, treats missing evidence as non-blocking, surfaces `desktop_trigger_request_preflight_smoke_status`, and turns failed or malformed supplied evidence into a release-only blocker while keeping merge readiness offline and side-effect free.

The previous v0.2.640-rc16 checkpoint added `scripts/desktop_trigger_request_preflight_smoke.rb`, a lightweight targeted smoke that records safe Runtime-status handoff evidence, invokes the Go-owned `desktop-trigger-request-preflight-preview`, verifies `blocked-missing-promotion` before formal promotion, verifies `ready-for-operator-request` with explicit promoted evidence, and keeps service dispatch, D-Bus calls, desktop launch, backend launch, Runtime writes, KDE writes, backend exposure, and host mutation disabled.

The previous v0.2.640-rc15 checkpoint added the Go-owned `desktop-trigger-request-preflight-preview` for the post-release real desktop-triggered `ShowRuntimeControlledLaunch` request lane. The preflight consumes service-call materialization, keeps KDE evidence-only, hides owner service call arguments, and blocks with `blocked-missing-promotion` until formal full checkpoint promotion is observed.

The previous v0.2.640-rc14 checkpoint kept the formal full checkpoint gate pending and added `docs/post-checkpoint-promotion-checklist-0640.md`, the human-operator runbook for promoting `v0.2.640` only after full smoke passes and the promotion packet allows the release. The earlier v0.2.640-rc13 checkpoint made `scripts/release_evidence_index.rb` consume `scripts/full_checkpoint_promotion_packet.rb`, keeping historical product smoke evidence separate from current formal release readiness.

## Main References

- [PRODUCT_OVERVIEW.md](PRODUCT_OVERVIEW.md) describes the current product state, architecture, and release notes.
- [docs/windows-app-smoke-profile-runbook.md](docs/windows-app-smoke-profile-runbook.md) is the operator runbook for repeatable local real Windows app smoke profiles.
- [docs/xnix-current-mainline.md](docs/xnix-current-mainline.md) is the Codex-owned implementation mainline.
- [docs/mainline-integration-checkpoint.md](docs/mainline-integration-checkpoint.md) captures merge-lane review rules.
- [AGENTS.md](AGENTS.md) defines repository contribution and safety rules.

Historical `docs/claude-code-*` files remain repository evidence, but new implementation work should start from the current mainline unless a user explicitly requests a handoff artifact.

## Focused Verification

Run targeted checks for small versions:

```text
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/owner ./internal/runtime/appidentity ./cmd/xnix-runtime-go -run 'TestDesktopTriggerDryRunRequestReview|TestPreviewLaunchEnvelopeGuard|TestKDEControlledLaunchActionSurfaceAudit|TestPreviewManagedLauncherAcceptance' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/owner ./cmd/xnix-runtime-go -run 'TestPreviewDesktopTriggerRequestPreflight|TestDesktopTriggerRequestPreflight' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/owner ./internal/runtime/appidentity ./cmd/xnix-runtime-go -run 'TestDesktopTriggerServiceCallMaterialization|TestDesktopTriggerDryRunRequestReview|TestKnownAppRuntimeStatusLaunchOwnerTrigger' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/appidentity -run 'TestPreviewDesktopTriggerStagedInvocationReadiness|TestPreviewKDEControlledLaunchSessionBusSmokePlanLinksActionToRestrictedSmoke|TestPreviewKDEControlledLaunchActionForwardsOnlyEvidenceHandle|TestPreviewKnownAppRuntimeStatusLaunchOwnerTriggerConsumesVerifiedHandoff' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./cmd/xnix-runtime-go -run 'TestDesktopTriggerStagedInvocationReadiness|TestKDEControlledLaunchSessionBusSmokePlanPreviewCommandLinksActionToRestrictedSmoke|TestKDEControlledLaunchSessionBusSmokePlanPreviewCommandRejectsMissingStateRoot|TestKDEControlledLaunchActionPreviewCommandForwardsOnlyEvidenceHandle|TestKDEControlledLaunchActionPreviewCommandRejectsMissingStateRoot|TestKnownAppRuntimeStatusLaunchOwnerTriggerPreviewCommandConsumesVerifiedHandoff' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/owner -run 'TestServiceCallDispatchesShowRuntimeControlledLaunch|TestServiceCallServesReadDispatchInProcess' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./cmd/xnix-runtime-owner -run 'TestRuntimeOwnerCommandRendersShowRuntimeControlledLaunchServiceCall|TestRuntimeOwnerCommandRendersRedactedAdapterProfileOwnerLocalReadDispatch' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/appidentity -run 'TestRecordKnownAppRuntimeStatusLaunchOwnerFixtureBlocksWithoutVerifiedArtifact' -count=1
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./cmd/xnix-runtime-go -run 'TestKnownAppRuntimeStatusLaunchOwnerFixtureRecordCommandBlocksWithoutArtifact|TestKnownAppRuntimeStatusLaunchOwnerFixtureRecordCommandRejectsMissingStateRoot' -count=1
ruby -Ilib test/test_staged_launcher_dispatch_smoke_script.rb
ruby -Ilib test/test_runtime_dbus_smoke_script.rb
ruby -Ilib test/test_runtime_status_owner_service_session_bus_smoke_script.rb
ruby -Ilib test/test_dbus_controlled_launch_owner_fixture_smoke_script.rb
ruby -Ilib test/test_kde_controlled_launch_action_stub.rb
ruby -Ilib test/test_kde_controlled_launch_action_smoke_script.rb
ruby -Ilib test/test_desktop_trigger_request_preflight_smoke_script.rb
ruby scripts/desktop_trigger_request_preflight_smoke.rb
ruby scripts/desktop_trigger_request_preflight_smoke.rb --format json
ruby -Ilib test/test_full_smoke_script.rb
ruby -Ilib test/test_full_checkpoint_promotion_packet.rb
ruby -Ilib test/test_merge_readiness_packet.rb
ruby -Ilib test/test_release_evidence_index.rb
gcc -std=c11 -Wall -Wextra -Werror runtime/dbus/xnix_compatd_smoke.c -o /tmp/xnix-dbus-smoke-check $(pkg-config --cflags --libs gio-2.0)
ruby scripts/verify_layout.rb
```

Run full Buildroot/QEMU smoke at the configured checkpoint cadence. The current formal full checkpoint candidate is `v0.2.640-rc45`; promote to `v0.2.640` only after `ruby scripts/full_smoke.rb` passes and the promotion packet allows the promotion.
