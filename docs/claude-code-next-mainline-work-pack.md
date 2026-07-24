# Claude Code Next Mainline Work Pack

> Last updated: 2026-07-24 | Baseline: v0.2.640-rc19 | Formal release: v0.2.640 remains blocked until full smoke passes

This document is a copy-first task pack for asking Claude Code to implement the next Xnix mainline work while Codex keeps review, merge, release-promotion, and host-safety decisions.

Use this file when the operator wants Claude Code to pick one bounded task, implement it in a separate branch, run only targeted verification, and return the result for Codex review.

Do not modify `docs/claude-code-implementation-packages.md` from any task in this pack.

## Current Mainline State

The current local baseline is `v0.2.640-rc19`.

For a short current handoff that Claude Code can execute without reading the full historical task pack first, use `docs/claude-code-active-work-order.md`.

Already present in the current checkpoint candidate:

- `desktop-trigger-staged-invocation-readiness-preview`.
- `owner-service-launch-envelope-guard-preview`.
- `kde-controlled-launch-action-surface-audit-preview`.
- `managed-launcher-acceptance-report-preview`.
- `desktop-trigger-dry-run-request-review-preview`.
- `desktop-trigger-service-call-materialization-preview`.
- The staged launcher dispatch smoke consumes materialized owner service call arguments with an explicit human-authorized smoke candidate flag.
- The private D-Bus controlled-launch fixture consumes the same materialization packet.
- Structured full-smoke failure reports exist when full smoke fails.
- `scripts/full_checkpoint_promotion_packet.rb` reads existing full-smoke reports and produces the release-promotion decision without running full smoke.
- `scripts/merge_readiness_packet.rb` consumes the promotion packet and exposes `full_checkpoint_promotion_status`.
- `scripts/release_evidence_index.rb` consumes the promotion packet and keeps historical product smoke evidence separate from current formal release readiness.
- `docs/post-checkpoint-promotion-checklist-0640.md` defines the human-owned promotion checklist.
- `desktop-trigger-request-preflight-preview` is present as a Go-owned post-release request preflight that remains fail-closed with `blocked-missing-promotion` until formal full checkpoint promotion is observed.
- `scripts/desktop_trigger_request_preflight_smoke.rb` is present as a targeted smoke for the request preflight lane; it verifies blocked and promoted fixture states without running Docker, QEMU, Wine, D-Bus, KDE, or a backend.
- `scripts/merge_readiness_packet.rb` consumes existing desktop-trigger request preflight smoke JSON evidence when `--desktop-trigger-request-preflight-smoke` is supplied, without running that smoke by default.
- `scripts/release_evidence_index.rb` classifies existing desktop-trigger request preflight smoke JSON evidence when `--desktop-trigger-request-preflight-smoke` is supplied, without running that smoke by default.
- `docs/claude-code-next-dispatch-brief.md` is present as the current short copy-first dispatch brief.

The formal `v0.2.640` release must not be promoted until this operator-owned command passes:

```text
ruby scripts/full_smoke.rb
```

Claude Code must not run that command unless the human operator gives explicit task-local approval.

## Global Rules for Claude Code

Before starting any task, run:

```text
git status --short --branch
```

If the checkout is dirty, stop and ask the operator whether to continue the dirty lane or wait for Codex review. Do not start unrelated work on top of unknown local changes.

Every task must:

- Implement exactly one selected task.
- Keep all source, comments, fixtures, tests, CLI output, and documentation in English.
- Prefer Go for durable Runtime product behavior and stable service boundaries.
- Use Ruby only for tests, smoke orchestration, reports, fixtures, and lightweight developer tooling.
- Keep KDE as a replaceable presentation shell.
- Keep the Xnix Compatibility Runtime as the policy owner.
- Update `VERSION`, `CHANGELOG.md`, `PRODUCT_OVERVIEW.md`, `docs/xnix-current-mainline.md`, layout verification, and mainline review gates when code changes.
- Add targeted tests for ready, blocked, malformed, redaction, and no-side-effect behavior where relevant.
- Run task-specific verification plus `ruby scripts/verify_layout.rb` and `git diff --check`.
- Report exact commands run and exact commands skipped.

