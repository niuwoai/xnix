# Xnix

Xnix is an atomic KDE Plasma desktop project focused on making existing Windows applications feel native on Linux.

The project is currently at `v0.2.640-rc126`.

## Product Direction

- KDE Plasma is the first flagship desktop shell.
- The Xnix AI Compatibility Runtime owns recipes, Wine/VM backend policy, permissions, snapshots, rollback, launch planning, diagnostics, and execution gates.
- KDE-facing components display state, collect user intent, and forward actions; they do not own compatibility decisions or raw backend launch commands.
- Important Runtime product logic is Go-first. C remains for low-level or already-owned Runtime policy surfaces. Ruby is the test and development-tool harness.
- Host impact must stay narrow: no privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation in normal development checks.

## Current Checkpoint

v0.2.640-rc126 lets the Go Runtime container X GUI smoke carry ad-hoc application identity through `--app-id`, `--display-name`, and `--app-version` when no recipe is being used. This gives the next real external Windows GUI app lane a Go-owned evidence identity without weakening recipe-backed digest validation or opening backend launch, host mounts, Docker socket mounts, host networking, or host-root mutation.

The previous v0.2.640-rc125 checkpoint surfaced the staged desktop launcher/session-gate facts from the real Notepad `.desktop` run inside the Go Runtime real GUI packet and KDE GUI evidence card. The packet now preserves launch authorization, session-gated dispatch, controlled execution session, digest verification, Runtime-owner/KDE read-model consumability, and post-review dispatch fields while keeping backend launch, raw command output, host mounts, Docker socket mounts, host networking, and host-root mutation closed.

The previous v0.2.640-rc124 checkpoint let the staged desktop Notepad smoke carry its real `.desktop`/managed-launcher run into Go Runtime real GUI packet evidence and a KDE page. The Go packet now accepts the delegated `windows-app-container-x-gui-smoke` payload returned by `xnix-compat-launch`, so Ruby does not need to wrap the launcher result into a synthetic report.

The previous v0.2.640-rc123 checkpoint let `compatibility-center-preview` and `kde-center-page-preview` consume the Go Runtime `real-winapp-gui-evidence-packet-preview` file directly through `--known-app-evidence-file`. The container GUI evidence packet harness now renders the KDE page from that Go-owned packet instead of the older GUI evidence projection.

The previous v0.2.640-rc122 checkpoint changed the container GUI evidence packet smoke harness to call the Go Runtime `real-winapp-gui-evidence-packet-preview` command and persist its packet before rendering the KDE page. Ruby still orchestrates the test, while Runtime packet semantics now stay in Go.

The previous v0.2.640-rc121 checkpoint added the Go Runtime `real-winapp-gui-evidence-packet-preview` command. It consumes a passed real Windows GUI smoke report and emits a KDE-safe desktop evidence packet for recipe-backed Notepad container X GUI runs, preserving the observed X window, recipe identity, network-isolated container state, zero host mounts, closed launch gates, and closed host-root mutation boundaries.

The previous v0.2.640-rc120 checkpoint fixed the formal staged launcher dispatch smoke for recipe-backed desktop activation. The Ruby smoke scripts now expect the Notepad desktop stage to include both the managed launcher executable and the packaged recipe registry/application recipe materials before running the controlled launcher path.

The previous v0.2.640-rc119 checkpoint made the staged desktop Notepad smoke easier to diagnose. The script now persists the delegated `xnix-compat-launch` JSON payload before enforcing pass/fail assertions, so skipped or failed Docker/Wine/Xvfb runs leave a concrete `/tmp` evidence file for debugging while the normal path still runs the real packaged-registry desktop entry without mutating the host root.

The previous v0.2.640-rc118 checkpoint added a repeatable staged desktop Notepad smoke. `scripts/staged_desktop_notepad_smoke.rb` stages the KDE `.desktop` entry and managed launcher into a temporary root, reads the real packaged-registry `Exec=` line, injects Runtime-owned launch/session receipts, and runs recipe-backed Notepad through the restricted Docker/Wine/Xvfb path without writing to the host root.

The previous v0.2.640-rc117 checkpoint made the packaged-registry desktop launcher shape executable in safe local staging. `xnix-compat-launch` now resolves `/usr/share/xnix/compatibility/recipes/registry.json` through `XNIX_STAGING_ROOT` when the host path is absent, so staged `.desktop` and KRunner launcher arguments can exercise recipe-backed Notepad dispatch without writing to the host root.

