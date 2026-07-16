# Claude Code Windows Compatibility Workstreams

> Last updated: 2026-07-16 | Product target: best Linux desktop for existing Windows applications

This document is the copy-first workstream board for handing Xnix Windows application compatibility work to Claude Code.

Use this file when the implementation goal is broader than one contract gap but still needs to remain branch-sized and reviewable. It translates the KDE-first product direction into independent implementation streams.

Do not modify `docs/claude-code-implementation-packages.md` from this document's scope. That file can be maintained by another agent.

## Product North Star

Xnix should become a Linux desktop where existing Windows applications feel like ordinary desktop applications.

The first official shell is KDE Plasma. GNOME and XFCE can be future shells, but they are not first-release integration targets.

The architecture is:

```text
KDE Plasma desktop
  Start menu, task manager, notifications, Dolphin, tray, settings

KDE integration layer
  Plasmoid, KRunner, KWin scripts, Dolphin actions, service menus

Xnix AI Compatibility Runtime
  Independent Go-first service and D-Bus API
  Recipes, Wine/VM management, permissions, snapshots, rollback, diagnostics

Compatibility backends
  Wine, Proton, Windows VM, and future isolated execution providers

Linux system
  Atomic base, system services, image build, QEMU smoke evidence
```

KDE is the replaceable user shell. The Runtime is the product core.

## Implementation Language Policy

- Go owns durable Runtime product logic, long-running services, state machines, package/artifact handling, execution planning, D-Bus owner behavior, and safety gates.
- C remains valid for low-level transport, ABI-shaped records, and already-owned Runtime policy surfaces.
- Ruby is for tests, smoke scripts, reports, and lightweight developer tooling.
- Do not add new core business logic in Ruby unless the work is explicitly test-only or one-shot tooling.

## User-Facing Safety Rules

Windows applications must appear as normal Linux applications.

KDE-facing output may show:

- Application name.
- Application icon.
- Normal launcher entry.
- Task-manager identity.
- Supported file types.
- User-safe compatibility mode names such as Automatic, Performance priority, or Compatibility priority.
- Permission states such as Documents allowed, Camera blocked, and Network allowed.
- Snapshot and repair status in user-safe language.

KDE-facing output must not show:

- Raw Windows executable paths.
- Raw Wine, Proton, VM, or backend commands.
- Prefix terminology.
- Runtime state-root paths.
- Host paths.
- File contents.
- Secrets, tokens, private keys, webhook URLs, or credentials.

## Global Workstream Rules

Every Claude Code branch must:

- Pick exactly one workstream from `CW1` through `CW11`.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Runtime logic.
- Keep unsafe production behavior disabled unless the selected workstream explicitly adds a gated test-only path.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for code changes.
- Add targeted success, failure, and blocked-state tests.
- Run the workstream-specific verification commands plus `ruby scripts/verify_layout.rb`.
- Stop and report if the work needs cross-workstream ownership or unsafe host behavior.

Unsafe behavior remains disabled by default:

- Production D-Bus ownership.
- Runtime write methods.
- Real Wine, Proton, VM, or backend launch.
- Real Portal transport calls.
- Default network artifact fetch.
- Host package-manager calls.
- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Host-root mutation.
- KDE-owned compatibility policy.

## Recommended First Wave

Start with these four workstreams:

| Order | Workstream | Why first |
| --- | --- | --- |
| 1 | `CW1` Runtime owner read boundary | Later features need a real Go-owned read boundary instead of static previews. |
| 2 | `CW2` Recipe and artifact trust pipeline | Desktop integration and execution must depend on verified inputs. |
| 3 | `CW3` Runtime state root and backend lifecycle | Execution, diagnostics, and KDE status need durable state without launching backends. |
| 4 | `CW10` Evidence and drift harness | The repository needs gates that resist more contract-only expansion. |

Then move to `CW4`, `CW5`, `CW6`, `CW7`, and finally `CW8` and `CW11`.

Do not assign execution-launch work before trust, lifecycle, permissions, and snapshots have real blocking evidence.

## Workstream Index