Every task must not:

- Run Docker, QEMU, Wine, Proton, Colima, or a real compatibility backend.
- Run `ruby scripts/full_smoke.rb`.
- Pull images, fetch packages, restart host services, use host package managers, or repair host networking.
- Use privileged containers, host networking, Docker socket mounts, broad host-directory mounts, or host-root mutation.
- Enable production D-Bus ownership.
- Add real XDG Portal transport calls.
- Write KDE configuration.
- Move Runtime policy into KDE, Ruby smoke scripts, or the C smoke adapter.
- Expose state-root paths, cache-root paths, launcher paths, executable paths, backend names, backend commands, raw launcher output, host paths, secrets, tokens, credentials, usernames, or environment variables in user-facing output.

## Dispatch Order

Use this order unless Codex or a reviewer asks for a repair branch:

| Order | Task | Suggested branch | When to dispatch |
| --- | --- | --- | --- |
| Done | `C8W11` Full checkpoint promotion packet | `codex/full-checkpoint-promotion-packet` | Present locally in v0.2.640-rc11. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W12` Merge readiness consumes promotion packet | `codex/merge-readiness-promotion-gate` | Present locally in v0.2.640-rc12. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W13` Release evidence index aligns with promotion packet | `codex/release-evidence-promotion-claim` | Present locally in v0.2.640-rc13. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C8W14` Operator promotion checklist refresh | `codex/operator-promotion-checklist-0640` | Present locally in v0.2.640-rc14. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C9W1` Post-release desktop-trigger request preflight | `codex/desktop-trigger-request-preflight` | Present locally in v0.2.640-rc15 as a read-only fail-closed preflight. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C9W2` Desktop-trigger request preflight smoke | `codex/desktop-trigger-request-preflight-smoke` | Present locally in v0.2.640-rc16 as a targeted smoke. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C9W3` Merge readiness consumes request preflight smoke evidence | `codex/merge-readiness-preflight-smoke-evidence` | Present locally in v0.2.640-rc17 as a read-only optional evidence input. Do not dispatch again unless a reviewer asks for repair. |
| Done | `C9W4` Release evidence index consumes request preflight smoke evidence | `codex/release-evidence-preflight-smoke-evidence` | Present locally in v0.2.640-rc18 as a read-only optional evidence claim. Do not dispatch again unless a reviewer asks for repair. |
| 1 | `C9W5` Runtime owner request receipt preview | `codex/runtime-owner-request-receipt-preview` | Dispatch after formal `v0.2.640` promotion or explicit Codex approval to continue pre-release modeling. |

Stop after each task. Return the branch or diff for Codex review. Do not chain tasks.

## Task C8W11: Full Checkpoint Promotion Packet

### Mission

Create a deterministic read-only promotion packet for deciding whether `v0.2.640` can be promoted after the human-authorized full smoke has already produced report files.

The packet must never run full smoke. It only reads existing evidence.

### Expected files

```text
scripts/full_checkpoint_promotion_packet.rb
test/test_full_checkpoint_promotion_packet.rb
scripts/verify_layout.rb
scripts/mainline_integration_review.rb
VERSION
CHANGELOG.md
PRODUCT_OVERVIEW.md
docs/xnix-current-mainline.md
```

Ruby is acceptable because this is release evidence orchestration, not Runtime product behavior.

### Required behavior

- Read `VERSION`.
- Read `output/full-smoke-report.json` when present.
- Optionally read `output/full-smoke-report.md` only as presence evidence.
- Emit JSON and Markdown formats.
- Report current version, expected formal release, release-candidate shape, report presence, full-smoke state, failure class, safe retry command, observed Docker/QEMU/Wine evidence, and promotion decision.
- Set `formal_release_ready=true` only when a valid existing full-smoke report says the full smoke passed and the repository version is promotable to the expected formal release.
- Keep explicit flags proving the packet did not run full smoke, Docker, QEMU, Wine, Colima, D-Bus, desktop launch, backend launch, network access, privileged containers, or host mutation.
- Redact host paths, secrets, tokens, usernames, and environment values.

