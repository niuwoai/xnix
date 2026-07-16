# Claude Code Fourth-Wave Task Batch

> Last updated: 2026-07-16 | Baseline under review: v0.2.293

This document contains a fourth batch of bounded Claude Code tasks for the KDE-first Windows application compatibility mainline.

Use this batch after the current mainline, second-wave, and third-wave work has been merged or deliberately skipped. The goal is to turn Runtime evidence into reviewable product surfaces: audit timelines, dry-run recovery plans, permission lifecycle summaries, diagnostics playbooks, and release-readiness reports. These tasks must not enable real execution, real Portal transport, host mutation, Docker, or QEMU.

Do not modify `docs/claude-code-implementation-packages.md` from any task in this batch. That file is Claude-owned.

## Global Rules for Every Fourth-Wave Task

Every Claude Code branch must:

- Pick exactly one task from `FW1` through `FW8`.
- Keep source, comments, fixtures, tests, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior.
- Use C only for low-level D-Bus transport, ABI-shaped records, or already-owned Runtime policy surfaces.
- Use Ruby only for tests, smoke scripts, reports, and lightweight developer tooling.
- Keep KDE as shell and presentation only.
- Keep the Xnix Compatibility Runtime as the policy and product-logic owner.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for code changes.
- Add success, failure, malformed-input, and blocked-state tests.
- Run the task-specific verification commands plus `ruby scripts/verify_layout.rb`.
- Stop and report if the task needs cross-lane ownership or unsafe host behavior.

Unsafe behavior remains disabled:

- Production D-Bus ownership.
- Runtime write methods.
- Real Wine, Proton, VM, or compatibility backend launch.
- Real XDG Portal transport calls.
- Default network artifact fetch.
- Host package-manager calls.
- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Docker or QEMU execution unless a human explicitly authorizes a restricted smoke.
- Host-root mutation.
- Raw backend commands in KDE-facing output.
- Raw executable paths, Runtime state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials in user-facing output.

## Dispatch Order

| Order | Task | Primary lane | Why now |
| --- | --- | --- | --- |
| 1 | `FW1` Runtime evidence audit timeline | `CW8 / CW10` | Reviewers need one chronological read model that explains which receipts exist, which are missing, and why launch is still blocked. |
| 2 | `FW2` Portal permission lifecycle dashboard | `CW5 / CW8` | Portal records should become a user-safe lifecycle summary before real Portal transport exists. |
| 3 | `FW3` Snapshot restore rehearsal plan | `CW6 / CW8` | Restore and rollback should have a dry-run plan before any restore operation can be enabled. |
| 4 | `FW4` Backend capability fixture probe harness | `CW3 / CW10` | Backend profiles need deterministic fixture probes without starting any backend. |
| 5 | `FW5` Settings change dependency review | `CW4 / CW5 / CW8` | Unified settings should explain exactly which Runtime evidence and receipts are required before a setting can change. |
| 6 | `FW6` Diagnostics repair playbook review queue | `CW7 / CW8` | AI diagnostics should feed human-reviewable repair playbooks without auto-repair. |
| 7 | `FW7` Desktop activation materialization audit | `CW4 / CW6` | Staged activation files need a reviewer-facing diff and rollback readiness report before host installation. |
| 8 | `FW8` Restricted release readiness packet | `CW10 / CW11` | The project needs a single safe release-readiness packet that summarizes local checks and explicitly names blocked heavy tests. |

Do not dispatch `FW8` as a real Docker or QEMU run. `FW8` is a report task only unless a human explicitly authorizes restricted execution.

## FW1: Runtime Evidence Audit Timeline

### Mission

Create a Runtime-owned audit timeline that joins action receipts, Portal receipts, execution ledger records, session records, diagnostics records, snapshot evidence, and write-gate state into one read-only chronological report.

### Start from

