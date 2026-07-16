# Claude Code Large Empty-Domain Assignments

> Last updated: 2026-07-16 | Baseline: v0.2.246

This document is a coarse assignment map for handing large, relatively independent Xnix domains to Claude Code.

Use it when the current problem is: "the contract exists, but the domain still has little durable implementation." It intentionally groups smaller briefs into branch-sized workstreams that can produce reviewable implementation evidence without asking one branch to finish the whole product.

## Relationship to Other Handoff Documents

- `docs/claude-code-domain-dispatch.md` is the first dispatch board for choosing large independent domains.
- `docs/claude-code-mainline-implementation-plan.md` is the strategic mainline plan.
- `docs/claude-code-contract-gap-work-packages.md` is the finer contract-gap package list.
- `docs/claude-code-empty-domain-implementation-packages.md` defines evidence levels and completion rules.
- `docs/claude-code-independent-implementation-briefs.md` contains smaller prompt-ready briefs.
- This file is the coarse dispatcher: pick one large assignment here, then let Claude Code drill down into the referenced finer packages.

## Single-Document Handoff Protocol

When handing work to Claude Code, start from this document unless a smaller referenced package is already chosen.

Use this protocol:

1. Pick exactly one assignment from `L1` through `L8`.
2. Copy the assignment's prompt into Claude Code.
3. Tell Claude Code to read the referenced finer packages, but not to expand scope beyond the selected assignment.
4. Require Claude Code to stop and report if completing the assignment would need a forbidden host operation, production write path, production bus ownership, real backend launch, or network fetch.
5. Require a final handoff using the completion statement template at the end of this file.

Parallel work is allowed only when assignments do not touch the same ownership boundary:

| Safe to run in parallel | Avoid running in parallel |
| --- | --- |
| `L2` with `L5`, if `L5` only consumes existing receipts | `L1` with any branch changing D-Bus owner routing |
| `L4` with `L7`, if report contracts are explicitly coordinated | `L2` with `L3`, when both change install-readiness semantics |
| `L6` with `L8`, if `L8` does not depend on new diagnostics | `L4` with `L8`, when both change QEMU or smoke safety gates |

If there is only one Claude Code runner available, use the first recommended wave order: `L1`, `L2`, `L3`, then `L7`.

## Evidence Targets

Each assignment should convert contract-only surfaces into at least one durable evidence level:

| Evidence level | Meaning | Example |
| --- | --- | --- |
| Fixture implemented | Logic runs against deterministic fixture data. | Recipe digest mismatch fails closed. |
| State-root implemented | Logic persists safe records under an explicit Runtime or test root. | Environment lifecycle state survives a restart inside a test root. |
| Smoke owned | A constrained smoke exercises the Runtime boundary. | A private session bus read is served by the Go owner. |
| Production gated | Production behavior is represented, observable, and disabled until prerequisites exist. | Write methods return stable disabled errors. |

Do not accept a branch that only adds new previews, new documentation, or new contract surfaces unless the selected assignment is `L7` and the branch adds evidence gates that prevent future empty contracts.

## Assignment Rules

Every Claude Code branch must:

- Implement exactly one large assignment unless this file explicitly says otherwise.
- Keep all source, comments, tests, fixtures, CLI output, and documentation in English.
- Prefer Go for durable Compatibility Runtime product behavior.
- Use C only for low-level Runtime transport, ABI-shaped records, or already-owned C policy surfaces.
- Use Ruby for tests, smoke scripts, reports, and developer tooling.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for every code change.
- Add or update targeted tests before claiming the assignment is complete.
- Run the package tests plus `ruby scripts/verify_layout.rb`.
- Run `ruby scripts/runtime_contract_drift_report.rb --format json` when touching D-Bus XML, Runtime route manifests, owner dispatch, smoke adapters, or Runtime CLI commands.
- Keep unsafe behavior disabled unless the assignment explicitly enables a gated test-only path.

Unsafe behavior that remains disabled by default:

