# Changelog

Xnix follows Semantic Versioning.

## [0.2.123] - 2026-07-15

### Added

- Added `xnix-runtime-go permission-review-preview --registry <path> --app <id>` for read-only KDE permission review previews from digest-verified registry recipes.
- Added Go Runtime permission review planning for documents, downloads, camera, network, clipboard, print, and screenshot decisions with Portal metadata and Runtime gate summaries.
- Added Go and Ruby harness coverage that verifies permission review previews remain Runtime-owned, KDE-targeted, Portal-review-required, backend-detail-free, host-root-safe, and honest about request objects, permission grants, direct access, settings persistence, and host permission changes remaining disabled.

### Changed

- Moved the user-visible permission review surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for review presentation only.

## [0.2.122] - 2026-07-15

### Added

- Added `xnix-runtime-go file-open-preview --registry <path> [--app <id>] file://...` for read-only Dolphin file-open previews from digest-verified registry recipes.
- Added Go Runtime file-open planning that resolves applications by selected-file extension or explicit application id, records Portal `OpenFile` mediation, and returns the managed `xnix-compat-open --app <id> %U` action.
- Added Go and Ruby harness coverage that verifies Dolphin file-open previews remain Runtime-owned, KDE-targeted, Portal-required, backend-detail-free, host-root-safe, and honest about request objects, permission grants, direct host file access, and backend launch remaining disabled.

### Changed

- Moved the Dolphin file-open entry point one step closer to Go-owned Runtime product logic while keeping Dolphin responsible for selection and presentation only.

## [0.2.121] - 2026-07-15

### Added

- Added `xnix-runtime-go compatibility-center-preview --registry <path>` for read-only KDE Compatibility Center previews across digest-verified registry recipes.
- Added Go Runtime Compatibility Center application cards with generated desktop identity, user-facing runtime mode, diagnostics state, known issue counts, repair record state, safe navigation actions, and disabled execution gates.
- Added Go and Ruby harness coverage that verifies Compatibility Center previews remain Runtime-owned, KDE-targeted, backend-detail-free, host-root-safe, and honest about diagnostics, actions, repairs, backend launch, and settings persistence remaining gated.

### Changed

- Moved the Compatibility Center overview entry point one step closer to Go-owned Runtime product logic while keeping KDE responsible for display and review only.

## [0.2.120] - 2026-07-15

### Added

- Added `xnix-runtime-go krunner-query-preview --registry <path> --query <text>` for read-only KDE KRunner query previews across digest-verified registry recipes.
- Added Go Runtime KRunner matching for application names, application identifiers, file extensions, and natural launch phrases while returning normal Linux application launcher actions.
- Added registry-wide Go recipe loading plus Go and Ruby harness coverage that verifies KRunner results remain Runtime-owned, KDE-targeted, backend-detail-free, host-root-safe, and honest about query execution and backend launch remaining disabled.

### Changed

- Moved the KRunner search entry point one step closer to Go-owned Runtime product logic while keeping KDE responsible for presenting and invoking approved launcher actions only.

## [0.2.119] - 2026-07-15

### Added

- Added `xnix-runtime-go settings-preview --registry <path> --app <id>` for read-only KDE unified-settings previews from the registry-backed desktop identity plan.
- Added Go Runtime settings modeling for run mode, file access, device access, network access, and snapshots with desktop identity fields and user-facing defaults.
- Added Go and Ruby harness coverage that verifies Runtime-owned settings previews remain KDE-targeted, user-visible, backend-detail-free, host-root-safe, and honest about settings persistence remaining disabled.

### Changed

- Moved the unified-settings entry point one step closer to Go-owned Runtime product logic while keeping KDE responsible for display and review only.

## [0.2.118] - 2026-07-15

### Added

- Added `xnix-runtime-go notification-preview --registry <path> --app <id> --event <event>` for read-only KDE notification previews from the registry-backed desktop identity plan.
- Added Go Runtime notification modeling for install failures, repair receipts, mode changes, and approval requests with urgency, category, actions, review state, and desktop identity fields.
- Added Go and Ruby harness coverage that verifies Runtime-owned notification previews remain KDE-targeted, user-visible, backend-detail-free, host-root-safe, and honest about action, repair, and settings execution gates remaining closed.

### Changed

- Moved the notification-center entry point one step closer to Go-owned Runtime product logic while keeping KDE responsible for display and user interaction only.

## [0.2.117] - 2026-07-15

### Added

- Added `xnix-runtime-go tray-status-preview --registry <path> --app <id>` for read-only KDE tray status previews from the registry-backed desktop identity plan.
- Added Go Runtime tray status modeling for registered compatibility applications, ready state, planned tray bridge state, KDE navigation actions, and disabled live bridge/persistence gates.
- Added Go and Ruby harness coverage that verifies Runtime-owned tray previews remain KDE-targeted, user-visible, backend-detail-free, host-root-safe, and honest about live tray bridging being planned rather than active.

### Changed

- Moved the system tray entry point one step closer to Go-owned Runtime product logic while keeping live backend tray bridging gated until a production bridge exists.

## [0.2.116] - 2026-07-15

### Added

- Added `xnix-runtime-go window-identity-preview --registry <path> --app <id>` for read-only KDE window identity previews from the registry-backed desktop identity plan.
- Added Go Runtime task-manager and KWin identity hints covering generated desktop files, launcher URLs, grouping keys, pinning, restore, taskbar visibility, switcher visibility, and bounded KWin window-management policy.
- Added Go and Ruby harness coverage that verifies Runtime-owned window identity previews remain KDE-targeted, backend-detail-free, host-root-safe, and usable by future task manager and KWin integrations.

### Changed

- Moved another durable desktop integration surface toward Go-owned Runtime product logic while keeping KDE responsible for presentation and interaction only.

## [0.2.115] - 2026-07-15

### Added

- Added `xnix-runtime-go mimeapps-preview --registry <path> --app <id>` for read-only `mimeapps.list` preview output from the registry-backed desktop identity plan.
- Added activation-staging file association renderer selection so `xnix-install-desktop-integration --file-association-source runtime-go` can stage MIME defaults rendered by the Go Runtime.
- Added Go, Ruby, and CLI coverage for Runtime Go MIME association rendering while preserving the default Ruby renderer for constrained development staging.

### Changed

- Extended activation safety reporting with the file association source so staged KDE activation results can distinguish Ruby MIME output from Runtime Go-rendered MIME output.

## [0.2.114] - 2026-07-15

### Added

- Added an activation-staging desktop entry renderer switch so `xnix-install-desktop-integration` can stage desktop entries rendered by `xnix-runtime-go desktop-entry-preview`.
- Added Runtime Go renderer validation in the activation installer, including managed launcher, application id, and backend-detail safety checks before staging the desktop entry.
- Added CLI and unit coverage for the `--desktop-entry-source runtime-go` path while keeping the default Ruby renderer available for constrained local development.

### Changed

- Extended activation safety reporting with the desktop entry source so KDE activation receipts can distinguish Ruby staging output from Runtime Go-rendered desktop entries.

## [0.2.113] - 2026-07-15

### Added

- Added Go Runtime desktop entry rendering so registry-backed desktop identity plans can produce standard KDE `.desktop` text without writing activation files.
- Added `xnix-runtime-go desktop-entry-preview --registry <path> --app <id>` for read-only launcher preview output using the managed Runtime launcher command.
- Added Go and Ruby harness coverage that assert the generated desktop entry remains user-visible, Runtime-owned, backend-detail-free, and safe for KDE launcher and task-manager use.

### Changed

- Moved standard desktop entry generation one step closer to Go-owned Runtime product logic while keeping Ruby desktop activation code as a staging harness.

## [0.2.112] - 2026-07-15

### Added

- Added Go Runtime registry-backed recipe loading for desktop identity planning with schema checks, application-id lookup, safe relative-path enforcement, SHA-256 verification, recipe-id matching, and signature-status reporting.
- Added `xnix-runtime-go desktop-identity-plan --registry <path> --app <id>` so the Go Runtime can resolve a managed application from the recipe library before producing a KDE-safe normal Linux application identity.
- Added Go unit coverage for registry digest verification, digest mismatch rejection, unsafe path rejection, and CLI registry lookup.

### Changed