- `internal/runtime/appidentity/kde_action_receipt.go`
- `internal/runtime/appidentity/execution_session_record_evidence.go`
- `internal/runtime/execution/ledger.go`
- `internal/runtime/portal/ledger.go`
- `internal/runtime/diagnostics/history.go`
- `internal/runtime/snapshot/store.go`
- `cmd/xnix-runtime-go/execution_ledger_commands.go`
- `test/test_compatibility_action_queue.rb`
- `test/test_runtime_write_gate.rb`

### Deliver

- A Go read model and CLI command such as `runtime-evidence-audit-preview`.
- A chronological `timeline` with stable event ids, event types, sources, safe summaries, evidence ids, and blocked reasons.
- Missing-evidence entries for absent Portal receipt, absent snapshot baseline, absent review receipt, absent execution ledger, absent session record, and blocked write gate.
- Tests proving malformed state-root records fail closed.
- Tests proving the report never exposes state-root paths, file contents, raw executable paths, backend commands, or secrets.

### Keep disabled

- Timeline persistence.
- Execution approval.
- Backend launch.
- Real Portal transport.
- Snapshot creation or restore.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/execution ./internal/runtime/portal ./internal/runtime/diagnostics ./internal/runtime/snapshot ./cmd/xnix-runtime-go
ruby -Ilib test/test_runtime_write_gate.rb
ruby -Ilib test/test_compatibility_action_queue.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement FW1 from docs/claude-code-fourth-wave-task-batch.md.

Target outcome:
- Add a Runtime evidence audit preview that joins review receipts, Portal receipts, execution ledger records, session records, diagnostics records, snapshot evidence, and Runtime write-gate state into a chronological read-only timeline.

Hard constraints:
- Do not persist the timeline, approve execution, start backends, call real Portal transport, create or restore snapshots, run Docker/QEMU, or mutate the host root.
- Do not expose raw backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add present-evidence, missing-evidence, malformed-record, blocked-write-gate, and no-sensitive-output tests.
- Run the FW1 verification commands and report exact commands run.
```

## FW2: Portal Permission Lifecycle Dashboard

### Mission

Turn fake-mode Portal permission records into a user-safe lifecycle dashboard that KDE can render without making real Portal calls.

### Start from

- `internal/runtime/portal/ledger.go`
- `internal/runtime/portal/request.go`
- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/execution_resource_grant.go`
- `lib/xnix/compatibility/compatibility_permission_review_plan.rb`
- `cmd/xnix-runtime-go/runtime_safety_commands.go`
- `test/test_portal_request_model.rb`
- `test/test_portal_access_policy.rb`
- `test/test_compatibility_permission_review_plan.rb`

### Deliver

- A Go read model and CLI command such as `portal-permission-lifecycle-preview`.
- Lifecycle states for missing, requested, granted, denied, expired, cancelled, malformed, and superseded fake-mode receipts.
- Per-resource summaries for Documents, Downloads, Clipboard, Camera, Screen, Remote Desktop, URI, Print, and Network.
- A KDE-safe summary that says what the user can review next without exposing host paths.
- Tests proving granted fake-mode receipts do not create real host permissions.

### Keep disabled

- Real XDG Portal transport.
- Host permission changes.
- Request-object creation.
- Permission grants.
- Execution approval.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/portal ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_compatibility_permission_review_plan.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement FW2 from docs/claude-code-fourth-wave-task-batch.md.

Target outcome:
- Add a Portal permission lifecycle preview that summarizes fake-mode Portal permission records for KDE without calling real Portal transport.

Hard constraints:
- Do not call real XDG Portal transport, change host permissions, create request objects, grant permissions, approve execution, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add missing, granted, denied, expired, cancelled, malformed, superseded, and no-host-permission tests.
- Run the FW2 verification commands and report exact commands run.
```

## FW3: Snapshot Restore Rehearsal Plan

### Mission

Create a dry-run restore rehearsal plan that explains whether a rollback could be prepared for an application without restoring anything.

### Start from

- `internal/runtime/snapshot/store.go`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/execution_transaction.go`
- `internal/runtime/appidentity/desktop_activation_transaction.go`
- `test/test_compatibility_snapshot_plan.rb`
- `test/test_desktop_activation_rollback.rb`
- `test/test_runtime_core.rb`