- Production D-Bus ownership.
- Write methods.
- Real backend launch.
- Real Portal transport calls.
- Network artifact fetches.
- Privileged containers.
- Host networking.
- Docker socket mounts.
- Broad host-directory mounts.
- Host-root mutation.
- Raw backend command, executable path, storage path, profile name, file content, secret, token, or private-key exposure in KDE-facing output.

## How to Pick the Next Assignment

| Need | Pick | Why |
| --- | --- | --- |
| The Runtime owner still feels like previews plus adapters | L1 | It turns read dispatch into a constrained process boundary. |
| Install inputs are modeled but not trustworthy enough | L2 | It joins recipe trust, artifact verification, cache, staging, and install gates. |
| Backends cannot yet have durable lifecycle state | L3 | It creates state-root lifecycle records before launch exists. |
| Execution safety prerequisites are not real enough | L4 | It builds Portal, snapshot, review, and rollback safety state. |
| KDE surfaces still consume static previews | L5 | It makes launcher, MIME, tray, notification, settings, and shell reads consume durable receipts. |
| Diagnostics and repair are mostly plans | L6 | It records fixture diagnostics and repair recommendations without live providers. |
| Contracts keep multiplying faster than implementation evidence | L7 | It adds gates that fail on orphan contracts and preview-only regressions. |
| The product needs bootable proof | L8 | It validates the atomic KDE/QEMU path without weakening host safety. |

## First Recommended Wave

Start with L1, L2, L3, and L7.

These assignments build ownership, trust, lifecycle, and evidence gates. They reduce the risk that later KDE or execution branches merely add more user-visible previews on top of empty implementation.

## L1: Runtime Owner Process Boundary

### Mission

Turn the Runtime owner from in-process preview dispatch into a constrained Go process that can own a private test session-bus name, serve read-only Runtime methods, and fail write methods closed.

### References

- Mainline package: `M1`
- Contract-gap packages: `P1`, `P2`
- Empty-domain package: `P1`

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

- A long-running smoke-owner mode for `xnix-runtime-owner`.
- Private session-bus ownership only inside constrained smoke tests.
- Read-only D-Bus method routing to Go owner handlers or explicit unsupported-read markers.
- Stable `WriteMethodDisabled` responses for every write method.
- Owner lifecycle logs for startup, route table version, bus claim mode, readiness, and shutdown.
- A parity gate proving the D-Bus XML, Go owner routes, C smoke adapter, Ruby D-Bus client, CLI route commands, and write gates agree.

### Do not deliver

- Production bus ownership.
- System service installation.
- Backend launch.
- Write enablement.
- Host-root writes.
- KDE-owned Runtime policy.

### Acceptance

- A constrained container smoke starts the Go owner on a private session bus.
- Read methods match existing safe payloads or return explicit unsupported-route readiness.
- Reserved write methods fail closed through the owner process.
- Contract drift reporting fails when a read method exists in only one layer.
- Owner output remains safe for logs and KDE-facing projections.

### Required verification

```text
go test ./...
ruby scripts/runtime_contract_drift_report.rb --format json
ruby scripts/container.rb runtime-owner-candidate-smoke
ruby scripts/container.rb runtime-dbus-smoke
ruby scripts/container.rb kde-center-dbus-smoke
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement L1 from docs/claude-code-large-empty-domain-assignments.md.

Target outcome:
- Add a constrained Go Runtime owner smoke process that can own a private test session-bus name, serve read-only Runtime methods, and fail write methods closed.

Constraints:
- Do not enable production D-Bus ownership, write methods, backend launch, service installation, network access, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, failure, and blocked-state tests.
- Run the required verification commands listed for L1.
- Stop and report blockers if implementation requires production ownership or host mutation.
```

## L2: Trust, Artifact, Cache, and Install Pipeline

### Mission

Turn recipe trust, artifact acquisition, cache, staging, and install gates into one production-shaped but fixture-first pipeline.

### References

- Mainline package: `M2`
- Contract-gap packages: `P3`, `P4`
- Empty-domain package: `P2`

### Start from

