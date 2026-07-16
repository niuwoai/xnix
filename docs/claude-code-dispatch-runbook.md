# Claude Code Dispatch Runbook

> Last updated: 2026-07-16 | Baseline: v0.2.293

This runbook explains how to hand the KDE-first Windows compatibility work to Claude Code without losing mainline convergence.

Use this runbook after reading:

- `docs/mainline-integration-checkpoint.md`
- `docs/claude-code-mainline-task-batch.md`
- `docs/claude-code-windows-compatibility-workstreams.md`
- `docs/claude-code-third-wave-task-batch.md` when the first and second waves are merged or explicitly skipped
- `docs/claude-code-fourth-wave-task-batch.md` when Runtime evidence needs to become reviewable product surfaces
- `docs/claude-code-fifth-wave-task-batch.md` when reviewable evidence needs to feed KDE journeys, onboarding, supportability, and restricted acceptance fixtures
- `docs/claude-code-sixth-wave-task-batch.md` when the next work should focus on support timelines, deactivation dry-runs, recipe conflict audits, Portal renewal previews, restore ranking, settings migration, fixture matrices, and merge readiness packets
- `docs/claude-code-seventh-wave-task-batch.md` when the next work should focus on upgrade impact, policy explanation cards, state-root retention, crash and hang summaries, backend fallback, KDE search visibility, permission evidence audits, and release evidence indexing

Do not modify `docs/claude-code-implementation-packages.md` from this runbook. That file is Claude-owned.

## Dispatch Principle

Assign one Claude Code branch to one task. Require a small, mergeable result with targeted verification.

Do not assign broad prompts such as:

- "Implement Windows compatibility."
- "Finish the Runtime."
- "Make KDE integration work."
- "Enable Wine launch."

Those prompts are too wide and will create cross-lane diffs that are hard to review.

## Current First Dispatch

Use this order unless a branch is already in review:

| Dispatch | Task | Lane | Branch suggestion | Acceptance focus |
| --- | --- | --- | --- | --- |
| 1 | `CB1` Runtime owner private bus service hardening | `CW1 / A1` | `codex/cb1-runtime-owner-private-bus` | Private read owner evidence, disabled writes, blocked production owner |
| 2 | `CB2` Recipe and artifact receipt end-to-end gate | `CW2 / A2` | `codex/cb2-recipe-artifact-receipts` | Verified local recipes and artifacts, tampered receipt failures |
| 3 | `CB3` Backend manager lifecycle repair states | `CW3 / A3` | `codex/cb3-backend-lifecycle-state` | Durable backend inventory and lifecycle states without launch |
| 4 | `CB8` Mainline integration review gate | `CW10 / A8` | `codex/cb8-mainline-review-gate` | Review reports classify every changed file into a lane |

Do not dispatch `CB9` until `CB1`, `CB2`, `CB3`, `CB4`, and `CB8` are merged.

## Secondary Dispatch

Use these after the first dispatch is stable:

| Dispatch | Task | Lane | Branch suggestion | Acceptance focus |
| --- | --- | --- | --- | --- |
| 5 | `CB4` Portal permission review ledger | `CW5 / A4` | `codex/cb4-portal-permission-ledger` | Fake-mode permission records and denied/granted receipts |
| 6 | `CB5` Snapshot and rollback receipt store | `CW6 / A4` | `codex/cb5-snapshot-rollback-receipts` | Content-addressed snapshot receipts and blocked rollback states |
| 7 | `CB7` AI diagnostic privacy boundary | `CW7 / A7` | `codex/cb7-ai-diagnostic-privacy` | Redaction, fixture diagnostics, review-only repair recommendations |
| 8 | `CB6` KDE entry-point evidence consumers | `CW4 / A5` | `codex/cb6-kde-evidence-consumers` | KDE reads durable Runtime receipts and hides backend details |

## Third-Wave Dispatch

Use `docs/claude-code-third-wave-task-batch.md` after the current mainline and second-wave branches have been merged or deliberately skipped. Prefer dispatching `TW1`, `TW2`, `TW3`, and `TW5` first because they improve merge review, action prerequisites, permission-to-execution evidence, and AI privacy without requiring Docker, QEMU, host package managers, real Portal calls, backend launch, or host-root mutation.

## Fourth-Wave Dispatch

Use `docs/claude-code-fourth-wave-task-batch.md` after the current mainline, second-wave, and third-wave branches have been merged or deliberately skipped. Prefer dispatching `FW1`, `FW2`, `FW3`, and `FW5` first because they convert existing Runtime evidence into audit timelines, permission lifecycle summaries, rollback rehearsal plans, and settings dependency reviews without requiring Docker, QEMU, host package managers, real Portal calls, backend launch, or host-root mutation.

## Fifth-Wave Dispatch

