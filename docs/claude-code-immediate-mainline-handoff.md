# Claude Code Immediate Mainline Handoff

> Last updated: 2026-07-24 | Baseline: v0.2.640-rc21 | Formal release: v0.2.640 is still blocked until full smoke passes

This is the short handoff document for asking Claude Code to perform the next Xnix mainline implementation work. It is intentionally separate from `docs/claude-code-implementation-packages.md`; do not edit that protected file while working from this handoff.

Use this document when the operator wants Claude Code to implement one bounded branch while Codex keeps mainline review, merge decisions, release promotion, and host-safety decisions.

## Current Product Direction

Xnix is converging on the best Linux desktop experience for existing Windows applications. KDE Plasma is the first supported desktop shell, but it must stay replaceable. The independent Xnix AI Compatibility Runtime owns compatibility policy, evidence, receipts, state roots, backend selection, launch planning, diagnostics, snapshots, rollback, and execution gates.

The current near-term milestone is still:

- keep the known Windows application path as the real mainline;
- preserve the staged `xnix-compat-launch` QEMU/Wine lane;
- route desktop-triggered launches through Runtime-owned evidence and owner-service materialization;
- keep KDE presentation-only;
- promote `v0.2.640` only after the human-authorized full smoke passes.

## Baseline State

The local baseline is `v0.2.640-rc21`.

Use `docs/claude-code-active-work-order.md` as the short current Claude Code handoff for the next bounded implementation task.

Already present:

- `desktop-trigger-staged-invocation-readiness-preview` from `v0.2.640-rc3`.
- `owner-service-launch-envelope-guard-preview` from `v0.2.640-rc4`.
- `kde-controlled-launch-action-surface-audit-preview` from `v0.2.640-rc5`.
- `managed-launcher-acceptance-report-preview` from `v0.2.640-rc6`.
- `desktop-trigger-dry-run-request-review-preview` from `v0.2.640-rc7`.
- `desktop-trigger-service-call-materialization-preview` from `v0.2.640-rc8`.
- The real staged launcher dispatch smoke consumes `desktop-trigger-service-call-materialization-preview` with an explicit human-authorized smoke candidate flag from `v0.2.640-rc9`.
- The private D-Bus controlled-launch fixture consumes the same `desktop-trigger-service-call-materialization-preview` path from `v0.2.640-rc10`.
- `scripts/full_checkpoint_promotion_packet.rb` reads existing full-smoke reports and produces the `v0.2.640` promotion decision without running full smoke from `v0.2.640-rc11`.
- `scripts/merge_readiness_packet.rb` consumes the promotion packet as a first-class release gate from `v0.2.640-rc12`.
- `scripts/release_evidence_index.rb` consumes the promotion packet from `v0.2.640-rc13`, keeping old product smoke evidence separate from current formal release readiness.
- `docs/post-checkpoint-promotion-checklist-0640.md` defines the human-owned formal promotion checklist from `v0.2.640-rc14`.
- `desktop-trigger-request-preflight-preview` is present from `v0.2.640-rc15`; it consumes service-call materialization but stays `blocked-missing-promotion` until formal full checkpoint promotion is observed.
- `scripts/desktop_trigger_request_preflight_smoke.rb` is present from `v0.2.640-rc16`; it verifies the Go preflight's blocked and promoted fixture states without running Docker, QEMU, Wine, D-Bus, KDE, or a backend.
- `scripts/merge_readiness_packet.rb` consumes existing desktop-trigger request preflight smoke JSON evidence from `v0.2.640-rc17` without running that smoke by default.
- `scripts/release_evidence_index.rb` classifies existing desktop-trigger request preflight smoke JSON evidence from `v0.2.640-rc18` without running that smoke by default.
- `docs/claude-code-next-dispatch-brief.md` is present from `v0.2.640-rc16` as the next short copy-first Claude dispatch brief.
- The existing staged smoke already reaches the Runtime owner service through `xnix-runtime-owner --service-call ShowRuntimeControlledLaunch ...` when QEMU/Wine prerequisites are available.

Do not ask Claude Code to reimplement completed checkpoint packets unless Codex or a reviewer explicitly asks for a repair branch.

## Dispatch Rules

Assign exactly one task from this document at a time.

Before starting, Claude Code must run:

```text
git status --short --branch
```

If the checkout is dirty, Claude Code must either continue the dirty lane or stop and ask for review direction. It must not start an unrelated branch on top of unknown local changes.

Claude Code must not modify:

```text
docs/claude-code-implementation-packages.md
```

Every implementation branch must:

