# Xnix Mainline Integration Checkpoint

> Last updated: 2026-07-16 | Baseline under review: v0.2.293

This checkpoint freezes new feature expansion long enough to merge Claude Code output, align it with the current product goal, and keep the mainline convergent.

## Product Goal

Xnix is moving toward the best Linux desktop for existing Windows applications.

The first official desktop shell is KDE Plasma. KDE is a replaceable shell and presentation layer, not the product core. The product core is the independent Xnix AI Compatibility Runtime.

The Runtime owns:

- Wine, Proton, Windows VM, and future backend inventory.
- Application recipes and artifact trust.
- Permissions, Portal requests, snapshots, rollback, diagnostics, and execution ledgers.
- D-Bus read ownership, write gates, state roots, and safety policy.

KDE owns only desktop presentation and interaction:

- Start menu entries.
- Task-manager identity.
- Dolphin and file association actions.
- Tray and notification status.
- Compatibility Center views.
- Unified settings surfaces.

## Integration Rule

Do not modify `docs/claude-code-implementation-packages.md` from this checkpoint. Treat that file as Claude-owned.

Claude-generated work may be merged when it supports the product goal and remains inside one reviewable lane. Do not hide unrelated cache, build, or temporary output inside a product commit.

Always exclude:

- `.gocache/`
- `tmp/`
- build artifacts
- local logs
- private configuration
- secrets, tokens, private keys, or credentials

## Mainline Lanes

Use these lanes to review the current mixed worktree and future Claude Code output.

| Lane | Purpose | Candidate files | Required evidence |
| --- | --- | --- | --- |
| `CW1 / A1` Runtime owner read boundary | Make the Go Runtime owner a real constrained read boundary while writes fail closed. | `cmd/xnix-runtime-owner/`, `internal/runtime/owner/`, owner route and drift scripts | Go owner tests, private session-bus smoke, drift report, disabled-write evidence |
| `CW2 / A2` Recipe and artifact trust | Verify recipes and staged artifacts before install, activation, or execution depends on them. | `internal/runtime/recipe/`, `internal/runtime/artifact/`, install-plan files | Digest mismatch tests, staging receipt tests, install-readiness blocked-state tests |
| `CW3 / A3` State root and backend lifecycle | Persist Runtime-owned backend state without launching Wine, Proton, or VM backends. | `internal/runtime/appidentity/backend_*`, environment lifecycle files | State-root record tests, hidden-path assertions, no backend-launch assertions |
| `CW5 / A4` Portal and permission safety | Record fake-mode permission requests and review states before real Portal transport. | `internal/runtime/portal/`, permission and execution preflight files | Granted, denied, expired, and blocked-state receipt tests |
| `CW8 / A6` Execution ledger and session evidence | Record reviewed execution transactions while launch remains disabled. | `internal/runtime/execution/`, KDE session-evidence consumers | Ledger tests, session-record tests, no launch or backend-start assertions |
| `CW10 / A8` Evidence and drift harness | Prevent more contract-only growth by requiring implementation evidence. | `scripts/implementation_evidence_report.rb`, `scripts/runtime_contract_drift_report.rb`, `scripts/verify_layout.rb` | Report tests, layout verification, contract-drift verification |
| `CW4 / A5` KDE entry points | Make KDE consume Runtime evidence through safe read models. | KDE shell, Compatibility Center, task-manager, KWin, tray, Dolphin, settings files | KDE-safe output tests, no backend detail or host path exposure |
| `CW11 / A9` Product image acceptance | Validate the product through restricted Docker or QEMU only after Runtime evidence is meaningful. | Buildroot, boot, QEMU, container scripts | Loopback-only SSH, persisted logs, no privileged container or host mutation |

## Current Merge Posture

The current worktree contains both planning documents and implementation changes. Review them in this order:

1. Planning and dispatch documents.
2. `CW1 / A1` Runtime owner boundary.
3. `CW2 / A2` trust and artifact pipeline.
4. `CW3 / A3` backend state root and lifecycle.
5. `CW5 / A4` Portal safety records.
6. `CW8 / A6` execution ledger and session evidence.
7. `CW10 / A8` evidence and drift gates.
8. KDE consumers that only read durable Runtime evidence.

Do not merge `CW11 / A9` product-image work until the restricted Docker or QEMU smoke can validate real Runtime evidence instead of static contracts.

## Safety Invariants

These behaviors remain disabled until a later explicitly scoped acceptance package enables a constrained test-only path:

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
- Host-root mutation.
- Raw backend commands in KDE-facing output.
- Raw executable paths, Runtime state-root paths, host paths, file contents, secrets, tokens, private keys, or credentials in user-facing output.

## Review Gate

A lane is mergeable only when it has all of the following:

1. A narrow file set that matches the lane.
2. Version, changelog, and product-overview updates when code changes.
3. Success, failure, and blocked-state tests.
4. Evidence that user-facing output hides backend details and host paths.
5. Evidence that unsafe behavior remains disabled.
6. Targeted Go and Ruby verification output.
7. `ruby scripts/verify_layout.rb` passing.
8. `git diff --check` passing.

If D-Bus XML, owner dispatch, Runtime route manifests, smoke adapters, or D-Bus clients change, also run:

```text
ruby scripts/runtime_contract_drift_report.rb --format json
```

If implementation evidence gates change, also run:

```text
ruby scripts/implementation_evidence_report.rb --format json
```

`scripts/mainline_integration_review.rb --format json` emits a `review_matrix` for every changed lane. Treat that matrix as the authoritative local checklist for lane-specific verification commands, required evidence, and safety guards before staging files from a Claude or Codex branch.

## Suggested Checkpoint Commit Shape

Prefer one reviewed integration checkpoint instead of silently mixing unrelated work:

```text
feat(runtime): integrate Windows compatibility foundation

- Align KDE-first Windows compatibility workstreams with the Runtime-owned product core.
- Add constrained Runtime owner, trust, state-root, Portal, execution, and evidence lanes.
- Keep production D-Bus ownership, Runtime writes, real backend launch, network fetch, privileged containers, and host-root mutation disabled.
```

If the worktree proves too broad for one review, split it by the lane order above.
