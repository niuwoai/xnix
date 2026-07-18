# Claude Code Current Dispatch Picks

> Last updated: 2026-07-19 | Baseline: v0.2.341

This is the current short dispatch sheet for Claude Code. It is intentionally smaller than the full wave documents and should be used when choosing the next independent branch.

Use `docs/claude-code-stability-release-train.md` as the cadence authority: one coherent small version per local commit, targeted tests on every small version, and, beginning with `v0.2.321`, a full build plus complete tests and remote push every thirty small versions, with the next gate at `v0.2.350`.

Local work already contains implementation evidence for all seventh-wave tasks (`S7W1` through `S7W8`) plus `S6W1`, `S6W2`, `S6W3`, `S6W4`, `S6W5`, `S6W6`, `S6W7`, `S6W8`, `F5W7`, the `v0.2.321` offline fixture artifact-receipt read model, the `v0.2.322` offline fixture snapshot-baseline read model, the `v0.2.323` production Runtime service activation preflight read model, the `v0.2.324` test-only launch materialization plan receipt, the `v0.2.325` backend adapter no-op contract preview, the `v0.2.326` restricted owner smoke execution receipt, the `v0.2.327` test-only materialization fan-out to KDE surfaces, the `v0.2.328` backend adapter fixture binding audit, the `v0.2.329` restricted owner smoke receipt fan-out to readiness and support surfaces, the `v0.2.330` materialization fan-out owner-route audit, the `v0.2.331` adapter contract owner-route audit, the `v0.2.332` restricted owner smoke receipt fan-out owner-route audit, the `v0.2.333` materialization fan-out receipt-consumption split, the `v0.2.334` backend adapter redacted profile owner-route split, the `v0.2.335` restricted owner smoke receipt opaque lookup split, the `v0.2.336` materialization fan-out opaque receipt route audit, the `v0.2.337` restricted owner smoke fan-out owner-local read route, the `v0.2.338` materialization fan-out owner-managed opaque receipt lookup, the `v0.2.339` materialization fan-out owner-local read route, the `v0.2.340` materialization fan-out owner smoke coverage read model, and the `v0.2.341` restricted owner smoke fan-out owner smoke coverage read model. Do not dispatch those again unless their local branches are discarded or a reviewer explicitly asks for a repair branch.

Do not modify `docs/claude-code-implementation-packages.md` from any task listed here.

## Good Claude Code Fit

A task is suitable for Claude Code now if it is:

- Bounded to one Runtime or KDE-facing product lane.
- Mostly Go Runtime read-model work, with Ruby only for tests, reports, or smoke scripts.
- Reviewable offline with local fixtures.
- Useful to the KDE-first Windows application compatibility mainline.
- Safe without Docker, QEMU, network fetch, package-manager calls, real Portal transport, compatibility backend launch, privileged containers, broad host mounts, or host-root mutation.
- Mergeable one branch at a time through `scripts/mainline_integration_review.rb`.

## Recommended Dispatch Order

| Order | Task | Source document | Suggested branch | Why it is suitable now |
| --- | --- | --- | --- | --- |
| 1 | Redacted adapter profile restricted owner smoke coverage | `docs/windows-app-compatibility-implementation-brief.md` | `codex/redacted-adapter-profile-owner-smoke-coverage` | Cover the new owner-local redacted adapter profile audit route in restricted owner smoke evidence before considering broader exposure. |
| 2 | Materialization fan-out smoke coverage audit refresh | `docs/windows-app-compatibility-implementation-brief.md` | `codex/materialization-fanout-smoke-coverage-audit-refresh` | Refresh the materialization owner-route audit so it recognizes v0.2.340 smoke coverage while still blocking production D-Bus exposure. |
| 3 | Restricted owner smoke fan-out smoke coverage audit refresh | `docs/windows-app-compatibility-implementation-brief.md` | `codex/restricted-smoke-fanout-smoke-coverage-audit-refresh` | Refresh the restricted owner smoke fan-out owner-route audit so it recognizes v0.2.341 smoke coverage while still blocking production D-Bus exposure. |

## Best First Pick

Start with the redacted adapter profile restricted owner smoke coverage.

