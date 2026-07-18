# Changelog

Xnix follows Semantic Versioning.

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

## [0.2.346] - 2026-07-19

### Added

- Added `production-dbus-gate-review-preview`, a read-only Go owner review packet that inventories the three smoke-covered owner-local routes and explicitly keeps production D-Bus ownership disabled.
- Added CLI, Go tests, layout, evidence, and product metadata coverage for the production D-Bus gate review while requiring human authorization, production service activation preflight, Runtime write-gate review, route-by-route D-Bus method review, desktop side-effect review, and rollback/diagnostics review before any production ownership claim.

## [0.2.345] - 2026-07-19

### Added

- Refreshed `backend-adapter-contract-owner-route-audit-preview` so the redacted adapter profile owner-route audit recognizes the v0.2.342 restricted owner smoke coverage evidence instead of still recommending smoke coverage as the next blocker.
- Added Go, CLI, layout, evidence, and product metadata coverage for the `redacted-profile-route-smoke-covered` audit state while keeping the full adapter contract fixture-local and leaving production D-Bus exposure, Runtime writes, adapter invocation, installation, download, command materialization, executable resolution, backend launch, raw-detail exposure, and host-root mutation disabled.

## [0.2.344] - 2026-07-19

### Added

- Refreshed `restricted-owner-smoke-receipt-fanout-owner-route-audit-preview` so the restricted owner smoke fan-out owner-route audit recognizes the v0.2.341 restricted owner smoke coverage evidence instead of still recommending smoke coverage as the next blocker.
- Added Go, CLI, layout, evidence, and product metadata coverage for the `owner-local-route-smoke-covered` restricted fan-out audit state while leaving production D-Bus exposure, Runtime writes, support bundle export, support case creation, notifications, backend launch, backend process start, path exposure, backend details, and host-root mutation disabled.

## [0.2.343] - 2026-07-19

### Added

- Refreshed `kde-test-launch-materialization-owner-route-audit-preview` so the materialization fan-out owner-route audit recognizes the v0.2.340 restricted owner smoke coverage evidence instead of still recommending smoke coverage as the next blocker.
- Added Go, CLI, layout, evidence, and product metadata coverage for the `owner-local-route-smoke-covered` audit state while leaving production D-Bus exposure, Runtime writes, backend launch, backend process start, path exposure, backend details, and host-root mutation disabled.

## [0.2.342] - 2026-07-19

### Added

- Added `backend-adapter-redacted-profile-owner-smoke-coverage-preview`, an owner smoke coverage read model that unwraps the restricted smoke-batch `GetBackendAdapterProfileAudit` service-call record and proves the v0.2.334 redacted adapter profile owner-local route is exercised through `Service.Call`.
- Added owner CLI, Go tests, layout, evidence, and product metadata coverage while keeping the full adapter contract fixture-local, preserving redacted user-safe profile output, and leaving production D-Bus exposure, Runtime writes, adapter invocation, install, download, command materialization, executable resolution, backend launch, raw-detail exposure, and host-root mutation disabled.

## [0.2.341] - 2026-07-19

### Added

- Added `restricted-owner-smoke-receipt-fanout-owner-smoke-coverage-preview`, an owner smoke coverage read model that unwraps the restricted smoke-batch `GetRestrictedOwnerSmokeReceiptFanOut` service-call record and proves the v0.2.337 owner-local restricted smoke receipt fan-out route is exercised through `Service.Call`.
- Added owner CLI, Go tests, layout, evidence, and product metadata coverage while keeping missing restricted owner smoke receipts fail-closed and leaving production D-Bus exposure, Runtime writes, support bundle export, support case creation, notifications, backend launch, backend process start, path exposure, and host-root mutation disabled.

## [0.2.340] - 2026-07-19

### Added

- Added `kde-test-launch-materialization-fanout-owner-smoke-coverage-preview`, an owner smoke coverage read model that unwraps the restricted smoke-batch `GetKDETestLaunchMaterializationFanOut` service-call record and proves the v0.2.339 owner-local materialization fan-out route is exercised through `Service.Call`.
- Added owner CLI, Go tests, layout, evidence, and product metadata coverage while keeping missing materialization receipts fail-closed and leaving production D-Bus exposure, Runtime writes, command materialization, executable resolution, backend launch, desktop side effects, path exposure, and host-root mutation disabled.

## [0.2.339] - 2026-07-19

### Added

- Added `kde-test-launch-materialization-fanout-owner-route-preview` and owner-local `GetKDETestLaunchMaterializationFanOut`, routing materialization fan-out through the opaque materialization receipt lookup without accepting caller registry, application, authorization, or state-root inputs.
- Updated materialization owner-route audit, owner dispatch, service-call, smoke-batch, private session-bus smoke, contract-drift, layout, and evidence coverage while keeping missing receipts fail-closed and leaving production D-Bus exposure, Runtime writes, command materialization, executable resolution, backend launch, desktop side effects, path exposure, and host-root mutation disabled.

## [0.2.338] - 2026-07-19

### Added

- Added `kde-test-launch-materialization-receipt-lookup-preview` and owner-local `GetKDETestLaunchMaterializationReceiptLookupPreview`, resolving the stable materialization opaque receipt id without accepting caller registry, application, state-root, or authorization inputs.
- Updated materialization owner-route audit, owner dispatch, service-call, smoke-batch, contract-drift, layout, and evidence coverage while keeping missing receipts fail-closed and leaving materialization fan-out owner routing, production D-Bus exposure, Runtime writes, launch, backend process start, path exposure, and host-root mutation disabled.

## [0.2.337] - 2026-07-19

### Added

- Added `restricted-owner-smoke-receipt-fanout-owner-route-preview` and owner-local `GetRestrictedOwnerSmokeReceiptFanOut`, routing restricted owner smoke fan-out through the opaque receipt lookup without accepting caller state-root paths.
- Updated owner dispatch, service-call, smoke-batch, private session-bus smoke, route checkpoint, contract-drift, layout, and evidence coverage while keeping missing receipts fail-closed and leaving production D-Bus exposure, Runtime writes, support side effects, launch, backend process start, path exposure, and host-root mutation disabled.

## [0.2.336] - 2026-07-19

### Added

- Updated `kde-test-launch-materialization-owner-route-audit-preview` to recognize the existing read-only materialization fan-out consumption split and to identify owner-managed opaque materialization receipt lookup as the remaining route blocker.
- Added Go, CLI, layout, and evidence coverage for the updated materialization route audit while keeping fan-out owner routing, production D-Bus exposure, Runtime writes, launch, backend process start, path exposure, and host-root mutation disabled.

## [0.2.335] - 2026-07-19

### Added

- Added `restricted-owner-smoke-receipt-lookup-preview`, a Go Runtime owner-managed opaque receipt lookup that resolves `restricted-owner-smoke-receipt-id` to a safe relative receipt slot without accepting caller state-root paths.
- Added owner-local read dispatch, service-call, smoke-batch, session-bus smoke, CLI, layout, drift, and evidence coverage for `GetRestrictedOwnerSmokeReceiptLookupPreview`; the existing fan-out remains CLI-only until it consumes lookup results, and production D-Bus exposure, Runtime writes, support side effects, backend launch, path exposure, and host-root mutation remain disabled.

## [0.2.334] - 2026-07-19

### Added

- Added a redacted backend adapter profile audit through `backend-adapter-redacted-profile-audit-preview`, exposing only user-safe compatibility profile states while omitting internal adapter identifiers and profile paths.
- Added owner-local read dispatch, service-call, smoke-batch, session-bus smoke, CLI, layout, and evidence coverage for `GetBackendAdapterProfileAudit` while keeping the full adapter contract fixture-local and leaving production D-Bus exposure, Runtime writes, adapter invocation, backend launch, command materialization, path exposure, networking, and host-root mutation disabled.

## [0.2.333] - 2026-07-19

### Added

- Added read-only KDE test launch materialization fan-out receipt consumption through `kde-test-launch-materialization-fanout-consume-preview`, allowing an existing materialization plan receipt to be projected to KDE surfaces without creating a new receipt.
- Added Go, CLI, layout, and evidence coverage that keeps the legacy test-only receipt creation path separate from the read-only consumption path while leaving fan-out writes, Runtime writes, command materialization, executable resolution, backend launch, process start, path exposure, and host-root mutation disabled.

## [0.2.332] - 2026-07-19

### Added

- Added the restricted owner smoke receipt fan-out owner-route audit preview and CLI command, deciding that the v0.2.329 readiness/support fan-out should remain CLI-only until owner-managed opaque receipt lookup exists.
- Added owner, CLI, layout, and evidence coverage that keeps caller state-root paths out of owner routing and keeps production D-Bus exposure, Runtime writes, support export/case/notification side effects, backend launch, path exposure, backend details, and host-root mutation disabled.

## [0.2.331] - 2026-07-18

### Added

- Added the backend adapter contract owner-route audit preview and CLI command, deciding that the v0.2.328 fixture adapter audit should remain fixture-local until a redacted owner-local adapter profile audit route exists.
- Added Go, CLI, layout, and evidence coverage that keeps owner dispatch, production D-Bus exposure, adapter invocation, install/download, backend launch, command materialization, path exposure, backend details, and host-root mutation disabled.

## [0.2.330] - 2026-07-18

### Added

- Added the KDE test launch materialization owner-route audit preview and CLI command, deciding that the v0.2.327 materialization fan-out must remain CLI-only until it can consume owner-managed opaque receipt evidence.
- Added Go, CLI, layout, and evidence coverage that keeps owner dispatch, production D-Bus exposure, service start, Runtime writes, backend launch, network access, state-root path exposure, backend details, and host-root mutation disabled.

## [0.2.329] - 2026-07-18

### Added

- Added the restricted owner smoke receipt fan-out preview and CLI command, letting Runtime readiness, service activation preflight, onboarding, support bundle, and support case surfaces consume the v0.2.326 receipt.
- Added owner, CLI, layout, and evidence coverage that keeps fan-out writes, service start, session-bus claim, production D-Bus ownership, Runtime write methods, support bundle export, support case creation, notifications, backend launch, network access, state-root path exposure, backend details, and host-root mutation disabled.

## [0.2.328] - 2026-07-18

### Added

- Added no-op backend adapter contract audits to the offline application fixture matrix, letting fixture-backed readiness rows consume KDE-safe adapter profiles from `backend-adapter-contract-preview`.
- Added Go, CLI, Ruby report, layout, and evidence coverage that keeps adapter invocation, install, download, command materialization, executable resolution, backend launch, backend process start, network access, state-root path exposure, backend details, and host-root mutation disabled.

## [0.2.327] - 2026-07-18

### Added

- Added the KDE test launch materialization fan-out preview and CLI command, projecting the test-only materialization receipt into Compatibility Center, task manager, tray, and notification read models.
- Added fan-out coverage and evidence guards that keep fan-out writes, task-manager activation, live tray bridging, notification delivery, Compatibility Center actions, Runtime writes, command materialization, executable resolution, backend launch, backend process start, network access, privileged containers, state-root path exposure, and host-root mutation disabled.

## [0.2.326] - 2026-07-18

### Added

- Added the restricted owner smoke execution receipt writer and CLI command, consuming Runtime service activation preflight and owner smoke-batch evidence into a digest-verified state-root receipt.
- Added owner and CLI coverage that keeps production service start, session-bus claim, production D-Bus ownership, Runtime write methods, backend launch, backend details, network access, privileged containers, state-root path exposure, and host-root mutation disabled.

## [0.2.325] - 2026-07-18

### Added

- Added the Go Runtime backend adapter no-op contract preview and CLI command, defining Wine, Proton, and Windows VM adapter boundaries as Runtime-internal no-op contracts before any real backend invocation is allowed.
- Added adapter contract tests and evidence hooks that keep invocation, install, download, launch, process start, command materialization, executable path resolution, raw command exposure, state-root path exposure, backend details in KDE-facing output, network access, privileged containers, and host-root mutation disabled.

## [0.2.324] - 2026-07-18

### Added

- Added a test-only launch materialization plan receipt under the controlled execution state root, joining restricted launch authorization, restricted preflight, and blocked execution transaction evidence.
- Added `kde-test-launch-materialization-record` to materialize only review-plan references while keeping command materialization, executable path resolution, backend selection, backend launch, process start, Runtime writes, network access, privileged containers, raw command exposure, and host-root mutation disabled.

## [0.2.323] - 2026-07-18

### Added

- Added the Go Runtime service activation preflight preview and CLI command to aggregate service binding, owner readiness, and owner smoke evidence before any production activation attempt.
- Added fail-closed production activation gates that allow only restricted owner smoke readiness while keeping service start, production bus claim, Runtime writes, backend launch, KDE ownership, network requirements, privileged containers, and host-root mutation disabled.

## [0.2.322] - 2026-07-18

### Added

- Added read-only snapshot baseline receipt evidence to the offline application fixture matrix, allowing verified controlled snapshot baselines to close the snapshot evidence gap without creating, restoring, deleting, or pruning snapshots.
- Added `--snapshot-state-root` to the Go Runtime fixture matrix preview and Ruby report wrapper so existing snapshot stores can be reviewed without exposing state roots, reading user file contents into KDE-facing output, launching backends, or mutating the host.

## [0.2.321] - 2026-07-18

### Added

- Added read-only artifact stage receipt evidence to the offline application fixture matrix, allowing valid local fixture receipts to close the artifact evidence gap without downloads, package-manager calls, backend launch, root-path exposure, or host-root mutation.
- Added `--artifact-receipt-root` to the Go Runtime fixture matrix preview and Ruby report wrapper so controlled receipt roots can be reviewed without staging artifacts from the matrix itself.

## [0.2.320-rc8] - 2026-07-18

### Changed

- Set the post-`v0.2.320` train cadence to targeted tests per small version and one full build, complete test suite, product smoke, and remote push every thirty small versions, with the next gate at `v0.2.350`.

### Fixed

- Separated product-image readiness from production Runtime activation after the full Go gate exposed stale assumptions about the removed development service wrapper.
- Kept Fedora Kinoite and KDE graphical-login smoke readiness valid without claiming a production D-Bus owner, Wine/Proton backend launch, or Windows application execution.
- Added structured q4 evidence for the authorized container build, clean qcow2 validation, persisted serial log, KVM graphical-login pass, and closed host boundary.
- Replaced stale hardcoded project versions in Ruby model tests with the canonical `VERSION` file and added a layout guard against future drift.

## [0.2.320-rc7] - 2026-07-18

### Fixed

- Corrected Fedora shell OSC stripping so a successful active KDE boot marker adjacent to a shell integration sequence remains visible to the smoke verifier.
- Added regression coverage using the exact control-sequence shape captured during the rc6 q4 KVM run, where both `graphical.target` and `plasmalogin.service` were active but the verifier timed out.

## [0.2.320-rc6] - 2026-07-18

### Fixed

- Replaced passive graphical-target serial matching with an authenticated serial probe that queries `graphical.target` and `plasmalogin.service`, emits deterministic pass/fail markers, and terminates QEMU after a result.
- Added isolated Podman graphroot, runroot, `slirp4netns`, repository host pin, and proxy-clearing options to the container build driver while rejecting host networking.
- Recorded the successful rc5 container and qcow2 builds, clean `qemu-img` validation, active KDE graphical target and Plasma Login Manager, and empty failed-unit set observed under KVM on q4.

## [0.2.320-rc5] - 2026-07-18

### Changed

- Aligned the Fedora 44 image with Plasma Login Manager and removed obsolete SDDM packages and configuration from the flagship build.
- Stopped installing and enabling the development `xnix-compatd` CLI wrapper as a production system D-Bus service; production ownership remains gated until a real owner exists.
- Recorded q4 evidence for the successful rc4 container compose, clean qcow2 validation, active graphical target, and active Plasma Login Manager while keeping Windows application execution unproven.

### Fixed

- Made KDE boot smoke checks ignore ANSI serial control sequences, prefer KVM when available, protect the source disk with QEMU snapshot mode, persist output incrementally, and allow five minutes for first boot.
- Added explicit custom Podman graphroot and runroot support for bootc-image-builder, including the static graphroot and compatibility storage mounts required on q4.

## [0.2.320-rc4] - 2026-07-18

### Changed

- Advanced the flagship atomic desktop base from Fedora Kinoite 41 to Fedora Kinoite 44 after q4 verification confirmed that the official amd64 image is available.
- Updated the checked-in Containerfile snapshot and KDE image test to pin the supported Fedora 44 base.

### Security

- Rejected Fedora 41 as a release base after its repositories moved to distant archives and its update lifecycle ended; qcow2 and KDE smoke evidence must now come from the supported Fedora 44 compose.

## [0.2.320-rc3] - 2026-07-17

### Fixed

- Fully qualified the bootc source as `localhost/xnix-kinoite:<version>` so the privileged builder resolves the Podman-local image through the mounted container store instead of `docker.io/library`.
- Added disk-command coverage that keeps local image resolution explicit.

### Changed

- Recorded the rc2 q3 manifest-generation result while keeping qcow2 generation, KDE startup, and Windows application execution claims pending.

## [0.2.320-rc2] - 2026-07-17

### Fixed

- Added an explicit, validated Btrfs root filesystem to Fedora Kinoite disk builds after bootc-image-builder rejected the otherwise successful q3 container compose for missing `DefaultRootFs` metadata.
- Added disk-model coverage that rejects unsupported root filesystems and requires `--rootfs btrfs` in the privileged builder command.

### Changed

- Recorded the successful q3 container compose while keeping qcow2 generation, KDE startup, and Windows application execution claims pending.

## [0.2.320-rc1] - 2026-07-17

### Fixed

- Removed the obsolete Fedora 41 `plasma-workspace-wayland` and `kwin-wayland` subpackages from the Kinoite layering transaction after the authorized q3 compose exposed rpm-ostree dependency conflicts.
- Added a KDE image regression guard that rejects both obsolete package requests while preserving the Plasma Wayland and KWin desktop policy.

### Changed

- Updated the release evidence notes to record the real q3 compose result and keep bootc disk, KDE startup, and Windows application execution claims pending.

## [0.2.320] - 2026-07-17

### Changed

- Advanced the first Go stability train to its full test, build, smoke, and release-evidence gate.
- Completed the six-step Buildroot build and constrained QEMU serial smoke with persisted v0.2.320 reports, successful boot markers, loopback-restricted networking, and no host-root mutation.
- Kept remote push blocked unless the complete gate produces persisted evidence without weakening host safety.

### Fixed

- Excluded the local `.gocache/` directory from Docker build contexts after the full gate exposed a one-gigabyte context stall.
- Added container and layout assertions that keep both Buildroot and Go caches outside future builder image contexts.
- Classified root Docker build-context safety changes in the CW11 product-image acceptance review lane.

### Security

- Refused the flagship Kinoite compose and disk-build path on the current host because the required Podman/Buildah toolchain is absent and bootc image production requires a privileged build host.

## [0.2.319] - 2026-07-17

### Added

- Added `kde-restricted-product-smoke-checkpoint-record` and the Go `KDERestrictedProductSmokeCheckpointRecord` model.
- Joined the restricted launch preflight packet to the fixed Kinoite product-image manifest and five restricted smoke evidence groups.
- Added eight checkpoint checks plus CLI authorization, path-redaction, blocked-state, and no-execution coverage.

### Changed

- Marked product-image metadata and the review packet ready for the v0.2.320 train gate while keeping product smoke separately human-authorized.
- Kept Docker, QEMU, product smoke, serial-log claims, release readiness, launch preflight, launch authorization, command materialization, backend selection and launch, compatibility process start, production bus ownership, host networking, Docker socket mounts, broad host mounts, privileged containers, path exposure, backend details, and host-root mutation disabled.

## [0.2.318] - 2026-07-17

### Added

- Added the digest-verified Go `RestrictedPreflightStore` for fail-closed restricted launch preflight packets under a controlled test state root.
- Added `kde-restricted-launch-preflight-record` and `KDERestrictedLaunchPreflightRecord` to consume the v0.2.317 preparation authorization and join ten core receipts.
- Added packet persistence, readback, tamper rejection, exact-blocker, explicit-authorization, redaction, and no-side-effect tests.

### Changed

- Advanced restricted test preparation to evidence-packet assembly while naming exactly `recipe-trust` and `runtime-write-gate` as the remaining launch blockers.
- Kept product-image readiness, production trust, Runtime writes, launch preflight, launch authorization, execution approval, process-start authorization, command materialization, executable-path resolution, launch-backend selection, backend launch, compatibility process start, production bus ownership, network access, privileged containers, backend details, and host-root mutation disabled.

## [0.2.317] - 2026-07-17

### Added

