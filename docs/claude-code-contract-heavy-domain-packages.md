# Claude Code Contract-Heavy Domain Packages

> Last updated: 2026-07-16 | Baseline: v0.2.293

This document splits the large contract-heavy areas of Xnix into independently assignable implementation packages for Claude Code.

Use this file when a domain already has contracts, previews, routes, or tests, but still lacks durable implementation evidence.

This file is intentionally a dispatcher. It does not replace:

- `docs/claude-code-implementation-packages.md`
- `docs/claude-code-empty-domain-implementation-packages.md`
- `docs/claude-code-windows-compatibility-workstreams.md`
- `docs/claude-code-mainline-implementation-plan.md`

## Assignment Rules

Assign one package per Claude Code branch.

Every branch must:

- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Runtime product logic.
- Use C only for low-level Runtime surfaces, ABI-shaped contracts, or already-owned C policy records.
- Use Ruby for tests, smoke scripts, developer tooling, and reports.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for code changes.
- Add success, failure, and blocked-state tests.
- Run package-specific verification plus `ruby scripts/verify_layout.rb`.
- Stop and report blockers before adding network fetch, production D-Bus ownership, real Portal calls, real backend launch, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.

Unsafe behavior remains disabled unless a package explicitly says otherwise:

- Production D-Bus ownership.
- Runtime write methods.
- Real Wine, Proton, or Windows VM launch.
- Real XDG Portal transport calls.
- Default network artifact fetch.
- Host package-manager calls.
- Host-root mutation.
- Raw backend commands in KDE/user-facing output.
- Raw executable paths, compatibility storage paths, secrets, tokens, private keys, or file contents in KDE/user-facing output.

## Package Index

| ID | Package | Primary owner | What stops being contract-only |
| --- | --- | --- | --- |
| `CH1` | Runtime owner service boundary | Go Runtime | Read-only owner process, route parity, disabled writes |
| `CH2` | Recipe, trust, artifact, and staging pipeline | Go Runtime | Verified inputs, local cache/stage receipts, install gates |
| `CH3` | Backend inventory and lifecycle state | Go Runtime | Persisted Wine/Proton/VM inventory and lifecycle transitions |
| `CH4` | Portal permission and snapshot safety plane | Go Runtime | Fake Portal records, permission decisions, snapshot receipts |
| `CH5` | KDE materialization and activation receipts | Go Runtime plus KDE files | Target-root desktop, MIME, service-menu, icon, rollback evidence |
| `CH6` | Execution transaction ledger | Go Runtime | Reviewed launch transactions that still block real execution |
| `CH7` | Diagnostics, repair, and AI boundary | Go Runtime | Redacted diagnostic records, repair recommendations, AI gates |
| `CH8` | Product image and QEMU acceptance | Build/image tooling | Atomic KDE image smoke and loopback-only SSH evidence |
| `CH9` | Evidence and drift harness | Ruby reports plus Go fixtures | CI-friendly detection of orphan contracts and preview-only regressions |
| `CH10` | Runtime packaging and service binding | Go Runtime plus system files | Installable daemon shape without production enablement |

Recommended first wave:

1. `CH1` Runtime owner service boundary.
2. `CH2` Recipe, trust, artifact, and staging pipeline.
3. `CH3` Backend inventory and lifecycle state.
4. `CH9` Evidence and drift harness.

Do not assign `CH6` real execution work before `CH2`, `CH3`, and `CH4` have durable blocking evidence.

## Implementation Evidence Standard

A package is not complete just because it adds a new model, preview, CLI command, D-Bus method, or documentation section.

Each package must add at least one durable evidence type:

- A state-root record with schema version, operation id, timestamps, status, blocked reasons, and relative receipt paths.
- A local fixture pipeline with digest verification and deterministic mismatch failures.
- A constrained smoke path that exercises the real adapter boundary without production privileges.
- A KDE-safe read model that consumes durable receipts rather than restating static preview data.
- A report gate that fails when a contracted domain remains orphaned or preview-only.

## CH1: Runtime Owner Service Boundary

### Mission

Turn the Runtime owner from static previews and local dispatch helpers into a constrained Go-owned read boundary.