### Deliver

- A Go read model and CLI command such as `snapshot-restore-rehearsal-preview`.
- Restore readiness states for absent baseline, corrupt baseline, stale baseline, compatible baseline, and blocked restore.
- Safe rollback scope descriptions for Runtime metadata, application state, activation artifacts, and user-visible desktop entries.
- Digest and schema checks for fixture restore points.
- Tests proving no restore operation, file overwrite, host-root mutation, or desktop cache refresh occurs.

### Keep disabled

- Snapshot creation.
- Restore execution.
- File overwrite.
- Desktop cache refresh.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/snapshot ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement FW3 from docs/claude-code-fourth-wave-task-batch.md.

Target outcome:
- Add a dry-run snapshot restore rehearsal preview that explains rollback readiness without restoring anything.

Hard constraints:
- Do not create snapshots, restore files, overwrite files, refresh desktop caches, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add absent-baseline, corrupt-baseline, stale-baseline, compatible-baseline, blocked-restore, and no-side-effect tests.
- Run the FW3 verification commands and report exact commands run.
```

## FW4: Backend Capability Fixture Probe Harness

### Mission

Add a deterministic fixture probe harness for backend capabilities so the Runtime can explain compatibility profiles without starting any backend.

### Start from

- `internal/runtime/appidentity/backend_manager.go`
- `internal/runtime/appidentity/backend_capability_matrix.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/backend_selection.go`
- `internal/runtime/owner/dispatch.go`
- `cmd/xnix-runtime-go/windows_compatibility_commands.go`
- `test/test_compatibility_backend_capability_matrix.rb`
- `test/test_compatibility_backend_selection_plan.rb`

### Deliver

- A Go read model and CLI command such as `backend-capability-probe-preview`.
- Fixture-backed probe results for profile availability, required artifacts, missing dependencies, isolation requirement, graphics capability, audio capability, input capability, and network policy.
- Deterministic degraded and unsupported states.
- Tests proving probe fixtures cannot start processes, invoke package managers, fetch network artifacts, or expose backend commands.

### Keep disabled

- Backend process start.
- VM start.
- Package-manager calls.
- Network fetch.
- Runtime profile creation.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/owner ./cmd/xnix-runtime-go
ruby -Ilib test/test_compatibility_backend_capability_matrix.rb
ruby -Ilib test/test_compatibility_backend_selection_plan.rb
ruby scripts/implementation_evidence_report.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement FW4 from docs/claude-code-fourth-wave-task-batch.md.

Target outcome:
- Add a fixture-backed backend capability probe preview that explains Runtime compatibility profile availability without starting any backend.

Hard constraints:
- Do not start Wine, Proton, a VM, or any backend process; do not call package managers; do not fetch network artifacts; do not create Runtime profiles; do not run Docker/QEMU; do not mutate the host root.
- Do not expose raw backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add available, degraded, unsupported, malformed-fixture, no-process-start, no-network, and no-package-manager tests.
- Run the FW4 verification commands and report exact commands run.
```

## FW5: Settings Change Dependency Review

### Mission

Make KDE unified-settings changes depend on explicit Runtime evidence, receipts, Portal policy, snapshot readiness, and write-gate state before any setting can become actionable.

### Start from

- `internal/runtime/appidentity/desktop_safety_policy.go`
- `internal/runtime/appidentity/kde_center_page.go`
- `lib/xnix/compatibility/settings_change_plan.rb`
- `lib/xnix/compatibility/compatibility_review_flow_plan.rb`
- `internal/runtime/appidentity/runtime_write_gate.go`
- `test/test_settings_change_plan.rb`
- `test/test_compatibility_review_flow_plan.rb`
- `test/test_runtime_write_gate.rb`