- Added the Go `RestrictedAuthorizationStore` with an exact `test-only` preparation scope and `authorize-restricted-test-preparation` directive.
- Added `kde-restricted-launch-authorization-record` and `KDERestrictedLaunchAuthorizationRecord` to join the explicit authorization receipt to the Sample Notepad blocked execution and session records.
- Added digest-verified authorization readback, eight boundary checks, invalid-directive, tamper, and managed-directory symlink tests.

### Changed

- Authorized only restricted test preparation; launch authorization, process-start authorization, and execution approval remain separate and false.
- Proved that recording preparation authorization does not satisfy production recipe trust, enable the Runtime write gate, or change the blocked execution/session state.
- Kept artifact acquisition, backend install and launch, compatibility process start, real Portal transport, production bus ownership, network access, privileged containers, raw commands, backend details, and host-root mutation disabled.

## [0.2.316] - 2026-07-17

### Added

- Added `kde-backend-lifecycle-evidence-record` and the Go `KDEBackendLifecycleEvidenceRecord` model for the Sample Notepad controlled state root.
- Added digest-verified backend-manager inventory readback and eight checks joining prerequisite, inventory, lifecycle, execution, session, KDE fan-out, backend ownership, and closed unsafe gates.
- Added backend-manager schema, fixed relative-path, count, SHA-256, disabled-gate, tamper, and managed-directory symlink validation.

### Changed

- Joined three internal managed-backend entries and three user-facing compatibility profiles to the ready lifecycle while keeping backend kinds hidden from KDE-facing output.
- Kept the execution and session receipts blocked by development recipe trust and the Runtime write gate even though environment, snapshot, and Portal evidence pass.
- Kept backend install, download, launch, process and VM start, raw commands, profile paths, real Portal transport, restore, live diagnostics, AI calls, repair, production bus ownership, network access, privileged containers, secrets, backend details, and host-root mutation disabled.

## [0.2.315] - 2026-07-17

### Added

- Added `kde-snapshot-diagnostics-evidence-record` and the Go `KDESnapshotDiagnosticsEvidenceRecord` model for the Sample Notepad controlled state root.
- Added one deterministic redacted diagnostic receipt, one content-addressed snapshot baseline, digest-verified readback, and eight convergence checks.
- Added diagnostic receipt schema, identity, relative-path, digest, and disabled-safety-gate validation with tamper rejection coverage.

### Changed

- Advanced the backend lifecycle from `staged` to `ready-with-launch-disabled` after Portal and snapshot evidence gates are satisfied.
- Recomputed the blocked execution and session receipts so environment, snapshot, and Portal gates pass while development recipe trust and the Runtime write gate remain closed.
- Kept snapshot restore and deletion, live diagnostic execution, AI provider calls, repair, real Portal transport, launch, compatibility process start, production bus ownership, network access, privileged containers, backend details, file contents, and host-root mutation disabled.

## [0.2.314] - 2026-07-17

### Added

- Added `kde-fake-portal-evidence-record` and the Go `KDEFakePortalEvidenceRecord` model for the Sample Notepad controlled state root.
- Added deterministic fake Portal create, grant, complete, readback, content-digest, execution-ledger, session-record, and four-surface KDE fan-out evidence.
- Added seven checks proving the blocked no-receipt baseline, completed receipt readback, Portal-only gate delta, lifecycle join, execution join, session join, and closed unsafe gates.

### Changed

- Advanced the staged lifecycle by satisfying only `portal-policy-review`; `snapshot-baseline` remains pending and the execution environment remains not ready.
- Made repeated Portal evidence recording reuse the completed fake receipt instead of adding duplicate permission records.
- Kept real Portal transport, host permission changes, execution approval, launch, execution, compatibility process start, production bus ownership, network access, privileged containers, backend details, and host-root mutation disabled.

## [0.2.313] - 2026-07-17

### Added

- Added `kde-fake-execution-evidence-record` and the Go `KDEFakeExecutionEvidenceRecord` orchestration model for the digest-verified `org.xnix.sample.notepad` fixture.
- Added controlled test-root persistence and readback for one staged lifecycle receipt, one blocked execution transaction, and one blocked session status record.
- Added six evidence checks for recipe identity, lifecycle readiness, transaction readback, session readback, four-surface KDE fan-out, and closed unsafe gates.

### Changed

- Started the `0.2.313`-`0.2.316` evidence-join band with an explicit `--mode test-only` write boundary and relative receipt paths that do not expose the configured state root.
- Preserved development-signature, staged lifecycle, missing Portal receipt, missing snapshot baseline, and Runtime write-gate blocked reasons for the next train slices.
- Kept real Portal calls, launch, execution, compatibility process start, production bus ownership, network access, privileged containers, backend detail exposure, and host-root mutation disabled.

## [0.2.312] - 2026-07-17

### Added

- Added `xnix-runtime-owner --kde-identity-checkpoint` and the Go `KDEOfflineIdentityCheckpoint` model to close the `0.2.309`-`0.2.312` identity band.
- Added one checkpoint for digest-verified recipe trust, exact nine-surface coverage, cross-surface identity parity, owner-local routing, in-process `Service.Call`, 61/66/5/66/4 route counts, and deterministic write denials.
- Added exact-count, six-check pass state, development-signature distinction, path-redaction, CLI, persistence, delivery, launch, execution, backend-process, production-owner, write-gate, network, privileged-container, and host-mutation tests.

### Changed

- Closed the offline KDE application identity band with `org.xnix.sample.notepad` behaving as one normal application identity across desktop entry, MIME, KRunner, task manager, KWin, tray, notifications, settings, and Compatibility Center evidence.
- Kept offline identity readiness separate from production signature readiness and production execution readiness.
- Kept desktop writes, MIME changes, KDE persistence, notification delivery, launch, execution, backend process start, production bus ownership, Runtime writes, network access, and host-root mutation disabled.

## [0.2.311] - 2026-07-17

### Added

- Extended the digest-verified offline KDE identity with unified-settings and Compatibility Center application-page evidence.
- Added the owner-local `GetKDEOfflineApplicationIdentityPreview` route across Go owner dispatch, `Service.Call`, owner CLI, smoke-batch, and private session-bus transcript coverage.
- Added settings/center identity parity, caller-path rejection, owner dispatch, service call, owner CLI, persistence, action, launch, backend-process, network, and host-mutation tests.

### Changed

- Increased the canonical identity from seven to nine KDE surfaces and owner smoke coverage from 65 to 66 read routes while keeping the formal production D-Bus ABI at 61 methods.
- Increased owner-local review routes from four to five and kept the owner route checkpoint exact-count audit current.
- Kept settings persistence, Compatibility Center persistence and actions, notification delivery, live tray bridging, production bus ownership, launch, backend process start, network access, and host-root mutation disabled.

## [0.2.310] - 2026-07-17

### Added

- Extended the verified offline KDE application identity with tray and notification-center surfaces for `org.xnix.sample.notepad`.
- Added deterministic notification namespace, event, category, action, and application-identity evidence plus tray registration and navigation evidence.
- Added tray/notification parity, deterministic notification id, disabled delivery, disabled action execution, disabled live bridge, and disabled persistence tests.

### Changed

- Increased the canonical offline identity from five to seven KDE surfaces while preserving one application id, display name, icon, desktop file, launcher URL, and grouping identity.
- Updated implementation evidence, layout verification, product documentation, and dispatch guidance for attention-surface identity.
- Kept live tray bridging, tray configuration persistence, notification delivery, notification action execution, launch, backend process start, network access, and host-root mutation disabled.

## [0.2.309] - 2026-07-17

### Added

- Added a Go `kde-offline-application-identity-preview` model and CLI for the digest-verified `org.xnix.sample.notepad` registry fixture.
- Added actual desktop-entry and MIME rendering evidence plus KRunner, task-manager, and KWin identity joins with content digests and cross-surface consistency checks.
- Added verified-registry, malformed-argument, identity-parity, path-redaction, raw-command, write, execution, network, backend-process, and host-mutation tests.

### Changed

- Started the `0.2.309`-`0.2.312` offline KDE application identity band with one canonical identity across five KDE surfaces.
- Updated implementation evidence, layout verification, product documentation, and dispatch guidance for the offline identity spine.
- Kept desktop writes, MIME default changes, KRunner index persistence, task-manager activation, KWin rule application, launch, execution, backend process start, network access, and host-root mutation disabled.

## [0.2.308] - 2026-07-17

### Added

- Added `xnix-runtime-owner --route-checkpoint` and the Go `RouteCheckpoint` model for the `0.2.305`-`0.2.308` owner route migration band.
- Added a full smoke-batch audit of 61 formal Go owner routes, four owner-local review routes, 65 read service calls, four deterministic write denials, and formal method parity.
- Added checkpoint schema, exact-count, pass-state, CLI, production-owner, service-start, Runtime-write, network, privileged-container, backend-detail, and host-mutation tests.

### Changed

- Closed the owner route migration band without adding a self-referential checkpoint read route or changing the 61-method production D-Bus ABI.
- Updated implementation evidence, layout verification, product documentation, and dispatch guidance for the owner route checkpoint.
- Kept production D-Bus ownership, system service start, Runtime writes, backend launch, Docker, QEMU, network access, privileged containers, and host-root mutation disabled.

## [0.2.307] - 2026-07-17

### Added

- Added the owner-local `GetRestrictedProductSmokePacketPreview` route across Go owner dispatch, `Service.Call`, owner CLI, smoke-batch, and restricted private session-bus coverage.
- Added no-argument enforcement so owner callers cannot supply repository or manifest paths to the restricted product smoke packet.
- Added packet-prepared, human-authorization, Docker, QEMU, product-smoke, serial-log, release-readiness, backend-launch, production-owner, and host-mutation gate tests.

### Changed

- Increased restricted owner smoke coverage from 64 to 65 read routes while keeping the formal production D-Bus read contract unchanged at 61 methods.
- Updated Runtime contract drift, implementation evidence, layout verification, product documentation, and dispatch guidance for the owner-managed restricted smoke packet.
- Kept execution authorization, Docker, QEMU, product smoke, release readiness, Runtime writes, production bus ownership, backend launch, and host-root mutation disabled.

## [0.2.306] - 2026-07-17

### Added

- Added `NewRegistrySignedRecipeVerificationPreview`, which projects verified Runtime recipe-store trust into owner-safe signature evidence without accepting paths or key material from callers.
- Added the owner-local `GetSignedRecipeVerificationPreview` route across Go owner dispatch, `Service.Call`, owner CLI, smoke-batch, and restricted private session-bus coverage.
- Added development-only, signed-without-production-key, malformed-argument, path-redaction, signature-material, private-key, network, write, backend-launch, and host-mutation gate tests.

### Changed

- Increased restricted owner smoke coverage from 63 to 64 read routes while keeping the formal production D-Bus read contract unchanged at 61 methods.
- Updated Runtime contract drift, implementation evidence, layout verification, product documentation, and dispatch guidance for owner-managed signed recipe verification.
- Kept production key configuration, production trust, private-key loading, recipe writes, Runtime writes, production bus ownership, backend launch, Docker, QEMU, and host-root mutation disabled.

## [0.2.305] - 2026-07-17

### Added

- Added the owner-local `GetKDENotificationDigestPreview` route to the Go Runtime owner dispatch and in-process `Service.Call` boundary.
- Added owner dispatch, service-call, owner CLI, smoke-batch, restricted private session-bus, malformed-event, and disabled-side-effect evidence for the notification digest route.

### Changed

- Increased restricted owner smoke coverage from 62 to 63 read routes while keeping the formal production D-Bus read contract unchanged at 61 methods.
- Updated Runtime contract drift, implementation evidence, layout verification, product documentation, and dispatch guidance for the new owner-local route.
- Kept notification delivery, Runtime writes, production bus ownership, backend launch, Docker, QEMU, network access, privileged containers, and host-root mutation disabled.

## [0.2.304] - 2026-07-17

### Changed

- Added a shared Go test-version reader and migrated Runtime app identity, Runtime CLI, owner, and owner CLI assertions to the canonical repository `VERSION` file.
- Updated the full smoke report test to use the canonical version, added layout enforcement that rejects hardcoded Xnix train versions in Go tests, and classified the shared test-version files in mainline review.
- Closed the first `0.2.301`-`0.2.304` stabilization review band and advanced the current dispatch recommendation to Runtime owner route migration at `0.2.305`.
- Kept production D-Bus ownership, Runtime writes, backend launch, Docker, QEMU, network access, privileged containers, and host-root mutation disabled.

## [0.2.303] - 2026-07-17

### Added

- Added the Go `restricted-product-smoke-packet-preview` model and CLI command. The packet joins Runtime owner, signed recipe and artifact trust, backend lifecycle, fake-mode Portal safety, KDE seven-entry-point, and image-manifest evidence into a dry-run readiness decision.
- Added the offline `scripts/restricted_product_smoke_packet.rb` JSON/Markdown wrapper plus ready, missing-evidence, manifest-escape, CLI, report-format, authorization-gate, loopback-networking, serial-log-requirement, and no-execution tests.

### Changed

- Updated layout, implementation, release, product, and dispatch evidence so restricted smoke packet preparation is recognized and the next train step closes the `0.2.301`-`0.2.304` review band.
- Kept Docker execution, QEMU execution, product smoke execution, live backend launch, privileged containers, Docker socket mounts, host networking, broad host mounts, serial-log claims, release readiness, and host-root mutation disabled.

## [0.2.302] - 2026-07-17

### Added

- Added a Go `recipe.SignedVerifier` boundary that cryptographically verifies Ed25519 signed recipe metadata against the recipe digest and application identity, with fail-closed states for missing metadata, digest drift, identity mismatch, invalid signatures, invalid public keys, unsupported schemas, and unsupported algorithms.
- Added the offline `signed-recipe-verifier-preview` CLI evidence command plus valid-fixture, tampered-signature, digest-mismatch, identity-mismatch, invalid-key, schema, missing-metadata, unsigned, argument-validation, path-redaction, and no-side-effect tests.

### Changed

- Updated layout, implementation, release, product, and dispatch evidence so signed recipe verifier coverage is recognized and the next train step advances to restricted product smoke packet preparation.
- Kept production key configuration, production trust readiness, private-key loading, network access, package-manager calls, recipe writes, registry migration, backend launch, path exposure, signature-material exposure, and host-root mutation disabled.

## [0.2.301] - 2026-07-17

### Added

- Added the Go Runtime `kde-notification-digest-preview` read model and CLI command for F5W7. The digest covers needs-review, blocked-action, permission-attention, diagnostic-issue, snapshot-warning, and readiness-change groups with deterministic deduplication keys, severity, user-safe labels, and next safe read-only routes.
- Added digest-group, deduplication, malformed-input, blocked-route, redaction, CLI, notification-disabled, live-tray-disabled, request-disabled, permission-disabled, backend-disabled, and no-host-mutation tests.

### Changed

- Updated layout, implementation, release, product, and dispatch evidence so KDE notification digest coverage is recognized and the next train step advances to signed recipe verifier evidence.
- Kept real notification delivery, live tray bridging, request creation, permission grants, backend launch, raw command exposure, state-root path exposure, and host-root mutation disabled.

## [0.2.300] - 2026-07-17

### Added

- Added the offline Ruby `scripts/merge_readiness_packet.rb` report command for S6W8. The packet joins layout verification, implementation evidence, Runtime contract drift, KDE-first presence smoke, mainline integration review, release evidence indexing, and the offline application fixture matrix into one reviewer-facing JSON/Markdown artifact with tool statuses, command strings, lane classification, protected-file status, unsafe-operation status, changed-file counts, required follow-up commands, merge blockers, and release blockers.
- Added pass, skipped, missing-command, malformed-JSON, protected-file, unclassified-file, unsafe-operation, JSON-output, and Markdown-output tests for the merge readiness packet.

### Changed

- Updated the layout verifier, implementation evidence report, mainline integration review classifier, release evidence index, product documentation, and current Claude Code dispatch picks so S6W8 is recognized as implemented and the next recommended Claude handoff starts with KDE notification digest evidence.
- Added the Go-first stabilization release train for `v0.2.301` through `v0.2.320`, including targeted-test cadence, local-commit cadence, twentieth-version full build gates, and twentieth-version-only remote pushes.
- Kept staging, committing, tagging, pushing, Docker, QEMU, network fetch, package-manager calls, backend launch, and host-root mutation disabled by default in the merge readiness flow.

## [0.2.299] - 2026-07-17

### Added

- Added the Go Runtime `offline-application-fixture-matrix-preview` read model and CLI command (S6W7). The matrix covers document editor, game, installer, launcher, network-heavy app, tray-heavy app, and unsupported app shapes using built-in offline fixtures, and records recipe trust, artifact readiness, compatibility profile mapping, Portal needs, snapshot readiness, diagnostic readiness, KDE journey coverage, missing evidence, and blocked unsafe actions. Network fetch, package-manager calls, artifact staging, backend launch, Docker, QEMU, request creation, settings persistence, file-content reads, path exposure, raw command exposure, backend detail exposure, privileged containers, and host-root mutation all stay disabled.
- Added the offline Ruby `scripts/offline_application_fixture_matrix.rb` JSON/Markdown report wrapper plus matrix-shape, filtered-shape, unsupported-shape, unknown-shape, Markdown-output, JSON-output, and no-side-effect tests.

### Changed

- Updated the layout verifier, implementation evidence report, mainline integration review classifier, release evidence index, product documentation, and current Claude Code dispatch picks so S6W7 is recognized as implemented and the next recommended Claude handoff starts with S6W8 merge readiness packets.

## [0.2.298] - 2026-07-17

### Added

- Added the Go Runtime `settings-profile-migration-preview` read model and CLI command (S6W6). The preview explains how user-facing compatibility settings move from older, current, or future schema versions into the supported settings schema, including run mode, priority, document access, downloads access, camera access, network access, snapshots, diagnostics privacy, defaulting rules, blocked setting ids, user-review requirements, and rollback notes. Settings persistence, resource grants, real Portal calls, backend launch, file-content reads, backend detail exposure, state-root path exposure, and host-root mutation all stay disabled.
- Added current-schema, old-schema, future-schema, blocked-setting, invalid-setting, review-required, rollback-note, CLI, argument-validation, redaction, and no-side-effect tests for the settings profile migration preview.

### Changed

- Updated the layout verifier, implementation evidence report, mainline integration review classifier, release evidence index, product documentation, and current Claude Code dispatch picks so S6W6 is recognized as implemented and the next recommended Claude handoff starts with S6W7 offline application fixture matrices.

## [0.2.297] - 2026-07-17

### Added

- Added the Go Runtime `portal-permission-renewal-preview` read model and CLI command (S6W4). The preview explains files, URI, print, clipboard, screen, camera, remote desktop, and network permission states as current, needs review, expiring soon, denied, revoked, missing receipt, or blocked by policy with KDE-safe review-only action labels. Real Portal transport, permission grants, permission revocation, receipt writes, settings persistence, execution approval, file-content reads, path exposure, raw command exposure, backend detail exposure, and host-root mutation all stay disabled.
- Added current, expiring, denied, revoked, missing-receipt, policy-blocked, malformed-ledger, missing-state-root-read-only, CLI, argument-validation, unsafe-input, and no-side-effect tests for the Portal permission renewal preview.

### Changed

- Updated the layout verifier, implementation evidence report, mainline integration review classifier, release evidence index, product documentation, and current Claude Code dispatch picks so S6W4 is recognized as implemented and the next recommended Claude handoff starts with S6W6 settings profile migration previews.

## [0.2.296] - 2026-07-17

### Added

- Added the Go Runtime `support-case-timeline-preview` read model and CLI command (S6W1). The preview joins diagnostic run history, blocked action evidence, repair recommendation categories, onboarding checklist gaps, KDE action dependency evidence, and KDE journey state into one redacted support-case timeline with stable event ids, event groups, severity, user-safe summaries, and next safe read-only checks. Support ticket creation, bundle export, AI provider calls, repair execution, action execution, backend process start, file-content reads, raw command exposure, local path exposure, state-root path exposure, and host-root mutation all stay disabled.
- Added no-history, joined-evidence, malformed-history, mixed-application, redaction, mismatched-application, CLI, missing-state-root-read-only, malformed-receipt, argument-validation, and no-side-effect tests for the support case timeline preview.

### Changed

- Updated the layout verifier, implementation evidence report, mainline integration review classifier, product documentation, and current Claude Code dispatch picks so S6W1 is recognized as implemented and the next recommended Claude handoff starts with S6W4 Portal permission renewal previews.

## [0.2.295] - 2026-07-17

### Changed

- Updated the current Claude Code dispatch sheet so it reflects the actual v0.2.294 implementation baseline: all seventh-wave tasks plus S6W2, S6W3, and S6W5 are now treated as already implemented, and the recommended next safe handoffs move to S6W1 support case timelines, S6W4 Portal permission renewal previews, S6W6 settings profile migration previews, S6W7 offline fixture matrices, S6W8 merge readiness packets, KDE notification digest previews, signed-recipe verifier evidence, and restricted product smoke packets.
- Updated the release evidence index follow-up guidance and source-file evidence lists so retention, permission audit, KDE search visibility, snapshot restore ranking, and crash/hang summaries are recognized as existing evidence instead of being recommended again. Docker, QEMU, network checks, package managers, backend launch, automatic staging, release tagging, real permission changes, notification sending, and host-root mutation remain disabled by default.