The previous v0.2.640-rc116 checkpoint aligned KRunner search actions with the packaged recipe-backed launcher. Recipe-backed container GUI matches now return `xnix-compat-launch --app APP --registry /usr/share/xnix/compatibility/recipes/registry.json`, so KDE search no longer drops the registry argument needed by installed Notepad desktop launches.

The previous v0.2.640-rc115 checkpoint made the real local Notepad GUI path less host-specific. The container X GUI Runtime now inspects the local Wine image platform and uses that platform when `--platform` is omitted, so Colima arm64 images can run recipe-backed Notepad without manually adding `--platform linux/arm64`.

The previous v0.2.640-rc114 checkpoint fixed the Go desktop identity Ruby smoke so its Dockerfile check follows the current `go test -timeout 90m ./...` validation command. The v0.2.640-rc113 packaged-registry Notepad launcher contract remains the active desktop-path checkpoint.

The previous v0.2.640-rc113 checkpoint aligned the packaged-registry Notepad launcher contract across the Ruby desktop-entry renderer, Ruby Runtime daemon previews, the C Runtime core, D-Bus smoke fixtures, and installer tests. Both Go and Ruby desktop material now route recipe-backed container GUI apps through `xnix-compat-launch --app org.xnix.sample.notepad --registry /usr/share/xnix/compatibility/recipes/registry.json %U`, while the desktop entry still hides Wine, Docker, and raw executable details.

The previous v0.2.640-rc112 checkpoint made recipe-backed Notepad desktop entries self-contained for installed systems. `desktop-entry-preview` now renders `xnix-compat-launch --app org.xnix.sample.notepad --registry /usr/share/xnix/compatibility/recipes/registry.json %U`, while activation staging plans the matching packaged recipe registry and application recipe files so KDE still calls only the Runtime launcher and never sees Wine, Docker, or raw executable details.

The previous v0.2.640-rc109 checkpoint tied the local real GUI proof to an application recipe. `org.xnix.sample.notepad` now carries `container_gui_smoke` hints, and `xnix-runtime-go windows-app-container-x-gui-smoke --registry PATH --recipe-app org.xnix.sample.notepad` resolves `notepad.exe` through the digest-verified recipe before running the restricted Docker/Wine/Xvfb smoke.

The previous v0.2.640-rc108 checkpoint made the local real GUI proof a one-shot evidence packet. `scripts/container_gui_evidence_packet.rb` runs the restricted `container-x-gui` Windows GUI smoke, saves the redacted report, projects Runtime GUI evidence, renders a KDE Center page, and prints a path-redacted packet proving X window observation while keeping desktop/backend launch, host networking, Docker socket mounts, broad host mounts, and host-root mutation closed.

The previous v0.2.640-rc107 checkpoint made the local real GUI proof durable. `scripts/winapp_smoke.rb --format json --report-output PATH --backend container-x-gui` now writes the same redacted report it prints, so the passed Windows GUI run can feed `gui-smoke-evidence-preview`, `compatibility-center-preview`, and `kde-center-page-preview` without manually copying stdout.

The previous v0.2.640-rc106 checkpoint connected the local container GUI proof to the Runtime/KDE read model. `gui-smoke-evidence-preview` now consumes passed `scripts/winapp_smoke.rb --backend container-x-gui` JSON reports as `winapp-smoke-container-x-gui` evidence, and `compatibility-center-preview` / `kde-center-page-preview` can bind that report to an explicit app id, display name, and app version before rendering safe GUI cards without exposing report paths, raw output, backend commands, host mounts, Docker sockets, host networking, or host-root mutation authority.

The previous v0.2.640-rc105 checkpoint promoted the local real GUI proof into maintained smoke entrypoints. `scripts/winapp_smoke.rb --backend container-x-gui` now produces JSON/Markdown evidence for Xvfb startup, Wine bootstrap, and `x_window_observed`, while `ruby scripts/container.rb winapp-container-x-gui-smoke` exposes the path through the constrained project container harness.