### Required tests

- Missing report produces a blocked packet.
- Malformed report produces a blocked packet.
- Failed report preserves `failure_class`, `failed_step`, `operator_action_required`, `project_defect_possible`, and `safe_retry_command`.
- Passing report allows promotion only for the matching release candidate line.
- Version mismatch blocks promotion.
- Markdown output includes the promotion decision and skipped unsafe operations.
- Output does not expose host paths or sensitive values.

### Verification

Claude Code may run:

```text
ruby -Ilib test/test_full_checkpoint_promotion_packet.rb
ruby scripts/full_checkpoint_promotion_packet.rb --format json
ruby scripts/full_checkpoint_promotion_packet.rb --format markdown
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

Claude Code must not run:

```text
ruby scripts/full_smoke.rb
ruby scripts/container.rb full-smoke
docker ...
qemu-system-...
wine ...
colima ...
```

### Copyable Claude Code prompt

```text
Implement C8W11 from docs/claude-code-next-mainline-work-pack.md.

Scope:
- Add a read-only full checkpoint promotion packet for v0.2.640.
- The packet may read existing output/full-smoke-report.json and output/full-smoke-report.md evidence.
- The packet must not run full_smoke, Docker, QEMU, Wine, Proton, Colima, D-Bus, desktop launch, backend execution, network access, or host mutation.
- Do not claim formal_release_ready unless an existing valid full-smoke PASS report is present and the current version belongs to the expected v0.2.640 release-candidate line.
- Do not modify docs/claude-code-implementation-packages.md.

Expected files:
- scripts/full_checkpoint_promotion_packet.rb
- test/test_full_checkpoint_promotion_packet.rb
- scripts/verify_layout.rb
- scripts/mainline_integration_review.rb
- VERSION
- CHANGELOG.md
- PRODUCT_OVERVIEW.md
- docs/xnix-current-mainline.md

Run the C8W11 verification commands from the work pack and report exact commands run and skipped.
```

## Task C8W12: Merge Readiness Consumes Promotion Packet

### Mission

After `C8W11` lands, make `scripts/merge_readiness_packet.rb` consume the new full checkpoint promotion packet as a first-class release gate.

This task must not duplicate promotion logic. It must call or read the promotion packet result and surface the decision in the merge readiness output.

### Expected files

```text
scripts/merge_readiness_packet.rb
test/test_merge_readiness_packet.rb
scripts/verify_layout.rb
scripts/mainline_integration_review.rb
VERSION
CHANGELOG.md
PRODUCT_OVERVIEW.md
docs/xnix-current-mainline.md
```

### Required behavior

- Add a `full_checkpoint_promotion` tool section.
- In offline mode, accept a fixture path such as `--full-checkpoint-promotion PATH`.
- Include promotion blockers in `release_blocking_reasons`.
- Keep `automatic_release_tagging_enabled=false`.
- Keep staging, committing, tagging, pushing, Docker, QEMU, Wine, Colima, backend launch, network access, and host mutation disabled.
- Preserve existing release evidence index behavior.

### Verification

```text
ruby -Ilib test/test_merge_readiness_packet.rb
ruby scripts/merge_readiness_packet.rb --format json --offline-only
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable Claude Code prompt

```text
Implement C8W12 from docs/claude-code-next-mainline-work-pack.md after C8W11 has landed.

Scope:
- Make scripts/merge_readiness_packet.rb consume the full checkpoint promotion packet as a release gate.
- Do not reimplement promotion logic inside merge_readiness_packet.
- Do not run full_smoke, Docker, QEMU, Wine, Proton, Colima, D-Bus, desktop launch, backend execution, network access, or host mutation.
- Do not modify docs/claude-code-implementation-packages.md.

Run the C8W12 verification commands and report exact commands run and skipped.
```