## [0.2.294] - 2026-07-16

### Added

- Added the Go Runtime `recipe-conflict-audit-preview` read model and CLI command (S6W3). The audit parses a recipe registry leniently (so it can surface duplicate ids and digest drift that the strict loader rejects) and reports conflict groups for duplicate application ids, stale recipe versions, unsupported capability claims, mismatched package-source pins, trust-policy blockers, and artifact digest drift, each with deterministic review-only resolution hints. Recipe writes, registry migration, artifact staging, network fetch, package-manager calls, backend launch, and host-root mutation all stay disabled. Added clean, duplicate, stale, missing-pin, unsupported-capability, digest-drift, untrusted, missing-file, and unreadable-registry tests.
- Added the Go Runtime `snapshot-restore-candidates-preview` read model and CLI command (S6W5). The preview reads the snapshot store read-only (via the new `snapshot.OpenReadOnly`) and ranks restore candidates as latest good, latest tested, last-known-running, manual, unknown, or blocked, each with relative evidence ids, reason codes, a compatibility-risk label, digest status, and active-session or digest-mismatch blockers. Restore execution, snapshot deletion, file-content reads, session termination, backend launch, state-root path exposure, and host-root mutation all stay disabled. Added ranking/classification, active-session, no-candidates, unsafe-input, CLI, empty-store-read-only, and path-redaction tests.
- Added the Go Runtime `compatibility-backend-fallback-preview` read model and CLI command (S7W5). The preview explains how Automatic mode would choose between Local compatibility and Isolated compatibility when evidence is missing, blocked, or risky, joining the recipe-derived recommendation, backend capability matrix, backend lifecycle status, and optional diagnostics-risk counts into per-candidate fallback states and user-safe reason codes. It keeps selection persistence, engine installation, backend launch, VM start, raw backend detail exposure, network access, and host-root mutation disabled, and never exposes backend names in KDE-facing output. Added local-primary-blocked, isolation-required, diagnostics-risk, automatic-orchestrator, CLI, and no-side-effect tests.
- Added the Go Runtime `kde-search-visibility-plan-preview` read model and CLI command (S7W6). The preview explains how an application appears across launcher, KRunner, file association, Dolphin action, Compatibility Center, settings, and task manager, deriving stable searchable labels and safety-filtered synonyms, recording MIME/extension evidence, and reporting disabled-action reasons and activation-receipt requirements. Unsafe synonyms (such as a raw `.exe` extension) and duplicates are dropped; desktop-file writes, MIME-default writes, KDE cache refresh, host-file indexing, and search-index persistence stay disabled. Added receipt-required, unsupported-MIME, hidden-application, unsafe-synonym, unsafe-plan, CLI, and no-side-effect tests.
- Added the Go Runtime `desktop-deactivation-dry-run-preview` read model and CLI command (S6W2). The dry run explains how a managed application would be removed from the launcher, desktop icon, MIME association, Dolphin service menu, KWin, tray, notification, Compatibility Center, and settings surfaces, reporting receipt requirements, digest checks, rollback safety, relative evidence ids, and blocked reasons for missing receipt, digest mismatch, unknown file owner, shared MIME association, and active session evidence. File deletion, MIME writes, KDE cache refresh, receipt rewriting, session termination, target-path exposure, and host-root mutation all stay disabled. Added missing-receipt, receipt-error-classification, receipt-backed-removable, active-session, shared-MIME, unsafe-plan, path-redaction, CLI, and no-side-effect tests.
- Added the Go Runtime `permission-evidence-audit-preview` read model and CLI command (S7W7). The audit joins the Runtime permission review plan, KDE resource bridge plan, recorded Portal receipts (rolled up per operation from a read-only ledger read), and execution preflight evidence into one row per capability (documents, downloads, uris, print, clipboard, screenshot, camera, remote desktop, and network policy) with a deterministic consistency state (consistent, setting-only, receipt-only, expired, denied, missing-review, blocked-by-policy, or unsupported), review-only remediation hints, and next safe read-only checks. It tolerates missing receipts, malformed ledger records (surfaced as review evidence, not a failure), and unrecognized permissions, and keeps real Portal transport, permission grants, permission revocation, receipt writes, settings persistence, execution approval, state-root path exposure, and host-root mutation disabled.
- Added the read-only `portal.ReadRequests` lister so audits can enumerate recorded Portal requests under a state root without creating the ledger directory or failing on a single malformed record.
- Added consistent, setting-only, receipt-only, expired, denied, missing-review, policy-blocked, malformed-ledger, unsupported, source-join, CLI, path-redaction, and no-side-effect tests for the permission evidence audit preview and the Portal read-only lister.
- Added the Go Runtime `crash-hang-signal-summary-preview` read model and CLI command (S7W4). The preview reads diagnostic run receipt metadata under an explicit Runtime state root and groups failing signals into crash, hang, timeout, missing-dependency, permission-denial, graphics-issue, network-issue, and regression-after-repair classes, each with occurrence counts, related diagnostic run ids, latest known state, recurrence hints, and a next safe read-only check. It tolerates missing history, malformed receipts (surfaced as blocked evidence rather than a hard failure), mixed application ids, duplicate signal ids, and privacy-sensitive fixture fields, redacting anything that resembles a host path, token, secret, or backend detail. Private log reads, file-content reads, AI provider calls, repair execution, backend launch, network access, state-root path exposure, and host-root mutation all stay disabled.
- Added the read-only `LenientHistory`/`listLenient` diagnostics readers so previews can inspect possibly-corrupt run ledgers without creating directories or failing on a single malformed receipt.
- Added no-history, group/recurrence, malformed-history, mixed-application, duplicate-signal-id, privacy-redaction, blocked-repair, regression, unsafe-input, CLI, and no-side-effect tests for the crash and hang signal summary preview.
- Added the Go Runtime `state-root-quota-retention-preview` read model and CLI command. The preview scans an explicit Runtime state root in read-only dry-run mode, groups snapshots, diagnostics, execution receipts, Portal receipts, artifact receipts, activation receipts, and unknown records, reports relative evidence ids, fixture metadata byte estimates, retention reason codes, blocked cleanup reasons, and user-safe recommendations, and keeps file deletion, directory creation, log truncation, receipt rewriting, snapshot deletion, state-root path exposure, backend detail exposure, network requirements, privileged containers, and host-root mutation disabled.
- Added missing-root, under-quota, over-quota, malformed-record, unknown-record, active-session-blocked, retention-exempt, unsafe-input, CLI, path-redaction, and no-side-effect tests for the state-root quota and retention preview.

### Changed

- Updated the layout verifier, implementation evidence report, and mainline integration review so the recipe conflict and pin audit is verified in the CW2 recipe and artifact trust lane, and the snapshot restore candidate ranking (plus the read-only `snapshot.OpenReadOnly` lister) is verified in the snapshot and rollback receipt lane.
- Updated the layout verifier, implementation evidence report, and mainline integration review so the compatibility backend fallback preview is verified in the CW3 backend lifecycle lane, and the KDE search visibility plan and desktop deactivation dry-run previews are verified in the CW4 KDE entry-point lane.
- Updated the layout verifier, implementation evidence report, and mainline integration review so the permission evidence audit preview and the Portal read-only lister are verified as part of the CW5 Portal permission safety lane.
- Updated the layout verifier, implementation evidence report, and mainline integration review so the crash and hang signal summary preview is verified as part of the CW7 diagnostics privacy lane.
- Updated the layout verifier and product documentation so state-root quota and retention evidence is verified alongside state-root planning, snapshots, diagnostics, execution, and Portal receipt evidence.

## [0.2.293] - 2026-07-16

### Added

- Added the offline Ruby `scripts/release_evidence_index.rb` report command for the S7W8 release evidence index. The report maps release-critical product claims to implemented, fixture-only, contract-only, blocked, skipped, or human-authorized evidence, names source files and verification commands, flags protected Claude-file and unclassified-file blockers, and keeps Docker, QEMU, network checks, package-manager calls, backend launch, automatic staging, release tagging, and host-root mutation disabled.
- Added release evidence index JSON, Markdown, malformed-report, protected-file, unclassified-file, skipped-heavy-smoke, and human-authorization tests.

### Changed

- Updated the layout verifier, implementation evidence report, and mainline integration review classifier so the release evidence index is verified as part of the CW10 evidence and drift harness.

## [0.2.292] - 2026-07-16

### Added

- Added the Go Runtime `runtime-policy-explanation-cards-preview` read model and CLI command. The preview builds a shared KDE-safe explanation card deck for install, launch, execution, Portal permission, snapshot, diagnostics, repair, settings, desktop activation, backend-readiness evidence, and unsupported production routes with stable card ids, severity, audience, related evidence ids, user-facing summaries, technical summaries, next safe read-only checks, and deduplicated blocker reasons while keeping action enablement, card persistence, request-object creation, permission grants, settings persistence, AI provider calls, backend process start, launch, execution, raw command exposure, backend detail exposure, privileged containers, network requirements, and host-root mutation disabled.

### Changed

- Updated the KDE-first presence smoke, layout verifier, implementation evidence report, mainline integration review classifier, README, and product overview so Runtime policy explanation cards are verified alongside upgrade impact, install queue, install readiness, onboarding, support bundle manifest, KDE journey, action dependency graph, route convergence, method parity, and write-gate evidence.

## [0.2.291] - 2026-07-16

### Added

- Added the Go Runtime `application-upgrade-impact-preview` read model and CLI command. The preview compares current and candidate registry recipes before upgrade writes, reports recipe metadata, package source, artifact digest, compatibility profile, Portal permission, snapshot, desktop activation, and diagnostics impact sections with deterministic unchanged, changed, needs-review, missing-evidence, blocked, and unsupported states, and keeps recipe writes, registry migration, artifact staging, downloads, host package-manager calls, settings persistence, desktop activation, backend process start, launch, raw command exposure, backend detail exposure, privileged containers, and host-root mutation disabled.
- Added optional semantic `version` metadata to application recipes so upgrade previews can classify newer, same-version, older, malformed, and unversioned candidate recipes without inferring version order from display names.

### Changed

- Updated the layout verifier and product documentation so application upgrade impact evidence is verified alongside install queue, install readiness, onboarding, support bundle manifest, KDE journey, action dependency graph, route convergence, method parity, and write-gate evidence.

## [0.2.290] - 2026-07-16

### Added

- Added the Go Runtime `multi-application-install-queue-preview` read model and CLI command. The preview aggregates several registry-backed compatibility install plans into one review-only queue with deterministic application ordering, missing-evidence counts, shared blocking reasons, next safe read-only checks, and explicit disabled safety gates while keeping queue persistence, request-object creation, artifact staging, artifact downloads, network requests, host package-manager calls, desktop activation, install execution, backend process start, launch, settings persistence, raw command exposure, backend detail exposure, privileged containers, and host-root mutation disabled.

### Changed

- Updated the KDE-first presence smoke, layout verifier, and mainline integration review classifier so the multi-application install queue is verified alongside install readiness, onboarding, support bundle manifest, KDE journey, action dependency graph, route convergence, method parity, and write-gate evidence.

## [0.2.289] - 2026-07-16

### Added

- Added the Claude Code seventh-wave task batch for application upgrade impact previews, Runtime policy explanation cards, state-root quota and retention previews, crash and hang signal summaries, compatibility backend fallback previews, KDE search visibility plans, permission evidence audits, and release evidence indexes while keeping Docker, QEMU, Runtime writes, backend launch, real Portal calls, network fetch, privileged containers, and host-root mutation disabled by default.

### Changed

- Updated the dispatch runbook, README, product overview, and layout verifier so the seventh-wave Claude Code task batch is a required KDE-first, Go-first handoff artifact with the protected `docs/claude-code-implementation-packages.md` boundary preserved.

## [0.2.288] - 2026-07-16

### Added

- Added the Go Runtime `support-bundle-manifest-preview` read model and CLI command. The preview joins diagnostic history, AI-safe diagnostic input, and review-only recommendation evidence into a redacted offline support bundle manifest with diagnostic run summaries, failing signal ids, recommendation categories, privacy redaction status, and explicit omitted-evidence counts while keeping archive creation, file-content reads, file-path exposure, AI provider calls, automatic repair, backend process start, raw command exposure, backend detail exposure, network access, privileged containers, and host-root mutation disabled.
- Added a read-only diagnostic run store opener so support previews can inspect existing diagnostic history without creating a missing Runtime state root.

### Changed

- Updated the KDE-first presence smoke, layout verifier, and mainline integration review classifier so support bundle manifest evidence is checked alongside onboarding, KDE journey, diagnostic, AI safety, route convergence, method parity, and write-gate evidence.
- Added the Claude Code sixth-wave task batch to the dispatch and layout-verification handoff set for support timelines, deactivation dry-runs, recipe conflict audits, Portal renewal previews, restore ranking, settings migration previews, fixture matrices, and merge readiness packets.

## [0.2.287] - 2026-07-16

### Added

- Added the Go Runtime `compatibility-onboarding-checklist-preview` read model and CLI command. The preview aggregates Runtime owner readiness, recipe trust, artifact staging, backend lifecycle, Portal review, snapshot baseline, diagnostics privacy, KDE entry-point, and production-activation evidence into a first-run checklist with ready, needs-review, missing-evidence, blocked, and not-yet-implemented states while keeping request creation, artifact staging, settings persistence, backend process start, launch, network fetch, host package-manager calls, AI provider calls, raw command exposure, file-content reads, backend detail exposure, privileged containers, and host-root mutation disabled.

### Changed

- Updated the KDE-first presence smoke and layout verifier so compatibility onboarding evidence is checked alongside KDE journey evidence, action dependency graph evidence, route convergence, method parity, and write-gate evidence.

## [0.2.286] - 2026-07-16

### Added

- Added the Go Runtime `kde-journey-evidence-preview` read model and CLI command. The preview stitches application readiness, KDE action dependency graph, task-manager identity, KWin window rules, tray status, notification, Compatibility Center, and unified-settings evidence into one Runtime-owned seven-entrypoint journey summary while keeping write methods, request creation, permission grants, settings persistence, launch, execution, KWin rule application, tray bridge activation, notification sending, state-root path exposure, raw command exposure, file-content reads, backend detail exposure, and host-root mutation disabled.

### Changed

- Updated the KDE-first presence smoke and layout verifier so KDE journey evidence is checked alongside action dependency graph evidence, Center page evidence, route convergence, method parity, and write-gate evidence.

## [0.2.285] - 2026-07-16

### Added

- Added the Claude Code fifth-wave task batch for KDE journey evidence stitching, compatibility onboarding checklists, offline support bundle manifests, multi-application install queue previews, Runtime state maintenance planning, recipe update migration previews, KDE notification digests, and restricted acceptance fixture suites while keeping Docker, QEMU, Runtime writes, backend launch, real Portal calls, network fetch, privileged containers, and host-root mutation disabled by default.

### Changed

- Updated the dispatch runbook and layout verifier so the fifth-wave Claude Code task batch is a required handoff artifact with the protected `docs/claude-code-implementation-packages.md` boundary preserved.

## [0.2.284] - 2026-07-16

### Added

- Added the Go Runtime `kde-action-dependency-graph-preview` read model and CLI command. The preview turns KDE Compatibility Center action queues into explicit Runtime evidence dependency graphs with action nodes, evidence nodes, gate nodes, missing-evidence ids, receipt-validation rules, and next safe read-only checks while keeping dependency graph persistence, settings persistence, permission grants, request creation, Runtime launch approval, execution, backend process start, state-root path exposure, raw command exposure, file-content reads, network access, backend detail exposure, and host-root mutation disabled.

### Changed

- Updated the KDE-first presence smoke and layout verifier so KDE action dependency graph evidence is inspected alongside action card decks, Runtime route convergence, method parity, and write-gate evidence.

## [0.2.283] - 2026-07-16

### Added

- Added the Claude Code fourth-wave task batch for Runtime evidence audit timelines, Portal permission lifecycle dashboards, snapshot restore rehearsal, backend capability fixture probes, settings change dependency reviews, diagnostics repair playbooks, desktop activation materialization audits, and restricted release-readiness packets while keeping Docker, QEMU, Runtime writes, backend launch, real Portal calls, network fetch, privileged containers, and host-root mutation disabled by default.

### Changed

- Updated the dispatch runbook and layout verifier so the fourth-wave Claude Code task batch is a required handoff artifact with the protected `docs/claude-code-implementation-packages.md` boundary preserved.

## [0.2.282] - 2026-07-16

### Added

- Added the Go Runtime `runtime-route-convergence-preview` read model and CLI command. The preview classifies every Runtime read route as Go product logic, C policy bridge, Ruby smoke bridge, fixture-only, contract-only, deprecated, unsupported, or unclassified, emits migration groups and next migration order, fails closed on missing route evidence, and keeps production D-Bus ownership, Runtime write methods, backend launch, Docker, QEMU, network access, privileged containers, backend detail exposure, and host-root mutation disabled.

### Changed

- Updated the KDE-first presence smoke and scaffold verifier so Runtime route convergence evidence is inspected alongside owner route manifest, method parity, and write-gate evidence.

## [0.2.281] - 2026-07-16

### Added

- Added the Claude Code third-wave task batch for Runtime route convergence, action dependency receipts, Portal-to-execution preflight evidence, snapshot baseline readiness, AI diagnostics privacy, KDE action-card decks, unsupported-route hardening, and restricted acceptance dashboards while keeping Docker, QEMU, Runtime writes, backend launch, real Portal calls, network fetch, privileged containers, and host-root mutation disabled by default.

### Changed

- Updated the dispatch runbook and layout verifier so the second-wave and third-wave Claude Code task batches are required handoff artifacts with the protected `docs/claude-code-implementation-packages.md` boundary preserved.

## [0.2.280] - 2026-07-16

### Changed

- Added Runtime execution session fan-out evidence so task manager, KWin, tray, and Compatibility Center consumers share one KDE-safe session record projection. The fan-out keeps live window observation, task-manager activation, KWin rule application, live tray bridging, launch, execution, backend process start, permission grants, state-root path exposure, network access, backend detail exposure, and host-root mutation disabled.

## [0.2.279] - 2026-07-16

### Changed

- Added structured Runtime receipt evidence to the desktop activation transaction preview. The transaction now reports planned commit and rollback receipt schemas, required file IDs, digest gates, commit and rollback availability, and safe path-exposure flags while keeping activation commit, receipt writes, rollback execution, KDE cache refresh, backend launch, execution, network access, privileged containers, and host-root mutation disabled.
- Strengthened layout, KDE-first smoke, implementation evidence, and mainline classification gates so desktop activation transaction receipt evidence remains visible to review tooling without enabling host writes or backend launch.

## [0.2.278] - 2026-07-16

### Changed

- Updated KDE Compatibility Center page sections and section details so they consume Runtime application readiness evidence. Each section now reports the readiness node IDs, readiness status, blocked state, and a safe readiness summary for its backing Runtime evidence while keeping section actions, request creation, launch, execution, backend process start, real Portal transport, snapshot creation, permission grants, network access, backend detail exposure, and host-root mutation disabled.

## [0.2.277] - 2026-07-16

### Changed

- Updated the KDE Compatibility Center page preview so it consumes the Go Runtime `application-readiness-preview` evidence graph directly. The page now exposes Runtime-owned install, backend lifecycle, Portal, snapshot, execution, and Launch write-gate evidence while keeping page persistence, request creation, launch approval, execution, backend process start, real Portal transport, snapshot creation, state-root path exposure, raw command exposure, network access, privileged containers, and host-root mutation disabled.

## [0.2.276] - 2026-07-16

### Added

- Added the Go Runtime `application-readiness-preview` model and CLI command. The preview joins install readiness, backend lifecycle, Portal policy, snapshot planning, execution readiness, and Launch write-gate evidence into one KDE-safe application readiness graph while keeping launch, backend process start, real Portal transport, snapshot creation, network access, privileged containers, raw command exposure, state-root exposure, and host-root mutation disabled.

## [0.2.275] - 2026-07-16

### Added

- Added the Go Runtime `desktop-safety-policy-preview` model and CLI command. The policy owns the KDE-first seven-entrypoint list, user-facing settings field ids, forbidden backend terminology, and disabled safety keys consumed by the KDE-first presence smoke.

### Changed

- Updated `scripts/kde_first_presence_smoke.rb` so Ruby remains the test harness while Go Runtime owns the desktop safety policy used to reject Wine, Prefix, backend command, host-path, Portal, execution, network, privileged-container, and host-root safety regressions.