- Updated the Go desktop identity harness to prefer registry-backed recipe lookup over direct recipe-file input while retaining direct input for constrained development tests.

## [0.2.111] - 2026-07-15

### Added

- Added a Go Runtime desktop identity planning package and CLI that turn an application recipe into a KDE-safe normal Linux application identity.
- Added Go unit tests and Docker image validation so Runtime Go logic is compiled and tested inside the constrained Colima development environment.
- Added Ruby harness coverage that verifies the Go source boundary, Docker build checks, managed launcher command, MIME identity, and backend terminology hiding when Go is available.

### Changed

- Updated Runtime implementation guidance so important product logic moves to Go first, with C reserved for low-level or already-owned Runtime policy surfaces and Ruby kept for tests and lightweight tooling.

## [0.2.110] - 2026-07-15

### Added

- Added Runtime-owned KDE shell integration plans that define KDE Plasma shell component boundaries for the start menu, task manager, Dolphin, tray, notifications, Compatibility Center, settings, KRunner, and KWin.
- Added `GetKDEShellIntegrationPlan` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, session-bus smoke coverage, Runtime dispatch, and KDE-safe Compatibility Center read model.
- Added Ruby development harness and C Runtime CLI coverage for KDE shell planning while keeping Plasma fork/source/config writes, component activation, backend launch, host-root mutation, privileged-container, and backend-detail gates disabled.

### Changed

- Extended Runtime diagnostics and Compatibility Center summaries with KDE shell component planning so KDE remains a replaceable desktop shell while Runtime owns compatibility policy.

## [0.2.109] - 2026-07-15

### Added

- Added Runtime-owned backend selection plans that explain recommended local or isolated compatibility profiles before any selection is committed.
- Added `GetBackendSelectionPlan` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, session-bus smoke coverage, Runtime dispatch, and KDE-safe Compatibility Center read model.
- Added Ruby development harness and C Runtime CLI coverage for backend selection planning while keeping selection, launch, activation, environment, request-object, state-root, snapshot, host-root, privileged-container, and backend-detail gates disabled.

### Changed

- Extended Runtime diagnostics and Compatibility Center summaries with pending backend selection state so KDE can explain Runtime profile recommendations without owning backend policy.

## [0.2.108] - 2026-07-15

### Added

- Added C Runtime backend capability matrix records and a CLI read method for local and isolated compatibility profiles.
- Added `GetBackendCapabilityMatrix` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, session-bus smoke coverage, Runtime dispatch, and KDE-safe Compatibility Center read model.
- Added Ruby development harness and unit coverage for application launch, package management, file bridge, clipboard bridge, print bridge, snapshot restore, and diagnostics capability planning.

### Changed

- Extended Runtime diagnostics and Compatibility Center summaries with backend capability matrix state while keeping selection, launch, activation, request-object, state-root, snapshot, host-root, privileged-container, and backend-detail gates disabled.

## [0.2.107] - 2026-07-15

### Added

- Added Runtime-owned compatibility review flow plans that connect settings change review, permission review, Portal request review, Runtime write gates, and review receipts into one KDE-visible confirmation flow.
- Added `GetCompatibilityReviewFlowPlan` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, session-bus smoke coverage, Runtime dispatch, and KDE-safe Compatibility Center read model.
- Added Ruby development harness and C Runtime CLI coverage for review flow planning while keeping all apply, request-object, permission-grant, settings-persistence, execution, host-root, and backend-detail gates disabled.

### Changed

- Extended Runtime diagnostics and Compatibility Center summaries with pending review flow counts so KDE can explain the full user confirmation path before any sensitive compatibility setting or desktop resource change is applied.

## [0.2.106] - 2026-07-14

### Added

- Added C Runtime compatibility permission review plan records and a CLI read method for documents, downloads, camera, network, clipboard, print, and screenshot permissions.
- Added `GetCompatibilityPermissionReviewPlan` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, session-bus smoke coverage, Runtime dispatch, and KDE-safe Compatibility Center read model.
- Added Ruby development harness and unit coverage that reject backend implementation terminology while asserting permission review plans remain non-mutating.

### Changed

- Extended Runtime diagnostics with compatibility permission review summaries that keep permission changes, request object creation, permission grants, settings persistence, host-root mutation, and backend detail exposure disabled until Runtime gates pass.

## [0.2.105] - 2026-07-14

### Added

- Added C Runtime compatibility mode switch plan records and a CLI read method for `automatic`, `prefer-performance`, `prefer-compatibility`, and `isolated-execution` user-facing modes.
- Added `GetCompatibilityModeSwitchPlan` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, session-bus smoke coverage, Runtime dispatch, and KDE-safe Compatibility Center read model.
- Added Ruby development harness and unit coverage that reject unsupported modes and assert mode-switch JSON does not expose backend implementation terminology.

### Changed

- Extended Runtime diagnostics with compatibility mode switch summaries that keep settings persistence, backend reconfiguration, backend process starts, launch enablement, host-root mutation, and backend detail exposure disabled until Runtime gates pass.

## [0.2.104] - 2026-07-14

### Added

- Added C Runtime desktop resource bridge plan records and a CLI read method for file, URI, print, clipboard, and screenshot bridge readiness.
- Added `GetDesktopResourceBridgePlan` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, session-bus smoke coverage, Runtime dispatch, and KDE-safe Compatibility Center read model.
- Added Ruby development harness and unit coverage for desktop resource bridge summaries while keeping the C Runtime as the product boundary.

### Changed

- Extended Runtime diagnostics with Portal-mediated desktop resource bridge summaries that keep bridge enablement, Portal request creation, backend process starts, direct host file access, direct clipboard access, direct print access, host-root mutation, and backend detail exposure disabled until Runtime gates pass.

## [0.2.103] - 2026-07-14

### Added

- Added C Runtime KDE application surface plan records and a CLI read method that combine launcher, task manager, Dolphin, tray, notification, Compatibility Center, and settings entry points for one managed compatibility application.
- Added `GetKDEApplicationSurfacePlan` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, session-bus smoke coverage, Runtime dispatch, and KDE-safe Compatibility Center read model.
- Added Ruby development harness and unit coverage for KDE application surface summaries while keeping the C Runtime as the product boundary.

### Changed

- Extended Runtime diagnostics with KDE application surface summaries that present Windows applications as normal Linux applications while keeping backend launch, backend process starts, desktop-file writes, MIME writes, host-root mutation, raw executable exposure, backend command exposure, and backend detail exposure disabled until Runtime gates pass.

## [0.2.102] - 2026-07-14

### Added

- Added C Runtime backend environment plan records and a CLI read method for local and isolated compatibility environment readiness.
- Added `GetBackendEnvironmentPlan` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, session-bus smoke coverage, and Compatibility Center read model.
- Added Ruby development harness and unit coverage for backend environment plan summaries while keeping C as the product boundary.

### Changed

- Extended Runtime diagnostics with backend environment summaries that keep environment creation, backend process starts, clipboard and print bridges, host storage exposure, launch enablement, network access, and backend detail exposure disabled until Runtime gates pass.

## [0.2.101] - 2026-07-14

### Added

- Added C Runtime backend lifecycle records and a CLI read method so KDE can display backend lifecycle state without starting local or isolated compatibility backends.
- Added `GetBackendLifecycle` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, session-bus smoke coverage, and Compatibility Center read model.
- Added Ruby development harness and unit coverage for backend lifecycle summaries while keeping product-side lifecycle logic mirrored by the C Runtime core.

### Changed

- Extended Runtime diagnostics with backend lifecycle summaries that keep backend process starts, launch enablement, host-root mutation, network access, and backend detail exposure disabled until Runtime gates pass.

## [0.2.100] - 2026-07-14

### Changed

- Split the C Runtime CLI probe output into `xnix_runtime_core_cli_probe.inc` so the main CLI translation unit stays below the repository file-size guardrail.
- Kept the C Runtime probe contract unchanged while preparing the CLI surface for additional Runtime-owned KDE and compatibility planning reads.

## [0.2.99] - 2026-07-14

### Added

- Added C Runtime launch intent records and a CLI read method so KDE launcher clicks can be modeled before any `Launch` write method is enabled.
- Added `GetLaunchIntent` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, and session-bus smoke coverage.
- Extended the Ruby development harness and launch request unit coverage to mirror the Runtime-owned launch intent boundary.