- `internal/runtime/recipe/`
- `internal/runtime/artifact/`
- `internal/runtime/appidentity/registry.go`
- `internal/runtime/appidentity/runtime_owner_recipe_trust.go`
- `internal/runtime/appidentity/artifact_manifest.go`
- `internal/runtime/appidentity/acquisition_preflight.go`
- `internal/runtime/appidentity/install_plan.go`
- `runtime/recipes/registry.json`
- `lib/xnix/compatibility/recipe_*`

### Deliver

- A read-only recipe store interface with local roots.
- Registry validation for schema, IDs, relative paths, digests, and declared signing state.
- A replaceable signature verifier boundary.
- Fixture artifact manifest parsing and digest verification.
- Runtime-root-scoped cache and staging receipts.
- Install readiness that joins recipe trust, artifact verification, cache state, staging state, and owner readiness.

### Do not deliver

- Remote downloads by default.
- Private keys or signing secrets.
- Host package-manager calls.
- Host-root mutation.
- Desktop activation writes.
- Backend process starts.

### Acceptance

- Invalid recipe digests fail closed.
- Invalid artifact digests block staging.
- Development fixtures remain usable but cannot pass production trust.
- Cache and staging stay under explicit Runtime or test roots.
- Install gates consume the same trust and artifact evidence shown to the Compatibility Center.

### Required verification

```text
go test ./...
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
Implement L2 from docs/claude-code-large-empty-domain-assignments.md.

Target outcome:
- Build a fixture-first trust, artifact, cache, staging, and install-readiness pipeline that fails closed and never mutates the host root.

Constraints:
- Do not add default network fetch, package-manager calls, private keys, production trust bypasses, desktop writes, backend launch, privileged containers, or host-root mutation.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, digest-mismatch, unsigned, invalid, and blocked-state tests.
- Run the required verification commands listed for L2.
```

## L3: Environment Lifecycle and State Root

### Mission

Convert backend environment plans into durable Runtime lifecycle state without starting compatibility backends.

### References

- Mainline package: `M3`
- Contract-gap package: `P5`
- Empty-domain package: `P3`

### Start from

- `internal/runtime/environment/`
- `internal/runtime/appidentity/backend_environment.go`
- `internal/runtime/appidentity/backend_binding.go`
- `internal/runtime/appidentity/backend_lifecycle.go`
- `internal/runtime/appidentity/execution_readiness.go`
- `internal/runtime/appidentity/application_state_root.go`

### Deliver

- A Runtime state-root lifecycle store.
- Environment states such as `missing`, `planned`, `staged`, `ready`, `repair-required`, and `blocked`.
- Transitions driven by recipe trust, artifact staging, backend binding, and Portal/snapshot prerequisites.
- KDE-safe readiness explanations.
- Deterministic blocked reasons for missing trust, missing artifacts, missing state root, or unsupported backend capability.

### Do not deliver

- Backend process starts.
- Launch transactions.
- Raw command exposure.
- Host storage path exposure.
- Host-root mutation.

### Acceptance

- Lifecycle state survives inside a test-controlled state root.
- Readiness changes when trust or staging evidence changes.
- Execution readiness consumes lifecycle state without enabling launch.
- KDE-facing output remains backend-detail safe.

### Required verification

```text
go test ./...
ruby -Ilib test/test_application_state_root.rb
ruby -Ilib test/test_compatibility_backend_environment_plan.rb
ruby -Ilib test/test_compatibility_backend_binding.rb
ruby -Ilib test/test_compatibility_backend_lifecycle.rb
ruby -Ilib test/test_compatibility_execution_readiness.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement L3 from docs/claude-code-large-empty-domain-assignments.md.

Target outcome:
- Add durable Runtime state-root environment lifecycle records and readiness transitions without starting any backend.

Constraints:
- Do not launch backends, create execution requests, expose raw commands or host paths, require network, use privileged containers, or mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, blocked, corrupted-state, and transition tests.
- Run the required verification commands listed for L3.
```

## L4: Portal, Snapshot, Review, and Rollback Safety Plane