| ID | Workstream | Primary layer | Main evidence |
| --- | --- | --- | --- |
| `CW1` | Runtime owner read boundary | Go Runtime | Private smoke owner, route parity, disabled writes |
| `CW2` | Recipe and artifact trust pipeline | Go Runtime | Digest-verified recipes/artifacts, staging receipts, install gates |
| `CW3` | Runtime state root and backend lifecycle | Go Runtime | Durable lifecycle states without backend launch |
| `CW4` | KDE seven entry points | KDE integration plus Runtime reads | Start menu, task manager, Dolphin, tray, notifications, center, settings |
| `CW5` | Portal permission broker | Go Runtime plus XDG Portal model | Fake-mode request records and review states |
| `CW6` | Snapshot and rollback store | Go Runtime | Content-addressed restore points and rollback receipts |
| `CW7` | AI diagnostics and repair boundary | Go Runtime | Redacted diagnostic inputs and review-only recommendations |
| `CW8` | Execution transaction ledger | Go Runtime | Reviewed transaction records that still block launch |
| `CW9` | KDE materialization writer | Go Runtime plus KDE files | Desktop, MIME, service-menu, icon, and activation receipts under target roots |
| `CW10` | Evidence and drift harness | Ruby tooling plus Go fixtures | Reports that fail on orphan contracts and preview-only regressions |
| `CW11` | Product image and QEMU acceptance | Build/image tooling | KDE-first image smoke with loopback-only SSH and persisted logs |

## CW1: Runtime Owner Read Boundary

### Mission

Make the Runtime's read-only product surface owned by a constrained Go process boundary.

### Start from

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `internal/runtime/appidentity/runtime_owner_*.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `runtime/dbus/xnix_compatd_smoke.c`
- `scripts/runtime_contract_drift_report.rb`

### Deliver

- A long-running smoke-owner mode for the Go Runtime owner.
- Private session-bus ownership only inside constrained smoke tests.
- Read-only D-Bus routing through Go owner handlers.
- Stable disabled responses for every write method.
- Route parity evidence across D-Bus XML, Go routes, C smoke adapter, Ruby client, and CLI commands.

### Do not deliver

- Production D-Bus ownership.
- Write enablement.
- Backend launch.
- System service installation.
- Host-root mutation.

### Required verification

```text
go test ./cmd/xnix-runtime-owner ./internal/runtime/owner ./cmd/xnix-runtime-go ./internal/runtime/appidentity
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/container.rb runtime-owner-candidate-smoke
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CW1 from docs/claude-code-windows-compatibility-workstreams.md.

Target outcome:
- Add or strengthen the constrained Go Runtime owner read boundary so read-only Runtime methods are served through Go owner routes and write methods fail closed.

Hard constraints:
- No production D-Bus ownership, write methods, system service installation, Wine/Proton/VM launch, network fetch, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add success, unsupported-read, disabled-write, lifecycle, and blocked-production-owner tests.
- Run the CW1 verification commands.
```

## CW2: Recipe and Artifact Trust Pipeline

### Mission

Make application recipes, package sources, artifact manifests, cache state, staging, and install readiness verifiable before desktop or execution work depends on them.

### Start from

- `internal/runtime/recipe/`
- `internal/runtime/artifact/`
- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/package_source.go`
- `internal/runtime/appidentity/acquisition_preflight.go`
- `internal/runtime/appidentity/artifact_manifest.go`
- `internal/runtime/appidentity/install_plan.go`
- `runtime/recipes/`

### Deliver

- Read-only recipe stores with explicit local roots.
- Registry validation for schema, IDs, relative paths, digests, and declared signing state.
- A replaceable signature verifier boundary without committed signing secrets.
- Local fixture artifact manifests with SHA-256 verification.
- Runtime-root-scoped cache and staging receipts.
- Install readiness that joins recipe trust, artifact verification, staging state, and owner readiness.

### Do not deliver

- Default network downloads.
- Host package-manager calls.
- Private keys.
- Production trust bypasses.
- Desktop writes.
- Backend launch.

### Required verification