### Changed

- Extended Runtime diagnostics with launch intent summaries that keep KDE from creating launch request objects, starting Wine/VM backends, mutating the host root, or exposing backend commands before Runtime gates pass.

## [0.2.98] - 2026-07-14

### Added

- Added C Runtime execution readiness records and a CLI read method that summarizes launch gates before KDE exposes launch intent.
- Added `GetExecutionReadiness` to the Runtime D-Bus contract, D-Bus client, smoke adapter, method parity manifest, and session-bus smoke coverage.
- Added a Ruby development harness and unit test for execution readiness while keeping C as the product boundary.

### Changed

- Extended Runtime diagnostics with execution readiness summaries that keep Launch blocked until backend binding, Portal review, snapshot baseline, and Runtime write gates are ready.

## [0.2.97] - 2026-07-14

### Added

- Added C Runtime compatibility test result records and a CLI read method for passed, pending, and blocked compatibility test outcome summaries.
- Added C Runtime probe ownership for compatibility test results so KDE and AI diagnostics can verify result coverage without owning test execution.

### Changed

- Marked compatibility test results as C Runtime-backed in the Ruby and D-Bus smoke layers while keeping test execution, host-root mutation, and backend detail exposure disabled.

## [0.2.96] - 2026-07-14

### Added

- Added C Runtime compatibility test plan records and a CLI read method for recipe validation, Portal preflight, snapshot preflight, and managed launch-binding checks.
- Added C Runtime probe ownership for compatibility test plans so KDE can verify test type coverage without owning test execution policy.

### Changed

- Marked compatibility test plans as C Runtime-backed in the Ruby and D-Bus smoke layers while keeping execution requests, test execution, host-root mutation, and backend detail exposure disabled.

## [0.2.95] - 2026-07-14

### Added

- Added C Runtime compatibility snapshot plan records and a CLI read method for Runtime-owned restore-point planning before repair, engine changes, and manual snapshots.
- Added C Runtime probe ownership for compatibility snapshot plans so KDE can verify snapshot reason coverage without owning snapshot or restore policy.

### Changed

- Marked compatibility snapshot plans as C Runtime-backed in the Ruby and D-Bus smoke layers while keeping snapshot creation, restore execution, host-root mutation, user-document capture, host-system capture, and backend detail exposure disabled.

## [0.2.94] - 2026-07-14

### Added

- Added C Runtime compatibility repair plan records and a CLI read method for diagnostic issue repair planning, approval requirements, snapshot summaries, rollback availability, notification mapping, and non-execution safety gates.
- Added C Runtime probe ownership for compatibility repair plans so KDE can verify repair issue coverage without owning repair execution policy.

### Changed

- Marked compatibility repair plans as C Runtime-backed in the Ruby and D-Bus smoke layers while keeping repair execution, backend launch, network access, host-root mutation, and backend detail exposure disabled.

## [0.2.93] - 2026-07-14

### Added

- Added C Runtime AI repair approval gate records and a CLI read method for review-first repair gates, required approval gates, approval-required actions, and blocked repair actions.
- Added C Runtime probe ownership for AI repair approval gates so KDE can verify Runtime-owned repair gating without owning repair execution policy.

### Changed

- Marked AI repair approval gates as C Runtime-backed in the Ruby and D-Bus smoke layers while keeping provider calls, network access, automatic execution, repair execution, host-root mutation, and backend detail exposure disabled.

## [0.2.92] - 2026-07-14

### Added

- Added C Runtime AI diagnostic recommendation records and a CLI read method for review-first recommendations, approval-required actions, and blocked AI actions.
- Added C Runtime probe ownership for AI diagnostic recommendations so KDE can verify recommendation counts and execution gates without owning AI policy.

### Changed

- Marked AI diagnostic recommendations as C Runtime-backed in the Ruby and D-Bus smoke layers while keeping provider calls, network access, automatic execution, host-root mutation, and backend detail exposure disabled.

## [0.2.91] - 2026-07-14

### Added

- Added C Runtime AI diagnostic input records and a CLI read method for Runtime-safe diagnostic context, signals, privacy boundaries, and AI task policy.
- Added C Runtime probe ownership for AI diagnostic input safety so KDE and tests can verify AI diagnostic boundaries without relying on Ruby product logic.

### Changed

- Marked AI diagnostic inputs as C Runtime-backed in the Ruby and D-Bus smoke layers while keeping provider calls, network access, host-root mutation, and backend detail exposure disabled.

## [0.2.90] - 2026-07-14

### Added

- Added C Runtime KDE integration status records and a CLI read method for the seven first-release KDE entry points.
- Added Runtime D-Bus contract coverage for `GetKDEIntegrationStatus` so KDE surfaces can read official desktop scope and entry-point readiness from the Runtime.

### Changed

- Updated `xnix-kde-integration-status` to consume Runtime-owned KDE integration status before falling back to local data.
- Extended Runtime method parity, daemon dispatch, D-Bus smoke coverage, and the D-Bus client to include KDE integration status while keeping KDE as a replaceable shell and backend policy Runtime-owned.

## [0.2.89] - 2026-07-14

### Added

- Added C Runtime KWin window rule plan records and a CLI read method for Runtime-owned identity and layout hints.
- Added Runtime D-Bus contract coverage for `GetKWinWindowRulePlan` so KDE KWin surfaces can read Runtime-owned window-rule plans without owning backend policy.

### Changed

- Updated the KDE-facing KWin window rule model to consume Runtime-owned window-rule plans before falling back to local task-manager identity data.
- Extended Runtime method parity, D-Bus smoke coverage, daemon diagnostics, the D-Bus client, and layout verification to include KWin window rule planning while keeping host-root mutation and backend details disabled.

## [0.2.88] - 2026-07-14

### Changed

- Updated the KDE KRunner read model to consume Runtime-owned KRunner query plans before falling back to local application-list matching.
- Normalized both nested local Runtime query plans and flat D-Bus smoke query plans into KDE-safe KRunner matches while preserving Runtime launch and execution gates.

### Fixed

- Prevented the KDE-facing KRunner model from duplicating Runtime query policy when `krunner_query_plan` is available.

## [0.2.87] - 2026-07-14

### Added

- Added C Runtime KRunner query plan records and a CLI read method for resolving desktop search queries to Runtime application identities and managed launcher actions.
- Added Runtime D-Bus contract coverage for `GetKRunnerQueryPlan` so KDE KRunner surfaces can read Runtime-owned query plans without enabling direct execution.

### Changed

- Extended Runtime method parity, D-Bus smoke coverage, daemon diagnostics, the D-Bus client, and layout verification to include KRunner query planning while keeping backend launch, query execution, host-root mutation, and backend details disabled.

## [0.2.86] - 2026-07-14

### Added

- Added C Runtime Compatibility Center summary records and a CLI read method for application state, known issues, repair records, and safe KDE navigation actions.
- Added Runtime D-Bus contract coverage for `GetCompatibilityCenterSummary` so KDE Compatibility Center surfaces can read a Runtime-owned center summary without enabling execution.

### Changed

- Extended Runtime method parity, D-Bus smoke coverage, daemon diagnostics, the D-Bus client, and the KDE Center read model to include Compatibility Center summaries while keeping action execution, repair execution, backend launch, settings persistence, host-root mutation, and backend details disabled.

## [0.2.85] - 2026-07-14

### Added

- Added C Runtime tray status plan records and a CLI read method for KDE system tray Runtime activity, attention state, tray bridge readiness, and user navigation actions.
- Added Runtime D-Bus contract coverage for `GetTrayStatus` so KDE tray surfaces can read Runtime status without enabling live tray bridging or backend policy ownership.

### Changed

- Extended Runtime method parity, D-Bus smoke coverage, daemon diagnostics, and the D-Bus client to include tray status planning while keeping live tray bridging, bridge persistence, host-root mutation, and backend details disabled.

## [0.2.84] - 2026-07-14

### Added

- Added C Runtime notification plan records and a CLI read method for KDE notification events covering install failures, repair receipts, mode changes, and approval requests.
- Added Runtime D-Bus contract coverage for `GetNotificationPlan` so KDE notification surfaces can read event plans without enabling execution or exposing backend details.

### Changed

