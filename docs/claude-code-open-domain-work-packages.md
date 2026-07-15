# Claude Code Open Domain Work Packages

> Last updated: 2026-07-15

This document breaks the current contract-heavy Xnix roadmap into large, independent implementation domains that Claude Code can own one branch at a time.

It intentionally does not replace `docs/claude-code-implementation-packages.md`. Treat this file as a higher-level backlog map for empty or mostly contract-only areas.

## Operating Rules

- Use one branch per package.
- Keep each package independently reviewable and shippable.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Keep all source, comments, tests, documentation, CLI output, and fixtures in English.
- Prefer Go for durable Runtime product logic, C for low-level or already-owned ABI-shaped surfaces, and Ruby for tests and developer tooling.
- Do not expose implementation details such as raw backend commands, compatibility storage paths, executable paths, or profile terminology in normal KDE/user-facing output.
- Do not use privileged containers, host networking, Docker socket mounts, broad host-directory mounts, or host-root mutation.
- Keep write methods disabled unless the package explicitly implements a gated write path.
- Add targeted tests first when the existing contract is clear enough.
- Run package-specific tests plus `ruby scripts/verify_layout.rb` before handoff.
- At every tenth code version, run the full build and QEMU smoke gate required by repository policy.

## Recommended Implementation Order

| Phase | Packages | Reason |
| --- | --- | --- |
| Foundation | D1, D2, D3 | The Runtime owner, trusted recipes, and artifacts decide what every later feature is allowed to do. |
| Controlled state | D4, D5, D6 | Environments, portals, and snapshots need safe state boundaries before execution can be real. |
| User-visible materialization | D7, D8, D9 | KDE can become concrete once Runtime-owned safety gates exist. |
| Execution and intelligence | D10, D11, D12 | Launch, tests, repair, AI, and live session UX should arrive after safety and ownership are explicit. |
| Product proof | D13 | Full image and QEMU acceptance should verify the complete path repeatedly. |

## D1: Go Runtime Owner Service

### Goal

Replace the current preview-and-adapter ownership model with a constrained Go Runtime owner process that can serve read-only Runtime methods on D-Bus.

### Current contract

- Runtime service binding, live owner gate, owner smoke plan, method parity, route manifest, recipe trust, owner process, owner readiness, and write-gate previews already describe the desired behavior.
- Production bus ownership remains gated.
- Write methods remain disabled.

### Deliverables

- A Go owner daemon entry point, for example `cmd/xnix-runtime-owner/`.
- A small internal owner package that maps D-Bus read methods to Go handlers or explicit adapter boundaries.
- A session-bus test mode that can claim the Runtime bus name inside a constrained container.
- Deterministic disabled errors for every write method.
- Owner logs that identify startup, bus claim mode, route table version, and shutdown without leaking backend details.

### Suggested files

