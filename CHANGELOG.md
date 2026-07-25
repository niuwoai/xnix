# Changelog

Xnix follows Semantic Versioning.

Older release entries are archived in [CHANGELOG_ARCHIVE.md](CHANGELOG_ARCHIVE.md).

## [0.2.640-rc113] - 2026-07-25

### Changed

- Changed the Ruby desktop-entry renderer and Runtime daemon desktop plan to emit the packaged recipe registry launcher for `container_gui_smoke` applications.
- Changed the C Runtime core and D-Bus smoke fixtures to expose the same packaged-registry Notepad launcher contract as the Go Runtime.
- Added Ruby desktop-entry coverage for recipe-backed container GUI applications while keeping raw Windows executable and backend details out of desktop files.

## [0.2.640-rc112] - 2026-07-25

### Added

- Added packaged recipe registry routing for `container_gui_smoke` applications in generated desktop entries, so `org.xnix.sample.notepad` launches through `xnix-compat-launch --app org.xnix.sample.notepad --registry /usr/share/xnix/compatibility/recipes/registry.json %U` instead of depending on a temporary development registry path.
- Added desktop activation staging for recipe-backed container GUI apps to plan the packaged recipe registry and application recipe under `/usr/share/xnix/compatibility/recipes/`.
- Added unit and CLI coverage proving the Notepad desktop entry hides Wine, Docker, raw executable, and temporary registry details while staging the Runtime-owned recipe materials needed by the launcher.

## [0.2.640-rc111] - 2026-07-25

### Added

- Added `org.xnix.sample.notepad` to the Go Runtime known Windows app catalog as a recipe-backed container GUI application.
- Added a recipe-backed container X GUI dispatch path to `xnix-compat-launch`, allowing the Runtime launcher to run the digest-verified Notepad recipe through the restricted Docker/Wine/Xvfb backend after the existing launch authorization, session-gated review, and controlled execution session gates pass.
- Added launcher coverage proving `xnix-compat-launch --app org.xnix.sample.notepad --registry PATH` resolves `notepad.exe` from the recipe, observes the X window, preserves recipe identity, and keeps Docker socket, host networking, broad host mount, raw command, backend detail, and host-root mutation exposure closed.

## [0.2.640-rc110] - 2026-07-25

### Changed

- Changed `gui-smoke-evidence-preview` to preserve recipe-backed container GUI smoke identity as `recipe_backed` and `recipe_app_id` when consuming `org.xnix.sample.notepad` real Notepad evidence.
- Changed KDE Center GUI evidence cards and the Compatibility Center plasmoid field contract to expose safe recipe-backed GUI evidence identity while keeping launch, backend, raw path, Docker socket, host networking, broad mount, and host-root mutation gates closed.
- Changed `scripts/container_gui_evidence_packet.rb` to require the KDE GUI card itself to carry recipe-backed Notepad evidence, not just the underlying smoke report.
- Verified a restricted recipe-backed local Docker/Wine/Xvfb Notepad packet with `kde_card_recipe_backed=true`, `kde_card_recipe_app_id=org.xnix.sample.notepad`, `x_window_observed=true`, `container_network_mode=none`, `container_host_mount_count=0`, no Docker socket mount, no broad host mount, no host networking, and no host-root mutation.

## [0.2.640-rc109] - 2026-07-25

### Added

- Added `container_gui_smoke` recipe hints for `org.xnix.sample.notepad`, letting the Go Runtime resolve the real container GUI smoke target from the registered application id instead of only from a raw executable name.
- Added `windows-app-container-x-gui-smoke --registry PATH --recipe-app ID` so recipe-backed local Docker/Wine/Xvfb GUI runs preserve application id, display name, version, and safe container evidence.
- Changed `scripts/winapp_smoke.rb` and `scripts/container_gui_evidence_packet.rb` to forward recipe-backed container GUI smoke requests and require the one-shot packet to prove the local GUI run was tied to the registered Notepad recipe.
- Verified a restricted recipe-backed local Docker/Wine/Xvfb `org.xnix.sample.notepad` run with `container_recipe_backed=true`, `container_application_id=org.xnix.sample.notepad`, `x_window_observed=true`, `container_network_mode=none`, `container_host_mount_count=0`, no Docker socket mount, no broad host mount, no host networking, and no host-root mutation.

## [0.2.640-rc108] - 2026-07-25

### Added

- Added `scripts/container_gui_evidence_packet.rb`, a one-shot local container GUI evidence packet that runs the restricted `container-x-gui` smoke, persists the redacted report, projects Runtime GUI evidence, renders a KDE Center page, and emits a path-redacted summary packet.
- Added script-level coverage for the packet path so the `winapp_smoke.rb`, `gui-smoke-evidence-preview`, and `kde-center-page-preview` handoff remains guarded without invoking Docker in unit tests.
- Verified a restricted local Docker/Wine/Xvfb `notepad.exe` run through `scripts/container_gui_evidence_packet.rb` with `x_window_observed=true`, `container_network_mode=none`, `container_host_mount_count=0`, no Docker socket mount, no broad host mount, no host networking, and no host-root mutation.

## [0.2.640-rc107] - 2026-07-25

### Added

- Added `scripts/winapp_smoke.rb --report-output PATH` for persisting JSON and Markdown Windows app smoke reports without exposing the output path in the report payload.
- Added container X GUI report-output coverage so the real local GUI app proof can be saved and then consumed by Go Runtime evidence and KDE Center read models.

## [0.2.640-rc106] - 2026-07-25

### Changed

- Changed `gui-smoke-evidence-preview` to consume passed `scripts/winapp_smoke.rb --backend container-x-gui` reports as `winapp-smoke-container-x-gui` evidence without exposing report paths, raw output, backend commands, host mounts, Docker sockets, host networking, or host-root mutation.
- Changed KDE Center GUI evidence cards and known-app evidence normalization to preserve local container GUI proof separately from q4/QEMU Wine GUI proof while keeping desktop and backend launch gates closed.
- Added `--known-app-gui-smoke-app`, `--known-app-gui-smoke-name`, and `--known-app-gui-smoke-version` identity binding for GUI smoke report consumption in Compatibility Center and KDE Center previews.

## [0.2.640-rc105] - 2026-07-25

### Added

- Added `container-x-gui` support to `scripts/winapp_smoke.rb`, forwarding to the Go Runtime `windows-app-container-x-gui-smoke` command and reporting Xvfb/X window evidence.
- Added the `ruby scripts/container.rb winapp-container-x-gui-smoke` entrypoint for running the real Windows GUI smoke through the constrained project container harness.

## [0.2.640-rc104] - 2026-07-25

### Added

- Added the Go Runtime `windows-app-container-x-gui-smoke` command for restricted Docker-based Wine/Xvfb GUI validation with X window observation.

### Changed

- Changed the local Wine smoke image definition and fallback builder to include Xvfb and X11 inspection tools so GUI Windows app smoke runs can observe real desktop windows without privileged containers, host networking, Docker socket mounts, or host directory mounts.

## [0.2.640-rc103] - 2026-07-25

### Fixed

- Fixed desktop activation staging for launcher-only known GUI apps without MIME/file associations, such as `org.xnix.apps.mines`.

### Changed

- Changed desktop activation bundle and staging output to mark file associations as not applicable when an app has no MIME types, while preserving start-menu launcher, KDE service-menu, manifest, managed-launcher, and receipt materials.

## [0.2.640-rc102] - 2026-07-25

### Added

- Added known-app identity selection to `xnix-runtime-go windows-app-guest-wine-gui-smoke` through `--app`, so Wine builtin GUI apps such as Mines can be launched by catalog id instead of guest filesystem path.
- Added `--known-app-id` forwarding to the local and q4 Wine GUI smoke harnesses so real runs can exercise the catalog-selected Go Runtime path.
- Added known app identity fields to GUI smoke JSON output for desktop and Runtime handoff consumers.

### Changed

- Changed the owner-controlled known GUI dispatch path to preserve catalog identity when it delegates to the guest Wine GUI runner.

## [0.2.640-rc101] - 2026-07-25

### Fixed

- Fixed KDE controlled-launch action readiness for owner-controlled Wine builtin GUI apps whose runtime evidence has a safe owner handoff but no managed executable copy.
- Kept owner-managed copy status visible as evidence instead of using it as a hard gate for builtin desktop action routing.

## [0.2.640-rc100] - 2026-07-25

### Added

- Added automatic KDE center page and owner-controlled KDE action preview evidence generation to the q4 remote Wine GUI smoke after a real GUI run passes.
- Added `--kde-page-output` and `--kde-action-output` options to `scripts/remote_wine_guest_gui_smoke.rb` for persisted desktop handoff evidence.

### Changed

- Changed the remote real GUI smoke plan to report whether KDE page and controlled-launch action evidence will be generated for direct versus owner-controlled runs.

## [0.2.640-rc99] - 2026-07-25

### Added

- Added configurable Runtime GUI evidence identity flags to `scripts/wine_guest_gui_smoke.rb` and forwarded them through `scripts/remote_wine_guest_gui_smoke.rb`.

### Changed

- Changed owner-controlled MessageBox GUI smoke planning so the Go GUI evidence projection and Runtime-status owner fixture can use `org.xnix.apps.messagebox` instead of borrowing the Mines identity.

## [0.2.640-rc98] - 2026-07-25

### Added

- Added a restricted KDE controlled-launch action smoke lane that validates the owner-controlled MessageBox GUI handoff through Runtime-status evidence, `gui-smoke-evidence-preview`, `kde-center-page-preview`, and `kde-controlled-launch-action-preview --kde-center-page-file`.

### Changed

- Changed the KDE controlled-launch action smoke default SKIP summary to report GUI center-page handoff validation while keeping actual session-bus or backend execution behind explicit opt-in environment variables.

## [0.2.640-rc97] - 2026-07-25

### Added

- Added Go Runtime support for selecting a safe owner-controlled GUI evidence card from a full `kde-center-page-preview` payload before producing a controlled-launch action preview.
- Added CLI support for `kde-controlled-launch-action-preview --kde-center-page-file FILE --app APP_ID`, allowing KDE to hand the page read model back to the Runtime instead of extracting owner handoff details itself.

### Changed

- Changed the KDE controlled-launch action path to reject unsafe center pages with opened launch/backend gates, backend detail exposure, host-root mutation, missing application identity, or non-unique GUI evidence card matches.

## [0.2.640-rc96] - 2026-07-25

### Added

- Added Go Runtime validation for consuming a safe KDE GUI evidence card as the input to `kde-controlled-launch-action-preview`.
- Added CLI support for `kde-controlled-launch-action-preview --kde-gui-card-file`, deriving the evidence handle only from the card's single safe forwarded argument when `--evidence-relative-path` is omitted.

### Changed

- Changed the KDE controlled-launch action preview path to verify GUI card identity, route metadata, D-Bus method, forwarded evidence handle, owner handoff readiness, and closed desktop/backend launch gates before returning a desktop action.

## [0.2.640-rc95] - 2026-07-25

### Added

- Added evidence-only desktop callable route metadata to handoff-ready owner-controlled GUI evidence cards, including the Runtime owner method, D-Bus action method, Runtime execution request type, and forwarded evidence handle.
- Added KDE bridge and Plasma stub coverage for the safe GUI handoff route while keeping owner service arguments hidden from desktop surfaces.

### Changed

- Changed KDE GUI evidence cards to forward only the safe `evidence-relative-path` handle for Runtime-controlled launch handoff, leaving desktop launch and backend launch disabled.

## [0.2.640-rc94] - 2026-07-25

### Changed

- Changed handoff-ready owner-controlled GUI evidence cards to expose `show-runtime-controlled-launch` as the safe Runtime-status primary action.
- Changed GUI evidence projection and Compatibility Center normalization to keep ordinary GUI smoke review-only while promoting only safe owner evidence handoffs.
- Changed the Plasma Compatibility Center stub to include `primary_action_id` and `primary_action_kind` in the GUI evidence card contract.

## [0.2.640-rc93] - 2026-07-25

### Added

- Added owner evidence handoff fields for owner-controlled GUI smoke reports, Runtime GUI evidence projections, known-app smoke summaries, KDE GUI evidence cards, and the KDE model bridge.
- Added safe relative owner evidence handoff validation for `runtime/kde-runtime-status-launch-evidence/` paths before exposing desktop-readable owner service-call readiness.

### Changed

- Changed the owner-controlled GUI smoke harness to preserve the owner fixture's reusable evidence handoff after the actual Runtime owner service launch passes.
- Changed the Plasma Compatibility Center stub to document owner evidence handoff readiness while keeping owner service arguments and backend details hidden.

## [0.2.640-rc92] - 2026-07-25

### Added

- Added structured `owner_controlled_runtime_launch_verified` and `owner_managed_copy_verified` fields to known-app GUI smoke evidence summaries and KDE GUI evidence cards.
- Added KDE Center owner-controlled GUI evidence counts for Runtime-owner managed GUI launches and owner-managed executable copy verification.

### Changed

- Changed the Compatibility Center normalization path to preserve owner-controlled GUI verification as structured evidence instead of relying on summary text.
- Changed the Plasma Compatibility Center stub to document the owner-controlled GUI evidence counters and card fields while keeping launch/backend execution disabled.

## [0.2.640-rc91] - 2026-07-25

### Added

- Added owner-controlled GUI evidence projection fields for seed projection, owner service handoff, delegated Wine GUI source, delegated execution/window observation, and owner-managed artifact copy.
- Added Runtime/KDE known-app evidence states for `owner-controlled-gui-qemu-wine-verified` and `validated-owner-controlled-gui-runtime-run`.

### Changed

- Changed GUI smoke evidence consumption to preserve owner-managed external executable copy proof through normalized Compatibility Center evidence.
- Changed Runtime owner and KDE-status launch execution results to surface delegated managed artifact copy evidence without exposing raw backend paths or commands.
- Verified the q4 owner-controlled MessageBox run with `owner_managed_copy_verified=true`, `owner_delegated_managed_artifact_copied=true`, `owner_delegated_controlled_session_window_observed=true`, `x_window_observed=true`, `compatibility_state=owner-controlled-gui-qemu-wine-verified`, and evidence persisted at `/tmp/xnix-run-materials/state/wine-gui-messagebox-owner-0.2.640-rc91-final-evidence.json`.

## [0.2.640-rc90] - 2026-07-25

### Added

- Added Runtime-owner-supplied external executable forwarding through `XNIX_RUNTIME_OWNER_GUI_EXECUTABLE` and `xnix-compat-launch --executable`, allowing owner-controlled launches to copy a Windows GUI executable into the QEMU guest before Wine starts it.
- Added launcher coverage for owner-managed external GUI executable copy delivery while preserving redacted product-facing output.

### Changed

- Changed owner-controlled external GUI smoke delivery to prefer owner-managed copy through the launcher instead of relying on a seed run's pre-copied guest path.
- Verified the q4 owner-controlled MessageBox run with `owner_external_gui_app_delivery=owner-managed-copy`, `executable_copied=true`, `owner_managed_launcher_invoked=true`, `owner_delegated_smoke_passed=true`, `owner_delegated_evidence_source=wine-guest-gui-smoke`, `x_window_observed=true`, and evidence persisted at `/tmp/xnix-run-materials/state/wine-gui-messagebox-owner-0.2.640-rc90.json`.

## [0.2.640-rc89] - 2026-07-25

### Added

- Added Runtime-owner-supplied guest GUI app forwarding through `XNIX_RUNTIME_OWNER_GUEST_GUI_APP` and `xnix-compat-launch --gui-app`, allowing owner-controlled smoke runs to start an already-copied external Windows GUI executable from the managed guest.
- Added owner-controlled GUI smoke plan fields for external GUI app requests while keeping the raw guest GUI app path out of product-facing JSON.
- Added `org.xnix.apps.messagebox` as a development recipe and known-app identity for the q4 `xnix-messagebox-smoke.exe` GUI fixture.

### Changed

- Changed guest-builtin GUI known-app dispatch to use an owner-supplied GUI app path when present, falling back to Wine's built-in `winemine.exe`.
- Verified the q4 owner-controlled external GUI `xnix-messagebox-smoke.exe` run with `executable_copied=true`, `owner_external_gui_app_requested=true`, `owner_managed_launcher_invoked=true`, `owner_delegated_smoke_passed=true`, `owner_delegated_evidence_source=wine-guest-gui-smoke`, `x_window_observed=true`, and evidence persisted at `/tmp/xnix-run-materials/state/wine-gui-messagebox-owner-0.2.640-rc89.json`.

## [0.2.640-rc88] - 2026-07-25

### Added

- Added `scripts/wine_guest_gui_smoke.rb --launch-mode owner-controlled-launch`, which runs a real Mines GUI evidence seed, records the Runtime-status launch owner fixture, invokes `xnix-runtime-owner --service-call ShowRuntimeControlledLaunch`, and verifies that the managed launcher starts `winemine.exe` through the Go-owned Wine GUI smoke lane.
- Added remote q4 owner-controlled GUI smoke planning and execution support, including q4 builds for `xnix-runtime-owner` and `xnix-compat-launch` alongside `xnix-runtime-go`.
- Added constrained `/tmp/xnix-*` remote scratch path support for q4 smoke runs when `/home` is full, while still rejecting unconstrained temporary or host paths.

### Fixed

- Fixed the owner-controlled smoke handoff to call `xnix-runtime-owner` with the positional `evidence-relative-path <relative>` service argument shape expected by the owner service boundary.
- Fixed owner-controlled smoke report parsing for the `runtime-owner-service-call` envelope returned by `xnix-runtime-owner`.
- Verified the q4 owner-controlled `winemine.exe` run with `owner_managed_launcher_invoked=true`, `owner_delegated_smoke_passed=true`, `owner_delegated_evidence_source=wine-guest-gui-smoke`, `x_window_observed=true`, and evidence persisted at `/tmp/xnix-run-materials/state/wine-gui-mines-owner-0.2.640-rc88.json`.

## [0.2.640-rc87] - 2026-07-25

### Added

- Added guest-builtin GUI dispatch support to `xnix-compat-launch` so `org.xnix.apps.mines` can execute through the Go-owned Wine GUI smoke lane after Runtime launch authorization, session-gated review, and controlled execution session consumption.
- Added owner-service forwarding for GUI launch parameters, including guest SSH endpoint, display selection, host X11 observation, and GUI wait duration, while preserving KDE-safe delegated evidence projection.

### Changed

- Changed Runtime-owner controlled launch results to retain `wine-guest-gui-smoke` as the delegated evidence source when a human-authorized Mines launch executes through the managed GUI path.
- Changed the Mines Runtime recipe and known-app catalog version to `0.2.640-rc87`.
- Verified a restricted q4 Xvfb/QEMU/Wine `winemine.exe` run with `x_window_observed=true`, `guest_x11_driver_available=true`, and evidence persisted at `/home/xnix-run-materials/state/wine-gui-mines-0.2.640-rc87.json`.

## [0.2.640-rc86] - 2026-07-25

### Added

- Added `org.xnix.apps.mines` to the Runtime-known Windows app identity catalog as a guest GUI application so Runtime-status launch request previews can recognize the same real GUI app identity used by the recipe registry.
- Added `known-app-runtime-status-launch-owner-fixture-record --gui-smoke-evidence-file FILE` so a passed `gui-smoke-evidence-preview` projection can seed Runtime-owned launch authorization, session, review, and KDE handoff evidence for Mines without exposing paths or enabling desktop/backend launch.

### Changed

- Changed Runtime-status launch delegated evidence to preserve `wine-guest-gui-smoke` as a supported evidence source, allowing real GUI window evidence to drive the same owner-service handoff and KDE action-trigger path as staged launcher smoke while keeping marker/checksum artifact claims separate.

## [0.2.640-rc85] - 2026-07-25

### Added

- Added `org.xnix.apps.mines` as the first backend-neutral Runtime recipe for a real Windows GUI smoke target.
- Added the Mines recipe to the development registry with digest verification so KDE/Compatibility Center pages can consume q4 GUI evidence against the real app identity.

## [0.2.640-rc84] - 2026-07-25

### Changed

- Changed `scripts/remote_wine_guest_gui_smoke.rb` to default to Runtime-only source sync so q4 real GUI smoke runs copy only `VERSION`, `go.mod`, `cmd`, `internal`, `runtime`, `scripts`, and `lib` before rebuilding the Go Runtime.
- Added `--source-sync-mode runtime|full` plus plan evidence for the remote GUI smoke harness so operators can explicitly choose the small development sync or the previous full-checkout sync.

## [0.2.640-rc83] - 2026-07-25

### Added

- Added `gui-smoke-evidence-preview --output FILE` so a passed real Wine guest GUI smoke report can persist the same redacted Runtime/KDE evidence projection that it prints to stdout.
- Changed `scripts/remote_wine_guest_gui_smoke.rb` to automatically project Runtime GUI evidence after a passed q4 Xvfb/QEMU/Wine run while constraining the evidence output path to `/home/xnix-*`.
- Changed Compatibility Center and KDE Center evidence loading so `--known-app-evidence-file` can consume the persisted `gui-smoke-evidence-preview` projection directly.
- Changed the remote GUI evidence defaults to backend-neutral application identifiers and display names so desktop-facing projections pass the existing KDE-safe backend-detail redaction guard.

## [0.2.640-rc82] - 2026-07-25

### Added

- Added `scripts/remote_wine_guest_gui_smoke.rb --remote-executable PATH` so q4 can repeat the real copied Windows GUI `.exe` smoke path through the maintained remote harness while keeping executable paths constrained to `/home/xnix-*`.

### Fixed

- Verified an unattended q4 `xnix-messagebox-smoke.exe` run copied the Windows GUI fixture into the Wine guest and observed an X window with `executable_copied=true`, `guest_x11_driver_available=true`, and `x_window_observed=true`.
- Verified the rc82 GUI report feeds `gui-smoke-evidence-preview` and `kde-center-page-preview`, producing `known_app_gui_evidence_count=1` without enabling desktop launch or backend launch.

## [0.2.640-rc81] - 2026-07-25

### Changed

- Changed the Go-owned Wine guest GUI smoke to poll X window state throughout the observation window instead of checking once after a fixed sleep.
- Changed Wine first-run installer suppression to match BusyBox `ps` output more reliably and to keep suppressing installer prompts while the GUI window is being observed.
- Added GUI window observation attempt counts and guest Wine X11 driver availability checks to the Runtime and Ruby harness reports for diagnosing real Xvfb/QEMU/Wine runs.

### Fixed

- Updated the q4 managed Wine guest run material to the rebuilt X11-driver image and verified an unattended `winemine.exe` GUI smoke pass with `guest_x11_driver_available=true`, `x_window_observed=true`, and the report persisted at `/home/xnix-run-materials/state/wine-gui-winemine-0.2.640-rc81-final.json`.

## [0.2.640-rc80] - 2026-07-25

### Added

- Added KDE Compatibility Center read-model projection for safe real Windows GUI evidence cards from the Runtime KDE Center page payload, including card counts, passed-evidence counts, display fields, Runtime dispatch status, and read-only action state.

### Changed

- Changed the KDE Compatibility Center plasmoid card contract to name the GUI evidence display fields expected from safe `known_app_gui_evidence_cards`.
- Updated the KDE package metadata, Claude collaboration index, and C Runtime core version define to match the Runtime version.

## [0.2.640-rc79] - 2026-07-25

### Added

- Added a KDE Compatibility Center plasmoid GUI evidence card entry that names the `known_app_gui_evidence_count` and `known_app_gui_evidence_cards` read-model fields for real Windows GUI smoke evidence while keeping desktop actions read-only.

### Changed

- Updated the KDE Compatibility Center plasmoid metadata version to match the Runtime version.
- Changed the layout verifier to read the expected version from `VERSION` so package metadata checks stay aligned after each small version bump.
- Updated the Claude collaboration index version to match the Runtime and KDE package version.
- Updated the C Runtime core version define to match the Runtime and KDE package version.

## [0.2.640-rc78] - 2026-07-25

### Added

- Added a Go-owned `gui-smoke-evidence-preview` read model that consumes executed Wine guest GUI smoke reports and projects KDE-safe real-window evidence without exposing report paths, guest paths, backend commands, or raw output.
- Added `--known-app-gui-smoke-report` support to Compatibility Center and KDE Center previews so real GUI QEMU/Wine smoke evidence can be counted and shown as a dedicated GUI evidence card.

## [0.2.640-rc77] - 2026-07-25

### Added

- Added a Go-owned local Windows GUI executable delivery path to `windows-app-guest-wine-gui-smoke`, allowing a validated local `.exe` to be copied into the QEMU guest over loopback SCP before Wine launches it on the observed X display.
- Added a Windows `MessageBoxW` GUI fixture source and `scripts/wine_guest_gui_smoke.rb --executable` pass-through so q4 can run a controlled local GUI `.exe` through the same Runtime-owned path.

## [0.2.640-rc76] - 2026-07-25

### Fixed

- Disabled Wine Mono, Gecko, and menu-builder bootstrap with the Wine-compatible `winemenubuilder.exe=d,mscoree,mshtml=` override, and added Go-owned suppression of only the first-run `control.exe appwiz.cpl install_mono` subprocess so guest GUI smoke can reach `winemine.exe` without interrupting core Wine prefix setup.

## [0.2.640-rc75] - 2026-07-24

### Added

- Added the Go-owned `windows-app-guest-wine-gui-smoke` command for guest Wine GUI app validation, including guest `wineboot`, GUI launch, host `xwininfo` observation, and KDE-safe evidence.

### Changed

- Changed `scripts/wine_guest_gui_smoke.rb` to delegate Wine GUI execution and window observation to the Go Runtime while Ruby remains the Xvfb/QEMU harness.
- Changed `scripts/remote_wine_guest_gui_smoke.rb --execute` to build the current Go Runtime on q4 before invoking the GUI smoke.

