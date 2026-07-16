# Claude Code Priority Implementation Packages

> Last updated: 2026-07-16 | Baseline: v0.2.252

This document is the short, practical handoff board for assigning large Xnix implementation packages to Claude Code.

Use this file when the repository already has contracts, previews, CLI commands, or tests, but the domain still lacks durable implementation evidence. The goal is to make Claude Code implement product behavior, not add more empty contracts.

This file does not replace:

- `docs/claude-code-domain-dispatch.md`
- `docs/claude-code-mainline-implementation-plan.md`
- `docs/claude-code-large-empty-domain-assignments.md`
- `docs/claude-code-empty-domain-implementation-packages.md`
- `docs/claude-code-contract-gap-work-packages.md`
- `docs/claude-code-independent-implementation-briefs.md`

Treat this file as the first page to copy from when giving Claude Code a focused branch.

## Handoff Rules

Give Claude Code exactly one package from this document at a time.

Every package branch must:

- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Compatibility Runtime product logic.
- Use C only for low-level Runtime transport, ABI-shaped records, or already-owned C policy surfaces.
- Use Ruby for tests, smoke scripts, reports, and developer tooling.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Add targeted success, failure, and blocked-state tests.
- Run the package-specific verification commands and `ruby scripts/verify_layout.rb`.
- Stop and report if implementation requires behavior outside the selected package.

Claude Code must not:

- Add new preview contracts as a substitute for implementation evidence.
- Mutate the host root.
- Use privileged containers.
- Use host networking.
- Mount the Docker socket.
- Add broad host-directory mounts.
- Start real Wine, Proton, VM, or compatibility backends unless a later package explicitly defines a gated test-only path.
- Enable production D-Bus ownership.
- Enable Runtime write methods.
- Perform default network artifact fetches.
- Call host package managers.
- Expose raw backend commands, executable paths, compatibility storage paths, profile terminology, host paths, file contents, secrets, tokens, or private keys in KDE-facing output.

## Recommended Assignment Order

Start with `P1`, `P2`, `P3`, and `P8`.

| Order | Package | Why first |
| --- | --- | --- |
| 1 | P1 Runtime owner read service | Other domains need a Runtime-owned read boundary instead of preview-only commands. |
| 2 | P2 Recipe, artifact, cache, and install trust | Activation, environments, and execution must start from trusted local inputs. |
| 3 | P3 Runtime state root and environment lifecycle | Execution readiness needs durable state before launch is meaningful. |
| 4 | P8 Evidence and drift harness | Prevents contract-only surfaces from multiplying silently. |
| 5 | P5 KDE activation materialization | User-visible desktop integration should consume receipts from trusted inputs. |
| 6 | P4 Portal, snapshot, and rollback safety | Required before launch approvals or restore flows can be credible. |
| 7 | P6 Execution transaction ledger | Should wait until trust, state, and safety gates can block unsafe launch. |
| 8 | P7 Diagnostics, repair, and AI boundary | Can run fixture-first, then feed execution readiness and repair review. |
| 9 | P9 Atomic KDE image and QEMU acceptance | Should validate real product evidence, not only preview contracts. |
| 10 | P10 Runtime packaging and service binding | Should follow enough owner and image evidence to avoid shipping a hollow service. |

Safe parallel work:

- `P1` and `P8` can usually run in parallel.
- `P2` and `P5` can run in parallel only if `P5` remains receipt-consuming and does not trust production inputs by itself.
- `P7` can run with `P3` when it uses fixture diagnostics and state-root-safe records.

Avoid parallel work:

- Do not run two branches that both change D-Bus XML, owner dispatch, or Runtime route parity.
- Do not run `P2` and `P3` together if both redefine install readiness semantics.
- Do not run `P6` until `P2`, `P3`, and `P4` have enough blocked-state evidence.

## P1: Runtime Owner Read Service

### Mission

Implement a constrained Go Runtime owner process that can own a private test session-bus name, serve read-only Runtime methods, and fail all write methods closed.