```text
go test ./internal/runtime/recipe ./internal/runtime/artifact ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_recipe_registry.rb
ruby -Ilib test/test_recipe_trust_policy.rb
ruby -Ilib test/test_recipe_install_gate.rb
ruby -Ilib test/test_compatibility_package_source.rb
ruby -Ilib test/test_compatibility_acquisition_preflight.rb
ruby -Ilib test/test_compatibility_artifact_manifest.rb
ruby -Ilib test/test_compatibility_install_plan.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CW2 from docs/claude-code-windows-compatibility-workstreams.md.

Target outcome:
- Build a Go-first local recipe and artifact trust pipeline that verifies digests, stages only under explicit roots, and fails closed for invalid inputs.

Hard constraints:
- No default network fetch, host package-manager call, private key, production trust bypass, desktop write, backend launch, privileged container, host networking, Docker socket mount, broad host mount, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add valid fixture, invalid digest, invalid path, unsigned, production-signature-blocked, and root-escape tests.
- Run the CW2 verification commands.
```

## CW3: Runtime State Root and Backend Lifecycle

### Mission

Track compatibility environment lifecycle under Runtime ownership without launching compatibility backends.

### Start from

- `internal/runtime/environment/`
- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/backend_selection.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `internal/runtime/appidentity/application_state_root.go`

### Deliver

- A Runtime state-root lifecycle store.
- States such as `missing`, `planned`, `staged`, `ready`, `repair-required`, and `blocked`.
- Readiness transitions driven by recipe trust, artifact staging, backend binding, Portal prerequisites, and snapshot prerequisites.
- KDE-safe readiness explanations that do not expose backend internals.

### Do not deliver

- Backend process starts.
- Launch transactions.
- Raw command exposure.
- Host path exposure.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/environment ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_application_state_root.rb
ruby -Ilib test/test_compatibility_backend_environment_plan.rb
ruby -Ilib test/test_compatibility_backend_binding.rb
ruby -Ilib test/test_compatibility_backend_lifecycle.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CW3 from docs/claude-code-windows-compatibility-workstreams.md.

Target outcome:
- Add durable Runtime state-root environment lifecycle records and readiness transitions without starting Wine, Proton, VM, or any other backend.

Hard constraints:
- Do not launch backends, create executable launch requests, expose raw commands or host paths, require network, use privileged containers, or mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add success, blocked, corrupted-state, transition, and root-escape tests.
- Run the CW3 verification commands.
```

## CW4: KDE Seven Entry Points

### Mission

Make the first-release KDE experience coherent across the seven official entry points while keeping KDE presentation-only.

### The seven entry points

1. Start menu.
2. Task manager.
3. File manager.
4. System tray.
5. Notification center.
6. AI Compatibility Center.
7. Unified settings.

### Start from

- `internal/runtime/appidentity/kde_*`
- `internal/runtime/appidentity/desktop_entry*.go`
- `internal/runtime/appidentity/file_association*.go`
- `internal/runtime/appidentity/task_manager_identity.go`
- `internal/runtime/appidentity/kwin_window_rule.go`
- `internal/runtime/appidentity/tray_status.go`
- `internal/runtime/appidentity/notification*.go`
- `internal/runtime/appidentity/settings*.go`
- `kde/`

### Deliver

- A documented Runtime read source or generated activation artifact for every KDE entry point.
- KDE-safe read models for Windows applications as normal desktop applications.
- File open, drag/drop, and sensitive desktop operations routed through Portal/review models.
- Notification and tray actions routed back to Runtime review or settings surfaces.

### Do not deliver

- KDE-owned backend policy.
- Direct Wine, Proton, or VM calls from KDE.
- Live tray bridging by default.
- Host KWin rule application.
- Production desktop writes.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_kde_integration_status.rb
ruby -Ilib test/test_kde_shell_integration_plan.rb
ruby -Ilib test/test_kde_application_surface_plan.rb
ruby -Ilib test/test_desktop_entry.rb
ruby -Ilib test/test_file_association_model.rb
ruby -Ilib test/test_task_manager_identity.rb
ruby -Ilib test/test_kwin_window_rule.rb
ruby -Ilib test/test_tray_status_model.rb
ruby -Ilib test/test_notification_request.rb
ruby -Ilib test/test_settings_model.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CW4 from docs/claude-code-windows-compatibility-workstreams.md.

Target outcome:
- Make KDE launcher, task-manager, Dolphin/file-manager, tray, notification, Compatibility Center, and settings reads consume Runtime-owned models or activation evidence while KDE remains presentation-only.

Hard constraints:
- Do not let KDE own compatibility policy, call Wine/Proton/VMs, write production desktop files, apply host KWin rules, start live tray bridges, or mutate the host root.
- Do not expose executable paths, backend commands, prefix terminology, state-root paths, host paths, file contents, secrets, or tokens in KDE-facing output.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add seven-entry-point success coverage plus missing-evidence, blocked-action, and backend-detail-redaction tests.
- Run the CW4 verification commands.
```