- keep all source, comments, tests, fixtures, CLI output, and documentation in English;
- prefer Go for durable Runtime product behavior and stable service boundaries;
- use Ruby only for tests, smoke orchestration, reports, fixtures, and lightweight tooling;
- keep KDE as a replaceable presentation shell;
- keep Runtime policy out of KDE, Ruby smoke scripts, and the C smoke adapter;
- update `VERSION`, `CHANGELOG.md`, `PRODUCT_OVERVIEW.md`, `docs/xnix-current-mainline.md`, and layout/review gates when code changes;
- add targeted tests for ready, blocked, malformed, stale/redaction, and no-side-effect behavior where relevant;
- run task-specific verification, `ruby scripts/verify_layout.rb`, and `git diff --check`;
- report exact commands run and exact commands skipped.

Claude Code must not:

- run Docker, QEMU, Wine, Proton, Colima, or a real compatibility backend;
- run `ruby scripts/full_smoke.rb`;
- pull images, fetch packages, restart host services, use host package managers, or repair host networking;
- use privileged containers, host networking, Docker socket mounts, broad host-directory mounts, or host-root mutation;
- enable production D-Bus ownership;
- add real Portal transport calls;
- write Runtime production state unless the selected task explicitly asks for a state-root-scoped preview writer;
- write KDE configuration;
- expose state-root paths, cache-root paths, launcher paths, executable paths, backend names, backend commands, raw launcher output, host paths, secrets, tokens, credentials, usernames, or environment variables in KDE-facing or user-facing output.

## Immediate Dispatch Queue

Use this order unless Codex or a reviewer asks for a repair branch.

| Order | Task | Suggested branch | Outcome |
| --- | --- | --- | --- |
| Done | `C8W9` Route staged smoke through materialized service-call arguments | `codex/staged-smoke-service-call-materialization` | Present locally in v0.2.640-rc9. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W10` Route the D-Bus controlled-launch fixture through the same materialization packet | `codex/dbus-fixture-service-call-materialization` | Present locally in v0.2.640-rc10. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W11` Add a release-promotion evidence packet for `v0.2.640` | `codex/full-checkpoint-promotion-packet` | Present locally in v0.2.640-rc11. Do not dispatch again unless a reviewer asks for repair. |

Stop after each task and return the branch for Codex review. Do not chain tasks.

## Task C8W9: Route Staged Smoke Through Materialized Service-Call Arguments

### Mission

Make the real staged launcher dispatch smoke consume the Go-owned `desktop-trigger-service-call-materialization-preview` packet before it invokes `xnix-runtime-owner`.

The important product distinction is this:

- the materialization preview may prepare evidence-only owner service call arguments for a human-authorized smoke candidate;
- the preview must not dispatch the owner service call itself;
- the preview must not claim that the formal full checkpoint is promoted unless the full checkpoint is actually promoted;
- the smoke harness may consume the materialized arguments only as part of an explicitly human-authorized smoke path.

This prevents the pipeline from deadlocking on "full checkpoint must be promoted before the full checkpoint can run" while still keeping formal release promotion honest.

### Expected Implementation Shape

Add or update only the relevant lane:

```text
internal/runtime/owner/desktop_trigger_service_call_materialization.go
internal/runtime/owner/desktop_trigger_service_call_materialization_test.go
cmd/xnix-runtime-go/desktop_trigger_service_call_materialization_commands.go
cmd/xnix-runtime-go/desktop_trigger_service_call_materialization_cli_test.go
scripts/staged_launcher_dispatch_smoke.rb
test/test_staged_launcher_dispatch_smoke_script.rb
scripts/verify_layout.rb
VERSION
CHANGELOG.md
PRODUCT_OVERVIEW.md
docs/xnix-current-mainline.md
```

Suggested behavior:

1. Add an explicit request field and CLI flag such as `HumanAuthorizedSmoke` / `--human-authorized-smoke`.
2. Keep `FullCheckpointPromoted` / `--full-checkpoint-promoted` as the only signal that may claim the formal full checkpoint has already been promoted.
3. Allow materialization when the dry-run review is blocked only by the missing full checkpoint and the new human-authorized smoke flag is set.
4. In that candidate path, emit `owner_service_call_args` and `owner_service_cli_args`, but keep:
   - `full_checkpoint_state=needs-full-checkpoint` or the existing equivalent blocked full-checkpoint state;
   - `human_authorization_required=true`;
   - `human_authorized_smoke=true`;
   - `full_checkpoint_promotion_claimed=false`;
   - `formal_release_ready=false` if such a field exists or is added;
   - `service_call_dispatch_enabled=false`;
   - `service_call_dispatched=false`;
   - `dbus_called=false`;
   - `runtime_state_written=false`;
   - `kde_configuration_written=false`;
   - `desktop_launch_enabled=false`;
   - `backend_launch_enabled=false`;
   - `execution_started=false`;
   - `host_root_modified=false`.
