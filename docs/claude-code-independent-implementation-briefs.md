# Claude Code Independent Implementation Briefs

> Last updated: 2026-07-15 | Companion to `docs/claude-code-open-domain-work-packages.md`

This file is a prompt-ready implementation backlog for Claude Code. It focuses on large, relatively independent domains where Xnix already has contracts, previews, or tests, but little durable implementation.

Use one brief per Claude Code branch. Do not ask one branch to implement multiple briefs unless the dependency is explicitly listed as part of the same acceptance path.

## How to Use These Briefs

1. Copy exactly one brief into Claude Code.
2. Create a branch for that brief only.
3. Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
4. Prefer Go for durable Runtime logic, C only for low-level or already-owned Runtime core surfaces, and Ruby for tests and developer tooling.
5. Keep host impact constrained: no privileged containers, host networking, Docker socket mounts, broad host-directory mounts, host-root mutation, or real backend launch unless the brief explicitly enables it.
6. Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
7. Run the listed tests plus `ruby scripts/verify_layout.rb`.
8. If the version lands on every tenth code version, run the full build and QEMU smoke gate.

## Brief Index

| Brief | Domain | Current shape | Main output |
| --- | --- | --- | --- |
| B1 | Runtime owner daemon | Mostly preview and adapter contracts | Constrained Go owner process |
| B2 | Recipe trust store | Development registry and trust previews | Production-shaped read-only trust chain |
| B3 | Artifact cache and staging | Acquisition/install planning contracts | Digest-verified local artifact pipeline |
| B4 | Environment lifecycle | Static backend/environment plans | Runtime-owned lifecycle state machine |
| B5 | Portal broker | Policy and request previews | Durable fake-first permission broker |
| B6 | Snapshot and rollback | Snapshot store plus repair contracts | Integrated restore-point workflow |
| B7 | KDE materialization writer | Desktop activation previews | Gated desktop/MIME/tray write receipts |
| B8 | KDE shell components | Data models and smoke contracts | Minimal runnable KDE adapters |
| B9 | Live session bridge | Window/task/KWin/tray previews | Safe live-session observation boundary |
| B10 | Execution transaction pipeline | Launch remains fully blocked | Gated transaction state without backend start |
| B11 | Test and repair pipeline | Plan/result/repair previews | Fixture-driven test execution records |
| B12 | AI diagnostics provider boundary | AI input/recommendation contracts | Provider interface with disabled production default |
| B13 | Atomic KDE image smoke | Product image mostly absent | Reproducible KDE image/QEMU acceptance path |

## B1: Runtime Owner Daemon

### Mission

Turn the current preview-first Runtime surface into a constrained Go process that can own a test D-Bus name and serve read-only methods.

### Start from