The v0.2.320 offline and constrained-host gates pass, including the full Buildroot build and QEMU serial smoke. Authorized q3 and q4 composes produced the Fedora 44 container, clean qcow2, persisted KVM graphical-login evidence, and closed host-boundary proof. The v0.2.321 and v0.2.322 matrix evidence can now consume controlled artifact and snapshot receipts, v0.2.323 turns production Runtime service activation into a fail-closed preflight, v0.2.324 materializes only a test-only review plan from restricted launch evidence, v0.2.325 defines the no-op backend adapter boundary before real Wine/Proton/VM launch paths are allowed, v0.2.326 records restricted owner smoke evidence without claiming production ownership, v0.2.327 fans the materialized review plan out to KDE-facing surfaces without enabling launch, v0.2.328 lets fixture-backed readiness rows consume no-op adapter contract profiles without invoking a backend, v0.2.329 lets readiness and support surfaces consume the restricted owner smoke receipt without claiming production ownership, v0.2.330 audits that the materialization fan-out should remain CLI-only until test receipt creation is split from read-only receipt consumption, v0.2.331 audits that the adapter contract should remain fixture-local until a redacted owner-local adapter profile audit route exists, v0.2.332 audits that the restricted owner smoke receipt fan-out should remain CLI-only until owner-managed opaque receipt lookup exists, v0.2.333 splits materialization receipt creation from read-only fan-out consumption, v0.2.334 creates the redacted adapter profile owner route while keeping the full adapter contract fixture-local, v0.2.335 introduces owner-managed opaque lookup for the restricted owner smoke receipt, v0.2.336 re-audits materialization fan-out to prove the read-only consume split exists while owner-managed opaque materialization receipt lookup is still missing, v0.2.337 routes restricted owner smoke fan-out through owner-managed opaque receipt lookup without caller state-root paths, v0.2.338 adds the matching owner-managed opaque lookup for materialization receipts, v0.2.339 lets materialization fan-out consume that owner lookup through an owner-local read route while staying non-launching and write-free, v0.2.340 proves that route appears in restricted owner smoke evidence through `Service.Call`, and v0.2.341 gives the restricted owner smoke fan-out route the same explicit smoke coverage. The next useful blocker is adding restricted owner smoke coverage for the redacted adapter profile route before any broader exposure. Windows application execution remains unproven.

The next branch must stay non-launching and must not claim the production bus, enable write methods, write KDE configuration, start a backend, call a real Portal, run QEMU, or mutate the host root.

## Good Branch Shape

Each Claude branch should:

- Implement exactly one listed task.
- Add a Go read model and a CLI preview command when the source task asks for one.
- Add success, malformed-input, blocked-state, redaction, and no-side-effect tests.
- Update `VERSION`, `CHANGELOG.md`, and `PRODUCT_OVERVIEW.md` for code changes.
- Wire new files into `scripts/verify_layout.rb`.
- Wire new evidence into `scripts/implementation_evidence_report.rb` and `scripts/mainline_integration_review.rb` when the task creates a new product surface.
- Report exact verification commands run and exact verification commands skipped.

## Do Not Dispatch Now

Avoid these prompts:

- "Finish Windows compatibility."
- "Make KDE integration real."
- "Enable Wine launch."
- "Install Proton or Wine."
- "Run QEMU acceptance."
- "Wire production D-Bus."
- "Call the real Portal."
- "Add real cleanup."
- "Delete stale state."
- "Renew real permissions."
- "Revoke real permissions."
- "Send notifications."
- "Auto-fix permissions."
- "Download packages."
- Any task that would require editing `docs/claude-code-implementation-packages.md`.

## Copyable Prompt Prefix

Use this prefix before the selected task prompt:

```text
You are implementing one bounded Xnix task from docs/claude-code-current-dispatch-picks.md.

Use the referenced source task document as the authority for detailed requirements.

Hard boundaries:
- Implement exactly one task.
- Do not modify docs/claude-code-implementation-packages.md.
- Keep all source, comments, tests, fixtures, CLI output, and docs in English.
- Prefer Go for durable Runtime product behavior.
- Use Ruby only for tests, smoke scripts, reports, and lightweight developer tooling.
- Keep KDE as shell and presentation only.
- Keep the Xnix Compatibility Runtime as the policy and product-logic owner.
- Do not enable production D-Bus ownership, Runtime writes, real Wine/Proton/VM launch, real Portal calls, network fetch, host package-manager calls, privileged containers, host networking, Docker socket mounts, broad host mounts, or host-root mutation.
- Do not expose raw executable paths, Runtime state-root paths, host paths, file contents, secrets, tokens, credentials, usernames, environment variables, raw backend commands, backend names, or backend details in KDE-facing output.

Before finishing:
- Update VERSION, CHANGELOG.md, and PRODUCT_OVERVIEW.md if code changes.
- Run the task-specific verification commands.
- Run ruby scripts/verify_layout.rb.
- Run git diff --check.
- Report exact commands run and any commands not run.
```

## Intake Reminder

When a Claude branch returns:

1. Run `ruby scripts/mainline_integration_review.rb --format json`.
2. Confirm `protected_claude_file_modified` is `false`.
3. Confirm `unclassified_file_count` is `0`.
4. Confirm the branch only touches the expected lane plus required shared plumbing and version metadata.
5. Run the task-specific tests listed in the source task document.
6. Exclude `.gocache/`, `.cache/`, `tmp/`, logs, build artifacts, and local config from staging.
7. Stage by lane, never with `git add .`.