- Extended Runtime method parity, D-Bus smoke coverage, daemon diagnostics, and the D-Bus client to include notification planning while keeping action execution, repair execution, settings persistence, and host-root mutation disabled.

## [0.2.83] - 2026-07-14

### Added

- Added C Runtime file association plan records and a CLI read method for mapping recipe MIME types to generated desktop entries and portal-mediated file opens.
- Added Runtime D-Bus contract coverage for `GetFileAssociationPlan` so KDE file-manager surfaces can read MIME association plans without writing files or seeing backend details.

### Changed

- Extended Runtime method parity, D-Bus smoke coverage, and diagnostics to include file association planning while keeping `mimeapps.list` changes staged and non-overwriting.

## [0.2.82] - 2026-07-14

### Added

- Added C Runtime task manager identity plan records and a CLI read method for grouping, pinning, switching, and restoring compatibility windows as normal KDE taskbar entries.
- Added Runtime D-Bus contract coverage for `GetTaskManagerIdentityPlan` so KDE surfaces can read window identity hints without seeing backend implementation details.

### Changed

- Extended Runtime method parity, D-Bus smoke coverage, and diagnostics to include task manager identity planning while keeping KWin policy identity-only and backend policy Runtime-owned.

## [0.2.81] - 2026-07-14

### Added

- Added C Runtime desktop entry plan records and a CLI read method for exposing compatibility applications as normal Linux `.desktop` launcher identities.
- Added Runtime D-Bus contract coverage for `GetDesktopEntryPlan` so KDE surfaces can read launcher identity without seeing backend commands, raw Windows executables, or prefix paths.

### Changed

- Extended Runtime method parity, D-Bus smoke coverage, and diagnostics to include desktop entry planning while keeping file writes, host-root mutation, and backend detail exposure disabled.

## [0.2.80] - 2026-07-14

### Added

- Added C Runtime Portal request plan records and a CLI read method for mapping desktop resource requests to XDG Desktop Portal calls without granting permissions.
- Added Runtime D-Bus contract coverage for `GetPortalRequestPlan` so KDE surfaces can request portal planning through the stable Runtime API.

### Changed

- Extended Runtime C ownership into Portal request planning while keeping request creation, permission grants, host permission changes, and host-root mutation disabled.

## [0.2.79] - 2026-07-14

### Added

- Added C Runtime desktop activation manifest records and a CLI read method for describing KDE launcher, task manager, file manager, tray, notification, Compatibility Center, and settings entry points.
- Added Runtime D-Bus contract coverage for `GetDesktopActivationManifest` so desktop shells can consume a stable read-only activation contract.

### Changed

- Extended Runtime C ownership into desktop activation planning while keeping desktop file writes, host-root mutation, and backend detail exposure disabled.

## [0.2.78] - 2026-07-14

### Added

- Added C Runtime compatibility run plan records and a CLI read method for mapping registered applications to Runtime-owned execution strategies.
- Added C Runtime run plan tests that verify recipe-based strategy selection, required preflight, disabled launch requests, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first compatibility run plan layer while keeping backend launch disabled until managed binding is ready.

## [0.2.77] - 2026-07-14

### Added

- Added C Runtime Compatibility Center action review receipt records and a CLI read method for recording KDE review intent without execution authority.
- Added C Runtime action review receipt tests that verify recorded decisions, deferred reviews, Runtime gate preservation, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first Compatibility Center action review receipt layer while keeping KDE review intent separate from Runtime execution approval.

## [0.2.76] - 2026-07-14

### Added

- Added C Runtime Compatibility Center action queue records and a CLI read method for KDE task-card planning.
- Added C Runtime action queue tests that verify action ordering, review counts, disabled execution, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first Compatibility Center action queue layer while keeping KDE limited to display and review intent.

## [0.2.75] - 2026-07-14

### Added

- Added C Runtime compatibility install plan records and a CLI read method for joining application, install readiness, and recipe install gate decisions.
- Added C Runtime compatibility install plan tests that verify development staging, production blocking, readiness aggregation, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first compatibility install plan layer while keeping downloads, package installation, desktop activation, and backend launch disabled.

## [0.2.74] - 2026-07-14

### Added

- Added C Runtime recipe install gate records and a CLI read method for evaluating production and development install decisions from the C Runtime core.
- Added C Runtime recipe install gate tests that verify production blocking, development staging, unknown recipe blocking, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first recipe install gate layer while keeping production installation disabled for development-only registries.

## [0.2.73] - 2026-07-14

### Added

- Added C Runtime recipe trust policy records and a CLI read method for reporting development-only registry trust, signature readiness, and production install blocking.
- Added C Runtime recipe trust tests that verify digest checks, pending signature validation, development registry warnings, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first recipe trust policy layer while keeping production recipe trust and production install disabled.

## [0.2.72] - 2026-07-14

### Added

- Added C Runtime method parity manifest policy records and a CLI read method for reporting read-only D-Bus method coverage across contract, dispatch, client, smoke adapter, and session smoke.
- Added C Runtime method parity tests that verify the 28 read-only methods, parity check counts, gated write methods, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first method parity manifest layer while keeping all write methods disabled.

## [0.2.71] - 2026-07-14

### Added

- Added C Runtime owner smoke plan policy records and a CLI read method for reporting the planned production owner smoke sequence.
- Added C Runtime owner smoke plan tests that verify smoke step ordering, readiness counts, disabled production bus claims, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first owner smoke plan layer while keeping production bus ownership and system service startup disabled.

## [0.2.70] - 2026-07-14

### Added

- Added C Runtime live owner gate policy records and a CLI read method for reporting production D-Bus ownership transition gates.
- Added C Runtime live owner gate tests that verify required gate ordering, pending production ownership, KDE non-ownership, host-root safety, network safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first live owner gate layer while keeping the smoke adapter separate from production Runtime ownership.

## [0.2.69] - 2026-07-14

### Added

- Added C Runtime service binding policy records and a CLI read method for reporting Runtime D-Bus activation, systemd hardening, wrapper, contract, and live-owner gates.
- Added C Runtime service binding tests that verify activation metadata, readiness counts, pending live ownership, host-root safety, network safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first service binding layer while keeping production D-Bus ownership disabled and KDE limited to read-only status consumption.

## [0.2.68] - 2026-07-14

### Added

- Added C Runtime compatibility settings change policy records and CLI read methods for listing planned settings-change gates and evaluating per-application settings change plans.
- Added C Runtime settings change tests that verify confirmation, Portal review, persistence blocking, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first compatibility settings change layer while keeping settings persistence, resource grants, Runtime restarts, and host-root mutation disabled.

## [0.2.67] - 2026-07-13

### Added

- Added C Runtime compatibility settings policy records and CLI read methods for listing settings gates and evaluating per-application KDE settings models.
- Added C Runtime settings tests that verify user-facing sections, default values, Runtime ownership, blocked persistence, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first compatibility settings layer while keeping settings persistence, host-root mutation, and backend terminology out of KDE-facing settings data.

## [0.2.66] - 2026-07-13

### Added

- Added C Runtime backend binding policy records and CLI read methods for listing managed backend binding gates and evaluating per-application launch-binding readiness.
- Added C Runtime backend binding tests that verify required preflight, blocked launch actions, blocked execution requests, host-root safety, network safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first managed backend binding layer while keeping launch enablement, execution request creation, privileged containers, and host-root mutation disabled.

## [0.2.65] - 2026-07-13

### Added

- Added C Runtime package source policy records and CLI read methods for listing package-source gates and evaluating per-application source selection readiness.
- Added C Runtime package source tests that verify source channels, required preflight, signed-source policy, blocked package-manager command exposure, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first compatibility package source layer while keeping package installation, host package manager invocation, network planning requirements, and host-root mutation disabled.

## [0.2.64] - 2026-07-13

### Added

- Added C Runtime acquisition preflight policy records and CLI read methods for listing acquisition gates and evaluating per-application acquisition readiness.
- Added C Runtime acquisition preflight tests that verify preflight checks, blocked network requests, blocked downloads, pending package-source readiness, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first compatibility acquisition preflight layer while keeping network requests, artifact downloads, installation, and host-root mutation disabled.

## [0.2.63] - 2026-07-13

### Added