- `cmd/xnix-runtime-owner/`
- `internal/runtime/owner/`
- `internal/runtime/appidentity/runtime_owner_*.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `runtime/dbus/xnix_compatd_smoke.c`
- `scripts/dbus_session_smoke.rb`
- `scripts/container.rb`

### Acceptance criteria

- A constrained container smoke starts the Go owner on a session bus.
- Read-only calls match the existing safe preview semantics.
- Write calls return explicit gated failures.
- No host root mutation, network dependency, privileged container, backend launch, or KDE policy ownership is introduced.
- Runtime owner readiness can distinguish preview-only, smoke-owner, and production-owner states.

### Required tests

- `go test ./...`
- `ruby scripts/verify_layout.rb`
- `ruby scripts/container.rb runtime-dbus-smoke`
- `ruby scripts/container.rb kde-center-dbus-smoke`

## D2: Production Recipe Store and Trust Chain

### Goal

Turn recipe loading and digest verification into a production-shaped trust chain while preserving development fixtures.

### Current contract

- Registry-backed recipe loading exists.
- Digest checks and trust previews exist.
- Production signature validation is still mostly modeled, not implemented.

### Deliverables

- A Go recipe store interface with local read-only roots.
- A signed registry metadata verifier boundary.
- Explicit trust states: `production-trusted`, `development-only`, `unsigned`, `invalid`, and `blocked`.
- Recipe trust diagnostics that feed Runtime owner readiness and Compatibility Center views.
- Fixture-only development registries that never pass production readiness.

### Suggested files

- `internal/runtime/recipe/`
- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `runtime/recipes/registry.json`
- `lib/xnix/compatibility/recipe_*`

### Acceptance criteria

- Invalid digests fail closed.
- Missing or invalid signatures cannot pass production owner readiness.
- No private keys, tokens, or signing secrets are committed.
- Development fixtures remain usable in tests and are visibly non-production.

### Required tests

- `go test ./...`
- `ruby -Ilib test/test_recipe_registry.rb`
- `ruby -Ilib test/test_recipe_trust_policy.rb`
- `ruby -Ilib test/test_recipe_install_gate.rb`
- `ruby scripts/verify_layout.rb`

## D3: Artifact Acquisition, Cache, and Staging

### Goal

Replace package acquisition placeholders with a safe artifact pipeline that can plan, verify, cache, and stage compatibility artifacts without mutating the host root.

### Current contract

- Package source, acquisition preflight, artifact manifest, and install plan previews exist.
- Actual fetch/cache/stage behavior is intentionally absent or disabled.

### Deliverables

- A Go artifact manifest parser.
- Digest-verified artifact references.
- A cache namespace that is local to the Runtime state root or test root.
- A dry-run acquisition path using local fixture sources first.
- A staging plan that records what would be staged and why.
- Clear separation between acquisition readiness and install permission.

### Suggested files

- `internal/runtime/artifact/`
- `internal/runtime/appidentity/*artifact*`
- `internal/runtime/appidentity/*acquisition*`
- `internal/runtime/appidentity/*install*`
- `scripts/container.rb`

### Acceptance criteria

- Digest mismatch blocks staging.
- Network fetch is disabled unless the package adds an explicit fixture-only network test mode.
- Cache and staging paths stay inside a controlled root.
- Host root mutation remains false.
- Install remains gated by recipe trust and owner readiness.

### Required tests

- `go test ./...`
- `ruby -Ilib test/test_compatibility_package_source.rb`
- `ruby -Ilib test/test_compatibility_acquisition_preflight.rb`
- `ruby -Ilib test/test_compatibility_artifact_manifest.rb`
- `ruby -Ilib test/test_compatibility_install_plan.rb`
- `ruby scripts/verify_layout.rb`

## D4: Backend Environment Lifecycle

### Goal

Implement a Runtime-owned lifecycle state machine for compatibility environments without starting real backends by default.

### Current contract

- Backend selection, environment, binding, capability, lifecycle, run plan, and execution readiness previews exist.
- Most lifecycle state is still static planning data.

### Deliverables

- A Go lifecycle package with states such as `missing`, `planned`, `staged`, `ready`, `repair-required`, and `blocked`.
- State persistence inside a controlled Runtime state root.
- A bridge from recipe trust and artifact staging into environment readiness.
- A backend binding resolver that can explain why an environment is ready or blocked.
- No raw backend commands in KDE-facing output.

### Suggested files

- `internal/runtime/environment/`
- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/execution_readiness.go`

### Acceptance criteria

- Lifecycle state survives within a test-controlled state root.
- Readiness changes when staged artifacts or recipe trust change.
- KDE/user-facing summaries remain backend-detail safe.
- Execution remains disabled until the execution package owns it.

### Required tests

- `go test ./...`
- `ruby -Ilib test/test_compatibility_backend_environment_plan.rb`
- `ruby -Ilib test/test_compatibility_backend_binding.rb`
- `ruby -Ilib test/test_compatibility_backend_lifecycle.rb`
- `ruby -Ilib test/test_compatibility_execution_readiness.rb`
- `ruby scripts/verify_layout.rb`

## D5: Portal Permission Broker

### Goal

Turn Portal access policy and fake-mode request brokering into a Runtime-owned permission request broker with durable request state.

### Current contract

- Portal access policy previews exist.
- A fake-mode Portal request broker exists.
- Real Portal transport calls are still pending.

### Deliverables

- Durable request records inside a controlled Runtime state root.
- Request states: `planned`, `requested`, `granted`, `denied`, `expired`, and `failed`.
- A transport interface with fake transport for tests and a disabled real transport boundary.
- Request correlation handles suitable for D-Bus callers.
- Permission summaries for Compatibility Center and execution readiness.

### Suggested files

- `internal/runtime/portal/`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/portal_request.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`

### Acceptance criteria

- Fake transport can create, complete, deny, and expire requests deterministically.
- Real Portal calls remain disabled unless explicitly enabled by a later package.
- No direct user-document access is introduced.
- Permission state can be consumed by execution readiness.

### Required tests

- `go test ./...`
- `ruby -Ilib test/test_portal_access_policy.rb`
- `ruby -Ilib test/test_portal_request_model.rb`
- `ruby scripts/verify_layout.rb`

## D6: Snapshot and Rollback Integration

### Goal

Connect the constrained snapshot store to Runtime state roots, repair planning, and rollback receipts.

### Current contract

- Snapshot previews exist.
- A constrained content-addressed snapshot store exists.
- Repair and rollback flows are still mostly planned.

### Deliverables

- Snapshot policies per application and state scope.
- Snapshot creation inside a controlled Runtime state root.
- Snapshot verification before restore.
- Rollback receipts that can be surfaced in Compatibility Center.
- Integration with repair approval gates and risky setting changes.

### Suggested files

- `internal/runtime/snapshot/`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/ai_repair_approval_gate.go`
- `lib/xnix/compatibility/desktop_activation_rollback.rb`

### Acceptance criteria

- Snapshots include only allowed Runtime state files.
- Restore refuses invalid snapshot identifiers and out-of-root paths.
- Repair plans can require snapshots before risky actions.
- Host root mutation remains false in previews and tests.

### Required tests

- `go test ./...`
- `ruby -Ilib test/test_compatibility_snapshot_plan.rb`
- `ruby -Ilib test/test_compatibility_repair_plan.rb`
- `ruby -Ilib test/test_desktop_activation_rollback.rb`
- `ruby scripts/verify_layout.rb`

## D7: KDE Materialization Writer

### Goal

Move from KDE previews to a gated writer that can materialize desktop entries, icons, MIME associations, activation manifests, and receipts into a staging root first.

### Current contract

- Desktop entry, desktop icon, MIME association, desktop activation bundle, preflight, staging, transaction, and status previews exist.
- Production writes remain disabled.

### Deliverables

- A Go writer that targets a staging root.
- Generated desktop entries and MIME files from the Runtime identity plan.
- Activation receipts and rollback receipts.
- Drift checks between previewed and staged files.
- A production write gate that remains closed unless explicitly enabled by owner readiness.

### Suggested files

- `internal/runtime/desktop/`
- `internal/runtime/appidentity/desktop_entry.go`
- `internal/runtime/appidentity/file_association.go`
- `internal/runtime/appidentity/desktop_activation_*.go`
- `scripts/install_runtime_activation.rb`

### Acceptance criteria

- Staging writes never escape the selected root.
- Preview and staged output match.
- Production host-root writes remain disabled by default.
- Rollback receipt describes exactly what would be removed.

### Required tests

- `go test ./...`
- `ruby -Ilib test/test_desktop_integration_manifest.rb`
- `ruby -Ilib test/test_file_association_model.rb`
- `ruby -Ilib test/test_desktop_activation_installer.rb`
- `ruby -Ilib test/test_desktop_activation_rollback.rb`
- `ruby scripts/verify_layout.rb`

## D8: KDE Shell Components

### Goal

Implement the first concrete KDE shell components behind the existing read-only Runtime contracts.

### Current contract

- KDE shell integration, application surface, Compatibility Center, KRunner, Dolphin, notification, tray, and unified settings previews exist.
- Most UI surfaces are still contracts, not real shell components.

### Deliverables

- A minimal Plasmoid or KDE-compatible surface for Compatibility Center status.
- A KRunner integration that reads Runtime application and action previews.
- A Dolphin service menu integration that routes through Runtime read models.
- Notification and tray adapters that remain preview-only until live event delivery exists.
- No backend policy embedded in KDE code.

### Suggested files

- `kde/plasmoids/`
- `kde/krunner/`
- `kde/dolphin/`
- `internal/runtime/appidentity/kde_*`
- `scripts/kde_first_presence_smoke.rb`

### Acceptance criteria

- KDE code calls Runtime read APIs or generated safe fixtures only.
- KDE code does not select backends, mutate Runtime state, or expose implementation details.
- Shell components fail closed when Runtime is unavailable.
- The seven first-release entry points remain traceable to Runtime read methods.

### Required tests

- `go test ./...`
- `ruby scripts/kde_first_presence_smoke.rb`
- `ruby scripts/verify_layout.rb`
- Relevant KDE model tests under `test/`

## D9: Window Identity and Live Session Bridge

### Goal

Make compatibility windows appear as normal KDE/Linux application windows while keeping session observation and backend execution gated.

### Current contract

- Window identity, task-manager identity, KWin rule, execution session, execution session status, tray status, and notification previews exist.
- Live window observation and live session registration remain disabled.

### Deliverables

- A Runtime-owned window identity resolver.
- A session registry inside a controlled state root.
- Task-manager grouping and restore keys that are stable across launches.
- KWin rule materialization into staging or test-only output.
- Live observation interfaces that default to disabled fake mode.

### Suggested files

- `internal/runtime/session/`
- `internal/runtime/appidentity/window_identity*.go`
- `internal/runtime/appidentity/execution_session*.go`
- `internal/runtime/appidentity/tray_status*.go`
- `runtime/core/xnix_runtime_core_*window*`

### Acceptance criteria

- A fake live session can be registered, listed, updated, and removed in tests.
- Task-manager and KWin output remains backend-detail safe.
- No real window-manager script is applied by default.
- Execution and backend process start remain disabled unless the execution package enables them.

### Required tests

- `go test ./...`
- `ruby -Ilib test/test_task_manager_identity.rb`
- `ruby -Ilib test/test_kwin_window_rule.rb`
- `ruby -Ilib test/test_tray_status_model.rb`
- `ruby scripts/verify_layout.rb`

## D10: Execution Transaction Pipeline

### Goal

Implement a gated execution transaction pipeline that turns launch intent into a reviewed, auditable Runtime action without immediately enabling broad production launch.

### Current contract

- Launch intent, execution request, execution review, execution decision, execution preflight, resource grant, execution transaction, execution session, and session status previews exist.
- Runtime write methods remain disabled.

### Deliverables

- A Go execution transaction model with request, review, decision, preflight, grant, commit, and receipt states.
- A fake backend runner for tests.
- Runtime write-gate integration for `Launch`.
- Portal permission and snapshot preconditions.
- Structured receipts suitable for Compatibility Center.

### Suggested files

- `internal/runtime/execution/`
- `internal/runtime/appidentity/execution_*.go`
- `internal/runtime/appidentity/launch_intent.go`
- `internal/runtime/appidentity/runtime_write_gate.go`

### Acceptance criteria

- Launch cannot proceed without recipe trust, environment readiness, permissions, and explicit decision state.
- Fake execution can complete without starting real compatibility backends.
- Production execution remains gated until explicitly enabled.
- Every decision produces an auditable receipt.

### Required tests

- `go test ./...`
- `ruby -Ilib test/test_launch_request.rb`
- `ruby -Ilib test/test_compatibility_execution_readiness.rb`
- `ruby -Ilib test/test_runtime_write_gate.rb`
- `ruby scripts/verify_layout.rb`

## D11: Compatibility Test and Repair Pipeline

### Goal

Turn compatibility test, result, repair, action queue, action review, and repair approval contracts into a safe automated pipeline.

### Current contract

- Test plan, test result, repair plan, action queue, action review receipt, and AI repair approval gate previews exist.
- Test execution and repair execution remain disabled.

### Deliverables

- A fake test runner for deterministic compatibility checks.
- Test result storage inside controlled Runtime state.
- Repair action planning with risk classes.
- Approval gates for repair actions.
- Queue and receipt updates for Compatibility Center.

### Suggested files

- `internal/runtime/testing/`
- `internal/runtime/repair/`
- `internal/runtime/appidentity/test_*.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/action_*.go`

### Acceptance criteria

- Tests can run in fake mode without launching real backends.
- Repair actions cannot execute without approval and snapshot preconditions.
- Compatibility Center can show queued, approved, rejected, and completed states.
- Risky actions remain blocked by default.

### Required tests

- `go test ./...`
- `ruby -Ilib test/test_compatibility_test_plan.rb`
- `ruby -Ilib test/test_compatibility_test_result.rb`
- `ruby -Ilib test/test_compatibility_repair_plan.rb`
- `ruby -Ilib test/test_action_review_receipt.rb`
- `ruby scripts/verify_layout.rb`

## D12: AI Diagnostics Provider Boundary

### Goal

Add a safe provider boundary for AI diagnostics without exposing file contents, secrets, backend details, or uncontrolled network access.

### Current contract

- AI diagnostic input, recommendation, and repair approval gate previews exist.
- Provider calls are disabled.

### Deliverables

- A provider interface with fake provider implementation.
- Strict redaction and context-size limits.
- Explicit user approval gates for any future provider call.
- Recommendation receipts that can be reviewed before action.
- Configuration through environment variables or local ignored config only.

### Suggested files

- `internal/runtime/ai/`
- `internal/runtime/appidentity/ai_diagnostics.go`
- `internal/runtime/appidentity/ai_repair_approval_gate.go`
- `lib/xnix/compatibility/ai_*`

### Acceptance criteria

- Fake provider works in tests.
- Real provider calls remain disabled by default.
- No secrets are committed or logged.
- File contents and host paths are excluded unless a later package explicitly adds approved file access.

### Required tests

- `go test ./...`
- `ruby -Ilib test/test_ai_diagnostic_input.rb`
- `ruby -Ilib test/test_ai_diagnostic_recommendation.rb`
- `ruby -Ilib test/test_ai_repair_approval_gate.rb`
- `ruby scripts/verify_layout.rb`

## D13: Atomic KDE Image and QEMU Product Smoke

### Goal

Make the flagship KDE image pipeline repeatedly buildable and smoke-testable without depending on broad host access.

### Current contract

- Buildroot/QEMU learning baseline exists.
- KDE image manifests and boot smoke contracts exist.
- Product-grade atomic image and full graphical smoke still need sustained implementation.

### Deliverables

- A pinned KDE image manifest and generated build context.
- A constrained image build script that refuses unsafe host access.
- A QEMU boot smoke that checks serial logs and loopback-only service exposure.
- Runtime owner and KDE entrypoint smoke checks inside the image when feasible.
- JSON and Markdown reports from smoke runs.

### Suggested files

- `image/kinoite/`
- `lib/xnix/image/`
- `scripts/build_kde_image.rb`
- `scripts/build_kde_disk.rb`
- `scripts/boot_kde_image.rb`
- `scripts/full_smoke.rb`

### Acceptance criteria

- A clean checkout can reproduce the build inputs.
- QEMU smoke persists serial logs.
- SSH or Runtime smoke is reachable only through loopback-bound forwarding.
- The smoke does not require privileged containers unless the command explicitly declares and gates that requirement.
- Reports are emitted as JSON and Markdown where practical.

### Required tests

- `ruby -Ilib test/test_kde_image.rb`
- `ruby -Ilib test/test_kde_disk.rb`
- `ruby scripts/verify_layout.rb`
- `ruby scripts/full_smoke.rb` when the version milestone requires the full gate.

## Package Handoff Template

Use this template when opening a Claude Code task:

```text
Implement package D<N>: <package name>.

Scope:
- Follow docs/claude-code-open-domain-work-packages.md package D<N>.
- Keep source/docs/comments/CLI output in English.
- Prefer Go for durable Runtime logic; Ruby only for tests/tooling.
- Do not expose backend implementation details in user-facing output.
- Do not enable host-root mutation, privileged containers, host networking, Docker socket mounts, broad host mounts, or real backend launch unless explicitly listed in the package acceptance criteria.

Required changes:
- Add or update the package implementation.
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Update scripts/verify_layout.rb if new required files or source invariants are added.

Required verification:
- Run package-specific tests listed in the package.
- Run ruby scripts/verify_layout.rb.
- If this lands on a tenth code version, run the full build and QEMU smoke gate.

Do not modify unrelated packages in the same branch.
```

## Completion Checklist

- [ ] Package has a single clear owner branch.
- [ ] Scope stays inside one domain.
- [ ] Runtime read/write boundaries remain explicit.
- [ ] User-facing output hides implementation details.
- [ ] Host safety constraints are preserved.
- [ ] Tests cover success, failure, and blocked/gated states.
- [ ] Version, changelog, and product overview are updated for code changes.
- [ ] Full smoke is run when repository version policy requires it.