### Start from

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `internal/runtime/appidentity/runtime_owner_*.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `runtime/dbus/xnix_compatd_smoke.c`
- `scripts/runtime_owner_candidate_smoke.rb`
- `scripts/dbus_session_smoke.rb`
- `scripts/runtime_contract_drift_report.rb`

### Deliver

- Long-running smoke-owner mode.
- Private session-bus ownership only inside constrained tests.
- Read-only D-Bus method routing through the Go owner table.
- Stable disabled responses for every write method.
- Explicit unsupported-read records for routes that are intentionally not implemented yet.
- Owner lifecycle logs for startup, route table version, bus mode, readiness, dispatch result, and shutdown.

### Acceptance

- A constrained container smoke starts the Go owner on a private session bus.
- Representative D-Bus reads return safe Go Runtime payloads.
- Every write method returns a deterministic disabled error.
- Contract drift reporting fails when D-Bus XML, Go routes, C smoke bridge, Ruby D-Bus client, or CLI route commands diverge.

### Required verification

```text
go test ./cmd/xnix-runtime-owner ./internal/runtime/owner ./cmd/xnix-runtime-go ./internal/runtime/appidentity
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/container.rb runtime-owner-candidate-smoke
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement P1 from docs/claude-code-priority-implementation-packages.md.

Target outcome:
- Add a constrained Go Runtime owner smoke service that can claim a private test session-bus name, serve existing read-only Runtime methods, and fail every write method closed.

Hard constraints:
- Do not enable production D-Bus ownership, write methods, system service installation, backend launch, network fetch, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, disabled-write, unsupported-read, lifecycle, and blocked-production-owner tests.
- Run the P1 required verification commands.
```

## P2: Recipe, Artifact, Cache, and Install Trust

### Mission

Turn recipe trust, package source, artifact manifest, cache planning, staging receipts, and install readiness into a local-first trust pipeline.

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

- Read-only recipe store with explicit local roots.
- Registry validation for schema, IDs, relative paths, digests, and declared signing state.
- Replaceable signature verifier boundary without committed production keys.
- Fixture artifact manifest parser and SHA-256 verifier.
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

### Copyable Claude Code prompt

```text
Implement P2 from docs/claude-code-priority-implementation-packages.md.

Target outcome:
- Build a local-first recipe, artifact, cache, staging, and install-readiness trust pipeline that verifies digests, stages only under explicit roots, and fails closed for invalid inputs.

Hard constraints:
- No default network fetch, host package-manager call, private key, production trust bypass, desktop write, backend launch, privileged container, host networking, Docker socket mount, broad host mount, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add valid fixture, invalid digest, invalid path, production-signature-blocked, root-escape, and blocked-install-readiness tests.
- Run the P2 required verification commands.
```

## P3: Runtime State Root and Environment Lifecycle

### Mission

Implement durable state-root records for application environment lifecycle without creating real environments or starting backends.

### Start from

- `internal/runtime/environment/`
- `internal/runtime/appidentity/state_root.go`
- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/backend_capability_matrix.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `cmd/xnix-runtime-go/`

### Deliver

- State-root-safe lifecycle store.
- Lifecycle states such as `missing`, `planned`, `staged`, `ready`, `repair-required`, and `blocked`.
- State transitions driven by trusted recipes, verified artifacts, and install readiness.
- Backend binding resolver that explains readiness without exposing backend details.
- CLI commands for reading lifecycle state from an explicit test root.

### Acceptance

- Lifecycle state survives inside an explicit test root.
- Root escape attempts fail closed.
- Readiness changes when recipe trust, artifact staging, or install readiness changes.
- KDE-facing summaries remain free of backend commands, storage paths, and profile terminology.
- No backend process starts.

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

### Copyable Claude Code prompt

```text
Implement P3 from docs/claude-code-priority-implementation-packages.md.

Target outcome:
- Add a Runtime state-root lifecycle store for compatibility environments, with durable blocked and ready states under explicit roots.

Hard constraints:
- Do not start Wine, Proton, VM, or any real compatibility backend.
- Do not expose raw backend commands, executable paths, compatibility storage paths, profile terminology, host paths, or file contents in KDE-facing output.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add persistence, root-escape, trust-change, artifact-change, and blocked-backend tests.
- Run the P3 required verification commands.
```

## P4: Portal, Snapshot, and Rollback Safety Plane

### Mission

Create safe state and review models for Portal permissions, snapshots, restore planning, rollback records, and recovery readiness without making real Portal calls or mutating user files.

### Start from

- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/execution_resource_grant.go`
- `internal/runtime/appidentity/execution_preflight.go`
- `internal/runtime/appidentity/execution_review.go`
- `internal/runtime/appidentity/execution_decision.go`
- `internal/runtime/appidentity/state_root.go`
- `internal/runtime/rootfs/`
- `cmd/xnix-runtime-go/`

### Deliver

- State-root-safe permission review records.
- Snapshot plan records with relative receipt paths.
- Restore and rollback readiness records that remain disabled.
- Explicit user-document exclusion unless a fixture test grants a safe fake Portal scope.
- Clear blocked reasons for missing permission, missing snapshot, or unsafe target.