## [0.2.274] - 2026-07-16

### Changed

- Strengthened the KDE-first presence smoke user-facing safety gate. The smoke now verifies the unified settings fields for run mode, priority, documents, downloads, camera, network, and snapshots, reports those fields in JSON/Markdown, and rejects user-visible Prefix, bottle, Wine, Proton, `.exe`, Program Files, QEMU, and host-path terminology across inspected KDE Runtime previews.

## [0.2.273] - 2026-07-16

### Changed

- Promoted the KDE-first presence smoke into the KDE activation and shell materialization implementation evidence domain. `scripts/implementation_evidence_report.rb` now treats `scripts/kde_first_presence_smoke.rb` and its tests as `M5` smoke-owned evidence and verifies the report schema, route baseline, seven-entry-point coverage, and disabled execution, backend launch, Portal transport, AI provider, privileged-container, and host-root gates.

## [0.2.272] - 2026-07-16

### Added

- Added JSON and Markdown report modes to `scripts/kde_first_presence_smoke.rb`. The KDE-first presence smoke now emits machine-readable and human-reviewable evidence for the seven KDE entry points, Runtime route baseline, disabled execution gates, disabled backend launch, host-root safety, network safety, and privileged-container safety while keeping the default text PASS output.

## [0.2.271] - 2026-07-16

### Added

- Added lane-level review matrix output to `scripts/mainline_integration_review.rb`. The report now includes required verification commands, required evidence, and safety guards for each changed mainline lane so Claude and Codex branches can be staged by explicit KDE-first Runtime-owned acceptance criteria instead of broad worktree guesses.

## [0.2.270] - 2026-07-16

### Added

- Added a Claude Code dispatch runbook that explains how to assign the `CB1` through `CB9` mainline task batch, review returned branches by lane, protect `docs/claude-code-implementation-packages.md`, avoid `git add .`, and keep Docker, QEMU, backend launch, real Portal calls, host package-manager calls, privileged containers, and host-root mutation disabled unless explicitly authorized.

## [0.2.269] - 2026-07-16

### Changed

- Strengthened the mainline integration review classifier so shared Runtime CLI plumbing, Windows compatibility workstream models, snapshot/rollback evidence, and AI diagnostic privacy files are assigned to explicit review lanes instead of remaining unclassified.

## [0.2.268] - 2026-07-16

### Added

- Added a copy-first Claude Code mainline task batch for the KDE-first Windows compatibility strategy. The task batch splits follow-up work into focused `CB1` through `CB9` assignments covering Runtime owner reads, recipe and artifact trust, backend lifecycle state, Portal permission reviews, snapshots, KDE evidence consumers, AI diagnostics, integration review gates, and later restricted product smoke readiness.
- Added the mainline integration review script and fixture-driven tests. `scripts/mainline_integration_review.rb` classifies the current worktree by checkpoint lane, excludes `.gocache/` and `tmp/`, flags protected Claude-owned file changes, emits JSON and Markdown, and never stages files, runs Docker, runs QEMU, launches backends, or mutates the host root.

## [0.2.267] - 2026-07-16

### Added

- Added the mainline integration checkpoint for merging Claude Code output against the KDE-first Windows compatibility goal. The checkpoint groups the current mixed worktree into review lanes for Runtime owner reads, trust and artifacts, state roots, Portal safety, execution evidence, KDE consumers, evidence gates, and later product-image acceptance while keeping unsafe behavior disabled.
- Added artifact staging receipt validation and install-readiness consumption. `compatibility-install-preview --artifact-receipt <path>` now reads a Go Runtime artifact staging receipt, verifies schema, relative path scope, receipt digest, required artifact staging, ownership flags, root-hiding flags, and unsafe side-effect gates before exposing desktop-safe receipt diagnostics.

### Changed

- Strengthened M2/CW2 recipe and artifact trust behavior so tampered or unsafe artifact receipts fail closed in install readiness while valid local fixture receipts can prove required artifacts are staged without enabling downloads, package-manager calls, desktop activation writes, backend launch, host-root mutation, or path exposure.

## [0.2.266] - 2026-07-16

### Added

- Added Runtime contract drift coverage for the Go Runtime owner private session-bus smoke transcript. The drift report now verifies `NewSessionBusSmokeTranscript`, `--session-bus-smoke`, the private session-bus schema, unsupported-read rejection, production-bus denial, host-root safety, backend-detail hiding, and write-method gating.

### Changed

- Strengthened the Runtime owner candidate restricted smoke script so it executes `xnix-runtime-owner --session-bus-smoke` and verifies the 71-step transcript covering startup, private bus claim, route-table readiness, 62 read dispatches, 4 disabled write methods, unsupported-read failure, and shutdown.

## [0.2.265] - 2026-07-16

### Added

- Added a Go Runtime owner private session-bus smoke transcript. `xnix-runtime-owner --session-bus-smoke` now emits JSONL evidence that wraps the full owner read-dispatch table, all disabled write-method responses, and unsupported-read rejection inside a restricted private session-bus event-loop model while keeping production bus ownership, system service startup, write dispatch enablement, network access, backend detail exposure, and host-root mutation disabled.

### Changed

- Strengthened M1/CW1 implementation evidence so the Runtime owner service domain requires private session-bus smoke transcript files and tokens in addition to in-process service-call, lifecycle, and smoke-batch evidence.

## [0.2.264] - 2026-07-16

### Added

- Added execution-session-record consumption to the Go KDE Compatibility Center page preview. `xnix-runtime-go kde-center-page-preview --session-root <path> --session-request-id <id>` now composes safe session evidence into the window and tray snapshots while exposing only relative evidence paths and keeping page persistence, request creation, task-manager activation, KWin rule application, live tray bridging, launch, backend processes, state-root path exposure, and host-root mutation disabled.

### Changed

- Strengthened M5 implementation evidence and scaffold verification so KDE Compatibility Center page previews require session-root CLI coverage and session-backed window/tray snapshot safety evidence.

## [0.2.263] - 2026-07-16

### Added

- Added KDE-facing execution session evidence consumption for task-manager identity, KWin window-rule, and tray-status Go read models. These previews can now read a safe `execution-session-record` from an explicit state root and expose only relative session evidence paths and blocked session state, without activating task-manager entries, applying KWin rules, enabling live tray bridges, observing windows, launching backends, exposing state-root paths, or mutating the host root.

### Changed

- Strengthened M5 implementation evidence and scaffold verification so KDE shell materialization requires execution-session-record evidence validation plus task-manager, KWin, and tray consumption tests.

## [0.2.262] - 2026-07-16

### Added

- Added a state-root-scoped Go Runtime execution session status record and `xnix-runtime-go execution-session-record --state-root <path> --request-id <id>`, deriving KDE-safe task-manager, tray, KWin, and Compatibility Center session evidence from an existing execution ledger transaction without creating a live session, observing windows, launching backends, applying KWin rules, activating tray bridges, exposing state-root paths, or mutating the host root.

### Changed

- Strengthened implementation evidence and scaffold verification so M6 requires execution session record schema, CLI coverage, and session safety tests in addition to transaction ledger and Portal receipt evidence.

## [0.2.261] - 2026-07-16

### Added

- Added optional Portal permission receipt consumption to the Go Runtime execution ledger. `xnix-runtime-go execution-ledger-record --portal-request <handle>` now reads a recorded Portal request from the same state root, uses granted/completed receipts to satisfy the Portal preflight gate, and records denied or not-granted receipts as blocking evidence without enabling launch, backend start, execution approval, host permission changes, real Portal calls, state-root path exposure, or host-root mutation.

### Changed

- Strengthened implementation evidence and scaffold verification so M6 requires execution ledger Portal receipt fields, CLI consumption coverage, and granted/denied receipt tests in addition to state-root transaction persistence.

## [0.2.260] - 2026-07-16

### Added

- Added a state-root-scoped Go Runtime Portal request ledger and `xnix-runtime-go portal-request-record`, allowing fake-mode Portal permission requests to be created, inspected, resolved, completed, cancelled, and expired under an explicit state root without making real Portal calls, approving execution, exposing state-root paths, changing host permissions, or mutating the host root.

### Changed

- Strengthened implementation evidence and scaffold verification so M4 requires Portal request record schema, state-root ledger implementation, CLI coverage, and permission-flow tests in addition to fake broker and snapshot-store evidence.

## [0.2.259] - 2026-07-16

### Added

- Added `xnix-runtime-go backend-lifecycle-record`, a Go Runtime CLI that can inspect, plan, stage, satisfy lifecycle gates, mark ready, flag repair, block, and retire backend lifecycle records under an explicit state root without launching backends, exposing state-root paths, requiring network, or mutating the host root.

### Changed

- Strengthened implementation evidence and scaffold verification so M3/CW3 requires state-root backend lifecycle record schema, Go implementation, CLI coverage, and transition tests in addition to backend manager inventory evidence.

## [0.2.258] - 2026-07-16

### Added

- Added a state-root-backed Go Runtime backend manager inventory record and `backend-manager-record --state-root <path>` CLI, persisting Wine, Proton, and Windows VM backend inventory under an explicit state root without installing, downloading, launching, starting VM processes, exposing paths or commands, requiring network, or mutating the host root.

### Changed

- Strengthened implementation evidence and scaffold verification so M3/CW3 requires backend manager state-root inventory persistence in addition to backend manager preview and lifecycle state-root evidence.

## [0.2.257] - 2026-07-16

### Added

- Added a Go Runtime backend manager preview and CLI command that tracks Wine, Proton, and Windows VM backends as internal Runtime-managed inventory while keeping installs, downloads, launches, backend processes, VM starts, raw commands, profile paths, network requirements, privileged containers, and host-root mutation disabled.

### Changed

- Strengthened implementation evidence and scaffold verification so M3/CW3 now requires backend manager inventory evidence, backend-safe user-facing profile tests, and CLI coverage in addition to state-root-backed backend lifecycle evidence.

## [0.2.256] - 2026-07-16

### Changed

- Strengthened implementation evidence reporting so M3/CW3 now requires Go backend lifecycle state-root preview methods, hidden state-root path safety fields, package tests, and CLI `--state-root` coverage before the environment lifecycle domain can remain marked state-root implemented.
- Updated version metadata and product documentation to `v0.2.256` while keeping production D-Bus ownership, write methods, backend launch, real Portal calls, network fetch, privileged containers, and host-root mutation disabled.

## [0.2.255] - 2026-07-16

### Added

- Added a copy-first Claude Code Windows compatibility workstream board that maps the KDE-first product target into `CW1` through `CW11` implementation streams with Go-first Runtime ownership, KDE seven-entry-point integration, safety gates, and verification commands.
- Added the Go Runtime `windows-compatibility-workstreams-preview` read model and CLI command so KDE-first Windows compatibility dispatch can be consumed as structured JSON without starting backends, enabling write methods, mutating the host root, or exposing backend details.
- Added owner-local Runtime read dispatch and service-call coverage for `GetWindowsCompatibilityWorkstreamsPreview`, so the Go owner can serve the KDE-first Windows compatibility workstream model without adding a production D-Bus method or expanding the formal read-only ABI.
- Added recipe trust diagnostics to the Go compatibility install readiness model, so install plans now expose stable `recipe_trust_diagnostics_ready` and `recipe_trust_blocking_reasons` fields for production-blocked development recipes without enabling downloads, installs, backend launch, or host-root writes.
- Added optional state-root backed Go backend lifecycle previews, so `backend-lifecycle-preview --state-root <path>` can project persisted environment lifecycle records, satisfied gates, pending gates, repair hints, and blocked reasons without exposing state-root paths, enabling launch, starting backends, or mutating the host root.
- Added Windows compatibility first-wave workstream guidance to the implementation evidence report, so JSON and Markdown outputs now expose `CW1`, `CW2`, `CW3`, and `CW10` next-dispatch evidence alongside the existing mainline dispatch.

### Changed

- Linked the Windows compatibility workstream board from the Windows application compatibility brief, the next-implementation assignment board, and the scaffold layout verifier so the KDE-first Claude Code handoff path remains discoverable and required.
- Updated Runtime contract drift reporting so D-Bus read-only method coverage remains fixed at the formal contract while owner-local read dispatch separately tracks the Windows compatibility workstream preview and smoke-batch evidence.
- Updated Runtime owner recipe trust previews to consume the Go recipe `Store` and replaceable `DigestVerifier` boundary instead of duplicating registry digest checks in the owner read model.
- Strengthened compatibility install plan tests and scaffold verification so recipe trust diagnostics must remain visible through both Go package and CLI output.
- Strengthened backend lifecycle tests and scaffold verification so Go lifecycle previews must keep their environment state-root projection and launch/backend/host-root safety gates.
- Strengthened scaffold verification so the Windows compatibility workstream board and Go Runtime preview must retain KDE-first shell selection, Go-first Runtime ownership, the `CW1` through `CW11` workstream set, the `CW1`/`CW2`/`CW3`/`CW10` first wave, seven KDE entry points, and the protected `docs/claude-code-implementation-packages.md` boundary.

## [0.2.254] - 2026-07-16

### Added

- Added receipt-backed KDE application surface previews through `kde-application-surface-preview --activation-root`, so the seven-entry-point KDE surface overview can prove staged activation evidence before any shell materialization remains gated.
- Added Go and CLI coverage proving KDE application surface previews consume staged activation receipts without exposing the activation root, enabling launch, writing desktop files or MIME data, mutating the host root, or exposing backend details.

### Changed

- Updated KDE activation implementation evidence so P5 tracks KDE application surface consumption of staged activation receipts alongside activation status, Compatibility Center, desktop entry, KRunner, Dolphin file-manager, desktop icon, MIME association, tray status, notification, settings, task-manager identity, and KWin window-rule reads.

## [0.2.253] - 2026-07-16

### Added

- Added receipt-gated KDE desktop entry previews through `desktop-entry-preview --activation-root`, so launcher rendering can require staged activation evidence while preserving standard `.desktop` output.
- Added Go and CLI coverage proving desktop entry previews can consume staged activation receipts without exposing the activation root, writing activation files, enabling launch, mutating the host root, or exposing backend details.

### Changed

- Updated KDE activation implementation evidence so P5 tracks desktop entry consumption of staged activation receipts alongside activation status, Compatibility Center, KRunner, Dolphin file-manager, desktop icon, MIME association, tray status, notification, settings, task-manager identity, and KWin window-rule reads.

## [0.2.252] - 2026-07-16

### Added

- Added receipt-gated KDE MIME association previews through `mimeapps-preview --activation-root`, so file association rendering can require staged activation evidence while preserving standard `mimeapps.list` output.
- Added Go and CLI coverage proving MIME association previews can consume staged activation receipts without exposing the activation root, writing activation files, overwriting associations, mutating the host root, or exposing backend details.

### Changed

- Updated KDE activation implementation evidence so P5 tracks MIME association consumption of staged activation receipts alongside activation status, Compatibility Center, KRunner, Dolphin file-manager, desktop icon, tray status, notification, settings, task-manager identity, and KWin window-rule reads.

## [0.2.251] - 2026-07-16

### Added

- Added receipt-backed KDE desktop icon previews through `desktop-icon-preview --activation-root`, so Plasma desktop shortcut reads can prove staged activation evidence before copy, placement, or launch actions remain gated.
- Added Go and CLI coverage proving desktop icon previews consume staged activation receipts without exposing the activation root, copying desktop files, persisting icon placement, enabling launch, mutating the host root, or exposing backend details.
- Added a copy-first Claude Code priority implementation package board for large independent follow-up branches covering Runtime ownership, recipe/artifact trust, environment lifecycle state, Portal/snapshot safety, KDE activation materialization, execution transactions, diagnostics/repair/AI, evidence harnesses, atomic KDE/QEMU acceptance, and Runtime packaging/service binding.

### Changed

- Updated KDE activation implementation evidence so P5 tracks desktop icon consumption of staged activation receipts alongside activation status, Compatibility Center, KRunner, Dolphin file-manager, tray status, notification, settings, task-manager identity, and KWin window-rule reads.
- Linked the new Claude Code priority implementation package board from the domain dispatch, mainline implementation plan, and product overview.

## [0.2.250] - 2026-07-16

### Added

- Added receipt-backed KDE KWin window-rule previews through `kwin-window-rule-preview --activation-root`, so KWin identity and layout reads can prove staged activation evidence before rule application remains gated.
- Added Go and CLI coverage proving KWin window-rule previews consume staged activation receipts without exposing the activation root, applying KWin rules, observing windows, enabling launch or execution, activating task-manager entries, mutating the host root, or exposing backend details.

### Changed

- Updated KDE activation implementation evidence so P5 tracks KWin window-rule consumption of staged activation receipts alongside activation status, Compatibility Center, KRunner, Dolphin file-manager, task-manager identity, tray status, notification, and settings reads.
- Extended the constrained QEMU boot-system timeout to 180 seconds so tenth-version full smoke can reach the same serial login markers on slower TCG hosts without weakening host-safety gates.

## [0.2.249] - 2026-07-16

### Added

- Added receipt-backed KDE task-manager identity previews through `task-manager-identity-preview --activation-root`, so Plasma taskbar grouping and restore reads can prove staged activation evidence before task-manager activation remains gated.
- Added Go and CLI coverage proving task-manager identity previews consume staged activation receipts without exposing the activation root, observing windows, enabling launch or execution, activating task-manager entries, mutating the host root, or exposing backend details.

### Changed

- Updated KDE activation implementation evidence so P5 tracks task-manager identity consumption of staged activation receipts alongside activation status, Compatibility Center, KRunner, Dolphin file-manager, tray status, notification, and settings reads.

## [0.2.248] - 2026-07-16

### Added

- Added receipt-backed KDE settings previews through `settings-preview --activation-root`, so unified settings reads can prove staged activation evidence before settings persistence remains gated.
- Added Go and CLI coverage proving settings previews consume staged activation receipts without exposing the activation root, persisting settings, changing permissions, mutating the host root, or exposing backend details.

### Changed

- Updated KDE activation implementation evidence so P5 tracks unified settings consumption of staged activation receipts alongside activation status, Compatibility Center, KRunner, Dolphin file-manager, tray status, and notification reads.

## [0.2.247] - 2026-07-16

### Added

- Added receipt-backed KDE notification previews through `notification-preview --activation-root`, so notification-center reads can prove staged activation evidence before delivery remains gated.
- Added Go and CLI coverage proving notification previews consume staged activation receipts without exposing the activation root, sending notifications, enabling notification actions, enabling repairs, persisting settings, mutating the host root, or exposing backend details.

### Changed

- Updated KDE activation implementation evidence so P5 tracks notification-center consumption of staged activation receipts alongside activation status, Compatibility Center, KRunner, Dolphin file-manager, and tray status reads.

## [0.2.246] - 2026-07-16

### Added

- Added receipt-backed KDE tray status previews through `tray-status-preview --activation-root`, so system tray reads can prove staged activation evidence before live tray bridging remains gated.
- Added Go and CLI coverage proving tray status consumes staged activation receipts without exposing the activation root, persisting bridge configuration, mutating the host root, enabling live tray bridging, or exposing backend details.

### Changed

- Updated KDE activation implementation evidence so P5 tracks system tray consumption of staged activation receipts alongside activation status, Compatibility Center, KRunner, and Dolphin file-manager reads.

## [0.2.245] - 2026-07-16

### Added

- Added receipt-backed Dolphin file-open and drag-and-drop previews through explicit `--activation-root` options, so file-manager routing can prove staged activation evidence before launch remains gated.
- Added Go and CLI coverage proving Dolphin file routing consumes staged activation receipts without exposing the activation root, granting file permissions, enabling direct host-file access, mutating the host root, or starting compatibility backends.
- Added a compact Claude Code mainline empty-domain handoff board for branch-sized implementation packages that should convert contract-heavy domains into durable evidence.

### Changed

- Updated KDE activation implementation evidence so P5 tracks Dolphin file-manager consumption of staged activation receipts alongside activation status, Compatibility Center, and KRunner reads.

## [0.2.244] - 2026-07-16

### Added

- Added receipt-backed KRunner query previews through `krunner-query-preview --activation-root`, so KDE search matches can prove staged activation evidence before launch remains gated.
- Added Go and CLI coverage proving KRunner search consumes staged activation receipts without exposing the activation root, mutating the host root, enabling query execution, or starting compatibility backends.

### Changed

- Updated KDE activation implementation evidence so P5 tracks KRunner consumption of staged activation receipts alongside activation status and Compatibility Center page reads.

## [0.2.243] - 2026-07-16

### Added

- Added receipt-backed activation snapshot support to KDE Compatibility Center page previews through an explicit `--activation-root` option.
- Added Go and CLI coverage proving `kde-center-page-preview` can compose staged activation receipt evidence without exposing the staging root or enabling launch.

### Changed