- `cmd/xnix-runtime-go/`
- `internal/runtime/appidentity/runtime_owner_*.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `scripts/container.rb`
- `scripts/dbus_session_smoke.rb`

### Deliver

- `cmd/xnix-runtime-owner/` with a test-mode session-bus owner.
- A small `internal/runtime/owner/` package that maps D-Bus read calls to existing Go preview handlers.
- Deterministic disabled errors for all write methods.
- Owner health/readiness output that distinguishes preview-only, smoke-owner, and production-owner states.
- Logs that include route-table version and shutdown reason without leaking backend commands, storage paths, executable names, or profile internals.

### Do not deliver

- Production bus ownership by default.
- Real backend launch.
- Host root writes.
- Privileged service setup.

### Acceptance

- A constrained container can start the owner on a session bus.
- Read methods return the same safe shapes as the corresponding CLI previews.
- Write methods fail closed with stable errors.
- Runtime owner readiness reflects the smoke-owner state.

### Tests

- `go test ./...`
- `ruby scripts/container.rb runtime-dbus-smoke`
- `ruby scripts/container.rb kde-center-dbus-smoke`
- `ruby scripts/verify_layout.rb`

## B2: Recipe Trust Store

### Mission

Replace development-only recipe trust assumptions with a production-shaped, read-only trust store and verifier boundary.

### Start from

- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `runtime/recipes/registry.json`
- `lib/xnix/compatibility/recipe_*`

### Deliver

- `internal/runtime/recipe/` with a store interface and local read-only implementation.
- Verifier boundary for registry metadata and recipe signatures.
- Trust states: `production-trusted`, `development-only`, `unsigned`, `invalid`, and `blocked`.
- Diagnostics that can feed owner readiness, install gates, and Compatibility Center pages.
- Fixture registries that remain useful but never pass production trust.

### Do not deliver

- Remote registry downloads.
- Private keys or signing secrets.
- Production write paths.

### Acceptance

- Invalid digest and invalid signature paths fail closed.
- Development fixtures are clearly non-production.
- Recipe trust state is consumed consistently by install gates and owner readiness.

### Tests

- `go test ./...`
- `ruby -Ilib test/test_recipe_registry.rb`
- `ruby -Ilib test/test_recipe_trust_policy.rb`
- `ruby -Ilib test/test_recipe_install_gate.rb`
- `ruby scripts/verify_layout.rb`

## B3: Artifact Cache and Staging

### Mission

Implement a safe artifact pipeline that can parse, verify, cache, and stage local fixture artifacts without mutating the host root.

### Start from

- `internal/runtime/appidentity/*artifact*`
- `internal/runtime/appidentity/*acquisition*`
- `internal/runtime/appidentity/*install*`
- `test/test_compatibility_artifact_manifest.rb`
- `test/test_compatibility_acquisition_preflight.rb`
- `test/test_compatibility_install_plan.rb`

### Deliver

- `internal/runtime/artifact/` with manifest parsing and digest verification.
- A Runtime state-root scoped cache namespace.
- Fixture-only acquisition and staging operations.
- Staging receipts that describe what was staged and why.
- Separate readiness for acquisition, staging, and install permission.

### Do not deliver

- Network artifact downloads by default.
- Host package-manager calls.
- Host root mutation.
- Desktop activation writes.

### Acceptance

- Digest mismatch blocks staging.
- Cache and staging paths remain under a configured test/runtime root.
- Install remains gated by recipe trust and owner readiness.
- All receipts are safe for KDE-facing summaries.

### Tests

- `go test ./...`
- `ruby -Ilib test/test_compatibility_package_source.rb`
- `ruby -Ilib test/test_compatibility_acquisition_preflight.rb`
- `ruby -Ilib test/test_compatibility_artifact_manifest.rb`
- `ruby -Ilib test/test_compatibility_install_plan.rb`
- `ruby scripts/verify_layout.rb`

## B4: Environment Lifecycle

### Mission

Convert backend environment planning into a Runtime-owned lifecycle state machine without starting real compatibility backends.

### Start from

- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/execution_readiness.go`

### Deliver

- `internal/runtime/environment/` with states such as `missing`, `planned`, `staged`, `ready`, `repair-required`, and `blocked`.
- Persistence under a controlled Runtime state root.
- Readiness transitions driven by recipe trust and artifact staging.
- Safe explanations for why an environment is ready or blocked.

### Do not deliver

- Real backend process starts.
- Shell commands exposed to KDE.
- Host storage path exposure.

### Acceptance

- State survives within a test-controlled root.
- Readiness changes when trust or staging state changes.
- User-facing output remains implementation-detail safe.

### Tests

- `go test ./...`
- `ruby -Ilib test/test_compatibility_backend_environment_plan.rb`
- `ruby -Ilib test/test_compatibility_backend_binding.rb`
- `ruby -Ilib test/test_compatibility_backend_lifecycle.rb`
- `ruby -Ilib test/test_compatibility_execution_readiness.rb`
- `ruby scripts/verify_layout.rb`

## B5: Portal Broker

### Mission

Turn Portal request previews into a durable, fake-first permission request broker that can later be wired to XDG Desktop Portal.

### Start from

- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/portal_request.go`
- `runtime/dbus/org.xnix.Compatibility1.xml`

### Deliver

- `internal/runtime/portal/` with request records and state transitions.
- Request states: `planned`, `requested`, `granted`, `denied`, `expired`, and `failed`.
- Fake transport for deterministic tests.
- Disabled real transport boundary.
- Correlation handles safe for D-Bus callers and Compatibility Center pages.

### Do not deliver

- Real host file access.
- Real Portal calls enabled by default.
- Permission grants outside the fake transport.

### Acceptance

- Fake requests can be created, completed, denied, expired, and inspected.
- Execution readiness can consume permission state.
- No direct user-document access appears in tests or logs.

### Tests

- `go test ./...`
- `ruby -Ilib test/test_portal_access_policy.rb`
- `ruby -Ilib test/test_portal_request_model.rb`
- `ruby scripts/verify_layout.rb`

## B6: Snapshot and Rollback Workflow

### Mission

Connect the constrained snapshot store to state roots, test results, repair plans, and rollback receipts.

### Start from

- `internal/runtime/snapshot/`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/desktop_activation_rollback.go`

### Deliver

- Snapshot baseline selection for application state roots.
- Restore-point receipts with content-addressed metadata.
- Rollback planning that can target staged Runtime state and desktop activation receipts.
- Repair-plan integration that requires a restore point before risky repair actions.

### Do not deliver

- User-document snapshots.
- Host-system rollback.
- Destructive rollback without an explicit test root and receipt.

### Acceptance

- Snapshot verify failures block rollback.
- Repair plans report missing restore points as blockers.
- Rollback tests stay inside a controlled root.

### Tests

- `go test ./...`
- `ruby -Ilib test/test_compatibility_snapshot_plan.rb`
- `ruby -Ilib test/test_compatibility_repair_plan.rb`
- `ruby -Ilib test/test_desktop_activation_rollback.rb`
- `ruby scripts/verify_layout.rb`

## B7: KDE Materialization Writer

### Mission

Move from KDE activation previews to a gated writer that can materialize desktop integration into a test root and produce receipts.

### Start from

- `internal/runtime/appidentity/desktop_activation_*.go`
- `internal/runtime/appidentity/file_association.go`
- `internal/runtime/appidentity/notification.go`
- `internal/runtime/appidentity/tray_status.go`

### Deliver

- Test-root-only desktop file writer.
- Test-root-only MIME association writer.
- Activation receipt records.
- Rollback receipts tied to B6 snapshot state.
- Write gates that stay disabled outside explicit test mode.

### Do not deliver

- Writes to the host desktop.
- Default MIME changes on the developer machine.
- Backend launch or execution.

### Acceptance

- Materialization works only under a configured test root.
- Receipts are enough to rollback test-root files.
- KDE-facing summaries do not expose implementation paths.

### Tests

- `go test ./...`
- `ruby -Ilib test/test_desktop_activation_installer.rb`
- `ruby -Ilib test/test_desktop_activation_rollback.rb`
- `ruby -Ilib test/test_desktop_integration_manifest.rb`
- `ruby scripts/verify_layout.rb`

## B8: KDE Shell Components

### Mission

Create minimal runnable KDE adapters that consume Runtime read models without owning Runtime policy.

### Start from

- `kde/plasmoids/org.xnix.compatibilitycenter/`
- `runtime/dbus/org.xnix.Compatibility1.xml`
- `test/test_kde_center_model.rb`
- `test/test_kde_integration_status.rb`
- `test/test_krunner_model.rb`

### Deliver

- Minimal Compatibility Center UI shell backed by fixture Runtime data.
- Minimal KRunner query adapter backed by Runtime search/query models.
- Minimal tray/status adapter backed by Runtime read models.
- Clear disabled states for actions that require write gates.

### Do not deliver

- Runtime policy in KDE code.
- Backend commands in UI.
- Production write actions.

### Acceptance

- KDE adapters can render from fixture/read-only Runtime payloads.
- Disabled actions are visually and semantically explicit.
- KDE remains replaceable presentation only.

### Tests

- `go test ./...`
- `ruby -Ilib test/test_kde_center_model.rb`
- `ruby -Ilib test/test_kde_integration_status.rb`
- `ruby -Ilib test/test_krunner_model.rb`
- `ruby scripts/container.rb kde-center-dbus-smoke`
- `ruby scripts/verify_layout.rb`

## B9: Live Session Bridge

### Mission

Define and implement the safe boundary for observing live session/window state without starting backends or controlling KWin by default.

### Start from

- `internal/runtime/appidentity/window_identity.go`
- `internal/runtime/appidentity/task_manager_identity.go`
- `internal/runtime/appidentity/kwin_window_rule.go`
- `internal/runtime/appidentity/execution_session*.go`

### Deliver

- `internal/runtime/session/` with session identity records.
- Fake window observation transport for tests.
- Disabled real KWin/task-manager transport boundary.
- Session status updates that can feed tray, task manager, and Compatibility Center views.

### Do not deliver

- Real KWin rule application by default.
- Real task-manager activation.
- Backend process starts.

### Acceptance

- Fake observations can update session status.
- Real observation and control transports remain disabled unless explicitly configured.
- No backend details leak through session records.

### Tests

- `go test ./...`
- `ruby -Ilib test/test_task_manager_identity.rb`
- `ruby -Ilib test/test_kwin_window_rule.rb`
- `ruby -Ilib test/test_tray_status_model.rb`
- `ruby scripts/verify_layout.rb`

## B10: Execution Transaction Pipeline

### Mission

Implement a gated launch transaction state machine that can reach `approved-for-dispatch` in tests without starting a real backend.

### Start from

- `internal/runtime/appidentity/launch_intent.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `internal/runtime/appidentity/runtime_write_gate.go`
- `internal/runtime/appidentity/execution_transaction*.go`

### Deliver

- `internal/runtime/execution/` with transaction records.
- Preflight collection for recipe trust, environment readiness, portal permissions, snapshot baseline, and write gates.
- State transitions: `draft`, `blocked`, `ready-for-review`, `approved-for-dispatch`, `dispatched-fake`, and `failed`.
- Fake dispatcher for tests.

### Do not deliver

- Real backend process execution.
- Runtime Launch write enablement in production.
- Host-root mutation.

### Acceptance

- A complete fake transaction can be approved and dispatched in test mode.
- Missing preflight inputs block dispatch with stable reasons.
- Production execution remains disabled.

### Tests

- `go test ./...`
- `ruby -Ilib test/test_launch_request.rb`
- `ruby -Ilib test/test_compatibility_execution_readiness.rb`
- `ruby -Ilib test/test_runtime_write_gate.rb`
- `ruby scripts/verify_layout.rb`

## B11: Compatibility Test and Repair Pipeline

### Mission

Turn compatibility test, result, and repair previews into a fixture-driven pipeline that records what would be tested and repaired.

### Start from

- `internal/runtime/appidentity/test_plan.go`
- `internal/runtime/appidentity/test_result.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/action_review_receipt.go`

### Deliver

- `internal/runtime/quality/` with test-run records and fixture runner.
- Result ingestion for pass, fail, warning, skipped, and blocked outcomes.
- Repair proposal records linked to test failures and snapshots.
- Review receipts before repair dispatch.

### Do not deliver

- Real backend smoke execution.
- Automatic repair execution.
- Host changes.

### Acceptance

- Fixture tests produce durable result records.
- Repair actions require review receipts and restore-point availability.
- AI diagnostics can consume result and repair records without reading user files.

### Tests

- `go test ./...`
- `ruby -Ilib test/test_compatibility_test_plan.rb`
- `ruby -Ilib test/test_compatibility_test_result.rb`
- `ruby -Ilib test/test_compatibility_repair_plan.rb`
- `ruby -Ilib test/test_action_review_receipt.rb`
- `ruby scripts/verify_layout.rb`

## B12: AI Diagnostics Provider Boundary

### Mission

Add a safe AI provider interface while keeping provider calls disabled by default and excluding user file contents from diagnostics.

### Start from

- `internal/runtime/appidentity/ai_diagnostic_input.go`
- `internal/runtime/appidentity/ai_diagnostic_recommendation.go`
- `internal/runtime/appidentity/ai_repair_approval_gate.go`

### Deliver

- `internal/runtime/ai/` with provider interface and disabled default provider.
- Fixture provider for deterministic tests.
- Input sanitizer that allows metadata and Runtime state but rejects file contents, secrets, host paths, and backend internals.
- Recommendation records that require repair approval gates before any action can be proposed as executable.

### Do not deliver

- Network provider calls enabled by default.
- API keys, tokens, or provider secrets.
- User document reads.
- Automatic repair execution.

### Acceptance

- Disabled provider returns a stable non-error diagnostic status.
- Fixture provider can produce deterministic recommendations.
- Sanitizer blocks forbidden input classes.

### Tests

- `go test ./...`
- `ruby -Ilib test/test_ai_diagnostic_input.rb`
- `ruby -Ilib test/test_ai_diagnostic_recommendation.rb`
- `ruby -Ilib test/test_ai_repair_approval_gate.rb`
- `ruby scripts/verify_layout.rb`

## B13: Atomic KDE Image Smoke

### Mission

Create the first reproducible product-image path for the KDE-first Xnix system, separate from the Buildroot learning baseline.

### Start from

- `boot/`
- `buildroot/`
- `scripts/container.rb`
- `scripts/full_smoke.rb`
- `docs/kde-first-compatibility-acceptance.md`
- `docs/kde-first-presence-smoke-spec.md`

### Deliver

- A documented atomic KDE image build plan.
- Containerized build/smoke commands that preserve current host-safety constraints.
- A QEMU smoke that proves the product image reaches a graphical/session readiness marker or a documented serial fallback marker.
- A clear separation between Buildroot learning smoke and KDE product smoke.

### Do not deliver

- Host package installation.
- Unbounded Docker resource usage.
- Privileged containers.
- Host network mode.
- Broad host mounts.

### Acceptance

- Clean checkout can run the documented product-image preparation path.
- QEMU smoke persists logs for diagnosis.
- Existing Buildroot smoke still passes.
- The product-image smoke can fail clearly when dependencies are missing instead of hanging.

### Tests

- `ruby scripts/verify_layout.rb`
- `ruby scripts/full_smoke.rb`
- Product-image smoke command added by this brief.

## Prompt Template

Use this when opening a Claude Code task:

```text
Implement B<N>: <brief title> from docs/claude-code-independent-implementation-briefs.md.

Scope:
- Work only on that brief and its listed dependencies.
- Keep all project-facing text in English.
- Prefer Go for Runtime implementation, Ruby for tests/tooling, and C only for low-level or already-owned Runtime core surfaces.
- Preserve host safety: no privileged containers, host networking, Docker socket mounts, broad host-directory mounts, host-root mutation, real backend launch, or production write enablement unless the brief explicitly requires it.
- Do not expose backend implementation details, raw commands, storage paths, executable paths, or profile internals in KDE/user-facing output.

Required:
- Add implementation and tests for this brief.
- Update VERSION, CHANGELOG.md, PRODUCT_OVERVIEW.md, and scripts/verify_layout.rb when new required files or invariants are introduced.
- Run the tests listed in the brief plus ruby scripts/verify_layout.rb.
- If the resulting version is a tenth code version, run the full build and QEMU smoke gate.

Do not modify docs/claude-code-implementation-packages.md unless explicitly asked.
```