### Acceptance

- Real Portal transport calls remain disabled.
- Snapshot creation and restore execution remain disabled.
- Permission and snapshot records stay inside explicit roots.
- Root escape, unsafe path, and missing-review cases fail closed.
- Execution preflight can consume the safety plane without enabling launch.

### Required verification

```text
go test ./internal/runtime/rootfs ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_snapshot_plan.rb
ruby -Ilib test/test_execution_resource_grant.rb
ruby -Ilib test/test_execution_preflight.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement P4 from docs/claude-code-priority-implementation-packages.md.

Target outcome:
- Build a state-root-safe Portal, snapshot, restore, and rollback safety plane that execution preflight can consume while all real writes remain disabled.

Hard constraints:
- No real Portal calls, snapshot creation, restore execution, host-root mutation, backend launch, broad host mounts, or file-content exposure.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add permission-granted fixture, permission-missing, snapshot-missing, rollback-blocked, root-escape, and execution-preflight integration tests.
- Run the P4 required verification commands.
```

## P5: KDE Activation Materialization

### Mission

Move KDE activation from static previews toward receipt-backed materialization under explicit roots, while keeping host-root writes and launch disabled.

### Start from

- `internal/runtime/appidentity/desktop_activation_*.go`
- `internal/runtime/appidentity/kde_entrypoints.go`
- `internal/runtime/appidentity/kde_shell_surface.go`
- `internal/runtime/appidentity/kde_action_*.go`
- `internal/runtime/appidentity/identity.go`
- `internal/runtime/appidentity/window_identity_routes.go`
- `cmd/xnix-runtime-go/desktop_activation_*`
- `cmd/xnix-runtime-go/kde_shell_commands.go`
- `cmd/xnix-runtime-go/window_identity_commands.go`
- `scripts/kde_first_presence_smoke.rb`
- `scripts/kde_center_dbus_smoke.rb`

### Deliver

- Receipt-backed reads for launcher, MIME, desktop icon, task manager, KWin, tray, notification, settings, and Compatibility Center surfaces.
- Explicit activation roots for staging and reading evidence.
- Materialization receipts with relative paths and digests.
- Drift checks proving KDE surfaces consume Runtime activation receipts rather than static fixture-only previews.
- Compatibility Center status that distinguishes preview-only, staged, receipt-backed, and host-installed states.

### Acceptance

- No host root writes occur.
- No KDE cache refresh is executed on the host.
- No launch path is enabled.
- Activation-root paths are never leaked in JSON or KDE-facing summaries.
- Missing or mismatched receipts fail closed.
- All seven KDE entry points can show receipt-backed readiness where applicable.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_kde_entrypoints.rb
ruby -Ilib test/test_kde_first_presence_smoke.rb
ruby scripts/container.rb kde-center-dbus-smoke
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement P5 from docs/claude-code-priority-implementation-packages.md.

Target outcome:
- Make KDE activation and all normal desktop entry points consume staged Runtime activation receipts under explicit roots, without writing to the host root or enabling launch.

Hard constraints:
- No host-root mutation, host KDE cache refresh, backend launch, production write methods, privileged container, host networking, Docker socket mount, broad host mount, or activation-root path exposure.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add receipt-backed, missing-receipt, mismatched-receipt, root-escape, and no-host-mutation tests for every touched KDE surface.
- Run the P5 required verification commands.
```

## P6: Execution Transaction Ledger

### Mission

Implement a reviewed execution transaction ledger that records launch intent, review, decision, resource grants, and session status while launch remains disabled.

### Start from

- `internal/runtime/appidentity/execution_request.go`
- `internal/runtime/appidentity/execution_review.go`
- `internal/runtime/appidentity/execution_decision.go`
- `internal/runtime/appidentity/execution_preflight.go`
- `internal/runtime/appidentity/execution_transaction.go`
- `internal/runtime/appidentity/execution_session.go`
- `internal/runtime/record/`
- `cmd/xnix-runtime-go/execution_ledger_commands.go`

### Deliver

- State-root-safe execution ledger records.
- Deterministic transaction IDs from safe inputs.
- Launch-intent, review, decision, preflight, and disabled-session records.
- Append-only audit receipts with relative paths.
- Compatibility Center read models for pending, blocked, approved-but-disabled, and failed-preflight states.

### Acceptance

- No backend launches.
- No real process execution.
- No write method enablement.
- Unsafe or missing prerequisites block transaction readiness.
- Ledger records do not expose raw commands, executable paths, storage paths, host paths, or file contents.

### Required verification

```text
go test ./internal/runtime/record ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_execution_ledger.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby -Ilib test/test_launch_intent.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement P6 from docs/claude-code-priority-implementation-packages.md.