## [0.2.640-rc74] - 2026-07-24

### Added

- Added `scripts/remote_wine_guest_gui_smoke.rb`, an execute-gated q4 harness that syncs this checkout, runs the Wine guest GUI smoke remotely, and persists JSON evidence under managed `/home/xnix-*` paths.
- Added plan-only coverage for the remote GUI smoke so the q4 command shape, safe path boundaries, and rebuilt-X11 Wine guest requirement are testable without SSH execution.

## [0.2.640-rc73] - 2026-07-24

### Added

- Added `scripts/wine_guest_gui_smoke.rb`, an execute-gated Xvfb/QEMU/Wine GUI smoke harness for proving a real Wine GUI Windows app creates an X window.
- Added plan-only coverage for the Wine guest GUI smoke so targeted checks can validate safety boundaries without launching QEMU.

### Changed

- Changed the Wine i386 guest Buildroot defconfig to include X.org client libraries, font/image dependencies, and Wine X11 integration prerequisites required for GUI Windows applications.

## [0.2.640-rc72] - 2026-07-24

### Added

- Added `--known-app-matrix-report` to `compatibility-center-preview` and `kde-center-page-preview` so product read models can consume aggregate q4 known-app matrix evidence.
- Added KDE Center page `known_app_matrix_evidence_cards` and matrix evidence counts for real-run known Windows app evidence while keeping launch and backend actions disabled.

### Changed

- Changed known-app smoke evidence normalization to preserve the `remote-known-winapp-matrix-smoke` source as `known-application-matrix-smoke` with `validated-real-runtime-run` card state.

## [0.2.640-rc71] - 2026-07-24

### Added

- Added the Go-owned `known-app-matrix-evidence-preview` Runtime command for consuming aggregate q4 known-app matrix JSON reports.
- Added KDE-safe matrix evidence projection for per-app known Windows app smoke summaries, including checksum, marker, redaction, serial-log, QEMU, and Wine execution counts without exposing remote paths or raw output.

### Changed

- Changed the remote known-app matrix smoke aggregate to preserve per-app `checksum_verified` evidence for downstream Runtime consumers.

## [0.2.640-rc70] - 2026-07-24

### Changed

- Changed `scripts/remote_known_winapp_matrix_smoke.rb --execute` to persist the aggregate known-app matrix JSON report to a managed remote `/home/xnix-*` state path by default.
- Added `--matrix-report-output` and dry-run coverage so operators and future desktop/release consumers can locate the aggregate matrix evidence without scraping stdout.

## [0.2.640-rc69] - 2026-07-24

### Added

- Added `scripts/remote_known_winapp_matrix_smoke.rb`, an execute-gated q4 matrix harness that syncs and builds the Go Runtime once, then runs the default `7zr` and `busybox-w32` known apps sequentially through Runtime-started QEMU/Wine.
- Added matrix dry-run and script-shape coverage proving per-app report paths, serial-log paths, aggregate pass counts, redacted output, loopback networking, and protected Claude document exclusion.

## [0.2.640-rc68] - 2026-07-24

### Added

- Added `busybox-w32` as the second pinned real Windows known app, including source URL, executable name, PE32 architecture, SHA-256 checksum, default `--help` smoke argument, and `BusyBox` marker.

### Changed

- Updated known-app downloads to send a Runtime User-Agent so official sources that block default Go HTTP clients can still be fetched through the managed catalog path.

## [0.2.640-rc67] - 2026-07-24

### Changed

- Changed `scripts/remote_known_winapp_guest_wine_smoke.rb` to default to a lightweight `runtime` source sync mode that copies only `go.mod`, `cmd/`, `internal/`, and `runtime/` into a versioned q4 source tree before building the Go Runtime.
- Added an explicit `--source-sync-mode full` fallback for operator-controlled whole-checkout q4 smoke syncs.

## [0.2.640-rc66] - 2026-07-24

### Added

- Added `--report-output` to `windows-known-app-run` so operator-controlled real Windows app runs can persist the same redacted JSON evidence emitted to stdout.
- Added `scripts/remote_known_winapp_guest_wine_smoke.rb`, an execute-gated q4 harness that syncs source, builds the Go Runtime with the managed remote Go toolchain, runs the known Windows app through Runtime-started QEMU/Wine, and records report plus serial-log evidence under the managed q4 run materials root.

## [0.2.640-rc65] - 2026-07-24

### Added

- Added `--redact-output` support to `windows-known-app-run`, including guest-wine runs, so real Windows app stdout and stderr can be omitted while preserving marker, byte-count, line-count, and KDE-safe output summary evidence.

### Changed

- Updated the known Windows app QEMU/Wine smoke harness to request redacted Runtime output by default.

## [0.2.640-rc64] - 2026-07-24

### Fixed

- Fixed Go-owned QEMU guest run results so `qemu_serial_log_written` reflects the serial log written during explicit guest shutdown instead of being lost by deferred cleanup.

## [0.2.640-rc63] - 2026-07-24

### Fixed

- Fixed Go-started known app QEMU guest runs so a `localhost` request is normalized to `127.0.0.1` for both QEMU host forwarding and the subsequent SSH/SCP Wine execution path.

## [0.2.640-rc62] - 2026-07-24

### Added

- Added `--port auto` support for `windows-known-app-run --backend guest-wine --start-qemu`, allowing the Go Runtime to allocate an available loopback SSH port for each QEMU guest run.
- Added safe guest port evidence fields (`guest_host`, `guest_port`, and `guest_port_auto`) to Go-owned known app guest-wine run results.

### Changed

- Updated the known Windows app QEMU/Wine smoke harness to use the Runtime-owned automatic loopback port instead of a fixed host port.

## [0.2.640-rc61] - 2026-07-24

### Added

- Added Go-owned QEMU guest lifecycle support for `windows-known-app-run --backend guest-wine --start-qemu`, including explicit QEMU binary, kernel, memory, CPU, boot-timeout, and serial-log controls.
- Added safe guest-start evidence fields for known app runs: `guest_start_attempted`, `guest_started`, `guest_start_mode`, `qemu_serial_log_written`, and `raw_qemu_path_exposed=false`.

### Changed

- Updated `scripts/known_winapp_guest_wine_smoke.rb` so Ruby remains the test harness while the Go Runtime starts, waits for, uses, and stops the QEMU Wine guest.
- Made the Go QEMU start path preflight the known app artifact before starting the guest, so missing or checksum-mismatched artifacts skip safely without launching a VM.

### Fixed

- Fixed Runtime tool resolution so bare command names such as `qemu-system-i386` are resolved from `PATH` instead of being misread as workspace-relative files.

## [0.2.640-rc60] - 2026-07-24

### Changed

- Routed `scripts/known_winapp_guest_wine_smoke.rb` through the Go-owned `windows-known-app-run --backend guest-wine` entrypoint instead of the older dispatch-smoke command.
- Added a targeted script guard proving the formal known Windows app QEMU/Wine smoke path now validates the unified `guest-wine` backend while preserving the explicit QEMU harness boundary.

### Fixed

- Fixed the full-suite launcher bundle CLI test to store runner settings in the smoke profile fixture instead of passing obsolete runner flags to `windows-app-launcher-bundle-record`.
- Fixed the mainline integration review classification for the Windows app smoke runbook and known-app guest smoke script guard.

## [0.2.640-rc59] - 2026-07-24

### Added

- Added `windows-known-app-run`, a Go-owned known app execution entrypoint with explicit `local` and `guest-wine` backends.
- Added unified run evidence for backend selection, backend readiness, checksum verification, launch attempt, runner availability, marker observation, local payloads, guest payloads, loopback networking, QEMU requirement, and host-boundary flags.
- Added targeted tests proving the local backend reuses prepare-and-launch and the guest-wine backend reaches the loopback Wine guest smoke path through fake SSH/SCP fixtures.

## [0.2.640-rc58] - 2026-07-24

### Added

- Added `windows-known-app-prepare-and-launch-profile`, a Go-owned one-step known app path that prepares or verifies the artifact, materializes the launch profile, and immediately calls `windows-app-launch-profile`.
- Added combined prepare and launch evidence with safe `prepare_payload`, `launch_payload`, launch-attempt, runner, PE architecture, redaction, and host-boundary fields.
- Added targeted tests proving missing artifacts do not launch and verified artifacts can launch through the redacted Runtime path when a runner is ready.

## [0.2.640-rc57] - 2026-07-24

### Added

- Added explicit runner settings to known-app launch profile materialization and one-step preparation with `--runner`, `--runner-bottle`, repeated `--runner-arg`, and `--skip-bootstrap`.
- Added path-safe runner summary evidence (`runner_configured`, `runner_bottle_configured`, `runner_argument_count`, and `skip_bootstrap`) while keeping raw runner paths, bottle names, and runner argv values out of product-facing JSON reports.

## [0.2.640-rc56] - 2026-07-24

### Added

- Added `windows-known-app-prepare-launch-profile`, combining offline known-app cache checks, optional explicit downloads, checksum verification, and launch-profile materialization.
- Added targeted tests for offline missing-artifact skips and explicit mock-download preparation into profile and launcher bundle outputs.

## [0.2.640-rc55] - 2026-07-24

### Added

- Added `windows-known-app-launch-profile-materialize`, connecting verified known Windows app artifacts to reusable launch profiles and managed `launch` mode launcher bundles.
- Added tests proving missing known app artifacts skip safely and verified artifacts can write profiles and launcher bundles without exposing raw paths.

## [0.2.640-rc54] - 2026-07-24

### Added

- Added the Go-owned `windows-app-launch-profile` entrypoint, which preflights a reusable Windows app profile before redacted Runtime execution.
- Added `--launcher-mode launch` to managed launcher bundles and `scripts/winapp_smoke.rb` so desktop launchers can call the product-style launch profile path.

## [0.2.640-rc53] - 2026-07-24

### Added

- Added `execute` and `preflight` modes to managed Windows app launcher bundles.
- Added `--launcher-mode execute|preflight` to `windows-app-launcher-bundle-record` and `scripts/winapp_smoke.rb`, enabling desktop launchers that validate a real Windows PE profile without starting Wine or the app.

## [0.2.640-rc52] - 2026-07-24

### Added

- Added runtime argv support to managed Windows app launcher bundles, enabling development launchers such as `go run ./cmd/xnix-runtime-go` without shell-command string parsing.
- Added path-safe `runtime_argument_count` and `raw_runtime_argv_exposed=false` evidence to launcher bundle records and reports.

## [0.2.640-rc51] - 2026-07-24

### Added

- Added Go-owned `windows-app-launcher-bundle-record` for writing a managed launcher script, `.desktop` file, and receipt from a reusable Windows app smoke profile.
- Added `scripts/winapp_smoke.rb --write-launcher-bundle` with path-safe launcher bundle evidence for future KDE click-to-launch wiring.

## [0.2.640-rc50] - 2026-07-24

### Added

- Added the Go-owned `windows-app-smoke-profile-render` command for generating reusable local Windows app smoke profile JSON.
- Added `scripts/winapp_smoke.rb --write-profile <path>` so operators can save repeatable real app settings without exposing the written profile path in reports.

## [0.2.640-rc49] - 2026-07-24

### Added

- Added `--stage-app-dir` to local Windows app smoke so the Go Runtime can copy an executable's application directory into `state_root/app-workspace` before launch.
- Added `stage_app_dir` profile support plus `application_workspace_mode`, `application_staged`, and staged file/byte-count evidence to Go Runtime, Ruby JSON, and Markdown reports.

## [0.2.640-rc48] - 2026-07-24

### Added

- Added `--skip-bootstrap` to local Windows app smoke so operators can bypass companion `wineboot` during real executable diagnostics.
- Added `skip_bootstrap` profile support and `wine_bootstrap_skipped` evidence to Go Runtime, preflight, Ruby JSON, and Markdown reports.

## [0.2.640-rc47] - 2026-07-24

### Added

- Added architecture-scoped local Wine prefixes under the managed state root for direct Windows app smoke execution.
- Added `wine_prefix_mode` and `wine_prefix_prepared` evidence to Go Runtime smoke, profile preflight, and Ruby JSON/Markdown smoke reports without exposing prefix paths.

## [0.2.640-rc46] - 2026-07-24

### Added

- Added architecture-derived Wine prefix selection so x86_64 executables use `WINEARCH=win64` and x86 executables use `WINEARCH=win32`.
- Added `wine_architecture` evidence to Go Runtime smoke, profile preflight, and Ruby JSON/Markdown smoke reports.

## [0.2.640-rc45] - 2026-07-24

### Added

- Added PE machine architecture parsing for direct and profile-backed Windows app smoke inputs.
- Added `executable_architecture` and `executable_architecture_supported` evidence to Go Runtime and Ruby smoke reports, blocking unsupported non-x86 Windows executables before runner execution.

## [0.2.640-rc44] - 2026-07-24

### Added

- Added Windows `MZ` executable signature validation to direct `windows-app-run-smoke` runs before runner resolution.
- Added report-level `executable_format` and `windows_executable_signature_observed` evidence to direct Windows app smoke JSON and Markdown reports.

## [0.2.640-rc43] - 2026-07-24

### Added

- Added Windows `MZ` executable signature validation to profile preflight before readiness can be reported.
- Added `executable_format` and `windows_executable_signature_observed` evidence to profile preflight reports.

## [0.2.640-rc42] - 2026-07-24

### Added

- Added `scripts/winapp_smoke.rb --preflight-only --profile <json>` for safe readiness reports without launching the Windows app.
- Embedded Go-owned profile preflight payloads into profile-backed `scripts/winapp_smoke.rb` reports before any smoke execution.

## [0.2.640-rc41] - 2026-07-24

### Added

- Added `windows-app-smoke-profile-preflight --profile <json>` to validate reusable real Windows app smoke profiles before execution.
- Added profile preflight tests proving safe readiness output, runner diagnostics integration, path redaction, and no Wine/Docker/QEMU execution.

## [0.2.640-rc40] - 2026-07-24

### Added

- Added a copy-ready Windows app smoke profile template for repeatable local real `.exe` smoke attempts.
- Added a Windows app smoke profile runbook and template checks covering safe entrypoints, success modes, working-directory reporting, and redaction boundaries.

## [0.2.640-rc39] - 2026-07-24

### Added

- Added `windows-app-run-smoke --profile <json>` for reusable local real Windows app smoke settings owned by the Go Runtime.
- Added `scripts/winapp_smoke.rb --profile <json>` report support while preserving path redaction and safe profile evidence fields.

## [0.2.640-rc38] - 2026-07-24

### Added

- Defaulted local Windows app smoke execution to the executable directory so sidecar DLLs, config files, and resources are discoverable.
- Added `--working-dir` support to the Go-owned smoke command and `scripts/winapp_smoke.rb`, reporting only `working_directory_mode` without exposing raw paths.

## [0.2.640-rc37] - 2026-07-24

### Added

- Added `--success-mode startup-window` to the Go-owned local Windows app smoke command for GUI-style apps that remain running through the configured startup timeout.
- Added `startup_window_observed` reporting and startup-window forwarding to `scripts/winapp_smoke.rb`.

## [0.2.640-rc36] - 2026-07-24

### Added

- Added `--success-mode marker|exit-code` to the Go-owned local Windows app smoke command, preserving marker validation by default while allowing explicit zero-exit real app evidence.
- Added success mode forwarding and `success_mode` reporting to `scripts/winapp_smoke.rb`.

## [0.2.640-rc35] - 2026-07-24

### Added

- Added local `--runner-bottle NAME` support to the Go-owned Windows app smoke command, expanding it to runner bottle arguments without exposing the raw bottle name.
- Added `--runner-bottle` forwarding to `scripts/winapp_smoke.rb` and updated safe command hints to prefer the simpler bottle selector.

## [0.2.640-rc34] - 2026-07-24

### Fixed

- Passed local Windows app smoke `--runner-arg` values into companion `wineboot` bootstrap before `--init`.
- Moved `runner_argument_count` reporting earlier so bootstrap failures still preserve safe runner-argument evidence without exposing raw values.

## [0.2.640-rc33] - 2026-07-24

### Added

- Added repeated `--runner-arg` support to the Go-owned local Windows app smoke command so compatibility runners can receive pre-executable arguments.
- Added `--runner-arg` forwarding and `runner_argument_count` reporting to `scripts/winapp_smoke.rb` without exposing raw runner argument values.

## [0.2.640-rc32] - 2026-07-24

### Added

- Added common macOS CrossOver system and user bundle Wine runner candidates to Go Runtime Windows app diagnostics.
- Added common Whisky user-library Wine runner candidates so local real `.exe` smoke attempts can discover an existing Whisky-managed Wine library without host mutation.

## [0.2.640-rc31] - 2026-07-24

### Added

- Added safe `runner_command_hints` to the Go-owned Windows app runner diagnostics result.
- Carried runner command hints into `scripts/winapp_smoke.rb` JSON and Markdown reports so real `.exe` smoke attempts can be rerun with `XNIX_WINDOWS_RUNNER` or `--runner` placeholder commands without exposing host paths.

## [0.2.640-rc30] - 2026-07-24

### Changed

- Turned `scripts/winapp_container_smoke.rb` into a compatibility wrapper around `scripts/winapp_smoke.rb --backend container`.
- Preserved the legacy container smoke environment knobs while making the unified smoke report the single implementation path for future restricted-container Windows app PASS attempts.

## [0.2.640-rc29] - 2026-07-24

### Added

- Added an explicit `--backend local|container` switch to `scripts/winapp_smoke.rb`, keeping the default local backend unchanged while allowing the same report path to opt into the restricted container smoke.
- Added container backend report fields for the embedded container payload, image availability, Docker execution evidence, and Wine bootstrap evidence.

## [0.2.640-rc28] - 2026-07-24

### Added

- Added managed Wine prefix bootstrap support to local `windows-app-run-smoke` when a companion `wineboot` is available beside the selected runner.
- Added bootstrap attempted, succeeded, and exit-code evidence to Runtime and `scripts/winapp_smoke.rb` reports while preserving custom runner compatibility when `wineboot` is unavailable.

## [0.2.640-rc27] - 2026-07-24

### Added

- Embedded Go-owned runner diagnostics into `scripts/winapp_smoke.rb` JSON and Markdown reports before invoking the real Windows app smoke.
- Added script coverage proving smoke reports carry diagnostics status, candidate counts, next-action evidence, and explicit-runner diagnostics without leaking executable or runner paths.

## [0.2.640-rc26] - 2026-07-24

### Added

- Added `XNIX_WINDOWS_RUNNER` as a Go Runtime-managed local Wine runner configuration path for both `windows-app-run-smoke` and `windows-app-runner-diagnostics`.
- Added tests proving the environment-configured runner path is honored without exposing raw host paths in diagnostics output.

## [0.2.640-rc25] - 2026-07-24

### Added

- Added the Go-owned `windows-app-runner-diagnostics` command for checking local Wine-compatible runner availability before attempting a real Windows app smoke.
- Added diagnostics evidence for explicit runner paths, discovered runner candidates, safe next actions, and closed Docker/QEMU/Colima/network/package-manager gates without exposing raw host paths.

## [0.2.640-rc24] - 2026-07-24

### Fixed

- Aligned local Windows app smoke execution with the Runtime-managed Wine environment by setting `WINEARCH=win64`, `WINEDEBUG=-all`, and `WINEDLLOVERRIDES=winemenubuilder.exe=d,mscoree=d,mshtml=d` alongside the isolated `WINEPREFIX`.
- Added Runtime coverage proving the local smoke runner receives the managed Wine environment before marker observation.

## [0.2.640-rc23] - 2026-07-24

### Fixed

- Broadened Windows app smoke runner discovery to try `wine`, `wine64`, Homebrew Wine paths, and common macOS Wine app bundle runner paths before reporting a runner-unavailable SKIP.
- Added Runtime coverage proving the smoke can discover `wine64` from `PATH` and run the fixture through the same marker-observation path.

## [0.2.640-rc22] - 2026-07-24

### Fixed

- Made `scripts/winapp_smoke.rb` clear its managed fixture executable output before rebuilding so stale local smoke artifacts cannot block repeated real Windows app smoke attempts.
- Added regression coverage proving the fixture smoke can recover from a stale managed `hello.exe` output while preserving custom executable behavior.

## [0.2.640-rc21] - 2026-07-24

### Added

- Added `--expected-marker` support to the Go-owned `windows-app-run-smoke` CLI.
- Added `--exe`, `--runner`, `--arg`, `--expected-marker`, `--timeout`, and `--state-root` support to `scripts/winapp_smoke.rb` so existing user-supplied Windows executables can use the same redacted evidence path as the fixture smoke.
- Added tests proving custom executable mode skips fixture builds, forwards runner/app arguments, preserves safe executable basenames, and avoids leaking executable or runner paths.

### Fixed

- Kept the default fixture smoke unchanged while making user-supplied executable smoke evidence path-safe and reportable.

## [0.2.640-rc20] - 2026-07-24

### Added

- Added JSON and Markdown report formats to `scripts/winapp_smoke.rb`.
- Added default redacted output reporting for non-text Windows app smoke reports, preserving marker and safe output-summary evidence without raw stdout or stderr.
- Added `test/test_winapp_smoke_script.rb` coverage for the report modes using a fake Go runner.

### Fixed

- Kept `windows-app-run-smoke` skip evidence consistent with `--redact-output` when a compatibility runner is unavailable.

## [0.2.640-rc19] - 2026-07-24

### Added

- Added `--redact-output` support to `windows-app-run-smoke` so real Windows app runner evidence can preserve marker, byte, and line-count diagnostics without exposing raw stdout or stderr.
- Added Go Runtime runner and CLI tests for KDE-safe redacted output summaries.

### Fixed

- Kept the local Windows app smoke path useful for developer marker checks by preserving raw output by default while giving product-facing consumers an explicit redacted mode.

## [0.2.640-rc18] - 2026-07-24

### Added

- Added optional `--desktop-trigger-request-preflight-smoke PATH` fixture support to `scripts/release_evidence_index.rb`.
- Added a `desktop-trigger-request-preflight-smoke` release evidence claim that classifies supplied preflight smoke evidence as implemented, blocked, or skipped.
- Added `docs/claude-code-active-work-order.md` as the short current Claude Code handoff for C9W5/C9W6.
- Added release evidence tests for supplied, missing, failed, and malformed preflight smoke evidence.

### Fixed

- Kept release evidence indexing from running the preflight smoke by default while still blocking failed or malformed supplied preflight smoke evidence.

## [0.2.640-rc17] - 2026-07-24

### Added

- Added `--format json` support to `scripts/desktop_trigger_request_preflight_smoke.rb` so targeted preflight smoke evidence can be consumed by read-only release tooling.
- Added optional `--desktop-trigger-request-preflight-smoke PATH` fixture support to `scripts/merge_readiness_packet.rb`.
- Added merge readiness tests for supplied, missing, failed, and malformed desktop-trigger request preflight smoke evidence.

### Fixed

- Kept merge readiness from running the preflight smoke by default, while making failed or malformed supplied preflight smoke evidence a release-only blocker instead of a merge blocker.

## [0.2.640-rc16] - 2026-07-24

### Added

- Added `scripts/desktop_trigger_request_preflight_smoke.rb`, a lightweight targeted smoke that records safe Runtime-status handoff evidence and invokes the Go-owned `desktop-trigger-request-preflight-preview`.
- Added container and script tests for the preflight smoke, including restricted container command coverage and redaction/no-side-effect assertions.
- Added `docs/claude-code-next-dispatch-brief.md` as a short copy-first handoff for the next Claude Code task batch.

### Fixed

- Kept the desktop-trigger request preflight automation fail-closed before formal promotion while verifying the promoted fixture path becomes `ready-for-operator-request` without dispatching service calls, calling D-Bus, launching desktop actions, starting backends, writing Runtime/KDE state, or mutating the host.

## [0.2.640-rc15] - 2026-07-24

### Added

- Added Go-owned `desktop-trigger-request-preflight-preview` to consume service-call materialization for the future real desktop-triggered `ShowRuntimeControlledLaunch` request lane.
- Added owner and CLI tests for ready, missing-promotion, unsafe-envelope, stale-evidence, malformed, no-dispatch, and redaction behavior.

### Fixed

- Kept post-release desktop-trigger request preparation fail-closed with `blocked-missing-promotion` until formal full checkpoint promotion is observed, while hiding owner service call arguments from KDE-facing output.

## [0.2.640-rc14] - 2026-07-24

### Added

- Added `docs/post-checkpoint-promotion-checklist-0640.md`, the human-operator checklist for promoting the active `v0.2.640` release candidate only after full smoke passes and the promotion packet allows the release.
- Added layout checks that keep the promotion checklist tied to full smoke, promotion packet, release evidence index, merge readiness, protected Claude file review, rollback notes, and formal tag commands.

### Fixed

- Kept formal `v0.2.640` promotion explicitly operator-owned and blocked until `ruby scripts/full_smoke.rb` passes with current PASS evidence.

## [0.2.640-rc13] - 2026-07-24

### Changed

- Made `scripts/release_evidence_index.rb` consume the full checkpoint promotion packet as part of release evidence indexing.
- Updated `product-image-qemu-acceptance` so historical product smoke evidence is not treated as current formal release readiness unless the promotion packet allows the release.
- Added a `full-checkpoint-promotion` claim with promotion decision, full-smoke state, formal readiness, and blocker evidence.

### Fixed

- Kept old authorized product image smoke evidence visible while blocking current `v0.2.640` product image acceptance when `full_checkpoint_promotion_packet` reports incomplete, missing, failed, malformed, or version-mismatched evidence.