- Updated KDE activation implementation evidence so P5 tracks Compatibility Center consumption of staged activation receipts in addition to activation status reads.

## [0.2.242] - 2026-07-16

### Added

- Added receipt-backed desktop activation status evidence so KDE-facing activation reads can consume staged Runtime activation receipts from an explicit root.
- Added safe receipt evidence fields for receipt schema, relative receipt path, installed file IDs, digest-gated rollback readiness, and KDE safety status without exposing the staging root.
- Added CLI and Go coverage for `desktop-activation-status-preview --activation-root`, including receipt-present, receipt-missing, and application-mismatch failure paths.

### Changed

- Updated implementation evidence reporting so the KDE activation and shell materialization domain tracks receipt-backed activation status, not only staged artifact writing.

## [0.2.241] - 2026-07-16

### Changed

- Included the full build and QEMU smoke report writer in the atomic KDE image and QEMU acceptance implementation-evidence domain.
- Added implementation-evidence checks for full-smoke JSON/Markdown report generation, persisted serial-log evidence, QEMU network restriction, and host-safety gates.
- Strengthened implementation-evidence tests so the M9 image/QEMU domain cannot regress to image scaffolding without the milestone smoke report evidence.
- Clarified the short Claude Code assignment board with a recommended first wave for large, independent contract-to-implementation domains.

## [0.2.240] - 2026-07-16

### Added

- Added a structured full build and QEMU smoke report writer that emits ignored JSON and Markdown artifacts under `output/` after the milestone boot attempt.
- Recorded full-smoke report fields for version, ordered build steps, persisted serial-log path, boot markers, host-safety gates, QEMU network restriction, and source-download network use.
- Added unit coverage and layout requirements for the full-smoke report so tenth-version QEMU milestone runs leave auditable evidence.

## [0.2.239] - 2026-07-16

### Changed

- Strengthened the restricted D-Bus session smoke so every read-only Runtime method must expose a parseable Go owner service-call envelope, not only string-level service-call evidence.
- Added contract drift coverage for session smoke service-call envelope assertions, including method preservation, dispatch readiness, write-gate status, KDE policy ownership, host mutation, network, and backend-detail gates.
- Updated D-Bus smoke script unit and layout checks to keep future D-Bus bridge work aligned with the in-process Go owner service boundary.

## [0.2.238] - 2026-07-16

### Changed

- Routed Runtime owner smoke-batch read and write records through the in-process Go owner `Service.Call` boundary instead of calling read dispatch and disabled-write helpers directly.
- Updated owner smoke-batch JSONL payloads so each record now carries a `xnix.runtime.owner_service_call.v1` envelope with nested read-dispatch or write-denial evidence.
- Strengthened contract drift, implementation evidence, layout, CLI, and restricted session smoke assertions to keep smoke-batch coverage aligned with the service-call boundary.

## [0.2.237] - 2026-07-16

### Added

- Added Go owner service-call bridge evidence to the constrained C D-Bus smoke adapter, alongside the existing owner read-dispatch evidence.
- Extended restricted D-Bus session smoke assertions so every Go-owner-backed read verifies `xnix.runtime.owner_service_call.v1` payload evidence from `xnix-runtime-owner --service-call`.

### Changed

- Updated Runtime contract drift, implementation evidence, layout, and D-Bus smoke-script gates to prevent regressions from the service-call bridge back to dispatch-only D-Bus evidence.
- Kept the existing `go_owner_dispatch_*` D-Bus fields for compatibility while adding `go_owner_service_call_*` fields for the next Runtime owner service transition.

## [0.2.236] - 2026-07-16

### Added

- Added an in-process Go Runtime owner service-call boundary that routes read-only Runtime methods through the existing owner read dispatch table and routes reserved write methods to deterministic disabled-write responses.
- Added `xnix-runtime-owner --service-call <method> [args...]` so the owner process can render service-call evidence for read dispatch and write denial without claiming D-Bus ownership.
- Extended the Runtime owner candidate smoke to verify service-call read dispatch and service-call write denial inside the restricted session-bus harness.

### Changed

- Updated Runtime owner implementation evidence reporting to track the new service-call boundary as M1 Runtime owner service evidence while keeping production bus ownership, system service start, write dispatch, network access, privileged containers, and host-root mutation disabled.

## [0.2.235] - 2026-07-16

### Added

- Exposed Runtime owner readiness and self-description reads on the read-only D-Bus contract: `GetRuntimeOwnerProcess`, `GetRuntimeOwnerRouteManifest`, `GetRuntimeOwnerRecipeTrust`, and `GetRuntimeOwnerReadiness`.
- Added constrained C D-Bus smoke read models and restricted session smoke assertions for those owner readiness reads, each carrying Go owner read-dispatch evidence while keeping production bus ownership, writes, service start, backend launch, network access, privileged containers, and host-root mutation disabled.

### Changed

- Expanded Runtime read-only parity, owner route, D-Bus client, C core parity, and implementation evidence gates from 57 to 61 read methods.
- Kept Ruby RuntimeDaemon support as a thin migration surface while Go owner dispatch remains the durable source for owner readiness payload evidence.

## [0.2.234] - 2026-07-16

### Added

- Extended the constrained C D-Bus smoke bridge so foundation catalog and status reads carry Go owner read-dispatch payload evidence: application list, application detail, diagnostics, engine catalog, and application state root.
- Added restricted session smoke assertions that prove foundation catalog and status reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without production D-Bus ownership, backend launch, state-root creation, network access, privileged containers, or host-root mutation.

### Changed

- Added a single-argument payload-only Go owner dispatch bridge helper for C D-Bus application dictionaries, allowing `ListApplications` and `GetApplication` to preserve their existing low-level shape while exposing owner dispatch JSON.
- Updated Runtime contract drift and implementation evidence reports to track foundation catalog and status reads as Go-owned payload handoffs through low-level C D-Bus transport.
- Clarified the Claude Code large empty-domain assignment handoff protocol so independent branches convert contract-only areas into verifiable implementation evidence instead of adding new empty contracts.

## [0.2.233] - 2026-07-16

### Added

- Extended the constrained C D-Bus smoke bridge so KDE Compatibility Center page reads carry Go owner read-dispatch payload evidence: full page, page sections, and section detail.
- Added restricted session smoke assertions that prove KDE Compatibility Center page reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without persisting pages, enabling section actions, sending notifications, granting resources, approving launches, production D-Bus ownership, or host-root mutation.

### Changed

- Added payload-only Go owner dispatch bridge helpers for C D-Bus builders that already own a specific KDE read-model source, preserving `go-kde-center-*` source labels while exposing owner dispatch JSON.
- Updated Runtime contract drift and implementation evidence reports to track KDE Compatibility Center page reads as Go-owned payload handoffs through low-level C D-Bus transport.

## [0.2.232] - 2026-07-16

### Added

- Added `docs/claude-code-contract-to-implementation-task-board.md`, a practical Claude Code task board for converting thin contract domains into durable implementation evidence without touching the existing implementation package guide.
- Extended the constrained C D-Bus smoke bridge so settings and review reads carry Go owner read-dispatch payload evidence: settings change plan, compatibility mode switch plan, permission review plan, review flow plan, Compatibility Center action queue, and action review receipt.
- Added restricted session smoke assertions that prove settings, review, and Compatibility Center action reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without persisting settings, granting permissions, creating Portal requests, executing actions, production D-Bus ownership, or host-root mutation.

### Changed

- Extended the C D-Bus Go owner bridge helper to support five-argument read dispatch for review-flow routes while preserving existing one-, two-, and three-argument bridges.
- Linked the contract-to-implementation task board from the Claude Code dispatch, next-assignment, and mainline handoff documents.
- Updated Runtime contract drift and implementation evidence reports to track settings/review and Compatibility Center action reads as Go-owned payload handoffs through low-level C D-Bus transport.

## [0.2.231] - 2026-07-16

### Added

- Extended the constrained C D-Bus smoke bridge so AI safety reads carry Go owner read-dispatch payload evidence: AI diagnostic input, AI diagnostic recommendation, AI repair approval gate, snapshot plan, and Portal access policy.
- Added restricted session smoke assertions that prove AI safety reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without calling AI providers, reading file contents, executing repairs, creating snapshots, granting Portal access, production D-Bus ownership, or host-root mutation.

### Changed

- Extended the C D-Bus Go owner bridge helper to support three-argument read dispatch for AI diagnostic routes while preserving the existing one- and two-argument read bridges.
- Updated Runtime contract drift and implementation evidence reports to track AI safety reads as Go-owned payload handoffs through low-level C D-Bus transport.

## [0.2.230] - 2026-07-16

### Added

- Extended the constrained C D-Bus smoke bridge so execution readiness reads carry Go owner read-dispatch payload evidence: run plan, repair plan, test plan, test result, execution readiness, and launch intent.
- Added restricted session smoke assertions that prove execution readiness reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without creating execution requests, running tests, executing repairs, enabling launch, production D-Bus ownership, or host-root mutation.
- Added `docs/claude-code-next-implementation-assignments.md`, a short copyable Claude Code assignment board for handing off large implementation packages without editing the existing implementation package guide.

### Changed

- Updated Runtime contract drift and implementation evidence reports to track execution readiness reads as Go-owned payload handoffs through low-level C D-Bus transport.
- Linked the short Claude Code assignment board from README, Product Overview, and the domain dispatch document.

## [0.2.229] - 2026-07-16

### Added

- Extended the constrained C D-Bus smoke bridge so backend lifecycle reads carry Go owner read-dispatch payload evidence: backend binding, backend capability matrix, backend selection plan, backend lifecycle, and backend environment plan.
- Added restricted session smoke assertions that prove backend lifecycle reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without creating environments, starting Wine/VM backends, enabling launch, production D-Bus ownership, or host-root mutation.

### Changed

- Updated Runtime contract drift and implementation evidence reports to track backend lifecycle reads as Go-owned payload handoffs through low-level C D-Bus transport.

## [0.2.228] - 2026-07-16

### Added

- Extended the constrained C D-Bus smoke bridge so install-input reads carry Go owner read-dispatch payload evidence: package source, acquisition preflight, artifact manifest, and compatibility install plan.
- Added restricted session smoke assertions that prove install-input reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without enabling downloads, package-manager calls, desktop activation writes, backend launch, production D-Bus ownership, or host-root mutation.

### Changed

- Updated Runtime contract drift and implementation evidence reports to track install-input reads as Go-owned payload handoffs through low-level C D-Bus transport.

## [0.2.227] - 2026-07-16

### Added

- Added `docs/claude-code-domain-dispatch.md`, a first-stop dispatch board for assigning large independent Claude Code implementation domains when Runtime contracts exist but durable implementation evidence is still thin.
- Added D1 through D10 dispatch prompts covering Runtime owner process boundaries, recipe/artifact/install trust, state-root lifecycle, Portal/snapshot safety, KDE activation, execution transactions, diagnostics/repair/AI, developer evidence gates, KDE/QEMU acceptance, and Runtime packaging/service binding.

### Changed

- Linked the new dispatch board from the mainline handoff, large empty-domain assignment map, README, Product Overview, and layout verification so future Claude Code branches can start from a convergent task index instead of wandering through overlapping package catalogs.

## [0.2.226] - 2026-07-16

### Added

- Extended the constrained C D-Bus smoke bridge so desktop activation reads carry Go owner read-dispatch payload evidence: activation manifest, activation transaction preview, and activation status.
- Added restricted session smoke assertions that prove desktop activation reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without enabling activation commits, desktop writes, KDE cache refresh, backend launch, production D-Bus ownership, or host-root mutation.

### Changed

- Updated Runtime contract drift and implementation evidence reports to track desktop activation reads as Go-owned payload handoffs through low-level C D-Bus transport.

## [0.2.225] - 2026-07-16

### Added

- Extended the constrained C D-Bus smoke bridge so second-ring KDE and Portal resource reads carry Go owner read-dispatch payload evidence: desktop icons, desktop resource bridges, KWin window rules, KRunner queries, and Portal request plans.
- Added restricted session smoke assertions that prove those second-ring KDE and Portal reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without enabling desktop writes, KWin rule application, query execution, Portal request creation, backend launch, production D-Bus ownership, or host-root mutation.

### Changed

- Updated Runtime contract drift and implementation evidence reports to track the second-ring KDE resource bridge as a Go-owned payload handoff through low-level C D-Bus transport.

## [0.2.224] - 2026-07-16

### Added

- Extended the constrained C D-Bus smoke bridge so the seven first-release KDE entry-point reads carry Go owner read-dispatch payload evidence: launcher, task manager, file manager, tray, notifications, Compatibility Center, and settings.
- Added restricted session smoke assertions that prove those seven KDE entry-point reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without enabling desktop writes, backend launch, production D-Bus ownership, or host-root mutation.
- Added `docs/claude-code-large-empty-domain-assignments.md`, a coarse Claude Code dispatcher that groups contract-heavy empty domains into large branch-sized assignments with explicit boundaries, tests, and copyable prompts.

### Changed

- Updated Runtime contract drift and implementation evidence reports to track the seven-entry-point bridge as a Go-owned payload handoff through low-level C D-Bus transport.
- Linked the large empty-domain assignment dispatcher from the existing Claude Code handoff documents so future implementation branches can choose between coarse and fine-grained task maps.

## [0.2.223] - 2026-07-16

### Added

- Extended the constrained C D-Bus smoke bridge so `GetKDEIntegrationStatus`, `GetKDEShellIntegrationPlan`, and `GetKDEApplicationSurfacePlan` carry Go owner read-dispatch payload evidence alongside their stable KDE shell fields.
- Added restricted session smoke assertions that prove KDE shell and application-surface reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without forking Plasma, enabling writes, launching backends, or mutating the host root.

### Changed

- Updated Runtime contract drift and implementation evidence reports to track KDE shell bridge coverage as a Go-owned payload handoff through low-level C D-Bus transport.
- Hardened the temporary Ruby D-Bus client parser so Go owner JSON bridge fields do not hide the outer KDE read model consumed by smoke clients.

## [0.2.222] - 2026-07-16

### Added

- Extended the constrained C D-Bus smoke bridge so `GetRuntimeServiceBinding`, `GetRuntimeLiveOwnerGate`, `GetRuntimeOwnerSmokePlan`, and `GetRuntimeMethodParityManifest` carry Go owner read-dispatch payload evidence alongside their stable smoke fields.
- Added restricted session smoke assertions that prove the Runtime owner self-description reads expose `xnix.runtime.owner_read_dispatch.v1` evidence without claiming production ownership, enabling writes, starting backends, or mutating the host root.

### Changed

- Refactored the D-Bus smoke bridge to use a shared Go owner read-dispatch helper while keeping `GetRuntimeWriteGate` write denial behavior explicit and disabled.
- Updated implementation evidence and contract drift reporting to track owner self-description bridge coverage as a Go-owned payload handoff through low-level C transport.

## [0.2.221] - 2026-07-16

### Added

- Added Go-owner-backed D-Bus smoke bridge evidence for `GetRuntimeWriteGate`, letting the existing C session-bus adapter return the Go owner read-dispatch JSON alongside its stable write-gate fields.
- Added restricted session smoke assertions that prove the D-Bus read path now includes `xnix.runtime.owner_read_dispatch.v1` payload evidence without enabling writes, starting a production owner, or mutating the host root.
- Expanded the Claude Code empty-domain handoff document with an implementation-evidence scale and branch completion template so contract-only domains can be assigned without encouraging more preview-only work.

### Changed

- Updated Runtime contract and layout gates so the C D-Bus bridge is tracked as a low-level transport boundary that can consume Go Runtime owner business payloads.

## [0.2.220] - 2026-07-16

### Added

- Added Runtime owner smoke-batch JSONL records that render the full owner read dispatch table plus every disabled write-method response as ordered restricted-session evidence.
- Added `xnix-runtime-owner --smoke-batch`, giving the future D-Bus owner event loop a batch-call harness without claiming a bus name, starting a service, enabling writes, requiring network, mutating the host root, or exposing backend details.
- Added Go and restricted-session smoke coverage for smoke-batch read coverage, write denial coverage, ordering, dispatch readiness, and disabled side-effect flags.

### Changed

- Updated Runtime owner candidate smoke evidence so owner smoke now verifies candidate JSON, single read dispatch, lifecycle JSONL, and full read/write smoke-batch output.

## [0.2.219] - 2026-07-16

### Added

- Extended Runtime owner in-process read dispatch to cover every D-Bus read-only Runtime method while preserving owner-local readiness probes.
- Added Go coverage that renders every supported owner read dispatch payload through the shared owner handler table without starting an event loop, claiming a bus name, enabling write methods, requiring network, mutating the host root, or exposing backend details.

### Changed

- Updated Runtime contract drift reporting to verify full owner read dispatch coverage for all D-Bus read-only methods plus the owner-local method group.
- Updated Runtime owner dispatch evidence to keep the future D-Bus event loop path aligned with Go-native handler coverage.

## [0.2.218] - 2026-07-16

### Added

- Added Runtime owner lifecycle events under `internal/runtime/owner`, projecting smoke-owner startup, route-table, readiness, and shutdown milestones as safe JSONL records.
- Added `xnix-runtime-owner --lifecycle-log`, which renders the owner smoke lifecycle without starting an event loop, claiming a D-Bus name, enabling write methods, requiring network, mutating the host root, or exposing backend details.
- Added Go and restricted-session smoke coverage for lifecycle JSONL ordering, route counts, shutdown reason, and disabled side-effect flags.

### Changed

- Updated layout verification to require Runtime owner lifecycle sources, tests, CLI routing, and restricted-session smoke coverage.
- Expanded the mainline Claude Code handoff plan with copyable prompts for Portal/snapshot safety, KDE activation materialization, execution transactions, diagnostics/repair/AI, and atomic KDE/QEMU acceptance packages.

## [0.2.217] - 2026-07-16

### Changed

- Extended `scripts/implementation_evidence_report.rb` with a machine-readable first-wave dispatch plan for `M1`, `M2`, `M3`, and `M8`, including suggested branch names, current evidence status, rationale, and minimal mergeable outcomes.
- Updated Markdown evidence output to include a First-Wave Dispatch section so Claude Code handoffs can be selected from the repository evidence report directly.
- Updated layout and evidence-report tests to guard the next-dispatch package list and keep Docker, QEMU, network, privileged containers, backend launch, and host-root mutation disabled for default reporting.

## [0.2.216] - 2026-07-16

### Changed

- Extended `scripts/implementation_evidence_report.rb` so every implementation domain now reports both its legacy package id and its `M1`-through-`M9` mainline package id from `docs/claude-code-mainline-implementation-plan.md`.
- Updated JSON and Markdown implementation evidence output to expose the mainline handoff document, prove the mainline plan is present, and show mainline package mappings for Claude Code dispatch.
- Updated layout and evidence-report tests to guard the new mainline implementation-plan mapping without requiring Docker, QEMU, network, privileged containers, backend launch, or host-root access.

## [0.2.215] - 2026-07-16

### Added

- Added a `diagnostic_history_route` object to the Go KDE Compatibility Center section detail read model, giving KDE a safe next-hop route to `diagnostic-history-preview` for the selected application.
- Added Go and CLI/Ruby coverage proving the Diagnostics section route preserves the application id, requires a Runtime state root without exposing it, and keeps AI provider calls, file reads, request creation, permission grants, launch, execution, repair execution, host-root mutation, and backend details disabled.
- Added `docs/claude-code-mainline-implementation-plan.md`, a mainline Claude Code handoff plan that splits the remaining contract-heavy work into independently reviewable implementation packages with explicit safety constraints and acceptance tests.

### Changed

- Updated layout verification to require the Diagnostics detail route, `GetDiagnosticHistoryPreview`, state-root safety flags, and the diagnostic history read model binding.
- Documented the recommended first Claude Code dispatch wave around the Runtime owner service, recipe/artifact trust pipeline, environment lifecycle state, and developer evidence harness.

## [0.2.214] - 2026-07-16

### Changed

- Routed the Go KDE Compatibility Center Diagnostics section to `GetDiagnostics` and the `diagnostic-history-preview` read model so the page consumes Runtime-owned diagnostic state and history first.
- Kept the Dolphin AI-safe file analysis link available from the Diagnostics section while leaving AI provider calls, file-content reads, request creation, permission grants, execution, backend launch, and host-root mutation disabled.
- Updated layout verification and Go/Ruby tests to require the Diagnostics section to stay bound to Runtime diagnostics/history rather than treating AI diagnostic input as the primary section model.

## [0.2.213] - 2026-07-16

### Added