The previous v0.2.640-rc104 checkpoint added `xnix-runtime-go windows-app-container-x-gui-smoke`, a restricted Docker fallback that starts Wine on an Xvfb desktop and verifies a real Windows GUI window with `xwininfo`. This gives the project a repeatable local `notepad.exe`-style GUI proof when q4 is unavailable, while the q4/QEMU/Wine guest lane remains the stricter product-grade run.

The previous v0.2.640-rc103 checkpoint fixed desktop activation staging for launcher-only known GUI apps such as Mines. The Runtime now treats empty MIME/file-association data as not applicable, while still staging the normal KDE launcher files, manifest, managed-launcher artifact, and activation receipt needed for a desktop entry.

The previous v0.2.640-rc102 checkpoint let the Go Runtime launch known Wine GUI apps by catalog id. `xnix-runtime-go windows-app-guest-wine-gui-smoke --app org.xnix.apps.mines` resolves the builtin Mines runner and emits the known app id, display name, and version in the JSON result. The local and q4 Wine GUI smoke harnesses can now pass `--known-app-id org.xnix.apps.mines` so real runs exercise the catalog-selected path instead of hand-written guest paths.

The previous v0.2.640-rc101 checkpoint fixed the KDE controlled-launch action route for owner-controlled Wine builtin GUI apps such as Mines. A real q4 `winemine.exe` desktop run can now keep `owner_managed_copy_verified=false` while still producing the safe KDE page and action evidence when the Runtime owner handoff is ready.

The previous v0.2.640-rc100 checkpoint made the q4 real GUI harness produce desktop-consumable evidence automatically after a successful run. `scripts/remote_wine_guest_gui_smoke.rb --execute` now writes the Runtime GUI evidence, a `kde-center-page-preview` JSON payload, and, for owner-controlled launches, a `kde-controlled-launch-action-preview` JSON payload from the same real MessageBox run.

The previous v0.2.640-rc99 checkpoint let the real Wine GUI smoke harness carry the requested application identity into owner-controlled Runtime evidence. MessageBox runs can now pass `--evidence-app-id org.xnix.apps.messagebox --evidence-display-name "Xnix MessageBox"` through the remote q4 harness, the local Xvfb/QEMU/Wine harness, Go `gui-smoke-evidence-preview`, and the Runtime-status owner fixture instead of reusing the Mines identity for the owner handoff.

The previous v0.2.640-rc98 checkpoint added a restricted KDE controlled-launch action smoke lane for the owner-controlled MessageBox GUI handoff. The smoke creates Runtime-status MessageBox handoff evidence, projects the owner GUI report through `gui-smoke-evidence-preview`, renders a full `kde-center-page-preview`, then feeds that full page back into `kde-controlled-launch-action-preview --kde-center-page-file FILE --app APP_ID` without launching a backend by default.

The previous v0.2.640-rc97 checkpoint let the Go Runtime select the safe owner-controlled GUI card from a full `kde-center-page-preview` payload through `kde-controlled-launch-action-preview --kde-center-page-file FILE --app APP_ID`. KDE can hand the page read model back to the Runtime, and the Runtime derives the handoff only from the matching card's single safe `evidence-relative-path` forwarded argument before returning the controlled-launch action preview.

The previous v0.2.640-rc96 checkpoint let the Go Runtime consume a safe KDE GUI card route through `kde-controlled-launch-action-preview --kde-gui-card-file`. The Runtime derives the handoff only from the card's single `evidence-relative-path` forwarded argument, verifies the persisted Runtime evidence and card route metadata, and returns the controlled-launch action preview without exposing owner service arguments or enabling desktop/backend launch.

The previous v0.2.640-rc95 checkpoint added the concrete evidence-only desktop route for handoff-ready owner-controlled GUI cards. KDE now receives `desktop_callable_route=kde-dbus-runtime-status-action`, `desktop_callable_runtime_method=ShowRuntimeControlledLaunch`, `desktop_callable_execution_type=known-app-kde-runtime-status-launch-execution`, `desktop_dbus_method=org.xnix.Compatibility1.ShowRuntimeControlledLaunch`, and a single safe `evidence-relative-path` argument, while owner service args, state roots, cache roots, launcher paths, raw executable paths, backend commands, desktop launch, and backend launch remain closed.