### Start from

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `internal/runtime/appidentity/runtime_owner_*.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `runtime/dbus/xnix_compatd_smoke.c`
- `scripts/dbus_session_smoke.rb`
- `scripts/runtime_contract_drift_report.rb`

### Deliver

- A long-running smoke-owner mode that can claim a private session-bus name inside constrained tests.
- Read-only Runtime method routing through Go owner handlers or stable unsupported-read markers.
- Deterministic disabled responses for every write method.
- Owner readiness states for `preview-only`, `smoke-owner`, and `production-owner-blocked`.
- Route parity evidence across D-Bus XML, owner routes, smoke adapters, Ruby clients, and Go CLI routes.

### Keep disabled

- Production D-Bus ownership.
- Runtime write methods.
- System service installation.
- Backend launch.
- Host-root mutation.
- KDE-owned Runtime policy.

### Verification

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
Implement CH1 from docs/claude-code-contract-heavy-domain-packages.md.

Target outcome:
- Add constrained Go Runtime owner service evidence for read-only Runtime methods while write methods fail closed.

Hard constraints:
- No production D-Bus ownership, write enablement, backend launch, system service installation, host-root mutation, privileged containers, host networking, Docker socket mounts, broad host mounts, or raw backend details in KDE-facing output.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add success, unsupported-read, disabled-write, lifecycle, and blocked-production-owner tests.
- Run the CH1 verification commands and report exact output.
```

## CH2: Recipe, Trust, Artifact, and Staging Pipeline

### Mission

Make compatibility inputs verifiable before any desktop activation or execution package depends on them.

### Start from

- `internal/runtime/recipe/`
- `internal/runtime/artifact/`
- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `internal/runtime/appidentity/package_source.go`
- `internal/runtime/appidentity/acquisition_preflight.go`
- `internal/runtime/appidentity/artifact_manifest.go`
- `internal/runtime/appidentity/install_plan.go`
- `runtime/recipes/`

### Deliver

- Read-only recipe store roots with schema, path, id, digest, and signing-state validation.
- A replaceable signature verifier boundary without committed keys or secrets.
- Local fixture artifact manifests with SHA-256 verification.
- Runtime state-root scoped cache and staging namespaces.
- Staging receipts that record relative paths, digests, source metadata, and blocked reasons.
- Install readiness that joins recipe trust, artifact verification, staging state, owner readiness, and backend lifecycle state.

### Keep disabled

- Default network downloads.
- Host package-manager calls.
- Private key material.
- Production trust without verifier evidence.
- Host-root writes.
- Backend launch.

### Verification

```text
go test ./internal/runtime/recipe ./internal/runtime/artifact ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_recipe_registry.rb
ruby -Ilib test/test_recipe_trust_policy.rb
ruby -Ilib test/test_recipe_install_gate.rb
ruby -Ilib test/test_compatibility_package_source.rb
ruby -Ilib test/test_compatibility_artifact_manifest.rb
ruby -Ilib test/test_compatibility_install_plan.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CH2 from docs/claude-code-contract-heavy-domain-packages.md.

Target outcome:
- Convert recipe trust, artifact verification, local cache planning, and staging into durable state-root evidence.

Hard constraints:
- No default network fetch, package-manager invocation, private keys, production trust claim, host-root mutation, backend launch, privileged containers, host networking, Docker socket mounts, or broad host mounts.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add digest success, digest mismatch, unsigned/invalid trust, blocked install, and state-root receipt tests.
- Run the CH2 verification commands and report exact output.
```

## CH3: Backend Inventory and Lifecycle State

### Mission

Turn backend selection and lifecycle contracts into persisted Runtime-owned state without starting Wine, Proton, or a Windows VM.

### Start from

- `internal/runtime/appidentity/backend_manager.go`
- `internal/runtime/appidentity/backend_selection.go`
- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `internal/runtime/environment/`
- `cmd/xnix-runtime-go/`

### Deliver

- A Runtime-internal backend inventory for Wine, Proton, Windows VM, and future backend ids.
- State-root records for inventory, selection intent, environment lifecycle, and binding readiness.
- Lifecycle transitions such as `missing`, `planned`, `staged`, `ready`, `repair-required`, and `blocked`.
- Backend binding explanations that are safe for KDE/user-facing output.
- Execution readiness that changes when recipe trust, staged artifacts, permissions, or lifecycle state changes.

### Keep disabled

- Backend installation.
- Backend download.
- Wine, Proton, or VM process launch.
- Raw backend commands.
- Raw profile/prefix paths.
- Backend detail exposure in KDE-facing output.
- Host-root mutation.

### Verification

