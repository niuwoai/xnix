# Xnix

Xnix is an atomic KDE Plasma desktop project focused on making existing Windows applications feel native on Linux.

The project is currently at `v0.2.640-rc14`.

## Product Direction

- KDE Plasma is the first flagship desktop shell.
- The Xnix AI Compatibility Runtime owns recipes, Wine/VM backend policy, permissions, snapshots, rollback, launch planning, diagnostics, and execution gates.
- KDE-facing components display state, collect user intent, and forward actions; they do not own compatibility decisions or raw backend launch commands.
- Important Runtime product logic is Go-first. C remains for low-level or already-owned Runtime policy surfaces. Ruby is the test and development-tool harness.
- Host impact must stay narrow: no privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation in normal development checks.

## Current Checkpoint

v0.2.640-rc14 keeps the formal full checkpoint gate pending and adds `docs/post-checkpoint-promotion-checklist-0640.md`, the human-operator runbook for promoting `v0.2.640` only after full smoke passes and the promotion packet allows the release.

The previous v0.2.640-rc13 checkpoint made `scripts/release_evidence_index.rb` consume `scripts/full_checkpoint_promotion_packet.rb`, keeping historical product smoke evidence separate from current formal release readiness. The earlier v0.2.640-rc12 checkpoint made merge readiness consume the same promotion packet as the single release-promotion gate and expose `full_checkpoint_promotion_status`.

## Main References

- [PRODUCT_OVERVIEW.md](PRODUCT_OVERVIEW.md) describes the current product state, architecture, and release notes.
- [docs/xnix-current-mainline.md](docs/xnix-current-mainline.md) is the Codex-owned implementation mainline.
- [docs/mainline-integration-checkpoint.md](docs/mainline-integration-checkpoint.md) captures merge-lane review rules.
- [AGENTS.md](AGENTS.md) defines repository contribution and safety rules.

Historical `docs/claude-code-*` files remain repository evidence, but new implementation work should start from the current mainline unless a user explicitly requests a handoff artifact.

## Focused Verification

Run targeted checks for small versions:

```text
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/owner ./internal/runtime/appidentity ./cmd/xnix-runtime-go -run 'TestDesktopTriggerDryRunRequestReview|TestPreviewLaunchEnvelopeGuard|TestKDEControlledLaunchActionSurfaceAudit|TestPreviewManagedLauncherAcceptance' -count=1
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
ruby -Ilib test/test_full_smoke_script.rb
ruby -Ilib test/test_full_checkpoint_promotion_packet.rb
ruby -Ilib test/test_merge_readiness_packet.rb
ruby -Ilib test/test_release_evidence_index.rb
gcc -std=c11 -Wall -Wextra -Werror runtime/dbus/xnix_compatd_smoke.c -o /tmp/xnix-dbus-smoke-check $(pkg-config --cflags --libs gio-2.0)
ruby scripts/verify_layout.rb
```

Run full Buildroot/QEMU smoke at the configured checkpoint cadence. The current formal full checkpoint candidate is `v0.2.640-rc14`; promote to `v0.2.640` only after `ruby scripts/full_smoke.rb` passes and the promotion packet allows the promotion.