- Added a Go Runtime diagnostic history preview under `internal/runtime/appidentity`, projecting persisted diagnostic run history into a KDE/Compatibility Center-safe read model with latest run state, repair issue summary, outcome counts, and disabled action flags.
- Added `xnix-runtime-go diagnostic-history-preview`, a constrained CLI that reads diagnostic run history from a controlled Runtime state root and renders a user-visible Compatibility Center card projection without exposing state-root paths, file contents, backend details, AI provider data, or executable repair actions.
- Added Go unit and CLI coverage for diagnostic history preview safety flags, latest-run projection, Compatibility Center history state mapping, local-path non-exposure, and disabled AI/backend/repair gates.

### Changed

- Updated layout verification to require the diagnostic history preview source, CLI command, tests, schema, KDE read-model projection, and disabled side-effect gates.

## [0.2.212] - 2026-07-16

### Added

- Added a Go Runtime diagnostic run history summary under `internal/runtime/diagnostics`, allowing Compatibility Center consumers to read persisted diagnostic run counts, latest run state, failing signal IDs, and review-first repair issue summaries from the controlled Runtime state root.
- Added `xnix-runtime-go diagnostic-run-history`, a constrained CLI that reads diagnostic run records with an optional application filter while keeping state-root paths, file contents, backend details, AI provider calls, auto-repair, backend launch, network access, and host-root mutation disabled.
- Added Go unit and CLI coverage for diagnostic history filtering, outcome counts, latest-run selection, safe relative receipt paths, required state-root arguments, and disabled AI/backend/repair side-effect flags.

### Changed

- Updated layout verification to require diagnostic run history source, CLI command, tests, schema, KDE-safe record projections, and disabled AI/backend/repair gates.
- Updated product documentation so Runtime diagnostics now include both state-root run recording and state-root diagnostic history reads.

## [0.2.211] - 2026-07-16

### Added

- Added a Go Runtime diagnostic run record store under `internal/runtime/diagnostics`, persisting fixture-driven diagnostic run receipts and latest test results inside a controlled Runtime state root.
- Added `xnix-runtime-go diagnostic-run-record`, a constrained CLI that reads a diagnostic fixture, records safe test signals, derives review-first repair recommendations, builds privacy-filtered AI diagnostic input, and keeps real AI provider calls, auto-repair, backend launch, root-path exposure, and host-root mutation disabled.
- Added Go unit and CLI coverage for diagnostic run receipt persistence, relative-path records, latest-result persistence, unsafe root rejection, unsafe run-id rejection, required CLI arguments, and disabled AI/backend/repair side-effect flags.

### Changed

- Updated implementation evidence reporting so the diagnostics, repair, and AI boundary domain now reports state-root implementation evidence.
- Updated layout verification to require diagnostic run record source, CLI command, tests, schema, state-root safety flags, fixture-path safety flags, and disabled AI/backend/repair gates.

## [0.2.210] - 2026-07-16

### Added

- Added a Go Runtime artifact stage receipt under `internal/runtime/artifact`, persisting local fixture artifact staging results inside the controlled cache root after digest verification.
- Added `xnix-runtime-go artifact-stage-record`, a constrained CLI that reads an artifact manifest, stages artifacts from a local fixture source, writes a receipt under the cache root, and keeps network fetch, package-manager calls, backend launch, root-path exposure, and host-root mutation disabled.
- Added Go unit and CLI coverage for artifact stage receipts, digest mismatch blocking, explicit cache/fixture root requirements, filesystem-root cache rejection, and disabled side-effect flags.
- Added Claude Code empty-domain handoff guidance that records current baseline evidence and provides a one-package assignment template for contract-to-implementation work.

### Changed

- Updated implementation evidence reporting so the recipe and artifact trust pipeline domain now reports state-root implementation evidence.
- Updated layout verification to require the artifact stage receipt package, CLI command, tests, schema, cache-root safety flags, and disabled network/package-manager/backend gates.

## [0.2.209] - 2026-07-16

### Added

- Added a Go Runtime execution ledger under `internal/runtime/execution`, persisting blocked-by-default execution transaction records inside a caller-supplied Runtime state root.
- Added `xnix-runtime-go execution-ledger-record`, a constrained CLI that creates, reviews, preflights, and records an execution transaction without launching backends, granting permissions, exposing the state root path, or mutating the host root.
- Added Go unit and CLI coverage for execution ledger persistence, relative-path receipts, sorted listing, unsafe root and request-id rejection, and disabled launch/backend/permission gates.

### Changed

- Updated implementation evidence reporting so the execution transaction ledger domain now reports state-root implementation evidence.
- Updated layout verification to require the execution ledger package, CLI command, tests, schema, state-root safety flags, and disabled launch/backend gates.

## [0.2.208] - 2026-07-16

### Added

- Added `scripts/implementation_evidence_report.rb`, a local JSON/Markdown evidence report that classifies the main empty-domain implementation packages by current implementation depth and flags orphan Runtime read methods.
- Added targeted tests for the implementation evidence report schema, domain coverage, production-gate evidence, disabled write-method flags, and Markdown output.
- Added `docs/claude-code-empty-domain-implementation-packages.md`, a large-package Claude Code handoff guide for contract-heavy domains that still need durable implementation evidence.

### Changed

- Updated the Claude Code handoff index documents to reference the new empty-domain implementation package guide.
- Updated layout verification to require the implementation evidence report, its tests, the empty-domain guide, stable domain identifiers, evidence statuses, and safety flags.

## [0.2.207] - 2026-07-16

### Added

- Added a Go Runtime desktop activation staging writer under `internal/runtime/activation`, allowing KDE desktop-entry, Dolphin service-menu, MIME association, manifest, and rollback receipt files to be materialized inside an explicit staging root.
- Added `xnix-runtime-go desktop-activation-stage`, a constrained CLI that requires `--staging-root`, refuses filesystem-root staging, refuses existing-file overwrites, and emits a safe JSON receipt without exposing the staging root path.
- Added Go unit and CLI coverage for staged file digests, rollback receipts, conflict rejection, blocked production mode, and disabled launch/backend/host-root gates.

### Changed

- Updated layout verification to require the Go activation staging writer, CLI command, tests, relative-path materialization, and staging safety flags.

## [0.2.206] - 2026-07-16

### Added

- Added `scripts/runtime_contract_drift_report.rb`, a read-only Runtime contract drift gate that compares the D-Bus XML contract, Go method parity list, Go owner route manifest, owner read dispatch methods, Runtime dispatch, D-Bus client, smoke adapter, session smoke, and Go CLI route commands.
- Added JSON and Markdown contract drift report output plus targeted tests for the stable check set and safety flags.

### Changed

- Updated layout verification to require the Runtime contract drift report, its tests, JSON/Markdown output support, the ten drift checks, and disabled write-method safety fields.

## [0.2.205] - 2026-07-16

### Added

- Added Go Runtime owner read dispatch previews for the Runtime owner method group, allowing `xnix-runtime-owner --dispatch-read <method> [args...]` to render selected read-only owner payloads in-process without starting a D-Bus event loop, claiming session or production bus ownership, enabling write methods, starting services, requiring network, mutating the host root, or exposing backend details.
- Extended the Runtime owner candidate restricted session smoke to verify `GetRuntimeWriteGate` through the new read dispatch envelope.
- Added `docs/claude-code-contract-gap-work-packages.md`, a branch-sized Claude Code backlog for contract-heavy Runtime, KDE, Portal, snapshot, execution, AI diagnostics, and image acceptance gaps.

### Changed

- Updated layout verification and Runtime owner CLI tests to cover the read-only dispatch schema, safety flags, and nested write-gate payload.

## [0.2.204] - 2026-07-16

### Added

- Added `scripts/runtime_owner_candidate_smoke.rb`, a restricted session-bus smoke that runs the Go Runtime owner candidate in `smoke-owner` mode, verifies full Go read-only route coverage, checks deterministic disabled write-method responses, and confirms the candidate does not claim session or production bus ownership, start services, require network, mutate the host root, require privileged containers, or expose backend details.
- Added the `runtime-owner-candidate-smoke` constrained container command and installed `xnix-runtime-owner` into the Docker smoke image for offline read-only Colima validation.

### Changed

- Updated the container command unit tests and layout verification to include the Runtime owner candidate smoke gate.

## [0.2.203] - 2026-07-16

### Added

- Added `xnix-runtime-owner` and the Go Runtime owner candidate package, exposing a safe smoke-owner/readiness JSON shape with full Go read-only route coverage and deterministic disabled write-method responses while keeping the production D-Bus loop, bus claims, service starts, host-root mutation, backend launch, and KDE ownership disabled.
- Added a coarse Claude Code contract implementation handoff guide that groups the remaining contract-heavy Runtime, KDE, execution, AI, image, and quality-gate domains into larger independently assignable packages.

### Changed

- Updated Runtime owner process readiness to recognize the Go owner candidate as present while continuing to report the production D-Bus event loop and production ownership as pending.
- Linked the new contract handoff guide from the existing Claude Code backlog documents.

## [0.2.202] - 2026-07-15

### Added

- Added Go Runtime `desktop-activation-manifest-preview` for `GetDesktopActivationManifest`, aggregating the KDE desktop activation bundle and seven first-release KDE entry points into a read-only manifest without writing desktop files, changing MIME defaults, persisting settings, sending notifications, activating task-manager entries, applying KWin rules, enabling tray bridges, launching backends, exposing backend details, or mutating the host root.

### Changed

- Migrated `GetDesktopActivationManifest` in `runtime-owner-route-manifest-preview` from the C Runtime adapter bucket to a native Go preview route, completing native Go owner preview coverage for all 57 read-only Runtime routes while production D-Bus ownership remains gated.
- Updated Runtime owner readiness and KDE-first route smoke expectations for the `57 total, 57 Go-routed, 0 C-backed, 0 Ruby legacy` baseline.

## [0.2.201] - 2026-07-15

### Added

- Added Go Runtime `compatibility-install-preview` for `GetCompatibilityInstallPlan`, aggregating artifact manifest, acquisition preflight, package source, state-root, and recipe install-gate readiness without downloading artifacts, installing packages, staging desktop activation, launching backends, exposing backend details, or mutating the host root.
- Added `docs/claude-code-independent-implementation-briefs.md` as a prompt-ready set of large, independent Claude Code implementation briefs for the contract-heavy Runtime, KDE, execution, AI, and product-image domains.

### Changed

- Migrated `GetCompatibilityInstallPlan` in `runtime-owner-route-manifest-preview` from the C Runtime adapter bucket to a native Go preview route, leaving only the desktop activation manifest as C-backed route evidence.
- Updated the Claude Code open-domain work-package map and product overview to point at the new prompt-ready implementation briefs.

## [0.2.200] - 2026-07-15

### Added

- Added Go Runtime `task-manager-identity-preview` and `kwin-window-rule-preview` for `GetTaskManagerIdentityPlan` and `GetKWinWindowRulePlan`, deriving KDE task-manager grouping, pinning, restore hints, and KWin identity/layout hints from the shared Go window identity model without observing windows, applying KWin rules, activating task-manager entries, enabling launch or execution, exposing backend details, or mutating the host root.
- Added `docs/claude-code-open-domain-work-packages.md` as a higher-level Claude Code work-package backlog for the mostly contract-only Runtime, KDE, execution, AI, and image domains.

### Changed

- Migrated the task-manager identity and KWin window-rule route group in `runtime-owner-route-manifest-preview` from C Runtime adapter routes to native Go preview routes, reducing the remaining C-backed read-only owner route count.
- Increased the constrained QEMU boot-system timeout used by the full smoke gate so the serial-login marker can appear on slower containerized boots.

## [0.2.199] - 2026-07-15

### Added

- Added Go Runtime `kde-integration-status-preview`, `kde-shell-integration-preview`, and `kde-application-surface-preview` for `GetKDEIntegrationStatus`, `GetKDEShellIntegrationPlan`, and `GetKDEApplicationSurfacePlan`, rendering KDE-first shell status, replaceable-shell integration, and per-application KDE surface plans without forking Plasma, writing shell configuration, enabling component activation, starting compatibility profiles, exposing implementation details, or mutating the host root.

### Changed

- Migrated the KDE shell status, shell integration, and application surface route group in `runtime-owner-route-manifest-preview` from C Runtime adapter routes to native Go preview routes, reducing the remaining C-backed read-only owner route count.

## [0.2.198] - 2026-07-15

### Added

- Added a constrained Go content-addressed snapshot store for Runtime-owned state roots, with create, list, verify, and rollback operations confined to the configured state root and guarded against host-root mutation.
- Added a Go Runtime Portal request broker package with explicit request creation, fake test transport, recoverable failure states, permission-state tracking, and deterministic request handles for future XDG Desktop Portal integration.

### Changed

- Aligned repository version metadata and product documentation with the current P7/P8 Runtime implementation commits while keeping production backend launch, real Portal calls, privileged operations, and host-root writes disabled.

## [0.2.197] - 2026-07-15

### Added

- Added Go Runtime `state-root-preview`, `snapshot-plan-preview`, and `portal-access-policy-preview` for `GetApplicationStateRoot`, `GetSnapshotPlan`, and `GetPortalAccessPolicy`, rendering Runtime-owned state storage, snapshot/restore planning, and XDG Desktop Portal policy without creating directories, creating snapshots, creating Portal request objects, granting permissions, requiring network access, exposing backend details, or mutating the host root.

### Changed

- Migrated the state-root, snapshot, and Portal access route group in `runtime-owner-route-manifest-preview` from C Runtime adapter routes to native Go preview routes, reducing the remaining C-backed read-only owner route count.
- Extended the KDE-first presence smoke to inspect the Go Runtime safety substrate previews and assert that user documents, direct desktop access, snapshot creation, restore execution, request creation, and host-root mutation remain disabled.

## [0.2.196] - 2026-07-15

### Added

- Added Go Runtime `ai-diagnostic-input-preview`, `ai-diagnostic-recommendation-preview`, and `ai-repair-approval-gate-preview` for `GetAIDiagnosticInput`, `GetAIDiagnosticRecommendation`, and `GetAIRepairApprovalGate`, composing Runtime-safe recipe, run-plan, test-result, and repair-plan metadata without calling an AI provider, reading user files, creating request objects, executing repairs, requiring network access, exposing backend details, or mutating the host root.

### Changed

- Migrated the AI diagnostics route group in `runtime-owner-route-manifest-preview` from C Runtime adapter routes to native Go preview routes, reducing the remaining C-backed read-only owner route count.
- Extended the KDE-first presence smoke to inspect the Go AI diagnostics previews and assert that AI explanation, recommendation, and repair approval flows remain review-only and non-executing.

## [0.2.195] - 2026-07-15

### Added

- Added Go Runtime `repair-plan-preview`, `test-plan-preview`, `test-result-preview`, and `runtime-write-gate-preview` for `GetRepairPlan`, `GetTestPlan`, `GetTestResult`, and `GetRuntimeWriteGate`, rendering repair, test, and write-gate planning from Runtime-owned models without executing repairs, starting tests, creating request objects, enabling write dispatch, requiring network access, exposing backend details, or mutating the host root.

### Changed

- Migrated `GetRepairPlan`, `GetTestPlan`, `GetTestResult`, and `GetRuntimeWriteGate` in `runtime-owner-route-manifest-preview` from C Runtime adapter routes to native Go preview routes, reducing the remaining C-backed read-only owner route count.
- Extended the KDE-first presence smoke to execute the Go `runtime-write-gate-preview --method Launch` path and assert that write dispatch, request creation, execution, network access, privileged containers, and host-root mutation remain disabled.

## [0.2.194] - 2026-07-15

### Added

- Added Go Runtime `backend-capability-matrix-preview` and `backend-lifecycle-preview` for `GetBackendCapabilityMatrix` and `GetBackendLifecycle`, rendering the desktop-safe backend capability matrix (local and isolated profiles across seven Runtime-owned capabilities) and the blocked backend lifecycle stages from registry recipes without enabling backend selection, starting any backend process, creating state roots or snapshots, requiring network access, exposing raw commands, or mutating the host root.
- Added `scripts/kde_first_presence_smoke.rb` and its unit coverage as a first-milestone KDE presence smoke harness, aggregating Go Runtime previews for one recipe across the seven KDE entry points while enforcing disabled execution, disabled backend launch, project-local Go cache use, no host-root mutation, no privileged containers, no network requirement, and no raw command or file-content exposure.

### Changed

- Migrated `GetBackendCapabilityMatrix` and `GetBackendLifecycle` in `runtime-owner-route-manifest-preview` from the C Runtime adapter bucket to native Go preview routes, completing the recommended backend group (`GetBackendEnvironmentPlan` was already native Go) and further lowering the remaining C-backed read-only owner route count.
- Updated CLI coverage, unit tests, layout verification, and product documentation for the Go backend capability and lifecycle migration plus the KDE-first presence smoke acceptance harness.

## [0.2.193] - 2026-07-15

### Added

- Added Go Runtime `package-source-preview`, `acquisition-preflight-preview`, and `artifact-manifest-preview` for `GetCompatibilityPackageSource`, `GetCompatibilityAcquisitionPreflight`, and `GetCompatibilityArtifactManifest`, rendering the desktop-safe compatibility package acquisition chain (Runtime-owned source channels, acquisition checks, and artifact groups) from registry recipes without downloading artifacts, installing packages, requiring network access, exposing raw commands or cache paths, or mutating the host root.

### Changed

- Migrated `GetCompatibilityPackageSource`, `GetCompatibilityAcquisitionPreflight`, and `GetCompatibilityArtifactManifest` in `runtime-owner-route-manifest-preview` from the C Runtime adapter bucket to native Go preview routes, advancing the acquisition slice while keeping the remaining C-backed policy routes explicit.
- Updated CLI coverage, unit tests, layout verification, and product documentation for the Go package acquisition migration.

## [0.2.192] - 2026-07-15

### Added

- Added Go Runtime `run-plan-preview` for `GetRunPlan`, rendering desktop-safe compatibility run planning from registry recipes without creating execution requests, launching profiles, granting permissions, exposing raw commands, or mutating the host root.

### Changed

- Migrated `GetRunPlan` in `runtime-owner-route-manifest-preview` from the C Runtime adapter bucket to a native Go preview route, completing the Go-routed run/launch/readiness planning slice while keeping remaining C-backed policy routes explicit.
- Updated CLI coverage, unit tests, layout verification, and product documentation for the Go run plan migration.

## [0.2.191] - 2026-07-15

### Added

- Added Go Runtime `engine-catalog-preview` for `GetEngineCatalog`, exposing Automatic, Local compatibility, and Isolated compatibility choices as desktop-safe Runtime metadata without installing engines, launching backends, persisting selection, or exposing raw commands.
- Added a Claude Code implementation package guide that splits contract-only Runtime domains into independently scoped work packages with boundaries, acceptance criteria, and required tests.

### Changed

- Migrated `GetEngineCatalog` in `runtime-owner-route-manifest-preview` from the C Runtime adapter bucket to a native Go preview route, reducing remaining C-backed read-only owner routes.
- Updated CLI coverage, unit tests, layout verification, and product documentation for the Go engine catalog migration.

## [0.2.190] - 2026-07-15

### Added

- Added Go Runtime `applications-preview` and `application-preview` for `ListApplications` and `GetApplication`, rendering registry-verified compatibility applications as normal KDE/Linux application identities without exposing backend terminology, raw executable details, or host paths.
- Added the **KDE Plasma atomic image pipeline** (Delivery Sequence step 1): `image/kinoite/manifest.json` as the single source of truth for a Fedora Kinoite-compatible KDE Plasma 6 image, `lib/xnix/image/kde_image.rb` to validate the manifest against on-disk artifacts and render a bootc `Containerfile` from it, and `lib/xnix/image/boot_smoke.rb` to model the graphical-login boot check.
- Added `scripts/build_kde_image.rb` (validate + Containerfile drift guard + honest podman/buildah toolchain detection; never reports an unperformed build) and `scripts/boot_kde_image.rb` (QEMU disk-image boot smoke scanning for manifest-declared serial markers, loopback-only networking).
- Added image config overlays (SDDM Plasma-Wayland greeter, xdg-desktop-portal KDE backend, os-release branding, systemd presets), `test/test_kde_image.rb`, layout verification coverage, and `docs/kde-image-pipeline.md`.
- Added the **disk-image production step** linking the container build to the boot smoke: `image/kinoite/disk-config.json` (bootc-image-builder source of truth, cross-checked against the image manifest name), `lib/xnix/image/disk_build.rb` to validate it and render the exact privileged `bootc-image-builder` podman command per output type (qcow2/raw/iso), and `scripts/build_kde_disk.rb` (validate + honest podman toolchain detection; writes the bib blueprint and never fakes a build).
- Wired `scripts/boot_kde_image.rb` to default its `--disk` to the disk-build qcow2 output so the pipeline connects end-to-end (manifest → Containerfile → container → qcow2 → boot smoke), added `test/test_kde_disk.rb`, and extended layout verification. The blueprint sets `console=ttyS0` so the boot-smoke serial markers surface. The image keeps `runtime_owned: true` / `kde_policy_owner: false` — it lays down KDE entry points and the Runtime's D-Bus activation without embedding compatibility policy.