- Added C Runtime artifact manifest policy records and CLI read methods for listing artifact manifest gates and evaluating per-application artifact readiness.
- Added C Runtime artifact manifest tests that verify artifact groups, required preflight, blocked downloads, blocked cache activation, host-root safety, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first compatibility artifact manifest layer while keeping artifact downloads, cache activation, installation, and host-root mutation disabled.

## [0.2.62] - 2026-07-13

### Added

- Added C Runtime install readiness policy records and CLI read methods for listing install gates and evaluating per-application install readiness by environment.
- Added C Runtime install readiness tests that verify blocked downloads, blocked installation, development recipe staging, production recipe blocking, phase order, and host-root safety.

### Changed

- Extended Runtime C ownership into the first compatibility install readiness layer while keeping package downloads, installation, desktop activation, and backend launch disabled.

## [0.2.61] - 2026-07-13

### Added

- Added C Runtime application state root policy records and CLI read methods for listing state-root policies and evaluating a per-application state root plan.
- Added C Runtime state-root policy tests that verify managed scopes, retention boundaries, Portal file requirements, restore confirmation, user-document exclusion, and host-root safety.

### Changed

- Extended Runtime C ownership into the first application state-root policy layer while keeping directory creation, host-root mutation, and backend-detail exposure disabled.

## [0.2.60] - 2026-07-13

### Added

- Added C Runtime snapshot policy records and CLI read methods for listing restore-point policies and evaluating per-application snapshot plans.
- Added C Runtime snapshot policy tests that verify restore scope, bounded retention, Runtime ownership, user-document preservation, host-system exclusion, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first snapshot and rollback policy layer while keeping host-root mutation disabled.

## [0.2.59] - 2026-07-13

### Added

- Added C Runtime Portal access policy records and CLI read methods for listing sensitive desktop operations and evaluating per-application Portal policy.
- Added C Runtime Portal policy tests that verify ask and deny defaults, Runtime ownership, KDE non-ownership, direct-access denial, and backend-detail filtering.

### Changed

- Extended Runtime C ownership into the first XDG Desktop Portal permission policy layer while keeping all host permission changes disabled.

## [0.2.58] - 2026-07-13

### Added

- Added C Runtime compatibility engine catalog records and CLI read methods for listing Runtime engine strategies and selecting an engine by recipe mode.
- Added C Runtime engine catalog tests that verify stable strategy order, blocked launch readiness, C ownership, and backend-detail filtering.

### Changed

- Extended Runtime C ownership from application metadata into the first engine selection layer while keeping all backend execution disabled.

## [0.2.57] - 2026-07-13

### Added

- Added C Runtime application catalog records and CLI read methods for listing and querying safe desktop application metadata.
- Added C Runtime catalog tests that compile the C core, verify the sample application record, reject unknown application ids, and guard against backend detail exposure.

### Changed

- Extended Runtime C ownership from identity and write gates into the first application recipe metadata layer while keeping Ruby as the test and development-tool harness.

## [0.2.56] - 2026-07-13

### Added

- Added a C Runtime core library and CLI for stable Runtime identity, reserved write-method enumeration, and write-gate decisions.
- Added C Runtime core tests that compile and execute the C implementation while keeping Ruby as the test harness.

### Changed

- Documented the Runtime implementation direction so important product logic moves into C while Ruby remains focused on tests and development tooling.

## [0.2.55] - 2026-07-13

### Added

- Added `scripts/runtime_activation_smoke.rb`, a constrained packaged Runtime activation smoke that installs into a temporary root and executes the staged libexec wrapper against installed Runtime assets.

### Changed

- Extended Runtime activation installation to include Runtime Ruby libraries, version metadata, recipe assets, D-Bus contract references, and smoke references required by the packaged wrapper.

## [0.2.54] - 2026-07-13

### Added

- Added Runtime-owned write gate modeling through `xnix-runtime-write-gate`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Changed gated Runtime write-method dispatch so install, launch, snapshot, and restore requests now fail with an explicit `WriteMethodDisabled` boundary instead of a generic unsupported-method error.

## [0.2.53] - 2026-07-13

### Added

- Added Runtime-owned read-only method parity manifests through `xnix-runtime-method-parity-manifest`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Extended production owner smoke readiness so XML contract, Runtime dispatch, D-Bus client, smoke adapter, and session smoke coverage must agree on the Runtime read-only method set before write methods are considered.

## [0.2.52] - 2026-07-13

### Added

- Added Runtime-owned production owner smoke planning through `xnix-runtime-owner-smoke-plan`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Extended live-owner readiness so the future production Runtime owner has explicit smoke steps for bus-name ownership, read-only method parity, write-method rejection, and smoke-adapter boundary checks.

## [0.2.51] - 2026-07-13

### Added

- Added Runtime-owned live D-Bus owner gate modeling through `xnix-runtime-live-owner-gate`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Extended Runtime service readiness so KDE can distinguish aligned activation files from gated production D-Bus ownership, while keeping the smoke adapter explicitly non-production.

## [0.2.50] - 2026-07-13

### Added

- Added Runtime-owned Compatibility Center action review receipts through `xnix-compat-action-review`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Extended Compatibility Center action handling so KDE review intent is recorded as a non-executing Runtime receipt that cannot launch backends, persist settings, grant resources, or mutate the host root.

## [0.2.49] - 2026-07-13

### Added

- Added Runtime-owned Compatibility Center action queues through `xnix-compat-action-queue`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, KDE Center summaries, and session-bus smoke coverage.

### Changed

- Extended KDE-facing Runtime read models so pending install readiness, settings changes, AI repair review, Runtime service ownership, and Portal policy review are grouped into non-executing Compatibility Center action cards.

## [0.2.48] - 2026-07-13

### Added

- Added Runtime-owned compatibility settings change planning through `xnix-compat-settings-change`, diagnostics, read-only dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, and session-bus smoke coverage.

### Changed

- Extended compatibility settings so KDE can request a safe Runtime plan for setting changes without persisting state, granting resources, mutating the host root, or exposing backend details.

## [0.2.47] - 2026-07-13

### Added

- Added Runtime-backed compatibility settings exposure through diagnostics, read-only dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, and session-bus smoke coverage.

### Changed

- Extended `xnix-compat-settings` so KDE can display Runtime-owned run mode, resource access, device, network, and snapshot settings without owning backend policy or exposing implementation terminology.

## [0.2.46] - 2026-07-13

### Added

- Added `xnix-compat-install-plan`, a Runtime-owned compatibility install plan that joins artifact manifest, acquisition preflight, package source, state root, and recipe install gate readiness before any desktop activation or backend launch.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for install plan status.

### Changed

- Extended Runtime and KDE read models so installation readiness is visible as a first-class Compatibility Center signal before artifact download, desktop activation, or backend launch is enabled.

## [0.2.45] - 2026-07-13

### Added

- Added `xnix-compat-artifact-manifest`, a Runtime-owned compatibility artifact manifest model for signed manifest readiness, artifact groups, digest verification, cache namespaces, and rollback references without downloading artifacts or exposing cache paths.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for artifact manifest status.

### Changed

- Extended Runtime and KDE read models so artifact manifest readiness is visible as a first-class Compatibility Center signal before artifact download, installation, or backend launch is enabled.

## [0.2.44] - 2026-07-13

### Added

- Added `xnix-compat-acquisition-preflight`, a Runtime-owned acquisition preflight model for package-source readiness, signed artifact manifests, cache capacity, network policy review, and rollback markers without downloading artifacts or mutating the host root.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for acquisition preflight status.

### Changed

- Extended Runtime and KDE read models so acquisition readiness is visible as a first-class Compatibility Center signal before package download, installation, or backend launch is enabled.

## [0.2.43] - 2026-07-13

### Added

- Added `xnix-compat-package-source`, a Runtime-owned compatibility package source model for signed source policy, source-channel planning, cache preflight, and blocked package installation before launch binding.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for package source status.

### Changed

- Extended Runtime and KDE read models so package source selection is visible as a first-class Compatibility Center signal before backend launch is enabled.

## [0.2.42] - 2026-07-13

### Added

- Added `xnix-compat-state-root`, a Runtime-owned application state root model for per-application state ownership, snapshot eligibility, Portal file boundaries, and restore confirmation without creating host directories or exposing storage paths.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for application state root status.

### Changed

- Extended Runtime and KDE read models so planned state roots are visible as a first-class Compatibility Center signal before backend launch is enabled.

