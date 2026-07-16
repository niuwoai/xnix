# Claude Code Mainline Empty-Domain Handoff

> Last updated: 2026-07-16 | Baseline: v0.2.253

This document is a compact handoff board for assigning large but independent Xnix implementation chunks to Claude Code.

Use it when the current state is: the contract, preview, CLI shape, or test expectation exists, but the domain still lacks durable implementation evidence.

This file intentionally does not replace the larger planning documents. It is the short practical entry point for choosing one branch-sized implementation task.

## Global Instructions for Claude Code

Every branch must:

- Implement exactly one package from this document.
- Prefer Go for durable Compatibility Runtime behavior.
- Use C only for low-level transport, ABI-shaped records, or already-owned C Runtime policy surfaces.
- Use Ruby for tests, smoke scripts, reports, and developer tooling.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Add targeted success, failure, and blocked-state tests.
- Run package-specific tests plus `ruby scripts/verify_layout.rb`.
- Stop and report if the package would require unsafe behavior outside its scope.

Unsafe behavior remains disabled unless a future package explicitly enables a constrained test-only path:

- Production D-Bus ownership.
- Runtime write methods.
- Real Wine, Proton, VM, or compatibility backend launch.
- Real XDG Desktop Portal transport calls.
- Default network artifact fetch.
- Host package-manager calls.
- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Host-root mutation.
- Raw backend commands, executable paths, compatibility storage paths, profile terminology, file contents, secrets, tokens, private keys, or host paths in KDE-facing output.

## Recommended First Wave

Start with these packages in order:

| Order | Package | Why |
| --- | --- | --- |
| 1 | `H1` Runtime owner read service | Gives later work a real Runtime-owned process boundary. |
| 2 | `H2` Recipe and artifact trust pipeline | Makes install inputs verifiable before activation or execution depends on them. |
| 3 | `H3` Runtime state root and backend lifecycle | Gives launch, diagnostics, and KDE views durable safe state to consume. |
| 4 | `H8` Evidence and drift harness | Prevents more contract-only surfaces from accumulating silently. |

Do not start `H6` execution transactions until `H2`, `H3`, and `H4` can produce blocking safety evidence.

Do not start `H9` image acceptance until the image smoke can validate real Runtime evidence instead of only static previews.

## Package Index

| ID | Package | Can start now | Primary area | Mergeable evidence |
| --- | --- | --- | --- | --- |
| H1 | Runtime owner read service | Yes | Go Runtime and D-Bus smoke | Private session-bus owner serves read-only methods and fails writes closed. |
| H2 | Recipe and artifact trust pipeline | Yes | Go Runtime trust and artifact code | Local recipes and fixture artifacts verify digests and produce safe receipts. |
| H3 | Runtime state root and backend lifecycle | Yes | Go Runtime state | Lifecycle records persist under explicit roots without backend launch. |
| H4 | Portal, snapshot, and rollback safety | Yes | Go Runtime safety plane | Fake Portal records and snapshot receipts prove review-first safety. |
| H5 | KDE activation materialization consumers | Yes, staged roots only | KDE-facing Runtime reads | KDE entry points consume staged receipts rather than static previews. |
| H6 | Execution transaction ledger | Later | Go Runtime execution safety | Reviewed transaction records still block launch. |
| H7 | Diagnostics, repair, and AI boundary | Yes, no-provider mode | Go Runtime diagnostics | Fixture diagnostics and review-only recommendations are recorded safely. |
| H8 | Evidence and drift harness | Yes | Ruby tooling plus Go fixtures | Reports fail on orphan contracts and preview-only regressions. |
| H9 | Atomic KDE image and QEMU acceptance | Later | Image tooling and smoke tests | Constrained image smoke validates KDE-first Runtime presence. |

## H1: Runtime Owner Read Service

### Mission

Turn the Runtime owner from previews and adapters into a constrained Go owner process that can claim a private test session-bus name, serve read-only Runtime methods, and fail write methods closed.

### Start from

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `internal/runtime/appidentity/runtime_owner_*.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `runtime/dbus/xnix_compatd_smoke.c`
- `scripts/dbus_session_smoke.rb`
- `scripts/runtime_owner_candidate_smoke.rb`
- `scripts/runtime_contract_drift_report.rb`

### Deliver

- Long-running smoke-owner mode.
- Private session-bus ownership only in constrained tests.
- Read-only D-Bus routing through Go owner handlers.
- Explicit unsupported-read results for intentionally unimplemented routes.
- Stable disabled responses for every write method.
- Safe lifecycle logs for startup, route table, bus mode, readiness, dispatch result, and shutdown.

### Acceptance

- A constrained container can start the owner on a private session bus.
- Representative D-Bus reads return safe Go Runtime payloads.
- Every write method returns a stable disabled error.
- Contract drift reporting stays clean.

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
Implement H1 from docs/claude-code-mainline-empty-domain-handoff.md.

Target outcome:
- Add a constrained Go Runtime owner service that can claim a private test session-bus name, serve read-only Runtime methods through the existing owner route table, and fail every write method closed.

Hard constraints:
- Do not enable production D-Bus ownership, system service installation, write methods, backend launch, network fetch, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add success, unsupported-read, disabled-write, lifecycle, and blocked-production-owner tests.
- Run the H1 verification commands.
```