5. Preserve the existing strict blocked behavior when the new human-authorized smoke flag is absent.
6. Update `scripts/staged_launcher_dispatch_smoke.rb` so the second Runtime-status launch path consumes `desktop-trigger-service-call-materialization-preview` and passes the returned `owner_service_cli_args` into `xnix-runtime-owner`.
7. Keep the script from reconstructing owner-only arguments in Ruby.

### Acceptance Criteria

- The staged smoke script contains and consumes `desktop-trigger-service-call-materialization-preview`.
- The staged smoke script passes the explicit human-authorized smoke candidate flag to the preview.
- The staged smoke script invokes `xnix-runtime-owner` only with materialized `owner_service_cli_args`.
- The materialization preview still blocks service-call argument emission when neither `--full-checkpoint-promoted` nor `--human-authorized-smoke` is present.
- The human-authorized smoke candidate path does not claim formal release readiness or full checkpoint promotion.
- The preview output does not expose state roots, cache roots, launcher paths, raw backend commands, raw launcher output, backend details, secrets, or host paths.
- No Docker, QEMU, Wine, Colima, backend launch, D-Bus call, production state write, KDE configuration write, host-root mutation, broad host mount, host networking, privileged container, or Docker socket mount is introduced.

### Targeted Verification

Claude Code may run these commands:

```text
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/owner ./internal/runtime/appidentity ./cmd/xnix-runtime-go -run 'TestDesktopTriggerServiceCallMaterialization|TestDesktopTriggerDryRunRequestReview|TestKnownAppRuntimeStatusLaunchOwnerTrigger' -count=1
ruby -Ilib test/test_staged_launcher_dispatch_smoke_script.rb
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
ruby scripts/runtime_contract_drift_report.rb --format json
git diff --check
```

Claude Code must not run:

```text
ruby scripts/staged_launcher_dispatch_smoke.rb
ruby scripts/full_smoke.rb
ruby scripts/container.rb full-smoke
docker ...
qemu-system-...
wine ...
colima ...
```

### Copyable Claude Code Prompt

```text
Implement C8W9 from docs/claude-code-immediate-mainline-handoff.md.

Scope:
- Route scripts/staged_launcher_dispatch_smoke.rb through desktop-trigger-service-call-materialization-preview before invoking xnix-runtime-owner.
- Add an explicit human-authorized smoke candidate signal to the materialization request/CLI so the staged smoke can materialize evidence-only owner_service_cli_args without claiming the formal full checkpoint has already passed.
- Preserve the existing blocked behavior when neither --full-checkpoint-promoted nor the new human-authorized smoke candidate flag is present.
- Do not run Docker, QEMU, Wine, Proton, Colima, full_smoke, or the real staged smoke script.
- Do not modify docs/claude-code-implementation-packages.md.

Expected files:
- internal/runtime/owner/desktop_trigger_service_call_materialization.go
- internal/runtime/owner/desktop_trigger_service_call_materialization_test.go
- cmd/xnix-runtime-go/desktop_trigger_service_call_materialization_commands.go
- cmd/xnix-runtime-go/desktop_trigger_service_call_materialization_cli_test.go
- scripts/staged_launcher_dispatch_smoke.rb
- test/test_staged_launcher_dispatch_smoke_script.rb
- scripts/verify_layout.rb
- VERSION
- CHANGELOG.md
- PRODUCT_OVERVIEW.md
- docs/xnix-current-mainline.md

Run the targeted verification listed under C8W9 and report exact commands run and skipped.
```

## Task C8W10: Route the D-Bus Fixture Through the Same Materialization Packet

### Mission

After `C8W9` lands, make the private D-Bus controlled-launch owner fixture consume the same Go-owned service-call materialization packet instead of reaching directly for the lower-level owner trigger preview.

The goal is to keep one Runtime-owned service-call argument source for both:

- the real staged launcher dispatch smoke; and
- the private D-Bus controlled-launch fixture lane.

### Expected Implementation Shape

Likely files:

```text
scripts/dbus_controlled_launch_owner_fixture_smoke.rb
test/test_dbus_controlled_launch_owner_fixture_smoke_script.rb
runtime/dbus/xnix_compatd_smoke.c
test/test_runtime_dbus_smoke_script.rb
scripts/verify_layout.rb
VERSION
CHANGELOG.md
PRODUCT_OVERVIEW.md
docs/xnix-current-mainline.md
```