## [0.2.41] - 2026-07-13

### Added

- Added `xnix-compat-backend-binding`, a Runtime-owned managed compatibility backend binding model that reports launch readiness, required preflight, blocked unsafe actions, and desktop-safe status without exposing implementation details.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke coverage, and focused tests for backend binding status.

### Changed

- Extended Runtime and KDE read models so pending backend binding is visible as a first-class Compatibility Center signal before any launch backend is enabled.

## [0.2.40] - 2026-07-13

### Added

- Added `xnix-runtime-service-binding`, a Runtime-owned service binding model for D-Bus activation files, systemd hardening, the packaged libexec wrapper, contract exposure, and live-owner readiness.
- Added read-only Runtime diagnostics, dispatch, D-Bus contract, D-Bus client, Compatibility Center summaries, session-bus smoke exposure, and focused coverage for Runtime service binding status.

### Changed

- Extended the Runtime foundation documentation to distinguish activation binding readiness from live production D-Bus ownership.
- Fixed the full smoke sequence so a clean milestone run builds the constrained Docker image and fetches Buildroot before configuring, downloading package sources, building, and booting QEMU.

## [0.2.39] - 2026-07-13

### Added

- Added `xnix-ai-repair-approval-gate`, a Runtime-owned approval gate model that blocks AI repair execution until Compatibility Center review, Runtime approval, and restore-point preflight gates pass.
- Added focused coverage for blocked gate decisions, required gates, approval-required actions, Runtime diagnostics, dispatch, D-Bus client reads, Compatibility Center summaries, and session-bus smoke exposure.

### Changed

- Extended Runtime diagnostics, read-only Runtime dispatch, D-Bus planning reads, and the KDE Compatibility Center model with AI repair approval gate summaries.

## [0.2.38] - 2026-07-13

### Added

- Added `xnix-ai-diagnostic-recommendation`, a Runtime-owned AI diagnostic recommendation model for review-first compatibility guidance.
- Added focused coverage for user-visible recommendations, approval-required actions, blocked AI tasks, Runtime diagnostics, dispatch, D-Bus client reads, Compatibility Center summaries, and session-bus smoke exposure.

### Changed

- Extended Runtime diagnostics, read-only Runtime dispatch, D-Bus planning reads, and the KDE Compatibility Center model with AI diagnostic recommendation summaries.

## [0.2.37] - 2026-07-13

### Added

- Added `xnix-ai-diagnostic-input`, a Runtime-owned AI diagnostic input model for safe recipe, run-plan, test-result, repair, signal, and privacy-boundary context.
- Added focused coverage for AI diagnostic safety boundaries, allowed and blocked AI tasks, Runtime diagnostics, dispatch, D-Bus client reads, Compatibility Center summaries, and session-bus smoke exposure.

### Changed

- Extended Runtime diagnostics, read-only Runtime dispatch, D-Bus planning reads, and the KDE Compatibility Center model with AI diagnostic input summaries.

## [0.2.36] - 2026-07-13

### Added

- Added `xnix-compat-test-result`, a Runtime-owned compatibility test result model for passed, pending, and blocked test outcomes.
- Added focused coverage for test result counts, diagnostic evidence, Compatibility Center summaries, Runtime diagnostics, dispatch, D-Bus client reads, and session-bus smoke exposure.

### Changed

- Extended Runtime diagnostics, read-only Runtime dispatch, D-Bus planning reads, and the KDE Compatibility Center model with compatibility test result summaries.

## [0.2.35] - 2026-07-13

### Added

- Added `xnix-compat-test-plan`, a Runtime-owned compatibility test plan model for preflight, smoke, and repair-readiness checks.
- Added focused coverage for recipe validation, Portal preflight, snapshot preflight, launch-binding readiness, Compatibility Center summaries, Runtime diagnostics, dispatch, D-Bus client reads, and session-bus smoke exposure.

### Changed

- Extended Runtime diagnostics, read-only Runtime dispatch, D-Bus planning reads, and the KDE Compatibility Center model with compatibility test plan summaries.

## [0.2.34] - 2026-07-13

### Added

- Added `xnix-portal-request-model`, a Runtime-owned XDG Desktop Portal request model for file, URI, print, screenshot, clipboard, camera, and remote-desktop operations.
- Added focused coverage for Portal D-Bus destination, interface, method, request handle tokens, Response completion, denied-request guidance, CLI validation, and backend-detail filtering.

### Changed

- Extended desktop integration manifests so settings artifacts expose both Portal policy evaluation and Portal request modeling.

## [0.2.33] - 2026-07-13

### Added

- Added `xnix-file-association-model`, a Runtime-owned file association model that maps recipe MIME types to generated desktop files and standard `mimeapps.list` content.
- Added activation staging for `usr/share/applications/mimeapps.list` so managed applications can become default handlers for their recipe file types inside a staging root.
- Added focused coverage for MIME defaults, portal-mediated file-open metadata, rollback tracking, and refusal to overwrite an existing `mimeapps.list`.

### Changed

- Extended desktop integration manifests so the file-manager artifact exposes file association generation alongside the Dolphin service menu.

## [0.2.32] - 2026-07-13

### Added

- Added `xnix-kwin-window-rule`, a KDE KWin window rule model that binds Runtime application identity to generated desktop files, task grouping, taskbar visibility, switcher visibility, and restore behavior.
- Added focused coverage for KWin rule identity matching, task-manager preservation, bounded window-manager scope, CLI validation, and backend-detail filtering.

### Changed

- Extended desktop integration manifests so the task-manager artifact exposes both the window identity model and the KWin window rule model.

## [0.2.31] - 2026-07-13

### Added

- Added `xnix-krunner-model`, a KDE KRunner query model that resolves Runtime-managed applications from names and file-oriented natural queries.
- Added focused coverage for KRunner launch delegation, desktop entry mapping, blank query behavior, extension hints, and backend-detail filtering.

### Changed

- Made `DBusRuntimeClient` parse D-Bus string arrays so KDE-facing read paths preserve list fields such as supported file extensions.

## [0.2.30] - 2026-07-13

### Added

- Added `DBusRuntimeClient` wrappers for Runtime planning reads covering engine catalog, run plans, repair plans, snapshot plans, and Portal access policy.
- Added focused D-Bus client coverage for planning reads and D-Bus boolean variant parsing.

### Changed

- Made the KDE-facing D-Bus client parse boolean `true` and `false` values into native booleans instead of strings.

### Verified

- Ran the full Buildroot and QEMU serial smoke test required for the 0.2.30 milestone.

## [0.2.29] - 2026-07-13

### Changed

- Extended the Runtime D-Bus contract with read-only planning methods for engine catalog, run plans, repair plans, snapshot plans, and Portal access policy.
- Routed the new read-only planning methods through `xnix-compatd dispatch`.
- Updated the Linux session-bus smoke adapter and smoke script to expose and call the new read-only planning methods.

### Added

- Added focused contract and dispatch coverage for the new planning methods while keeping write operations unsupported until backends exist.

## [0.2.28] - 2026-07-13

### Added

- Added `xnix-compat-engine-catalog`, a Runtime-owned catalog for automatic, local, and isolated compatibility engine choices.
- Added focused coverage for engine selection, Runtime ownership, pending backend readiness, CLI output, and backend-term filtering.

### Changed

- Made compatibility run plans select their engine summary through the engine catalog.
- Added Runtime probe capability reporting for the compatibility engine catalog.

## [0.2.27] - 2026-07-13

### Added

- Added `xnix-compat-snapshot-plan`, a Runtime-owned snapshot planning model for repair, compatibility-engine changes, and manual restore points.
- Added focused coverage for snapshot scope, restore availability, bounded retention, user-document preservation, CLI validation, and backend-term filtering.

### Changed

- Attached snapshot plan summaries to repair plans, Runtime diagnostics, and the KDE Compatibility Center model when a repair requires a restore point.

## [0.2.26] - 2026-07-13

### Added

- Added `xnix-compat-repair-plan`, a Runtime-owned repair planning model for pending engine setup, Portal approval, recipe trust blocking, and applied repair records.
- Added focused coverage for user approval requirements, snapshot requirements, rollback availability, notification mapping, CLI validation, and backend-term filtering.

### Changed