## CW5: Portal Permission Broker

### Mission

Create a Runtime-owned permission broker model for desktop-sensitive operations.

### Start from

- `internal/runtime/portal/`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/portal_request.go`
- `internal/runtime/appidentity/permission_review_plan.go`
- `internal/runtime/appidentity/review_flow_plan.go`
- `internal/runtime/appidentity/execution_preflight.go`

### Deliver

- Fake-mode request records for file, URI, print, clipboard, screenshot, camera, and remote-desktop operations.
- Completion, denial, timeout, and failure states.
- Permission review records that remain separate from execution approval.
- KDE-safe status and action projections.

### Do not deliver

- Real Portal calls by default.
- Direct permission grants.
- File-content reads.
- Host path exposure.
- Backend launch.

### Required verification

```text
go test ./internal/runtime/portal ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_compatibility_permission_review_plan.rb
ruby -Ilib test/test_compatibility_review_flow_plan.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CW5 from docs/claude-code-windows-compatibility-workstreams.md.

Target outcome:
- Add or strengthen fake-mode Runtime Portal permission request records and KDE-safe review states for file, URI, print, clipboard, screenshot, camera, and remote-desktop operations.

Hard constraints:
- Do not call real Portal transports, grant permissions, read file contents, expose host paths, launch backends, require network, or mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add allowed, denied, timeout, failed, redacted, and execution-preflight-blocked tests.
- Run the CW5 verification commands.
```

## CW6: Snapshot and Rollback Store

### Mission

Provide restore-point evidence for risky compatibility changes before execution and repair become real.

### Start from

- `internal/runtime/snapshot/`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/action_review_receipt.go`
- `internal/runtime/appidentity/desktop_activation_rollback.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/execution_preflight.go`

### Deliver

- Content-addressed snapshot records under explicit roots.
- Snapshot list, verify, and rollback evidence for test-controlled Runtime state.
- Rollback receipts that can be joined to repair and execution preflight.
- KDE-safe summaries without file contents or host paths.

### Do not deliver

- Host filesystem snapshots.
- Btrfs, ZFS, or system snapshot integration.
- Privileged rollback.
- Restore execution on host state.

### Required verification

```text
go test ./internal/runtime/snapshot ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_action_review_receipt.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CW6 from docs/claude-code-windows-compatibility-workstreams.md.

Target outcome:
- Add constrained snapshot and rollback evidence under explicit Runtime/test roots so repair and execution preflight can require verified restore points.

Hard constraints:
- Do not snapshot host roots, integrate privileged filesystem snapshots, restore host state, expose file contents or host paths, launch backends, or mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add create, list, verify, rollback, missing-snapshot, invalid-digest, and root-escape tests.
- Run the CW6 verification commands.
```

## CW7: AI Diagnostics and Repair Boundary

### Mission

Use Runtime evidence for AI-assisted compatibility diagnostics while keeping provider calls disabled and repair review-first.

### Start from

- `internal/runtime/diagnostics/`
- `internal/runtime/appidentity/diagnostics*.go`
- `internal/runtime/appidentity/ai_diagnostic_*.go`
- `internal/runtime/appidentity/ai_repair_approval_gate.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/test_plan.go`
- `internal/runtime/appidentity/test_result.go`

### Deliver

- Fixture diagnostic run records under controlled state roots.
- Redacted AI diagnostic input records.
- Disabled-by-default provider boundary.
- Review-only repair recommendations.
- Compatibility Center summaries that avoid raw paths, file contents, backend commands, secrets, and tokens.

### Do not deliver

- Real AI provider calls by default.
- Auto-repair.
- File-content reads.
- Network requirements.
- Backend launch.

### Required verification