## [0.2.640-rc12] - 2026-07-24

### Changed

- Made `scripts/merge_readiness_packet.rb` consume the full checkpoint promotion packet as a first-class release gate.
- Added `--full-checkpoint-promotion` fixture support and exposed `full_checkpoint_promotion_status` in merge readiness output.
- Updated merge readiness tests for denied and passing promotion packet fixtures while keeping promotion failures release-only rather than merge-blocking.

### Fixed

- Kept release readiness dependent on the promotion packet's `promotion_allowed` and `formal_release_ready` values instead of duplicating full-smoke promotion logic inside merge readiness.

## [0.2.640-rc11] - 2026-07-24

### Added

- Added `scripts/full_checkpoint_promotion_packet.rb`, a read-only release-promotion evidence packet that reads existing full-smoke JSON and Markdown reports and decides whether `v0.2.640` can be promoted.
- Added targeted tests for missing, malformed, failed, incomplete, passing, version-mismatched, Markdown, redaction, and no-side-effect promotion packet behavior.
- Added `docs/claude-code-next-mainline-work-pack.md` as the next copy-first Claude Code task pack for C8W11-C8W14 and post-release C9W1 work.

### Changed

- Updated layout and mainline integration review gates so the full checkpoint promotion packet is classified with the product image and QEMU acceptance lane.

### Fixed

- Kept formal release readiness false unless an existing valid full-smoke PASS report matches the active `v0.2.640` release-candidate line.

## [0.2.640-rc10] - 2026-07-24

### Changed

- Routed the private D-Bus controlled-launch fixture and C smoke adapter through `desktop-trigger-service-call-materialization-preview`, matching the staged launcher dispatch smoke's Go-owned service-call materialization path.
- Renamed the D-Bus fixture's exposed Go metadata payload from owner-trigger evidence to service-call materialization evidence while preserving the owner service-call envelope payload.

### Fixed

- Removed the D-Bus controlled-launch fixture's direct dependency on `known-app-runtime-status-launch-owner-trigger-preview` for owner service call arguments, keeping the lower-level trigger consumed inside the Go materialization packet instead.

## [0.2.640-rc9] - 2026-07-24

### Added

- Added an explicit human-authorized smoke candidate path to `desktop-trigger-service-call-materialization-preview`, allowing evidence-only `ShowRuntimeControlledLaunch` owner service call arguments to be materialized without claiming formal full-checkpoint promotion.
- Added Go and CLI tests for the human-authorized smoke candidate path, checkpoint-promotion honesty, release-readiness non-claiming, redaction, and disabled side effects.
- Added `docs/claude-code-immediate-mainline-handoff.md` as the current copy-first Claude Code handoff for the next bounded C8W9-C8W11 mainline tasks.

### Changed

- Routed `scripts/staged_launcher_dispatch_smoke.rb` through `desktop-trigger-service-call-materialization-preview` before invoking `xnix-runtime-owner`, so the real staged launcher path consumes the same Go-owned service-call materialization gate that future desktop-triggered launch work will rely on.
- Updated layout and smoke-script unit gates to require the staged smoke materialization path, the explicit human-authorized smoke flag, and release-promotion non-claim fields.

### Fixed

- Avoided the full-checkpoint deadlock where the staged smoke needed materialized owner service arguments but materialization previously required full-checkpoint promotion before the smoke could run.
- Fixed desktop-trigger dry-run review classification so malformed evidence handles are not masked as missing evidence when the full-checkpoint gate is otherwise ready.

## [0.2.640-rc8] - 2026-07-24

### Added

- Added `desktop-trigger-service-call-materialization-preview`, a Go-owned materialization packet that consumes the accepted desktop-trigger dry-run review and the Runtime-status owner trigger to emit evidence-only `ShowRuntimeControlledLaunch` owner service call arguments for a future human-authorized staged launch smoke.
- Added targeted Go and CLI tests for ready materialization, blocked full-checkpoint dependency, unsafe envelope rejection without emitted service call arguments, malformed evidence, stale evidence, redaction, and no-side-effect behavior.

### Changed

- Updated layout and mainline review gates so the desktop-trigger service call materialization lane is classified with the Runtime owner read boundary.

### Fixed

- Kept owner service dispatch, D-Bus calls, request-object writes, permission grants, Runtime state writes, KDE configuration writes, desktop launch, backend launch, network access, and host mutation disabled while still exposing only safe evidence-relative-path service call arguments for operator-owned smoke execution.

## [0.2.640-rc7] - 2026-07-24

### Added

- Added `desktop-trigger-dry-run-request-review-preview`, a Go-owned review packet that combines the Runtime owner launch envelope guard, KDE controlled-launch action surface audit, managed launcher acceptance report, Runtime-status evidence digest state, and the full-checkpoint dependency before any desktop-triggered staged launch attempt.
- Added targeted Go and CLI tests for accepted-review, missing full checkpoint, unsafe envelope, unsafe KDE action surface, malformed evidence, stale evidence, redaction, and no-side-effect behavior.
- Added `docs/claude-code-next-delegation-pack.md` as a copy-first Claude Code handoff pack for the next bounded tasks.

### Changed

- Updated layout and mainline review gates so the desktop-trigger dry-run request review is classified with the Runtime owner read boundary.
- Updated current Claude Code dispatch guidance to point at the new short delegation pack.

### Fixed

- Kept the desktop-trigger request review fail-closed without dispatching service calls, calling D-Bus, writing request objects, creating permission grants, writing Runtime state, writing KDE configuration, enabling desktop launch, starting a backend, requiring network access, or mutating the host root.

## [0.2.640-rc6] - 2026-07-24

### Added

- Added `managed-launcher-acceptance-report-preview`, a Go-owned user-safe acceptance report that joins Runtime-status launch evidence, known-app identity, artifact digest evidence, launch authorization, session-gated review, controlled execution session, managed launcher request readiness, guest smoke evidence, KDE-safe projection state, and the full-checkpoint dependency.
- Added `docs/claude-code-current-execution-queue.md` as the current copy-first Claude Code execution queue for `C8W5`, `C8W6`, and `C8W7`.

### Changed

- Updated Claude Code dispatch guidance so the short current handoff points at the new execution queue.
- Updated layout and mainline review gates so the managed launcher acceptance report is classified with the KDE-first Runtime-owned evidence lane.

### Fixed

- Kept the acceptance report fail-closed for missing, malformed, stale, incomplete, or not-yet-promoted evidence while preserving disabled D-Bus calls, desktop launch, backend launch, Runtime writes, KDE configuration writes, network access, and host mutation.

## [0.2.640-rc5] - 2026-07-24

### Added

- Added `kde-controlled-launch-action-surface-audit-preview`, a Go-owned audit that validates the KDE controlled-launch desktop action metadata before any desktop-triggered smoke can be attempted.

### Changed

- The audit checks action identity, public D-Bus method, evidence-only argument shape, missing or malformed metadata, unsafe owner-only launch inputs, backend terms, redaction, and no-side-effect behavior while keeping D-Bus calls, KDE configuration writes, Runtime writes, desktop launch, backend launch, network access, and host mutation disabled.
- Added targeted Go and CLI tests that audit the repository KDE action metadata and block unsafe owner arguments, backend terms, malformed metadata, missing evidence handles, and unsafe evidence handles.

## [0.2.640-rc4] - 2026-07-24

### Added

- Added `owner-service-launch-envelope-guard-preview`, a Go-owned fail-closed Runtime owner guard that validates future desktop-triggered launch envelopes before service dispatch.

### Changed

- The guard accepts only safe evidence handles from KDE, binds them to Runtime-status launch evidence, rejects owner-only state/cache/launcher/timeout/executable/backend/receipt/session/dispatch inputs, and keeps service dispatch, D-Bus ownership, writes, backend launch, network access, and host mutation disabled.
- Added targeted Go and CLI tests for accepted, missing-evidence, mismatched-route, stale, replay, owner-argument, malformed, redaction, and no-side-effect behavior.
- Added `docs/claude-code-live-delegation-brief.md` as the current one-task-at-a-time Claude Code delegation brief.

## [0.2.640-rc3] - 2026-07-24

### Added

- Added `desktop-trigger-staged-invocation-readiness-preview`, a Go-owned read-only packet that combines KDE controlled-launch action metadata, the D-Bus smoke plan, Runtime-status handoff evidence, owner trigger state, managed launcher request readiness, known-app artifact evidence, guest smoke evidence, and the full-checkpoint release gate.

### Changed

- Added stale-digest, missing-evidence, malformed-evidence, missing-artifact, blocked guest-smoke, redaction, no-side-effect, and CLI tests for the desktop-trigger staged invocation readiness packet.
- Added `docs/claude-code-next-actions-handoff.md` as the short bounded Claude Code dispatch sheet for the next implementation branches.

## [0.2.640-rc2] - 2026-07-24

### Changed

- Added structured full-smoke failure classification for Docker daemon blockers, Docker Hub EOF failures, missing local base images, Buildroot build failures, QEMU serial failures, known Windows app smoke failures, fixture Windows app smoke failures, and KDE action fixture failures.
- Updated `scripts/full_smoke.rb` so failed runs write JSON and Markdown reports before exiting nonzero, preserving the safe retry command and redacted failure summary without claiming formal release readiness from partial smoke evidence.
- Extended full-smoke report tests to cover the KDE action smoke lane and failed-report release readiness guards.

## [0.2.640-rc1] - 2026-07-24

### Changed

- Promoted `kde-controlled-launch-action-dbus-fixture-smoke` into the formal full-smoke gate candidate so v0.2.640 requires the KDE controlled-launch action D-Bus fixture lane to PASS once the external Docker pull blocker is resolved.
- Extended the constrained container command model and tests with explicit D-Bus fixture execution environment wiring for the KDE action smoke while preserving no networking, no privileged container, no Docker socket, managed cache volume use, and the controlled-launch scratch tmpfs.
- Extended the full-smoke report to record KDE action smoke inclusion and pass status.

### Blocked

- Attempted the v0.2.640 full smoke, but Colima Docker could not pull `debian:bookworm-slim` from Docker Hub; Docker daemon requests to Docker auth/layer endpoints failed with EOF before project build/test steps could run.

## [0.2.639] - 2026-07-24

### Changed

- Extended `kde-controlled-launch-session-bus-smoke-plan-preview` to expose the D-Bus controlled-launch owner fixture command, restricted container command, and fixture PASS/SKIP markers.
- Updated `kde_controlled_launch_action_smoke.rb` to validate both the private session-bus smoke lane and the D-Bus fixture lane from Go-owned metadata, with a separate explicit execution opt-in for the D-Bus fixture path.

## [0.2.638] - 2026-07-24

### Added

- Added `kde_controlled_launch_action_smoke.rb`, a restricted KDE controlled-launch action smoke harness that consumes the Go-owned session-bus smoke plan before any optional execution.
- Added the `kde-controlled-launch-action-smoke` container command and targeted script guard.

### Changed

- Updated the KDE controlled-launch lane documentation and layout guards to require plan-consuming smoke orchestration, explicit execution opt-in, evidence-only KDE forwarding, hidden owner-service arguments, disabled KDE state-root access, disabled receipt reconstruction, and closed host/container gates.

## [0.2.637] - 2026-07-24

### Added

- Added `kde-controlled-launch-session-bus-smoke-plan-preview`, a Go-owned KDE controlled-launch smoke plan that links the action stub to the restricted private session-bus D-Bus smoke chain without executing it.

### Changed

- Updated the KDE controlled-launch action stub, targeted tests, and layout guards to require the restricted smoke plan command while preserving evidence-only KDE forwarding, hidden owner-service arguments, disabled KDE state-root access, disabled receipt reconstruction, and no execution start from the preview.

## [0.2.636] - 2026-07-24

### Added

- Added `kde-controlled-launch-action-preview`, a Go-owned KDE controlled-launch action preview that consumes verified Runtime-status launch evidence and exposes only an evidence-relative-path handle to KDE.
- Added `kde/actions/xnix-runtime-status-controlled-launch.desktop`, a declarative KDE action stub for the public `org.xnix.Compatibility1.ShowRuntimeControlledLaunch` D-Bus route.

### Changed

- Updated layout and targeted guards to require the KDE controlled-launch action stub, safe Go CLI output, hidden owner-service arguments, disabled KDE state-root access, disabled receipt reconstruction, and no execution start from KDE.

## [0.2.635] - 2026-07-24

### Changed

- Updated the C D-Bus smoke adapter to call `known-app-runtime-status-launch-owner-trigger-preview` and consume Go-provided owner-service call arguments before dispatching `ShowRuntimeControlledLaunch`.
- Extended the D-Bus controlled-launch fixture smoke to parse `go_owner_trigger_json` and verify trigger digest parity, owner-service args, evidence-only KDE forwarding, and disabled unsafe gates.
- Added Claude's next Runtime work handoff document with prioritized implementation packages for the remaining D-Bus/KDE controlled-launch lane.

## [0.2.634] - 2026-07-24

### Changed

- Updated `dbus_controlled_launch_owner_fixture_smoke.rb` to consume `known-app-runtime-status-launch-owner-trigger-preview` before invoking the public D-Bus controlled-launch method.
- Added D-Bus smoke guards for the Go trigger preview runtime/read methods, evidence digest parity, owner-service CLI args, and disabled unsafe gates.
- Extended static layout checks so the D-Bus fixture lane cannot drift back to a fixture-only or Ruby-reconstructed trigger path.

## [0.2.633] - 2026-07-24

### Added

- Added `known-app-runtime-status-launch-owner-trigger-preview`, a Go-owned Runtime-status launch owner trigger preview that reads verified handoff evidence and emits desktop and owner-service trigger metadata.

### Changed

- Updated the staged launcher dispatch smoke to consume `owner_service_cli_args` from the Go preview before calling `xnix-runtime-owner`.
- Extended script and layout guards so the real staged owner-service launch lane cannot drift back to a Ruby-reconstructed method/evidence pair.

## [0.2.632] - 2026-07-24

### Changed

- Updated `dbus_controlled_launch_owner_fixture_smoke.rb` to consume the Go-owned desktop trigger bridge by deriving the D-Bus method from `desktop_dbus_method` and the evidence handoff from `owner_service_call_args`.
- Tightened D-Bus smoke guards so the nested owner service call must preserve the same Runtime method provided by the fixture trigger.
- Extended layout and script checks to prevent the D-Bus fixture smoke from drifting back to a Ruby-reconstructed method/evidence launch pair.

## [0.2.631] - 2026-07-24

### Changed

- Extended `known-app-runtime-status-launch-owner-fixture-record` ready output with the KDE D-Bus Runtime-status route, `ShowRuntimeControlledLaunch` method, Runtime launch execution type, desktop trigger readiness, and evidence-only owner service call arguments.
- Updated the D-Bus controlled-launch owner fixture smoke to verify the Go-owned desktop trigger bridge before invoking `org.xnix.Compatibility1.ShowRuntimeControlledLaunch`.
- Tightened layout and unit guards so blocked fixtures do not expose ready desktop triggers, while ready fixtures must map to `ShowRuntimeControlledLaunch evidence-relative-path <relative>`.

## [0.2.630] - 2026-07-24

### Added

- Added `known-app-runtime-status-launch-owner-fixture-record`, a Go Runtime owner helper that prepares D-Bus controlled-launch fixture state behind one reusable command.
- Added targeted Go CLI and appidentity tests for the helper's safe blocked path when the managed known-app artifact is unavailable.

### Changed

- Updated `dbus_controlled_launch_owner_fixture_smoke.rb` to delegate launch authorization, controlled session, review receipt, and Runtime-status evidence handoff setup to the Go helper instead of assembling the fixture through multiple Ruby-owned command calls.
- Extended layout and script guards to verify the helper keeps KDE evidence-only, receipt reconstruction disabled, KDE state-root access disabled, local/backend details hidden, and host/container safety gates closed.

## [0.2.629] - 2026-07-24

### Added

- Added `dbus_controlled_launch_owner_fixture_smoke.rb`, a restricted private-session D-Bus smoke that prepares Runtime-owned launch, controlled session, review receipt, and Runtime-status evidence handoff state in a managed container tmpfs scratch mount before calling `org.xnix.Compatibility1.ShowRuntimeControlledLaunch`.
- Added the `dbus-controlled-launch-owner-fixture-smoke` container command and targeted Ruby guard coverage for the offline, read-only, Docker-socket-free tested Runtime image path.

### Changed

- Extended layout and container guards to verify the D-Bus controlled-launch owner fixture parses the nested Go owner `runtime-owner-service-call` payload while keeping KDE evidence-only, receipt reconstruction disabled, state-root access disabled, backend details hidden, and host/container safety gates closed.

## [0.2.628] - 2026-07-24

### Added

- Added the `ShowRuntimeControlledLaunch` method to the `org.xnix.Compatibility1` smoke D-Bus contract and introspection include.
- Added a C smoke-adapter action bridge that accepts only a safe relative Runtime-status evidence path and forwards it to `xnix-runtime-owner --service-call ShowRuntimeControlledLaunch evidence-relative-path <relative>`.

### Changed

- Extended D-Bus smoke and layout guards to verify controlled launch action evidence, evidence-only KDE handoff fields, Go owner service-call visibility, and disabled desktop state-root/receipt reconstruction gates.
- Classified `ShowRuntimeControlledLaunch` as a controlled desktop action in evidence and drift reports instead of treating it as a read-only Runtime method.

## [0.2.627] - 2026-07-24

### Added

- Added `runtime_status_owner_service_session_bus_smoke.rb`, a private `dbus-run-session` wrapper around the staged launcher dispatch smoke for the Runtime-status owner service launch path.
- Added the `runtime-status-owner-service-session-bus-smoke` container command and targeted Ruby tests for its offline, read-only, Docker-socket-free execution contract.

### Changed

- Updated the current mainline to target a real smoke D-Bus adapter method for `ShowRuntimeControlledLaunch` after the private-session wrapper boundary.

## [0.2.626] - 2026-07-24

### Changed

- Routed the staged launcher dispatch smoke's trigger-fed Runtime-status launch through `xnix-runtime-owner --service-call ShowRuntimeControlledLaunch`, so the real staged `xnix-compat-launch` lane exercises the owner service boundary when QEMU/Wine prerequisites are available.
- Extended staged launcher smoke assertions to verify the `runtime-owner-service-call` envelope, `desktop-action-dispatch`, owner-supplied inputs, delegated receipt/session fields, and `kde-dbus-runtime-status-action` route.

## [0.2.625] - 2026-07-24

### Added

- Added the Go Runtime Owner `ShowRuntimeControlledLaunch` service method, returning `desktop-action-dispatch` evidence for desktop/DBus callers that forward only a Runtime-status evidence handoff.
- Added targeted owner service and CLI tests that exercise the service boundary with Runtime-supplied state root, cache root, launcher, and timeout inputs.

### Changed

- Aligned Runtime-owner action output with delegated receipt/session fields and Compatibility Center projection evidence while preserving redaction of state-root paths, launcher paths, raw launcher output, backend details, Docker socket mounts, broad host mounts, and host-root mutation.

## [0.2.624] - 2026-07-24

### Added

- Added Runtime-owner adapter configuration for `show-runtime-controlled-launch`, using `XNIX_RUNTIME_OWNER_*` values for state root, cache root, managed launcher, and timeout settings.

### Changed

- Tightened the desktop-callable launch route so KDE forwards only the Runtime-status evidence handoff while owner-only launch inputs are supplied by the Runtime boundary instead of desktop CLI flags.

## [0.2.623] - 2026-07-24

### Added

- Added the desktop-callable `show-runtime-controlled-launch` CLI route that runs the trigger-fed Runtime launch execution from a safe Runtime-status evidence handoff.

### Changed

- Updated the staged launcher dispatch smoke to call `show-runtime-controlled-launch` for the second trigger-fed execution and verify that KDE forwards only the evidence handle without reconstructing app, receipt, session, card, action, or dispatch fields.

## [0.2.622] - 2026-07-24

### Added

- Added trigger-fed execution planning for `known-app-kde-runtime-status-launch-execution`, allowing the Runtime-owned execution wrapper to consume a persisted Runtime-status launch evidence handoff by opaque id or safe relative path.

### Changed

- Extended the staged launcher dispatch smoke so, after the first managed launcher pass records the Runtime-status handoff, a second Runtime execution call derives its launcher plan from `known-app-kde-runtime-status-launch-action-trigger-preview` instead of reconstructing launch request fields from smoke-local variables.

## [0.2.621] - 2026-07-24

### Added

- Added the Go-owned `known-app-kde-runtime-status-launch-action-trigger-preview` command, which consumes the persisted Runtime-status launch evidence handoff and assembles a Runtime-owned managed launcher request for the `show-runtime-controlled-launch` action.

### Changed

- Extended the staged launcher dispatch smoke so the KDE Runtime-status action is forwarded through the Runtime handoff trigger preview instead of reconstructing launch request fields directly from smoke-local variables.

## [0.2.620] - 2026-07-24

### Changed

- Recorded v0.2.620 as a twentieth small-version formal full Buildroot/QEMU validation checkpoint for the Runtime-status launch evidence handoff reader.
- Captured restricted full-smoke evidence in `output/full-smoke-report.json`, `output/full-smoke-report.md`, and `output/serial.log`, with Buildroot/QEMU boot, SSH/sshd serial evidence, the i386 Wine guest, the pinned 7zr known Windows app lane, staged managed launcher dispatch, and the fixture Windows app smoke all passing while Docker socket mounting and host-root mutation remained disabled.

## [0.2.619] - 2026-07-24

### Added

- Added the Go-owned `known-app-kde-runtime-status-launch-evidence-preview` read model, which consumes the persisted Runtime-status launch evidence handoff by opaque id or safe relative path and verifies its digest before exposing KDE-safe next-action evidence.

### Changed

- Extended delegated Runtime-status launch evidence normalization so the handoff preview exposes a complete `known_app_smoke_evidence` payload with `validated-post-review-dispatch`, `show-runtime-controlled-launch`, and `runtime-status` action state.
- Extended the staged launcher dispatch smoke so it reads back the Runtime-owned evidence handoff before feeding Center previews, proving the next desktop-triggered action can consume persisted handoff state instead of relying only on the smoke-local projection file.

## [0.2.618] - 2026-07-24

### Added

- Added the Go-owned `known-app-kde-runtime-status-launch-evidence-record` command to persist the Runtime-status launch projection as a KDE-readable handoff artifact under the Runtime state root.
- Added `RecordKnownAppKDERuntimeStatusLaunchEvidence`, which validates the delegated evidence projection, verifies it can feed normalized known-app smoke evidence, writes only the projection JSON, and returns a relative evidence path plus SHA-256 digest.

### Changed

- Extended the staged launcher dispatch smoke so the Runtime wrapper's `compatibility_center_known_app_evidence` projection is recorded as a state-root handoff before Compatibility Center and KDE Center read models consume the projection file.

## [0.2.617] - 2026-07-24

### Added

- Added Runtime-projected known-app evidence inputs to `compatibility-center-preview` and `kde-center-page-preview` through `--known-app-evidence-json` and `--known-app-evidence-file`.
- Added Go normalization from `known-app-kde-runtime-status-launch-delegated-evidence` into Compatibility Center known-app smoke evidence, including launch receipt, launch gate, controlled session, launcher session gate, and post-review dispatch fields.

### Changed

- Simplified the staged launcher dispatch smoke so it writes the Runtime wrapper projection once and passes that evidence file to both Center previews instead of reconstructing legacy known-app evidence flags in Ruby.

## [0.2.616] - 2026-07-24

### Added

- Added a Go-owned `known-app-kde-runtime-status-launch-delegated-evidence` projection that converts redacted managed launcher output into the Compatibility Center and KDE Center evidence shape without exposing state-root paths, managed launcher paths, raw launcher output, backend details, Docker socket mounts, broad host mounts, or host-root mutation.

### Changed

- Simplified the staged launcher dispatch smoke so it consumes the Runtime wrapper's `compatibility_center_known_app_evidence` projection directly instead of manually remapping delegated launcher fields in Ruby.

## [0.2.615] - 2026-07-24

### Changed

- Switched the staged launcher dispatch smoke so its QEMU/Wine lane now calls the Go-owned `known-app-kde-runtime-status-launch-execution` wrapper, which then invokes the staged `xnix-compat-launch` executable instead of letting the smoke call the launcher directly.
- Extended Runtime execution wrapper output with redacted delegated launcher evidence used by Compatibility Center and KDE Center page previews, including delegated request type, guest boundary, artifact verification, marker observation, post-review dispatch state, opaque receipt ids, controlled session consumption, and safety flags without exposing raw launcher output or state-root paths.

## [0.2.614] - 2026-07-24

### Added

- Added the Go-owned `known-app-kde-runtime-status-launch-execution` entrypoint, which consumes the KDE Runtime-status launch request, accepts the Runtime-owned state root only at the Runtime boundary, revalidates launch authorization, session-gated review, controlled execution session, post-review state, and managed guest boundary evidence, then invokes an existing managed launcher executable.
- Added Runtime-status launch execution plan and CLI tests proving that the state root is injected into the launcher argv without exposing state-root paths, raw launcher output, backend details, Docker socket mounts, broad host mounts, or host-root mutation in the returned desktop-safe JSON.

## [0.2.613] - 2026-07-24

### Added

- Added the Go-owned `known-app-kde-runtime-status-launch-request-preview` command, which validates post-review KDE card state and collects the three Runtime-generated opaque ids needed for the managed launcher.
- Extended KDE Center page known-app cards and staged launcher dispatch smoke with Runtime-status launch request route metadata, managed launcher argv assembly, Runtime-owned state-root handling, and disabled direct launch, backend launch, request writes, permission grants, path exposure, Docker socket mounts, broad host mounts, and host-root mutation.