### Deliver

- A Go read model and CLI command such as `settings-change-dependency-preview`.
- Per-setting dependency rows for run mode, priority, documents, downloads, camera, network, snapshots, diagnostics, and notifications.
- Missing-evidence and blocked-reason summaries for each setting.
- Tests proving no setting is persisted, no permission is granted, no request object is created, and no backend is launched.

### Keep disabled

- Settings persistence.
- Permission grants.
- Portal request creation.
- Runtime write methods.
- Backend launch.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./cmd/xnix-runtime-go
ruby -Ilib test/test_settings_change_plan.rb
ruby -Ilib test/test_compatibility_review_flow_plan.rb
ruby -Ilib test/test_runtime_write_gate.rb
ruby scripts/kde_first_presence_smoke.rb --format json
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement FW5 from docs/claude-code-fourth-wave-task-batch.md.

Target outcome:
- Add a settings change dependency preview that explains which Runtime evidence and receipts are required before KDE unified-settings changes can become actionable.

Hard constraints:
- Do not persist settings, grant permissions, create Portal requests, enable Runtime writes, start backends, run Docker/QEMU, or mutate the host root.
- Do not expose backend terms in KDE-facing output; do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add per-setting dependency, missing-evidence, blocked-write-gate, invalid-field, invalid-value, and no-side-effect tests.
- Run the FW5 verification commands and report exact commands run.
```

## FW6: Diagnostics Repair Playbook Review Queue

### Mission

Turn diagnostics and AI recommendations into a human-reviewable repair playbook queue without applying repairs.

### Start from

- `internal/runtime/appidentity/ai_diagnostics.go`
- `internal/runtime/appidentity/diagnostic_history.go`
- `internal/runtime/diagnostics/record.go`
- `internal/runtime/diagnostics/history.go`
- `internal/runtime/appidentity/repair_plan.go`
- `test/test_ai_diagnostic_input.rb`
- `test/test_ai_diagnostic_recommendation.rb`
- `test/test_ai_repair_approval_gate.rb`
- `test/test_compatibility_repair_plan.rb`

### Deliver

- A Go read model and CLI command such as `diagnostic-repair-playbook-preview`.
- Playbook items with diagnostic signal ids, redacted evidence ids, suggested repair category, user confirmation requirement, risk level, blocked reason, and next read-only check.
- Fixture-based privacy tests for file path, username, environment variable, token-shaped, and command-shaped data.
- Tests proving no AI provider call, no auto-repair, no backend launch, and no host mutation occurs.

### Keep disabled

- AI provider calls.
- Automatic repair.
- Backend launch.
- Package-manager calls.
- File writes outside explicit state roots.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/diagnostics ./cmd/xnix-runtime-go
ruby -Ilib test/test_ai_diagnostic_input.rb
ruby -Ilib test/test_ai_diagnostic_recommendation.rb
ruby -Ilib test/test_ai_repair_approval_gate.rb
ruby -Ilib test/test_compatibility_repair_plan.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement FW6 from docs/claude-code-fourth-wave-task-batch.md.

Target outcome:
- Add a diagnostic repair playbook preview that turns diagnostics and AI recommendations into a human-reviewable queue without applying repairs.

Hard constraints:
- Do not call AI providers, apply repairs, start backends, call package managers, write outside explicit state roots, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, credentials, usernames, or environment variables.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add playbook, privacy-redaction, malformed-fixture, provider-disabled, auto-repair-disabled, and no-host-mutation tests.
- Run the FW6 verification commands and report exact commands run.
```

## FW7: Desktop Activation Materialization Audit

### Mission

Add a reviewer-facing materialization audit for staged KDE activation artifacts before any host installation is allowed.

### Start from