If the C adapter still needs a lower-level helper for a narrow smoke boundary, document why in `docs/xnix-current-mainline.md` and keep the adapter policy-free.

### Acceptance Criteria

- The D-Bus fixture verifies `desktop-trigger-service-call-materialization-preview`.
- It consumes evidence-only `owner_service_call_args` or `owner_service_cli_args` from Go-owned materialization.
- It still calls only `org.xnix.Compatibility1.ShowRuntimeControlledLaunch` with a safe evidence-relative-path handle.
- KDE and the C adapter do not reconstruct receipt ids, session ids, state roots, launcher paths, backend commands, or owner-only inputs.
- No production D-Bus ownership, real Portal transport, Docker, QEMU, Wine, backend launch, Runtime production write, KDE configuration write, or host mutation is enabled.

### Targeted Verification

Claude Code may run:

```text
ruby -Ilib test/test_dbus_controlled_launch_owner_fixture_smoke_script.rb
ruby -Ilib test/test_runtime_dbus_smoke_script.rb
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
ruby scripts/runtime_contract_drift_report.rb --format json
git diff --check
```

Claude Code must not run:

```text
ruby scripts/dbus_controlled_launch_owner_fixture_smoke.rb
ruby scripts/kde_controlled_launch_action_smoke.rb
ruby scripts/full_smoke.rb
docker ...
qemu-system-...
wine ...
colima ...
```

### Copyable Claude Code Prompt

```text
Implement C8W10 from docs/claude-code-immediate-mainline-handoff.md after C8W9 has landed.

Scope:
- Route the private D-Bus controlled-launch owner fixture through desktop-trigger-service-call-materialization-preview.
- Keep the C adapter and KDE-facing route policy-free and evidence-only.
- Do not run Docker, QEMU, Wine, Proton, Colima, full_smoke, or real fixture smoke execution.
- Do not modify docs/claude-code-implementation-packages.md.

Run the targeted verification listed under C8W10 and report exact commands run and skipped.
```

## Task C8W11: Add a Release-Promotion Evidence Packet

### Mission

Create a deterministic release-promotion packet for `v0.2.640` so Codex and the operator can decide whether to promote the release after the human-authorized full smoke passes.

This task must not run the full smoke. It should only read already-produced reports when they exist and clearly state what remains operator-owned.

### Expected Implementation Shape

Add a Go or Ruby read-only report command that can summarize:

- current repository version;
- expected formal release version;
- latest full-smoke report presence;
- full-smoke pass/fail/blocked state;
- structured failure class when the report exists;
- whether Docker/QEMU/Wine evidence was observed in an already-produced report;
- whether `formal_release_ready` can be claimed;
- exact operator command required when evidence is missing.

Likely files:

```text
scripts/full_checkpoint_promotion_packet.rb
test/test_full_checkpoint_promotion_packet.rb
scripts/verify_layout.rb
VERSION
CHANGELOG.md
PRODUCT_OVERVIEW.md
docs/xnix-current-mainline.md
```

Ruby is acceptable here because this is release evidence orchestration, not Runtime product behavior.

### Acceptance Criteria

- Missing full-smoke reports produce a clear blocked packet, not a false PASS.
- Failed full-smoke reports preserve the failure class and safe retry command from existing structured reports.
- Passing full-smoke reports are required before `formal_release_ready=true`.
- The packet does not run Docker, QEMU, Wine, Colima, full smoke, or any compatibility backend.
- The packet redacts host paths and does not expose secrets, tokens, usernames, or environment values.

### Targeted Verification

Claude Code may run:

```text
ruby -Ilib test/test_full_checkpoint_promotion_packet.rb
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
```

Claude Code must not run:

```text
ruby scripts/full_smoke.rb
docker ...
qemu-system-...
wine ...
colima ...
```

### Copyable Claude Code Prompt

```text
Implement C8W11 from docs/claude-code-immediate-mainline-handoff.md after C8W9 and C8W10 have landed or been explicitly skipped.

Scope:
- Add a read-only release-promotion evidence packet for v0.2.640.
- Do not run full_smoke, Docker, QEMU, Wine, Proton, Colima, or backend execution.
- Do not claim formal_release_ready unless an existing full-smoke PASS report is present and validated.
- Do not modify docs/claude-code-implementation-packages.md.

Run the targeted verification listed under C8W11 and report exact commands run and skipped.
```