## Task C8W13: Release Evidence Index Aligns With Promotion Packet

### Mission

After `C8W11` lands, align `scripts/release_evidence_index.rb` with the promotion packet so release-critical claims distinguish:

- implemented evidence;
- fixture-only evidence;
- human-authorized full-smoke evidence;
- blocked promotion;
- formal release readiness.

### Expected files

```text
scripts/release_evidence_index.rb
test/test_release_evidence_index.rb
scripts/verify_layout.rb
scripts/mainline_integration_review.rb
VERSION
CHANGELOG.md
PRODUCT_OVERVIEW.md
docs/xnix-current-mainline.md
```

### Required behavior

- Accept an optional promotion packet fixture.
- Add or update the `product-image-qemu-acceptance` claim so it can reference the full checkpoint promotion decision.
- Never claim formal readiness from targeted tests, fixture-only evidence, or missing full-smoke reports.
- Keep automatic release tagging disabled.
- Preserve existing claim ids unless a migration is necessary and tested.

### Verification

```text
ruby -Ilib test/test_release_evidence_index.rb
ruby scripts/release_evidence_index.rb --format json
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable Claude Code prompt

```text
Implement C8W13 from docs/claude-code-next-mainline-work-pack.md after C8W11 has landed.

Scope:
- Align scripts/release_evidence_index.rb with the full checkpoint promotion packet.
- Keep formal release readiness false unless a valid promotion packet proves an existing full-smoke PASS report.
- Do not run full_smoke, Docker, QEMU, Wine, Proton, Colima, D-Bus, desktop launch, backend execution, network access, or host mutation.
- Do not modify docs/claude-code-implementation-packages.md.

Run the C8W13 verification commands and report exact commands run and skipped.
```

## Task C8W14: Operator Promotion Checklist Refresh

### Mission

Create or refresh a human-operator checklist for promoting `v0.2.640` after the full checkpoint promotion packet says promotion is allowed.

This is documentation and release-tooling alignment. It must not run full smoke or create tags.

### Expected files

```text
docs/post-checkpoint-promotion-checklist-0640.md
docs/xnix-current-mainline.md
PRODUCT_OVERVIEW.md
```

Only update code if a tiny read-only validation hook is necessary.

### Required content

- Required PASS evidence for full smoke, restricted container safety, QEMU serial log, `sshd`, known Windows app smoke, fixture Windows app smoke, KDE controlled-launch action smoke, D-Bus fixture smoke, layout verification, mainline review, release evidence index, merge readiness packet, and diff hygiene.
- Exact human-owned promotion steps from the active release candidate to `0.2.640`.
- Explicit statement that promotion is forbidden until `ruby scripts/full_smoke.rb` passes.
- Explicit statement that Claude Code must not tag, push, run full smoke, or repair the host unless instructed.
- Rollback note for failed promotion review.

### Verification

```text
ruby scripts/verify_layout.rb
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable Claude Code prompt

```text
Implement C8W14 from docs/claude-code-next-mainline-work-pack.md.

Scope:
- Create or refresh docs/post-checkpoint-promotion-checklist-0640.md as the human promotion checklist for v0.2.640.
- Do not run full_smoke, Docker, QEMU, Wine, Proton, Colima, D-Bus, desktop launch, backend execution, network access, tagging, pushing, or host mutation.
- Do not claim the release is ready unless a valid existing promotion packet proves full-smoke PASS evidence.
- Do not modify docs/claude-code-implementation-packages.md.

Run the C8W14 verification commands and report exact commands run and skipped.
```

## Task C9W1: Post-Release Desktop-Trigger Request Preflight

Status: present locally in `v0.2.640-rc15`. Do not dispatch again unless Codex or a reviewer asks for a repair branch.

### Mission