- Extended Runtime diagnostics and KDE Compatibility Center models with desktop-safe repair summaries.
- Reworded pending engine diagnostics so desktop-facing diagnostics do not expose implementation-specific backend terms.

## [0.2.25] - 2026-07-13

### Added

- Added `xnix-compat-run-plan`, a Runtime-owned compatibility run plan model for automatic, local, and isolated execution strategies.
- Added focused coverage for desktop-safe run strategy mapping, pending backend binding, Portal preflight, snapshot preflight, and backend-term filtering.

### Changed

- Extended launch request models with a desktop-safe run plan summary so KDE launchers can show Runtime planning status without exposing backend implementation details.
- Added Runtime probe capability reporting for compatibility run planning.

## [0.2.24] - 2026-07-13

### Added

- Added `xnix-portal-access-policy`, a Runtime-owned policy model for file, URI, print, screenshot, clipboard, camera, and remote-desktop Portal access decisions.
- Added focused coverage for Portal mediation, direct-access denial, default ask or deny decisions, CLI validation, and backend-term filtering.

### Changed

- Connected the desktop integration manifest settings artifact to the Portal access policy command so KDE settings can discover Runtime-owned sensitive desktop operation policy.

## [0.2.23] - 2026-07-13

### Changed

- Enforced recipe install gate preflight in `xnix-install-desktop-integration` before staged desktop activation.
- Made the desktop activation CLI use registry-backed recipe loading and block development-only registries in production mode by default.
- Added an explicit development install mode for safe local staging after digest verification.

### Added

- Added focused coverage that blocked production activation writes no desktop files and development staging still succeeds under the verified development registry.

## [0.2.22] - 2026-07-13

### Added

- Added `xnix-recipe-install-gate`, an install gate that evaluates a recipe registry, target application id, and production or development mode before activation.
- Added focused coverage for production blocking of development-only registries, development staging allowance, verified production registries, unknown recipe blocking, and desktop-safe output.

### Changed

- Documented recipe install gates as the enforcement layer above recipe trust policy decisions.

## [0.2.21] - 2026-07-13

### Added

- Added `xnix-recipe-trust-policy`, a policy model that evaluates recipe registry trust as production-trusted, development-only, or untrusted.
- Added focused coverage for development-only blocking reasons, production-trusted decisions, and desktop-safe summaries.

### Changed

- Documented recipe trust policy as the desktop-facing explanation layer above raw recipe trust probe signals.

## [0.2.20] - 2026-07-13

### Added

- Added Runtime probe reporting for recipe trust status, including registry-backed loading, digest verification, development registry state, and signed-recipe validation status.
- Added focused coverage for recipe trust capability reporting.

### Changed

- Documented the Runtime probe as the observable trust boundary for KDE surfaces and diagnostics.

### Verified

- Completed the twentieth-version full Buildroot build and QEMU serial smoke test.

## [0.2.19] - 2026-07-13

### Added

- Added a registry-backed recipe store that verifies `registry.json` before loading Runtime-managed recipes.
- Added focused coverage for registered recipe loading, digest mismatch rejection, and no-registry development fallback.

### Changed

- Made `xnix-compatd` use registry-backed recipe loading by default when a recipe registry is present.

## [0.2.18] - 2026-07-13

### Added

- Added `xnix-recipe-registry`, a recipe registry verifier for schema version, recipe identifiers, safe relative paths, SHA-256 digests, and signature status.
- Added a development recipe registry for the bundled sample application recipe.
- Added focused coverage for digest verification, digest mismatch rejection, unsafe path rejection, and explicit reporting that production signed-recipe validation is not yet enabled.

### Changed

- Documented recipe registry verification as the next application-library trust layer before production signed recipe storage.

## [0.2.17] - 2026-07-12

### Added

- Added `xnix-rollback-desktop-integration`, a receipt-based rollback command for staged KDE desktop activation files.
- Added SHA-256 activation receipts for generated launchers, Dolphin service menus, and persisted desktop integration manifests.
- Added focused coverage for checksum-verified removal, changed-file preservation, receipt cleanup, and host-root rejection.

### Changed

- Extended staged desktop activation results with rollback receipt metadata so Runtime-managed desktop integration can be audited and reverted.

## [0.2.16] - 2026-07-12

### Added

- Added `xnix-install-desktop-integration`, a staged desktop activation installer that writes generated application launchers, Dolphin service menus, and desktop integration manifests under a target root.
- Added focused coverage for staging-root safety, installed file paths, file modes, manifest persistence, and backend-term filtering.

### Changed

- Documented desktop activation as a staging-root operation so tests can verify KDE integration output without modifying the host desktop.

## [0.2.15] - 2026-07-12

### Added

- Added `xnix-desktop-integration-manifest`, a KDE recipe activation manifest that groups launcher, task manager, Dolphin, tray, notification, Compatibility Center, and settings artifacts.
- Added focused coverage to keep the manifest ordered, portal-aware, privilege-free, and free of backend implementation terminology.

### Changed

- Documented the desktop activation path as a single Runtime-owned manifest over the seven KDE entry-point models.

## [0.2.14] - 2026-07-12

### Added

- Added `xnix-compat-window-identity`, a KDE Task Manager and KWin-facing identity model for compatibility application windows.
- Added focused coverage for desktop file mapping, grouping, pinning, restore behavior, and KWin identity-only metadata.

### Changed

- Marked the Task Manager KDE entry point as an initial integration, completing initial coverage for all seven first-release KDE entry points.

## [0.2.13] - 2026-07-12

### Added

- Added `xnix-compat-tray-status`, a KDE-facing system tray status model for Runtime activity, compatibility attention state, and bridged tray applications.
- Added focused coverage to keep tray status user-facing and free of backend implementation terminology.

### Changed

- Marked the System Tray KDE entry point as an initial integration in the seven-entry-point status model.

## [0.2.12] - 2026-07-12

### Added

- Added `xnix-compat-settings`, a KDE-facing compatibility settings model for run mode, file access, device access, network access, and snapshots.
- Added focused coverage to keep settings user-facing and free of backend implementation terminology.

### Changed

- Marked the Settings KDE entry point as an initial integration in the seven-entry-point status model.

## [0.2.11] - 2026-07-12

### Added

- Added `xnix-compat-notify`, a Runtime event notification request entry point for KDE-facing desktop notifications.
- Added a notification request model for install failures, automatic repairs, compatibility mode changes, and approval-required events.
- Added focused coverage for notification urgency, actions, event validation, and desktop-facing output safety.

### Changed

- Marked the Notifications KDE entry point as an initial integration in the seven-entry-point status model.

## [0.2.10] - 2026-07-12

### Added

- Added a KDE seven-entry-point integration status model covering launcher, task manager, file manager, system tray, notifications, Compatibility Center, and settings.
- Added `xnix-kde-integration-status` for reporting which KDE entry points have an initial Runtime-backed integration and which remain planned.
- Added focused coverage for the first-release KDE scope and entry-point evidence.

### Verified

- Completed the tenth-version full Buildroot build and QEMU serial smoke test.

## [0.2.9] - 2026-07-12

### Added

- Added `xnix-compat-launch`, the managed desktop launcher entry point used by generated application `.desktop` files.
- Added a launch request model that validates Runtime application ids, preserves optional file URIs, marks file launches as portal-mediated, and targets the Runtime `Launch` method without invoking a backend directly.
- Added focused coverage for launcher request validation and desktop-facing output safety.

### Changed

- Documented the KDE launcher path as a Runtime request model rather than a direct Wine or VM command path.

## [0.2.8] - 2026-07-12

### Added

- Added a Dolphin service menu entry for opening selected files through Xnix Compatibility without exposing backend commands.
- Added `xnix-compat-open`, a file-open request entry point that validates `file://` URIs, resolves a Runtime application recipe by file extension, and emits a portal-required Runtime launch request model.
- Added focused coverage for Dolphin service menu packaging and file-open request validation.

### Changed

- Documented the file manager integration as a Runtime request path rather than a direct Wine or VM launcher.

## [0.2.7] - 2026-07-12

### Added

- Added a thin Runtime D-Bus client for KDE-facing read models, using the session bus and the stable `org.xnix.Compatibility1` contract.
- Added `--source auto|local|dbus` to `xnix-kde-center-model`, preferring D-Bus when a Runtime session service is available and falling back to the local read model otherwise.
- Added a constrained KDE Compatibility Center D-Bus smoke path that verifies Plasma-facing model data can be read from the Runtime session bus.

