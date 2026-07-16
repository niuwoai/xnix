# Windows App Compatibility Implementation Brief

> Last updated: 2026-07-15 | Scope: KDE-first Linux desktop for existing Windows applications

This brief splits the Xnix product goal into large, independent implementation packages. It is written for handoff to another coding agent or contributor without requiring them to understand the whole repository first.

The product goal is not "a KDE theme" and not "a from-scratch desktop." Xnix should become a Linux system that makes existing Windows applications feel like ordinary Linux desktop applications while an independent Runtime manages compatibility, AI diagnostics, permissions, state, snapshots, and rollback.

For copy-first Claude Code workstream prompts, use `docs/claude-code-windows-compatibility-workstreams.md`. For current mainline merge review and convergence rules, use `docs/mainline-integration-checkpoint.md`.

## Product Architecture Decision

Xnix should use a mature desktop environment as the user shell and keep the product core in an independent Runtime.

First official desktop:

- KDE Plasma.

Later possible editions:

- GNOME for controlled AI terminal appliances.
- XFCE for low-resource or thin-client deployments.

Do not implement first-release deep support for KDE, GNOME, and XFCE at the same time. The first product release should support KDE Plasma officially; other desktops may run but should not promise full AI and Windows application integration.

## Component Boundary

```text
KDE Plasma desktop
  Start menu, task manager, notifications, file manager, tray, settings

KDE integration plugins
  Plasmoid, KRunner, KWin scripts, Dolphin actions, service menus

Xnix Compatibility Runtime
  Independent service and D-Bus API
  Application recipes, compatibility planning, package sources
  AI diagnostics, permissions, snapshots, rollback

Compatibility backends
  Local compatibility layer and isolated compatibility environments

Linux system
  Atomic base, system services, image and QEMU smoke
```

KDE code must handle presentation and user interaction. Runtime code must own compatibility policy and state.

## Implementation Language Policy

- Use Go for durable Runtime product logic, long-running services, state machines, package acquisition, execution planning, and APIs.
- Use C for low-level policy surfaces, ABI-shaped boundaries, and already implemented C Runtime records.
- Use Ruby for tests, smoke scripts, scaffolding, and occasional developer tools.
- Do not move new core business logic into Ruby unless it is explicitly test-only or one-shot tooling.

## User Experience Rules

Installed Windows applications must appear as ordinary Linux applications.

KDE-facing UI should show:

- Application name.
- Application icon.
- Normal launcher entry.
- Normal task-manager identity.
- Supported file types.
- Compatibility mode in user language such as Automatic, Performance priority, or Compatibility priority.
- Permission states such as Documents allowed, Camera blocked, Network allowed.

KDE-facing UI must not show:

- Raw executable paths.
- Raw backend commands.
- Internal state-root paths.
- Compatibility storage implementation terms.
- Developer-only environment terms.

## Safety Rules for Every Package

- No privileged containers.
- No host networking.
- No Docker socket mounts.
- No broad host-directory mounts.
- No host root mutation.
- No real backend launch unless a package explicitly enables a gated execution path.
- No real network artifact fetch unless a package explicitly enables a constrained fixture or reviewed production path.
- No secrets, signing keys, credentials, or private webhook URLs in the repository.
- Every project-facing file must be English.

## Implementation Package Map