## [0.2.612] - 2026-07-24

### Added

- Extended Compatibility Center known-app staged launcher evidence with post-review dispatch consumption fields, including a validated post-review card state, Runtime-status primary action, and opaque session-gated review receipt id.
- Extended KDE Center page known-app session gate cards and the staged launcher dispatch smoke so the UI read model can show launcher-consumed post-review dispatch state while KDE remains presentation-only.

## [0.2.611] - 2026-07-24

### Added

- Updated the Go `xnix-compat-launch` managed launcher so controlled dispatch now requires `--review-receipt-id` and consumes `known-app-session-gated-controlled-dispatch-request-preview` before invoking the staged known Windows app runner.
- Reordered the staged launcher dispatch smoke so the Runtime records and accepts the session-gated launch review receipt before QEMU execution, then verifies launcher-side post-review controlled dispatch state consumption in the final dispatch result.

## [0.2.610] - 2026-07-24

### Added

- Added the Go-owned `known-app-session-gated-controlled-dispatch-request-preview` command, which requires an accepted session-gated launch review receipt before consuming the existing launch authorization gate and creating post-review controlled dispatch request state.
- Extended the staged launcher dispatch smoke, CLI tests, and layout verifier to cover review receipt evidence, launch receipt reuse, post-review controlled dispatch state creation, and disabled direct launch, desktop launch, backend launch, execution, permission grants, raw path exposure, Docker socket mounts, broad host mounts, and host-root mutation.

## [0.2.609] - 2026-07-24

### Added

- Added the Go-owned `known-app-session-gated-launch-review-gate-preview` command, which revalidates the controlled session evidence, reads the Runtime-owned session-gated launch review receipt, verifies its digest, and accepts only approved matching receipts.
- Extended the staged launcher dispatch smoke and targeted tests to verify review receipt consumption, gate readiness, dispatch state advancement readiness, and disabled dispatch-object creation, direct launch, desktop launch, backend launch, execution, path exposure, Docker socket mounts, broad host mounts, and host-root mutation.

## [0.2.608] - 2026-07-24

### Added

- Added the Go-owned `known-app-session-gated-launch-review-receipt-record` command, which re-consumes digest-verified session evidence and records a Runtime-owned review receipt under the explicit state root.
- Extended the staged launcher dispatch smoke and targeted tests to verify relative receipt evidence, receipt digests, read-before-write consumption, and disabled launch, dispatch, request-object, permission-grant, backend, path-exposure, Docker socket, broad host mount, and host-root mutation flags.

## [0.2.607] - 2026-07-24

### Added

- Added the Go-owned `known-app-session-gated-launch-review-preview` command, which re-consumes digest-verified controlled session evidence before exposing a Runtime-owned read-before-write review route for the KDE `review-session-gated-dispatch` card action.
- Extended KDE Center page session-gated known-app cards and the staged launcher dispatch smoke with Runtime review route metadata while keeping desktop launch, backend launch, execution, review receipt recording, raw paths, backend details, Docker socket mounts, broad host mounts, and host-root mutation disabled.

## [0.2.606] - 2026-07-24

### Added

- Added KDE Center page read-model projection for session-gated known-app staged launcher evidence, including safe `known_app_session_gate_cards` with opaque controlled session ids, relative session evidence, digest verification, Runtime-owner consumability, and KDE read-model consumability.
- Extended the staged managed launcher dispatch smoke to verify both Compatibility Center summary evidence and KDE Center page session-gated cards while keeping desktop launch, backend launch, backend details, host paths, Docker socket mounts, broad host mounts, and host-root mutation disabled.

## [0.2.605] - 2026-07-24

### Added

- Extended Compatibility Center known-app staged launcher evidence with launcher-side controlled session gate consumption fields, including opaque session id, relative session evidence, digest verification, Runtime-owner consumability, and KDE read-model consumability.
- Updated the staged managed launcher dispatch smoke to pass launcher session gate evidence into `compatibility-center-preview` and verify the new session-gated card state while keeping desktop launch, backend launch, path exposure, Docker socket mounts, broad host mounts, and host-root mutation disabled.

## [0.2.604] - 2026-07-24

### Added

- Updated the Go `xnix-compat-launch` dispatch path so a controlled managed-guest dispatch now consumes the digest-verified known-app controlled execution session before invoking the existing staged 7zr dispatch lane.
- Extended the staged managed launcher dispatch smoke to pass the opaque session id into the staged launcher and verify launcher-side session consumption evidence in the final dispatch result while keeping state-root path exposure, live desktop activation, backend process starts, Docker socket mounts, broad host mounts, and host-root mutation disabled.

## [0.2.603] - 2026-07-24

### Added

- Added the Go-owned `known-app-controlled-execution-session-consume-preview` command, which reads the persisted known-app ledger/session records, verifies the session digest, and exposes Runtime-owner and KDE read-model consumption evidence.
- Extended the staged managed launcher dispatch smoke to consume the recorded controlled execution session before QEMU dispatch and verify fan-out readiness for task manager, KWin, tray, and Compatibility Center consumers while keeping live desktop activation, dispatch, execution, path exposure, Docker socket mounts, broad host mounts, and host-root mutation disabled.

## [0.2.602] - 2026-07-24

### Added

- Added the Go-owned `known-app-controlled-execution-session-record` command, which validates the known-app controlled execution-session handoff and persists a Runtime execution ledger transaction plus session status record under the explicit state root.
- Extended the staged managed launcher dispatch smoke to record the controlled execution session before QEMU dispatch and verify relative transaction/session evidence, digest evidence, and disabled live-session, dispatch, execution, backend, path-exposure, Docker socket, broad host mount, and host-root mutation flags.

## [0.2.601] - 2026-07-24

### Added

- Added the Go-owned `known-app-controlled-execution-session-preview` command, which consumes the existing known-app launch authorization receipt, launch gate, and controlled dispatch request path before materializing a portable Runtime execution-session handoff.
- Extended the staged managed launcher dispatch smoke to validate the controlled execution-session handoff before QEMU dispatch while keeping session registration, window observation, direct launch, desktop launch, backend launch, backend process start, dispatch start, execution start, receipt path exposure, state-root path exposure, Docker socket mounts, broad host mounts, and host-root mutation disabled.

## [0.2.600] - 2026-07-24

### Changed

- Completed the v0.2.600 formal full Buildroot/QEMU checkpoint after moving the managed launcher dispatch path behind the Go-owned controlled dispatch request gate.
- Verified the full checkpoint with `PASS: staged managed launcher dispatch smoke (7zr 26.02)`, `PASS: QEMU guest real Windows app Wine smoke`, and `PASS: full build and QEMU serial smoke test`.

## [0.2.599] - 2026-07-24

### Changed

- Updated the Go `xnix-compat-launch` entrypoint so any `--guest-boundary` dispatch path now requires `--state-root` and `--receipt-id`, consumes `PreviewKnownAppControlledDispatchRequest`, and refuses dispatch unless the controlled request object is created.
- Reordered the staged launcher dispatch smoke so the Runtime receipt, launch gate, and controlled dispatch request are created before the staged launcher invokes the existing QEMU/Wine dispatch smoke runner.

## [0.2.598] - 2026-07-24

### Added

- Added the Go-owned `known-app-controlled-dispatch-request-preview` command, which turns an accepted Runtime launch gate into a portable controlled dispatch request object only after receipt acceptance, guest-boundary acceptance, and managed artifact verification.
- Extended the staged launcher dispatch smoke to validate the new controlled dispatch request object after a real 7zr PASS while keeping dispatch start, execution start, direct launch, desktop launch, backend launch, backend process start, receipt path exposure, state-root path exposure, Docker socket mounts, broad host mounts, and host-root mutation disabled.

## [0.2.597] - 2026-07-24

### Added

- Wired `scripts/staged_launcher_dispatch_smoke.rb` to record a Go-owned known-app launch authorization receipt and consume it through `known-app-launch-gate-preview` after a real staged 7zr dispatch PASS.
- Extended Compatibility Center known-app smoke evidence with launch-gate consumption, receipt acceptance, guest-boundary acceptance, controlled-dispatch readiness, and matching aggregate counts while keeping direct launch, desktop launch, backend launch, backend process start, execution start, receipt path exposure, state-root path exposure, Docker socket mounts, broad host mounts, and host-root mutation disabled.

## [0.2.596] - 2026-07-24

### Added

- Added the Go-owned `known-app-launch-gate-preview` command, which consumes an opaque known-app launch authorization receipt from an explicit state root and fails closed for missing, malformed, mismatched, or unsafe receipts without exposing receipt paths or state-root paths.
- Wired the launch gate to require the existing `managed-known-app-guest-smoke` boundary before dispatch readiness can advance, while keeping direct launch, desktop launch, backend launch, backend process start, execution start, Docker socket mounts, broad host mounts, raw artifact path exposure, backend detail exposure, and host-root mutation disabled.

## [0.2.595] - 2026-07-24

### Added

- Added the Go-owned `known-app-launch-authorization-receipt-preview` command, which records an opaque Runtime launch authorization receipt for the tested 7zr staged launcher lane under an explicit state root while exposing only a receipt id to KDE-facing consumers.
- Extended Compatibility Center known-app evidence with recorded authorization receipt state, receipt id, launch-gate state, and recorded-count fields so KDE can show `validated-launch-authorization-recorded` without enabling direct launch, desktop launch, backend launch, backend process start, path exposure, Docker socket mounts, broad host mounts, or host-root mutation.

## [0.2.594] - 2026-07-24

### Added

- Added concrete Compatibility Center card/action state for staged launcher known-app evidence, marking the 7zr smoke as `validated-launch-authorization-required` with `review-launch-authorization` as the safe next action.
- Extended staged launcher dispatch smoke validation to prove the Center projection keeps direct launch and backend launch disabled after a real PASS.

## [0.2.593] - 2026-07-24

### Added

- Added staged launcher dispatch source fields to Go Compatibility Center known-app smoke evidence, including staged launcher verification, Runtime dispatch verification, launch authorization requirement, and disabled desktop launch state.
- Extended the staged launcher dispatch smoke to validate that a passing 7zr run is immediately projected into redacted Compatibility Center evidence without enabling backend launch or exposing implementation details.

## [0.2.592] - 2026-07-24

### Changed

- Promoted `staged-launcher-dispatch-smoke` into the formal full-smoke known Windows app lane so the next full checkpoint validates the staged desktop launcher path instead of the older direct known-app guest harness.
- Updated the full smoke report model and tests to treat `staged-launcher-dispatch-smoke` as known Windows app pass evidence while retaining the fixture Wine guest smoke as a lower-level regression.

## [0.2.591] - 2026-07-24

### Added

- Added a staged managed launcher dispatch smoke that preflights the managed 7-Zip artifact, stages a freshly built Go launcher, invokes the staged launcher with the authorized `managed-known-app-guest-smoke` boundary when QEMU Wine guest prerequisites exist, and validates redacted `windows-known-app-dispatch-smoke` output.
- Added the restricted container entrypoint `ruby scripts/container.rb staged-launcher-dispatch-smoke` plus script-level tests so the staged desktop launcher path can become the main real Windows app smoke lane without changing the full-smoke cadence.

## [0.2.590] - 2026-07-24

### Added

- Added a staged managed launcher smoke script that builds the Go launcher, stages it through `desktop-activation-stage --managed-launcher-bin`, executes the staged `usr/local/bin/xnix-compat-launch --app 7zr`, validates redacted launch bridge JSON, and keeps execution, backend, and host mutation gates closed.

## [0.2.589] - 2026-07-24

### Added

- Added explicit `--managed-launcher-bin` support to the Go `desktop-activation-stage` command so a prebuilt Go `xnix-compat-launch` executable can be copied into a controlled staging root.
- Added staged executable support to the Go activation writer at `usr/local/bin/xnix-compat-launch` with mode `0755`, digest evidence, and no host-root mutation.
- Added matching `managed_launcher_bin` support to the Ruby desktop activation installer for legacy staging workflows.
- Extended launcher artifact receipts to record `staged_executable`, `binary_copied`, and `executable_staged` state while preserving the Runtime bridge and dispatch gate evidence.

## [0.2.588] - 2026-07-24

### Added

- Added a managed launcher artifact receipt to Go desktop activation staging previews so KDE desktop entries have an auditable Go `xnix-compat-launch` command source.
- Added the same managed launcher artifact receipt to the Go controlled staging writer under `usr/share/xnix/compatibility/launcher-artifacts/xnix-compat-launch.json` without copying binaries or mutating the host root.
- Added the managed launcher artifact receipt to the Ruby desktop activation installer staging path, keeping legacy test harness compatibility while documenting that the desktop launcher command is supplied by `cmd/xnix-compat-launch`.
- Extended targeted activation tests and layout checks to require the launcher artifact, Runtime bridge method, and gated dispatch boundary.

## [0.2.587] - 2026-07-24

### Added

- Added the Go-owned `cmd/xnix-compat-launch` executable entrypoint for the managed desktop launcher path.
- Wired the launcher entrypoint to `PreviewKnownPortableLaunchBridge` by default so desktop invocation returns a redacted Runtime-owned bridge result instead of Ruby product logic.
- Added an optional `--guest-boundary managed-known-app-guest-smoke` path that continues into `RunKnownPortableDispatchSmoke` only after bridge materialization proves the managed artifact is verified.
- Added Docker build and layout verification coverage for the Go launcher while retaining the legacy Ruby `bin/xnix-compat-launch` wrapper for existing generic recipe launch tests.

## [0.2.586] - 2026-07-24

### Added

- Added the Go-owned `windows-known-app-launch-bridge-preview` CLI, turning the managed `xnix-compat-launch --app 7zr` launcher invocation into a Runtime-owned gated dispatch-smoke request model.
- Added launcher argv validation so bridge previews accept only the Runtime-owned managed launcher shape and block unexpected launcher arguments before dispatch-smoke request materialization.
- Added launch bridge readiness states that keep dispatch-smoke request materialization blocked until the managed artifact checksum verifies, then mark the request model ready for a smoke harness or future Runtime service owner to supply the controlled guest boundary.
- Added targeted Go tests proving launcher argv acceptance/rejection, verified-artifact request materialization, unknown-app rejection, and redacted CLI output without raw `.exe` paths, Wine/QEMU details, backend internals, host paths, backend starts, or host mutation.

## [0.2.585] - 2026-07-24

### Added

- Added the Go-owned `windows-known-app-dispatch-smoke` CLI, which consumes `windows-known-app-dispatch-preview` and invokes the managed known-app guest runner only after artifact verification and an explicit smoke-harness guest boundary are present.
- Added redacted dispatch-smoke result fields for guest reachability, managed runtime readiness, artifact copy, marker observation, pass status, exit code, and duration without exposing raw `.exe` paths, Wine/QEMU command details, backend internals, or host paths.
- Updated the known Windows app QEMU smoke harness to call the gated dispatch command instead of directly invoking the lower-level known-app guest runner.
- Added targeted Go tests proving artifact and guest-boundary gates block safely, verified artifacts can invoke the fake managed guest runner, unknown apps are rejected, and CLI output preserves the managed safety boundary.

## [0.2.584] - 2026-07-24

### Added

- Added the Go-owned `windows-known-app-dispatch-preview` CLI, mapping the pinned 7-Zip Runtime launch request into a managed dispatch preview for the known-app guest smoke lane.
- Added dispatch readiness states that keep dispatch blocked until the managed artifact checksum verifies, then mark the dispatch preview ready without starting execution or backend processes.
- Added targeted Go tests proving the dispatch preview consumes the launch request, preserves the safe managed launcher argv, rejects unknown apps, stays dry-run, and avoids raw `.exe` paths, Wine/QEMU details, backend internals, host paths, host networking, Docker socket mounts, broad host mounts, and host-root mutation.

## [0.2.583] - 2026-07-24

### Added

- Added the Go-owned `windows-known-app-launch-request-preview` CLI, turning the pinned 7-Zip KDE launcher entrypoint into a Runtime-owned launch request object.
- Added launch request readiness states that block dispatch until the managed 7-Zip artifact checksum verifies, then mark the request ready for future managed dispatch through the existing known-app guest smoke lane.
- Added targeted Go tests proving the launch request consumes the KDE launcher identity, preserves the managed launcher argv, rejects unknown apps, stays dry-run, and avoids backend process starts, raw `.exe` paths, Wine/QEMU details, backend internals, host paths, host networking, Docker socket mounts, broad host mounts, and host-root mutation.

## [0.2.582] - 2026-07-24

### Added

- Added the Go-owned `windows-known-app-kde-launcher-preview` CLI, wiring the pinned 7-Zip managed launch surface into a KDE Plasma launcher entrypoint model.
- Added KDE launcher readiness states that show the 7-Zip launcher as visible but preparation-required until the managed cache artifact verifies, then visible and launch-enabled once verification passes.
- Added targeted Go tests proving the KDE launcher model consumes the managed launch surface, preserves the `xnix-compat-launch --app 7zr` argv, rejects unknown apps, and avoids raw `.exe` paths, Wine/QEMU details, backend internals, host paths, host-root mutation, broad mounts, host networking, and Docker socket exposure.

## [0.2.581] - 2026-07-23

### Added

- Added the Go-owned `windows-known-app-managed-launch-preview` Runtime CLI for the pinned 7-Zip app, exposing a KDE-safe managed launch surface without raw `.exe` paths, Wine/QEMU command details, backend internals, or host paths.
- Added managed launch preview checks for missing, checksum-mismatched, and verified known-app cache states so desktop entrypoints can distinguish preparation-required from launch-ready states.
- Added targeted Go tests covering the redacted managed launch model, CLI output, managed launcher argv, unknown-app rejection, and safety flags.

## [0.2.580] - 2026-07-23

### Changed

- Promoted the formal full checkpoint cadence from every tenth patch version to every twentieth patch version, matching the current development strategy for targeted small-version validation.
- Extended `scripts/full_smoke.rb` to build both runtime and tools images explicitly, build the SSH-enabled Wine guest, fetch the pinned known Windows application, and require concrete `PASS` output from both the known-app QEMU Wine smoke and the fixture QEMU Wine smoke before accepting the checkpoint.
- Extended the full smoke report with Wine guest and Windows app smoke pass evidence while preserving the no-privileged-container, no-host-network, no-Docker-socket, no-host-root-mutation safety boundary.
- Added owner-package source/token and nested-preview validation caching so full Go validation can finish under the formal checkpoint instead of spending tens of minutes repeatedly scanning the same historical preview chains.
- Verified the v0.2.580 restricted full smoke: `build-tools`, `build`, Buildroot source/config/download/build, SSH key preparation, Wine guest config/download/build, known Windows app fetch, base QEMU boot, known 7-Zip QEMU Wine smoke, and fixture QEMU Wine smoke all completed with PASS evidence.

## [0.2.579] - 2026-07-23

### Added

- Added a Go-owned Compatibility Center known-app smoke evidence summary so the KDE read model can show that a managed real Windows application smoke has passed without exposing `.exe` paths, Wine/QEMU command details, backend internals, or host paths.
- Added `compatibility-center-preview` CLI flags for consuming redacted known-app smoke status, marker observation, and checksum verification evidence produced by the Runtime smoke lane.
- Kept KDE presentation-only and all action execution, backend launch, settings persistence, host-root mutation, raw artifact path exposure, and backend detail exposure disabled while surfacing the real-app smoke result.

## [0.2.578] - 2026-07-23

### Added

- Added a Go-owned known Windows app catalog with the pinned 7-Zip `7zr.exe` 26.02 x86 standalone console executable, official source metadata, SHA256 verification, and the `7-Zip` stdout marker for QEMU Wine guest validation.
- Added `windows-known-app-fetch` and `windows-known-app-guest-wine-smoke` Runtime commands so known portable app acquisition and guest execution are owned by Go while preserving redacted JSON evidence.
- Added `scripts/container.rb fetch-known-winapp` for explicit bridge-networked artifact fetch into the managed cache volume, and `scripts/container.rb known-winapp-guest-wine-smoke` for no-network QEMU Wine guest execution from that verified cache.
- Verified `PASS: known Windows app fetched and verified (7zr 26.02)`, `PASS: known Windows app QEMU guest Wine smoke (7zr 26.02)`, and the retained `PASS: QEMU guest real Windows app Wine smoke` regression under the restricted Colima/Docker boundary.

## [0.2.577] - 2026-07-23

### Added

- Added `--expected-marker` to the Go-owned `windows-app-guest-wine-smoke` CLI so non-fixture Windows executables can be validated by their own stdout marker while keeping redacted JSON evidence.
- Extended `scripts/winapp_guest_wine_smoke.rb` with `XNIX_WINAPP_GUEST_EXE`, `XNIX_WINAPP_GUEST_ARGS`, and `XNIX_WINAPP_GUEST_MARKER`, allowing the same restricted QEMU Wine guest path to run a container-visible external Windows executable instead of always rebuilding the tiny fixture.
- Kept the default fixture path intact as the regression baseline for `PASS: QEMU guest real Windows app Wine smoke`, and verified the new external executable path with `PASS: QEMU guest external Windows app Wine smoke`.

## [0.2.576] - 2026-07-23

### Fixed

- Fixed the QEMU guest Wine smoke harness to set `GOTMPDIR` under the repository cache before invoking Go, avoiding `/tmp/go-build...` execution denial inside the restricted tools container.
- Added a Buildroot post-build Wine NLS install step so the guest rootfs includes `l_intl.nls` and related files under `/usr/share/wine/nls`, matching the path searched by `wineserver`.
- Increased the Wine guest smoke SSH boot wait window to 180 seconds so the heavier Wine/glibc initramfs can boot under 2-CPU Colima before the Runtime smoke begins.
- Persisted the QEMU serial console from `winapp-guest-wine-smoke` to `.cache/xnix/winapp-guest-wine-smoke/qemu-serial.log` so the next guest-side Wine failure can be diagnosed from kernel/init logs.
- Switched the i386 Wine guest QEMU CPU model from `max` to `qemu32` after serial logs showed 32-bit Wine processes triggering kernel Oops with the overly broad emulated CPU feature set.
- Recorded evidence that the retained v0.2.575 no-network Wine guest build completed successfully, the v0.2.576 incremental rootfs rebuild installed 76 Wine NLS files, and `ruby scripts/container.rb winapp-guest-wine-smoke` now returns `PASS: QEMU guest real Windows app Wine smoke`.
- Extended the Buildroot and layout verifiers to keep the `GOTMPDIR` boundary and Wine NLS rootfs install path visible.

## [0.2.575] - 2026-07-23

### Changed

- Added a Buildroot source patch application step to `scripts/fetch_buildroot.rb` so repository-owned Buildroot metadata fixes are applied after extracting the pinned Buildroot source archive.
- Added a Buildroot 2025.02.15 Wine package metadata patch that removes `--without-mingw` from the host-wine cut-down configure options, allowing Wine 10 to detect the already installed clang, lld, and llvm-dlltool PE toolchain on aarch64 Colima builders.
- Extended the layout verifier to require the Buildroot source patch file and the idempotent patch application path; local evidence shows `fetch-sources` applied the patch, the v0.2.575 host-wine configure command no longer includes `--without-mingw`, Wine detected clang/lld PE support, and `host-wine` entered its build step.

## [0.2.574] - 2026-07-23

### Changed

- Added `clang`, `lld`, and `llvm` to the tools Docker image so Buildroot `host-wine` can satisfy Wine's aarch64 PE cross-compilation requirement inside Colima.
- Extended the layout verifier to keep the aarch64 `host-wine` toolchain dependency visible in the builder image contract.
- Recorded local evidence that the v0.2.573 offline Wine guest build moved past the previous missing-source failure and then failed while configuring `host-wine` with `PE cross-compilation is required for aarch64`; `xnix-builder-tools:0.2.574` built with the requested LLVM tools, and the v0.2.574 detached offline build exposed Buildroot's host-wine `--without-mingw` metadata blocker.

## [0.2.573] - 2026-07-23

### Changed

- Added `scripts/container.rb download-wine-guest` so Wine guest Buildroot sources can be fetched in an explicit bridge-networked preparation step before the restricted offline build.
- Kept `build-ssh-wine-guest` and `start-build-ssh-wine-guest` networkless, preserving the no-network build boundary while allowing missing sources such as `gettext-tiny` to be fetched through the dedicated preparation command.
- Recorded local evidence that the v0.2.572 detached offline Wine guest build reached OpenSSH, host-patchelf, and urandom-scripts before failing on the missing `host-gettext-tiny` tarball because DNS is unavailable in the networkless build container; after `download-wine-guest`, the v0.2.573 detached offline build started and progressed into `host-flex` compilation.

## [0.2.572] - 2026-07-23

### Changed

- Added `scripts/container.rb start-build-ssh-wine-guest` so the heavy Wine-capable Buildroot guest can build as a detached observed Docker container instead of flooding the interactive Codex session.
- Updated the QEMU guest Wine smoke skip guidance to point at the detached Wine guest build path while preserving the existing restricted container boundary: no host networking, no privileged mode, no Docker socket mount, no broad host mount, and only the managed cache volume mounted.
- Recorded local evidence that `ruby scripts/container.rb build-tools` creates `xnix-builder-tools:0.2.572`, and `ruby scripts/container.rb start-build-ssh-wine-guest` starts a retained container that reaches the Buildroot i386 GCC toolchain build stage; no concrete Buildroot package failure has been observed yet.

## [0.2.571] - 2026-07-23

### Changed