```text
go test ./internal/runtime/environment ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_backend_selection_plan.rb
ruby -Ilib test/test_compatibility_backend_environment_plan.rb
ruby -Ilib test/test_compatibility_backend_binding.rb
ruby -Ilib test/test_compatibility_backend_lifecycle.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CH3 from docs/claude-code-contract-heavy-domain-packages.md.

Target outcome:
- Add persisted Runtime state-root backend inventory and lifecycle evidence for Wine, Proton, and Windows VM without starting any backend.

Hard constraints:
- No backend install, download, launch, VM start, raw command exposure, profile path exposure, host-root mutation, network requirement, privileged container, host networking, Docker socket mount, or broad host mount.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add state-root inventory, lifecycle transition, blocked binding, and KDE-safe output tests.
- Run the CH3 verification commands and report exact output.
```

## CH4: Portal Permission and Snapshot Safety Plane

### Mission

Build the safety plane that execution and KDE settings must rely on before any launch path becomes real.

### Start from

- `internal/runtime/portal/`
- `internal/runtime/snapshot/`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/execution_resource_grant.go`
- `internal/runtime/appidentity/snapshot_plan.go`
- `lib/xnix/compatibility/portal_*`
- `lib/xnix/compatibility/compatibility_snapshot_plan.rb`

### Deliver

- A fake-mode Portal request state machine with request, review, granted, denied, expired, and blocked states.
- Permission records scoped to application id, operation id, resource kind, and user decision.
- Snapshot receipt records under an explicit state root.
- Snapshot rollback preflight that verifies digest and ownership before claiming rollback readiness.
- A joined safety read model for KDE settings and Compatibility Center cards.

### Keep disabled

- Real Portal transport calls.
- Direct file, URI, print, clipboard, screen, camera, or remote-desktop access.
- Snapshot writes outside explicit state roots.
- Rollback execution against the host.
- Backend launch.

### Verification

```text
go test ./internal/runtime/portal ./internal/runtime/snapshot ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_compatibility_permission_review_plan.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CH4 from docs/claude-code-contract-heavy-domain-packages.md.

Target outcome:
- Add fake-mode Portal permission records and state-root snapshot receipts that execution can depend on later.

Hard constraints:
- No real Portal calls, direct desktop resource access, host rollback, backend launch, host-root mutation, privileged containers, host networking, Docker socket mounts, or broad host mounts.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add allow, deny, expired, blocked, snapshot digest, and rollback-preflight tests.
- Run the CH4 verification commands and report exact output.
```

## CH5: KDE Materialization and Activation Receipts

### Mission

Move KDE integration from previews toward target-root materialization that is reversible and receipt-verified.

### Start from

- `internal/runtime/appidentity/desktop_activation_*.go`
- `internal/runtime/appidentity/kde_*`
- `kde/`
- `bin/xnix-install-desktop-integration`
- `bin/xnix-rollback-desktop-integration`
- `lib/xnix/compatibility/desktop_*`
- `lib/xnix/compatibility/file_association_model.rb`

### Deliver

- Target-root writers for desktop entries, MIME association files, Dolphin service menus, icons, and activation manifests.
- Receipts with relative paths, file modes, SHA-256 digests, renderer versions, and rollback eligibility.
- Drift checks between expected activation artifacts and materialized files.
- Rollback that removes only unchanged receipt-owned files.
- KDE status models that consume activation receipts instead of only static previews.

### Keep disabled

- Host-root activation by default.
- KDE cache refresh on the host.
- Launcher execution.
- Backend launch.
- Raw command/path exposure.
- Unchecked rollback.

### Verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby -Ilib test/test_file_association_model.rb
ruby -Ilib test/test_dolphin_service_menu.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CH5 from docs/claude-code-contract-heavy-domain-packages.md.

Target outcome:
- Add target-root KDE materialization and rollback receipt evidence for launcher, MIME, Dolphin, icon, and manifest artifacts.

Hard constraints:
- No host-root activation by default, host KDE cache refresh, launch, backend launch, raw path exposure, raw backend command exposure, privileged containers, host networking, Docker socket mounts, or broad host mounts.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add write, digest, drift, rollback, unchanged-file, changed-file, and KDE status consumption tests.
- Run the CH5 verification commands and report exact output.
```

## CH6: Execution Transaction Ledger

### Mission

Prepare execution as an auditable transaction system while keeping real backend launch disabled.

### Start from

- `internal/runtime/execution/`
- `internal/runtime/appidentity/execution_*.go`
- `internal/runtime/appidentity/launch_intent.go`
- `internal/runtime/appidentity/run_plan.go`
- `cmd/xnix-runtime-go/execution_ledger_commands.go`

### Deliver