```text
go test ./internal/runtime/diagnostics ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_test_plan.rb
ruby -Ilib test/test_compatibility_test_result.rb
ruby -Ilib test/test_compatibility_repair_plan.rb
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_ai_repair_approval_gate.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CW7 from docs/claude-code-windows-compatibility-workstreams.md.

Target outcome:
- Add or strengthen fixture diagnostics, redacted AI inputs, disabled-provider behavior, and review-only repair recommendations.

Hard constraints:
- Do not call real AI providers, auto-repair, read user file contents, require network, launch backends, expose secrets or host paths, or mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add success, failure, redaction, provider-disabled, approval-required, and snapshot-required tests.
- Run the CW7 verification commands.
```

## CW8: Execution Transaction Ledger

### Mission

Connect launch intent, review, permissions, snapshots, lifecycle readiness, resource grants, transactions, and session status into a blocked-by-default execution ledger.

### Start from

- `internal/runtime/execution/`
- `internal/runtime/appidentity/launch_intent.go`
- `internal/runtime/appidentity/execution_request.go`
- `internal/runtime/appidentity/execution_review.go`
- `internal/runtime/appidentity/execution_decision.go`
- `internal/runtime/appidentity/execution_preflight.go`
- `internal/runtime/appidentity/execution_resource_grant.go`
- `internal/runtime/appidentity/execution_transaction.go`
- `internal/runtime/appidentity/execution_session*.go`

### Deliver

- Execution request IDs and transaction records under explicit roots.
- Separate records for user review, permission state, snapshot state, backend lifecycle readiness, and resource grants.
- Deterministic blocked reasons.
- Session status projections that remain preview-only.

### Do not deliver

- Real process launch.
- Wine, Proton, VM, or backend starts.
- Permission grant bypasses.
- Snapshot bypasses.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/execution ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_launch_request.rb
ruby -Ilib test/test_compatibility_run_plan.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CW8 from docs/claude-code-windows-compatibility-workstreams.md.

Target outcome:
- Add or strengthen a blocked-by-default execution transaction ledger that joins review, Portal permission, snapshot, backend lifecycle, and resource-grant evidence without launching any backend.

Hard constraints:
- Do not start Wine, Proton, VM, or any backend; do not bypass review, permission, snapshot, or lifecycle gates; do not expose raw backend commands or host paths; do not mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add approved-but-blocked, missing-permission, missing-snapshot, backend-not-ready, denied-review, root-escape, and session-preview tests.
- Run the CW8 verification commands.
```

## CW9: KDE Materialization Writer

### Mission

Turn Runtime activation plans into deterministic desktop artifacts under explicit target roots.

### Start from

- `internal/runtime/appidentity/desktop_activation_*.go`
- `internal/runtime/appidentity/desktop_entry*.go`
- `internal/runtime/appidentity/desktop_icon*.go`
- `internal/runtime/appidentity/file_association*.go`
- `internal/runtime/appidentity/desktop_resource_bridge.go`
- `kde/`

### Deliver

- Target-root-only desktop entry writer.
- Target-root-only MIME association writer.
- Target-root-only Dolphin service-menu writer.
- Target-root-only icon and manifest writer.
- Activation and rollback receipts.
- Drift checks between receipts and KDE-safe read models.

### Do not deliver

- Host desktop database mutation.
- Host MIME default changes.
- Service installation.
- Notification delivery.
- KWin rule application.
- Backend launch.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby -Ilib test/test_desktop_entry.rb
ruby -Ilib test/test_file_association_model.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CW9 from docs/claude-code-windows-compatibility-workstreams.md.

Target outcome:
- Add or strengthen target-root-only KDE materialization for desktop entries, MIME records, Dolphin service menus, icons, activation manifests, and rollback receipts.

Hard constraints:
- Do not mutate the host desktop database, write host MIME defaults, install services, deliver notifications, apply KWin rules, start backends, or mutate host roots.
- Generated artifacts must expose normal application identity only, without executable paths, backend commands, prefix terminology, state-root paths, host paths, file contents, secrets, or tokens.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add deterministic-write, rollback, unchanged-file, modified-file-preservation, missing-receipt, digest-mismatch, and root-escape tests.
- Run the CW9 verification commands.
```