## H2: Recipe and Artifact Trust Pipeline

### Mission

Turn recipe trust, package sources, artifact manifests, cache planning, staging receipts, and install readiness into a local-only pipeline that fails closed.

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
- `cmd/xnix-runtime-go/`

### Deliver

- Read-only recipe store with local roots.
- Registry validation for schema, IDs, relative paths, SHA-256 digests, and declared signing state.
- Replaceable signature verifier boundary without committed production keys.
- Fixture artifact manifest parsing and digest verification.
- Runtime-root-scoped cache and staging receipts.
- Install readiness that joins recipe trust, artifact verification, cache state, staging state, and owner readiness.

### Acceptance

- Invalid recipe digests fail closed.
- Invalid artifact digests block staging.
- Development fixtures remain usable but cannot pass production trust.
- Cache and staging paths stay under explicit Runtime or test roots.
- KDE-facing output exposes relative receipts and user-safe status only.

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
Implement H2 from docs/claude-code-mainline-empty-domain-handoff.md.

Target outcome:
- Build a local-only recipe, artifact, cache, staging, and install-readiness trust pipeline that verifies digests, stages only under explicit roots, and fails closed for invalid inputs.

Hard constraints:
- No default network fetch, host package-manager calls, private keys, production trust bypasses, desktop writes, backend launch, privileged containers, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add valid fixture, invalid digest, invalid path, production-signature-blocked, and root-escape tests.
- Run the H2 verification commands.
```

## H3: Runtime State Root and Backend Lifecycle

### Mission

Create durable Runtime state-root records for application environment lifecycle without creating environments or starting backends.

### Start from

- `internal/runtime/environment/`
- `internal/runtime/appidentity/state_root.go`
- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `cmd/xnix-runtime-go/`

### Deliver

- State-root schema for application lifecycle records.
- Lifecycle states such as `missing`, `planned`, `staged`, `ready`, `repair-required`, `blocked`, and `retired`.
- Atomic writes under explicit state roots.
- Relative receipt paths in all user-facing output.
- Read models consumed by backend binding, execution readiness, diagnostics, and owner readiness.

### Acceptance

- State-root writes cannot escape the configured root.
- Lifecycle transitions reject invalid jumps.
- Backend launch remains disabled.
- KDE-facing output never exposes state-root paths or backend internals.

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
Implement H3 from docs/claude-code-mainline-empty-domain-handoff.md.

Target outcome:
- Add a Runtime state-root-scoped backend lifecycle store that records environment readiness without creating environments or launching backends.

Hard constraints:
- Do not create Wine prefixes, start VMs, launch backends, expose state-root paths, mutate the host root, use privileged containers, or enable Runtime writes.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add lifecycle success, invalid transition, root-escape, missing-state, and blocked-readiness tests.
- Run the H3 verification commands.
```

## H4: Portal, Snapshot, and Rollback Safety

### Mission

Build a review-first safety plane for permissions, fake Portal records, snapshots, restore planning, and rollback evidence before any launch path becomes real.

### Start from

- `internal/runtime/portal/`
- `internal/runtime/snapshot/`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/execution_resource_grant.go`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/execution_preflight.go`

### Deliver

- Fake-mode Portal request records under explicit test roots.
- Snapshot metadata and receipt records under controlled roots.
- Restore preflight and rollback evidence that can be consumed by execution readiness.
- Safety summaries that do not expose file contents, host paths, or backend details.

### Acceptance

- Portal grants are records only; no real Portal transport call is made.
- Snapshot and rollback records cannot escape configured roots.
- Restore stays review-gated.
- Execution readiness can identify missing permission or snapshot safety evidence.

### Required verification

```text
go test ./internal/runtime/portal ./internal/runtime/snapshot ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_compatibility_repair_plan.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement H4 from docs/claude-code-mainline-empty-domain-handoff.md.

Target outcome:
- Build fake-mode Portal, snapshot, restore, and rollback safety records that execution readiness can consume while real desktop access remains disabled.

Hard constraints:
- Do not call real XDG Desktop Portal APIs, read user file contents, expose host paths, perform restore, launch backends, mutate the host root, or enable Runtime writes.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add permission-record, denied-record, snapshot-receipt, rollback-blocked, and root-escape tests.
- Run the H4 verification commands.
```