The previous v0.2.640-rc94 checkpoint promoted handoff-ready owner-controlled GUI evidence cards from review-only status to the evidence-only Runtime-controlled launch action. When `owner_evidence_handoff_ready=true`, the known-app GUI evidence summary and KDE GUI card now expose `primary_action_id=show-runtime-controlled-launch`, `primary_action_kind=runtime-status`, and `primary_action_label=Show Runtime-controlled launch`, while desktop launch, backend launch, owner service args, state roots, cache roots, launcher paths, raw executable paths, and backend commands remain closed.

The previous v0.2.640-rc93 checkpoint carried the Runtime owner evidence handoff from the actual owner-controlled Windows GUI smoke into the safe desktop read model. The GUI smoke harness records `owner_evidence_handoff_ready` and the safe relative `owner_evidence_relative_path` produced by the owner fixture, `gui-smoke-evidence-preview` validates that handoff under `runtime/kde-runtime-status-launch-evidence/`, and KDE GUI cards expose owner service-call readiness without exposing state roots, cache roots, launcher paths, raw executable paths, backend commands, or owner service arguments.

The previous v0.2.640-rc92 checkpoint made owner-controlled Windows GUI evidence explicit in the KDE/Compatibility Center read model. KDE Center pages expose `known_app_owner_controlled_gui_evidence_count`, `known_app_owner_managed_copy_verified_count`, and per-card `owner_controlled_runtime_launch_verified` / `owner_managed_copy_verified` fields for `wine-guest-gui-smoke` evidence, so the desktop shell can distinguish a Runtime-owner managed app launch from a generic GUI smoke without gaining launch, backend, raw path, or host mutation authority.

The previous v0.2.640-rc91 checkpoint projected owner-controlled external GUI launch evidence into the Runtime/KDE read model. `gui-smoke-evidence-preview` verifies the owner seed projection, owner service handoff, managed launcher invocation, delegated `wine-guest-gui-smoke` source, delegated execution/window observation, and owner-managed artifact copy before marking `org.xnix.apps.messagebox` as `owner-controlled-gui-qemu-wine-verified`. Unsafe delegated host mutation, Docker socket, broad mount, raw command, backend detail, and raw path details remain closed. The q4 owner-controlled MessageBox run passed with `owner_managed_copy_verified=true`, `owner_delegated_managed_artifact_copied=true`, `owner_delegated_controlled_session_window_observed=true`, and evidence persisted at `/tmp/xnix-run-materials/state/wine-gui-messagebox-owner-0.2.640-rc91-final-evidence.json`.

The previous v0.2.640-rc90 checkpoint moved external GUI executable delivery into the owner-controlled launch path. `XNIX_RUNTIME_OWNER_GUI_EXECUTABLE` is consumed only by the Runtime owner boundary and forwarded to `xnix-compat-launch --executable`, letting the managed launcher copy the external `.exe` into the QEMU guest with SCP before starting it through the Go-owned Wine GUI smoke lane. `XNIX_RUNTIME_OWNER_GUEST_GUI_APP` remains available for already-present guest paths, but the q4 MessageBox path now proves owner-managed copy delivery.

The previous v0.2.640-rc89 checkpoint extended the owner-controlled GUI path from Wine's built-in Mines target to a Runtime-owner-supplied external Windows GUI executable already copied into the managed guest. `XNIX_RUNTIME_OWNER_GUEST_GUI_APP` is consumed only by the Runtime owner boundary and forwarded to `xnix-compat-launch --gui-app`, so KDE still forwards only the evidence handle while the launcher can start the controlled guest `.exe` through Wine without exposing raw paths in product JSON. `org.xnix.apps.messagebox` is now a development recipe identity for the q4 `xnix-messagebox-smoke.exe` GUI fixture.

The previous v0.2.640-rc88 checkpoint made the real Mines GUI path repeatable through the Runtime owner entrypoint. `scripts/wine_guest_gui_smoke.rb --launch-mode owner-controlled-launch` now starts Xvfb and QEMU, generates a real GUI smoke evidence seed, records the Runtime-status launch owner fixture, calls `xnix-runtime-owner --service-call ShowRuntimeControlledLaunch`, and verifies that `xnix-compat-launch` launches `winemine.exe` through Wine with `wine-guest-gui-smoke` delegated evidence and a host-observed X window. `scripts/remote_wine_guest_gui_smoke.rb --launch-mode owner-controlled-launch --execute` builds `xnix-runtime-go`, `xnix-runtime-owner`, and `xnix-compat-launch` on q4 before running the same restricted chain.