- `internal/runtime/appidentity/desktop_activation_staging.go`
- `internal/runtime/appidentity/desktop_activation_transaction.go`
- `internal/runtime/appidentity/desktop_activation_status.go`
- `internal/runtime/activation/stage.go`
- `cmd/xnix-runtime-go/desktop_activation_stage_commands.go`
- `test/test_desktop_activation_installer.rb`
- `test/test_desktop_activation_rollback.rb`
- `test/test_desktop_integration_manifest.rb`

### Deliver

- A Go read model and CLI command such as `desktop-activation-audit-preview`.
- A safe diff of staged desktop entry, service menu, MIME association, icon, manifest, and rollback receipt ids.
- Digest validation and missing-file diagnostics without exposing file contents or host paths.
- Tests proving the audit does not install files, refresh KDE caches, change MIME defaults, or mutate the host root.

### Keep disabled

- Host desktop activation.
- KDE cache refresh.
- MIME default writes.
- Icon installation.
- Rollback execution.
- Host-root mutation.

### Required verification

```text
go test ./internal/runtime/appidentity ./internal/runtime/activation ./cmd/xnix-runtime-go
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement FW7 from docs/claude-code-fourth-wave-task-batch.md.

Target outcome:
- Add a desktop activation audit preview that reviews staged KDE activation artifacts and rollback receipts before host installation is allowed.

Hard constraints:
- Do not install files into the host, refresh KDE caches, change MIME defaults, install icons, execute rollback, run Docker/QEMU, or mutate the host root.
- Do not expose raw executable paths, staging-root paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add complete-stage, missing-file, digest-mismatch, malformed-manifest, rollback-unavailable, and no-host-install tests.
- Run the FW7 verification commands and report exact commands run.
```

## FW8: Restricted Release Readiness Packet

### Mission

Create a safe release-readiness packet that summarizes local verification evidence and explicitly lists blocked heavy tests without running Docker or QEMU.

### Start from

- `scripts/verify_layout.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/implementation_evidence_report.rb`
- `scripts/mainline_integration_review.rb`
- `scripts/kde_first_presence_smoke.rb`
- `docs/kde-first-presence-smoke-spec.md`
- `docs/kde-first-compatibility-acceptance.md`
- `test/test_implementation_evidence_report.rb`
- `test/test_mainline_integration_review.rb`

### Deliver

- A Ruby report script such as `scripts/release_readiness_packet.rb` with JSON and Markdown output.
- A report schema that includes version, dirty-worktree summary, protected-file status, required local commands, last known safe command outputs when provided by fixture input, blocked heavy checks, and human authorization requirements.
- Fixture tests that prove the script does not execute Docker, QEMU, Colima, network fetch, package managers, or host mutation commands.
- Documentation in `PRODUCT_OVERVIEW.md` describing the report as a read-only readiness packet, not a release command.

### Keep disabled

- Docker execution.
- QEMU execution.
- Colima start or stop.
- Network fetch.
- Package-manager calls.
- Git staging, committing, tagging, or pushing.
- Host-root mutation.

### Required verification

```text
ruby -Ilib test/test_implementation_evidence_report.rb
ruby -Ilib test/test_mainline_integration_review.rb
ruby scripts/release_readiness_packet.rb --format json
ruby scripts/release_readiness_packet.rb --format markdown
ruby scripts/verify_layout.rb
git diff --check
```

### Copyable Claude Code prompt

```text
Implement FW8 from docs/claude-code-fourth-wave-task-batch.md.

Target outcome:
- Add a read-only release readiness packet script that summarizes safe local evidence and names blocked heavy tests without running them.

Hard constraints:
- Do not run Docker, QEMU, Colima, network fetches, package managers, git staging, commits, tags, pushes, or host-root mutation.
- Do not expose host paths, file contents, secrets, tokens, private keys, or credentials.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Do not modify docs/claude-code-implementation-packages.md.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md for code changes.
- Add JSON, Markdown, fixture-input, blocked-heavy-check, protected-file, dirty-worktree, and no-command-execution tests.
- Run the FW8 verification commands and report exact commands run.
```