- Split the Docker builder into a lightweight `tools` target and the existing full `tested-runtime` target so Buildroot and Wine guest bring-up can run without triggering hidden `go test ./...` full validation during small-version work.
- Added a versioned `xnix-builder-tools` image path plus `scripts/container.rb build-tools`; Buildroot configure/build, source retrieval, SSH key preparation, QEMU boot smoke, and guest Wine smoke now use the tools image while Runtime D-Bus smokes continue to use the fully validated Runtime image.
- Moved the QEMU guest Wine smoke working directory and Go caches under `.cache` so the script remains writable inside the restricted read-only container with only the managed cache volume mounted.
- Recorded local evidence that `ruby scripts/container.rb build-tools` builds `xnix-builder-tools:0.2.571` without entering the full Go suite and `ruby scripts/container.rb configure-wine-guest` writes the isolated Wine guest Buildroot configuration.

## [0.2.570] - 2026-07-23

### Added

- Added an isolated Buildroot `xnix_wine_i386_defconfig` for the first Wine-capable QEMU guest baseline, with an i386 target, glibc toolchain, OpenSSH server, and Buildroot Wine package enabled separately from the default x86_64 SSH guest.
- Added Wine guest build commands in `scripts/container.rb` (`configure-wine-guest`, `build-wine-guest`, and `build-ssh-wine-guest`) so the heavier Wine image path is explicit and does not bloat the normal tiny QEMU smoke baseline.
- Added a dedicated i386 QEMU helper profile and updated the QEMU guest Wine smoke harness to build a 32-bit Windows PE fixture and boot the Wine guest kernel, while safely reporting `SKIP` until the Wine guest image exists.

## [0.2.569] - 2026-07-23

### Added

- Added a Go-owned `windows-app-guest-wine-smoke` command that copies a real Windows `.exe` into a loopback SSH QEMU guest, checks for guest-side Wine, runs the executable inside the guest, captures the smoke marker, and reports redacted JSON without exposing host paths.
- Added `scripts/winapp_guest_wine_smoke.rb` and a `scripts/container.rb winapp-guest-wine-smoke` entrypoint so Ruby remains the QEMU/test harness while the Runtime owns the guest Wine execution path in Go.
- Added unit and CLI coverage for the guest Wine smoke path with fake SSH/SCP clients, including the current expected `guest wine runner unavailable` skip mode for guests that boot but do not yet include Wine.

## [0.2.568] - 2026-07-23

### Changed

- Hardened the containerized Windows app smoke path for Apple/Colima by defaulting the local Wine container to `linux/amd64`, adding platform-specific image validation, and shrinking the Docker build context with local build/cache artifact exclusions.
- Added a Docker build fallback for environments without `docker buildx`; the fallback prepares and commits a local Wine image for the requested platform while keeping the preparation path explicit and noninteractive.
- Moved Wine smoke state into a container tmpfs, reduced host exposure to one read-only application mount, raised the restricted container resource ceiling to 2 CPUs and 2 GB for Wine bootstrap, and added explicit Wine bootstrap timeout evidence in the Runtime JSON result.
- Recorded current Colima evidence: the local `xnix-wine-smoke:local` image builds as `linux/amd64`, Docker exposes 2 CPUs after restart, and the real Windows PE fixture reaches Wine prefix bootstrap but times out in the current macOS Colima/QEMU-user environment instead of producing the smoke marker.

## [0.2.567] - 2026-07-23

### Added

- Added a Go-owned `windows-app-container-run-smoke` command that runs the real Windows PE smoke fixture through a local Wine container image with `--pull never`, `--network none`, dropped capabilities, no new privileges, bounded CPU/memory/PID limits, two narrow repository-local mounts, and redacted JSON evidence.
- Added `scripts/winapp_container_smoke.rb` to build the Windows PE fixture and run the containerized smoke through the Runtime command while reporting `PASS` for a working local Wine image or `SKIP` when the image is unavailable.
- Added `containers/wine-smoke.Dockerfile` and `scripts/build_wine_smoke_image.rb` as the explicit local-image preparation path for turning the container smoke from `SKIP` into a real Wine-backed `PASS`.

## [0.2.566] - 2026-07-23

### Added

- Added a Go-owned `windows-app-run-smoke` Runtime command that runs an explicit Windows `.exe` through an isolated compatibility state root, captures redacted JSON evidence, and keeps host networking, privileged containers, Docker socket mounts, broad host mounts, and host-root mutation disabled.
- Added a cross-compiled Windows PE smoke fixture under `test/fixtures/winapp/hello` plus `scripts/winapp_smoke.rb`, which builds the fixture with Go, stores local artifacts under ignored repository-local state, and reports `PASS` when a compatibility runner is available or `SKIP` when it is not installed.
- Added unit and CLI coverage for the Windows app smoke runner using a fake compatibility runner so CI can verify execution, marker capture, state-root isolation, and host-path redaction without requiring Wine.

## [0.2.565] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-call-gate-preview` to the v0.2.564 storage record writer enablement predecessor.
- Added explicit storage record writer enablement evidence consumption before modeling review-only storage record writer call gates while keeping call gates not callable, disabled, ungranted, and unaccepted; enablements not callable, disabled, ungranted, and unaccepted; storage records unwritten; receipt persistence and writes disabled; receipt consumer enablement disabled; dry-run result persistence disabled; request-object dispatch disabled; notification actions disabled; path exposure disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.565 as a targeted-validation-only follow-up after the v0.2.560 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.580 unless explicitly requested earlier.

## [0.2.564] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-enablement-preview` to the v0.2.563 storage record writer grant predecessor.
- Added explicit storage record writer grant evidence consumption before modeling review-only storage record writer enablements while keeping enablements not callable, disabled, ungranted, and unaccepted; grants not callable, disabled, ungranted, and unaccepted; accepted receipt gates disabled and unaccepted; storage records unwritten; receipt persistence and writes disabled; receipt consumer enablement disabled; dry-run result persistence disabled; request-object dispatch disabled; notification actions disabled; path exposure disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.564 as a targeted-validation-only follow-up after the v0.2.560 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.580 unless explicitly requested earlier.

## [0.2.563] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-grant-preview` to the v0.2.562 storage record writer accepted receipt gate predecessor.
- Added explicit storage record writer accepted receipt gate evidence consumption before modeling review-only storage record writer grants while keeping grants not callable, disabled, ungranted, and unaccepted; accepted receipt gates not callable, disabled, ungranted, and unaccepted; authorization receipts unaccepted; storage records unwritten; receipt persistence and writes disabled; receipt consumer enablement disabled; dry-run result persistence disabled; request-object dispatch disabled; notification actions disabled; path exposure disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.563 as a targeted-validation-only follow-up after the v0.2.560 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.580 unless explicitly requested earlier.

## [0.2.562] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-preview` to the v0.2.561 storage record writer authorization receipt acceptance predecessor.
- Added explicit storage record writer accepted receipt gate evidence consumption while keeping accepted receipt gates not callable, disabled, ungranted, and unaccepted; authorization receipt acceptances not callable, disabled, ungranted, and unaccepted; authorization receipts absent, unwritten, unpersisted, not callable, disabled, and ungranted; storage records unwritten; receipt persistence and writes disabled; receipt consumer enablement disabled; dry-run result persistence disabled; request-object dispatch disabled; notification actions disabled; path exposure disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.562 as a targeted-validation-only follow-up after the v0.2.560 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.580 unless explicitly requested earlier.

## [0.2.561] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-preview` to the v0.2.560 storage record writer authorization receipt predecessor.
- Added explicit storage record writer authorization receipt acceptance evidence consumption while keeping authorization receipt acceptances not callable, disabled, ungranted, and unaccepted; authorization receipts absent, unwritten, unpersisted, not callable, disabled, and ungranted; storage record writer authorizations not callable, disabled, ungranted, and unauthorized; storage records unwritten; receipt persistence and writes disabled; receipt consumer enablement disabled; dry-run result persistence disabled; request-object dispatch disabled; notification actions disabled; path exposure disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.561 as a targeted-validation-only follow-up after the v0.2.560 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.580 unless explicitly requested earlier.

## [0.2.560] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-preview` to the v0.2.559 storage record writer authorization predecessor.
- Added explicit storage record writer authorization receipt evidence consumption while keeping authorization receipts absent, unwritten, unpersisted, not callable, disabled, and ungranted; storage record writer authorizations not callable, disabled, ungranted, and unauthorized; storage records unwritten; receipt persistence and writes disabled; receipt consumer enablement disabled; dry-run result persistence disabled; request-object dispatch disabled; notification actions disabled; path exposure disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.560 as the next formal full Buildroot/QEMU checkpoint after v0.2.540; targeted validation and the restricted full build plus QEMU serial smoke passed before release tagging.

## [0.2.559] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-writer-authorization-preview` to the v0.2.558 closed storage record writer implementation predecessor.
- Added explicit closed storage record writer implementation authorization evidence consumption to the storage record writer authorization preview while keeping storage record writer authorizations not callable, disabled, ungranted, and unauthorized; closed writers not callable and unimplemented; storage record contracts not writable; storage records unwritten; storage persistence gates not passed; receipt persistence disabled; receipt writes disabled; receipt consumer enablement disabled; dry-run result persistence disabled; request-object dispatch disabled; notification actions disabled; path exposure disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.559 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.558] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-closed-storage-record-writer-implementation-preview` to the v0.2.557 storage record contract predecessor.
- Added explicit storage record contract authorization evidence consumption to the closed storage record writer implementation preview while keeping closed writers not callable, not enabled, and unimplemented; storage record contracts not writable; storage records unwritten; storage persistence gates not passed; receipt persistence disabled; receipt writes disabled; receipt consumer enablement disabled; dry-run result persistence disabled; request-object dispatch disabled; notification actions disabled; path exposure disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.558 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.557] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-record-contract-preview` to the v0.2.556 storage persistence gate audit predecessor.
- Added explicit storage persistence gate evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt storage record contract preview while keeping storage record contracts not writable; storage records unwritten; storage persistence gates not passed; storage persistence disabled; receipt persistence disabled; receipt writes disabled; receipt consumer enablement disabled; dry-run result persistence disabled; request-object dispatch disabled; notification actions disabled; path exposure disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.557 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.556] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-storage-persistence-gate-audit-preview` to the v0.2.555 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt closed persistence implementation predecessor.
- Added explicit closed persistence implementation evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt storage persistence gate audit preview while keeping storage persistence gates not passed; storage persistence disabled; closed persistence implementations disabled and unimplemented; persistence authorizations not callable, disabled, ungranted, and unauthorized; consumer enablement receipts absent, unpersisted, disabled, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.556 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.555] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-closed-persistence-implementation-preview` to the v0.2.554 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt persistence authorization predecessor.
- Added explicit receipt persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt closed persistence implementation preview while keeping closed persistence implementations disabled and unimplemented; persistence authorizations not callable, disabled, ungranted, and unauthorized; consumer enablement receipts absent, unpersisted, disabled, unaccepted, and unauthorized; consumer enablement gates disabled; consumer enablements disabled; consumer authorizations disabled; receipt consumption gates disabled; receipt acceptances disabled; call receipts disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.555 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.554] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-persistence-authorization-preview` to the v0.2.553 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt predecessor.
- Added explicit receipt consumer enablement receipt evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt persistence authorization preview while keeping persistence authorizations not callable, disabled, ungranted, and unauthorized; consumer enablement receipts absent, unpersisted, disabled, unaccepted, and unauthorized; consumer enablement gates disabled; consumer enablements disabled; consumer authorizations disabled; receipt consumption gates disabled; receipt acceptances disabled; call receipts disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.554 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.553] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-preview` to the v0.2.552 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement gate predecessor.
- Added explicit receipt consumer enablement gate evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt preview while keeping consumer enablement receipts absent, unpersisted, disabled, unaccepted, and unauthorized; consumer enablement gates disabled; consumer enablements disabled; consumer authorizations disabled; receipt consumption gates disabled; receipt acceptances disabled; call receipts disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.553 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.552] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-gate-preview` to the v0.2.551 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement predecessor.
- Added explicit receipt consumer enablement evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement gate preview while keeping consumer enablement gates not callable, disabled, ungranted, unaccepted, and unauthorized; consumer enablements disabled; consumer authorizations disabled; receipt consumption gates disabled; receipt acceptances disabled; call receipts disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.552 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.551] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-preview` to the v0.2.550 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer authorization predecessor.
- Added explicit receipt consumer authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement preview while keeping consumer enablements not callable, disabled, ungranted, unaccepted, and unauthorized; consumer authorizations disabled; receipt consumption gates disabled; receipt acceptances disabled; call receipts disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.551 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.550] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-authorization-preview` to the v0.2.549 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumption gate predecessor.
- Added explicit receipt consumption gate evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer authorization preview while keeping consumer authorizations not callable, disabled, ungranted, unaccepted, and unauthorized; receipt consumption gates disabled; receipt acceptances disabled; call receipts disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.550 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.549] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumption-gate-preview` to the v0.2.548 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt acceptance predecessor.
- Added explicit receipt acceptance evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumption gate preview while keeping receipt consumption gates not callable, disabled, ungranted, unaccepted, and unauthorized; receipt acceptances disabled; call receipts disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.549 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.548] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-acceptance-preview` to the v0.2.547 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt predecessor.
- Added explicit call receipt evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt acceptance preview while keeping receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; call receipts disabled; calls disabled; storage record writer calls disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.548 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.547] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-preview` to the v0.2.546 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call predecessor.
- Added explicit call evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt preview while keeping call receipts not callable, disabled, ungranted, unaccepted, and unauthorized; calls disabled; call-gate calls disabled; storage record writer calls disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.547 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.546] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-preview` to the v0.2.545 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call-gate predecessor.
- Added explicit call-gate evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call preview while keeping calls not callable, disabled, ungranted, unaccepted, and unauthorized; call-gate calls disabled; enablement calls disabled; storage record writer calls disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.546 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.545] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-gate-preview` to the v0.2.544 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate enablement predecessor.
- Added explicit enablement evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call-gate preview while keeping call-gate calls not callable, disabled, ungranted, unaccepted, and unauthorized; enablement calls disabled; grant calls disabled; storage record writer calls disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.545 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.544] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-enablement-preview` to the v0.2.543 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate grant predecessor.
- Added explicit grant evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate enablement preview while keeping enablement calls not callable, disabled, ungranted, unaccepted, and unauthorized; grant calls disabled; authorization calls disabled; storage record writer calls disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.544 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.543] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-grant-preview` to the v0.2.542 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate authorization predecessor.
- Added explicit authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate grant preview while keeping grant calls not callable, disabled, ungranted, unaccepted, and unauthorized; authorization calls disabled; storage record writer calls disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.543 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.542] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-authorization-preview` to the v0.2.541 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate predecessor.
- Added explicit accepted receipt gate evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate authorization preview while keeping authorization calls not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer calls disabled; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.542 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.541] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-preview` to the v0.2.540 storage record writer call authorization receipt accepted receipt gate call authorization receipt acceptance predecessor.
- Added explicit call authorization receipt acceptance evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate preview while keeping storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate call authorization receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate call authorization receipts not callable, disabled, ungranted, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.541 as a targeted-validation-only follow-up after the v0.2.540 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.560 unless explicitly requested earlier.

## [0.2.540] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-acceptance-preview` to the v0.2.539 storage record writer call authorization receipt accepted receipt gate call authorization receipt predecessor.
- Added explicit call authorization receipt evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt acceptance preview while keeping storage record writer call authorization receipt accepted receipt gate call authorization receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate call authorization receipts not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate call authorizations not callable, disabled, ungranted, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.540 as the formal restricted full Buildroot/QEMU checkpoint after v0.2.520.

## [0.2.539] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-preview` to the v0.2.538 storage record writer call authorization receipt accepted receipt gate call authorization predecessor.
- Added explicit accepted receipt gate call authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt preview while keeping storage record writer call authorization receipt accepted receipt gate call authorization receipts not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate call authorizations not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate call gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.539 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.538] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-preview` to the v0.2.537 storage record writer call authorization receipt accepted receipt gate call gate predecessor.
- Added explicit accepted receipt gate call gate evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization preview while keeping storage record writer call authorization receipt accepted receipt gate call authorizations not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate call gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate enablements not callable, disabled, ungranted, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.538 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.537] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-gate-preview` to the v0.2.536 storage record writer call authorization receipt accepted receipt gate enablement predecessor.
- Added explicit accepted receipt gate enablement evidence consumption to the storage record writer call authorization receipt accepted receipt gate call gate preview while keeping storage record writer call authorization receipt accepted receipt gate call gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate enablements not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate grants not callable, disabled, ungranted, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.537 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.536] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-enablement-preview` to the v0.2.535 storage record writer call authorization receipt accepted receipt gate grant predecessor.
- Added explicit accepted receipt gate grant evidence consumption to the storage record writer call authorization receipt accepted receipt gate enablement preview while keeping storage record writer call authorization receipt accepted receipt gate enablements not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate grants not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate authorizations not callable, disabled, ungranted, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.536 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.535] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-grant-preview` to the v0.2.534 storage record writer call authorization receipt accepted receipt gate authorization predecessor.
- Added explicit accepted receipt gate authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate grant preview while keeping storage record writer call authorization receipt accepted receipt gate grants not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gate authorizations not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.535 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.534] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-authorization-preview` to the v0.2.533 storage record writer call authorization receipt accepted receipt gate predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate evidence consumption to the storage record writer call authorization receipt accepted receipt gate authorization preview while keeping storage record writer call authorization receipt accepted receipt gate authorizations not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt accepted receipt gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.534 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.533] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-preview` to the v0.2.532 storage record writer call authorization receipt acceptance predecessor.
- Added explicit storage record writer call authorization receipt acceptance evidence consumption to the storage record writer call authorization receipt accepted receipt gate preview while keeping storage record writer call authorization receipt accepted receipt gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipts not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorizations not callable, disabled, ungranted, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.533 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.532] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-acceptance-preview` to the v0.2.531 storage record writer call authorization receipt predecessor.
- Added explicit storage record writer call authorization receipt acceptance call persistence authorization evidence consumption to the storage record writer call authorization receipt acceptance preview while keeping storage record writer call authorization receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorization receipts not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorizations not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer enablements not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer grants not callable, disabled, ungranted, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Ignored the repo-local `.local-cache/` directory used for sandbox-safe Go build cache during targeted validation.
- Recorded v0.2.532 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.531] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-preview` to the v0.2.530 storage record writer call authorization predecessor.
- Added explicit storage record writer call authorization call-authorization receipt call persistence authorization evidence consumption to the storage record writer call authorization receipt preview while keeping storage record writer call authorization receipts not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call authorizations not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer enablements not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer grants not callable, disabled, ungranted, unaccepted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.531 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.530] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-preview` to the v0.2.529 storage record writer call gate predecessor.
- Added explicit storage record writer call gate call-authorization call persistence authorization evidence consumption to the storage record writer call authorization preview while keeping storage record writer call authorizations not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer call gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer enablements not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer grants not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer accepted receipt gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer authorization receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer authorization receipts not callable, disabled, ungranted, and unauthorized; storage record writer authorizations not callable, disabled, ungranted, and unauthorized; closed storage record writers not callable; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.530 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.529] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-gate-preview` to the v0.2.528 storage record writer enablement predecessor.
- Added explicit storage record writer enablement call-gate call persistence authorization evidence consumption to the storage record writer call gate preview while keeping storage record writer call gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer enablements not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer grants not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer accepted receipt gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer authorization receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer authorization receipts not callable, disabled, ungranted, and unauthorized; storage record writer authorizations not callable, disabled, ungranted, and unauthorized; closed storage record writers not callable; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.529 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.528] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-enablement-preview` to the v0.2.527 storage record writer grant predecessor.
- Added explicit storage record writer grant enablement call persistence authorization evidence consumption to the storage record writer enablement preview while keeping storage record writer enablements not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer grants not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer accepted receipt gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer authorization receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer authorization receipts not callable, disabled, ungranted, and unauthorized; storage record writer authorizations not callable, disabled, ungranted, and unauthorized; closed storage record writers not callable; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.528 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.527] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-grant-preview` to the v0.2.526 storage record writer accepted receipt gate predecessor.
- Added explicit storage record writer accepted receipt gate grant call persistence authorization evidence consumption to the storage record writer grant preview while keeping storage record writer grants not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer accepted receipt gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer authorization receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer authorization receipts not callable, disabled, ungranted, and unauthorized; storage record writer authorizations not callable, disabled, ungranted, and unauthorized; closed storage record writers not callable; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.527 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.526] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-preview` to the v0.2.525 storage record writer authorization receipt acceptance predecessor.
- Added explicit storage record writer authorization receipt acceptance accepted receipt gate call persistence authorization evidence consumption to the storage record writer accepted receipt gate preview while keeping storage record writer accepted receipt gates not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer authorization receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer authorization receipts not callable, disabled, ungranted, and unauthorized; storage record writer authorizations not callable, disabled, ungranted, and unauthorized; closed storage record writers not callable; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.526 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.525] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-preview` to the v0.2.524 storage record writer authorization receipt predecessor.
- Added explicit storage record writer authorization receipt accepted receipt gate call persistence authorization evidence consumption to the storage record writer authorization receipt acceptance preview while keeping storage record writer authorization receipt acceptances not callable, disabled, ungranted, unaccepted, and unauthorized; storage record writer authorization receipts not callable, disabled, ungranted, and unauthorized; storage record writer authorizations not callable, disabled, ungranted, and unauthorized; closed storage record writers not callable; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.525 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.524] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-preview` to the v0.2.523 storage record writer authorization predecessor.
- Added explicit storage record writer authorization accepted receipt gate call persistence authorization evidence consumption to the storage record writer authorization receipt preview while keeping storage record writer authorization receipts not callable, disabled, ungranted, and unauthorized; storage record writer authorizations not callable, disabled, ungranted, and unauthorized; closed storage record writers not callable; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.524 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.523] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-preview` to the v0.2.522 closed storage record writer implementation predecessor.
- Added explicit closed storage record writer implementation accepted receipt gate call persistence authorization evidence consumption to the storage record writer authorization preview while keeping storage record writer authorizations not callable, disabled, ungranted, and unauthorized; closed storage record writers not callable; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.523 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.522] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-storage-record-writer-implementation-preview` to the v0.2.521 storage record contract predecessor.
- Added explicit storage record contract accepted receipt gate call persistence authorization evidence consumption to the closed storage record writer implementation preview while keeping closed storage record writer implementations not callable, disabled, unimplemented, ungranted, and unauthorized; storage record contracts not writable; storage records unwritten; receipt writes disabled; receipt persistence disabled; consumer enablement disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.
- Recorded v0.2.522 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.521] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-contract-preview` to the v0.2.520 storage persistence gate audit predecessor.
- Added explicit storage persistence gate accepted receipt gate call persistence authorization evidence consumption to the storage record contract preview while keeping storage record contracts not writable, storage records unwritten, storage persistence gates not passed, receipt writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.521 as a targeted-validation-only follow-up after the v0.2.520 formal full Buildroot/QEMU checkpoint; the next formal full checkpoint remains v0.2.540 unless explicitly requested earlier.

## [0.2.520] - 2026-07-23

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-persistence-gate-audit-preview` to the v0.2.519 closed persistence implementation predecessor.
- Added explicit accepted receipt gate call persistence authorization evidence consumption to the storage persistence gate audit preview while keeping storage persistence gates not passed, storage persistence disabled, storage records absent, receipt writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Extended the restricted Docker image Go test timeout so the formal full checkpoint can cover the current full owner-package test volume instead of failing during the image build test stage.
- Captured operator-authorized restricted full Buildroot/QEMU serial smoke evidence for v0.2.520 in `output/full-smoke-report.json`, `output/full-smoke-report.md`, and `output/serial.log`, with expected boot markers present and host-root mutation, privileged containers, Docker socket mounts, host networking, and unrestricted QEMU networking disabled.

## [0.2.519] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-persistence-implementation-preview` to the v0.2.518 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt persistence authorization predecessor.
- Added explicit accepted receipt gate call persistence authorization evidence consumption to the closed persistence implementation preview while keeping closed persistence implementations not callable, disabled, unimplemented, ungranted, and unauthorized; persistence authorizations disabled; consumer enablement receipts absent, unpersisted, unaccepted, and disabled; receipt writes disabled; result persistence disabled; dispatch disabled; notification actions disabled; production ownership disabled; backend launch disabled; and host mutation disabled.

## [0.2.518] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-persistence-authorization-preview` to the v0.2.517 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt consumer enablement gate consumer authorization consumption gate acceptance call receipt call gate enablement grant authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt persistence authorization preview while keeping persistence authorizations not callable, disabled, ungranted, and unauthorized; consumer enablement receipts absent, unpersisted, unaccepted, and disabled; consumer enablement gates disabled; consumer enablement disabled; consumer authorization disabled; receipt consumption disabled; receipt acceptance disabled; receipts unaccepted; storage record writes disabled; path exposure disabled; backend launch disabled; and host mutation disabled.

## [0.2.517] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-preview` to the v0.2.516 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement gate predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement gate consumer enablement consumer authorization consumption gate acceptance call receipt call gate enablement grant authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt preview while keeping consumer enablement receipts absent, unpersisted, unaccepted, and disabled; consumer enablement gates disabled; consumer enablement disabled; consumer authorization disabled; receipt consumption disabled; receipt acceptance disabled; receipts unaccepted; consumption gates disabled; receipt acceptances disabled; call receipts disabled; accepted receipt gate calls disabled; accepted receipt gate call gates disabled; accepted receipt gate enablements disabled; accepted receipt gate grants disabled; accepted receipt gate authorizations disabled; accepted receipt gates disabled; writer calls disabled; storage record writes disabled; path exposure disabled; backend launch disabled; and host mutation disabled.

