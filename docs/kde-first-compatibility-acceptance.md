# KDE-First Windows App Compatibility Acceptance Matrix

> Last updated: 2026-07-15 | Scope: first official Xnix desktop experience

This document defines how to prove that Xnix is moving toward "the Linux desktop that best supports existing Windows applications." It focuses on observable user experience, Runtime ownership, and safety gates rather than implementation volume.

The first official desktop is KDE Plasma. GNOME and XFCE can be future shells, but they are not first-release acceptance targets.

## Product-Level Acceptance Statement

Xnix is acceptable for the first KDE-focused milestone when a digest-verified application recipe can appear throughout KDE as a normal Linux application while all compatibility decisions remain owned by the Xnix Compatibility Runtime.

The user should be able to see, search, open, pin, inspect, and configure the application without seeing raw compatibility implementation details. The Runtime may still block actual execution until backend lifecycle, permissions, snapshots, and launch gates are production-ready, but every block must be explicit, user-facing, and safe.

## Global Evidence Rules

Every acceptance item needs evidence from at least one of these sources:

- Go unit or CLI test.
- Ruby integration or smoke test.
- `ruby scripts/verify_layout.rb`.
- Runtime D-Bus contract or session-bus smoke.
- Generated KDE activation artifact under an explicit target root.
- Product documentation that names the current limitation and next gate.

Do not count a green test as evidence unless the test checks the relevant user-facing or Runtime boundary requirement.

## Seven KDE Entry Points

| Entry point | User-facing expectation | Runtime-owned evidence | Must not happen |
| --- | --- | --- | --- |
| Start menu | The Windows application appears with a normal name, icon, categories, and launcher action | Desktop identity, desktop entry, activation bundle, and route manifest evidence | KDE must not expose raw executable paths or backend commands |
| Task manager | Windows application windows can be grouped, pinned, restored, and identified like normal windows | Task manager identity and KWin window identity plans | KDE must not own backend selection or launch permission |
| File manager | Dolphin can offer "open with compatible application" and AI analysis actions through managed intents | File association, Dolphin file-open, drag-and-drop, Portal request, and action card evidence | Dolphin must not read arbitrary files directly or bypass Portal review |
| System tray | Compatibility status and attention states are visible without revealing backend internals | Tray status and backend binding evidence | Tray UI must not start backend processes directly |
| Notification center | Install failure, repair suggestion, approval, and mode-change notifications are understandable | Notification plan, action queue, action review, and repair approval evidence | Notifications must not execute repairs or launch applications |
| AI Compatibility Center | Users can see compatibility state, known issues, repair history, and next safe actions | Compatibility Center summary/page/card deck, diagnostics, AI recommendation, and approval gate evidence | AI flows must not send secrets, file contents, host paths, or raw commands by default |
| Unified settings | Users can choose normal settings such as mode, performance priority, file access, camera, network, and snapshots | Settings, settings-change, permission review, Portal policy, and snapshot evidence | Settings must not expose developer-only environment terminology |

## Runtime Ownership Gates

| Gate | Acceptance requirement | Evidence to inspect |
| --- | --- | --- |
| Owner routing | Read-only product methods should move steadily into Go owner handlers or thin Go owner adapters | `runtime-owner-route-manifest-preview`, Go route tests, method parity tests |
| Production owner | Production D-Bus ownership remains gated until a long-running Go owner can pass smoke tests | Runtime owner readiness, process, service binding, live owner gate, owner smoke plan |
| Recipe trust | Development recipes are usable but never production trusted without production signature evidence | Recipe registry, recipe trust policy, recipe install gate, owner recipe trust |
| Package acquisition | Package source, artifact manifests, acquisition preflight, and install plans are safe before downloads or installs | Package source, artifact manifest, acquisition preflight, install plan tests |
| Backend lifecycle | Compatibility environments have explicit lifecycle state before launch is possible | Backend selection, environment, binding, lifecycle, diagnostics |
| Portal permissions | File, URI, print, clipboard, screenshot, camera, and remote desktop access use request objects and review | Portal access policy, Portal request model, permission review, execution preflight |
| Snapshot baseline | Risky launch, repair, install, and restore flows require restore-point evidence | Snapshot plan, rollback receipt, repair plan, execution preflight |
| Execution transaction | Launch approval, resource grants, backend readiness, Portal permission, snapshot readiness, and process start are separate states | Launch intent, execution request, review, decision, preflight, resource grant, transaction, session, session status |
| AI diagnostics | AI inputs and recommendations are redacted, review-first, and provider-neutral by default | AI diagnostic input, recommendation, repair approval gate, Compatibility Center cards |
| KDE materialization | Desktop files, MIME records, service menus, icons, and receipts are written only under explicit target roots | Desktop activation staging, transaction, status, installer, rollback tests |