### Mission

Build the safety prerequisites that must exist before real execution can be enabled: fake-first Portal requests, content-addressed snapshots, review receipts, and rollback records.

### References

- Mainline package: `M4`
- Contract-gap packages: `P6`, `P7`
- Empty-domain package: `P4`

### Start from

- `internal/runtime/appidentity/portal_access_policy.go`
- `internal/runtime/appidentity/portal_request.go`
- `internal/runtime/appidentity/snapshot_plan.go`
- `internal/runtime/appidentity/action_review_receipt.go`
- `internal/runtime/appidentity/permission_review_plan.go`
- `internal/runtime/appidentity/review_flow_plan.go`
- snapshot and Portal Ruby tests under `test/`

### Deliver

- Durable fake-mode Portal request state.
- Explicit disabled boundary for real Portal transport.
- Content-addressed snapshot records under a controlled root.
- Review receipts that can be joined to permission, snapshot, and rollback decisions.
- Rollback readiness that can be consumed by execution and KDE status surfaces.

### Do not deliver

- Real Portal method calls.
- Permission grants.
- File-content reads.
- Snapshotting host roots.
- Restore execution.
- Host-root mutation.

### Acceptance

- Fake Portal requests persist and can be inspected under a test state root.
- Snapshot records include digest, schema version, app id, operation id, and relative receipt paths.
- Rollback readiness changes when snapshot evidence is missing or invalid.
- KDE-facing output avoids file contents and raw host paths.

### Required verification

```text
go test ./...
ruby -Ilib test/test_portal_access_policy.rb
ruby -Ilib test/test_portal_request_model.rb
ruby -Ilib test/test_compatibility_snapshot_plan.rb
ruby -Ilib test/test_action_review_receipt.rb
ruby -Ilib test/test_compatibility_permission_review_plan.rb
ruby -Ilib test/test_compatibility_review_flow_plan.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement L4 from docs/claude-code-large-empty-domain-assignments.md.

Target outcome:
- Add durable fake-mode Portal, snapshot, review, and rollback safety records under a controlled Runtime state root.

Constraints:
- Do not call real Portal transports, grant permissions, read user file contents, snapshot host roots, execute restores, launch backends, or mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, denied, missing-snapshot, invalid-snapshot, and blocked-state tests.
- Run the required verification commands listed for L4.
```

## L5: KDE Activation and Shell Materialization

### Mission

Move KDE-facing surfaces from static previews toward durable activation receipts and safe shell adapters while keeping KDE presentation-only.

### References

- Mainline package: `M5`
- Contract-gap packages: `P9`, `P10`
- Empty-domain package: `P5`

### Start from

- `internal/runtime/appidentity/desktop_activation_*.go`
- `internal/runtime/appidentity/desktop_entry*.go`
- `internal/runtime/appidentity/file_association*.go`
- `internal/runtime/appidentity/task_manager_identity.go`
- `internal/runtime/appidentity/kwin_window_rule.go`
- `internal/runtime/appidentity/tray_status.go`
- `internal/runtime/appidentity/notification*.go`
- `kde/`
- `scripts/kde_first_presence_smoke.rb`

### Deliver

- Activation receipt consumption for launcher, MIME, Dolphin, tray, notifications, settings, KRunner, task manager, and KWin views.
- Drift checks between staged activation files and KDE-safe read models.
- Minimal shell adapter fixtures that prove KDE code consumes Runtime-owned read models.
- Optional digest-verified icon handling, if it stays fixture-only and root-scoped.

### Do not deliver

- KDE-owned backend policy.
- Direct Wine or VM calls from KDE.
- Live tray bridging by default.
- KWin rule application to the host session.
- Production desktop writes.
- Host-root mutation.

### Acceptance

- KDE status surfaces can report whether activation receipts exist and match expected digests.
- Mismatched or missing activation receipts produce safe blocked states.
- KDE-facing output never exposes raw executable paths, compatibility storage paths, backend commands, or profile names.
- Shell adapter tests run without a live Plasma session.