Target outcome:
- Add a state-root-safe execution transaction ledger that records reviewed launch intent and disabled session status without starting any backend.

Hard constraints:
- No real launch, backend process start, write-method enablement, raw command exposure, executable path exposure, compatibility storage path exposure, host path exposure, or file-content exposure.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add approved-but-disabled, missing-trust, missing-state, missing-permission, root-escape, and append-only ledger tests.
- Run the P6 required verification commands.
```

## P7: Diagnostics, Repair, and AI Boundary

### Mission

Implement fixture-first diagnostics, repair recommendations, and AI-safe context records without calling live AI providers or applying repairs automatically.

### Start from

- `internal/runtime/appidentity/ai_diagnostics.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/test_plan.go`
- `internal/runtime/appidentity/test_result.go`
- `internal/runtime/appidentity/diagnostic_history_test.go`
- `cmd/xnix-runtime-go/ai_diagnostics_commands.go`
- `cmd/xnix-runtime-go/diagnostic_record_commands.go`
- `scripts/runtime_activation_smoke.rb`

### Deliver

- Fixture diagnostic signal runner.
- State-root-safe diagnostic history records.
- AI-safe input projection that excludes file contents, secrets, host paths, and backend details.
- Review-only repair recommendations.
- Repair approval gates that remain disabled until execution and write gates are ready.
- Compatibility Center cards for failed checks and repair suggestions.

### Acceptance

- No live AI provider calls.
- No auto-repair.
- No backend launch.
- Diagnostic records stay under explicit roots.
- Sensitive terms and host paths are filtered from AI-safe context and KDE-facing summaries.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_ai_diagnostics.rb
ruby -Ilib test/test_runtime_diagnostic_run_record.rb
ruby -Ilib test/test_runtime_diagnostic_history.rb
ruby -Ilib test/test_repair_plan.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement P7 from docs/claude-code-priority-implementation-packages.md.

Target outcome:
- Add fixture-first diagnostic history, AI-safe input projection, and review-only repair recommendations under explicit Runtime roots.

Hard constraints:
- No live AI provider calls, auto-repair, backend launch, file-content reads, secret exposure, host path exposure, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add passing fixture, failing fixture, privacy-filtering, repair-review-only, approval-blocked, and root-escape tests.
- Run the P7 required verification commands.
```

## P8: Evidence and Drift Harness

### Mission

Strengthen developer reports so the repository fails fast when new contracts, D-Bus methods, CLI commands, or read models are added without implementation evidence.

### Start from

- `scripts/implementation_evidence_report.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/full_smoke.rb`
- `scripts/container.rb`
- `test/test_implementation_evidence_report.rb`
- `test/test_runtime_contract_drift_report.rb`
- `test/test_full_smoke_script.rb`
- `docs/claude-code-mainline-implementation-plan.md`

### Deliver

- Evidence classification for contract-only, fixture-implemented, state-root-implemented, smoke-owned, and production-gated domains.
- Orphan contract detection for D-Bus XML, Go owner routes, Runtime CLI commands, C smoke bridge methods, Ruby smoke clients, and docs.
- Markdown and JSON reports with next actionable packages.
- CI-friendly failure modes for preview-only regressions.
- Report tests that do not require Docker, QEMU, network, or privileged containers.

### Acceptance

- Adding a new read method in only one layer fails drift reporting.
- Adding a new package without evidence classification fails evidence reporting.
- Reports identify the owning package from this document when possible.
- The harness itself does not require Docker, QEMU, network, backend launch, or host-root mutation.

### Required verification

```text
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/implementation_evidence_report.rb --format markdown
ruby scripts/runtime_contract_drift_report.rb --format json
ruby -Ilib test/test_implementation_evidence_report.rb
ruby -Ilib test/test_runtime_contract_drift_report.rb
ruby -Ilib test/test_full_smoke_script.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement P8 from docs/claude-code-priority-implementation-packages.md.

Target outcome:
- Strengthen implementation evidence and contract drift reports so new contracts cannot land without classified implementation evidence and route parity.

Hard constraints:
- Do not require Docker, QEMU, network, backend launch, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation for the report tests.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add tests for orphan D-Bus methods, orphan CLI routes, contract-only domains, missing package ownership, and report format stability.
- Run the P8 required verification commands.
```