## First Milestone Definition

The first milestone is not "real Windows app execution." It is "safe KDE-native application presence."

Minimum acceptance:

- A sample recipe appears as a normal KDE application identity.
- The seven KDE entry points have Runtime-backed read models or generated artifacts.
- D-Bus read-only parity covers the KDE-facing Runtime methods.
- Write methods return explicit disabled errors.
- No backend process starts.
- No host-root mutation occurs.
- No privileged container, host networking, or Docker socket mount is needed.
- File access and AI analysis paths require Portal or privacy-gate evidence.
- The Compatibility Center can explain why launch is blocked and what must become ready.

## Second Milestone Definition

The second milestone is "safe preparation for execution."

Minimum acceptance:

- Recipe trust distinguishes production-signed, development-only, unsigned, and invalid recipes.
- Package acquisition can verify and stage local fixture artifacts under a controlled root.
- Backend lifecycle can reach at least `staged` and `ready` in a test state root without starting real backends.
- Portal request objects can be created and completed through a fake broker.
- Snapshot creation and rollback work inside a test root.
- Execution transactions can move through request, review, preflight, blocked, and ready states without launching.

## Third Milestone Definition

The third milestone is "gated real compatibility execution."

Do not start this milestone until the second milestone evidence is strong.

Minimum acceptance:

- A production-shaped Go Runtime owner owns the D-Bus name in a constrained smoke.
- Backend lifecycle can start a reviewed compatibility backend in an isolated test environment.
- Launch is allowed only after recipe trust, package readiness, Portal permission, snapshot baseline, and user review pass.
- Session status is observable and recoverable.
- Repair and rollback are available before risky changes.
- KDE still never exposes raw backend commands or internal state-root paths.

## Evidence Commands

Use these commands as a starting point. Add package-specific tests when a package changes its area.

```text
go test ./...
ruby scripts/verify_layout.rb
ruby -Ilib test/test_runtime_method_parity_manifest.rb
ruby -Ilib test/test_runtime_service_binding.rb
ruby -Ilib test/test_kde_center_model.rb
ruby -Ilib test/test_kde_integration_status.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
```

Run constrained container smokes when D-Bus, activation packaging, or image behavior changes:

```text
ruby scripts/container.rb build
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
```

Run full smoke only for boot, image, QEMU, or tenth-version milestone changes:

```text
ruby scripts/full_smoke.rb
```

## Non-Acceptance Examples

These do not prove progress toward the product goal by themselves:

- A KDE widget mock that does not consume Runtime state.
- A launcher that exposes a raw compatibility command.
- A passing unit test that only checks JSON shape but not safety flags.
- A backend lifecycle status that is hard-coded and unrelated to state-root evidence.
- A Portal plan that describes permission but cannot create or track a request in test mode.
- An AI recommendation that cannot prove redaction boundaries.
- An installer that writes to a temporary root but cannot produce rollback receipts.

## Current Strategic Bias

When choosing between two useful tasks, prefer the one that makes this statement more true:

> A normal user can treat a Windows application as a normal KDE application, while the Xnix Compatibility Runtime quietly and safely owns all compatibility complexity.