The previous v0.2.640-rc87 checkpoint promoted `org.xnix.apps.mines` from GUI evidence handoff into the first Runtime-owner-managed GUI launch execution path. `xnix-compat-launch` can now dispatch guest-builtin GUI known apps through the Go-owned Wine GUI smoke lane after human authorization and controlled execution session gates, while `ShowRuntimeControlledLaunch` forwards only owner-side guest parameters and preserves `wine-guest-gui-smoke` delegated evidence for KDE-safe projections.

The previous v0.2.640-rc86 checkpoint connected the first real Windows GUI app evidence to the Runtime-owned desktop trigger path. `known-app-runtime-status-launch-owner-fixture-record --gui-smoke-evidence-file FILE` can consume a passed `gui-smoke-evidence-preview` projection for `org.xnix.apps.mines`, record Runtime launch authorization/session/review handoff evidence, and let the owner trigger plus KDE action trigger assemble a `ShowRuntimeControlledLaunch` request from an evidence handle. KDE still receives no Runtime state-root, cache-root, launcher path, raw executable path, backend command, or receipt/session reconstruction authority, and the preview does not start a process.

The previous v0.2.640-rc85 checkpoint gave the first real Windows GUI app its own Runtime recipe entry. `org.xnix.apps.mines` now appears in the development recipe registry with a backend-neutral identity, so q4 GUI evidence for `Mines` can land on its own KDE/Compatibility Center page instead of borrowing the sample Notepad recipe.

The previous v0.2.640-rc81 checkpoint tightened and verified the real Windows GUI app smoke loop after q4 showed QEMU, guest SSH, Wine boot, and executable copy succeeding while Xvfb saw no child window. `windows-app-guest-wine-gui-smoke` now polls X window state throughout the observation window, keeps suppressing Wine first-run installer prompts while observing, verifies `guest_x11_driver_available`, and reports `x_window_observation_attempts` for failed real Xvfb/QEMU/Wine runs. The q4 managed Wine guest run material was updated to the rebuilt X11-driver image, and an unattended `winemine.exe` smoke passed with `x_window_observed=true` in `/home/xnix-run-materials/state/wine-gui-winemine-0.2.640-rc81-final.json`.

The previous v0.2.640-rc80 checkpoint let `xnix-kde-center-model` render safe real Windows GUI evidence cards from Runtime KDE Center page payloads. The model now exposes GUI evidence counts, passed-evidence counts, display fields, Runtime dispatch status, and read-only action state while filtering cards that expose backend details, host mutation, KDE policy ownership, or raw artifact paths.

The previous v0.2.640-rc79 checkpoint carried real Windows GUI evidence into the KDE shell package. The Compatibility Center plasmoid now has a read-only GUI evidence card entry that points KDE surfaces at `kde-center-page-preview`, `known_app_gui_evidence_count`, and `known_app_gui_evidence_cards`, while keeping launch, backend startup, repair, snapshot, restore, and host mutation disabled.

The previous v0.2.640-rc78 checkpoint connected the real Windows GUI smoke lane to product-facing Runtime read models. `xnix-runtime-go gui-smoke-evidence-preview --gui-smoke-report FILE` consumes executed Wine guest GUI smoke JSON, verifies Go-owned QEMU/Wine/X11 evidence, and emits KDE-safe GUI run evidence for Compatibility Center and KDE Center previews.

The previous v0.2.640-rc77 checkpoint turned the guest Wine GUI smoke from an in-guest-only launcher into a repeatable local-artifact delivery path. `xnix-runtime-go windows-app-guest-wine-gui-smoke --executable FILE.exe` validates a local Windows executable, copies it into the QEMU guest over loopback SCP, launches the copied guest path with Wine on the X display, and observes the resulting X window.

The previous v0.2.640-rc75 checkpoint moved the Wine GUI smoke core into Go. `xnix-runtime-go windows-app-guest-wine-gui-smoke` owns guest `wineboot`, GUI app launch, host `xwininfo` observation, and KDE-safe evidence, while `scripts/wine_guest_gui_smoke.rb` remains the Xvfb/QEMU harness. The remote q4 GUI harness builds the current Go Runtime before running the smoke so it does not depend on stale binaries.