### Changed

- Migrated `ListApplications` and `GetApplication` in `runtime-owner-route-manifest-preview` from the C Runtime adapter bucket to native Go preview routes, increasing Go-routed read-only owner methods while keeping C-backed policy routes explicit.
- Updated owner route summaries, CLI coverage, unit tests, layout verification, and product documentation for the Go application catalog migration.

## [0.2.189] - 2026-07-15

### Added

- Added Go Runtime `diagnostics-preview` for `GetDiagnostics` to aggregate Runtime-safe application status, execution readiness, launch gates, KDE action queues, Compatibility Center summaries, KDE section coverage, and AI diagnostics safety without starting execution, reading file contents, calling an AI provider, or exposing backend details.

### Changed

- Migrated `GetDiagnostics` in `runtime-owner-route-manifest-preview` from the Ruby legacy dispatch bucket to a native Go preview route, reducing legacy Runtime owner routes to zero while leaving C Runtime adapter migration pending.
- Wired CLI, unit, and layout verification coverage for the Go diagnostics route and updated owner-readiness route summaries.

## [0.2.188] - 2026-07-15

### Added

- Added Go Runtime `runtime-owner-route-manifest-preview` to classify every read-only Runtime method as a native Go preview route, C Runtime policy route requiring a Go owner adapter, or legacy Runtime dispatch route requiring migration before production ownership.
- Wired owner route evidence into `runtime-owner-readiness-preview` so production owner readiness now gates on read-only owner route migration instead of relying on method parity alone.
- Added Go unit, CLI, and layout verification coverage for route counts, route statuses, legacy dispatch detection, owner-readiness aggregation, and host-safe preview flags.

## [0.2.187] - 2026-07-15

### Added

- Added Go Runtime `runtime-owner-process-preview` to verify the packaged Runtime owner activation chain, report the current Ruby wrapper as a transition entrypoint, and keep the production Go owner loop pending without starting services or claiming the D-Bus name.
- Wired owner process evidence into `runtime-owner-readiness-preview` so the long-running Runtime owner gate is derived from process checks instead of a hard-coded pending state.

### Changed

- Split Runtime-owner Go CLI runners into `runtime_owner_commands.go` so future owner commands do not continue growing the main CLI file toward the file-size guardrail.

## [0.2.186] - 2026-07-15

### Added

- Added Go Runtime `runtime-owner-recipe-trust-preview` to read the Runtime recipe registry, verify recipe SHA-256 digests, summarize signature status counts, and gate production owner promotion on production-signed recipes without enabling installation or host mutation.
- Wired owner recipe trust evidence into `runtime-owner-readiness-preview` so production owner readiness now reports recipe trust as pass, pending, or blocked from registry evidence instead of a hard-coded pending gate.
- Added Go unit, CLI, and layout verification coverage for development-only, signed, and missing-registry trust states plus owner-readiness aggregation.

## [0.2.185] - 2026-07-15

### Added

- Added Go Runtime `runtime-owner-readiness-preview` to aggregate service binding, live owner gates, owner smoke planning, and method parity evidence into one production-owner readiness view without starting services or claiming the production D-Bus name.
- Added Go unit, CLI, and layout verification coverage for owner readiness checks, pending production owner gates, write-method closure, KDE non-ownership, host safety, and desktop-safe reporting.

## [0.2.184] - 2026-07-15

### Added

- Added Go Runtime `runtime-method-parity-manifest-preview` to verify read-only Runtime D-Bus method coverage across the XML contract, Runtime dispatch, KDE-facing D-Bus client, split smoke adapter files, and session smoke script without enabling write dispatch.
- Added Go unit, CLI, and layout verification coverage for method parity readiness, gated write methods, missing-source blocking, and desktop-safe reporting before production owner smoke.

## [0.2.183] - 2026-07-15

### Added

- Added Go Runtime `runtime-owner-smoke-plan-preview` to derive the future production owner smoke sequence from the Go live owner gate while keeping smoke execution, host system-service startup, production bus claims, and backend launch disabled.
- Added Go unit, CLI, and layout verification coverage for owner smoke planning, including activation validation, packaged owner startup, bus-name assertion, read-only method parity, write-method rejection, smoke-adapter boundary checks, and KDE-safe summary reporting.

## [0.2.182] - 2026-07-15

### Added

- Added Go Runtime `runtime-live-owner-gate-preview` to derive production D-Bus owner transition gates from the Go service binding preview while keeping long-running owner startup, bus-name acquisition, method parity, and production recipe trust pending.
- Added Go unit, CLI, and layout verification coverage proving the live-owner gate cannot start services, claim the production bus, treat the smoke adapter as production, let KDE own Runtime policy, mutate the host root, or expose backend details.

## [0.2.181] - 2026-07-15

### Added

- Added Go Runtime `runtime-service-binding-preview` to report D-Bus activation, systemd hardening, libexec wrapper, contract, smoke-adapter, and live-owner readiness without starting services, claiming the production bus, mutating the host root, requiring network access, or exposing backend details.
- Added Go unit, CLI, and layout verification coverage for the service binding preview so KDE can consume a Go-backed read model while the existing Ruby and C service binding surfaces remain compatible.

## [0.2.180] - 2026-07-15

### Changed

- Split Portal, package-source, backend, execution, settings, action queue, Compatibility Center summary, Runtime service binding, owner-gate, method-parity, and write-gate builders out of the Linux D-Bus smoke adapter into `runtime/dbus/xnix_compatd_runtime_models.inc`, reducing `xnix_compatd_smoke.c` below the 2,000-line review threshold while preserving the existing session-bus read-model behavior.
- Updated layout verification, Runtime method parity, activation packaging, activation install tests, and D-Bus smoke-script tests so the dispatcher, introspection include, KDE Center include, and Runtime read-model include are all treated as one smoke adapter contract.

## [0.2.179] - 2026-07-15

### Changed

- Split KDE Compatibility Center page, section, and section-detail builders out of the Linux D-Bus smoke adapter into `runtime/dbus/xnix_compatd_kde_center.inc`, reducing `xnix_compatd_smoke.c` away from the 3,000-line hard limit while preserving the existing session-bus read-model behavior.
- Updated layout verification and D-Bus smoke-script unit coverage so Runtime method parity checks include the new KDE Center smoke adapter include alongside the dispatcher and introspection include.

## [0.2.178] - 2026-07-15

### Added

- Added a Go-backed notification plan snapshot to KDE Compatibility Center page previews so Plasma can show approval-required Runtime events, urgency, categories, and safe notification actions without sending notifications, enabling action execution, enabling repair execution, persisting settings, mutating the host root, or exposing backend details.
- Added a read-only Notifications section to KDE Compatibility Center page navigation and section-detail routing, backed by `GetNotificationPlan`, with Ruby daemon, C smoke adapter, D-Bus client, Go, and Ruby coverage updated for the twelve-section page model.

## [0.2.177] - 2026-07-15

### Added

- Added a Go-backed tray status snapshot to KDE Compatibility Center page previews so Plasma can show registered compatibility applications, attention state, tray bridge state, and safe navigation actions without enabling live tray bridging, persisting bridge configuration, mutating the host root, or exposing backend details.
- Added a read-only Tray section to KDE Compatibility Center page navigation and section-detail routing, backed by `GetTrayStatus`, with Ruby daemon, C smoke adapter, D-Bus client, Go, and Ruby coverage updated for the eleven-section page model.

## [0.2.176] - 2026-07-15

### Added

- Added a Go-backed file association snapshot to KDE Compatibility Center page previews so Plasma can explain MIME associations, Dolphin file-open routing, and Portal requirements without writing MIME defaults, reading host files directly, granting permissions, or starting execution.
- Added a read-only Files section to KDE Compatibility Center page navigation and section-detail routing, backed by `GetFileAssociationPlan`, with Ruby daemon, C smoke adapter, D-Bus client, Go, and Ruby coverage updated for the ten-section page model.

## [0.2.175] - 2026-07-15

### Added

- Added a Go-backed window identity snapshot to KDE Compatibility Center page previews so Plasma can explain task-manager grouping, pinning, restore, switcher visibility, and KWin identity hints without activating task-manager entries, applying KWin rules, starting execution, or exposing backend details.
- Added a read-only Window section to KDE Compatibility Center page navigation and section-detail routing, backed by `GetTaskManagerIdentityPlan`, with Ruby daemon, C smoke adapter, D-Bus client, Go, and Ruby coverage updated for the nine-section page model.

## [0.2.174] - 2026-07-15

### Added

- Added a read-only Launch section to KDE Compatibility Center page navigation and section-detail routing, backed by `GetLaunchIntent`, so Plasma can explain managed desktop-entry launches without treating the page as a Launch approval surface.
- Added Go, Ruby daemon, C smoke adapter, D-Bus client, and Ruby coverage proving the Launch section remains navigation-only, non-executing, and blocked from request creation, permission grants, host-root mutation, and backend detail exposure.

## [0.2.173] - 2026-07-15

### Added

- Added a Go-backed launch intent snapshot to the KDE Compatibility Center page preview so Plasma can explain managed desktop-entry launches without creating Launch requests, request objects, permission grants, or backend processes.
- Added Ruby daemon, C smoke adapter, D-Bus client, Go, and Ruby coverage proving the page-level launch intent remains Runtime-owned, non-executing, and hidden from backend implementation details.

## [0.2.172] - 2026-07-15

### Added

- Added a Go-backed execution readiness snapshot to the KDE Compatibility Center page preview so Plasma can show desktop-entry visibility, launch-readiness gates, blocked execution state, and closed launch/request/backend-binding gates without starting a compatibility backend.
- Added a read-only Execution section to KDE Compatibility Center page sections and section-detail routing, backed by `GetExecutionReadiness`, with Ruby daemon, C smoke adapter, D-Bus client, Go, and Ruby coverage updated for the seven-section page model.

## [0.2.171] - 2026-07-15

### Added

- Added a Go-backed backend selection snapshot to the KDE Compatibility Center page preview so Plasma can show the Runtime-recommended compatibility profile without committing selection, creating environments, launching backends, mutating the host root, or exposing backend details.
- Added a read-only Backend section to KDE Compatibility Center page sections and section-detail routing, backed by `GetBackendSelectionPlan`, with Ruby daemon, C smoke adapter, D-Bus client, Go, and Ruby coverage updated for the six-section page model.

## [0.2.170] - 2026-07-15

### Added

- Added a Go-backed desktop activation status snapshot to the KDE Compatibility Center page preview so Plasma can show activation readiness, the Runtime status method, status counts, and closed commit/launch/host-mutation gates directly on the application page.
- Added a read-only Activation section to KDE Compatibility Center page sections and section-detail routing, backed by `GetDesktopActivationStatus`, with Ruby daemon, C smoke adapter, D-Bus client, Go, and Ruby coverage updated for the five-section page model.

## [0.2.169] - 2026-07-15

### Added

- Added `GetDesktopActivationStatus(application_id, mode)` to the Runtime D-Bus contract, Ruby daemon dispatch, KDE-facing D-Bus client, C smoke adapter, C Runtime method-parity policy, and Linux session smoke so KDE can read the Go-owned desktop activation status through the standard Runtime boundary.
- Added D-Bus client, daemon, contract, C Runtime, and session-smoke coverage proving the activation status read model remains Go-backed, read-only, non-committing, non-launching, and blocked from host-root mutation, network access, privileged containers, and backend detail exposure.

## [0.2.168] - 2026-07-15

### Changed

- Split the Linux D-Bus smoke adapter introspection XML into `runtime/dbus/xnix_compatd_introspection.inc`, reducing `xnix_compatd_smoke.c` below the repository line-count warning zone while preserving the existing Runtime session-bus contract behavior.
- Updated Runtime method parity, layout verification, activation packaging, and D-Bus smoke-script unit checks so the split smoke adapter still proves method coverage across both the C dispatcher and introspection include.

## [0.2.167] - 2026-07-15

### Added

- Added Go Runtime `desktop-activation-status-preview` to summarize desktop activation transaction readiness, staging readiness, commit gates, KDE surface readiness, blocked reasons, and next safe actions without opening commit, launch, backend, network, privileged-container, host-root, or desktop-write gates.
- Added Go unit and CLI coverage proving the activation status preview derives from the Go-owned transaction preview, stays safe for KDE Compatibility Center display, and remains separate from the future D-Bus method until the oversized C smoke adapter is split.

## [0.2.166] - 2026-07-15

### Added

- Added `GetDesktopActivationTransactionPreview` to the Runtime D-Bus contract, method parity manifest, Ruby daemon dispatch, KDE-facing D-Bus client, C smoke adapter, session smoke, and C Runtime parity policy so KDE can consume the Go-owned desktop activation transaction read model through the standard Runtime boundary.
- Added D-Bus client, daemon, contract, C Runtime, and session-smoke coverage proving the activation transaction preview remains read-only, Go-backed, non-committing, and blocked from desktop writes, MIME writes, KDE cache refresh, launch, execution, networking, privileged containers, host-root mutation, and backend detail exposure.

## [0.2.165] - 2026-07-15

### Added

- Added Go Runtime `desktop-activation-transaction-preview` to turn desktop activation staging output into an auditable KDE activation transaction and rollback sequence without committing writes, refreshing KDE caches, exposing target paths, launching backends, or mutating the host root.
- Split oversized Go CLI tests into focused desktop activation and KDE Compatibility Center test files so the main CLI test file stays below the repository line-count limit while preserving existing coverage.

## [0.2.164] - 2026-07-15

### Added

- Added Go Runtime `desktop-activation-staging-preview` to plan the KDE activation staging file set, relative paths, modes, SHA-256 digests, activated entry points, installer command preview, and rollback receipt preview without writing files or exposing host paths.
- Added Go unit and CLI coverage proving development staging can produce an auditable planned file set while production staging remains blocked for development-only registries, with desktop writes, MIME writes, manifest writes, receipt writes, settings persistence, notifications, task-manager activation, KWin rules, tray bridges, launch, execution, networking, privileged containers, host-root mutation, raw backend command exposure, compatibility storage exposure, and backend detail exposure disabled.

## [0.2.163] - 2026-07-15

### Added

- Added Go Runtime `desktop-activation-preflight-preview` to connect desktop activation bundle readiness, recipe trust, recipe install gates, backend binding, staging-root requirements, and host-root write gates before any KDE activation files are written.
- Added Go unit and CLI coverage proving production activation blocks development-only registries while development staging can be marked eligible without writing desktop files, MIME defaults, manifests, receipts, settings, notifications, task-manager entries, KWin rules, tray bridges, launching, networking, privileged containers, host-root mutation, raw backend command exposure, compatibility storage exposure, or backend detail exposure.

## [0.2.162] - 2026-07-15

### Added

- Added Go Runtime `backend-binding-preview` to connect backend selection, backend environment, and launch preflight gates into a Runtime-owned compatibility profile binding plan.
- Added Go unit and CLI coverage proving binding commits, persistence, execution request creation, environment creation, backend startup, state-root creation, Portal request creation, snapshot creation, launch, host-root mutation, network and privileged-container requirements, raw backend command exposure, compatibility storage exposure, and backend detail exposure stay disabled.

## [0.2.161] - 2026-07-15

### Added

- Added Go Runtime `backend-environment-preview` to plan local and isolated compatibility environments plus Portal-mediated bridge capabilities from the Runtime boundary.
- Added Go unit and CLI coverage proving environment creation, backend startup, backend binding, launch, request creation, host storage exposure, clipboard and print bridge activation, host-root mutation, privileged containers, raw backend command exposure, compatibility storage exposure, and backend detail exposure stay disabled.

## [0.2.160] - 2026-07-15

### Added

- Added Go Runtime `desktop-activation-bundle-preview` to aggregate launcher, MIME, desktop icon, task-manager, KWin, tray, notification, settings, and Compatibility Center materials for presenting a compatibility application as a normal Linux application.
- Added Go unit and CLI coverage proving the activation bundle keeps desktop writes, MIME writes, settings persistence, notifications, task-manager activation, KWin rule application, live tray bridging, backend launch, execution, host-root mutation, raw executable exposure, compatibility storage exposure, and backend details disabled.

## [0.2.159] - 2026-07-15

### Added

- Added a privacy-safe Dolphin AI analysis input summary to KDE Diagnostics section detail previews when file selections are present.
- Added Go coverage proving Diagnostics detail exposes only file count, extension, selection mode, and disclosure metadata while keeping file paths, file contents, AI provider calls, network access, permission grants, request creation, backend launch, and host-root mutation disabled.

## [0.2.158] - 2026-07-15

### Added

- Added Dolphin AI-safe file analysis metadata to the KDE Compatibility Center Diagnostics section and section-detail preview.
- Routed the Diagnostics section to the `GetAIDiagnosticInput` read model while keeping section navigation read-only and blocking AI provider calls, network access, file reads, path exposure, request creation, permission grants, execution, backend launch, and host-root mutation.

## [0.2.157] - 2026-07-15

### Added

- Added Dolphin AI analysis aggregation to KDE action card deck and KDE Compatibility Center page summaries.
- Added Go and CLI coverage proving only the Dolphin file-manager card contributes the AI-safe link and that the aggregated page/deck metadata remains navigation-only with provider calls, network access, file reads, path exposure, permission grants, request creation, backend launch, and host-root mutation disabled.

## [0.2.156] - 2026-07-15

### Added

- Added a shared KDE Dolphin AI analysis link for the file-manager entry point, entrypoint action preview, and Compatibility Center action card preview.
- Added Go coverage proving the linked Dolphin AI action is navigation-only and keeps AI provider calls, network access, file content reads, file path exposure, request-object creation, permission grants, backend launch, and host-root mutation disabled.

## [0.2.155] - 2026-07-15

### Added

- Added Go Runtime `dolphin-ai-analysis-preview` for privacy-safe Dolphin “send to AI” file-review intents.
- Added Go unit and CLI coverage proving Dolphin AI previews expose only file count and extension while keeping file paths, file contents, AI provider calls, network access, permission grants, backend launch, host-root mutation, and backend details disabled.

## [0.2.154] - 2026-07-15

### Added

- Added Go Runtime `dolphin-drop-preview` for read-only Dolphin drag-and-drop file routing through the Runtime.
- Added Go unit and CLI coverage proving Dolphin drops stay Portal-mediated and do not create request objects, grant file permissions, read files directly, launch backends, mutate the host root, or expose backend details.

## [0.2.153] - 2026-07-15

### Added

- Added Go Runtime `desktop-icon-preview` for read-only KDE desktop icon planning from the standard desktop-entry identity.
- Added the read-only `GetDesktopIconPlan` Runtime D-Bus method so KDE can show a normal desktop shortcut without copying files, persisting placement, launching backends, or exposing backend commands.
- Added daemon, D-Bus client, C Runtime core, C smoke-adapter, session-smoke, dispatch, contract, Go CLI, and method-parity coverage for desktop icon reads.

### Changed

- Increased the Runtime read-only method parity count from 54 to 55.

## [0.2.152] - 2026-07-15

### Added

- Added Go Runtime `kde-center-page-section-detail-preview` for read-only KDE Compatibility Center section detail routing.
- Added the read-only `GetKDECenterPageSectionDetail` Runtime D-Bus method so KDE can open one selected Compatibility Center section and discover its backing Runtime read model.
- Added daemon, D-Bus client, C smoke-adapter, session-smoke, dispatch, contract, and method-parity coverage for KDE center page section detail reads.

### Changed

- Increased the Runtime read-only method parity count from 53 to 54.

## [0.2.151] - 2026-07-15

### Added

- Added Go Runtime `kde-center-page-sections-preview` for read-only KDE Compatibility Center page section navigation.
- Added the read-only `GetKDECenterPageSections` Runtime D-Bus method so KDE can discover the Overview, Actions, Settings, and Diagnostics page sections and their backing Runtime read methods.
- Added daemon, D-Bus client, C smoke-adapter, session-smoke, dispatch, contract, and method-parity coverage for KDE center page section reads.

### Changed

- Increased the Runtime read-only method parity count from 52 to 53.

## [0.2.150] - 2026-07-15

### Added

- Added the read-only `GetKDECenterPage` Runtime D-Bus method so KDE can fetch a Runtime-owned Compatibility Center application page from the session bus.
- Added daemon, D-Bus client, C smoke-adapter, session-smoke, dispatch, contract, and method-parity coverage for KDE center page reads.
- Added diagnostics coverage that summarizes KDE center pages while keeping page persistence, card actions, settings persistence, grants, notifications, launch, execution, host-root mutation, and backend detail exposure disabled.

### Changed

- Increased the Runtime read-only method parity count from 51 to 52.

## [0.2.149] - 2026-07-15

### Added