| Package | Goal | Primary layer | Can start now |
| --- | --- | --- | --- |
| W1 Runtime owner route migration | Move remaining read-only Runtime product methods toward Go ownership | Go Runtime | Yes |
| W2 KDE seven-entry-point integration | Make KDE consume Runtime models for launcher, taskbar, file manager, tray, notifications, center, and settings | KDE integration and Runtime API | Yes |
| W3 Recipe trust and application identity | Make application recipes production-shaped and digest/signature aware | Go Runtime | Yes |
| W4 Package source and artifact acquisition | Safely plan, verify, cache, and stage compatibility artifacts | Go Runtime | After W3 basics |
| W5 Backend lifecycle state machine | Track compatibility environment readiness without exposing backend details | Go Runtime | After W4 model |
| W6 Portal permission broker | Own desktop permission requests and completion state | Go Runtime and XDG Portal | Yes in fake mode |
| W7 Snapshot and rollback store | Provide restore points for risky compatibility changes | Go Runtime | After W5 state root |
| W8 Execution transaction pipeline | Connect launch intent, review, preflight, permissions, snapshots, and sessions | Go Runtime | After W5, W6, W7 |
| W9 AI compatibility diagnostics | Feed safe Runtime evidence into AI repair and recommendation flows | Go Runtime | Yes in no-provider mode |
| W10 KDE materialization writer | Write desktop entries, MIME records, service menus, icons, and receipts under explicit roots | Go or Ruby writer with Runtime policy | After W2 and W3 |
| W11 Product image and QEMU smoke | Prove the KDE-first product image path without breaking the learning baseline | Build and packaging | Later |

## W1: Runtime Owner Route Migration

### Goal

Move remaining read-only Runtime methods from static C policy records or legacy adapter paths into native Go owner previews or thin Go adapters.

### Suggested scope per version

Migrate 3 to 5 related read-only methods at a time.

Good slices:

- Package acquisition family: package source, artifact manifest, acquisition preflight, install plan.
- Backend family: capability matrix, lifecycle, environment plan, binding.
- Portal family: access policy, request plan, permission review.
- Execution family: execution readiness, launch intent, run plan, repair plan.

### Acceptance criteria

- `runtime-owner-route-manifest-preview` shows the migrated methods as Go routes.
- C-backed route count decreases.
- Write methods remain disabled.
- No backend launch, package installation, network fetch, or host-root mutation is enabled.

### Required tests

- `go test ./...`
- `ruby scripts/verify_layout.rb`
- Runtime method parity tests.
- Ruby tests for the migrated method family.

## W2: KDE Seven-Entry-Point Integration

### Goal

Make the first-release KDE integration coherent around seven official entry points:

1. Start menu.
2. Task manager.
3. File manager.
4. System tray.
5. Notification center.
6. AI Compatibility Center.
7. Unified settings.

### Scope

- Keep KDE as a presentation shell only.
- Make every KDE-facing surface consume Runtime D-Bus read models or generated activation artifacts.
- Ensure a Windows application can be shown as a normal Linux application identity.
- Keep user action submission separate from Runtime approval and execution.

### Acceptance criteria

- Each entry point has a documented Runtime method or generated artifact source.
- KDE-facing strings avoid backend details.
- File open and drag/drop plans require Portal mediation.
- Notifications and tray actions route back to Runtime review or settings surfaces.

### Required tests

- `go test ./...`
- KDE model Ruby tests.
- Runtime D-Bus client tests.
- `ruby scripts/verify_layout.rb`

## W3: Recipe Trust and Application Identity

### Goal

Turn local development recipes into a production-shaped trust model without committing production secrets.

### Scope

- Verify recipe digests.
- Model production-signed, development-only, unsigned, and invalid states.
- Keep development recipes usable but never production trusted.
- Feed trust evidence into install gates, owner readiness, diagnostics, and Compatibility Center.

### Acceptance criteria

- Invalid recipe digest fails closed.
- Missing production signature blocks production activation.
- No signing key or secret is committed.
- Application identity is derived from trusted recipe metadata, not raw executable paths.

### Required tests

- `go test ./...`
- Recipe registry tests.
- Recipe trust policy tests.
- Recipe install gate tests.
- `ruby scripts/verify_layout.rb`

## W4: Package Source and Artifact Acquisition

### Goal

Build a safe package acquisition pipeline that can plan, verify, cache, and stage artifacts before any installation is possible.

### Scope

- Parse artifact manifests.
- Verify digests.
- Model cache namespaces without exposing host paths.
- Support local fixture acquisition first.
- Stage only under explicit project-controlled or container-controlled roots.