## H5: KDE Activation Materialization Consumers

### Mission

Make KDE-facing entry points consume staged activation receipts and Runtime read models instead of relying on static previews.

### Start from

- `internal/runtime/appidentity/desktop_activation_*.go`
- `internal/runtime/appidentity/kde_*`
- `internal/runtime/appidentity/identity.go`
- `kde/`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `scripts/kde_first_presence_smoke.rb`
- `scripts/kde_center_dbus_smoke.rb`

### Deliver

- Receipt-backed reads for launcher, task manager, Dolphin, KRunner, tray, notifications, settings, and Compatibility Center surfaces.
- Explicit staged-root input for test-only receipt consumption.
- KDE-safe status fields for missing, stale, mismatched, and safe receipts.
- No raw staged root path in user-facing output.

### Acceptance

- KDE entry-point previews distinguish receipt-backed state from preview-only state.
- Missing or mismatched receipts fail closed.
- No KDE surface exposes backend implementation details or host paths.
- No desktop files are written unless an explicit staging command already owns that behavior.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby scripts/container.rb kde-center-dbus-smoke
ruby scripts/container.rb kde-first-presence-smoke
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement H5 from docs/claude-code-mainline-empty-domain-handoff.md.

Target outcome:
- Make KDE entry-point read models consume staged activation receipts through explicit test roots and fail closed on missing, unsafe, or mismatched evidence.

Hard constraints:
- Do not write host desktop files, mutate the host root, refresh the host KDE cache, launch backends, enable Runtime writes, expose staged-root paths, or expose backend details.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add receipt-backed, missing-receipt, mismatched-receipt, root-hidden, and backend-detail-hidden tests.
- Run the H5 verification commands.
```

## H6: Execution Transaction Ledger

### Mission

Connect launch intent, review, preflight, permissions, resources, transactions, sessions, and session status into a blocked-by-default execution ledger.

### Start from

- `internal/runtime/appidentity/launch_intent.go`
- `internal/runtime/appidentity/execution_*.go`
- `internal/runtime/execution/`
- `cmd/xnix-runtime-go/execution_ledger_commands.go`

### Deliver

- Transaction records under explicit Runtime/test roots.
- State transitions for requested, reviewed, preflight-blocked, ready-but-gated, denied, and expired.
- Dependencies on recipe trust, lifecycle state, Portal evidence, snapshot evidence, and write gates.
- Session status records that remain synthetic until real launch is explicitly enabled by a future package.

### Acceptance

- Unsafe prerequisites keep launch blocked.
- Transaction records cannot escape configured roots.
- Write gates remain disabled.
- No backend process starts.

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
Implement H6 from docs/claude-code-mainline-empty-domain-handoff.md only after H2, H3, and H4 have enough blocking safety evidence.

Target outcome:
- Add a blocked-by-default execution transaction ledger that records reviewed launch intent and safety decisions without starting backends.

Hard constraints:
- Do not enable real launch, write methods, Wine/Proton/VM start, real Portal calls, host-root mutation, privileged containers, or backend-detail exposure.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add blocked-prerequisite, reviewed-but-gated, denied, expired, root-escape, and no-backend-launch tests.
- Run the H6 verification commands.
```

## H7: Diagnostics, Repair, and AI Boundary

### Mission

Record deterministic fixture diagnostics and repair recommendations without reading private file contents or calling live AI providers.

### Start from

- `internal/runtime/diagnostics/`
- `internal/runtime/appidentity/diagnostics.go`
- `internal/runtime/appidentity/ai_diagnostics.go`
- `internal/runtime/appidentity/repair_plan.go`
- `cmd/xnix-runtime-go/ai_diagnostics_commands.go`
- `cmd/xnix-runtime-go/diagnostic_record_commands.go`

### Deliver

- Diagnostic run records under explicit roots.
- Fixture signal collectors for recipe, trust, artifact, lifecycle, permission, and snapshot states.
- AI-safe diagnostic input redaction.
- Review-only repair recommendations and approval gates.
- No-provider mode that never calls a network AI service.

### Acceptance

- Diagnostic records contain no secrets, file contents, host paths, raw backend commands, or private keys.
- AI provider calls remain disabled.
- Repair recommendations cannot execute by themselves.
- Compatibility Center can consume the diagnostic summary safely.

### Required verification

```text
go test ./internal/runtime/diagnostics ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_ai_repair_approval_gate.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement H7 from docs/claude-code-mainline-empty-domain-handoff.md.

Target outcome:
- Add deterministic no-provider diagnostic records and review-only repair recommendations that are safe for Compatibility Center display.

Hard constraints:
- Do not call live AI providers, read user file contents, expose host paths, execute repairs, launch backends, enable Runtime writes, or mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add redaction, no-provider, fixture-signal, approval-gated, and no-repair-execution tests.
- Run the H7 verification commands.
```