- Added `xnix-runtime-go kde-center-page-preview --registry <path> --app <id> --decision <decision> [file://...]` for read-only Runtime previews of a single KDE Compatibility Center application page.
- Added Go Runtime KDE center page modeling that composes the registry Compatibility Center summary, seven-card KDE action deck, unified settings snapshot, page header, and navigation targets for Plasma rendering.
- Added Go and Ruby harness coverage that verifies KDE center page previews remain Runtime-owned, Go-backed, KDE-targeted, backend-detail-free, host-root-safe, and honest about page persistence, deck persistence, card action enablement, settings persistence, review receipts, request objects, resource grants, notifications, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved the first-release KDE Compatibility Center flow from a renderable action deck toward a renderable per-application page while keeping Runtime write and execution gates closed.

## [0.2.148] - 2026-07-15

### Added

- Added `xnix-runtime-go kde-action-card-deck-preview --registry <path> --app <id> --decision <decision> [file://...]` for read-only Runtime previews of the full KDE Compatibility Center action-card deck.
- Added Go Runtime KDE action card deck modeling that aggregates the seven first-release KDE entrypoint cards, keeps launcher first, preserves the Dolphin Portal gate, and reports waiting, deferred, rejected, navigation, and disabled-action counts for Plasma rendering.
- Added Go and Ruby harness coverage that verifies KDE action card deck previews remain Runtime-owned, Go-backed, KDE-targeted, backend-detail-free, host-root-safe, and honest about deck persistence, card persistence, card action enablement, status persistence, review receipts, queue mutation, request objects, resource grants, notifications, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved the first-release KDE Compatibility Center flow from single-card rendering toward a renderable card deck while keeping Runtime write and execution gates closed.

## [0.2.147] - 2026-07-15

### Added

- Added `xnix-runtime-go kde-action-card-preview --registry <path> --app <id> --action <id> --decision <decision> [file://...]` for read-only Runtime previews of KDE Compatibility Center action cards derived from action status.
- Added Go Runtime KDE action card modeling with user-facing title, subtitle, badge, badge tone, navigation-only actions, disabled execution/grant actions, and safe detail rows for Plasma rendering.
- Added Go and Ruby harness coverage that verifies KDE action card previews remain Runtime-owned, Go-backed, KDE-targeted, backend-detail-free, host-root-safe, and honest about card persistence, status persistence, review receipts, queue mutation, request objects, resource grants, notifications, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved the first-release KDE Compatibility Center queue flow from raw status previews toward renderable user cards while keeping Runtime write and execution gates closed.

## [0.2.146] - 2026-07-15

### Added

- Added `xnix-runtime-go kde-action-status-preview --registry <path> --app <id> --action <id> --decision <decision> [file://...]` for read-only Runtime previews of user-visible Compatibility Center action status after receipt preview.
- Added Go Runtime KDE action status modeling that maps approved, reviewed, deferred, and rejected action decisions into KDE-safe waiting, deferred, or rejected states without persisting status.
- Added Go and Ruby harness coverage that verifies KDE action status previews remain Runtime-owned, Go-backed, KDE-targeted, backend-detail-free, host-root-safe, and honest about status persistence, review receipts, queue mutation, request objects, resource grants, notifications, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved the first-release KDE Compatibility Center queue flow from receipt shape toward user-visible status cards while keeping Runtime write and execution gates closed.

## [0.2.145] - 2026-07-15

### Added

- Added `xnix-runtime-go kde-action-receipt-preview --registry <path> --app <id> --action <id> --decision <decision> [file://...]` for read-only Runtime previews of the review receipt shape for one queued KDE action.
- Added Go Runtime KDE action receipt modeling that derives from action preflight state, builds a stable receipt identifier, and exposes required receipt fields without recording the receipt.
- Added Go and Ruby harness coverage that verifies KDE action receipt previews remain Runtime-owned, Go-backed, KDE-targeted, backend-detail-free, host-root-safe, and honest about decision recording, review receipts, queue mutation, request objects, resource grants, notifications, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved the first-release KDE Compatibility Center queue flow from preflight visibility toward a durable review-receipt contract while keeping Runtime write gates closed.

## [0.2.144] - 2026-07-15

### Added

- Added `xnix-runtime-go kde-action-preflight-preview --registry <path> --app <id> --action <id> --decision <decision> [file://...]` for read-only Runtime previews of preflight gates after one queued KDE action has been reviewed.
- Added Go Runtime KDE action preflight modeling that joins action review intent with execution preflight state before any review receipt, request object, permission grant, or launch approval can exist.
- Added Go and Ruby harness coverage that verifies KDE action preflight previews remain Runtime-owned, Go-backed, KDE-targeted, backend-detail-free, host-root-safe, and honest about review receipts, queue mutation, request objects, resource grants, notifications, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved the first-release KDE Compatibility Center queue flow from review preview toward explicit preflight-gate visibility while keeping all execution and persistence gates closed.

## [0.2.143] - 2026-07-15

### Added

- Added `xnix-runtime-go kde-action-review-preview --registry <path> --app <id> --action <id> --decision <decision> [file://...]` for read-only Runtime previews of user review intent on one queued KDE action.
- Added Go Runtime KDE action review modeling that selects a queued launcher, task-manager, Dolphin, tray, notification, Compatibility Center, or settings action and keeps the user decision visible without recording a durable receipt.
- Added Go and Ruby harness coverage that verifies KDE action review previews remain Runtime-owned, Go-backed, KDE-targeted, backend-detail-free, host-root-safe, and honest about decision recording, review receipts, queue mutation, settings persistence, notifications, permission grants, request objects, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved the first-release KDE Compatibility Center queue flow one step closer to user-review handling while keeping the Runtime's execution and persistence gates closed.

## [0.2.142] - 2026-07-15

### Added

- Added `xnix-runtime-go kde-action-queue-preview --registry <path> --app <id> --decision <decision> [file://...]` for read-only Runtime previews of Compatibility Center queue cards derived from the seven KDE entrypoint actions.
- Added Go Runtime KDE action queue modeling that keeps launcher, task-manager, Dolphin, tray, notification, Compatibility Center, and settings follow-up visible to KDE without persisting queues or creating launch request objects.
- Added Go and Ruby harness coverage that verifies KDE action queue previews remain Runtime-owned, Go-backed, KDE-targeted, Portal-aware for file actions, backend-detail-free, host-root-safe, and honest about queue persistence, desktop file writes, MIME default writes, settings persistence, notifications, task-manager activation, KWin rule application, tray bridging, request objects, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved the first-release KDE entrypoint action contract toward a Compatibility Center queue read model while keeping KDE responsible for presentation and user interaction only.

## [0.2.141] - 2026-07-15

### Added

- Added `xnix-runtime-go kde-entrypoint-action-preview --registry <path> --app <id> --entrypoint <id> --decision <decision> [file://...]` for read-only Runtime previews of a single KDE entrypoint action.
- Added Go Runtime action routing for launcher, task-manager, Dolphin, tray, notification, Compatibility Center, and settings entry points while keeping all execution and host mutation disabled.
- Added Go and Ruby harness coverage that verifies KDE entrypoint action previews remain Runtime-owned, Go-backed, KDE-targeted, Portal-aware for file actions, backend-detail-free, host-root-safe, and honest about desktop file writes, MIME default writes, settings persistence, notifications, task-manager activation, KWin rule application, tray bridging, request objects, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved the first-release KDE entrypoint contract from visibility-only planning toward user-action routing without granting KDE backend policy ownership.

## [0.2.140] - 2026-07-15

### Added

- Added `xnix-runtime-go kde-entrypoints-preview --registry <path> --app <id> --decision <decision> [file://...]` for read-only Runtime previews of the seven first-release KDE entry points from digest-verified registry recipes.
- Added Go Runtime KDE entrypoint aggregation that combines session status with launcher, task-manager, Dolphin, tray, notification, Compatibility Center, and settings surface planning for one normal application contract.
- Added Go and Ruby harness coverage that verifies KDE entrypoint previews remain Runtime-owned, Go-backed, KDE-targeted, official-desktop-only, backend-detail-free, host-root-safe, and honest about desktop file writes, MIME default writes, settings persistence, notifications, task-manager activation, KWin rule application, tray bridging, backend process starts, request objects, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved the first-release KDE application surface contract one step closer to Go-owned Runtime product logic while keeping KDE responsible for display, navigation, and user interaction only.

## [0.2.139] - 2026-07-15

### Added

- Added `xnix-runtime-go execution-session-status-preview --registry <path> --app <id> --decision <decision> [file://...]` for read-only Runtime execution session status previews from digest-verified registry recipes.
- Added Go Runtime session status planning that turns session identity plans into KDE-visible status, desktop-surface, user-visible state, and live-session gate summaries without observing or persisting live state.
- Added Go and Ruby harness coverage that verifies execution-session-status previews remain Runtime-owned, Go-backed, KDE-targeted, Compatibility Center-ready, backend-detail-free, host-root-safe, and honest about live session creation, session registration, live state observation, status persistence, task-manager activation, KWin rule application, tray bridging, backend process starts, request objects, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved launch-session status observability one step closer to Go-owned Runtime product logic while keeping KDE responsible for display and navigation only.

## [0.2.138] - 2026-07-15

### Added

- Added `xnix-runtime-go execution-session-preview --registry <path> --app <id> --decision <decision> [file://...]` for read-only Runtime execution session identity previews from digest-verified registry recipes.
- Added Go Runtime session identity planning that combines launch transactions, KDE window identity, task-manager hints, KWin rule hints, and tray status into one KDE-visible session read model.
- Added Go and Ruby harness coverage that verifies execution-session previews remain Runtime-owned, Go-backed, KDE-targeted, Compatibility Center-ready, backend-detail-free, host-root-safe, and honest about live session creation, window observation, task-manager activation, KWin rule application, tray bridging, backend process starts, request objects, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved launch-session desktop identity one step closer to Go-owned Runtime product logic while keeping KDE responsible for display and window-shell integration only.

## [0.2.137] - 2026-07-15

### Added

- Added `xnix-runtime-go execution-transaction-preview --registry <path> --app <id> --decision <decision> [file://...]` for read-only Runtime launch transaction previews from digest-verified registry recipes.
- Added Go Runtime launch transaction planning that combines execution resource grants, execution readiness, backend binding, snapshot baseline, and the Runtime Launch write gate into one KDE-visible transaction read model.
- Added Go and Ruby harness coverage that verifies execution-transaction previews remain Runtime-owned, Go-backed, KDE-targeted, Compatibility Center-ready, backend-detail-free, host-root-safe, and honest about transaction commits, Portal approvals, resource grants, snapshot creation, backend binding, request objects, Runtime launch approval, and execution remaining disabled.

### Changed

- Moved launch commit orchestration one step closer to Go-owned Runtime product logic while keeping KDE responsible for display and review only.

## [0.2.136] - 2026-07-15

### Added

- Added `xnix-runtime-go execution-resource-grant-preview --registry <path> --app <id> --decision <decision> [file://...]` for read-only Runtime launch resource-grant previews from digest-verified registry recipes.
- Added Go Runtime launch resource planning that connects execution preflight, compatibility permission review, and desktop resource bridge summaries into one KDE-visible resource access read model.
- Added Go and Ruby harness coverage that verifies execution-resource-grant previews remain Runtime-owned, Go-backed, KDE-targeted, Compatibility Center-ready, Portal-aware, backend-detail-free, host-root-safe, and honest about grant objects, request objects, permission grants, resource bridges, settings persistence, launch approval, network requirements, and execution remaining disabled.

### Changed

- Moved launch-time resource access planning one step closer to Go-owned Runtime product logic while keeping KDE responsible for display and review only.

## [0.2.135] - 2026-07-15

### Added

- Added `xnix-runtime-go execution-preflight-preview --registry <path> --app <id> --decision <decision> [file://...]` for read-only Runtime launch preflight previews from digest-verified registry recipes.
- Added Go Runtime launch-preflight planning that turns a Compatibility Center launch decision into explicit user decision, Portal review, snapshot baseline, backend binding, and Runtime Launch write-gate checks before execution can be considered.
- Added Go and Ruby harness coverage that verifies execution-preflight previews remain Runtime-owned, Go-backed, KDE-targeted, Compatibility Center-ready, Portal-aware, backend-detail-free, host-root-safe, and honest about preflight, Runtime approval, request objects, permissions, backend binding, and execution remaining disabled.

### Changed

- Moved the user-visible launch preflight surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for displaying gate status only.

## [0.2.134] - 2026-07-15

### Added

- Added `xnix-runtime-go execution-decision-preview --registry <path> --app <id> --decision <decision> [file://...]` for read-only KDE Compatibility Center launch decision previews from digest-verified registry recipes.
- Added Go Runtime decision planning that captures reviewed, approved, deferred, and rejected launch-review intent without recording review receipts, persisting queues, granting Runtime launch approval, or starting execution.
- Added Go and Ruby harness coverage that verifies execution-decision previews remain Runtime-owned, Go-backed, KDE-targeted, Compatibility Center-ready, Portal-aware, backend-detail-free, host-root-safe, and honest about user intent being captured while Runtime approval, review receipts, queue mutation, permissions, and execution remain disabled.

### Changed

- Moved the user-visible launch review decision surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for forwarding the decision intent only.

## [0.2.133] - 2026-07-15

### Added

- Added `xnix-runtime-go execution-review-preview --registry <path> --app <id> [file://...]` for read-only KDE Compatibility Center launch-review card previews from digest-verified registry recipes.
- Added Go Runtime review-card planning that turns a blocked execution request preview into a non-persistent Compatibility Center request-review queue candidate with review card metadata, gate summaries, selected file routing, and write-gate denial state.
- Added Go and Ruby harness coverage that verifies execution-review previews remain Runtime-owned, Go-backed, KDE-targeted, Compatibility Center-ready, Portal-aware, backend-detail-free, host-root-safe, and honest about review receipts, queue persistence, request persistence, permissions, and execution remaining disabled.

### Changed

- Moved the user-visible launch review card surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for display and navigation only.

## [0.2.132] - 2026-07-15

### Added

- Added `xnix-runtime-go execution-request-preview --registry <path> --app <id> [file://...]` for read-only KDE execution-request intake previews from digest-verified registry recipes.
- Added Go Runtime request-intake planning that turns KDE launch intent into a blocked, non-persistent execution request preview with launch-intent summary, Runtime gate summary, selected file routing, write-gate denial metadata, and safety flags before any request object can be created.
- Added Go and Ruby harness coverage that verifies execution-request previews remain Runtime-owned, Go-backed, KDE-targeted, Compatibility Center-ready, Portal-aware, backend-detail-free, host-root-safe, and honest about request persistence, permissions, execution, and profile binding remaining disabled.

### Changed

- Moved the user-visible execution-request intake surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for forwarding launch intent and presenting review state only.

## [0.2.131] - 2026-07-15

### Added

- Added `xnix-runtime-go launch-intent-preview --registry <path> --app <id> [file://...]` for read-only KDE desktop-launch intent previews from digest-verified registry recipes.
- Added Go Runtime launch-intent planning that captures desktop launcher clicks, selected file URIs, recommended compatibility profile, run-plan safety, write-gate denial metadata, and blocked actions before any execution request can be created.
- Added Go and Ruby harness coverage that verifies launch-intent previews remain Runtime-owned, Go-backed, KDE-targeted, Portal-aware, desktop-entry-safe, backend-detail-free, host-root-safe, and honest about launch, request objects, permissions, execution, and profile binding remaining disabled.

### Changed

- Moved the user-visible KDE launcher intent surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for showing and forwarding launch intent only.

## [0.2.130] - 2026-07-15

### Added

- Added `xnix-runtime-go execution-readiness-preview --registry <path> --app <id>` for read-only KDE Compatibility Center launch-readiness previews from digest-verified registry recipes.
- Added Go Runtime execution-readiness planning that exposes application identity, the recommended compatibility profile, Runtime gates, gate counts, blocked actions, and launch safety flags before any execution request can be created.
- Added Go and Ruby harness coverage that verifies execution-readiness previews remain Runtime-owned, Go-backed, KDE-targeted, AI-diagnostic-safe, desktop-entry-visible, backend-detail-free, host-root-safe, and honest about launch, execution requests, profile binding, Portal review, and snapshot gates remaining blocked or required.

### Changed

- Moved the user-visible launch-readiness surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for displaying launch status and routing launch intent through Runtime gates.

## [0.2.129] - 2026-07-15

### Added

- Added `xnix-runtime-go backend-selection-preview --registry <path> --app <id>` for read-only KDE Compatibility Center backend-selection previews from digest-verified registry recipes.
- Added Go Runtime backend-selection planning for local and isolated compatibility profiles with recommendation, required review, preflight, and blocked-action metadata.
- Added Go and Ruby harness coverage that verifies backend-selection previews remain Runtime-owned, Go-backed, KDE-targeted, backend-detail-free, host-root-safe, and honest about profile selection, launch, environment creation, request objects, state roots, snapshots, privileged containers, and capability activation remaining disabled.

### Changed

- Moved the user-visible compatibility profile selection surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for presentation and confirmation only.

## [0.2.128] - 2026-07-15

### Added

- Added `xnix-runtime-go mode-switch-preview --registry <path> --app <id> --mode <mode>` for read-only KDE unified-settings mode switch previews from digest-verified registry recipes.
- Added Go Runtime mode-switch planning for automatic, prefer-performance, prefer-compatibility, and isolated-execution modes with user confirmation, Portal review, snapshot, environment-plan, and write-gate metadata.
- Added Go and Ruby harness coverage that verifies mode-switch previews remain Runtime-owned, Go-backed, KDE-targeted, backend-detail-free, host-root-safe, and honest about settings persistence, backend reconfiguration, backend processes, and launch remaining disabled.

### Changed

- Moved the user-visible compatibility mode switch surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for display and confirmation only.

## [0.2.127] - 2026-07-15

### Added

- Added `xnix-runtime-go review-flow-preview --registry <path> --app <id> [--section <section>] [--field <field>] [--value <value>] [--operation <operation>]` for read-only KDE Compatibility Center review-flow previews from digest-verified registry recipes.
- Added Go Runtime review-flow planning that connects settings-change review, permission review, Portal request review, Runtime write gates, and review receipts into one KDE-visible confirmation flow.
- Added Go and Ruby harness coverage that verifies review-flow previews remain Runtime-owned, Go-backed, KDE-targeted, backend-detail-free, host-root-safe, and honest about apply, request objects, permission grants, settings persistence, execution, and receipt recording remaining disabled or pending.

### Changed

- Moved the user-visible Compatibility Center review-flow surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for displaying confirmation steps only.

## [0.2.126] - 2026-07-15

### Added

- Added `xnix-runtime-go settings-change-preview --registry <path> --app <id> --section <section> --field <field> --value <value>` for read-only KDE unified-settings change previews from digest-verified registry recipes.
- Added Go Runtime settings-change planning for run mode, resource access, device access, network access, and snapshots with user confirmation, Portal policy review, restore-point, persistence, and blocked-action metadata.
- Added Go and Ruby harness coverage that verifies settings-change previews remain Runtime-owned, KDE-targeted, backend-detail-free, host-root-safe, and honest about apply, persistence, resource grants, and review gates remaining disabled or required.

### Changed

- Moved the user-visible settings-change review surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for presentation and confirmation only.

## [0.2.125] - 2026-07-15

### Added

- Added `xnix-runtime-go portal-request-preview --registry <path> --app <id> --operation <operation> [--reason <text>]` for read-only KDE Portal request previews from digest-verified registry recipes.
- Added Go Runtime Portal request planning for file-open, URI-open, print, screenshot, clipboard, camera, and remote-desktop operations with deterministic handle tokens, Portal D-Bus endpoints, completion metadata, and denied-action guidance.
- Added Go and Ruby harness coverage that verifies Portal request previews remain Runtime-owned, KDE-targeted, backend-detail-free, host-root-safe, host-path-free, and honest about request objects, permission grants, host permission changes, direct access, and denied operations.

### Changed

- Moved the user-visible Portal request planning surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for presenting review and denial guidance only.

## [0.2.124] - 2026-07-15

### Added

- Added `xnix-runtime-go desktop-resource-bridge-preview --registry <path> --app <id>` for read-only KDE desktop resource bridge previews from digest-verified registry recipes.
- Added Go Runtime desktop resource bridge planning for file-open, URI-open, print, clipboard, and screenshot bridges with Portal metadata and required Runtime gate summaries.
- Added Go and Ruby harness coverage that verifies bridge previews remain Runtime-owned, KDE-targeted, Portal-mediated, backend-detail-free, host-root-safe, host-path-free, and honest about bridges, request objects, backend processes, and direct desktop resource access remaining disabled.

### Changed

- Moved the user-visible desktop resource bridge surface one step closer to Go-owned Runtime product logic while keeping KDE responsible for bridge presentation and review only.

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