### Required verification

```text
go test ./...
ruby -Ilib test/test_desktop_activation_installer.rb
ruby -Ilib test/test_desktop_activation_rollback.rb
ruby -Ilib test/test_desktop_integration_manifest.rb
ruby -Ilib test/test_kde_integration_status.rb
ruby -Ilib test/test_kde_shell_integration_plan.rb
ruby -Ilib test/test_kde_application_surface_plan.rb
ruby -Ilib test/test_task_manager_identity.rb
ruby -Ilib test/test_kwin_window_rule.rb
ruby -Ilib test/test_file_association_model.rb
ruby -Ilib test/test_tray_status_model.rb
ruby -Ilib test/test_krunner_model.rb
ruby scripts/kde_first_presence_smoke.rb
ruby scripts/verify_layout.rb
```

### Copyable prompt

```text
Implement L5 from docs/claude-code-large-empty-domain-assignments.md.

Target outcome:
- Make KDE-facing launcher, MIME, Dolphin, tray, notification, settings, KRunner, task-manager, and KWin reads consume durable activation receipt evidence where available.

Constraints:
- Do not let KDE own Runtime policy, call Wine or VMs, apply host KWin rules, write production desktop files, start live tray bridges, use privileged containers, or mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, missing-receipt, digest-mismatch, and blocked-state tests.
- Run the required verification commands listed for L5.
```

## L6: Diagnostics, Repair, and AI Provider Boundary

### Mission

Replace diagnostic and repair previews with fixture-driven records, review-first repair recommendations, and an explicitly disabled production AI provider boundary.

### References

- Mainline package: `M7`
- Contract-gap package: `P11`
- Empty-domain package: `P7`

### Start from

- `internal/runtime/appidentity/diagnostics*.go`
- `internal/runtime/appidentity/test_plan.go`
- `internal/runtime/appidentity/test_result.go`
- `internal/runtime/appidentity/repair_plan.go`
- `internal/runtime/appidentity/ai_diagnostic_*.go`
- `internal/runtime/appidentity/ai_repair_approval_gate.go`
- diagnostic and AI tests under `test/`

### Deliver

- Fixture diagnostic run records under a controlled state root.
- Diagnostic history summaries consumed by Compatibility Center reads.
- Repair recommendations that require user review and never auto-execute.
- An AI provider interface with fake and disabled implementations.
- Privacy filters that block file contents, raw paths, backend commands, secrets, and tokens from AI inputs.

### Do not deliver

- Real provider calls by default.
- Auto-repair.
- File-content reads.
- Network requirements.
- Backend launch.
- Host-root mutation.

### Acceptance

- Diagnostic success and failure fixtures produce durable records.
- Diagnostic history can be read without exposing state-root paths or file contents.
- AI provider calls are disabled unless explicitly configured for a future gated test path.
- Repair recommendations are review-only and carry approval gates.

### Required verification

```text
go test ./...
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
Implement L6 from docs/claude-code-large-empty-domain-assignments.md.

Target outcome:
- Add fixture-driven diagnostic records, diagnostic history reads, review-first repair recommendations, and a disabled-by-default AI provider boundary.

Constraints:
- Do not call real AI providers, auto-repair, read user file contents, require network, launch backends, expose secrets or paths, or mutate the host root.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, failure, redaction, provider-disabled, and approval-required tests.
- Run the required verification commands listed for L6.
```

## L7: Developer Verification and Evidence Gates

### Mission

Prevent empty-domain regressions by making implementation evidence measurable, reportable, and CI-friendly.

### References

- Mainline package: `M8`
- Contract-gap package: `P2`
- Empty-domain package: `P9`

### Start from

- `scripts/implementation_evidence_report.rb`
- `scripts/runtime_contract_drift_report.rb`
- `scripts/verify_layout.rb`
- `test/test_implementation_evidence_report.rb`
- `test/test_runtime_contract_drift_report.rb`

### Deliver