## P9: Atomic KDE Image and QEMU Acceptance

### Mission

Make the KDE-first image pipeline validate real Runtime evidence in QEMU while preserving host safety.

### Start from

- `scripts/build_kde_image.rb`
- `scripts/build_kde_disk.rb`
- `scripts/boot_kde_image.rb`
- `scripts/full_smoke.rb`
- `scripts/container.rb`
- `docs/kde-image-pipeline.md`
- `docs/kde-first-presence-smoke-spec.md`
- `docs/kde-first-compatibility-acceptance.md`
- `buildroot/`

### Deliver

- Deterministic build and boot records.
- QEMU serial evidence with persisted logs.
- KDE-first Runtime presence checks.
- Restricted loopback networking only.
- Full-smoke Markdown and JSON reports that include Runtime evidence, activation readiness, and host-safety flags.
- Timeout and cleanup behavior that cannot leave broad host mounts or privileged containers behind.

### Acceptance

- QEMU boot evidence is reproducible from a clean checkout when dependencies are available.
- Host root remains unmodified.
- No privileged container, host networking, Docker socket mount, or broad host mount is required.
- Failure reports keep serial logs and explain the blocked stage.
- Full smoke still fails closed when core Runtime evidence is missing.

### Required verification

```text
ruby scripts/full_smoke.rb
ruby scripts/container.rb build
ruby scripts/container.rb boot-system
ruby -Ilib test/test_full_smoke_script.rb
ruby -Ilib test/test_full_smoke_report.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement P9 from docs/claude-code-priority-implementation-packages.md.

Target outcome:
- Make the KDE image and QEMU smoke pipeline validate real Runtime evidence with persisted logs and explicit host-safety reporting.

Hard constraints:
- No privileged container, host networking, Docker socket mount, broad host mount, host-root mutation, or unsafe QEMU network exposure.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, boot-timeout, missing-runtime-evidence, cleanup, and host-safety report tests.
- Run the P9 required verification commands.
```

## P10: Runtime Packaging and Service Binding

### Mission

Prepare installable Runtime packaging and service binding records without enabling production ownership or host installation by default.

### Start from

- `internal/runtime/appidentity/runtime_service_binding.go`
- `internal/runtime/appidentity/runtime_owner_process.go`
- `internal/runtime/appidentity/runtime_owner_readiness.go`
- `cmd/xnix-runtime-go/runtime_service_binding_cli_test.go`
- `cmd/xnix-runtime-go/runtime_owner_process_cli_test.go`
- `buildroot/board/xnix/rootfs-overlay/`
- `scripts/install_runtime_activation.rb`
- `scripts/container.rb`

### Deliver

- Service binding records for system, session, and smoke modes.
- Rootfs overlay packaging for the constrained image.
- Installation plan records with disabled production ownership.
- Smoke-only service activation inside explicit image or test roots.
- Readiness that distinguishes packaged, smoke-installed, and production-enabled states.

### Acceptance

- No service is installed on the host.
- Production D-Bus ownership remains disabled.
- Packaging records use explicit image/test roots.
- Service files and activation records are reproducible in the image root.
- Runtime owner readiness consumes packaging state without enabling writes.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby scripts/container.rb build
ruby scripts/container.rb runtime-dbus-smoke
ruby -Ilib test/test_runtime_service_binding.rb
ruby -Ilib test/test_runtime_owner_process.rb
ruby scripts/verify_layout.rb
```

### Copyable Claude Code prompt

```text
Implement P10 from docs/claude-code-priority-implementation-packages.md.

Target outcome:
- Add Runtime packaging and service binding records for image/test roots while production ownership and host installation remain disabled.

Hard constraints:
- Do not install services on the host, enable production D-Bus ownership, enable write methods, launch backends, use privileged containers, use host networking, mount Docker socket, broad-mount the host, or mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add packaged, smoke-installed, production-disabled, root-escape, and readiness-integration tests.
- Run the P10 required verification commands.
```

## Completion Statement Template

Require Claude Code to finish with this exact structure:

```text
Implemented package: P<N> <name>
Branch:
Version:
Commit:
Tag:

Changed files:
- ...

Implementation evidence:
- ...

Verification run:
- <command>: PASS
- <command>: PASS

Safety confirmation:
- Host root modified: false
- Privileged container required: false
- Host networking used: false
- Docker socket mounted: false
- Broad host mount used: false
- Backend launch enabled: false
- Production D-Bus ownership enabled: false
- Runtime write methods enabled: false

Known follow-up:
- ...
```