- Reviewed execution transactions under an explicit state root.
- Joined preflight gates from recipe trust, artifact staging, backend lifecycle, Portal permissions, snapshots, and user review.
- Session records that can represent `blocked`, `ready-for-test`, `launch-disabled`, and `completed-without-launch`.
- CLI commands to create, inspect, list, and explain transactions.
- KDE-safe execution summaries for Compatibility Center cards.

### Keep disabled

- Real launch.
- Backend process start.
- VM start.
- Direct file access.
- Raw command output in KDE/user-facing output.
- Host-root mutation.

### Verification

```text
go test ./internal/runtime/execution ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_run_plan.rb
ruby -Ilib test/test_launch_request.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CH6 from docs/claude-code-contract-heavy-domain-packages.md.

Target outcome:
- Add state-root execution transaction ledger evidence that joins safety gates but keeps real launch disabled.

Hard constraints:
- No real launch, backend process start, VM start, direct desktop resource access, raw command exposure, host-root mutation, privileged containers, host networking, Docker socket mounts, or broad host mounts.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add create, inspect, list, blocked, ready-with-launch-disabled, missing-permission, missing-snapshot, and KDE-safe summary tests.
- Run the CH6 verification commands and report exact output.
```

## CH7: Diagnostics, Repair, and AI Boundary

### Mission

Make diagnostics and repair reproducible, redacted, and approval-gated before any AI-assisted repair can affect state.

### Start from

- `internal/runtime/diagnostics/`
- `internal/runtime/appidentity/diagnostics.go`
- `internal/runtime/appidentity/ai_diagnostics.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/kde_ai_analysis.go`
- `internal/runtime/appidentity/kde_action_review.go`
- `internal/runtime/appidentity/kde_action_receipt.go`

### Deliver

- Fixture diagnostic runners with deterministic success and failure records.
- Redacted diagnostic input records that exclude secrets, file contents, raw executable paths, and backend commands.
- Repair recommendation records that require user review.
- AI provider policy states such as `disabled`, `fake-fixture`, `configured-but-blocked`, and `production-gated`.
- Compatibility Center consumption of diagnostic history and repair receipts.

### Keep disabled

- Real external AI calls by default.
- Automatic repair.
- Settings persistence.
- File mutation.
- Backend launch.
- Secret logging.

### Verification

```text
go test ./internal/runtime/diagnostics ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_ai_repair_approval_gate.rb
ruby -Ilib test/test_compatibility_repair_plan.rb
ruby -Ilib test/test_action_review_receipt.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CH7 from docs/claude-code-contract-heavy-domain-packages.md.

Target outcome:
- Add reproducible diagnostic records, redacted AI inputs, review-only repair recommendations, and Compatibility Center history consumption.

Hard constraints:
- No real external AI call by default, automatic repair, file mutation, backend launch, settings persistence, secret logging, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add redaction, fake-provider, provider-disabled, repair-review, receipt, and history-consumption tests.
- Run the CH7 verification commands and report exact output.
```

## CH8: Product Image and QEMU Acceptance

### Mission

Keep the product honest by proving that the image can boot and expose observable services under constrained test conditions.

### Start from

- `image/kinoite/`
- `buildroot/`
- `boot/`
- `lib/xnix/image/`
- `lib/xnix/qemu.rb`
- `lib/xnix/sshd.rb`
- `scripts/build_kde_image.rb`
- `scripts/boot_kde_image.rb`
- `scripts/full_smoke.rb`

### Deliver

- Image manifest validation for KDE-first product components.
- QEMU smoke report records with persisted serial logs.
- Loopback-only SSH forwarding checks.
- Clear split between Buildroot learning baseline and atomic KDE product image.
- Safe cleanup that only removes owned test artifacts.

### Keep disabled

- Host networking.
- Privileged containers.
- Docker socket mounts.
- Broad host mounts.
- Host-root writes.
- Non-loopback SSH exposure.
- Unbounded cleanup.

### Verification

```text
ruby -Ilib test/test_kde_image.rb
ruby -Ilib test/test_kde_disk.rb
ruby -Ilib test/test_qemu.rb
ruby -Ilib test/test_sshd.rb
ruby -Ilib test/test_full_smoke_report.rb
ruby scripts/container.rb full-smoke
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CH8 from docs/claude-code-contract-heavy-domain-packages.md.

Target outcome:
- Add or strengthen constrained product-image and QEMU smoke evidence with persisted logs and loopback-only SSH checks.

Hard constraints:
- No privileged containers, host networking, Docker socket mounts, broad host mounts, host-root mutation, non-loopback SSH exposure, unsafe cleanup, or host package-manager calls.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add image-manifest, QEMU command, serial-log, loopback-SSH, failure-report, and cleanup tests.
- Run the CH8 verification commands and report exact output.
```