Add the next read-only preflight for a future real desktop-triggered staged `ShowRuntimeControlledLaunch` request. In `v0.2.640-rc15`, this is implemented fail-closed before formal promotion: it returns `blocked-missing-promotion` unless promoted evidence is explicitly supplied.

This task moves beyond fixture validation, but it still must not execute the request.

### Expected files

```text
internal/runtime/owner/desktop_trigger_request_preflight.go
internal/runtime/owner/desktop_trigger_request_preflight_test.go
cmd/xnix-runtime-go/desktop_trigger_request_preflight_commands.go
cmd/xnix-runtime-go/desktop_trigger_request_preflight_cli_test.go
scripts/verify_layout.rb
scripts/mainline_integration_review.rb
VERSION
CHANGELOG.md
PRODUCT_OVERVIEW.md
docs/xnix-current-mainline.md
```

### Required behavior

- Consume the promoted full checkpoint state, materialized service call shape, envelope guard, action surface audit, and managed launcher acceptance report.
- Return deterministic states such as `ready-for-operator-request`, `blocked-missing-promotion`, `blocked-unsafe-envelope`, `blocked-unsafe-action`, `blocked-stale-evidence`, `malformed`, and `unsupported`.
- Emit only a safe request preflight summary and opaque evidence handles.
- Keep request creation, permission grants, service dispatch, D-Bus calls, desktop launch, backend launch, Runtime writes, KDE configuration writes, network access, and host mutation disabled.

### Verification

```text
GOCACHE=/Users/rocky/Sites/xnix/.cache/go-build go test ./internal/runtime/owner ./internal/runtime/appidentity ./cmd/xnix-runtime-go -run 'TestPreviewDesktopTriggerRequestPreflight|TestDesktopTriggerRequestPreflight|TestPreviewDesktopTriggerServiceCallMaterialization|TestDesktopTriggerServiceCallMaterialization|TestPreviewLaunchEnvelopeGuard|TestKDEControlledLaunchActionSurfaceAudit|TestPreviewManagedLauncherAcceptance' -count=1
ruby scripts/verify_layout.rb
ruby scripts/mainline_integration_review.rb --format json
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

### Copyable Claude Code prompt

```text
Implement C9W1 from docs/claude-code-next-mainline-work-pack.md only after formal v0.2.640 promotion has landed or Codex explicitly skips the promotion dependency.

Scope:
- Add a Go-owned read-only desktop-trigger request preflight for a future real staged ShowRuntimeControlledLaunch request.
- Do not create request objects, grant permissions, dispatch service calls, call D-Bus, launch desktop actions, start compatibility backends, run Docker, run QEMU, run Wine, use Colima, use network access, or mutate host state.
- Do not expose state roots, cache roots, launcher paths, executable paths, backend names, backend commands, raw launcher output, host paths, secrets, tokens, credentials, usernames, or environment variables.
- Do not modify docs/claude-code-implementation-packages.md.

Run the C9W1 verification commands and report exact commands run and skipped.
```

## Task C9W2: Human-Authorized Desktop-Trigger Staged Request Smoke

### Mission

After formal `v0.2.640` is promoted or Codex explicitly skips that dependency, wire a bounded smoke path that asks the operator for approval before using the preflight-approved desktop-trigger request lane.

This task must still keep production D-Bus ownership, real desktop launch, broad host mounts, host networking, Docker socket mounts, and host-root mutation disabled unless Codex explicitly expands the scope.

### Expected starting evidence

```text
desktop-trigger-request-preflight-preview
desktop-trigger-service-call-materialization-preview
scripts/staged_launcher_dispatch_smoke.rb
scripts/dbus_controlled_launch_owner_fixture_smoke.rb
```

### Required behavior

- Require `desktop-trigger-request-preflight-preview` to return `ready-for-operator-request` before any staged request smoke proceeds.
- Keep KDE-facing inputs limited to an opaque evidence id or safe evidence-relative-path.
- Keep Runtime owner-only inputs inside the Go Runtime owner boundary.
- Report exact commands run and skipped.