### Out of scope

- System package manager writes.
- Global cache writes.
- Privileged installation.
- Unreviewed network fetch.

### Acceptance criteria

- Digest mismatch blocks staging.
- Dry-run acquisition reports planned artifact records.
- Install plan consumes acquisition readiness.
- Host root mutation remains false.

### Required tests

- `go test ./...`
- Artifact, acquisition, install, and package-source tests.
- `ruby scripts/verify_layout.rb`

## W5: Backend Lifecycle State Machine

### Goal

Create a Runtime-owned lifecycle model for compatibility environments while keeping real backend processes disabled by default.

### States

- `missing`
- `planned`
- `staged`
- `ready`
- `repair-required`
- `blocked`

### Scope

- Persist lifecycle state under a test-controlled Runtime state root.
- Connect backend selection, environment plan, binding status, and diagnostics to the same state evidence.
- Provide repair hints without executing repair.

### Acceptance criteria

- State survives within a controlled test root.
- KDE sees user-facing readiness, not backend internals.
- Execution remains blocked until W8.
- The state model refuses paths outside its configured root.

### Required tests

- `go test ./...`
- Backend selection, environment, binding, and lifecycle tests.
- `ruby scripts/verify_layout.rb`

## W6: Portal Permission Broker

### Goal

Implement a Runtime-owned request broker for XDG Desktop Portal-style permissions.

### Scope

- Track file, URI, print, clipboard, screenshot, camera, and remote-desktop request objects.
- Add a fake broker for container tests.
- Connect request state to permission review and execution preflight.
- Keep user review separate from execution approval.

### Acceptance criteria

- Request creation is explicit and test-gated.
- Completion, denial, timeout, and failure states are visible in diagnostics.
- KDE never receives direct host paths or bypasses review.

### Required tests

- `go test ./...`
- Portal policy and request tests.
- Permission review tests.
- `ruby scripts/verify_layout.rb`

## W7: Snapshot and Rollback Store

### Goal

Provide a constrained restore-point mechanism for Runtime-managed state roots.

### Scope

- Create, list, verify, and roll back snapshots under an explicit test root.
- Use metadata and content digests.
- Connect snapshot readiness to repair and execution preflight.
- Emit rollback receipts.

### Out of scope

- Host filesystem snapshots.
- Btrfs, ZFS, or system snapshot integration.
- Privileged rollback.

### Acceptance criteria

- Snapshot operations refuse paths outside the configured root.
- Rollback restores fixture state in tests.
- Repair and execution preflight can require a verified restore point.

### Required tests

- `go test ./...`
- Snapshot and rollback tests.
- `ruby scripts/verify_layout.rb`

## W8: Execution Transaction Pipeline

### Goal

Connect launch intent, review, permission, snapshot, backend lifecycle, resource grant, transaction, session, and status contracts into one blocked-by-default pipeline.

### Scope

- Generate execution request IDs.
- Track review decisions separately from launch.
- Require backend lifecycle readiness, Portal permission state, and snapshot state.
- Provide deterministic blocked reasons.
- Keep real process launch disabled.

### Acceptance criteria

- User approval alone does not launch.
- Missing permission, snapshot, or backend readiness blocks launch with a clear reason.
- Session status is preview-only unless a future package enables real execution.

### Required tests

- `go test ./...`
- Execution request, review, decision, preflight, resource grant, transaction, session, and status tests.
- `ruby scripts/verify_layout.rb`

## W9: AI Compatibility Diagnostics

### Goal

Use Runtime evidence to produce safe AI diagnostic inputs and recommendation flows without sending secrets, user files, or host paths to an AI provider by default.

### Scope

- Collect structured compatibility evidence from recipes, logs, run plans, test results, and repair history.
- Redact host paths, user names, tokens, file contents, and backend commands.
- Keep provider calls disabled unless a future reviewed integration enables them.
- Route repair suggestions through explicit approval gates.