## CH9: Evidence and Drift Harness

### Mission

Prevent the project from accumulating more contracts without implementation evidence.

### Start from

- `scripts/runtime_contract_drift_report.rb`
- `scripts/implementation_evidence_report.rb`
- `scripts/verify_layout.rb`
- `test/test_runtime_contract_drift_report.rb`
- `test/test_implementation_evidence_report.rb`
- `docs/claude-code-mainline-implementation-plan.md`
- `docs/claude-code-windows-compatibility-workstreams.md`

### Deliver

- Orphan contract detection for Runtime read methods, CLI commands, D-Bus routes, Ruby client methods, and smoke adapter routes.
- Evidence-level classification for each mainline and Windows compatibility workstream.
- CI-friendly JSON and Markdown outputs.
- Failure gates for domains that regress from state-root, fixture, or smoke evidence back to preview-only.
- A next-dispatch section that recommends the next branch-sized Claude Code package.

### Keep disabled

- Network access.
- Docker/QEMU execution as part of report generation.
- Host-root mutation.
- Secret scanning output that prints secret values.

### Verification

```text
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby -Ilib test/test_implementation_evidence_report.rb
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/runtime_contract_drift_report.rb --format markdown
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/implementation_evidence_report.rb --format markdown
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CH9 from docs/claude-code-contract-heavy-domain-packages.md.

Target outcome:
- Strengthen evidence and drift reports so contract-only Runtime, KDE, backend, Portal, execution, and diagnostic surfaces cannot silently grow.

Hard constraints:
- No Docker, QEMU, network, host-root mutation, privileged containers, host networking, Docker socket mounts, broad host mounts, or secret-value output during report generation.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add orphan-contract, evidence-regression, JSON, Markdown, and next-dispatch tests.
- Run the CH9 verification commands and report exact output.
```

## CH10: Runtime Packaging and Service Binding

### Mission

Shape the Runtime daemon as an installable system component without enabling production ownership or writes prematurely.

### Start from

- `runtime/systemd/xnix-compatd.service`
- `runtime/dbus/org.xnix.Compatibility1.service`
- `libexec/xnix/compatd`
- `internal/runtime/appidentity/runtime_service_binding.go`
- `internal/runtime/appidentity/runtime_owner_process.go`
- `internal/runtime/activation/`
- `scripts/install_runtime_activation.rb`
- `scripts/runtime_activation_smoke.rb`

### Deliver

- Service binding records that explain installed, staged, smoke, and disabled daemon states.
- Target-root service file staging with receipts and rollback metadata.
- Versioned Runtime binary/library activation records.
- Smoke-mode service binding checks that do not install into the host.
- Clear blocked state for production enablement.

### Keep disabled

- Host systemd installation.
- Host D-Bus service installation.
- Production bus ownership.
- Runtime write methods.
- Backend launch.
- Host-root mutation by default.

### Verification

```text
go test ./internal/runtime/activation ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_runtime_activation.rb
ruby -Ilib test/test_runtime_activation_install.rb
ruby -Ilib test/test_runtime_activation_smoke_script.rb
ruby -Ilib test/test_runtime_service_binding.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement CH10 from docs/claude-code-contract-heavy-domain-packages.md.

Target outcome:
- Add target-root Runtime packaging and service-binding receipts while production installation remains blocked.

Hard constraints:
- No host systemd installation, host D-Bus installation, production bus ownership, write enablement, backend launch, host-root mutation, privileged containers, host networking, Docker socket mounts, or broad host mounts.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add target-root staging, receipt, rollback, version, smoke-binding, and production-blocked tests.
- Run the CH10 verification commands and report exact output.
```

## Branch Handoff Template

Copy this template when assigning any package:

```text
Implement <PACKAGE_ID> from docs/claude-code-contract-heavy-domain-packages.md.

Target outcome:
- <one durable implementation outcome>

Start from:
- <key files>

Hard constraints:
- Keep production D-Bus ownership, Runtime write methods, real backend launch, real Portal calls, default network fetch, privileged containers, host networking, Docker socket mounts, broad host mounts, and host-root mutation disabled unless the package explicitly enables a constrained test-only path.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Runtime logic and Ruby for tests/tooling.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add success, failure, and blocked-state tests.
- Run package verification plus ruby scripts/verify_layout.rb.
- Report exact commands and results.
- Stop and report blockers before widening scope or touching unrelated packages.
```