## [0.2.516] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-gate-preview` to the v0.2.515 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement consumer authorization consumption gate acceptance call receipt call gate enablement grant authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement gate preview while keeping consumer enablement gates disabled, consumer enablement disabled, consumer authorization disabled, receipt consumption disabled, receipt acceptance disabled, receipts unaccepted, consumption gates disabled, receipt acceptances disabled, call receipts disabled, accepted receipt gate calls disabled, accepted receipt gate call gates disabled, accepted receipt gate enablements disabled, accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.515] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-preview` to the v0.2.514 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer authorization predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer authorization consumption gate acceptance call receipt call gate enablement grant authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement preview while keeping consumer enablement disabled, consumer authorization disabled, receipt consumption disabled, receipt acceptance disabled, receipts unaccepted, consumption gates disabled, receipt acceptances disabled, call receipts disabled, accepted receipt gate calls disabled, accepted receipt gate call gates disabled, accepted receipt gate enablements disabled, accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.514] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-authorization-preview` to the v0.2.513 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumption gate predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumption gate acceptance call receipt call gate enablement grant authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer authorization preview while keeping consumer authorization disabled, receipt consumption disabled, receipt acceptance disabled, receipts unaccepted, consumption gates disabled, receipt acceptances disabled, call receipts disabled, accepted receipt gate calls disabled, accepted receipt gate call gates disabled, accepted receipt gate enablements disabled, accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.513] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumption-gate-preview` to the v0.2.512 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt acceptance predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt acceptance call receipt call gate enablement grant authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumption gate preview while keeping consumer authorization disabled, receipt consumption disabled, receipt acceptance disabled, receipts unaccepted, consumption gates disabled, receipt acceptances disabled, call receipts disabled, accepted receipt gate calls disabled, accepted receipt gate call gates disabled, accepted receipt gate enablements disabled, accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.512] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-acceptance-preview` to the v0.2.511 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt call gate enablement grant authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt acceptance preview while keeping receipt consumption disabled, receipt acceptance disabled, receipts unaccepted, call receipt acceptances disabled, call receipts disabled, accepted receipt gate calls disabled, accepted receipt gate call gates disabled, accepted receipt gate enablements disabled, accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.511] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-preview` to the v0.2.510 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call gate enablement grant authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt preview while keeping receipt acceptance disabled, call receipts disabled, accepted receipt gate calls disabled, accepted receipt gate call gates disabled, accepted receipt gate enablements disabled, accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.510] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-preview` to the v0.2.509 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call gate predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call gate enablement grant authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call preview while keeping accepted receipt gate calls disabled, accepted receipt gate call gates disabled, accepted receipt gate enablements disabled, accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.509] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-gate-preview` to the v0.2.508 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate enablement predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call gate enablement grant authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call gate preview while keeping accepted receipt gate call gates disabled, accepted receipt gate enablements disabled, accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.508] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-enablement-preview` to the v0.2.507 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate grant predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate grant authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate enablement preview while keeping accepted receipt gate enablements disabled, accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.507] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-grant-preview` to the v0.2.506 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate authorization predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate authorization acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate grant preview while keeping accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.506] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-authorization-preview` to the v0.2.505 storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate authorization preview while keeping accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.505] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-preview` to the v0.2.504 storage record writer call authorization receipt accepted receipt gate call authorization receipt acceptance predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt acceptance call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate preview while keeping accepted receipt gate calls disabled, acceptance calls disabled, call authorization receipts disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.504] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-acceptance-preview` to the v0.2.503 storage record writer call authorization receipt accepted receipt gate call authorization receipt predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization receipt call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt acceptance preview while keeping acceptance calls disabled, call authorization receipts disabled, call authorizations disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.503] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-preview` to the v0.2.502 storage record writer call authorization receipt accepted receipt gate call authorization predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call authorization call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization receipt preview while keeping call authorization receipts disabled, call authorizations disabled, call gates disabled, enablements disabled, grants disabled, authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.502] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-preview` to the v0.2.501 storage record writer call authorization receipt accepted receipt gate call gate predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate call gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call authorization preview while keeping call authorizations disabled, call gates disabled, enablements disabled, grants disabled, authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.501] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-gate-preview` to the v0.2.500 storage record writer call authorization receipt accepted receipt gate enablement predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate enablement grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate call gate preview while keeping call gates disabled, enablements disabled, grants disabled, authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.500] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-enablement-preview` to the v0.2.499 storage record writer call authorization receipt accepted receipt gate grant predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate grant authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate enablement preview while keeping accepted receipt gate enablements disabled, accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

### Fixed

- Corrected stale CLI test assertions for storage record writer call authorization receipt accepted receipt gate call-gate and call preview predecessor evidence tokens exposed by the full checkpoint test suite.

## [0.2.499] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-grant-preview` to the v0.2.498 storage record writer call authorization receipt accepted receipt gate authorization predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate authorization acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate grant preview while keeping accepted receipt gate grants disabled, accepted receipt gate authorizations disabled, accepted receipt gates disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.498] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-authorization-preview` to the v0.2.497 storage record writer call authorization receipt accepted receipt gate predecessor.
- Added explicit storage record writer call authorization receipt accepted receipt gate acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate authorization preview while keeping accepted receipt gate authorizations disabled, accepted receipt gates disabled, call authorization receipt acceptance disabled, call authorization receipts disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.497] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-preview` to the v0.2.496 storage record writer call authorization receipt acceptance predecessor.
- Added explicit storage record writer call authorization receipt acceptance call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt accepted receipt gate preview while keeping accepted receipt gates disabled, call authorization receipt acceptance disabled, call authorization receipts disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.496] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-acceptance-preview` to the v0.2.495 storage record writer call authorization receipt predecessor.
- Added explicit storage record writer call authorization receipt call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt acceptance preview while keeping call authorization receipt acceptance disabled, call authorization receipts disabled, call authorizations disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.495] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-preview` to the v0.2.494 storage record writer call authorization predecessor.
- Added explicit storage record writer call authorization call persistence authorization evidence consumption to the storage record writer call authorization receipt preview while keeping call authorization receipts disabled, call authorizations disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.494] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-preview` to the v0.2.493 storage record writer call gate predecessor.
- Added explicit storage record writer call gate call persistence authorization evidence consumption to the storage record writer call authorization preview while keeping call authorizations disabled, writer calls disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.493] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-gate-preview` to the v0.2.492 storage record writer enablement predecessor.
- Added explicit storage record writer call gate call persistence authorization evidence tracking while keeping call gates disabled, writer calls disabled, writer enablement disabled, writer grants disabled, storage record writes disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.492] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-enablement-preview` to the v0.2.491 storage record writer grant predecessor.
- Added explicit storage record writer enablement call persistence authorization evidence tracking while keeping writer enablement disabled, writer grants disabled, authorization receipts unaccepted, writer authorization ungranted, storage record writer calls disabled, storage records unwritten, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.491] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-grant-preview` to the v0.2.490 storage record writer accepted receipt gate predecessor.
- Added explicit storage record writer grant call persistence authorization evidence tracking while keeping writer grants disabled, authorization receipts unaccepted, writer authorization ungranted, storage record writer calls disabled, storage records unwritten, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.490] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-preview` to the v0.2.489 storage record writer authorization receipt acceptance predecessor.
- Added explicit storage record writer accepted receipt gate call persistence authorization evidence tracking while keeping accepted receipt gates disabled, authorization receipts unaccepted, writer authorization ungranted, storage record writer calls disabled, storage records unwritten, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.489] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-preview` to the v0.2.488 storage record writer authorization receipt predecessor.
- Added explicit storage record writer authorization receipt accepted receipt gate call persistence authorization evidence tracking while keeping authorization receipt acceptance disabled, accepted authorization receipts absent, writer authorization ungranted, storage record writer calls disabled, storage records unwritten, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.488] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-preview` to the v0.2.487 storage record writer authorization predecessor.
- Added explicit storage record writer authorization accepted receipt gate call persistence authorization evidence tracking while keeping authorization receipts disabled, authorization grants absent, writer authorization ungranted, storage record writer calls disabled, storage records unwritten, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.487] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-preview` to the v0.2.486 closed storage record writer implementation predecessor.
- Added explicit closed storage record writer implementation accepted receipt gate call persistence authorization evidence tracking while keeping writer authorization disabled, authorization grants absent, storage record writer calls disabled, storage records unwritten, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.486] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-storage-record-writer-implementation-preview` to the v0.2.485 storage record contract predecessor.
- Added explicit storage record contract accepted receipt gate call persistence authorization evidence tracking to the closed storage record writer implementation preview while keeping closed writers disabled, storage record writers uncalled, storage records unwritten, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.485] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-contract-preview` to the v0.2.484 storage persistence gate audit predecessor.
- Added explicit accepted receipt gate call persistence authorization evidence tracking to the storage record contract preview while keeping storage record contracts read-only, storage records unwritten, storage persistence gates unpassed, receipt persistence disabled, notification actions disabled, path exposure disabled, backend launch disabled, and host mutation disabled.

## [0.2.484] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-persistence-gate-audit-preview` to the v0.2.483 closed persistence implementation predecessor.
- Added explicit accepted receipt gate call persistence authorization evidence tracking to the storage persistence gate audit while keeping storage gate passage, storage persistence, receipt writes, notification actions, production ownership, backend launch, path exposure, and host mutation disabled.

## [0.2.483] - 2026-07-22

### Changed

- Reconnected `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-persistence-implementation-preview` to the v0.2.482 accepted receipt gate call receipt consumer enablement receipt persistence authorization predecessor.
- Updated the closed persistence implementation evidence checks to consume the storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt persistence authorization source while keeping closed persistence implementation, receipt persistence authorization, receipt persistence, storage writes, notification actions, production ownership, backend launch, path exposure, and host mutation disabled.

## [0.2.482] - 2026-07-22

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-persistence-authorization-preview`, a read-only Go owner preview that consumes the v0.2.481 accepted receipt gate call receipt consumer enablement receipt preview before modeling future accepted receipt gate call receipt consumer enablement receipt persistence authorization boundaries.
- Wired the accepted receipt gate call receipt consumer enablement receipt persistence authorization preview into the Runtime CLI while keeping persistence authorization disabled and ungranted, consumer enablement receipts absent and unpersisted, storage record writes disabled, receipt persistence disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.

## [0.2.481] - 2026-07-22

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-receipt-preview`, a read-only Go owner preview that consumes the v0.2.480 accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement gate preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement receipt preview into the Runtime CLI while keeping consumer enablement receipts absent, unpersisted, unaccepted, and disabled; consumer enablement gates disabled; KDE consumers disabled; Runtime consumers disabled; storage record writes disabled; receipt persistence disabled; notification actions disabled; path exposure disabled; production ownership disabled; backend launch disabled; and host mutation disabled.

## [0.2.480] - 2026-07-22

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-gate-preview`, a read-only Go owner preview that consumes the v0.2.479 accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement gate boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement gate preview into the Runtime CLI while keeping consumer enablement gates disabled, consumer enablement disabled, KDE consumers disabled, Runtime consumers disabled, consumer authorization disabled, receipt consumption disabled, receipts unconsumed, storage record writes disabled, receipt persistence disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.

## [0.2.479] - 2026-07-22

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-enablement-preview`, a read-only Go owner preview that consumes the v0.2.478 accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer authorization preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer enablement preview into the Runtime CLI while keeping consumer enablement disabled, KDE consumers disabled, Runtime consumers disabled, consumer authorization disabled, receipt consumption disabled, receipts unconsumed, storage record writes disabled, receipt persistence disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.

## [0.2.478] - 2026-07-22

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumer-authorization-preview`, a read-only Go owner preview that consumes the v0.2.477 accepted receipt gate call authorization receipt accepted receipt gate call receipt consumption gate preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer authorization boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumer authorization preview into the Runtime CLI while keeping consumer authorization disabled, receipt consumption disabled, receipts unconsumed, receipt acceptance disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.

## [0.2.477] - 2026-07-22

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-consumption-gate-preview`, a read-only Go owner preview that consumes the v0.2.476 accepted receipt gate call authorization receipt accepted receipt gate call receipt acceptance preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate call receipt consumption gate boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt consumption gate preview into the Runtime CLI while keeping receipt consumption disabled, receipts unconsumed, receipt acceptance disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.

## [0.2.476] - 2026-07-22

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-acceptance-preview`, a read-only Go owner preview that consumes the v0.2.475 accepted receipt gate call authorization receipt accepted receipt gate call receipt preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate call receipt acceptance boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt acceptance preview into the Runtime CLI while keeping receipt acceptance disabled, receipts unaccepted, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.

## [0.2.475] - 2026-07-22

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-receipt-preview`, a read-only Go owner preview that consumes the v0.2.474 accepted receipt gate call authorization receipt accepted receipt gate call preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate call receipt boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call receipt preview into the Runtime CLI while keeping call receipts disabled, calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.

## [0.2.474] - 2026-07-22

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-preview`, a read-only Go owner preview that consumes the v0.2.473 accepted receipt gate call authorization receipt accepted receipt gate call-gate preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate call boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call preview into the Runtime CLI while keeping calls disabled, call-gates disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.

## [0.2.473] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-call-gate-preview`, a read-only Go owner preview that consumes the v0.2.472 accepted receipt gate call authorization receipt accepted receipt gate enablement preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate call-gate boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate call-gate preview into the Runtime CLI while keeping call-gate calls disabled, call-gates disabled, enablement disabled, grants absent, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Advanced the Codex-owned current mainline to the next accepted receipt gate call authorization receipt accepted receipt gate call preview while preserving call-gate predecessor evidence.

## [0.2.472] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-enablement-preview`, a read-only Go owner preview that consumes the v0.2.471 accepted receipt gate call authorization receipt accepted receipt gate grant preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate enablement boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate enablement preview into the Runtime CLI while keeping enablement calls disabled, enablement disabled, grants absent, grant calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Advanced the Codex-owned current mainline to the next accepted receipt gate call authorization receipt accepted receipt gate call-gate preview while preserving enablement predecessor evidence.

## [0.2.471] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-grant-preview`, a read-only Go owner preview that consumes the v0.2.470 accepted receipt gate call authorization receipt accepted receipt gate authorization preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate grant boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate grant preview into the Runtime CLI while keeping grant calls disabled, grants absent, authorization calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Advanced the Codex-owned current mainline to the next accepted receipt gate call authorization receipt accepted receipt gate enablement preview while preserving grant predecessor evidence.

## [0.2.470] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-authorization-preview`, a read-only Go owner preview that consumes the v0.2.469 accepted receipt gate call authorization receipt accepted receipt gate preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate authorization boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate authorization preview into the Runtime CLI while keeping authorization calls disabled, grants absent, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.

### Fixed

- Raised the Docker build full Go test timeout so milestone full smoke can complete the long owner/CLI safety-gate test suite instead of failing at the default 10-minute test timeout.
- Replaced full JSON encoding in Runtime owner backend-term validation with recursive field scanning so deeply nested safety-gate previews remain checked without exhausting the restricted full smoke test window.

## [0.2.469] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-accepted-receipt-gate-preview`, a read-only Go owner preview that consumes the v0.2.468 accepted receipt gate call authorization receipt acceptance preview before modeling future accepted receipt gate call authorization receipt accepted receipt gate boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt accepted receipt gate preview into the Runtime CLI while keeping accepted receipt gate calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.469 as targeted-validation-only follow-up work under the twenty-small-version testing cadence; formal full Buildroot/QEMU smoke was not rerun for this non-boundary checkpoint.

### Fixed

- Restored the accepted receipt gate call authorization receipt acceptance predecessor evidence in the current mainline so the v0.2.469 preview consumes the intended v0.2.468 boundary instead of failing closed due to missing documentation evidence.

## [0.2.468] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-acceptance-preview`, a read-only Go owner preview that consumes the v0.2.467 accepted receipt gate call authorization receipt preview before modeling future accepted receipt gate call authorization receipt acceptance boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt acceptance preview into the Runtime CLI while keeping accepted receipt gate call authorization receipt acceptance calls disabled, accepted receipt gate call authorization receipt calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.468 as targeted-validation-only follow-up work under the twenty-small-version testing cadence; formal full Buildroot/QEMU smoke was not rerun for this non-boundary checkpoint.

## [0.2.467] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-receipt-preview`, a read-only Go owner preview that consumes the v0.2.466 accepted receipt gate call authorization preview before modeling future accepted receipt gate call authorization receipt boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization receipt preview into the Runtime CLI while keeping accepted receipt gate call authorization receipt calls disabled, accepted receipt gate call authorization calls disabled, accepted receipt gate call gate calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.467 as targeted-validation-only follow-up work under the twenty-small-version testing cadence; formal full Buildroot/QEMU smoke was not rerun for this non-boundary checkpoint.

## [0.2.466] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-authorization-preview`, a read-only Go owner preview that consumes the v0.2.465 storage record writer call authorization receipt accepted receipt gate call gate preview before modeling future accepted receipt gate call authorization boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call authorization preview into the Runtime CLI while keeping accepted receipt gate call authorization calls disabled, accepted receipt gate call gate calls disabled, accepted receipt gate enablement calls disabled, accepted receipt gate grant calls disabled, accepted receipt gate authorization calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.466 as targeted-validation-only follow-up work under the twenty-small-version testing cadence; formal full Buildroot/QEMU smoke was not rerun for this non-boundary checkpoint.

## [0.2.465] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-call-gate-preview`, a read-only Go owner preview that consumes the v0.2.464 storage record writer call authorization receipt accepted receipt gate enablement preview before modeling future accepted receipt gate call gate boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate call gate preview into the Runtime CLI while keeping accepted receipt gate call gate calls disabled, accepted receipt gate enablement calls disabled, accepted receipt gate grant calls disabled, accepted receipt gate authorization calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.465 as targeted-validation-only follow-up work under the twenty-small-version testing cadence; formal full Buildroot/QEMU smoke was not rerun for this non-boundary checkpoint.

## [0.2.464] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-enablement-preview`, a read-only Go owner preview that consumes the v0.2.463 storage record writer call authorization receipt accepted receipt gate grant preview before modeling future accepted receipt gate enablement boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate enablement preview into the Runtime CLI while keeping accepted receipt gate enablement disabled, accepted receipt gate grant calls disabled, accepted receipt gate authorization calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.464 as targeted-validation-only follow-up work under the twenty-small-version testing cadence; formal full Buildroot/QEMU smoke was not rerun for this non-boundary checkpoint.

## [0.2.463] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-grant-preview`, a read-only Go owner preview that consumes the v0.2.462 storage record writer call authorization receipt accepted receipt gate authorization preview before modeling future accepted receipt gate grant boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate grant preview into the Runtime CLI while keeping accepted receipt gate grant calls disabled, accepted receipt gate authorization calls disabled, accepted receipt gate calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.463 as targeted-validation-only follow-up work under the twenty-small-version testing cadence; formal full Buildroot/QEMU smoke was not rerun for this non-boundary checkpoint.

## [0.2.462] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-authorization-preview`, a read-only Go owner preview that consumes the v0.2.461 storage record writer call authorization receipt accepted receipt gate preview before modeling future accepted receipt gate authorization boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate authorization preview into the Runtime CLI while keeping accepted receipt gate authorization calls disabled, accepted receipt gate calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.462 as targeted-validation-only follow-up work under the twenty-small-version testing cadence; formal full Buildroot/QEMU smoke was not rerun for this non-boundary checkpoint.

## [0.2.461] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-accepted-receipt-gate-preview`, a read-only Go owner preview that consumes the v0.2.460 storage record writer call authorization receipt acceptance preview before modeling future storage record writer call authorization receipt accepted receipt gate boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt accepted receipt gate preview into the Runtime CLI while keeping accepted receipt gate calls disabled, call authorization receipt acceptance disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.461 as targeted-validation-only follow-up work under the twenty-small-version testing cadence; formal full Buildroot/QEMU smoke was not rerun for this non-boundary checkpoint.

## [0.2.460] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-acceptance-preview`, a read-only Go owner preview that consumes the v0.2.459 storage record writer call authorization receipt preview before modeling future storage record writer call authorization receipt acceptance boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt acceptance preview into the Runtime CLI while keeping call authorization receipt acceptance calls disabled, call authorization receipt calls disabled, writer calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.460 as a twentieth small-version boundary with operator-authorized restricted full Buildroot/QEMU smoke passing after targeted validation.