Use `docs/claude-code-fifth-wave-task-batch.md` after the current mainline, second-wave, third-wave, and fourth-wave branches have been merged or deliberately skipped. Prefer dispatching `F5W1`, `F5W2`, `F5W3`, and `F5W8` first because they connect Runtime evidence to KDE user journeys, onboarding, redacted supportability, and restricted acceptance fixtures without requiring Docker, QEMU, host package managers, real Portal calls, backend launch, or host-root mutation.

## Sixth-Wave Dispatch

Use `docs/claude-code-sixth-wave-task-batch.md` after the current mainline, second-wave, third-wave, fourth-wave, and fifth-wave branches have been merged or deliberately skipped. Prefer dispatching `S6W1`, `S6W2`, `S6W3`, and `S6W8` first because they improve support timelines, desktop lifecycle closure, recipe conflict evidence, and merge readiness without requiring Docker, QEMU, host package managers, real Portal calls, backend launch, or host-root mutation.

## Seventh-Wave Dispatch

Use `docs/claude-code-seventh-wave-task-batch.md` after the current mainline, second-wave, third-wave, fourth-wave, fifth-wave, and sixth-wave branches have been merged or deliberately skipped. Prefer dispatching `S7W1`, `S7W2`, `S7W4`, and `S7W8` first because they improve upgrade safety, policy explanation quality, diagnostic signal summaries, and release evidence review without requiring Docker, QEMU, host package managers, real Portal calls, backend launch, or host-root mutation.

## Prompt Prefix

Prepend this prefix to the selected task prompt from `docs/claude-code-mainline-task-batch.md`:

```text
You are implementing one bounded Xnix task. Stay inside the selected task unless a required shared file is listed in the task.

Hard boundaries:
- Do not modify docs/claude-code-implementation-packages.md.
- Keep all source, comments, tests, CLI output, and docs in English.
- Prefer Go for durable Runtime product logic.
- Use Ruby only for tests, reports, smoke scripts, and lightweight tooling.
- Do not enable production D-Bus ownership, Runtime writes, real Wine/Proton/VM launch, real Portal calls, network fetch, host package-manager calls, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Do not expose Wine prefix terminology, raw executable paths, backend commands, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials in KDE-facing output.

Before finishing:
- Run the task-specific verification commands.
- Run ruby scripts/verify_layout.rb.
- Report exact commands run and any commands not run.
```

## Intake Checklist for Claude Results

When Claude returns a branch, check:

- The branch changed one task lane plus clearly justified shared plumbing.
- `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` are updated for code changes.
- `docs/claude-code-implementation-packages.md` has no diff.
- `.gocache/`, `tmp/`, build artifacts, local logs, private config, and secrets are not staged.
- The branch adds success, failure, and blocked-state tests.
- KDE-facing output hides backend details and host paths.
- Unsafe behavior remains disabled.
- The task-specific verification commands were run.
- `ruby scripts/verify_layout.rb` passed.

## Local Review Commands

Use these commands before staging a Claude branch:

```text
ruby scripts/mainline_integration_review.rb --format json
ruby scripts/mainline_integration_review.rb --format markdown
ruby scripts/verify_layout.rb
git diff --check
git diff -- docs/claude-code-implementation-packages.md
```

The JSON report includes a `review_matrix` for every changed lane. Use that matrix as the lane-specific checklist for required verification commands, required evidence, and safety guards before staging.

If Runtime owner routes, D-Bus XML, D-Bus clients, smoke adapters, or Runtime CLI route commands changed, also run:

```text
ruby scripts/runtime_contract_drift_report.rb --format json
```

If implementation evidence gates changed, also run:

```text
ruby scripts/implementation_evidence_report.rb --format json
```

## Staging Rule

Never run `git add .` for Claude output.

Stage by lane using the review report:

1. Run `ruby scripts/mainline_integration_review.rb --format json`.
2. Confirm `protected_claude_file_modified` is `false`.
3. Confirm `unclassified_file_count` is `0`.
4. Confirm the lane you are staging has the expected files.
5. Confirm the lane's `review_matrix` commands and safety guards have been satisfied or explicitly reported as not run.
6. Exclude `.gocache/`, `tmp/`, build artifacts, logs, and local config.
7. Stage only the lane files plus required version and documentation files.

## Completion Format Required from Claude

Claude should return:

```text
Task:
Branch:
Version:
Lane:
Files changed:
Shared files changed and why:
What moved from contract-only to implementation evidence:
Safety gates kept disabled:
Verification commands run:
Commands not run and why:
Known follow-up:
```

## Stop Conditions

Stop review and ask for human direction if a branch:

- Modifies `docs/claude-code-implementation-packages.md`.
- Requires Docker or QEMU without explicit authorization.
- Starts Wine, Proton, a VM, or any compatibility backend.
- Calls real XDG Portal transport.
- Fetches network artifacts by default.
- Invokes the host package manager.
- Requires privileged containers, host networking, Docker socket mounts, or broad host mounts.
- Mutates the host root.
- Exposes backend commands, raw executable paths, state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials in user-facing output.