## CW10: Evidence and Drift Harness

### Mission

Prevent Xnix from accumulating more contract-only Windows compatibility surfaces without implementation evidence.

### Start from

- `scripts/implementation_evidence_report.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/verify_layout.rb`
- `test/test_implementation_evidence_report.rb`
- `test/test_runtime_contract_drift_report.rb`
- `docs/claude-code-*.md`

### Deliver

- Orphan contract detection for Runtime reads, CLI commands, smoke adapters, docs, and KDE entry points.
- Evidence-level reporting for contract-only, fixture-implemented, state-root-implemented, smoke-owned, and production-gated surfaces.
- Assignment ownership checks that map new files and changed surfaces to expected workstreams.
- JSON and Markdown reports suitable for CI.

### Do not deliver

- Product behavior changes.
- QEMU requirements for normal report runs.
- Docker requirements for normal report runs.
- Network requirements.
- Host-root access.

### Required verification

```text
go test ./...
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/implementation_evidence_report.rb --format markdown
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/runtime_contract_drift_report.rb --format markdown
ruby -Ilib test/test_implementation_evidence_report.rb
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CW10 from docs/claude-code-windows-compatibility-workstreams.md.

Target outcome:
- Add CI-friendly evidence and drift gates that fail on orphan Runtime contracts, KDE entry-point drift, preview-only regressions, and missing workstream ownership.

Hard constraints:
- Do not require Docker, QEMU, network, backend launch, privileged containers, or host-root access for ordinary report runs.
- Do not change product behavior except where fixtures are required for report evidence.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add deterministic-output, missing-evidence, orphan-contract, KDE-entrypoint-drift, and workstream-ownership tests.
- Run the CW10 verification commands.
```

## CW11: Product Image and QEMU Acceptance

### Mission

Prove the KDE-first Windows compatibility product shape in a constrained image/QEMU smoke without weakening host safety.

### Start from

- `boot/`
- `buildroot/`
- `scripts/container.rb`
- `scripts/full_smoke.rb`
- `scripts/qemu_smoke.rb`
- `docs/kde-first-compatibility-acceptance.md`
- `docs/kde-first-presence-smoke-spec.md`

### Deliver

- Product-image metadata for the KDE-first path.
- Acceptance checks for Runtime files, D-Bus service definitions, KDE integration assets, seven entry-point evidence, and safety defaults.
- Persisted QEMU serial logs.
- Loopback-only SSH smoke, if SSH is part of the check.
- Clear separation between the Buildroot learning baseline and the KDE-first product image path.

### Do not deliver

- Host package installation.
- Host network exposure.
- Privileged containers.
- Docker socket mounts.
- Broad host mounts.
- Production backend launch.
- SSH exposure beyond loopback-bound forwarded ports.

### Required verification

```text
ruby scripts/verify_layout.rb
ruby scripts/container.rb build
ruby scripts/full_smoke.rb
```

Run the repository's full QEMU milestone gate if the version lands on a tenth code version.

### Copyable prompt

```text
Implement CW11 from docs/claude-code-windows-compatibility-workstreams.md.

Target outcome:
- Add or strengthen a reproducible KDE-first product image/QEMU acceptance path with persisted smoke evidence and strict host-safety defaults.

Hard constraints:
- Do not install host packages, expose host networking, use privileged containers, mount the Docker socket, broadly mount host directories, launch production backends, or mutate the host root.
- Keep SSH smoke exposure loopback-bound and key-authenticated.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add build-plan, host-risk-denial, serial-log, loopback-SSH, Runtime-presence, KDE-asset-presence, and smoke-report tests.
- Run the CW11 verification commands.
- If the resulting version is a tenth code version, run the full QEMU milestone gate.
```

## Completion Statement Template

Each Claude Code branch should finish with this statement:

```text
Workstream:
- <CW id and name>

Converted contracts:
- <contract or preview surfaces that now have implementation evidence>

Implementation evidence added:
- <fixture/state-root/smoke/report evidence>

Still gated:
- <unsafe production behavior that remains disabled>

Verification run:
- <commands and results>

Files intentionally not touched:
- docs/claude-code-implementation-packages.md
- <other unrelated areas>
```