## [0.2.459] - 2026-07-21

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-receipt-preview`, a read-only Go owner preview that consumes the v0.2.458 storage record writer call authorization preview before modeling future storage record writer call authorization receipt boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization receipt preview into the Runtime CLI while keeping call authorization receipt calls disabled, writer calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.459 as targeted-validation-only follow-up work under the current twenty-small-version testing cadence; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.458] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-authorization-preview`, a read-only Go owner preview that consumes the v0.2.457 storage record writer call gate preview plus writer enablement, writer grant, accepted receipt gate, authorization receipt acceptance, authorization receipt, writer authorization, closed writer implementation, storage record contract, and storage persistence gate evidence before modeling future storage record writer call authorization boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call authorization preview into the Runtime CLI while keeping writer call authorization disabled, writer calls disabled, call gates disabled, writer enablement disabled, writer grants disabled, accepted receipt gates disabled, authorization receipts unaccepted, writer authorization ungranted, closed writer calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.458 as targeted-validation-only follow-up work under the current twenty-small-version testing cadence; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.457] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-call-gate-preview`, a read-only Go owner preview that consumes the v0.2.456 storage record writer enablement preview plus writer grant, accepted receipt gate, authorization receipt acceptance, authorization receipt, writer authorization, closed writer implementation, storage record contract, and storage persistence gate evidence before modeling future storage record writer call gate boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer call gate preview into the Runtime CLI while keeping writer calls disabled, writer call gates disabled, writer enablement disabled, writer grants disabled, accepted receipt gates disabled, authorization receipts unaccepted, writer authorization ungranted, closed writer calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.457 as targeted-validation-only follow-up work under the current twenty-small-version testing cadence; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.456] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-enablement-preview`, a read-only Go owner preview that consumes the v0.2.455 storage record writer grant preview plus accepted receipt gate, authorization receipt acceptance, authorization receipt, writer authorization, closed writer implementation, storage record contract, and storage persistence gate evidence before modeling future storage record writer enablement boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer enablement preview into the Runtime CLI while keeping writer enablement disabled, writer grants disabled, accepted receipt gates disabled, authorization receipts unaccepted, writer authorization ungranted, closed writer calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.456 as targeted-validation-only follow-up work under the current twenty-small-version testing cadence; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.455] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-grant-preview`, a read-only Go owner preview that consumes the v0.2.454 storage record writer accepted receipt gate preview plus authorization receipt acceptance, authorization receipt, writer authorization, closed writer implementation, storage record contract, and storage persistence gate evidence before modeling future storage record writer grant boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer grant preview into the Runtime CLI while keeping writer grants disabled, accepted receipt gates disabled, authorization receipts unaccepted, writer authorization ungranted, closed writer calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.455 as targeted-validation-only follow-up work under the current twenty-small-version testing cadence; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.454] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-accepted-receipt-gate-preview`, a read-only Go owner preview that consumes the v0.2.453 storage record writer authorization receipt acceptance preview plus writer authorization receipt, writer authorization, closed writer implementation, storage record contract, and storage persistence gate evidence before modeling future accepted receipt gate boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer accepted receipt gate preview into the Runtime CLI while keeping accepted receipt gates disabled, authorization receipts unaccepted, writer authorization ungranted, closed writer calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.454 as targeted-validation-only follow-up work under the current twenty-small-version testing cadence; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.453] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-acceptance-preview`, a read-only Go owner preview that consumes the v0.2.452 storage record writer authorization receipt preview plus the writer authorization, closed writer implementation, storage record contract, and storage persistence gate before modeling future writer authorization receipt acceptance boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer authorization receipt acceptance preview into the Runtime CLI while keeping authorization receipt acceptance disabled, accepted authorization receipts absent, writer authorization ungranted, closed writer calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.453 as targeted-validation-only follow-up work under the current twenty-small-version testing cadence; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.452] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-receipt-preview`, a read-only Go owner preview that consumes the v0.2.451 storage record writer authorization preview plus the closed storage record writer implementation, storage record contract, and storage persistence gate before modeling future writer authorization receipt boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer authorization receipt preview into the Runtime CLI while keeping authorization receipt calls disabled, authorization receipt grants absent, writer authorization ungranted, closed writer calls disabled, storage record writes disabled, receipt persistence disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.452 as targeted-validation-only follow-up work under the current twenty-small-version testing cadence; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.451] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-writer-authorization-preview`, a read-only Go owner preview that consumes the v0.2.450 closed storage record writer implementation preview plus the storage record contract and storage persistence gate before modeling future storage record writer authorization.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record writer authorization preview into the Runtime CLI while keeping storage record writer authorization disabled, authorization grants absent, closed writer calls disabled, storage record writes disabled, receipt persistence disabled, receipt writes disabled, consumer enablement disabled, notification actions disabled, path exposure disabled, production ownership disabled, backend launch disabled, and host mutation disabled.
- Recorded v0.2.451 as targeted-validation-only follow-up work under the current twenty-small-version testing cadence; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.450] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-storage-record-writer-implementation-preview`, a read-only Go owner preview that consumes the v0.2.449 storage record contract plus storage gate, closed persistence implementation, and persistence authorization predecessor evidence before modeling future closed storage record writer implementation boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed storage record writer implementation preview into the Runtime CLI while keeping writer calls, writer implementation enablement, storage record writes, storage record contract writability, storage gate passage, storage persistence, receipt persistence, receipt writes, consumer enablement, consumer authorization grants, receipt consumption, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, path exposure, production ownership, Runtime writes, desktop side effects, backend launch, unsafe-data exposure, and host mutation disabled.
- Recorded v0.2.450 as targeted-validation-only follow-up work under the current twenty-small-version testing cadence; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.449] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-record-contract-preview`, a read-only Go owner preview that consumes the v0.2.448 storage persistence gate audit plus the direct receipt consumer enablement receipt predecessor chain before modeling future receipt storage record contracts.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage record contract preview into the Runtime CLI while keeping storage gate passage, storage record writes, storage persistence, receipt persistence, receipt writes, consumer enablement, consumer authorization grants, receipt consumption, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, path exposure, production ownership, Runtime writes, desktop side effects, backend launch, unsafe-data exposure, and host mutation disabled.
- Recorded v0.2.449 as targeted-validation-only follow-up work after the operator-authorized v0.2.444 restricted container heavy smoke; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.448] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-storage-persistence-gate-audit-preview`, a read-only Go owner preview that consumes the v0.2.447 notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed persistence implementation preview before modeling future receipt storage persistence gate boundaries.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt storage persistence gate audit preview into the Runtime CLI while keeping storage gate passage, storage persistence, receipt persistence, receipt writes, consumer enablement, consumer authorization grants, receipt consumption, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, path exposure, production ownership, Runtime writes, desktop side effects, backend launch, unsafe-data exposure, and host mutation disabled.
- Recorded v0.2.448 as targeted-validation-only follow-up work after the operator-authorized v0.2.444 restricted container heavy smoke; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.447] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-persistence-implementation-preview`, a read-only Go owner preview that consumes the v0.2.446 notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt persistence authorization preview before modeling future closed persistence receipt implementation shapes.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed persistence implementation preview into the Runtime CLI while keeping closed persistence implementation, receipt persistence authorization grants, receipt persistence, receipt writes, consumer enablement, consumer authorization grants, receipt consumption, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.
- Recorded v0.2.447 as targeted-validation-only follow-up work after the operator-authorized v0.2.444 restricted container heavy smoke; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.446] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-persistence-authorization-preview`, a read-only Go owner preview that consumes the v0.2.445 notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation preview before modeling future receipt persistence authorization.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt persistence authorization preview into the Runtime CLI while keeping persistence authorization grants, receipt persistence, receipt writes, consumer enablement, consumer authorization grants, receipt consumption, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.
- Recorded v0.2.446 as targeted-validation-only follow-up work after the operator-authorized v0.2.444 restricted container heavy smoke; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.445] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-closed-implementation-preview`, a read-only Go owner preview that consumes the v0.2.444 notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt enablement audit before modeling future closed receipt implementation shapes.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt closed implementation preview into the Runtime CLI while keeping receipt implementation calls, receipt enablement, receipt grants, accepted receipt gates, acceptance gates, receipt acceptance, receipt persistence, receipt writes, consumer enablement, consumer authorization grants, receipt consumption, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.
- Recorded v0.2.445 as targeted-validation-only follow-up work after the operator-authorized v0.2.444 restricted container heavy smoke; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.444] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-enablement-audit-preview`, a read-only Go owner audit that consumes the v0.2.443 notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt grant audit before modeling future receipt enablement for persistence receipt consumer enablement receipts.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt enablement audit into the Runtime CLI while keeping receipt enablement, receipt grants, accepted receipt gates, acceptance gates, receipt acceptance, enablement receipt persistence, enablement receipt writes, consumer enablement, consumer authorization grants, receipt consumption, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.
- Recorded v0.2.444 as targeted-validation-only follow-up work after the operator-authorized v0.2.442 restricted container heavy smoke; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.443] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-grant-audit-preview`, a read-only Go owner audit that consumes the v0.2.442 notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt accepted receipt gate audit before modeling future receipt grants for persistence receipt consumer enablement receipts.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt grant audit into the Runtime CLI while keeping receipt grants, accepted receipt gates, acceptance gates, receipt acceptance, enablement receipt persistence, enablement receipt writes, consumer enablement, consumer authorization grants, receipt consumption, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.
- Recorded v0.2.443 as targeted-validation-only follow-up work after the operator-authorized v0.2.442 restricted container heavy smoke; formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and was not rerun for this non-boundary checkpoint.

## [0.2.442] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-accepted-receipt-gate-audit-preview`, a read-only Go owner audit that consumes the v0.2.441 notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt acceptance gate audit before modeling future accepted receipt gates for persistence receipt consumer enablement receipts.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt accepted receipt gate audit into the Runtime CLI while keeping accepted receipt gates, acceptance gates, receipt acceptance, enablement receipt persistence, enablement receipt writes, consumer enablement, consumer authorization grants, receipt consumption, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.442` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, Runtime activation smoke, and the restricted product smoke packet while keeping privileged containers, host networking, Docker socket mounts inside containers, broad host mounts, backend launch, QEMU execution, and host-root mutation disabled.

## [0.2.441] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-acceptance-gate-audit-preview`, a read-only Go owner audit that consumes the v0.2.440 notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt acceptance authorization audit before modeling future acceptance gates for persistence receipt consumer enablement receipts.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt acceptance gate audit into the Runtime CLI while keeping acceptance gates, acceptance authorization grants, receipt acceptance, enablement receipt persistence, enablement receipt writes, persistence authorization grants, writer authorization grants, consumer enablement, consumer authorization grants, receipt consumption, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.
- Recorded the operator-authorized restricted full Buildroot/QEMU serial smoke evidence for v0.2.440 in `output/full-smoke-report.json`, `output/full-smoke-report.md`, and `output/serial.log`; v0.2.441 itself uses targeted validation only and does not rerun formal full Buildroot/QEMU smoke.

## [0.2.440] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-acceptance-authorization-audit-preview`, a read-only Go owner audit that consumes the v0.2.439 notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt persistence authorization audit before modeling future acceptance authorization for persistence receipt consumer enablement receipts.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt acceptance authorization audit into the Runtime CLI while keeping acceptance authorization grants, receipt acceptance, enablement receipt persistence, enablement receipt writes, persistence authorization grants, writer authorization grants, consumer enablement, consumer authorization grants, receipt consumption, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.
- Because v0.2.440 is a twentieth small-version boundary, captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.440` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, Runtime activation smoke, and the restricted product smoke packet through the prebuilt Runtime CLI while formal Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb` and requires separate operator authorization.

## [0.2.439] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-persistence-authorization-audit-preview`, a read-only Go owner audit that consumes the v0.2.438 notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt writer authorization audit before modeling future persistence authorization for persistence receipt consumer enablement receipts.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt persistence authorization audit into the Runtime CLI while keeping persistence authorization grants, enablement receipt persistence, enablement receipt writes, writer authorization grants, consumer enablement, consumer authorization grants, receipt consumption, receipt acceptance, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.439` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, Runtime activation smoke, and the restricted product smoke packet through the prebuilt Runtime CLI while the `go run` wrapper remains blocked by the container `noexec` temporary filesystem and formal Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb`.

## [0.2.438] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-writer-authorization-audit-preview`, a read-only Go owner audit that consumes the v0.2.437 notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt audit before modeling future writer authorization for persistence receipt consumer enablement receipts.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt writer authorization audit into the Runtime CLI while keeping writer authorization grants, receipt consumer enablement receipt writes, receipt consumer enablement, consumer authorization grants, receipt consumption, receipt acceptance, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.438` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, Runtime activation smoke, and the restricted product smoke packet through the prebuilt Runtime CLI while the `go run` wrapper remains blocked by the container `noexec` temporary filesystem and formal Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb`.

## [0.2.437] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-receipt-audit-preview`, a read-only Go owner audit that consumes the v0.2.436 notification action request-object dispatch dry-run result persistence receipt consumer enablement gate audit before modeling future persistence receipt consumer enablement receipts.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement receipt audit into the Runtime CLI while keeping consumer enablement receipts absent, receipt writes, consumer enablement, consumer authorization grants, receipt consumption, receipt acceptance, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.

## [0.2.436] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-enablement-gate-audit-preview`, a read-only Go owner audit that consumes the v0.2.435 notification action request-object dispatch dry-run result persistence receipt consumer authorization audit before modeling future persistence receipt consumer enablement gates.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer enablement gate audit into the Runtime CLI while keeping consumer enablement, consumer authorization grants, receipt consumption, receipt acceptance, receipt writes, dry-run result persistence, result visibility persistence, dispatch execution, request-object creation, notification actions, notification delivery, production ownership, Runtime writes, desktop side effects, path exposure, backend launch, unsafe-data exposure, and host mutation disabled.

## [0.2.435] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumer-authorization-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.434 notification action request-object dispatch dry-run result persistence receipt consumption gate audit, and the inherited receipt acceptance, persistence receipt, persistence authorization, visibility, dispatch, request-object, notification action, delivery grant, delivery execution authorization, and consent receipt consumer chain before modeling future dry-run result persistence receipt consumer authorization readiness.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumer authorization audit into the Runtime CLI while keeping consumer authorization disabled, receipt consumption disabled, receipt acceptance disabled, persistence receipts absent, receipt writes, result writes, result visibility persistence, dry-run execution, dispatch authorization grants, request object creation, request object dispatch, request object persistence, notification actions, action cards, delivery, notification sending, Notification Center events, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.434] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-consumption-gate-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.433 notification action request-object dispatch dry-run result persistence receipt acceptance audit, and the inherited persistence receipt, persistence authorization, visibility, dispatch, request-object, notification action, delivery grant, delivery execution authorization, and consent receipt consumer chain before modeling future dry-run result persistence receipt consumption gate readiness.
- Wired the notification action request-object dispatch dry-run result persistence receipt consumption gate audit into the Runtime CLI while keeping receipt consumption disabled, receipt acceptance disabled, persistence receipts absent, receipt writes, result writes, result visibility persistence, dry-run execution, dispatch authorization grants, request object creation, request object dispatch, request object persistence, notification actions, action cards, delivery, notification sending, Notification Center events, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.433] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-acceptance-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.432 notification action request-object dispatch dry-run result persistence receipt audit, and the inherited persistence authorization, visibility, dispatch, request-object, notification action, delivery grant, delivery execution authorization, and consent receipt consumer chain before modeling future dry-run result persistence receipt acceptance readiness.
- Wired the notification action request-object dispatch dry-run result persistence receipt acceptance audit into the Runtime CLI while keeping receipt acceptance disabled, persistence receipts absent, receipt writes, receipt consumption, result writes, result visibility persistence, dry-run execution, dispatch authorization grants, request object creation, request object dispatch, request object persistence, notification actions, action cards, delivery, notification sending, Notification Center events, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.432] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-receipt-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.431 notification action request-object dispatch dry-run result persistence authorization audit, v0.2.430 notification action request-object dispatch dry-run result visibility audit, v0.2.429 notification action request-object dispatch dry-run audit, v0.2.428 notification action request-object dispatch authorization audit, v0.2.427 notification action request-object audit, v0.2.426 notification action enablement audit, v0.2.425 notification delivery grant audit, v0.2.424 notification delivery execution authorization audit, and the v0.2.420-v0.2.423 consent receipt consumer chain before modeling future dry-run result persistence receipt readiness.
- Wired the notification action request-object dispatch dry-run result persistence receipt audit into the Runtime CLI while keeping persistence receipts absent, receipt writes, result writes, result visibility persistence, dry-run execution, dispatch authorization grants, request object creation, request object dispatch, request object persistence, notification actions, action cards, delivery, notification sending, Notification Center events, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.432` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, Runtime activation smoke, and the restricted product smoke packet while formal Buildroot/QEMU and loopback SSH smokes remain gated by `scripts/full_smoke.rb`.

## [0.2.431] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-persistence-authorization-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.430 notification action request-object dispatch dry-run result visibility audit, v0.2.429 notification action request-object dispatch dry-run audit, v0.2.428 notification action request-object dispatch authorization audit, v0.2.427 notification action request-object audit, v0.2.426 notification action enablement audit, v0.2.425 notification delivery grant audit, v0.2.424 notification delivery execution authorization audit, and the v0.2.420-v0.2.423 consent receipt consumer chain before modeling future dry-run result persistence authorization readiness.
- Wired the notification action request-object dispatch dry-run result persistence authorization audit into the Runtime CLI while keeping result persistence authorization grants, operator persistence approval, dry-run result persistence, result visibility persistence, dry-run execution, dispatch authorization grants, request object creation, request object dispatch, request object persistence, notification actions, action cards, delivery, notification sending, Notification Center events, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.430] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-result-visibility-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.429 notification action request-object dispatch dry-run audit, v0.2.428 notification action request-object dispatch authorization audit, v0.2.427 notification action request-object audit, v0.2.426 notification action enablement audit, v0.2.425 notification delivery grant audit, v0.2.424 notification delivery execution authorization audit, v0.2.423 notification delivery consent collection receipt consumer enablement gate audit, v0.2.422 notification delivery consent collection receipt consumer authorization audit, v0.2.421 notification delivery consent collection receipt consumption gate audit, and v0.2.420 notification delivery consent collection receipt acceptance audit before modeling future KDE and Runtime diagnostics redacted dry-run result visibility readiness.
- Wired the notification action request-object dispatch dry-run result visibility audit into the Runtime CLI while keeping user visibility, result visibility persistence, dry-run result persistence, dry-run execution, dispatch authorization grants, request object creation, request object dispatch, request object persistence, notification actions, action cards, delivery grant issuance, delivery authorization grants, delivery execution enablement, notification sending, Notification Center events, consent collection, consent receipt creation, consent receipt acceptance, consent receipt consumption, consumer authorization grants, consumer enablement, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.430` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, and Runtime activation smoke while formal Buildroot/QEMU and loopback SSH smokes remain gated by `scripts/full_smoke.rb`.

## [0.2.429] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-dry-run-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.428 notification action request-object dispatch authorization audit, v0.2.427 notification action request-object audit, v0.2.426 notification action enablement audit, v0.2.425 notification delivery grant audit, v0.2.424 notification delivery execution authorization audit, v0.2.423 notification delivery consent collection receipt consumer enablement gate audit, v0.2.422 notification delivery consent collection receipt consumer authorization audit, v0.2.421 notification delivery consent collection receipt consumption gate audit, and v0.2.420 notification delivery consent collection receipt acceptance audit before modeling future KDE Notification Center action request-object dispatch dry-run planning readiness.
- Wired the notification action request-object dispatch dry-run audit into the Runtime CLI while keeping dispatch dry-run execution, dry-run result persistence, dispatch authorization grants, request object creation, request object dispatch, request object persistence, notification actions, action cards, delivery grant issuance, delivery authorization grants, delivery execution enablement, notification sending, Notification Center events, consent collection, consent receipt creation, consent receipt acceptance, consent receipt consumption, consumer authorization grants, consumer enablement, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.429` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, and Runtime activation smoke while formal Buildroot/QEMU and loopback SSH smokes remain gated by `scripts/full_smoke.rb`.

## [0.2.428] - 2026-07-20

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-dispatch-authorization-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.427 notification action request-object audit, v0.2.426 notification action enablement audit, v0.2.425 notification delivery grant audit, v0.2.424 notification delivery execution authorization audit, v0.2.423 notification delivery consent collection receipt consumer enablement gate audit, v0.2.422 notification delivery consent collection receipt consumer authorization audit, v0.2.421 notification delivery consent collection receipt consumption gate audit, and v0.2.420 notification delivery consent collection receipt acceptance audit before modeling future KDE Notification Center action request-object dispatch authorization readiness.
- Wired the notification action request-object dispatch authorization audit into the Runtime CLI while keeping dispatch authorization grants, request object creation, request object dispatch, request object persistence, notification actions, action cards, delivery grant issuance, delivery authorization grants, delivery execution enablement, notification sending, Notification Center events, consent collection, consent receipt creation, consent receipt acceptance, consent receipt consumption, consumer authorization grants, consumer enablement, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.428` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, and Runtime activation smoke while formal Buildroot/QEMU and loopback SSH smokes remain gated by `scripts/full_smoke.rb`.

## [0.2.427] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-request-object-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.426 notification action enablement audit, v0.2.425 notification delivery grant audit, v0.2.424 notification delivery execution authorization audit, v0.2.423 notification delivery consent collection receipt consumer enablement gate audit, v0.2.422 notification delivery consent collection receipt consumer authorization audit, v0.2.421 notification delivery consent collection receipt consumption gate audit, and v0.2.420 notification delivery consent collection receipt acceptance audit before modeling future KDE Notification Center action request-object readiness.
- Wired the notification action request-object audit into the Runtime CLI while keeping request object creation, request object dispatch, request object persistence, notification actions, action cards, delivery grant issuance, delivery authorization grants, delivery execution enablement, notification sending, Notification Center events, consent collection, consent receipt creation, consent receipt acceptance, consent receipt consumption, consumer authorization grants, consumer enablement, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.427` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, and Runtime activation smoke while formal Buildroot/QEMU and loopback SSH smokes remain gated by `scripts/full_smoke.rb`.

## [0.2.426] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-action-enablement-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.425 notification delivery grant audit, v0.2.424 notification delivery execution authorization audit, v0.2.423 notification delivery consent collection receipt consumer enablement gate audit, v0.2.422 notification delivery consent collection receipt consumer authorization audit, v0.2.421 notification delivery consent collection receipt consumption gate audit, and v0.2.420 notification delivery consent collection receipt acceptance audit before modeling future KDE Notification Center action enablement readiness.
- Wired the notification action enablement audit into the Runtime CLI while keeping notification actions, action cards, request object creation, delivery grant issuance, delivery authorization grants, delivery execution enablement, notification sending, Notification Center events, consent collection, consent receipt creation, consent receipt acceptance, consent receipt consumption, consumer authorization grants, consumer enablement, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.425] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.424 notification delivery execution authorization audit, v0.2.423 notification delivery consent collection receipt consumer enablement gate audit, v0.2.422 notification delivery consent collection receipt consumer authorization audit, v0.2.421 notification delivery consent collection receipt consumption gate audit, and v0.2.420 notification delivery consent collection receipt acceptance audit before modeling future KDE Notification Center delivery grant readiness.
- Realigned the existing notification delivery grant Runtime CLI command to the execution-authorization-backed grant audit while keeping delivery grant issuance, delivery authorization grants, delivery execution enablement, notification sending, Notification Center events, notification actions, action cards, consent collection, consent receipt creation, consent receipt acceptance, consent receipt consumption, consumer authorization grants, consumer enablement, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.424] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-execution-authorization-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.423 notification delivery consent collection receipt consumer enablement gate audit, v0.2.422 notification delivery consent collection receipt consumer authorization audit, v0.2.421 notification delivery consent collection receipt consumption gate audit, and v0.2.420 notification delivery consent collection receipt acceptance audit before modeling future KDE Notification Center delivery execution authorization readiness.
- Wired the notification delivery execution authorization audit into the Runtime CLI while keeping delivery execution authorization grants, delivery execution enablement, delivery grant issuance, delivery authorization grants, notification sending, Notification Center events, notification actions, action cards, consent collection, consent receipt creation, consent receipt acceptance, consent receipt consumption, consumer authorization grants, consumer enablement, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.423] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-enablement-gate-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.422 notification delivery consent collection receipt consumer authorization audit, v0.2.421 notification delivery consent collection receipt consumption gate audit, and v0.2.420 notification delivery consent collection receipt acceptance audit before modeling future KDE Notification Center consent collection receipt consumer enablement gate readiness.
- Wired the notification delivery consent collection receipt consumer enablement gate audit into the Runtime CLI while keeping consent collection, consent persistence, consent receipt creation, consent receipt persistence, consent receipt acceptance, consent receipt consumption, consumer authorization grants, KDE consumer authorization, Runtime consumer authorization, KDE consumer enablement, Runtime consumer enablement, grant issuance, delivery authorization grants, notification sending, Notification Center events, notification actions, action cards, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Kept predecessor consent collection receipt consumption gate and consumer authorization audits mainline-ready after the safe next task advances by recognizing current mainline continuity anchors instead of requiring those predecessors to remain the active safe next task.

## [0.2.422] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumer-authorization-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline, v0.2.421 notification delivery consent collection receipt consumption gate audit, and v0.2.420 notification delivery consent collection receipt acceptance audit before modeling future KDE Notification Center consent collection receipt consumer authorization readiness.
- Wired the notification delivery consent collection receipt consumer authorization audit into the Runtime CLI while keeping consent collection, consent persistence, consent receipt creation, consent receipt persistence, consent receipt acceptance, consent receipt consumption, consumer authorization grants, KDE consumer authorization, Runtime consumer authorization, consumer enablement, grant issuance, delivery authorization grants, notification sending, Notification Center events, notification actions, action cards, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.421] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-consumption-gate-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.420 notification delivery consent collection receipt acceptance audit before modeling future KDE Notification Center consent collection receipt consumption gate readiness.
- Wired the notification delivery consent collection receipt consumption gate audit into the Runtime CLI while keeping consent collection, consent persistence, consent receipt creation, consent receipt persistence, consent receipt acceptance, consent receipt consumption, consumption authorization, grant issuance, delivery authorization grants, notification sending, Notification Center events, notification actions, action cards, receipt presence, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.420] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-acceptance-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.419 notification delivery consent collection receipt audit before modeling future KDE Notification Center consent collection receipt acceptance readiness.
- Wired the notification delivery consent collection receipt acceptance audit into the Runtime CLI while keeping consent collection, consent persistence, consent receipt creation, consent receipt persistence, consent receipt acceptance, consent receipt consumption, grant issuance, delivery authorization grants, notification sending, Notification Center events, notification actions, action cards, receipt presence, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Captured operator-authorized restricted full Buildroot/QEMU serial smoke evidence for v0.2.420 in `output/full-smoke-report.json`, `output/full-smoke-report.md`, and `output/serial.log` while keeping privileged containers, host networking, Docker socket mounts inside containers, broad host mounts, backend launch, and host-root mutation disabled.

## [0.2.419] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-receipt-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.418 notification delivery consent collection gate audit before modeling future KDE Notification Center consent collection receipt readiness.
- Wired the notification delivery consent collection receipt audit into the Runtime CLI while keeping consent collection, consent persistence, consent receipt creation, consent receipt persistence, consent receipt acceptance, consent receipt consumption, grant issuance, delivery authorization grants, notification sending, Notification Center events, notification actions, action cards, receipt presence, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.419` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, Runtime activation smoke, and the restricted product smoke packet while formal full Buildroot/QEMU smoke remains gated by `scripts/full_smoke.rb`.

## [0.2.418] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-consent-collection-gate-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.417 notification delivery explicit consent audit before modeling future KDE Notification Center consent collection gates.
- Wired the notification delivery consent collection gate audit into the Runtime CLI while keeping consent collection, consent persistence, consent receipt creation, consent receipt persistence, grant issuance, delivery authorization grants, notification sending, Notification Center events, notification actions, action cards, receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.417] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-explicit-consent-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.416 notification delivery grant audit before modeling future KDE Notification Center explicit consent readiness.
- Wired the notification delivery explicit consent audit into the Runtime CLI while keeping consent collection, grant issuance, delivery authorization grants, notification sending, Notification Center events, notification actions, action cards, receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

### Fixed

- Shortened fail-closed owner-audit test temp directory prefixes so the restricted Docker heavy smoke `go test ./...` path works on Linux filesystems with component-length limits.

## [0.2.416] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-grant-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.415 notification delivery authorization audit before modeling future KDE Notification Center delivery grant readiness.
- Wired the notification delivery grant audit into the Runtime CLI while keeping grant issuance, delivery authorization grants, notification sending, Notification Center events, notification actions, action cards, explicit consent collection, receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.415] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-delivery-authorization-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.414 notification presentation audit before modeling future KDE Notification Center delivery authorization.
- Wired the notification delivery authorization audit into the Runtime CLI while keeping delivery grants, notification sending, Notification Center events, notification actions, action cards, explicit consent collection, receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.415` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, and Runtime activation smoke while keeping formal full Buildroot/QEMU smoke gated by `scripts/full_smoke.rb`.

## [0.2.414] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-notification-presentation-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.413 presentation consent audit before modeling future KDE Notification Center presentation.
- Wired the notification presentation audit into the Runtime CLI while keeping notification delivery, notification actions, action cards, explicit consent collection, receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup routes, lookup route dispatch authorization, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.414` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, and Runtime activation smoke while keeping formal full Buildroot/QEMU smoke gated by `scripts/full_smoke.rb`.

## [0.2.413] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-presentation-consent-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.412 redacted presentation eligibility audit before modeling future KDE-safe notification and action-card consent boundaries.
- Wired the presentation consent audit into the Runtime CLI while keeping explicit consent collection, notification triggers, action cards, receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup routes, lookup route dispatch authorization, dry-run result redaction passage, dispatch calls, status persistence writes, storage writes, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.412] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redacted-presentation-eligibility-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.411 route enablement lookup route dispatch dry-run result redaction boundary audit before modeling future KDE-safe redacted presentation eligibility.
- Wired the redacted presentation eligibility audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup routes, lookup route dispatch authorization, dry-run result redaction passage, dispatch calls, status persistence writes, storage writes, notification actions, support side effects, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Captured operator-authorized restricted container heavy smoke evidence for the local `xnix-builder:0.2.412` image, Runtime owner candidate restricted session smoke, Runtime D-Bus session smoke, KDE Compatibility Center D-Bus smoke, and Runtime activation smoke while keeping full Buildroot/QEMU smoke gated by the formal milestone script.