### Acceptance criteria

- Diagnostic inputs are provider-neutral and safe to display.
- Recommendations never execute repair directly.
- Approval gate requires user review and restore-point readiness.

### Required tests

- `go test ./...`
- AI diagnostic input, recommendation, and repair approval tests.
- `ruby scripts/verify_layout.rb`

## W10: KDE Materialization Writer

### Goal

Turn Runtime activation previews into deterministic files under explicit target roots.

### Scope

- Write desktop entries.
- Write Dolphin service menus.
- Write MIME association records.
- Write icons and activation manifests.
- Emit rollback receipts.
- Keep all writes under an explicit target root.

### Acceptance criteria

- Staging writes are deterministic and reversible.
- Rollback removes only unchanged files listed in receipts.
- No real host `/usr`, `/etc`, or home-directory writes happen by default.
- Generated desktop entries expose normal application identity only.

### Required tests

- `go test ./...`
- Desktop activation installer and rollback tests.
- Constrained container build and Runtime D-Bus smokes.
- `ruby scripts/verify_layout.rb`

## W11: Product Image and QEMU Smoke

### Goal

Create a reproducible KDE-first product image smoke while preserving the Buildroot/QEMU learning baseline.

### Scope

- Keep Buildroot as the low-level learning baseline.
- Add product-image metadata for the KDE-first path.
- Verify Runtime files, D-Bus service definitions, KDE integration assets, and safety defaults in an image root.
- Keep QEMU and SSH loopback-bound.
- Persist serial logs for diagnosis.

### Acceptance criteria

- A clean checkout can reproduce the smoke inputs.
- The product smoke does not require privileged host installation.
- SSH forwarding stays loopback-only.
- Buildroot smoke and product-image smoke remain separate.

### Required tests

- `ruby scripts/full_smoke.rb` only when boot/image behavior changes.
- `ruby scripts/verify_layout.rb`
- Relevant constrained container smoke tests.

## Recommended Work Order

1. Keep W1 moving in small route-migration waves.
2. Run W3 next to W1 if it avoids file conflicts.
3. Start W6 in fake mode early because Portal permission state affects execution.
4. Start W2 once the Runtime methods it needs are stable enough.
5. Start W4, then W5, then W7.
6. Start W8 only after backend lifecycle, permissions, and snapshots can provide real evidence.
7. Start W10 after KDE surfaces and recipe trust are stable.
8. Start W11 when the product-image path is ready to prove integration.

## Stop Conditions for a Coding Agent

Stop and report instead of guessing when a task requires:

- Enabling real backend launch.
- Enabling production D-Bus ownership.
- Writing to the host root.
- Downloading network artifacts outside a reviewed fixture.
- Using privileged containers or host networking.
- Adding secrets or signing keys.
- Exposing backend internals to KDE-facing output.
- Combining unrelated packages into one version.

## Handoff Prompt Template

```text
Work in /Users/rocky/Sites/xnix.
Implement package <PACKAGE_ID> from docs/windows-app-compatibility-implementation-brief.md.
Follow AGENTS.md strictly.
Keep all project-facing text in English.
Use Go for durable Runtime product logic, C only where low-level or existing C Runtime boundaries require it, and Ruby mainly for tests or light tooling.
Do not edit docs/claude-code-implementation-packages.md if another agent owns it.
Do not use privileged containers, host networking, Docker socket mounts, broad host mounts, or host root mutation.
Do not enable real backend launch, real installation, or production D-Bus ownership unless the package explicitly requires a gated test path.
Make the smallest coherent version bump if code changes, update VERSION, CHANGELOG.md, PRODUCT_OVERVIEW.md, and scripts/verify_layout.rb as required.
Run package-required tests plus ruby scripts/verify_layout.rb.
Do not commit unrelated .claude/ or tmp/ content.
Report changed files, tests, and remaining blockers.
```