The previous v0.2.640-rc74 checkpoint made the q4 Wine GUI lane reproducible from the repository. `scripts/remote_wine_guest_gui_smoke.rb --execute` syncs the checkout to a managed `/home/xnix-*` source root, runs `scripts/wine_guest_gui_smoke.rb` remotely against the q4 Wine guest kernel, and persists JSON evidence under the q4 run-materials state directory. The plan-only mode records that the pass condition requires a rebuilt Wine guest with X11 support.

The previous v0.2.640-rc73 checkpoint opened the next real Windows GUI-app lane. The Wine i386 guest defconfig now includes the X.org client libraries and font/image dependencies that make Wine build with its X11 driver, and `scripts/wine_guest_gui_smoke.rb --execute` can run a real Wine GUI app through QEMU against an Xvfb desktop while keeping SSH loopback-bound and reporting the temporary display-network exception explicitly.

The previous v0.2.640-rc72 checkpoint wired the real known-app matrix evidence into the product read models. `compatibility-center-preview` and `kde-center-page-preview` now accept `--known-app-matrix-report FILE`, consume the Go-owned matrix evidence projection, and surface KDE-safe matrix cards and counts without exposing remote paths or raw output.

The previous v0.2.640-rc71 checkpoint added a Go-owned known-app matrix evidence consumer. `xnix-runtime-go known-app-matrix-evidence-preview --matrix-report FILE` reads the aggregate q4 matrix JSON, validates the executed two-app QEMU/Wine evidence, preserves checksum, marker, redaction, serial-log, QEMU, and Wine counts, and projects KDE-safe per-app `known_app_smoke_evidence` without exposing remote paths or raw output.

The previous v0.2.640-rc70 checkpoint made the q4 real Windows app matrix evidence easier for downstream consumers to find. `scripts/remote_known_winapp_matrix_smoke.rb --execute` persists the aggregate matrix JSON to `/home/xnix-run-materials/state/known-run-matrix-<version>.json` by default, or to an operator-selected `--matrix-report-output` path under `/home/xnix-*`.

The previous v0.2.640-rc69 checkpoint added the remote known-app matrix smoke for the q4 real Windows app lane, recording per-app JSON and serial-log evidence plus aggregate pass/fail counts. The earlier v0.2.640-rc68 checkpoint added BusyBox-w32 as the second pinned real known app.

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
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/winapp ./cmd/xnix-runtime-go -run 'TestRunGuestGUISmoke|TestWindowsAppGuestWineGUISmoke' -count=1
ruby -Ilib test/test_staged_launcher_dispatch_smoke_script.rb
ruby -Ilib test/test_runtime_dbus_smoke_script.rb
ruby -Ilib test/test_runtime_status_owner_service_session_bus_smoke_script.rb
ruby -Ilib test/test_dbus_controlled_launch_owner_fixture_smoke_script.rb
ruby -Ilib test/test_kde_controlled_launch_action_stub.rb
ruby -Ilib test/test_kde_controlled_launch_action_smoke_script.rb
ruby -Ilib test/test_desktop_trigger_request_preflight_smoke_script.rb
ruby -Ilib test/test_wine_guest_gui_smoke_script.rb
ruby -Ilib test/test_remote_wine_guest_gui_smoke_script.rb
ruby -Ilib test/test_kde_center_model.rb
ruby -Ilib test/test_kde_compatibility_center_plasmoid.rb
ruby scripts/desktop_trigger_request_preflight_smoke.rb
ruby scripts/desktop_trigger_request_preflight_smoke.rb --format json
ruby -Ilib test/test_full_smoke_script.rb
ruby -Ilib test/test_full_checkpoint_promotion_packet.rb
ruby -Ilib test/test_merge_readiness_packet.rb
ruby -Ilib test/test_release_evidence_index.rb
gcc -std=c11 -Wall -Wextra -Werror runtime/dbus/xnix_compatd_smoke.c -o /tmp/xnix-dbus-smoke-check $(pkg-config --cflags --libs gio-2.0)
ruby scripts/verify_layout.rb
```

Run full Buildroot/QEMU smoke at the configured checkpoint cadence. The current formal full checkpoint candidate is `v0.2.640-rc60`; promote to `v0.2.640` only after `ruby scripts/full_smoke.rb` passes and the promotion packet allows the promotion. v0.2.640-rc100 received an explicitly approved restricted full-smoke run before v0.2.640-rc101 continued the post-boundary GUI action fix.