### Changed

- Kept the KDE read model desktop-facing while making the Runtime source explicit as either local fallback data or D-Bus session data.

## [0.2.6] - 2026-07-12

### Added

- Added a read-only KDE Compatibility Center model that converts Runtime applications and diagnostics into a safe desktop presentation model.
- Added a `xnix-kde-center-model` CLI for Plasma-facing model smoke tests.
- Added focused coverage to prevent the KDE presentation model from exposing backend storage or implementation terms.

### Changed

- Updated the Plasma package metadata to the current project version and documented the Runtime-backed Compatibility Center model boundary.

## [0.2.5] - 2026-07-12

### Added

- Added a Linux-only D-Bus session smoke adapter that owns `org.xnix.Compatibility1`, exposes `/org/xnix/Compatibility1`, and answers read-only Runtime calls over `gdbus`.
- Added a constrained container command for the Runtime D-Bus session smoke path.
- Added focused coverage for the D-Bus smoke script, adapter source, and container command.

### Changed

- Made the container image build prefer the local base-image cache with `--pull=false` so repeated verification is less sensitive to registry metadata failures.

### Verified

- Built `xnix-builder:0.2.5` and ran the Runtime D-Bus session smoke path in the constrained no-network, read-only-root container.

## [0.2.4] - 2026-07-12

### Fixed

- Made the Runtime activation installer test compatible with the constrained container's `noexec` temporary filesystem by checking executable mode bits instead of attempting an executable-path predicate.

### Verified

- Rebuilt the constrained Docker image with D-Bus tooling and ran the Runtime activation smoke path in a no-network, read-only-root container.

## [0.2.3] - 2026-07-12

### Added

- Added a root-staging installer for Runtime activation files, covering the packaged libexec wrapper, D-Bus system service, and systemd unit.
- Added container image dependencies for future Linux D-Bus smoke tests: `dbus` and `libglib2.0-bin`.
- Added focused coverage for activation-file installation into an unprivileged temporary root.

## [0.2.2] - 2026-07-12

### Added

- Added a Runtime method-dispatch layer for read-only D-Bus contract methods: `ListApplications`, `GetApplication`, and `GetDiagnostics`.
- Added a packaged libexec wrapper that matches the D-Bus and systemd activation path.
- Added focused tests for Runtime method dispatch and activation-path consistency.

### Changed

- Kept launch, install, snapshot, and restore methods rejected until Wine/VM backends and request signaling exist.

## [0.2.1] - 2026-07-12

### Added

- Added a runnable Compatibility Runtime daemon core with JSON probe, application listing, per-application diagnostics, introspection output, and managed recipe-store loading.
- Added a bundled sample recipe and focused tests for recipe loading and runtime daemon behavior.

### Changed

- Documented that the Runtime core is executable while the real D-Bus binding, Wine backend, VM backend, and signed recipe trust policy remain pending.

## [0.2.0] - 2026-07-12

### Added

- Added the independent Compatibility Runtime foundation: validated application recipes, ordinary Linux desktop entry generation, a D-Bus API contract, a hardened systemd service definition, and a Plasma 6 Compatibility Center package skeleton.
- Added the atomic KDE Plasma product architecture, with Wine/Proton and Windows VM backends explicitly separated from the desktop shell.

### Changed

- Repositioned Buildroot/QEMU as the verified learning baseline and selected a Fedora Kinoite-compatible atomic KDE desktop as the flagship product direction.

## [0.1.16-rc1] - 2026-07-12

### Fixed

- Removed the unsupported `UsePAM` directive after the QEMU serial boot showed that Buildroot OpenSSH rejected it during `sshd` startup.

### Verified

- Completed the constrained Buildroot, QEMU serial, DHCP, and loopback key-authenticated `sshd` smoke path.

## [0.1.16] - 2026-07-12

### Added

- Added an isolated QEMU SSH smoke-test path using a disposable test key stored only in the Docker-managed cache volume and a container-loopback-only forwarding rule.

### Changed

- Updated the documented constrained runtime limit to 4 GiB inside the user-authorized 6 GiB Colima VM.

## [0.1.15-rc2] - 2026-07-12

### Fixed

- Raised the constrained build-container memory budget to 4 GiB after Colima was explicitly authorized to use up to 6 GiB, while retaining single-job compilation and a 2 GiB VM safety reserve.

## [0.1.15-rc1] - 2026-07-12

### Fixed

- Limited Buildroot to one job after the initial cross-GCC automata generator was killed by the constrained container memory limit.

## [0.1.15] - 2026-07-12

### Added

- Human-readable serial boot failure summaries for QEMU smoke-test diagnostics.

## [0.1.14] - 2026-07-12

### Added

- Named persistent full-build container command so Buildroot logs remain available after a failed build.

## [0.1.13] - 2026-07-12

### Added

- OpenSSH server hardening overlay that disables password authentication and root password login.

## [0.1.12] - 2026-07-12

### Added

- Loopback-only non-interactive SSH probe command and focused unit coverage.

## [0.1.11] - 2026-07-12

### Added

- OpenSSH `sshd` serial-service evidence verifier and focused unit coverage.

## [0.1.10-rc1] - 2026-07-12

### Fixed

- Added a constrained Buildroot dependency-download phase before the offline full build, allowing the v0.1.10 verification to obtain its declared sources without granting network access to compilation or QEMU.

## [0.1.10] - 2026-07-12

### Added

- Full-build and QEMU serial-smoke runner protected by the tenth-version milestone gate.

## [0.1.9] - 2026-07-12

### Added

- Version gate that permits full-build and smoke-test execution only at every tenth formal code version.
- Focused unit coverage for milestone and pre-release version handling.

## [0.1.8] - 2026-07-12

### Added

- Optional loopback-only QEMU SSH forwarding command and focused port-exposure unit coverage.

## [0.1.7] - 2026-07-12

### Added

- Xnix hostname configuration and serial-log boot evidence verifier.
- Focused unit coverage for successful and incomplete boot logs.

## [0.1.6] - 2026-07-12

### Added

- Restricted x86_64 QEMU TCG command generation for the future Xnix boot smoke test.
- Focused unit coverage for QEMU memory, CPU, console, and network restrictions.

## [0.1.5] - 2026-07-12

### Added

- Buildroot configuration and full-build command generation against the Docker-managed source cache.
- Focused unit coverage for the x86_64 defconfig, external tree, output directory, and offline build isolation.

## [0.1.4] - 2026-07-12

### Added

- Docker-managed Buildroot source cache volume for constrained source retrieval without a host-directory mount.
- Networked source-retrieval container command with targeted isolation unit coverage.

## [0.1.3-rc3] - 2026-07-12

### Fixed

- Removed the runtime bind mount that is unavailable when Colima host-directory mounts are disabled; project files are now copied into the tool image during its build.

## [0.1.3-rc2] - 2026-07-12

### Fixed

- Replaced the invalid `rw` field in the Docker bind-mount specification with Docker's compatible default writable bind mount.

## [0.1.3-rc1] - 2026-07-12

### Fixed

- Removed unsupported Docker Buildx CPU and memory arguments from the image-build command; resource isolation remains enforced by the 1 GiB, one-CPU Colima VM.

## [0.1.3] - 2026-07-12

### Added

- Resource-constrained Docker command builder for Xnix build and offline runtime containers.
- Unit tests that enforce container isolation and resource-limit arguments without invoking Docker.

## [0.1.2] - 2026-07-12

### Added

- Pinned Buildroot 2025.02.15 source URL and SHA-256 lock data.
- Container-only Buildroot source retrieval utility with offline lock verification.

## [0.1.1] - 2026-07-12

### Added

- Buildroot external-tree skeleton and x86_64 system configuration.
- Restricted Docker build definition with QEMU software emulation tooling.
- Focused Ruby layout verification for the initial system scaffold.

## [0.1.0] - 2026-07-12

### Added

- Initial Xnix project rules, automated-contributor guide, product overview, and ignore rules.
- Linux LTS and Buildroot primary implementation path.
- Constrained Colima and Docker test policy with focused tests for small versions and full QEMU smoke tests every tenth code version.

### Changed

- Standardized all project-facing content on English only.