## [0.2.411] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-redaction-boundary-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.410 route enablement lookup route dispatch dry-run result capture audit before modeling future KDE-safe dry-run result redaction boundaries.
- Wired the route enablement lookup route dispatch dry-run result redaction boundary audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup route enablement grants, lookup route grant authorization, lookup routes, lookup route dispatch authorization, dry-run result capture passage, dry-run result redaction passage, dispatch calls, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.410] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-result-capture-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.409 route enablement lookup route dispatch dry-run execution gate audit before modeling future lookup route dispatch dry-run result capture.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run result capture audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup route enablement grants, lookup route grant authorization, lookup routes, lookup route dispatch authorization, dispatch dry-run result capture passage, dispatch calls, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.
- Captured operator-authorized restricted full build, QEMU serial boot, and loopback SSH smoke evidence for v0.2.410 without privileged containers, host networking, Docker socket mounts inside containers, broad host mounts, backend launch, or host-root mutation.

## [0.2.409] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-dry-run-execution-gate-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.408 route enablement lookup route dispatch authorization audit before modeling future lookup route dispatch dry-run execution gates.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch dry-run execution gate audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup route enablement grants, lookup route grant authorization, lookup routes, lookup route dispatch authorization, dispatch dry-run execution gate passage, dispatch calls, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.408] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-authorization-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.407 route enablement lookup route dispatch gate audit before modeling future lookup route dispatch authorization.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch authorization audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup route enablement grants, lookup route grant authorization, lookup routes, lookup route dispatch authorization, dispatch authorization passage, dispatch calls, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.407] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-dispatch-gate-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.406 route enablement lookup route enablement audit before modeling future lookup route dispatch gates.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route dispatch gate audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup route enablement grants, lookup route grant authorization, lookup routes, lookup route dispatch authorization, dispatch gate passage, dispatch calls, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.406] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-enablement-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.405 route enablement lookup route grant audit before modeling future lookup route enablement.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route enablement audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup route enablement grants, lookup route grant authorization, lookup routes, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.405] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-lookup-route-grant-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.404 route enablement accepted receipt gate audit before modeling future lookup route grants.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection route enablement lookup route grant audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup route grants, lookup route grant authorization, lookup routes, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.404] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-accepted-receipt-gate-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.403 route enablement receipt acceptance audit before modeling future accepted receipt gates.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection route enablement accepted receipt gate audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup route grants, lookup routes, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

### Fixed

- Shortened missing-source Go test temporary-directory prefixes so restricted Docker build smoke can run the long route enablement receipt acceptance and accepted receipt gate audit tests on overlay filesystems.

## [0.2.403] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-acceptance-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.402 route enablement receipt gate before modeling future route enablement receipt acceptance.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection route enablement receipt acceptance audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, route enablement acceptance, lookup route grants, lookup routes, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.402] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-receipt-gate-preview`, a read-only Go owner gate that consumes the Xnix current mainline and v0.2.401 route enablement evidence gate before modeling future operator-reviewed route enablement receipt gates.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection route enablement receipt gate into the Runtime CLI while keeping receipt creation, receipt acceptance, receipt consumption, receipt persistence, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, lookup route enablement, lookup routes, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.401] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-enablement-evidence-gate-preview`, a read-only Go owner gate that consumes the Xnix current mainline and v0.2.400 route authorization evidence gate before gating future lookup route enablement evidence.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection route enablement evidence gate into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, lookup route enablement, lookup routes, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.400] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-route-authorization-evidence-gate-preview`, a read-only Go owner gate that consumes the Xnix current mainline and v0.2.399 storage-root authorization receipt dry-run lookup result consumer projection evidence audit preview before gating KDE-safe Compatibility Center and Runtime diagnostics route authorization evidence.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection route authorization evidence gate into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, lookup route authorization, lookup routes, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

### Verification

- Operator-authorized restricted full build and QEMU serial smoke passed for the v0.2.400 twentieth small-version boundary, with `output/full-smoke-report.json`, `output/full-smoke-report.md`, and `output/serial.log` showing expected boot markers and constrained execution.

## [0.2.399] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-evidence-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.398 storage-root authorization receipt dry-run lookup result consumer projection preview before auditing KDE-safe Compatibility Center and Runtime diagnostics projection evidence.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection evidence audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, lookup routes, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.398] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-consumer-projection-preview`, a read-only Go owner preview that consumes the Xnix current mainline and v0.2.397 storage-root authorization receipt dry-run lookup result boundary preview before modeling KDE-safe Compatibility Center and Runtime diagnostics result consumer projections.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result consumer projection preview into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, status persistence writes, storage writes, consumer enablement, lookup routes, dry-run dispatch, result persistence, raw result exposure, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.397] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-result-boundary-preview`, a read-only Go owner preview that consumes the Xnix current mainline and v0.2.396 storage-root authorization receipt dry-run lookup evidence preview before modeling redacted lookup result identity, readiness, and failure boundaries.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup result boundary preview into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, durable record writes, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.396] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-dry-run-lookup-evidence-preview`, a read-only Go owner preview that consumes the Xnix current mainline and v0.2.395 storage-root authorization receipt consumption audit before modeling future dry-run lookup evidence for storage-root authorization receipts.
- Wired the KDE-safe redacted status storage-root authorization receipt dry-run lookup evidence preview into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, durable record writes, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.395] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-authorization-receipt-consumption-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.394 storage-root policy audit before modeling accepted receipt, receipt scope, storage-root policy grant, and record-writer call grant boundaries.
- Wired the KDE-safe redacted status storage-root authorization receipt consumption audit into the Runtime CLI while keeping receipt presence, receipt acceptance, receipt consumption, storage-root policy grants, record-writer call authorization, storage-root resolution, writer calls, durable record writes, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.394] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-storage-root-policy-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.393 closed record writer implementation preview before modeling storage-root ownership, namespace, retention, redaction, and record-writer call guard boundaries.
- Wired the KDE-safe redacted status closed record writer storage-root policy audit into the Runtime CLI while keeping storage-root resolution, creation, mounting, ownership grants, namespace grants, retention enforcement, redaction enforcement, writer calls, closed record writers, durable record writes, storage gate passage, storage authorization, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.393] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-record-writer-implementation-preview`, a read-only Go owner preview that consumes the Xnix current mainline and v0.2.392 persistence record contract preview before modeling the future closed redacted status record writer implementation.
- Wired the KDE-safe redacted status closed record writer implementation preview into the Runtime CLI while keeping writer calls, closed record writers, durable record writes, storage gate passage, storage authorization, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.392] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-record-contract-preview`, a read-only Go owner preview that consumes the Xnix current mainline and v0.2.391 storage persistence gate audit before modeling the durable redacted status record contract.
- Wired the KDE-safe redacted status persistence record contract preview into the Runtime CLI while keeping durable record writes, storage gate passage, writer calls, persistence writers, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.391] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-storage-persistence-gate-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.390 closed persistence writer preview before modeling the gate that must pass before redacted status storage can write.
- Wired the KDE-safe redacted status storage persistence gate audit into the Runtime CLI while keeping gate passage, writer calls, persistence writers, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.390] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-persistence-writer-preview`, a read-only Go owner preview that consumes the Xnix current mainline and v0.2.389 writer persistence authorization preview before modeling the disabled redacted status storage writer shape.
- Wired the KDE-safe redacted status closed persistence writer preview into the Runtime CLI while keeping writer calls, persistence writers, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.389] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-persistence-authorization-preview`, a read-only Go owner preview that consumes the Xnix current mainline and v0.2.388 closed writer implementation preview before modeling the persistence authorization boundary.
- Wired the KDE-safe redacted status writer persistence authorization preview into the Runtime CLI while keeping callable writers, real writer authorization grants, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.388] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-closed-writer-implementation-preview`, a read-only Go owner preview that consumes the Xnix current mainline and v0.2.387 writer enablement audit before modeling the closed writer call shape.
- Wired the KDE-safe redacted status closed writer implementation preview into the Runtime CLI while keeping callable writers, real writer authorization grants, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.387] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-enablement-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.386 writer grant audit before modeling the future writer enablement boundary.
- Wired the KDE-safe redacted status writer enablement audit into the Runtime CLI while keeping real writer authorization grants, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.386] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-grant-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.385 accepted receipt gate audit before modeling the future writer grant boundary.
- Wired the KDE-safe redacted status writer grant audit into the Runtime CLI while keeping writer authorization grants, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.385] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-accepted-receipt-gate-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.384 writer authorization receipt acceptance audit before modeling the future accepted receipt gate.
- Wired the KDE-safe redacted status writer accepted receipt gate audit into the Runtime CLI while keeping receipt acceptance, writer authorization grants, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.384] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-acceptance-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.383 writer authorization receipt audit before modeling the future acceptance authorization boundary.
- Wired the KDE-safe redacted status writer authorization receipt acceptance audit into the Runtime CLI while keeping receipt acceptance, writer authorization grants, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.383] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-receipt-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.382 writer authorization gate before modeling opaque writer authorization receipt records.
- Wired the KDE-safe redacted status writer authorization receipt audit into the Runtime CLI while keeping receipt acceptance, writer authorization grants, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.382] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-writer-authorization-gate-audit-preview`, a read-only Go owner audit that consumes the Xnix current mainline and v0.2.380 redacted status write model before modeling the future writer authorization gate.
- Wired the KDE-safe redacted status writer authorization gate audit into the Runtime CLI while keeping writer authorization grants, status writers, status persistence writes, consumer enablement, lookup routes, dry-run execution, request creation, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.381] - 2026-07-19

### Added

- Added `docs/xnix-current-mainline.md` as the Xnix-owned current mainline source for new implementation work, replacing external-agent dispatch as the default planning source while preserving historical handoff documents as repository evidence.
- Added `current-mainline-ownership-audit-preview`, a read-only Go owner audit and Runtime CLI command that verifies the autonomous mainline source, KDE-only first-release scope, Go-first Runtime ownership, Ruby test harness boundary, C low-level boundary, seven KDE entrypoints, explicit full-gate authorization, and closed runtime/host safety gates.
- Classified the Xnix current mainline document and current-mainline Runtime CLI audit files in the mainline integration review so merge readiness can converge without staging unclassified side work.

## [0.2.380] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-redacted-status-persistence-write-model-audit-preview`, a read-only Go owner audit that defines the future redacted status record shape for Compatibility Center and Runtime diagnostics.
- Wired the KDE-safe redacted status persistence write-model audit into the Runtime CLI while keeping status persistence authorization grants, real status writes, consumer enablement, lookup route enablement, lookup enablement, dry-run execution, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.379] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-persistence-authorization-audit-preview`, a read-only Go owner audit that models the authorization boundary required before KDE-safe Compatibility Center or Runtime diagnostics status summaries can be stored durably.
- Wired the KDE-safe status persistence authorization audit into the Runtime CLI while keeping persistence authorization grants, status persistence writes, consumer enablement, lookup route enablement, lookup enablement, dry-run execution, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.378] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-kde-safe-status-fanout-audit-preview`, a read-only Go owner audit that projects the closed consumer enablement gate into KDE-safe Compatibility Center and Runtime diagnostics status summaries without enabling consumers.
- Wired the KDE-safe status fan-out audit into the Runtime CLI while keeping consumer enablement, lookup route grants, route enablement, lookup enablement, status persistence, Runtime diagnostics persistence, raw result exposure, dry-run execution, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.377] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-enablement-gate-audit-preview`, a read-only Go owner audit that models the final consumer enablement gate after the receipt consumer authorization audit before KDE or Runtime consumers can be enabled.
- Wired the dry-run result lookup consumer enablement receipt consumer enablement gate audit into the Runtime CLI while keeping consumer enablement, consumer authorization grants, lookup route grants, route enablement, lookup enablement, lookup persistence, redacted summary persistence, raw result exposure, dry-run execution, result persistence, Runtime diagnostics persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.376] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumer-authorization-audit-preview`, a read-only Go owner audit that models the explicit consumer authorization boundary after the receipt consumption gate before KDE or Runtime consumers can be enabled.
- Wired the dry-run result lookup consumer enablement receipt consumer authorization audit into the Runtime CLI while keeping consumer authorization grants, consumer enablement, lookup route authorization grants, route enablement, lookup enablement, lookup persistence, redacted summary persistence, raw result exposure, dry-run execution, result persistence, Runtime diagnostics persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.375] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-consumption-gate-audit-preview`, a read-only Go owner audit that models the final fail-closed consumption gate before future accepted consumer enablement authorization receipts can authorize KDE or Runtime consumers.
- Wired the dry-run result lookup consumer enablement receipt consumption gate audit into the Runtime CLI while keeping receipt acceptance, receipt consumption, consumer authorization, consumer enablement, lookup route authorization grants, route enablement, lookup enablement, lookup persistence, redacted summary persistence, raw result exposure, dry-run execution, result persistence, Runtime diagnostics persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.374] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-acceptance-authorization-audit-preview`, a read-only Go owner audit that models the acceptance authorization boundary required before future persisted consumer enablement authorization receipts can be accepted and consumed by consumer enablement gates.
- Wired the dry-run result lookup consumer enablement receipt acceptance authorization audit into the Runtime CLI while keeping receipt acceptance, receipt consumption, consumer authorization, consumer enablement, lookup route authorization grants, route enablement, lookup enablement, lookup persistence, redacted summary persistence, raw result exposure, dry-run execution, result persistence, Runtime diagnostics persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.373] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-persistence-authorization-audit-preview`, a read-only Go owner audit that models the persistence, replay, expiry, and revocation authorization boundary required before future consumer enablement authorization receipts can be persisted or accepted.
- Wired the dry-run result lookup consumer enablement receipt persistence authorization audit into the Runtime CLI while keeping receipt writes, receipt acceptance, persistence, replay, expiry writes, revocation writes, consumer authorization, consumer enablement, lookup route authorization grants, route enablement, lookup enablement, lookup persistence, redacted summary persistence, raw result exposure, dry-run execution, result persistence, Runtime diagnostics persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.372] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-receipt-writer-authorization-audit-preview`, a read-only Go owner audit that models the writer authorization boundary required before future consumer enablement authorization receipts can be written.
- Wired the dry-run result lookup consumer enablement receipt writer authorization audit into the Runtime CLI while keeping receipt writes, receipt acceptance, persistence, replay, consumer authorization, consumer enablement, lookup route authorization grants, route enablement, lookup enablement, lookup persistence, redacted summary persistence, raw result exposure, dry-run execution, result persistence, Runtime diagnostics persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.371] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-authorization-receipt-audit-preview`, a read-only Go owner audit that models the opaque authorization receipt boundary required before consumer enablement gates can authorize KDE or Runtime consumers.
- Wired the dry-run result lookup consumer enablement authorization receipt audit into the Runtime CLI while keeping receipt presence, receipt acceptance, consumer authorization, consumer enablement, lookup route authorization grants, route enablement, lookup enablement, lookup persistence, redacted summary persistence, raw result exposure, dry-run execution, result persistence, Runtime diagnostics persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.370] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-enablement-gate-audit-preview`, a read-only Go owner audit that combines lookup route authorization and consumer redaction evidence before any KDE or Runtime dry-run result consumer can be enabled.
- Wired the dry-run result lookup consumer enablement gate audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping consumer authorization, consumer enablement, route authorization grants, route enablement, lookup enablement, lookup persistence, redacted summary persistence, raw result exposure, dry-run execution, result persistence, Runtime diagnostics persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.369] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-consumer-redaction-audit-preview`, a read-only Go owner audit that models KDE and Runtime consumer redaction boundaries before any surface can consume opaque dry-run result identifiers.
- Wired the dry-run result lookup consumer redaction audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping consumer authorization, consumer enablement, route authorization grants, route enablement, lookup enablement, lookup persistence, redacted summary persistence, raw result exposure, dry-run execution, result persistence, Runtime diagnostics persistence, request creation, dispatch, Portal requests, notifications, Compatibility Center navigation, support writes, production ownership, Runtime writes, desktop side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.368] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-lookup-route-authorization-audit-preview`, a read-only Go owner audit that models the owner-local route authorization boundary required before KDE or Runtime surfaces can consume opaque dry-run result identifiers.
- Wired the dry-run result lookup route authorization audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping route authorization grants, route enablement, route persistence, lookup enablement, lookup persistence, result persistence, dry-run execution, request creation, dispatch, Portal requests, receipt writes, production ownership, Runtime writes, desktop side effects, support side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.367] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-opaque-lookup-audit-preview`, a read-only Go owner audit that models owner-managed opaque identifiers for future stored dry-run result summaries.
- Wired the dry-run result opaque lookup audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping lookup enablement, lookup persistence, result persistence, retention enforcement, redaction enforcement, dry-run execution, request creation, dispatch, Portal requests, receipt writes, production ownership, Runtime writes, desktop side effects, support side effects, path exposure, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.366] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-retention-redaction-policy-audit-preview`, a read-only Go owner audit that models retention windows and redaction rules for future stored dry-run result summaries.
- Wired the dry-run result retention redaction policy audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping policy persistence, retention enforcement, redaction enforcement, dry-run execution, dry-run result persistence, visibility persistence, request creation, dispatch, Portal requests, receipt writes, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.365] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-persistence-authorization-audit-preview`, a read-only Go owner audit that models the authorization boundary required before redacted dry-run result summaries can be stored.
- Wired the dry-run result persistence authorization audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping dry-run execution, dry-run result persistence, visibility persistence, dispatch grants, request-object creation, request dispatch, request persistence, notification actions, Portal requests, Compatibility Center navigation, receipt writes, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.364] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dry-run-result-visibility-audit-preview`, a read-only Go owner audit that models how future dispatch dry-run results would be redacted for KDE and summarized for Runtime diagnostics.
- Wired the dry-run result visibility audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping dry-run execution, dry-run result persistence, visibility persistence, dispatch grants, request-object creation, request dispatch, request persistence, notification actions, Portal requests, Compatibility Center navigation, receipt writes, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.363] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dispatch-dry-run-audit-preview`, a read-only Go owner audit that models the dry-run execution boundary required after notification action dispatch authorization.
- Wired the dispatch dry-run audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping dry-run execution, dry-run result persistence, dispatch grants, request-object creation, request dispatch, request persistence, notification actions, Portal requests, Compatibility Center navigation, receipt writes, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.362] - 2026-07-19

### Added

- Added `production-receipt-notification-action-dispatch-authorization-audit-preview`, a read-only Go owner audit that models the authorization boundary required before notification action request objects can be dispatched.
- Wired the dispatch authorization audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping dispatch grants, request-object creation, request dispatch, request persistence, notification actions, Portal requests, Compatibility Center navigation, receipt writes, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.361] - 2026-07-19

### Added

- Added `production-receipt-notification-action-request-object-audit-preview`, a read-only Go owner audit that maps review, renew, open Compatibility Center, dismiss, and support-info notification actions to future Runtime request-object kinds.
- Wired the request-object audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping request-object creation, request dispatch, request persistence, notification actions, Portal requests, Compatibility Center navigation, receipt writes, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.360] - 2026-07-19

### Added

- Added `production-receipt-notification-action-safety-audit-preview`, a read-only Go owner audit that models review, renew, open Compatibility Center, dismiss, and support-info actions for future receipt-related notifications.
- Wired the notification action safety audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping notification actions, notification delivery, Portal requests, Runtime request objects, Compatibility Center navigation, receipt writes, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.359] - 2026-07-19

### Added

- Added `production-receipt-notification-delivery-gate-audit-preview`, a read-only Go owner audit that models the separate gate required before expiring, expired, revoked, and missing-review receipt states can notify users.
- Wired the notification delivery gate audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping notification delivery, notification actions, Portal requests, Runtime request objects, receipt writes, receipt persistence, receipt lookup, replay, expiry writes, revocation writes, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.358] - 2026-07-19

### Added

- Added `production-receipt-revocation-visibility-audit-preview`, a read-only Go owner audit that models how current, expiring, expired, revoked, and missing-review authorization receipt states would surface to production gates and KDE-safe status views.
- Wired the revocation visibility audit into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping receipt revocation writes, expiry writes, persistence, lookup writes, replay, notification delivery, Compatibility Center persistence, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.357] - 2026-07-19

### Added

- Added `production-receipt-persistence-threat-review-preview`, a read-only Go owner review that models storage confidentiality, expiry, revocation, replay protection, and audit visibility requirements for future authorization receipts.
- Wired the persistence threat review into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping receipt persistence, lookup writes, replay, expiry writes, revocation writes, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.356] - 2026-07-19

### Added

- Added `production-receipt-writer-authorization-review-preview`, a read-only Go owner review that models the separate operator and writer authorization boundary required before any future authorization receipt can be written, persisted, accepted, or replayed.
- Wired the writer authorization review into the Runtime CLI, layout verifier, implementation evidence report, and product documentation while keeping receipt writes, receipt persistence, receipt replay, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.355] - 2026-07-19

### Added

- Added `production-receipt-acceptance-propagation-preflight-preview`, a read-only Go owner preflight that models how a future accepted opaque authorization receipt would propagate across the production D-Bus gate, method review, service activation preflight, Runtime write gate, rollback diagnostics review, and desktop side-effect review while keeping real receipt acceptance, receipt writes, production ownership, Runtime writes, desktop side effects, support side effects, engine launch, unsafe data exposure, and host mutation disabled.
- Wired the propagation preflight into the Runtime CLI, layout verifier, implementation evidence report, and product documentation.

## [0.2.354] - 2026-07-19

### Added

- Added `production-authorization-consumption-audit-preview`, a read-only Go owner audit that verifies production D-Bus gate review, method review, service activation preflight, Runtime write gate, rollback diagnostics review, and desktop side-effect review all consume the consolidated opaque authorization receipt boundary.
- Added CLI, Go tests, layout, evidence, and product metadata coverage for production authorization consumption auditing while keeping authorization acceptance, receipt writes, production ownership, Runtime writes, desktop side effects, support side effects, restore, cleanup, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.353] - 2026-07-19

### Added

- Added `production-human-authorization-receipt-consolidation-preview`, a read-only Go owner model that consolidates the production human authorization receipt boundary across production D-Bus gate, method review, service activation preflight, write gate, rollback diagnostics, and desktop side-effect reviews.
- Added CLI, Go tests, layout, evidence, and product metadata coverage for the owner-managed opaque authorization receipt boundary while keeping receipt writes, receipt acceptance, service start, bus ownership, Runtime writes, desktop side effects, support side effects, restore, cleanup, engine launch, unsafe data exposure, and host mutation disabled.

## [0.2.352] - 2026-07-19

### Added

- Added `production-desktop-side-effect-review-preview`, a read-only Go owner review that inventories the seven KDE first-release entry points before any production desktop side-effect decision.
- Added CLI, Go tests, layout, evidence, and product metadata coverage for KDE desktop side-effect review while keeping desktop files, MIME defaults, shell configuration, settings, KRunner indexes, task-manager activation, KWin rules, tray bridges, notifications, Compatibility Center persistence, Portal requests, Runtime request objects, engine launch, unsafe data exposure, production bus ownership, and host mutation disabled.

## [0.2.351] - 2026-07-19

### Added

- Added `production-rollback-diagnostics-review-preview`, a read-only Go owner review that consumes production D-Bus gate, method review, service activation, support, snapshot restore, and state-retention safety surfaces before any production ownership decision.
- Added CLI, Go tests, layout, evidence, and product metadata coverage for rollback and diagnostics review while keeping service start, bus ownership, writes, support export, support case creation, notifications, restore, cleanup, file-content reads, unsafe data exposure, network use, privileged containers, engine launch, and host mutation disabled.

## [0.2.350] - 2026-07-19

### Added

- Added `production-dbus-method-review-preview`, a read-only Go owner review that inventories 61 read-only D-Bus contract methods, 3 owner-local candidate routes, and 4 reserved write methods before any production D-Bus exposure decision.
- Added CLI, Go tests, layout, evidence, and product metadata coverage for route-by-route production D-Bus method review while registering no new production methods, claiming no bus ownership, enabling no writes, launching no backend, sending no notifications, and mutating no host state.

## [0.2.349] - 2026-07-19

### Added

- Added production gate consumption to `runtime-write-gate-preview`, letting the write gate consume `runtime-service-activation-preflight-preview`, the production D-Bus gate review, and the human authorization preflight while write dispatch remains disabled.
- Added write-gate checks for production D-Bus gate review, production service activation preflight, and the still-pending human authorization receipt while preserving disabled request creation, execution, backend launch, network use, privileged containers, and host mutation.

## [0.2.348] - 2026-07-19

### Added

- Added production D-Bus gate consumption to `runtime-service-activation-preflight-preview`, including a source-backed gate summary for the production D-Bus gate review and human authorization preflight.
- Added service activation checks for production D-Bus gate review, human authorization preflight, and the still-pending human authorization receipt while keeping production activation, service start, bus ownership, Runtime writes, backend launch, network use, privileged containers, and host mutation disabled.

## [0.2.347] - 2026-07-19

### Added

- Added `production-dbus-human-authorization-preflight-preview`, a read-only Go owner preflight that defines the future human authorization receipt shape required by the production D-Bus gate without granting authorization.
- Wired the production D-Bus gate review to consume the preflight shape while keeping authorization ungranted, production readiness false, service start disabled, bus ownership disabled, Runtime writes disabled, backend launch disabled, support side effects disabled, and host mutation disabled.