- Orphan contract detection for Runtime read methods, CLI commands, smoke adapters, and docs.
- Evidence-level regression guards for contract-only, fixture-implemented, state-root-implemented, smoke-owned, and production-gated domains.
- JSON and Markdown reports suitable for CI artifacts.
- Package ownership checks that map changed files to expected assignment IDs.
- Clear next-dispatch guidance that prefers implementation over new previews.

### Do not deliver

- Product behavior changes.
- QEMU requirements for ordinary report runs.
- Network requirements.
- Docker requirements for ordinary report runs.
- Host-root access.

### Acceptance

- Adding a new Runtime method without implementation evidence fails a report or test.
- Reports remain deterministic.
- Reports can run on a clean checkout without Docker, QEMU, network, privileged containers, backend launch, or host-root mutation.
- The report identifies the next recommended implementation package.

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
Implement L7 from docs/claude-code-large-empty-domain-assignments.md.

Target outcome:
- Add CI-friendly evidence gates that fail on orphan contracts, preview-only regressions, and missing assignment ownership.

Constraints:
- Do not require Docker, QEMU, network, backend launch, privileged containers, or host-root access for ordinary report runs.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add success, missing-evidence, orphan-contract, and deterministic-output tests.
- Run the required verification commands listed for L7.
```

## L8: Atomic KDE Image and QEMU Acceptance

### Mission

Create the product-level acceptance path that proves the KDE-first Xnix shape can be built and boot-smoked without weakening host safety.

### References

- Mainline package: `M9`
- Contract-gap package: `P12`
- Empty-domain package: `P8`

### Start from

- `boot/`
- `buildroot/`
- `scripts/container.rb`
- `scripts/full_smoke.rb`
- `scripts/qemu_smoke.rb`
- `docs/`

### Deliver

- A reproducible image manifest for the atomic KDE target.
- A constrained QEMU acceptance plan for boot, serial logging, loopback-only SSH, Runtime service presence, and KDE package presence.
- A no-host-mutation build/test path suitable for Colima.
- Clear separation between the low-level Buildroot learning baseline and the KDE-first product image target.
- JSON and Markdown smoke reports, if a new report script is added.

### Do not deliver

- Host package installation.
- Host network exposure.
- Privileged containers.
- Broad host mounts.
- Docker socket mounts.
- Production backend launch.
- SSH exposure beyond loopback-bound forwarded ports.

### Acceptance

- A clean checkout can reproduce the build plan.
- QEMU serial logs are persisted for diagnosis.
- SSH, if enabled in a smoke, is loopback-bound and key-authenticated.
- The smoke fails if host-risk flags are requested.
- Product documentation explains what is verified and what remains pending.

### Required verification

```text
ruby scripts/verify_layout.rb
ruby scripts/container.rb build
ruby scripts/full_smoke.rb
```

Run the repository's full QEMU milestone gate if the version lands on a tenth code version.

### Copyable prompt

```text
Implement L8 from docs/claude-code-large-empty-domain-assignments.md.

Target outcome:
- Add a reproducible atomic KDE image/QEMU acceptance path that keeps host impact minimal and produces persisted smoke evidence.

Constraints:
- Do not install host packages, expose host networking, use privileged containers, mount the Docker socket, broadly mount host directories, launch production backends, or mutate the host root.
- Keep SSH smoke exposure loopback-bound and key-authenticated.
- Keep source, comments, tests, fixtures, CLI output, and docs in English.

Required handoff:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md.
- Add build-plan, host-risk-denial, serial-log, and smoke-report tests.
- Run the required verification commands listed for L8.
- If the resulting version is a tenth code version, run the full QEMU milestone gate.
```

## Completion Statement Template

Each Claude Code branch should end with this handoff statement:

```text
Assignment:
- <L1-L8 name>

Converted contracts:
- <contract surfaces that now have implementation evidence>

Implementation evidence added:
- <fixture/state-root/smoke/report evidence>

Still gated:
- <unsafe production behavior that remains disabled>

Verification run:
- <commands and results>

Files intentionally not touched:
- <unrelated package areas>
```