## H8: Evidence and Drift Harness

### Mission

Make the repository fail when new Runtime contracts, previews, or KDE surfaces are added without implementation evidence.

### Start from

- `scripts/implementation_evidence_report.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/verify_layout.rb`
- `test/test_implementation_evidence_report.rb`
- `test/test_runtime_contract_drift_report.rb`
- `docs/claude-code-mainline-implementation-plan.md`
- `docs/claude-code-domain-dispatch.md`

### Deliver

- JSON and Markdown evidence reports with stable exit codes.
- Classification for contract-only, fixture-implemented, state-root-implemented, smoke-owned, and production-gated evidence.
- Orphan-contract detection for D-Bus methods, CLI commands, read models, and KDE entry points.
- Required-token gates for each main domain.
- Actionable next-dispatch guidance.

### Acceptance

- Adding a D-Bus method without route, smoke, client, and evidence coverage fails.
- Adding a CLI preview without tests or domain evidence fails.
- Reports remain deterministic and safe to commit.
- No Docker, QEMU, network, or backend launch is required.

### Required verification

```text
ruby -Ilib test/test_implementation_evidence_report.rb
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/implementation_evidence_report.rb --format markdown
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement H8 from docs/claude-code-mainline-empty-domain-handoff.md.

Target outcome:
- Strengthen implementation evidence and contract drift reports so new empty contracts, orphan methods, and preview-only regressions fail in developer verification.

Hard constraints:
- Do not add product behavior, enable writes, require Docker/QEMU/network, launch backends, mutate host roots, or change Runtime semantics except reporting and verification.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add report success, orphan-contract, preview-only-regression, missing-token, and stable-output tests.
- Run the H8 verification commands.
```

## H9: Atomic KDE Image and QEMU Acceptance

### Mission

Validate the KDE-first product path in a constrained image and QEMU smoke once enough Runtime evidence exists to test real integration instead of static previews.

### Start from

- `image/kinoite/`
- `scripts/build_kde_image.rb`
- `scripts/build_kde_disk.rb`
- `scripts/boot_kde_image.rb`
- `scripts/full_smoke.rb`
- `lib/xnix/image/`
- `internal/runtime/image/`
- `docs/kde-image-pipeline.md`
- `docs/kde-first-presence-smoke-spec.md`

### Deliver

- Constrained image build checks that avoid privileged host impact.
- QEMU boot smoke that records serial logs.
- KDE-first presence smoke that validates Runtime-owned application identity, activation receipts, D-Bus reads, and blocked unsafe actions.
- JSON and Markdown smoke reports.
- Clear skip/block messages when Docker, QEMU, or Colima is unavailable.

### Acceptance

- The smoke never uses privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Serial logs and smoke reports are persisted for diagnosis.
- SSH or QEMU forwarding binds to loopback only.
- Unsafe Runtime operations remain blocked in the product image.

### Required verification

```text
go test ./internal/runtime/image ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_kde_image.rb
ruby -Ilib test/test_kde_disk.rb
ruby -Ilib test/test_qemu.rb
ruby -Ilib test/test_full_smoke_script.rb
ruby scripts/container.rb kde-first-presence-smoke
ruby scripts/full_smoke.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement H9 from docs/claude-code-mainline-empty-domain-handoff.md only after the Runtime owner, trust pipeline, lifecycle, and KDE receipt consumers have enough evidence to validate.

Target outcome:
- Add or strengthen constrained atomic KDE image and QEMU acceptance so the smoke proves Runtime-owned KDE presence and blocked unsafe actions.

Hard constraints:
- Do not use privileged containers, host networking, Docker socket mounts, broad host mounts, host-root mutation, external network fetch by default, or backend launch.
- Bind guest forwarding to loopback only and persist serial logs for diagnosis.
- Keep source, comments, tests, fixtures, CLI output, and documentation in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add unavailable-dependency, loopback-forwarding, serial-log, blocked-unsafe-action, and report-generation tests.
- Run the H9 verification commands.
```

## Completion Statement Template

Claude Code must finish each branch with this statement:

```text
Implemented package: H<id> <name>
Version: v<version>
Branch: <branch>

Changed:
- <short implementation summary>

Safety preserved:
- No production D-Bus ownership.
- No Runtime write enablement.
- No backend launch.
- No host-root mutation.
- No privileged container, host networking, Docker socket mount, or broad host mount.
- No secrets, file contents, host paths, or backend details exposed in KDE-facing output.

Verification:
- <command>: PASS
- <command>: PASS

Known remaining gaps:
- <gap or "None for this package">
```
